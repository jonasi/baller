#! /usr/bin/env bash

gen() {
    local spec=$(curl -s "http://stats.nba.com/stats/${2}" | go run res2spec/main.go)

    echo "$spec" | gen_api "$1"
    echo "$spec" | gen_cli "$1"
}

gen_api() {
    cat <<EOF > "$1".go
package baller

import (
    "net/url"
)
EOF

    go run spec2go/main.go >> "$1".go
    go fmt "$1".go
}

gen_cli() {
    cat <<EOF > "baller/$1.go"
package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)
EOF

    go run spec2cli/main.go >> "baller/$1".go
    go fmt "baller/$1".go
}

gen boxscore "boxscore?GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
