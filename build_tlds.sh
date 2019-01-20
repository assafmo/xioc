#!/bin/bash

echo -e "package xioc\n" > tlds.go
echo "// KnownTLDs is a set of all known TLDs" >> tlds.go
echo "var KnownTLDs = map[string]bool{" >> tlds.go

curl -s --compressed https://data.iana.org/TLD/tlds-alpha-by-domain.txt |
    grep -vF '#' |
    awk '{print "\""tolower($0)"\": true,"}' >> tlds.go

echo '"onion": true,' >> tlds.go

echo "}" >> tlds.go
