package schema

import (
	"math/rand/v2"
	"sort"
)

// ShrinkRandomly will deterministically, pseudo-randomly shrink the provided
// schema by removing the last field from one of the structs or oneofs.
// This function is used for testing schema changes.
func ShrinkRandomly(r *rand.Rand, schem *Schema) {
	var structNames []string
	for structName := range schem.Structs {
		structNames = append(structNames, structName)
	}
	sort.Strings(structNames)

	for {
		structName := structNames[r.IntN(len(structNames))]
		str := schem.Structs[structName]
		if shrinkStruct(r, schem, str) {
			return
		}
	}
}

func shrinkStruct(r *rand.Rand, schem *Schema, str *Struct) bool {
	if r.IntN(10) == 0 && len(str.Fields) > 0 {
		str.Fields = str.Fields[0 : len(str.Fields)-1]
		return true
	}

	for _, field := range str.Fields {
		if field.Struct != "" {
			if r.IntN(3) == 0 {
				childStruct := schem.Structs[field.Struct]
				changed := shrinkStruct(r, schem, childStruct)
				if changed {
					return true
				}
			}
		}
	}

	return false
}
