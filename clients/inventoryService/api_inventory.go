/*
InventoryService

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


// InventoryAPIService InventoryAPI service
type InventoryAPIService service

type ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	stockItemId string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest) ApiVersion(apiVersion string) ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest) XApiVersion(xApiVersion string) ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV2InventoryServiceInventoryStockItemIdDetailsGetExecute(r)
}

/*
ApiV2InventoryServiceInventoryStockItemIdDetailsGet Method for ApiV2InventoryServiceInventoryStockItemIdDetailsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stockItemId
 @return ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest
*/
func (a *InventoryAPIService) ApiV2InventoryServiceInventoryStockItemIdDetailsGet(ctx context.Context, stockItemId string) ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest {
	return ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest{
		ApiService: a,
		ctx: ctx,
		stockItemId: stockItemId,
	}
}

// Execute executes the request
func (a *InventoryAPIService) ApiV2InventoryServiceInventoryStockItemIdDetailsGetExecute(r ApiApiV2InventoryServiceInventoryStockItemIdDetailsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.ApiV2InventoryServiceInventoryStockItemIdDetailsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/InventoryService/Inventory/{stockItemId}/Details"
	localVarPath = strings.Replace(localVarPath, "{"+"stockItemId"+"}", url.PathEscape(parameterValueToString(r.stockItemId, "stockItemId")), -1)

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