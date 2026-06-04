package main

import (
	"fmt"
)

type Interpreter struct {
	ast AST 
}

func (i *Interpreter) interpretOp(n Binary) int {
	left := i.walk(n.left)
	right := i.walk(n.right)

	
	switch n.op.tokenType {
		case plus:
			return left + right
		case minus:
			return left - right
		case mul:
			return left * right
		case div:
			return left / right
		default:
			panic("operador desconhecido")
	}
}

func (i *Interpreter) walk(currNode Node) int {
	switch n := currNode.(type) {
		case Value:
			return n.value
		case Binary:
			return i.interpretOp(n)
		default:
			panic("nó desconhecido")
	}
}

func (i *Interpreter) interpret(ast AST) {
	
	result := i.walk(i.ast.root)
	fmt.Println(result)
}

