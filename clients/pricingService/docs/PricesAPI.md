# \PricesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFinalPrice**](PricesAPI.md#GetFinalPrice) | **Get** /api/v2/PricingService/Prices/{itemId}/FinalPrice | Gets the final price for an item
[**GetPrice**](PricesAPI.md#GetPrice) | **Get** /api/v2/PricingService/Prices/{itemId}/Price | Gets the calculated price for an item
[**GetTotalSavingsInUsd**](PricesAPI.md#GetTotalSavingsInUsd) | **Get** /api/v2/PricingService/Prices/{itemId}/TotalSavings | Gets total savings for an item
[**GetTotalTaxesInUsd**](PricesAPI.md#GetTotalTaxesInUsd) | **Get** /api/v2/PricingService/Prices/{itemId}/TotalTaxes | Gets total taxes for an item



## GetFinalPrice

> MoneyEnvelope GetFinalPrice(ctx, itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the final price for an item



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
	resp, r, err := apiClient.PricesAPI.GetFinalPrice(context.Background(), itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.GetFinalPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinalPrice`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.GetFinalPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinalPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## GetPrice

> ItemPriceCalculationEnvelope GetPrice(ctx, itemId).PriceListId(priceListId).DiscountsListId(discountsListId).Quantity(quantity).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the calculated price for an item



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
	quantity := float64(1.2) // float64 |  (optional) (default to 1)
	currencyId := "currencyId_example" // string |  (optional) (default to "USD.USA")
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.GetPrice(context.Background(), itemId).PriceListId(priceListId).DiscountsListId(discountsListId).Quantity(quantity).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.GetPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrice`: ItemPriceCalculationEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.GetPrice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **priceListId** | **string** |  | 
 **discountsListId** | **string** |  | 
 **quantity** | **float64** |  | [default to 1]
 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPriceCalculationEnvelope**](ItemPriceCalculationEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTotalSavingsInUsd

> MoneyEnvelope GetTotalSavingsInUsd(ctx, itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets total savings for an item



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
	resp, r, err := apiClient.PricesAPI.GetTotalSavingsInUsd(context.Background(), itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.GetTotalSavingsInUsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalSavingsInUsd`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.GetTotalSavingsInUsd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalSavingsInUsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## GetTotalTaxesInUsd

> MoneyEnvelope GetTotalTaxesInUsd(ctx, itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets total taxes for an item



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
	resp, r, err := apiClient.PricesAPI.GetTotalTaxesInUsd(context.Background(), itemId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.GetTotalTaxesInUsd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTotalTaxesInUsd`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.GetTotalTaxesInUsd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTotalTaxesInUsdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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

