package lexer

import "github.com/theredditbandit/jason/token"

type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // points to the next char
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // reached the end of file
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '"':
		l.readChar() // advance position to first character after the first DQUOTE
		tok.Literal = l.readString()
		tok.Type = token.STRING
		l.readChar() // consumes the second double quote
		return tok
	case ':':
		tok = newToken(token.COLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	default:
		if string(l.ch) == `n` { // null char
			tok.Literal = l.readNull()
			tok.Type = token.NULL
			return tok
		} else {
			tok.Literal = l.readInt()
			tok.Type = token.INT
			return tok
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readNull() string {
	for i := 0; i < 4; i++ {
		l.readChar() // length of null
	}
	return "null"
}

func (l *Lexer) readInt() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position
	for l.ch != '"' {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
