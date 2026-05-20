package ast

func (BlockStmt) IsStmt()        {}
func (ShortVarDeclStmt) IsStmt() {}
func (AssignStmt) IsStmt()       {}
func (RangeLoopStmt) IsStmt()    {}
func (IfStmt) IsStmt()           {}
func (ReturnStmt) IsStmt()       {}
func (ExprStmt) IsStmt()         {}

/*
BlockStmt groups statements with optional leading layout nodes.

[Context]
Body statements execute in order; leading nodes may be layout comments before the first statement.
*/
type BlockStmt struct {
	Leading []Node
	Stmts   []Stmt
}

func (BlockStmt) NodeKind() NodeKind { return NodeKindBlockStmt }

/*
ShortVarDeclStmt declares and initializes with := syntax.

[Context]
Rhs is rendered as an expression; used for addr := call() and command := slice[i].
*/
type ShortVarDeclStmt struct {
	Name string
	Rhs  Expr
}

func (ShortVarDeclStmt) NodeKind() NodeKind { return NodeKindShortVarDeclStmt }

/*
AssignStmt assigns an expression to a name with = syntax.

[Context]
Included for completeness; VulkanGen primarily uses ShortVarDeclStmt.
*/
type AssignStmt struct {
	Name string
	Rhs  Expr
}

func (AssignStmt) NodeKind() NodeKind { return NodeKindAssignStmt }

/*
RangeLoopStmt iterates over a slice with for i := range collection.

[Context]
Binds elementName to collection[i] before executing body statements.
*/
type RangeLoopStmt struct {
	Collection  string
	ElementName string
	Body        BlockStmt
}

func (RangeLoopStmt) NodeKind() NodeKind { return NodeKindRangeLoopStmt }

/*
IfStmt conditionally executes a body.

[Context]
When Init is non-nil, renders Go if init; condition { } form. Condition is always an Expr.
*/
type IfStmt struct {
	Init      Stmt
	Condition Expr
	Body      BlockStmt
}

func (IfStmt) NodeKind() NodeKind { return NodeKindIfStmt }

/*
ReturnStmt returns a value from the enclosing function.

[Context]
Value is an Expr, typically IdentExpr.
*/
type ReturnStmt struct {
	Value Expr
}

func (ReturnStmt) NodeKind() NodeKind { return NodeKindReturnStmt }

/*
ExprStmt executes an expression for side effects.

[Context]
Renders a call or other expr as a statement line in Go.
*/
type ExprStmt struct {
	Expr Expr
}

func (ExprStmt) NodeKind() NodeKind { return NodeKindExprStmt }

func StmtBlock(stmts ...Stmt) BlockStmt {
	return BlockStmt{Stmts: stmts}
}

func StmtBlockWithLeading(leading []Node, stmts ...Stmt) BlockStmt {
	return BlockStmt{Leading: leading, Stmts: stmts}
}

func StmtShortVarDecl(name string, rhs Expr) ShortVarDeclStmt {
	return ShortVarDeclStmt{Name: name, Rhs: rhs}
}

func StmtRangeLoop(collection string, elementName string, body BlockStmt) RangeLoopStmt {
	return RangeLoopStmt{Collection: collection, ElementName: elementName, Body: body}
}

func StmtIf(condition Expr, body BlockStmt) IfStmt {
	return IfStmt{Condition: condition, Body: body}
}

func StmtIfWithInit(init Stmt, condition Expr, body BlockStmt) IfStmt {
	return IfStmt{Init: init, Condition: condition, Body: body}
}

func StmtReturn(value Expr) ReturnStmt {
	return ReturnStmt{Value: value}
}

func StmtExpr(expr Expr) ExprStmt {
	return ExprStmt{Expr: expr}
}
