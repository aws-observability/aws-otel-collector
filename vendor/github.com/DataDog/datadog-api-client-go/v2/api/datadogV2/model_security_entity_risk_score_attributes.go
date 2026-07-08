// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityEntityRiskScoreAttributes Attributes of an entity risk score.
type SecurityEntityRiskScoreAttributes struct {
	// Cloud account IDs associated with the entity.
	AccountIds []string `json:"accountIds"`
	// Configuration risks associated with the entity
	ConfigRisks SecurityEntityConfigRisks `json:"configRisks"`
	// Metadata about the entity from cloud providers
	EntityMetadata SecurityEntityMetadata `json:"entityMetadata"`
	// Human-readable name of the entity.
	EntityName *string `json:"entityName,omitempty"`
	// Cloud providers associated with the entity.
	EntityProviders []string `json:"entityProviders"`
	// Roles associated with the entity.
	EntityRoles []string `json:"entityRoles,omitempty"`
	// Sub-types associated with the entity.
	EntitySubTypes []string `json:"entitySubTypes"`
	// Type of the entity (for example, aws_iam_user, aws_ec2_instance).
	EntityType *string `json:"entityType,omitempty"`
	// All types associated with the entity.
	EntityTypes []string `json:"entityTypes,omitempty"`
	// Timestamp when the entity was first detected (Unix milliseconds).
	FirstDetected int64 `json:"firstDetected"`
	// Title of the most recent signal detected for this entity.
	LastActivityTitle string `json:"lastActivityTitle"`
	// Timestamp when the entity was last detected (Unix milliseconds).
	LastDetected int64 `json:"lastDetected"`
	// Current risk score for the entity.
	RiskScore int64 `json:"riskScore"`
	// Change in risk score compared to previous period.
	RiskScoreEvolution int64 `json:"riskScoreEvolution"`
	// Severity level based on risk score
	Severity SecurityEntityRiskScoreAttributesSeverity `json:"severity"`
	// Number of security signals detected for this entity.
	SignalsDetected int64 `json:"signalsDetected"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityEntityRiskScoreAttributes instantiates a new SecurityEntityRiskScoreAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityEntityRiskScoreAttributes(accountIds []string, configRisks SecurityEntityConfigRisks, entityMetadata SecurityEntityMetadata, entityProviders []string, entitySubTypes []string, firstDetected int64, lastActivityTitle string, lastDetected int64, riskScore int64, riskScoreEvolution int64, severity SecurityEntityRiskScoreAttributesSeverity, signalsDetected int64) *SecurityEntityRiskScoreAttributes {
	this := SecurityEntityRiskScoreAttributes{}
	this.AccountIds = accountIds
	this.ConfigRisks = configRisks
	this.EntityMetadata = entityMetadata
	this.EntityProviders = entityProviders
	this.EntitySubTypes = entitySubTypes
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

// GetAccountIds returns the AccountIds field value.
func (o *SecurityEntityRiskScoreAttributes) GetAccountIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetAccountIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIds, true
}

// SetAccountIds sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetAccountIds(v []string) {
	o.AccountIds = v
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

// GetEntitySubTypes returns the EntitySubTypes field value.
func (o *SecurityEntityRiskScoreAttributes) GetEntitySubTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EntitySubTypes
}

// GetEntitySubTypesOk returns a tuple with the EntitySubTypes field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntitySubTypesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitySubTypes, true
}

// SetEntitySubTypes sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetEntitySubTypes(v []string) {
	o.EntitySubTypes = v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *SecurityEntityRiskScoreAttributes) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *SecurityEntityRiskScoreAttributes) HasEntityType() bool {
	return o != nil && o.EntityType != nil
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *SecurityEntityRiskScoreAttributes) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityTypes returns the EntityTypes field value if set, zero value otherwise.
func (o *SecurityEntityRiskScoreAttributes) GetEntityTypes() []string {
	if o == nil || o.EntityTypes == nil {
		var ret []string
		return ret
	}
	return o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetEntityTypesOk() (*[]string, bool) {
	if o == nil || o.EntityTypes == nil {
		return nil, false
	}
	return &o.EntityTypes, true
}

// HasEntityTypes returns a boolean if a field has been set.
func (o *SecurityEntityRiskScoreAttributes) HasEntityTypes() bool {
	return o != nil && o.EntityTypes != nil
}

// SetEntityTypes gets a reference to the given []string and assigns it to the EntityTypes field.
func (o *SecurityEntityRiskScoreAttributes) SetEntityTypes(v []string) {
	o.EntityTypes = v
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
func (o *SecurityEntityRiskScoreAttributes) GetRiskScore() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetRiskScoreOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskScore, true
}

// SetRiskScore sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetRiskScore(v int64) {
	o.RiskScore = v
}

// GetRiskScoreEvolution returns the RiskScoreEvolution field value.
func (o *SecurityEntityRiskScoreAttributes) GetRiskScoreEvolution() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.RiskScoreEvolution
}

// GetRiskScoreEvolutionOk returns a tuple with the RiskScoreEvolution field value
// and a boolean to check if the value has been set.
func (o *SecurityEntityRiskScoreAttributes) GetRiskScoreEvolutionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskScoreEvolution, true
}

// SetRiskScoreEvolution sets field value.
func (o *SecurityEntityRiskScoreAttributes) SetRiskScoreEvolution(v int64) {
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
	toSerialize["accountIds"] = o.AccountIds
	toSerialize["configRisks"] = o.ConfigRisks
	toSerialize["entityMetadata"] = o.EntityMetadata
	if o.EntityName != nil {
		toSerialize["entityName"] = o.EntityName
	}
	toSerialize["entityProviders"] = o.EntityProviders
	if o.EntityRoles != nil {
		toSerialize["entityRoles"] = o.EntityRoles
	}
	toSerialize["entitySubTypes"] = o.EntitySubTypes
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.EntityTypes != nil {
		toSerialize["entityTypes"] = o.EntityTypes
	}
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
		AccountIds         *[]string                                  `json:"accountIds"`
		ConfigRisks        *SecurityEntityConfigRisks                 `json:"configRisks"`
		EntityMetadata     *SecurityEntityMetadata                    `json:"entityMetadata"`
		EntityName         *string                                    `json:"entityName,omitempty"`
		EntityProviders    *[]string                                  `json:"entityProviders"`
		EntityRoles        []string                                   `json:"entityRoles,omitempty"`
		EntitySubTypes     *[]string                                  `json:"entitySubTypes"`
		EntityType         *string                                    `json:"entityType,omitempty"`
		EntityTypes        []string                                   `json:"entityTypes,omitempty"`
		FirstDetected      *int64                                     `json:"firstDetected"`
		LastActivityTitle  *string                                    `json:"lastActivityTitle"`
		LastDetected       *int64                                     `json:"lastDetected"`
		RiskScore          *int64                                     `json:"riskScore"`
		RiskScoreEvolution *int64                                     `json:"riskScoreEvolution"`
		Severity           *SecurityEntityRiskScoreAttributesSeverity `json:"severity"`
		SignalsDetected    *int64                                     `json:"signalsDetected"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountIds == nil {
		return fmt.Errorf("required field accountIds missing")
	}
	if all.ConfigRisks == nil {
		return fmt.Errorf("required field configRisks missing")
	}
	if all.EntityMetadata == nil {
		return fmt.Errorf("required field entityMetadata missing")
	}
	if all.EntityProviders == nil {
		return fmt.Errorf("required field entityProviders missing")
	}
	if all.EntitySubTypes == nil {
		return fmt.Errorf("required field entitySubTypes missing")
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accountIds", "configRisks", "entityMetadata", "entityName", "entityProviders", "entityRoles", "entitySubTypes", "entityType", "entityTypes", "firstDetected", "lastActivityTitle", "lastDetected", "riskScore", "riskScoreEvolution", "severity", "signalsDetected"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountIds = *all.AccountIds
	if all.ConfigRisks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ConfigRisks = *all.ConfigRisks
	if all.EntityMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EntityMetadata = *all.EntityMetadata
	o.EntityName = all.EntityName
	o.EntityProviders = *all.EntityProviders
	o.EntityRoles = all.EntityRoles
	o.EntitySubTypes = *all.EntitySubTypes
	o.EntityType = all.EntityType
	o.EntityTypes = all.EntityTypes
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
