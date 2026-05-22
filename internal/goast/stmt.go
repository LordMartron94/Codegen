package goast

import "codegen/internal/ast"

func (BlockStmt) IsStmt()        {}
func (ShortVarDeclStmt) IsStmt() {}
func (AssignStmt) IsStmt()       {}
func (RangeLoopStmt) IsStmt()    {}
func (IfStmt) IsStmt()           {}
func (ReturnStmt) IsStmt()       {}
func (ExprStmt) IsStmt()         {}
func (VarDeclStmt) IsStmt()      {}

/*
BlockStmt groups statements with optional leading layout nodes.
*/
type BlockStmt struct {
	Leading []ast.Node
	Stmts   []ast.Stmt
}

func (BlockStmt) NodeKind() ast.NodeKind { return KindBlockStmt }

/*
VarDeclStmt declares a variable with var syntax.
*/
type VarDeclStmt struct {
	Name       string
	Type       TypeExpr
	IsSliceVar bool
}

func (VarDeclStmt) NodeKind() ast.NodeKind { return KindVarDeclStmt }

/*
ShortVarDeclStmt declares and initializes with := syntax.
*/
type ShortVarDeclStmt struct {
	Name string
	Rhs  ast.Expr
}

func (ShortVarDeclStmt) NodeKind() ast.NodeKind { return KindShortVarDeclStmt }

/*
AssignStmt assigns an expression to a name with = syntax.
*/
type AssignStmt struct {
	Name string
	Rhs  ast.Expr
}

func (AssignStmt) NodeKind() ast.NodeKind { return KindAssignStmt }

/*
RangeLoopStmt iterates over a slice with for i := range collection.
*/
type RangeLoopStmt struct {
	Collection  string
	ElementName string
	Body        BlockStmt
}

func (RangeLoopStmt) NodeKind() ast.NodeKind { return KindRangeLoopStmt }

/*
IfStmt conditionally executes a body.
*/
type IfStmt struct {
	Init      ast.Stmt
	Condition ast.Expr
	Body      BlockStmt
}

func (IfStmt) NodeKind() ast.NodeKind { return KindIfStmt }

/*
ReturnStmt returns a value from the enclosing function.
*/
type ReturnStmt struct {
	Value ast.Expr
}

func (ReturnStmt) NodeKind() ast.NodeKind { return KindReturnStmt }

/*
ExprStmt executes an expression for side effects.
*/
type ExprStmt struct {
	Expr ast.Expr
}

func (ExprStmt) NodeKind() ast.NodeKind { return KindExprStmt }

func StmtBlock(stmts ...ast.Stmt) BlockStmt {
	return BlockStmt{Stmts: stmts}
}

func StmtBlockWithLeading(leading []ast.Node, stmts ...ast.Stmt) BlockStmt {
	return BlockStmt{Leading: leading, Stmts: stmts}
}

func StmtVarDecl(name string, typ TypeExpr, isSlice bool) VarDeclStmt {
	return VarDeclStmt{Name: name, Type: typ, IsSliceVar: isSlice}
}

func StmtShortVarDecl(name string, rhs ast.Expr) ShortVarDeclStmt {
	return ShortVarDeclStmt{Name: name, Rhs: rhs}
}

func StmtRangeLoop(collection string, elementName string, body BlockStmt) RangeLoopStmt {
	return RangeLoopStmt{Collection: collection, ElementName: elementName, Body: body}
}

func StmtIf(condition ast.Expr, body BlockStmt) IfStmt {
	return IfStmt{Condition: condition, Body: body}
}

func StmtIfWithInit(init ast.Stmt, condition ast.Expr, body BlockStmt) IfStmt {
	return IfStmt{Init: init, Condition: condition, Body: body}
}

func StmtReturn(value ast.Expr) ReturnStmt {
	return ReturnStmt{Value: value}
}

func StmtExpr(expr ast.Expr) ExprStmt {
	return ExprStmt{Expr: expr}
}
