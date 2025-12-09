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

// TestOptimizationApi service type
type TestOptimizationApi datadog.Service

// SearchFlakyTestsOptionalParameters holds optional parameters for SearchFlakyTests.
type SearchFlakyTestsOptionalParameters struct {
	Body *FlakyTestsSearchRequest
}

// NewSearchFlakyTestsOptionalParameters creates an empty struct for parameters.
func NewSearchFlakyTestsOptionalParameters() *SearchFlakyTestsOptionalParameters {
	this := SearchFlakyTestsOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SearchFlakyTestsOptionalParameters) WithBody(body FlakyTestsSearchRequest) *SearchFlakyTestsOptionalParameters {
	r.Body = &body
	return r
}

// SearchFlakyTests Search flaky tests.
// List endpoint returning flaky tests from Flaky Test Management. Results are paginated.
func (a *TestOptimizationApi) SearchFlakyTests(ctx _context.Context, o ...SearchFlakyTestsOptionalParameters) (FlakyTestsSearchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue FlakyTestsSearchResponse
		optionalParams      SearchFlakyTestsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type SearchFlakyTestsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.SearchFlakyTests"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.TestOptimizationApi.SearchFlakyTests")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/test/flaky-test-management/tests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	if optionalParams.Body != nil {
		localVarPostBody = &optionalParams.Body
	}
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

// SearchFlakyTestsWithPagination provides a paginated version of SearchFlakyTests returning a channel with all items.
func (a *TestOptimizationApi) SearchFlakyTestsWithPagination(ctx _context.Context, o ...SearchFlakyTestsOptionalParameters) (<-chan datadog.PaginationResult[FlakyTest], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(10)
	if len(o) == 0 {
		o = append(o, SearchFlakyTestsOptionalParameters{})
	}
	if o[0].Body == nil {
		o[0].Body = NewFlakyTestsSearchRequest()
	}
	if o[0].Body.Data == nil {
		o[0].Body.Data = NewFlakyTestsSearchRequestData()
	}
	if o[0].Body.Data.Attributes == nil {
		o[0].Body.Data.Attributes = NewFlakyTestsSearchRequestAttributes()
	}
	if o[0].Body.Data.Attributes.Page == nil {
		o[0].Body.Data.Attributes.Page = NewFlakyTestsSearchPageOptions()
	}
	if o[0].Body.Data.Attributes.Page.Limit != nil {
		pageSize_ = *o[0].Body.Data.Attributes.Page.Limit
	}
	o[0].Body.Data.Attributes.Page.Limit = &pageSize_

	items := make(chan datadog.PaginationResult[FlakyTest], pageSize_)
	go func() {
		for {
			resp, _, err := a.SearchFlakyTests(ctx, o...)
			if err != nil {
				var returnItem FlakyTest
				items <- datadog.PaginationResult[FlakyTest]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[FlakyTest]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) < int(pageSize_) {
				break
			}
			cursorMeta, ok := resp.GetMetaOk()
			if !ok {
				break
			}
			cursorMetaPagination, ok := cursorMeta.GetPaginationOk()
			if !ok {
				break
			}
			cursorMetaPaginationNextPage, ok := cursorMetaPagination.GetNextPageOk()
			if !ok {
				break
			}

			o[0].Body.Data.Attributes.Page.Cursor = cursorMetaPaginationNextPage
		}
		close(items)
	}()
	return items, cancel
}

// NewTestOptimizationApi Returns NewTestOptimizationApi.
func NewTestOptimizationApi(client *datadog.APIClient) *TestOptimizationApi {
	return &TestOptimizationApi{
		Client: client,
	}
}
