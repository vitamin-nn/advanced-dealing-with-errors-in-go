package astpain

import (
	"go/ast"
)

// GetDeferredFunctionName возвращает имя функции, вызов которой был отложен через defer,
// если входящий node является *ast.DeferStmt.
func GetDeferredFunctionName(node ast.Node) string {
	defStmt, ok := node.(*ast.DeferStmt)
	if !ok {
		return ""
	}

	switch defStmt.Call.Fun.(type) {
	case *ast.FuncLit:
		return "anonymous func"
	case *ast.Ident:
		return defStmt.Call.Fun.(*ast.Ident).Name
	case *ast.SelectorExpr:
		return getFuncNameRecursive(defStmt.Call.Fun.(*ast.SelectorExpr))
	}

	return ""
}

func getFuncNameRecursive(node *ast.SelectorExpr) string {
	var s string

	switch node.X.(type) {
	case *ast.SelectorExpr:
		s = getFuncNameRecursive(node.X.(*ast.SelectorExpr))
	case *ast.Ident:
		s = node.X.(*ast.Ident).Name
	}

	return s + "." + node.Sel.Name
}
