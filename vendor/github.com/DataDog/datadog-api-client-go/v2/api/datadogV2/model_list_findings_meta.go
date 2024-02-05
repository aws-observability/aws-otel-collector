// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListFindingsMeta Metadata for pagination.
type ListFindingsMeta struct {
	// Pagination and findings count information.
	Page *ListFindingsPage `json:"page,omitempty"`
	// The point in time corresponding to the listed findings.
	SnapshotTimestamp *int64 `json:"snapshot_timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewListFindingsMeta instantiates a new ListFindingsMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListFindingsMeta() *ListFindingsMeta {
	this := ListFindingsMeta{}
	return &this
}

// NewListFindingsMetaWithDefaults instantiates a new ListFindingsMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListFindingsMetaWithDefaults() *ListFindingsMeta {
	this := ListFindingsMeta{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ListFindingsMeta) GetPage() ListFindingsPage {
	if o == nil || o.Page == nil {
		var ret ListFindingsPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFindingsMeta) GetPageOk() (*ListFindingsPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ListFindingsMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given ListFindingsPage and assigns it to the Page field.
func (o *ListFindingsMeta) SetPage(v ListFindingsPage) {
	o.Page = &v
}

// GetSnapshotTimestamp returns the SnapshotTimestamp field value if set, zero value otherwise.
func (o *ListFindingsMeta) GetSnapshotTimestamp() int64 {
	if o == nil || o.SnapshotTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotTimestamp
}

// GetSnapshotTimestampOk returns a tuple with the SnapshotTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFindingsMeta) GetSnapshotTimestampOk() (*int64, bool) {
	if o == nil || o.SnapshotTimestamp == nil {
		return nil, false
	}
	return o.SnapshotTimestamp, true
}

// HasSnapshotTimestamp returns a boolean if a field has been set.
func (o *ListFindingsMeta) HasSnapshotTimestamp() bool {
	return o != nil && o.SnapshotTimestamp != nil
}

// SetSnapshotTimestamp gets a reference to the given int64 and assigns it to the SnapshotTimestamp field.
func (o *ListFindingsMeta) SetSnapshotTimestamp(v int64) {
	o.SnapshotTimestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListFindingsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.SnapshotTimestamp != nil {
		toSerialize["snapshot_timestamp"] = o.SnapshotTimestamp
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListFindingsMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page              *ListFindingsPage `json:"page,omitempty"`
		SnapshotTimestamp *int64            `json:"snapshot_timestamp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	o.SnapshotTimestamp = all.SnapshotTimestamp

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
