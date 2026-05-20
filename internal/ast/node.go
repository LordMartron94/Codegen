package ast

/*
NodeKind identifies an AST node variant for engine registration and dispatch.

[Context]
Kinds are grouped by family: layout, declaration, statement, and expression nodes share one flat enumeration so a single registry can render any subtree. Kinds from NodeKindGoDocComment through NodeKindFuncDecl and NodeKindVarDeclStmt are Go-backend shapes (structs live in internal/goast); expression and neutral layout kinds are language-agnostic.
*/
type NodeKind uint16

const (
	// Layout nodes.
	NodeKindBlankLine NodeKind = iota
	NodeKindLineComment
	NodeKindBlockComment
	NodeKindGoDocComment

	// File and declaration nodes.
	NodeKindFile
	NodeKindPackageDecl
	NodeKindImportBlock
	NodeKindTypeStructDecl
	NodeKindFuncDecl

	// Statement nodes.
	NodeKindBlockStmt
	NodeKindVarDeclStmt
	NodeKindShortVarDeclStmt
	NodeKindAssignStmt
	NodeKindRangeLoopStmt
	NodeKindIfStmt
	NodeKindReturnStmt
	NodeKindExprStmt

	// Expression nodes.
	NodeKindIdentExpr
	NodeKindSelectorExpr
	NodeKindIndexExpr
	NodeKindAddressOfExpr
	NodeKindCallExpr
	NodeKindStringLitExpr
	NodeKindIntLitExpr
	NodeKindNilLitExpr
	NodeKindBinaryExpr
	NodeKindCompositeLitExpr

	_nodeKindCount
)

/*
NodeKindCount is the exclusive upper bound for NodeKind values used to size engine registries.

[Context]
EngineCreate pre-allocates a renderer slice of this length for O(1) index dispatch without map hashing.
*/
const NodeKindCount = _nodeKindCount

/*
Node is the root interface for all AST nodes rendered by a codegen engine.

[Context]
Concrete node types are plain structs with a NodeKind discriminator; renderers switch on NodeKind via the engine registry.
*/
type Node interface {
	NodeKind() NodeKind
}

/*
Expr is an expression node producing a value in the target language.

[Context]
All value positions in statements and composite literals use Expr implementations.
*/
type Expr interface {
	Node
	IsExpr()
}

/*
Stmt is a statement node executed inside function bodies and blocks.

[Context]
BlockStmt holds Stmt children; some Stmt nodes embed Expr for conditions and right-hand sides. Go-specific stmts such as VarDeclStmt live in internal/goast and implement this interface.
*/
type Stmt interface {
	Node
	IsStmt()
}

/*
FileElement wraps a top-level file child (layout or declaration).

[Context]
Files use an ordered element list so comments and blank lines can appear between declarations.
*/
type FileElement struct {
	Node Node
}

func FileElementFrom(node Node) FileElement {
	return FileElement{Node: node}
}
