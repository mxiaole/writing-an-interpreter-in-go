package lexer

import (
	"testing"

	"github.com/mxiaole/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := "=+{},;"

	l := New(input)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
	}

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("%d, expectedLiteral: %q, gotLiteral: %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("%d, expectedLiteral: %q, gotLiteral: %q", i, tt.expectedType, tok.Literal)
		}
	}
}
