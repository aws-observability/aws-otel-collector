// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionIntegrationUpdate - The definition of `ActionConnectionIntegrationUpdate` object.
type ActionConnectionIntegrationUpdate struct {
	AWSIntegrationUpdate          *AWSIntegrationUpdate
	AnthropicIntegrationUpdate    *AnthropicIntegrationUpdate
	AsanaIntegrationUpdate        *AsanaIntegrationUpdate
	AzureIntegrationUpdate        *AzureIntegrationUpdate
	CircleCIIntegrationUpdate     *CircleCIIntegrationUpdate
	ClickupIntegrationUpdate      *ClickupIntegrationUpdate
	CloudflareIntegrationUpdate   *CloudflareIntegrationUpdate
	ConfigCatIntegrationUpdate    *ConfigCatIntegrationUpdate
	DatadogIntegrationUpdate      *DatadogIntegrationUpdate
	FastlyIntegrationUpdate       *FastlyIntegrationUpdate
	FreshserviceIntegrationUpdate *FreshserviceIntegrationUpdate
	GCPIntegrationUpdate          *GCPIntegrationUpdate
	GeminiIntegrationUpdate       *GeminiIntegrationUpdate
	GitlabIntegrationUpdate       *GitlabIntegrationUpdate
	GreyNoiseIntegrationUpdate    *GreyNoiseIntegrationUpdate
	HTTPIntegrationUpdate         *HTTPIntegrationUpdate
	LaunchDarklyIntegrationUpdate *LaunchDarklyIntegrationUpdate
	NotionIntegrationUpdate       *NotionIntegrationUpdate
	OktaIntegrationUpdate         *OktaIntegrationUpdate
	OpenAIIntegrationUpdate       *OpenAIIntegrationUpdate
	ServiceNowIntegrationUpdate   *ServiceNowIntegrationUpdate
	SplitIntegrationUpdate        *SplitIntegrationUpdate
	StatsigIntegrationUpdate      *StatsigIntegrationUpdate
	VirusTotalIntegrationUpdate   *VirusTotalIntegrationUpdate

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns AWSIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func AWSIntegrationUpdateAsActionConnectionIntegrationUpdate(v *AWSIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{AWSIntegrationUpdate: v}
}

// AnthropicIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns AnthropicIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func AnthropicIntegrationUpdateAsActionConnectionIntegrationUpdate(v *AnthropicIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{AnthropicIntegrationUpdate: v}
}

// AsanaIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns AsanaIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func AsanaIntegrationUpdateAsActionConnectionIntegrationUpdate(v *AsanaIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{AsanaIntegrationUpdate: v}
}

// AzureIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns AzureIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func AzureIntegrationUpdateAsActionConnectionIntegrationUpdate(v *AzureIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{AzureIntegrationUpdate: v}
}

// CircleCIIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns CircleCIIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func CircleCIIntegrationUpdateAsActionConnectionIntegrationUpdate(v *CircleCIIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{CircleCIIntegrationUpdate: v}
}

// ClickupIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns ClickupIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func ClickupIntegrationUpdateAsActionConnectionIntegrationUpdate(v *ClickupIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{ClickupIntegrationUpdate: v}
}

// CloudflareIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns CloudflareIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func CloudflareIntegrationUpdateAsActionConnectionIntegrationUpdate(v *CloudflareIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{CloudflareIntegrationUpdate: v}
}

// ConfigCatIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns ConfigCatIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func ConfigCatIntegrationUpdateAsActionConnectionIntegrationUpdate(v *ConfigCatIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{ConfigCatIntegrationUpdate: v}
}

// DatadogIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns DatadogIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func DatadogIntegrationUpdateAsActionConnectionIntegrationUpdate(v *DatadogIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{DatadogIntegrationUpdate: v}
}

// FastlyIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns FastlyIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func FastlyIntegrationUpdateAsActionConnectionIntegrationUpdate(v *FastlyIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{FastlyIntegrationUpdate: v}
}

// FreshserviceIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns FreshserviceIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func FreshserviceIntegrationUpdateAsActionConnectionIntegrationUpdate(v *FreshserviceIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{FreshserviceIntegrationUpdate: v}
}

// GCPIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns GCPIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func GCPIntegrationUpdateAsActionConnectionIntegrationUpdate(v *GCPIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{GCPIntegrationUpdate: v}
}

// GeminiIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns GeminiIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func GeminiIntegrationUpdateAsActionConnectionIntegrationUpdate(v *GeminiIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{GeminiIntegrationUpdate: v}
}

// GitlabIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns GitlabIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func GitlabIntegrationUpdateAsActionConnectionIntegrationUpdate(v *GitlabIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{GitlabIntegrationUpdate: v}
}

// GreyNoiseIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns GreyNoiseIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func GreyNoiseIntegrationUpdateAsActionConnectionIntegrationUpdate(v *GreyNoiseIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{GreyNoiseIntegrationUpdate: v}
}

// HTTPIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns HTTPIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func HTTPIntegrationUpdateAsActionConnectionIntegrationUpdate(v *HTTPIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{HTTPIntegrationUpdate: v}
}

// LaunchDarklyIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns LaunchDarklyIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func LaunchDarklyIntegrationUpdateAsActionConnectionIntegrationUpdate(v *LaunchDarklyIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{LaunchDarklyIntegrationUpdate: v}
}

// NotionIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns NotionIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func NotionIntegrationUpdateAsActionConnectionIntegrationUpdate(v *NotionIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{NotionIntegrationUpdate: v}
}

// OktaIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns OktaIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func OktaIntegrationUpdateAsActionConnectionIntegrationUpdate(v *OktaIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{OktaIntegrationUpdate: v}
}

// OpenAIIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns OpenAIIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func OpenAIIntegrationUpdateAsActionConnectionIntegrationUpdate(v *OpenAIIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{OpenAIIntegrationUpdate: v}
}

// ServiceNowIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns ServiceNowIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func ServiceNowIntegrationUpdateAsActionConnectionIntegrationUpdate(v *ServiceNowIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{ServiceNowIntegrationUpdate: v}
}

// SplitIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns SplitIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func SplitIntegrationUpdateAsActionConnectionIntegrationUpdate(v *SplitIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{SplitIntegrationUpdate: v}
}

// StatsigIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns StatsigIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func StatsigIntegrationUpdateAsActionConnectionIntegrationUpdate(v *StatsigIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{StatsigIntegrationUpdate: v}
}

// VirusTotalIntegrationUpdateAsActionConnectionIntegrationUpdate is a convenience function that returns VirusTotalIntegrationUpdate wrapped in ActionConnectionIntegrationUpdate.
func VirusTotalIntegrationUpdateAsActionConnectionIntegrationUpdate(v *VirusTotalIntegrationUpdate) ActionConnectionIntegrationUpdate {
	return ActionConnectionIntegrationUpdate{VirusTotalIntegrationUpdate: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionConnectionIntegrationUpdate) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.AWSIntegrationUpdate)
	if err == nil {
		if obj.AWSIntegrationUpdate != nil && obj.AWSIntegrationUpdate.UnparsedObject == nil {
			jsonAWSIntegrationUpdate, _ := datadog.Marshal(obj.AWSIntegrationUpdate)
			if string(jsonAWSIntegrationUpdate) == "{}" { // empty struct
				obj.AWSIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AWSIntegrationUpdate = nil
		}
	} else {
		obj.AWSIntegrationUpdate = nil
	}

	// try to unmarshal data into AnthropicIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.AnthropicIntegrationUpdate)
	if err == nil {
		if obj.AnthropicIntegrationUpdate != nil && obj.AnthropicIntegrationUpdate.UnparsedObject == nil {
			jsonAnthropicIntegrationUpdate, _ := datadog.Marshal(obj.AnthropicIntegrationUpdate)
			if string(jsonAnthropicIntegrationUpdate) == "{}" { // empty struct
				obj.AnthropicIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AnthropicIntegrationUpdate = nil
		}
	} else {
		obj.AnthropicIntegrationUpdate = nil
	}

	// try to unmarshal data into AsanaIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.AsanaIntegrationUpdate)
	if err == nil {
		if obj.AsanaIntegrationUpdate != nil && obj.AsanaIntegrationUpdate.UnparsedObject == nil {
			jsonAsanaIntegrationUpdate, _ := datadog.Marshal(obj.AsanaIntegrationUpdate)
			if string(jsonAsanaIntegrationUpdate) == "{}" { // empty struct
				obj.AsanaIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AsanaIntegrationUpdate = nil
		}
	} else {
		obj.AsanaIntegrationUpdate = nil
	}

	// try to unmarshal data into AzureIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.AzureIntegrationUpdate)
	if err == nil {
		if obj.AzureIntegrationUpdate != nil && obj.AzureIntegrationUpdate.UnparsedObject == nil {
			jsonAzureIntegrationUpdate, _ := datadog.Marshal(obj.AzureIntegrationUpdate)
			if string(jsonAzureIntegrationUpdate) == "{}" { // empty struct
				obj.AzureIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.AzureIntegrationUpdate = nil
		}
	} else {
		obj.AzureIntegrationUpdate = nil
	}

	// try to unmarshal data into CircleCIIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.CircleCIIntegrationUpdate)
	if err == nil {
		if obj.CircleCIIntegrationUpdate != nil && obj.CircleCIIntegrationUpdate.UnparsedObject == nil {
			jsonCircleCIIntegrationUpdate, _ := datadog.Marshal(obj.CircleCIIntegrationUpdate)
			if string(jsonCircleCIIntegrationUpdate) == "{}" { // empty struct
				obj.CircleCIIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.CircleCIIntegrationUpdate = nil
		}
	} else {
		obj.CircleCIIntegrationUpdate = nil
	}

	// try to unmarshal data into ClickupIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.ClickupIntegrationUpdate)
	if err == nil {
		if obj.ClickupIntegrationUpdate != nil && obj.ClickupIntegrationUpdate.UnparsedObject == nil {
			jsonClickupIntegrationUpdate, _ := datadog.Marshal(obj.ClickupIntegrationUpdate)
			if string(jsonClickupIntegrationUpdate) == "{}" { // empty struct
				obj.ClickupIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ClickupIntegrationUpdate = nil
		}
	} else {
		obj.ClickupIntegrationUpdate = nil
	}

	// try to unmarshal data into CloudflareIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.CloudflareIntegrationUpdate)
	if err == nil {
		if obj.CloudflareIntegrationUpdate != nil && obj.CloudflareIntegrationUpdate.UnparsedObject == nil {
			jsonCloudflareIntegrationUpdate, _ := datadog.Marshal(obj.CloudflareIntegrationUpdate)
			if string(jsonCloudflareIntegrationUpdate) == "{}" { // empty struct
				obj.CloudflareIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.CloudflareIntegrationUpdate = nil
		}
	} else {
		obj.CloudflareIntegrationUpdate = nil
	}

	// try to unmarshal data into ConfigCatIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.ConfigCatIntegrationUpdate)
	if err == nil {
		if obj.ConfigCatIntegrationUpdate != nil && obj.ConfigCatIntegrationUpdate.UnparsedObject == nil {
			jsonConfigCatIntegrationUpdate, _ := datadog.Marshal(obj.ConfigCatIntegrationUpdate)
			if string(jsonConfigCatIntegrationUpdate) == "{}" { // empty struct
				obj.ConfigCatIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ConfigCatIntegrationUpdate = nil
		}
	} else {
		obj.ConfigCatIntegrationUpdate = nil
	}

	// try to unmarshal data into DatadogIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.DatadogIntegrationUpdate)
	if err == nil {
		if obj.DatadogIntegrationUpdate != nil && obj.DatadogIntegrationUpdate.UnparsedObject == nil {
			jsonDatadogIntegrationUpdate, _ := datadog.Marshal(obj.DatadogIntegrationUpdate)
			if string(jsonDatadogIntegrationUpdate) == "{}" { // empty struct
				obj.DatadogIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.DatadogIntegrationUpdate = nil
		}
	} else {
		obj.DatadogIntegrationUpdate = nil
	}

	// try to unmarshal data into FastlyIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.FastlyIntegrationUpdate)
	if err == nil {
		if obj.FastlyIntegrationUpdate != nil && obj.FastlyIntegrationUpdate.UnparsedObject == nil {
			jsonFastlyIntegrationUpdate, _ := datadog.Marshal(obj.FastlyIntegrationUpdate)
			if string(jsonFastlyIntegrationUpdate) == "{}" { // empty struct
				obj.FastlyIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.FastlyIntegrationUpdate = nil
		}
	} else {
		obj.FastlyIntegrationUpdate = nil
	}

	// try to unmarshal data into FreshserviceIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.FreshserviceIntegrationUpdate)
	if err == nil {
		if obj.FreshserviceIntegrationUpdate != nil && obj.FreshserviceIntegrationUpdate.UnparsedObject == nil {
			jsonFreshserviceIntegrationUpdate, _ := datadog.Marshal(obj.FreshserviceIntegrationUpdate)
			if string(jsonFreshserviceIntegrationUpdate) == "{}" { // empty struct
				obj.FreshserviceIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.FreshserviceIntegrationUpdate = nil
		}
	} else {
		obj.FreshserviceIntegrationUpdate = nil
	}

	// try to unmarshal data into GCPIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.GCPIntegrationUpdate)
	if err == nil {
		if obj.GCPIntegrationUpdate != nil && obj.GCPIntegrationUpdate.UnparsedObject == nil {
			jsonGCPIntegrationUpdate, _ := datadog.Marshal(obj.GCPIntegrationUpdate)
			if string(jsonGCPIntegrationUpdate) == "{}" { // empty struct
				obj.GCPIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GCPIntegrationUpdate = nil
		}
	} else {
		obj.GCPIntegrationUpdate = nil
	}

	// try to unmarshal data into GeminiIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.GeminiIntegrationUpdate)
	if err == nil {
		if obj.GeminiIntegrationUpdate != nil && obj.GeminiIntegrationUpdate.UnparsedObject == nil {
			jsonGeminiIntegrationUpdate, _ := datadog.Marshal(obj.GeminiIntegrationUpdate)
			if string(jsonGeminiIntegrationUpdate) == "{}" { // empty struct
				obj.GeminiIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GeminiIntegrationUpdate = nil
		}
	} else {
		obj.GeminiIntegrationUpdate = nil
	}

	// try to unmarshal data into GitlabIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.GitlabIntegrationUpdate)
	if err == nil {
		if obj.GitlabIntegrationUpdate != nil && obj.GitlabIntegrationUpdate.UnparsedObject == nil {
			jsonGitlabIntegrationUpdate, _ := datadog.Marshal(obj.GitlabIntegrationUpdate)
			if string(jsonGitlabIntegrationUpdate) == "{}" { // empty struct
				obj.GitlabIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GitlabIntegrationUpdate = nil
		}
	} else {
		obj.GitlabIntegrationUpdate = nil
	}

	// try to unmarshal data into GreyNoiseIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.GreyNoiseIntegrationUpdate)
	if err == nil {
		if obj.GreyNoiseIntegrationUpdate != nil && obj.GreyNoiseIntegrationUpdate.UnparsedObject == nil {
			jsonGreyNoiseIntegrationUpdate, _ := datadog.Marshal(obj.GreyNoiseIntegrationUpdate)
			if string(jsonGreyNoiseIntegrationUpdate) == "{}" { // empty struct
				obj.GreyNoiseIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.GreyNoiseIntegrationUpdate = nil
		}
	} else {
		obj.GreyNoiseIntegrationUpdate = nil
	}

	// try to unmarshal data into HTTPIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.HTTPIntegrationUpdate)
	if err == nil {
		if obj.HTTPIntegrationUpdate != nil && obj.HTTPIntegrationUpdate.UnparsedObject == nil {
			jsonHTTPIntegrationUpdate, _ := datadog.Marshal(obj.HTTPIntegrationUpdate)
			if string(jsonHTTPIntegrationUpdate) == "{}" { // empty struct
				obj.HTTPIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.HTTPIntegrationUpdate = nil
		}
	} else {
		obj.HTTPIntegrationUpdate = nil
	}

	// try to unmarshal data into LaunchDarklyIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.LaunchDarklyIntegrationUpdate)
	if err == nil {
		if obj.LaunchDarklyIntegrationUpdate != nil && obj.LaunchDarklyIntegrationUpdate.UnparsedObject == nil {
			jsonLaunchDarklyIntegrationUpdate, _ := datadog.Marshal(obj.LaunchDarklyIntegrationUpdate)
			if string(jsonLaunchDarklyIntegrationUpdate) == "{}" { // empty struct
				obj.LaunchDarklyIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.LaunchDarklyIntegrationUpdate = nil
		}
	} else {
		obj.LaunchDarklyIntegrationUpdate = nil
	}

	// try to unmarshal data into NotionIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.NotionIntegrationUpdate)
	if err == nil {
		if obj.NotionIntegrationUpdate != nil && obj.NotionIntegrationUpdate.UnparsedObject == nil {
			jsonNotionIntegrationUpdate, _ := datadog.Marshal(obj.NotionIntegrationUpdate)
			if string(jsonNotionIntegrationUpdate) == "{}" { // empty struct
				obj.NotionIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.NotionIntegrationUpdate = nil
		}
	} else {
		obj.NotionIntegrationUpdate = nil
	}

	// try to unmarshal data into OktaIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.OktaIntegrationUpdate)
	if err == nil {
		if obj.OktaIntegrationUpdate != nil && obj.OktaIntegrationUpdate.UnparsedObject == nil {
			jsonOktaIntegrationUpdate, _ := datadog.Marshal(obj.OktaIntegrationUpdate)
			if string(jsonOktaIntegrationUpdate) == "{}" { // empty struct
				obj.OktaIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.OktaIntegrationUpdate = nil
		}
	} else {
		obj.OktaIntegrationUpdate = nil
	}

	// try to unmarshal data into OpenAIIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.OpenAIIntegrationUpdate)
	if err == nil {
		if obj.OpenAIIntegrationUpdate != nil && obj.OpenAIIntegrationUpdate.UnparsedObject == nil {
			jsonOpenAIIntegrationUpdate, _ := datadog.Marshal(obj.OpenAIIntegrationUpdate)
			if string(jsonOpenAIIntegrationUpdate) == "{}" { // empty struct
				obj.OpenAIIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIIntegrationUpdate = nil
		}
	} else {
		obj.OpenAIIntegrationUpdate = nil
	}

	// try to unmarshal data into ServiceNowIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.ServiceNowIntegrationUpdate)
	if err == nil {
		if obj.ServiceNowIntegrationUpdate != nil && obj.ServiceNowIntegrationUpdate.UnparsedObject == nil {
			jsonServiceNowIntegrationUpdate, _ := datadog.Marshal(obj.ServiceNowIntegrationUpdate)
			if string(jsonServiceNowIntegrationUpdate) == "{}" { // empty struct
				obj.ServiceNowIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.ServiceNowIntegrationUpdate = nil
		}
	} else {
		obj.ServiceNowIntegrationUpdate = nil
	}

	// try to unmarshal data into SplitIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.SplitIntegrationUpdate)
	if err == nil {
		if obj.SplitIntegrationUpdate != nil && obj.SplitIntegrationUpdate.UnparsedObject == nil {
			jsonSplitIntegrationUpdate, _ := datadog.Marshal(obj.SplitIntegrationUpdate)
			if string(jsonSplitIntegrationUpdate) == "{}" { // empty struct
				obj.SplitIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.SplitIntegrationUpdate = nil
		}
	} else {
		obj.SplitIntegrationUpdate = nil
	}

	// try to unmarshal data into StatsigIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.StatsigIntegrationUpdate)
	if err == nil {
		if obj.StatsigIntegrationUpdate != nil && obj.StatsigIntegrationUpdate.UnparsedObject == nil {
			jsonStatsigIntegrationUpdate, _ := datadog.Marshal(obj.StatsigIntegrationUpdate)
			if string(jsonStatsigIntegrationUpdate) == "{}" { // empty struct
				obj.StatsigIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.StatsigIntegrationUpdate = nil
		}
	} else {
		obj.StatsigIntegrationUpdate = nil
	}

	// try to unmarshal data into VirusTotalIntegrationUpdate
	err = datadog.Unmarshal(data, &obj.VirusTotalIntegrationUpdate)
	if err == nil {
		if obj.VirusTotalIntegrationUpdate != nil && obj.VirusTotalIntegrationUpdate.UnparsedObject == nil {
			jsonVirusTotalIntegrationUpdate, _ := datadog.Marshal(obj.VirusTotalIntegrationUpdate)
			if string(jsonVirusTotalIntegrationUpdate) == "{}" { // empty struct
				obj.VirusTotalIntegrationUpdate = nil
			} else {
				match++
			}
		} else {
			obj.VirusTotalIntegrationUpdate = nil
		}
	} else {
		obj.VirusTotalIntegrationUpdate = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSIntegrationUpdate = nil
		obj.AnthropicIntegrationUpdate = nil
		obj.AsanaIntegrationUpdate = nil
		obj.AzureIntegrationUpdate = nil
		obj.CircleCIIntegrationUpdate = nil
		obj.ClickupIntegrationUpdate = nil
		obj.CloudflareIntegrationUpdate = nil
		obj.ConfigCatIntegrationUpdate = nil
		obj.DatadogIntegrationUpdate = nil
		obj.FastlyIntegrationUpdate = nil
		obj.FreshserviceIntegrationUpdate = nil
		obj.GCPIntegrationUpdate = nil
		obj.GeminiIntegrationUpdate = nil
		obj.GitlabIntegrationUpdate = nil
		obj.GreyNoiseIntegrationUpdate = nil
		obj.HTTPIntegrationUpdate = nil
		obj.LaunchDarklyIntegrationUpdate = nil
		obj.NotionIntegrationUpdate = nil
		obj.OktaIntegrationUpdate = nil
		obj.OpenAIIntegrationUpdate = nil
		obj.ServiceNowIntegrationUpdate = nil
		obj.SplitIntegrationUpdate = nil
		obj.StatsigIntegrationUpdate = nil
		obj.VirusTotalIntegrationUpdate = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionConnectionIntegrationUpdate) MarshalJSON() ([]byte, error) {
	if obj.AWSIntegrationUpdate != nil {
		return datadog.Marshal(&obj.AWSIntegrationUpdate)
	}

	if obj.AnthropicIntegrationUpdate != nil {
		return datadog.Marshal(&obj.AnthropicIntegrationUpdate)
	}

	if obj.AsanaIntegrationUpdate != nil {
		return datadog.Marshal(&obj.AsanaIntegrationUpdate)
	}

	if obj.AzureIntegrationUpdate != nil {
		return datadog.Marshal(&obj.AzureIntegrationUpdate)
	}

	if obj.CircleCIIntegrationUpdate != nil {
		return datadog.Marshal(&obj.CircleCIIntegrationUpdate)
	}

	if obj.ClickupIntegrationUpdate != nil {
		return datadog.Marshal(&obj.ClickupIntegrationUpdate)
	}

	if obj.CloudflareIntegrationUpdate != nil {
		return datadog.Marshal(&obj.CloudflareIntegrationUpdate)
	}

	if obj.ConfigCatIntegrationUpdate != nil {
		return datadog.Marshal(&obj.ConfigCatIntegrationUpdate)
	}

	if obj.DatadogIntegrationUpdate != nil {
		return datadog.Marshal(&obj.DatadogIntegrationUpdate)
	}

	if obj.FastlyIntegrationUpdate != nil {
		return datadog.Marshal(&obj.FastlyIntegrationUpdate)
	}

	if obj.FreshserviceIntegrationUpdate != nil {
		return datadog.Marshal(&obj.FreshserviceIntegrationUpdate)
	}

	if obj.GCPIntegrationUpdate != nil {
		return datadog.Marshal(&obj.GCPIntegrationUpdate)
	}

	if obj.GeminiIntegrationUpdate != nil {
		return datadog.Marshal(&obj.GeminiIntegrationUpdate)
	}

	if obj.GitlabIntegrationUpdate != nil {
		return datadog.Marshal(&obj.GitlabIntegrationUpdate)
	}

	if obj.GreyNoiseIntegrationUpdate != nil {
		return datadog.Marshal(&obj.GreyNoiseIntegrationUpdate)
	}

	if obj.HTTPIntegrationUpdate != nil {
		return datadog.Marshal(&obj.HTTPIntegrationUpdate)
	}

	if obj.LaunchDarklyIntegrationUpdate != nil {
		return datadog.Marshal(&obj.LaunchDarklyIntegrationUpdate)
	}

	if obj.NotionIntegrationUpdate != nil {
		return datadog.Marshal(&obj.NotionIntegrationUpdate)
	}

	if obj.OktaIntegrationUpdate != nil {
		return datadog.Marshal(&obj.OktaIntegrationUpdate)
	}

	if obj.OpenAIIntegrationUpdate != nil {
		return datadog.Marshal(&obj.OpenAIIntegrationUpdate)
	}

	if obj.ServiceNowIntegrationUpdate != nil {
		return datadog.Marshal(&obj.ServiceNowIntegrationUpdate)
	}

	if obj.SplitIntegrationUpdate != nil {
		return datadog.Marshal(&obj.SplitIntegrationUpdate)
	}

	if obj.StatsigIntegrationUpdate != nil {
		return datadog.Marshal(&obj.StatsigIntegrationUpdate)
	}

	if obj.VirusTotalIntegrationUpdate != nil {
		return datadog.Marshal(&obj.VirusTotalIntegrationUpdate)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ActionConnectionIntegrationUpdate) GetActualInstance() interface{} {
	if obj.AWSIntegrationUpdate != nil {
		return obj.AWSIntegrationUpdate
	}

	if obj.AnthropicIntegrationUpdate != nil {
		return obj.AnthropicIntegrationUpdate
	}

	if obj.AsanaIntegrationUpdate != nil {
		return obj.AsanaIntegrationUpdate
	}

	if obj.AzureIntegrationUpdate != nil {
		return obj.AzureIntegrationUpdate
	}

	if obj.CircleCIIntegrationUpdate != nil {
		return obj.CircleCIIntegrationUpdate
	}

	if obj.ClickupIntegrationUpdate != nil {
		return obj.ClickupIntegrationUpdate
	}

	if obj.CloudflareIntegrationUpdate != nil {
		return obj.CloudflareIntegrationUpdate
	}

	if obj.ConfigCatIntegrationUpdate != nil {
		return obj.ConfigCatIntegrationUpdate
	}

	if obj.DatadogIntegrationUpdate != nil {
		return obj.DatadogIntegrationUpdate
	}

	if obj.FastlyIntegrationUpdate != nil {
		return obj.FastlyIntegrationUpdate
	}

	if obj.FreshserviceIntegrationUpdate != nil {
		return obj.FreshserviceIntegrationUpdate
	}

	if obj.GCPIntegrationUpdate != nil {
		return obj.GCPIntegrationUpdate
	}

	if obj.GeminiIntegrationUpdate != nil {
		return obj.GeminiIntegrationUpdate
	}

	if obj.GitlabIntegrationUpdate != nil {
		return obj.GitlabIntegrationUpdate
	}

	if obj.GreyNoiseIntegrationUpdate != nil {
		return obj.GreyNoiseIntegrationUpdate
	}

	if obj.HTTPIntegrationUpdate != nil {
		return obj.HTTPIntegrationUpdate
	}

	if obj.LaunchDarklyIntegrationUpdate != nil {
		return obj.LaunchDarklyIntegrationUpdate
	}

	if obj.NotionIntegrationUpdate != nil {
		return obj.NotionIntegrationUpdate
	}

	if obj.OktaIntegrationUpdate != nil {
		return obj.OktaIntegrationUpdate
	}

	if obj.OpenAIIntegrationUpdate != nil {
		return obj.OpenAIIntegrationUpdate
	}

	if obj.ServiceNowIntegrationUpdate != nil {
		return obj.ServiceNowIntegrationUpdate
	}

	if obj.SplitIntegrationUpdate != nil {
		return obj.SplitIntegrationUpdate
	}

	if obj.StatsigIntegrationUpdate != nil {
		return obj.StatsigIntegrationUpdate
	}

	if obj.VirusTotalIntegrationUpdate != nil {
		return obj.VirusTotalIntegrationUpdate
	}

	// all schemas are nil
	return nil
}
