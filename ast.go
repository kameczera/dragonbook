package main

type Node interface {  }

type varType int

const (
	Integer varType = iota
)

type Number struct {
	value int
}

type Variable struct {
	identifier string
}

type Binary struct {
	op Token
	left Node
	right Node
}

type Assign struct {
	variable Variable
	value Node
}

type Print struct {
	value Node
}

type AST struct {
	root Node
}
