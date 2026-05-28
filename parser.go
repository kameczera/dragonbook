package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
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

func createTree(tokens []Token) AST {
	parser := Parser{ tokens: tokens }
	return AST { root: parser.expr() }
}

func (p *Parser) term() Node {
	left := p.factor()

	for matchTokenType(p.current(), mul, div) {
		op := p.advance()
		right := p.factor()

		left = Binary{
			op:    op,
			left:  left,
			right: right,
		}
	}
	return left
}

func (p *Parser) expr() Node {
	left := p.term()

	for matchTokenType(p.current(), plus, minus) {
		op := p.advance()
		right := p.term()

		left = Binary{
			op:    op,
			left:  left,
			right: right,
		}
	}
	return left
}

func (p *Parser) factor() Node {
	token := p.advance()
	var node Node
	if matchTokenType(token, number) {
		value, err := strconv.Atoi(token.value)
		if err != nil {
			panic("numero invalido: " + token.value)
		}
		node = Value { value: value }
	} else if matchTokenType(token, leftParen) {
		node = p.expr()

		if !matchTokenType(p.current(), rightParen) {
			panic("esperado ')'")
		}
		p.advance()
		return node
	} else {
		panic("token posicao invalida")
	}
	return node
}


func printTree(currNode Node, indent string) {
	switch n := currNode.(type) {
		case Value:
			fmt.Println(indent + "Value: ", n.value)

		case Binary:
			fmt.Println(indent + "Binary: ", n.op.value)
			printTree(n.left, indent + "   ")
			printTree(n.right, indent + "   ")
		default:
			fmt.Println("Error")
	}
}

type Visit func(currNode Node)

func prefixVisit(currNode Node) {
	switch n := currNode.(type) {
		case Value:
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
		case Value:
			fmt.Print(n.value)

		case Binary:
			infixVisit(n.left)
			fmt.Print(n.op.value)
			infixVisit(n.right)
		default:
			fmt.Print("Error")
	}
	
}

func posfixVisit(currNode Node) {
	switch n := currNode.(type) {
		case Value:
			fmt.Print(n.value)

		case Binary:
			posfixVisit(n.left)
			posfixVisit(n.right)
			fmt.Print(n.op.value)
		default:
			fmt.Print("Error")
	}
}