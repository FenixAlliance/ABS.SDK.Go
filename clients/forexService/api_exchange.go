/*
ForexService

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.1.4089
Contact: support@fenix-alliance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ExchangeAPIService ExchangeAPI service
type ExchangeAPIService service

type ApiApiV2ForexServiceExchangeHistoryGetRequest struct {
	ctx context.Context
	ApiService *ExchangeAPIService
	amount *float64
	sourceCurrencyId *string
	targetCurrencyId *string
	date *string
}

func (r ApiApiV2ForexServiceExchangeHistoryGetRequest) Amount(amount float64) ApiApiV2ForexServiceExchangeHistoryGetRequest {
	r.amount = &amount
	return r
}

func (r ApiApiV2ForexServiceExchangeHistoryGetRequest) SourceCurrencyId(sourceCurrencyId string) ApiApiV2ForexServiceExchangeHistoryGetRequest {
	r.sourceCurrencyId = &sourceCurrencyId
	return r
}

func (r ApiApiV2ForexServiceExchangeHistoryGetRequest) TargetCurrencyId(targetCurrencyId string) ApiApiV2ForexServiceExchangeHistoryGetRequest {
	r.targetCurrencyId = &targetCurrencyId
	return r
}

func (r ApiApiV2ForexServiceExchangeHistoryGetRequest) Date(date string) ApiApiV2ForexServiceExchangeHistoryGetRequest {
	r.date = &date
	return r
}

func (r ApiApiV2ForexServiceExchangeHistoryGetRequest) Execute() (*MoneyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2ForexServiceExchangeHistoryGetExecute(r)
}

/*
ApiV2ForexServiceExchangeHistoryGet Method for ApiV2ForexServiceExchangeHistoryGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2ForexServiceExchangeHistoryGetRequest
*/
func (a *ExchangeAPIService) ApiV2ForexServiceExchangeHistoryGet(ctx context.Context) ApiApiV2ForexServiceExchangeHistoryGetRequest {
	return ApiApiV2ForexServiceExchangeHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MoneyEnvelope
func (a *ExchangeAPIService) ApiV2ForexServiceExchangeHistoryGetExecute(r ApiApiV2ForexServiceExchangeHistoryGetRequest) (*MoneyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MoneyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangeAPIService.ApiV2ForexServiceExchangeHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/ForexService/Exchange/History"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.amount == nil {
		return localVarReturnValue, nil, reportError("amount is required and must be specified")
	}
	if r.sourceCurrencyId == nil {
		return localVarReturnValue, nil, reportError("sourceCurrencyId is required and must be specified")
	}
	if r.targetCurrencyId == nil {
		return localVarReturnValue, nil, reportError("targetCurrencyId is required and must be specified")
	}
	if r.date == nil {
		return localVarReturnValue, nil, reportError("date is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "sourceCurrencyId", r.sourceCurrencyId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "targetCurrencyId", r.targetCurrencyId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV2ForexServiceExchangeLatestGetRequest struct {
	ctx context.Context
	ApiService *ExchangeAPIService
	amount *float64
	sourceCurrencyId *string
	targetCurrencyId *string
}

func (r ApiApiV2ForexServiceExchangeLatestGetRequest) Amount(amount float64) ApiApiV2ForexServiceExchangeLatestGetRequest {
	r.amount = &amount
	return r
}

func (r ApiApiV2ForexServiceExchangeLatestGetRequest) SourceCurrencyId(sourceCurrencyId string) ApiApiV2ForexServiceExchangeLatestGetRequest {
	r.sourceCurrencyId = &sourceCurrencyId
	return r
}

func (r ApiApiV2ForexServiceExchangeLatestGetRequest) TargetCurrencyId(targetCurrencyId string) ApiApiV2ForexServiceExchangeLatestGetRequest {
	r.targetCurrencyId = &targetCurrencyId
	return r
}

func (r ApiApiV2ForexServiceExchangeLatestGetRequest) Execute() (*MoneyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2ForexServiceExchangeLatestGetExecute(r)
}

/*
ApiV2ForexServiceExchangeLatestGet Method for ApiV2ForexServiceExchangeLatestGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2ForexServiceExchangeLatestGetRequest
*/
func (a *ExchangeAPIService) ApiV2ForexServiceExchangeLatestGet(ctx context.Context) ApiApiV2ForexServiceExchangeLatestGetRequest {
	return ApiApiV2ForexServiceExchangeLatestGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MoneyEnvelope
func (a *ExchangeAPIService) ApiV2ForexServiceExchangeLatestGetExecute(r ApiApiV2ForexServiceExchangeLatestGetRequest) (*MoneyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MoneyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExchangeAPIService.ApiV2ForexServiceExchangeLatestGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/ForexService/Exchange/Latest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.amount == nil {
		return localVarReturnValue, nil, reportError("amount is required and must be specified")
	}
	if r.sourceCurrencyId == nil {
		return localVarReturnValue, nil, reportError("sourceCurrencyId is required and must be specified")
	}
	if r.targetCurrencyId == nil {
		return localVarReturnValue, nil, reportError("targetCurrencyId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "sourceCurrencyId", r.sourceCurrencyId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "targetCurrencyId", r.targetCurrencyId, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
