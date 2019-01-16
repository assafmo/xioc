package main // import "github.com/assafmo/xioc"

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func ExtractIPv4s(text string) []string {
	return []string{}
}

func ExtractIPv6s(text string) []string {
	return []string{}
}

func ExtractDomains(text string) []string {
	return []string{}
}

func ExtractURLs(text string) []string {
	return []string{}
}

func ExtractEmails(text string) []string {
	return []string{}
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
