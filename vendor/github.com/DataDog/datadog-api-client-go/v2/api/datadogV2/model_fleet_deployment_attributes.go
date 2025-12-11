// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentAttributes Attributes of a deployment in the response.
type FleetDeploymentAttributes struct {
	// Ordered list of configuration file operations to perform on the target hosts.
	ConfigOperations []FleetDeploymentOperation `json:"config_operations,omitempty"`
	// Estimated completion time of the deployment as a Unix timestamp (seconds since epoch).
	EstimatedEndTimeUnix *int64 `json:"estimated_end_time_unix,omitempty"`
	// Query used to filter and select target hosts for the deployment. Uses the Datadog query syntax.
	FilterQuery *string `json:"filter_query,omitempty"`
	// Current high-level status of the deployment (for example, "pending", "running",
	// "completed", "failed").
	HighLevelStatus *string `json:"high_level_status,omitempty"`
	// Paginated list of hosts in this deployment with their individual statuses. Only included
	// when fetching a single deployment by ID. Use the `limit` and `page` query parameters to
	// navigate through pages. Pagination metadata is included in the response `meta.hosts` field.
	Hosts []FleetDeploymentHost `json:"hosts,omitempty"`
	// List of packages to deploy to target hosts. Present only for package upgrade deployments.
	Packages []FleetDeploymentPackage `json:"packages,omitempty"`
	// Total number of hosts targeted by this deployment.
	TotalHosts *int64 `json:"total_hosts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentAttributes instantiates a new FleetDeploymentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentAttributes() *FleetDeploymentAttributes {
	this := FleetDeploymentAttributes{}
	return &this
}

// NewFleetDeploymentAttributesWithDefaults instantiates a new FleetDeploymentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentAttributesWithDefaults() *FleetDeploymentAttributes {
	this := FleetDeploymentAttributes{}
	return &this
}

// GetConfigOperations returns the ConfigOperations field value if set, zero value otherwise.
func (o *FleetDeploymentAttributes) GetConfigOperations() []FleetDeploymentOperation {
	if o == nil || o.ConfigOperations == nil {
		var ret []FleetDeploymentOperation
		return ret
	}
	return o.ConfigOperations
}

// GetConfigOperationsOk returns a tuple with the ConfigOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentAttributes) GetConfigOperationsOk() (*[]FleetDeploymentOperation, bool) {
	if o == nil || o.ConfigOperations == nil {
		return nil, false
	}
	return &o.ConfigOperations, true
}

// HasConfigOperations returns a boolean if a field has been set.
func (o *FleetDeploymentAttributes) HasConfigOperations() bool {
	return o != nil && o.ConfigOperations != nil
}

// SetConfigOperations gets a reference to the given []FleetDeploymentOperation and assigns it to the ConfigOperations field.
func (o *FleetDeploymentAttributes) SetConfigOperations(v []FleetDeploymentOperation) {
	o.ConfigOperations = v
}

// GetEstimatedEndTimeUnix returns the EstimatedEndTimeUnix field value if set, zero value otherwise.
func (o *FleetDeploymentAttributes) GetEstimatedEndTimeUnix() int64 {
	if o == nil || o.EstimatedEndTimeUnix == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedEndTimeUnix
}

// GetEstimatedEndTimeUnixOk returns a tuple with the EstimatedEndTimeUnix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentAttributes) GetEstimatedEndTimeUnixOk() (*int64, bool) {
	if o == nil || o.EstimatedEndTimeUnix == nil {
		return nil, false
	}
	return o.EstimatedEndTimeUnix, true
}

// HasEstimatedEndTimeUnix returns a boolean if a field has been set.
func (o *FleetDeploymentAttributes) HasEstimatedEndTimeUnix() bool {
	return o != nil && o.EstimatedEndTimeUnix != nil
}

// SetEstimatedEndTimeUnix gets a reference to the given int64 and assigns it to the EstimatedEndTimeUnix field.
func (o *FleetDeploymentAttributes) SetEstimatedEndTimeUnix(v int64) {
	o.EstimatedEndTimeUnix = &v
}

// GetFilterQuery returns the FilterQuery field value if set, zero value otherwise.
func (o *FleetDeploymentAttributes) GetFilterQuery() string {
	if o == nil || o.FilterQuery == nil {
		var ret string
		return ret
	}
	return *o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentAttributes) GetFilterQueryOk() (*string, bool) {
	if o == nil || o.FilterQuery == nil {
		return nil, false
	}
	return o.FilterQuery, true
}

// HasFilterQuery returns a boolean if a field has been set.
func (o *FleetDeploymentAttributes) HasFilterQuery() bool {
	return o != nil && o.FilterQuery != nil
}

// SetFilterQuery gets a reference to the given string and assigns it to the FilterQuery field.
func (o *FleetDeploymentAttributes) SetFilterQuery(v string) {
	o.FilterQuery = &v
}

// GetHighLevelStatus returns the HighLevelStatus field value if set, zero value otherwise.
func (o *FleetDeploymentAttributes) GetHighLevelStatus() string {
	if o == nil || o.HighLevelStatus == nil {
		var ret string
		return ret
	}
	return *o.HighLevelStatus
}

// GetHighLevelStatusOk returns a tuple with the HighLevelStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentAttributes) GetHighLevelStatusOk() (*string, bool) {
	if o == nil || o.HighLevelStatus == nil {
		return nil, false
	}
	return o.HighLevelStatus, true
}

// HasHighLevelStatus returns a boolean if a field has been set.
func (o *FleetDeploymentAttributes) HasHighLevelStatus() bool {
	return o != nil && o.HighLevelStatus != nil
}

// SetHighLevelStatus gets a reference to the given string and assigns it to the HighLevelStatus field.
func (o *FleetDeploymentAttributes) SetHighLevelStatus(v string) {
	o.HighLevelStatus = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *FleetDeploymentAttributes) GetHosts() []FleetDeploymentHost {
	if o == nil || o.Hosts == nil {
		var ret []FleetDeploymentHost
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentAttributes) GetHostsOk() (*[]FleetDeploymentHost, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *FleetDeploymentAttributes) HasHosts() bool {
	return o != nil && o.Hosts != nil
}

// SetHosts gets a reference to the given []FleetDeploymentHost and assigns it to the Hosts field.
func (o *FleetDeploymentAttributes) SetHosts(v []FleetDeploymentHost) {
	o.Hosts = v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *FleetDeploymentAttributes) GetPackages() []FleetDeploymentPackage {
	if o == nil || o.Packages == nil {
		var ret []FleetDeploymentPackage
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentAttributes) GetPackagesOk() (*[]FleetDeploymentPackage, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return &o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *FleetDeploymentAttributes) HasPackages() bool {
	return o != nil && o.Packages != nil
}

// SetPackages gets a reference to the given []FleetDeploymentPackage and assigns it to the Packages field.
func (o *FleetDeploymentAttributes) SetPackages(v []FleetDeploymentPackage) {
	o.Packages = v
}

// GetTotalHosts returns the TotalHosts field value if set, zero value otherwise.
func (o *FleetDeploymentAttributes) GetTotalHosts() int64 {
	if o == nil || o.TotalHosts == nil {
		var ret int64
		return ret
	}
	return *o.TotalHosts
}

// GetTotalHostsOk returns a tuple with the TotalHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentAttributes) GetTotalHostsOk() (*int64, bool) {
	if o == nil || o.TotalHosts == nil {
		return nil, false
	}
	return o.TotalHosts, true
}

// HasTotalHosts returns a boolean if a field has been set.
func (o *FleetDeploymentAttributes) HasTotalHosts() bool {
	return o != nil && o.TotalHosts != nil
}

// SetTotalHosts gets a reference to the given int64 and assigns it to the TotalHosts field.
func (o *FleetDeploymentAttributes) SetTotalHosts(v int64) {
	o.TotalHosts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfigOperations != nil {
		toSerialize["config_operations"] = o.ConfigOperations
	}
	if o.EstimatedEndTimeUnix != nil {
		toSerialize["estimated_end_time_unix"] = o.EstimatedEndTimeUnix
	}
	if o.FilterQuery != nil {
		toSerialize["filter_query"] = o.FilterQuery
	}
	if o.HighLevelStatus != nil {
		toSerialize["high_level_status"] = o.HighLevelStatus
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	if o.TotalHosts != nil {
		toSerialize["total_hosts"] = o.TotalHosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigOperations     []FleetDeploymentOperation `json:"config_operations,omitempty"`
		EstimatedEndTimeUnix *int64                     `json:"estimated_end_time_unix,omitempty"`
		FilterQuery          *string                    `json:"filter_query,omitempty"`
		HighLevelStatus      *string                    `json:"high_level_status,omitempty"`
		Hosts                []FleetDeploymentHost      `json:"hosts,omitempty"`
		Packages             []FleetDeploymentPackage   `json:"packages,omitempty"`
		TotalHosts           *int64                     `json:"total_hosts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config_operations", "estimated_end_time_unix", "filter_query", "high_level_status", "hosts", "packages", "total_hosts"})
	} else {
		return err
	}
	o.ConfigOperations = all.ConfigOperations
	o.EstimatedEndTimeUnix = all.EstimatedEndTimeUnix
	o.FilterQuery = all.FilterQuery
	o.HighLevelStatus = all.HighLevelStatus
	o.Hosts = all.Hosts
	o.Packages = all.Packages
	o.TotalHosts = all.TotalHosts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
