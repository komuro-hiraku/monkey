package evaluator

import (
	"github.com/komuro-hiraku/monkey/ast"
	"github.com/komuro-hiraku/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	//文
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// 式
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	// 真偽値
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	}
	return nil
}

// 再帰的にStatementを評価していく
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}
	return result
}
