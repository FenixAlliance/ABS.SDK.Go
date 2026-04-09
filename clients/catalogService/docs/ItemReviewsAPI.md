# \ItemReviewsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemReviewAsync**](ItemReviewsAPI.md#CreateItemReviewAsync) | **Post** /api/v2/CatalogService/ItemReviews | Create a new item review
[**DeleteItemReviewAsync**](ItemReviewsAPI.md#DeleteItemReviewAsync) | **Delete** /api/v2/CatalogService/ItemReviews/{itemReviewId} | Delete an item review
[**GetItemReviewByIdAsync**](ItemReviewsAPI.md#GetItemReviewByIdAsync) | **Get** /api/v2/CatalogService/ItemReviews/{itemReviewId} | Get item review by ID
[**GetItemReviewsAsync**](ItemReviewsAPI.md#GetItemReviewsAsync) | **Get** /api/v2/CatalogService/ItemReviews | Get all item reviews
[**UpdateItemReviewAsync**](ItemReviewsAPI.md#UpdateItemReviewAsync) | **Put** /api/v2/CatalogService/ItemReviews/{itemReviewId} | Update an item review



## CreateItemReviewAsync

> ItemReviewDtoEnvelope CreateItemReviewAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemReviewCreateDto(itemReviewCreateDto).Execute()

Create a new item review



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
	itemReviewCreateDto := *openapiclient.NewItemReviewCreateDto() // ItemReviewCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemReviewsAPI.CreateItemReviewAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemReviewCreateDto(itemReviewCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReviewsAPI.CreateItemReviewAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemReviewAsync`: ItemReviewDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemReviewsAPI.CreateItemReviewAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemReviewAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemReviewCreateDto** | [**ItemReviewCreateDto**](ItemReviewCreateDto.md) |  | 

### Return type

[**ItemReviewDtoEnvelope**](ItemReviewDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItemReviewAsync

> DeleteItemReviewAsync(ctx, itemReviewId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item review



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
	itemReviewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemReviewsAPI.DeleteItemReviewAsync(context.Background(), itemReviewId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReviewsAPI.DeleteItemReviewAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemReviewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemReviewAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemReviewByIdAsync

> ItemReviewDtoEnvelope GetItemReviewByIdAsync(ctx, itemReviewId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item review by ID



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
	itemReviewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemReviewsAPI.GetItemReviewByIdAsync(context.Background(), itemReviewId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReviewsAPI.GetItemReviewByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemReviewByIdAsync`: ItemReviewDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemReviewsAPI.GetItemReviewByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemReviewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemReviewByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemReviewDtoEnvelope**](ItemReviewDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemReviewsAsync

> ItemReviewDtoListEnvelope GetItemReviewsAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all item reviews



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemReviewsAPI.GetItemReviewsAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReviewsAPI.GetItemReviewsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemReviewsAsync`: ItemReviewDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemReviewsAPI.GetItemReviewsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemReviewsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemReviewDtoListEnvelope**](ItemReviewDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItemReviewAsync

> UpdateItemReviewAsync(ctx, itemReviewId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemReviewUpdateDto(itemReviewUpdateDto).Execute()

Update an item review



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
	itemReviewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemReviewUpdateDto := *openapiclient.NewItemReviewUpdateDto() // ItemReviewUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemReviewsAPI.UpdateItemReviewAsync(context.Background(), itemReviewId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemReviewUpdateDto(itemReviewUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReviewsAPI.UpdateItemReviewAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemReviewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemReviewAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemReviewUpdateDto** | [**ItemReviewUpdateDto**](ItemReviewUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

