// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListStreamQuery Updated list stream widget.
type ListStreamQuery struct {
	// Specifies the field for logs pattern clustering. Usable only with logs_pattern_stream.
	ClusteringPatternFieldPath *string `json:"clustering_pattern_field_path,omitempty"`
	// Compute configuration for the List Stream Widget. Compute can be used only with the logs_transaction_stream (from 1 to 5 items) list stream source.
	Compute []ListStreamComputeItems `json:"compute,omitempty"`
	// Source from which to query items to display in the stream.
	DataSource ListStreamSource `json:"data_source"`
	// Size to use to display an event.
	EventSize *WidgetEventSize `json:"event_size,omitempty"`
	// Group by configuration for the List Stream Widget. Group by can be used only with logs_pattern_stream (up to 4 items) or logs_transaction_stream (one group by item is required) list stream source.
	GroupBy []ListStreamGroupByItems `json:"group_by,omitempty"`
	// List of indexes.
	Indexes []string `json:"indexes,omitempty"`
	// Widget query.
	QueryString string `json:"query_string"`
	// Which column and order to sort by
	Sort *WidgetFieldSort `json:"sort,omitempty"`
	// Option for storage location. Feature in Private Beta.
	Storage *string `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListStreamQuery instantiates a new ListStreamQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListStreamQuery(dataSource ListStreamSource, queryString string) *ListStreamQuery {
	this := ListStreamQuery{}
	this.DataSource = dataSource
	this.QueryString = queryString
	return &this
}

// NewListStreamQueryWithDefaults instantiates a new ListStreamQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListStreamQueryWithDefaults() *ListStreamQuery {
	this := ListStreamQuery{}
	var dataSource ListStreamSource = LISTSTREAMSOURCE_APM_ISSUE_STREAM
	this.DataSource = dataSource
	return &this
}

// GetClusteringPatternFieldPath returns the ClusteringPatternFieldPath field value if set, zero value otherwise.
func (o *ListStreamQuery) GetClusteringPatternFieldPath() string {
	if o == nil || o.ClusteringPatternFieldPath == nil {
		var ret string
		return ret
	}
	return *o.ClusteringPatternFieldPath
}

// GetClusteringPatternFieldPathOk returns a tuple with the ClusteringPatternFieldPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetClusteringPatternFieldPathOk() (*string, bool) {
	if o == nil || o.ClusteringPatternFieldPath == nil {
		return nil, false
	}
	return o.ClusteringPatternFieldPath, true
}

// HasClusteringPatternFieldPath returns a boolean if a field has been set.
func (o *ListStreamQuery) HasClusteringPatternFieldPath() bool {
	return o != nil && o.ClusteringPatternFieldPath != nil
}

// SetClusteringPatternFieldPath gets a reference to the given string and assigns it to the ClusteringPatternFieldPath field.
func (o *ListStreamQuery) SetClusteringPatternFieldPath(v string) {
	o.ClusteringPatternFieldPath = &v
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *ListStreamQuery) GetCompute() []ListStreamComputeItems {
	if o == nil || o.Compute == nil {
		var ret []ListStreamComputeItems
		return ret
	}
	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetComputeOk() (*[]ListStreamComputeItems, bool) {
	if o == nil || o.Compute == nil {
		return nil, false
	}
	return &o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ListStreamQuery) HasCompute() bool {
	return o != nil && o.Compute != nil
}

// SetCompute gets a reference to the given []ListStreamComputeItems and assigns it to the Compute field.
func (o *ListStreamQuery) SetCompute(v []ListStreamComputeItems) {
	o.Compute = v
}

// GetDataSource returns the DataSource field value.
func (o *ListStreamQuery) GetDataSource() ListStreamSource {
	if o == nil {
		var ret ListStreamSource
		return ret
	}
	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetDataSourceOk() (*ListStreamSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value.
func (o *ListStreamQuery) SetDataSource(v ListStreamSource) {
	o.DataSource = v
}

// GetEventSize returns the EventSize field value if set, zero value otherwise.
func (o *ListStreamQuery) GetEventSize() WidgetEventSize {
	if o == nil || o.EventSize == nil {
		var ret WidgetEventSize
		return ret
	}
	return *o.EventSize
}

// GetEventSizeOk returns a tuple with the EventSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetEventSizeOk() (*WidgetEventSize, bool) {
	if o == nil || o.EventSize == nil {
		return nil, false
	}
	return o.EventSize, true
}

// HasEventSize returns a boolean if a field has been set.
func (o *ListStreamQuery) HasEventSize() bool {
	return o != nil && o.EventSize != nil
}

// SetEventSize gets a reference to the given WidgetEventSize and assigns it to the EventSize field.
func (o *ListStreamQuery) SetEventSize(v WidgetEventSize) {
	o.EventSize = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ListStreamQuery) GetGroupBy() []ListStreamGroupByItems {
	if o == nil || o.GroupBy == nil {
		var ret []ListStreamGroupByItems
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetGroupByOk() (*[]ListStreamGroupByItems, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ListStreamQuery) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []ListStreamGroupByItems and assigns it to the GroupBy field.
func (o *ListStreamQuery) SetGroupBy(v []ListStreamGroupByItems) {
	o.GroupBy = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *ListStreamQuery) GetIndexes() []string {
	if o == nil || o.Indexes == nil {
		var ret []string
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetIndexesOk() (*[]string, bool) {
	if o == nil || o.Indexes == nil {
		return nil, false
	}
	return &o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *ListStreamQuery) HasIndexes() bool {
	return o != nil && o.Indexes != nil
}

// SetIndexes gets a reference to the given []string and assigns it to the Indexes field.
func (o *ListStreamQuery) SetIndexes(v []string) {
	o.Indexes = v
}

// GetQueryString returns the QueryString field value.
func (o *ListStreamQuery) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value.
func (o *ListStreamQuery) SetQueryString(v string) {
	o.QueryString = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ListStreamQuery) GetSort() WidgetFieldSort {
	if o == nil || o.Sort == nil {
		var ret WidgetFieldSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetSortOk() (*WidgetFieldSort, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ListStreamQuery) HasSort() bool {
	return o != nil && o.Sort != nil
}

// SetSort gets a reference to the given WidgetFieldSort and assigns it to the Sort field.
func (o *ListStreamQuery) SetSort(v WidgetFieldSort) {
	o.Sort = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ListStreamQuery) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListStreamQuery) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ListStreamQuery) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *ListStreamQuery) SetStorage(v string) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListStreamQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ClusteringPatternFieldPath != nil {
		toSerialize["clustering_pattern_field_path"] = o.ClusteringPatternFieldPath
	}
	if o.Compute != nil {
		toSerialize["compute"] = o.Compute
	}
	toSerialize["data_source"] = o.DataSource
	if o.EventSize != nil {
		toSerialize["event_size"] = o.EventSize
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	if o.Indexes != nil {
		toSerialize["indexes"] = o.Indexes
	}
	toSerialize["query_string"] = o.QueryString
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListStreamQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusteringPatternFieldPath *string                  `json:"clustering_pattern_field_path,omitempty"`
		Compute                    []ListStreamComputeItems `json:"compute,omitempty"`
		DataSource                 *ListStreamSource        `json:"data_source"`
		EventSize                  *WidgetEventSize         `json:"event_size,omitempty"`
		GroupBy                    []ListStreamGroupByItems `json:"group_by,omitempty"`
		Indexes                    []string                 `json:"indexes,omitempty"`
		QueryString                *string                  `json:"query_string"`
		Sort                       *WidgetFieldSort         `json:"sort,omitempty"`
		Storage                    *string                  `json:"storage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataSource == nil {
		return fmt.Errorf("required field data_source missing")
	}
	if all.QueryString == nil {
		return fmt.Errorf("required field query_string missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"clustering_pattern_field_path", "compute", "data_source", "event_size", "group_by", "indexes", "query_string", "sort", "storage"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ClusteringPatternFieldPath = all.ClusteringPatternFieldPath
	o.Compute = all.Compute
	if !all.DataSource.IsValid() {
		hasInvalidField = true
	} else {
		o.DataSource = *all.DataSource
	}
	if all.EventSize != nil && !all.EventSize.IsValid() {
		hasInvalidField = true
	} else {
		o.EventSize = all.EventSize
	}
	o.GroupBy = all.GroupBy
	o.Indexes = all.Indexes
	o.QueryString = *all.QueryString
	if all.Sort != nil && all.Sort.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sort = all.Sort
	o.Storage = all.Storage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
