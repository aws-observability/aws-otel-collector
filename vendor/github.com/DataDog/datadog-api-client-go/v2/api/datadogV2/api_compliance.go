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

// ComplianceApi service type
type ComplianceApi datadog.Service

// GetRuleBasedViewOptionalParameters holds optional parameters for GetRuleBasedView.
type GetRuleBasedViewOptionalParameters struct {
	Framework                            *string
	Version                              *string
	QueryFindingsWithoutFrameworkVersion *bool
	IncludeRulesWithoutFindings          *bool
	IsCustom                             *bool
	Query                                *string
}

// NewGetRuleBasedViewOptionalParameters creates an empty struct for parameters.
func NewGetRuleBasedViewOptionalParameters() *GetRuleBasedViewOptionalParameters {
	this := GetRuleBasedViewOptionalParameters{}
	return &this
}

// WithFramework sets the corresponding parameter name and returns the struct.
func (r *GetRuleBasedViewOptionalParameters) WithFramework(framework string) *GetRuleBasedViewOptionalParameters {
	r.Framework = &framework
	return r
}

// WithVersion sets the corresponding parameter name and returns the struct.
func (r *GetRuleBasedViewOptionalParameters) WithVersion(version string) *GetRuleBasedViewOptionalParameters {
	r.Version = &version
	return r
}

// WithQueryFindingsWithoutFrameworkVersion sets the corresponding parameter name and returns the struct.
func (r *GetRuleBasedViewOptionalParameters) WithQueryFindingsWithoutFrameworkVersion(queryFindingsWithoutFrameworkVersion bool) *GetRuleBasedViewOptionalParameters {
	r.QueryFindingsWithoutFrameworkVersion = &queryFindingsWithoutFrameworkVersion
	return r
}

// WithIncludeRulesWithoutFindings sets the corresponding parameter name and returns the struct.
func (r *GetRuleBasedViewOptionalParameters) WithIncludeRulesWithoutFindings(includeRulesWithoutFindings bool) *GetRuleBasedViewOptionalParameters {
	r.IncludeRulesWithoutFindings = &includeRulesWithoutFindings
	return r
}

// WithIsCustom sets the corresponding parameter name and returns the struct.
func (r *GetRuleBasedViewOptionalParameters) WithIsCustom(isCustom bool) *GetRuleBasedViewOptionalParameters {
	r.IsCustom = &isCustom
	return r
}

// WithQuery sets the corresponding parameter name and returns the struct.
func (r *GetRuleBasedViewOptionalParameters) WithQuery(query string) *GetRuleBasedViewOptionalParameters {
	r.Query = &query
	return r
}

// GetRuleBasedView Get the rule-based view of compliance findings.
// Get an aggregated view of compliance rules with their pass, fail, and muted finding counts.
// Supports filtering by compliance framework, framework version, and additional query filters.
func (a *ComplianceApi) GetRuleBasedView(ctx _context.Context, to int64, o ...GetRuleBasedViewOptionalParameters) (RuleBasedViewResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue RuleBasedViewResponse
		optionalParams      GetRuleBasedViewOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type GetRuleBasedViewOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.GetRuleBasedView"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ComplianceApi.GetRuleBasedView")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/compliance_findings/rule_based_view"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("to", datadog.ParameterToString(to, ""))
	if optionalParams.Framework != nil {
		localVarQueryParams.Add("framework", datadog.ParameterToString(*optionalParams.Framework, ""))
	}
	if optionalParams.Version != nil {
		localVarQueryParams.Add("version", datadog.ParameterToString(*optionalParams.Version, ""))
	}
	if optionalParams.QueryFindingsWithoutFrameworkVersion != nil {
		localVarQueryParams.Add("query_findings_without_framework_version", datadog.ParameterToString(*optionalParams.QueryFindingsWithoutFrameworkVersion, ""))
	}
	if optionalParams.IncludeRulesWithoutFindings != nil {
		localVarQueryParams.Add("include_rules_without_findings", datadog.ParameterToString(*optionalParams.IncludeRulesWithoutFindings, ""))
	}
	if optionalParams.IsCustom != nil {
		localVarQueryParams.Add("is_custom", datadog.ParameterToString(*optionalParams.IsCustom, ""))
	}
	if optionalParams.Query != nil {
		localVarQueryParams.Add("query", datadog.ParameterToString(*optionalParams.Query, ""))
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 503 {
			var v JSONAPIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

// NewComplianceApi Returns NewComplianceApi.
func NewComplianceApi(client *datadog.APIClient) *ComplianceApi {
	return &ComplianceApi{
		Client: client,
	}
}
