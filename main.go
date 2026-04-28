package main

import (
	"fmt"
	"strconv"
)

func getNumber(char byte) Token {
	// TODO: Pegar mais de um numero
	return Token{
		tokenType: number,
		value:     string(char),
	}
}

func getToken(char byte) Token {
	switch char {
	case '+':
		return Token{tokenType: plus, value: "+"}
	case '-':
		return Token{tokenType: minus, value: "-"}
	case '(':
		return Token{tokenType: leftParen, value: "("}
	case ')':
		return Token{tokenType: rightParen, value: ")"}
	default:
		return getNumber(char)
	}
}

func main() {
	input := "2 + 5 + 3 - 2 * 7"

	tokens := []Token{}

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			continue
		}

		tokens = append(tokens, getToken(input[i]))
	}
	printTokens(tokens)
	ast := createTree(tokens)
	fmt.Println("\nast:")
	printTree(ast.root, "")
}
