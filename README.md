# xioc

Extract domains, ips, urls, emails, md5, sha1 and sha256 from text.

[![CircleCI](https://circleci.com/gh/assafmo/xioc.svg?style=shield&circle-token=53b168115c42a883184dd01267d549aed80c2f49)](https://circleci.com/gh/assafmo/xioc)
[![Coverage Status](https://coveralls.io/repos/github/assafmo/xioc/badge.svg?branch=master)](https://coveralls.io/github/assafmo/xioc?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/assafmo/xioc)](https://goreportcard.com/report/github.com/assafmo/xioc)
[![GoDoc](https://godoc.org/github.com/assafmo/xioc?status.svg)](https://godoc.org/github.com/assafmo/xioc)

## Installation

```bash
go get -u github.com/assafmo/xioc
```

## Command line

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
$ lynx -dump https://unit42.paloaltonetworks.com/digital-quartermaster-scenario-demonstrated-in-attacks-against-the-mongolian-government/ | xioc
sha256  5beb50d95c1e720143ca0004f5172cb8881d75f6c9f434ceaff59f34fa1fe378
domain  energy.gov.mn
email   altangadas@energy.gov.mn
sha256  10090692ff40758a08bd66f806e0f2c831b4b9742bbf3d19c250e778de638f57
domain  bpo.gov.mn
email   ganbat_g@bpo.gov.mn
sha256  44dbf05bc81d17542a656525772e0f0973b603704f213278036d8ffc999bb79a
sha256  91ffe6fab7b33ff47b184b59356408951176c670cad3afcde79aa8464374acd3
sha256  6f3d4fb64de9ae61776fd19a8eba3d1d828e7e26bb89ace00c7843a57c5f6e8a
domain  masm.gov.mn
# ...
```

```bash
$ lynx -dump https://unit42.paloaltonetworks.com/digital-quartermaster-scenario-demonstrated-in-attacks-against-the-mongolian-government/ | xioc -o email,sha256
sha256  5beb50d95c1e720143ca0004f5172cb8881d75f6c9f434ceaff59f34fa1fe378
email   altangadas@energy.gov.mn
sha256  10090692ff40758a08bd66f806e0f2c831b4b9742bbf3d19c250e778de638f57
email   ganbat_g@bpo.gov.mn
sha256  44dbf05bc81d17542a656525772e0f0973b603704f213278036d8ffc999bb79a
sha256  91ffe6fab7b33ff47b184b59356408951176c670cad3afcde79aa8464374acd3
sha256  6f3d4fb64de9ae61776fd19a8eba3d1d828e7e26bb89ace00c7843a57c5f6e8a
email   bilguun@masm.gov.mn
sha256  e88ea5eb642eaf832f8399d0337ba9eb1563862ddee68c26a74409a7384b9bb9
email   davaa_ayush@yahoo.com
# ...
```

## Library

[![GoDoc](https://godoc.org/github.com/assafmo/xioc/xioc?status.svg)](https://godoc.org/github.com/assafmo/xioc/xioc)

## Sources

- Test email address: http://codefool.tumblr.com/post/15288874550/list-of-valid-and-invalid-email-addresses
- Domains can start with a number: https://serverfault.com/a/638270
- IPv6 Examples: http://www.gestioip.net/docu/ipv6_address_examples.html
