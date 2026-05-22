package goengine

import (
	"fmt"

	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

/*
GoRenderTypeExpr writes a Go type expression to the emitter.

[Parameters]
ctx — render context with emitter.
typ — Go type expression tree.

[Returns]
An error when a type node is malformed or an unknown kind is encountered.

[Context]
Owned by the Go backend; the core engine has no type-expression rendering logic.
*/
func GoRenderTypeExpr(ctx *engine.RenderContext, typ goast.TypeExpr) error {
	switch typ.Kind {
	case goast.TypeExprKindNamed:
		emit.EmitterWrite(ctx.Emitter, typ.Named)
	case goast.TypeExprKindPointer:
		emit.EmitterWrite(ctx.Emitter, "*")
		if typ.Elem == nil {
			return fmt.Errorf("goengine: pointer type missing element")
		}
		return GoRenderTypeExpr(ctx, *typ.Elem)
	case goast.TypeExprKindSlice:
		emit.EmitterWrite(ctx.Emitter, "[]")
		if typ.Elem == nil {
			return fmt.Errorf("goengine: slice type missing element")
		}
		return GoRenderTypeExpr(ctx, *typ.Elem)
	case goast.TypeExprKindArray:
		if typ.ArrayLen == "" || typ.Elem == nil {
			return fmt.Errorf("goengine: array type missing length or element")
		}
		emit.EmitterWriteFormatted(ctx.Emitter, "[%s]", typ.ArrayLen)
		return GoRenderTypeExpr(ctx, *typ.Elem)
	case goast.TypeExprKindFunc:
		emit.EmitterWrite(ctx.Emitter, "func(")
		for i, param := range typ.Params {
			if i != 0 {
				emit.EmitterWrite(ctx.Emitter, ", ")
			}
			emit.EmitterWriteFormatted(ctx.Emitter, "%s ", param.Name)
			if err := GoRenderTypeExpr(ctx, param.Type); err != nil {
				return err
			}
		}
		emit.EmitterWrite(ctx.Emitter, ")")
		if len(typ.Returns) > 0 {
			emit.EmitterWrite(ctx.Emitter, " ")
			for i, ret := range typ.Returns {
				if i != 0 {
					emit.EmitterWrite(ctx.Emitter, ", ")
				}
				if err := GoRenderTypeExpr(ctx, ret); err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("goengine: unknown type expr kind %d", typ.Kind)
	}
	return nil
}
