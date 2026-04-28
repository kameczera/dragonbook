package main

type TokenType int

const (
	number TokenType = iota
	plus
	minus
	leftParen
	rightParen
	eof
)

type Token struct {
	tokenType TokenType
	value     string
}

