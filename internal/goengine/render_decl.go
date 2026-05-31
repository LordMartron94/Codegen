package goengine

import (
	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

func renderFile(ctx *engine.RenderContext, node ast.Node) error {
	file := node.(goast.File)
	for _, element := range file.Elements {
		if err := engine.EngineRenderNode(ctx, element.Node); err != nil {
			return err
		}
	}
	return nil
}

func renderPackageDecl(ctx *engine.RenderContext, node ast.Node) error {
	decl := node.(goast.PackageDecl)
	emit.EmitterWriteLineFormatted(ctx.Emitter, "package %s", decl.Name)
	return nil
}

func renderImportBlock(ctx *engine.RenderContext, node ast.Node) error {
	block := node.(goast.ImportBlock)
	emit.EmitterWriteLine(ctx.Emitter, "import (")
	emit.EmitterIndent(ctx.Emitter)
	for _, path := range block.Paths {
		emit.EmitterWriteLineFormatted(ctx.Emitter, "%q", path)
	}
	emit.EmitterDedent(ctx.Emitter)
	emit.EmitterWriteLine(ctx.Emitter, ")")
	return nil
}

func renderTypeStructDecl(ctx *engine.RenderContext, node ast.Node) error {
	decl := node.(goast.TypeStructDecl)
	if decl.Doc != "" {
		if err := renderDocumentation(ctx, decl.Name, decl.Doc); err != nil {
			return err
		}
	}
	emit.EmitterWriteLineFormatted(ctx.Emitter, "type %s struct {", decl.Name)
	emit.EmitterIndent(ctx.Emitter)
	for _, field := range decl.Fields {
		for _, leading := range field.Leading {
			if err := engine.EngineRenderNode(ctx, leading); err != nil {
				return err
			}
		}
		if field.Doc != "" {
			if err := renderDocumentation(ctx, field.Name, field.Doc); err != nil {
				return err
			}
		}
		emit.EmitterWriteFormatted(ctx.Emitter, "%s ", field.Name)
		if field.FuncType != nil {
			if err := renderFuncTypeSig(ctx, *field.FuncType); err != nil {
				return err
			}
			emit.EmitterLineBreak(ctx.Emitter)
			continue
		}
		if err := GoRenderTypeExpr(ctx, field.Type); err != nil {
			return err
		}
		emit.EmitterLineBreak(ctx.Emitter)
	}
	emit.EmitterDedent(ctx.Emitter)
	emit.EmitterWriteLine(ctx.Emitter, "}")
	return nil
}

func renderTypeDefinedDecl(ctx *engine.RenderContext, node ast.Node) error {
	decl := node.(goast.TypeDefinedDecl)
	for _, leading := range decl.Leading {
		if err := engine.EngineRenderNode(ctx, leading); err != nil {
			return err
		}
	}
	if decl.Doc != "" {
		if err := renderDocumentation(ctx, decl.Name, decl.Doc); err != nil {
			return err
		}
	}
	if decl.IsAlias {
		emit.EmitterWriteFormatted(ctx.Emitter, "type %s = ", decl.Name)
	} else {
		emit.EmitterWriteFormatted(ctx.Emitter, "type %s ", decl.Name)
	}
	if err := GoRenderTypeExpr(ctx, decl.Underlying); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderConstGroupDecl(ctx *engine.RenderContext, node ast.Node) error {
	decl := node.(goast.ConstGroupDecl)
	for _, leading := range decl.Leading {
		if err := engine.EngineRenderNode(ctx, leading); err != nil {
			return err
		}
	}
	if decl.Doc != "" {
		if err := renderDocumentation(ctx, "", decl.Doc); err != nil {
			return err
		}
	}
	emit.EmitterWriteLine(ctx.Emitter, "const (")
	emit.EmitterIndent(ctx.Emitter)
	for specIndex, spec := range decl.Specs {
		if spec.Doc != "" && decl.SeparateDocumentedSpecs && specIndex > 0 {
			emit.EmitterLineBreak(ctx.Emitter)
		}
		if spec.Doc != "" {
			if err := renderDocumentation(ctx, spec.Name, spec.Doc); err != nil {
				return err
			}
		}
		emit.EmitterWriteFormatted(ctx.Emitter, "%s", spec.Name)
		if spec.Type != nil {
			emit.EmitterWrite(ctx.Emitter, " ")
			if err := GoRenderTypeExpr(ctx, *spec.Type); err != nil {
				return err
			}
		}
		emit.EmitterWriteFormatted(ctx.Emitter, " = %s", spec.Value)
		emit.EmitterLineBreak(ctx.Emitter)
	}
	emit.EmitterDedent(ctx.Emitter)
	emit.EmitterWriteLine(ctx.Emitter, ")")
	return nil
}

func renderVarDecl(ctx *engine.RenderContext, node ast.Node) error {
	decl := node.(goast.VarDecl)
	if decl.Doc != "" {
		if err := renderDocumentation(ctx, decl.Name, decl.Doc); err != nil {
			return err
		}
	}
	emit.EmitterWriteFormatted(ctx.Emitter, "var %s", decl.Name)
	if decl.Type != nil {
		emit.EmitterWrite(ctx.Emitter, " ")
		if err := GoRenderTypeExpr(ctx, *decl.Type); err != nil {
			return err
		}
		emit.EmitterLineBreak(ctx.Emitter)
		return nil
	}
	emit.EmitterWrite(ctx.Emitter, " = ")
	if decl.Init == nil {
		emit.EmitterWrite(ctx.Emitter, "nil")
		emit.EmitterLineBreak(ctx.Emitter)
		return nil
	}
	if err := engine.EngineRenderExpr(ctx, decl.Init); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderFuncDecl(ctx *engine.RenderContext, node ast.Node) error {
	decl := node.(goast.FuncDecl)
	emit.EmitterWriteFormatted(ctx.Emitter, "func %s(", decl.Name)
	for i, param := range decl.Params {
		if i != 0 {
			emit.EmitterWrite(ctx.Emitter, ", ")
		}
		emit.EmitterWriteFormatted(ctx.Emitter, "%s ", param.Name)
		if err := GoRenderTypeExpr(ctx, param.Type); err != nil {
			return err
		}
	}
	emit.EmitterWrite(ctx.Emitter, ")")
	if len(decl.Returns) > 0 {
		emit.EmitterWrite(ctx.Emitter, " ")
		for i, ret := range decl.Returns {
			if i != 0 {
				emit.EmitterWrite(ctx.Emitter, ", ")
			}
			if err := GoRenderTypeExpr(ctx, ret); err != nil {
				return err
			}
		}
	}
	emit.EmitterWrite(ctx.Emitter, " {")
	emit.EmitterLineBreak(ctx.Emitter)
	emit.EmitterIndent(ctx.Emitter)
	if err := renderBlockStmt(ctx, decl.Body); err != nil {
		return err
	}
	emit.EmitterDedent(ctx.Emitter)
	emit.EmitterWriteLine(ctx.Emitter, "}")
	return nil
}

func renderFuncTypeSig(ctx *engine.RenderContext, sig goast.FuncTypeSig) error {
	emit.EmitterWrite(ctx.Emitter, "func(")
	for i, param := range sig.Params {
		if i != 0 {
			emit.EmitterWrite(ctx.Emitter, ", ")
		}
		emit.EmitterWriteFormatted(ctx.Emitter, "%s ", param.Name)
		if err := GoRenderTypeExpr(ctx, param.Type); err != nil {
			return err
		}
	}
	emit.EmitterWrite(ctx.Emitter, ")")
	if len(sig.Returns) > 0 {
		emit.EmitterWrite(ctx.Emitter, " ")
		for i, ret := range sig.Returns {
			if i != 0 {
				emit.EmitterWrite(ctx.Emitter, ", ")
			}
			if err := GoRenderTypeExpr(ctx, ret); err != nil {
				return err
			}
		}
	}
	return nil
}
