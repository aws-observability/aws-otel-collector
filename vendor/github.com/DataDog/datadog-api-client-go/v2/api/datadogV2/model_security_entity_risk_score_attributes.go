// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityEntityRiskScoreAttributes Attributes of an entity risk score
type SecurityEntityRiskScoreAttributes struct {
	// Configuration risks associated with the entity
	ConfigRisks SecurityEntityConfigRisks `json:"configRisks"`
	// Unique identifier for the entity
	EntityId string `json:"entityID"`
	// Metadata about the entity from cloud providers
	EntityMetadata SecurityEntityMetadata `json:"entityMetadata"`
	// Human-readable name of the entity
	EntityName *string `json:"entityName,omitempty"`
	// Cloud providers associated with the entity
	EntityProviders []string `json:"entityProviders"`
	// Roles associated with the entity
	EntityRoles []string `json:"entityRoles,omitempty"`
	// Type of the entity (e.g., aws_iam_user, aws_ec2_instance)
	EntityType string `json:"entityType"`
	// Timestamp when the entity was first detected (Unix milliseconds)
	FirstDetected int64 `json:"firstDetected"`
	// Title of the most recent signal detected for this entity
	LastActivityTitle string `json:"lastActivityTitle"`
	// Timestamp when the entity was last detected (Unix milliseconds)
	LastDetected int64 `json:"lastDetected"`
	// Current risk score for the entity
	RiskScore float64 `json:"riskScore"`
	// Change in risk score compared to previous period
	RiskScoreEvolution float64 `json:"riskScoreEvolution"`
	// Severity level based on risk score
	Severity SecurityEntityRiskScoreAttributesSeverity `json:"severity"`
	// Number of security signals detected for this entity
	SignalsDetected int64 `json:"signalsDetected"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityEntityRiskScoreAttributes instantiates a new SecurityEntityRiskScoreAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityEntityRiskScoreAttributes(configRisks SecurityEntityConfigRisks, entityId string, entityMetadata SecurityEntityMetadata, entityProviders []string, entityType string, firstDetected int64, lastActivityTitle string, lastDetected int64, riskScore float64, riskScoreEvolution float64, severity SecurityEntityRiskScoreAttributesSeverity, signalsDetected int64) *SecurityEntityRiskScoreAttributes {
	this := SecurityEntityRiskScoreAttributes{}
	this.ConfigRisks = configRisks
	this.EntityId = entityId
	this.EntityMetadata = entityMetadata
	this.EntityProviders = entityProviders
	this.EntityType = entityType
	this.FirstDetected = firstDetected
	this.LastActivityTitle = lastActivityTitle
	this.LastDetected = lastDetected
	this.RiskScore = riskScore
	this.RiskScoreEvolution = riskScoreEvolution
	this.Severity = severity
	this.SignalsDetected = signalsDetected
	return &this
}

// NewSecurityEntityRiskScoreAttributesWithDefaults instantiates a new SecurityEntityRiskScoreAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityEntityRiskScoreAttributesWithDefaults() *SecurityEntityRiskScoreAttributes {
	this := SecurityEntityRiskScoreAttributes{}
	return &this
}

// GetConfigRisks returns the ConfigRisks field value.
func (o *SecurityEntityRiskScoreAttributes) GetConfigRisks() SecurityEntityConfigRisks {
	if o == nil {
		var ret SecurityEntityConfigRisks
		return ret
	}
	return o.ConfigRisks
}

// GetConfigRisksOk returns a tuple with the ConfigRisks field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetConfigRisksOk() (*SecurityEntityConfigRisks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigRisks, true
}

// SetConfigRisks sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetConfigRisks(v SecurityEntityConfigRisks) {
	o.ConfigRisks = v
}

// GetEntityId returns the EntityId field value.
func (o *SecurityEntityRiskScoreAttributes) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityMetadata returns the EntityMetadata field value.
func (o *SecurityEntityRiskScoreAttributes) GetEntityMetadata() SecurityEntityMetadata {
	if o == nil {
		var ret SecurityEntityMetadata
		return ret
	}
	return o.EntityMetadata
}

// GetEntityMetadataOk returns a tuple with the EntityMetadata field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntityMetadataOk() (*SecurityEntityMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityMetadata, true
}

// SetEntityMetadata sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetEntityMetadata(v SecurityEntityMetadata) {
	o.EntityMetadata = v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *SecurityEntityRiskScoreAttributes) GetEntityName() string {
	if o == nil || o.EntityName == nil {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntityNameOk() (*string, bool) {
	if o == nil || o.EntityName == nil {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *SecurityEntityRiskScoreAttributes) HasEntityName() bool {
	return o != nil && o.EntityName != nil
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *SecurityEntityRiskScoreAttributes) SetEntityName(v string) {
	o.EntityName = &v
}

// GetEntityProviders returns the EntityProviders field value.
func (o *SecurityEntityRiskScoreAttributes) GetEntityProviders() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EntityProviders
}

// GetEntityProvidersOk returns a tuple with the EntityProviders field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntityProvidersOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityProviders, true
}

// SetEntityProviders sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetEntityProviders(v []string) {
	o.EntityProviders = v
}

// GetEntityRoles returns the EntityRoles field value if set, zero value otherwise.
func (o *SecurityEntityRiskScoreAttributes) GetEntityRoles() []string {
	if o == nil || o.EntityRoles == nil {
		var ret []string
		return ret
	}
	return o.EntityRoles
}

// GetEntityRolesOk returns a tuple with the EntityRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntityRolesOk() (*[]string, bool) {
	if o == nil || o.EntityRoles == nil {
		return nil, false
	}
	return &o.EntityRoles, true
}

// HasEntityRoles returns a boolean if a field has been set.
func (o *SecurityEntityRiskScoreAttributes) HasEntityRoles() bool {
	return o != nil && o.EntityRoles != nil
}

// SetEntityRoles gets a reference to the given []string and assigns it to the EntityRoles field.
func (o *SecurityEntityRiskScoreAttributes) SetEntityRoles(v []string) {
	o.EntityRoles = v
}

// GetEntityType returns the EntityType field value.
func (o *SecurityEntityRiskScoreAttributes) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetEntityType(v string) {
	o.EntityType = v
}

// GetFirstDetected returns the FirstDetected field value.
func (o *SecurityEntityRiskScoreAttributes) GetFirstDetected() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.FirstDetected
}

// GetFirstDetectedOk returns a tuple with the FirstDetected field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetFirstDetectedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstDetected, true
}

// SetFirstDetected sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetFirstDetected(v int64) {
	o.FirstDetected = v
}

// GetLastActivityTitle returns the LastActivityTitle field value.
func (o *SecurityEntityRiskScoreAttributes) GetLastActivityTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LastActivityTitle
}

// GetLastActivityTitleOk returns a tuple with the LastActivityTitle field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetLastActivityTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastActivityTitle, true
}

// SetLastActivityTitle sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetLastActivityTitle(v string) {
	o.LastActivityTitle = v
}

// GetLastDetected returns the LastDetected field value.
func (o *SecurityEntityRiskScoreAttributes) GetLastDetected() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LastDetected
}

// GetLastDetectedOk returns a tuple with the LastDetected field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetLastDetectedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastDetected, true
}

// SetLastDetected sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetLastDetected(v int64) {
	o.LastDetected = v
}

// GetRiskScore returns the RiskScore field value.
func (o *SecurityEntityRiskScoreAttributes) GetRiskScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetRiskScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskScore, true
}

// SetRiskScore sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetRiskScore(v float64) {
	o.RiskScore = v
}

// GetRiskScoreEvolution returns the RiskScoreEvolution field value.
func (o *SecurityEntityRiskScoreAttributes) GetRiskScoreEvolution() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.RiskScoreEvolution
}

// GetRiskScoreEvolutionOk returns a tuple with the RiskScoreEvolution field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetRiskScoreEvolutionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskScoreEvolution, true
}

// SetRiskScoreEvolution sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetRiskScoreEvolution(v float64) {
	o.RiskScoreEvolution = v
}

// GetSeverity returns the Severity field value.
func (o *SecurityEntityRiskScoreAttributes) GetSeverity() SecurityEntityRiskScoreAttributesSeverity {
	if o == nil {
		var ret SecurityEntityRiskScoreAttributesSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetSeverityOk() (*SecurityEntityRiskScoreAttributesSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetSeverity(v SecurityEntityRiskScoreAttributesSeverity) {
	o.Severity = v
}

// GetSignalsDetected returns the SignalsDetected field value.
func (o *SecurityEntityRiskScoreAttributes) GetSignalsDetected() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SignalsDetected
}

// GetSignalsDetectedOk returns a tuple with the SignalsDetected field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetSignalsDetectedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalsDetected, true
}

// SetSignalsDetected sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetSignalsDetected(v int64) {
	o.SignalsDetected = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityEntityRiskScoreAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["configRisks"] = o.ConfigRisks
	toSerialize["entityID"] = o.EntityId
	toSerialize["entityMetadata"] = o.EntityMetadata
	if o.EntityName != nil {
		toSerialize["entityName"] = o.EntityName
	}
	toSerialize["entityProviders"] = o.EntityProviders
	if o.EntityRoles != nil {
		toSerialize["entityRoles"] = o.EntityRoles
	}
	toSerialize["entityType"] = o.EntityType
	toSerialize["firstDetected"] = o.FirstDetected
	toSerialize["lastActivityTitle"] = o.LastActivityTitle
	toSerialize["lastDetected"] = o.LastDetected
	toSerialize["riskScore"] = o.RiskScore
	toSerialize["riskScoreEvolution"] = o.RiskScoreEvolution
	toSerialize["severity"] = o.Severity
	toSerialize["signalsDetected"] = o.SignalsDetected

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityEntityRiskScoreAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigRisks        *SecurityEntityConfigRisks                 `json:"configRisks"`
		EntityId           *string                                    `json:"entityID"`
		EntityMetadata     *SecurityEntityMetadata                    `json:"entityMetadata"`
		EntityName         *string                                    `json:"entityName,omitempty"`
		EntityProviders    *[]string                                  `json:"entityProviders"`
		EntityRoles        []string                                   `json:"entityRoles,omitempty"`
		EntityType         *string                                    `json:"entityType"`
		FirstDetected      *int64                                     `json:"firstDetected"`
		LastActivityTitle  *string                                    `json:"lastActivityTitle"`
		LastDetected       *int64                                     `json:"lastDetected"`
		RiskScore          *float64                                   `json:"riskScore"`
		RiskScoreEvolution *float64                                   `json:"riskScoreEvolution"`
		Severity           *SecurityEntityRiskScoreAttributesSeverity `json:"severity"`
		SignalsDetected    *int64                                     `json:"signalsDetected"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigRisks == nil {
		return fmt.Errorf("required field configRisks missing")
	}
	if all.EntityId == nil {
		return fmt.Errorf("required field entityID missing")
	}
	if all.EntityMetadata == nil {
		return fmt.Errorf("required field entityMetadata missing")
	}
	if all.EntityProviders == nil {
		return fmt.Errorf("required field entityProviders missing")
	}
	if all.EntityType == nil {
		return fmt.Errorf("required field entityType missing")
	}
	if all.FirstDetected == nil {
		return fmt.Errorf("required field firstDetected missing")
	}
	if all.LastActivityTitle == nil {
		return fmt.Errorf("required field lastActivityTitle missing")
	}
	if all.LastDetected == nil {
		return fmt.Errorf("required field lastDetected missing")
	}
	if all.RiskScore == nil {
		return fmt.Errorf("required field riskScore missing")
	}
	if all.RiskScoreEvolution == nil {
		return fmt.Errorf("required field riskScoreEvolution missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.SignalsDetected == nil {
		return fmt.Errorf("required field signalsDetected missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configRisks", "entityID", "entityMetadata", "entityName", "entityProviders", "entityRoles", "entityType", "firstDetected", "lastActivityTitle", "lastDetected", "riskScore", "riskScoreEvolution", "severity", "signalsDetected"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConfigRisks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfigRisks = *all.ConfigRisks
	o.EntityId = *all.EntityId
	if all.EntityMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EntityMetadata = *all.EntityMetadata
	o.EntityName = all.EntityName
	o.EntityProviders = *all.EntityProviders
	o.EntityRoles = all.EntityRoles
	o.EntityType = *all.EntityType
	o.FirstDetected = *all.FirstDetected
	o.LastActivityTitle = *all.LastActivityTitle
	o.LastDetected = *all.LastDetected
	o.RiskScore = *all.RiskScore
	o.RiskScoreEvolution = *all.RiskScoreEvolution
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.SignalsDetected = *all.SignalsDetected

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
