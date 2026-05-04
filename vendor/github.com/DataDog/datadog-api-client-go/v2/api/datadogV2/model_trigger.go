// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Trigger - One of the triggers that can start the execution of a workflow.
type Trigger struct {
	APITriggerWrapper                *APITriggerWrapper
	AppTriggerWrapper                *AppTriggerWrapper
	CaseTriggerWrapper               *CaseTriggerWrapper
	ChangeEventTriggerWrapper        *ChangeEventTriggerWrapper
	DatabaseMonitoringTriggerWrapper *DatabaseMonitoringTriggerWrapper
	DatastoreTriggerWrapper          *DatastoreTriggerWrapper
	DashboardTriggerWrapper          *DashboardTriggerWrapper
	FormTriggerWrapper               *FormTriggerWrapper
	GithubWebhookTriggerWrapper      *GithubWebhookTriggerWrapper
	IncidentTriggerWrapper           *IncidentTriggerWrapper
	MonitorTriggerWrapper            *MonitorTriggerWrapper
	NotebookTriggerWrapper           *NotebookTriggerWrapper
	OnCallTriggerWrapper             *OnCallTriggerWrapper
	ScheduleTriggerWrapper           *ScheduleTriggerWrapper
	SecurityTriggerWrapper           *SecurityTriggerWrapper
	SelfServiceTriggerWrapper        *SelfServiceTriggerWrapper
	SlackTriggerWrapper              *SlackTriggerWrapper
	SoftwareCatalogTriggerWrapper    *SoftwareCatalogTriggerWrapper
	WorkflowTriggerWrapper           *WorkflowTriggerWrapper

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// APITriggerWrapperAsTrigger is a convenience function that returns APITriggerWrapper wrapped in Trigger.
func APITriggerWrapperAsTrigger(v *APITriggerWrapper) Trigger {
	return Trigger{APITriggerWrapper: v}
}

// AppTriggerWrapperAsTrigger is a convenience function that returns AppTriggerWrapper wrapped in Trigger.
func AppTriggerWrapperAsTrigger(v *AppTriggerWrapper) Trigger {
	return Trigger{AppTriggerWrapper: v}
}

// CaseTriggerWrapperAsTrigger is a convenience function that returns CaseTriggerWrapper wrapped in Trigger.
func CaseTriggerWrapperAsTrigger(v *CaseTriggerWrapper) Trigger {
	return Trigger{CaseTriggerWrapper: v}
}

// ChangeEventTriggerWrapperAsTrigger is a convenience function that returns ChangeEventTriggerWrapper wrapped in Trigger.
func ChangeEventTriggerWrapperAsTrigger(v *ChangeEventTriggerWrapper) Trigger {
	return Trigger{ChangeEventTriggerWrapper: v}
}

// DatabaseMonitoringTriggerWrapperAsTrigger is a convenience function that returns DatabaseMonitoringTriggerWrapper wrapped in Trigger.
func DatabaseMonitoringTriggerWrapperAsTrigger(v *DatabaseMonitoringTriggerWrapper) Trigger {
	return Trigger{DatabaseMonitoringTriggerWrapper: v}
}

// DatastoreTriggerWrapperAsTrigger is a convenience function that returns DatastoreTriggerWrapper wrapped in Trigger.
func DatastoreTriggerWrapperAsTrigger(v *DatastoreTriggerWrapper) Trigger {
	return Trigger{DatastoreTriggerWrapper: v}
}

// DashboardTriggerWrapperAsTrigger is a convenience function that returns DashboardTriggerWrapper wrapped in Trigger.
func DashboardTriggerWrapperAsTrigger(v *DashboardTriggerWrapper) Trigger {
	return Trigger{DashboardTriggerWrapper: v}
}

// FormTriggerWrapperAsTrigger is a convenience function that returns FormTriggerWrapper wrapped in Trigger.
func FormTriggerWrapperAsTrigger(v *FormTriggerWrapper) Trigger {
	return Trigger{FormTriggerWrapper: v}
}

// GithubWebhookTriggerWrapperAsTrigger is a convenience function that returns GithubWebhookTriggerWrapper wrapped in Trigger.
func GithubWebhookTriggerWrapperAsTrigger(v *GithubWebhookTriggerWrapper) Trigger {
	return Trigger{GithubWebhookTriggerWrapper: v}
}

// IncidentTriggerWrapperAsTrigger is a convenience function that returns IncidentTriggerWrapper wrapped in Trigger.
func IncidentTriggerWrapperAsTrigger(v *IncidentTriggerWrapper) Trigger {
	return Trigger{IncidentTriggerWrapper: v}
}

// MonitorTriggerWrapperAsTrigger is a convenience function that returns MonitorTriggerWrapper wrapped in Trigger.
func MonitorTriggerWrapperAsTrigger(v *MonitorTriggerWrapper) Trigger {
	return Trigger{MonitorTriggerWrapper: v}
}

// NotebookTriggerWrapperAsTrigger is a convenience function that returns NotebookTriggerWrapper wrapped in Trigger.
func NotebookTriggerWrapperAsTrigger(v *NotebookTriggerWrapper) Trigger {
	return Trigger{NotebookTriggerWrapper: v}
}

// OnCallTriggerWrapperAsTrigger is a convenience function that returns OnCallTriggerWrapper wrapped in Trigger.
func OnCallTriggerWrapperAsTrigger(v *OnCallTriggerWrapper) Trigger {
	return Trigger{OnCallTriggerWrapper: v}
}

// ScheduleTriggerWrapperAsTrigger is a convenience function that returns ScheduleTriggerWrapper wrapped in Trigger.
func ScheduleTriggerWrapperAsTrigger(v *ScheduleTriggerWrapper) Trigger {
	return Trigger{ScheduleTriggerWrapper: v}
}

// SecurityTriggerWrapperAsTrigger is a convenience function that returns SecurityTriggerWrapper wrapped in Trigger.
func SecurityTriggerWrapperAsTrigger(v *SecurityTriggerWrapper) Trigger {
	return Trigger{SecurityTriggerWrapper: v}
}

// SelfServiceTriggerWrapperAsTrigger is a convenience function that returns SelfServiceTriggerWrapper wrapped in Trigger.
func SelfServiceTriggerWrapperAsTrigger(v *SelfServiceTriggerWrapper) Trigger {
	return Trigger{SelfServiceTriggerWrapper: v}
}

// SlackTriggerWrapperAsTrigger is a convenience function that returns SlackTriggerWrapper wrapped in Trigger.
func SlackTriggerWrapperAsTrigger(v *SlackTriggerWrapper) Trigger {
	return Trigger{SlackTriggerWrapper: v}
}

// SoftwareCatalogTriggerWrapperAsTrigger is a convenience function that returns SoftwareCatalogTriggerWrapper wrapped in Trigger.
func SoftwareCatalogTriggerWrapperAsTrigger(v *SoftwareCatalogTriggerWrapper) Trigger {
	return Trigger{SoftwareCatalogTriggerWrapper: v}
}

// WorkflowTriggerWrapperAsTrigger is a convenience function that returns WorkflowTriggerWrapper wrapped in Trigger.
func WorkflowTriggerWrapperAsTrigger(v *WorkflowTriggerWrapper) Trigger {
	return Trigger{WorkflowTriggerWrapper: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *Trigger) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into APITriggerWrapper
	err = datadog.Unmarshal(data, &obj.APITriggerWrapper)
	if err == nil {
		if obj.APITriggerWrapper != nil && obj.APITriggerWrapper.UnparsedObject == nil {
			jsonAPITriggerWrapper, _ := datadog.Marshal(obj.APITriggerWrapper)
			if string(jsonAPITriggerWrapper) == "{}" { // empty struct
				obj.APITriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.APITriggerWrapper = nil
		}
	} else {
		obj.APITriggerWrapper = nil
	}

	// try to unmarshal data into AppTriggerWrapper
	err = datadog.Unmarshal(data, &obj.AppTriggerWrapper)
	if err == nil {
		if obj.AppTriggerWrapper != nil && obj.AppTriggerWrapper.UnparsedObject == nil {
			jsonAppTriggerWrapper, _ := datadog.Marshal(obj.AppTriggerWrapper)
			if string(jsonAppTriggerWrapper) == "{}" { // empty struct
				obj.AppTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.AppTriggerWrapper = nil
		}
	} else {
		obj.AppTriggerWrapper = nil
	}

	// try to unmarshal data into CaseTriggerWrapper
	err = datadog.Unmarshal(data, &obj.CaseTriggerWrapper)
	if err == nil {
		if obj.CaseTriggerWrapper != nil && obj.CaseTriggerWrapper.UnparsedObject == nil {
			jsonCaseTriggerWrapper, _ := datadog.Marshal(obj.CaseTriggerWrapper)
			if string(jsonCaseTriggerWrapper) == "{}" { // empty struct
				obj.CaseTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.CaseTriggerWrapper = nil
		}
	} else {
		obj.CaseTriggerWrapper = nil
	}

	// try to unmarshal data into ChangeEventTriggerWrapper
	err = datadog.Unmarshal(data, &obj.ChangeEventTriggerWrapper)
	if err == nil {
		if obj.ChangeEventTriggerWrapper != nil && obj.ChangeEventTriggerWrapper.UnparsedObject == nil {
			jsonChangeEventTriggerWrapper, _ := datadog.Marshal(obj.ChangeEventTriggerWrapper)
			if string(jsonChangeEventTriggerWrapper) == "{}" { // empty struct
				obj.ChangeEventTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.ChangeEventTriggerWrapper = nil
		}
	} else {
		obj.ChangeEventTriggerWrapper = nil
	}

	// try to unmarshal data into DatabaseMonitoringTriggerWrapper
	err = datadog.Unmarshal(data, &obj.DatabaseMonitoringTriggerWrapper)
	if err == nil {
		if obj.DatabaseMonitoringTriggerWrapper != nil && obj.DatabaseMonitoringTriggerWrapper.UnparsedObject == nil {
			jsonDatabaseMonitoringTriggerWrapper, _ := datadog.Marshal(obj.DatabaseMonitoringTriggerWrapper)
			if string(jsonDatabaseMonitoringTriggerWrapper) == "{}" { // empty struct
				obj.DatabaseMonitoringTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.DatabaseMonitoringTriggerWrapper = nil
		}
	} else {
		obj.DatabaseMonitoringTriggerWrapper = nil
	}

	// try to unmarshal data into DatastoreTriggerWrapper
	err = datadog.Unmarshal(data, &obj.DatastoreTriggerWrapper)
	if err == nil {
		if obj.DatastoreTriggerWrapper != nil && obj.DatastoreTriggerWrapper.UnparsedObject == nil {
			jsonDatastoreTriggerWrapper, _ := datadog.Marshal(obj.DatastoreTriggerWrapper)
			if string(jsonDatastoreTriggerWrapper) == "{}" { // empty struct
				obj.DatastoreTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.DatastoreTriggerWrapper = nil
		}
	} else {
		obj.DatastoreTriggerWrapper = nil
	}

	// try to unmarshal data into DashboardTriggerWrapper
	err = datadog.Unmarshal(data, &obj.DashboardTriggerWrapper)
	if err == nil {
		if obj.DashboardTriggerWrapper != nil && obj.DashboardTriggerWrapper.UnparsedObject == nil {
			jsonDashboardTriggerWrapper, _ := datadog.Marshal(obj.DashboardTriggerWrapper)
			if string(jsonDashboardTriggerWrapper) == "{}" { // empty struct
				obj.DashboardTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.DashboardTriggerWrapper = nil
		}
	} else {
		obj.DashboardTriggerWrapper = nil
	}

	// try to unmarshal data into FormTriggerWrapper
	err = datadog.Unmarshal(data, &obj.FormTriggerWrapper)
	if err == nil {
		if obj.FormTriggerWrapper != nil && obj.FormTriggerWrapper.UnparsedObject == nil {
			jsonFormTriggerWrapper, _ := datadog.Marshal(obj.FormTriggerWrapper)
			if string(jsonFormTriggerWrapper) == "{}" { // empty struct
				obj.FormTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.FormTriggerWrapper = nil
		}
	} else {
		obj.FormTriggerWrapper = nil
	}

	// try to unmarshal data into GithubWebhookTriggerWrapper
	err = datadog.Unmarshal(data, &obj.GithubWebhookTriggerWrapper)
	if err == nil {
		if obj.GithubWebhookTriggerWrapper != nil && obj.GithubWebhookTriggerWrapper.UnparsedObject == nil {
			jsonGithubWebhookTriggerWrapper, _ := datadog.Marshal(obj.GithubWebhookTriggerWrapper)
			if string(jsonGithubWebhookTriggerWrapper) == "{}" { // empty struct
				obj.GithubWebhookTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.GithubWebhookTriggerWrapper = nil
		}
	} else {
		obj.GithubWebhookTriggerWrapper = nil
	}

	// try to unmarshal data into IncidentTriggerWrapper
	err = datadog.Unmarshal(data, &obj.IncidentTriggerWrapper)
	if err == nil {
		if obj.IncidentTriggerWrapper != nil && obj.IncidentTriggerWrapper.UnparsedObject == nil {
			jsonIncidentTriggerWrapper, _ := datadog.Marshal(obj.IncidentTriggerWrapper)
			if string(jsonIncidentTriggerWrapper) == "{}" { // empty struct
				obj.IncidentTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.IncidentTriggerWrapper = nil
		}
	} else {
		obj.IncidentTriggerWrapper = nil
	}

	// try to unmarshal data into MonitorTriggerWrapper
	err = datadog.Unmarshal(data, &obj.MonitorTriggerWrapper)
	if err == nil {
		if obj.MonitorTriggerWrapper != nil && obj.MonitorTriggerWrapper.UnparsedObject == nil {
			jsonMonitorTriggerWrapper, _ := datadog.Marshal(obj.MonitorTriggerWrapper)
			if string(jsonMonitorTriggerWrapper) == "{}" { // empty struct
				obj.MonitorTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.MonitorTriggerWrapper = nil
		}
	} else {
		obj.MonitorTriggerWrapper = nil
	}

	// try to unmarshal data into NotebookTriggerWrapper
	err = datadog.Unmarshal(data, &obj.NotebookTriggerWrapper)
	if err == nil {
		if obj.NotebookTriggerWrapper != nil && obj.NotebookTriggerWrapper.UnparsedObject == nil {
			jsonNotebookTriggerWrapper, _ := datadog.Marshal(obj.NotebookTriggerWrapper)
			if string(jsonNotebookTriggerWrapper) == "{}" { // empty struct
				obj.NotebookTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.NotebookTriggerWrapper = nil
		}
	} else {
		obj.NotebookTriggerWrapper = nil
	}

	// try to unmarshal data into OnCallTriggerWrapper
	err = datadog.Unmarshal(data, &obj.OnCallTriggerWrapper)
	if err == nil {
		if obj.OnCallTriggerWrapper != nil && obj.OnCallTriggerWrapper.UnparsedObject == nil {
			jsonOnCallTriggerWrapper, _ := datadog.Marshal(obj.OnCallTriggerWrapper)
			if string(jsonOnCallTriggerWrapper) == "{}" { // empty struct
				obj.OnCallTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.OnCallTriggerWrapper = nil
		}
	} else {
		obj.OnCallTriggerWrapper = nil
	}

	// try to unmarshal data into ScheduleTriggerWrapper
	err = datadog.Unmarshal(data, &obj.ScheduleTriggerWrapper)
	if err == nil {
		if obj.ScheduleTriggerWrapper != nil && obj.ScheduleTriggerWrapper.UnparsedObject == nil {
			jsonScheduleTriggerWrapper, _ := datadog.Marshal(obj.ScheduleTriggerWrapper)
			if string(jsonScheduleTriggerWrapper) == "{}" { // empty struct
				obj.ScheduleTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.ScheduleTriggerWrapper = nil
		}
	} else {
		obj.ScheduleTriggerWrapper = nil
	}

	// try to unmarshal data into SecurityTriggerWrapper
	err = datadog.Unmarshal(data, &obj.SecurityTriggerWrapper)
	if err == nil {
		if obj.SecurityTriggerWrapper != nil && obj.SecurityTriggerWrapper.UnparsedObject == nil {
			jsonSecurityTriggerWrapper, _ := datadog.Marshal(obj.SecurityTriggerWrapper)
			if string(jsonSecurityTriggerWrapper) == "{}" { // empty struct
				obj.SecurityTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.SecurityTriggerWrapper = nil
		}
	} else {
		obj.SecurityTriggerWrapper = nil
	}

	// try to unmarshal data into SelfServiceTriggerWrapper
	err = datadog.Unmarshal(data, &obj.SelfServiceTriggerWrapper)
	if err == nil {
		if obj.SelfServiceTriggerWrapper != nil && obj.SelfServiceTriggerWrapper.UnparsedObject == nil {
			jsonSelfServiceTriggerWrapper, _ := datadog.Marshal(obj.SelfServiceTriggerWrapper)
			if string(jsonSelfServiceTriggerWrapper) == "{}" { // empty struct
				obj.SelfServiceTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.SelfServiceTriggerWrapper = nil
		}
	} else {
		obj.SelfServiceTriggerWrapper = nil
	}

	// try to unmarshal data into SlackTriggerWrapper
	err = datadog.Unmarshal(data, &obj.SlackTriggerWrapper)
	if err == nil {
		if obj.SlackTriggerWrapper != nil && obj.SlackTriggerWrapper.UnparsedObject == nil {
			jsonSlackTriggerWrapper, _ := datadog.Marshal(obj.SlackTriggerWrapper)
			if string(jsonSlackTriggerWrapper) == "{}" { // empty struct
				obj.SlackTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.SlackTriggerWrapper = nil
		}
	} else {
		obj.SlackTriggerWrapper = nil
	}

	// try to unmarshal data into SoftwareCatalogTriggerWrapper
	err = datadog.Unmarshal(data, &obj.SoftwareCatalogTriggerWrapper)
	if err == nil {
		if obj.SoftwareCatalogTriggerWrapper != nil && obj.SoftwareCatalogTriggerWrapper.UnparsedObject == nil {
			jsonSoftwareCatalogTriggerWrapper, _ := datadog.Marshal(obj.SoftwareCatalogTriggerWrapper)
			if string(jsonSoftwareCatalogTriggerWrapper) == "{}" { // empty struct
				obj.SoftwareCatalogTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.SoftwareCatalogTriggerWrapper = nil
		}
	} else {
		obj.SoftwareCatalogTriggerWrapper = nil
	}

	// try to unmarshal data into WorkflowTriggerWrapper
	err = datadog.Unmarshal(data, &obj.WorkflowTriggerWrapper)
	if err == nil {
		if obj.WorkflowTriggerWrapper != nil && obj.WorkflowTriggerWrapper.UnparsedObject == nil {
			jsonWorkflowTriggerWrapper, _ := datadog.Marshal(obj.WorkflowTriggerWrapper)
			if string(jsonWorkflowTriggerWrapper) == "{}" { // empty struct
				obj.WorkflowTriggerWrapper = nil
			} else {
				match++
			}
		} else {
			obj.WorkflowTriggerWrapper = nil
		}
	} else {
		obj.WorkflowTriggerWrapper = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.APITriggerWrapper = nil
		obj.AppTriggerWrapper = nil
		obj.CaseTriggerWrapper = nil
		obj.ChangeEventTriggerWrapper = nil
		obj.DatabaseMonitoringTriggerWrapper = nil
		obj.DatastoreTriggerWrapper = nil
		obj.DashboardTriggerWrapper = nil
		obj.FormTriggerWrapper = nil
		obj.GithubWebhookTriggerWrapper = nil
		obj.IncidentTriggerWrapper = nil
		obj.MonitorTriggerWrapper = nil
		obj.NotebookTriggerWrapper = nil
		obj.OnCallTriggerWrapper = nil
		obj.ScheduleTriggerWrapper = nil
		obj.SecurityTriggerWrapper = nil
		obj.SelfServiceTriggerWrapper = nil
		obj.SlackTriggerWrapper = nil
		obj.SoftwareCatalogTriggerWrapper = nil
		obj.WorkflowTriggerWrapper = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj Trigger) MarshalJSON() ([]byte, error) {
	if obj.APITriggerWrapper != nil {
		return datadog.Marshal(&obj.APITriggerWrapper)
	}

	if obj.AppTriggerWrapper != nil {
		return datadog.Marshal(&obj.AppTriggerWrapper)
	}

	if obj.CaseTriggerWrapper != nil {
		return datadog.Marshal(&obj.CaseTriggerWrapper)
	}

	if obj.ChangeEventTriggerWrapper != nil {
		return datadog.Marshal(&obj.ChangeEventTriggerWrapper)
	}

	if obj.DatabaseMonitoringTriggerWrapper != nil {
		return datadog.Marshal(&obj.DatabaseMonitoringTriggerWrapper)
	}

	if obj.DatastoreTriggerWrapper != nil {
		return datadog.Marshal(&obj.DatastoreTriggerWrapper)
	}

	if obj.DashboardTriggerWrapper != nil {
		return datadog.Marshal(&obj.DashboardTriggerWrapper)
	}

	if obj.FormTriggerWrapper != nil {
		return datadog.Marshal(&obj.FormTriggerWrapper)
	}

	if obj.GithubWebhookTriggerWrapper != nil {
		return datadog.Marshal(&obj.GithubWebhookTriggerWrapper)
	}

	if obj.IncidentTriggerWrapper != nil {
		return datadog.Marshal(&obj.IncidentTriggerWrapper)
	}

	if obj.MonitorTriggerWrapper != nil {
		return datadog.Marshal(&obj.MonitorTriggerWrapper)
	}

	if obj.NotebookTriggerWrapper != nil {
		return datadog.Marshal(&obj.NotebookTriggerWrapper)
	}

	if obj.OnCallTriggerWrapper != nil {
		return datadog.Marshal(&obj.OnCallTriggerWrapper)
	}

	if obj.ScheduleTriggerWrapper != nil {
		return datadog.Marshal(&obj.ScheduleTriggerWrapper)
	}

	if obj.SecurityTriggerWrapper != nil {
		return datadog.Marshal(&obj.SecurityTriggerWrapper)
	}

	if obj.SelfServiceTriggerWrapper != nil {
		return datadog.Marshal(&obj.SelfServiceTriggerWrapper)
	}

	if obj.SlackTriggerWrapper != nil {
		return datadog.Marshal(&obj.SlackTriggerWrapper)
	}

	if obj.SoftwareCatalogTriggerWrapper != nil {
		return datadog.Marshal(&obj.SoftwareCatalogTriggerWrapper)
	}

	if obj.WorkflowTriggerWrapper != nil {
		return datadog.Marshal(&obj.WorkflowTriggerWrapper)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *Trigger) GetActualInstance() interface{} {
	if obj.APITriggerWrapper != nil {
		return obj.APITriggerWrapper
	}

	if obj.AppTriggerWrapper != nil {
		return obj.AppTriggerWrapper
	}

	if obj.CaseTriggerWrapper != nil {
		return obj.CaseTriggerWrapper
	}

	if obj.ChangeEventTriggerWrapper != nil {
		return obj.ChangeEventTriggerWrapper
	}

	if obj.DatabaseMonitoringTriggerWrapper != nil {
		return obj.DatabaseMonitoringTriggerWrapper
	}

	if obj.DatastoreTriggerWrapper != nil {
		return obj.DatastoreTriggerWrapper
	}

	if obj.DashboardTriggerWrapper != nil {
		return obj.DashboardTriggerWrapper
	}

	if obj.FormTriggerWrapper != nil {
		return obj.FormTriggerWrapper
	}

	if obj.GithubWebhookTriggerWrapper != nil {
		return obj.GithubWebhookTriggerWrapper
	}

	if obj.IncidentTriggerWrapper != nil {
		return obj.IncidentTriggerWrapper
	}

	if obj.MonitorTriggerWrapper != nil {
		return obj.MonitorTriggerWrapper
	}

	if obj.NotebookTriggerWrapper != nil {
		return obj.NotebookTriggerWrapper
	}

	if obj.OnCallTriggerWrapper != nil {
		return obj.OnCallTriggerWrapper
	}

	if obj.ScheduleTriggerWrapper != nil {
		return obj.ScheduleTriggerWrapper
	}

	if obj.SecurityTriggerWrapper != nil {
		return obj.SecurityTriggerWrapper
	}

	if obj.SelfServiceTriggerWrapper != nil {
		return obj.SelfServiceTriggerWrapper
	}

	if obj.SlackTriggerWrapper != nil {
		return obj.SlackTriggerWrapper
	}

	if obj.SoftwareCatalogTriggerWrapper != nil {
		return obj.SoftwareCatalogTriggerWrapper
	}

	if obj.WorkflowTriggerWrapper != nil {
		return obj.WorkflowTriggerWrapper
	}

	// all schemas are nil
	return nil
}
