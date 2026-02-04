// =====================================================================================================================
// = LICENSE:       Copyright (c) 2026 Kevin De Coninck
// =
// =                Permission is hereby granted, free of charge, to any person
// =                obtaining a copy of this software and associated documentation
// =                files (the "Software"), to deal in the Software without
// =                restriction, including without limitation the rights to use,
// =                copy, modify, merge, publish, distribute, sublicense, and/or sell
// =                copies of the Software, and to permit persons to whom the
// =                Software is furnished to do so, subject to the following
// =                conditions:
// =
// =                The above copyright notice and this permission notice shall be
// =                included in all copies or substantial portions of the Software.
// =
// =                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// =                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// =                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// =                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// =                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// =                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// =                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// =                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package token defines the lexical atoms of the lux language.
package token

import "strconv"

// Type represents the category of a lexical token.
type Type int

// The different types of a lexical token.
const (
	Error Type = iota
	EOF
	Ident
	Number
	String
	Bool
	Dot
	LBrace
	RBrace
	LBracket
	RBracket
	Colon
	Equals
	Comma
)

// Maps a [Type] to its human-readable name.
var tokenMap = map[Type]string{
	Error:    "Error",
	EOF:      "EOF",
	Ident:    "Identifier",
	Number:   "Number",
	String:   "String",
	Bool:     "Boolean",
	Dot:      ".",
	LBrace:   "{",
	RBrace:   "}",
	LBracket: "[",
	RBracket: "]",
	Colon:    ":",
	Equals:   "=",
	Comma:    ",",
}

// String returns the string representation of the token type.
func (t Type) String() string {
	value, ok := tokenMap[t]

	if ok {
		return value
	}

	return "Unknown(" + strconv.Itoa(int(t)) + ")"
}
