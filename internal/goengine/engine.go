package goengine

import (
	"codegen/internal/emit"
	"codegen/internal/engine"
	"codegen/internal/goast"
)

/*
GoEngineCreate constructs an engine with all Go node renderers registered.

[Returns]
Configured engine for EngineRender.

[Context]
Registers renderers for every Kind defined in internal/goast. The core engine has no predefined node kinds.
*/
func GoEngineCreate() *engine.Engine {
	eng := engine.EngineCreate()

	engine.EngineRegister(eng, goast.KindBlankLine, renderBlankLine)
	engine.EngineRegister(eng, goast.KindLineComment, renderLineComment)
	engine.EngineRegister(eng, goast.KindBlockComment, renderBlockComment)
	engine.EngineRegister(eng, goast.KindGoDocComment, renderGoDocComment)

	engine.EngineRegister(eng, goast.KindFile, renderFile)
	engine.EngineRegister(eng, goast.KindPackageDecl, renderPackageDecl)
	engine.EngineRegister(eng, goast.KindImportBlock, renderImportBlock)
	engine.EngineRegister(eng, goast.KindTypeStructDecl, renderTypeStructDecl)
	engine.EngineRegister(eng, goast.KindTypeDefinedDecl, renderTypeDefinedDecl)
	engine.EngineRegister(eng, goast.KindConstGroupDecl, renderConstGroupDecl)
	engine.EngineRegister(eng, goast.KindVarDecl, renderVarDecl)
	engine.EngineRegister(eng, goast.KindFuncDecl, renderFuncDecl)

	engine.EngineRegister(eng, goast.KindBlockStmt, renderBlockStmtNode)
	engine.EngineRegister(eng, goast.KindVarDeclStmt, renderVarDeclStmt)
	engine.EngineRegister(eng, goast.KindShortVarDeclStmt, renderShortVarDeclStmt)
	engine.EngineRegister(eng, goast.KindAssignStmt, renderAssignStmt)
	engine.EngineRegister(eng, goast.KindRangeLoopStmt, renderRangeLoopStmt)
	engine.EngineRegister(eng, goast.KindIfStmt, renderIfStmt)
	engine.EngineRegister(eng, goast.KindReturnStmt, renderReturnStmt)
	engine.EngineRegister(eng, goast.KindExprStmt, renderExprStmt)

	engine.EngineRegister(eng, goast.KindIdentExpr, renderIdentExpr)
	engine.EngineRegister(eng, goast.KindSelectorExpr, renderSelectorExpr)
	engine.EngineRegister(eng, goast.KindIndexExpr, renderIndexExpr)
	engine.EngineRegister(eng, goast.KindAddressOfExpr, renderAddressOfExpr)
	engine.EngineRegister(eng, goast.KindCallExpr, renderCallExpr)
	engine.EngineRegister(eng, goast.KindStringLitExpr, renderStringLitExpr)
	engine.EngineRegister(eng, goast.KindIntLitExpr, renderIntLitExpr)
	engine.EngineRegister(eng, goast.KindNilLitExpr, renderNilLitExpr)
	engine.EngineRegister(eng, goast.KindBoolLitExpr, renderBoolLitExpr)
	engine.EngineRegister(eng, goast.KindBinaryExpr, renderBinaryExpr)
	engine.EngineRegister(eng, goast.KindCompositeLitExpr, renderCompositeLitExpr)

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
