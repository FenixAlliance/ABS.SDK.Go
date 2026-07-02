# \RatesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistoricalCurrencyRateAsync**](RatesAPI.md#GetHistoricalCurrencyRateAsync) | **Get** /api/v2/ForexService/Rates/History/{currencyId} | Get historical rate for a currency
[**GetHistoricalCurrencyRatesAsync**](RatesAPI.md#GetHistoricalCurrencyRatesAsync) | **Get** /api/v2/ForexService/Rates/History | Get historical currency rates
[**GetLatestCurrencyRateAsync**](RatesAPI.md#GetLatestCurrencyRateAsync) | **Get** /api/v2/ForexService/Rates/Latest/{currencyId} | Get latest rate for a currency
[**GetLatestCurrencyRatesModelAsync**](RatesAPI.md#GetLatestCurrencyRatesModelAsync) | **Get** /api/v2/ForexService/Rates/Latest | Get latest currency rates



## GetHistoricalCurrencyRateAsync

> ExchangeRateEnvelope GetHistoricalCurrencyRateAsync(ctx, currencyId).Date(date).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get historical rate for a currency



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
	resp, r, err := apiClient.RatesAPI.GetHistoricalCurrencyRateAsync(context.Background(), currencyId).Date(date).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.GetHistoricalCurrencyRateAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricalCurrencyRateAsync`: ExchangeRateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.GetHistoricalCurrencyRateAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalCurrencyRateAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **time.Time** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExchangeRateEnvelope**](ExchangeRateEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHistoricalCurrencyRatesAsync

> ForexRatesDtoEnvelope GetHistoricalCurrencyRatesAsync(ctx).Date(date).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get historical currency rates



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
	resp, r, err := apiClient.RatesAPI.GetHistoricalCurrencyRatesAsync(context.Background()).Date(date).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.GetHistoricalCurrencyRatesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistoricalCurrencyRatesAsync`: ForexRatesDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.GetHistoricalCurrencyRatesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoricalCurrencyRatesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ForexRatesDtoEnvelope**](ForexRatesDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestCurrencyRateAsync

> ExchangeRateEnvelope GetLatestCurrencyRateAsync(ctx, currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get latest rate for a currency



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
	resp, r, err := apiClient.RatesAPI.GetLatestCurrencyRateAsync(context.Background(), currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.GetLatestCurrencyRateAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestCurrencyRateAsync`: ExchangeRateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.GetLatestCurrencyRateAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestCurrencyRateAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExchangeRateEnvelope**](ExchangeRateEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestCurrencyRatesModelAsync

> ForexRatesDtoEnvelope GetLatestCurrencyRatesModelAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get latest currency rates



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
	resp, r, err := apiClient.RatesAPI.GetLatestCurrencyRatesModelAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.GetLatestCurrencyRatesModelAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestCurrencyRatesModelAsync`: ForexRatesDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.GetLatestCurrencyRatesModelAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestCurrencyRatesModelAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ForexRatesDtoEnvelope**](ForexRatesDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

