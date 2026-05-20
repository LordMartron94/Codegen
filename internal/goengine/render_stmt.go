package goengine

import (
	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

func renderBlockStmtNode(ctx *engine.RenderContext, node ast.Node) error {
	return renderBlockStmt(ctx, node.(ast.BlockStmt))
}

func renderBlockStmt(ctx *engine.RenderContext, block ast.BlockStmt) error {
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
	stmt := node.(ast.ShortVarDeclStmt)
	emit.EmitterWriteFormatted(ctx.Emitter, "%s := ", stmt.Name)
	if err := engine.EngineRenderExpr(ctx, stmt.Rhs); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderAssignStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(ast.AssignStmt)
	emit.EmitterWriteFormatted(ctx.Emitter, "%s = ", stmt.Name)
	if err := engine.EngineRenderExpr(ctx, stmt.Rhs); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderRangeLoopStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(ast.RangeLoopStmt)
	emit.EmitterWriteFormatted(ctx.Emitter, "for i := range %s {", stmt.Collection)
	emit.EmitterLineBreak(ctx.Emitter)
	emit.EmitterIndent(ctx.Emitter)

	elementRhs := ast.ExprIndex(
		ast.ExprIdent(stmt.Collection),
		ast.ExprIdent("i"),
	)
	if err := engine.EngineRenderStmt(ctx, ast.StmtShortVarDecl(stmt.ElementName, elementRhs)); err != nil {
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
	stmt := node.(ast.IfStmt)
	emit.EmitterWrite(ctx.Emitter, "if ")
	if stmt.Init != nil {
		if initDecl, ok := stmt.Init.(ast.ShortVarDeclStmt); ok {
			emit.EmitterWriteFormatted(ctx.Emitter, "%s := ", initDecl.Name)
			if err := engine.EngineRenderExpr(ctx, initDecl.Rhs); err != nil {
				return err
			}
			emit.EmitterWrite(ctx.Emitter, "; ")
		} else {
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
	stmt := node.(ast.ReturnStmt)
	emit.EmitterWrite(ctx.Emitter, "return ")
	if err := engine.EngineRenderExpr(ctx, stmt.Value); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}

func renderExprStmt(ctx *engine.RenderContext, node ast.Node) error {
	stmt := node.(ast.ExprStmt)
	if err := engine.EngineRenderExpr(ctx, stmt.Expr); err != nil {
		return err
	}
	emit.EmitterLineBreak(ctx.Emitter)
	return nil
}
