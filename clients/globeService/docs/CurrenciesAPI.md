# \CurrenciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2GlobeServiceCurrenciesCurrencyIdGet**](CurrenciesAPI.md#ApiV2GlobeServiceCurrenciesCurrencyIdGet) | **Get** /api/v2/GlobeService/Currencies/{currencyId} | 
[**ApiV2GlobeServiceCurrenciesGet**](CurrenciesAPI.md#ApiV2GlobeServiceCurrenciesGet) | **Get** /api/v2/GlobeService/Currencies | 



## ApiV2GlobeServiceCurrenciesCurrencyIdGet

> CurrencyDtoEnvelope ApiV2GlobeServiceCurrenciesCurrencyIdGet(ctx, currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.CurrenciesAPI.ApiV2GlobeServiceCurrenciesCurrencyIdGet(context.Background(), currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesAPI.ApiV2GlobeServiceCurrenciesCurrencyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCurrenciesCurrencyIdGet`: CurrencyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurrenciesAPI.ApiV2GlobeServiceCurrenciesCurrencyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCurrenciesCurrencyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CurrencyDtoEnvelope**](CurrencyDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCurrenciesGet

> CurrencyDtoListEnvelope ApiV2GlobeServiceCurrenciesGet(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.CurrenciesAPI.ApiV2GlobeServiceCurrenciesGet(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesAPI.ApiV2GlobeServiceCurrenciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCurrenciesGet`: CurrencyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurrenciesAPI.ApiV2GlobeServiceCurrenciesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCurrenciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CurrencyDtoListEnvelope**](CurrencyDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

