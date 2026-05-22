package gocode

import (
	"codegen"
	"codegen/internal/ast"
	"codegen/internal/goast"
)

/*
GoDocComment is a Go documentation block attached to a declaration or field.
*/
type GoDocComment = goast.GoDocComment

/*
PackageDecl declares the package clause of a generated Go file.
*/
type PackageDecl = goast.PackageDecl

/*
ImportBlock declares a parenthesized Go import block.
*/
type ImportBlock = goast.ImportBlock

/*
TypeStructDecl declares a Go struct type with fields.
*/
type TypeStructDecl = goast.TypeStructDecl

/*
StructFieldDecl describes one field in a Go struct.
*/
type StructFieldDecl = goast.StructFieldDecl

/*
FuncTypeSig is a Go function type used as a struct field type.
*/
type FuncTypeSig = goast.FuncTypeSig

/*
FuncDecl declares a Go function with signature and body block.
*/
type FuncDecl = goast.FuncDecl

/*
VarDeclStmt declares a variable with Go var syntax.
*/
type VarDeclStmt = goast.VarDeclStmt

/*
ShortVarDeclStmt declares and initializes a variable with Go := syntax.
*/
type ShortVarDeclStmt = ast.ShortVarDeclStmt

/*
RangeLoopStmt iterates a Go slice with for i := range collection.
*/
type RangeLoopStmt = ast.RangeLoopStmt

/*
NilLitExpr is the Go nil literal.
*/
type NilLitExpr = ast.NilLitExpr

/*
FieldInit pairs a field name with an initializing expression in a Go composite literal.
*/
type FieldInit = ast.FieldInit

/*
CompositeLitExpr is a Go struct or slice composite literal.
*/
type CompositeLitExpr = ast.CompositeLitExpr

/*
TypeExpr describes a Go type in generated code.
*/
type TypeExpr = goast.TypeExpr

/*
TypeExprKind identifies a Go type-expression variant.
*/
type TypeExprKind = goast.TypeExprKind

/*
ParamType pairs a parameter name with its Go type.
*/
type ParamType = goast.ParamType

const (
	/* NodeKindGoDocComment identifies a Go block documentation comment node. */
	NodeKindGoDocComment = ast.NodeKindGoDocComment
	/* NodeKindPackageDecl identifies a Go package declaration. */
	NodeKindPackageDecl = ast.NodeKindPackageDecl
	/* NodeKindImportBlock identifies a Go import block. */
	NodeKindImportBlock = ast.NodeKindImportBlock
	/* NodeKindTypeStructDecl identifies a Go struct type declaration. */
	NodeKindTypeStructDecl = ast.NodeKindTypeStructDecl
	/* NodeKindFuncDecl identifies a Go function declaration. */
	NodeKindFuncDecl = ast.NodeKindFuncDecl
	/* NodeKindVarDeclStmt identifies a Go var declaration statement. */
	NodeKindVarDeclStmt = ast.NodeKindVarDeclStmt
	/* NodeKindShortVarDeclStmt identifies a Go := declaration statement. */
	NodeKindShortVarDeclStmt = ast.NodeKindShortVarDeclStmt
	/* NodeKindRangeLoopStmt identifies a Go for-range loop statement. */
	NodeKindRangeLoopStmt = ast.NodeKindRangeLoopStmt
	/* NodeKindNilLitExpr identifies a Go nil literal expression. */
	NodeKindNilLitExpr = ast.NodeKindNilLitExpr
	/* NodeKindCompositeLitExpr identifies a Go composite literal expression. */
	NodeKindCompositeLitExpr = ast.NodeKindCompositeLitExpr

	/* TypeExprKindNamed identifies a named Go type. */
	TypeExprKindNamed = goast.TypeExprKindNamed
	/* TypeExprKindPointer identifies a Go pointer type. */
	TypeExprKindPointer = goast.TypeExprKindPointer
	/* TypeExprKindSlice identifies a Go slice type. */
	TypeExprKindSlice = goast.TypeExprKindSlice
	/* TypeExprKindFunc identifies a Go function type. */
	TypeExprKindFunc = goast.TypeExprKindFunc
)

/*
FileElementFrom wraps any node as a FileElement.

[Returns]
A FileElement suitable for DeclFile.
*/
func FileElementFrom(node codegen.Node) codegen.FileElement {
	return codegen.FileElementFrom(node)
}

/*
DeclFile constructs a file root node from ordered top-level elements.
*/
func DeclFile(elements ...codegen.FileElement) codegen.File {
	return codegen.DeclFile(elements...)
}

/*
TypeExprFromGoTypeString parses a Go type spelling into a TypeExpr tree.

[Parameters]
typeString — type identifier such as "*uint32" or "[]commandMapping".

[Returns]
The parsed TypeExpr, or an error when typeString is empty.
*/
func TypeExprFromGoTypeString(typeString string) (TypeExpr, error) {
	return goast.TypeExprFromGoTypeString(typeString)
}

/*
DeclPackage constructs a Go package declaration node.
*/
func DeclPackage(name string) PackageDecl {
	return goast.DeclPackage(name)
}

/*
DeclImportBlock constructs a parenthesized Go import block.
*/
func DeclImportBlock(paths ...string) ImportBlock {
	return goast.DeclImportBlock(paths...)
}

/*
DeclTypeStruct constructs a Go struct type declaration with fields.
*/
func DeclTypeStruct(name string, fields []StructFieldDecl) TypeStructDecl {
	return goast.DeclTypeStruct(name, fields)
}

/*
StructFieldType declares a struct field with a plain Go type.
*/
func StructFieldType(name string, typ TypeExpr) StructFieldDecl {
	return goast.StructFieldType(name, typ)
}

/*
StructFieldFunc declares a struct field with an embedded Go function type.
*/
func StructFieldFunc(name string, sig FuncTypeSig) StructFieldDecl {
	return goast.StructFieldFunc(name, sig)
}

/*
StructFieldFuncDoc declares a documented struct field with an embedded Go function type.
*/
func StructFieldFuncDoc(name string, sig FuncTypeSig, doc string, leading ...codegen.Node) StructFieldDecl {
	return goast.StructFieldFuncDoc(name, sig, doc, leading...)
}

/*
TypeExprNamed constructs a named Go type identifier.
*/
func TypeExprNamed(name string) TypeExpr {
	return goast.TypeExprNamed(name)
}

/*
TypeExprPointer constructs a Go pointer type expression.
*/
func TypeExprPointer(elem TypeExpr) TypeExpr {
	return goast.TypeExprPointer(elem)
}

/*
DeclFunc constructs a Go function declaration with body.
*/
func DeclFunc(name string, params []ParamType, returns []TypeExpr, body codegen.BlockStmt) FuncDecl {
	return goast.DeclFunc(name, params, returns, body)
}

/*
ExprNil constructs a Go nil literal expression.
*/
func ExprNil() *NilLitExpr {
	return ast.ExprNil()
}

/*
ExprCompositeLit constructs a Go struct composite literal.
*/
func ExprCompositeLit(typeName string, fields []FieldInit) *CompositeLitExpr {
	return ast.ExprCompositeLit(typeName, fields)
}

/*
ExprSliceCompositeLit constructs a Go slice of composite literals.
*/
func ExprSliceCompositeLit(elementTypeName string, elements []codegen.Expr) *CompositeLitExpr {
	return ast.ExprSliceCompositeLit(elementTypeName, elements)
}

/*
StmtVarDecl constructs a Go var declaration statement.
*/
func StmtVarDecl(name string, typ TypeExpr, isSlice bool) VarDeclStmt {
	return goast.StmtVarDecl(name, typ, isSlice)
}

/*
StmtShortVarDecl constructs a Go short variable declaration with :=.
*/
func StmtShortVarDecl(name string, rhs codegen.Expr) ShortVarDeclStmt {
	return ast.StmtShortVarDecl(name, rhs)
}

/*
StmtRangeLoop constructs a Go for i := range collection loop.
*/
func StmtRangeLoop(collection string, elementName string, body codegen.BlockStmt) RangeLoopStmt {
	return ast.StmtRangeLoop(collection, elementName, body)
}
