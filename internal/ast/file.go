package ast

/*
File is the root node for a generated source file.

[Context]
Elements are rendered in order, interleaving layout nodes and top-level declarations.
*/
type File struct {
	Elements []FileElement
}

func (File) NodeKind() NodeKind { return NodeKindFile }

func DeclFile(elements ...FileElement) File {
	return File{Elements: elements}
}
