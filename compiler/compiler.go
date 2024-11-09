package compiler

func Compiler(code string) string {
	// First, we tokenize the raw code into a tokens array
	tokens := Tokenizer(code)

	// Then, we create the Abstract Syntax Tree using the parser
	ast := Parser(tokens)

	// Then, we transform the original AST into the new language one
	newAST := Transformer(ast)

	// Finally we use the code generator to traverse the new AST, returning
	// the result code
	return CodeGenerator(ASTNode(newAST))
}
