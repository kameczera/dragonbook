package main

import (
	"fmt"
)

func main() {
	input := "(2<5) + 1 * 7 asdf * 2"

	lexer := getTokens(input)

	lexer.printTokens()
	ast := createTree(lexer.tokens)
	fmt.Println("\nast:")
	printTree(ast.root, "")
	// fmt.Println("\n")
	// prefixVisit(ast.root)
	// fmt.Println("\n")
	// infixVisit(ast.root)
	// fmt.Println("\n")
	// posfixVisit(ast.root)	
	// fmt.Println("\n\n")
	interpret(ast)
}
