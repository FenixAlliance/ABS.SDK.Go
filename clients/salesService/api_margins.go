/*
SalesService

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


// MarginsAPIService MarginsAPI service
type MarginsAPIService service

type ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest struct {
	ctx context.Context
	ApiService *MarginsAPIService
	marginId string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest) ApiVersion(apiVersion string) ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest) XApiVersion(xApiVersion string) ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV2SalesServiceMarginsMarginIdDetailsGetExecute(r)
}

/*
ApiV2SalesServiceMarginsMarginIdDetailsGet Method for ApiV2SalesServiceMarginsMarginIdDetailsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param marginId
 @return ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest
*/
func (a *MarginsAPIService) ApiV2SalesServiceMarginsMarginIdDetailsGet(ctx context.Context, marginId string) ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest {
	return ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest{
		ApiService: a,
		ctx: ctx,
		marginId: marginId,
	}
}

// Execute executes the request
func (a *MarginsAPIService) ApiV2SalesServiceMarginsMarginIdDetailsGetExecute(r ApiApiV2SalesServiceMarginsMarginIdDetailsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarginsAPIService.ApiV2SalesServiceMarginsMarginIdDetailsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/SalesService/Margins/{marginId}/Details"
	localVarPath = strings.Replace(localVarPath, "{"+"marginId"+"}", url.PathEscape(parameterValueToString(r.marginId, "marginId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
