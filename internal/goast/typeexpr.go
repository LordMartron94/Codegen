package goast

import "strings"

/*
TypeExprKind identifies a type-expression variant.

[Context]
Type expressions describe Go type syntax without embedding raw type strings as syntax fragments beyond identifier spellings.
*/
type TypeExprKind uint8

const (
	TypeExprKindNamed TypeExprKind = iota
	TypeExprKindPointer
	TypeExprKindSlice
	TypeExprKindArray
	TypeExprKindFunc
)

/*
TypeExpr describes a Go type in generated code.

[Context]
Only NamedType uses a bare identifier string; PointerType and SliceType compose recursively.
*/
type TypeExpr struct {
	Kind TypeExprKind

	Named    string
	ArrayLen string
	Elem     *TypeExpr
	Params   []ParamType
	Returns  []TypeExpr
}

/*
ParamType pairs a parameter name with its type for function types.

[Context]
Used in struct fields and FuncDecl signatures that embed function types.
*/
type ParamType struct {
	Name string
	Type TypeExpr
}

func TypeExprNamed(name string) TypeExpr {
	return TypeExpr{Kind: TypeExprKindNamed, Named: name}
}

func TypeExprPointer(elem TypeExpr) TypeExpr {
	return TypeExpr{Kind: TypeExprKindPointer, Elem: &elem}
}

func TypeExprSlice(elem TypeExpr) TypeExpr {
	return TypeExpr{Kind: TypeExprKindSlice, Elem: &elem}
}

func TypeExprArray(length string, elem TypeExpr) TypeExpr {
	return TypeExpr{Kind: TypeExprKindArray, ArrayLen: length, Elem: &elem}
}

func TypeExprFunc(params []ParamType, returns []TypeExpr) TypeExpr {
	return TypeExpr{Kind: TypeExprKindFunc, Params: params, Returns: returns}
}

/*
TypeExprFromGoTypeString parses a Go type identifier string into a TypeExpr tree.

[Context]
Supports leading '*' and '[]' prefixes; the remainder is treated as a named type identifier. Used by domain generators that store types as strings in tables.

[Parameters]
typeString — Go type spelling such as "*uint32" or "VulkanLogicalDevice".

[Returns]
A TypeExpr representing the parsed type.

[Errors]
Returns an error when typeString is empty.
*/
func TypeExprFromGoTypeString(typeString string) (TypeExpr, error) {
	if typeString == "" {
		return TypeExpr{}, errEmptyTypeString
	}

	remaining := typeString
	if len(remaining) > 0 && remaining[0] == '[' {
		closeIndex := strings.Index(remaining, "]")
		if closeIndex <= 1 {
			return TypeExpr{}, errMalformedArrayType
		}
		length := remaining[1:closeIndex]
		inner, err := TypeExprFromGoTypeString(remaining[closeIndex+1:])
		if err != nil {
			return TypeExpr{}, err
		}
		return TypeExprArray(length, inner), nil
	}
	if len(remaining) >= 2 && remaining[:2] == "[]" {
		inner, err := TypeExprFromGoTypeString(remaining[2:])
		if err != nil {
			return TypeExpr{}, err
		}
		return TypeExprSlice(inner), nil
	}
	if remaining[0] == '*' {
		inner, err := TypeExprFromGoTypeString(remaining[1:])
		if err != nil {
			return TypeExpr{}, err
		}
		return TypeExprPointer(inner), nil
	}
	return TypeExprNamed(remaining), nil
}

var (
	errEmptyTypeString    = &typeExprError{message: "type string is empty"}
	errMalformedArrayType = &typeExprError{message: "malformed array type"}
)

type typeExprError struct {
	message string
}

func (e *typeExprError) Error() string {
	return e.message
}
