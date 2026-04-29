// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InvestigationConclusion A full explanation of the finding, including root cause analysis and supporting evidence.
type InvestigationConclusion struct {
	// A full explanation of the finding, including root cause analysis and supporting evidence.
	Description string `json:"description"`
	// A summary of the finding, including affected components and timeframe.
	Summary string `json:"summary"`
	// The title of the conclusion.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInvestigationConclusion instantiates a new InvestigationConclusion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInvestigationConclusion(description string, summary string, title string) *InvestigationConclusion {
	this := InvestigationConclusion{}
	this.Description = description
	this.Summary = summary
	this.Title = title
	return &this
}

// NewInvestigationConclusionWithDefaults instantiates a new InvestigationConclusion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInvestigationConclusionWithDefaults() *InvestigationConclusion {
	this := InvestigationConclusion{}
	return &this
}

// GetDescription returns the Description field value.
func (o *InvestigationConclusion) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InvestigationConclusion) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *InvestigationConclusion) SetDescription(v string) {
	o.Description = v
}

// GetSummary returns the Summary field value.
func (o *InvestigationConclusion) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *InvestigationConclusion) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *InvestigationConclusion) SetSummary(v string) {
	o.Summary = v
}

// GetTitle returns the Title field value.
func (o *InvestigationConclusion) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *InvestigationConclusion) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *InvestigationConclusion) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InvestigationConclusion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["summary"] = o.Summary
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InvestigationConclusion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description"`
		Summary     *string `json:"summary"`
		Title       *string `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "summary", "title"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Summary = *all.Summary
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
