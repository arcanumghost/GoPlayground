package strutil

import "strings"

// Reverse all characters
func Reverse(s string) string {
	runes := []rune(s)
	returnVal := make([]rune, len(runes))
	for i := 0; i < len(runes); i = i + 1 {
		returnVal[i] = runes[len(runes)-1-i]
	}
	return string(returnVal)
}

// Reverse by word
func ReverseWords(s string) string {
	mysplit := strings.Split(s, " ")
	returnVal := make([]string, len(mysplit))
	for i := 0; i < len(mysplit); i = i + 1 {
		returnVal[i] = mysplit[len(mysplit)-1-i]
	}
	return strings.Join(returnVal, " ")
}
