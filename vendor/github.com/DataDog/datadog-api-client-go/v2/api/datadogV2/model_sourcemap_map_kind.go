// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapMapKind The type of source map.
type SourcemapMapKind string

// List of SourcemapMapKind.
const (
	SOURCEMAPMAPKIND_JS      SourcemapMapKind = "js"
	SOURCEMAPMAPKIND_JVM     SourcemapMapKind = "jvm"
	SOURCEMAPMAPKIND_IOS     SourcemapMapKind = "ios"
	SOURCEMAPMAPKIND_REACT   SourcemapMapKind = "react"
	SOURCEMAPMAPKIND_FLUTTER SourcemapMapKind = "flutter"
	SOURCEMAPMAPKIND_ELF     SourcemapMapKind = "elf"
	SOURCEMAPMAPKIND_NDK     SourcemapMapKind = "ndk"
	SOURCEMAPMAPKIND_IL2CPP  SourcemapMapKind = "il2cpp"
)

var allowedSourcemapMapKindEnumValues = []SourcemapMapKind{
	SOURCEMAPMAPKIND_JS,
	SOURCEMAPMAPKIND_JVM,
	SOURCEMAPMAPKIND_IOS,
	SOURCEMAPMAPKIND_REACT,
	SOURCEMAPMAPKIND_FLUTTER,
	SOURCEMAPMAPKIND_ELF,
	SOURCEMAPMAPKIND_NDK,
	SOURCEMAPMAPKIND_IL2CPP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SourcemapMapKind) GetAllowedValues() []SourcemapMapKind {
	return allowedSourcemapMapKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SourcemapMapKind) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SourcemapMapKind(value)
	return nil
}

// NewSourcemapMapKindFromValue returns a pointer to a valid SourcemapMapKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSourcemapMapKindFromValue(v string) (*SourcemapMapKind, error) {
	ev := SourcemapMapKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SourcemapMapKind: valid values are %v", v, allowedSourcemapMapKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SourcemapMapKind) IsValid() bool {
	for _, existing := range allowedSourcemapMapKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SourcemapMapKind value.
func (v SourcemapMapKind) Ptr() *SourcemapMapKind {
	return &v
}
