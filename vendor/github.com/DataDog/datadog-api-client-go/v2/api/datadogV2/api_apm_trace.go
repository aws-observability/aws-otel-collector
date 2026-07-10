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

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APMTraceApi service type
type APMTraceApi datadog.Service

// GetPrunedTraceByIDOptionalParameters holds optional parameters for GetPrunedTraceByID.
type GetPrunedTraceByIDOptionalParameters struct {
	ExpandSpanId          *int64
	TimeHint              *int32
	ForceSource           *string
	IncludePath           *[]string
	TagInclude            *[]string
	TagExclude            *[]string
	OnlyServiceEntrySpans *bool
}

// NewGetPrunedTraceByIDOptionalParameters creates an empty struct for parameters.
func NewGetPrunedTraceByIDOptionalParameters() *GetPrunedTraceByIDOptionalParameters {
	this := GetPrunedTraceByIDOptionalParameters{}
	return &this
}

// WithExpandSpanId sets the corresponding parameter name and returns the struct.
func (r *GetPrunedTraceByIDOptionalParameters) WithExpandSpanId(expandSpanId int64) *GetPrunedTraceByIDOptionalParameters {
	r.ExpandSpanId = &expandSpanId
	return r
}

// WithTimeHint sets the corresponding parameter name and returns the struct.
func (r *GetPrunedTraceByIDOptionalParameters) WithTimeHint(timeHint int32) *GetPrunedTraceByIDOptionalParameters {
	r.TimeHint = &timeHint
	return r
}

// WithForceSource sets the corresponding parameter name and returns the struct.
func (r *GetPrunedTraceByIDOptionalParameters) WithForceSource(forceSource string) *GetPrunedTraceByIDOptionalParameters {
	r.ForceSource = &forceSource
	return r
}

// WithIncludePath sets the corresponding parameter name and returns the struct.
func (r *GetPrunedTraceByIDOptionalParameters) WithIncludePath(includePath []string) *GetPrunedTraceByIDOptionalParameters {
	r.IncludePath = &includePath
	return r
}

// WithTagInclude sets the corresponding parameter name and returns the struct.
func (r *GetPrunedTraceByIDOptionalParameters) WithTagInclude(tagInclude []string) *GetPrunedTraceByIDOptionalParameters {
	r.TagInclude = &tagInclude
	return r
}

// WithTagExclude sets the corresponding parameter name and returns the struct.
func (r *GetPrunedTraceByIDOptionalParameters) WithTagExclude(tagExclude []string) *GetPrunedTraceByIDOptionalParameters {
	r.TagExclude = &tagExclude
	return r
}

// WithOnlyServiceEntrySpans sets the corresponding parameter name and returns the struct.
func (r *GetPrunedTraceByIDOptionalParameters) WithOnlyServiceEntrySpans(onlyServiceEntrySpans bool) *GetPrunedTraceByIDOptionalParameters {
	r.OnlyServiceEntrySpans = &onlyServiceEntrySpans
	return r
}

// GetPrunedTraceByID Get a pruned trace by ID.
// Retrieve a pruned, hierarchical view of an APM trace by its trace ID.
// The trace is summarized as a tree of spans rooted at the trace root and reduced in size
// to keep rendering large traces in the UI practical.
// This endpoint is rate limited to `60` requests per minute per organization.
func (a *APMTraceApi) GetPrunedTraceByID(ctx _context.Context, traceId string, o ...GetPrunedTraceByIDOptionalParameters) (PrunedTraceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue PrunedTraceResponse
		optionalParams      GetPrunedTraceByIDOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetPrunedTraceByIDOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetPrunedTraceByID"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.APMTraceApi.GetPrunedTraceByID")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/pruned_trace/{trace_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{trace_id}", _neturl.PathEscape(datadog.ParameterToString(traceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.ExpandSpanId != nil {
		localVarQueryParams.Add("expand_span_id", datadog.ParameterToString(*optionalParams.ExpandSpanId, ""))
	}
	if optionalParams.TimeHint != nil {
		localVarQueryParams.Add("time_hint", datadog.ParameterToString(*optionalParams.TimeHint, ""))
	}
	if optionalParams.ForceSource != nil {
		localVarQueryParams.Add("force_source", datadog.ParameterToString(*optionalParams.ForceSource, ""))
	}
	if optionalParams.IncludePath != nil {
		t := *optionalParams.IncludePath
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("include_path", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("include_path", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.TagInclude != nil {
		t := *optionalParams.TagInclude
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("tag_include", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("tag_include", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.TagExclude != nil {
		t := *optionalParams.TagExclude
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("tag_exclude", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("tag_exclude", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.OnlyServiceEntrySpans != nil {
		localVarQueryParams.Add("only_service_entry_spans", datadog.ParameterToString(*optionalParams.OnlyServiceEntrySpans, ""))
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 504 {
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

// GetTraceByIDOptionalParameters holds optional parameters for GetTraceByID.
type GetTraceByIDOptionalParameters struct {
	IncludeFields *[]string
}

// NewGetTraceByIDOptionalParameters creates an empty struct for parameters.
func NewGetTraceByIDOptionalParameters() *GetTraceByIDOptionalParameters {
	this := GetTraceByIDOptionalParameters{}
	return &this
}

// WithIncludeFields sets the corresponding parameter name and returns the struct.
func (r *GetTraceByIDOptionalParameters) WithIncludeFields(includeFields []string) *GetTraceByIDOptionalParameters {
	r.IncludeFields = &includeFields
	return r
}

// GetTraceByID Get a trace by ID.
// Retrieve a full APM trace by its trace ID, including every span in the trace.
// Traces are returned from live storage when available and fall back to longer-term storage.
// This endpoint is rate limited to `60` requests per minute per organization.
func (a *APMTraceApi) GetTraceByID(ctx _context.Context, traceId string, o ...GetTraceByIDOptionalParameters) (TraceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue TraceResponse
		optionalParams      GetTraceByIDOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetTraceByIDOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetTraceByID"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.APMTraceApi.GetTraceByID")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/trace/{trace_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{trace_id}", _neturl.PathEscape(datadog.ParameterToString(traceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.IncludeFields != nil {
		t := *optionalParams.IncludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("include_fields", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("include_fields", datadog.ParameterToString(t, "multi"))
		}
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 413 {
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

// NewAPMTraceApi Returns NewAPMTraceApi.
func NewAPMTraceApi(client *datadog.APIClient) *APMTraceApi {
	return &APMTraceApi{
		Client: client,
	}
}
