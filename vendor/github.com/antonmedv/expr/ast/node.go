package ast

import (
	"reflect"
	"regexp"

<<<<<<< HEAD
	"github.com/antonmedv/expr/builtin"
=======
>>>>>>> main
	"github.com/antonmedv/expr/file"
)

// Node represents items of abstract syntax tree.
type Node interface {
	Location() file.Location
	SetLocation(file.Location)
	Type() reflect.Type
	SetType(reflect.Type)
<<<<<<< HEAD
=======
	String() string
>>>>>>> main
}

func Patch(node *Node, newNode Node) {
	newNode.SetType((*node).Type())
	newNode.SetLocation((*node).Location())
	*node = newNode
}

type base struct {
	loc      file.Location
	nodeType reflect.Type
}

func (n *base) Location() file.Location {
	return n.loc
}

func (n *base) SetLocation(loc file.Location) {
	n.loc = loc
}

func (n *base) Type() reflect.Type {
	return n.nodeType
}

func (n *base) SetType(t reflect.Type) {
	n.nodeType = t
}

type NilNode struct {
	base
}

type IdentifierNode struct {
	base
	Value       string
<<<<<<< HEAD
	Deref       bool
=======
>>>>>>> main
	FieldIndex  []int
	Method      bool // true if method, false if field
	MethodIndex int  // index of method, set only if Method is true
}

type IntegerNode struct {
	base
	Value int
}

type FloatNode struct {
	base
	Value float64
}

type BoolNode struct {
	base
	Value bool
}

type StringNode struct {
	base
	Value string
}

type ConstantNode struct {
	base
<<<<<<< HEAD
	Value interface{}
=======
	Value any
>>>>>>> main
}

type UnaryNode struct {
	base
	Operator string
	Node     Node
}

type BinaryNode struct {
	base
	Regexp   *regexp.Regexp
	Operator string
	Left     Node
	Right    Node
}

type ChainNode struct {
	base
	Node Node
}

type MemberNode struct {
	base
	Node       Node
	Property   Node
	Name       string // Name of the filed or method. Used for error reporting.
	Optional   bool
<<<<<<< HEAD
	Deref      bool
	FieldIndex []int

	// TODO: Replace with a single MethodIndex field of &int type.
=======
	FieldIndex []int

	// TODO: Combine Method and MethodIndex into a single MethodIndex field of &int type.
>>>>>>> main
	Method      bool
	MethodIndex int
}

type SliceNode struct {
	base
	Node Node
	From Node
	To   Node
}

type CallNode struct {
	base
	Callee    Node
	Arguments []Node
	Typed     int
	Fast      bool
<<<<<<< HEAD
	Func      *builtin.Function
=======
	Func      *Function
>>>>>>> main
}

type BuiltinNode struct {
	base
	Name      string
	Arguments []Node
<<<<<<< HEAD
=======
	Throws    bool
	Map       Node
>>>>>>> main
}

type ClosureNode struct {
	base
	Node Node
}

type PointerNode struct {
	base
<<<<<<< HEAD
=======
	Name string
>>>>>>> main
}

type ConditionalNode struct {
	base
	Cond Node
	Exp1 Node
	Exp2 Node
}

<<<<<<< HEAD
=======
type VariableDeclaratorNode struct {
	base
	Name  string
	Value Node
	Expr  Node
}

>>>>>>> main
type ArrayNode struct {
	base
	Nodes []Node
}

type MapNode struct {
	base
	Pairs []Node
}

type PairNode struct {
	base
	Key   Node
	Value Node
}
