package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "2 + 5 + 3 / 2 + 7"

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
