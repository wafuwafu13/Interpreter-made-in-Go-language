package parser

import (
//	"fmt"
	"testing"
	"Interpreter-made-in-Go-language/ast"
	"Interpreter-made-in-Go-language/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5 + 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input) // &{
		                  //   let x = 5;
		                  //   let y = 10;
		                  //   let foobar = 838383;
						  //    0 1 10
						  // }
	
	p := New(l) // &{0xc0000103f0 [] {LET let} {IDENT x}}

	program := p.ParseProgram() // inputの解析結果

	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	// let から ; までの1文の解析結果が3つ格納されている
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}	

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.Letstatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name,TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestRerutnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l) // &{0xc0000103f0 [] {RETURN return} {INT 5}}

	program := p.ParseProgram()
	checkParserErrors(t, p)

	// return文の解析結果が3つ格納されている
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}	

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
	
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}