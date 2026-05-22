package goengine

import (
	"strings"

	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

func renderOptionsFromContext(ctx *engine.RenderContext) goast.RenderOptions {
	if ctx == nil || ctx.Data == nil {
		return goast.RenderOptionsDefault()
	}
	options, ok := ctx.Data.(goast.RenderOptions)
	if !ok {
		return goast.RenderOptionsDefault()
	}
	return options
}

func renderDocumentation(ctx *engine.RenderContext, subject string, doc string) error {
	doc = goast.GoDocFormatExported(subject, doc)
	if doc == "" {
		return nil
	}

	options := renderOptionsFromContext(ctx)
	switch options.DocCommentStyle {
	case goast.CommentStyleBlock:
		return engine.EngineRenderNode(ctx, goast.LayoutGoDocCommentFromText(doc))
	default:
		for _, line := range strings.Split(doc, "\n") {
			emit.EmitterWriteLineFormatted(ctx.Emitter, "// %s", line)
		}
		return nil
	}
}
