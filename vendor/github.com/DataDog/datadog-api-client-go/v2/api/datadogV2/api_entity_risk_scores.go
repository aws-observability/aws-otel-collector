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

// EntityRiskScoresApi service type
type EntityRiskScoresApi datadog.Service

// ListEntityRiskScoresOptionalParameters holds optional parameters for ListEntityRiskScores.
type ListEntityRiskScoresOptionalParameters struct {
	From        *int64
	To          *int64
	PageSize    *int32
	PageNumber  *int32
	PageQueryId *string
	FilterSort  *string
	FilterQuery *string
	EntityType  *[]string
}

// NewListEntityRiskScoresOptionalParameters creates an empty struct for parameters.
func NewListEntityRiskScoresOptionalParameters() *ListEntityRiskScoresOptionalParameters {
	this := ListEntityRiskScoresOptionalParameters{}
	return &this
}

// WithFrom sets the corresponding parameter name and returns the struct.
func (r *ListEntityRiskScoresOptionalParameters) WithFrom(from int64) *ListEntityRiskScoresOptionalParameters {
	r.From = &from
	return r
}

// WithTo sets the corresponding parameter name and returns the struct.
func (r *ListEntityRiskScoresOptionalParameters) WithTo(to int64) *ListEntityRiskScoresOptionalParameters {
	r.To = &to
	return r
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListEntityRiskScoresOptionalParameters) WithPageSize(pageSize int32) *ListEntityRiskScoresOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListEntityRiskScoresOptionalParameters) WithPageNumber(pageNumber int32) *ListEntityRiskScoresOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithPageQueryId sets the corresponding parameter name and returns the struct.
func (r *ListEntityRiskScoresOptionalParameters) WithPageQueryId(pageQueryId string) *ListEntityRiskScoresOptionalParameters {
	r.PageQueryId = &pageQueryId
	return r
}

// WithFilterSort sets the corresponding parameter name and returns the struct.
func (r *ListEntityRiskScoresOptionalParameters) WithFilterSort(filterSort string) *ListEntityRiskScoresOptionalParameters {
	r.FilterSort = &filterSort
	return r
}

// WithFilterQuery sets the corresponding parameter name and returns the struct.
func (r *ListEntityRiskScoresOptionalParameters) WithFilterQuery(filterQuery string) *ListEntityRiskScoresOptionalParameters {
	r.FilterQuery = &filterQuery
	return r
}

// WithEntityType sets the corresponding parameter name and returns the struct.
func (r *ListEntityRiskScoresOptionalParameters) WithEntityType(entityType []string) *ListEntityRiskScoresOptionalParameters {
	r.EntityType = &entityType
	return r
}

// ListEntityRiskScores List Entity Risk Scores.
// Get a list of entity risk scores for your organization. Entity risk scores provide security risk assessment for entities like cloud resources, identities, or services based on detected signals, misconfigurations, and identity risks.
func (a *EntityRiskScoresApi) ListEntityRiskScores(ctx _context.Context, o ...ListEntityRiskScoresOptionalParameters) (SecurityEntityRiskScoresResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityEntityRiskScoresResponse
		optionalParams      ListEntityRiskScoresOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListEntityRiskScoresOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListEntityRiskScores"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.EntityRiskScoresApi.ListEntityRiskScores")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security-entities/risk-scores"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.From != nil {
		localVarQueryParams.Add("from", datadog.ParameterToString(*optionalParams.From, ""))
	}
	if optionalParams.To != nil {
		localVarQueryParams.Add("to", datadog.ParameterToString(*optionalParams.To, ""))
	}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.PageQueryId != nil {
		localVarQueryParams.Add("page[queryId]", datadog.ParameterToString(*optionalParams.PageQueryId, ""))
	}
	if optionalParams.FilterSort != nil {
		localVarQueryParams.Add("filter[sort]", datadog.ParameterToString(*optionalParams.FilterSort, ""))
	}
	if optionalParams.FilterQuery != nil {
		localVarQueryParams.Add("filter[query]", datadog.ParameterToString(*optionalParams.FilterQuery, ""))
	}
	if optionalParams.EntityType != nil {
		t := *optionalParams.EntityType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("entityType", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("entityType", datadog.ParameterToString(t, "multi"))
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

// NewEntityRiskScoresApi Returns NewEntityRiskScoresApi.
func NewEntityRiskScoresApi(client *datadog.APIClient) *EntityRiskScoresApi {
	return &EntityRiskScoresApi{
		Client: client,
	}
}
