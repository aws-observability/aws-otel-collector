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

// ContainerImagesApi service type
type ContainerImagesApi datadog.Service

// ListContainerImagesOptionalParameters holds optional parameters for ListContainerImages.
type ListContainerImagesOptionalParameters struct {
	FilterTags *string
	GroupBy    *string
	Sort       *string
	PageSize   *int32
	PageCursor *string
}

// NewListContainerImagesOptionalParameters creates an empty struct for parameters.
func NewListContainerImagesOptionalParameters() *ListContainerImagesOptionalParameters {
	this := ListContainerImagesOptionalParameters{}
	return &this
}

// WithFilterTags sets the corresponding parameter name and returns the struct.
func (r *ListContainerImagesOptionalParameters) WithFilterTags(filterTags string) *ListContainerImagesOptionalParameters {
	r.FilterTags = &filterTags
	return r
}

// WithGroupBy sets the corresponding parameter name and returns the struct.
func (r *ListContainerImagesOptionalParameters) WithGroupBy(groupBy string) *ListContainerImagesOptionalParameters {
	r.GroupBy = &groupBy
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListContainerImagesOptionalParameters) WithSort(sort string) *ListContainerImagesOptionalParameters {
	r.Sort = &sort
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListContainerImagesOptionalParameters) WithPageSize(pageSize int32) *ListContainerImagesOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageCursor sets the corresponding parameter name and returns the struct.
func (r *ListContainerImagesOptionalParameters) WithPageCursor(pageCursor string) *ListContainerImagesOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

// ListContainerImages Get all Container Images.
// Get all Container Images for your organization.
// **Note**: To enrich the data returned by this endpoint with security scans, see the new [api/v2/security/scanned-assets-metadata](https://docs.datadoghq.com/api/latest/security-monitoring/#list-scanned-assets-metadata) endpoint.
func (a *ContainerImagesApi) ListContainerImages(ctx _context.Context, o ...ListContainerImagesOptionalParameters) (ContainerImagesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ContainerImagesResponse
		optionalParams      ListContainerImagesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListContainerImagesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ContainerImagesApi.ListContainerImages")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/container_images"

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

// ListContainerImagesWithPagination provides a paginated version of ListContainerImages returning a channel with all items.
func (a *ContainerImagesApi) ListContainerImagesWithPagination(ctx _context.Context, o ...ListContainerImagesOptionalParameters) (<-chan datadog.PaginationResult[ContainerImageItem], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(1000)
	if len(o) == 0 {
		o = append(o, ListContainerImagesOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_

	items := make(chan datadog.PaginationResult[ContainerImageItem], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListContainerImages(ctx, o...)
			if err != nil {
				var returnItem ContainerImageItem
				items <- datadog.PaginationResult[ContainerImageItem]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[ContainerImageItem]{Item: item, Error: nil}:
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

// NewContainerImagesApi Returns NewContainerImagesApi.
func NewContainerImagesApi(client *datadog.APIClient) *ContainerImagesApi {
	return &ContainerImagesApi{
		Client: client,
	}
}
