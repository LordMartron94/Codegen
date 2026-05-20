package ast

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
BinaryOp identifies a binary expression operator.

[Context]
Go engines map operators to Go token spellings.
*/
type BinaryOp uint8

const (
	BinaryOpEq BinaryOp = iota
	BinaryOpNe
)

/*
IdentExpr references a variable or package identifier by name.

[Context]
Renders as a bare identifier in Go output.
*/
type IdentExpr struct {
	Name string
}

func (e *IdentExpr) NodeKind() NodeKind { return NodeKindIdentExpr }

/*
SelectorExpr selects a field or method on a base expression.

[Context]
Renders as base.selector in Go output.
*/
type SelectorExpr struct {
	Base Expr
	Name string
}

func (e *SelectorExpr) NodeKind() NodeKind { return NodeKindSelectorExpr }

/*
IndexExpr indexes a base expression with an index expression.

[Context]
Renders as base[index] in Go output.
*/
type IndexExpr struct {
	Base  Expr
	Index Expr
}

func (e *IndexExpr) NodeKind() NodeKind { return NodeKindIndexExpr }

/*
AddressOfExpr takes the address of its operand expression.

[Context]
Renders as &operand in Go output.
*/
type AddressOfExpr struct {
	Operand Expr
}

func (e *AddressOfExpr) NodeKind() NodeKind { return NodeKindAddressOfExpr }

/*
CallExpr invokes a callee with arguments.

[Context]
Callee is typically IdentExpr or SelectorExpr; arguments are Expr nodes.
*/
type CallExpr struct {
	Callee Expr
	Args   []Expr
}

func (e *CallExpr) NodeKind() NodeKind { return NodeKindCallExpr }

/*
StringLitExpr is a string literal value.

[Context]
Renders as a quoted Go string with proper escaping.
*/
type StringLitExpr struct {
	Value string
}

func (e *StringLitExpr) NodeKind() NodeKind { return NodeKindStringLitExpr }

/*
IntLitExpr is an integer literal value.

[Context]
Renders as a decimal integer in Go output.
*/
type IntLitExpr struct {
	Value int64
}

func (e *IntLitExpr) NodeKind() NodeKind { return NodeKindIntLitExpr }

/*
NilLitExpr is the nil literal.

[Context]
Renders as nil in Go output.
*/
type NilLitExpr struct{}

func (e *NilLitExpr) NodeKind() NodeKind { return NodeKindNilLitExpr }

/*
BinaryExpr combines two expressions with a binary operator.

[Context]
Used for conditions such as addr == 0 or err != nil.
*/
type BinaryExpr struct {
	Op    BinaryOp
	Left  Expr
	Right Expr
}

func (e *BinaryExpr) NodeKind() NodeKind { return NodeKindBinaryExpr }

/*
FieldInit pairs a field name with an initializing expression in a composite literal.

[Context]
Used by CompositeLitExpr for struct literal fields.
*/
type FieldInit struct {
	Name  string
	Value Expr
}

/*
CompositeLitExpr is a composite literal for a named type or slice of composites.

[Context]
When IsSlice is false, Fields supplies struct field initializers. When IsSlice is true, Elements holds each composite entry (typically nested CompositeLitExpr values).
*/
type CompositeLitExpr struct {
	TypeName string
	IsSlice  bool
	Fields   []FieldInit
	Elements []Expr
}

func (e *CompositeLitExpr) NodeKind() NodeKind { return NodeKindCompositeLitExpr }

func ExprIdent(name string) *IdentExpr {
	return &IdentExpr{Name: name}
}

func ExprSelector(base Expr, name string) *SelectorExpr {
	return &SelectorExpr{Base: base, Name: name}
}

func ExprIndex(base Expr, index Expr) *IndexExpr {
	return &IndexExpr{Base: base, Index: index}
}

func ExprAddressOf(operand Expr) *AddressOfExpr {
	return &AddressOfExpr{Operand: operand}
}

func ExprCall(callee Expr, args ...Expr) *CallExpr {
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

func ExprBinary(op BinaryOp, left Expr, right Expr) *BinaryExpr {
	return &BinaryExpr{Op: op, Left: left, Right: right}
}

func ExprCompositeLit(typeName string, fields []FieldInit) *CompositeLitExpr {
	return &CompositeLitExpr{TypeName: typeName, IsSlice: false, Fields: fields}
}

func ExprSliceCompositeLit(elementTypeName string, elements []Expr) *CompositeLitExpr {
	return &CompositeLitExpr{TypeName: elementTypeName, IsSlice: true, Elements: elements}
}
