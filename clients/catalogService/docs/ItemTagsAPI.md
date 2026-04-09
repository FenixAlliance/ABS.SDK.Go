# \ItemTagsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemTagAsync**](ItemTagsAPI.md#CreateItemTagAsync) | **Post** /api/v2/CatalogService/ItemTags | Create a new item tag
[**DeleteItemTagAsync**](ItemTagsAPI.md#DeleteItemTagAsync) | **Delete** /api/v2/CatalogService/ItemTags/{itemTagId} | Delete an item tag
[**GetItemTagByIdAsync**](ItemTagsAPI.md#GetItemTagByIdAsync) | **Get** /api/v2/CatalogService/ItemTags/{itemTagId} | Get item tag by ID
[**GetItemTagsAsync**](ItemTagsAPI.md#GetItemTagsAsync) | **Get** /api/v2/CatalogService/ItemTags | Get all item tags
[**UpdateItemTagAsync**](ItemTagsAPI.md#UpdateItemTagAsync) | **Put** /api/v2/CatalogService/ItemTags/{itemTagId} | Update an item tag



## CreateItemTagAsync

> ItemTagDtoEnvelope CreateItemTagAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemTagCreateDto(itemTagCreateDto).Execute()

Create a new item tag



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
	itemTagCreateDto := *openapiclient.NewItemTagCreateDto("Title_example", "BusinessID_example") // ItemTagCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemTagsAPI.CreateItemTagAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemTagCreateDto(itemTagCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTagsAPI.CreateItemTagAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemTagAsync`: ItemTagDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemTagsAPI.CreateItemTagAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemTagAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemTagCreateDto** | [**ItemTagCreateDto**](ItemTagCreateDto.md) |  | 

### Return type

[**ItemTagDtoEnvelope**](ItemTagDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItemTagAsync

> DeleteItemTagAsync(ctx, itemTagId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item tag



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
	itemTagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemTagsAPI.DeleteItemTagAsync(context.Background(), itemTagId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTagsAPI.DeleteItemTagAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemTagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemTagAsyncRequest struct via the builder pattern


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


## GetItemTagByIdAsync

> ItemTagDtoEnvelope GetItemTagByIdAsync(ctx, itemTagId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item tag by ID



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
	itemTagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemTagsAPI.GetItemTagByIdAsync(context.Background(), itemTagId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTagsAPI.GetItemTagByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemTagByIdAsync`: ItemTagDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemTagsAPI.GetItemTagByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemTagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemTagByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemTagDtoEnvelope**](ItemTagDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemTagsAsync

> ItemTagDtoListEnvelope GetItemTagsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all item tags



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
	resp, r, err := apiClient.ItemTagsAPI.GetItemTagsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTagsAPI.GetItemTagsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemTagsAsync`: ItemTagDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemTagsAPI.GetItemTagsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemTagsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemTagDtoListEnvelope**](ItemTagDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItemTagAsync

> UpdateItemTagAsync(ctx, itemTagId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemTagUpdateDto(itemTagUpdateDto).Execute()

Update an item tag



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
	itemTagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemTagUpdateDto := *openapiclient.NewItemTagUpdateDto("Title_example") // ItemTagUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemTagsAPI.UpdateItemTagAsync(context.Background(), itemTagId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemTagUpdateDto(itemTagUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTagsAPI.UpdateItemTagAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemTagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemTagAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemTagUpdateDto** | [**ItemTagUpdateDto**](ItemTagUpdateDto.md) |  | 

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

