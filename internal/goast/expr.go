package goast

import "codegen/internal/ast"

func (*IdentExpr) IsExpr()        {}
func (*SelectorExpr) IsExpr()     {}
func (*IndexExpr) IsExpr()        {}
func (*AddressOfExpr) IsExpr()    {}
func (*CallExpr) IsExpr()         {}
func (*StringLitExpr) IsExpr()    {}
func (*IntLitExpr) IsExpr()       {}
func (*NilLitExpr) IsExpr()       {}
func (*BinaryExpr) IsExpr()       {}
func (*CompositeLitExpr) IsExpr() {}

/*
BinaryOp identifies a binary expression operator in Go output.
*/
type BinaryOp uint8

const (
	BinaryOpEq BinaryOp = iota
	BinaryOpNe
)

/*
IdentExpr references a variable or package identifier by name.
*/
type IdentExpr struct {
	Name string
}

func (e *IdentExpr) NodeKind() ast.NodeKind { return KindIdentExpr }

/*
SelectorExpr selects a field or method on a base expression.
*/
type SelectorExpr struct {
	Base ast.Expr
	Name string
}

func (e *SelectorExpr) NodeKind() ast.NodeKind { return KindSelectorExpr }

/*
IndexExpr indexes a base expression with an index expression.
*/
type IndexExpr struct {
	Base  ast.Expr
	Index ast.Expr
}

func (e *IndexExpr) NodeKind() ast.NodeKind { return KindIndexExpr }

/*
AddressOfExpr takes the address of its operand expression.
*/
type AddressOfExpr struct {
	Operand ast.Expr
}

func (e *AddressOfExpr) NodeKind() ast.NodeKind { return KindAddressOfExpr }

/*
CallExpr invokes a callee with arguments.
*/
type CallExpr struct {
	Callee ast.Expr
	Args   []ast.Expr
}

func (e *CallExpr) NodeKind() ast.NodeKind { return KindCallExpr }

/*
StringLitExpr is a string literal value.
*/
type StringLitExpr struct {
	Value string
}

func (e *StringLitExpr) NodeKind() ast.NodeKind { return KindStringLitExpr }

/*
IntLitExpr is an integer literal value.
*/
type IntLitExpr struct {
	Value int64
}

func (e *IntLitExpr) NodeKind() ast.NodeKind { return KindIntLitExpr }

/*
NilLitExpr is the nil literal.
*/
type NilLitExpr struct{}

func (e *NilLitExpr) NodeKind() ast.NodeKind { return KindNilLitExpr }

/*
BinaryExpr combines two expressions with a binary operator.
*/
type BinaryExpr struct {
	Op    BinaryOp
	Left  ast.Expr
	Right ast.Expr
}

func (e *BinaryExpr) NodeKind() ast.NodeKind { return KindBinaryExpr }

/*
FieldInit pairs a field name with an initializing expression in a composite literal.
*/
type FieldInit struct {
	Name  string
	Value ast.Expr
}

/*
CompositeLitExpr is a composite literal for a named type or slice of composites.
*/
type CompositeLitExpr struct {
	TypeName string
	IsSlice  bool
	Fields   []FieldInit
	Elements []ast.Expr
}

func (e *CompositeLitExpr) NodeKind() ast.NodeKind { return KindCompositeLitExpr }

func ExprIdent(name string) *IdentExpr {
	return &IdentExpr{Name: name}
}

func ExprSelector(base ast.Expr, name string) *SelectorExpr {
	return &SelectorExpr{Base: base, Name: name}
}

func ExprIndex(base ast.Expr, index ast.Expr) *IndexExpr {
	return &IndexExpr{Base: base, Index: index}
}

func ExprAddressOf(operand ast.Expr) *AddressOfExpr {
	return &AddressOfExpr{Operand: operand}
}

func ExprCall(callee ast.Expr, args ...ast.Expr) *CallExpr {
	return &CallExpr{Callee: callee, Args: args}
}

func ExprStringLit(value string) *StringLitExpr {
	return &StringLitExpr{Value: value}
}

func ExprIntLit(value int64) *IntLitExpr {
	return &IntLitExpr{Value: value}
}

func ExprNil() *NilLitExpr {
	return &NilLitExpr{}
}

func ExprBinary(op BinaryOp, left ast.Expr, right ast.Expr) *BinaryExpr {
	return &BinaryExpr{Op: op, Left: left, Right: right}
}

func ExprCompositeLit(typeName string, fields []FieldInit) *CompositeLitExpr {
	return &CompositeLitExpr{TypeName: typeName, IsSlice: false, Fields: fields}
}

func ExprSliceCompositeLit(elementTypeName string, elements []ast.Expr) *CompositeLitExpr {
	return &CompositeLitExpr{TypeName: elementTypeName, IsSlice: true, Elements: elements}
}
