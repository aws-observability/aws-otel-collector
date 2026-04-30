// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JsonPatchOperationOp The operation to perform.
type JsonPatchOperationOp string

// List of JsonPatchOperationOp.
const (
	JSONPATCHOPERATIONOP_ADD     JsonPatchOperationOp = "add"
	JSONPATCHOPERATIONOP_REMOVE  JsonPatchOperationOp = "remove"
	JSONPATCHOPERATIONOP_REPLACE JsonPatchOperationOp = "replace"
	JSONPATCHOPERATIONOP_MOVE    JsonPatchOperationOp = "move"
	JSONPATCHOPERATIONOP_COPY    JsonPatchOperationOp = "copy"
	JSONPATCHOPERATIONOP_TEST    JsonPatchOperationOp = "test"
)

var allowedJsonPatchOperationOpEnumValues = []JsonPatchOperationOp{
	JSONPATCHOPERATIONOP_ADD,
	JSONPATCHOPERATIONOP_REMOVE,
	JSONPATCHOPERATIONOP_REPLACE,
	JSONPATCHOPERATIONOP_MOVE,
	JSONPATCHOPERATIONOP_COPY,
	JSONPATCHOPERATIONOP_TEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *JsonPatchOperationOp) GetAllowedValues() []JsonPatchOperationOp {
	return allowedJsonPatchOperationOpEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *JsonPatchOperationOp) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = JsonPatchOperationOp(value)
	return nil
}

// NewJsonPatchOperationOpFromValue returns a pointer to a valid JsonPatchOperationOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewJsonPatchOperationOpFromValue(v string) (*JsonPatchOperationOp, error) {
	ev := JsonPatchOperationOp(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for JsonPatchOperationOp: valid values are %v", v, allowedJsonPatchOperationOpEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v JsonPatchOperationOp) IsValid() bool {
	for _, existing := range allowedJsonPatchOperationOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JsonPatchOperationOp value.
func (v JsonPatchOperationOp) Ptr() *JsonPatchOperationOp {
	return &v
}
