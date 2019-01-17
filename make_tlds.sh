#!/bin/bash

curl -s --compressed https://data.iana.org/TLD/tlds-alpha-by-domain.txt |
    grep -vF '#' |
    awk '{print "\""tolower($0)"\""} END{print "\"onion\""}' |
    jq -s . > tlds.json