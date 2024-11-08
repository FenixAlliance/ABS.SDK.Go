/*
SystemService

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


// LicensesAPIService LicensesAPI service
type LicensesAPIService service

type ApiApiLicensingLicensesGeneratePostRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
	licenseKeyRequest *LicenseKeyRequest
}

func (r ApiApiLicensingLicensesGeneratePostRequest) TenantId(tenantId string) ApiApiLicensingLicensesGeneratePostRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiLicensingLicensesGeneratePostRequest) ApiVersion(apiVersion string) ApiApiLicensingLicensesGeneratePostRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiLicensingLicensesGeneratePostRequest) XApiVersion(xApiVersion string) ApiApiLicensingLicensesGeneratePostRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiLicensingLicensesGeneratePostRequest) LicenseKeyRequest(licenseKeyRequest LicenseKeyRequest) ApiApiLicensingLicensesGeneratePostRequest {
	r.licenseKeyRequest = &licenseKeyRequest
	return r
}

func (r ApiApiLicensingLicensesGeneratePostRequest) Execute() (*StringEnvelope, *http.Response, error) {
	return r.ApiService.ApiLicensingLicensesGeneratePostExecute(r)
}

/*
ApiLicensingLicensesGeneratePost Method for ApiLicensingLicensesGeneratePost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiLicensingLicensesGeneratePostRequest
*/
func (a *LicensesAPIService) ApiLicensingLicensesGeneratePost(ctx context.Context) ApiApiLicensingLicensesGeneratePostRequest {
	return ApiApiLicensingLicensesGeneratePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StringEnvelope
func (a *LicensesAPIService) ApiLicensingLicensesGeneratePostExecute(r ApiApiLicensingLicensesGeneratePostRequest) (*StringEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StringEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.ApiLicensingLicensesGeneratePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Licensing/Licenses/Generate"

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
	localVarPostBody = r.licenseKeyRequest
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

type ApiApiLicensingLicensesValidateAttributesGetRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
	licenseKey *LicenseKey
}

func (r ApiApiLicensingLicensesValidateAttributesGetRequest) TenantId(tenantId string) ApiApiLicensingLicensesValidateAttributesGetRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiLicensingLicensesValidateAttributesGetRequest) ApiVersion(apiVersion string) ApiApiLicensingLicensesValidateAttributesGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiLicensingLicensesValidateAttributesGetRequest) XApiVersion(xApiVersion string) ApiApiLicensingLicensesValidateAttributesGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiLicensingLicensesValidateAttributesGetRequest) LicenseKey(licenseKey LicenseKey) ApiApiLicensingLicensesValidateAttributesGetRequest {
	r.licenseKey = &licenseKey
	return r
}

func (r ApiApiLicensingLicensesValidateAttributesGetRequest) Execute() (*LicenseAttributesListEnvelope, *http.Response, error) {
	return r.ApiService.ApiLicensingLicensesValidateAttributesGetExecute(r)
}

/*
ApiLicensingLicensesValidateAttributesGet Method for ApiLicensingLicensesValidateAttributesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiLicensingLicensesValidateAttributesGetRequest
*/
func (a *LicensesAPIService) ApiLicensingLicensesValidateAttributesGet(ctx context.Context) ApiApiLicensingLicensesValidateAttributesGetRequest {
	return ApiApiLicensingLicensesValidateAttributesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LicenseAttributesListEnvelope
func (a *LicensesAPIService) ApiLicensingLicensesValidateAttributesGetExecute(r ApiApiLicensingLicensesValidateAttributesGetRequest) (*LicenseAttributesListEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LicenseAttributesListEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.ApiLicensingLicensesValidateAttributesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Licensing/Licenses/Validate/Attributes"

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
	localVarPostBody = r.licenseKey
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

type ApiApiLicensingLicensesValidateErrorsGetRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
	licenseKey *LicenseKey
}

func (r ApiApiLicensingLicensesValidateErrorsGetRequest) TenantId(tenantId string) ApiApiLicensingLicensesValidateErrorsGetRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiLicensingLicensesValidateErrorsGetRequest) ApiVersion(apiVersion string) ApiApiLicensingLicensesValidateErrorsGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiLicensingLicensesValidateErrorsGetRequest) XApiVersion(xApiVersion string) ApiApiLicensingLicensesValidateErrorsGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiLicensingLicensesValidateErrorsGetRequest) LicenseKey(licenseKey LicenseKey) ApiApiLicensingLicensesValidateErrorsGetRequest {
	r.licenseKey = &licenseKey
	return r
}

func (r ApiApiLicensingLicensesValidateErrorsGetRequest) Execute() (*LicenseValidationErrorListEnvelope, *http.Response, error) {
	return r.ApiService.ApiLicensingLicensesValidateErrorsGetExecute(r)
}

/*
ApiLicensingLicensesValidateErrorsGet Method for ApiLicensingLicensesValidateErrorsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiLicensingLicensesValidateErrorsGetRequest
*/
func (a *LicensesAPIService) ApiLicensingLicensesValidateErrorsGet(ctx context.Context) ApiApiLicensingLicensesValidateErrorsGetRequest {
	return ApiApiLicensingLicensesValidateErrorsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LicenseValidationErrorListEnvelope
func (a *LicensesAPIService) ApiLicensingLicensesValidateErrorsGetExecute(r ApiApiLicensingLicensesValidateErrorsGetRequest) (*LicenseValidationErrorListEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LicenseValidationErrorListEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.ApiLicensingLicensesValidateErrorsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Licensing/Licenses/Validate/Errors"

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
	localVarPostBody = r.licenseKey
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

type ApiApiLicensingLicensesValidateGetRequest struct {
	ctx context.Context
	ApiService *LicensesAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
	licenseKey *LicenseKey
}

func (r ApiApiLicensingLicensesValidateGetRequest) TenantId(tenantId string) ApiApiLicensingLicensesValidateGetRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiLicensingLicensesValidateGetRequest) ApiVersion(apiVersion string) ApiApiLicensingLicensesValidateGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiLicensingLicensesValidateGetRequest) XApiVersion(xApiVersion string) ApiApiLicensingLicensesValidateGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiLicensingLicensesValidateGetRequest) LicenseKey(licenseKey LicenseKey) ApiApiLicensingLicensesValidateGetRequest {
	r.licenseKey = &licenseKey
	return r
}

func (r ApiApiLicensingLicensesValidateGetRequest) Execute() (*BooleanEnvelope, *http.Response, error) {
	return r.ApiService.ApiLicensingLicensesValidateGetExecute(r)
}

/*
ApiLicensingLicensesValidateGet Method for ApiLicensingLicensesValidateGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiLicensingLicensesValidateGetRequest
*/
func (a *LicensesAPIService) ApiLicensingLicensesValidateGet(ctx context.Context) ApiApiLicensingLicensesValidateGetRequest {
	return ApiApiLicensingLicensesValidateGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BooleanEnvelope
func (a *LicensesAPIService) ApiLicensingLicensesValidateGetExecute(r ApiApiLicensingLicensesValidateGetRequest) (*BooleanEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BooleanEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesAPIService.ApiLicensingLicensesValidateGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Licensing/Licenses/Validate"

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
	localVarPostBody = r.licenseKey
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
