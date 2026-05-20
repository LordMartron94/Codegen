package goast

import "codegen/internal/ast"

func (VarDeclStmt) IsStmt() {}

/*
VarDeclStmt declares a variable with var syntax.

[Context]
Renders as var name Type or var name []ElemType in Go.
*/
type VarDeclStmt struct {
	Name       string
	Type       TypeExpr
	IsSliceVar bool
}

func (VarDeclStmt) NodeKind() ast.NodeKind { return ast.NodeKindVarDeclStmt }

func StmtVarDecl(name string, typ TypeExpr, isSlice bool) VarDeclStmt {
	return VarDeclStmt{Name: name, Type: typ, IsSliceVar: isSlice}
}
