package main

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
	left := p.term()

	for matchTokenType(p.current(), plus, minus) {
		 op := p.advance()
