package ast

import "Interpreter-made-in-Go-language/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// ノードが関連づけられているトークンのリテラル値を返す
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral() // [0]はlet
	} else {
		return ""
	}
}

// let <identifier> = <expression>
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // 識別子(変数名)
	Value Expression // 値を生成する式
}

// Statementインターフェイスを満たす 
// parser.go:63:29: cannot use p.parseLetStatement() (type *ast.LetStatement) as type ast.Statement in return argument: 
// *ast.LetStatement does not implement ast.Statement (missing ast.statementNode method)エラーを避ける
func (ls *LetStatement) statementNode() {}
// Nodeインターフェイスを満たす
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// return <expression>
type ReturnStatement struct {
	Token       token.Token // 'return'トークン
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }