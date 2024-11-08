/*
MarketingService

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
	"strings"
)


// MarketingListsAPIService MarketingListsAPI service
type MarketingListsAPIService service

type ApiApiV2MarketingServiceMarketingListsCountGetRequest struct {
	ctx context.Context
	ApiService *MarketingListsAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2MarketingServiceMarketingListsCountGetRequest) TenantId(tenantId string) ApiApiV2MarketingServiceMarketingListsCountGetRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsCountGetRequest) ApiVersion(apiVersion string) ApiApiV2MarketingServiceMarketingListsCountGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsCountGetRequest) XApiVersion(xApiVersion string) ApiApiV2MarketingServiceMarketingListsCountGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsCountGetRequest) Execute() (*Int32Envelope, *http.Response, error) {
	return r.ApiService.ApiV2MarketingServiceMarketingListsCountGetExecute(r)
}

/*
ApiV2MarketingServiceMarketingListsCountGet Method for ApiV2MarketingServiceMarketingListsCountGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2MarketingServiceMarketingListsCountGetRequest
*/
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsCountGet(ctx context.Context) ApiApiV2MarketingServiceMarketingListsCountGetRequest {
	return ApiApiV2MarketingServiceMarketingListsCountGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Int32Envelope
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsCountGetExecute(r ApiApiV2MarketingServiceMarketingListsCountGetRequest) (*Int32Envelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Int32Envelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingListsAPIService.ApiV2MarketingServiceMarketingListsCountGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/MarketingService/MarketingLists/Count"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return localVarReturnValue, nil, reportError("tenantId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "form", "")
	if r.apiVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "api-version", r.apiVersion, "form", "")
	}
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
	if r.xApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", r.xApiVersion, "simple", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiApiV2MarketingServiceMarketingListsGetRequest struct {
	ctx context.Context
	ApiService *MarketingListsAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2MarketingServiceMarketingListsGetRequest) TenantId(tenantId string) ApiApiV2MarketingServiceMarketingListsGetRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsGetRequest) ApiVersion(apiVersion string) ApiApiV2MarketingServiceMarketingListsGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsGetRequest) XApiVersion(xApiVersion string) ApiApiV2MarketingServiceMarketingListsGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsGetRequest) Execute() (*MarketingListDtoListEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2MarketingServiceMarketingListsGetExecute(r)
}

/*
ApiV2MarketingServiceMarketingListsGet Method for ApiV2MarketingServiceMarketingListsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2MarketingServiceMarketingListsGetRequest
*/
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsGet(ctx context.Context) ApiApiV2MarketingServiceMarketingListsGetRequest {
	return ApiApiV2MarketingServiceMarketingListsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MarketingListDtoListEnvelope
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsGetExecute(r ApiApiV2MarketingServiceMarketingListsGetRequest) (*MarketingListDtoListEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MarketingListDtoListEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingListsAPIService.ApiV2MarketingServiceMarketingListsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/MarketingService/MarketingLists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return localVarReturnValue, nil, reportError("tenantId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "form", "")
	if r.apiVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "api-version", r.apiVersion, "form", "")
	}
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
	if r.xApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", r.xApiVersion, "simple", "")
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

type ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest struct {
	ctx context.Context
	ApiService *MarketingListsAPIService
	tenantId *string
	marketinglistId string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest) TenantId(tenantId string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest) ApiVersion(apiVersion string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest) XApiVersion(xApiVersion string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest) Execute() (*EmptyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2MarketingServiceMarketingListsMarketinglistIdDeleteExecute(r)
}

/*
ApiV2MarketingServiceMarketingListsMarketinglistIdDelete Method for ApiV2MarketingServiceMarketingListsMarketinglistIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketinglistId
 @return ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest
*/
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsMarketinglistIdDelete(ctx context.Context, marketinglistId string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest {
	return ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		marketinglistId: marketinglistId,
	}
}

// Execute executes the request
//  @return EmptyEnvelope
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsMarketinglistIdDeleteExecute(r ApiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest) (*EmptyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmptyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingListsAPIService.ApiV2MarketingServiceMarketingListsMarketinglistIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/MarketingService/MarketingLists/{marketinglistId}"
	localVarPath = strings.Replace(localVarPath, "{"+"marketinglistId"+"}", url.PathEscape(parameterValueToString(r.marketinglistId, "marketinglistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return localVarReturnValue, nil, reportError("tenantId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "form", "")
	if r.apiVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "api-version", r.apiVersion, "form", "")
	}
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
	if r.xApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", r.xApiVersion, "simple", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest struct {
	ctx context.Context
	ApiService *MarketingListsAPIService
	tenantId *string
	marketinglistId string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest) TenantId(tenantId string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest) ApiVersion(apiVersion string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest) XApiVersion(xApiVersion string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest) Execute() (*MarketingListDtoEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2MarketingServiceMarketingListsMarketinglistIdGetExecute(r)
}

/*
ApiV2MarketingServiceMarketingListsMarketinglistIdGet Method for ApiV2MarketingServiceMarketingListsMarketinglistIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketinglistId
 @return ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest
*/
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsMarketinglistIdGet(ctx context.Context, marketinglistId string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest {
	return ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest{
		ApiService: a,
		ctx: ctx,
		marketinglistId: marketinglistId,
	}
}

// Execute executes the request
//  @return MarketingListDtoEnvelope
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsMarketinglistIdGetExecute(r ApiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest) (*MarketingListDtoEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MarketingListDtoEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingListsAPIService.ApiV2MarketingServiceMarketingListsMarketinglistIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/MarketingService/MarketingLists/{marketinglistId}"
	localVarPath = strings.Replace(localVarPath, "{"+"marketinglistId"+"}", url.PathEscape(parameterValueToString(r.marketinglistId, "marketinglistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return localVarReturnValue, nil, reportError("tenantId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "form", "")
	if r.apiVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "api-version", r.apiVersion, "form", "")
	}
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
	if r.xApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", r.xApiVersion, "simple", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest struct {
	ctx context.Context
	ApiService *MarketingListsAPIService
	tenantId *string
	marketinglistId string
	marketingListUpdateDto *MarketingListUpdateDto
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest) TenantId(tenantId string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest) MarketingListUpdateDto(marketingListUpdateDto MarketingListUpdateDto) ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest {
	r.marketingListUpdateDto = &marketingListUpdateDto
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest) ApiVersion(apiVersion string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest) XApiVersion(xApiVersion string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest) Execute() (*EmptyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2MarketingServiceMarketingListsMarketinglistIdPutExecute(r)
}

/*
ApiV2MarketingServiceMarketingListsMarketinglistIdPut Method for ApiV2MarketingServiceMarketingListsMarketinglistIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marketinglistId
 @return ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest
*/
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsMarketinglistIdPut(ctx context.Context, marketinglistId string) ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest {
	return ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest{
		ApiService: a,
		ctx: ctx,
		marketinglistId: marketinglistId,
	}
}

// Execute executes the request
//  @return EmptyEnvelope
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsMarketinglistIdPutExecute(r ApiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest) (*EmptyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmptyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingListsAPIService.ApiV2MarketingServiceMarketingListsMarketinglistIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/MarketingService/MarketingLists/{marketinglistId}"
	localVarPath = strings.Replace(localVarPath, "{"+"marketinglistId"+"}", url.PathEscape(parameterValueToString(r.marketinglistId, "marketinglistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return localVarReturnValue, nil, reportError("tenantId is required and must be specified")
	}
	if r.marketingListUpdateDto == nil {
		return localVarReturnValue, nil, reportError("marketingListUpdateDto is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "form", "")
	if r.apiVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "api-version", r.apiVersion, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	if r.xApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", r.xApiVersion, "simple", "")
	}
	// body params
	localVarPostBody = r.marketingListUpdateDto
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiApiV2MarketingServiceMarketingListsPostRequest struct {
	ctx context.Context
	ApiService *MarketingListsAPIService
	tenantId *string
	marketingListCreateDto *MarketingListCreateDto
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2MarketingServiceMarketingListsPostRequest) TenantId(tenantId string) ApiApiV2MarketingServiceMarketingListsPostRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsPostRequest) MarketingListCreateDto(marketingListCreateDto MarketingListCreateDto) ApiApiV2MarketingServiceMarketingListsPostRequest {
	r.marketingListCreateDto = &marketingListCreateDto
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsPostRequest) ApiVersion(apiVersion string) ApiApiV2MarketingServiceMarketingListsPostRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsPostRequest) XApiVersion(xApiVersion string) ApiApiV2MarketingServiceMarketingListsPostRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2MarketingServiceMarketingListsPostRequest) Execute() (*EmptyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2MarketingServiceMarketingListsPostExecute(r)
}

/*
ApiV2MarketingServiceMarketingListsPost Method for ApiV2MarketingServiceMarketingListsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2MarketingServiceMarketingListsPostRequest
*/
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsPost(ctx context.Context) ApiApiV2MarketingServiceMarketingListsPostRequest {
	return ApiApiV2MarketingServiceMarketingListsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmptyEnvelope
func (a *MarketingListsAPIService) ApiV2MarketingServiceMarketingListsPostExecute(r ApiApiV2MarketingServiceMarketingListsPostRequest) (*EmptyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmptyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingListsAPIService.ApiV2MarketingServiceMarketingListsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/MarketingService/MarketingLists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return localVarReturnValue, nil, reportError("tenantId is required and must be specified")
	}
	if r.marketingListCreateDto == nil {
		return localVarReturnValue, nil, reportError("marketingListCreateDto is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "form", "")
	if r.apiVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "api-version", r.apiVersion, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	if r.xApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-api-version", r.xApiVersion, "simple", "")
	}
	// body params
	localVarPostBody = r.marketingListCreateDto
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
		if localVarHTTPResponse.StatusCode == 400 {
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
