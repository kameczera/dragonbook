package main

import (
	"fmt"
)

func main() {
	// input := "(2<5) + 1 * 7"
	input := "abc = 1324;print(abc + 123)"
	lexer := getTokens(input)

	lexer.printTokens()
	asts := createTrees(lexer.tokens)
	for _, ast := range asts {
		fmt.Println("\nast:")
		printTree(ast.root, "")
		interpret(ast)
	}
	// fmt.Println("\n")
	// prefixVisit(ast.root)
	// fmt.Println("\n")
	// infixVisit(ast.root)
	// fmt.Println("\n")
	// posfixVisit(ast.root)	
	// fmt.Println("\n\n")
}
