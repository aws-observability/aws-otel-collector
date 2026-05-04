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

// SpaApi service type
type SpaApi datadog.Service

// GetSPARecommendationsOptionalParameters holds optional parameters for GetSPARecommendations.
type GetSPARecommendationsOptionalParameters struct {
	BypassCache *string
}

// NewGetSPARecommendationsOptionalParameters creates an empty struct for parameters.
func NewGetSPARecommendationsOptionalParameters() *GetSPARecommendationsOptionalParameters {
	this := GetSPARecommendationsOptionalParameters{}
	return &this
}

// WithBypassCache sets the corresponding parameter name and returns the struct.
func (r *GetSPARecommendationsOptionalParameters) WithBypassCache(bypassCache string) *GetSPARecommendationsOptionalParameters {
	r.BypassCache = &bypassCache
	return r
}

// GetSPARecommendations Get SPA Recommendations.
// This endpoint is currently experimental and restricted to Datadog internal use only. Retrieve resource recommendations for a Spark job. The caller (Spark Gateway or DJM UI) provides a service name and SPA returns structured recommendations for driver and executor resources. The version with a shard should be preferred, where possible, as it gives more accurate results.
func (a *SpaApi) GetSPARecommendations(ctx _context.Context, service string, o ...GetSPARecommendationsOptionalParameters) (RecommendationDocument, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue RecommendationDocument
		optionalParams      GetSPARecommendationsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetSPARecommendationsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetSPARecommendations"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SpaApi.GetSPARecommendations")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/spa/recommendations/{service}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{service}", _neturl.PathEscape(datadog.ParameterToString(service, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.BypassCache != nil {
		localVarQueryParams.Add("bypass_cache", datadog.ParameterToString(*optionalParams.BypassCache, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
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

// GetSPARecommendationsWithShardOptionalParameters holds optional parameters for GetSPARecommendationsWithShard.
type GetSPARecommendationsWithShardOptionalParameters struct {
	BypassCache *string
}

// NewGetSPARecommendationsWithShardOptionalParameters creates an empty struct for parameters.
func NewGetSPARecommendationsWithShardOptionalParameters() *GetSPARecommendationsWithShardOptionalParameters {
	this := GetSPARecommendationsWithShardOptionalParameters{}
	return &this
}

// WithBypassCache sets the corresponding parameter name and returns the struct.
func (r *GetSPARecommendationsWithShardOptionalParameters) WithBypassCache(bypassCache string) *GetSPARecommendationsWithShardOptionalParameters {
	r.BypassCache = &bypassCache
	return r
}

// GetSPARecommendationsWithShard Get SPA Recommendations with a shard parameter.
// This endpoint is currently experimental and restricted to Datadog internal use only. Retrieve resource recommendations for a Spark job. The caller (Spark Gateway or DJM UI) provides a service name and shard identifier, and SPA returns structured recommendations for driver and executor resources.
func (a *SpaApi) GetSPARecommendationsWithShard(ctx _context.Context, shard string, service string, o ...GetSPARecommendationsWithShardOptionalParameters) (RecommendationDocument, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue RecommendationDocument
		optionalParams      GetSPARecommendationsWithShardOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetSPARecommendationsWithShardOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetSPARecommendationsWithShard"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SpaApi.GetSPARecommendationsWithShard")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/spa/recommendations/{service}/{shard}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{shard}", _neturl.PathEscape(datadog.ParameterToString(shard, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{service}", _neturl.PathEscape(datadog.ParameterToString(service, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.BypassCache != nil {
		localVarQueryParams.Add("bypass_cache", datadog.ParameterToString(*optionalParams.BypassCache, ""))
	}
	localVarHeaderParams["Accept"] = "application/json"

	datadog.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
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

// NewSpaApi Returns NewSpaApi.
func NewSpaApi(client *datadog.APIClient) *SpaApi {
	return &SpaApi{
		Client: client,
	}
}
