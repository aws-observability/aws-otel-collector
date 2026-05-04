// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationDataAttributesSource The source of the degradation.
type DegradationDataAttributesSource struct {
	// Timestamp of when the source was created.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the source.
	SourceId string `json:"source_id"`
	// The type of the source.
	Type DegradationDataAttributesSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationDataAttributesSource instantiates a new DegradationDataAttributesSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationDataAttributesSource(createdAt time.Time, sourceId string, typeVar DegradationDataAttributesSourceType) *DegradationDataAttributesSource {
	this := DegradationDataAttributesSource{}
	this.CreatedAt = createdAt
	this.SourceId = sourceId
	this.Type = typeVar
	return &this
}

// NewDegradationDataAttributesSourceWithDefaults instantiates a new DegradationDataAttributesSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationDataAttributesSourceWithDefaults() *DegradationDataAttributesSource {
	this := DegradationDataAttributesSource{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DegradationDataAttributesSource) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributesSource) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DegradationDataAttributesSource) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetSourceId returns the SourceId field value.
func (o *DegradationDataAttributesSource) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributesSource) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value.
func (o *DegradationDataAttributesSource) SetSourceId(v string) {
	o.SourceId = v
}

// GetType returns the Type field value.
func (o *DegradationDataAttributesSource) GetType() DegradationDataAttributesSourceType {
	if o == nil {
		var ret DegradationDataAttributesSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributesSource) GetTypeOk() (*DegradationDataAttributesSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DegradationDataAttributesSource) SetType(v DegradationDataAttributesSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationDataAttributesSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["source_id"] = o.SourceId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DegradationDataAttributesSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time                           `json:"created_at"`
		SourceId  *string                              `json:"source_id"`
		Type      *DegradationDataAttributesSourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.SourceId == nil {
		return fmt.Errorf("required field source_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "source_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.SourceId = *all.SourceId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
