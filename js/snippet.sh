#! /usr/bin/env bash

IFS=$'\n'       # make newlines the only separator
set -f
echo "// +build webface" > ../javascript.go
echo "" >> ../javascript.go
echo "package samcatweb" >> ../javascript.go
echo "func defaultJS() string {" >> ../javascript.go
printf '    r := '>> ../javascript.go
for line in $(cat script.js); do
    echo "    \"$line\n\" +" >> ../javascript.go
done
echo "    \"\n\"" >> ../javascript.go
echo "    return r" >> ../javascript.go
echo "}" >> ../javascript.go
gofmt -w ../javascript.go
