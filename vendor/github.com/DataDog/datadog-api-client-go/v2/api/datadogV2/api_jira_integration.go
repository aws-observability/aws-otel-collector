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
	"github.com/google/uuid"
)

// JiraIntegrationApi service type
type JiraIntegrationApi datadog.Service

// CreateJiraIssueTemplate Create Jira issue template.
// Create a new Jira issue template.
func (a *JiraIntegrationApi) CreateJiraIssueTemplate(ctx _context.Context, body JiraIssueTemplateCreateRequest) (JiraIssueTemplateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue JiraIssueTemplateResponse
	)

	operationId := "v2.CreateJiraIssueTemplate"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.JiraIntegrationApi.CreateJiraIssueTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integration/jira/issue-templates"

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
		if localVarHTTPResponse.StatusCode == 400 {
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

// DeleteJiraAccount Delete Jira account.
// Delete a Jira account by ID.
func (a *JiraIntegrationApi) DeleteJiraAccount(ctx _context.Context, accountId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	operationId := "v2.DeleteJiraAccount"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.JiraIntegrationApi.DeleteJiraAccount")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integration/jira/accounts/{account_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{account_id}", _neturl.PathEscape(datadog.ParameterToString(accountId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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

// DeleteJiraIssueTemplate Delete Jira issue template.
// Delete a Jira issue template by ID.
func (a *JiraIntegrationApi) DeleteJiraIssueTemplate(ctx _context.Context, issueTemplateId uuid.UUID) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	operationId := "v2.DeleteJiraIssueTemplate"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.JiraIntegrationApi.DeleteJiraIssueTemplate")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integration/jira/issue-templates/{issue_template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{issue_template_id}", _neturl.PathEscape(datadog.ParameterToString(issueTemplateId, "")))

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
		if localVarHTTPResponse.StatusCode == 429 {
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

// GetJiraIssueTemplate Get Jira issue template.
// Get a Jira issue template by ID.
func (a *JiraIntegrationApi) GetJiraIssueTemplate(ctx _context.Context, issueTemplateId uuid.UUID) (JiraIssueTemplateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue JiraIssueTemplateResponse
	)

	operationId := "v2.GetJiraIssueTemplate"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.JiraIntegrationApi.GetJiraIssueTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integration/jira/issue-templates/{issue_template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{issue_template_id}", _neturl.PathEscape(datadog.ParameterToString(issueTemplateId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
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

// ListJiraAccounts List Jira accounts.
// Get all Jira accounts for the organization.
func (a *JiraIntegrationApi) ListJiraAccounts(ctx _context.Context) (JiraAccountsResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue JiraAccountsResponse
	)

	operationId := "v2.ListJiraAccounts"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.JiraIntegrationApi.ListJiraAccounts")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integration/jira/accounts"

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
		if localVarHTTPResponse.StatusCode == 404 {
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

// ListJiraIssueTemplates List Jira issue templates.
// Get all Jira issue templates for the organization.
func (a *JiraIntegrationApi) ListJiraIssueTemplates(ctx _context.Context) (JiraIssueTemplatesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue JiraIssueTemplatesResponse
	)

	operationId := "v2.ListJiraIssueTemplates"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.JiraIntegrationApi.ListJiraIssueTemplates")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integration/jira/issue-templates"

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

// UpdateJiraIssueTemplate Update Jira issue template.
// Update a Jira issue template by ID.
func (a *JiraIntegrationApi) UpdateJiraIssueTemplate(ctx _context.Context, issueTemplateId uuid.UUID, body JiraIssueTemplateUpdateRequest) (JiraIssueTemplateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		localVarReturnValue JiraIssueTemplateResponse
	)

	operationId := "v2.UpdateJiraIssueTemplate"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.JiraIntegrationApi.UpdateJiraIssueTemplate")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/integration/jira/issue-templates/{issue_template_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{issue_template_id}", _neturl.PathEscape(datadog.ParameterToString(issueTemplateId, "")))

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
		if localVarHTTPResponse.StatusCode == 400 {
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

// NewJiraIntegrationApi Returns NewJiraIntegrationApi.
func NewJiraIntegrationApi(client *datadog.APIClient) *JiraIntegrationApi {
	return &JiraIntegrationApi{
		Client: client,
	}
}
