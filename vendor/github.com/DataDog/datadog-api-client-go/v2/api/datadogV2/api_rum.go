// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_fmt "fmt"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApi service type
type RUMApi datadog.Service

// AggregateRUMEvents Aggregate RUM events.
// The API endpoint to aggregate RUM events into buckets of computed metrics and timeseries.
func (a *RUMApi) AggregateRUMEvents(ctx _context.Context, body RUMAggregateRequest) (RUMAnalyticsAggregateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue RUMAnalyticsAggregateResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.AggregateRUMEvents")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/analytics/aggregate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// CreateRUMApplication Create a new RUM application.
// Create a new RUM application in your organization.
func (a *RUMApi) CreateRUMApplication(ctx _context.Context, body RUMApplicationCreateRequest) (RUMApplicationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue RUMApplicationResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.CreateRUMApplication")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/applications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// DeleteRUMApplication Delete a RUM application.
// Delete an existing RUM application in your organization.
func (a *RUMApi) DeleteRUMApplication(ctx _context.Context, id string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.DeleteRUMApplication")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/applications/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// DeleteSourcemapsOptionalParameters holds optional parameters for DeleteSourcemaps.
type DeleteSourcemapsOptionalParameters struct {
	FilterService       *[]string
	FilterVersion       *[]string
	FilterVariant       *[]string
	FilterId            *[]string
	FilterBuildId       *[]string
	FilterUuid          *[]string
	FilterPlatform      *[]string
	FilterBuildNumber   *[]string
	FilterBundleName    *[]string
	FilterArch          *[]string
	FilterSymbolSource  *[]string
	FilterOrigin        *[]string
	FilterOriginVersion *[]string
	FilterFilename      *string
	FilterDebugId       *string
	FilterGnuBuildId    *string
	FilterGoBuildId     *string
	FilterFileHash      *string
}

// NewDeleteSourcemapsOptionalParameters creates an empty struct for parameters.
func NewDeleteSourcemapsOptionalParameters() *DeleteSourcemapsOptionalParameters {
	this := DeleteSourcemapsOptionalParameters{}
	return &this
}

// WithFilterService sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterService(filterService []string) *DeleteSourcemapsOptionalParameters {
	r.FilterService = &filterService
	return r
}

// WithFilterVersion sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterVersion(filterVersion []string) *DeleteSourcemapsOptionalParameters {
	r.FilterVersion = &filterVersion
	return r
}

// WithFilterVariant sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterVariant(filterVariant []string) *DeleteSourcemapsOptionalParameters {
	r.FilterVariant = &filterVariant
	return r
}

// WithFilterId sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterId(filterId []string) *DeleteSourcemapsOptionalParameters {
	r.FilterId = &filterId
	return r
}

// WithFilterBuildId sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterBuildId(filterBuildId []string) *DeleteSourcemapsOptionalParameters {
	r.FilterBuildId = &filterBuildId
	return r
}

// WithFilterUuid sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterUuid(filterUuid []string) *DeleteSourcemapsOptionalParameters {
	r.FilterUuid = &filterUuid
	return r
}

// WithFilterPlatform sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterPlatform(filterPlatform []string) *DeleteSourcemapsOptionalParameters {
	r.FilterPlatform = &filterPlatform
	return r
}

// WithFilterBuildNumber sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterBuildNumber(filterBuildNumber []string) *DeleteSourcemapsOptionalParameters {
	r.FilterBuildNumber = &filterBuildNumber
	return r
}

// WithFilterBundleName sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterBundleName(filterBundleName []string) *DeleteSourcemapsOptionalParameters {
	r.FilterBundleName = &filterBundleName
	return r
}

// WithFilterArch sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterArch(filterArch []string) *DeleteSourcemapsOptionalParameters {
	r.FilterArch = &filterArch
	return r
}

// WithFilterSymbolSource sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterSymbolSource(filterSymbolSource []string) *DeleteSourcemapsOptionalParameters {
	r.FilterSymbolSource = &filterSymbolSource
	return r
}

// WithFilterOrigin sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterOrigin(filterOrigin []string) *DeleteSourcemapsOptionalParameters {
	r.FilterOrigin = &filterOrigin
	return r
}

// WithFilterOriginVersion sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterOriginVersion(filterOriginVersion []string) *DeleteSourcemapsOptionalParameters {
	r.FilterOriginVersion = &filterOriginVersion
	return r
}

// WithFilterFilename sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterFilename(filterFilename string) *DeleteSourcemapsOptionalParameters {
	r.FilterFilename = &filterFilename
	return r
}

// WithFilterDebugId sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterDebugId(filterDebugId string) *DeleteSourcemapsOptionalParameters {
	r.FilterDebugId = &filterDebugId
	return r
}

// WithFilterGnuBuildId sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterGnuBuildId(filterGnuBuildId string) *DeleteSourcemapsOptionalParameters {
	r.FilterGnuBuildId = &filterGnuBuildId
	return r
}

// WithFilterGoBuildId sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterGoBuildId(filterGoBuildId string) *DeleteSourcemapsOptionalParameters {
	r.FilterGoBuildId = &filterGoBuildId
	return r
}

// WithFilterFileHash sets the corresponding parameter name and returns the struct.
func (r *DeleteSourcemapsOptionalParameters) WithFilterFileHash(filterFileHash string) *DeleteSourcemapsOptionalParameters {
	r.FilterFileHash = &filterFileHash
	return r
}

// DeleteSourcemaps Delete source maps.
// Deletes source maps matching the specified filter criteria. Supports
// dry-run mode to preview which source maps would be deleted without
// performing the actual deletion.
func (a *RUMApi) DeleteSourcemaps(ctx _context.Context, mapkind SourcemapMapKind, dryRun bool, o ...DeleteSourcemapsOptionalParameters) (SourcemapsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		localVarReturnValue SourcemapsResponse
		optionalParams      DeleteSourcemapsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type DeleteSourcemapsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.DeleteSourcemaps"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.DeleteSourcemaps")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourcemaps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("mapkind", datadog.ParameterToString(mapkind, ""))
	localVarQueryParams.Add("dry_run", datadog.ParameterToString(dryRun, ""))
	if optionalParams.FilterService != nil {
		t := *optionalParams.FilterService
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[service]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[service]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterVersion != nil {
		t := *optionalParams.FilterVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[version]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[version]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterVariant != nil {
		t := *optionalParams.FilterVariant
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[variant]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[variant]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterId != nil {
		t := *optionalParams.FilterId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[id]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[id]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBuildId != nil {
		t := *optionalParams.FilterBuildId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[build_id]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[build_id]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterUuid != nil {
		t := *optionalParams.FilterUuid
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[uuid]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[uuid]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterPlatform != nil {
		t := *optionalParams.FilterPlatform
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[platform]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[platform]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBuildNumber != nil {
		t := *optionalParams.FilterBuildNumber
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[build_number]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[build_number]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBundleName != nil {
		t := *optionalParams.FilterBundleName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[bundle_name]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[bundle_name]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterArch != nil {
		t := *optionalParams.FilterArch
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[arch]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[arch]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterSymbolSource != nil {
		t := *optionalParams.FilterSymbolSource
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[symbol_source]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[symbol_source]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterOrigin != nil {
		t := *optionalParams.FilterOrigin
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[origin]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[origin]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterOriginVersion != nil {
		t := *optionalParams.FilterOriginVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[origin_version]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[origin_version]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterFilename != nil {
		localVarQueryParams.Add("filter[filename]", datadog.ParameterToString(*optionalParams.FilterFilename, ""))
	}
	if optionalParams.FilterDebugId != nil {
		localVarQueryParams.Add("filter[debug_id]", datadog.ParameterToString(*optionalParams.FilterDebugId, ""))
	}
	if optionalParams.FilterGnuBuildId != nil {
		localVarQueryParams.Add("filter[gnu_build_id]", datadog.ParameterToString(*optionalParams.FilterGnuBuildId, ""))
	}
	if optionalParams.FilterGoBuildId != nil {
		localVarQueryParams.Add("filter[go_build_id]", datadog.ParameterToString(*optionalParams.FilterGoBuildId, ""))
	}
	if optionalParams.FilterFileHash != nil {
		localVarQueryParams.Add("filter[file_hash]", datadog.ParameterToString(*optionalParams.FilterFileHash, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetRUMApplication Get a RUM application.
// Get the RUM application with given ID in your organization.
func (a *RUMApi) GetRUMApplication(ctx _context.Context, id string) (RUMApplicationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue RUMApplicationResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.GetRUMApplication")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/applications/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetRUMApplications List all the RUM applications.
// List all the RUM applications in your organization.
func (a *RUMApi) GetRUMApplications(ctx _context.Context) (RUMApplicationsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue RUMApplicationsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.GetRUMApplications")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/applications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetServiceRepositoryInfo Get service repository information.
// Returns the repository URL and commit SHA associated with a given service and version.
func (a *RUMApi) GetServiceRepositoryInfo(ctx _context.Context, body ServiceRepositoryInfoRequest) (ServiceRepositoryInfoResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue ServiceRepositoryInfoResponse
	)

	operationId := "v2.GetServiceRepositoryInfo"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.GetServiceRepositoryInfo")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourcemaps/service_repository_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 500 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// GetSourcemaps Get a JavaScript source map.
// Retrieves the content of a specific JavaScript source map file by its
// filename, service name, and version.
func (a *RUMApi) GetSourcemaps(ctx _context.Context, filename string, service string, version string) (SourcemapFileResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SourcemapFileResponse
	)

	operationId := "v2.GetSourcemaps"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.GetSourcemaps")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourcemaps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("filename", datadog.ParameterToString(filename, ""))
	localVarQueryParams.Add("service", datadog.ParameterToString(service, ""))
	localVarQueryParams.Add("version", datadog.ParameterToString(version, ""))
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListRUMEventsOptionalParameters holds optional parameters for ListRUMEvents.
type ListRUMEventsOptionalParameters struct {
	FilterQuery *string
	FilterFrom  *time.Time
	FilterTo    *time.Time
	Sort        *RUMSort
	PageCursor  *string
	PageLimit   *int32
}

// NewListRUMEventsOptionalParameters creates an empty struct for parameters.
func NewListRUMEventsOptionalParameters() *ListRUMEventsOptionalParameters {
	this := ListRUMEventsOptionalParameters{}
	return &this
}

// WithFilterQuery sets the corresponding parameter name and returns the struct.
func (r *ListRUMEventsOptionalParameters) WithFilterQuery(filterQuery string) *ListRUMEventsOptionalParameters {
	r.FilterQuery = &filterQuery
	return r
}

// WithFilterFrom sets the corresponding parameter name and returns the struct.
func (r *ListRUMEventsOptionalParameters) WithFilterFrom(filterFrom time.Time) *ListRUMEventsOptionalParameters {
	r.FilterFrom = &filterFrom
	return r
}

// WithFilterTo sets the corresponding parameter name and returns the struct.
func (r *ListRUMEventsOptionalParameters) WithFilterTo(filterTo time.Time) *ListRUMEventsOptionalParameters {
	r.FilterTo = &filterTo
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListRUMEventsOptionalParameters) WithSort(sort RUMSort) *ListRUMEventsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithPageCursor sets the corresponding parameter name and returns the struct.
func (r *ListRUMEventsOptionalParameters) WithPageCursor(pageCursor string) *ListRUMEventsOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListRUMEventsOptionalParameters) WithPageLimit(pageLimit int32) *ListRUMEventsOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// ListRUMEvents Get a list of RUM events.
// List endpoint returns events that match a RUM search query.
// [Results are paginated][1].
//
// Use this endpoint to see your latest RUM events.
//
// [1]: https://docs.datadoghq.com/logs/guide/collect-multiple-logs-with-pagination
func (a *RUMApi) ListRUMEvents(ctx _context.Context, o ...ListRUMEventsOptionalParameters) (RUMEventsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue RUMEventsResponse
		optionalParams      ListRUMEventsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListRUMEventsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.ListRUMEvents")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterQuery != nil {
		localVarQueryParams.Add("filter[query]", datadog.ParameterToString(*optionalParams.FilterQuery, ""))
	}
	if optionalParams.FilterFrom != nil {
		localVarQueryParams.Add("filter[from]", datadog.ParameterToString(*optionalParams.FilterFrom, ""))
	}
	if optionalParams.FilterTo != nil {
		localVarQueryParams.Add("filter[to]", datadog.ParameterToString(*optionalParams.FilterTo, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.PageCursor != nil {
		localVarQueryParams.Add("page[cursor]", datadog.ParameterToString(*optionalParams.PageCursor, ""))
	}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListRUMEventsWithPagination provides a paginated version of ListRUMEvents returning a channel with all items.
func (a *RUMApi) ListRUMEventsWithPagination(ctx _context.Context, o ...ListRUMEventsOptionalParameters) (<-chan datadog.PaginationResult[RUMEvent], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(10)
	if len(o) == 0 {
		o = append(o, ListRUMEventsOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_

	items := make(chan datadog.PaginationResult[RUMEvent], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListRUMEvents(ctx, o...)
			if err != nil {
				var returnItem RUMEvent
				items <- datadog.PaginationResult[RUMEvent]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[RUMEvent]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) == 0 {
				break
			}
			cursorMeta, ok := resp.GetMetaOk()
			if !ok {
				break
			}
			cursorMetaPage, ok := cursorMeta.GetPageOk()
			if !ok {
				break
			}
			cursorMetaPageAfter, ok := cursorMetaPage.GetAfterOk()
			if !ok {
				break
			}

			o[0].PageCursor = cursorMetaPageAfter
		}
		close(items)
	}()
	return items, cancel
}

// ListSourcemapsOptionalParameters holds optional parameters for ListSourcemaps.
type ListSourcemapsOptionalParameters struct {
	Mapkind             *SourcemapMapKind
	PageSize            *int64
	PageNumber          *int64
	FilterService       *[]string
	FilterVersion       *[]string
	FilterVariant       *[]string
	FilterId            *[]string
	FilterBuildId       *[]string
	FilterUuid          *[]string
	FilterPlatform      *[]string
	FilterBuildNumber   *[]string
	FilterBundleName    *[]string
	FilterArch          *[]string
	FilterSymbolSource  *[]string
	FilterOrigin        *[]string
	FilterOriginVersion *[]string
	FilterFilename      *string
	FilterDebugId       *string
	FilterGnuBuildId    *string
	FilterGoBuildId     *string
	FilterFileHash      *string
}

// NewListSourcemapsOptionalParameters creates an empty struct for parameters.
func NewListSourcemapsOptionalParameters() *ListSourcemapsOptionalParameters {
	this := ListSourcemapsOptionalParameters{}
	return &this
}

// WithMapkind sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithMapkind(mapkind SourcemapMapKind) *ListSourcemapsOptionalParameters {
	r.Mapkind = &mapkind
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithPageSize(pageSize int64) *ListSourcemapsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithPageNumber(pageNumber int64) *ListSourcemapsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithFilterService sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterService(filterService []string) *ListSourcemapsOptionalParameters {
	r.FilterService = &filterService
	return r
}

// WithFilterVersion sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterVersion(filterVersion []string) *ListSourcemapsOptionalParameters {
	r.FilterVersion = &filterVersion
	return r
}

// WithFilterVariant sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterVariant(filterVariant []string) *ListSourcemapsOptionalParameters {
	r.FilterVariant = &filterVariant
	return r
}

// WithFilterId sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterId(filterId []string) *ListSourcemapsOptionalParameters {
	r.FilterId = &filterId
	return r
}

// WithFilterBuildId sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterBuildId(filterBuildId []string) *ListSourcemapsOptionalParameters {
	r.FilterBuildId = &filterBuildId
	return r
}

// WithFilterUuid sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterUuid(filterUuid []string) *ListSourcemapsOptionalParameters {
	r.FilterUuid = &filterUuid
	return r
}

// WithFilterPlatform sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterPlatform(filterPlatform []string) *ListSourcemapsOptionalParameters {
	r.FilterPlatform = &filterPlatform
	return r
}

// WithFilterBuildNumber sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterBuildNumber(filterBuildNumber []string) *ListSourcemapsOptionalParameters {
	r.FilterBuildNumber = &filterBuildNumber
	return r
}

// WithFilterBundleName sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterBundleName(filterBundleName []string) *ListSourcemapsOptionalParameters {
	r.FilterBundleName = &filterBundleName
	return r
}

// WithFilterArch sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterArch(filterArch []string) *ListSourcemapsOptionalParameters {
	r.FilterArch = &filterArch
	return r
}

// WithFilterSymbolSource sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterSymbolSource(filterSymbolSource []string) *ListSourcemapsOptionalParameters {
	r.FilterSymbolSource = &filterSymbolSource
	return r
}

// WithFilterOrigin sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterOrigin(filterOrigin []string) *ListSourcemapsOptionalParameters {
	r.FilterOrigin = &filterOrigin
	return r
}

// WithFilterOriginVersion sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterOriginVersion(filterOriginVersion []string) *ListSourcemapsOptionalParameters {
	r.FilterOriginVersion = &filterOriginVersion
	return r
}

// WithFilterFilename sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterFilename(filterFilename string) *ListSourcemapsOptionalParameters {
	r.FilterFilename = &filterFilename
	return r
}

// WithFilterDebugId sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterDebugId(filterDebugId string) *ListSourcemapsOptionalParameters {
	r.FilterDebugId = &filterDebugId
	return r
}

// WithFilterGnuBuildId sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterGnuBuildId(filterGnuBuildId string) *ListSourcemapsOptionalParameters {
	r.FilterGnuBuildId = &filterGnuBuildId
	return r
}

// WithFilterGoBuildId sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterGoBuildId(filterGoBuildId string) *ListSourcemapsOptionalParameters {
	r.FilterGoBuildId = &filterGoBuildId
	return r
}

// WithFilterFileHash sets the corresponding parameter name and returns the struct.
func (r *ListSourcemapsOptionalParameters) WithFilterFileHash(filterFileHash string) *ListSourcemapsOptionalParameters {
	r.FilterFileHash = &filterFileHash
	return r
}

// ListSourcemaps List source maps.
// Retrieves a paginated list of source maps matching the specified filter criteria.
func (a *RUMApi) ListSourcemaps(ctx _context.Context, o ...ListSourcemapsOptionalParameters) (ListSourcemapsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListSourcemapsResponse
		optionalParams      ListSourcemapsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListSourcemapsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListSourcemaps"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.ListSourcemaps")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourcemaps/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Mapkind != nil {
		localVarQueryParams.Add("mapkind", datadog.ParameterToString(*optionalParams.Mapkind, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.FilterService != nil {
		t := *optionalParams.FilterService
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[service]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[service]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterVersion != nil {
		t := *optionalParams.FilterVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[version]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[version]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterVariant != nil {
		t := *optionalParams.FilterVariant
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[variant]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[variant]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterId != nil {
		t := *optionalParams.FilterId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[id]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[id]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBuildId != nil {
		t := *optionalParams.FilterBuildId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[build_id]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[build_id]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterUuid != nil {
		t := *optionalParams.FilterUuid
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[uuid]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[uuid]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterPlatform != nil {
		t := *optionalParams.FilterPlatform
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[platform]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[platform]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBuildNumber != nil {
		t := *optionalParams.FilterBuildNumber
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[build_number]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[build_number]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBundleName != nil {
		t := *optionalParams.FilterBundleName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[bundle_name]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[bundle_name]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterArch != nil {
		t := *optionalParams.FilterArch
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[arch]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[arch]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterSymbolSource != nil {
		t := *optionalParams.FilterSymbolSource
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[symbol_source]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[symbol_source]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterOrigin != nil {
		t := *optionalParams.FilterOrigin
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[origin]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[origin]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterOriginVersion != nil {
		t := *optionalParams.FilterOriginVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[origin_version]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[origin_version]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterFilename != nil {
		localVarQueryParams.Add("filter[filename]", datadog.ParameterToString(*optionalParams.FilterFilename, ""))
	}
	if optionalParams.FilterDebugId != nil {
		localVarQueryParams.Add("filter[debug_id]", datadog.ParameterToString(*optionalParams.FilterDebugId, ""))
	}
	if optionalParams.FilterGnuBuildId != nil {
		localVarQueryParams.Add("filter[gnu_build_id]", datadog.ParameterToString(*optionalParams.FilterGnuBuildId, ""))
	}
	if optionalParams.FilterGoBuildId != nil {
		localVarQueryParams.Add("filter[go_build_id]", datadog.ParameterToString(*optionalParams.FilterGoBuildId, ""))
	}
	if optionalParams.FilterFileHash != nil {
		localVarQueryParams.Add("filter[file_hash]", datadog.ParameterToString(*optionalParams.FilterFileHash, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 413 || localVarHTTPResponse.StatusCode == 500 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RestoreSourcemapsOptionalParameters holds optional parameters for RestoreSourcemaps.
type RestoreSourcemapsOptionalParameters struct {
	FilterService       *[]string
	FilterVersion       *[]string
	FilterVariant       *[]string
	FilterId            *[]string
	FilterBuildId       *[]string
	FilterUuid          *[]string
	FilterPlatform      *[]string
	FilterBuildNumber   *[]string
	FilterBundleName    *[]string
	FilterArch          *[]string
	FilterSymbolSource  *[]string
	FilterOrigin        *[]string
	FilterOriginVersion *[]string
	FilterFilename      *string
	FilterDebugId       *string
	FilterGnuBuildId    *string
	FilterGoBuildId     *string
	FilterFileHash      *string
}

// NewRestoreSourcemapsOptionalParameters creates an empty struct for parameters.
func NewRestoreSourcemapsOptionalParameters() *RestoreSourcemapsOptionalParameters {
	this := RestoreSourcemapsOptionalParameters{}
	return &this
}

// WithFilterService sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterService(filterService []string) *RestoreSourcemapsOptionalParameters {
	r.FilterService = &filterService
	return r
}

// WithFilterVersion sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterVersion(filterVersion []string) *RestoreSourcemapsOptionalParameters {
	r.FilterVersion = &filterVersion
	return r
}

// WithFilterVariant sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterVariant(filterVariant []string) *RestoreSourcemapsOptionalParameters {
	r.FilterVariant = &filterVariant
	return r
}

// WithFilterId sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterId(filterId []string) *RestoreSourcemapsOptionalParameters {
	r.FilterId = &filterId
	return r
}

// WithFilterBuildId sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterBuildId(filterBuildId []string) *RestoreSourcemapsOptionalParameters {
	r.FilterBuildId = &filterBuildId
	return r
}

// WithFilterUuid sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterUuid(filterUuid []string) *RestoreSourcemapsOptionalParameters {
	r.FilterUuid = &filterUuid
	return r
}

// WithFilterPlatform sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterPlatform(filterPlatform []string) *RestoreSourcemapsOptionalParameters {
	r.FilterPlatform = &filterPlatform
	return r
}

// WithFilterBuildNumber sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterBuildNumber(filterBuildNumber []string) *RestoreSourcemapsOptionalParameters {
	r.FilterBuildNumber = &filterBuildNumber
	return r
}

// WithFilterBundleName sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterBundleName(filterBundleName []string) *RestoreSourcemapsOptionalParameters {
	r.FilterBundleName = &filterBundleName
	return r
}

// WithFilterArch sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterArch(filterArch []string) *RestoreSourcemapsOptionalParameters {
	r.FilterArch = &filterArch
	return r
}

// WithFilterSymbolSource sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterSymbolSource(filterSymbolSource []string) *RestoreSourcemapsOptionalParameters {
	r.FilterSymbolSource = &filterSymbolSource
	return r
}

// WithFilterOrigin sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterOrigin(filterOrigin []string) *RestoreSourcemapsOptionalParameters {
	r.FilterOrigin = &filterOrigin
	return r
}

// WithFilterOriginVersion sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterOriginVersion(filterOriginVersion []string) *RestoreSourcemapsOptionalParameters {
	r.FilterOriginVersion = &filterOriginVersion
	return r
}

// WithFilterFilename sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterFilename(filterFilename string) *RestoreSourcemapsOptionalParameters {
	r.FilterFilename = &filterFilename
	return r
}

// WithFilterDebugId sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterDebugId(filterDebugId string) *RestoreSourcemapsOptionalParameters {
	r.FilterDebugId = &filterDebugId
	return r
}

// WithFilterGnuBuildId sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterGnuBuildId(filterGnuBuildId string) *RestoreSourcemapsOptionalParameters {
	r.FilterGnuBuildId = &filterGnuBuildId
	return r
}

// WithFilterGoBuildId sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterGoBuildId(filterGoBuildId string) *RestoreSourcemapsOptionalParameters {
	r.FilterGoBuildId = &filterGoBuildId
	return r
}

// WithFilterFileHash sets the corresponding parameter name and returns the struct.
func (r *RestoreSourcemapsOptionalParameters) WithFilterFileHash(filterFileHash string) *RestoreSourcemapsOptionalParameters {
	r.FilterFileHash = &filterFileHash
	return r
}

// RestoreSourcemaps Restore source maps.
// Restores previously deleted source maps matching the specified filter
// criteria. Supports dry-run mode to preview which source maps would be
// restored without performing the actual restoration.
func (a *RUMApi) RestoreSourcemaps(ctx _context.Context, mapkind SourcemapMapKind, dryRun bool, o ...RestoreSourcemapsOptionalParameters) (SourcemapsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue SourcemapsResponse
		optionalParams      RestoreSourcemapsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type RestoreSourcemapsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.RestoreSourcemaps"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.RestoreSourcemaps")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sourcemaps/restore"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("mapkind", datadog.ParameterToString(mapkind, ""))
	localVarQueryParams.Add("dry_run", datadog.ParameterToString(dryRun, ""))
	if optionalParams.FilterService != nil {
		t := *optionalParams.FilterService
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[service]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[service]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterVersion != nil {
		t := *optionalParams.FilterVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[version]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[version]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterVariant != nil {
		t := *optionalParams.FilterVariant
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[variant]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[variant]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterId != nil {
		t := *optionalParams.FilterId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[id]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[id]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBuildId != nil {
		t := *optionalParams.FilterBuildId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[build_id]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[build_id]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterUuid != nil {
		t := *optionalParams.FilterUuid
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[uuid]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[uuid]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterPlatform != nil {
		t := *optionalParams.FilterPlatform
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[platform]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[platform]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBuildNumber != nil {
		t := *optionalParams.FilterBuildNumber
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[build_number]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[build_number]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterBundleName != nil {
		t := *optionalParams.FilterBundleName
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[bundle_name]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[bundle_name]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterArch != nil {
		t := *optionalParams.FilterArch
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[arch]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[arch]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterSymbolSource != nil {
		t := *optionalParams.FilterSymbolSource
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[symbol_source]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[symbol_source]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterOrigin != nil {
		t := *optionalParams.FilterOrigin
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[origin]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[origin]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterOriginVersion != nil {
		t := *optionalParams.FilterOriginVersion
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[origin_version]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[origin_version]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.FilterFilename != nil {
		localVarQueryParams.Add("filter[filename]", datadog.ParameterToString(*optionalParams.FilterFilename, ""))
	}
	if optionalParams.FilterDebugId != nil {
		localVarQueryParams.Add("filter[debug_id]", datadog.ParameterToString(*optionalParams.FilterDebugId, ""))
	}
	if optionalParams.FilterGnuBuildId != nil {
		localVarQueryParams.Add("filter[gnu_build_id]", datadog.ParameterToString(*optionalParams.FilterGnuBuildId, ""))
	}
	if optionalParams.FilterGoBuildId != nil {
		localVarQueryParams.Add("filter[go_build_id]", datadog.ParameterToString(*optionalParams.FilterGoBuildId, ""))
	}
	if optionalParams.FilterFileHash != nil {
		localVarQueryParams.Add("filter[file_hash]", datadog.ParameterToString(*optionalParams.FilterFileHash, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 500 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// SearchRUMEvents Search RUM events.
// List endpoint returns RUM events that match a RUM search query.
// [Results are paginated][1].
//
// Use this endpoint to build complex RUM events filtering and search.
//
// [1]: https://docs.datadoghq.com/logs/guide/collect-multiple-logs-with-pagination
func (a *RUMApi) SearchRUMEvents(ctx _context.Context, body RUMSearchEventsRequest) (RUMEventsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue RUMEventsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.SearchRUMEvents")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/events/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// SearchRUMEventsWithPagination provides a paginated version of SearchRUMEvents returning a channel with all items.
func (a *RUMApi) SearchRUMEventsWithPagination(ctx _context.Context, body RUMSearchEventsRequest) (<-chan datadog.PaginationResult[RUMEvent], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(10)
	if body.Page == nil {
		body.Page = NewRUMQueryPageOptions()
	}
	if body.Page.Limit == nil {
		// int32
		body.Page.Limit = &pageSize_
	} else {
		pageSize_ = *body.Page.Limit
	}

	items := make(chan datadog.PaginationResult[RUMEvent], pageSize_)
	go func() {
		for {
			resp, _, err := a.SearchRUMEvents(ctx, body)
			if err != nil {
				var returnItem RUMEvent
				items <- datadog.PaginationResult[RUMEvent]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[RUMEvent]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) == 0 {
				break
			}
			cursorMeta, ok := resp.GetMetaOk()
			if !ok {
				break
			}
			cursorMetaPage, ok := cursorMeta.GetPageOk()
			if !ok {
				break
			}
			cursorMetaPageAfter, ok := cursorMetaPage.GetAfterOk()
			if !ok {
				break
			}

			body.Page.Cursor = cursorMetaPageAfter
		}
		close(items)
	}()
	return items, cancel
}

// UpdateRUMApplication Update a RUM application.
// Update the RUM application with given ID in your organization.
func (a *RUMApi) UpdateRUMApplication(ctx _context.Context, id string, body RUMApplicationUpdateRequest) (RUMApplicationResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue RUMApplicationResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RUMApi.UpdateRUMApplication")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/applications/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
	if a.Client.Cfg.DelegatedTokenConfig != nil {
		err = datadog.UseDelegatedTokenAuth(ctx, &localVarHeaderParams, a.Client.Cfg.DelegatedTokenConfig)
		if err != nil {
			return localVarReturnValue, nil, err
		}
	} else {
		datadog.SetAuthKeys(
			ctx,
			&localVarHeaderParams,
			[2]string{"apiKeyAuth", "DD-API-KEY"},
			[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
		)
	}
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := datadog.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 422 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewRUMApi Returns NewRUMApi.
func NewRUMApi(client *datadog.APIClient) *RUMApi {
	return &RUMApi{
		Client: client,
	}
}
