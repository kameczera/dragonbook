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
		case greater: {
			if left > right {
				return 1
			} else {
				return 0	
			}
		}
		case greaterEqual:
			if left >= right {
				return 1
			} else {
				return 0
			}
		case less:
			if left < right {
				return 1
			} else {
				return 0
			}
		case lessEqual:
			if left <= right {
				return 1
			} else {
				return 0
			}
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

func interpret(ast AST){
	interpreter := Interpreter{ ast: ast }
	result := interpreter.walk(ast.root)
	fmt.Println(result)
}

