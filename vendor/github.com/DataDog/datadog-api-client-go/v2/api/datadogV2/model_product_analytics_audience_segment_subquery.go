// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsAudienceSegmentSubquery A segment-based audience query.
type ProductAnalyticsAudienceSegmentSubquery struct {
	// Name of this query, referenced in the formula.
	Name string `json:"name"`
	// UUID of the segment to filter by.
	SegmentId uuid.UUID `json:"segment_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsAudienceSegmentSubquery instantiates a new ProductAnalyticsAudienceSegmentSubquery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsAudienceSegmentSubquery(name string, segmentId uuid.UUID) *ProductAnalyticsAudienceSegmentSubquery {
	this := ProductAnalyticsAudienceSegmentSubquery{}
	this.Name = name
	this.SegmentId = segmentId
	return &this
}

// NewProductAnalyticsAudienceSegmentSubqueryWithDefaults instantiates a new ProductAnalyticsAudienceSegmentSubquery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsAudienceSegmentSubqueryWithDefaults() *ProductAnalyticsAudienceSegmentSubquery {
	this := ProductAnalyticsAudienceSegmentSubquery{}
	return &this
}

// GetName returns the Name field value.
func (o *ProductAnalyticsAudienceSegmentSubquery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAudienceSegmentSubquery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsAudienceSegmentSubquery) SetName(v string) {
	o.Name = v
}

// GetSegmentId returns the SegmentId field value.
func (o *ProductAnalyticsAudienceSegmentSubquery) GetSegmentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAudienceSegmentSubquery) GetSegmentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentId, true
}

// SetSegmentId sets field value.
func (o *ProductAnalyticsAudienceSegmentSubquery) SetSegmentId(v uuid.UUID) {
	o.SegmentId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsAudienceSegmentSubquery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["segment_id"] = o.SegmentId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsAudienceSegmentSubquery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string    `json:"name"`
		SegmentId *uuid.UUID `json:"segment_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SegmentId == nil {
		return fmt.Errorf("required field segment_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "segment_id"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.SegmentId = *all.SegmentId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
