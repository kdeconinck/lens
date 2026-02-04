// =====================================================================================================================
// = LICENSE:       Copyright (c) 2025 Kevin De Coninck
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

/// Note: Lux is a simple configuration language inspired by JSON and TOML.
///       Below is a production-ready example of a lux file.
///
///       version = 1.0
///
///       extension: ".cs" {
///           tokens: {
///               access: [ "public", "private", "internal" ]
///               kind: [ "class", "interface", "enum" ]
///               name: alpha
///           }
///
///           rule: "The name of an 'interface' must start with an 'I'." {
///               match: [ access, "interface", name: interfaceName ]
///               interfaceName: {
///                   starts_with: "I"
///               }
///               enabled: true
///               pass: [ "public interface IUserRepository" ]
///               fail: ["public interface UserRepository"]
///           }
///       }

// Package scanner implements a scanner for the lux language.
package scanner

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/kdeconinck/lens/internal/pkg/lux/token"
	"github.com/kdeconinck/lens/internal/pkg/text"
)

// Scanner transforms a [text.Input] into a stream of [token.Token]s.
type Scanner struct {
	input      *text.Input
	tokenStart int
	pos        int
}

// New initializes a [Scanner] with the provided input.
func New(input *text.Input) *Scanner {
	return &Scanner{
		input: input,
	}
}

// NextToken scans the next token from the input.
// It skips whitespace and comments automatically.
func (scanner *Scanner) NextToken() token.Token {
	scanner.skipWhitespace()
	scanner.tokenStart = scanner.pos

	r := scanner.peek()

	if r == 0 {
		return scanner.emit(token.EOF, "")
	}

	switch r {
	case '.':
		scanner.consume()

		return scanner.emit(token.Dot, "")

	case '{':
		scanner.consume()

		return scanner.emit(token.LBrace, "")

	case '}':
		scanner.consume()

		return scanner.emit(token.RBrace, "")

	case '[':
		scanner.consume()

		return scanner.emit(token.LBracket, "")

	case ']':
		scanner.consume()

		return scanner.emit(token.RBracket, "")

	case ':':
		scanner.consume()

		return scanner.emit(token.Colon, "")

	case '=':
		scanner.consume()

		return scanner.emit(token.Equals, "")

	case ',':
		scanner.consume()

		return scanner.emit(token.Comma, "")

	case '"':
		scanner.consume()

		return scanner.scanString()
	}

	if unicode.IsDigit(r) {
		return scanner.scanNumber()
	}

	if unicode.IsLetter(r) || r == '_' {
		return scanner.scanIdentifier()
	}

	scanner.consume()

	return scanner.emit(token.Error, fmt.Sprintf("Invalid character '%s'.", string(r)))
}

// Keep reading data until the termination of the string.
// A string is terminated if:
// - We hit the string termination character (quote).
// - We hit EOF (this means an error, beacused the string isn't properly terminated).
// - We hit a newline (this means an error, beacused the string isn't properly terminated).
func (scanner *Scanner) scanString() token.Token {
	for {
		r := scanner.peek()

		if r == '"' {
			value := scanner.input.Content[scanner.tokenStart+1 : scanner.pos]

			scanner.consume()

			return scanner.emit(token.String, value)
		}

		if r == 0 || r == '\n' {
			return scanner.emit(token.Error, "Unclosed string literal.")
		}

		scanner.consume()
	}
}

// Keep reading data until the termination of the number.
// A number is terminated if:
// - We hit a character isn't a number (including EOF).
func (scanner *Scanner) scanNumber() token.Token {
	for {
		r := scanner.peek()

		if !unicode.IsDigit(r) {
			break
		}

		scanner.consume()
	}

	value := scanner.input.Content[scanner.tokenStart:scanner.pos]

	return scanner.emit(token.Number, value)
}

// Keep reading data until the termination of the identifier.
// An identifier is terminated if:
// - We hit a character isn't a letter (including EOF).
// - We hit a character isn't a digit (including EOF).
// - We hit a character isn't an underscore "_" (including EOF).
//
// If the identifier is a boolean value (either "true" or "false"), a "Bool" token is emitted, in any other case, an
// "Ident" token is emitted.
func (scanner *Scanner) scanIdentifier() token.Token {
	for {
		r := scanner.peek()

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			break
		}

		scanner.consume()
	}

	value := scanner.input.Content[scanner.tokenStart:scanner.pos]

	if value == "true" || value == "false" {
		return scanner.emit(token.Bool, value)
	}

	return scanner.emit(token.Ident, value)
}

// Keep reading data until EOF or a non-whitespace character is encountered.
func (scanner *Scanner) skipWhitespace() {
	for {
		r := scanner.peek()

		if r == 0 {
			break
		}

		if !unicode.IsSpace(r) {
			break
		}

		scanner.consume()
	}
}

// Look at the next character without consuming it.
func (scanner *Scanner) peek() rune {
	if scanner.pos >= len(scanner.input.Content) {
		return 0
	}

	r, _ := utf8.DecodeRuneInString(scanner.input.Content[scanner.pos:])

	return r
}

// Consumes the next character.
func (scanner *Scanner) consume() rune {
	r, width := utf8.DecodeRuneInString(scanner.input.Content[scanner.pos:])
	scanner.pos += width

	return r
}

// Emit a token that represents the scanned data.
func (scanner *Scanner) emit(t token.Type, lit string) token.Token {
	return token.Token{
		Type:    t,
		Literal: lit,
		Span: text.Span{
			Start: scanner.tokenStart,
			End:   scanner.pos,
		},
	}
}
