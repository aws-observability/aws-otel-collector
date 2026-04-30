// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationDataAttributes The attributes of a degradation.
type DegradationDataAttributes struct {
	// Components affected by the degradation.
	ComponentsAffected []DegradationDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
	// Timestamp of when the degradation was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Description of the degradation.
	Description *string `json:"description,omitempty"`
	// Timestamp of when the degradation was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The source of the degradation.
	Source *DegradationDataAttributesSource `json:"source,omitempty"`
	// The status of the degradation.
	Status *CreateDegradationRequestDataAttributesStatus `json:"status,omitempty"`
	// Title of the degradation.
	Title *string `json:"title,omitempty"`
	// Past updates made to the degradation.
	Updates []DegradationDataAttributesUpdatesItems `json:"updates,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationDataAttributes instantiates a new DegradationDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationDataAttributes() *DegradationDataAttributes {
	this := DegradationDataAttributes{}
	return &this
}

// NewDegradationDataAttributesWithDefaults instantiates a new DegradationDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationDataAttributesWithDefaults() *DegradationDataAttributes {
	this := DegradationDataAttributes{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value if set, zero value otherwise.
func (o *DegradationDataAttributes) GetComponentsAffected() []DegradationDataAttributesComponentsAffectedItems {
	if o == nil || o.ComponentsAffected == nil {
		var ret []DegradationDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributes) GetComponentsAffectedOk() (*[]DegradationDataAttributesComponentsAffectedItems, bool) {
	if o == nil || o.ComponentsAffected == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// HasComponentsAffected returns a boolean if a field has been set.
func (o *DegradationDataAttributes) HasComponentsAffected() bool {
	return o != nil && o.ComponentsAffected != nil
}

// SetComponentsAffected gets a reference to the given []DegradationDataAttributesComponentsAffectedItems and assigns it to the ComponentsAffected field.
func (o *DegradationDataAttributes) SetComponentsAffected(v []DegradationDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DegradationDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DegradationDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DegradationDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DegradationDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DegradationDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DegradationDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *DegradationDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *DegradationDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *DegradationDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DegradationDataAttributes) GetSource() DegradationDataAttributesSource {
	if o == nil || o.Source == nil {
		var ret DegradationDataAttributesSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributes) GetSourceOk() (*DegradationDataAttributesSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DegradationDataAttributes) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given DegradationDataAttributesSource and assigns it to the Source field.
func (o *DegradationDataAttributes) SetSource(v DegradationDataAttributesSource) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DegradationDataAttributes) GetStatus() CreateDegradationRequestDataAttributesStatus {
	if o == nil || o.Status == nil {
		var ret CreateDegradationRequestDataAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributes) GetStatusOk() (*CreateDegradationRequestDataAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DegradationDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given CreateDegradationRequestDataAttributesStatus and assigns it to the Status field.
func (o *DegradationDataAttributes) SetStatus(v CreateDegradationRequestDataAttributesStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DegradationDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DegradationDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DegradationDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *DegradationDataAttributes) GetUpdates() []DegradationDataAttributesUpdatesItems {
	if o == nil || o.Updates == nil {
		var ret []DegradationDataAttributesUpdatesItems
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataAttributes) GetUpdatesOk() (*[]DegradationDataAttributesUpdatesItems, bool) {
	if o == nil || o.Updates == nil {
		return nil, false
	}
	return &o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *DegradationDataAttributes) HasUpdates() bool {
	return o != nil && o.Updates != nil
}

// SetUpdates gets a reference to the given []DegradationDataAttributesUpdatesItems and assigns it to the Updates field.
func (o *DegradationDataAttributes) SetUpdates(v []DegradationDataAttributesUpdatesItems) {
	o.Updates = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentsAffected != nil {
		toSerialize["components_affected"] = o.ComponentsAffected
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Updates != nil {
		toSerialize["updates"] = o.Updates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DegradationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentsAffected []DegradationDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
		CreatedAt          *time.Time                                         `json:"created_at,omitempty"`
		Description        *string                                            `json:"description,omitempty"`
		ModifiedAt         *time.Time                                         `json:"modified_at,omitempty"`
		Source             *DegradationDataAttributesSource                   `json:"source,omitempty"`
		Status             *CreateDegradationRequestDataAttributesStatus      `json:"status,omitempty"`
		Title              *string                                            `json:"title,omitempty"`
		Updates            []DegradationDataAttributesUpdatesItems            `json:"updates,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components_affected", "created_at", "description", "modified_at", "source", "status", "title", "updates"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentsAffected = all.ComponentsAffected
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.ModifiedAt = all.ModifiedAt
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Title = all.Title
	o.Updates = all.Updates

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
