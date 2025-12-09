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

// CloudNetworkMonitoringApi service type
type CloudNetworkMonitoringApi datadog.Service

// GetAggregatedConnectionsOptionalParameters holds optional parameters for GetAggregatedConnections.
type GetAggregatedConnectionsOptionalParameters struct {
	From    *int64
	To      *int64
	GroupBy *string
	Tags    *string
	Limit   *int32
}

// NewGetAggregatedConnectionsOptionalParameters creates an empty struct for parameters.
func NewGetAggregatedConnectionsOptionalParameters() *GetAggregatedConnectionsOptionalParameters {
	this := GetAggregatedConnectionsOptionalParameters{}
	return &this
}

// WithFrom sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedConnectionsOptionalParameters) WithFrom(from int64) *GetAggregatedConnectionsOptionalParameters {
	r.From = &from
	return r
}

// WithTo sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedConnectionsOptionalParameters) WithTo(to int64) *GetAggregatedConnectionsOptionalParameters {
	r.To = &to
	return r
}

// WithGroupBy sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedConnectionsOptionalParameters) WithGroupBy(groupBy string) *GetAggregatedConnectionsOptionalParameters {
	r.GroupBy = &groupBy
	return r
}

// WithTags sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedConnectionsOptionalParameters) WithTags(tags string) *GetAggregatedConnectionsOptionalParameters {
	r.Tags = &tags
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedConnectionsOptionalParameters) WithLimit(limit int32) *GetAggregatedConnectionsOptionalParameters {
	r.Limit = &limit
	return r
}

// GetAggregatedConnections Get all aggregated connections.
// Get all aggregated connections.
func (a *CloudNetworkMonitoringApi) GetAggregatedConnections(ctx _context.Context, o ...GetAggregatedConnectionsOptionalParameters) (SingleAggregatedConnectionResponseArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SingleAggregatedConnectionResponseArray
		optionalParams      GetAggregatedConnectionsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetAggregatedConnectionsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.CloudNetworkMonitoringApi.GetAggregatedConnections")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/network/connections/aggregate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.From != nil {
		localVarQueryParams.Add("from", datadog.ParameterToString(*optionalParams.From, ""))
	}
	if optionalParams.To != nil {
		localVarQueryParams.Add("to", datadog.ParameterToString(*optionalParams.To, ""))
	}
	if optionalParams.GroupBy != nil {
		localVarQueryParams.Add("group_by", datadog.ParameterToString(*optionalParams.GroupBy, ""))
	}
	if optionalParams.Tags != nil {
		localVarQueryParams.Add("tags", datadog.ParameterToString(*optionalParams.Tags, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", datadog.ParameterToString(*optionalParams.Limit, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 429 {
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

// GetAggregatedDnsOptionalParameters holds optional parameters for GetAggregatedDns.
type GetAggregatedDnsOptionalParameters struct {
	From    *int64
	To      *int64
	GroupBy *string
	Tags    *string
	Limit   *int32
}

// NewGetAggregatedDnsOptionalParameters creates an empty struct for parameters.
func NewGetAggregatedDnsOptionalParameters() *GetAggregatedDnsOptionalParameters {
	this := GetAggregatedDnsOptionalParameters{}
	return &this
}

// WithFrom sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedDnsOptionalParameters) WithFrom(from int64) *GetAggregatedDnsOptionalParameters {
	r.From = &from
	return r
}

// WithTo sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedDnsOptionalParameters) WithTo(to int64) *GetAggregatedDnsOptionalParameters {
	r.To = &to
	return r
}

// WithGroupBy sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedDnsOptionalParameters) WithGroupBy(groupBy string) *GetAggregatedDnsOptionalParameters {
	r.GroupBy = &groupBy
	return r
}

// WithTags sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedDnsOptionalParameters) WithTags(tags string) *GetAggregatedDnsOptionalParameters {
	r.Tags = &tags
	return r
}

// WithLimit sets the corresponding parameter name and returns the struct.
func (r *GetAggregatedDnsOptionalParameters) WithLimit(limit int32) *GetAggregatedDnsOptionalParameters {
	r.Limit = &limit
	return r
}

// GetAggregatedDns Get all aggregated DNS traffic.
// Get all aggregated DNS traffic.
func (a *CloudNetworkMonitoringApi) GetAggregatedDns(ctx _context.Context, o ...GetAggregatedDnsOptionalParameters) (SingleAggregatedDnsResponseArray, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SingleAggregatedDnsResponseArray
		optionalParams      GetAggregatedDnsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetAggregatedDnsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.CloudNetworkMonitoringApi.GetAggregatedDns")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/network/dns/aggregate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.From != nil {
		localVarQueryParams.Add("from", datadog.ParameterToString(*optionalParams.From, ""))
	}
	if optionalParams.To != nil {
		localVarQueryParams.Add("to", datadog.ParameterToString(*optionalParams.To, ""))
	}
	if optionalParams.GroupBy != nil {
		localVarQueryParams.Add("group_by", datadog.ParameterToString(*optionalParams.GroupBy, ""))
	}
	if optionalParams.Tags != nil {
		localVarQueryParams.Add("tags", datadog.ParameterToString(*optionalParams.Tags, ""))
	}
	if optionalParams.Limit != nil {
		localVarQueryParams.Add("limit", datadog.ParameterToString(*optionalParams.Limit, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 429 {
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

// NewCloudNetworkMonitoringApi Returns NewCloudNetworkMonitoringApi.
func NewCloudNetworkMonitoringApi(client *datadog.APIClient) *CloudNetworkMonitoringApi {
	return &CloudNetworkMonitoringApi{
		Client: client,
	}
}
