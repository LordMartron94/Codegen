/*
Package goast provides Go-specific AST node shapes for the codegen library.

[Context]
Type expressions, Go declarations, Go documentation layout, and var declarations live here. The language-agnostic engine in internal/engine dispatches by ast.NodeKind; internal/goengine registers renderers for Go kinds and uses types from this package.

[Boundary]
Does not perform rendering or registry management; see internal/goengine and internal/engine.
*/
package goast
