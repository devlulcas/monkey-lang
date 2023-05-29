package repl

import (
	"bufio"
	"fmt"
	"io"

	"lucasrego.tech/monkey-lang/lexer"
	"lucasrego.tech/monkey-lang/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tkn := l.NextToken(); tkn.Type != token.EOF; tkn = l.NextToken() {
			fmt.Printf("%+v\n", tkn)
		}
	}
}
