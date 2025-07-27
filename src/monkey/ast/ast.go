package ast

import (
	"monkey/token"
)

/* Interface Definitions */
// Base node
type Node interface {
	TokenLiteral() string // Where TokenLiteral is the literal value of the token
}

// Statement node (those that do not evaluate)
type Statement interface {
	Node
	statementNode() // dummy methods to guide Go compiler
}

// Expression node (those that evaluate and produce a value)
type Expression interface {
	Node
	expressionNode() // dummy methods to guide Go compiler
}

/* Program root node */
type Program struct {
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
