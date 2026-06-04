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
