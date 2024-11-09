package compiler

import (
	"log"
)

// - ASTNode: Node for the Abstract Syntax Tree
type ASTNode struct {
	kind      string
	value     string
	name      string
	callee    *ASTNode
	expresion *ASTNode
	body      []ASTNode
	params    []ASTNode
	arguments *[]ASTNode
	context   *[]ASTNode
}

func (node *ASTNode) ToStringHelper(margin int) string {
	marginStr := "\n"
	for i := 0; i <= margin; i++ {
		marginStr += "\t"
	}

	if node.kind == "CallExpression" {
		res := marginStr + "\tkind: CallExpression," + marginStr + "\tname: " + node.name + "," + marginStr + "\tparams: {"
		for _, val := range node.params {
			res += val.ToStringHelper(margin + 1)
		}

		return res + marginStr + "}"
	}

	if node.kind == "NumberLiteral" {
		res := marginStr + "\tkind: NumberLiteral," + marginStr + "\tvalue: " + node.value
		return res + marginStr + "\t}"
	}

	return ""
}

// - AST: For simplicity we alias ASTNode as AST
type AST ASTNode

func (ast *AST) ToString() string {
	res := "{\n\tkind: Program,\n\tbody: {"

	for _, val := range ast.body {
		res += val.ToStringHelper(0)
	}

	return res + "\n}"
}

// - Parser: Iterate over array of Tokens and generate the Abstract Syntax Tree
func Parser(tokens []Token) AST {
	// Instantiate the AST
	ast := AST{
		kind: "Program",
		body: []ASTNode{},
	}

	// Iterating over the tokens and appending them, if necessary, to the AST
	for parserCounter := 0; parserCounter < len(tokens); {
		ast.body = append(ast.body, ParserWalker(tokens, &parserCounter))
	}

	return ast
}

func ParserWalker(tokens []Token, counter *int) ASTNode {
	// Store number tokens as NumberLiteral in the AST
	if tokens[*counter].kind == "number" {
		newNode := ASTNode{
			kind:  "NumberLiteral",
			value: tokens[*counter].value,
		}

		(*counter)++

		return newNode
	}

	// If there is an openning parenthesis, search for names and their expressions
	if tokens[*counter].kind == "paren" && tokens[*counter].value == "(" {
		(*counter)++

		newNode := ASTNode{
			kind:   "CallExpression",
			name:   tokens[*counter].value,
			params: []ASTNode{},
		}

		// Add parameters until finding a closing parenthesis
		(*counter)++
		for tokens[*counter].kind != "paren" || (tokens[*counter].kind == "paren" && tokens[*counter].value != ")") {
			newNode.params = append(newNode.params, ParserWalker(tokens, counter))
		}

		(*counter)++
		return newNode
	}

	// Again, if we haven't recognized the token type by now we're going to
	// throw an error.
	log.Fatal(tokens[*counter].kind)
	return ASTNode{}
}
