package gocode

import (
	"codegen"
	"codegen/internal/ast"
	"codegen/internal/goast"
)

/*
Kind identifies a Go AST node variant. See internal/goast for all kind constants.
*/
type Kind = goast.Kind

/*
CommentStyle selects how documentation text is rendered in Go output.
*/
type CommentStyle = goast.CommentStyle

/*
RenderOptions configures Go backend rendering behavior.
*/
type RenderOptions = goast.RenderOptions

const (
	CommentStyleLine  = goast.CommentStyleLine
	CommentStyleBlock = goast.CommentStyleBlock
)

/*
File is the root node for a generated Go source file.
*/
type File = goast.File

/*
GoDocComment is a Go documentation block attached to a declaration or field.
*/
type GoDocComment = goast.GoDocComment

/*
BlankLine inserts vertical spacing in generated Go output.
*/
type BlankLine = goast.BlankLine

/*
LineComment is a single-line // comment.
*/
type LineComment = goast.LineComment

/*
BlockComment is a multi-line // comment block.
*/
type BlockComment = goast.BlockComment

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
TypeDefinedDecl declares a Go defined type or type alias.
*/
type TypeDefinedDecl = goast.TypeDefinedDecl

/*
ConstSpec is one entry in a Go const block.
*/
type ConstSpec = goast.ConstSpec

/*
ConstGroupDecl declares a parenthesized Go const block.
*/
type ConstGroupDecl = goast.ConstGroupDecl

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
BlockStmt groups statements with optional leading layout nodes.
*/
type BlockStmt = goast.BlockStmt

/*
VarDeclStmt declares a variable with Go var syntax.
*/
type VarDeclStmt = goast.VarDeclStmt

/*
ShortVarDeclStmt declares and initializes a variable with Go := syntax.
*/
type ShortVarDeclStmt = goast.ShortVarDeclStmt

/*
AssignStmt assigns an expression to a name with = syntax.
*/
type AssignStmt = goast.AssignStmt

/*
RangeLoopStmt iterates a Go slice with for i := range collection.
*/
type RangeLoopStmt = goast.RangeLoopStmt

/*
IfStmt conditionally executes a body, optionally with an init statement.
*/
type IfStmt = goast.IfStmt

/*
ReturnStmt returns a value from the enclosing function.
*/
type ReturnStmt = goast.ReturnStmt

/*
ExprStmt executes an expression for side effects.
*/
type ExprStmt = goast.ExprStmt

/*
IdentExpr references an identifier by name.
*/
type IdentExpr = goast.IdentExpr

/*
SelectorExpr selects a field or method on a base expression.
*/
type SelectorExpr = goast.SelectorExpr

/*
IndexExpr indexes a base expression.
*/
type IndexExpr = goast.IndexExpr

/*
AddressOfExpr takes the address of its operand.
*/
type AddressOfExpr = goast.AddressOfExpr

/*
CallExpr invokes a callee with arguments.
*/
type CallExpr = goast.CallExpr

/*
StringLitExpr is a string literal.
*/
type StringLitExpr = goast.StringLitExpr

/*
IntLitExpr is an integer literal.
*/
type IntLitExpr = goast.IntLitExpr

/*
NilLitExpr is the Go nil literal.
*/
type NilLitExpr = goast.NilLitExpr

/*
BinaryExpr combines two expressions with a binary operator.
*/
type BinaryExpr = goast.BinaryExpr

/*
BinaryOp identifies a binary operator.
*/
type BinaryOp = goast.BinaryOp

/*
FieldInit pairs a field name with an initializing expression in a Go composite literal.
*/
type FieldInit = goast.FieldInit

/*
CompositeLitExpr is a Go struct or slice composite literal.
*/
type CompositeLitExpr = goast.CompositeLitExpr

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
	KindBlankLine        = goast.KindBlankLine
	KindLineComment      = goast.KindLineComment
	KindBlockComment     = goast.KindBlockComment
	KindGoDocComment     = goast.KindGoDocComment
	KindFile             = goast.KindFile
	KindPackageDecl      = goast.KindPackageDecl
	KindImportBlock      = goast.KindImportBlock
	KindTypeStructDecl   = goast.KindTypeStructDecl
	KindTypeDefinedDecl  = goast.KindTypeDefinedDecl
	KindConstGroupDecl   = goast.KindConstGroupDecl
	KindFuncDecl         = goast.KindFuncDecl
	KindBlockStmt        = goast.KindBlockStmt
	KindVarDeclStmt      = goast.KindVarDeclStmt
	KindShortVarDeclStmt = goast.KindShortVarDeclStmt
	KindAssignStmt       = goast.KindAssignStmt
	KindRangeLoopStmt    = goast.KindRangeLoopStmt
	KindIfStmt           = goast.KindIfStmt
	KindReturnStmt       = goast.KindReturnStmt
	KindExprStmt         = goast.KindExprStmt
	KindIdentExpr        = goast.KindIdentExpr
	KindSelectorExpr     = goast.KindSelectorExpr
	KindIndexExpr        = goast.KindIndexExpr
	KindAddressOfExpr    = goast.KindAddressOfExpr
	KindCallExpr         = goast.KindCallExpr
	KindStringLitExpr    = goast.KindStringLitExpr
	KindIntLitExpr       = goast.KindIntLitExpr
	KindNilLitExpr       = goast.KindNilLitExpr
	KindBinaryExpr       = goast.KindBinaryExpr
	KindCompositeLitExpr = goast.KindCompositeLitExpr
	BinaryOpEq           = goast.BinaryOpEq
	BinaryOpNe           = goast.BinaryOpNe
	TypeExprKindNamed    = goast.TypeExprKindNamed
	TypeExprKindPointer  = goast.TypeExprKindPointer
	TypeExprKindSlice    = goast.TypeExprKindSlice
	TypeExprKindFunc     = goast.TypeExprKindFunc
)

func FileElementFrom(node codegen.Node) codegen.FileElement {
	return codegen.FileElementFrom(node)
}

func DeclFile(elements ...codegen.FileElement) File {
	return goast.DeclFile(elements...)
}

func TypeExprFromGoTypeString(typeString string) (TypeExpr, error) {
	return goast.TypeExprFromGoTypeString(typeString)
}

func DeclPackage(name string) PackageDecl {
	return goast.DeclPackage(name)
}

func DeclImportBlock(paths ...string) ImportBlock {
	return goast.DeclImportBlock(paths...)
}

func DeclTypeStruct(name string, fields []StructFieldDecl) TypeStructDecl {
	return goast.DeclTypeStruct(name, fields)
}

func DeclTypeDefined(name string, underlying TypeExpr, isAlias bool, doc string, leading ...codegen.Node) TypeDefinedDecl {
	return goast.DeclTypeDefined(name, underlying, isAlias, doc, leading...)
}

func RenderOptionsDefault() RenderOptions {
	return goast.RenderOptionsDefault()
}

func RenderOptionsEnumBindings() RenderOptions {
	return goast.RenderOptionsEnumBindings()
}

/*
GoDocFormatExported formats documentation so it begins with the exported identifier name.
*/
func GoDocFormatExported(subject string, doc string) string {
	return goast.GoDocFormatExported(subject, doc)
}

func DeclConstGroup(specs []ConstSpec, doc string, separateDocumentedSpecs bool, leading ...codegen.Node) ConstGroupDecl {
	return goast.DeclConstGroup(specs, doc, separateDocumentedSpecs, leading...)
}

func ConstSpecNew(name string, typ *TypeExpr, value string, doc string) ConstSpec {
	return goast.ConstSpecNew(name, typ, value, doc)
}

func TypeExprNamedPtr(name string) *TypeExpr {
	typ := TypeExprNamed(name)
	return &typ
}

func StructFieldType(name string, typ TypeExpr) StructFieldDecl {
	return goast.StructFieldType(name, typ)
}

func StructFieldFunc(name string, sig FuncTypeSig) StructFieldDecl {
	return goast.StructFieldFunc(name, sig)
}

func StructFieldFuncDoc(name string, sig FuncTypeSig, doc string, leading ...codegen.Node) StructFieldDecl {
	return goast.StructFieldFuncDoc(name, sig, doc, leading...)
}

func TypeExprNamed(name string) TypeExpr {
	return goast.TypeExprNamed(name)
}

func TypeExprPointer(elem TypeExpr) TypeExpr {
	return goast.TypeExprPointer(elem)
}

func DeclFunc(name string, params []ParamType, returns []TypeExpr, body BlockStmt) FuncDecl {
	return goast.DeclFunc(name, params, returns, body)
}

func ExprIdent(name string) *IdentExpr {
	return goast.ExprIdent(name)
}

func ExprSelector(base ast.Expr, name string) *SelectorExpr {
	return goast.ExprSelector(base, name)
}

func ExprIndex(base ast.Expr, index ast.Expr) *IndexExpr {
	return goast.ExprIndex(base, index)
}

func ExprAddressOf(operand ast.Expr) *AddressOfExpr {
	return goast.ExprAddressOf(operand)
}

func ExprCall(callee ast.Expr, args ...ast.Expr) *CallExpr {
	return goast.ExprCall(callee, args...)
}

func ExprStringLit(value string) *StringLitExpr {
	return goast.ExprStringLit(value)
}

func ExprIntLit(value int64) *IntLitExpr {
	return goast.ExprIntLit(value)
}

func ExprNil() *NilLitExpr {
	return goast.ExprNil()
}

func ExprBinary(op BinaryOp, left ast.Expr, right ast.Expr) *BinaryExpr {
	return goast.ExprBinary(op, left, right)
}

func ExprCompositeLit(typeName string, fields []FieldInit) *CompositeLitExpr {
	return goast.ExprCompositeLit(typeName, fields)
}

func ExprSliceCompositeLit(elementTypeName string, elements []ast.Expr) *CompositeLitExpr {
	return goast.ExprSliceCompositeLit(elementTypeName, elements)
}

func StmtBlock(stmts ...ast.Stmt) BlockStmt {
	return goast.StmtBlock(stmts...)
}

func StmtVarDecl(name string, typ TypeExpr, isSlice bool) VarDeclStmt {
	return goast.StmtVarDecl(name, typ, isSlice)
}

func StmtShortVarDecl(name string, rhs ast.Expr) ShortVarDeclStmt {
	return goast.StmtShortVarDecl(name, rhs)
}

func StmtRangeLoop(collection string, elementName string, body BlockStmt) RangeLoopStmt {
	return goast.StmtRangeLoop(collection, elementName, body)
}

func StmtIf(condition ast.Expr, body BlockStmt) IfStmt {
	return goast.StmtIf(condition, body)
}

func StmtIfWithInit(init ast.Stmt, condition ast.Expr, body BlockStmt) IfStmt {
	return goast.StmtIfWithInit(init, condition, body)
}

func StmtReturn(value ast.Expr) ReturnStmt {
	return goast.StmtReturn(value)
}

func StmtExpr(expr ast.Expr) ExprStmt {
	return goast.StmtExpr(expr)
}
