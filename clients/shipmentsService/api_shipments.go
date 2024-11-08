/*
ShipmentsService

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


// ShipmentsAPIService ShipmentsAPI service
type ShipmentsAPIService service

type ApiApiV2ShipmentsServiceShipmentsGetRequest struct {
	ctx context.Context
	ApiService *ShipmentsAPIService
	tenantId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2ShipmentsServiceShipmentsGetRequest) TenantId(tenantId string) ApiApiV2ShipmentsServiceShipmentsGetRequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiApiV2ShipmentsServiceShipmentsGetRequest) ApiVersion(apiVersion string) ApiApiV2ShipmentsServiceShipmentsGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2ShipmentsServiceShipmentsGetRequest) XApiVersion(xApiVersion string) ApiApiV2ShipmentsServiceShipmentsGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2ShipmentsServiceShipmentsGetRequest) Execute() (*ShipmentDtoListEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2ShipmentsServiceShipmentsGetExecute(r)
}

/*
ApiV2ShipmentsServiceShipmentsGet Method for ApiV2ShipmentsServiceShipmentsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV2ShipmentsServiceShipmentsGetRequest
*/
func (a *ShipmentsAPIService) ApiV2ShipmentsServiceShipmentsGet(ctx context.Context) ApiApiV2ShipmentsServiceShipmentsGetRequest {
	return ApiApiV2ShipmentsServiceShipmentsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ShipmentDtoListEnvelope
func (a *ShipmentsAPIService) ApiV2ShipmentsServiceShipmentsGetExecute(r ApiApiV2ShipmentsServiceShipmentsGetRequest) (*ShipmentDtoListEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShipmentDtoListEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentsAPIService.ApiV2ShipmentsServiceShipmentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/ShipmentsService/Shipments"

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
