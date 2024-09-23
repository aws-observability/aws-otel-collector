// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentIntegrationMetadataAttributes Incident integration metadata's attributes for a create request.
type IncidentIntegrationMetadataAttributes struct {
	// Timestamp when the incident todo was created.
	Created *time.Time `json:"created,omitempty"`
	// UUID of the incident this integration metadata is connected to.
	IncidentId *string `json:"incident_id,omitempty"`
	// A number indicating the type of integration this metadata is for. 1 indicates Slack;
	// 8 indicates Jira.
	IntegrationType int32 `json:"integration_type"`
	// Incident integration metadata's metadata attribute.
	Metadata IncidentIntegrationMetadataMetadata `json:"metadata"`
	// Timestamp when the incident todo was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// A number indicating the status of this integration metadata. 0 indicates unknown;
	// 1 indicates pending; 2 indicates complete; 3 indicates manually created;
	// 4 indicates manually updated; 5 indicates failed.
	Status *int32 `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentIntegrationMetadataAttributes instantiates a new IncidentIntegrationMetadataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentIntegrationMetadataAttributes(integrationType int32, metadata IncidentIntegrationMetadataMetadata) *IncidentIntegrationMetadataAttributes {
	this := IncidentIntegrationMetadataAttributes{}
	this.IntegrationType = integrationType
	this.Metadata = metadata
	return &this
}

// NewIncidentIntegrationMetadataAttributesWithDefaults instantiates a new IncidentIntegrationMetadataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentIntegrationMetadataAttributesWithDefaults() *IncidentIntegrationMetadataAttributes {
	this := IncidentIntegrationMetadataAttributes{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IncidentIntegrationMetadataAttributes) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IncidentIntegrationMetadataAttributes) HasCreated() bool {
	return o != nil && o.Created != nil
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IncidentIntegrationMetadataAttributes) SetCreated(v time.Time) {
	o.Created = &v
}

// GetIncidentId returns the IncidentId field value if set, zero value otherwise.
func (o *IncidentIntegrationMetadataAttributes) GetIncidentId() string {
	if o == nil || o.IncidentId == nil {
		var ret string
		return ret
	}
	return *o.IncidentId
}

// GetIncidentIdOk returns a tuple with the IncidentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataAttributes) GetIncidentIdOk() (*string, bool) {
	if o == nil || o.IncidentId == nil {
		return nil, false
	}
	return o.IncidentId, true
}

// HasIncidentId returns a boolean if a field has been set.
func (o *IncidentIntegrationMetadataAttributes) HasIncidentId() bool {
	return o != nil && o.IncidentId != nil
}

// SetIncidentId gets a reference to the given string and assigns it to the IncidentId field.
func (o *IncidentIntegrationMetadataAttributes) SetIncidentId(v string) {
	o.IncidentId = &v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *IncidentIntegrationMetadataAttributes) GetIntegrationType() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataAttributes) GetIntegrationTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *IncidentIntegrationMetadataAttributes) SetIntegrationType(v int32) {
	o.IntegrationType = v
}

// GetMetadata returns the Metadata field value.
func (o *IncidentIntegrationMetadataAttributes) GetMetadata() IncidentIntegrationMetadataMetadata {
	if o == nil {
		var ret IncidentIntegrationMetadataMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataAttributes) GetMetadataOk() (*IncidentIntegrationMetadataMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *IncidentIntegrationMetadataAttributes) SetMetadata(v IncidentIntegrationMetadataMetadata) {
	o.Metadata = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IncidentIntegrationMetadataAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IncidentIntegrationMetadataAttributes) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *IncidentIntegrationMetadataAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IncidentIntegrationMetadataAttributes) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataAttributes) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IncidentIntegrationMetadataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *IncidentIntegrationMetadataAttributes) SetStatus(v int32) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentIntegrationMetadataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Created != nil {
		if o.Created.Nanosecond() == 0 {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.IncidentId != nil {
		toSerialize["incident_id"] = o.IncidentId
	}
	toSerialize["integration_type"] = o.IntegrationType
	toSerialize["metadata"] = o.Metadata
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentIntegrationMetadataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Created         *time.Time                           `json:"created,omitempty"`
		IncidentId      *string                              `json:"incident_id,omitempty"`
		IntegrationType *int32                               `json:"integration_type"`
		Metadata        *IncidentIntegrationMetadataMetadata `json:"metadata"`
		Modified        *time.Time                           `json:"modified,omitempty"`
		Status          *int32                               `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IntegrationType == nil {
		return fmt.Errorf("required field integration_type missing")
	}
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created", "incident_id", "integration_type", "metadata", "modified", "status"})
	} else {
		return err
	}
	o.Created = all.Created
	o.IncidentId = all.IncidentId
	o.IntegrationType = *all.IntegrationType
	o.Metadata = *all.Metadata
	o.Modified = all.Modified
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
