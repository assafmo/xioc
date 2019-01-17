#!/bin/bash

echo -e "package main\n" > tlds.go
echo "var knownTLDs = map[string]bool{" >> tlds.go

curl -s --compressed https://data.iana.org/TLD/tlds-alpha-by-domain.txt |
    grep -vF '#' |
    awk '{print "\""tolower($0)"\": true,"}' >> tlds.go

echo '"onion": true,' >> tlds.go

echo "}" >> tlds.go
