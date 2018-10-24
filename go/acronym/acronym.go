// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"bytes"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	f := func(c rune) bool {
		if c == ' ' || c == '-' {
			return true
		}
		return false
	}

	var buffer bytes.Buffer
	arr := strings.FieldsFunc(s, f)

	for _, elem := range arr {
		buffer.WriteByte(elem[0])
	}

	return strings.ToUpper(buffer.String())
}
