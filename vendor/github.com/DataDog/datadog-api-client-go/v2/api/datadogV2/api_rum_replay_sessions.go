// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumReplaySessionsApi service type
type RumReplaySessionsApi datadog.Service

// GetSegmentsOptionalParameters holds optional parameters for GetSegments.
type GetSegmentsOptionalParameters struct {
	Source      *string
	Ts          *int64
	MaxListSize *int32
	Paging      *string
}

// NewGetSegmentsOptionalParameters creates an empty struct for parameters.
func NewGetSegmentsOptionalParameters() *GetSegmentsOptionalParameters {
	this := GetSegmentsOptionalParameters{}
	return &this
}

// WithSource sets the corresponding parameter name and returns the struct.
func (r *GetSegmentsOptionalParameters) WithSource(source string) *GetSegmentsOptionalParameters {
	r.Source = &source
	return r
}

// WithTs sets the corresponding parameter name and returns the struct.
func (r *GetSegmentsOptionalParameters) WithTs(ts int64) *GetSegmentsOptionalParameters {
	r.Ts = &ts
	return r
}

// WithMaxListSize sets the corresponding parameter name and returns the struct.
func (r *GetSegmentsOptionalParameters) WithMaxListSize(maxListSize int32) *GetSegmentsOptionalParameters {
	r.MaxListSize = &maxListSize
	return r
}

// WithPaging sets the corresponding parameter name and returns the struct.
func (r *GetSegmentsOptionalParameters) WithPaging(paging string) *GetSegmentsOptionalParameters {
	r.Paging = &paging
	return r
}

// GetSegments Get segments.
// Get segments for a view.
func (a *RumReplaySessionsApi) GetSegments(ctx _context.Context, viewId string, sessionId string, o ...GetSegmentsOptionalParameters) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodGet
		localVarPostBody   interface{}
		optionalParams     GetSegmentsOptionalParameters
	)

	if len(o) > 1 {
		return nil, datadog.ReportError("only one argument of type GetSegmentsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RumReplaySessionsApi.GetSegments")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/rum/replay/sessions/{session_id}/views/{view_id}/segments"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{view_id}", _neturl.PathEscape(datadog.ParameterToString(viewId, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{session_id}", _neturl.PathEscape(datadog.ParameterToString(sessionId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Source != nil {
		localVarQueryParams.Add("source", datadog.ParameterToString(*optionalParams.Source, ""))
	}
	if optionalParams.Ts != nil {
		localVarQueryParams.Add("ts", datadog.ParameterToString(*optionalParams.Ts, ""))
	}
	if optionalParams.MaxListSize != nil {
		localVarQueryParams.Add("max_list_size", datadog.ParameterToString(*optionalParams.MaxListSize, ""))
	}
	if optionalParams.Paging != nil {
		localVarQueryParams.Add("paging", datadog.ParameterToString(*optionalParams.Paging, ""))
	}
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
		if localVarHTTPResponse.StatusCode == 429 {
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

// NewRumReplaySessionsApi Returns NewRumReplaySessionsApi.
func NewRumReplaySessionsApi(client *datadog.APIClient) *RumReplaySessionsApi {
	return &RumReplaySessionsApi{
		Client: client,
	}
}
