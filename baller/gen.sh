#! /usr/bin/env bash

v=$(head -n $2 "$1")

echo "$v" > "$1"

curl -s http://stats.nba.com/stats/${3} | go run ../res2spec/main.go | go run ../spec2cli/main.go >> "$1" 

go fmt "$1"
