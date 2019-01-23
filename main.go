package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/assafmo/xioc/xioc"
)

type extractFunction func(text string) []string

var availableFunctions = map[string]extractFunction{
	"domain": xioc.ExtractDomains,
	"email":  xioc.ExtractEmails,
	"ip4":    xioc.ExtractIPv4s,
	"ip6":    xioc.ExtractIPv6s,
	"url":    xioc.ExtractURLs,
	"md5":    xioc.ExtractMD5s,
	"sha1":   xioc.ExtractSHA1s,
	"sha256": xioc.ExtractSHA256s,
}

const version = "1.1.10"

var versionFlag bool
var onlyFlag string

func init() {
	flag.BoolVar(&versionFlag, "v", false, "Print version and exit")
	flag.StringVar(&onlyFlag, "o", "", `Extract only specified types.
Types must be comma seperated. E.g: xioc -o "ip4,domain,url,md5"
Available types:
	- ip4
	- ip6
	- domain
	- url
	- email
	- md5
	- sha1
	- sha256`)

	flag.Parse()
}

func main() {
	// if -v flag, print version and exit
	if versionFlag {
		fmt.Printf("xioc v%s\n\n", version)
		fmt.Println("Extract domains, ips, urls, emails, md5, sha1 and sha256 from text.")
		fmt.Println("For more info visit https://github.com/assafmo/xioc")
		return
	}

	functions := availableFunctions
	if onlyFlag != "" {
		functions = map[string]extractFunction{}

		types := strings.Split(onlyFlag, ",")
		for _, t := range types {
			if f, ok := availableFunctions[t]; ok {
				functions[t] = f
			} else {
				fmt.Printf(`Unknown extraction type "%s"`+"\n", t)
				os.Exit(1)
			}
		}
	}

	fi, _ := os.Stdin.Stat()

	if (fi.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Cannot read data from stdin.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 64*1024)     // 64KiB initial size
	scanner.Buffer(buf, 5*1024*1024) // 5MiB max size

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
