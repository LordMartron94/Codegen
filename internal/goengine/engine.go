package goengine

import (
	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
)

/*
GoEngineCreate constructs an engine with all Go renderers registered.

[Returns]
An engine ready to render ast.File trees produced by clients such as VulkanGen.

[Context]
Registers layout, declaration, statement, and expression node kinds in dependency-safe order.
*/
func GoEngineCreate() *engine.Engine {
	eng := engine.EngineCreate()

	engine.EngineRegister(eng, ast.NodeKindBlankLine, renderBlankLine)
	engine.EngineRegister(eng, ast.NodeKindLineComment, renderLineComment)
	engine.EngineRegister(eng, ast.NodeKindBlockComment, renderBlockComment)
	engine.EngineRegister(eng, ast.NodeKindGoDocComment, renderGoDocComment)

	engine.EngineRegister(eng, ast.NodeKindFile, renderFile)
	engine.EngineRegister(eng, ast.NodeKindPackageDecl, renderPackageDecl)
	engine.EngineRegister(eng, ast.NodeKindImportBlock, renderImportBlock)
	engine.EngineRegister(eng, ast.NodeKindTypeStructDecl, renderTypeStructDecl)
	engine.EngineRegister(eng, ast.NodeKindFuncDecl, renderFuncDecl)

	engine.EngineRegister(eng, ast.NodeKindBlockStmt, renderBlockStmtNode)
	engine.EngineRegister(eng, ast.NodeKindVarDeclStmt, renderVarDeclStmt)
	engine.EngineRegister(eng, ast.NodeKindShortVarDeclStmt, renderShortVarDeclStmt)
	engine.EngineRegister(eng, ast.NodeKindAssignStmt, renderAssignStmt)
	engine.EngineRegister(eng, ast.NodeKindRangeLoopStmt, renderRangeLoopStmt)
	engine.EngineRegister(eng, ast.NodeKindIfStmt, renderIfStmt)
	engine.EngineRegister(eng, ast.NodeKindReturnStmt, renderReturnStmt)
	engine.EngineRegister(eng, ast.NodeKindExprStmt, renderExprStmt)

	engine.EngineRegister(eng, ast.NodeKindIdentExpr, renderIdentExpr)
	engine.EngineRegister(eng, ast.NodeKindSelectorExpr, renderSelectorExpr)
	engine.EngineRegister(eng, ast.NodeKindIndexExpr, renderIndexExpr)
	engine.EngineRegister(eng, ast.NodeKindAddressOfExpr, renderAddressOfExpr)
	engine.EngineRegister(eng, ast.NodeKindCallExpr, renderCallExpr)
	engine.EngineRegister(eng, ast.NodeKindStringLitExpr, renderStringLitExpr)
	engine.EngineRegister(eng, ast.NodeKindIntLitExpr, renderIntLitExpr)
	engine.EngineRegister(eng, ast.NodeKindNilLitExpr, renderNilLitExpr)
	engine.EngineRegister(eng, ast.NodeKindBinaryExpr, renderBinaryExpr)
	engine.EngineRegister(eng, ast.NodeKindCompositeLitExpr, renderCompositeLitExpr)

	return eng
}

/*
GoEmitterCreate constructs an emitter configured for Go output.

[Returns]
Emitter with tab indentation per level, matching gofmt conventions.
*/
func GoEmitterCreate() emit.Emitter {
	return emit.EmitterCreateTabs()
}
