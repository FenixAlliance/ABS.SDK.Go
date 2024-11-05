# \ExchangeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ForexServiceExchangeHistoryGet**](ExchangeAPI.md#ApiV2ForexServiceExchangeHistoryGet) | **Get** /api/v2/ForexService/Exchange/History | 
[**ApiV2ForexServiceExchangeLatestGet**](ExchangeAPI.md#ApiV2ForexServiceExchangeLatestGet) | **Get** /api/v2/ForexService/Exchange/Latest | 



## ApiV2ForexServiceExchangeHistoryGet

> MoneyEnvelope ApiV2ForexServiceExchangeHistoryGet(ctx).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Date(date).Execute()



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
	resp, r, err := apiClient.ExchangeAPI.ApiV2ForexServiceExchangeHistoryGet(context.Background()).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeAPI.ApiV2ForexServiceExchangeHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ForexServiceExchangeHistoryGet`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExchangeAPI.ApiV2ForexServiceExchangeHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ForexServiceExchangeHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float64** |  | 
 **sourceCurrencyId** | **string** |  | 
 **targetCurrencyId** | **string** |  | 
 **date** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ForexServiceExchangeLatestGet

> MoneyEnvelope ApiV2ForexServiceExchangeLatestGet(ctx).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Execute()



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
	resp, r, err := apiClient.ExchangeAPI.ApiV2ForexServiceExchangeLatestGet(context.Background()).Amount(amount).SourceCurrencyId(sourceCurrencyId).TargetCurrencyId(targetCurrencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExchangeAPI.ApiV2ForexServiceExchangeLatestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ForexServiceExchangeLatestGet`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExchangeAPI.ApiV2ForexServiceExchangeLatestGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ForexServiceExchangeLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amount** | **float64** |  | 
 **sourceCurrencyId** | **string** |  | 
 **targetCurrencyId** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

