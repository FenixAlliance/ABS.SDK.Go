# \RatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ForexServiceRatesHistoryCurrencyIdGet**](RatesAPI.md#ApiV2ForexServiceRatesHistoryCurrencyIdGet) | **Get** /api/v2/ForexService/Rates/History/{currencyId} | 
[**ApiV2ForexServiceRatesHistoryGet**](RatesAPI.md#ApiV2ForexServiceRatesHistoryGet) | **Get** /api/v2/ForexService/Rates/History | 
[**ApiV2ForexServiceRatesLatestCurrencyIdGet**](RatesAPI.md#ApiV2ForexServiceRatesLatestCurrencyIdGet) | **Get** /api/v2/ForexService/Rates/Latest/{currencyId} | 
[**ApiV2ForexServiceRatesLatestGet**](RatesAPI.md#ApiV2ForexServiceRatesLatestGet) | **Get** /api/v2/ForexService/Rates/Latest | 



## ApiV2ForexServiceRatesHistoryCurrencyIdGet

> ExchangeRateEnvelope ApiV2ForexServiceRatesHistoryCurrencyIdGet(ctx, currencyId).Date(date).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	currencyId := "currencyId_example" // string | 
	date := time.Now() // time.Time |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatesAPI.ApiV2ForexServiceRatesHistoryCurrencyIdGet(context.Background(), currencyId).Date(date).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.ApiV2ForexServiceRatesHistoryCurrencyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ForexServiceRatesHistoryCurrencyIdGet`: ExchangeRateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.ApiV2ForexServiceRatesHistoryCurrencyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ForexServiceRatesHistoryCurrencyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **time.Time** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExchangeRateEnvelope**](ExchangeRateEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ForexServiceRatesHistoryGet

> ForexRatesDtoEnvelope ApiV2ForexServiceRatesHistoryGet(ctx).Date(date).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	date := time.Now() // time.Time |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatesAPI.ApiV2ForexServiceRatesHistoryGet(context.Background()).Date(date).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.ApiV2ForexServiceRatesHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ForexServiceRatesHistoryGet`: ForexRatesDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.ApiV2ForexServiceRatesHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ForexServiceRatesHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ForexRatesDtoEnvelope**](ForexRatesDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ForexServiceRatesLatestCurrencyIdGet

> ExchangeRateEnvelope ApiV2ForexServiceRatesLatestCurrencyIdGet(ctx, currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	currencyId := "currencyId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatesAPI.ApiV2ForexServiceRatesLatestCurrencyIdGet(context.Background(), currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.ApiV2ForexServiceRatesLatestCurrencyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ForexServiceRatesLatestCurrencyIdGet`: ExchangeRateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.ApiV2ForexServiceRatesLatestCurrencyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ForexServiceRatesLatestCurrencyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExchangeRateEnvelope**](ExchangeRateEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ForexServiceRatesLatestGet

> ForexRatesDtoEnvelope ApiV2ForexServiceRatesLatestGet(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatesAPI.ApiV2ForexServiceRatesLatestGet(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.ApiV2ForexServiceRatesLatestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ForexServiceRatesLatestGet`: ForexRatesDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.ApiV2ForexServiceRatesLatestGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ForexServiceRatesLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ForexRatesDtoEnvelope**](ForexRatesDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

