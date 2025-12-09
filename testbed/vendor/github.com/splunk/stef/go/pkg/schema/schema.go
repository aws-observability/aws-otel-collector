package schema

import (
	"errors"
	"fmt"
)

// Schema is a STEF schema description, serializable in JSON format.
type Schema struct {
	PackageName []string             `json:"package,omitempty"`
	Structs     map[string]*Struct   `json:"structs"`
	Multimaps   map[string]*Multimap `json:"multimaps"`
	Enums       map[string]*Enum
}

type Compatibility int

const (
	CompatibilityExact Compatibility = iota
	CompatibilitySuperset
	CompatibilityIncompatible
)

// Compatible checks backward compatibility of this schema with oldSchema.
// If the schemas are incompatible returns CompatibilityIncompatible and an error.
func (d *Schema) Compatible(oldSchema *Schema) (Compatibility, error) {
	// Exact compatibility is only possible if the number of structs is exactly the same.
	exact := len(d.Structs) == len(oldSchema.Structs)

	for name, oldStruc := range oldSchema.Structs {
		newStruc, ok := d.Structs[name]
		if !ok {
			return CompatibilityIncompatible,
				fmt.Errorf("struct %s does not exist in new schema", name)
		}
		comp, err := d.compatibleStruct(name, newStruc, oldStruc)
		if err != nil {
			return CompatibilityIncompatible, err
		}
		if comp == CompatibilitySuperset {
			exact = false
		}
	}

	for name, oldMap := range oldSchema.Multimaps {
		newMap, ok := d.Multimaps[name]
		if !ok {
			return CompatibilityIncompatible,
				fmt.Errorf("multimap %s does not exist in new schema", name)
		}
		comp, err := d.compatibleMultimap(name, newMap, oldMap)
		if err != nil {
			return CompatibilityIncompatible, err
		}
		if comp == CompatibilitySuperset {
			exact = false
		}
	}

	if exact {
		return CompatibilityExact, nil
	}

	return CompatibilitySuperset, nil
}

func (d *Schema) compatibleStruct(
	name string, newStruct *Struct, oldStruc *Struct,
) (Compatibility, error) {
	if len(newStruct.Fields) < len(oldStruc.Fields) {
		return CompatibilityIncompatible, fmt.Errorf("new struct %s has fewer fields than old struct", name)
	}

	if newStruct.OneOf != oldStruc.OneOf {
		return CompatibilityIncompatible, fmt.Errorf("new struct %s has different oneof flag than theold struct", name)
	}

	if newStruct.DictName != oldStruc.DictName {
		return CompatibilityIncompatible, fmt.Errorf(
			"new struct %s dictionary name is %s, old struct dictionary name is %s",
			name, newStruct.DictName, oldStruc.DictName,
		)
	}

	// Exact compatibility is only possible if the number of fields is exactly the same.
	exact := len(newStruct.Fields) == len(oldStruc.Fields)

	for i := range oldStruc.Fields {
		newField := newStruct.Fields[i]
		oldField := oldStruc.Fields[i]
		if err := isCompatibleField(name, i, newField, oldField); err != nil {
			return CompatibilityIncompatible, err
		}
	}

	if exact {
		return CompatibilityExact, nil
	}

	return CompatibilitySuperset, nil
}

func (d *Schema) compatibleMultimap(
	name string, newMap *Multimap, oldMap *Multimap,
) (Compatibility, error) {
	if !isCompatibleFieldType(&newMap.Key.Type, &oldMap.Key.Type) {
		return CompatibilityIncompatible,
			fmt.Errorf("multimap %s key type does not match", name)
	}
	if !isCompatibleFieldType(&newMap.Value.Type, &oldMap.Value.Type) {
		return CompatibilityIncompatible,
			fmt.Errorf("multimap %s value type does not match", name)
	}
	return CompatibilityExact, nil
}

func isCompatibleField(
	structName string, fieldIndex int, newField *StructField, oldField *StructField,
) error {
	if newField.Optional != oldField.Optional {
		return fmt.Errorf(
			"field %d in new struct %s has different optional flag than the old struct",
			fieldIndex, structName,
		)
	}

	if !isCompatibleFieldType(&newField.FieldType, &oldField.FieldType) {
		return fmt.Errorf(
			"field %d in new struct %s has a different type than the old struct",
			fieldIndex, structName,
		)
	}

	return nil
}

func isCompatibleFieldType(
	newField *FieldType, oldField *FieldType,
) bool {
	if (newField.Primitive == nil) != (oldField.Primitive == nil) {
		return false
	}

	if newField.Primitive != nil {
		if *newField.Primitive != *oldField.Primitive {
			return false
		}
	}

	if (newField.Array == nil) != (oldField.Array == nil) {
		return false
	}

	if newField.Array != nil {
		if !isCompatibleFieldType(&newField.Array.ElemType, &oldField.Array.ElemType) {
			return false
		}
	}

	if newField.Struct != oldField.Struct {
		return false
	}

	if newField.MultiMap != oldField.MultiMap {
		return false
	}

	if newField.DictName != oldField.DictName {
		return false
	}

	return true
}

// PrunedForRoot produces a pruned copy of the schema that includes the specified root
// struct and parts of schema reachable from that root. Unreachable parts of the schema
// are excluded.
func (d *Schema) PrunedForRoot(rootStructName string) (*Schema, error) {
	out := Schema{
		Structs:   map[string]*Struct{},
		Multimaps: map[string]*Multimap{},
		Enums:     map[string]*Enum{},
	}
	if err := d.copyPrunedStruct(rootStructName, &out); err != nil {
		return nil, err
	}

	err := out.ResolveRefs()
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// ResolveRefs resolves all type references by names to structs, multimaps,
// arrays and enums in the schema to pointers to appropriate objects in the schema.
func (d *Schema) ResolveRefs() error {
	for _, v := range d.Structs {
		for i := range v.Fields {
			field := v.Fields[i]
			if err := d.resolveFieldType(&field.FieldType); err != nil {
				return err
			}
		}
	}
	for _, v := range d.Multimaps {
		if err := d.resolveFieldType(&v.Key.Type); err != nil {
			return err
		}
		if err := d.resolveFieldType(&v.Value.Type); err != nil {
			return err
		}
	}

	return d.computeRecursive()
}

func (d *Schema) resolveFieldType(fieldType *FieldType) error {
	var typeName string
	if fieldType.Struct != "" {
		typeName = fieldType.Struct
	} else if fieldType.MultiMap != "" {
		typeName = fieldType.MultiMap
	} else if fieldType.Enum != "" {
		typeName = fieldType.Enum
	}

	if typeName != "" {
		matches := 0
		var isStruct bool
		fieldType.StructDef, isStruct = d.Structs[typeName]
		if isStruct {
			matches++
		}

		var isMultimap bool
		fieldType.MultimapDef, isMultimap = d.Multimaps[typeName]
		if isMultimap {
			fieldType.MultiMap = typeName
			fieldType.Struct = ""
			matches++
		}

		_, isEnum := d.Enums[typeName]
		if isEnum {
			// All enums are uint64.
			fieldType.Primitive = &PrimitiveType{Type: PrimitiveTypeUint64}
			fieldType.Enum = typeName
			fieldType.Struct = ""
			matches++
		}

		if matches == 0 {
			return errors.New("unknown type: " + typeName)
		}
		if matches > 1 {
			return errors.New("ambiguous type: " + typeName)
		}
		return nil
	}

	if fieldType.Array != nil {
		if err := d.resolveFieldType(&fieldType.Array.ElemType); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (d *Schema) copyPrunedFieldType(fieldType *FieldType, dst *Schema) error {
	if fieldType.Struct != "" {
		if err := d.copyPrunedStruct(fieldType.Struct, dst); err != nil {
			return err
		}
	} else if fieldType.MultiMap != "" {
		if err := d.copyPrunedMultiMap(fieldType.MultiMap, dst); err != nil {
			return err
		}
	} else if fieldType.Array != nil {
		if err := d.copyPrunedFieldType(&fieldType.Array.ElemType, dst); err != nil {
			return err
		}
	} else if fieldType.Enum != "" {
		if err := d.copyPrunedEnum(fieldType.Enum, dst); err != nil {
			return err
		}
	}
	return nil
}

func (d *Schema) copyPrunedStruct(strucName string, dst *Schema) error {
	if dst.Structs[strucName] != nil {
		// already copied
		return nil
	}

	srcStruc := d.Structs[strucName]
	if srcStruc == nil {
		return fmt.Errorf("no struct named %s found", strucName)
	}

	dstStruc := &Struct{
		Name:     strucName,
		OneOf:    srcStruc.OneOf,
		DictName: srcStruc.DictName,
		IsRoot:   srcStruc.IsRoot,
		Fields:   make([]*StructField, len(srcStruc.Fields)),
	}
	dst.Structs[strucName] = dstStruc

	for i := range srcStruc.Fields {
		dstStruc.Fields[i] = srcStruc.Fields[i].Clone()
		if err := d.copyPrunedFieldType(&dstStruc.Fields[i].FieldType, dst); err != nil {
			return err
		}
	}

	return nil
}

func (d *Schema) copyPrunedMultiMap(multiMapName string, dst *Schema) error {
	if dst.Multimaps[multiMapName] != nil {
		// already copied
		return nil
	}

	srcMultiMap := d.Multimaps[multiMapName]
	if srcMultiMap == nil {
		return fmt.Errorf("no multimap named %s found", multiMapName)
	}

	dstMultimap := &Multimap{
		Name:  multiMapName,
		Key:   srcMultiMap.Key.Clone(),
		Value: srcMultiMap.Value.Clone(),
	}
	dst.Multimaps[multiMapName] = dstMultimap

	if err := d.copyPrunedFieldType(&dstMultimap.Key.Type, dst); err != nil {
		return err
	}

	if err := d.copyPrunedFieldType(&dstMultimap.Value.Type, dst); err != nil {
		return err
	}

	return nil
}

func (d *Schema) copyPrunedEnum(enumName string, dst *Schema) error {
	if dst.Enums[enumName] != nil {
		// already copied
		return nil
	}

	srcEnum := d.Enums[enumName]
	if srcEnum == nil {
		return fmt.Errorf("no enum named %s found", enumName)
	}

	dstEnum := &Enum{
		Name: enumName,
	}
	for i := range srcEnum.Fields {
		dstEnum.Fields = append(dstEnum.Fields, srcEnum.Fields[i])
	}

	dst.Enums[enumName] = dstEnum

	return nil
}

// FieldCount returns the number of fields in the specified struct.
func (d *Schema) FieldCount(structName string) (int, error) {
	struc := d.Structs[structName]
	if struc == nil {
		return 0, fmt.Errorf("struct %s not found", structName)
	}
	return len(struc.Fields), nil
}

func (d *Schema) computeRecursive() error {
	for _, struc := range d.Structs {
		if !struc.IsRoot {
			continue
		}
		stack := recurseStack{asMap: map[string]bool{}}
		computeRecursiveStruct(struc, &stack)
	}
	return nil
}

func markRecursive(typeName string, stack *recurseStack) {
	startIdx := findLast(stack.asStack, typeName)
	if startIdx == -1 {
		panic("invalid state")
	}
	for i := startIdx; i < len(stack.fields); i++ {
		stack.fields[i].SetRecursive()
	}
}

func findLast(stack []string, name string) int {
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == name {
			return i
		}
	}
	return -1
}

func computeRecursiveStruct(struc *Struct, stack *recurseStack) {
	stack.asStack = append(stack.asStack, struc.Name)
	stack.asMap[struc.Name] = true

	for _, field := range struc.Fields {
		stack.fields = append(stack.fields, field)
		computeRecursiveType(field.FieldType, stack)
		stack.fields = stack.fields[:len(stack.fields)-1]
	}

	stack.asStack = stack.asStack[:len(stack.asStack)-1]
	delete(stack.asMap, struc.Name)
}

func computeRecursiveMultimap(multimap *Multimap, stack *recurseStack) {
	stack.asStack = append(stack.asStack, multimap.Name)
	stack.asMap[multimap.Name] = true

	stack.fields = append(stack.fields, &multimap.Key)
	computeRecursiveType(multimap.Key.Type, stack)
	stack.fields = stack.fields[:len(stack.fields)-1]

	stack.fields = append(stack.fields, &multimap.Value)
	computeRecursiveType(multimap.Value.Type, stack)
	stack.fields = stack.fields[:len(stack.fields)-1]

	stack.asStack = stack.asStack[:len(stack.asStack)-1]
	delete(stack.asMap, multimap.Name)
}

func computeRecursiveType(typ FieldType, stack *recurseStack) {
	if typ.Primitive != nil {
		// Primitive types are not recursive.
		return
	}
	if typ.Struct != "" {
		if stack.asMap[typ.Struct] {
			markRecursive(typ.Struct, stack)
		} else {
			computeRecursiveStruct(typ.StructDef, stack)
		}
		return
	}
	if typ.MultiMap != "" {
		if stack.asMap[typ.MultiMap] {
			markRecursive(typ.MultiMap, stack)
		} else {
			computeRecursiveMultimap(typ.MultimapDef, stack)
		}
		return
	}
	if typ.Array != nil {
		computeRecursiveType(typ.Array.ElemType, stack)
		return
	}

	panic("unknown type")
}

type Struct struct {
	Name     string         `json:"name,omitempty"`
	OneOf    bool           `json:"oneof,omitempty"`
	DictName string         `json:"dict,omitempty"`
	IsRoot   bool           `json:"root,omitempty"`
	Fields   []*StructField `json:"fields"`

	// recursive is true if this struct type definition is self-recursive
	// or mutually recursive.
	recursive bool
}

func (s Struct) Recursive() bool {
	return s.recursive
}

type StructField struct {
	FieldType
	Name     string `json:"name,omitempty"`
	Optional bool   `json:"optional,omitempty"`
}

func (s *StructField) Clone() *StructField {
	clone := &StructField{
		FieldType: s.FieldType.Clone(),
		Name:      s.Name,
		Optional:  s.Optional,
	}
	return clone
}

func (s *StructField) SetRecursive() {
	s.FieldType.SetRecursive()
}

func (s *StructField) Recursive() bool {
	return s.FieldType.Recursive()
}

type PrimitiveFieldType int

const (
	PrimitiveTypeInt64 PrimitiveFieldType = iota
	PrimitiveTypeUint64
	PrimitiveTypeFloat64
	PrimitiveTypeBool
	PrimitiveTypeString
	PrimitiveTypeBytes
)

type PrimitiveType struct {
	Type PrimitiveFieldType
}

type ArrayType struct {
	ElemType FieldType `json:"elem,omitempty"`

	// recursive is true if this array type definition is self-recursive
	// or mutually recursive.
	recursive bool
}

// FieldType describes a field type of in a struct, multimap key or value or array element.
type FieldType struct {
	// Only one of the following 5 fields should be set.
	Primitive *PrimitiveType `json:"primitive,omitempty"`
	Array     *ArrayType     `json:"array,omitempty"`
	Struct    string         `json:"struct,omitempty"`
	MultiMap  string         `json:"multimap,omitempty"`
	Enum      string

	// DictName is the name of the dictionary if the field dictionary-encoded.
	DictName string `json:"dict,omitempty"`

	// StructDef and MultimapDef are pointers to the struct or multimap definitions
	// that are named in Struct or MultiMap fields.
	StructDef   *Struct
	MultimapDef *Multimap
}

func (t *FieldType) Clone() FieldType {
	clone := FieldType{
		Primitive: t.Primitive,
		Array:     t.Array,
		Struct:    t.Struct,
		MultiMap:  t.MultiMap,
		Enum:      t.Enum,
		DictName:  t.DictName,
	}
	if t.Array != nil {
		clone.Array = &ArrayType{ElemType: t.Array.ElemType.Clone()}
	}
	return clone
}

func (t *FieldType) SetRecursive() {
	switch {
	case t.Primitive != nil:
		panic("cannot set recursive on Primitive")
	case t.Array != nil:
		t.Array.recursive = true
	case t.StructDef != nil:
		t.StructDef.recursive = true
	case t.MultimapDef != nil:
		t.MultimapDef.recursive = true
	default:
		panic("invalid FieldType")
	}
}

func (t *FieldType) Recursive() bool {
	switch {
	case t.Primitive != nil:
		return false
	case t.Array != nil:
		return t.Array.recursive
	case t.StructDef != nil:
		return t.StructDef.recursive
	case t.MultimapDef != nil:
		return t.MultimapDef.recursive
	default:
		panic("invalid FieldType")
	}
}

// MultimapField is either a key or a value field in a multimap.
type MultimapField struct {
	Type FieldType `json:"type"`
}

func (m *MultimapField) SetRecursive() {
	m.Type.SetRecursive()
}

func (m *MultimapField) Clone() MultimapField {
	return MultimapField{
		Type: m.Type.Clone(),
	}
}

type Multimap struct {
	Name  string        `json:"name,omitempty"`
	Key   MultimapField `json:"key"`
	Value MultimapField `json:"value"`

	// recursive is true if this multimap type definition is self-recursive
	// or mutually recursive.
	recursive bool
}

type Enum struct {
	Name   string
	Fields []EnumField
}

type EnumField struct {
	Name  string
	Value uint64
}
