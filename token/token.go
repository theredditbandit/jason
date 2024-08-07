package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"
	STR   = "STR"

	LBRACE = "{"
	RBRACE = "}"
	COMMA  = ","
	COLON  = ":"
)
