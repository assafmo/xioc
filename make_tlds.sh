#!/bin/bash

echo -e "package main\n" > tlds.go
echo "var knownTLDs = []string{" >> tlds.go

curl -s --compressed https://data.iana.org/TLD/tlds-alpha-by-domain.txt |
    grep -vF '#' |
    awk '{print "\""tolower($0)"\","} END{print "\"onion\","}' >> tlds.go

echo "}" >> tlds.go
