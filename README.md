# xioc

Extract indicators of compromise from text, including "escaped" ones like `hxxp://banana.com`, `1.1.1[.]1` and `phish at malicious dot com`.

[![CircleCI](https://circleci.com/gh/assafmo/xioc.svg?style=shield&circle-token=53b168115c42a883184dd01267d549aed80c2f49)](https://circleci.com/gh/assafmo/xioc)
[![Coverage Status](https://coveralls.io/repos/github/assafmo/xioc/badge.svg?branch=master)](https://coveralls.io/github/assafmo/xioc?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/assafmo/xioc)](https://goreportcard.com/report/github.com/assafmo/xioc)
[![GoDoc](https://godoc.org/github.com/assafmo/xioc/xioc?status.svg)](https://godoc.org/github.com/assafmo/xioc/xioc)

## Installation

- Download a precompiled binary from https://github.com/assafmo/xioc/releases
- Or... Use `go get`:

  ```bash
  go get -u github.com/assafmo/xioc
  ```

- Or... Use snap install (Ubuntu):

  ```bash
  snap install xioc
  ```

## Features

- Extract IOCs (indicators of compromise) from an input text:
  - IPv4
  - IPv6
  - Domain
  - URL
  - Email
  - MD5
  - SHA1
  - SHA256
- Translate some kinds of "escaping"/"defanging" techniques:
  - `(dot)`, `[dot]`, `(.)`, `[.]`, `{.}` to `.`.
  - `(at)`, `[at]`, `(@)`, `[@]`, `{@}` to `@`.
  - `hxxp`, `hzzzp`, `hxxxp`, `hXXp`, `h__p`, `h**p` to `http`.
- Command line interface
- Go library

## Command line usage

```bash
$ xioc -h
Usage of xioc:
  -o string
        Extract only specified types.
        Types must be comma seperated. E.g: xioc -o "ip4,domain,url,md5"
        Available types:
                - ip4
                - ip6
                - domain
                - url
                - email
                - md5
                - sha1
                - sha256
  -v    Print version and exit
```

```bash
$ REPORT="https://unit42.paloaltonetworks.com/digital-quartermaster-scenario-demonstrated-in-attacks-against-the-mongolian-government/"
$ lynx -dump "$REPORT" | xioc
sha256  5beb50d95c1e720143ca0004f5172cb8881d75f6c9f434ceaff59f34fa1fe378
domain  energy.gov.mn
email   altangadas@energy.gov.mn
sha256  10090692ff40758a08bd66f806e0f2c831b4b9742bbf3d19c250e778de638f57
# ...
```

```bash
$ REPORT="https://unit42.paloaltonetworks.com/digital-quartermaster-scenario-demonstrated-in-attacks-against-the-mongolian-government/"
$ lynx -dump "$REPORT" | xioc -o email,sha256
sha256  5beb50d95c1e720143ca0004f5172cb8881d75f6c9f434ceaff59f34fa1fe378
email   altangadas@energy.gov.mn
sha256  10090692ff40758a08bd66f806e0f2c831b4b9742bbf3d19c250e778de638f57
email   ganbat_g@bpo.gov.mn
# ...
```

## Library usage

Full API:  
[![GoDoc](https://godoc.org/github.com/assafmo/xioc/xioc?status.svg)](https://godoc.org/github.com/assafmo/xioc/xioc)

```golang
package main

import (
	"fmt"

	"github.com/assafmo/xioc/xioc"
)

func main() {
	input := `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	banana.com
	hxxp://i.robot.com/robots.txt
	1.2.3.4
	1.1.1[.]1
	info at gmail dot com
	hxxps://m.twitter[dot]com/`

	fmt.Println(xioc.ExtractDomains(input)) // => [i.robot.com m.twitter.com gmail.com banana.com]
	fmt.Println(xioc.ExtractSHA256s(input)) // => [e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855]
	fmt.Println(xioc.ExtractMD5s(input))    // => []
	fmt.Println(xioc.ExtractIPv4s(input))   // => [1.2.3.4 1.1.1.1]
	fmt.Println(xioc.ExtractURLs(input))    // => [http://i.robot.com/robots.txt https://m.twitter.com/]
	fmt.Println(xioc.ExtractEmails(input))  // => [info@gmail.com]
}
```

## Sources

- Test email address: http://codefool.tumblr.com/post/15288874550/list-of-valid-and-invalid-email-addresses
- Domains can start with a number: https://serverfault.com/a/638270
- IPv6 Examples: http://www.gestioip.net/docu/ipv6_address_examples.html
- Fang and defang IOCs: https://github.com/ioc-fang/ioc_fanger
- Indicator of Compromise (De)Fanging Project: https://ioc-fang.hightower.space/
- InQuest/python-iocextract test data: https://github.com/InQuest/python-iocextract/tree/master/test_data
- Email address can be treated as case-insensitive: https://stackoverflow.com/a/9808332
