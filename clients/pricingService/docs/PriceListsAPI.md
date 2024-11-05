# \PriceListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2PricingServicePriceListsCountGet**](PriceListsAPI.md#ApiV2PricingServicePriceListsCountGet) | **Get** /api/v2/PricingService/PriceLists/Count | 
[**ApiV2PricingServicePriceListsGet**](PriceListsAPI.md#ApiV2PricingServicePriceListsGet) | **Get** /api/v2/PricingService/PriceLists | 
[**ApiV2PricingServicePriceListsPost**](PriceListsAPI.md#ApiV2PricingServicePriceListsPost) | **Post** /api/v2/PricingService/PriceLists | 
[**ApiV2PricingServicePriceListsPriceListIdDelete**](PriceListsAPI.md#ApiV2PricingServicePriceListsPriceListIdDelete) | **Delete** /api/v2/PricingService/PriceLists/{priceListId} | 
[**ApiV2PricingServicePriceListsPriceListIdPricesPost**](PriceListsAPI.md#ApiV2PricingServicePriceListsPriceListIdPricesPost) | **Post** /api/v2/PricingService/PriceLists/{priceListId}/Prices | 
[**ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete**](PriceListsAPI.md#ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete) | **Delete** /api/v2/PricingService/PriceLists/{priceListId}/Prices/{priceId} | 
[**ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut**](PriceListsAPI.md#ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut) | **Put** /api/v2/PricingService/PriceLists/{priceListId}/Prices/{priceId} | 
[**ApiV2PricingServicePriceListsPriceListIdPut**](PriceListsAPI.md#ApiV2PricingServicePriceListsPriceListIdPut) | **Put** /api/v2/PricingService/PriceLists/{priceListId} | 
[**GetPriceListAsync**](PriceListsAPI.md#GetPriceListAsync) | **Get** /api/v2/PricingService/PriceLists/{priceListId} | 
[**GetPriceListPriceAsync**](PriceListsAPI.md#GetPriceListPriceAsync) | **Get** /api/v2/PricingService/PriceLists/{priceListId}/Prices/{priceId} | 
[**GetPriceListPricesAsync**](PriceListsAPI.md#GetPriceListPricesAsync) | **Get** /api/v2/PricingService/PriceLists/{priceListId}/Prices | 



## ApiV2PricingServicePriceListsCountGet

> Int32Envelope ApiV2PricingServicePriceListsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.ApiV2PricingServicePriceListsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePriceListsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.ApiV2PricingServicePriceListsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePriceListsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServicePriceListsGet

> PriceListDtoListEnvelope ApiV2PricingServicePriceListsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.ApiV2PricingServicePriceListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePriceListsGet`: PriceListDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.ApiV2PricingServicePriceListsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePriceListsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PriceListDtoListEnvelope**](PriceListDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServicePriceListsPost

> EmptyEnvelope ApiV2PricingServicePriceListsPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PriceListCreateDto(priceListCreateDto).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	priceListCreateDto := *openapiclient.NewPriceListCreateDto("Name_example") // PriceListCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PriceListCreateDto(priceListCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.ApiV2PricingServicePriceListsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePriceListsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.ApiV2PricingServicePriceListsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePriceListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **priceListCreateDto** | [**PriceListCreateDto**](PriceListCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServicePriceListsPriceListIdDelete

> EmptyEnvelope ApiV2PricingServicePriceListsPriceListIdDelete(ctx, priceListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdDelete(context.Background(), priceListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePriceListsPriceListIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePriceListsPriceListIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServicePriceListsPriceListIdPricesPost

> EmptyEnvelope ApiV2PricingServicePriceListsPriceListIdPricesPost(ctx, priceListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPriceCreateDto(itemPriceCreateDto).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemPriceCreateDto := *openapiclient.NewItemPriceCreateDto("ItemId_example") // ItemPriceCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPost(context.Background(), priceListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPriceCreateDto(itemPriceCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePriceListsPriceListIdPricesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePriceListsPriceListIdPricesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPriceCreateDto** | [**ItemPriceCreateDto**](ItemPriceCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete

> EmptyEnvelope ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete(ctx, priceListId, priceId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete(context.Background(), priceListId, priceId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPriceIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePriceListsPriceListIdPricesPriceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut

> EmptyEnvelope ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut(ctx, priceListId, priceId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPriceUpdateDto(itemPriceUpdateDto).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemPriceUpdateDto := *openapiclient.NewItemPriceUpdateDto() // ItemPriceUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut(context.Background(), priceListId, priceId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPriceUpdateDto(itemPriceUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPricesPriceIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePriceListsPriceListIdPricesPriceIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPriceUpdateDto** | [**ItemPriceUpdateDto**](ItemPriceUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServicePriceListsPriceListIdPut

> EmptyEnvelope ApiV2PricingServicePriceListsPriceListIdPut(ctx, priceListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PriceListUpdateDto(priceListUpdateDto).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	priceListUpdateDto := *openapiclient.NewPriceListUpdateDto("Name_example") // PriceListUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPut(context.Background(), priceListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PriceListUpdateDto(priceListUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServicePriceListsPriceListIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.ApiV2PricingServicePriceListsPriceListIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServicePriceListsPriceListIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **priceListUpdateDto** | [**PriceListUpdateDto**](PriceListUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListAsync

> PriceListDtoEnvelope GetPriceListAsync(ctx, priceListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListAsync(context.Background(), priceListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.GetPriceListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceListAsync`: PriceListDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.GetPriceListAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PriceListDtoEnvelope**](PriceListDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListPriceAsync

> ItemPriceDtoEnvelope GetPriceListPriceAsync(ctx, priceListId, priceId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListPriceAsync(context.Background(), priceListId, priceId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.GetPriceListPriceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceListPriceAsync`: ItemPriceDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.GetPriceListPriceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListPriceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPriceDtoEnvelope**](ItemPriceDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListPricesAsync

> ItemPriceDtoListEnvelope GetPriceListPricesAsync(ctx, priceListId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	priceListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListPricesAsync(context.Background(), priceListId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.GetPriceListPricesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceListPricesAsync`: ItemPriceDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.GetPriceListPricesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListPricesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPriceDtoListEnvelope**](ItemPriceDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

