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

// UT: Verify that byte offsets are correctly translated into line and column coordinates.
func TestInput_LineCol(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When the offset is at the beginning, the representation matches '1:1'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		input := newInput("Hello")

		// Act.
		got, want := input.LineCol(0), newLocation(1, 1)

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When the offset is at the beginning, the representation matches '1:1'.\n"+
			"\033[32mExpected: %s\033[0m\n"+
			"\033[31mActual:   %s\033[0m\n\n", want, got)
	})

	t.Run("When the offset is after a newline, the representation matches '2:1'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		input := newInput("Hello\nWorld")

		// Act.
		got, want := input.LineCol(6), newLocation(2, 1)

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When the offset is after a newline, the representation matches '2:1'.\n"+
			"\033[32mExpected: %s\033[0m\n"+
			"\033[31mActual:   %s\033[0m\n\n", want, got)
	})

	t.Run("When the offset is after 2 newline characters, the representation matches '3:1'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		input := newInput("Hello\n\nWorld")

		// Act.
		got, want := input.LineCol(7), newLocation(3, 1)

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When the offset is after 2 newline characters, the representation matches '3:1'.\n"+
			"\033[32mExpected: %s\033[0m\n"+
			"\033[31mActual:   %s\033[0m\n\n", want, got)
	})

	t.Run("When the offset in the middle of a line, the representation matches the correct line and column.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		input := newInput("Hello\nWorld")

		// Act.
		got, want := input.LineCol(8), newLocation(2, 3)

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When the offset in the middle of a line, the representation matches the correct line and column.\n"+
			"\033[32mExpected: %s\033[0m\n"+
			"\033[31mActual:   %s\033[0m\n\n", want, got)
	})

	t.Run("When the offset is out of bounds, the representation matches the maximum line and column.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		input := newInput("Hello")

		// Act.
		got, want := input.LineCol(120), newLocation(1, 5)

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  When the offset is out of bounds, the representation matches the maximum line and column.\n"+
			"\033[32mExpected: %s\033[0m\n"+
			"\033[31mActual:   %s\033[0m\n\n", want, got)
	})

	t.Run("An offset can be used to handle a multi-byte character.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		// Note: The "rocket" emoji uses 4 bytes.
		input := newInput("A🚀C")

		// Act.
		got, want := input.LineCol(3), newLocation(1, 2)

		// Assert.
		assert.Equalf(t, got, want, "\n\n"+
			"UT Name:  An offset can be used to handle a multi-byte character.\n"+
			"\033[32mExpected: %s\033[0m\n"+
			"\033[31mActual:   %s\033[0m\n\n", want, got)
	})
}

// Returns a new input with the given content.
func newInput(content string) text.Input {
	return text.Input{
		Content: content,
	}
}
