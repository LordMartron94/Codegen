/*
Package codegen provides a language-agnostic AST and registration-based code generation engines.

[Context]
Build an AST tree, configure an engine with renderers (or use codegen/go for Go), and render to source text through an injected emitter. The core engine uses a dense slice registry indexed by NodeKind for O(1) dispatch. Go-specific shapes (types, declarations) live in internal/goast; neutral layout and expression nodes live in internal/ast.

[Example]

	eng := gocode.GoEngineCreate()
	emitter := gocode.GoEmitterCreate()
	file := DeclFile()
	_ = EngineRender(eng, &emitter, file)
	output := EmitterRender(&emitter)
*/
package codegen

import (
	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

/*
Node is the root interface for every AST value rendered by an engine.

[Context]
All layout, declaration, statement, and expression nodes implement Node and expose a NodeKind discriminator.
*/
type Node = ast.Node

/*
NodeKind identifies an AST variant for engine registration and dispatch.
*/
type NodeKind = ast.NodeKind

/*
Expr is an expression node that produces a value in the target language.
*/
type Expr = ast.Expr

/*
Stmt is a statement node executed inside function bodies and blocks.
*/
type Stmt = ast.Stmt

/*
File is the root node for a generated source file with ordered top-level elements.
*/
type File = ast.File

/*
FileElement wraps one top-level file child (layout node or declaration).
*/
type FileElement = ast.FileElement

/*
BlankLine inserts vertical spacing in generated output.
*/
type BlankLine = ast.BlankLine

/*
LineComment is a single-line annotation (for example a // comment in Go).
*/
type LineComment = ast.LineComment

/*
BlockComment is a multi-line line-comment block (// per line in Go).
*/
type BlockComment = ast.BlockComment

/*
GoDocComment is a Go block documentation comment attached to declarations or fields.
*/
type GoDocComment = goast.GoDocComment

/*
PackageDecl declares the package clause of a generated file.
*/
type PackageDecl = goast.PackageDecl

/*
ImportBlock declares a parenthesized import block.
*/
type ImportBlock = goast.ImportBlock

/*
TypeStructDecl declares a struct type with fields.
*/
type TypeStructDecl = goast.TypeStructDecl

/*
StructFieldDecl describes one field in a struct, optionally with documentation and leading layout.
*/
type StructFieldDecl = goast.StructFieldDecl

/*
FuncTypeSig is a function type used as a struct field type.
*/
type FuncTypeSig = goast.FuncTypeSig

/*
FuncDecl declares a function with signature and body block.
*/
type FuncDecl = goast.FuncDecl

/*
BlockStmt groups statements and optional leading layout nodes.
*/
type BlockStmt = ast.BlockStmt

/*
VarDeclStmt declares a variable with var syntax.
*/
type VarDeclStmt = goast.VarDeclStmt

/*
ShortVarDeclStmt declares and initializes a variable with := syntax.
*/
type ShortVarDeclStmt = ast.ShortVarDeclStmt

/*
AssignStmt assigns an expression to a name with = syntax.
*/
type AssignStmt = ast.AssignStmt

/*
RangeLoopStmt iterates a slice with for i := range collection.
*/
type RangeLoopStmt = ast.RangeLoopStmt

/*
IfStmt conditionally executes a body, optionally with an init statement.
*/
type IfStmt = ast.IfStmt

/*
ReturnStmt returns a value from the enclosing function.
*/
type ReturnStmt = ast.ReturnStmt

/*
ExprStmt executes an expression for side effects.
*/
type ExprStmt = ast.ExprStmt

/*
IdentExpr references an identifier by name.
*/
type IdentExpr = ast.IdentExpr

/*
SelectorExpr selects a field or method on a base expression.
*/
type SelectorExpr = ast.SelectorExpr

/*
IndexExpr indexes a base expression.
*/
type IndexExpr = ast.IndexExpr

/*
AddressOfExpr takes the address of its operand.
*/
type AddressOfExpr = ast.AddressOfExpr

/*
CallExpr invokes a callee with arguments.
*/
type CallExpr = ast.CallExpr

/*
StringLitExpr is a string literal.
*/
type StringLitExpr = ast.StringLitExpr

/*
IntLitExpr is an integer literal.
*/
type IntLitExpr = ast.IntLitExpr

/*
NilLitExpr is the nil literal.
*/
type NilLitExpr = ast.NilLitExpr

/*
BinaryExpr combines two expressions with a binary operator.
*/
type BinaryExpr = ast.BinaryExpr

/*
BinaryOp identifies a binary operator (equality or inequality).
*/
type BinaryOp = ast.BinaryOp

/*
FieldInit pairs a field name with an initializing expression in a composite literal.
*/
type FieldInit = ast.FieldInit

/*
CompositeLitExpr is a struct or slice composite literal.
*/
type CompositeLitExpr = ast.CompositeLitExpr

/*
TypeExpr describes a Go type in generated code without embedding syntax fragments.
*/
type TypeExpr = goast.TypeExpr

/*
TypeExprKind identifies a type-expression variant.
*/
type TypeExprKind = goast.TypeExprKind

/*
ParamType pairs a parameter name with its type.
*/
type ParamType = goast.ParamType

/*
Emitter is a language-neutral buffered text sink with indentation support.
*/
type Emitter = emit.Emitter

/*
Engine dispatches AST nodes to registered render functions.
*/
type Engine = engine.Engine

/*
RenderContext carries engine and emitter state during a render pass.
*/
type RenderContext = engine.RenderContext

/*
RenderFn renders a single AST node kind.
*/
type RenderFn = engine.RenderFn

const (
	/* NodeKindBlankLine identifies a vertical spacing node. */
	NodeKindBlankLine = ast.NodeKindBlankLine
	/* NodeKindLineComment identifies a single-line comment node. */
	NodeKindLineComment = ast.NodeKindLineComment
	/* NodeKindBlockComment identifies a multi-line line-comment block. */
	NodeKindBlockComment = ast.NodeKindBlockComment
	/* NodeKindGoDocComment identifies a Go block documentation comment node. */
	NodeKindGoDocComment = ast.NodeKindGoDocComment

	/* NodeKindCount is the exclusive upper bound for sizing engine renderer slices. */
	NodeKindCount = ast.NodeKindCount

	/* BinaryOpEq is the == operator. */
	BinaryOpEq = ast.BinaryOpEq
	/* BinaryOpNe is the != operator. */
	BinaryOpNe = ast.BinaryOpNe
)

/*
EngineCreate constructs an engine with a pre-allocated renderer slice of length NodeKindCount.

[Returns]
An engine ready for EngineRegister calls.

[Context]
Uses dense slice indexing by int(kind) for O(1) dispatch without map hashing.
*/
func EngineCreate() *Engine {
	return engine.EngineCreate()
}

/*
EngineRegister binds a renderer function to a node kind index on the given engine.

[Parameters]
eng — target engine instance.
kind — node discriminator; used as a direct slice index.
fn — renderer invoked when that kind is encountered.
*/
func EngineRegister(eng *Engine, kind NodeKind, fn RenderFn) {
	engine.EngineRegister(eng, kind, fn)
}

/*
EngineRender renders a root AST node tree into the emitter buffer.

[Parameters]
eng — engine with renderers registered for every kind in the tree.
emitter — output sink to write into.
root — root node, typically a File.

[Returns]
An error when a renderer is missing or rendering fails.

[Side Effects]
Writes rendered text into emitter.
*/
func EngineRender(eng *Engine, emitter *Emitter, root Node) error {
	return engine.EngineRender(eng, emitter, root)
}

/*
EngineRenderNode renders a single node via direct registry slice lookup.

[Parameters]
ctx — render context with engine and emitter.
node — AST node to render.

[Returns]
An error when no renderer is registered or rendering fails.

[Context]
Hot path: one index into the renderer slice, no map traversal.
*/
func EngineRenderNode(ctx *RenderContext, node Node) error {
	return engine.EngineRenderNode(ctx, node)
}

/*
EmitterCreate constructs an emitter that indents with a fixed number of spaces per level.

[Parameters]
indentWidth — spaces written per indent level.

[Returns]
A fresh emitter.
*/
func EmitterCreate(indentWidth int) Emitter {
	return emit.EmitterCreate(indentWidth)
}

/*
EmitterRender returns the accumulated text written to the emitter.

[Parameters]
emitter — emitter that received writes.

[Returns]
The full rendered string.
*/
func EmitterRender(emitter *Emitter) string {
	return emit.EmitterRender(emitter)
}

/*
FileElementFrom wraps any node as a FileElement for DeclFile.

[Parameters]
node — layout or declaration node.

[Returns]
A FileElement containing node.
*/
func FileElementFrom(node Node) FileElement {
	return ast.FileElementFrom(node)
}

/*
TypeExprFromGoTypeString parses a Go type spelling into a TypeExpr tree.

[Parameters]
typeString — type identifier such as "*uint32" or "[]commandMapping".

[Returns]
The parsed TypeExpr, or an error when typeString is empty.

[Errors]
Returns an error when typeString is empty.
*/
func TypeExprFromGoTypeString(typeString string) (TypeExpr, error) {
	return goast.TypeExprFromGoTypeString(typeString)
}

/*
DeclFile constructs a file root node from ordered top-level elements.

[Parameters]
elements — layout nodes and declarations in render order.

[Returns]
A File AST root.
*/
func DeclFile(elements ...FileElement) File {
	return ast.DeclFile(elements...)
}

/*
DeclPackage constructs a package declaration node.

[Parameters]
name — package name written in the output file.

[Returns]
A PackageDecl node.
*/
func DeclPackage(name string) PackageDecl {
	return goast.DeclPackage(name)
}

/*
DeclImportBlock constructs a parenthesized import block.

[Parameters]
paths — import paths in block order.

[Returns]
An ImportBlock node.
*/
func DeclImportBlock(paths ...string) ImportBlock {
	return goast.DeclImportBlock(paths...)
}

/*
DeclTypeStruct constructs a struct type declaration with fields.

[Parameters]
name — struct type name.
fields — field declarations including optional documentation.

[Returns]
A TypeStructDecl node.
*/
func DeclTypeStruct(name string, fields []StructFieldDecl) TypeStructDecl {
	return goast.DeclTypeStruct(name, fields)
}

/*
StructFieldType declares a struct field with a plain type.

[Parameters]
name — field name.
typ — field type expression.

[Returns]
A StructFieldDecl without documentation.
*/
func StructFieldType(name string, typ TypeExpr) StructFieldDecl {
	return goast.StructFieldType(name, typ)
}

/*
StructFieldFunc declares a struct field with an embedded function type.

[Parameters]
name — field name.
sig — function type signature.

[Returns]
A StructFieldDecl without documentation.
*/
func StructFieldFunc(name string, sig FuncTypeSig) StructFieldDecl {
	return goast.StructFieldFunc(name, sig)
}

/*
StructFieldFuncDoc declares a documented struct field with an embedded function type.

[Parameters]
name — field identifier.
sig — function type signature.
doc — Go documentation block body rendered before the field.
leading — optional layout nodes (BlankLine, LineComment) before the doc.

[Returns]
A StructFieldDecl with documentation and leading layout.
*/
func StructFieldFuncDoc(name string, sig FuncTypeSig, doc string, leading ...Node) StructFieldDecl {
	return goast.StructFieldFuncDoc(name, sig, doc, leading...)
}

/*
TypeExprNamed constructs a named type identifier.

[Parameters]
name — type name spelling.

[Returns]
A TypeExpr with Kind TypeExprKindNamed.
*/
func TypeExprNamed(name string) TypeExpr {
	return goast.TypeExprNamed(name)
}

/*
TypeExprPointer constructs a pointer type expression.

[Parameters]
elem — pointed-to type.

[Returns]
A TypeExpr with Kind TypeExprKindPointer.
*/
func TypeExprPointer(elem TypeExpr) TypeExpr {
	return goast.TypeExprPointer(elem)
}

/*
DeclFunc constructs a function declaration with body.

[Parameters]
name — function name.
params — parameter name/type pairs.
returns — result types (empty for no return).
body — function body block.

[Returns]
A FuncDecl node.
*/
func DeclFunc(name string, params []ParamType, returns []TypeExpr, body BlockStmt) FuncDecl {
	return goast.DeclFunc(name, params, returns, body)
}

/*
ExprIdent constructs an identifier expression.

[Parameters]
name — identifier spelling.

[Returns]
An IdentExpr node.
*/
func ExprIdent(name string) *IdentExpr {
	return ast.ExprIdent(name)
}

/*
ExprSelector constructs a selector expression on a base.

[Parameters]
base — left-hand expression.
name — selected identifier.

[Returns]
A SelectorExpr node.
*/
func ExprSelector(base Expr, name string) *SelectorExpr {
	return ast.ExprSelector(base, name)
}

/*
ExprIndex constructs an index expression.

[Parameters]
base — indexed expression.
index — index expression.

[Returns]
An IndexExpr node.
*/
func ExprIndex(base Expr, index Expr) *IndexExpr {
	return ast.ExprIndex(base, index)
}

/*
ExprAddressOf constructs an address-of expression.

[Parameters]
operand — expression whose address is taken.

[Returns]
An AddressOfExpr node.
*/
func ExprAddressOf(operand Expr) *AddressOfExpr {
	return ast.ExprAddressOf(operand)
}

/*
ExprCall constructs a function or method call expression.

[Parameters]
callee — called expression.
args — argument expressions.

[Returns]
A CallExpr node.
*/
func ExprCall(callee Expr, args ...Expr) *CallExpr {
	return ast.ExprCall(callee, args...)
}

/*
ExprStringLit constructs a string literal expression.

[Parameters]
value — string value (escaped by the Go renderer).

[Returns]
A StringLitExpr node.
*/
func ExprStringLit(value string) *StringLitExpr {
	return ast.ExprStringLit(value)
}

/*
ExprIntLit constructs an integer literal expression.

[Parameters]
value — integer value.

[Returns]
An IntLitExpr node.
*/
func ExprIntLit(value int64) *IntLitExpr {
	return ast.ExprIntLit(value)
}

/*
ExprNil constructs a nil literal expression.

[Returns]
A NilLitExpr node.
*/
func ExprNil() *NilLitExpr {
	return ast.ExprNil()
}

/*
ExprBinary constructs a binary expression.

[Parameters]
op — binary operator.
left — left-hand operand.
right — right-hand operand.

[Returns]
A BinaryExpr node.
*/
func ExprBinary(op BinaryOp, left Expr, right Expr) *BinaryExpr {
	return ast.ExprBinary(op, left, right)
}

/*
ExprCompositeLit constructs a struct composite literal.

[Parameters]
typeName — composite type name.
fields — field initializer pairs.

[Returns]
A CompositeLitExpr node with IsSlice false.
*/
func ExprCompositeLit(typeName string, fields []FieldInit) *CompositeLitExpr {
	return ast.ExprCompositeLit(typeName, fields)
}

/*
ExprSliceCompositeLit constructs a slice of composite literals.

[Parameters]
elementTypeName — slice element type name.
elements — composite literal expressions for each element.

[Returns]
A CompositeLitExpr node with IsSlice true.
*/
func ExprSliceCompositeLit(elementTypeName string, elements []Expr) *CompositeLitExpr {
	return ast.ExprSliceCompositeLit(elementTypeName, elements)
}

/*
StmtBlock constructs a statement block.

[Parameters]
stmts — statements executed in order.

[Returns]
A BlockStmt node.
*/
func StmtBlock(stmts ...Stmt) BlockStmt {
	return ast.StmtBlock(stmts...)
}

/*
StmtVarDecl constructs a var declaration statement.

[Parameters]
name — variable name.
typ — variable type.
isSlice — when true, renders as var name []typ.

[Returns]
A VarDeclStmt node.
*/
func StmtVarDecl(name string, typ TypeExpr, isSlice bool) VarDeclStmt {
	return goast.StmtVarDecl(name, typ, isSlice)
}

/*
StmtShortVarDecl constructs a short variable declaration with :=.

[Parameters]
name — variable name.
rhs — right-hand expression.

[Returns]
A ShortVarDeclStmt node.
*/
func StmtShortVarDecl(name string, rhs Expr) ShortVarDeclStmt {
	return ast.StmtShortVarDecl(name, rhs)
}

/*
StmtRangeLoop constructs a for i := range collection loop.

[Parameters]
collection — slice identifier to range over.
elementName — name bound to collection[i] inside the body.
body — loop body block.

[Returns]
A RangeLoopStmt node.
*/
func StmtRangeLoop(collection string, elementName string, body BlockStmt) RangeLoopStmt {
	return ast.StmtRangeLoop(collection, elementName, body)
}

/*
StmtIf constructs an if statement with a condition expression.

[Parameters]
condition — boolean condition expression.
body — conditional body block.

[Returns]
An IfStmt node without init statement.
*/
func StmtIf(condition Expr, body BlockStmt) IfStmt {
	return ast.StmtIf(condition, body)
}

/*
StmtIfWithInit constructs an if statement with init; condition form.

[Parameters]
init — statement executed before the condition (typically a short var decl).
condition — boolean condition evaluated after init.
body — conditional body block.

[Returns]
An IfStmt node with init.
*/
func StmtIfWithInit(init Stmt, condition Expr, body BlockStmt) IfStmt {
	return ast.StmtIfWithInit(init, condition, body)
}

/*
StmtReturn constructs a return statement.

[Parameters]
value — returned expression.

[Returns]
A ReturnStmt node.
*/
func StmtReturn(value Expr) ReturnStmt {
	return ast.StmtReturn(value)
}

/*
StmtExpr constructs an expression statement.

[Parameters]
expr — expression executed for side effects.

[Returns]
An ExprStmt node.
*/
func StmtExpr(expr Expr) ExprStmt {
	return ast.StmtExpr(expr)
}
