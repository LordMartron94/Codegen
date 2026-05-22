package goast

import "codegen/internal/ast"

/*
BlankLine inserts vertical spacing in generated Go output.
*/
type BlankLine struct{}

func (BlankLine) NodeKind() ast.NodeKind { return KindBlankLine }

/*
LineComment is a single-line // comment.
*/
type LineComment struct {
	Text string
}

func (LineComment) NodeKind() ast.NodeKind { return KindLineComment }

/*
BlockComment is a multi-line // comment block.
*/
type BlockComment struct {
	Lines []string
}

func (BlockComment) NodeKind() ast.NodeKind { return KindBlockComment }

/*
GoDocComment is a Go documentation block attached to a declaration or field.
*/
type GoDocComment struct {
	Lines []string
}

func (GoDocComment) NodeKind() ast.NodeKind { return KindGoDocComment }

func LayoutBlankLineNew() BlankLine {
	return BlankLine{}
}

func LayoutLineCommentNew(text string) LineComment {
	return LineComment{Text: text}
}

func LayoutBlockCommentNew(lines ...string) BlockComment {
	return BlockComment{Lines: lines}
}

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
