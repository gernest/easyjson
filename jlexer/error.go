package jlexer

import "strconv"

// LexerError implements the error interface and represents all possible errors that can be
// generated during parsing the JSON data.
type LexerError struct {
	Reason string
	Offset int
	Data   string
}

func (l *LexerError) Error() string {
	return "parse error: " + l.Reason +
		"near offset " + strconv.FormatInt(int64(l.Offset), 10) + "of '" + l.Data + "'"
}
