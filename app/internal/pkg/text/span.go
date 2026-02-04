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

// Span represents a discrete region within a body of text.
//
// It uses a half-open interval [Start, End) based on 0-indexed byte offsets.
// This design allows for high performance slicing of strings and easy calculating of length (End - Start).
type Span struct {
	// Start is the 0-based byte offset of the first byte in the span.
	Start int

	// End if the 0-based byte offset of the first byte following the span.
	End int
}

// String returns the string representation of the span.
func (span Span) String() string {
	return fmt.Sprintf("%d..%d", span.Start, span.End)
}
