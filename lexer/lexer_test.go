package lexer

import (
	"testing"

	"github.com/theredditbandit/jason/token"
)

func TestLexer(t *testing.T) {
	input := `
    {
    "abcd":12345,
	"key":"Value",
	"key2":9,
	"Space separated key":"Space separated alphanumeric value 3"
    }
    `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "abcd"},
		{token.COLON, ":"},
		{token.INT, "12345"},
		{token.COMMA, ","},
		{token.STRING, "key"},
		{token.COLON, ":"},
		{token.STRING, "Value"},
		{token.COMMA, ","},
		{token.STRING, "key2"},
		{token.COLON, ":"},
		{token.INT, "9"},
		{token.COMMA, ","},
		{token.STRING, "Space separated key"},
		{token.COLON, ":"},
		{token.STRING, "Space separated alphanumeric value 3"},
		{token.RBRACE, "}"},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests [%d] - tokenType wrong. expected=%q, got=%q for=%v", i, tt.expectedType, tok.Type, tt)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests [%d] - expectedLiteral wrong. expected=%q, got=%q for=%q", i, tt.expectedLiteral, tok.Literal, tok.Type)
		}
	}

}
