package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
	statements []Node
	tokens []Token
	pos int
}

func (p *Parser) current() Token {
	if p.pos >= len(p.tokens) {
		return Token{eof, " "}
	}
	return p.tokens[p.pos]
}

func (p *Parser) advance() Token {
	t := p.current()
	p.pos++
	return t
}

func matchTokenType(token Token, types ...TokenType) bool {
	for _, val := range types {
		if token.tokenType == val {
			return true
		}
	}
	return false
}

func createTrees(tokens []Token) []AST {
	parser := Parser{ tokens: tokens }
	asts := []AST{}
	for parser.pos < len(parser.tokens) {
		asts = append(asts, AST { root: parser.statement() })
	}
	return asts
}

func (p *Parser) statement() Node {
	node := p.expr()

	semic := p.advance()
	if !matchTokenType(semic, semicolon) {
		panic(fmt.Sprintf("esperado uma ';', mas encontrou um %s", semic.value))
	}
	return node
}

func (p *Parser) expr() Node {
	left := p.term()

	for matchTokenType(p.current(), plus, minus) {
		op := p.advance()
		right := p.term()

		left = Binary{
			op: op,
			left: left,
			right: right,
		}
	}
	return left
}

func (p *Parser) term() Node {
	left := p.rel()

	for matchTokenType(p.current(), mul, div) {
		op := p.advance()
		right := p.factor()

		left = Binary{
			op: op,
			left: left,
			right: right,
		}
	} 

	return left
}

func (p *Parser) rel() Node {
	left := p.assign()

	for matchTokenType(p.current(), greater, greaterEqual, less, lessEqual) {
		op := p.advance()
		right := p.factor()
		left = Binary {
			left: left,
			op: op,
			right: right,
		}
	}
	return left
}

func (p *Parser) assign() Node {
	variable := p.factor()

	for matchTokenType(p.current(), equal) {
		p.advance()
		v, ok := variable.(Variable)
		if !ok {
			panic(fmt.Sprintf("esperado uma variable, mas encontrou um %T", variable))
		}
		value := p.expr()
		variable = Assign {
			variable: v,
			value: value,
		}
	}
	return variable
}

func (p *Parser) factor() Node {
	token := p.advance()
	var node Node
	if matchTokenType(token, number) {
		value, err := strconv.Atoi(token.value)
		if err != nil {
			panic("numero invalido: " + token.value)
		}
		node = Number { value: value }
	} else if matchTokenType(token, id) {
		node = Variable { identifier: token.value }
	} else if matchTokenType(token, leftParen) {
		node = p.expr()

		if !matchTokenType(p.advance(), rightParen) {
			panic("esperado ')'")
		}
	
	} else if matchTokenType(token, printStmt) {
		if !matchTokenType(p.advance(), leftParen) {
			panic("esperado '('")
		}
		node = p.expr()
		if !matchTokenType(p.advance(), rightParen) {
			panic("esperado ')'")
		}
		node = Print { value: node }
	} else {
		panic("token posicao invalida")
	}
	return node
}


func printTree(currNode Node, indent string) {
	switch n := currNode.(type) {
		case Number:
			fmt.Println(indent + "Number: ", n.value)
		case Variable:
			fmt.Println(indent + "Variable: ", n.identifier)
		case Assign:
			fmt.Println(indent + "Assign: ", n.variable)
			printTree(n.value, indent + "   ")
		case Binary:
			fmt.Println(indent + "Binary: ", n.op.value)
			printTree(n.left, indent + "   ")
			printTree(n.right, indent + "   ")
		case Print:
			fmt.Println(indent + "Print: ")
			printTree(n.value, indent + "   ")
		default:
			fmt.Println("Error")
	}
}

type Visit func(currNode Node)

func prefixVisit(currNode Node) {
	switch n := currNode.(type) {
		case Number:
			fmt.Print(n.value)

		case Binary:
			fmt.Print(n.op.value)
			prefixVisit(n.left)
			prefixVisit(n.right)
		default:
			fmt.Print("Error")
	}
}

func infixVisit(currNode Node) {
	switch n := currNode.(type) {
		case Number:
			fmt.Print(n.value)

		case Binary:
			infixVisit(n.left)
			fmt.Print(n.op.value)
			infixVisit(n.right)
		case Variable:
			fmt.Print(n.identifier)
		case Assign:
			
		default:
			fmt.Print("Error")
	}
	
}

func posfixVisit(currNode Node) {
	switch n := currNode.(type) {
		case Number:
			fmt.Print(n.value)

		case Binary:
			posfixVisit(n.left)
			posfixVisit(n.right)
			fmt.Print(n.op.value)
		default:
			fmt.Print("Error")
	}
}
