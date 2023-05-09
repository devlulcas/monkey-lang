package lexer

type Lexer struct {
	input        string // Código fonte
	position     int    // Posição atual no input (caractere atual)
	readPosition int    // Posição de leitura atual (após o caractere atual)
	ch           byte   // Caractere que está sendo analisado
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

// Lê o próximo caractere e avança nossa posição no input
func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
