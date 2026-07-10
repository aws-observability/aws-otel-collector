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

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardsApi service type
type DashboardsApi datadog.Service

// GetDashboardUsage Get usage stats for a dashboard.
// Get usage statistics for a single dashboard. The response includes view counts, the most recent view and edit times, widget counts, and the dashboard quality score. View-count fields depend on Real User Monitoring (RUM) and are `null` or `0` in orgs without RUM.
func (a *DashboardsApi) GetDashboardUsage(ctx _context.Context, dashboardId string) (DashboardUsageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue DashboardUsageResponse
	)

	operationId := "v2.GetDashboardUsage"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DashboardsApi.GetDashboardUsage")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dashboards/{dashboard_id}/usage"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{dashboard_id}", _neturl.PathEscape(datadog.ParameterToString(dashboardId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// ListDashboardsUsageOptionalParameters holds optional parameters for ListDashboardsUsage.
type ListDashboardsUsageOptionalParameters struct {
	PageLimit          *int64
	PageOffset         *int64
	FilterEditedBefore *string
	FilterViewedBefore *string
}

// NewListDashboardsUsageOptionalParameters creates an empty struct for parameters.
func NewListDashboardsUsageOptionalParameters() *ListDashboardsUsageOptionalParameters {
	this := ListDashboardsUsageOptionalParameters{}
	return &this
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListDashboardsUsageOptionalParameters) WithPageLimit(pageLimit int64) *ListDashboardsUsageOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListDashboardsUsageOptionalParameters) WithPageOffset(pageOffset int64) *ListDashboardsUsageOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithFilterEditedBefore sets the corresponding parameter name and returns the struct.
func (r *ListDashboardsUsageOptionalParameters) WithFilterEditedBefore(filterEditedBefore string) *ListDashboardsUsageOptionalParameters {
	r.FilterEditedBefore = &filterEditedBefore
	return r
}

// WithFilterViewedBefore sets the corresponding parameter name and returns the struct.
func (r *ListDashboardsUsageOptionalParameters) WithFilterViewedBefore(filterViewedBefore string) *ListDashboardsUsageOptionalParameters {
	r.FilterViewedBefore = &filterViewedBefore
	return r
}

// ListDashboardsUsage Get usage stats for all dashboards.
// Get paginated usage statistics for every dashboard in the caller's organization. Use `page[limit]` and `page[offset]` to walk the result set. Use `filter[edited_before]` or `filter[viewed_before]` to narrow results by recency. View-count fields depend on Real User Monitoring (RUM) and are `null` or `0` in orgs without RUM.
func (a *DashboardsApi) ListDashboardsUsage(ctx _context.Context, o ...ListDashboardsUsageOptionalParameters) (ListDashboardsUsageResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListDashboardsUsageResponse
		optionalParams      ListDashboardsUsageOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListDashboardsUsageOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListDashboardsUsage"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.DashboardsApi.ListDashboardsUsage")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dashboards/usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.FilterEditedBefore != nil {
		localVarQueryParams.Add("filter[edited_before]", datadog.ParameterToString(*optionalParams.FilterEditedBefore, ""))
	}
	if optionalParams.FilterViewedBefore != nil {
		localVarQueryParams.Add("filter[viewed_before]", datadog.ParameterToString(*optionalParams.FilterViewedBefore, ""))
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

// ListDashboardsUsageWithPagination provides a paginated version of ListDashboardsUsage returning a channel with all items.
func (a *DashboardsApi) ListDashboardsUsageWithPagination(ctx _context.Context, o ...ListDashboardsUsageOptionalParameters) (<-chan datadog.PaginationResult[DashboardUsage], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(250)
	if len(o) == 0 {
		o = append(o, ListDashboardsUsageOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_

	items := make(chan datadog.PaginationResult[DashboardUsage], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListDashboardsUsage(ctx, o...)
			if err != nil {
				var returnItem DashboardUsage
				items <- datadog.PaginationResult[DashboardUsage]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[DashboardUsage]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) < int(pageSize_) {
				break
			}
			if o[0].PageOffset == nil {
				o[0].PageOffset = &pageSize_
			} else {
				pageOffset_ := *o[0].PageOffset + pageSize_
				o[0].PageOffset = &pageOffset_
			}
		}
		close(items)
	}()
	return items, cancel
}

// NewDashboardsApi Returns NewDashboardsApi.
func NewDashboardsApi(client *datadog.APIClient) *DashboardsApi {
	return &DashboardsApi{
		Client: client,
	}
}
