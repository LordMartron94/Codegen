package goengine

import (
	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

func renderBlankLine(ctx *engine.RenderContext, node ast.Node) error {
	_ = node.(goast.BlankLine)
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderLineComment(ctx *engine.RenderContext, node ast.Node) error {
	comment := node.(goast.LineComment)
	emit.EmitterWriteLineFormatted(ctx.Emitter, "// %s", comment.Text)
	return nil
}

func renderBlockComment(ctx *engine.RenderContext, node ast.Node) error {
	comment := node.(goast.BlockComment)
	for _, line := range comment.Lines {
		emit.EmitterWriteLineFormatted(ctx.Emitter, "// %s", line)
	}
	return nil
}

func renderBuildConstraint(ctx *engine.RenderContext, node ast.Node) error {
	constraint := node.(goast.BuildConstraint)
	emit.EmitterWriteLineFormatted(ctx.Emitter, "//go:build %s", constraint.Tag)
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
