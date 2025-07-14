package token

/*
	Token type constants
*/
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers & literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	//	delimiters
	COMMA     = ","
	SEMICOLON = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

/*
	Token data structure definition
*/
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
