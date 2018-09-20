#! /usr/bin/env bash

IFS=$'\n'       # make newlines the only separator
set -f
echo "// +build webface" > ../style.go
echo "" >> ../style.go
echo "package samcatwebstyle" >> ../style.go
echo "func defaultCSS() string {" >> ../style.go
printf '    r := '>> ../style.go
for line in $(cat styles.css); do
    echo "    \"$line\n\" +" >> ../style.go
done
echo "    \"\n\"" >> ../style.go
echo "    return r" >> ../style.go
echo "}" >> ../style.go
gofmt -w ../style.go
