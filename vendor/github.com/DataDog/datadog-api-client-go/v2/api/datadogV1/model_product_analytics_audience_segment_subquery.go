// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsAudienceSegmentSubquery Product Analytics audience segment subquery.
type ProductAnalyticsAudienceSegmentSubquery struct {
	// The name of the segment subquery.
	Name *string `json:"name,omitempty"`
	// The unique identifier of the segment.
	SegmentId *string `json:"segment_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsAudienceSegmentSubquery instantiates a new ProductAnalyticsAudienceSegmentSubquery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsAudienceSegmentSubquery() *ProductAnalyticsAudienceSegmentSubquery {
	this := ProductAnalyticsAudienceSegmentSubquery{}
	return &this
}

// NewProductAnalyticsAudienceSegmentSubqueryWithDefaults instantiates a new ProductAnalyticsAudienceSegmentSubquery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsAudienceSegmentSubqueryWithDefaults() *ProductAnalyticsAudienceSegmentSubquery {
	this := ProductAnalyticsAudienceSegmentSubquery{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsAudienceSegmentSubquery) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAudienceSegmentSubquery) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsAudienceSegmentSubquery) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsAudienceSegmentSubquery) SetName(v string) {
	o.Name = &v
}

// GetSegmentId returns the SegmentId field value if set, zero value otherwise.
func (o *ProductAnalyticsAudienceSegmentSubquery) GetSegmentId() string {
	if o == nil || o.SegmentId == nil {
		var ret string
		return ret
	}
	return *o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsAudienceSegmentSubquery) GetSegmentIdOk() (*string, bool) {
	if o == nil || o.SegmentId == nil {
		return nil, false
	}
	return o.SegmentId, true
}

// HasSegmentId returns a boolean if a field has been set.
func (o *ProductAnalyticsAudienceSegmentSubquery) HasSegmentId() bool {
	return o != nil && o.SegmentId != nil
}

// SetSegmentId gets a reference to the given string and assigns it to the SegmentId field.
func (o *ProductAnalyticsAudienceSegmentSubquery) SetSegmentId(v string) {
	o.SegmentId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsAudienceSegmentSubquery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SegmentId != nil {
		toSerialize["segment_id"] = o.SegmentId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsAudienceSegmentSubquery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name,omitempty"`
		SegmentId *string `json:"segment_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "segment_id"})
	} else {
		return err
	}
	o.Name = all.Name
	o.SegmentId = all.SegmentId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
