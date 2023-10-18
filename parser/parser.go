package parser

import (
	"lucasrego.tech/monkey-lang/ast"
	"lucasrego.tech/monkey-lang/lexer"
	"lucasrego.tech/monkey-lang/token"
)

// A estrutura para o parser funciona de forma semelhante ao lexer, mas ao invés de lidarmos com caracteres usamos tokens
type Parser struct {
	l         *lexer.Lexer // Um ponteiro para uma instância do lexer que usamos para ficar chamando NextToken()
	curToken  token.Token  // Funciona como `position` dentro do módulo lexer
	peekToken token.Token  // Funciona como `readPosition` dentro do módulo lexer
}

// Instancia um novo parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Lê dois token para que curToken e peekToken sejam setados
	// | 0 | curToken = nada                       | peekToken = nada                       |
	// | 1 | curToken = nada                       | peekToken = primeiro token do programa |
	// | 2 | curToken = primeiro token do programa | peekToken = segundo token do programa  |
	p.nextToken()
	p.nextToken()

	return p
}

// Avança na leitura de tokens
func (p *Parser) nextToken() {
	// Vai puxando os tokens de um a um. Você pode imaginar uma string de código fonte como uma fila.
	// O código a seguir pega o token que tá na frente da fila (p.peekToken) e coloca numa salinha (p.curToken)
	// para ser examinado. Logo depois colocamos mais um token da fila de código fonte usando o lexer (p.l.NextToken)
	// na ponta da fila para ser o próximo a ser examinado.
	p.curToken = p.peekToken()
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
