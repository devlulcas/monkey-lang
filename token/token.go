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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	EQ       = "=="
	NOT_EQ   = "!="
	BANG     = "!"
	LT       = "<"
	GT       = ">"

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
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var reservedKeywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdentifier(identifier string) TokenType {
	// Se achar na lista ok, é uma palavra reservada
	if tkn, ok := reservedKeywords[identifier]; ok {
		return tkn
	}

	// Se não é só um identificador definido pelo usuário
	return IDENT
}
