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

func (PackageDecl) NodeKind() ast.NodeKind { return KindPackageDecl }

/*
ImportBlock groups import paths.

[Context]
Renders as a parenthesized import block in Go.
*/
type ImportBlock struct {
	Paths []string
}

func (ImportBlock) NodeKind() ast.NodeKind { return KindImportBlock }

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

func (TypeStructDecl) NodeKind() ast.NodeKind { return KindTypeStructDecl }

/*
TypeDefinedDecl declares a defined type or type alias in Go.

[Context]
When IsAlias is false, renders as type Name Underlying. When IsAlias is true, renders as type Name = Underlying.
*/
type TypeDefinedDecl struct {
	Leading    []ast.Node
	Doc        string
	Name       string
	Underlying TypeExpr
	IsAlias    bool
}

func (TypeDefinedDecl) NodeKind() ast.NodeKind { return KindTypeDefinedDecl }

/*
ConstSpec is one entry in a parenthesized const block.

[Context]
Value is emitted verbatim as the right-hand literal from the registry (for example 0, -1, 0x7FFFFFFF).
*/
type ConstSpec struct {
	Doc   string
	Name  string
	Type  *TypeExpr
	Value string
}

/*
ConstGroupDecl declares a parenthesized const block.

[Context]
Renders as const ( ... ) in Go. Each spec may include an optional typed name and doc as line comments.
*/
type ConstGroupDecl struct {
	Leading []ast.Node
	Doc     string
	Specs   []ConstSpec
}

func (ConstGroupDecl) NodeKind() ast.NodeKind { return KindConstGroupDecl }

/*
FuncDecl declares a function with signature and body.

[Context]
Renders as func Name(params) result { body } in Go.
*/
type FuncDecl struct {
	Name    string
	Params  []ParamType
	Returns []TypeExpr
	Body    BlockStmt
}

func (FuncDecl) NodeKind() ast.NodeKind { return KindFuncDecl }

func DeclPackage(name string) PackageDecl {
	return PackageDecl{Name: name}
}

func DeclImportBlock(paths ...string) ImportBlock {
	return ImportBlock{Paths: paths}
}

func DeclTypeStruct(name string, fields []StructFieldDecl) TypeStructDecl {
	return TypeStructDecl{Name: name, Fields: fields}
}

func DeclFunc(name string, params []ParamType, returns []TypeExpr, body BlockStmt) FuncDecl {
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

func DeclTypeDefined(name string, underlying TypeExpr, isAlias bool, doc string, leading ...ast.Node) TypeDefinedDecl {
	return TypeDefinedDecl{Name: name, Underlying: underlying, IsAlias: isAlias, Doc: doc, Leading: leading}
}

func DeclConstGroup(specs []ConstSpec, doc string, leading ...ast.Node) ConstGroupDecl {
	return ConstGroupDecl{Specs: specs, Doc: doc, Leading: leading}
}

func ConstSpecNew(name string, typ *TypeExpr, value string, doc string) ConstSpec {
	return ConstSpec{Name: name, Type: typ, Value: value, Doc: doc}
}
