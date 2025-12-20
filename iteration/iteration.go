package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func Replace(stringToReplace, replacedString string) string {
	return strings.ReplaceAll(stringToReplace, stringToReplace, replacedString)
}

func Upper(stringInput string) string {
	return strings.ToUpper(stringInput)
}
