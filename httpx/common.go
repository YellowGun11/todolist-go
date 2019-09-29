package httpx

import (
	"fmt"
	"strings"
)

func MergeURL(first, second string) string {
	firstEnd := len(first) - 1
	if strings.HasSuffix(first, "/") {
		firstEnd -= 1
	}
	secondBegin := 0
	if strings.HasPrefix(second, "/") {
		secondBegin += 1
	}
	return fmt.Sprintf("%s/%s", first[:firstEnd], second[secondBegin:])
}
