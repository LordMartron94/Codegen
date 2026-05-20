package goast

import "codegen/internal/ast"

/*
GoDocComment is a Go documentation block attached to a declaration or field.

[Context]
Go engines render Lines as a block documentation comment; used for generated API documentation on struct fields and declarations.
*/
type GoDocComment struct {
	Lines []string
}

func (GoDocComment) NodeKind() ast.NodeKind { return ast.NodeKindGoDocComment }

func LayoutGoDocCommentNew(lines ...string) GoDocComment {
	return GoDocComment{Lines: lines}
}

func LayoutGoDocCommentFromText(text string) GoDocComment {
	if text == "" {
		return GoDocComment{}
	}
	lines := make([]string, 0)
	start := 0
	for i := 0; i <= len(text); i++ {
		if i == len(text) || text[i] == '\n' {
			lines = append(lines, text[start:i])
			start = i + 1
		}
	}
	return GoDocComment{Lines: lines}
}
