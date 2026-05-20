package ast

/*
BlankLine inserts vertical spacing in generated output.

[Context]
Rendered by layout engines as an empty line; language-agnostic.
*/
type BlankLine struct{}

func (BlankLine) NodeKind() NodeKind { return NodeKindBlankLine }

/*
LineComment is a single-line annotation in generated output.

[Context]
Go engines render Text as a // comment; other engines may map to different comment syntax.
*/
type LineComment struct {
	Text string
}

func (LineComment) NodeKind() NodeKind { return NodeKindLineComment }

/*
BlockComment is a multi-line annotation in generated output.

[Context]
Go engines render each entry in Lines; other engines choose block or repeated line comments.
*/
type BlockComment struct {
	Lines []string
}

func (BlockComment) NodeKind() NodeKind { return NodeKindBlockComment }

func LayoutBlankLineNew() BlankLine {
	return BlankLine{}
}

func LayoutLineCommentNew(text string) LineComment {
	return LineComment{Text: text}
}

func LayoutBlockCommentNew(lines ...string) BlockComment {
	return BlockComment{Lines: lines}
}
