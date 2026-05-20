package goengine

import (
	"fmt"
	"strconv"

	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
)

func renderIdentExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.IdentExpr)
	emit.EmitterWrite(ctx.Emitter, expr.Name)
	return nil
}

func renderSelectorExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.SelectorExpr)
	if err := engine.EngineRenderExpr(ctx, expr.Base); err != nil {
		return err
	}
	emit.EmitterWriteFormatted(ctx.Emitter, ".%s", expr.Name)
	return nil
}

func renderIndexExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.IndexExpr)
	if err := engine.EngineRenderExpr(ctx, expr.Base); err != nil {
		return err
	}
	emit.EmitterWrite(ctx.Emitter, "[")
	if err := engine.EngineRenderExpr(ctx, expr.Index); err != nil {
		return err
	}
	emit.EmitterWrite(ctx.Emitter, "]")
	return nil
}

func renderAddressOfExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.AddressOfExpr)
	emit.EmitterWrite(ctx.Emitter, "&")
	return engine.EngineRenderExpr(ctx, expr.Operand)
}

func renderCallExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.CallExpr)
	if err := engine.EngineRenderExpr(ctx, expr.Callee); err != nil {
		return err
	}
	emit.EmitterWrite(ctx.Emitter, "(")
	for i, arg := range expr.Args {
		if i != 0 {
			emit.EmitterWrite(ctx.Emitter, ", ")
		}
		if err := engine.EngineRenderExpr(ctx, arg); err != nil {
			return err
		}
	}
	emit.EmitterWrite(ctx.Emitter, ")")
	return nil
}

func renderStringLitExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.StringLitExpr)
	emit.EmitterWrite(ctx.Emitter, strconv.Quote(expr.Value))
	return nil
}

func renderIntLitExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.IntLitExpr)
	emit.EmitterWriteFormatted(ctx.Emitter, "%d", expr.Value)
	return nil
}

func renderNilLitExpr(ctx *engine.RenderContext, node ast.Node) error {
	_ = node.(*ast.NilLitExpr)
	emit.EmitterWrite(ctx.Emitter, "nil")
	return nil
}

func renderBinaryExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.BinaryExpr)
	if err := engine.EngineRenderExpr(ctx, expr.Left); err != nil {
		return err
	}
	switch expr.Op {
	case ast.BinaryOpEq:
		emit.EmitterWrite(ctx.Emitter, " == ")
	case ast.BinaryOpNe:
		emit.EmitterWrite(ctx.Emitter, " != ")
	default:
		return fmt.Errorf("goengine: unknown binary op %d", expr.Op)
	}
	return engine.EngineRenderExpr(ctx, expr.Right)
}

func renderCompositeLitExpr(ctx *engine.RenderContext, node ast.Node) error {
	expr := node.(*ast.CompositeLitExpr)
	if expr.IsSlice {
		emit.EmitterWriteFormatted(ctx.Emitter, "[]%s{", expr.TypeName)
		emit.EmitterLineBreak(ctx.Emitter)
		emit.EmitterIndent(ctx.Emitter)
		for _, element := range expr.Elements {
			if err := engine.EngineRenderExpr(ctx, element); err != nil {
				return err
			}
			emit.EmitterWrite(ctx.Emitter, ",")
			emit.EmitterLineBreak(ctx.Emitter)
		}
		emit.EmitterDedent(ctx.Emitter)
		emit.EmitterWrite(ctx.Emitter, "}")
		return nil
	}

	emit.EmitterWriteLine(ctx.Emitter, "{")
	emit.EmitterIndent(ctx.Emitter)
	for _, field := range expr.Fields {
		emit.EmitterWriteFormatted(ctx.Emitter, "%s: ", field.Name)
		if err := engine.EngineRenderExpr(ctx, field.Value); err != nil {
			return err
		}
		emit.EmitterWrite(ctx.Emitter, ",")
		emit.EmitterLineBreak(ctx.Emitter)
	}
	emit.EmitterDedent(ctx.Emitter)
	emit.EmitterWrite(ctx.Emitter, "}")
	return nil
}
