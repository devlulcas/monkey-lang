package lexer

import (
	"lucasrego.tech/monkey-lang/token"
)

type Lexer struct {
	input        string // Código fonte
	position     int    // Posição atual no input (caractere atual)
	readPosition int    // Posição de leitura atual (após o caractere atual)
	// O próximo passo seria usar runas ao invés de bytes, para suportar UTF-8
	ch byte // Caractere que está sendo analisado
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Inicializa o lexer
	return l
}

// Lê o próximo caractere e avança nossa posição no input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// EOF (End of File) - Fim do arquivo
		l.ch = 0
	} else {
		// Lê o próximo caractere
		l.ch = l.input[l.readPosition]
	}

	// Avança a posição atual e a posição de leitura
	l.position = l.readPosition
	l.readPosition += 1
}

// Lê o próximo caractere, porém não avança nossa posição no input (só damos uma curiadinha)
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tkn token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			initialCh := l.ch

			// Consome o primeiro '='  e vai pro próximo '='
			l.readChar()

			tkn = token.Token{
				Type:    token.EQ,
				Literal: string(initialCh) + string(l.ch),
			}
		} else {
			tkn = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			initialCh := l.ch

			// Consome o primeiro '!'  e vai pro '='
			l.readChar()

			tkn = token.Token{
				Type:    token.NOT_EQ,
				Literal: string(initialCh) + string(l.ch),
			}
		} else {
			tkn = newToken(token.BANG, l.ch)
		}
	case ';':
		tkn = newToken(token.SEMICOLON, l.ch)
	case '(':
		tkn = newToken(token.LPAREN, l.ch)
	case ')':
		tkn = newToken(token.RPAREN, l.ch)
	case '{':
		tkn = newToken(token.LBRACE, l.ch)
	case '}':
		tkn = newToken(token.RBRACE, l.ch)
	case ',':
		tkn = newToken(token.COMMA, l.ch)
	case '+':
		tkn = newToken(token.PLUS, l.ch)
	case '-':
		tkn = newToken(token.MINUS, l.ch)
	case '*':
		tkn = newToken(token.ASTERISK, l.ch)
	case '/':
		tkn = newToken(token.SLASH, l.ch)
	case '<':
		tkn = newToken(token.LT, l.ch)
	case '>':
		tkn = newToken(token.GT, l.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// Além de caracteres especiais leremos identificadores e palavras-chave
			tkn.Literal = l.readIdentifier()
			tkn.Type = token.LookupIdentifier(tkn.Literal)

			// Como já consumimos toda a extensão de caracteres do identificador
			// não precisamos ler mais caracteres com o readChar(), então é melhor
			// retornar antecipadamente
			return tkn
		} else if isDigit(l.ch) {
			// Precisamos ler dígitos
			tkn.Literal = l.readNumber()
			tkn.Type = token.INT

			return tkn
		} else {
			tkn = newToken(token.ILLEGAL, l.ch)
		}
	}

	// Consome o caractere lido
	l.readChar()

	return tkn
}

func (l *Lexer) eatWhitespace() {
	// Espaços, tabulações, quebra de linha e recuos não significam nada para nós
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		// Nhom nhom consome caracteres em branco
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	// Inicio do possível identificador/palavra-chave
	position := l.position

	// Consome caracteres até encontrar um que não seja uma letra
	for isLetter(l.ch) {
		l.readChar()
	}

	// Captura o literal do inicio até o fim
	return l.input[position:l.position]
}

// No momento essa função é basicamente igual a de identificadores
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	// Aceitamos caracteres ASCII [A-Za-z_]
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	// Aceitamos caracteres entre [0-9_]
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
