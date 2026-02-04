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

// QA: Verify the implementation of the "token" package.
package token_test

import (
	"testing"

	"github.com/kdeconinck/lens/internal/pkg/assert"
	"github.com/kdeconinck/lens/internal/pkg/lux/token"
)

// UT: Get the human-readable representation of a single string.
func Test_TypeString(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	for tcName, tc := range map[string]struct {
		typeInput token.Type
		want      string
	}{
		"When the token is 'Error' it's displayed as 'Error'.": {
			typeInput: token.Error,
			want:      "Error",
		},
		"When the token is 'EOF' it's displayed as 'EOF'.": {
			typeInput: token.EOF,
			want:      "EOF",
		},
		"When the token is 'Ident' it's displayed as 'Identifier'.": {
			typeInput: token.Ident,
			want:      "Identifier",
		},
		"When the token is 'Number' it's displayed as 'Number'.": {
			typeInput: token.Number,
			want:      "Number",
		},
		"When the token is 'String' it's displayed as 'String'.": {
			typeInput: token.String,
			want:      "String",
		},
		"When the token is 'Bool' it's displayed as 'Boolean'.": {
			typeInput: token.Bool,
			want:      "Boolean",
		},
		"When the token is 'Dot' it's displayed as '.'.": {
			typeInput: token.Dot,
			want:      ".",
		},
		"When the token is 'LBrace' it's displayed as '{'.": {
			typeInput: token.LBrace,
			want:      "{",
		},
		"When the token is 'RBrace' it's displayed as '}'.": {
			typeInput: token.RBrace,
			want:      "}",
		},
		"When the token is 'LBracket' it's displayed as '['.": {
			typeInput: token.LBracket,
			want:      "[",
		},
		"When the token is 'RBracket' it's displayed as ']'.": {
			typeInput: token.RBracket,
			want:      "]",
		},
		"When the token is 'Colon' it's displayed as ':'.": {
			typeInput: token.Colon,
			want:      ":",
		},
		"When the token is 'Equals' it's displayed as '='.": {
			typeInput: token.Equals,
			want:      "=",
		},
		"When the token is 'Comma' it's displayed as ','.": {
			typeInput: token.Comma,
			want:      ",",
		},
		"When the token is NOT known it's displayed as 'Unknown(xxx)'.": {
			typeInput: token.Type(100),
			want:      "Unknown(100)",
		},
	} {
		t.Run(tcName, func(t *testing.T) {
			t.Parallel() // Enable parallel execution.

			// Act.
			got := tc.typeInput.String()

			// Assert.
			assert.Equalf(t, got, tc.want, "\n\n"+
				"UT Name:  %s.\n"+
				"\033[32mExpected: %s.\033[0m\n"+
				"\033[31mActual:   %s.\033[0m\n\n", tcName, tc.want, got)
		})
	}
}
