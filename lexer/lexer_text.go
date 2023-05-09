package lexer

import (
	"testing"

	"lucasrego.tech/monkey-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lex := New(input)

	for i, testToken := range tests {
		tkn := lex.NextToken()

		if tkn.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, testToken.expectedLiteral, tkn.Literal)
		}
	}
}
