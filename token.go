package main

type TokenType int

const (
	number TokenType = iota
	plus
	minus
	mul
	div
	leftParen
	rightParen
	eof
)

type Token struct {
	tokenType TokenType
	value     string
}

