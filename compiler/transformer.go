package compiler

func Transformer(ast AST) AST {
	newAST := AST{
		kind: "Program",
		body: []ASTNode{},
	}

	// In order to iterate over the old and the new AST, we assign the new one
	// into the old one.
	ast.context = &newAST.body

	// Define the functions of the visitor and traverse over ast.
	Traverser(ast, map[string]func(n *ASTNode, p ASTNode){
		"NumberLiteral": func(n *ASTNode, p ASTNode) {
			*p.context = append(*p.context, ASTNode{
				kind:  "NumberLiteral",
				value: n.value,
			})
		},
		"CallExpression": func(n *ASTNode, p ASTNode) {
			expression := ASTNode{
				kind: "CallExpression",
				callee: &ASTNode{
					kind: "Identifier",
					name: n.name,
				},
				arguments: &[]ASTNode{},
			}

			n.context = expression.arguments

			if p.kind != "CallExpression" {
				expressionStatement := ASTNode{
					kind:      "ExpressionStatement",
					expresion: &expression,
				}

				*p.context = append(*p.context, expressionStatement)
			} else {
				*p.context = append(*p.context, expression)
			}
		},
	})

	return newAST
}
