package repl

import (
	"bufio"
	"flare/lexer"
	"flare/parser"
	"fmt"
	"io"
)

const PROOMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROOMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		if line == ".bye" {
			io.WriteString(out, "Bye!!\n")
			return
		}

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "lmao desu! we wanu into some syntaxu meowe like skillu issue hewe x3")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
