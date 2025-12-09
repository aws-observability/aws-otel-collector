package schema

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/splunk/stef/go/pkg/internal"
)

// WireSchema caries only the parts of the schema, which are necessary to be
// communicated between readers and writers that work with evolving versions
// of the same schema.
//
// WireSchema allows readers and writers to perform compatibility checks
// of their schema version with the schema version that a peer they communicate
// with uses.
//
// The only valid way to evolve a STEF schema is by adding new fields at the end
// of the existing structs/oneofs. This means that in order to correctly read/write an
// evolved schema the only necessary information is the number of the fields in
// in each struct/oneof. This is precisely the information that is recorded in WireSchema.
//
// The full schema information can be recorded in a schema.Schema, however that
// full information is not necessary for wire compatibility checks. Instead, we
// use the much simpler and more compact WireSchema to serve that role.
type WireSchema struct {
	// structCounts is a linearized list of the number of fields in each struct/oneof
	// in the schema. The order of the structs in this list is the same as the order
	// in which the structs are encoded/decoded in the schema.
	structCounts []uint
}

// NewWireSchema creates a new WireSchema from a schema for the given root.
// The wire schema will only include the structs that are reachable
// from the specified root struct.
func NewWireSchema(schema *Schema, rootStructName string) WireSchema {
	// Compose the struct counts tree from the schema, for the given root struct.
	rootStruc := schema.Structs[rootStructName]
	stack := recurseStack{asMap: map[string]bool{}}

	rootType := &FieldType{
		Struct:    rootStruc.Name,
		StructDef: rootStruc,
		DictName:  rootStruc.DictName,
	}
	tree := structCountTree{}

	schemaToStructCountTree(rootType, &tree, &stack)

	w := WireSchema{}

	// setStructCountsFromTree recursively linearizes the structCounts from the tree,
	// in depth-first traversal order. This traversal order matches the order in which
	// structs are encoded/decoded in the schema. This ordering ensures that when
	// encoder/decoder ask for struct counts via WireSchemaIter, the counts are returned
	// in the same order as we linearize them in structCounts field.
	w.setStructCountsFromTree(&tree)

	return w
}

func (w *WireSchema) setStructCountsFromTree(s *structCountTree) {
	w.structCounts = append(w.structCounts, s.fieldCount)
	for i := range s.structFields {
		w.setStructCountsFromTree(&s.structFields[i])
	}
}

const (
	maxStructCount = 1024
)

var (
	errStructCountLimit = errors.New("struct count limit exceeded")
)

// Serialize the schema to binary format.
func (w *WireSchema) Serialize(dst *bytes.Buffer) error {
	if err := internal.WriteUvarint(uint64(len(w.structCounts)), dst); err != nil {
		return err
	}
	for _, count := range w.structCounts {
		if err := internal.WriteUvarint(uint64(count), dst); err != nil {
			return err
		}
	}

	return nil
}

// Deserialize the schema from binary format.
func (w *WireSchema) Deserialize(src io.ByteReader) error {
	count, err := binary.ReadUvarint(src)
	if err != nil {
		return err
	}

	if count > maxStructCount {
		return errStructCountLimit
	}

	w.structCounts = make([]uint, int(count))

	for i := 0; i < int(count); i++ {
		fieldCount, err := binary.ReadUvarint(src)
		if err != nil {
			return err
		}

		w.structCounts[i] = uint(fieldCount)
	}

	return nil
}

// Compatible checks backward compatibility of this schema with oldSchema.
// If the schemas are incompatible returns CompatibilityIncompatible and an error.
func (w *WireSchema) Compatible(oldSchema *WireSchema) (Compatibility, error) {
	if len(w.structCounts) > len(oldSchema.structCounts) {
		return CompatibilitySuperset, nil
	}

	if len(w.structCounts) < len(oldSchema.structCounts) {
		return CompatibilityIncompatible,
			fmt.Errorf(
				"new schema has fewers structs than old schema (%d vs %d)", len(w.structCounts),
				len(oldSchema.structCounts),
			)
	}

	newFieldTotal := uint(0)
	oldFieldTotal := uint(0)
	for i := range w.structCounts {
		newFieldTotal += w.structCounts[i]
		oldFieldTotal += oldSchema.structCounts[i]
	}

	if newFieldTotal > oldFieldTotal {
		return CompatibilitySuperset, nil
	}

	if newFieldTotal < oldFieldTotal {
		return CompatibilityIncompatible,
			fmt.Errorf(
				"new schema has fewers fields than old schema (%d vs %d)", newFieldTotal, oldFieldTotal,
			)
	}

	return CompatibilityExact, nil
}

// WireSchemaIter is an iterator over the structs in a WireSchema.
type WireSchemaIter struct {
	schema    *WireSchema
	structIdx int
}

func NewWireSchemaIter(schema *WireSchema) WireSchemaIter {
	return WireSchemaIter{
		schema:    schema,
		structIdx: 0,
	}
}

// NextFieldCount returns the field count for the next struct in the schema.
func (i *WireSchemaIter) NextFieldCount() (fieldCount uint, err error) {
	if i.structIdx >= len(i.schema.structCounts) {
		return 0, errors.New("struct count limit exceeded")
	}

	count := i.schema.structCounts[i.structIdx]
	i.structIdx++

	return count, nil
}

func (i *WireSchemaIter) Done() bool {
	return i.structIdx >= len(i.schema.structCounts)
}
