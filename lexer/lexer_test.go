package lexer

import (
	"testing"

	"github.com/mxiaole/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := "=+{},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
            t.Fatalf("%d, expectedType: %q, gotType: %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
            t.Fatalf("%d, expectedLiteral: %q, gotLiteral: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
