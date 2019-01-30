# Changelog

## v1.1.11 (Jan 30, 2019)

- fix crash if stdin line bigger than 5MiB (issue #7)

## v1.1.10 (Jan 23, 2019)

- de-defang `{.}` => `.`
- de-defang `{@}` => `@`
- allow whitespace before and after special defangs like `{.}`, `[@]`, `(at)`, etc.

## v1.1.7 (Jan 21, 2019)

- increase max read buffer from 64KiB to 5MiB

## v1.1.6 (Jan 21, 2019)

- domain: support internationalized domain names (IDNs)

## v1.1.5 (Jan 21, 2019)

- url: catch "hzzzp" and "hxxxp"
- url: http is case insensitive

## v1.1.2 (Jan 21, 2019)

- url: must contain a domain or an ip

## v1.1.1 (Jan 21, 2019)

- url: catch "://(space)"

## v1.1.0 (Jan 20, 2019)

- -o flag: extract only selected types

## v1.0.0 (Jan 20, 2019)

- extract ip4, ip6, domain, url, email, md5, sha1, sha256
