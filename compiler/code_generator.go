package compiler

import (
	"log"
	"strings"
)

func CodeGenerator(node ASTNode) string {
	switch node.kind {
	case "Program":
		res := ""
		for _, val := range node.body {
			res += CodeGenerator(val) + "\n"
		}
		return res

	case "ExpressionStatement":
		return CodeGenerator(*node.expresion) + ";"

	case "CallExpression":
		var params []string
		for _, node := range *node.arguments {
			params = append(params, CodeGenerator(node))
		}
		paramsJoined := strings.Join(params, ", ")
		return CodeGenerator(*node.callee) + "(" + paramsJoined + ")"

	case "Identifier":
		return node.name

	case "NumberLiteral":
		return node.value

	default:
		log.Println("Code Generator Error")
		return ""
	}
}
