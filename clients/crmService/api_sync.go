/*
CrmService

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


// SyncAPIService SyncAPI service
type SyncAPIService service

type ApiApiV2CrmServiceSyncMePostRequest struct {
	ctx context.Context
	ApiService *SyncAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2CrmServiceSyncMePostRequest) TenantId(tenantId string) ApiApiV2CrmServiceSyncMePostRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2CrmServiceSyncMePostRequest) ApiVersion(apiVersion string) ApiApiV2CrmServiceSyncMePostRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2CrmServiceSyncMePostRequest) XApiVersion(xApiVersion string) ApiApiV2CrmServiceSyncMePostRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2CrmServiceSyncMePostRequest) Execute() (*ContactDtoListEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2CrmServiceSyncMePostExecute(r)
}

/*
ApiV2CrmServiceSyncMePost Method for ApiV2CrmServiceSyncMePost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2CrmServiceSyncMePostRequest
*/
func (a *SyncAPIService) ApiV2CrmServiceSyncMePost(ctx context.Context) ApiApiV2CrmServiceSyncMePostRequest {
	return ApiApiV2CrmServiceSyncMePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ContactDtoListEnvelope
func (a *SyncAPIService) ApiV2CrmServiceSyncMePostExecute(r ApiApiV2CrmServiceSyncMePostRequest) (*ContactDtoListEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContactDtoListEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyncAPIService.ApiV2CrmServiceSyncMePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/CrmService/Sync/Me"

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

type ApiApiV2CrmServiceSyncPostRequest struct {
	ctx context.Context
	ApiService *SyncAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2CrmServiceSyncPostRequest) TenantId(tenantId string) ApiApiV2CrmServiceSyncPostRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2CrmServiceSyncPostRequest) ApiVersion(apiVersion string) ApiApiV2CrmServiceSyncPostRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2CrmServiceSyncPostRequest) XApiVersion(xApiVersion string) ApiApiV2CrmServiceSyncPostRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2CrmServiceSyncPostRequest) Execute() (*ContactDtoListEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2CrmServiceSyncPostExecute(r)
}

/*
ApiV2CrmServiceSyncPost Method for ApiV2CrmServiceSyncPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2CrmServiceSyncPostRequest
*/
func (a *SyncAPIService) ApiV2CrmServiceSyncPost(ctx context.Context) ApiApiV2CrmServiceSyncPostRequest {
	return ApiApiV2CrmServiceSyncPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ContactDtoListEnvelope
func (a *SyncAPIService) ApiV2CrmServiceSyncPostExecute(r ApiApiV2CrmServiceSyncPostRequest) (*ContactDtoListEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContactDtoListEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyncAPIService.ApiV2CrmServiceSyncPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/CrmService/Sync"

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

type ApiApiV2CrmServiceSyncTenantPostRequest struct {
	ctx context.Context
	ApiService *SyncAPIService
	tenantId *string
	relatedTenantId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2CrmServiceSyncTenantPostRequest) TenantId(tenantId string) ApiApiV2CrmServiceSyncTenantPostRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2CrmServiceSyncTenantPostRequest) RelatedTenantId(relatedTenantId string) ApiApiV2CrmServiceSyncTenantPostRequest {
	r.relatedTenantId = &relatedTenantId
	return r
}

func (r ApiApiV2CrmServiceSyncTenantPostRequest) ApiVersion(apiVersion string) ApiApiV2CrmServiceSyncTenantPostRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2CrmServiceSyncTenantPostRequest) XApiVersion(xApiVersion string) ApiApiV2CrmServiceSyncTenantPostRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2CrmServiceSyncTenantPostRequest) Execute() (*EmptyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2CrmServiceSyncTenantPostExecute(r)
}

/*
ApiV2CrmServiceSyncTenantPost Method for ApiV2CrmServiceSyncTenantPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2CrmServiceSyncTenantPostRequest
*/
func (a *SyncAPIService) ApiV2CrmServiceSyncTenantPost(ctx context.Context) ApiApiV2CrmServiceSyncTenantPostRequest {
	return ApiApiV2CrmServiceSyncTenantPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmptyEnvelope
func (a *SyncAPIService) ApiV2CrmServiceSyncTenantPostExecute(r ApiApiV2CrmServiceSyncTenantPostRequest) (*EmptyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmptyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyncAPIService.ApiV2CrmServiceSyncTenantPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/CrmService/Sync/Tenant"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return localVarReturnValue, nil, reportError("tenantId is required and must be specified")
	}
	if r.relatedTenantId == nil {
		return localVarReturnValue, nil, reportError("relatedTenantId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "relatedTenantId", r.relatedTenantId, "form", "")
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

type ApiApiV2CrmServiceSyncUserPostRequest struct {
	ctx context.Context
	ApiService *SyncAPIService
	tenantId *string
	relatedUserId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2CrmServiceSyncUserPostRequest) TenantId(tenantId string) ApiApiV2CrmServiceSyncUserPostRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2CrmServiceSyncUserPostRequest) RelatedUserId(relatedUserId string) ApiApiV2CrmServiceSyncUserPostRequest {
	r.relatedUserId = &relatedUserId
	return r
}

func (r ApiApiV2CrmServiceSyncUserPostRequest) ApiVersion(apiVersion string) ApiApiV2CrmServiceSyncUserPostRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2CrmServiceSyncUserPostRequest) XApiVersion(xApiVersion string) ApiApiV2CrmServiceSyncUserPostRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2CrmServiceSyncUserPostRequest) Execute() (*ContactDtoListEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2CrmServiceSyncUserPostExecute(r)
}

/*
ApiV2CrmServiceSyncUserPost Method for ApiV2CrmServiceSyncUserPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2CrmServiceSyncUserPostRequest
*/
func (a *SyncAPIService) ApiV2CrmServiceSyncUserPost(ctx context.Context) ApiApiV2CrmServiceSyncUserPostRequest {
	return ApiApiV2CrmServiceSyncUserPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ContactDtoListEnvelope
func (a *SyncAPIService) ApiV2CrmServiceSyncUserPostExecute(r ApiApiV2CrmServiceSyncUserPostRequest) (*ContactDtoListEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContactDtoListEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyncAPIService.ApiV2CrmServiceSyncUserPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/CrmService/Sync/User"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return localVarReturnValue, nil, reportError("tenantId is required and must be specified")
	}
	if r.relatedUserId == nil {
		return localVarReturnValue, nil, reportError("relatedUserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tenantId", r.tenantId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "relatedUserId", r.relatedUserId, "form", "")
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