package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	VALUE = "VALUE"
	KEY   = "KEY"

	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	ARRAY  = "ARRAY"

	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	DQUOTE    = "\""
	SQUOTE    = "'"
	LBRACKET  = "["
	RBRACKET  = "]"
	NULL      = "NULL"
)
