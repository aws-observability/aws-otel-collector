// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionQueryWithRelationshipsResponse Response containing information about a single restriction query.
type RestrictionQueryWithRelationshipsResponse struct {
	// Restriction query object returned by the API.
	Data *RestrictionQueryWithRelationships `json:"data,omitempty"`
	// Array of objects related to the restriction query.
	Included []RestrictionQueryResponseIncludedItem `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestrictionQueryWithRelationshipsResponse instantiates a new RestrictionQueryWithRelationshipsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestrictionQueryWithRelationshipsResponse() *RestrictionQueryWithRelationshipsResponse {
	this := RestrictionQueryWithRelationshipsResponse{}
	return &this
}

// NewRestrictionQueryWithRelationshipsResponseWithDefaults instantiates a new RestrictionQueryWithRelationshipsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestrictionQueryWithRelationshipsResponseWithDefaults() *RestrictionQueryWithRelationshipsResponse {
	this := RestrictionQueryWithRelationshipsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RestrictionQueryWithRelationshipsResponse) GetData() RestrictionQueryWithRelationships {
	if o == nil || o.Data == nil {
		var ret RestrictionQueryWithRelationships
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryWithRelationshipsResponse) GetDataOk() (*RestrictionQueryWithRelationships, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RestrictionQueryWithRelationshipsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given RestrictionQueryWithRelationships and assigns it to the Data field.
func (o *RestrictionQueryWithRelationshipsResponse) SetData(v RestrictionQueryWithRelationships) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *RestrictionQueryWithRelationshipsResponse) GetIncluded() []RestrictionQueryResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []RestrictionQueryResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictionQueryWithRelationshipsResponse) GetIncludedOk() (*[]RestrictionQueryResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *RestrictionQueryWithRelationshipsResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []RestrictionQueryResponseIncludedItem and assigns it to the Included field.
func (o *RestrictionQueryWithRelationshipsResponse) SetIncluded(v []RestrictionQueryResponseIncludedItem) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestrictionQueryWithRelationshipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestrictionQueryWithRelationshipsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *RestrictionQueryWithRelationships     `json:"data,omitempty"`
		Included []RestrictionQueryResponseIncludedItem `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
