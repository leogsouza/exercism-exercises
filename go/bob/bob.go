// Package Bob answers a question

package bob

import (
	"strings"
	"unicode"
)

// Answers for each situation
const (
	ASKQUESTION  = "Sure."
	YELL         = "Whoa, chill out!"
	YELLQUESTION = "Calm down, I know what I'm doing!"
	FINE         = "Fine. Be that way!"
	ANYTHING     = "Whatever."
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	switch {
	case remark == "":
		return FINE
	case anyChar(remark, unicode.IsUpper) && !anyChar(remark, unicode.IsLower) && !isQuestion(remark):
		return YELL
	case anyChar(remark, unicode.IsUpper) && !anyChar(remark, unicode.IsLower) && isQuestion(remark):
		return YELLQUESTION
	case isQuestion(remark):
		return ASKQUESTION
	default:
		return ANYTHING
	}
}

func isQuestion(remark string) bool {

	return remark[len(remark)-1] == '?'
}

func anyChar(remark string, check func(rune) bool) bool {
	for _, letter := range remark {
		if check(letter) {
			return true
		}
	}

	return false
}
