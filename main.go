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

type Node interface {  }

type Value struct {
	value int
}

type Binary struct {
	op Token
	left Node
	right Node
}

type AST struct {
	root Node
}

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

func (p *Parser) expr() Node {
	left := p.factor()

	for matchTokenType(p.current(), plus, minus) {
		 op := p.advance()
		 right := p.factor()

		 left = Binary {
			 op: op,
			 left: left,
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

func printTokens(tokens []Token) {
	for _, t := range tokens {
		fmt.Println(t)
	}
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
