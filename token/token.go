package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identificadores + literais
	IDENT = "IDENT"
	INT   = "INT"

	// Operadores
	ASSIGN = "="
	PLUS   = "+"

	// Delimitadores
	COMMA     = ","
	SEMICOLON = ";"

	// Caracteres especiais
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Palavras chave
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
