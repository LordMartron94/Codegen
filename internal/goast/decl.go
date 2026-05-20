package goast

import "codegen/internal/ast"

/*
PackageDecl declares the package clause.

[Context]
Renders as package <name> in Go.
*/
type PackageDecl struct {
	Name string
}

func (PackageDecl) NodeKind() ast.NodeKind { return ast.NodeKindPackageDecl }

/*
ImportBlock groups import paths.

[Context]
Renders as a parenthesized import block in Go.
*/
type ImportBlock struct {
	Paths []string
}

func (ImportBlock) NodeKind() ast.NodeKind { return ast.NodeKindImportBlock }

/*
StructFieldDecl describes one field in a struct type.

[Context]
Either Type holds a plain field type or FuncType holds an embedded function field signature.
*/
type StructFieldDecl struct {
	Leading  []ast.Node
	Doc      string
	Name     string
	Type     TypeExpr
	FuncType *FuncTypeSig
}

/*
FuncTypeSig is a function type used as a struct field type.

[Context]
Renders as func(params) returns in Go struct fields.
*/
type FuncTypeSig struct {
	Params  []ParamType
	Returns []TypeExpr
}

/*
TypeStructDecl declares a struct type.

[Context]
Renders as type Name struct { ... } in Go.
*/
type TypeStructDecl struct {
	Name   string
	Fields []StructFieldDecl
}

func (TypeStructDecl) NodeKind() ast.NodeKind { return ast.NodeKindTypeStructDecl }

/*
FuncDecl declares a function with signature and body.

[Context]
Renders as func Name(params) result { body } in Go.
*/
type FuncDecl struct {
	Name    string
	Params  []ParamType
	Returns []TypeExpr
	Body    ast.BlockStmt
}

func (FuncDecl) NodeKind() ast.NodeKind { return ast.NodeKindFuncDecl }

func DeclPackage(name string) PackageDecl {
	return PackageDecl{Name: name}
}

func DeclImportBlock(paths ...string) ImportBlock {
	return ImportBlock{Paths: paths}
}

func DeclTypeStruct(name string, fields []StructFieldDecl) TypeStructDecl {
	return TypeStructDecl{Name: name, Fields: fields}
}

func DeclFunc(name string, params []ParamType, returns []TypeExpr, body ast.BlockStmt) FuncDecl {
	return FuncDecl{Name: name, Params: params, Returns: returns, Body: body}
}

func StructFieldType(name string, typ TypeExpr) StructFieldDecl {
	return StructFieldDecl{Name: name, Type: typ}
}

func StructFieldFunc(name string, sig FuncTypeSig) StructFieldDecl {
	return StructFieldDecl{Name: name, FuncType: &sig}
}

func StructFieldFuncDoc(name string, sig FuncTypeSig, doc string, leading ...ast.Node) StructFieldDecl {
	return StructFieldDecl{Name: name, FuncType: &sig, Doc: doc, Leading: leading}
}

func StructFieldTypeDoc(name string, typ TypeExpr, doc string, leading ...ast.Node) StructFieldDecl {
	return StructFieldDecl{Name: name, Type: typ, Doc: doc, Leading: leading}
}
