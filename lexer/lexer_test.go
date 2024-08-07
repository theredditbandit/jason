package lexer

import (
	"testing"

	"github.com/theredditbandit/jason/token"
)

func TestLexer(t *testing.T) {
	input := `
    {
    "a":1,
    }
    `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.IDENT, "a"},
		{token.COLON, ":"},
		{token.INT, "1"},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests [%d] - tokenType wrong. expected=%q, got=%q for=%q", i, tt.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests [%d] - expectedLiteral wrong. expected=%q, got=%q for=%q", i, tt.expectedLiteral, tok.Literal, tok.Type)
		}
	}

}
