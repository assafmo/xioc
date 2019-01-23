#!/bin/bash

# build into ./release/
rm -rf release
mkdir -p release

VERSION=$(git describe --tags $(git rev-list --tags --max-count=1))

# https://golang.org/doc/install/source#environment
GOOS=linux   GOARCH=amd64 go build -o "release/xioc-linux64-${VERSION}"
GOOS=windows GOARCH=amd64 go build -o "release/xioc-win64-${VERSION}.exe"
GOOS=darwin  GOARCH=amd64 go build -o "release/xioc-macos64-${VERSION}"

(
    cd release
    find -type f | 
    parallel --bar 'zip "$(echo "{}" | sed "s/.exe//").zip" "{}" && rm -f "{}"'
)

# publish ubuntu snap

rm -rf snap *.snap*
snapcraft
snapcraft push *.snap
REV=$(snapcraft list-revisions xioc | head -2 | tail -1 | awk '{print $1}')
snapcraft release xioc "$REV" stable
snapcraft clean
rm -rf snap *.snap*