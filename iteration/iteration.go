package iteration

import "strings"

const repeatCount = 5

func Repeat(c string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(c)
	}
	return repeated.String()
}

func Upper(s string) string {
	return strings.ToUpper(s)
}
