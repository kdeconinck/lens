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

import "unicode/utf8"

// Input encapsulates a string.
type Input struct {
	// Content represents the full string.
	Content string
}

// LineCol translates a 0-based byte offset into 1-based line and column numbers.
//
// LineCol is UTF-8 aware; it treats multibyte characters as a single visual column unit. It iterates through the
// content to find the position, making it an O(n) operation relative to the offset.
func (input Input) LineCol(offset int) (loc Location) {
	loc.Line, loc.Column = 1, 1

	if offset > len(input.Content) {
		offset = len(input.Content) - 1
	}

	// Backtrack if the offset is in the middle of a multi-byte character.
	for offset > 0 && !utf8.RuneStart(input.Content[offset]) {
		offset--
	}

	currentByte := 0

	for _, r := range input.Content {
		if currentByte >= offset {
			break
		}

		loc.advance(r)

		currentByte += len(string(r))
	}

	return loc
}
