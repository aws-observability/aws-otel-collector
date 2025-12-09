// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionConnectionIntegration - The definition of `ActionConnectionIntegration` object.
type ActionConnectionIntegration struct {
	AWSIntegration          *AWSIntegration
	AnthropicIntegration    *AnthropicIntegration
	AsanaIntegration        *AsanaIntegration
	AzureIntegration        *AzureIntegration
	CircleCIIntegration     *CircleCIIntegration
	ClickupIntegration      *ClickupIntegration
	CloudflareIntegration   *CloudflareIntegration
	ConfigCatIntegration    *ConfigCatIntegration
	DatadogIntegration      *DatadogIntegration
	FastlyIntegration       *FastlyIntegration
	FreshserviceIntegration *FreshserviceIntegration
	GCPIntegration          *GCPIntegration
	GeminiIntegration       *GeminiIntegration
	GitlabIntegration       *GitlabIntegration
	GreyNoiseIntegration    *GreyNoiseIntegration
	HTTPIntegration         *HTTPIntegration
	LaunchDarklyIntegration *LaunchDarklyIntegration
	NotionIntegration       *NotionIntegration
	OktaIntegration         *OktaIntegration
	OpenAIIntegration       *OpenAIIntegration
	ServiceNowIntegration   *ServiceNowIntegration
	SplitIntegration        *SplitIntegration
	StatsigIntegration      *StatsigIntegration
	VirusTotalIntegration   *VirusTotalIntegration

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AWSIntegrationAsActionConnectionIntegration is a convenience function that returns AWSIntegration wrapped in ActionConnectionIntegration.
func AWSIntegrationAsActionConnectionIntegration(v *AWSIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{AWSIntegration: v}
}

// AnthropicIntegrationAsActionConnectionIntegration is a convenience function that returns AnthropicIntegration wrapped in ActionConnectionIntegration.
func AnthropicIntegrationAsActionConnectionIntegration(v *AnthropicIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{AnthropicIntegration: v}
}

// AsanaIntegrationAsActionConnectionIntegration is a convenience function that returns AsanaIntegration wrapped in ActionConnectionIntegration.
func AsanaIntegrationAsActionConnectionIntegration(v *AsanaIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{AsanaIntegration: v}
}

// AzureIntegrationAsActionConnectionIntegration is a convenience function that returns AzureIntegration wrapped in ActionConnectionIntegration.
func AzureIntegrationAsActionConnectionIntegration(v *AzureIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{AzureIntegration: v}
}

// CircleCIIntegrationAsActionConnectionIntegration is a convenience function that returns CircleCIIntegration wrapped in ActionConnectionIntegration.
func CircleCIIntegrationAsActionConnectionIntegration(v *CircleCIIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{CircleCIIntegration: v}
}

// ClickupIntegrationAsActionConnectionIntegration is a convenience function that returns ClickupIntegration wrapped in ActionConnectionIntegration.
func ClickupIntegrationAsActionConnectionIntegration(v *ClickupIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{ClickupIntegration: v}
}

// CloudflareIntegrationAsActionConnectionIntegration is a convenience function that returns CloudflareIntegration wrapped in ActionConnectionIntegration.
func CloudflareIntegrationAsActionConnectionIntegration(v *CloudflareIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{CloudflareIntegration: v}
}

// ConfigCatIntegrationAsActionConnectionIntegration is a convenience function that returns ConfigCatIntegration wrapped in ActionConnectionIntegration.
func ConfigCatIntegrationAsActionConnectionIntegration(v *ConfigCatIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{ConfigCatIntegration: v}
}

// DatadogIntegrationAsActionConnectionIntegration is a convenience function that returns DatadogIntegration wrapped in ActionConnectionIntegration.
func DatadogIntegrationAsActionConnectionIntegration(v *DatadogIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{DatadogIntegration: v}
}

// FastlyIntegrationAsActionConnectionIntegration is a convenience function that returns FastlyIntegration wrapped in ActionConnectionIntegration.
func FastlyIntegrationAsActionConnectionIntegration(v *FastlyIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{FastlyIntegration: v}
}

// FreshserviceIntegrationAsActionConnectionIntegration is a convenience function that returns FreshserviceIntegration wrapped in ActionConnectionIntegration.
func FreshserviceIntegrationAsActionConnectionIntegration(v *FreshserviceIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{FreshserviceIntegration: v}
}

// GCPIntegrationAsActionConnectionIntegration is a convenience function that returns GCPIntegration wrapped in ActionConnectionIntegration.
func GCPIntegrationAsActionConnectionIntegration(v *GCPIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{GCPIntegration: v}
}

// GeminiIntegrationAsActionConnectionIntegration is a convenience function that returns GeminiIntegration wrapped in ActionConnectionIntegration.
func GeminiIntegrationAsActionConnectionIntegration(v *GeminiIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{GeminiIntegration: v}
}

// GitlabIntegrationAsActionConnectionIntegration is a convenience function that returns GitlabIntegration wrapped in ActionConnectionIntegration.
func GitlabIntegrationAsActionConnectionIntegration(v *GitlabIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{GitlabIntegration: v}
}

// GreyNoiseIntegrationAsActionConnectionIntegration is a convenience function that returns GreyNoiseIntegration wrapped in ActionConnectionIntegration.
func GreyNoiseIntegrationAsActionConnectionIntegration(v *GreyNoiseIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{GreyNoiseIntegration: v}
}

// HTTPIntegrationAsActionConnectionIntegration is a convenience function that returns HTTPIntegration wrapped in ActionConnectionIntegration.
func HTTPIntegrationAsActionConnectionIntegration(v *HTTPIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{HTTPIntegration: v}
}

// LaunchDarklyIntegrationAsActionConnectionIntegration is a convenience function that returns LaunchDarklyIntegration wrapped in ActionConnectionIntegration.
func LaunchDarklyIntegrationAsActionConnectionIntegration(v *LaunchDarklyIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{LaunchDarklyIntegration: v}
}

// NotionIntegrationAsActionConnectionIntegration is a convenience function that returns NotionIntegration wrapped in ActionConnectionIntegration.
func NotionIntegrationAsActionConnectionIntegration(v *NotionIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{NotionIntegration: v}
}

// OktaIntegrationAsActionConnectionIntegration is a convenience function that returns OktaIntegration wrapped in ActionConnectionIntegration.
func OktaIntegrationAsActionConnectionIntegration(v *OktaIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{OktaIntegration: v}
}

// OpenAIIntegrationAsActionConnectionIntegration is a convenience function that returns OpenAIIntegration wrapped in ActionConnectionIntegration.
func OpenAIIntegrationAsActionConnectionIntegration(v *OpenAIIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{OpenAIIntegration: v}
}

// ServiceNowIntegrationAsActionConnectionIntegration is a convenience function that returns ServiceNowIntegration wrapped in ActionConnectionIntegration.
func ServiceNowIntegrationAsActionConnectionIntegration(v *ServiceNowIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{ServiceNowIntegration: v}
}

// SplitIntegrationAsActionConnectionIntegration is a convenience function that returns SplitIntegration wrapped in ActionConnectionIntegration.
func SplitIntegrationAsActionConnectionIntegration(v *SplitIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{SplitIntegration: v}
}

// StatsigIntegrationAsActionConnectionIntegration is a convenience function that returns StatsigIntegration wrapped in ActionConnectionIntegration.
func StatsigIntegrationAsActionConnectionIntegration(v *StatsigIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{StatsigIntegration: v}
}

// VirusTotalIntegrationAsActionConnectionIntegration is a convenience function that returns VirusTotalIntegration wrapped in ActionConnectionIntegration.
func VirusTotalIntegrationAsActionConnectionIntegration(v *VirusTotalIntegration) ActionConnectionIntegration {
	return ActionConnectionIntegration{VirusTotalIntegration: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ActionConnectionIntegration) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSIntegration
	err = datadog.Unmarshal(data, &obj.AWSIntegration)
	if err == nil {
		if obj.AWSIntegration != nil && obj.AWSIntegration.UnparsedObject == nil {
			jsonAWSIntegration, _ := datadog.Marshal(obj.AWSIntegration)
			if string(jsonAWSIntegration) == "{}" { // empty struct
				obj.AWSIntegration = nil
			} else {
				match++
			}
		} else {
			obj.AWSIntegration = nil
		}
	} else {
		obj.AWSIntegration = nil
	}

	// try to unmarshal data into AnthropicIntegration
	err = datadog.Unmarshal(data, &obj.AnthropicIntegration)
	if err == nil {
		if obj.AnthropicIntegration != nil && obj.AnthropicIntegration.UnparsedObject == nil {
			jsonAnthropicIntegration, _ := datadog.Marshal(obj.AnthropicIntegration)
			if string(jsonAnthropicIntegration) == "{}" { // empty struct
				obj.AnthropicIntegration = nil
			} else {
				match++
			}
		} else {
			obj.AnthropicIntegration = nil
		}
	} else {
		obj.AnthropicIntegration = nil
	}

	// try to unmarshal data into AsanaIntegration
	err = datadog.Unmarshal(data, &obj.AsanaIntegration)
	if err == nil {
		if obj.AsanaIntegration != nil && obj.AsanaIntegration.UnparsedObject == nil {
			jsonAsanaIntegration, _ := datadog.Marshal(obj.AsanaIntegration)
			if string(jsonAsanaIntegration) == "{}" { // empty struct
				obj.AsanaIntegration = nil
			} else {
				match++
			}
		} else {
			obj.AsanaIntegration = nil
		}
	} else {
		obj.AsanaIntegration = nil
	}

	// try to unmarshal data into AzureIntegration
	err = datadog.Unmarshal(data, &obj.AzureIntegration)
	if err == nil {
		if obj.AzureIntegration != nil && obj.AzureIntegration.UnparsedObject == nil {
			jsonAzureIntegration, _ := datadog.Marshal(obj.AzureIntegration)
			if string(jsonAzureIntegration) == "{}" { // empty struct
				obj.AzureIntegration = nil
			} else {
				match++
			}
		} else {
			obj.AzureIntegration = nil
		}
	} else {
		obj.AzureIntegration = nil
	}

	// try to unmarshal data into CircleCIIntegration
	err = datadog.Unmarshal(data, &obj.CircleCIIntegration)
	if err == nil {
		if obj.CircleCIIntegration != nil && obj.CircleCIIntegration.UnparsedObject == nil {
			jsonCircleCIIntegration, _ := datadog.Marshal(obj.CircleCIIntegration)
			if string(jsonCircleCIIntegration) == "{}" { // empty struct
				obj.CircleCIIntegration = nil
			} else {
				match++
			}
		} else {
			obj.CircleCIIntegration = nil
		}
	} else {
		obj.CircleCIIntegration = nil
	}

	// try to unmarshal data into ClickupIntegration
	err = datadog.Unmarshal(data, &obj.ClickupIntegration)
	if err == nil {
		if obj.ClickupIntegration != nil && obj.ClickupIntegration.UnparsedObject == nil {
			jsonClickupIntegration, _ := datadog.Marshal(obj.ClickupIntegration)
			if string(jsonClickupIntegration) == "{}" { // empty struct
				obj.ClickupIntegration = nil
			} else {
				match++
			}
		} else {
			obj.ClickupIntegration = nil
		}
	} else {
		obj.ClickupIntegration = nil
	}

	// try to unmarshal data into CloudflareIntegration
	err = datadog.Unmarshal(data, &obj.CloudflareIntegration)
	if err == nil {
		if obj.CloudflareIntegration != nil && obj.CloudflareIntegration.UnparsedObject == nil {
			jsonCloudflareIntegration, _ := datadog.Marshal(obj.CloudflareIntegration)
			if string(jsonCloudflareIntegration) == "{}" { // empty struct
				obj.CloudflareIntegration = nil
			} else {
				match++
			}
		} else {
			obj.CloudflareIntegration = nil
		}
	} else {
		obj.CloudflareIntegration = nil
	}

	// try to unmarshal data into ConfigCatIntegration
	err = datadog.Unmarshal(data, &obj.ConfigCatIntegration)
	if err == nil {
		if obj.ConfigCatIntegration != nil && obj.ConfigCatIntegration.UnparsedObject == nil {
			jsonConfigCatIntegration, _ := datadog.Marshal(obj.ConfigCatIntegration)
			if string(jsonConfigCatIntegration) == "{}" { // empty struct
				obj.ConfigCatIntegration = nil
			} else {
				match++
			}
		} else {
			obj.ConfigCatIntegration = nil
		}
	} else {
		obj.ConfigCatIntegration = nil
	}

	// try to unmarshal data into DatadogIntegration
	err = datadog.Unmarshal(data, &obj.DatadogIntegration)
	if err == nil {
		if obj.DatadogIntegration != nil && obj.DatadogIntegration.UnparsedObject == nil {
			jsonDatadogIntegration, _ := datadog.Marshal(obj.DatadogIntegration)
			if string(jsonDatadogIntegration) == "{}" { // empty struct
				obj.DatadogIntegration = nil
			} else {
				match++
			}
		} else {
			obj.DatadogIntegration = nil
		}
	} else {
		obj.DatadogIntegration = nil
	}

	// try to unmarshal data into FastlyIntegration
	err = datadog.Unmarshal(data, &obj.FastlyIntegration)
	if err == nil {
		if obj.FastlyIntegration != nil && obj.FastlyIntegration.UnparsedObject == nil {
			jsonFastlyIntegration, _ := datadog.Marshal(obj.FastlyIntegration)
			if string(jsonFastlyIntegration) == "{}" { // empty struct
				obj.FastlyIntegration = nil
			} else {
				match++
			}
		} else {
			obj.FastlyIntegration = nil
		}
	} else {
		obj.FastlyIntegration = nil
	}

	// try to unmarshal data into FreshserviceIntegration
	err = datadog.Unmarshal(data, &obj.FreshserviceIntegration)
	if err == nil {
		if obj.FreshserviceIntegration != nil && obj.FreshserviceIntegration.UnparsedObject == nil {
			jsonFreshserviceIntegration, _ := datadog.Marshal(obj.FreshserviceIntegration)
			if string(jsonFreshserviceIntegration) == "{}" { // empty struct
				obj.FreshserviceIntegration = nil
			} else {
				match++
			}
		} else {
			obj.FreshserviceIntegration = nil
		}
	} else {
		obj.FreshserviceIntegration = nil
	}

	// try to unmarshal data into GCPIntegration
	err = datadog.Unmarshal(data, &obj.GCPIntegration)
	if err == nil {
		if obj.GCPIntegration != nil && obj.GCPIntegration.UnparsedObject == nil {
			jsonGCPIntegration, _ := datadog.Marshal(obj.GCPIntegration)
			if string(jsonGCPIntegration) == "{}" { // empty struct
				obj.GCPIntegration = nil
			} else {
				match++
			}
		} else {
			obj.GCPIntegration = nil
		}
	} else {
		obj.GCPIntegration = nil
	}

	// try to unmarshal data into GeminiIntegration
	err = datadog.Unmarshal(data, &obj.GeminiIntegration)
	if err == nil {
		if obj.GeminiIntegration != nil && obj.GeminiIntegration.UnparsedObject == nil {
			jsonGeminiIntegration, _ := datadog.Marshal(obj.GeminiIntegration)
			if string(jsonGeminiIntegration) == "{}" { // empty struct
				obj.GeminiIntegration = nil
			} else {
				match++
			}
		} else {
			obj.GeminiIntegration = nil
		}
	} else {
		obj.GeminiIntegration = nil
	}

	// try to unmarshal data into GitlabIntegration
	err = datadog.Unmarshal(data, &obj.GitlabIntegration)
	if err == nil {
		if obj.GitlabIntegration != nil && obj.GitlabIntegration.UnparsedObject == nil {
			jsonGitlabIntegration, _ := datadog.Marshal(obj.GitlabIntegration)
			if string(jsonGitlabIntegration) == "{}" { // empty struct
				obj.GitlabIntegration = nil
			} else {
				match++
			}
		} else {
			obj.GitlabIntegration = nil
		}
	} else {
		obj.GitlabIntegration = nil
	}

	// try to unmarshal data into GreyNoiseIntegration
	err = datadog.Unmarshal(data, &obj.GreyNoiseIntegration)
	if err == nil {
		if obj.GreyNoiseIntegration != nil && obj.GreyNoiseIntegration.UnparsedObject == nil {
			jsonGreyNoiseIntegration, _ := datadog.Marshal(obj.GreyNoiseIntegration)
			if string(jsonGreyNoiseIntegration) == "{}" { // empty struct
				obj.GreyNoiseIntegration = nil
			} else {
				match++
			}
		} else {
			obj.GreyNoiseIntegration = nil
		}
	} else {
		obj.GreyNoiseIntegration = nil
	}

	// try to unmarshal data into HTTPIntegration
	err = datadog.Unmarshal(data, &obj.HTTPIntegration)
	if err == nil {
		if obj.HTTPIntegration != nil && obj.HTTPIntegration.UnparsedObject == nil {
			jsonHTTPIntegration, _ := datadog.Marshal(obj.HTTPIntegration)
			if string(jsonHTTPIntegration) == "{}" { // empty struct
				obj.HTTPIntegration = nil
			} else {
				match++
			}
		} else {
			obj.HTTPIntegration = nil
		}
	} else {
		obj.HTTPIntegration = nil
	}

	// try to unmarshal data into LaunchDarklyIntegration
	err = datadog.Unmarshal(data, &obj.LaunchDarklyIntegration)
	if err == nil {
		if obj.LaunchDarklyIntegration != nil && obj.LaunchDarklyIntegration.UnparsedObject == nil {
			jsonLaunchDarklyIntegration, _ := datadog.Marshal(obj.LaunchDarklyIntegration)
			if string(jsonLaunchDarklyIntegration) == "{}" { // empty struct
				obj.LaunchDarklyIntegration = nil
			} else {
				match++
			}
		} else {
			obj.LaunchDarklyIntegration = nil
		}
	} else {
		obj.LaunchDarklyIntegration = nil
	}

	// try to unmarshal data into NotionIntegration
	err = datadog.Unmarshal(data, &obj.NotionIntegration)
	if err == nil {
		if obj.NotionIntegration != nil && obj.NotionIntegration.UnparsedObject == nil {
			jsonNotionIntegration, _ := datadog.Marshal(obj.NotionIntegration)
			if string(jsonNotionIntegration) == "{}" { // empty struct
				obj.NotionIntegration = nil
			} else {
				match++
			}
		} else {
			obj.NotionIntegration = nil
		}
	} else {
		obj.NotionIntegration = nil
	}

	// try to unmarshal data into OktaIntegration
	err = datadog.Unmarshal(data, &obj.OktaIntegration)
	if err == nil {
		if obj.OktaIntegration != nil && obj.OktaIntegration.UnparsedObject == nil {
			jsonOktaIntegration, _ := datadog.Marshal(obj.OktaIntegration)
			if string(jsonOktaIntegration) == "{}" { // empty struct
				obj.OktaIntegration = nil
			} else {
				match++
			}
		} else {
			obj.OktaIntegration = nil
		}
	} else {
		obj.OktaIntegration = nil
	}

	// try to unmarshal data into OpenAIIntegration
	err = datadog.Unmarshal(data, &obj.OpenAIIntegration)
	if err == nil {
		if obj.OpenAIIntegration != nil && obj.OpenAIIntegration.UnparsedObject == nil {
			jsonOpenAIIntegration, _ := datadog.Marshal(obj.OpenAIIntegration)
			if string(jsonOpenAIIntegration) == "{}" { // empty struct
				obj.OpenAIIntegration = nil
			} else {
				match++
			}
		} else {
			obj.OpenAIIntegration = nil
		}
	} else {
		obj.OpenAIIntegration = nil
	}

	// try to unmarshal data into ServiceNowIntegration
	err = datadog.Unmarshal(data, &obj.ServiceNowIntegration)
	if err == nil {
		if obj.ServiceNowIntegration != nil && obj.ServiceNowIntegration.UnparsedObject == nil {
			jsonServiceNowIntegration, _ := datadog.Marshal(obj.ServiceNowIntegration)
			if string(jsonServiceNowIntegration) == "{}" { // empty struct
				obj.ServiceNowIntegration = nil
			} else {
				match++
			}
		} else {
			obj.ServiceNowIntegration = nil
		}
	} else {
		obj.ServiceNowIntegration = nil
	}

	// try to unmarshal data into SplitIntegration
	err = datadog.Unmarshal(data, &obj.SplitIntegration)
	if err == nil {
		if obj.SplitIntegration != nil && obj.SplitIntegration.UnparsedObject == nil {
			jsonSplitIntegration, _ := datadog.Marshal(obj.SplitIntegration)
			if string(jsonSplitIntegration) == "{}" { // empty struct
				obj.SplitIntegration = nil
			} else {
				match++
			}
		} else {
			obj.SplitIntegration = nil
		}
	} else {
		obj.SplitIntegration = nil
	}

	// try to unmarshal data into StatsigIntegration
	err = datadog.Unmarshal(data, &obj.StatsigIntegration)
	if err == nil {
		if obj.StatsigIntegration != nil && obj.StatsigIntegration.UnparsedObject == nil {
			jsonStatsigIntegration, _ := datadog.Marshal(obj.StatsigIntegration)
			if string(jsonStatsigIntegration) == "{}" { // empty struct
				obj.StatsigIntegration = nil
			} else {
				match++
			}
		} else {
			obj.StatsigIntegration = nil
		}
	} else {
		obj.StatsigIntegration = nil
	}

	// try to unmarshal data into VirusTotalIntegration
	err = datadog.Unmarshal(data, &obj.VirusTotalIntegration)
	if err == nil {
		if obj.VirusTotalIntegration != nil && obj.VirusTotalIntegration.UnparsedObject == nil {
			jsonVirusTotalIntegration, _ := datadog.Marshal(obj.VirusTotalIntegration)
			if string(jsonVirusTotalIntegration) == "{}" { // empty struct
				obj.VirusTotalIntegration = nil
			} else {
				match++
			}
		} else {
			obj.VirusTotalIntegration = nil
		}
	} else {
		obj.VirusTotalIntegration = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AWSIntegration = nil
		obj.AnthropicIntegration = nil
		obj.AsanaIntegration = nil
		obj.AzureIntegration = nil
		obj.CircleCIIntegration = nil
		obj.ClickupIntegration = nil
		obj.CloudflareIntegration = nil
		obj.ConfigCatIntegration = nil
		obj.DatadogIntegration = nil
		obj.FastlyIntegration = nil
		obj.FreshserviceIntegration = nil
		obj.GCPIntegration = nil
		obj.GeminiIntegration = nil
		obj.GitlabIntegration = nil
		obj.GreyNoiseIntegration = nil
		obj.HTTPIntegration = nil
		obj.LaunchDarklyIntegration = nil
		obj.NotionIntegration = nil
		obj.OktaIntegration = nil
		obj.OpenAIIntegration = nil
		obj.ServiceNowIntegration = nil
		obj.SplitIntegration = nil
		obj.StatsigIntegration = nil
		obj.VirusTotalIntegration = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ActionConnectionIntegration) MarshalJSON() ([]byte, error) {
	if obj.AWSIntegration != nil {
		return datadog.Marshal(&obj.AWSIntegration)
	}

	if obj.AnthropicIntegration != nil {
		return datadog.Marshal(&obj.AnthropicIntegration)
	}

	if obj.AsanaIntegration != nil {
		return datadog.Marshal(&obj.AsanaIntegration)
	}

	if obj.AzureIntegration != nil {
		return datadog.Marshal(&obj.AzureIntegration)
	}

	if obj.CircleCIIntegration != nil {
		return datadog.Marshal(&obj.CircleCIIntegration)
	}

	if obj.ClickupIntegration != nil {
		return datadog.Marshal(&obj.ClickupIntegration)
	}

	if obj.CloudflareIntegration != nil {
		return datadog.Marshal(&obj.CloudflareIntegration)
	}

	if obj.ConfigCatIntegration != nil {
		return datadog.Marshal(&obj.ConfigCatIntegration)
	}

	if obj.DatadogIntegration != nil {
		return datadog.Marshal(&obj.DatadogIntegration)
	}

	if obj.FastlyIntegration != nil {
		return datadog.Marshal(&obj.FastlyIntegration)
	}

	if obj.FreshserviceIntegration != nil {
		return datadog.Marshal(&obj.FreshserviceIntegration)
	}

	if obj.GCPIntegration != nil {
		return datadog.Marshal(&obj.GCPIntegration)
	}

	if obj.GeminiIntegration != nil {
		return datadog.Marshal(&obj.GeminiIntegration)
	}

	if obj.GitlabIntegration != nil {
		return datadog.Marshal(&obj.GitlabIntegration)
	}

	if obj.GreyNoiseIntegration != nil {
		return datadog.Marshal(&obj.GreyNoiseIntegration)
	}

	if obj.HTTPIntegration != nil {
		return datadog.Marshal(&obj.HTTPIntegration)
	}

	if obj.LaunchDarklyIntegration != nil {
		return datadog.Marshal(&obj.LaunchDarklyIntegration)
	}

	if obj.NotionIntegration != nil {
		return datadog.Marshal(&obj.NotionIntegration)
	}

	if obj.OktaIntegration != nil {
		return datadog.Marshal(&obj.OktaIntegration)
	}

	if obj.OpenAIIntegration != nil {
		return datadog.Marshal(&obj.OpenAIIntegration)
	}

	if obj.ServiceNowIntegration != nil {
		return datadog.Marshal(&obj.ServiceNowIntegration)
	}

	if obj.SplitIntegration != nil {
		return datadog.Marshal(&obj.SplitIntegration)
	}

	if obj.StatsigIntegration != nil {
		return datadog.Marshal(&obj.StatsigIntegration)
	}

	if obj.VirusTotalIntegration != nil {
		return datadog.Marshal(&obj.VirusTotalIntegration)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ActionConnectionIntegration) GetActualInstance() interface{} {
	if obj.AWSIntegration != nil {
		return obj.AWSIntegration
	}

	if obj.AnthropicIntegration != nil {
		return obj.AnthropicIntegration
	}

	if obj.AsanaIntegration != nil {
		return obj.AsanaIntegration
	}

	if obj.AzureIntegration != nil {
		return obj.AzureIntegration
	}

	if obj.CircleCIIntegration != nil {
		return obj.CircleCIIntegration
	}

	if obj.ClickupIntegration != nil {
		return obj.ClickupIntegration
	}

	if obj.CloudflareIntegration != nil {
		return obj.CloudflareIntegration
	}

	if obj.ConfigCatIntegration != nil {
		return obj.ConfigCatIntegration
	}

	if obj.DatadogIntegration != nil {
		return obj.DatadogIntegration
	}

	if obj.FastlyIntegration != nil {
		return obj.FastlyIntegration
	}

	if obj.FreshserviceIntegration != nil {
		return obj.FreshserviceIntegration
	}

	if obj.GCPIntegration != nil {
		return obj.GCPIntegration
	}

	if obj.GeminiIntegration != nil {
		return obj.GeminiIntegration
	}

	if obj.GitlabIntegration != nil {
		return obj.GitlabIntegration
	}

	if obj.GreyNoiseIntegration != nil {
		return obj.GreyNoiseIntegration
	}

	if obj.HTTPIntegration != nil {
		return obj.HTTPIntegration
	}

	if obj.LaunchDarklyIntegration != nil {
		return obj.LaunchDarklyIntegration
	}

	if obj.NotionIntegration != nil {
		return obj.NotionIntegration
	}

	if obj.OktaIntegration != nil {
		return obj.OktaIntegration
	}

	if obj.OpenAIIntegration != nil {
		return obj.OpenAIIntegration
	}

	if obj.ServiceNowIntegration != nil {
		return obj.ServiceNowIntegration
	}

	if obj.SplitIntegration != nil {
		return obj.SplitIntegration
	}

	if obj.StatsigIntegration != nil {
		return obj.StatsigIntegration
	}

	if obj.VirusTotalIntegration != nil {
		return obj.VirusTotalIntegration
	}

	// all schemas are nil
	return nil
}
