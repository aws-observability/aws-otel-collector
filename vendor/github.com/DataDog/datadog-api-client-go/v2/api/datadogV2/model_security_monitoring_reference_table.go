// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringReferenceTable Reference tables used in the queries.
type SecurityMonitoringReferenceTable struct {
	// Whether to include or exclude the matched values.
	CheckPresence *bool `json:"checkPresence,omitempty"`
	// The name of the column in the reference table.
	ColumnName *string `json:"columnName,omitempty"`
	// The field in the log to match against the reference table.
	LogFieldPath *string `json:"logFieldPath,omitempty"`
	// The name of the query to apply the reference table to.
	RuleQueryName *string `json:"ruleQueryName,omitempty"`
	// The name of the reference table.
	TableName *string `json:"tableName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringReferenceTable instantiates a new SecurityMonitoringReferenceTable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringReferenceTable() *SecurityMonitoringReferenceTable {
	this := SecurityMonitoringReferenceTable{}
	return &this
}

// NewSecurityMonitoringReferenceTableWithDefaults instantiates a new SecurityMonitoringReferenceTable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringReferenceTableWithDefaults() *SecurityMonitoringReferenceTable {
	this := SecurityMonitoringReferenceTable{}
	return &this
}

// GetCheckPresence returns the CheckPresence field value if set, zero value otherwise.
func (o *SecurityMonitoringReferenceTable) GetCheckPresence() bool {
	if o == nil || o.CheckPresence == nil {
		var ret bool
		return ret
	}
	return *o.CheckPresence
}

// GetCheckPresenceOk returns a tuple with the CheckPresence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringReferenceTable) GetCheckPresenceOk() (*bool, bool) {
	if o == nil || o.CheckPresence == nil {
		return nil, false
	}
	return o.CheckPresence, true
}

// HasCheckPresence returns a boolean if a field has been set.
func (o *SecurityMonitoringReferenceTable) HasCheckPresence() bool {
	return o != nil && o.CheckPresence != nil
}

// SetCheckPresence gets a reference to the given bool and assigns it to the CheckPresence field.
func (o *SecurityMonitoringReferenceTable) SetCheckPresence(v bool) {
	o.CheckPresence = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *SecurityMonitoringReferenceTable) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringReferenceTable) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *SecurityMonitoringReferenceTable) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *SecurityMonitoringReferenceTable) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetLogFieldPath returns the LogFieldPath field value if set, zero value otherwise.
func (o *SecurityMonitoringReferenceTable) GetLogFieldPath() string {
	if o == nil || o.LogFieldPath == nil {
		var ret string
		return ret
	}
	return *o.LogFieldPath
}

// GetLogFieldPathOk returns a tuple with the LogFieldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringReferenceTable) GetLogFieldPathOk() (*string, bool) {
	if o == nil || o.LogFieldPath == nil {
		return nil, false
	}
	return o.LogFieldPath, true
}

// HasLogFieldPath returns a boolean if a field has been set.
func (o *SecurityMonitoringReferenceTable) HasLogFieldPath() bool {
	return o != nil && o.LogFieldPath != nil
}

// SetLogFieldPath gets a reference to the given string and assigns it to the LogFieldPath field.
func (o *SecurityMonitoringReferenceTable) SetLogFieldPath(v string) {
	o.LogFieldPath = &v
}

// GetRuleQueryName returns the RuleQueryName field value if set, zero value otherwise.
func (o *SecurityMonitoringReferenceTable) GetRuleQueryName() string {
	if o == nil || o.RuleQueryName == nil {
		var ret string
		return ret
	}
	return *o.RuleQueryName
}

// GetRuleQueryNameOk returns a tuple with the RuleQueryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringReferenceTable) GetRuleQueryNameOk() (*string, bool) {
	if o == nil || o.RuleQueryName == nil {
		return nil, false
	}
	return o.RuleQueryName, true
}

// HasRuleQueryName returns a boolean if a field has been set.
func (o *SecurityMonitoringReferenceTable) HasRuleQueryName() bool {
	return o != nil && o.RuleQueryName != nil
}

// SetRuleQueryName gets a reference to the given string and assigns it to the RuleQueryName field.
func (o *SecurityMonitoringReferenceTable) SetRuleQueryName(v string) {
	o.RuleQueryName = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *SecurityMonitoringReferenceTable) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringReferenceTable) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *SecurityMonitoringReferenceTable) HasTableName() bool {
	return o != nil && o.TableName != nil
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *SecurityMonitoringReferenceTable) SetTableName(v string) {
	o.TableName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringReferenceTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CheckPresence != nil {
		toSerialize["checkPresence"] = o.CheckPresence
	}
	if o.ColumnName != nil {
		toSerialize["columnName"] = o.ColumnName
	}
	if o.LogFieldPath != nil {
		toSerialize["logFieldPath"] = o.LogFieldPath
	}
	if o.RuleQueryName != nil {
		toSerialize["ruleQueryName"] = o.RuleQueryName
	}
	if o.TableName != nil {
		toSerialize["tableName"] = o.TableName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringReferenceTable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CheckPresence *bool   `json:"checkPresence,omitempty"`
		ColumnName    *string `json:"columnName,omitempty"`
		LogFieldPath  *string `json:"logFieldPath,omitempty"`
		RuleQueryName *string `json:"ruleQueryName,omitempty"`
		TableName     *string `json:"tableName,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"checkPresence", "columnName", "logFieldPath", "ruleQueryName", "tableName"})
	} else {
		return err
	}
	o.CheckPresence = all.CheckPresence
	o.ColumnName = all.ColumnName
	o.LogFieldPath = all.LogFieldPath
	o.RuleQueryName = all.RuleQueryName
	o.TableName = all.TableName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
