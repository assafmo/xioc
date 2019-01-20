#!/bin/bash

echo -e "package xioc\n" > xioc/tlds.go
echo "// KnownTLDs is a set of all known TLDs" >> xioc/tlds.go
echo "var KnownTLDs = map[string]bool{" >> xioc/tlds.go

curl -s --compressed https://data.iana.org/TLD/tlds-alpha-by-domain.txt |
    grep -vF '#' |
    awk '{print "\""tolower($0)"\": true,"}' >> xioc/tlds.go

echo '"onion": true,' >> xioc/tlds.go

echo "}" >> xioc/tlds.go
