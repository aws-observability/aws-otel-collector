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

// CSMAgentsApi service type
type CSMAgentsApi datadog.Service

// ListAllCSMAgentsOptionalParameters holds optional parameters for ListAllCSMAgents.
type ListAllCSMAgentsOptionalParameters struct {
	Page           *int32
	Size           *int32
	Query          *string
	OrderDirection *OrderDirection
}

// NewListAllCSMAgentsOptionalParameters creates an empty struct for parameters.
func NewListAllCSMAgentsOptionalParameters() *ListAllCSMAgentsOptionalParameters {
	this := ListAllCSMAgentsOptionalParameters{}
	return &this
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *ListAllCSMAgentsOptionalParameters) WithPage(page int32) *ListAllCSMAgentsOptionalParameters {
	r.Page = &page
	return r
}

// WithSize sets the corresponding parameter name and returns the struct.
func (r *ListAllCSMAgentsOptionalParameters) WithSize(size int32) *ListAllCSMAgentsOptionalParameters {
	r.Size = &size
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *ListAllCSMAgentsOptionalParameters) WithQuery(query string) *ListAllCSMAgentsOptionalParameters {
	r.Query = &query
	return r
}

// WithOrderDirection sets the corresponding parameter name and returns the struct.
func (r *ListAllCSMAgentsOptionalParameters) WithOrderDirection(orderDirection OrderDirection) *ListAllCSMAgentsOptionalParameters {
	r.OrderDirection = &orderDirection
	return r
}

// ListAllCSMAgents Get all CSM Agents.
// Get the list of all CSM Agents running on your hosts and containers.
func (a *CSMAgentsApi) ListAllCSMAgents(ctx _context.Context, o ...ListAllCSMAgentsOptionalParameters) (CsmAgentsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CsmAgentsResponse
		optionalParams      ListAllCSMAgentsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListAllCSMAgentsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.CSMAgentsApi.ListAllCSMAgents")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/csm/onboarding/agents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*optionalParams.Page, ""))
	}
	if optionalParams.Size != nil {
		localVarQueryParams.Add("size", datadog.ParameterToString(*optionalParams.Size, ""))
	}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", datadog.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.OrderDirection != nil {
		localVarQueryParams.Add("order_direction", datadog.ParameterToString(*optionalParams.OrderDirection, ""))
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
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

// ListAllCSMServerlessAgentsOptionalParameters holds optional parameters for ListAllCSMServerlessAgents.
type ListAllCSMServerlessAgentsOptionalParameters struct {
	Page           *int32
	Size           *int32
	Query          *string
	OrderDirection *OrderDirection
}

// NewListAllCSMServerlessAgentsOptionalParameters creates an empty struct for parameters.
func NewListAllCSMServerlessAgentsOptionalParameters() *ListAllCSMServerlessAgentsOptionalParameters {
	this := ListAllCSMServerlessAgentsOptionalParameters{}
	return &this
}

// WithPage sets the corresponding parameter name and returns the struct.
func (r *ListAllCSMServerlessAgentsOptionalParameters) WithPage(page int32) *ListAllCSMServerlessAgentsOptionalParameters {
	r.Page = &page
	return r
}

// WithSize sets the corresponding parameter name and returns the struct.
func (r *ListAllCSMServerlessAgentsOptionalParameters) WithSize(size int32) *ListAllCSMServerlessAgentsOptionalParameters {
	r.Size = &size
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *ListAllCSMServerlessAgentsOptionalParameters) WithQuery(query string) *ListAllCSMServerlessAgentsOptionalParameters {
	r.Query = &query
	return r
}

// WithOrderDirection sets the corresponding parameter name and returns the struct.
func (r *ListAllCSMServerlessAgentsOptionalParameters) WithOrderDirection(orderDirection OrderDirection) *ListAllCSMServerlessAgentsOptionalParameters {
	r.OrderDirection = &orderDirection
	return r
}

// ListAllCSMServerlessAgents Get all CSM Serverless Agents.
// Get the list of all CSM Serverless Agents running on your hosts and containers.
func (a *CSMAgentsApi) ListAllCSMServerlessAgents(ctx _context.Context, o ...ListAllCSMServerlessAgentsOptionalParameters) (CsmAgentsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue CsmAgentsResponse
		optionalParams      ListAllCSMServerlessAgentsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListAllCSMServerlessAgentsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.CSMAgentsApi.ListAllCSMServerlessAgents")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/csm/onboarding/serverless/agents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.Page != nil {
		localVarQueryParams.Add("page", datadog.ParameterToString(*optionalParams.Page, ""))
	}
	if optionalParams.Size != nil {
		localVarQueryParams.Add("size", datadog.ParameterToString(*optionalParams.Size, ""))
	}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", datadog.ParameterToString(*optionalParams.Query, ""))
	}
	if optionalParams.OrderDirection != nil {
		localVarQueryParams.Add("order_direction", datadog.ParameterToString(*optionalParams.OrderDirection, ""))
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
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

// NewCSMAgentsApi Returns NewCSMAgentsApi.
func NewCSMAgentsApi(client *datadog.APIClient) *CSMAgentsApi {
	return &CSMAgentsApi{
		Client: client,
	}
}
