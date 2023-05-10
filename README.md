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

Nessa primeira versão não vamos nos preocupar com a linha e coluna do token, mas vamos nos preocupar com o tipo e o valor do token.

Alguns token tem significado especial para a gente, como as palavras chaves e os operadores, mas outros tokens não tem significado especial, como os identificadores e os números.

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

## MELHORIAS

- [ ] Suporte a unicode
- [ ] Suporte a números de ponto flutuante
- [ ] Suporte a números hexadecimais
- [ ] Suporte a números binários
- [ ] Suporte a números octais
- [ ] Suporte a "_" em números para melhorar a legibilidade
- [ ] Suporte a comentários com `#`

