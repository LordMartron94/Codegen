package goast

import "codegen/internal/ast"

/*
Kind identifies a Go AST node variant for engine registration and dispatch.

[Context]
All Go node kinds are defined in this frontend package. The core codegen engine has no knowledge of these values.
*/
type Kind = ast.NodeKind

const (
	/* KindBlankLine inserts vertical spacing. */
	KindBlankLine Kind = iota
	/* KindLineComment is a single-line // comment. */
	KindLineComment
	/* KindBlockComment is a multi-line // comment block. */
	KindBlockComment
	/* KindGoDocComment is a Go block documentation comment. */
	KindGoDocComment

	/* KindFile is a generated Go source file root. */
	KindFile
	/* KindPackageDecl is a package clause. */
	KindPackageDecl
	/* KindImportBlock is a parenthesized import block. */
	KindImportBlock
	/* KindTypeStructDecl is a struct type declaration. */
	KindTypeStructDecl
	/* KindTypeDefinedDecl is a defined type or type alias declaration. */
	KindTypeDefinedDecl
	/* KindConstGroupDecl is a parenthesized const block. */
	KindConstGroupDecl
	/* KindVarDecl is a package-level var declaration. */
	KindVarDecl
	/* KindFuncDecl is a function declaration. */
	KindFuncDecl

	/* KindBlockStmt is a statement block. */
	KindBlockStmt
	/* KindVarDeclStmt is a var declaration statement. */
	KindVarDeclStmt
	/* KindShortVarDeclStmt is a := declaration statement. */
	KindShortVarDeclStmt
	/* KindAssignStmt is an assignment statement. */
	KindAssignStmt
	/* KindRangeLoopStmt is a for-range loop statement. */
	KindRangeLoopStmt
	/* KindIfStmt is a conditional statement. */
	KindIfStmt
	/* KindReturnStmt is a return statement. */
	KindReturnStmt
	/* KindExprStmt is an expression statement. */
	KindExprStmt

	/* KindIdentExpr is an identifier expression. */
	KindIdentExpr
	/* KindSelectorExpr is a selector expression. */
	KindSelectorExpr
	/* KindIndexExpr is an index expression. */
	KindIndexExpr
	/* KindAddressOfExpr is an address-of expression. */
	KindAddressOfExpr
	/* KindCallExpr is a call expression. */
	KindCallExpr
	/* KindStringLitExpr is a string literal expression. */
	KindStringLitExpr
	/* KindIntLitExpr is an integer literal expression. */
	KindIntLitExpr
	/* KindNilLitExpr is a nil literal expression. */
	KindNilLitExpr
	/* KindBoolLitExpr is a boolean literal expression. */
	KindBoolLitExpr
	/* KindBinaryExpr is a binary expression. */
	KindBinaryExpr
	/* KindCompositeLitExpr is a composite literal expression. */
	KindCompositeLitExpr
)
