// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuesSearchRequestDataAttributes Object describing a search issue request.
type IssuesSearchRequestDataAttributes struct {
	// Start date (inclusive) of the query in milliseconds since the Unix epoch.
	From int64 `json:"from"`
	// The attribute to sort the search results by.
	OrderBy *IssuesSearchRequestDataAttributesOrderBy `json:"order_by,omitempty"`
	// Persona for the search. Either track(s) or persona(s) must be specified.
	Persona *IssuesSearchRequestDataAttributesPersona `json:"persona,omitempty"`
	// Search query following the event search syntax.
	Query string `json:"query"`
	// End date (exclusive) of the query in milliseconds since the Unix epoch.
	To int64 `json:"to"`
	// Track of the events to query. Either track(s) or persona(s) must be specified.
	Track *IssuesSearchRequestDataAttributesTrack `json:"track,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssuesSearchRequestDataAttributes instantiates a new IssuesSearchRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssuesSearchRequestDataAttributes(from int64, query string, to int64) *IssuesSearchRequestDataAttributes {
	this := IssuesSearchRequestDataAttributes{}
	this.From = from
	this.Query = query
	this.To = to
	return &this
}

// NewIssuesSearchRequestDataAttributesWithDefaults instantiates a new IssuesSearchRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssuesSearchRequestDataAttributesWithDefaults() *IssuesSearchRequestDataAttributes {
	this := IssuesSearchRequestDataAttributes{}
	return &this
}

// GetFrom returns the From field value.
func (o *IssuesSearchRequestDataAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *IssuesSearchRequestDataAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *IssuesSearchRequestDataAttributes) SetFrom(v int64) {
	o.From = v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *IssuesSearchRequestDataAttributes) GetOrderBy() IssuesSearchRequestDataAttributesOrderBy {
	if o == nil || o.OrderBy == nil {
		var ret IssuesSearchRequestDataAttributesOrderBy
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchRequestDataAttributes) GetOrderByOk() (*IssuesSearchRequestDataAttributesOrderBy, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *IssuesSearchRequestDataAttributes) HasOrderBy() bool {
	return o != nil && o.OrderBy != nil
}

// SetOrderBy gets a reference to the given IssuesSearchRequestDataAttributesOrderBy and assigns it to the OrderBy field.
func (o *IssuesSearchRequestDataAttributes) SetOrderBy(v IssuesSearchRequestDataAttributesOrderBy) {
	o.OrderBy = &v
}

// GetPersona returns the Persona field value if set, zero value otherwise.
func (o *IssuesSearchRequestDataAttributes) GetPersona() IssuesSearchRequestDataAttributesPersona {
	if o == nil || o.Persona == nil {
		var ret IssuesSearchRequestDataAttributesPersona
		return ret
	}
	return *o.Persona
}

// GetPersonaOk returns a tuple with the Persona field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchRequestDataAttributes) GetPersonaOk() (*IssuesSearchRequestDataAttributesPersona, bool) {
	if o == nil || o.Persona == nil {
		return nil, false
	}
	return o.Persona, true
}

// HasPersona returns a boolean if a field has been set.
func (o *IssuesSearchRequestDataAttributes) HasPersona() bool {
	return o != nil && o.Persona != nil
}

// SetPersona gets a reference to the given IssuesSearchRequestDataAttributesPersona and assigns it to the Persona field.
func (o *IssuesSearchRequestDataAttributes) SetPersona(v IssuesSearchRequestDataAttributesPersona) {
	o.Persona = &v
}

// GetQuery returns the Query field value.
func (o *IssuesSearchRequestDataAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *IssuesSearchRequestDataAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *IssuesSearchRequestDataAttributes) SetQuery(v string) {
	o.Query = v
}

// GetTo returns the To field value.
func (o *IssuesSearchRequestDataAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *IssuesSearchRequestDataAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *IssuesSearchRequestDataAttributes) SetTo(v int64) {
	o.To = v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *IssuesSearchRequestDataAttributes) GetTrack() IssuesSearchRequestDataAttributesTrack {
	if o == nil || o.Track == nil {
		var ret IssuesSearchRequestDataAttributesTrack
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchRequestDataAttributes) GetTrackOk() (*IssuesSearchRequestDataAttributesTrack, bool) {
	if o == nil || o.Track == nil {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *IssuesSearchRequestDataAttributes) HasTrack() bool {
	return o != nil && o.Track != nil
}

// SetTrack gets a reference to the given IssuesSearchRequestDataAttributesTrack and assigns it to the Track field.
func (o *IssuesSearchRequestDataAttributes) SetTrack(v IssuesSearchRequestDataAttributesTrack) {
	o.Track = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssuesSearchRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}
	if o.Persona != nil {
		toSerialize["persona"] = o.Persona
	}
	toSerialize["query"] = o.Query
	toSerialize["to"] = o.To
	if o.Track != nil {
		toSerialize["track"] = o.Track
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssuesSearchRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From    *int64                                    `json:"from"`
		OrderBy *IssuesSearchRequestDataAttributesOrderBy `json:"order_by,omitempty"`
		Persona *IssuesSearchRequestDataAttributesPersona `json:"persona,omitempty"`
		Query   *string                                   `json:"query"`
		To      *int64                                    `json:"to"`
		Track   *IssuesSearchRequestDataAttributesTrack   `json:"track,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "order_by", "persona", "query", "to", "track"})
	} else {
		return err
	}

	hasInvalidField := false
	o.From = *all.From
	if all.OrderBy != nil && !all.OrderBy.IsValid() {
		hasInvalidField = true
	} else {
		o.OrderBy = all.OrderBy
	}
	if all.Persona != nil && !all.Persona.IsValid() {
		hasInvalidField = true
	} else {
		o.Persona = all.Persona
	}
	o.Query = *all.Query
	o.To = *all.To
	if all.Track != nil && !all.Track.IsValid() {
		hasInvalidField = true
	} else {
		o.Track = all.Track
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
