package main

type TokenType int

const (
	number TokenType = iota
	plus
	minus
	mul
	div
	greater
	less
	greaterEqual
	lessEqual
	leftParen
	rightParen
	ifStmt
	forStmt
	id
	equal
	eof
)

type Token struct {
	tokenType TokenType
	value     string
}

