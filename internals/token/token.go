package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifiers + literals
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}
