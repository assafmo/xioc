package xioc

import (
	"log"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/idna"
)

func init() {
	for tld := range KnownTLDs {
		if strings.HasPrefix(tld, "xn--") {
			u, err := idna.ToUnicode(tld)
			if err != nil {
				log.Fatalf(`cannot encode IDNA TLD "%s" to unicode: %v`, tld, err)
			}
			KnownTLDs[u] = true
		}
	}
}

var dot = `(\.| dot |\(dot\)|\[dot\]|\(\.\)|\[\.\])`
var dotRegex = regexp.MustCompile(`(?i)` + dot)

func replaceDot(s string) string {
	return dotRegex.ReplaceAllString(s, ".")
}

var at = `(@| at |\(at\)|\[at\]|\(@\)|\[@\])`
var atRegex = regexp.MustCompile(`(?i)` + at)

func replaceAt(s string) string {
	return atRegex.ReplaceAllString(s, "@")
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

// ExtractIPv4s extracts IPv4 addresses from an input string
func ExtractIPv4s(text string) []string {
	ips := ip4Regex.FindAllString(text, -1)
	return filterOnlyValidIPs(ips)
}

var ip6Regex = regexp.MustCompile(`(?i)[a-f0-9:]+`)

// ExtractIPv6s extracts IPv6 addresses from an input string
func ExtractIPv6s(text string) []string {
	ips := ip6Regex.FindAllString(text, -1)
	return filterOnlyValidIPs(ips)
}

func hasKnownTLD(input string) bool {
	domainParts := strings.Split(input, ".")
	return KnownTLDs[domainParts[len(domainParts)-1]]
}

var emailRegex = regexp.MustCompile(`(?i)\b\S+?` + at + `\S+?` + dot + `\S+\b`)

// ExtractEmails extracts email addresses from an input string
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

		domain := strings.Split(email, "@")[1]
		isIP := len(filterOnlyValidIPs([]string{domain})) == 1

		if !hasKnownTLD(email) && !isIP {
			continue
		}

		if !resultSet[email] {
			resultSet[email] = true
			result = append(result, email)
		}
	}

	return result
}

var urlRegex = regexp.MustCompile(`(?i)(h..ps?|ftp)\[?:\]?//\s?\S+`)

// ExtractURLs extracts ftp and http addresses from an input string
func ExtractURLs(text string) []string {
	urls := urlRegex.FindAllString(text, -1)

	resultSet := map[string]bool{}
	result := []string{}
	for _, url := range urls {
		url = replaceDot(url)

		url = strings.Replace(url, "hxxp", "http", -1)
		url = strings.Replace(url, "hXXp", "http", -1)
		url = strings.Replace(url, "h__p", "http", -1)
		url = strings.Replace(url, "h**p", "http", -1)
		url = strings.Replace(url, "http[:]//", "http://", -1)
		url = strings.Replace(url, "https[:]//", "https://", -1)
		url = strings.Replace(url, ":// ", "://", -1)
		url = strings.Replace(url, "[com]", "com", -1)

		if !strings.Contains(url, ".") && // check for domain without mutual calls
			len(ExtractIPv4s(url)) == 0 &&
			len(ExtractIPv6s(url)) == 0 {
			continue
		}

		if !resultSet[url] {
			resultSet[url] = true
			result = append(result, url)
		}
	}

	return result

}

var domainRegex = regexp.MustCompile(`(?i)([\p{L}\p{N}][\p{L}\p{N}\-]*` + dot + `)+[a-z]{2,}`)

// ExtractDomains extracts domain names from an input string
func ExtractDomains(text string) []string {
	domains := []string{}

	urls := ExtractURLs(text)
	for _, u := range urls {
		parsedURL, err := url.Parse(u)
		if err != nil {
			continue
		}
		domains = append(domains, parsedURL.Hostname())
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

		if strings.ContainsAny(domain, `!#$%^&*()+=,@:/'\"[]`+"`") ||
			strings.Contains(domain, "..") ||
			strings.Contains(domain, ".-") {
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

// ExtractMD5s extracts md5 hex strings from an input string
func ExtractMD5s(text string) []string {
	return md5Regex.FindAllString(text, -1)
}

// ExtractSHA1s extracts sha1 hex strings from an input string
func ExtractSHA1s(text string) []string {
	return sha1Regex.FindAllString(text, -1)
}

// ExtractSHA256s extracts sha256 hex strings from an input string
func ExtractSHA256s(text string) []string {
	return sha256Regex.FindAllString(text, -1)
}
