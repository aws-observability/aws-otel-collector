// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsSchemaData Configuration of the schema data to use.
type LogsSchemaData struct {
	// Class name of the schema to use.
	ClassName string `json:"class_name"`
	// Class UID of the schema to use.
	ClassUid int64 `json:"class_uid"`
	// Optional list of profiles to modify the schema.
	Profiles []string `json:"profiles,omitempty"`
	// Type of schema to use.
	SchemaType string `json:"schema_type"`
	// Version of the schema to use.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsSchemaData instantiates a new LogsSchemaData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsSchemaData(className string, classUid int64, schemaType string, version string) *LogsSchemaData {
	this := LogsSchemaData{}
	this.ClassName = className
	this.ClassUid = classUid
	this.SchemaType = schemaType
	this.Version = version
	return &this
}

// NewLogsSchemaDataWithDefaults instantiates a new LogsSchemaData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsSchemaDataWithDefaults() *LogsSchemaData {
	this := LogsSchemaData{}
	return &this
}

// GetClassName returns the ClassName field value.
func (o *LogsSchemaData) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaData) GetClassNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value.
func (o *LogsSchemaData) SetClassName(v string) {
	o.ClassName = v
}

// GetClassUid returns the ClassUid field value.
func (o *LogsSchemaData) GetClassUid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ClassUid
}

// GetClassUidOk returns a tuple with the ClassUid field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaData) GetClassUidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassUid, true
}

// SetClassUid sets field value.
func (o *LogsSchemaData) SetClassUid(v int64) {
	o.ClassUid = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *LogsSchemaData) GetProfiles() []string {
	if o == nil || o.Profiles == nil {
		var ret []string
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsSchemaData) GetProfilesOk() (*[]string, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *LogsSchemaData) HasProfiles() bool {
	return o != nil && o.Profiles != nil
}

// SetProfiles gets a reference to the given []string and assigns it to the Profiles field.
func (o *LogsSchemaData) SetProfiles(v []string) {
	o.Profiles = v
}

// GetSchemaType returns the SchemaType field value.
func (o *LogsSchemaData) GetSchemaType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SchemaType
}

// GetSchemaTypeOk returns a tuple with the SchemaType field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaData) GetSchemaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaType, true
}

// SetSchemaType sets field value.
func (o *LogsSchemaData) SetSchemaType(v string) {
	o.SchemaType = v
}

// GetVersion returns the Version field value.
func (o *LogsSchemaData) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *LogsSchemaData) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *LogsSchemaData) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsSchemaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["class_name"] = o.ClassName
	toSerialize["class_uid"] = o.ClassUid
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}
	toSerialize["schema_type"] = o.SchemaType
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsSchemaData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClassName  *string  `json:"class_name"`
		ClassUid   *int64   `json:"class_uid"`
		Profiles   []string `json:"profiles,omitempty"`
		SchemaType *string  `json:"schema_type"`
		Version    *string  `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClassName == nil {
		return fmt.Errorf("required field class_name missing")
	}
	if all.ClassUid == nil {
		return fmt.Errorf("required field class_uid missing")
	}
	if all.SchemaType == nil {
		return fmt.Errorf("required field schema_type missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"class_name", "class_uid", "profiles", "schema_type", "version"})
	} else {
		return err
	}
	o.ClassName = *all.ClassName
	o.ClassUid = *all.ClassUid
	o.Profiles = all.Profiles
	o.SchemaType = *all.SchemaType
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
