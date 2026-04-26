package main

import "fmt"

type tokenType int

const (
	number tokenType = iota
	plus
	minus
)

type token struct {
	tokenType tokenType
	value     string
}

func getNumber(char byte) token {
	return token{
		tokenType: number,
		value:     string(char),
	}
}

func getToken(char byte) token {
	switch char {
	case '+':
		return token{tokenType: plus, value: "+"}
	case '-':
		return token{tokenType: minus, value: "-"}
	default:
		return getNumber(char)
	}
}

func main() {
	input := "2 + 5 + 3 - 2"

	tokens := []token{}

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			continue
		}

		tokens = append(tokens, getToken(input[i]))
	}

	for _, t := range tokens {
		fmt.Println(t)
	}
}
