package main

import (
	"fmt"
)

type Lexer struct {
	program string
	pos int
	tokens []Token
}

func getTokens(program string) Lexer {
	lexer := Lexer { program: program, pos: 0 }

	for {
		if lexer.pos >= len(lexer.program) {
			break
		}
		if program[lexer.pos] == ' ' {
			lexer.pos++
			continue
		}
		lexer.tokens = append(lexer.tokens, lexer.getToken())
		lexer.pos++
	}
	return lexer
}

func (l *Lexer) getNumber() Token {	
	init := l.pos
	for l.pos < len(l.program) {
		if isNumeric(l.program[l.pos]) {
			return Token {
				tokenType: number,
				value: l.program[init:l.pos + 1],
			}
		}
		l.pos++
	}
	panic("not a number")
}

func isNumeric(c byte) bool {
	if c < '0' || c > '9' {
		return false
	}
	return true
}

func (l *Lexer) getToken() Token {
	switch l.program[l.pos] {
		case '+':
			return Token{tokenType: plus, value: "+"}
		case '-':
			return Token{tokenType: minus, value: "-"}
		case '*':
			return Token{tokenType: mul, value: "*"}
		case '/':
			return Token{tokenType: div, value: "/"}
		case '(':
			return Token{tokenType: leftParen, value: "("}
		case ')':
			return Token{tokenType: rightParen, value: ")"}
		default:
			if isNumeric(l.program[l.pos]) {
				return l.getNumber()
			}
	}
	panic("token not found")
}

func (l *Lexer) printTokens() {
	for _, t := range l.tokens {
		fmt.Println(t)
	}
}
