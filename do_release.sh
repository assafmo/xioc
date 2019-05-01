#!/bin/bash

# build into ./release/

set -e
set -v

go get -v -u -t -d ./...

go test -race -cover ./...

rm -rf release
mkdir -p release

VERSION=$(git describe --tags $(git rev-list --tags --max-count=1))

# https://golang.org/doc/install/source#environment
GOOS=linux   GOARCH=amd64 go build -o "release/xioc-linux64-${VERSION}"
GOOS=windows GOARCH=amd64 go build -o "release/xioc-win64-${VERSION}.exe"
GOOS=darwin  GOARCH=amd64 go build -o "release/xioc-macos64-${VERSION}"

(
    # zip
    cd release
    find -type f | 
        parallel --bar 'zip "$(echo "{}" | sed "s/.exe//").zip" "{}" && rm -f "{}"'

    # deb
    mkdir -p ./deb/DEBIAN
    cat > ./deb/DEBIAN/control <<EOF 
Package: xioc
Architecture: amd64
Maintainer: Assaf Morami <assaf.morami@gmail.com>
Priority: optional
Version: $(echo "${VERSION}" | tr -d v)
Homepage: https://github.com/assafmo/xioc
Description: Extract indicators of compromise from text, including "escaped" ones. 
EOF

    mkdir -p ./deb/bin
    unzip -o -d ./deb/bin xioc-linux64-*
    mv -f ./deb/bin/xioc-linux64-* ./deb/bin/xioc

    dpkg-deb --build ./deb/ .
)

# publish ubuntu snap

rm -rf snap *.snap* *_source.tar.bz2
snapcraft
snapcraft push *.snap
REV=$(snapcraft list-revisions xioc | head -2 | tail -1 | awk '{print $1}')
snapcraft release xioc "$REV" stable
snapcraft clean
rm -rf snap *.snap* *_source.tar.bz2