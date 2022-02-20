package iteration

import (
	"fmt"
	"strings"
)

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func UpgradeRepeat(charcter string, count int) string {
	var b strings.Builder
	for i := count; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString(charcter)
	return b.String()
}
