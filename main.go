package main // import "github.com/assafmo/xioc"

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func init() {
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		for _, f := range []ExtractFunction{
			ExtractDomains,
			ExtractEmails,
			ExtractIPv4s,
			ExtractIPv6s,
			ExtractURLs,
			ExtractMD5s,
			ExtractSHA1s,
			ExtractSHA256s,
			ExtractSHA512s,
		} {
			result := f(text)
			if len(result) > 0 {
				fmt.Println(result)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

type ExtractFunction func(text string) []string

var dot = `(\.| dot |\(dot\)|\[dot\]|\(\.\)|\[\.\])`
var dotRegex = regexp.MustCompile(`(?i)` + dot)

func replaceDot(s string) string {
	return dotRegex.ReplaceAllString(s, ".")
}

var at = `(@| at |\(at\)|\[at\]|\(@\)|\[@\])`
var atRegex = regexp.MustCompile(`(?i)` + at)

func replaceAt(s string) string {
	return atRegex.ReplaceAllString(s, ".")
}

func filterOnlyValidIPs(ips []string) []string {
	resultSet := map[string]bool{}
	result := []string{}
	for _, ip := range ips {
		ip = strings.ToLower(replaceDot(ip))

		if resultSet[ip] {
			continue
		}

		if net.ParseIP(ip) == nil {
			continue
		}

		if !resultSet[ip] {
			resultSet[ip] = true
			result = append(result, ip)
		}
	}
	return result
}

var ip4Regex = regexp.MustCompile(`(?i)([0-9]|` + dot + `)+`)

func ExtractIPv4s(text string) []string {
	ips := ip4Regex.FindAllString(text, -1)
	return filterOnlyValidIPs(ips)
}

var ip6Regex = regexp.MustCompile(`(?i)[a-f0-9:]+`)

func ExtractIPv6s(text string) []string {
	ips := ip6Regex.FindAllString(text, -1)
	return filterOnlyValidIPs(ips)
}

func hasKnownTLD(input string) bool {
	domainParts := strings.Split(input, ".")
	return knownTLDs[domainParts[len(domainParts)-1]]
}

var emailRegex = regexp.MustCompile(`(?i)\b\S+?` + at + `\S+?` + dot + `\S+\b`)

func ExtractEmails(text string) []string {
	emails := emailRegex.FindAllString(text, -1)

	resultSet := map[string]bool{}
	result := []string{}
	for _, email := range emails {
		email = strings.ToLower(email)
		email = replaceAt(email)
		email = replaceDot(email)

		e, err := mail.ParseAddress(email)
		if err != nil {
			continue
		}

		email = e.Address

		if resultSet[email] {
			continue
		}

		if !hasKnownTLD(email) {
			continue
		}

		if !resultSet[email] {
			resultSet[email] = true
			result = append(result, email)
		}
	}

	return result
}

var urlRegex = regexp.MustCompile(`(?i)(h..ps?|ftp)://\S+`)

func ExtractURLs(text string) []string {
	urls := urlRegex.FindAllString(text, -1)

	resultSet := map[string]bool{}
	result := []string{}
	for _, u := range urls {
		u = replaceDot(u)

		u = strings.Replace(u, "hxxp", "http", -1)
		u = strings.Replace(u, "hXXp", "http", -1)
		u = strings.Replace(u, "h__p", "http", -1)
		u = strings.Replace(u, "h**p", "http", -1)
		u = strings.Replace(u, "[com]", "com", -1)

		parsedUrl, err := url.Parse(u)
		if err != nil {
			continue
		}

		u = parsedUrl.String()

		if !resultSet[u] {
			resultSet[u] = true
			result = append(result, u)
		}
	}

	return result

}

var domainRegex = regexp.MustCompile(`(?i)([\p{L}\p{N}][\p{L}\p{N}\-]*` + dot + `)+[a-z]{2,}`)

func ExtractDomains(text string) []string {
	domains := []string{}

	urls := ExtractURLs(text)
	for _, u := range urls {
		parsedUrl, err := url.Parse(u)
		if err != nil {
			continue
		}

		domains = append(domains, parsedUrl.Hostname())
	}

	emails := ExtractEmails(text)
	for _, email := range emails {
		domain := strings.Split(email, "@")[1]
		domains = append(domains, domain)
	}

	domains = append(domains, domainRegex.FindAllString(text, -1)...)

	resultSet := map[string]bool{}
	result := []string{}
	for _, domain := range domains {
		domain = strings.ToLower(domain)
		domain = replaceDot(domain)

		if resultSet[domain] {
			continue
		}

		if strings.ContainsAny(domain, `!#$%^&*()+=,@:/'\"[]`+"`") {
			continue
		}

		if strings.Contains(domain, "..") {
			continue
		}

		if strings.Contains(domain, ".-") {
			continue
		}

		if net.ParseIP(domain) != nil {
			continue
		}

		if !hasKnownTLD(domain) {
			continue
		}

		if !resultSet[domain] {
			resultSet[domain] = true
			result = append(result, domain)
		}
	}

	return result
}

var md5Regex = regexp.MustCompile(`(?i)\b[a-f0-9]{32}\b`)
var sha1Regex = regexp.MustCompile(`(?i)\b[a-f0-9]{40}\b`)
var sha256Regex = regexp.MustCompile(`(?i)\b[a-f0-9]{64}\b`)
var sha512Regex = regexp.MustCompile(`(?i)\b[a-f0-9]{128}\b`)

func ExtractMD5s(text string) []string {
	return md5Regex.FindAllString(text, -1)
}

func ExtractSHA1s(text string) []string {
	return sha1Regex.FindAllString(text, -1)
}

func ExtractSHA256s(text string) []string {
	return sha256Regex.FindAllString(text, -1)
}

func ExtractSHA512s(text string) []string {
	return sha512Regex.FindAllString(text, -1)
}
