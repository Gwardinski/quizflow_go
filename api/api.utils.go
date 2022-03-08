package main

import (
	"strconv"
	"strings"
)

func byteStringToIntSlice(byteString string) (ids []int) {
	trimmed := strings.Trim(byteString, "{}")
	if trimmed != "" {
		strings := strings.Split(trimmed, ",")
		for _, s := range strings {
			i, _ := strconv.Atoi(s)
			ids = append(ids, i)
		}
	}
	return ids
}
