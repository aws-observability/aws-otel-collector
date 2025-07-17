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
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringApi service type
type SecurityMonitoringApi datadog.Service

// CancelHistoricalJob Cancel a historical job.
// Cancel a historical job.
func (a *SecurityMonitoringApi) CancelHistoricalJob(ctx _context.Context, jobId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPatch
		localVarPostBody   interface{}
	)

	operationId := "v2.CancelHistoricalJob"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.CancelHistoricalJob")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/siem-historical-detections/jobs/{job_id}/cancel"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{job_id}", _neturl.PathEscape(datadog.ParameterToString(jobId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 429 {
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

// ConvertExistingSecurityMonitoringRule Convert an existing rule from JSON to Terraform.
// Convert an existing rule from JSON to Terraform for datadog provider
// resource datadog_security_monitoring_rule.
func (a *SecurityMonitoringApi) ConvertExistingSecurityMonitoringRule(ctx _context.Context, ruleId string) (SecurityMonitoringRuleConvertResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringRuleConvertResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ConvertExistingSecurityMonitoringRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/{rule_id}/convert"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{rule_id}", _neturl.PathEscape(datadog.ParameterToString(ruleId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// ConvertJobResultToSignal Convert a job result to a signal.
// Convert a job result to a signal.
func (a *SecurityMonitoringApi) ConvertJobResultToSignal(ctx _context.Context, body ConvertJobResultsToSignalsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	operationId := "v2.ConvertJobResultToSignal"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ConvertJobResultToSignal")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/siem-historical-detections/jobs/signal_convert"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "*/*"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// ConvertSecurityMonitoringRuleFromJSONToTerraform Convert a rule from JSON to Terraform.
// Convert a rule that doesn't (yet) exist from JSON to Terraform for datadog provider
// resource datadog_security_monitoring_rule.
func (a *SecurityMonitoringApi) ConvertSecurityMonitoringRuleFromJSONToTerraform(ctx _context.Context, body SecurityMonitoringRuleConvertPayload) (SecurityMonitoringRuleConvertResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringRuleConvertResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ConvertSecurityMonitoringRuleFromJSONToTerraform")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/convert"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// CreateCustomFramework Create a custom framework.
// Create a custom framework.
func (a *SecurityMonitoringApi) CreateCustomFramework(ctx _context.Context, body CreateCustomFrameworkRequest) (CreateCustomFrameworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue CreateCustomFrameworkResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.CreateCustomFramework")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/cloud_security_management/custom_frameworks"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 429 || localVarHTTPResponse.StatusCode == 500 {
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

// CreateSecurityFilter Create a security filter.
// Create a security filter.
//
// See the [security filter guide](https://docs.datadoghq.com/security_platform/guide/how-to-setup-security-filters-using-security-monitoring-api/)
// for more examples.
func (a *SecurityMonitoringApi) CreateSecurityFilter(ctx _context.Context, body SecurityFilterCreateRequest) (SecurityFilterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue SecurityFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.CreateSecurityFilter")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/security_filters"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 429 {
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

// CreateSecurityMonitoringRule Create a detection rule.
// Create a detection rule.
func (a *SecurityMonitoringApi) CreateSecurityMonitoringRule(ctx _context.Context, body SecurityMonitoringRuleCreatePayload) (SecurityMonitoringRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.CreateSecurityMonitoringRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules"

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

// CreateSecurityMonitoringSuppression Create a suppression rule.
// Create a new suppression rule.
func (a *SecurityMonitoringApi) CreateSecurityMonitoringSuppression(ctx _context.Context, body SecurityMonitoringSuppressionCreateRequest) (SecurityMonitoringSuppressionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSuppressionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.CreateSecurityMonitoringSuppression")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/suppressions"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 429 {
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

// CreateSignalNotificationRule Create a new signal-based notification rule.
// Create a new notification rule for security signals and return the created rule.
func (a *SecurityMonitoringApi) CreateSignalNotificationRule(ctx _context.Context, body CreateNotificationRuleParameters) (NotificationRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue NotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.CreateSignalNotificationRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/signals/notification_rules"

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

// CreateVulnerabilityNotificationRule Create a new vulnerability-based notification rule.
// Create a new notification rule for security vulnerabilities and return the created rule.
func (a *SecurityMonitoringApi) CreateVulnerabilityNotificationRule(ctx _context.Context, body CreateNotificationRuleParameters) (NotificationRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue NotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.CreateVulnerabilityNotificationRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/vulnerabilities/notification_rules"

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

// DeleteCustomFramework Delete a custom framework.
// Delete a custom framework.
func (a *SecurityMonitoringApi) DeleteCustomFramework(ctx _context.Context, handle string, version string) (DeleteCustomFrameworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		localVarReturnValue DeleteCustomFrameworkResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.DeleteCustomFramework")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/cloud_security_management/custom_frameworks/{handle}/{version}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{handle}", _neturl.PathEscape(datadog.ParameterToString(handle, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{version}", _neturl.PathEscape(datadog.ParameterToString(version, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 429 || localVarHTTPResponse.StatusCode == 500 {
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

// DeleteHistoricalJob Delete an existing job.
// Delete an existing job.
func (a *SecurityMonitoringApi) DeleteHistoricalJob(ctx _context.Context, jobId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	operationId := "v2.DeleteHistoricalJob"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.DeleteHistoricalJob")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/siem-historical-detections/jobs/{job_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{job_id}", _neturl.PathEscape(datadog.ParameterToString(jobId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 429 {
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

// DeleteSecurityFilter Delete a security filter.
// Delete a specific security filter.
func (a *SecurityMonitoringApi) DeleteSecurityFilter(ctx _context.Context, securityFilterId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.DeleteSecurityFilter")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/security_filters/{security_filter_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{security_filter_id}", _neturl.PathEscape(datadog.ParameterToString(securityFilterId, "")))

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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// DeleteSecurityMonitoringRule Delete an existing rule.
// Delete an existing rule. Default rules cannot be deleted.
func (a *SecurityMonitoringApi) DeleteSecurityMonitoringRule(ctx _context.Context, ruleId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.DeleteSecurityMonitoringRule")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/{rule_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{rule_id}", _neturl.PathEscape(datadog.ParameterToString(ruleId, "")))

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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// DeleteSecurityMonitoringSuppression Delete a suppression rule.
// Delete a specific suppression rule.
func (a *SecurityMonitoringApi) DeleteSecurityMonitoringSuppression(ctx _context.Context, suppressionId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.DeleteSecurityMonitoringSuppression")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/suppressions/{suppression_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{suppression_id}", _neturl.PathEscape(datadog.ParameterToString(suppressionId, "")))

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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// DeleteSignalNotificationRule Delete a signal-based notification rule.
// Delete a notification rule for security signals.
func (a *SecurityMonitoringApi) DeleteSignalNotificationRule(ctx _context.Context, id string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.DeleteSignalNotificationRule")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/signals/notification_rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// DeleteVulnerabilityNotificationRule Delete a vulnerability-based notification rule.
// Delete a notification rule for security vulnerabilities.
func (a *SecurityMonitoringApi) DeleteVulnerabilityNotificationRule(ctx _context.Context, id string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.DeleteVulnerabilityNotificationRule")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/vulnerabilities/notification_rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// EditSecurityMonitoringSignalAssignee Modify the triage assignee of a security signal.
// Modify the triage assignee of a security signal.
func (a *SecurityMonitoringApi) EditSecurityMonitoringSignalAssignee(ctx _context.Context, signalId string, body SecurityMonitoringSignalAssigneeUpdateRequest) (SecurityMonitoringSignalTriageUpdateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSignalTriageUpdateResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.EditSecurityMonitoringSignalAssignee")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/signals/{signal_id}/assignee"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{signal_id}", _neturl.PathEscape(datadog.ParameterToString(signalId, "")))

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

// EditSecurityMonitoringSignalIncidents Change the related incidents of a security signal.
// Change the related incidents for a security signal.
func (a *SecurityMonitoringApi) EditSecurityMonitoringSignalIncidents(ctx _context.Context, signalId string, body SecurityMonitoringSignalIncidentsUpdateRequest) (SecurityMonitoringSignalTriageUpdateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSignalTriageUpdateResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/signals/{signal_id}/incidents"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{signal_id}", _neturl.PathEscape(datadog.ParameterToString(signalId, "")))

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

// EditSecurityMonitoringSignalState Change the triage state of a security signal.
// Change the triage state of a security signal.
func (a *SecurityMonitoringApi) EditSecurityMonitoringSignalState(ctx _context.Context, signalId string, body SecurityMonitoringSignalStateUpdateRequest) (SecurityMonitoringSignalTriageUpdateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSignalTriageUpdateResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.EditSecurityMonitoringSignalState")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/signals/{signal_id}/state"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{signal_id}", _neturl.PathEscape(datadog.ParameterToString(signalId, "")))

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

// GetCustomFramework Get a custom framework.
// Get a custom framework.
func (a *SecurityMonitoringApi) GetCustomFramework(ctx _context.Context, handle string, version string) (GetCustomFrameworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue GetCustomFrameworkResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetCustomFramework")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/cloud_security_management/custom_frameworks/{handle}/{version}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{handle}", _neturl.PathEscape(datadog.ParameterToString(handle, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{version}", _neturl.PathEscape(datadog.ParameterToString(version, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 429 || localVarHTTPResponse.StatusCode == 500 {
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

// GetFindingOptionalParameters holds optional parameters for GetFinding.
type GetFindingOptionalParameters struct {
	SnapshotTimestamp *int64
}

// NewGetFindingOptionalParameters creates an empty struct for parameters.
func NewGetFindingOptionalParameters() *GetFindingOptionalParameters {
	this := GetFindingOptionalParameters{}
	return &this
}

// WithSnapshotTimestamp sets the corresponding parameter name and returns the struct.
func (r *GetFindingOptionalParameters) WithSnapshotTimestamp(snapshotTimestamp int64) *GetFindingOptionalParameters {
	r.SnapshotTimestamp = &snapshotTimestamp
	return r
}

// GetFinding Get a finding.
// Returns a single finding with message and resource configuration.
func (a *SecurityMonitoringApi) GetFinding(ctx _context.Context, findingId string, o ...GetFindingOptionalParameters) (GetFindingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue GetFindingResponse
		optionalParams      GetFindingOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetFindingOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetFinding"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetFinding")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/posture_management/findings/{finding_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{finding_id}", _neturl.PathEscape(datadog.ParameterToString(findingId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.SnapshotTimestamp != nil {
		localVarQueryParams.Add("snapshot_timestamp", datadog.ParameterToString(*optionalParams.SnapshotTimestamp, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v JSONAPIErrorResponse
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

// GetHistoricalJob Get a job's details.
// Get a job's details.
func (a *SecurityMonitoringApi) GetHistoricalJob(ctx _context.Context, jobId string) (HistoricalJobResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue HistoricalJobResponse
	)

	operationId := "v2.GetHistoricalJob"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetHistoricalJob")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/siem-historical-detections/jobs/{job_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{job_id}", _neturl.PathEscape(datadog.ParameterToString(jobId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// GetResourceEvaluationFiltersOptionalParameters holds optional parameters for GetResourceEvaluationFilters.
type GetResourceEvaluationFiltersOptionalParameters struct {
	CloudProvider *string
	AccountId     *string
	SkipCache     *bool
}

// NewGetResourceEvaluationFiltersOptionalParameters creates an empty struct for parameters.
func NewGetResourceEvaluationFiltersOptionalParameters() *GetResourceEvaluationFiltersOptionalParameters {
	this := GetResourceEvaluationFiltersOptionalParameters{}
	return &this
}

// WithCloudProvider sets the corresponding parameter name and returns the struct.
func (r *GetResourceEvaluationFiltersOptionalParameters) WithCloudProvider(cloudProvider string) *GetResourceEvaluationFiltersOptionalParameters {
	r.CloudProvider = &cloudProvider
	return r
}

// WithAccountId sets the corresponding parameter name and returns the struct.
func (r *GetResourceEvaluationFiltersOptionalParameters) WithAccountId(accountId string) *GetResourceEvaluationFiltersOptionalParameters {
	r.AccountId = &accountId
	return r
}

// WithSkipCache sets the corresponding parameter name and returns the struct.
func (r *GetResourceEvaluationFiltersOptionalParameters) WithSkipCache(skipCache bool) *GetResourceEvaluationFiltersOptionalParameters {
	r.SkipCache = &skipCache
	return r
}

// GetResourceEvaluationFilters List resource filters.
// List resource filters.
func (a *SecurityMonitoringApi) GetResourceEvaluationFilters(ctx _context.Context, o ...GetResourceEvaluationFiltersOptionalParameters) (GetResourceEvaluationFiltersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue GetResourceEvaluationFiltersResponse
		optionalParams      GetResourceEvaluationFiltersOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetResourceEvaluationFiltersOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetResourceEvaluationFilters")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/cloud_security_management/resource_filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.CloudProvider != nil {
		localVarQueryParams.Add("cloud_provider", datadog.ParameterToString(*optionalParams.CloudProvider, ""))
	}
	if optionalParams.AccountId != nil {
		localVarQueryParams.Add("account_id", datadog.ParameterToString(*optionalParams.AccountId, ""))
	}
	if optionalParams.SkipCache != nil {
		localVarQueryParams.Add("skip_cache", datadog.ParameterToString(*optionalParams.SkipCache, ""))
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

// GetRuleVersionHistoryOptionalParameters holds optional parameters for GetRuleVersionHistory.
type GetRuleVersionHistoryOptionalParameters struct {
	PageSize   *int64
	PageNumber *int64
}

// NewGetRuleVersionHistoryOptionalParameters creates an empty struct for parameters.
func NewGetRuleVersionHistoryOptionalParameters() *GetRuleVersionHistoryOptionalParameters {
	this := GetRuleVersionHistoryOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *GetRuleVersionHistoryOptionalParameters) WithPageSize(pageSize int64) *GetRuleVersionHistoryOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *GetRuleVersionHistoryOptionalParameters) WithPageNumber(pageNumber int64) *GetRuleVersionHistoryOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// GetRuleVersionHistory Get a rule's version history.
// Get a rule's version history.
func (a *SecurityMonitoringApi) GetRuleVersionHistory(ctx _context.Context, ruleId string, o ...GetRuleVersionHistoryOptionalParameters) (GetRuleVersionHistoryResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue GetRuleVersionHistoryResponse
		optionalParams      GetRuleVersionHistoryOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetRuleVersionHistoryOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetRuleVersionHistory"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetRuleVersionHistory")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/{rule_id}/version_history"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{rule_id}", _neturl.PathEscape(datadog.ParameterToString(ruleId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
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

// GetSBOMOptionalParameters holds optional parameters for GetSBOM.
type GetSBOMOptionalParameters struct {
	FilterRepoDigest *string
}

// NewGetSBOMOptionalParameters creates an empty struct for parameters.
func NewGetSBOMOptionalParameters() *GetSBOMOptionalParameters {
	this := GetSBOMOptionalParameters{}
	return &this
}

// WithFilterRepoDigest sets the corresponding parameter name and returns the struct.
func (r *GetSBOMOptionalParameters) WithFilterRepoDigest(filterRepoDigest string) *GetSBOMOptionalParameters {
	r.FilterRepoDigest = &filterRepoDigest
	return r
}

// GetSBOM Get SBOM.
// Get a single SBOM related to an asset by its type and name.
func (a *SecurityMonitoringApi) GetSBOM(ctx _context.Context, assetType AssetType, filterAssetName string, o ...GetSBOMOptionalParameters) (GetSBOMResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue GetSBOMResponse
		optionalParams      GetSBOMOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetSBOMOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetSBOM"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetSBOM")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/sboms/{asset_type}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{asset_type}", _neturl.PathEscape(datadog.ParameterToString(assetType, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("filter[asset_name]", datadog.ParameterToString(filterAssetName, ""))
	if optionalParams.FilterRepoDigest != nil {
		localVarQueryParams.Add("filter[repo_digest]", datadog.ParameterToString(*optionalParams.FilterRepoDigest, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// GetSecurityFilter Get a security filter.
// Get the details of a specific security filter.
//
// See the [security filter guide](https://docs.datadoghq.com/security_platform/guide/how-to-setup-security-filters-using-security-monitoring-api/)
// for more examples.
func (a *SecurityMonitoringApi) GetSecurityFilter(ctx _context.Context, securityFilterId string) (SecurityFilterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetSecurityFilter")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/security_filters/{security_filter_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{security_filter_id}", _neturl.PathEscape(datadog.ParameterToString(securityFilterId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// GetSecurityMonitoringRule Get a rule's details.
// Get a rule's details.
func (a *SecurityMonitoringApi) GetSecurityMonitoringRule(ctx _context.Context, ruleId string) (SecurityMonitoringRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetSecurityMonitoringRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/{rule_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{rule_id}", _neturl.PathEscape(datadog.ParameterToString(ruleId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
		if localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// GetSecurityMonitoringSignal Get a signal's details.
// Get a signal's details.
func (a *SecurityMonitoringApi) GetSecurityMonitoringSignal(ctx _context.Context, signalId string) (SecurityMonitoringSignalResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSignalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetSecurityMonitoringSignal")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/signals/{signal_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{signal_id}", _neturl.PathEscape(datadog.ParameterToString(signalId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
		if localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// GetSecurityMonitoringSuppression Get a suppression rule.
// Get the details of a specific suppression rule.
func (a *SecurityMonitoringApi) GetSecurityMonitoringSuppression(ctx _context.Context, suppressionId string) (SecurityMonitoringSuppressionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSuppressionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetSecurityMonitoringSuppression")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/suppressions/{suppression_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{suppression_id}", _neturl.PathEscape(datadog.ParameterToString(suppressionId, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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
		if localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// GetSignalNotificationRule Get details of a signal-based notification rule.
// Get the details of a notification rule for security signals.
func (a *SecurityMonitoringApi) GetSignalNotificationRule(ctx _context.Context, id string) (NotificationRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue NotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetSignalNotificationRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/signals/notification_rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// GetSignalNotificationRules Get the list of signal-based notification rules.
// Returns the list of notification rules for security signals.
func (a *SecurityMonitoringApi) GetSignalNotificationRules(ctx _context.Context) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetSignalNotificationRules")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/signals/notification_rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// GetVulnerabilityNotificationRule Get details of a vulnerability notification rule.
// Get the details of a notification rule for security vulnerabilities.
func (a *SecurityMonitoringApi) GetVulnerabilityNotificationRule(ctx _context.Context, id string) (NotificationRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue NotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetVulnerabilityNotificationRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/vulnerabilities/notification_rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// GetVulnerabilityNotificationRules Get the list of vulnerability notification rules.
// Returns the list of notification rules for security vulnerabilities.
func (a *SecurityMonitoringApi) GetVulnerabilityNotificationRules(ctx _context.Context) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.GetVulnerabilityNotificationRules")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/vulnerabilities/notification_rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// ListFindingsOptionalParameters holds optional parameters for ListFindings.
type ListFindingsOptionalParameters struct {
	PageLimit                 *int64
	SnapshotTimestamp         *int64
	PageCursor                *string
	FilterTags                *string
	FilterEvaluationChangedAt *string
	FilterMuted               *bool
	FilterRuleId              *string
	FilterRuleName            *string
	FilterResourceType        *string
	FilterDiscoveryTimestamp  *string
	FilterEvaluation          *FindingEvaluation
	FilterStatus              *FindingStatus
	FilterVulnerabilityType   *[]FindingVulnerabilityType
	DetailedFindings          *bool
}

// NewListFindingsOptionalParameters creates an empty struct for parameters.
func NewListFindingsOptionalParameters() *ListFindingsOptionalParameters {
	this := ListFindingsOptionalParameters{}
	return &this
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithPageLimit(pageLimit int64) *ListFindingsOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// WithSnapshotTimestamp sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithSnapshotTimestamp(snapshotTimestamp int64) *ListFindingsOptionalParameters {
	r.SnapshotTimestamp = &snapshotTimestamp
	return r
}

// WithPageCursor sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithPageCursor(pageCursor string) *ListFindingsOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

// WithFilterTags sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterTags(filterTags string) *ListFindingsOptionalParameters {
	r.FilterTags = &filterTags
	return r
}

// WithFilterEvaluationChangedAt sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterEvaluationChangedAt(filterEvaluationChangedAt string) *ListFindingsOptionalParameters {
	r.FilterEvaluationChangedAt = &filterEvaluationChangedAt
	return r
}

// WithFilterMuted sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterMuted(filterMuted bool) *ListFindingsOptionalParameters {
	r.FilterMuted = &filterMuted
	return r
}

// WithFilterRuleId sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterRuleId(filterRuleId string) *ListFindingsOptionalParameters {
	r.FilterRuleId = &filterRuleId
	return r
}

// WithFilterRuleName sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterRuleName(filterRuleName string) *ListFindingsOptionalParameters {
	r.FilterRuleName = &filterRuleName
	return r
}

// WithFilterResourceType sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterResourceType(filterResourceType string) *ListFindingsOptionalParameters {
	r.FilterResourceType = &filterResourceType
	return r
}

// WithFilterDiscoveryTimestamp sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterDiscoveryTimestamp(filterDiscoveryTimestamp string) *ListFindingsOptionalParameters {
	r.FilterDiscoveryTimestamp = &filterDiscoveryTimestamp
	return r
}

// WithFilterEvaluation sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterEvaluation(filterEvaluation FindingEvaluation) *ListFindingsOptionalParameters {
	r.FilterEvaluation = &filterEvaluation
	return r
}

// WithFilterStatus sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterStatus(filterStatus FindingStatus) *ListFindingsOptionalParameters {
	r.FilterStatus = &filterStatus
	return r
}

// WithFilterVulnerabilityType sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithFilterVulnerabilityType(filterVulnerabilityType []FindingVulnerabilityType) *ListFindingsOptionalParameters {
	r.FilterVulnerabilityType = &filterVulnerabilityType
	return r
}

// WithDetailedFindings sets the corresponding parameter name and returns the struct.
func (r *ListFindingsOptionalParameters) WithDetailedFindings(detailedFindings bool) *ListFindingsOptionalParameters {
	r.DetailedFindings = &detailedFindings
	return r
}

// ListFindings List findings.
// Get a list of findings. These include both misconfigurations and identity risks.
//
// **Note**: To filter and return only identity risks, add the following query parameter: `?filter[tags]=dd_rule_type:ciem`
//
// ### Filtering
//
// Filters can be applied by appending query parameters to the URL.
//
//   - Using a single filter: `?filter[attribute_key]=attribute_value`
//   - Chaining filters: `?filter[attribute_key]=attribute_value&filter[attribute_key]=attribute_value...`
//   - Filtering on tags: `?filter[tags]=tag_key:tag_value&filter[tags]=tag_key_2:tag_value_2`
//
// Here, `attribute_key` can be any of the filter keys described further below.
//
// Query parameters of type `integer` support comparison operators (`>`, `>=`, `<`, `<=`). This is particularly useful when filtering by `evaluation_changed_at` or `resource_discovery_timestamp`. For example: `?filter[evaluation_changed_at]=>20123123121`.
//
// You can also use the negation operator on strings. For example, use `filter[resource_type]=-aws*` to filter for any non-AWS resources.
//
// The operator must come after the equal sign. For example, to filter with the `>=` operator, add the operator after the equal sign: `filter[evaluation_changed_at]=>=1678809373257`.
//
// Query parameters must be only among the documented ones and with values of correct types. Duplicated query parameters (e.g. `filter[status]=low&filter[status]=info`) are not allowed.
//
// ### Additional extension fields
//
// Additional extension fields are available for some findings.
//
// The data is available when you include the query parameter `?detailed_findings=true` in the request.
//
// The following fields are available for findings:
// - `external_id`: The resource external ID related to the finding.
// - `description`: The description and remediation steps for the finding.
// - `datadog_link`: The Datadog relative link for the finding.
//
// ### Response
//
// The response includes an array of finding objects, pagination metadata, and a count of items that match the query.
//
// Each finding object contains the following:
//
// - The finding ID that can be used in a `GetFinding` request to retrieve the full finding details.
// - Core attributes, including status, evaluation, high-level resource details, muted state, and rule details.
// - `evaluation_changed_at` and `resource_discovery_date` time stamps.
// - An array of associated tags.
func (a *SecurityMonitoringApi) ListFindings(ctx _context.Context, o ...ListFindingsOptionalParameters) (ListFindingsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListFindingsResponse
		optionalParams      ListFindingsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListFindingsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListFindings"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ListFindings")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/posture_management/findings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
	}
	if optionalParams.SnapshotTimestamp != nil {
		localVarQueryParams.Add("snapshot_timestamp", datadog.ParameterToString(*optionalParams.SnapshotTimestamp, ""))
	}
	if optionalParams.PageCursor != nil {
		localVarQueryParams.Add("page[cursor]", datadog.ParameterToString(*optionalParams.PageCursor, ""))
	}
	if optionalParams.FilterTags != nil {
		localVarQueryParams.Add("filter[tags]", datadog.ParameterToString(*optionalParams.FilterTags, ""))
	}
	if optionalParams.FilterEvaluationChangedAt != nil {
		localVarQueryParams.Add("filter[evaluation_changed_at]", datadog.ParameterToString(*optionalParams.FilterEvaluationChangedAt, ""))
	}
	if optionalParams.FilterMuted != nil {
		localVarQueryParams.Add("filter[muted]", datadog.ParameterToString(*optionalParams.FilterMuted, ""))
	}
	if optionalParams.FilterRuleId != nil {
		localVarQueryParams.Add("filter[rule_id]", datadog.ParameterToString(*optionalParams.FilterRuleId, ""))
	}
	if optionalParams.FilterRuleName != nil {
		localVarQueryParams.Add("filter[rule_name]", datadog.ParameterToString(*optionalParams.FilterRuleName, ""))
	}
	if optionalParams.FilterResourceType != nil {
		localVarQueryParams.Add("filter[resource_type]", datadog.ParameterToString(*optionalParams.FilterResourceType, ""))
	}
	if optionalParams.FilterDiscoveryTimestamp != nil {
		localVarQueryParams.Add("filter[discovery_timestamp]", datadog.ParameterToString(*optionalParams.FilterDiscoveryTimestamp, ""))
	}
	if optionalParams.FilterEvaluation != nil {
		localVarQueryParams.Add("filter[evaluation]", datadog.ParameterToString(*optionalParams.FilterEvaluation, ""))
	}
	if optionalParams.FilterStatus != nil {
		localVarQueryParams.Add("filter[status]", datadog.ParameterToString(*optionalParams.FilterStatus, ""))
	}
	if optionalParams.FilterVulnerabilityType != nil {
		t := *optionalParams.FilterVulnerabilityType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[vulnerability_type]", datadog.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[vulnerability_type]", datadog.ParameterToString(t, "multi"))
		}
	}
	if optionalParams.DetailedFindings != nil {
		localVarQueryParams.Add("detailed_findings", datadog.ParameterToString(*optionalParams.DetailedFindings, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v JSONAPIErrorResponse
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

// ListFindingsWithPagination provides a paginated version of ListFindings returning a channel with all items.
func (a *SecurityMonitoringApi) ListFindingsWithPagination(ctx _context.Context, o ...ListFindingsOptionalParameters) (<-chan datadog.PaginationResult[Finding], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(100)
	if len(o) == 0 {
		o = append(o, ListFindingsOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_

	items := make(chan datadog.PaginationResult[Finding], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListFindings(ctx, o...)
			if err != nil {
				var returnItem Finding
				items <- datadog.PaginationResult[Finding]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[Finding]{Item: item, Error: nil}:
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
			cursorMetaPage, ok := cursorMeta.GetPageOk()
			if !ok {
				break
			}
			cursorMetaPageCursor, ok := cursorMetaPage.GetCursorOk()
			if !ok {
				break
			}

			o[0].PageCursor = cursorMetaPageCursor
		}
		close(items)
	}()
	return items, cancel
}

// ListHistoricalJobsOptionalParameters holds optional parameters for ListHistoricalJobs.
type ListHistoricalJobsOptionalParameters struct {
	PageSize    *int64
	PageNumber  *int64
	Sort        *string
	FilterQuery *string
}

// NewListHistoricalJobsOptionalParameters creates an empty struct for parameters.
func NewListHistoricalJobsOptionalParameters() *ListHistoricalJobsOptionalParameters {
	this := ListHistoricalJobsOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListHistoricalJobsOptionalParameters) WithPageSize(pageSize int64) *ListHistoricalJobsOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListHistoricalJobsOptionalParameters) WithPageNumber(pageNumber int64) *ListHistoricalJobsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListHistoricalJobsOptionalParameters) WithSort(sort string) *ListHistoricalJobsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithFilterQuery sets the corresponding parameter name and returns the struct.
func (r *ListHistoricalJobsOptionalParameters) WithFilterQuery(filterQuery string) *ListHistoricalJobsOptionalParameters {
	r.FilterQuery = &filterQuery
	return r
}

// ListHistoricalJobs List historical jobs.
// List historical jobs.
func (a *SecurityMonitoringApi) ListHistoricalJobs(ctx _context.Context, o ...ListHistoricalJobsOptionalParameters) (ListHistoricalJobsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListHistoricalJobsResponse
		optionalParams      ListHistoricalJobsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListHistoricalJobsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListHistoricalJobs"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ListHistoricalJobs")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/siem-historical-detections/jobs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.FilterQuery != nil {
		localVarQueryParams.Add("filter[query]", datadog.ParameterToString(*optionalParams.FilterQuery, ""))
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

// ListSecurityFilters Get all security filters.
// Get the list of configured security filters with their definitions.
func (a *SecurityMonitoringApi) ListSecurityFilters(ctx _context.Context) (SecurityFiltersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityFiltersResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ListSecurityFilters")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/security_filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// ListSecurityMonitoringRulesOptionalParameters holds optional parameters for ListSecurityMonitoringRules.
type ListSecurityMonitoringRulesOptionalParameters struct {
	PageSize   *int64
	PageNumber *int64
}

// NewListSecurityMonitoringRulesOptionalParameters creates an empty struct for parameters.
func NewListSecurityMonitoringRulesOptionalParameters() *ListSecurityMonitoringRulesOptionalParameters {
	this := ListSecurityMonitoringRulesOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListSecurityMonitoringRulesOptionalParameters) WithPageSize(pageSize int64) *ListSecurityMonitoringRulesOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListSecurityMonitoringRulesOptionalParameters) WithPageNumber(pageNumber int64) *ListSecurityMonitoringRulesOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// ListSecurityMonitoringRules List rules.
// List rules.
func (a *SecurityMonitoringApi) ListSecurityMonitoringRules(ctx _context.Context, o ...ListSecurityMonitoringRulesOptionalParameters) (SecurityMonitoringListRulesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringListRulesResponse
		optionalParams      ListSecurityMonitoringRulesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListSecurityMonitoringRulesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ListSecurityMonitoringRules")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
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

// ListSecurityMonitoringSignalsOptionalParameters holds optional parameters for ListSecurityMonitoringSignals.
type ListSecurityMonitoringSignalsOptionalParameters struct {
	FilterQuery *string
	FilterFrom  *time.Time
	FilterTo    *time.Time
	Sort        *SecurityMonitoringSignalsSort
	PageCursor  *string
	PageLimit   *int32
}

// NewListSecurityMonitoringSignalsOptionalParameters creates an empty struct for parameters.
func NewListSecurityMonitoringSignalsOptionalParameters() *ListSecurityMonitoringSignalsOptionalParameters {
	this := ListSecurityMonitoringSignalsOptionalParameters{}
	return &this
}

// WithFilterQuery sets the corresponding parameter name and returns the struct.
func (r *ListSecurityMonitoringSignalsOptionalParameters) WithFilterQuery(filterQuery string) *ListSecurityMonitoringSignalsOptionalParameters {
	r.FilterQuery = &filterQuery
	return r
}

// WithFilterFrom sets the corresponding parameter name and returns the struct.
func (r *ListSecurityMonitoringSignalsOptionalParameters) WithFilterFrom(filterFrom time.Time) *ListSecurityMonitoringSignalsOptionalParameters {
	r.FilterFrom = &filterFrom
	return r
}

// WithFilterTo sets the corresponding parameter name and returns the struct.
func (r *ListSecurityMonitoringSignalsOptionalParameters) WithFilterTo(filterTo time.Time) *ListSecurityMonitoringSignalsOptionalParameters {
	r.FilterTo = &filterTo
	return r
}

// WithSort sets the corresponding parameter name and returns the struct.
func (r *ListSecurityMonitoringSignalsOptionalParameters) WithSort(sort SecurityMonitoringSignalsSort) *ListSecurityMonitoringSignalsOptionalParameters {
	r.Sort = &sort
	return r
}

// WithPageCursor sets the corresponding parameter name and returns the struct.
func (r *ListSecurityMonitoringSignalsOptionalParameters) WithPageCursor(pageCursor string) *ListSecurityMonitoringSignalsOptionalParameters {
	r.PageCursor = &pageCursor
	return r
}

// WithPageLimit sets the corresponding parameter name and returns the struct.
func (r *ListSecurityMonitoringSignalsOptionalParameters) WithPageLimit(pageLimit int32) *ListSecurityMonitoringSignalsOptionalParameters {
	r.PageLimit = &pageLimit
	return r
}

// ListSecurityMonitoringSignals Get a quick list of security signals.
// The list endpoint returns security signals that match a search query.
// Both this endpoint and the POST endpoint can be used interchangeably when listing
// security signals.
func (a *SecurityMonitoringApi) ListSecurityMonitoringSignals(ctx _context.Context, o ...ListSecurityMonitoringSignalsOptionalParameters) (SecurityMonitoringSignalsListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSignalsListResponse
		optionalParams      ListSecurityMonitoringSignalsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListSecurityMonitoringSignalsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ListSecurityMonitoringSignals")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/signals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.FilterQuery != nil {
		localVarQueryParams.Add("filter[query]", datadog.ParameterToString(*optionalParams.FilterQuery, ""))
	}
	if optionalParams.FilterFrom != nil {
		localVarQueryParams.Add("filter[from]", datadog.ParameterToString(*optionalParams.FilterFrom, ""))
	}
	if optionalParams.FilterTo != nil {
		localVarQueryParams.Add("filter[to]", datadog.ParameterToString(*optionalParams.FilterTo, ""))
	}
	if optionalParams.Sort != nil {
		localVarQueryParams.Add("sort", datadog.ParameterToString(*optionalParams.Sort, ""))
	}
	if optionalParams.PageCursor != nil {
		localVarQueryParams.Add("page[cursor]", datadog.ParameterToString(*optionalParams.PageCursor, ""))
	}
	if optionalParams.PageLimit != nil {
		localVarQueryParams.Add("page[limit]", datadog.ParameterToString(*optionalParams.PageLimit, ""))
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

// ListSecurityMonitoringSignalsWithPagination provides a paginated version of ListSecurityMonitoringSignals returning a channel with all items.
func (a *SecurityMonitoringApi) ListSecurityMonitoringSignalsWithPagination(ctx _context.Context, o ...ListSecurityMonitoringSignalsOptionalParameters) (<-chan datadog.PaginationResult[SecurityMonitoringSignal], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(10)
	if len(o) == 0 {
		o = append(o, ListSecurityMonitoringSignalsOptionalParameters{})
	}
	if o[0].PageLimit != nil {
		pageSize_ = *o[0].PageLimit
	}
	o[0].PageLimit = &pageSize_

	items := make(chan datadog.PaginationResult[SecurityMonitoringSignal], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListSecurityMonitoringSignals(ctx, o...)
			if err != nil {
				var returnItem SecurityMonitoringSignal
				items <- datadog.PaginationResult[SecurityMonitoringSignal]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[SecurityMonitoringSignal]{Item: item, Error: nil}:
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
			cursorMetaPage, ok := cursorMeta.GetPageOk()
			if !ok {
				break
			}
			cursorMetaPageAfter, ok := cursorMetaPage.GetAfterOk()
			if !ok {
				break
			}

			o[0].PageCursor = cursorMetaPageAfter
		}
		close(items)
	}()
	return items, cancel
}

// ListSecurityMonitoringSuppressions Get all suppression rules.
// Get the list of all suppression rules.
func (a *SecurityMonitoringApi) ListSecurityMonitoringSuppressions(ctx _context.Context) (SecurityMonitoringSuppressionsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSuppressionsResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ListSecurityMonitoringSuppressions")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/suppressions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
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

// ListVulnerabilitiesOptionalParameters holds optional parameters for ListVulnerabilities.
type ListVulnerabilitiesOptionalParameters struct {
	PageToken                                *string
	PageNumber                               *int64
	FilterType                               *VulnerabilityType
	FilterCvssBaseScoreOp                    *float64
	FilterCvssBaseSeverity                   *VulnerabilitySeverity
	FilterCvssBaseVector                     *string
	FilterCvssDatadogScoreOp                 *float64
	FilterCvssDatadogSeverity                *VulnerabilitySeverity
	FilterCvssDatadogVector                  *string
	FilterStatus                             *VulnerabilityStatus
	FilterTool                               *VulnerabilityTool
	FilterLibraryName                        *string
	FilterLibraryVersion                     *string
	FilterAdvisoryId                         *string
	FilterRisksExploitationProbability       *bool
	FilterRisksPocExploitAvailable           *bool
	FilterRisksExploitAvailable              *bool
	FilterRisksEpssScoreOp                   *float64
	FilterRisksEpssSeverity                  *VulnerabilitySeverity
	FilterLanguage                           *string
	FilterEcosystem                          *VulnerabilityEcosystem
	FilterCodeLocationLocation               *string
	FilterCodeLocationFilePath               *string
	FilterCodeLocationMethod                 *string
	FilterFixAvailable                       *bool
	FilterRepoDigests                        *string
	FilterAssetName                          *string
	FilterAssetType                          *AssetType
	FilterAssetVersionFirst                  *string
	FilterAssetVersionLast                   *string
	FilterAssetRepositoryUrl                 *string
	FilterAssetRisksInProduction             *bool
	FilterAssetRisksUnderAttack              *bool
	FilterAssetRisksIsPubliclyAccessible     *bool
	FilterAssetRisksHasPrivilegedAccess      *bool
	FilterAssetRisksHasAccessToSensitiveData *bool
	FilterAssetEnvironments                  *string
	FilterAssetArch                          *string
	FilterAssetOperatingSystemName           *string
	FilterAssetOperatingSystemVersion        *string
}

// NewListVulnerabilitiesOptionalParameters creates an empty struct for parameters.
func NewListVulnerabilitiesOptionalParameters() *ListVulnerabilitiesOptionalParameters {
	this := ListVulnerabilitiesOptionalParameters{}
	return &this
}

// WithPageToken sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithPageToken(pageToken string) *ListVulnerabilitiesOptionalParameters {
	r.PageToken = &pageToken
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithPageNumber(pageNumber int64) *ListVulnerabilitiesOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithFilterType sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterType(filterType VulnerabilityType) *ListVulnerabilitiesOptionalParameters {
	r.FilterType = &filterType
	return r
}

// WithFilterCvssBaseScoreOp sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCvssBaseScoreOp(filterCvssBaseScoreOp float64) *ListVulnerabilitiesOptionalParameters {
	r.FilterCvssBaseScoreOp = &filterCvssBaseScoreOp
	return r
}

// WithFilterCvssBaseSeverity sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCvssBaseSeverity(filterCvssBaseSeverity VulnerabilitySeverity) *ListVulnerabilitiesOptionalParameters {
	r.FilterCvssBaseSeverity = &filterCvssBaseSeverity
	return r
}

// WithFilterCvssBaseVector sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCvssBaseVector(filterCvssBaseVector string) *ListVulnerabilitiesOptionalParameters {
	r.FilterCvssBaseVector = &filterCvssBaseVector
	return r
}

// WithFilterCvssDatadogScoreOp sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCvssDatadogScoreOp(filterCvssDatadogScoreOp float64) *ListVulnerabilitiesOptionalParameters {
	r.FilterCvssDatadogScoreOp = &filterCvssDatadogScoreOp
	return r
}

// WithFilterCvssDatadogSeverity sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCvssDatadogSeverity(filterCvssDatadogSeverity VulnerabilitySeverity) *ListVulnerabilitiesOptionalParameters {
	r.FilterCvssDatadogSeverity = &filterCvssDatadogSeverity
	return r
}

// WithFilterCvssDatadogVector sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCvssDatadogVector(filterCvssDatadogVector string) *ListVulnerabilitiesOptionalParameters {
	r.FilterCvssDatadogVector = &filterCvssDatadogVector
	return r
}

// WithFilterStatus sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterStatus(filterStatus VulnerabilityStatus) *ListVulnerabilitiesOptionalParameters {
	r.FilterStatus = &filterStatus
	return r
}

// WithFilterTool sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterTool(filterTool VulnerabilityTool) *ListVulnerabilitiesOptionalParameters {
	r.FilterTool = &filterTool
	return r
}

// WithFilterLibraryName sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterLibraryName(filterLibraryName string) *ListVulnerabilitiesOptionalParameters {
	r.FilterLibraryName = &filterLibraryName
	return r
}

// WithFilterLibraryVersion sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterLibraryVersion(filterLibraryVersion string) *ListVulnerabilitiesOptionalParameters {
	r.FilterLibraryVersion = &filterLibraryVersion
	return r
}

// WithFilterAdvisoryId sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAdvisoryId(filterAdvisoryId string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAdvisoryId = &filterAdvisoryId
	return r
}

// WithFilterRisksExploitationProbability sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterRisksExploitationProbability(filterRisksExploitationProbability bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterRisksExploitationProbability = &filterRisksExploitationProbability
	return r
}

// WithFilterRisksPocExploitAvailable sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterRisksPocExploitAvailable(filterRisksPocExploitAvailable bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterRisksPocExploitAvailable = &filterRisksPocExploitAvailable
	return r
}

// WithFilterRisksExploitAvailable sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterRisksExploitAvailable(filterRisksExploitAvailable bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterRisksExploitAvailable = &filterRisksExploitAvailable
	return r
}

// WithFilterRisksEpssScoreOp sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterRisksEpssScoreOp(filterRisksEpssScoreOp float64) *ListVulnerabilitiesOptionalParameters {
	r.FilterRisksEpssScoreOp = &filterRisksEpssScoreOp
	return r
}

// WithFilterRisksEpssSeverity sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterRisksEpssSeverity(filterRisksEpssSeverity VulnerabilitySeverity) *ListVulnerabilitiesOptionalParameters {
	r.FilterRisksEpssSeverity = &filterRisksEpssSeverity
	return r
}

// WithFilterLanguage sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterLanguage(filterLanguage string) *ListVulnerabilitiesOptionalParameters {
	r.FilterLanguage = &filterLanguage
	return r
}

// WithFilterEcosystem sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterEcosystem(filterEcosystem VulnerabilityEcosystem) *ListVulnerabilitiesOptionalParameters {
	r.FilterEcosystem = &filterEcosystem
	return r
}

// WithFilterCodeLocationLocation sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCodeLocationLocation(filterCodeLocationLocation string) *ListVulnerabilitiesOptionalParameters {
	r.FilterCodeLocationLocation = &filterCodeLocationLocation
	return r
}

// WithFilterCodeLocationFilePath sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCodeLocationFilePath(filterCodeLocationFilePath string) *ListVulnerabilitiesOptionalParameters {
	r.FilterCodeLocationFilePath = &filterCodeLocationFilePath
	return r
}

// WithFilterCodeLocationMethod sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterCodeLocationMethod(filterCodeLocationMethod string) *ListVulnerabilitiesOptionalParameters {
	r.FilterCodeLocationMethod = &filterCodeLocationMethod
	return r
}

// WithFilterFixAvailable sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterFixAvailable(filterFixAvailable bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterFixAvailable = &filterFixAvailable
	return r
}

// WithFilterRepoDigests sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterRepoDigests(filterRepoDigests string) *ListVulnerabilitiesOptionalParameters {
	r.FilterRepoDigests = &filterRepoDigests
	return r
}

// WithFilterAssetName sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetName(filterAssetName string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetName = &filterAssetName
	return r
}

// WithFilterAssetType sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetType(filterAssetType AssetType) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetType = &filterAssetType
	return r
}

// WithFilterAssetVersionFirst sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetVersionFirst(filterAssetVersionFirst string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetVersionFirst = &filterAssetVersionFirst
	return r
}

// WithFilterAssetVersionLast sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetVersionLast(filterAssetVersionLast string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetVersionLast = &filterAssetVersionLast
	return r
}

// WithFilterAssetRepositoryUrl sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetRepositoryUrl(filterAssetRepositoryUrl string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetRepositoryUrl = &filterAssetRepositoryUrl
	return r
}

// WithFilterAssetRisksInProduction sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetRisksInProduction(filterAssetRisksInProduction bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetRisksInProduction = &filterAssetRisksInProduction
	return r
}

// WithFilterAssetRisksUnderAttack sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetRisksUnderAttack(filterAssetRisksUnderAttack bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetRisksUnderAttack = &filterAssetRisksUnderAttack
	return r
}

// WithFilterAssetRisksIsPubliclyAccessible sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetRisksIsPubliclyAccessible(filterAssetRisksIsPubliclyAccessible bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetRisksIsPubliclyAccessible = &filterAssetRisksIsPubliclyAccessible
	return r
}

// WithFilterAssetRisksHasPrivilegedAccess sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetRisksHasPrivilegedAccess(filterAssetRisksHasPrivilegedAccess bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetRisksHasPrivilegedAccess = &filterAssetRisksHasPrivilegedAccess
	return r
}

// WithFilterAssetRisksHasAccessToSensitiveData sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetRisksHasAccessToSensitiveData(filterAssetRisksHasAccessToSensitiveData bool) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetRisksHasAccessToSensitiveData = &filterAssetRisksHasAccessToSensitiveData
	return r
}

// WithFilterAssetEnvironments sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetEnvironments(filterAssetEnvironments string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetEnvironments = &filterAssetEnvironments
	return r
}

// WithFilterAssetArch sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetArch(filterAssetArch string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetArch = &filterAssetArch
	return r
}

// WithFilterAssetOperatingSystemName sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetOperatingSystemName(filterAssetOperatingSystemName string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetOperatingSystemName = &filterAssetOperatingSystemName
	return r
}

// WithFilterAssetOperatingSystemVersion sets the corresponding parameter name and returns the struct.
func (r *ListVulnerabilitiesOptionalParameters) WithFilterAssetOperatingSystemVersion(filterAssetOperatingSystemVersion string) *ListVulnerabilitiesOptionalParameters {
	r.FilterAssetOperatingSystemVersion = &filterAssetOperatingSystemVersion
	return r
}

// ListVulnerabilities List vulnerabilities.
// Get a list of vulnerabilities.
//
// ### Pagination
//
// Pagination is enabled by default in both `vulnerabilities` and `assets`. The size of the page varies depending on the endpoint and cannot be modified. To automate the request of the next page, you can use the links section in the response.
//
// This endpoint will return paginated responses. The pages are stored in the links section of the response:
//
// ```JSON
//
//	{
//	  "data": [...],
//	  "meta": {...},
//	  "links": {
//	    "self": "https://.../api/v2/security/vulnerabilities",
//	    "first": "https://.../api/v2/security/vulnerabilities?page[number]=1&page[token]=abc",
//	    "last": "https://.../api/v2/security/vulnerabilities?page[number]=43&page[token]=abc",
//	    "next": "https://.../api/v2/security/vulnerabilities?page[number]=2&page[token]=abc"
//	  }
//	}
//
// ```
//
// - `links.previous` is empty if the first page is requested.
// - `links.next` is empty if the last page is requested.
//
// #### Token
//
// Vulnerabilities can be created, updated or deleted at any point in time.
//
// Upon the first request, a token is created to ensure consistency across subsequent paginated requests.
//
// A token is valid only for 24 hours.
//
// #### First request
//
// We consider a request to be the first request when there is no `page[token]` parameter.
//
// The response of this first request contains the newly created token in the `links` section.
//
// This token can then be used in the subsequent paginated requests.
//
// #### Subsequent requests
//
// Any request containing valid `page[token]` and `page[number]` parameters will be considered a subsequent request.
//
// If the `token` is invalid, a `404` response will be returned.
//
// If the page `number` is invalid, a `400` response will be returned.
//
// ### Filtering
//
// The request can include some filter parameters to filter the data to be retrieved. The format of the filter parameters follows the [JSON:API format](https://jsonapi.org/format/#fetching-filtering): `filter[$prop_name]`, where `prop_name` is the property name in the entity being filtered by.
//
// All filters can include multiple values, where data will be filtered with an OR clause: `filter[title]=Title1,Title2` will filter all vulnerabilities where title is equal to `Title1` OR `Title2`.
//
// String filters are case sensitive.
//
// Boolean filters accept `true` or `false` as values.
//
// Number filters must include an operator as a second filter input: `filter[$prop_name][$operator]`. For example, for the vulnerabilities endpoint: `filter[cvss.base.score][lte]=8`.
//
// Available operators are: `eq` (==), `lt` (<), `lte` (<=), `gt` (>) and `gte` (>=).
//
// ### Metadata
//
// Following [JSON:API format](https://jsonapi.org/format/#document-meta), object including non-standard meta-information.
//
// This endpoint includes the meta member in the response. For more details on each of the properties included in this section, check the endpoints response tables.
//
// ```JSON
//
//	{
//	  "data": [...],
//	  "meta": {
//	    "total": 1500,
//	    "count": 18732,
//	    "token": "some_token"
//	  },
//	  "links": {...}
//	}
//
// ```
func (a *SecurityMonitoringApi) ListVulnerabilities(ctx _context.Context, o ...ListVulnerabilitiesOptionalParameters) (ListVulnerabilitiesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListVulnerabilitiesResponse
		optionalParams      ListVulnerabilitiesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListVulnerabilitiesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListVulnerabilities"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ListVulnerabilities")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/vulnerabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageToken != nil {
		localVarQueryParams.Add("page[token]", datadog.ParameterToString(*optionalParams.PageToken, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.FilterType != nil {
		localVarQueryParams.Add("filter[type]", datadog.ParameterToString(*optionalParams.FilterType, ""))
	}
	if optionalParams.FilterCvssBaseScoreOp != nil {
		localVarQueryParams.Add("filter[cvss.base.score][`$op`]", datadog.ParameterToString(*optionalParams.FilterCvssBaseScoreOp, ""))
	}
	if optionalParams.FilterCvssBaseSeverity != nil {
		localVarQueryParams.Add("filter[cvss.base.severity]", datadog.ParameterToString(*optionalParams.FilterCvssBaseSeverity, ""))
	}
	if optionalParams.FilterCvssBaseVector != nil {
		localVarQueryParams.Add("filter[cvss.base.vector]", datadog.ParameterToString(*optionalParams.FilterCvssBaseVector, ""))
	}
	if optionalParams.FilterCvssDatadogScoreOp != nil {
		localVarQueryParams.Add("filter[cvss.datadog.score][`$op`]", datadog.ParameterToString(*optionalParams.FilterCvssDatadogScoreOp, ""))
	}
	if optionalParams.FilterCvssDatadogSeverity != nil {
		localVarQueryParams.Add("filter[cvss.datadog.severity]", datadog.ParameterToString(*optionalParams.FilterCvssDatadogSeverity, ""))
	}
	if optionalParams.FilterCvssDatadogVector != nil {
		localVarQueryParams.Add("filter[cvss.datadog.vector]", datadog.ParameterToString(*optionalParams.FilterCvssDatadogVector, ""))
	}
	if optionalParams.FilterStatus != nil {
		localVarQueryParams.Add("filter[status]", datadog.ParameterToString(*optionalParams.FilterStatus, ""))
	}
	if optionalParams.FilterTool != nil {
		localVarQueryParams.Add("filter[tool]", datadog.ParameterToString(*optionalParams.FilterTool, ""))
	}
	if optionalParams.FilterLibraryName != nil {
		localVarQueryParams.Add("filter[library.name]", datadog.ParameterToString(*optionalParams.FilterLibraryName, ""))
	}
	if optionalParams.FilterLibraryVersion != nil {
		localVarQueryParams.Add("filter[library.version]", datadog.ParameterToString(*optionalParams.FilterLibraryVersion, ""))
	}
	if optionalParams.FilterAdvisoryId != nil {
		localVarQueryParams.Add("filter[advisory_id]", datadog.ParameterToString(*optionalParams.FilterAdvisoryId, ""))
	}
	if optionalParams.FilterRisksExploitationProbability != nil {
		localVarQueryParams.Add("filter[risks.exploitation_probability]", datadog.ParameterToString(*optionalParams.FilterRisksExploitationProbability, ""))
	}
	if optionalParams.FilterRisksPocExploitAvailable != nil {
		localVarQueryParams.Add("filter[risks.poc_exploit_available]", datadog.ParameterToString(*optionalParams.FilterRisksPocExploitAvailable, ""))
	}
	if optionalParams.FilterRisksExploitAvailable != nil {
		localVarQueryParams.Add("filter[risks.exploit_available]", datadog.ParameterToString(*optionalParams.FilterRisksExploitAvailable, ""))
	}
	if optionalParams.FilterRisksEpssScoreOp != nil {
		localVarQueryParams.Add("filter[risks.epss.score][`$op`]", datadog.ParameterToString(*optionalParams.FilterRisksEpssScoreOp, ""))
	}
	if optionalParams.FilterRisksEpssSeverity != nil {
		localVarQueryParams.Add("filter[risks.epss.severity]", datadog.ParameterToString(*optionalParams.FilterRisksEpssSeverity, ""))
	}
	if optionalParams.FilterLanguage != nil {
		localVarQueryParams.Add("filter[language]", datadog.ParameterToString(*optionalParams.FilterLanguage, ""))
	}
	if optionalParams.FilterEcosystem != nil {
		localVarQueryParams.Add("filter[ecosystem]", datadog.ParameterToString(*optionalParams.FilterEcosystem, ""))
	}
	if optionalParams.FilterCodeLocationLocation != nil {
		localVarQueryParams.Add("filter[code_location.location]", datadog.ParameterToString(*optionalParams.FilterCodeLocationLocation, ""))
	}
	if optionalParams.FilterCodeLocationFilePath != nil {
		localVarQueryParams.Add("filter[code_location.file_path]", datadog.ParameterToString(*optionalParams.FilterCodeLocationFilePath, ""))
	}
	if optionalParams.FilterCodeLocationMethod != nil {
		localVarQueryParams.Add("filter[code_location.method]", datadog.ParameterToString(*optionalParams.FilterCodeLocationMethod, ""))
	}
	if optionalParams.FilterFixAvailable != nil {
		localVarQueryParams.Add("filter[fix_available]", datadog.ParameterToString(*optionalParams.FilterFixAvailable, ""))
	}
	if optionalParams.FilterRepoDigests != nil {
		localVarQueryParams.Add("filter[repo_digests]", datadog.ParameterToString(*optionalParams.FilterRepoDigests, ""))
	}
	if optionalParams.FilterAssetName != nil {
		localVarQueryParams.Add("filter[asset.name]", datadog.ParameterToString(*optionalParams.FilterAssetName, ""))
	}
	if optionalParams.FilterAssetType != nil {
		localVarQueryParams.Add("filter[asset.type]", datadog.ParameterToString(*optionalParams.FilterAssetType, ""))
	}
	if optionalParams.FilterAssetVersionFirst != nil {
		localVarQueryParams.Add("filter[asset.version.first]", datadog.ParameterToString(*optionalParams.FilterAssetVersionFirst, ""))
	}
	if optionalParams.FilterAssetVersionLast != nil {
		localVarQueryParams.Add("filter[asset.version.last]", datadog.ParameterToString(*optionalParams.FilterAssetVersionLast, ""))
	}
	if optionalParams.FilterAssetRepositoryUrl != nil {
		localVarQueryParams.Add("filter[asset.repository_url]", datadog.ParameterToString(*optionalParams.FilterAssetRepositoryUrl, ""))
	}
	if optionalParams.FilterAssetRisksInProduction != nil {
		localVarQueryParams.Add("filter[asset.risks.in_production]", datadog.ParameterToString(*optionalParams.FilterAssetRisksInProduction, ""))
	}
	if optionalParams.FilterAssetRisksUnderAttack != nil {
		localVarQueryParams.Add("filter[asset.risks.under_attack]", datadog.ParameterToString(*optionalParams.FilterAssetRisksUnderAttack, ""))
	}
	if optionalParams.FilterAssetRisksIsPubliclyAccessible != nil {
		localVarQueryParams.Add("filter[asset.risks.is_publicly_accessible]", datadog.ParameterToString(*optionalParams.FilterAssetRisksIsPubliclyAccessible, ""))
	}
	if optionalParams.FilterAssetRisksHasPrivilegedAccess != nil {
		localVarQueryParams.Add("filter[asset.risks.has_privileged_access]", datadog.ParameterToString(*optionalParams.FilterAssetRisksHasPrivilegedAccess, ""))
	}
	if optionalParams.FilterAssetRisksHasAccessToSensitiveData != nil {
		localVarQueryParams.Add("filter[asset.risks.has_access_to_sensitive_data]", datadog.ParameterToString(*optionalParams.FilterAssetRisksHasAccessToSensitiveData, ""))
	}
	if optionalParams.FilterAssetEnvironments != nil {
		localVarQueryParams.Add("filter[asset.environments]", datadog.ParameterToString(*optionalParams.FilterAssetEnvironments, ""))
	}
	if optionalParams.FilterAssetArch != nil {
		localVarQueryParams.Add("filter[asset.arch]", datadog.ParameterToString(*optionalParams.FilterAssetArch, ""))
	}
	if optionalParams.FilterAssetOperatingSystemName != nil {
		localVarQueryParams.Add("filter[asset.operating_system.name]", datadog.ParameterToString(*optionalParams.FilterAssetOperatingSystemName, ""))
	}
	if optionalParams.FilterAssetOperatingSystemVersion != nil {
		localVarQueryParams.Add("filter[asset.operating_system.version]", datadog.ParameterToString(*optionalParams.FilterAssetOperatingSystemVersion, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// ListVulnerableAssetsOptionalParameters holds optional parameters for ListVulnerableAssets.
type ListVulnerableAssetsOptionalParameters struct {
	PageToken                           *string
	PageNumber                          *int64
	FilterName                          *string
	FilterType                          *AssetType
	FilterVersionFirst                  *string
	FilterVersionLast                   *string
	FilterRepositoryUrl                 *string
	FilterRisksInProduction             *bool
	FilterRisksUnderAttack              *bool
	FilterRisksIsPubliclyAccessible     *bool
	FilterRisksHasPrivilegedAccess      *bool
	FilterRisksHasAccessToSensitiveData *bool
	FilterEnvironments                  *string
	FilterArch                          *string
	FilterOperatingSystemName           *string
	FilterOperatingSystemVersion        *string
}

// NewListVulnerableAssetsOptionalParameters creates an empty struct for parameters.
func NewListVulnerableAssetsOptionalParameters() *ListVulnerableAssetsOptionalParameters {
	this := ListVulnerableAssetsOptionalParameters{}
	return &this
}

// WithPageToken sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithPageToken(pageToken string) *ListVulnerableAssetsOptionalParameters {
	r.PageToken = &pageToken
	return r
}

// WithPageNumber sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithPageNumber(pageNumber int64) *ListVulnerableAssetsOptionalParameters {
	r.PageNumber = &pageNumber
	return r
}

// WithFilterName sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterName(filterName string) *ListVulnerableAssetsOptionalParameters {
	r.FilterName = &filterName
	return r
}

// WithFilterType sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterType(filterType AssetType) *ListVulnerableAssetsOptionalParameters {
	r.FilterType = &filterType
	return r
}

// WithFilterVersionFirst sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterVersionFirst(filterVersionFirst string) *ListVulnerableAssetsOptionalParameters {
	r.FilterVersionFirst = &filterVersionFirst
	return r
}

// WithFilterVersionLast sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterVersionLast(filterVersionLast string) *ListVulnerableAssetsOptionalParameters {
	r.FilterVersionLast = &filterVersionLast
	return r
}

// WithFilterRepositoryUrl sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterRepositoryUrl(filterRepositoryUrl string) *ListVulnerableAssetsOptionalParameters {
	r.FilterRepositoryUrl = &filterRepositoryUrl
	return r
}

// WithFilterRisksInProduction sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterRisksInProduction(filterRisksInProduction bool) *ListVulnerableAssetsOptionalParameters {
	r.FilterRisksInProduction = &filterRisksInProduction
	return r
}

// WithFilterRisksUnderAttack sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterRisksUnderAttack(filterRisksUnderAttack bool) *ListVulnerableAssetsOptionalParameters {
	r.FilterRisksUnderAttack = &filterRisksUnderAttack
	return r
}

// WithFilterRisksIsPubliclyAccessible sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterRisksIsPubliclyAccessible(filterRisksIsPubliclyAccessible bool) *ListVulnerableAssetsOptionalParameters {
	r.FilterRisksIsPubliclyAccessible = &filterRisksIsPubliclyAccessible
	return r
}

// WithFilterRisksHasPrivilegedAccess sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterRisksHasPrivilegedAccess(filterRisksHasPrivilegedAccess bool) *ListVulnerableAssetsOptionalParameters {
	r.FilterRisksHasPrivilegedAccess = &filterRisksHasPrivilegedAccess
	return r
}

// WithFilterRisksHasAccessToSensitiveData sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterRisksHasAccessToSensitiveData(filterRisksHasAccessToSensitiveData bool) *ListVulnerableAssetsOptionalParameters {
	r.FilterRisksHasAccessToSensitiveData = &filterRisksHasAccessToSensitiveData
	return r
}

// WithFilterEnvironments sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterEnvironments(filterEnvironments string) *ListVulnerableAssetsOptionalParameters {
	r.FilterEnvironments = &filterEnvironments
	return r
}

// WithFilterArch sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterArch(filterArch string) *ListVulnerableAssetsOptionalParameters {
	r.FilterArch = &filterArch
	return r
}

// WithFilterOperatingSystemName sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterOperatingSystemName(filterOperatingSystemName string) *ListVulnerableAssetsOptionalParameters {
	r.FilterOperatingSystemName = &filterOperatingSystemName
	return r
}

// WithFilterOperatingSystemVersion sets the corresponding parameter name and returns the struct.
func (r *ListVulnerableAssetsOptionalParameters) WithFilterOperatingSystemVersion(filterOperatingSystemVersion string) *ListVulnerableAssetsOptionalParameters {
	r.FilterOperatingSystemVersion = &filterOperatingSystemVersion
	return r
}

// ListVulnerableAssets List vulnerable assets.
// Get a list of vulnerable assets.
//
// ### Pagination
//
// Please review the [Pagination section for the "List Vulnerabilities"](#pagination) endpoint.
//
// ### Filtering
//
// Please review the [Filtering section for the "List Vulnerabilities"](#filtering) endpoint.
//
// ### Metadata
//
// Please review the [Metadata section for the "List Vulnerabilities"](#metadata) endpoint.
func (a *SecurityMonitoringApi) ListVulnerableAssets(ctx _context.Context, o ...ListVulnerableAssetsOptionalParameters) (ListVulnerableAssetsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListVulnerableAssetsResponse
		optionalParams      ListVulnerableAssetsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListVulnerableAssetsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListVulnerableAssets"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ListVulnerableAssets")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/assets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageToken != nil {
		localVarQueryParams.Add("page[token]", datadog.ParameterToString(*optionalParams.PageToken, ""))
	}
	if optionalParams.PageNumber != nil {
		localVarQueryParams.Add("page[number]", datadog.ParameterToString(*optionalParams.PageNumber, ""))
	}
	if optionalParams.FilterName != nil {
		localVarQueryParams.Add("filter[name]", datadog.ParameterToString(*optionalParams.FilterName, ""))
	}
	if optionalParams.FilterType != nil {
		localVarQueryParams.Add("filter[type]", datadog.ParameterToString(*optionalParams.FilterType, ""))
	}
	if optionalParams.FilterVersionFirst != nil {
		localVarQueryParams.Add("filter[version.first]", datadog.ParameterToString(*optionalParams.FilterVersionFirst, ""))
	}
	if optionalParams.FilterVersionLast != nil {
		localVarQueryParams.Add("filter[version.last]", datadog.ParameterToString(*optionalParams.FilterVersionLast, ""))
	}
	if optionalParams.FilterRepositoryUrl != nil {
		localVarQueryParams.Add("filter[repository_url]", datadog.ParameterToString(*optionalParams.FilterRepositoryUrl, ""))
	}
	if optionalParams.FilterRisksInProduction != nil {
		localVarQueryParams.Add("filter[risks.in_production]", datadog.ParameterToString(*optionalParams.FilterRisksInProduction, ""))
	}
	if optionalParams.FilterRisksUnderAttack != nil {
		localVarQueryParams.Add("filter[risks.under_attack]", datadog.ParameterToString(*optionalParams.FilterRisksUnderAttack, ""))
	}
	if optionalParams.FilterRisksIsPubliclyAccessible != nil {
		localVarQueryParams.Add("filter[risks.is_publicly_accessible]", datadog.ParameterToString(*optionalParams.FilterRisksIsPubliclyAccessible, ""))
	}
	if optionalParams.FilterRisksHasPrivilegedAccess != nil {
		localVarQueryParams.Add("filter[risks.has_privileged_access]", datadog.ParameterToString(*optionalParams.FilterRisksHasPrivilegedAccess, ""))
	}
	if optionalParams.FilterRisksHasAccessToSensitiveData != nil {
		localVarQueryParams.Add("filter[risks.has_access_to_sensitive_data]", datadog.ParameterToString(*optionalParams.FilterRisksHasAccessToSensitiveData, ""))
	}
	if optionalParams.FilterEnvironments != nil {
		localVarQueryParams.Add("filter[environments]", datadog.ParameterToString(*optionalParams.FilterEnvironments, ""))
	}
	if optionalParams.FilterArch != nil {
		localVarQueryParams.Add("filter[arch]", datadog.ParameterToString(*optionalParams.FilterArch, ""))
	}
	if optionalParams.FilterOperatingSystemName != nil {
		localVarQueryParams.Add("filter[operating_system.name]", datadog.ParameterToString(*optionalParams.FilterOperatingSystemName, ""))
	}
	if optionalParams.FilterOperatingSystemVersion != nil {
		localVarQueryParams.Add("filter[operating_system.version]", datadog.ParameterToString(*optionalParams.FilterOperatingSystemVersion, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 {
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

// MuteFindings Mute or unmute a batch of findings.
// Mute or unmute findings.
func (a *SecurityMonitoringApi) MuteFindings(ctx _context.Context, body BulkMuteFindingsRequest) (BulkMuteFindingsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue BulkMuteFindingsResponse
	)

	operationId := "v2.MuteFindings"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.MuteFindings")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/posture_management/findings"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 422 || localVarHTTPResponse.StatusCode == 429 {
			var v JSONAPIErrorResponse
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

// PatchSignalNotificationRule Patch a signal-based notification rule.
// Partially update the notification rule. All fields are optional; if a field is not provided, it is not updated.
func (a *SecurityMonitoringApi) PatchSignalNotificationRule(ctx _context.Context, id string, body PatchNotificationRuleParameters) (NotificationRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue NotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.PatchSignalNotificationRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/signals/notification_rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v JSONAPIErrorResponse
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

// PatchVulnerabilityNotificationRule Patch a vulnerability-based notification rule.
// Partially update the notification rule. All fields are optional; if a field is not provided, it is not updated.
func (a *SecurityMonitoringApi) PatchVulnerabilityNotificationRule(ctx _context.Context, id string, body PatchNotificationRuleParameters) (NotificationRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue NotificationRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.PatchVulnerabilityNotificationRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security/vulnerabilities/notification_rules/{id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{id}", _neturl.PathEscape(datadog.ParameterToString(id, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v JSONAPIErrorResponse
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

// RunHistoricalJob Run a historical job.
// Run a historical job.
func (a *SecurityMonitoringApi) RunHistoricalJob(ctx _context.Context, body RunHistoricalJobRequest) (JobCreateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue JobCreateResponse
	)

	operationId := "v2.RunHistoricalJob"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.RunHistoricalJob")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/siem-historical-detections/jobs"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// SearchSecurityMonitoringSignalsOptionalParameters holds optional parameters for SearchSecurityMonitoringSignals.
type SearchSecurityMonitoringSignalsOptionalParameters struct {
	Body *SecurityMonitoringSignalListRequest
}

// NewSearchSecurityMonitoringSignalsOptionalParameters creates an empty struct for parameters.
func NewSearchSecurityMonitoringSignalsOptionalParameters() *SearchSecurityMonitoringSignalsOptionalParameters {
	this := SearchSecurityMonitoringSignalsOptionalParameters{}
	return &this
}

// WithBody sets the corresponding parameter name and returns the struct.
func (r *SearchSecurityMonitoringSignalsOptionalParameters) WithBody(body SecurityMonitoringSignalListRequest) *SearchSecurityMonitoringSignalsOptionalParameters {
	r.Body = &body
	return r
}

// SearchSecurityMonitoringSignals Get a list of security signals.
// Returns security signals that match a search query.
// Both this endpoint and the GET endpoint can be used interchangeably for listing
// security signals.
func (a *SecurityMonitoringApi) SearchSecurityMonitoringSignals(ctx _context.Context, o ...SearchSecurityMonitoringSignalsOptionalParameters) (SecurityMonitoringSignalsListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSignalsListResponse
		optionalParams      SearchSecurityMonitoringSignalsOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type SearchSecurityMonitoringSignalsOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.SearchSecurityMonitoringSignals")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/signals/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	if optionalParams.Body != nil {
		localVarPostBody = &optionalParams.Body
	}
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

// SearchSecurityMonitoringSignalsWithPagination provides a paginated version of SearchSecurityMonitoringSignals returning a channel with all items.
func (a *SecurityMonitoringApi) SearchSecurityMonitoringSignalsWithPagination(ctx _context.Context, o ...SearchSecurityMonitoringSignalsOptionalParameters) (<-chan datadog.PaginationResult[SecurityMonitoringSignal], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int32(10)
	if len(o) == 0 {
		o = append(o, SearchSecurityMonitoringSignalsOptionalParameters{})
	}
	if o[0].Body == nil {
		o[0].Body = NewSecurityMonitoringSignalListRequest()
	}
	if o[0].Body.Page == nil {
		o[0].Body.Page = NewSecurityMonitoringSignalListRequestPage()
	}
	if o[0].Body.Page.Limit != nil {
		pageSize_ = *o[0].Body.Page.Limit
	}
	o[0].Body.Page.Limit = &pageSize_

	items := make(chan datadog.PaginationResult[SecurityMonitoringSignal], pageSize_)
	go func() {
		for {
			resp, _, err := a.SearchSecurityMonitoringSignals(ctx, o...)
			if err != nil {
				var returnItem SecurityMonitoringSignal
				items <- datadog.PaginationResult[SecurityMonitoringSignal]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[SecurityMonitoringSignal]{Item: item, Error: nil}:
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
			cursorMetaPage, ok := cursorMeta.GetPageOk()
			if !ok {
				break
			}
			cursorMetaPageAfter, ok := cursorMetaPage.GetAfterOk()
			if !ok {
				break
			}

			o[0].Body.Page.Cursor = cursorMetaPageAfter
		}
		close(items)
	}()
	return items, cancel
}

// TestExistingSecurityMonitoringRule Test an existing rule.
// Test an existing rule.
func (a *SecurityMonitoringApi) TestExistingSecurityMonitoringRule(ctx _context.Context, ruleId string, body SecurityMonitoringRuleTestRequest) (SecurityMonitoringRuleTestResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringRuleTestResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.TestExistingSecurityMonitoringRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/{rule_id}/test"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{rule_id}", _neturl.PathEscape(datadog.ParameterToString(ruleId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// TestSecurityMonitoringRule Test a rule.
// Test a rule.
func (a *SecurityMonitoringApi) TestSecurityMonitoringRule(ctx _context.Context, body SecurityMonitoringRuleTestRequest) (SecurityMonitoringRuleTestResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringRuleTestResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.TestSecurityMonitoringRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/test"

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// UpdateCustomFramework Update a custom framework.
// Update a custom framework.
func (a *SecurityMonitoringApi) UpdateCustomFramework(ctx _context.Context, handle string, version string, body UpdateCustomFrameworkRequest) (UpdateCustomFrameworkResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue UpdateCustomFrameworkResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.UpdateCustomFramework")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/cloud_security_management/custom_frameworks/{handle}/{version}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{handle}", _neturl.PathEscape(datadog.ParameterToString(handle, "")))
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{version}", _neturl.PathEscape(datadog.ParameterToString(version, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 429 || localVarHTTPResponse.StatusCode == 500 {
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

// UpdateResourceEvaluationFilters Update resource filters.
// Update resource filters.
func (a *SecurityMonitoringApi) UpdateResourceEvaluationFilters(ctx _context.Context, body UpdateResourceEvaluationFiltersRequest) (UpdateResourceEvaluationFiltersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue UpdateResourceEvaluationFiltersResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.UpdateResourceEvaluationFilters")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/cloud_security_management/resource_filters"

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

// UpdateSecurityFilter Update a security filter.
// Update a specific security filter.
// Returns the security filter object when the request is successful.
func (a *SecurityMonitoringApi) UpdateSecurityFilter(ctx _context.Context, securityFilterId string, body SecurityFilterUpdateRequest) (SecurityFilterResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue SecurityFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.UpdateSecurityFilter")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/security_filters/{security_filter_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{security_filter_id}", _neturl.PathEscape(datadog.ParameterToString(securityFilterId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 429 {
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

// UpdateSecurityMonitoringRule Update an existing rule.
// Update an existing rule. When updating `cases`, `queries` or `options`, the whole field
// must be included. For example, when modifying a query all queries must be included.
// Default rules can only be updated to be enabled, to change notifications, or to update
// the tags (default tags cannot be removed).
func (a *SecurityMonitoringApi) UpdateSecurityMonitoringRule(ctx _context.Context, ruleId string, body SecurityMonitoringRuleUpdatePayload) (SecurityMonitoringRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringRuleResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.UpdateSecurityMonitoringRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/{rule_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{rule_id}", _neturl.PathEscape(datadog.ParameterToString(ruleId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 401 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 429 {
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

// UpdateSecurityMonitoringSuppression Update a suppression rule.
// Update a specific suppression rule.
func (a *SecurityMonitoringApi) UpdateSecurityMonitoringSuppression(ctx _context.Context, suppressionId string, body SecurityMonitoringSuppressionUpdateRequest) (SecurityMonitoringSuppressionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue SecurityMonitoringSuppressionResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.UpdateSecurityMonitoringSuppression")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/configuration/suppressions/{suppression_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{suppression_id}", _neturl.PathEscape(datadog.ParameterToString(suppressionId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 404 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 429 {
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

// ValidateSecurityMonitoringRule Validate a detection rule.
// Validate a detection rule.
func (a *SecurityMonitoringApi) ValidateSecurityMonitoringRule(ctx _context.Context, body SecurityMonitoringRuleValidatePayload) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.SecurityMonitoringApi.ValidateSecurityMonitoringRule")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/security_monitoring/rules/validation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "*/*"

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

// NewSecurityMonitoringApi Returns NewSecurityMonitoringApi.
func NewSecurityMonitoringApi(client *datadog.APIClient) *SecurityMonitoringApi {
	return &SecurityMonitoringApi{
		Client: client,
	}
}
