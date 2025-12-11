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

// ServiceScorecardsApi service type
type ServiceScorecardsApi datadog.Service

// CreateScorecardOutcomesBatch Create outcomes batch.
// Sets multiple service-rule outcomes in a single batched request.
func (a *ServiceScorecardsApi) CreateScorecardOutcomesBatch(ctx _context.Context, body OutcomesBatchRequest) (OutcomesBatchResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue OutcomesBatchResponse
	)

	operationId := "v2.CreateScorecardOutcomesBatch"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ServiceScorecardsApi.CreateScorecardOutcomesBatch")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scorecard/outcomes/batch"

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

// CreateScorecardRule Create a new rule.
// Creates a new rule.
func (a *ServiceScorecardsApi) CreateScorecardRule(ctx _context.Context, body CreateRuleRequest) (CreateRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue CreateRuleResponse
	)

	operationId := "v2.CreateScorecardRule"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ServiceScorecardsApi.CreateScorecardRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scorecard/rules"

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

// DeleteScorecardRule Delete a rule.
// Deletes a single rule.
func (a *ServiceScorecardsApi) DeleteScorecardRule(ctx _context.Context, ruleId string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodDelete
		localVarPostBody   interface{}
	)

	operationId := "v2.DeleteScorecardRule"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ServiceScorecardsApi.DeleteScorecardRule")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scorecard/rules/{rule_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{rule_id}", _neturl.PathEscape(datadog.ParameterToString(ruleId, "")))

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

// ListScorecardOutcomesOptionalParameters holds optional parameters for ListScorecardOutcomes.
type ListScorecardOutcomesOptionalParameters struct {
	PageSize                 *int64
	PageOffset               *int64
	Include                  *string
	FieldsOutcome            *string
	FieldsRule               *string
	FilterOutcomeServiceName *string
	FilterOutcomeState       *string
	FilterRuleEnabled        *bool
	FilterRuleId             *string
	FilterRuleName           *string
}

// NewListScorecardOutcomesOptionalParameters creates an empty struct for parameters.
func NewListScorecardOutcomesOptionalParameters() *ListScorecardOutcomesOptionalParameters {
	this := ListScorecardOutcomesOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithPageSize(pageSize int64) *ListScorecardOutcomesOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithPageOffset(pageOffset int64) *ListScorecardOutcomesOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithInclude(include string) *ListScorecardOutcomesOptionalParameters {
	r.Include = &include
	return r
}

// WithFieldsOutcome sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithFieldsOutcome(fieldsOutcome string) *ListScorecardOutcomesOptionalParameters {
	r.FieldsOutcome = &fieldsOutcome
	return r
}

// WithFieldsRule sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithFieldsRule(fieldsRule string) *ListScorecardOutcomesOptionalParameters {
	r.FieldsRule = &fieldsRule
	return r
}

// WithFilterOutcomeServiceName sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithFilterOutcomeServiceName(filterOutcomeServiceName string) *ListScorecardOutcomesOptionalParameters {
	r.FilterOutcomeServiceName = &filterOutcomeServiceName
	return r
}

// WithFilterOutcomeState sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithFilterOutcomeState(filterOutcomeState string) *ListScorecardOutcomesOptionalParameters {
	r.FilterOutcomeState = &filterOutcomeState
	return r
}

// WithFilterRuleEnabled sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithFilterRuleEnabled(filterRuleEnabled bool) *ListScorecardOutcomesOptionalParameters {
	r.FilterRuleEnabled = &filterRuleEnabled
	return r
}

// WithFilterRuleId sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithFilterRuleId(filterRuleId string) *ListScorecardOutcomesOptionalParameters {
	r.FilterRuleId = &filterRuleId
	return r
}

// WithFilterRuleName sets the corresponding parameter name and returns the struct.
func (r *ListScorecardOutcomesOptionalParameters) WithFilterRuleName(filterRuleName string) *ListScorecardOutcomesOptionalParameters {
	r.FilterRuleName = &filterRuleName
	return r
}

// ListScorecardOutcomes List all rule outcomes.
// Fetches all rule outcomes.
func (a *ServiceScorecardsApi) ListScorecardOutcomes(ctx _context.Context, o ...ListScorecardOutcomesOptionalParameters) (OutcomesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue OutcomesResponse
		optionalParams      ListScorecardOutcomesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListScorecardOutcomesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListScorecardOutcomes"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ServiceScorecardsApi.ListScorecardOutcomes")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scorecard/outcomes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
	if optionalParams.FieldsOutcome != nil {
		localVarQueryParams.Add("fields[outcome]", datadog.ParameterToString(*optionalParams.FieldsOutcome, ""))
	}
	if optionalParams.FieldsRule != nil {
		localVarQueryParams.Add("fields[rule]", datadog.ParameterToString(*optionalParams.FieldsRule, ""))
	}
	if optionalParams.FilterOutcomeServiceName != nil {
		localVarQueryParams.Add("filter[outcome][service_name]", datadog.ParameterToString(*optionalParams.FilterOutcomeServiceName, ""))
	}
	if optionalParams.FilterOutcomeState != nil {
		localVarQueryParams.Add("filter[outcome][state]", datadog.ParameterToString(*optionalParams.FilterOutcomeState, ""))
	}
	if optionalParams.FilterRuleEnabled != nil {
		localVarQueryParams.Add("filter[rule][enabled]", datadog.ParameterToString(*optionalParams.FilterRuleEnabled, ""))
	}
	if optionalParams.FilterRuleId != nil {
		localVarQueryParams.Add("filter[rule][id]", datadog.ParameterToString(*optionalParams.FilterRuleId, ""))
	}
	if optionalParams.FilterRuleName != nil {
		localVarQueryParams.Add("filter[rule][name]", datadog.ParameterToString(*optionalParams.FilterRuleName, ""))
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

// ListScorecardOutcomesWithPagination provides a paginated version of ListScorecardOutcomes returning a channel with all items.
func (a *ServiceScorecardsApi) ListScorecardOutcomesWithPagination(ctx _context.Context, o ...ListScorecardOutcomesOptionalParameters) (<-chan datadog.PaginationResult[OutcomesResponseDataItem], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(10)
	if len(o) == 0 {
		o = append(o, ListScorecardOutcomesOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_

	items := make(chan datadog.PaginationResult[OutcomesResponseDataItem], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListScorecardOutcomes(ctx, o...)
			if err != nil {
				var returnItem OutcomesResponseDataItem
				items <- datadog.PaginationResult[OutcomesResponseDataItem]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[OutcomesResponseDataItem]{Item: item, Error: nil}:
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

// ListScorecardRulesOptionalParameters holds optional parameters for ListScorecardRules.
type ListScorecardRulesOptionalParameters struct {
	PageSize              *int64
	PageOffset            *int64
	Include               *string
	FilterRuleId          *string
	FilterRuleEnabled     *bool
	FilterRuleCustom      *bool
	FilterRuleName        *string
	FilterRuleDescription *string
	FieldsRule            *string
	FieldsScorecard       *string
}

// NewListScorecardRulesOptionalParameters creates an empty struct for parameters.
func NewListScorecardRulesOptionalParameters() *ListScorecardRulesOptionalParameters {
	this := ListScorecardRulesOptionalParameters{}
	return &this
}

// WithPageSize sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithPageSize(pageSize int64) *ListScorecardRulesOptionalParameters {
	r.PageSize = &pageSize
	return r
}

// WithPageOffset sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithPageOffset(pageOffset int64) *ListScorecardRulesOptionalParameters {
	r.PageOffset = &pageOffset
	return r
}

// WithInclude sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithInclude(include string) *ListScorecardRulesOptionalParameters {
	r.Include = &include
	return r
}

// WithFilterRuleId sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithFilterRuleId(filterRuleId string) *ListScorecardRulesOptionalParameters {
	r.FilterRuleId = &filterRuleId
	return r
}

// WithFilterRuleEnabled sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithFilterRuleEnabled(filterRuleEnabled bool) *ListScorecardRulesOptionalParameters {
	r.FilterRuleEnabled = &filterRuleEnabled
	return r
}

// WithFilterRuleCustom sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithFilterRuleCustom(filterRuleCustom bool) *ListScorecardRulesOptionalParameters {
	r.FilterRuleCustom = &filterRuleCustom
	return r
}

// WithFilterRuleName sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithFilterRuleName(filterRuleName string) *ListScorecardRulesOptionalParameters {
	r.FilterRuleName = &filterRuleName
	return r
}

// WithFilterRuleDescription sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithFilterRuleDescription(filterRuleDescription string) *ListScorecardRulesOptionalParameters {
	r.FilterRuleDescription = &filterRuleDescription
	return r
}

// WithFieldsRule sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithFieldsRule(fieldsRule string) *ListScorecardRulesOptionalParameters {
	r.FieldsRule = &fieldsRule
	return r
}

// WithFieldsScorecard sets the corresponding parameter name and returns the struct.
func (r *ListScorecardRulesOptionalParameters) WithFieldsScorecard(fieldsScorecard string) *ListScorecardRulesOptionalParameters {
	r.FieldsScorecard = &fieldsScorecard
	return r
}

// ListScorecardRules List all rules.
// Fetch all rules.
func (a *ServiceScorecardsApi) ListScorecardRules(ctx _context.Context, o ...ListScorecardRulesOptionalParameters) (ListRulesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		localVarReturnValue ListRulesResponse
		optionalParams      ListScorecardRulesOptionalParameters
	)

	if len(o) > 1 {
		return localVarReturnValue, nil, datadog.ReportError("only one argument of type ListScorecardRulesOptionalParameters is allowed")
	}
	if len(o) == 1 {
		optionalParams = o[0]
	}

	operationId := "v2.ListScorecardRules"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ServiceScorecardsApi.ListScorecardRules")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scorecard/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if optionalParams.PageSize != nil {
		localVarQueryParams.Add("page[size]", datadog.ParameterToString(*optionalParams.PageSize, ""))
	}
	if optionalParams.PageOffset != nil {
		localVarQueryParams.Add("page[offset]", datadog.ParameterToString(*optionalParams.PageOffset, ""))
	}
	if optionalParams.Include != nil {
		localVarQueryParams.Add("include", datadog.ParameterToString(*optionalParams.Include, ""))
	}
	if optionalParams.FilterRuleId != nil {
		localVarQueryParams.Add("filter[rule][id]", datadog.ParameterToString(*optionalParams.FilterRuleId, ""))
	}
	if optionalParams.FilterRuleEnabled != nil {
		localVarQueryParams.Add("filter[rule][enabled]", datadog.ParameterToString(*optionalParams.FilterRuleEnabled, ""))
	}
	if optionalParams.FilterRuleCustom != nil {
		localVarQueryParams.Add("filter[rule][custom]", datadog.ParameterToString(*optionalParams.FilterRuleCustom, ""))
	}
	if optionalParams.FilterRuleName != nil {
		localVarQueryParams.Add("filter[rule][name]", datadog.ParameterToString(*optionalParams.FilterRuleName, ""))
	}
	if optionalParams.FilterRuleDescription != nil {
		localVarQueryParams.Add("filter[rule][description]", datadog.ParameterToString(*optionalParams.FilterRuleDescription, ""))
	}
	if optionalParams.FieldsRule != nil {
		localVarQueryParams.Add("fields[rule]", datadog.ParameterToString(*optionalParams.FieldsRule, ""))
	}
	if optionalParams.FieldsScorecard != nil {
		localVarQueryParams.Add("fields[scorecard]", datadog.ParameterToString(*optionalParams.FieldsScorecard, ""))
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

// ListScorecardRulesWithPagination provides a paginated version of ListScorecardRules returning a channel with all items.
func (a *ServiceScorecardsApi) ListScorecardRulesWithPagination(ctx _context.Context, o ...ListScorecardRulesOptionalParameters) (<-chan datadog.PaginationResult[ListRulesResponseDataItem], func()) {
	ctx, cancel := _context.WithCancel(ctx)
	pageSize_ := int64(10)
	if len(o) == 0 {
		o = append(o, ListScorecardRulesOptionalParameters{})
	}
	if o[0].PageSize != nil {
		pageSize_ = *o[0].PageSize
	}
	o[0].PageSize = &pageSize_

	items := make(chan datadog.PaginationResult[ListRulesResponseDataItem], pageSize_)
	go func() {
		for {
			resp, _, err := a.ListScorecardRules(ctx, o...)
			if err != nil {
				var returnItem ListRulesResponseDataItem
				items <- datadog.PaginationResult[ListRulesResponseDataItem]{Item: returnItem, Error: err}
				break
			}
			respData, ok := resp.GetDataOk()
			if !ok {
				break
			}
			results := *respData

			for _, item := range results {
				select {
				case items <- datadog.PaginationResult[ListRulesResponseDataItem]{Item: item, Error: nil}:
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

// UpdateScorecardOutcomesAsync Update Scorecard outcomes asynchronously.
// Updates multiple scorecard rule outcomes in a single batched request.
func (a *ServiceScorecardsApi) UpdateScorecardOutcomesAsync(ctx _context.Context, body UpdateOutcomesAsyncRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	operationId := "v2.UpdateScorecardOutcomesAsync"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ServiceScorecardsApi.UpdateScorecardOutcomesAsync")
	if err != nil {
		return nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scorecard/outcomes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "*/*"

	// body params
	localVarPostBody = &body
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
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 409 || localVarHTTPResponse.StatusCode == 429 {
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

// UpdateScorecardRule Update an existing rule.
// Updates an existing rule.
func (a *ServiceScorecardsApi) UpdateScorecardRule(ctx _context.Context, ruleId string, body UpdateRuleRequest) (UpdateRuleResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPut
		localVarPostBody    interface{}
		localVarReturnValue UpdateRuleResponse
	)

	operationId := "v2.UpdateScorecardRule"
	isOperationEnabled := a.Client.Cfg.IsUnstableOperationEnabled(operationId)
	if !isOperationEnabled {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: _fmt.Sprintf("Unstable operation '%s' is disabled", operationId)}
	}
	if isOperationEnabled && a.Client.Cfg.Debug {
		_log.Printf("WARNING: Using unstable operation '%s'", operationId)
	}

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, "v2.ServiceScorecardsApi.UpdateScorecardRule")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scorecard/rules/{rule_id}"
	localVarPath = datadog.ReplacePathParameter(localVarPath, "{rule_id}", _neturl.PathEscape(datadog.ParameterToString(ruleId, "")))

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

// NewServiceScorecardsApi Returns NewServiceScorecardsApi.
func NewServiceScorecardsApi(client *datadog.APIClient) *ServiceScorecardsApi {
	return &ServiceScorecardsApi{
		Client: client,
	}
}
