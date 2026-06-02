package goengine

import (
	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

func renderBlockStmtNode(ctx *engine.RenderContext, node ast.Node) error {
	return renderBlockStmt(ctx, node.(goast.BlockStmt))
}

func renderBlockStmt(ctx *engine.RenderContext, block goast.BlockStmt) error {
	for _, leading := range block.Leading {
		if err := engine.EngineRenderNode(ctx, leading); err != nil {
			return err
		}
	}
	for _, stmt := range block.Stmts {
		if err := engine.EngineRenderStmt(ctx, stmt); err != nil {
			return err
		}
	}
	return nil
}

func renderVarDeclStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(goast.VarDeclStmt)
	if stmt.IsSliceVar {
		emit.EmitterWriteFormatted(ctx.Emitter, "var %s []", stmt.Name)
		if stmt.Type.Kind != goast.TypeExprKindNamed {
			return GoRenderTypeExpr(ctx, stmt.Type)
		}
		emit.EmitterWriteLine(ctx.Emitter, stmt.Type.Named)
		return nil
	}
	emit.EmitterWriteFormatted(ctx.Emitter, "var %s ", stmt.Name)
	if err := GoRenderTypeExpr(ctx, stmt.Type); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderShortVarDeclStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(goast.ShortVarDeclStmt)
	emit.EmitterWriteFormatted(ctx.Emitter, "%s := ", stmt.Name)
	if err := engine.EngineRenderExpr(ctx, stmt.Rhs); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderAssignStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(goast.AssignStmt)
	if stmt.Lhs != nil {
		if err := engine.EngineRenderExpr(ctx, stmt.Lhs); err != nil {
			return err
		}
	} else {
		emit.EmitterWrite(ctx.Emitter, stmt.Name)
	}
	emit.EmitterWrite(ctx.Emitter, " = ")
	if err := engine.EngineRenderExpr(ctx, stmt.Rhs); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderRangeLoopStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(goast.RangeLoopStmt)
	emit.EmitterWriteFormatted(ctx.Emitter, "for i := range %s {", stmt.Collection)
	emit.EmitterLineBreak(ctx.Emitter)
	emit.EmitterIndent(ctx.Emitter)

	elementRhs := goast.ExprIndex(
		goast.ExprIdent(stmt.Collection),
		goast.ExprIdent("i"),
	)
	if err := engine.EngineRenderStmt(ctx, goast.StmtShortVarDecl(stmt.ElementName, elementRhs)); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)

	if err := renderBlockStmt(ctx, stmt.Body); err != nil {
		return err
	}

	emit.EmitterDedent(ctx.Emitter)
	emit.EmitterWriteLine(ctx.Emitter, "}")
	return nil
}

func renderIfStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(goast.IfStmt)
	emit.EmitterWrite(ctx.Emitter, "if ")
	if stmt.Init != nil {
		switch init := stmt.Init.(type) {
		case goast.ShortVarDeclStmt:
			emit.EmitterWriteFormatted(ctx.Emitter, "%s := ", init.Name)
			if err := engine.EngineRenderExpr(ctx, init.Rhs); err != nil {
				return err
			}
			emit.EmitterWrite(ctx.Emitter, "; ")
		case goast.AssignStmt:
			if init.Lhs != nil {
				if err := engine.EngineRenderExpr(ctx, init.Lhs); err != nil {
					return err
				}
			} else {
				emit.EmitterWrite(ctx.Emitter, init.Name)
			}
			emit.EmitterWrite(ctx.Emitter, " = ")
			if err := engine.EngineRenderExpr(ctx, init.Rhs); err != nil {
				return err
			}
			emit.EmitterWrite(ctx.Emitter, "; ")
		default:
			if err := engine.EngineRenderStmt(ctx, stmt.Init); err != nil {
				return err
			}
			emit.EmitterWrite(ctx.Emitter, "; ")
		}
	}
	if err := engine.EngineRenderExpr(ctx, stmt.Condition); err != nil {
		return err
	}
	emit.EmitterWrite(ctx.Emitter, " {")
	emit.EmitterLineBreak(ctx.Emitter)
	emit.EmitterIndent(ctx.Emitter)
	if err := renderBlockStmt(ctx, stmt.Body); err != nil {
		return err
	}
	emit.EmitterDedent(ctx.Emitter)
	emit.EmitterWriteLine(ctx.Emitter, "}")
	return nil
}

func renderReturnStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(goast.ReturnStmt)
	emit.EmitterWrite(ctx.Emitter, "return ")
	if err := engine.EngineRenderExpr(ctx, stmt.Value); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderExprStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(goast.ExprStmt)
	if err := engine.EngineRenderExpr(ctx, stmt.Expr); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}
