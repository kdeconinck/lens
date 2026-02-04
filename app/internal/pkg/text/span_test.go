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

// QA: Verify the implementation of the "text" package.
package text_test

import (
	"testing"

	"github.com/kdeconinck/lens/internal/pkg/assert"
	"github.com/kdeconinck/lens/internal/pkg/text"
)

// UT: Get the human-readable representation of a span
func Test_SpanString(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	// Arrange.
	span := newSpan(5, 7)

	// Act.
	got, want := span.String(), "5..7"

	// Assert.
	assert.Equalf(t, got, want, "\n\n"+
		"UT Name:  When formatting a 'Span' it's displayed as 'Start..End'.\n"+
		"\033[32mExpected: %s\033[0m\n"+
		"\033[31mActual:   %s\033[0m\n\n", want, got)
}

// Returns a new span with the given start and end positions.
func newSpan(start, end int) text.Span {
	return text.Span{
		Start: start,
		End:   end,
	}
}
