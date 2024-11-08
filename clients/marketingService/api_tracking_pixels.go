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


// TrackingPixelsAPIService TrackingPixelsAPI service
type TrackingPixelsAPIService service

type ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest struct {
	ctx context.Context
	ApiService *TrackingPixelsAPIService
	pixelId string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest) ApiVersion(apiVersion string) ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest) XApiVersion(xApiVersion string) ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest) Execute() (*OrderDtoEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2MarketingServiceTrackingPixelsPixelIdGetExecute(r)
}

/*
ApiV2MarketingServiceTrackingPixelsPixelIdGet Method for ApiV2MarketingServiceTrackingPixelsPixelIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pixelId
 @return ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest
*/
func (a *TrackingPixelsAPIService) ApiV2MarketingServiceTrackingPixelsPixelIdGet(ctx context.Context, pixelId string) ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest {
	return ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest{
		ApiService: a,
		ctx: ctx,
		pixelId: pixelId,
	}
}

// Execute executes the request
//  @return OrderDtoEnvelope
func (a *TrackingPixelsAPIService) ApiV2MarketingServiceTrackingPixelsPixelIdGetExecute(r ApiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest) (*OrderDtoEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderDtoEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TrackingPixelsAPIService.ApiV2MarketingServiceTrackingPixelsPixelIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/MarketingService/TrackingPixels/{pixelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pixelId"+"}", url.PathEscape(parameterValueToString(r.pixelId, "pixelId")), -1)

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
