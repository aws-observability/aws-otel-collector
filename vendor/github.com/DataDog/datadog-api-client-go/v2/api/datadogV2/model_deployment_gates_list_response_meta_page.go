// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesListResponseMetaPage Pagination information for a list of deployment gates.
type DeploymentGatesListResponseMetaPage struct {
	// The cursor used for the current page.
	Cursor *string `json:"cursor,omitempty"`
	// The cursor to use to fetch the next page. This is absent when there are no more pages.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The number of results per page.
	Size *int64 `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesListResponseMetaPage instantiates a new DeploymentGatesListResponseMetaPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesListResponseMetaPage() *DeploymentGatesListResponseMetaPage {
	this := DeploymentGatesListResponseMetaPage{}
	var size int64 = 50
	this.Size = &size
	return &this
}

// NewDeploymentGatesListResponseMetaPageWithDefaults instantiates a new DeploymentGatesListResponseMetaPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesListResponseMetaPageWithDefaults() *DeploymentGatesListResponseMetaPage {
	this := DeploymentGatesListResponseMetaPage{}
	var size int64 = 50
	this.Size = &size
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *DeploymentGatesListResponseMetaPage) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesListResponseMetaPage) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *DeploymentGatesListResponseMetaPage) HasCursor() bool {
	return o != nil && o.Cursor != nil
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *DeploymentGatesListResponseMetaPage) SetCursor(v string) {
	o.Cursor = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *DeploymentGatesListResponseMetaPage) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesListResponseMetaPage) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *DeploymentGatesListResponseMetaPage) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *DeploymentGatesListResponseMetaPage) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DeploymentGatesListResponseMetaPage) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesListResponseMetaPage) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DeploymentGatesListResponseMetaPage) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *DeploymentGatesListResponseMetaPage) SetSize(v int64) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesListResponseMetaPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentGatesListResponseMetaPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cursor     *string `json:"cursor,omitempty"`
		NextCursor *string `json:"next_cursor,omitempty"`
		Size       *int64  `json:"size,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cursor", "next_cursor", "size"})
	} else {
		return err
	}
	o.Cursor = all.Cursor
	o.NextCursor = all.NextCursor
	o.Size = all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
