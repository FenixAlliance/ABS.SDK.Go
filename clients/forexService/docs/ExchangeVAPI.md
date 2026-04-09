# \ExchangeVAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExchangeAmountHistoricalV3Async**](ExchangeVAPI.md#ExchangeAmountHistoricalV3Async) | **Get** /api/v3/ForexService/Exchange/History | Exchange currency at historical rates (v3)
[**ExchangeAmountV3Async**](ExchangeVAPI.md#ExchangeAmountV3Async) | **Get** /api/v3/ForexService/Exchange/Latest | Exchange currency at latest rates (v3)



## ExchangeAmountHistoricalV3Async

> ExchangeRateEnvelope ExchangeAmountHistoricalV3Async(ctx).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Date(date).Execute()

Exchange currency at historical rates (v3)



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
	amount := float64(1.2) // float64 | 
	sourceCurrencyId := "sourceCurrencyId_example" // string | 
	targetCurrencyId := "targetCurrencyId_example" // string | 
	date := time.Now() // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangeVAPI.ExchangeAmountHistoricalV3Async(context.Background()).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeVAPI.ExchangeAmountHistoricalV3Async``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeAmountHistoricalV3Async`: ExchangeRateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExchangeVAPI.ExchangeAmountHistoricalV3Async`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeAmountHistoricalV3AsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float64** |  | 
 **sourceCurrencyId** | **string** |  | 
 **targetCurrencyId** | **string** |  | 
 **date** | **string** |  | 

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


## ExchangeAmountV3Async

> ExchangeRateEnvelope ExchangeAmountV3Async(ctx).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Execute()

Exchange currency at latest rates (v3)



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
	amount := float64(1.2) // float64 | 
	sourceCurrencyId := "sourceCurrencyId_example" // string | 
	targetCurrencyId := "targetCurrencyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExchangeVAPI.ExchangeAmountV3Async(context.Background()).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeVAPI.ExchangeAmountV3Async``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeAmountV3Async`: ExchangeRateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExchangeVAPI.ExchangeAmountV3Async`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeAmountV3AsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float64** |  | 
 **sourceCurrencyId** | **string** |  | 
 **targetCurrencyId** | **string** |  | 

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

