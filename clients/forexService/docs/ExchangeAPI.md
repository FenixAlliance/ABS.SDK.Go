# \ExchangeAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExchangeAmountAsync**](ExchangeAPI.md#ExchangeAmountAsync) | **Get** /api/v2/ForexService/Exchange/Latest | Exchange currency at latest rates
[**ExchangeAmountHistoricalAsync**](ExchangeAPI.md#ExchangeAmountHistoricalAsync) | **Get** /api/v2/ForexService/Exchange/History | Exchange currency at historical rates



## ExchangeAmountAsync

> MoneyEnvelope ExchangeAmountAsync(ctx).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Execute()

Exchange currency at latest rates



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
	resp, r, err := apiClient.ExchangeAPI.ExchangeAmountAsync(context.Background()).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeAPI.ExchangeAmountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeAmountAsync`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExchangeAPI.ExchangeAmountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeAmountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float64** |  | 
 **sourceCurrencyId** | **string** |  | 
 **targetCurrencyId** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExchangeAmountHistoricalAsync

> MoneyEnvelope ExchangeAmountHistoricalAsync(ctx).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Date(date).Execute()

Exchange currency at historical rates



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
	resp, r, err := apiClient.ExchangeAPI.ExchangeAmountHistoricalAsync(context.Background()).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeAPI.ExchangeAmountHistoricalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeAmountHistoricalAsync`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExchangeAPI.ExchangeAmountHistoricalAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeAmountHistoricalAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float64** |  | 
 **sourceCurrencyId** | **string** |  | 
 **targetCurrencyId** | **string** |  | 
 **date** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

