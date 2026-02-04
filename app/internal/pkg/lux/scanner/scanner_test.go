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

// QA: Verify the implementation of the "scanner" package.
package scanner_test

import (
	"testing"

	"github.com/kdeconinck/lens/internal/pkg/assert"
	"github.com/kdeconinck/lens/internal/pkg/lux/scanner"
	"github.com/kdeconinck/lens/internal/pkg/lux/token"
	"github.com/kdeconinck/lens/internal/pkg/text"
)

// UT: Convert a string into a set of lexical tokens.
func TestScanner_NextToken(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When scanning NO data, the 'EOF' token is returned.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.EOF, 0, 0),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  Whitespace characters are ignored.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning whitespace, the whitespace is ignored.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner(" \t ")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.EOF, 3, 3),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  Whitespace characters are ignored.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a '.', the 'Dot' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner(".")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.Dot, 0, 1),
			newToken(token.EOF, 1, 1),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a '.', the 'Dot' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a '{', the 'LBrace' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("{")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.LBrace, 0, 1),
			newToken(token.EOF, 1, 1),
		)

		// Assert.
		for idx, want := range wantTokens {
			got := scanner.NextToken()

			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a '{', the 'LBrace' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a '}', the 'RBrace' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("}")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.RBrace, 0, 1),
			newToken(token.EOF, 1, 1),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a '}', the 'RBrace' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a '[', the 'LBracket' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("[")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.LBracket, 0, 1),
			newToken(token.EOF, 1, 1),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a '[', the 'LBracket' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a ']', the 'RBracket' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("]")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.RBracket, 0, 1),
			newToken(token.EOF, 1, 1),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a ']', the 'RBracket' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a ':', the 'Colon' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner(":")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.Colon, 0, 1),
			newToken(token.EOF, 1, 1),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a ':', the 'Colon' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a '=', the 'Equals' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("=")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.Equals, 0, 1),
			newToken(token.EOF, 1, 1),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a '=', the 'Equals' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a ',', the 'Comma' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner(",")

		// Act.
		wantTokens := newTokenSet(
			newToken(token.Comma, 0, 1),
			newToken(token.EOF, 1, 1),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a ',', the 'Comma' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a 'string', the 'String' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("\"Hello\"")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.String, "Hello", 0, 7),
			newToken(token.EOF, 7, 7),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a 'string', the 'String' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a 'string' (multi-byte), the 'String' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("\"A🚀C\"")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.String, "A🚀C", 0, 8),
			newToken(token.EOF, 8, 8),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a 'string' (multi-byte), the 'String' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning an invalid 'string' (newline), the 'Error' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("\"Hello\n")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Error, "Unclosed string literal.", 0, 6),
			newToken(token.EOF, 7, 7),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a 'string', the 'String' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning an invalid 'string' (EOF), the 'Error' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("\"Hello")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Error, "Unclosed string literal.", 0, 6),
			newToken(token.EOF, 6, 6),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a 'string', the 'String' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a 'number', the 'Number' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("123")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Number, "123", 0, 3),
			newToken(token.EOF, 3, 3),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a 'number', the 'Number' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning an 'identifier' (normal), the 'Ident' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("apples")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Ident, "apples", 0, 6),
			newToken(token.EOF, 6, 6),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning an 'identifier' (normal), the 'Ident' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning an 'identifier' (with underscores), the 'Ident' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("_apples_and_bananas_")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Ident, "_apples_and_bananas_", 0, 20),
			newToken(token.EOF, 20, 20),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning an 'identifier' (with underscores), the 'Ident' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning an 'identifier' (with numbers), the 'Ident' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("carrots3nuts7")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Ident, "carrots3nuts7", 0, 13),
			newToken(token.EOF, 13, 13),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning an 'identifier' (with numbers), the 'Ident' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning 'true', the 'Bool' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("true")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Bool, "true", 0, 4),
			newToken(token.EOF, 4, 4),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning an 'identifier' (with underscores), the 'Ident' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning 'false', the 'Bool' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("false")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Bool, "false", 0, 5),
			newToken(token.EOF, 5, 5),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning an 'identifier' (with underscores), the 'Ident' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning anything that's not valid, the 'Error' token is returned'.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		scanner := newScanner("@")

		// Act.
		wantTokens := newTokenSet(
			newValueToken(token.Error, "Invalid character '@'.", 0, 1),
			newToken(token.EOF, 1, 1),
		)

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning anything that's not valid, the 'Error' token is returned'.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})

	t.Run("When scanning a complete Lux file, all tokens are correct.", func(t *testing.T) {
		t.Parallel() // Enable parallel execution.

		// Arrange.
		input := `version = 1.0
extension: ".cs" {
    tokens: {
        access: [ "public", "private", "internal" ]
        kind: [ "class", "interface", "enum" ]
        name: alpha
    }
    rule: "The name of an 'interface' must start with an 'I'." {
        match: [ access, "interface", name: interfaceName ]
        interfaceName: {
            starts_with: "I"
        }
        enabled: true
        pass: [ "public interface IUserRepository" ]
        fail: ["public interface UserRepository"]
    }
}`
		scanner := newScanner(input)

		// Act.
		wantTokens := []token.Token{
			newValueToken(token.Ident, "version", 0, 7),
			newToken(token.Equals, 8, 9),
			newValueToken(token.Number, "1", 10, 11),
			newToken(token.Dot, 11, 12),
			newValueToken(token.Number, "0", 12, 13),
			newValueToken(token.Ident, "extension", 14, 23),
			newToken(token.Colon, 23, 24),
			newValueToken(token.String, ".cs", 25, 30),
			newToken(token.LBrace, 31, 32),
			newValueToken(token.Ident, "tokens", 37, 43),
			newToken(token.Colon, 43, 44),
			newToken(token.LBrace, 45, 46),
			newValueToken(token.Ident, "access", 55, 61),
			newToken(token.Colon, 61, 62),
			newToken(token.LBracket, 63, 64),
			newValueToken(token.String, "public", 65, 73),
			newToken(token.Comma, 73, 74),
			newValueToken(token.String, "private", 75, 84),
			newToken(token.Comma, 84, 85),
			newValueToken(token.String, "internal", 86, 96),
			newToken(token.RBracket, 97, 98),
			newValueToken(token.Ident, "kind", 107, 111),
			newToken(token.Colon, 111, 112),
			newToken(token.LBracket, 113, 114),
			newValueToken(token.String, "class", 115, 122),
			newToken(token.Comma, 122, 123),
			newValueToken(token.String, "interface", 124, 135),
			newToken(token.Comma, 135, 136),
			newValueToken(token.String, "enum", 137, 143),
			newToken(token.RBracket, 144, 145),
			newValueToken(token.Ident, "name", 154, 158),
			newToken(token.Colon, 158, 159),
			newValueToken(token.Ident, "alpha", 160, 165),
			newToken(token.RBrace, 170, 171),
			newValueToken(token.Ident, "rule", 176, 180),
			newToken(token.Colon, 180, 181),
			newValueToken(token.String, "The name of an 'interface' must start with an 'I'.", 182, 234),
			newToken(token.LBrace, 235, 236),
			newValueToken(token.Ident, "match", 245, 250),
			newToken(token.Colon, 250, 251),
			newToken(token.LBracket, 252, 253),
			newValueToken(token.Ident, "access", 254, 260),
			newToken(token.Comma, 260, 261),
			newValueToken(token.String, "interface", 262, 273),
			newToken(token.Comma, 273, 274),
			newValueToken(token.Ident, "name", 275, 279),
			newToken(token.Colon, 279, 280),
			newValueToken(token.Ident, "interfaceName", 281, 294),
			newToken(token.RBracket, 295, 296),
			newValueToken(token.Ident, "interfaceName", 305, 318),
			newToken(token.Colon, 318, 319),
			newToken(token.LBrace, 320, 321),
			newValueToken(token.Ident, "starts_with", 334, 345),
			newToken(token.Colon, 345, 346),
			newValueToken(token.String, "I", 347, 350),
			newToken(token.RBrace, 359, 360),
			newValueToken(token.Ident, "enabled", 369, 376),
			newToken(token.Colon, 376, 377),
			newValueToken(token.Bool, "true", 378, 382),
			newValueToken(token.Ident, "pass", 391, 395),
			newToken(token.Colon, 395, 396),
			newToken(token.LBracket, 397, 398),
			newValueToken(token.String, "public interface IUserRepository", 399, 433),
			newToken(token.RBracket, 434, 435),
			newValueToken(token.Ident, "fail", 444, 448),
			newToken(token.Colon, 448, 449),
			newToken(token.LBracket, 450, 451),
			newValueToken(token.String, "public interface UserRepository", 451, 484),
			newToken(token.RBracket, 484, 485),
			newToken(token.RBrace, 490, 491),
			newToken(token.RBrace, 492, 493),
			newToken(token.EOF, 493, 493),
		}

		for idx, want := range wantTokens {
			got := scanner.NextToken()

			// Assert.
			assert.Equalf(t, got, want, "\n\n"+
				"UT Name:  When scanning a complete Lux file, all tokens are correct.\n"+
				"\033[32mExpected: #%d - %s\033[0m\n"+
				"\033[31mActual:   #%d - %s\033[0m\n\n", idx, want, idx, got)
		}
	})
}

// Returns a new input with the given content.
func newScanner(content string) *scanner.Scanner {
	input := &text.Input{
		Content: content,
	}

	return scanner.New(input)
}

// Returns a new set of tokens.
func newTokenSet(toks ...token.Token) []token.Token {
	return toks
}

// Returns a new token with the given type and span.
func newToken(tType token.Type, start, end int) token.Token {
	return token.Token{
		Type: tType,
		Span: newSpan(start, end),
	}
}

// Returns a new token with the given type, value and span.
func newValueToken(tType token.Type, value string, start, end int) token.Token {
	return token.Token{
		Type:    tType,
		Literal: value,
		Span:    newSpan(start, end),
	}
}

// Returns a new span with the given start and end positions.
func newSpan(start, end int) text.Span {
	return text.Span{
		Start: start,
		End:   end,
	}
}
