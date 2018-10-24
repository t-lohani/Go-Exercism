// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

const (
	QUESTION      = "Sure."
	YELL          = "Whoa, chill out!"
	QUESTION_YELL = "Calm down, I know what I'm doing!"
	QUIET         = "Fine. Be that way!"
	ORELSE        = "Whatever."
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return QUIET
	}

	isQuestion := strings.HasSuffix(remark, "?")

	upperCount := 0
	lowerCount := 0

	for _, char := range remark {
		if unicode.IsLower(char) {
			lowerCount++
		} else if unicode.IsUpper(char) {
			upperCount++
		}
	}

	if isQuestion {
		if lowerCount == 0 && upperCount > 0 {
			return QUESTION_YELL
		} else {
			return QUESTION
		}
	} else if lowerCount == 0 && upperCount > 0 {
		return YELL
	}

	return ORELSE
}
