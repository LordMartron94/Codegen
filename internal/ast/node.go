package ast

/*
NodeKind is an opaque discriminator assigned by a language frontend.

[Context]
The core engine indexes renderers by int(kind). Each frontend owns its own kind numbering via iota or explicit constants in its AST package. The core does not declare any predefined kind values.
*/
type NodeKind uint32

/*
Node is the root interface for every AST value rendered by a codegen engine.

[Context]
Concrete node types live in language frontends (for example internal/goast). Each implements NodeKind() with a frontend-specific kind value.
*/
type Node interface {
	NodeKind() NodeKind
}

/*
Expr is an expression node that produces a value in the target language.
*/
type Expr interface {
	Node
	IsExpr()
}

/*
Stmt is a statement node executed inside function bodies and blocks.
*/
type Stmt interface {
	Node
	IsStmt()
}

/*
FileElement wraps one top-level file child.

[Context]
Source file roots in language frontends hold ordered FileElement slices so layout and declarations can interleave.
*/
type FileElement struct {
	Node Node
}

/*
FileElementFrom wraps any node as a FileElement.

[Parameters]
node — layout or declaration node from a language frontend.

[Returns]
A FileElement containing node.
*/
func FileElementFrom(node Node) FileElement {
	return FileElement{Node: node}
}
