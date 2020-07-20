package parser

import (
	"fmt"
	"Interpreter-made-in-Go-language/ast"
	"Interpreter-made-in-Go-language/lexer"
	"Interpreter-made-in-Go-language/token"
)

type Parser struct {
	l *lexer.Lexer
    errors []string
	curToken  token.Token // 現在のトークン
	peekToken token.Token // 次のトークン
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// 2つのトークンを読み込む。 curTokenとpeekTokenの両方がセットされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// ASTのルートノードを生成
	program := &ast.Program{} // &{[]}
	program.Statements = []ast.Statement{} // []

	for p.curToken.Type != token.EOF { // !p.curTokenIs(token.EOF)で代替できる
		stmt := p.parseStatement() // &{{LET let} 0xc00010e390 <nil>}
		if stmt != nil {
			program.Statements = append(program.Statements, stmt) // let から ; までの1文の解析結果を格納していく
		}
		p.nextToken() // 次のLETから始まるトークンを読んでいく
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// let から ; までを解析する
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// token.Letトークンに基づいて*ast.LetStatementノードを構築
	stmt := &ast.LetStatement{Token: p.curToken} // &{{LET let} <nil> <nil>}

	// token.IDENTトークン(add, foobar, x, y, ...)を期待
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal} // &{{IDENT x} x}

	// 等号を期待
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっているので式はみていない
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken} // &{{RETURN return} <nil>}

	p.nextToken()

	// TODO: セミコロンに遭遇するまで式を読み飛ばしてしまっているので式はみていない
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken() // 予想通りだったら次のトークンを読んでいく
		return true
	} else {
		p.peekError(t)
		return false
	}
}