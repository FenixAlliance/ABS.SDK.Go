/*
PricingService

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


// PricesAPIService PricesAPI service
type PricesAPIService service

type ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest struct {
	ctx context.Context
	ApiService *PricesAPIService
	itemId string
	currencyId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest) CurrencyId(currencyId string) ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest {
	r.currencyId = &currencyId
	return r
}

func (r ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest) ApiVersion(apiVersion string) ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest) XApiVersion(xApiVersion string) ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest) Execute() (*MoneyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2PricingServicePricesItemIdFinalPriceGetExecute(r)
}

/*
ApiV2PricingServicePricesItemIdFinalPriceGet Method for ApiV2PricingServicePricesItemIdFinalPriceGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId
 @return ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest
*/
func (a *PricesAPIService) ApiV2PricingServicePricesItemIdFinalPriceGet(ctx context.Context, itemId string) ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest {
	return ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return MoneyEnvelope
func (a *PricesAPIService) ApiV2PricingServicePricesItemIdFinalPriceGetExecute(r ApiApiV2PricingServicePricesItemIdFinalPriceGetRequest) (*MoneyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MoneyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesAPIService.ApiV2PricingServicePricesItemIdFinalPriceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/PricingService/Prices/{itemId}/FinalPrice"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.currencyId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currencyId", r.currencyId, "form", "")
	} else {
		var defaultValue string = "USD.USA"
		r.currencyId = &defaultValue
	}
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiApiV2PricingServicePricesItemIdPriceGetRequest struct {
	ctx context.Context
	ApiService *PricesAPIService
	itemId string
	priceListId *string
	discountsListId *string
	currencyId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2PricingServicePricesItemIdPriceGetRequest) PriceListId(priceListId string) ApiApiV2PricingServicePricesItemIdPriceGetRequest {
	r.priceListId = &priceListId
	return r
}

func (r ApiApiV2PricingServicePricesItemIdPriceGetRequest) DiscountsListId(discountsListId string) ApiApiV2PricingServicePricesItemIdPriceGetRequest {
	r.discountsListId = &discountsListId
	return r
}

func (r ApiApiV2PricingServicePricesItemIdPriceGetRequest) CurrencyId(currencyId string) ApiApiV2PricingServicePricesItemIdPriceGetRequest {
	r.currencyId = &currencyId
	return r
}

func (r ApiApiV2PricingServicePricesItemIdPriceGetRequest) ApiVersion(apiVersion string) ApiApiV2PricingServicePricesItemIdPriceGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2PricingServicePricesItemIdPriceGetRequest) XApiVersion(xApiVersion string) ApiApiV2PricingServicePricesItemIdPriceGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2PricingServicePricesItemIdPriceGetRequest) Execute() (*PriceCalculationDtoEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2PricingServicePricesItemIdPriceGetExecute(r)
}

/*
ApiV2PricingServicePricesItemIdPriceGet Method for ApiV2PricingServicePricesItemIdPriceGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId
 @return ApiApiV2PricingServicePricesItemIdPriceGetRequest
*/
func (a *PricesAPIService) ApiV2PricingServicePricesItemIdPriceGet(ctx context.Context, itemId string) ApiApiV2PricingServicePricesItemIdPriceGetRequest {
	return ApiApiV2PricingServicePricesItemIdPriceGetRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return PriceCalculationDtoEnvelope
func (a *PricesAPIService) ApiV2PricingServicePricesItemIdPriceGetExecute(r ApiApiV2PricingServicePricesItemIdPriceGetRequest) (*PriceCalculationDtoEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PriceCalculationDtoEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesAPIService.ApiV2PricingServicePricesItemIdPriceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/PricingService/Prices/{itemId}/Price"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.priceListId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "priceListId", r.priceListId, "form", "")
	}
	if r.discountsListId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "discountsListId", r.discountsListId, "form", "")
	}
	if r.currencyId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currencyId", r.currencyId, "form", "")
	} else {
		var defaultValue string = "USD.USA"
		r.currencyId = &defaultValue
	}
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest struct {
	ctx context.Context
	ApiService *PricesAPIService
	itemId string
	currencyId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest) CurrencyId(currencyId string) ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest {
	r.currencyId = &currencyId
	return r
}

func (r ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest) ApiVersion(apiVersion string) ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest) XApiVersion(xApiVersion string) ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest) Execute() (*MoneyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2PricingServicePricesItemIdTotalSavingsGetExecute(r)
}

/*
ApiV2PricingServicePricesItemIdTotalSavingsGet Method for ApiV2PricingServicePricesItemIdTotalSavingsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId
 @return ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest
*/
func (a *PricesAPIService) ApiV2PricingServicePricesItemIdTotalSavingsGet(ctx context.Context, itemId string) ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest {
	return ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return MoneyEnvelope
func (a *PricesAPIService) ApiV2PricingServicePricesItemIdTotalSavingsGetExecute(r ApiApiV2PricingServicePricesItemIdTotalSavingsGetRequest) (*MoneyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MoneyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesAPIService.ApiV2PricingServicePricesItemIdTotalSavingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/PricingService/Prices/{itemId}/TotalSavings"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.currencyId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currencyId", r.currencyId, "form", "")
	} else {
		var defaultValue string = "USD.USA"
		r.currencyId = &defaultValue
	}
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest struct {
	ctx context.Context
	ApiService *PricesAPIService
	itemId string
	currencyId *string
	apiVersion *string
	xApiVersion *string
}

func (r ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest) CurrencyId(currencyId string) ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest {
	r.currencyId = &currencyId
	return r
}

func (r ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest) ApiVersion(apiVersion string) ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest {
	r.apiVersion = &apiVersion
	return r
}

func (r ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest) XApiVersion(xApiVersion string) ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest {
	r.xApiVersion = &xApiVersion
	return r
}

func (r ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest) Execute() (*MoneyEnvelope, *http.Response, error) {
	return r.ApiService.ApiV2PricingServicePricesItemIdTotalTaxesGetExecute(r)
}

/*
ApiV2PricingServicePricesItemIdTotalTaxesGet Method for ApiV2PricingServicePricesItemIdTotalTaxesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId
 @return ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest
*/
func (a *PricesAPIService) ApiV2PricingServicePricesItemIdTotalTaxesGet(ctx context.Context, itemId string) ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest {
	return ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return MoneyEnvelope
func (a *PricesAPIService) ApiV2PricingServicePricesItemIdTotalTaxesGetExecute(r ApiApiV2PricingServicePricesItemIdTotalTaxesGetRequest) (*MoneyEnvelope, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MoneyEnvelope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PricesAPIService.ApiV2PricingServicePricesItemIdTotalTaxesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/PricingService/Prices/{itemId}/TotalTaxes"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.currencyId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currencyId", r.currencyId, "form", "")
	} else {
		var defaultValue string = "USD.USA"
		r.currencyId = &defaultValue
	}
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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