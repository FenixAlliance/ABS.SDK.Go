# \PricesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2PricingServicePricesItemIdFinalPriceGet**](PricesAPI.md#ApiV2PricingServicePricesItemIdFinalPriceGet) | **Get** /api/v2/PricingService/Prices/{itemId}/FinalPrice | 
[**ApiV2PricingServicePricesItemIdPriceGet**](PricesAPI.md#ApiV2PricingServicePricesItemIdPriceGet) | **Get** /api/v2/PricingService/Prices/{itemId}/Price | 
[**ApiV2PricingServicePricesItemIdTotalSavingsGet**](PricesAPI.md#ApiV2PricingServicePricesItemIdTotalSavingsGet) | **Get** /api/v2/PricingService/Prices/{itemId}/TotalSavings | 
[**ApiV2PricingServicePricesItemIdTotalTaxesGet**](PricesAPI.md#ApiV2PricingServicePricesItemIdTotalTaxesGet) | **Get** /api/v2/PricingService/Prices/{itemId}/TotalTaxes | 



## ApiV2PricingServicePricesItemIdFinalPriceGet

> MoneyEnvelope ApiV2PricingServicePricesItemIdFinalPriceGet(ctx, itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	currencyId := "currencyId_example" // string |  (optional) (default to "USD.USA")
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.ApiV2PricingServicePricesItemIdFinalPriceGet(context.Background(), itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.ApiV2PricingServicePricesItemIdFinalPriceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePricesItemIdFinalPriceGet`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.ApiV2PricingServicePricesItemIdFinalPriceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePricesItemIdFinalPriceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2PricingServicePricesItemIdPriceGet

> PriceCalculationDtoEnvelope ApiV2PricingServicePricesItemIdPriceGet(ctx, itemId).PriceListId(priceListId).DiscountsListId(discountsListId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	discountsListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	currencyId := "currencyId_example" // string |  (optional) (default to "USD.USA")
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.ApiV2PricingServicePricesItemIdPriceGet(context.Background(), itemId).PriceListId(priceListId).DiscountsListId(discountsListId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.ApiV2PricingServicePricesItemIdPriceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePricesItemIdPriceGet`: PriceCalculationDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.ApiV2PricingServicePricesItemIdPriceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePricesItemIdPriceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **priceListId** | **string** |  | 
 **discountsListId** | **string** |  | 
 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PriceCalculationDtoEnvelope**](PriceCalculationDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServicePricesItemIdTotalSavingsGet

> MoneyEnvelope ApiV2PricingServicePricesItemIdTotalSavingsGet(ctx, itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	currencyId := "currencyId_example" // string |  (optional) (default to "USD.USA")
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.ApiV2PricingServicePricesItemIdTotalSavingsGet(context.Background(), itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.ApiV2PricingServicePricesItemIdTotalSavingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePricesItemIdTotalSavingsGet`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.ApiV2PricingServicePricesItemIdTotalSavingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePricesItemIdTotalSavingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2PricingServicePricesItemIdTotalTaxesGet

> MoneyEnvelope ApiV2PricingServicePricesItemIdTotalTaxesGet(ctx, itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	currencyId := "currencyId_example" // string |  (optional) (default to "USD.USA")
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.ApiV2PricingServicePricesItemIdTotalTaxesGet(context.Background(), itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.ApiV2PricingServicePricesItemIdTotalTaxesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePricesItemIdTotalTaxesGet`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.ApiV2PricingServicePricesItemIdTotalTaxesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePricesItemIdTotalTaxesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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

