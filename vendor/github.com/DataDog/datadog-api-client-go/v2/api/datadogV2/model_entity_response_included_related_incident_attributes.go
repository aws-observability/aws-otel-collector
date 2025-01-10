// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRelatedIncidentAttributes Incident attributes.
type EntityResponseIncludedRelatedIncidentAttributes struct {
	// Incident creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Incident URL.
	HtmlUrl *string `json:"htmlURL,omitempty"`
	// Incident provider.
	Provider *string `json:"provider,omitempty"`
	// Incident status.
	Status *string `json:"status,omitempty"`
	// Incident title.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseIncludedRelatedIncidentAttributes instantiates a new EntityResponseIncludedRelatedIncidentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseIncludedRelatedIncidentAttributes() *EntityResponseIncludedRelatedIncidentAttributes {
	this := EntityResponseIncludedRelatedIncidentAttributes{}
	return &this
}

// NewEntityResponseIncludedRelatedIncidentAttributesWithDefaults instantiates a new EntityResponseIncludedRelatedIncidentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseIncludedRelatedIncidentAttributesWithDefaults() *EntityResponseIncludedRelatedIncidentAttributes {
	this := EntityResponseIncludedRelatedIncidentAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EntityResponseIncludedRelatedIncidentAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl == nil {
		var ret string
		return ret
	}
	return *o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetHtmlUrlOk() (*string, bool) {
	if o == nil || o.HtmlUrl == nil {
		return nil, false
	}
	return o.HtmlUrl, true
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) HasHtmlUrl() bool {
	return o != nil && o.HtmlUrl != nil
}

// SetHtmlUrl gets a reference to the given string and assigns it to the HtmlUrl field.
func (o *EntityResponseIncludedRelatedIncidentAttributes) SetHtmlUrl(v string) {
	o.HtmlUrl = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *EntityResponseIncludedRelatedIncidentAttributes) SetProvider(v string) {
	o.Provider = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EntityResponseIncludedRelatedIncidentAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedIncidentAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EntityResponseIncludedRelatedIncidentAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseIncludedRelatedIncidentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.HtmlUrl != nil {
		toSerialize["htmlURL"] = o.HtmlUrl
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseIncludedRelatedIncidentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time `json:"createdAt,omitempty"`
		HtmlUrl   *string    `json:"htmlURL,omitempty"`
		Provider  *string    `json:"provider,omitempty"`
		Status    *string    `json:"status,omitempty"`
		Title     *string    `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "htmlURL", "provider", "status", "title"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.HtmlUrl = all.HtmlUrl
	o.Provider = all.Provider
	o.Status = all.Status
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
