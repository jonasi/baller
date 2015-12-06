#! /usr/bin/env bash

CACHE_DIR=.cache
set -e

gen() {
    local cf="${CACHE_DIR}/${2}?${3}"

    if [[ ! -f "$cf" ]]; then
        curl -s "http://stats.nba.com/stats/${2}?${3}" > "$cf"
    fi

    local spec=$(res2spec -name="$1" -path="$2" -response="$cf")

    echo "$spec" | gen_api "$1"
    echo "$spec" | gen_cli "$1"
}

gen_api() {
    local f="gen_${1}.go"
    cat <<EOF > "$f"
package baller

import (
    "net/url"
)
EOF

    spec2go >> "$f"
    go fmt "$f"
}

gen_cli() {
    local f="baller/gen_${1}.go"
    cat <<EOF > "$f"
package main

import (
	"flag"
	"github.com/jonasi/baller"
	"os"
)
EOF

    spec2cli >> "$f"
    go fmt "$f"
}

mkdir -p "$CACHE_DIR"
go install github.com/jonasi/baller/{spec2go,spec2cli,res2spec}

rm -rf gen_*.go
rm -rf baller/gen_*.go

gen boxscore boxscore "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_advanced boxscoreadvanced "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_advanced_v2 boxscoreadvancedv2 "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_four_factors boxscorefourfactors "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_four_factors_v2 boxscorefourfactorsv2 "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_misc boxscoremisc "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_misc_v2 boxscoremiscv2 "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_player_track_v2 boxscoreplayertrackv2 "GameID=0021500277"
gen boxscore_scoring boxscorescoring "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_scoring_v2 boxscorescoringv2 "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_summary_v2 boxscoresummaryv2 "GameID=0021500277"
gen boxscore_traditional_v2 boxscoretranditionalv2 "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_usage boxscoreusage "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen boxscore_usage_v2 boxscoreusagev2 "GameID=0021500277&StartPeriod=0&EndPeriod=4&StartRange=0&EndRange=0&RangeType=1"
gen common_team_years commonteamyears "LeagueID=00"
gen common_all_players commonallplayers "LeagueID=00&Season=2015-16&IsOnlyCurrentSeason=1"
