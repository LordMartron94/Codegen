package goengine

import (
	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

func renderBlankLine(ctx *engine.RenderContext, node ast.Node) error {
	_ = node.(ast.BlankLine)
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderLineComment(ctx *engine.RenderContext, node ast.Node) error {
	comment := node.(ast.LineComment)
	emit.EmitterWriteLineFormatted(ctx.Emitter, "// %s", comment.Text)
	return nil
}

func renderBlockComment(ctx *engine.RenderContext, node ast.Node) error {
	comment := node.(ast.BlockComment)
	for _, line := range comment.Lines {
		emit.EmitterWriteLineFormatted(ctx.Emitter, "// %s", line)
	}
	return nil
}

func renderGoDocComment(ctx *engine.RenderContext, node ast.Node) error {
	comment := node.(goast.GoDocComment)
	emit.EmitterWriteLine(ctx.Emitter, "/*")
	for _, line := range comment.Lines {
		emit.EmitterWriteLine(ctx.Emitter, line)
	}
	emit.EmitterWriteLine(ctx.Emitter, "*/")
	return nil
}
