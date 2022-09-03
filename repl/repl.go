package repl

import (
	"Interpreter-made-in-Go-language/evaluator"
	"Interpreter-made-in-Go-language/lexer"
	"Interpreter-made-in-Go-language/object"
	"Interpreter-made-in-Go-language/parser"
	"Interpreter-made-in-Go-language/token"
	"fmt"
	"os"
)

const PROMPT = ">> "

func Start() {
	// scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	b, err := os.ReadFile("./MONKEY")
	if err != nil {
		fmt.Print(err)
	}

	content := string(b)
	fmt.Print(content, "\n")

	l := lexer.New(content)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(content, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		fmt.Print(evaluated.Inspect(), "\n")
	}

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out string, errors []string) {
	fmt.Print(MONKEY_FACE)
	fmt.Print("Woops! We ran into some monkey business here!\n")
	fmt.Print("parser errors:\n")
	for _, msg := range errors {
		fmt.Print("\t" + msg + "\n")
	}
}
