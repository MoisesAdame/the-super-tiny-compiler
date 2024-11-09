package compiler

import (
	"log"
)

type Visitor map[string]func(n *ASTNode, p ASTNode)

func Traverser(ast AST, visitor Visitor) {
	NodeTraverser(ASTNode(ast), ASTNode{}, visitor)
}

func NodeTraverser(childNode ASTNode, parentNode ASTNode, visitor Visitor) {
	// Try to use the predifined methods in the Visitor's map
	for key, method := range visitor {
		if key == childNode.kind {
			method(&childNode, parentNode)
		}
	}

	// Considering the node's type, execute a given function.
	switch childNode.kind {
	case "Program":
		ArrayTraverser(childNode.body, childNode, visitor)

	case "CallExpression":
		ArrayTraverser(childNode.params, childNode, visitor)

	case "NumberLiteral":
		break

	default:
		log.Fatal(childNode.kind)
	}
}

func ArrayTraverser(array []ASTNode, parentNode ASTNode, visitor Visitor) {
	for _, childNode := range array {
		NodeTraverser(childNode, parentNode, visitor)
	}
}
