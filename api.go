/*
Package codegen provides a language-agnostic codegen engine and emitter.

[Context]
Language frontends (for example codegen/go) define their own AST node kinds, node structs, and renderers. Register those renderers on an engine created here, then render through an emitter. The core does not declare any node kind constants or language-specific AST shapes.

[Example]

	eng := gocode.GoEngineCreate()
	emitter := gocode.GoEmitterCreate()
	file := gocode.DeclFile(...)
	_ = codegen.EngineRender(eng, &emitter, file)
	output := codegen.EmitterRender(&emitter)
*/
package codegen

import (
	"codegen/internal/ast"
	"codegen/internal/emit"
	"codegen/internal/engine"
)

/*
Node is the root interface for every AST value rendered by an engine.
*/
type Node = ast.Node

/*
NodeKind is an opaque discriminator assigned by a language frontend.
*/
type NodeKind = ast.NodeKind

/*
Expr is an expression node that produces a value in the target language.
*/
type Expr = ast.Expr

/*
Stmt is a statement node executed inside function bodies and blocks.
*/
type Stmt = ast.Stmt

/*
FileElement wraps one top-level file child from a language frontend.
*/
type FileElement = ast.FileElement

/*
Emitter is a language-neutral buffered text sink with indentation support.
*/
type Emitter = emit.Emitter

/*
Engine dispatches AST nodes to registered render functions.
*/
type Engine = engine.Engine

/*
RenderContext carries engine and emitter state during a render pass.
*/
type RenderContext = engine.RenderContext

/*
RenderFn renders a single AST node kind.
*/
type RenderFn = engine.RenderFn

/*
EngineCreate constructs an engine with an empty renderer registry.

[Returns]
An engine ready for EngineRegister calls from a language frontend.
*/
func EngineCreate() *Engine {
	return engine.EngineCreate()
}

/*
EngineRegister binds a renderer function to a node kind index on the given engine.

[Parameters]
eng — target engine instance.
kind — node discriminator assigned by the language frontend.
fn — renderer invoked when that kind is encountered.
*/
func EngineRegister(eng *Engine, kind NodeKind, fn RenderFn) {
	engine.EngineRegister(eng, kind, fn)
}

/*
EngineRender renders a root AST node tree into the emitter buffer.

[Parameters]
eng — engine with renderers registered for every kind in the tree.
emitter — output sink to write into.
root — root node from a language frontend.

[Returns]
An error when a renderer is missing or rendering fails.

[Side Effects]
Writes rendered text into emitter.
*/
func EngineRender(eng *Engine, emitter *Emitter, root Node) error {
	return engine.EngineRender(eng, emitter, root)
}

/*
EngineRenderNode renders a single node via direct registry slice lookup.

[Parameters]
ctx — render context with engine and emitter.
node — AST node to render.

[Returns]
An error when no renderer is registered or rendering fails.
*/
func EngineRenderNode(ctx *RenderContext, node Node) error {
	return engine.EngineRenderNode(ctx, node)
}

/*
EmitterCreate constructs an emitter that indents with a fixed number of spaces per level.

[Parameters]
indentWidth — spaces written per indent level.

[Returns]
A fresh emitter.
*/
func EmitterCreate(indentWidth int) Emitter {
	return emit.EmitterCreate(indentWidth)
}

/*
EmitterRender returns the accumulated text written to the emitter.

[Parameters]
emitter — emitter that received writes.

[Returns]
The full rendered string.
*/
func EmitterRender(emitter *Emitter) string {
	return emit.EmitterRender(emitter)
}

/*
FileElementFrom wraps any node as a FileElement.

[Parameters]
node — layout or declaration node from a language frontend.

[Returns]
A FileElement containing node.
*/
func FileElementFrom(node Node) FileElement {
	return ast.FileElementFrom(node)
}
