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

// ContainersApi service type
type ContainersApi datadog.Service

// ListContainersOptionalParameters holds optional parameters for ListContainers.
type ListContainersOptionalParameters struct {
	FilterTags *string
	GroupBy    *string
	Sort       *string
	PageSize   *int32
	PageCursor *string
}

// NewListContainersOptionalParameters creates an empty struct for parameters.
func NewListContainersOptionalParameters() *ListContainersOptionalParameters {
	this := ListContainersOptionalParameters{}
	return &this
}

// WithFilterTags sets the corresponding parameter name and returns the struct.
func (r *ListContainersOptionalParameters) WithFilterTags(filterTags string) *ListContainersOptionalParameters {
	r.FilterTags = &filterTags
	return r
}

// WithGroupBy sets the corresponding parameter name and returns the struct.
func (r *ListContainersOptionalParameters) WithGroupBy(groupBy string) *ListContainersOptionalParameters {
	r.GroupBy = &groupBy
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListContainersOptionalParameters) WithSort(sort string) *ListContainersOptionalParameters {
	r.Sort = &sort
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListContainersOptionalParameters) WithPageSize(pageSize int32) *ListContainersOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageCursor sets the corresponding parameter name and returns the struct.
func (r *ListContainersOptionalParameters) WithPageCursor(pageCursor string) *ListContainersOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

// ListContainers Get All Containers.
// Get all containers for your organization.
func (a *ContainersApi) ListContainers(ctx _context.Context, o ...ListContainersOptionalParameters) (ContainersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ContainersResponse
		optionalParams      ListContainersOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListContainersOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ContainersApi.ListContainers")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/containers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterTags != nil {
		localVarQueryParams.Add("filter[tags]", datadog.ParameterToString(*optionalParams.FilterTags, ""))
	}
	if optionalParams.GroupBy != nil {
		localVarQueryParams.Add("group_by", datadog.ParameterToString(*optionalParams.GroupBy, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageCursor != nil {
		localVarQueryParams.Add("page[cursor]", datadog.ParameterToString(*optionalParams.PageCursor, ""))
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

// ListContainersWithPagination provides a paginated version of ListContainers returning a channel with all items.
func (a *ContainersApi) ListContainersWithPagination(ctx _context.Context, o ...ListContainersOptionalParameters) (<-chan datadog.PaginationResult[ContainerItem], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(1000)
	if len(o) == 0 {
		o = append(o, ListContainersOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_

	items := make(chan datadog.PaginationResult[ContainerItem], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListContainers(ctx, o...)
			if err != nil {
				var returnItem ContainerItem
				items <- datadog.PaginationResult[ContainerItem]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[ContainerItem]{Item: item, Error: nil}:
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
			cursorMetaPaginationNextCursor, ok := cursorMetaPagination.GetNextCursorOk()
			if !ok {
				break
			}

			o[0].PageCursor = cursorMetaPaginationNextCursor
		}
		close(items)
	}()
	return items, cancel
}

// NewContainersApi Returns NewContainersApi.
func NewContainersApi(client *datadog.APIClient) *ContainersApi {
	return &ContainersApi{
		Client: client,
	}
}
