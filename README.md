# MONKEY LANG

Monkey é uma linguagem de programação feita para aprender a criar linguagens de programação.

```monkey
let add = fn(x, y) {
  x + y;
};

let result = add(1, 2);
```

```
                        .="=.
                      _/.-.-.\_     _
                     ( ( o o ) )    ))
                      |/  "  \|    //
      .-------.        \'---'/    //
     _|~~ ~~  |_       /`"""`\\  ((
   =(_|_______|_)=    / /_,_\ \\  \\
     |:::::::::|      \_\\_'__/ \  ))
     |:::::::[]|       /`  /`~\  |//
     |o=======.|      /   /    \  /
jgs  `"""""""""`  ,--`,--'\/\    /
                   '-- "--'  '--'
```

[ASCII arts by ASCII Art Archive](https://www.asciiart.eu/animals/monkeys)

## ANALISE LÉXICA

Para podermos trabalhar com o código fonte precisamos antes transformar texto em algo que possamos trabalhar, e é isso que a analise léxica faz, ela transforma o código fonte em tokens.

Código fonte --[Analise Léxica]-> Tokens --[Analise Sintática]-> AST --[Avaliação]-> Resultado

### Tokens

Tokens são peganas estruturas de dados que carregam informações sobre o código fonte, como o tipo do token, o valor do token e a linha e coluna que o token se encontra.

Nessa primeira versão não vamos nos preocupar com a linha e coluna do token, mas vamos nos preocupar com o tipo e o valor do token. Contudo a informação da linha e da coluna é importante para gerar mensagens de erro mais precisas.

Alguns tokens tem significado especial para a gente, como as palavras chaves e os operadores, mas outros tokens não tem significado especial, como os identificadores e os números.

Veja o exemplo abaixo:

```monkey
let i = 42;
```

Ao processar esse input, vamos gerar os seguintes tokens:

| Tipo  | Valor |
| ----- | ----- |
| LET   | let   |
| IDENT | i     |
| =     | =     |
| INT   | 42    |
| ;     | ;     |

Os tipos podem variar de acordo com a sua implementação, mas o conceito é o mesmo.

#### Palavras Chaves

Palavras chaves são palavras que tem significado especial para a linguagem, como `let`, `fn`, `true`, `false`, `if`, `else`, `return`, etc.

#### Identificadores

Identificadores são nomes que damos para variáveis, funções, etc. No estágio atual somente identificadores declarados com caracteres de `a` a `z` e `_` são suportados.

#### Números

Números são números, como `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`. No momento só suportamos inteiros.

#### Espaços em branco

Espaços em branco são espaços, tabulações e quebras de linha. Eles são usados para separar tokens, mas não tem significado especial então a gente só DEVORA eles e segue em frente.

```
================================================.
     .-.   .-.     .--.                         |
    | OO| | OO|   / _.-' .-.   .-.  .-.   .''.  |
    |   | |   |   \  '-. '-'   '-'  '-'   '..'  |
    '^^^' '^^^'    '--'                         |
===============.  .-.  .================.  .-.  |
               | |   | |                |  '-'  |
               | |   | |                |       |
               | ':-:' |                |  .-.  |
l42            |  '-'  |                |  '-'  |
==============='       '================'       |
```

### Lexer

O lexer é a ferramenta que nos permite transformar o código fonte em tokens, ele lê o código fonte e gera os tokens.
Nós pegamos a string de entrada e transformamos em uma lista de tokens.

Fazemos isso percorrendo a string de entrada e consumindo os caracteres um a um. Por isso comentei que o lexer devora os espaços em branco, ele consome os caracteres que não tem significado especial e gera os tokens que tem significado especial. Fazemos isso até encontrar um carácter que não sabemos o que fazer, nesse caso geramos um erro, ou até o fim da string.

## REPL

- `R`ead
- `E`valuate
- `P`rint
- `L`oop

O REPL é uma ferramenta que nos permite executar código fonte de forma interativa, ele lê o código fonte, avalia o código fonte, imprime o resultado e repete tudo de novo.

REPLs são muito úteis para testar código e para aprender uma linguagem de programação.

Eles foram popularizados pelo [Lisp](<https://en.wikipedia.org/wiki/Lisp_(programming_language)>) e hoje em dia são muito comuns em diversas linguagens de programação.

```bash
go run main.go
```

## PARSING

Um parser transform um determinado input em uma estrutura de dados hierárquica, no nosso caso o parser transforma os tokens em uma árvore sintática abstrata (AST).

Uma AST é uma estrutura de dados que representa o código fonte de forma hierárquica, ela é muito útil para avaliar o código fonte.

Existem ferramentas capazes de gerar parsers a partir de uma gramática, mas nesse projeto o parser é feito manualmente.

Essa "gramática" citadada acima costuma ser uma [Context Free Grammar (Gramática livre de contexto)](https://en.wikipedia.org/wiki/Context-free_grammar) (CFG), que é uma forma de descrever uma linguagem.

### Gramática

Uma gramática é uma forma de descrever uma linguagem, ela é composta por regras que descrevem como a linguagem funciona e como você pode construir sentenças válidas.

Existem diversas formas de escrever uma gramática, mas a mais comum é a [Backus-Naur Form](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form) (BNF) e a [Extended Backus-Naur Form](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form) (EBNF).

```bnf
<regra> ::= <expressão>
```

Exemplo simples:

```bnf
<operador> ::= +|-|/|*
```

A regra acima descreve um operador, que pode ser `+`, `-`, `/` ou `*`.

Exemplo mais complexo:

```bnf
	chunk ::= block

	block ::= {stat} [retstat]

	stat ::=  ‘;’ |
		 varlist ‘=’ explist |
		 functioncall |
		 label |
		 break |
		 goto Name |
		 do block end |
		 while exp do block end |
		 repeat block until exp |
		 if exp then block {elseif exp then block} [else block] end |
		 for Name ‘=’ exp ‘,’ exp [‘,’ exp] do block end |
		 for namelist in explist do block end |
		 function funcname funcbody |
		 local function Name funcbody |
		 local attnamelist [‘=’ explist]

	attnamelist ::=  Name attrib {‘,’ Name attrib}

	attrib ::= [‘<’ Name ‘>’]

	retstat ::= return [explist] [‘;’]

	label ::= ‘::’ Name ‘::’

	funcname ::= Name {‘.’ Name} [‘:’ Name]

	varlist ::= var {‘,’ var}

	var ::=  Name | prefixexp ‘[’ exp ‘]’ | prefixexp ‘.’ Name

	namelist ::= Name {‘,’ Name}

	explist ::= exp {‘,’ exp}

	exp ::=  nil | false | true | Numeral | LiteralString | ‘...’ | functiondef |
		 prefixexp | tableconstructor | exp binop exp | unop exp

	prefixexp ::= var | functioncall | ‘(’ exp ‘)’

	functioncall ::=  prefixexp args | prefixexp ‘:’ Name args

	args ::=  ‘(’ [explist] ‘)’ | tableconstructor | LiteralString

	functiondef ::= function funcbody

	funcbody ::= ‘(’ [parlist] ‘)’ block end

	parlist ::= namelist [‘,’ ‘...’] | ‘...’

	tableconstructor ::= ‘{’ [fieldlist] ‘}’

	fieldlist ::= field {fieldsep field} [fieldsep]

	field ::= ‘[’ exp ‘]’ ‘=’ exp | Name ‘=’ exp | exp

	fieldsep ::= ‘,’ | ‘;’

	binop ::=  ‘+’ | ‘-’ | ‘*’ | ‘/’ | ‘//’ | ‘^’ | ‘%’ |
		 ‘&’ | ‘~’ | ‘|’ | ‘>>’ | ‘<<’ | ‘..’ |
		 ‘<’ | ‘<=’ | ‘>’ | ‘>=’ | ‘==’ | ‘~=’ |
		 and | or

	unop ::= ‘-’ | not | ‘#’ | ‘~’
```

O exemplo acima é [toda a sintaxe](https://www.lua.org/manual/5.4/manual.html#9) da linguagem [Lua](https://www.lua.org/).

Um gerador de parser pegaria uma sintaxe como essa e geraria um parser em uma linguagem de programação como C ou Java.

Mas enfim, vamos voltar para o nosso parser.

### Escrevendo um parser

Existem diversas estratégias para escrever um parser, mas a mais comum é a [Recursive Descent](https://en.wikipedia.org/wiki/Recursive_descent_parser), que é uma estratégia top-down, ou seja, a gente começa do topo da árvore e vai descendo até as folhas. É como vamos fazer o nosso.

Outras variações são [LL](https://en.wikipedia.org/wiki/LL_parser) e [LL(k)](<https://en.wikipedia.org/wiki/LL_parser#LL(k)_parsers>), que são parsers top-down que usam lookahead para decidir qual regra aplicar.

Parsers bottom-up são mais complexos e geralmente são gerados a partir de uma gramática, como é o caso do [LALR](https://en.wikipedia.org/wiki/LALR_parser) e [LR](https://en.wikipedia.org/wiki/LR_parser).

## ESTUDAR

- [ ] Context Free Grammar
- [ ] Backus-Naur Form
- [ ] Extended Backus-Naur Form

## MELHORIAS

- [ ] Suporte a unicode (byte -> rune)
- [ ] Suporte a números de ponto flutuante
- [ ] Suporte a números hexadecimais
- [ ] Suporte a números binários
- [ ] Suporte a números octais
- [ ] Suporte a "\_" em números para melhorar a legibilidade
- [ ] Suporte a comentários com `#`
