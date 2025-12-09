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
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{entity_id}", _neturl.PathEscape(datadog.ParameterToString(entityId, "")))

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

// DeleteCatalogKind Delete a single kind.
// Delete a single kind in Software Catalog.
func (a *SoftwareCatalogApi) DeleteCatalogKind(ctx _context.Context, kindId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SoftwareCatalogApi.DeleteCatalogKind")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/catalog/kind/{kind_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{kind_id}", _neturl.PathEscape(datadog.ParameterToString(kindId, "")))

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
	FilterId              *string
	FilterRef             *string
	FilterName            *string
	FilterKind            *string
	FilterOwner           *string
	FilterRelationType    *RelationType
	FilterExcludeSnapshot *string
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

// WithFilterId sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFilterId(filterId string) *ListCatalogEntityOptionalParameters {
	r.FilterId = &filterId
	return r
}

// WithFilterRef sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFilterRef(filterRef string) *ListCatalogEntityOptionalParameters {
	r.FilterRef = &filterRef
	return r
}

// WithFilterName sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFilterName(filterName string) *ListCatalogEntityOptionalParameters {
	r.FilterName = &filterName
	return r
}

// WithFilterKind sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFilterKind(filterKind string) *ListCatalogEntityOptionalParameters {
	r.FilterKind = &filterKind
	return r
}

// WithFilterOwner sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFilterOwner(filterOwner string) *ListCatalogEntityOptionalParameters {
	r.FilterOwner = &filterOwner
	return r
}

// WithFilterRelationType sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFilterRelationType(filterRelationType RelationType) *ListCatalogEntityOptionalParameters {
	r.FilterRelationType = &filterRelationType
	return r
}

// WithFilterExcludeSnapshot sets the corresponding parameter name and returns the struct.
func (r *ListCatalogEntityOptionalParameters) WithFilterExcludeSnapshot(filterExcludeSnapshot string) *ListCatalogEntityOptionalParameters {
	r.FilterExcludeSnapshot = &filterExcludeSnapshot
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
	if optionalParams.FilterId != nil {
		localVarQueryParams.Add("filter[id]", datadog.ParameterToString(*optionalParams.FilterId, ""))
	}
	if optionalParams.FilterRef != nil {
		localVarQueryParams.Add("filter[ref]", datadog.ParameterToString(*optionalParams.FilterRef, ""))
	}
	if optionalParams.FilterName != nil {
		localVarQueryParams.Add("filter[name]", datadog.ParameterToString(*optionalParams.FilterName, ""))
	}
	if optionalParams.FilterKind != nil {
		localVarQueryParams.Add("filter[kind]", datadog.ParameterToString(*optionalParams.FilterKind, ""))
	}
	if optionalParams.FilterOwner != nil {
		localVarQueryParams.Add("filter[owner]", datadog.ParameterToString(*optionalParams.FilterOwner, ""))
	}
	if optionalParams.FilterRelationType != nil {
		localVarQueryParams.Add("filter[relation][type]", datadog.ParameterToString(*optionalParams.FilterRelationType, ""))
	}
	if optionalParams.FilterExcludeSnapshot != nil {
		localVarQueryParams.Add("filter[exclude_snapshot]", datadog.ParameterToString(*optionalParams.FilterExcludeSnapshot, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// ListCatalogKindOptionalParameters holds optional parameters for ListCatalogKind.
type ListCatalogKindOptionalParameters struct {
	PageOffset *int64
	PageLimit  *int64
	FilterId   *string
	FilterName *string
}

// NewListCatalogKindOptionalParameters creates an empty struct for parameters.
func NewListCatalogKindOptionalParameters() *ListCatalogKindOptionalParameters {
	this := ListCatalogKindOptionalParameters{}
	return &this
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListCatalogKindOptionalParameters) WithPageOffset(pageOffset int64) *ListCatalogKindOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListCatalogKindOptionalParameters) WithPageLimit(pageLimit int64) *ListCatalogKindOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithFilterId sets the corresponding parameter name and returns the struct.
func (r *ListCatalogKindOptionalParameters) WithFilterId(filterId string) *ListCatalogKindOptionalParameters {
	r.FilterId = &filterId
	return r
}

// WithFilterName sets the corresponding parameter name and returns the struct.
func (r *ListCatalogKindOptionalParameters) WithFilterName(filterName string) *ListCatalogKindOptionalParameters {
	r.FilterName = &filterName
	return r
}

// ListCatalogKind Get a list of entity kinds.
// Get a list of entity kinds from Software Catalog.
func (a *SoftwareCatalogApi) ListCatalogKind(ctx _context.Context, o ...ListCatalogKindOptionalParameters) (ListKindCatalogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListKindCatalogResponse
		optionalParams      ListCatalogKindOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListCatalogKindOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SoftwareCatalogApi.ListCatalogKind")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/catalog/kind"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	if optionalParams.FilterId != nil {
		localVarQueryParams.Add("filter[id]", datadog.ParameterToString(*optionalParams.FilterId, ""))
	}
	if optionalParams.FilterName != nil {
		localVarQueryParams.Add("filter[name]", datadog.ParameterToString(*optionalParams.FilterName, ""))
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

// ListCatalogKindWithPagination provides a paginated version of ListCatalogKind returning a channel with all items.
func (a *SoftwareCatalogApi) ListCatalogKindWithPagination(ctx _context.Context, o ...ListCatalogKindOptionalParameters) (<-chan datadog.PaginationResult[KindData], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(100)
	if len(o) == 0 {
		o = append(o, ListCatalogKindOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_

	items := make(chan datadog.PaginationResult[KindData], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListCatalogKind(ctx, o...)
			if err != nil {
				var returnItem KindData
				items <- datadog.PaginationResult[KindData]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[KindData]{Item: item, Error: nil}:
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

// ListCatalogRelationOptionalParameters holds optional parameters for ListCatalogRelation.
type ListCatalogRelationOptionalParameters struct {
	PageOffset    *int64
	PageLimit     *int64
	FilterType    *RelationType
	FilterFromRef *string
	FilterToRef   *string
	Include       *RelationIncludeType
}

// NewListCatalogRelationOptionalParameters creates an empty struct for parameters.
func NewListCatalogRelationOptionalParameters() *ListCatalogRelationOptionalParameters {
	this := ListCatalogRelationOptionalParameters{}
	return &this
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListCatalogRelationOptionalParameters) WithPageOffset(pageOffset int64) *ListCatalogRelationOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListCatalogRelationOptionalParameters) WithPageLimit(pageLimit int64) *ListCatalogRelationOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithFilterType sets the corresponding parameter name and returns the struct.
func (r *ListCatalogRelationOptionalParameters) WithFilterType(filterType RelationType) *ListCatalogRelationOptionalParameters {
	r.FilterType = &filterType
	return r
}

// WithFilterFromRef sets the corresponding parameter name and returns the struct.
func (r *ListCatalogRelationOptionalParameters) WithFilterFromRef(filterFromRef string) *ListCatalogRelationOptionalParameters {
	r.FilterFromRef = &filterFromRef
	return r
}

// WithFilterToRef sets the corresponding parameter name and returns the struct.
func (r *ListCatalogRelationOptionalParameters) WithFilterToRef(filterToRef string) *ListCatalogRelationOptionalParameters {
	r.FilterToRef = &filterToRef
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListCatalogRelationOptionalParameters) WithInclude(include RelationIncludeType) *ListCatalogRelationOptionalParameters {
	r.Include = &include
	return r
}

// ListCatalogRelation Get a list of entity relations.
// Get a list of entity relations from Software Catalog.
func (a *SoftwareCatalogApi) ListCatalogRelation(ctx _context.Context, o ...ListCatalogRelationOptionalParameters) (ListRelationCatalogResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListRelationCatalogResponse
		optionalParams      ListCatalogRelationOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListCatalogRelationOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SoftwareCatalogApi.ListCatalogRelation")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/catalog/relation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	if optionalParams.FilterType != nil {
		localVarQueryParams.Add("filter[type]", datadog.ParameterToString(*optionalParams.FilterType, ""))
	}
	if optionalParams.FilterFromRef != nil {
		localVarQueryParams.Add("filter[from_ref]", datadog.ParameterToString(*optionalParams.FilterFromRef, ""))
	}
	if optionalParams.FilterToRef != nil {
		localVarQueryParams.Add("filter[to_ref]", datadog.ParameterToString(*optionalParams.FilterToRef, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
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

// ListCatalogRelationWithPagination provides a paginated version of ListCatalogRelation returning a channel with all items.
func (a *SoftwareCatalogApi) ListCatalogRelationWithPagination(ctx _context.Context, o ...ListCatalogRelationOptionalParameters) (<-chan datadog.PaginationResult[RelationResponse], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(100)
	if len(o) == 0 {
		o = append(o, ListCatalogRelationOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_

	items := make(chan datadog.PaginationResult[RelationResponse], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListCatalogRelation(ctx, o...)
			if err != nil {
				var returnItem RelationResponse
				items <- datadog.PaginationResult[RelationResponse]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[RelationResponse]{Item: item, Error: nil}:
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

// PreviewCatalogEntities Preview catalog entities.

func (a *SoftwareCatalogApi) PreviewCatalogEntities(ctx _context.Context) (EntityResponseArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue EntityResponseArray
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SoftwareCatalogApi.PreviewCatalogEntities")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/catalog/entity/preview"

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

// UpsertCatalogKind Create or update kinds.
// Create or update kinds in Software Catalog.
func (a *SoftwareCatalogApi) UpsertCatalogKind(ctx _context.Context, body UpsertCatalogKindRequest) (UpsertCatalogKindResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue UpsertCatalogKindResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SoftwareCatalogApi.UpsertCatalogKind")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/catalog/kind"

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

// NewSoftwareCatalogApi Returns NewSoftwareCatalogApi.
func NewSoftwareCatalogApi(client *datadog.APIClient) *SoftwareCatalogApi {
	return &SoftwareCatalogApi{
		Client: client,
	}
}
