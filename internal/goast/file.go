package goast

import "codegen/internal/ast"

/*
File is the root node for a generated Go source file.

[Context]
Elements are rendered in order, interleaving layout nodes and top-level declarations.
*/
type File struct {
	Elements []ast.FileElement
}

func (File) NodeKind() ast.NodeKind { return KindFile }

/*
DeclFile constructs a Go file root node from ordered top-level elements.
*/
func DeclFile(elements ...ast.FileElement) File {
	return File{Elements: elements}
}
