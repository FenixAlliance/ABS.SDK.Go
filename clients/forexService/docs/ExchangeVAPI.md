# \ExchangeVAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV3ForexServiceExchangeHistoryGet**](ExchangeVAPI.md#ApiV3ForexServiceExchangeHistoryGet) | **Get** /api/v3/ForexService/Exchange/History | 
[**ApiV3ForexServiceExchangeLatestGet**](ExchangeVAPI.md#ApiV3ForexServiceExchangeLatestGet) | **Get** /api/v3/ForexService/Exchange/Latest | 



## ApiV3ForexServiceExchangeHistoryGet

> ExchangeRateEnvelope ApiV3ForexServiceExchangeHistoryGet(ctx).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Date(date).Execute()



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
	resp, r, err := apiClient.ExchangeVAPI.ApiV3ForexServiceExchangeHistoryGet(context.Background()).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeVAPI.ApiV3ForexServiceExchangeHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV3ForexServiceExchangeHistoryGet`: ExchangeRateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExchangeVAPI.ApiV3ForexServiceExchangeHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ForexServiceExchangeHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float64** |  | 
 **sourceCurrencyId** | **string** |  | 
 **targetCurrencyId** | **string** |  | 
 **date** | **string** |  | 

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


## ApiV3ForexServiceExchangeLatestGet

> ExchangeRateEnvelope ApiV3ForexServiceExchangeLatestGet(ctx).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Execute()



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
	resp, r, err := apiClient.ExchangeVAPI.ApiV3ForexServiceExchangeLatestGet(context.Background()).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeVAPI.ApiV3ForexServiceExchangeLatestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV3ForexServiceExchangeLatestGet`: ExchangeRateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExchangeVAPI.ApiV3ForexServiceExchangeLatestGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV3ForexServiceExchangeLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float64** |  | 
 **sourceCurrencyId** | **string** |  | 
 **targetCurrencyId** | **string** |  | 

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

