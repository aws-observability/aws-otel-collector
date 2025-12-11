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

// RestrictionPoliciesApi service type
type RestrictionPoliciesApi datadog.Service

// DeleteRestrictionPolicy Delete a restriction policy.
// Deletes the restriction policy associated with a specified resource.
func (a *RestrictionPoliciesApi) DeleteRestrictionPolicy(ctx _context.Context, resourceId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RestrictionPoliciesApi.DeleteRestrictionPolicy")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/restriction_policy/{resource_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{resource_id}", _neturl.PathEscape(datadog.ParameterToString(resourceId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 429 {
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

// GetRestrictionPolicy Get a restriction policy.
// Retrieves the restriction policy associated with a specified resource.
func (a *RestrictionPoliciesApi) GetRestrictionPolicy(ctx _context.Context, resourceId string) (RestrictionPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue RestrictionPolicyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RestrictionPoliciesApi.GetRestrictionPolicy")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/restriction_policy/{resource_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{resource_id}", _neturl.PathEscape(datadog.ParameterToString(resourceId, "")))

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

// UpdateRestrictionPolicyOptionalParameters holds optional parameters for UpdateRestrictionPolicy.
type UpdateRestrictionPolicyOptionalParameters struct {
	AllowSelfLockout *bool
}

// NewUpdateRestrictionPolicyOptionalParameters creates an empty struct for parameters.
func NewUpdateRestrictionPolicyOptionalParameters() *UpdateRestrictionPolicyOptionalParameters {
	this := UpdateRestrictionPolicyOptionalParameters{}
	return &this
}

// WithAllowSelfLockout sets the corresponding parameter name and returns the struct.
func (r *UpdateRestrictionPolicyOptionalParameters) WithAllowSelfLockout(allowSelfLockout bool) *UpdateRestrictionPolicyOptionalParameters {
	r.AllowSelfLockout = &allowSelfLockout
	return r
}

// UpdateRestrictionPolicy Update a restriction policy.
// Updates the restriction policy associated with a resource.
//
// #### Supported resources
// Restriction policies can be applied to the following resources:
// - Dashboards: `dashboard`
// - Integration Services: `integration-service`
// - Integration Webhooks: `integration-webhook`
// - Notebooks: `notebook`
// - Powerpacks: `powerpack`
// - Reference Tables: `reference-table`
// - Security Rules: `security-rule`
// - Service Level Objectives: `slo`
// - Synthetic Global Variables: `synthetics-global-variable`
// - Synthetic Tests: `synthetics-test`
// - Synthetic Private Locations: `synthetics-private-location`
// - Monitors: `monitor`
// - Workflows: `workflow`
// - App Builder Apps: `app-builder-app`
// - Connections: `connection`
// - Connection Groups: `connection-group`
// - RUM Applications: `rum-application`
// - Cross Org Connections: `cross-org-connection`
// - Spreadsheets: `spreadsheet`
// - On-Call Schedules: `on-call-schedule`
// - On-Call Escalation Policies: `on-call-escalation-policy`
// - On-Call Team Routing Rules: `on-call-team-routing-rules`
// - Logs Pipelines: `logs-pipeline`
//
// #### Supported relations for resources
// Resource Type               | Supported Relations
// ----------------------------|--------------------------
// Dashboards                  | `viewer`, `editor`
// Integration Services        | `viewer`, `editor`
// Integration Webhooks        | `viewer`, `editor`
// Notebooks                   | `viewer`, `editor`
// Powerpacks                  | `viewer`, `editor`
// Security Rules              | `viewer`, `editor`
// Service Level Objectives    | `viewer`, `editor`
// Synthetic Global Variables  | `viewer`, `editor`
// Synthetic Tests             | `viewer`, `editor`
// Synthetic Private Locations | `viewer`, `editor`
// Monitors                    | `viewer`, `editor`
// Reference Tables            | `viewer`, `editor`
// Workflows                   | `viewer`, `runner`, `editor`
// App Builder Apps            | `viewer`, `editor`
// Connections                 | `viewer`, `resolver`, `editor`
// Connection Groups           | `viewer`, `editor`
// RUM Application             | `viewer`, `editor`
// Cross Org Connections       | `viewer`, `editor`
// Spreadsheets                | `viewer`, `editor`
// On-Call Schedules           | `viewer`, `overrider`, `editor`
// On-Call Escalation Policies | `viewer`, `editor`
// On-Call Team Routing Rules  | `viewer`, `editor`
// Logs Pipelines              | `viewer`, `processors_editor`, `editor`
func (a *RestrictionPoliciesApi) UpdateRestrictionPolicy(ctx _context.Context, resourceId string, body RestrictionPolicyUpdateRequest, o ...UpdateRestrictionPolicyOptionalParameters) (RestrictionPolicyResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue RestrictionPolicyResponse
		optionalParams      UpdateRestrictionPolicyOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type UpdateRestrictionPolicyOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.RestrictionPoliciesApi.UpdateRestrictionPolicy")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/restriction_policy/{resource_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{resource_id}", _neturl.PathEscape(datadog.ParameterToString(resourceId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.AllowSelfLockout != nil {
		localVarQueryParams.Add("allow_self_lockout", datadog.ParameterToString(*optionalParams.AllowSelfLockout, ""))
	}
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

// NewRestrictionPoliciesApi Returns NewRestrictionPoliciesApi.
func NewRestrictionPoliciesApi(client *datadog.APIClient) *RestrictionPoliciesApi {
	return &RestrictionPoliciesApi{
		Client: client,
	}
}
