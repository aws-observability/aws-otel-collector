// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	_context "context"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SoftwareCatalogApi service type
type SoftwareCatalogApi datadog.Service

// DeleteCatalogEntity Delete a single entity.
// Delete a single entity in Software Catalog.
func (a *SoftwareCatalogApi) DeleteCatalogEntity(ctx _context.Context, entityId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SoftwareCatalogApi.DeleteCatalogEntity")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/catalog/entity/{entity_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity_id"+"}", _neturl.PathEscape(datadog.ParameterToString(entityId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "*/*"

	datadog.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// ListCatalogEntityOptionalParameters holds optional parameters for ListCatalogEntity.
type ListCatalogEntityOptionalParameters struct {
	PageOffset            *int64
	PageLimit             *int64
	FitlerId              *string
	FitlerRef             *string
	FitlerName            *string
	FitlerKind            *string
	FitlerOwner           *string
	FitlerRelationType    *RelationType
	FitlerExcludeSnapshot *string
	Include               *IncludeType
}

// NewListCatalogEntityOptionalParameters creates an empty struct for parameters.
func NewListCatalogEntityOptionalParameters() *ListCatalogEntityOptionalParameters {
	this := ListCatalogEntityOptionalParameters{}
	return &this
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithPageOffset(pageOffset int64) *ListCatalogEntityOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithPageLimit(pageLimit int64) *ListCatalogEntityOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithFitlerId sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFitlerId(fitlerId string) *ListCatalogEntityOptionalParameters {
	r.FitlerId = &fitlerId
	return r
}

// WithFitlerRef sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFitlerRef(fitlerRef string) *ListCatalogEntityOptionalParameters {
	r.FitlerRef = &fitlerRef
	return r
}

// WithFitlerName sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFitlerName(fitlerName string) *ListCatalogEntityOptionalParameters {
	r.FitlerName = &fitlerName
	return r
}

// WithFitlerKind sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFitlerKind(fitlerKind string) *ListCatalogEntityOptionalParameters {
	r.FitlerKind = &fitlerKind
	return r
}

// WithFitlerOwner sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFitlerOwner(fitlerOwner string) *ListCatalogEntityOptionalParameters {
	r.FitlerOwner = &fitlerOwner
	return r
}

// WithFitlerRelationType sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFitlerRelationType(fitlerRelationType RelationType) *ListCatalogEntityOptionalParameters {
	r.FitlerRelationType = &fitlerRelationType
	return r
}

// WithFitlerExcludeSnapshot sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFitlerExcludeSnapshot(fitlerExcludeSnapshot string) *ListCatalogEntityOptionalParameters {
	r.FitlerExcludeSnapshot = &fitlerExcludeSnapshot
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithInclude(include IncludeType) *ListCatalogEntityOptionalParameters {
	r.Include = &include
	return r
}

// ListCatalogEntity Get a list of entities.
// Get a list of entities from Software Catalog.
func (a *SoftwareCatalogApi) ListCatalogEntity(ctx _context.Context, o ...ListCatalogEntityOptionalParameters) (ListEntityCatalogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListEntityCatalogResponse
		optionalParams      ListCatalogEntityOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListCatalogEntityOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SoftwareCatalogApi.ListCatalogEntity")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/catalog/entity"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	if optionalParams.FitlerId != nil {
		localVarQueryParams.Add("fitler[id]", datadog.ParameterToString(*optionalParams.FitlerId, ""))
	}
	if optionalParams.FitlerRef != nil {
		localVarQueryParams.Add("fitler[ref]", datadog.ParameterToString(*optionalParams.FitlerRef, ""))
	}
	if optionalParams.FitlerName != nil {
		localVarQueryParams.Add("fitler[name]", datadog.ParameterToString(*optionalParams.FitlerName, ""))
	}
	if optionalParams.FitlerKind != nil {
		localVarQueryParams.Add("fitler[kind]", datadog.ParameterToString(*optionalParams.FitlerKind, ""))
	}
	if optionalParams.FitlerOwner != nil {
		localVarQueryParams.Add("fitler[owner]", datadog.ParameterToString(*optionalParams.FitlerOwner, ""))
	}
	if optionalParams.FitlerRelationType != nil {
		localVarQueryParams.Add("fitler[relation][type]", datadog.ParameterToString(*optionalParams.FitlerRelationType, ""))
	}
	if optionalParams.FitlerExcludeSnapshot != nil {
		localVarQueryParams.Add("fitler[exclude_snapshot]", datadog.ParameterToString(*optionalParams.FitlerExcludeSnapshot, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
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

// ListCatalogEntityWithPagination provides a paginated version of ListCatalogEntity returning a channel with all items.
func (a *SoftwareCatalogApi) ListCatalogEntityWithPagination(ctx _context.Context, o ...ListCatalogEntityOptionalParameters) (<-chan datadog.PaginationResult[EntityData], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(100)
	if len(o) == 0 {
		o = append(o, ListCatalogEntityOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_
	page_ := int64(0)
	o[0].PageOffset = &page_

	items := make(chan datadog.PaginationResult[EntityData], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListCatalogEntity(ctx, o...)
			if err != nil {
				var returnItem EntityData
				items <- datadog.PaginationResult[EntityData]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[EntityData]{Item: item, Error: nil}:
				case <-ctx.Done():
					close(items)
					return
				}
			}
			if len(results) < int(pageSize_) {
				break
			}
			pageOffset_ := *o[0].PageOffset + 1
			o[0].PageOffset = &pageOffset_
		}
		close(items)
	}()
	return items, cancel
}

// UpsertCatalogEntity Create or update entities.
// Create or update entities in Software Catalog.
func (a *SoftwareCatalogApi) UpsertCatalogEntity(ctx _context.Context, body UpsertCatalogEntityRequest) (UpsertCatalogEntityResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue UpsertCatalogEntityResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SoftwareCatalogApi.UpsertCatalogEntity")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/catalog/entity"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = &body
	datadog.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"apiKeyAuth", "DD-API-KEY"},
		[2]string{"appKeyAuth", "DD-APPLICATION-KEY"},
	)
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

// NewSoftwareCatalogApi Returns NewSoftwareCatalogApi.
func NewSoftwareCatalogApi(client *datadog.APIClient) *SoftwareCatalogApi {
	return &SoftwareCatalogApi{
		Client: client,
	}
}
