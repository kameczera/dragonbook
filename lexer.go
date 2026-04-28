package main

type Lexer struct {
	program string
	pos int
}

func (l *Lexer) advance() byte {
	if l.pos >= len(l.program) {
		
	}
}

func (l *Lexer) getToken(char byte) Token {
	switch char {
	case '+':
		return Token{tokenType: plus, value: "+"}
	case '-':
		return Token{tokenType: minus, value: "-"}
	case '*':
		return Token{tokenType: mul, value: "*"}
	case '/':
		return Token{tokenType: div, value: "/"}
	case '(':
		return Token{tokenType: leftParen, value: "("}
	case ')':
		return Token{tokenType: rightParen, value: ")"}
	default:
		return getNumber(char)
	}
}

func printTokens(tokens []Token) {
	for _, t := range tokens {
		fmt.Println(t)
	}
}

func getNumber(char byte) Token {
	// TODO: Pegar mais de um numero
	return Token {
		tokenType: number,
		value:     string(char),
	}
}
