package baller

import (
	"strconv"
)

func encodeString(v string) string {
	return v
}

func encodeInt(v int) string {
	return strconv.Itoa(v)
}

func encodeBool(v bool) string {
	if v {
		return "1"
	}

	return "0"
}
