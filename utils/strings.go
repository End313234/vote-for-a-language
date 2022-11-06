package utils

import "strings"

// In that case, it is not correct to use
// https://pkg.go.dev/golang.org/x/text/cases#Title
// since it makes the first letter uppercase but
// lowercase the other ones.
//
// This behavior might negatively impact while
// searching for languages that might contain an
// uppercase letter in the middle of the name.
//
// E.g.: NodeJs
func UpperCaseFirstLetter(word string) string {
	return strings.ToUpper(string(word[0])) + word[1:]
}
