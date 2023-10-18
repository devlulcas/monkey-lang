package ast

import (
	"lucasrego.tech/monkey-lang/token"
)

// Os nós na nossa árvore de sintaxe abstrata implementam essa interface
// para recuperamos o token a qual ele está associado. Essa função é usada apenas em debug e nos nossos testes
type Node interface {
	TokenLiteral() string
}

// Statements são nós que não retornam valor
type Statement interface {
	Node
	statementNode()
}

// Expressions são nós que retornam valor
type Expression interface {
	Node
	expressionNode()
}

// Este é o nó raiz do nosso programa (por isso ele se chama `Program`)
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Estrutura para armazenar um identificador. Essa estrutura implementa a interface de expressões.
// Iremos usar essa estrutura para reduzir a variedade de nós criados.
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Este é o nó que representa statements como `let boo = 1;`
type LetStatement struct {
	Token token.Token // O token que representa o `let` token.LET
	Name  *Identifier // O identificador usado como nome `boo` (O nome da variável em si, não retorna valores, mas usamos Identifier mesmo assim para simplificar o programa)
	Value Expression  // O valor, no exemplo temos `1`, mas pode ser qualquer expressão
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

//
