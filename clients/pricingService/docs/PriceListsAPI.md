# \PriceListsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePriceListAsync**](PriceListsAPI.md#CreatePriceListAsync) | **Post** /api/v2/PricingService/PriceLists | Creates a new price list
[**CreatePriceListPricesAsync**](PriceListsAPI.md#CreatePriceListPricesAsync) | **Post** /api/v2/PricingService/PriceLists/{priceListId}/Prices | Creates a price list entry
[**DeletePriceListAsync**](PriceListsAPI.md#DeletePriceListAsync) | **Delete** /api/v2/PricingService/PriceLists/{priceListId} | Deletes a price list
[**DeletePriceListPriceAsync**](PriceListsAPI.md#DeletePriceListPriceAsync) | **Delete** /api/v2/PricingService/PriceLists/{priceListId}/Prices/{priceId} | Deletes a price list entry
[**GetPriceListAsync**](PriceListsAPI.md#GetPriceListAsync) | **Get** /api/v2/PricingService/PriceLists/{priceListId} | Gets a price list by ID
[**GetPriceListPriceAsync**](PriceListsAPI.md#GetPriceListPriceAsync) | **Get** /api/v2/PricingService/PriceLists/{priceListId}/Prices/{priceId} | Gets a price list entry by ID
[**GetPriceListPricesAsync**](PriceListsAPI.md#GetPriceListPricesAsync) | **Get** /api/v2/PricingService/PriceLists/{priceListId}/Prices | Retrieves prices in a price list
[**GetPriceListsAsync**](PriceListsAPI.md#GetPriceListsAsync) | **Get** /api/v2/PricingService/PriceLists | Retrieves all price lists
[**GetPriceListsCountAsync**](PriceListsAPI.md#GetPriceListsCountAsync) | **Get** /api/v2/PricingService/PriceLists/Count | Counts price lists
[**UpdatePriceListAsync**](PriceListsAPI.md#UpdatePriceListAsync) | **Put** /api/v2/PricingService/PriceLists/{priceListId} | Updates a price list
[**UpdatePriceListPriceAsync**](PriceListsAPI.md#UpdatePriceListPriceAsync) | **Put** /api/v2/PricingService/PriceLists/{priceListId}/Prices/{priceId} | Updates a price list entry



## CreatePriceListAsync

> EmptyEnvelope CreatePriceListAsync(ctx).TenantId(tenantId).PriceListCreateDto(priceListCreateDto).Execute()

Creates a new price list



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
	priceListCreateDto := *openapiclient.NewPriceListCreateDto("Name_example") // PriceListCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.CreatePriceListAsync(context.Background()).TenantId(tenantId).PriceListCreateDto(priceListCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.CreatePriceListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePriceListAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.CreatePriceListAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePriceListAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **priceListCreateDto** | [**PriceListCreateDto**](PriceListCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePriceListPricesAsync

> EmptyEnvelope CreatePriceListPricesAsync(ctx, priceListId).TenantId(tenantId).ItemPriceCreateDto(itemPriceCreateDto).Execute()

Creates a price list entry



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
	itemPriceCreateDto := *openapiclient.NewItemPriceCreateDto("ItemId_example") // ItemPriceCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.CreatePriceListPricesAsync(context.Background(), priceListId).TenantId(tenantId).ItemPriceCreateDto(itemPriceCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.CreatePriceListPricesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePriceListPricesAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.CreatePriceListPricesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePriceListPricesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **itemPriceCreateDto** | [**ItemPriceCreateDto**](ItemPriceCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePriceListAsync

> EmptyEnvelope DeletePriceListAsync(ctx, priceListId).TenantId(tenantId).Execute()

Deletes a price list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.DeletePriceListAsync(context.Background(), priceListId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.DeletePriceListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePriceListAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.DeletePriceListAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePriceListAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePriceListPriceAsync

> EmptyEnvelope DeletePriceListPriceAsync(ctx, priceListId, priceId).TenantId(tenantId).Execute()

Deletes a price list entry



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.DeletePriceListPriceAsync(context.Background(), priceListId, priceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.DeletePriceListPriceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePriceListPriceAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.DeletePriceListPriceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePriceListPriceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListAsync

> PriceListDtoEnvelope GetPriceListAsync(ctx, priceListId).TenantId(tenantId).Execute()

Gets a price list by ID



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListAsync(context.Background(), priceListId).TenantId(tenantId).Execute()
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


### Return type

[**PriceListDtoEnvelope**](PriceListDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListPriceAsync

> ItemPriceDtoEnvelope GetPriceListPriceAsync(ctx, priceListId, priceId).TenantId(tenantId).Execute()

Gets a price list entry by ID



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListPriceAsync(context.Background(), priceListId, priceId).TenantId(tenantId).Execute()
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



### Return type

[**ItemPriceDtoEnvelope**](ItemPriceDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListPricesAsync

> ItemPriceDtoListEnvelope GetPriceListPricesAsync(ctx, priceListId).TenantId(tenantId).ItemId(itemId).Execute()

Retrieves prices in a price list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListPricesAsync(context.Background(), priceListId).TenantId(tenantId).ItemId(itemId).Execute()
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

### Return type

[**ItemPriceDtoListEnvelope**](ItemPriceDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListsAsync

> PriceListDtoListEnvelope GetPriceListsAsync(ctx).TenantId(tenantId).Execute()

Retrieves all price lists



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.GetPriceListsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceListsAsync`: PriceListDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.GetPriceListsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**PriceListDtoListEnvelope**](PriceListDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriceListsCountAsync

> Int32Envelope GetPriceListsCountAsync(ctx).TenantId(tenantId).Execute()

Counts price lists



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.GetPriceListsCountAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.GetPriceListsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceListsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.GetPriceListsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceListsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePriceListAsync

> EmptyEnvelope UpdatePriceListAsync(ctx, priceListId).TenantId(tenantId).PriceListUpdateDto(priceListUpdateDto).Execute()

Updates a price list



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
	priceListUpdateDto := *openapiclient.NewPriceListUpdateDto("Name_example") // PriceListUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.UpdatePriceListAsync(context.Background(), priceListId).TenantId(tenantId).PriceListUpdateDto(priceListUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.UpdatePriceListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePriceListAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.UpdatePriceListAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceListAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **priceListUpdateDto** | [**PriceListUpdateDto**](PriceListUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePriceListPriceAsync

> EmptyEnvelope UpdatePriceListPriceAsync(ctx, priceListId, priceId).TenantId(tenantId).ItemPriceUpdateDto(itemPriceUpdateDto).Execute()

Updates a price list entry



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
	itemPriceUpdateDto := *openapiclient.NewItemPriceUpdateDto() // ItemPriceUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceListsAPI.UpdatePriceListPriceAsync(context.Background(), priceListId, priceId).TenantId(tenantId).ItemPriceUpdateDto(itemPriceUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceListsAPI.UpdatePriceListPriceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePriceListPriceAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PriceListsAPI.UpdatePriceListPriceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceListId** | **string** |  | 
**priceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceListPriceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **itemPriceUpdateDto** | [**ItemPriceUpdateDto**](ItemPriceUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

