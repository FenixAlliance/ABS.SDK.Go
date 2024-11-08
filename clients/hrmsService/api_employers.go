/*
HrmsService

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


// EmployersAPIService EmployersAPI service
type EmployersAPIService service

type ApiCreateEmployerAsyncRequest struct {
	ctx context.Context
	ApiService *EmployersAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
	employerProfileCreateDto *EmployerProfileCreateDto
}

func (r ApiCreateEmployerAsyncRequest) TenantId(tenantId string) ApiCreateEmployerAsyncRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiCreateEmployerAsyncRequest) ApiVersion(apiVersion string) ApiCreateEmployerAsyncRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiCreateEmployerAsyncRequest) XApiVersion(xApiVersion string) ApiCreateEmployerAsyncRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiCreateEmployerAsyncRequest) EmployerProfileCreateDto(employerProfileCreateDto EmployerProfileCreateDto) ApiCreateEmployerAsyncRequest {
	r.employerProfileCreateDto = &employerProfileCreateDto
	return r
}

func (r ApiCreateEmployerAsyncRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateEmployerAsyncExecute(r)
}

/*
CreateEmployerAsync Method for CreateEmployerAsync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateEmployerAsyncRequest
*/
func (a *EmployersAPIService) CreateEmployerAsync(ctx context.Context) ApiCreateEmployerAsyncRequest {
	return ApiCreateEmployerAsyncRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *EmployersAPIService) CreateEmployerAsyncExecute(r ApiCreateEmployerAsyncRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployersAPIService.CreateEmployerAsync")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/HrmsService/Employers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return nil, reportError("tenantId is required and must be specified")
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
	localVarPostBody = r.employerProfileCreateDto
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteEmployerAsyncRequest struct {
	ctx context.Context
	ApiService *EmployersAPIService
	tenantId *string
	employerId string
	apiVersion *string
	xApiVersion *string
}

func (r ApiDeleteEmployerAsyncRequest) TenantId(tenantId string) ApiDeleteEmployerAsyncRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiDeleteEmployerAsyncRequest) ApiVersion(apiVersion string) ApiDeleteEmployerAsyncRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiDeleteEmployerAsyncRequest) XApiVersion(xApiVersion string) ApiDeleteEmployerAsyncRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiDeleteEmployerAsyncRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEmployerAsyncExecute(r)
}

/*
DeleteEmployerAsync Method for DeleteEmployerAsync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param employerId
 @return ApiDeleteEmployerAsyncRequest
*/
func (a *EmployersAPIService) DeleteEmployerAsync(ctx context.Context, employerId string) ApiDeleteEmployerAsyncRequest {
	return ApiDeleteEmployerAsyncRequest{
		ApiService: a,
		ctx: ctx,
		employerId: employerId,
	}
}

// Execute executes the request
func (a *EmployersAPIService) DeleteEmployerAsyncExecute(r ApiDeleteEmployerAsyncRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployersAPIService.DeleteEmployerAsync")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/HrmsService/Employers/{employerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"employerId"+"}", url.PathEscape(parameterValueToString(r.employerId, "employerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return nil, reportError("tenantId is required and must be specified")
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetEmployerByIdAsyncRequest struct {
	ctx context.Context
	ApiService *EmployersAPIService
	tenantId *string
	employerId string
	apiVersion *string
	xApiVersion *string
}

func (r ApiGetEmployerByIdAsyncRequest) TenantId(tenantId string) ApiGetEmployerByIdAsyncRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiGetEmployerByIdAsyncRequest) ApiVersion(apiVersion string) ApiGetEmployerByIdAsyncRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiGetEmployerByIdAsyncRequest) XApiVersion(xApiVersion string) ApiGetEmployerByIdAsyncRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiGetEmployerByIdAsyncRequest) Execute() (*EmployerProfileDtoEnvelope, *http.Response, error) {
	return r.ApiService.GetEmployerByIdAsyncExecute(r)
}

/*
GetEmployerByIdAsync Method for GetEmployerByIdAsync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param employerId
 @return ApiGetEmployerByIdAsyncRequest
*/
func (a *EmployersAPIService) GetEmployerByIdAsync(ctx context.Context, employerId string) ApiGetEmployerByIdAsyncRequest {
	return ApiGetEmployerByIdAsyncRequest{
		ApiService: a,
		ctx: ctx,
		employerId: employerId,
	}
}

// Execute executes the request
//  @return EmployerProfileDtoEnvelope
func (a *EmployersAPIService) GetEmployerByIdAsyncExecute(r ApiGetEmployerByIdAsyncRequest) (*EmployerProfileDtoEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployerProfileDtoEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployersAPIService.GetEmployerByIdAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/HrmsService/Employers/{employerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"employerId"+"}", url.PathEscape(parameterValueToString(r.employerId, "employerId")), -1)

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

type ApiGetEmployersAsyncRequest struct {
	ctx context.Context
	ApiService *EmployersAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiGetEmployersAsyncRequest) TenantId(tenantId string) ApiGetEmployersAsyncRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiGetEmployersAsyncRequest) ApiVersion(apiVersion string) ApiGetEmployersAsyncRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiGetEmployersAsyncRequest) XApiVersion(xApiVersion string) ApiGetEmployersAsyncRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiGetEmployersAsyncRequest) Execute() (*EmployerProfileDtoListEnvelope, *http.Response, error) {
	return r.ApiService.GetEmployersAsyncExecute(r)
}

/*
GetEmployersAsync Method for GetEmployersAsync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEmployersAsyncRequest
*/
func (a *EmployersAPIService) GetEmployersAsync(ctx context.Context) ApiGetEmployersAsyncRequest {
	return ApiGetEmployersAsyncRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EmployerProfileDtoListEnvelope
func (a *EmployersAPIService) GetEmployersAsyncExecute(r ApiGetEmployersAsyncRequest) (*EmployerProfileDtoListEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EmployerProfileDtoListEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployersAPIService.GetEmployersAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/HrmsService/Employers"

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

type ApiGetEmployersCountAsyncRequest struct {
	ctx context.Context
	ApiService *EmployersAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiGetEmployersCountAsyncRequest) TenantId(tenantId string) ApiGetEmployersCountAsyncRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiGetEmployersCountAsyncRequest) ApiVersion(apiVersion string) ApiGetEmployersCountAsyncRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiGetEmployersCountAsyncRequest) XApiVersion(xApiVersion string) ApiGetEmployersCountAsyncRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiGetEmployersCountAsyncRequest) Execute() (*Int32Envelope, *http.Response, error) {
	return r.ApiService.GetEmployersCountAsyncExecute(r)
}

/*
GetEmployersCountAsync Method for GetEmployersCountAsync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEmployersCountAsyncRequest
*/
func (a *EmployersAPIService) GetEmployersCountAsync(ctx context.Context) ApiGetEmployersCountAsyncRequest {
	return ApiGetEmployersCountAsyncRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Int32Envelope
func (a *EmployersAPIService) GetEmployersCountAsyncExecute(r ApiGetEmployersCountAsyncRequest) (*Int32Envelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Int32Envelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployersAPIService.GetEmployersCountAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/HrmsService/Employers/Count"

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

type ApiUpdateEmployerAsyncRequest struct {
	ctx context.Context
	ApiService *EmployersAPIService
	tenantId *string
	employerId string
	apiVersion *string
	xApiVersion *string
	body *map[string]interface{}
}

func (r ApiUpdateEmployerAsyncRequest) TenantId(tenantId string) ApiUpdateEmployerAsyncRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiUpdateEmployerAsyncRequest) ApiVersion(apiVersion string) ApiUpdateEmployerAsyncRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiUpdateEmployerAsyncRequest) XApiVersion(xApiVersion string) ApiUpdateEmployerAsyncRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiUpdateEmployerAsyncRequest) Body(body map[string]interface{}) ApiUpdateEmployerAsyncRequest {
	r.body = &body
	return r
}

func (r ApiUpdateEmployerAsyncRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateEmployerAsyncExecute(r)
}

/*
UpdateEmployerAsync Method for UpdateEmployerAsync

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param employerId
 @return ApiUpdateEmployerAsyncRequest
*/
func (a *EmployersAPIService) UpdateEmployerAsync(ctx context.Context, employerId string) ApiUpdateEmployerAsyncRequest {
	return ApiUpdateEmployerAsyncRequest{
		ApiService: a,
		ctx: ctx,
		employerId: employerId,
	}
}

// Execute executes the request
func (a *EmployersAPIService) UpdateEmployerAsyncExecute(r ApiUpdateEmployerAsyncRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployersAPIService.UpdateEmployerAsync")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/HrmsService/Employers/{employerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"employerId"+"}", url.PathEscape(parameterValueToString(r.employerId, "employerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tenantId == nil {
		return nil, reportError("tenantId is required and must be specified")
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
	localVarPostBody = r.body
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
