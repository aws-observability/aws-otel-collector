// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentImportRequestAttributes The incident's attributes for an import request.
type IncidentImportRequestAttributes struct {
	// Timestamp when the incident was declared.
	Declared *time.Time `json:"declared,omitempty"`
	// Timestamp when the incident was detected.
	Detected *time.Time `json:"detected,omitempty"`
	// A condensed view of the user-defined fields for which to create initial selections.
	Fields map[string]IncidentImportFieldAttributes `json:"fields,omitempty"`
	// A unique identifier that represents the incident type. If not provided, the default incident type is used.
	IncidentTypeUuid *string `json:"incident_type_uuid,omitempty"`
	// Timestamp when the incident was resolved. Can only be set when the state field is set to 'resolved'.
	Resolved *time.Time `json:"resolved,omitempty"`
	// The title of the incident that summarizes what happened.
	Title string `json:"title"`
	// The visibility of the incident.
	Visibility *IncidentImportVisibility `json:"visibility,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentImportRequestAttributes instantiates a new IncidentImportRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentImportRequestAttributes(title string) *IncidentImportRequestAttributes {
	this := IncidentImportRequestAttributes{}
	this.Title = title
	var visibility IncidentImportVisibility = INCIDENTIMPORTVISIBILITY_ORGANIZATION
	this.Visibility = &visibility
	return &this
}

// NewIncidentImportRequestAttributesWithDefaults instantiates a new IncidentImportRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentImportRequestAttributesWithDefaults() *IncidentImportRequestAttributes {
	this := IncidentImportRequestAttributes{}
	var visibility IncidentImportVisibility = INCIDENTIMPORTVISIBILITY_ORGANIZATION
	this.Visibility = &visibility
	return &this
}

// GetDeclared returns the Declared field value if set, zero value otherwise.
func (o *IncidentImportRequestAttributes) GetDeclared() time.Time {
	if o == nil || o.Declared == nil {
		var ret time.Time
		return ret
	}
	return *o.Declared
}

// GetDeclaredOk returns a tuple with the Declared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestAttributes) GetDeclaredOk() (*time.Time, bool) {
	if o == nil || o.Declared == nil {
		return nil, false
	}
	return o.Declared, true
}

// HasDeclared returns a boolean if a field has been set.
func (o *IncidentImportRequestAttributes) HasDeclared() bool {
	return o != nil && o.Declared != nil
}

// SetDeclared gets a reference to the given time.Time and assigns it to the Declared field.
func (o *IncidentImportRequestAttributes) SetDeclared(v time.Time) {
	o.Declared = &v
}

// GetDetected returns the Detected field value if set, zero value otherwise.
func (o *IncidentImportRequestAttributes) GetDetected() time.Time {
	if o == nil || o.Detected == nil {
		var ret time.Time
		return ret
	}
	return *o.Detected
}

// GetDetectedOk returns a tuple with the Detected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestAttributes) GetDetectedOk() (*time.Time, bool) {
	if o == nil || o.Detected == nil {
		return nil, false
	}
	return o.Detected, true
}

// HasDetected returns a boolean if a field has been set.
func (o *IncidentImportRequestAttributes) HasDetected() bool {
	return o != nil && o.Detected != nil
}

// SetDetected gets a reference to the given time.Time and assigns it to the Detected field.
func (o *IncidentImportRequestAttributes) SetDetected(v time.Time) {
	o.Detected = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *IncidentImportRequestAttributes) GetFields() map[string]IncidentImportFieldAttributes {
	if o == nil || o.Fields == nil {
		var ret map[string]IncidentImportFieldAttributes
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestAttributes) GetFieldsOk() (*map[string]IncidentImportFieldAttributes, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *IncidentImportRequestAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]IncidentImportFieldAttributes and assigns it to the Fields field.
func (o *IncidentImportRequestAttributes) SetFields(v map[string]IncidentImportFieldAttributes) {
	o.Fields = v
}

// GetIncidentTypeUuid returns the IncidentTypeUuid field value if set, zero value otherwise.
func (o *IncidentImportRequestAttributes) GetIncidentTypeUuid() string {
	if o == nil || o.IncidentTypeUuid == nil {
		var ret string
		return ret
	}
	return *o.IncidentTypeUuid
}

// GetIncidentTypeUuidOk returns a tuple with the IncidentTypeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestAttributes) GetIncidentTypeUuidOk() (*string, bool) {
	if o == nil || o.IncidentTypeUuid == nil {
		return nil, false
	}
	return o.IncidentTypeUuid, true
}

// HasIncidentTypeUuid returns a boolean if a field has been set.
func (o *IncidentImportRequestAttributes) HasIncidentTypeUuid() bool {
	return o != nil && o.IncidentTypeUuid != nil
}

// SetIncidentTypeUuid gets a reference to the given string and assigns it to the IncidentTypeUuid field.
func (o *IncidentImportRequestAttributes) SetIncidentTypeUuid(v string) {
	o.IncidentTypeUuid = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *IncidentImportRequestAttributes) GetResolved() time.Time {
	if o == nil || o.Resolved == nil {
		var ret time.Time
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestAttributes) GetResolvedOk() (*time.Time, bool) {
	if o == nil || o.Resolved == nil {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *IncidentImportRequestAttributes) HasResolved() bool {
	return o != nil && o.Resolved != nil
}

// SetResolved gets a reference to the given time.Time and assigns it to the Resolved field.
func (o *IncidentImportRequestAttributes) SetResolved(v time.Time) {
	o.Resolved = &v
}

// GetTitle returns the Title field value.
func (o *IncidentImportRequestAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *IncidentImportRequestAttributes) SetTitle(v string) {
	o.Title = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *IncidentImportRequestAttributes) GetVisibility() IncidentImportVisibility {
	if o == nil || o.Visibility == nil {
		var ret IncidentImportVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentImportRequestAttributes) GetVisibilityOk() (*IncidentImportVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *IncidentImportRequestAttributes) HasVisibility() bool {
	return o != nil && o.Visibility != nil
}

// SetVisibility gets a reference to the given IncidentImportVisibility and assigns it to the Visibility field.
func (o *IncidentImportRequestAttributes) SetVisibility(v IncidentImportVisibility) {
	o.Visibility = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentImportRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Declared != nil {
		if o.Declared.Nanosecond() == 0 {
			toSerialize["declared"] = o.Declared.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["declared"] = o.Declared.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Detected != nil {
		if o.Detected.Nanosecond() == 0 {
			toSerialize["detected"] = o.Detected.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["detected"] = o.Detected.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.IncidentTypeUuid != nil {
		toSerialize["incident_type_uuid"] = o.IncidentTypeUuid
	}
	if o.Resolved != nil {
		if o.Resolved.Nanosecond() == 0 {
			toSerialize["resolved"] = o.Resolved.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["resolved"] = o.Resolved.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["title"] = o.Title
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentImportRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Declared         *time.Time                               `json:"declared,omitempty"`
		Detected         *time.Time                               `json:"detected,omitempty"`
		Fields           map[string]IncidentImportFieldAttributes `json:"fields,omitempty"`
		IncidentTypeUuid *string                                  `json:"incident_type_uuid,omitempty"`
		Resolved         *time.Time                               `json:"resolved,omitempty"`
		Title            *string                                  `json:"title"`
		Visibility       *IncidentImportVisibility                `json:"visibility,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"declared", "detected", "fields", "incident_type_uuid", "resolved", "title", "visibility"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Declared = all.Declared
	o.Detected = all.Detected
	o.Fields = all.Fields
	o.IncidentTypeUuid = all.IncidentTypeUuid
	o.Resolved = all.Resolved
	o.Title = *all.Title
	if all.Visibility != nil && !all.Visibility.IsValid() {
		hasInvalidField = true
	} else {
		o.Visibility = all.Visibility
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
