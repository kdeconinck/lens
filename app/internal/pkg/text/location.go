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

// Package text provides foundational types for managing and measuring text content.
//
// It serves as a coordinate system for UTF-8 encoded strings, defining how raw byte segments [Span]s map to
// human-readable dimensions like line and columns.
package text

import "fmt"

// Location represents a human-readable position (line & column) inside a string.
type Location struct {
	// Line is the 1-based line number.
	Line int

	// Column is the 1-based column number, representing the visual position within the line-based UTF-8 characters.
	Column int
}

// String returns the string representation of the location.
func (loc Location) String() string {
	return fmt.Sprintf("%d:%d", loc.Line, loc.Column)
}

// Updates the Location based on the provided rune.
//
// If the rune is a newline ('\n'), it increments the line count and resets the column to 1.
// For all other characters, it increments the column count.
func (loc *Location) advance(r rune) {
	if r == '\n' {
		loc.Line += 1
		loc.Column = 1
	} else {
		loc.Column++
	}
}
