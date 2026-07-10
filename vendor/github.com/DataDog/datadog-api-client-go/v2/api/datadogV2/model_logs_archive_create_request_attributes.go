// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveCreateRequestAttributes The attributes associated with the archive.
type LogsArchiveCreateRequestAttributes struct {
	// The type of compression for the archive.
	CompressionMethod *LogsArchiveAttributesCompressionMethod `json:"compression_method,omitempty"`
	// An archive's destination.
	Destination LogsArchiveCreateRequestDestination `json:"destination"`
	// To store the tags in the archive, set the value "true".
	// If it is set to "false", the tags will be deleted when the logs are sent to the archive.
	IncludeTags *bool `json:"include_tags,omitempty"`
	// An array of attributes to use as lookup keys for the archive.
	LookupAttributes []string `json:"lookup_attributes,omitempty"`
	// The archive name.
	Name string `json:"name"`
	// An array of attributes to use as partition keys for the archive. The attribute used most frequently for querying should be first.
	PartitioningAttributes []string `json:"partitioning_attributes,omitempty"`
	// The archive query/filter. Logs matching this query are included in the archive.
	Query string `json:"query"`
	// Maximum scan size for rehydration from this archive.
	RehydrationMaxScanSizeInGb datadog.NullableInt64 `json:"rehydration_max_scan_size_in_gb,omitempty"`
	// An array of tags to add to rehydrated logs from an archive.
	RehydrationTags []string `json:"rehydration_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArchiveCreateRequestAttributes instantiates a new LogsArchiveCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArchiveCreateRequestAttributes(destination LogsArchiveCreateRequestDestination, name string, query string) *LogsArchiveCreateRequestAttributes {
	this := LogsArchiveCreateRequestAttributes{}
	var compressionMethod LogsArchiveAttributesCompressionMethod = LOGSARCHIVEATTRIBUTESCOMPRESSIONMETHOD_GZIP
	this.CompressionMethod = &compressionMethod
	this.Destination = destination
	var includeTags bool = false
	this.IncludeTags = &includeTags
	this.Name = name
	this.Query = query
	return &this
}

// NewLogsArchiveCreateRequestAttributesWithDefaults instantiates a new LogsArchiveCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArchiveCreateRequestAttributesWithDefaults() *LogsArchiveCreateRequestAttributes {
	this := LogsArchiveCreateRequestAttributes{}
	var compressionMethod LogsArchiveAttributesCompressionMethod = LOGSARCHIVEATTRIBUTESCOMPRESSIONMETHOD_GZIP
	this.CompressionMethod = &compressionMethod
	var includeTags bool = false
	this.IncludeTags = &includeTags
	return &this
}

// GetCompressionMethod returns the CompressionMethod field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequestAttributes) GetCompressionMethod() LogsArchiveAttributesCompressionMethod {
	if o == nil || o.CompressionMethod == nil {
		var ret LogsArchiveAttributesCompressionMethod
		return ret
	}
	return *o.CompressionMethod
}

// GetCompressionMethodOk returns a tuple with the CompressionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetCompressionMethodOk() (*LogsArchiveAttributesCompressionMethod, bool) {
	if o == nil || o.CompressionMethod == nil {
		return nil, false
	}
	return o.CompressionMethod, true
}

// HasCompressionMethod returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestAttributes) HasCompressionMethod() bool {
	return o != nil && o.CompressionMethod != nil
}

// SetCompressionMethod gets a reference to the given LogsArchiveAttributesCompressionMethod and assigns it to the CompressionMethod field.
func (o *LogsArchiveCreateRequestAttributes) SetCompressionMethod(v LogsArchiveAttributesCompressionMethod) {
	o.CompressionMethod = &v
}

// GetDestination returns the Destination field value.
func (o *LogsArchiveCreateRequestAttributes) GetDestination() LogsArchiveCreateRequestDestination {
	if o == nil {
		var ret LogsArchiveCreateRequestDestination
		return ret
	}
	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetDestinationOk() (*LogsArchiveCreateRequestDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value.
func (o *LogsArchiveCreateRequestAttributes) SetDestination(v LogsArchiveCreateRequestDestination) {
	o.Destination = v
}

// GetIncludeTags returns the IncludeTags field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequestAttributes) GetIncludeTags() bool {
	if o == nil || o.IncludeTags == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTags
}

// GetIncludeTagsOk returns a tuple with the IncludeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetIncludeTagsOk() (*bool, bool) {
	if o == nil || o.IncludeTags == nil {
		return nil, false
	}
	return o.IncludeTags, true
}

// HasIncludeTags returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestAttributes) HasIncludeTags() bool {
	return o != nil && o.IncludeTags != nil
}

// SetIncludeTags gets a reference to the given bool and assigns it to the IncludeTags field.
func (o *LogsArchiveCreateRequestAttributes) SetIncludeTags(v bool) {
	o.IncludeTags = &v
}

// GetLookupAttributes returns the LookupAttributes field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequestAttributes) GetLookupAttributes() []string {
	if o == nil || o.LookupAttributes == nil {
		var ret []string
		return ret
	}
	return o.LookupAttributes
}

// GetLookupAttributesOk returns a tuple with the LookupAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetLookupAttributesOk() (*[]string, bool) {
	if o == nil || o.LookupAttributes == nil {
		return nil, false
	}
	return &o.LookupAttributes, true
}

// HasLookupAttributes returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestAttributes) HasLookupAttributes() bool {
	return o != nil && o.LookupAttributes != nil
}

// SetLookupAttributes gets a reference to the given []string and assigns it to the LookupAttributes field.
func (o *LogsArchiveCreateRequestAttributes) SetLookupAttributes(v []string) {
	o.LookupAttributes = v
}

// GetName returns the Name field value.
func (o *LogsArchiveCreateRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LogsArchiveCreateRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetPartitioningAttributes returns the PartitioningAttributes field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequestAttributes) GetPartitioningAttributes() []string {
	if o == nil || o.PartitioningAttributes == nil {
		var ret []string
		return ret
	}
	return o.PartitioningAttributes
}

// GetPartitioningAttributesOk returns a tuple with the PartitioningAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetPartitioningAttributesOk() (*[]string, bool) {
	if o == nil || o.PartitioningAttributes == nil {
		return nil, false
	}
	return &o.PartitioningAttributes, true
}

// HasPartitioningAttributes returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestAttributes) HasPartitioningAttributes() bool {
	return o != nil && o.PartitioningAttributes != nil
}

// SetPartitioningAttributes gets a reference to the given []string and assigns it to the PartitioningAttributes field.
func (o *LogsArchiveCreateRequestAttributes) SetPartitioningAttributes(v []string) {
	o.PartitioningAttributes = v
}

// GetQuery returns the Query field value.
func (o *LogsArchiveCreateRequestAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *LogsArchiveCreateRequestAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRehydrationMaxScanSizeInGb returns the RehydrationMaxScanSizeInGb field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogsArchiveCreateRequestAttributes) GetRehydrationMaxScanSizeInGb() int64 {
	if o == nil || o.RehydrationMaxScanSizeInGb.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RehydrationMaxScanSizeInGb.Get()
}

// GetRehydrationMaxScanSizeInGbOk returns a tuple with the RehydrationMaxScanSizeInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LogsArchiveCreateRequestAttributes) GetRehydrationMaxScanSizeInGbOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RehydrationMaxScanSizeInGb.Get(), o.RehydrationMaxScanSizeInGb.IsSet()
}

// HasRehydrationMaxScanSizeInGb returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestAttributes) HasRehydrationMaxScanSizeInGb() bool {
	return o != nil && o.RehydrationMaxScanSizeInGb.IsSet()
}

// SetRehydrationMaxScanSizeInGb gets a reference to the given datadog.NullableInt64 and assigns it to the RehydrationMaxScanSizeInGb field.
func (o *LogsArchiveCreateRequestAttributes) SetRehydrationMaxScanSizeInGb(v int64) {
	o.RehydrationMaxScanSizeInGb.Set(&v)
}

// SetRehydrationMaxScanSizeInGbNil sets the value for RehydrationMaxScanSizeInGb to be an explicit nil.
func (o *LogsArchiveCreateRequestAttributes) SetRehydrationMaxScanSizeInGbNil() {
	o.RehydrationMaxScanSizeInGb.Set(nil)
}

// UnsetRehydrationMaxScanSizeInGb ensures that no value is present for RehydrationMaxScanSizeInGb, not even an explicit nil.
func (o *LogsArchiveCreateRequestAttributes) UnsetRehydrationMaxScanSizeInGb() {
	o.RehydrationMaxScanSizeInGb.Unset()
}

// GetRehydrationTags returns the RehydrationTags field value if set, zero value otherwise.
func (o *LogsArchiveCreateRequestAttributes) GetRehydrationTags() []string {
	if o == nil || o.RehydrationTags == nil {
		var ret []string
		return ret
	}
	return o.RehydrationTags
}

// GetRehydrationTagsOk returns a tuple with the RehydrationTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveCreateRequestAttributes) GetRehydrationTagsOk() (*[]string, bool) {
	if o == nil || o.RehydrationTags == nil {
		return nil, false
	}
	return &o.RehydrationTags, true
}

// HasRehydrationTags returns a boolean if a field has been set.
func (o *LogsArchiveCreateRequestAttributes) HasRehydrationTags() bool {
	return o != nil && o.RehydrationTags != nil
}

// SetRehydrationTags gets a reference to the given []string and assigns it to the RehydrationTags field.
func (o *LogsArchiveCreateRequestAttributes) SetRehydrationTags(v []string) {
	o.RehydrationTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArchiveCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompressionMethod != nil {
		toSerialize["compression_method"] = o.CompressionMethod
	}
	toSerialize["destination"] = o.Destination
	if o.IncludeTags != nil {
		toSerialize["include_tags"] = o.IncludeTags
	}
	if o.LookupAttributes != nil {
		toSerialize["lookup_attributes"] = o.LookupAttributes
	}
	toSerialize["name"] = o.Name
	if o.PartitioningAttributes != nil {
		toSerialize["partitioning_attributes"] = o.PartitioningAttributes
	}
	toSerialize["query"] = o.Query
	if o.RehydrationMaxScanSizeInGb.IsSet() {
		toSerialize["rehydration_max_scan_size_in_gb"] = o.RehydrationMaxScanSizeInGb.Get()
	}
	if o.RehydrationTags != nil {
		toSerialize["rehydration_tags"] = o.RehydrationTags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArchiveCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompressionMethod          *LogsArchiveAttributesCompressionMethod `json:"compression_method,omitempty"`
		Destination                *LogsArchiveCreateRequestDestination    `json:"destination"`
		IncludeTags                *bool                                   `json:"include_tags,omitempty"`
		LookupAttributes           []string                                `json:"lookup_attributes,omitempty"`
		Name                       *string                                 `json:"name"`
		PartitioningAttributes     []string                                `json:"partitioning_attributes,omitempty"`
		Query                      *string                                 `json:"query"`
		RehydrationMaxScanSizeInGb datadog.NullableInt64                   `json:"rehydration_max_scan_size_in_gb,omitempty"`
		RehydrationTags            []string                                `json:"rehydration_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Destination == nil {
		return fmt.Errorf("required field destination missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"compression_method", "destination", "include_tags", "lookup_attributes", "name", "partitioning_attributes", "query", "rehydration_max_scan_size_in_gb", "rehydration_tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CompressionMethod != nil && !all.CompressionMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.CompressionMethod = all.CompressionMethod
	}
	o.Destination = *all.Destination
	o.IncludeTags = all.IncludeTags
	o.LookupAttributes = all.LookupAttributes
	o.Name = *all.Name
	o.PartitioningAttributes = all.PartitioningAttributes
	o.Query = *all.Query
	o.RehydrationMaxScanSizeInGb = all.RehydrationMaxScanSizeInGb
	o.RehydrationTags = all.RehydrationTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
