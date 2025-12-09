package schema

// structCountTree is a tree structure that stores the number of fields of all struct/oneof
// in a schema. The tree's hierarchy corresponds to the hierarchy of structs/oneof in the schema.
// Non-struct/oneof types are omitted in the tree, i.e. the children of such types are
// parented by the nearest parent struct/oneof.
type structCountTree struct {
	// The name of the struct/oneof that this node represents.
	structName string

	// The number of fields in the struct/oneof.
	fieldCount uint

	// The children of this node, which are structs/oneofs that are reachable from this node.
	structFields []structCountTree
}

// schemaToStructCountTree recursively traverses the schema and builds a tree of struct counts.
// The tree is built in a depth-first traversal order, where each node represents a struct/oneof
// and its children represent the structs/oneofs that are reachable from it.
// The stack is used to detect and avoid infinite recursion in case of recursive types.
func schemaToStructCountTree(src *FieldType, dst *structCountTree, stack *recurseStack) {
	switch {
	case src.Primitive != nil:
		return

	case src.StructDef != nil:
		if stack.asMap[src.StructDef.Name] {
			// We have seen this struct before, no need to record its field count again.
			return
		}

		dst.structName = src.StructDef.Name
		dst.fieldCount = uint(len(src.StructDef.Fields))

		stack.asStack = append(stack.asStack, src.StructDef.Name)
		stack.asMap[src.StructDef.Name] = true

		for _, field := range src.StructDef.Fields {
			subDst := structCountTree{}
			schemaToStructCountTree(&field.FieldType, &subDst, stack)
			if subDst.structName != "" {
				// There is a struct in the children nodes.
				dst.structFields = append(dst.structFields, subDst)
			}
		}

		stack.asStack = stack.asStack[:len(stack.asStack)-1]

	case src.Array != nil:
		// Ignore the array and continue with the element type.
		schemaToStructCountTree(&src.Array.ElemType, dst, stack)

	case src.MultimapDef != nil:
		if stack.asMap[src.MultimapDef.Name] {
			return
		}

		stack.asStack = append(stack.asStack, src.MultimapDef.Name)
		stack.asMap[src.MultimapDef.Name] = true

		// Continue with both key and value types.
		schemaToStructCountTree(&src.MultimapDef.Key.Type, dst, stack)
		schemaToStructCountTree(&src.MultimapDef.Value.Type, dst, stack)

		stack.asStack = stack.asStack[:len(stack.asStack)-1]
		delete(stack.asMap, src.MultimapDef.Name)

	default:
		panic("unknown FieldType")
	}
}
