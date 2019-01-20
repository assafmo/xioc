package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/assafmo/xioc/xioc"
)

type extractFunction func(text string) []string

var functions = map[string]extractFunction{
	"domain": xioc.ExtractDomains,
	"email":  xioc.ExtractEmails,
	"ip4":    xioc.ExtractIPv4s,
	"ip6":    xioc.ExtractIPv6s,
	"url":    xioc.ExtractURLs,
	"md5":    xioc.ExtractMD5s,
	"sha1":   xioc.ExtractSHA1s,
	"sha256": xioc.ExtractSHA256s,
}

const version = "1.0.0"

var isPrintVersion bool

func init() {
	flag.BoolVar(&isPrintVersion, "v", false, "Print version and exit")
	flag.Parse()
}

func main() {
	// if -v flag, print version and exit
	if isPrintVersion {
		fmt.Printf("xioc v%s\n\n", version)
		fmt.Println("Extract domains, ips, urls, emails, md5, sha1 and sha256 from text.")
		fmt.Println("For more info visit https://github.com/assafmo/xioc")
		return
	}

	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Cannot read data from stdin.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		for iocType, f := range functions {
			results := f(text)
			for _, ioc := range results {
				fmt.Printf("%s\t%s\n", iocType, ioc)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
