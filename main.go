package main

import (
	"fmt"
)

func main() {
	input := "2 + 5 + 3 - 2 * 7"

	lexer := getTokens(input)

	lexer.printTokens()
	ast := createTree(lexer.tokens)
	fmt.Println("\nast:")
	printTree(ast.root, "")
}
