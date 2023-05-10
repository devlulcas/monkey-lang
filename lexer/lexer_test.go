package lexer

import (
	"testing"

	"lucasrego.tech/monkey-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);

	! == < > !=

	true false
	
	- + * /

	if else return 
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// let five = 5;
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// let ten = 10;
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		// let add = fn(x, y) { x + y; };
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		// let result = add(five, ten);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// ! == < > !=
		{token.BANG, "!"},
		{token.EQ, "=="},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.NOT_EQ, "!="},

		// true false
		{token.TRUE, "true"},
		{token.FALSE, "false"},

		// - + * /
		{token.MINUS, "-"},
		{token.PLUS, "+"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},

		// if else return
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.RETURN, "return"},

		// EOF
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
