// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JobDefinition Definition of a historical job.
type JobDefinition struct {
	// Calculated fields.
	CalculatedFields []CalculatedField `json:"calculatedFields,omitempty"`
	// Cases used for generating job results.
	Cases []SecurityMonitoringRuleCaseCreate `json:"cases"`
	// Starting time of data analyzed by the job.
	From int64 `json:"from"`
	// Additional grouping to perform on top of the existing groups in the query section. Must be a subset of the existing groups.
	GroupSignalsBy []string `json:"groupSignalsBy,omitempty"`
	// Index used to load the data.
	Index string `json:"index"`
	// Message for generated results.
	Message string `json:"message"`
	// Job name.
	Name string `json:"name"`
	// Job options.
	Options *HistoricalJobOptions `json:"options,omitempty"`
	// Queries for selecting logs analyzed by the job.
	Queries []HistoricalJobQuery `json:"queries"`
	// Reference tables used in the queries.
	ReferenceTables []SecurityMonitoringReferenceTable `json:"referenceTables,omitempty"`
	// Tags for generated signals.
	Tags []string `json:"tags,omitempty"`
	// Cases for generating results from third-party detection method. Only available for third-party detection method.
	ThirdPartyCases []SecurityMonitoringThirdPartyRuleCaseCreate `json:"thirdPartyCases,omitempty"`
	// Ending time of data analyzed by the job.
	To int64 `json:"to"`
	// Job type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJobDefinition instantiates a new JobDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJobDefinition(cases []SecurityMonitoringRuleCaseCreate, from int64, index string, message string, name string, queries []HistoricalJobQuery, to int64) *JobDefinition {
	this := JobDefinition{}
	this.Cases = cases
	this.From = from
	this.Index = index
	this.Message = message
	this.Name = name
	this.Queries = queries
	this.To = to
	return &this
}

// NewJobDefinitionWithDefaults instantiates a new JobDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJobDefinitionWithDefaults() *JobDefinition {
	this := JobDefinition{}
	return &this
}

// GetCalculatedFields returns the CalculatedFields field value if set, zero value otherwise.
func (o *JobDefinition) GetCalculatedFields() []CalculatedField {
	if o == nil || o.CalculatedFields == nil {
		var ret []CalculatedField
		return ret
	}
	return o.CalculatedFields
}

// GetCalculatedFieldsOk returns a tuple with the CalculatedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetCalculatedFieldsOk() (*[]CalculatedField, bool) {
	if o == nil || o.CalculatedFields == nil {
		return nil, false
	}
	return &o.CalculatedFields, true
}

// HasCalculatedFields returns a boolean if a field has been set.
func (o *JobDefinition) HasCalculatedFields() bool {
	return o != nil && o.CalculatedFields != nil
}

// SetCalculatedFields gets a reference to the given []CalculatedField and assigns it to the CalculatedFields field.
func (o *JobDefinition) SetCalculatedFields(v []CalculatedField) {
	o.CalculatedFields = v
}

// GetCases returns the Cases field value.
func (o *JobDefinition) GetCases() []SecurityMonitoringRuleCaseCreate {
	if o == nil {
		var ret []SecurityMonitoringRuleCaseCreate
		return ret
	}
	return o.Cases
}

// GetCasesOk returns a tuple with the Cases field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetCasesOk() (*[]SecurityMonitoringRuleCaseCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cases, true
}

// SetCases sets field value.
func (o *JobDefinition) SetCases(v []SecurityMonitoringRuleCaseCreate) {
	o.Cases = v
}

// GetFrom returns the From field value.
func (o *JobDefinition) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *JobDefinition) SetFrom(v int64) {
	o.From = v
}

// GetGroupSignalsBy returns the GroupSignalsBy field value if set, zero value otherwise.
func (o *JobDefinition) GetGroupSignalsBy() []string {
	if o == nil || o.GroupSignalsBy == nil {
		var ret []string
		return ret
	}
	return o.GroupSignalsBy
}

// GetGroupSignalsByOk returns a tuple with the GroupSignalsBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetGroupSignalsByOk() (*[]string, bool) {
	if o == nil || o.GroupSignalsBy == nil {
		return nil, false
	}
	return &o.GroupSignalsBy, true
}

// HasGroupSignalsBy returns a boolean if a field has been set.
func (o *JobDefinition) HasGroupSignalsBy() bool {
	return o != nil && o.GroupSignalsBy != nil
}

// SetGroupSignalsBy gets a reference to the given []string and assigns it to the GroupSignalsBy field.
func (o *JobDefinition) SetGroupSignalsBy(v []string) {
	o.GroupSignalsBy = v
}

// GetIndex returns the Index field value.
func (o *JobDefinition) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value.
func (o *JobDefinition) SetIndex(v string) {
	o.Index = v
}

// GetMessage returns the Message field value.
func (o *JobDefinition) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *JobDefinition) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value.
func (o *JobDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *JobDefinition) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *JobDefinition) GetOptions() HistoricalJobOptions {
	if o == nil || o.Options == nil {
		var ret HistoricalJobOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetOptionsOk() (*HistoricalJobOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *JobDefinition) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given HistoricalJobOptions and assigns it to the Options field.
func (o *JobDefinition) SetOptions(v HistoricalJobOptions) {
	o.Options = &v
}

// GetQueries returns the Queries field value.
func (o *JobDefinition) GetQueries() []HistoricalJobQuery {
	if o == nil {
		var ret []HistoricalJobQuery
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetQueriesOk() (*[]HistoricalJobQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *JobDefinition) SetQueries(v []HistoricalJobQuery) {
	o.Queries = v
}

// GetReferenceTables returns the ReferenceTables field value if set, zero value otherwise.
func (o *JobDefinition) GetReferenceTables() []SecurityMonitoringReferenceTable {
	if o == nil || o.ReferenceTables == nil {
		var ret []SecurityMonitoringReferenceTable
		return ret
	}
	return o.ReferenceTables
}

// GetReferenceTablesOk returns a tuple with the ReferenceTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetReferenceTablesOk() (*[]SecurityMonitoringReferenceTable, bool) {
	if o == nil || o.ReferenceTables == nil {
		return nil, false
	}
	return &o.ReferenceTables, true
}

// HasReferenceTables returns a boolean if a field has been set.
func (o *JobDefinition) HasReferenceTables() bool {
	return o != nil && o.ReferenceTables != nil
}

// SetReferenceTables gets a reference to the given []SecurityMonitoringReferenceTable and assigns it to the ReferenceTables field.
func (o *JobDefinition) SetReferenceTables(v []SecurityMonitoringReferenceTable) {
	o.ReferenceTables = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *JobDefinition) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *JobDefinition) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *JobDefinition) SetTags(v []string) {
	o.Tags = v
}

// GetThirdPartyCases returns the ThirdPartyCases field value if set, zero value otherwise.
func (o *JobDefinition) GetThirdPartyCases() []SecurityMonitoringThirdPartyRuleCaseCreate {
	if o == nil || o.ThirdPartyCases == nil {
		var ret []SecurityMonitoringThirdPartyRuleCaseCreate
		return ret
	}
	return o.ThirdPartyCases
}

// GetThirdPartyCasesOk returns a tuple with the ThirdPartyCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetThirdPartyCasesOk() (*[]SecurityMonitoringThirdPartyRuleCaseCreate, bool) {
	if o == nil || o.ThirdPartyCases == nil {
		return nil, false
	}
	return &o.ThirdPartyCases, true
}

// HasThirdPartyCases returns a boolean if a field has been set.
func (o *JobDefinition) HasThirdPartyCases() bool {
	return o != nil && o.ThirdPartyCases != nil
}

// SetThirdPartyCases gets a reference to the given []SecurityMonitoringThirdPartyRuleCaseCreate and assigns it to the ThirdPartyCases field.
func (o *JobDefinition) SetThirdPartyCases(v []SecurityMonitoringThirdPartyRuleCaseCreate) {
	o.ThirdPartyCases = v
}

// GetTo returns the To field value.
func (o *JobDefinition) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *JobDefinition) SetTo(v int64) {
	o.To = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JobDefinition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDefinition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JobDefinition) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *JobDefinition) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JobDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CalculatedFields != nil {
		toSerialize["calculatedFields"] = o.CalculatedFields
	}
	toSerialize["cases"] = o.Cases
	toSerialize["from"] = o.From
	if o.GroupSignalsBy != nil {
		toSerialize["groupSignalsBy"] = o.GroupSignalsBy
	}
	toSerialize["index"] = o.Index
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	toSerialize["queries"] = o.Queries
	if o.ReferenceTables != nil {
		toSerialize["referenceTables"] = o.ReferenceTables
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ThirdPartyCases != nil {
		toSerialize["thirdPartyCases"] = o.ThirdPartyCases
	}
	toSerialize["to"] = o.To
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JobDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CalculatedFields []CalculatedField                            `json:"calculatedFields,omitempty"`
		Cases            *[]SecurityMonitoringRuleCaseCreate          `json:"cases"`
		From             *int64                                       `json:"from"`
		GroupSignalsBy   []string                                     `json:"groupSignalsBy,omitempty"`
		Index            *string                                      `json:"index"`
		Message          *string                                      `json:"message"`
		Name             *string                                      `json:"name"`
		Options          *HistoricalJobOptions                        `json:"options,omitempty"`
		Queries          *[]HistoricalJobQuery                        `json:"queries"`
		ReferenceTables  []SecurityMonitoringReferenceTable           `json:"referenceTables,omitempty"`
		Tags             []string                                     `json:"tags,omitempty"`
		ThirdPartyCases  []SecurityMonitoringThirdPartyRuleCaseCreate `json:"thirdPartyCases,omitempty"`
		To               *int64                                       `json:"to"`
		Type             *string                                      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cases == nil {
		return fmt.Errorf("required field cases missing")
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Index == nil {
		return fmt.Errorf("required field index missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"calculatedFields", "cases", "from", "groupSignalsBy", "index", "message", "name", "options", "queries", "referenceTables", "tags", "thirdPartyCases", "to", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CalculatedFields = all.CalculatedFields
	o.Cases = *all.Cases
	o.From = *all.From
	o.GroupSignalsBy = all.GroupSignalsBy
	o.Index = *all.Index
	o.Message = *all.Message
	o.Name = *all.Name
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	o.Queries = *all.Queries
	o.ReferenceTables = all.ReferenceTables
	o.Tags = all.Tags
	o.ThirdPartyCases = all.ThirdPartyCases
	o.To = *all.To
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
