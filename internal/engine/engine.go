package engine

import (
	"fmt"

	"codegen/internal/ast"
	"codegen/internal/emit"
)

/*
RenderContext carries engine and emitter state for a render pass.

[Context]
Passed to every registered renderer so nested nodes can recurse through EngineRenderNode.
*/
type RenderContext struct {
	Engine  *Engine
	Emitter *emit.Emitter
	/* Data holds optional frontend-specific render configuration (for example goast.RenderOptions). */
	Data any
}

/*
RenderFn renders a single AST node kind.

[Context]
Registered per NodeKind index; returns an error when rendering fails instead of panicking for client-facing failures.
*/
type RenderFn func(ctx *RenderContext, node ast.Node) error

/*
Engine dispatches AST nodes to registered render functions via a dense slice registry.

[Context]
Registry is indexed by int(NodeKind) for O(1) cache-local dispatch. Language backends register renderers; the engine package has no language-specific logic.
*/
type Engine struct {
	registry []RenderFn
}

/*
EngineCreate constructs an engine with an empty renderer registry.

[Returns]
An engine ready for EngineRegister calls. Slots are nil until registered.

[Context]
Language frontends register only the node kinds they define. The registry slice grows on demand when a kind index exceeds the current length.
*/
func EngineCreate() *Engine {
	return &Engine{
		registry: make([]RenderFn, 0),
	}
}

/*
EngineRegister binds a renderer to a node kind index.

[Parameters]
engine — target engine.
kind — node discriminator assigned by the language frontend.
fn — renderer invoked when EngineRenderNode encounters kind.

[Context]
Grows the registry slice if kind exceeds the current length.
*/
func EngineRegister(engine *Engine, kind ast.NodeKind, fn RenderFn) {
	kindIdx := int(kind)
	if kindIdx >= len(engine.registry) {
		newRegistry := make([]RenderFn, kindIdx+1)
		copy(newRegistry, engine.registry)
		engine.registry = newRegistry
	}
	engine.registry[kindIdx] = fn
}

/*
EngineRenderNode renders a single node using a direct registry index lookup.

[Parameters]
ctx — render context with engine and emitter.
node — AST node to render.

[Errors]
Returns an error when node is nil, no renderer is registered for the kind, or the renderer returns an error.

[Context]
Hot path: one slice index read per node, no map hashing.
*/
func EngineRenderNode(ctx *RenderContext, node ast.Node) error {
	if node == nil {
		return fmt.Errorf("codegen engine: nil node")
	}

	kindIdx := int(node.NodeKind())
	if kindIdx >= len(ctx.Engine.registry) || ctx.Engine.registry[kindIdx] == nil {
		return fmt.Errorf("codegen engine: no renderer for node kind %d", kindIdx)
	}

	return ctx.Engine.registry[kindIdx](ctx, node)
}

/*
EngineRender renders a root node tree to the emitter buffer.

[Parameters]
engine — configured engine with renderers.
emitter — output sink.
root — root node, typically ast.File.

[Returns]
Error from the root renderer or nested render failures.

[Side Effects]
Mutates emitter buffer content.
*/
func EngineRender(engine *Engine, emitter *emit.Emitter, root ast.Node) error {
	ctx := &RenderContext{Engine: engine, Emitter: emitter}
	return EngineRenderNode(ctx, root)
}

/*
EngineRenderExpr renders an expression node.

[Context]
Convenience wrapper around EngineRenderNode for expr subtrees.
*/
func EngineRenderExpr(ctx *RenderContext, expr ast.Expr) error {
	return EngineRenderNode(ctx, expr)
}

/*
EngineRenderStmt renders a statement node.

[Context]
Convenience wrapper around EngineRenderNode for statement subtrees.
*/
func EngineRenderStmt(ctx *RenderContext, stmt ast.Stmt) error {
	return EngineRenderNode(ctx, stmt)
}
