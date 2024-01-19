// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SplitGraphSourceWidgetDefinition - The original widget we are splitting on.
type SplitGraphSourceWidgetDefinition struct {
	ChangeWidgetDefinition      *ChangeWidgetDefinition
	GeomapWidgetDefinition      *GeomapWidgetDefinition
	QueryValueWidgetDefinition  *QueryValueWidgetDefinition
	ScatterPlotWidgetDefinition *ScatterPlotWidgetDefinition
	SunburstWidgetDefinition    *SunburstWidgetDefinition
	TableWidgetDefinition       *TableWidgetDefinition
	TimeseriesWidgetDefinition  *TimeseriesWidgetDefinition
	ToplistWidgetDefinition     *ToplistWidgetDefinition
	TreeMapWidgetDefinition     *TreeMapWidgetDefinition

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ChangeWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns ChangeWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func ChangeWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *ChangeWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{ChangeWidgetDefinition: v}
}

// GeomapWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns GeomapWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func GeomapWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *GeomapWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{GeomapWidgetDefinition: v}
}

// QueryValueWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns QueryValueWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func QueryValueWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *QueryValueWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{QueryValueWidgetDefinition: v}
}

// ScatterPlotWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns ScatterPlotWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func ScatterPlotWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *ScatterPlotWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{ScatterPlotWidgetDefinition: v}
}

// SunburstWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns SunburstWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func SunburstWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *SunburstWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{SunburstWidgetDefinition: v}
}

// TableWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns TableWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func TableWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *TableWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{TableWidgetDefinition: v}
}

// TimeseriesWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns TimeseriesWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func TimeseriesWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *TimeseriesWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{TimeseriesWidgetDefinition: v}
}

// ToplistWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns ToplistWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func ToplistWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *ToplistWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{ToplistWidgetDefinition: v}
}

// TreeMapWidgetDefinitionAsSplitGraphSourceWidgetDefinition is a convenience function that returns TreeMapWidgetDefinition wrapped in SplitGraphSourceWidgetDefinition.
func TreeMapWidgetDefinitionAsSplitGraphSourceWidgetDefinition(v *TreeMapWidgetDefinition) SplitGraphSourceWidgetDefinition {
	return SplitGraphSourceWidgetDefinition{TreeMapWidgetDefinition: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *SplitGraphSourceWidgetDefinition) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChangeWidgetDefinition
	err = datadog.Unmarshal(data, &obj.ChangeWidgetDefinition)
	if err == nil {
		if obj.ChangeWidgetDefinition != nil && obj.ChangeWidgetDefinition.UnparsedObject == nil {
			jsonChangeWidgetDefinition, _ := datadog.Marshal(obj.ChangeWidgetDefinition)
			if string(jsonChangeWidgetDefinition) == "{}" { // empty struct
				obj.ChangeWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.ChangeWidgetDefinition = nil
		}
	} else {
		obj.ChangeWidgetDefinition = nil
	}

	// try to unmarshal data into GeomapWidgetDefinition
	err = datadog.Unmarshal(data, &obj.GeomapWidgetDefinition)
	if err == nil {
		if obj.GeomapWidgetDefinition != nil && obj.GeomapWidgetDefinition.UnparsedObject == nil {
			jsonGeomapWidgetDefinition, _ := datadog.Marshal(obj.GeomapWidgetDefinition)
			if string(jsonGeomapWidgetDefinition) == "{}" { // empty struct
				obj.GeomapWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.GeomapWidgetDefinition = nil
		}
	} else {
		obj.GeomapWidgetDefinition = nil
	}

	// try to unmarshal data into QueryValueWidgetDefinition
	err = datadog.Unmarshal(data, &obj.QueryValueWidgetDefinition)
	if err == nil {
		if obj.QueryValueWidgetDefinition != nil && obj.QueryValueWidgetDefinition.UnparsedObject == nil {
			jsonQueryValueWidgetDefinition, _ := datadog.Marshal(obj.QueryValueWidgetDefinition)
			if string(jsonQueryValueWidgetDefinition) == "{}" { // empty struct
				obj.QueryValueWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.QueryValueWidgetDefinition = nil
		}
	} else {
		obj.QueryValueWidgetDefinition = nil
	}

	// try to unmarshal data into ScatterPlotWidgetDefinition
	err = datadog.Unmarshal(data, &obj.ScatterPlotWidgetDefinition)
	if err == nil {
		if obj.ScatterPlotWidgetDefinition != nil && obj.ScatterPlotWidgetDefinition.UnparsedObject == nil {
			jsonScatterPlotWidgetDefinition, _ := datadog.Marshal(obj.ScatterPlotWidgetDefinition)
			if string(jsonScatterPlotWidgetDefinition) == "{}" { // empty struct
				obj.ScatterPlotWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.ScatterPlotWidgetDefinition = nil
		}
	} else {
		obj.ScatterPlotWidgetDefinition = nil
	}

	// try to unmarshal data into SunburstWidgetDefinition
	err = datadog.Unmarshal(data, &obj.SunburstWidgetDefinition)
	if err == nil {
		if obj.SunburstWidgetDefinition != nil && obj.SunburstWidgetDefinition.UnparsedObject == nil {
			jsonSunburstWidgetDefinition, _ := datadog.Marshal(obj.SunburstWidgetDefinition)
			if string(jsonSunburstWidgetDefinition) == "{}" { // empty struct
				obj.SunburstWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.SunburstWidgetDefinition = nil
		}
	} else {
		obj.SunburstWidgetDefinition = nil
	}

	// try to unmarshal data into TableWidgetDefinition
	err = datadog.Unmarshal(data, &obj.TableWidgetDefinition)
	if err == nil {
		if obj.TableWidgetDefinition != nil && obj.TableWidgetDefinition.UnparsedObject == nil {
			jsonTableWidgetDefinition, _ := datadog.Marshal(obj.TableWidgetDefinition)
			if string(jsonTableWidgetDefinition) == "{}" { // empty struct
				obj.TableWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.TableWidgetDefinition = nil
		}
	} else {
		obj.TableWidgetDefinition = nil
	}

	// try to unmarshal data into TimeseriesWidgetDefinition
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetDefinition)
	if err == nil {
		if obj.TimeseriesWidgetDefinition != nil && obj.TimeseriesWidgetDefinition.UnparsedObject == nil {
			jsonTimeseriesWidgetDefinition, _ := datadog.Marshal(obj.TimeseriesWidgetDefinition)
			if string(jsonTimeseriesWidgetDefinition) == "{}" { // empty struct
				obj.TimeseriesWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetDefinition = nil
		}
	} else {
		obj.TimeseriesWidgetDefinition = nil
	}

	// try to unmarshal data into ToplistWidgetDefinition
	err = datadog.Unmarshal(data, &obj.ToplistWidgetDefinition)
	if err == nil {
		if obj.ToplistWidgetDefinition != nil && obj.ToplistWidgetDefinition.UnparsedObject == nil {
			jsonToplistWidgetDefinition, _ := datadog.Marshal(obj.ToplistWidgetDefinition)
			if string(jsonToplistWidgetDefinition) == "{}" { // empty struct
				obj.ToplistWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.ToplistWidgetDefinition = nil
		}
	} else {
		obj.ToplistWidgetDefinition = nil
	}

	// try to unmarshal data into TreeMapWidgetDefinition
	err = datadog.Unmarshal(data, &obj.TreeMapWidgetDefinition)
	if err == nil {
		if obj.TreeMapWidgetDefinition != nil && obj.TreeMapWidgetDefinition.UnparsedObject == nil {
			jsonTreeMapWidgetDefinition, _ := datadog.Marshal(obj.TreeMapWidgetDefinition)
			if string(jsonTreeMapWidgetDefinition) == "{}" { // empty struct
				obj.TreeMapWidgetDefinition = nil
			} else {
				match++
			}
		} else {
			obj.TreeMapWidgetDefinition = nil
		}
	} else {
		obj.TreeMapWidgetDefinition = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ChangeWidgetDefinition = nil
		obj.GeomapWidgetDefinition = nil
		obj.QueryValueWidgetDefinition = nil
		obj.ScatterPlotWidgetDefinition = nil
		obj.SunburstWidgetDefinition = nil
		obj.TableWidgetDefinition = nil
		obj.TimeseriesWidgetDefinition = nil
		obj.ToplistWidgetDefinition = nil
		obj.TreeMapWidgetDefinition = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj SplitGraphSourceWidgetDefinition) MarshalJSON() ([]byte, error) {
	if obj.ChangeWidgetDefinition != nil {
		return datadog.Marshal(&obj.ChangeWidgetDefinition)
	}

	if obj.GeomapWidgetDefinition != nil {
		return datadog.Marshal(&obj.GeomapWidgetDefinition)
	}

	if obj.QueryValueWidgetDefinition != nil {
		return datadog.Marshal(&obj.QueryValueWidgetDefinition)
	}

	if obj.ScatterPlotWidgetDefinition != nil {
		return datadog.Marshal(&obj.ScatterPlotWidgetDefinition)
	}

	if obj.SunburstWidgetDefinition != nil {
		return datadog.Marshal(&obj.SunburstWidgetDefinition)
	}

	if obj.TableWidgetDefinition != nil {
		return datadog.Marshal(&obj.TableWidgetDefinition)
	}

	if obj.TimeseriesWidgetDefinition != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetDefinition)
	}

	if obj.ToplistWidgetDefinition != nil {
		return datadog.Marshal(&obj.ToplistWidgetDefinition)
	}

	if obj.TreeMapWidgetDefinition != nil {
		return datadog.Marshal(&obj.TreeMapWidgetDefinition)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *SplitGraphSourceWidgetDefinition) GetActualInstance() interface{} {
	if obj.ChangeWidgetDefinition != nil {
		return obj.ChangeWidgetDefinition
	}

	if obj.GeomapWidgetDefinition != nil {
		return obj.GeomapWidgetDefinition
	}

	if obj.QueryValueWidgetDefinition != nil {
		return obj.QueryValueWidgetDefinition
	}

	if obj.ScatterPlotWidgetDefinition != nil {
		return obj.ScatterPlotWidgetDefinition
	}

	if obj.SunburstWidgetDefinition != nil {
		return obj.SunburstWidgetDefinition
	}

	if obj.TableWidgetDefinition != nil {
		return obj.TableWidgetDefinition
	}

	if obj.TimeseriesWidgetDefinition != nil {
		return obj.TimeseriesWidgetDefinition
	}

	if obj.ToplistWidgetDefinition != nil {
		return obj.ToplistWidgetDefinition
	}

	if obj.TreeMapWidgetDefinition != nil {
		return obj.TreeMapWidgetDefinition
	}

	// all schemas are nil
	return nil
}
