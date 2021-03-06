/*
Object Structural: Facade

“Provide a unified interface to a set of interfaces in a subsystem. Facade
defines a higher-level interface that makes the subsystem easier to use”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we provide the compiler a unified interface 'compile' that
acts as a 'facade' hiding the internal compiler subsystems.
*/

package main

import(
  "fmt"
)

type Token struct{}

type AbstractSyntaxTree struct{}

type ByteCode struct{}

type Scanner struct{}

type Parser struct{}

type Compiler struct{}

type ByteCodeGenerator struct{}

func (s Scanner) scan(source string) []Token {
  return []Token{}
}

func (p Parser) parse(tokens []Token) AbstractSyntaxTree {
  return AbstractSyntaxTree{}
}

func (b ByteCodeGenerator) generate(tree AbstractSyntaxTree) []ByteCode {
  return []ByteCode{}
}

func (c Compiler) compile(source string) []ByteCode {
  tokens := Scanner{}.scan(source)
  ast := Parser{}.parse(tokens)
  byteCode := ByteCodeGenerator{}.generate(ast)
  return byteCode
}

func main() {
  compiler := Compiler{}
  byteCode := compiler.compile("func main() {}")
  fmt.Println(byteCode)
}
