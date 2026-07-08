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

// GovernanceInsightsApi service type
type GovernanceInsightsApi datadog.Service

// ListGovernanceInsightsOptionalParameters holds optional parameters for ListGovernanceInsights.
type ListGovernanceInsightsOptionalParameters struct {
	WithValues    *bool
	OrgUuid       *string
	FilterProduct *[]string
}

// NewListGovernanceInsightsOptionalParameters creates an empty struct for parameters.
func NewListGovernanceInsightsOptionalParameters() *ListGovernanceInsightsOptionalParameters {
	this := ListGovernanceInsightsOptionalParameters{}
	return &this
}

// WithWithValues sets the corresponding parameter name and returns the struct.
func (r *ListGovernanceInsightsOptionalParameters) WithWithValues(withValues bool) *ListGovernanceInsightsOptionalParameters {
	r.WithValues = &withValues
	return r
}

// WithOrgUuid sets the corresponding parameter name and returns the struct.
func (r *ListGovernanceInsightsOptionalParameters) WithOrgUuid(orgUuid string) *ListGovernanceInsightsOptionalParameters {
	r.OrgUuid = &orgUuid
	return r
}

// WithFilterProduct sets the corresponding parameter name and returns the struct.
func (r *ListGovernanceInsightsOptionalParameters) WithFilterProduct(filterProduct []string) *ListGovernanceInsightsOptionalParameters {
	r.FilterProduct = &filterProduct
	return r
}

// ListGovernanceInsights List governance insights.
// Retrieve the list of governance insights available to the organization. By default, only
// insight metadata is returned; pass `withValues=true` to also compute and include each
// insight's current and previous values. Insights can be filtered by product.
func (a *GovernanceInsightsApi) ListGovernanceInsights(ctx _context.Context, o ...ListGovernanceInsightsOptionalParameters) (GovernanceInsightsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue GovernanceInsightsResponse
		optionalParams      ListGovernanceInsightsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListGovernanceInsightsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListGovernanceInsights"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.GovernanceInsightsApi.ListGovernanceInsights")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/governance/insights"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.WithValues != nil {
		localVarQueryParams.Add("withValues", datadog.ParameterToString(*optionalParams.WithValues, ""))
	}
	if optionalParams.OrgUuid != nil {
		localVarQueryParams.Add("orgUuid", datadog.ParameterToString(*optionalParams.OrgUuid, ""))
	}
	if optionalParams.FilterProduct != nil {
		t := *optionalParams.FilterProduct
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[product]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[product]", datadog.ParameterToString(t, "multi"))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 {
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

// NewGovernanceInsightsApi Returns NewGovernanceInsightsApi.
func NewGovernanceInsightsApi(client *datadog.APIClient) *GovernanceInsightsApi {
	return &GovernanceInsightsApi{
		Client: client,
	}
}
