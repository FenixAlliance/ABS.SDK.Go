# \ItemAttachmentsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemAttachmentAsync**](ItemAttachmentsAPI.md#CreateItemAttachmentAsync) | **Post** /api/v2/CatalogService/ItemAttachments | Create a new item attachment
[**DeleteItemAttachmentAsync**](ItemAttachmentsAPI.md#DeleteItemAttachmentAsync) | **Delete** /api/v2/CatalogService/ItemAttachments/{itemAttachmentId} | Delete an item attachment
[**GetItemAttachmentByIdAsync**](ItemAttachmentsAPI.md#GetItemAttachmentByIdAsync) | **Get** /api/v2/CatalogService/ItemAttachments/{itemAttachmentId} | Get item attachment by ID
[**GetItemAttachmentsAsync**](ItemAttachmentsAPI.md#GetItemAttachmentsAsync) | **Get** /api/v2/CatalogService/ItemAttachments | Get all item attachments
[**UpdateItemAttachmentAsync**](ItemAttachmentsAPI.md#UpdateItemAttachmentAsync) | **Put** /api/v2/CatalogService/ItemAttachments/{itemAttachmentId} | Update an item attachment



## CreateItemAttachmentAsync

> ItemAttachmentDtoEnvelope CreateItemAttachmentAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemAttachmentCreateDto(itemAttachmentCreateDto).Execute()

Create a new item attachment



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
	itemAttachmentCreateDto := *openapiclient.NewItemAttachmentCreateDto() // ItemAttachmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttachmentsAPI.CreateItemAttachmentAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemAttachmentCreateDto(itemAttachmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttachmentsAPI.CreateItemAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemAttachmentAsync`: ItemAttachmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttachmentsAPI.CreateItemAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemAttachmentCreateDto** | [**ItemAttachmentCreateDto**](ItemAttachmentCreateDto.md) |  | 

### Return type

[**ItemAttachmentDtoEnvelope**](ItemAttachmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItemAttachmentAsync

> EmptyEnvelope DeleteItemAttachmentAsync(ctx, itemAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item attachment



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
	itemAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttachmentsAPI.DeleteItemAttachmentAsync(context.Background(), itemAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttachmentsAPI.DeleteItemAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteItemAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttachmentsAPI.DeleteItemAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## GetItemAttachmentByIdAsync

> ItemAttachmentDtoEnvelope GetItemAttachmentByIdAsync(ctx, itemAttachmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item attachment by ID



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
	itemAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttachmentsAPI.GetItemAttachmentByIdAsync(context.Background(), itemAttachmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttachmentsAPI.GetItemAttachmentByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemAttachmentByIdAsync`: ItemAttachmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttachmentsAPI.GetItemAttachmentByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemAttachmentByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemAttachmentDtoEnvelope**](ItemAttachmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemAttachmentsAsync

> ItemAttachmentDtoListEnvelope GetItemAttachmentsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all item attachments



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
	resp, r, err := apiClient.ItemAttachmentsAPI.GetItemAttachmentsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttachmentsAPI.GetItemAttachmentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemAttachmentsAsync`: ItemAttachmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttachmentsAPI.GetItemAttachmentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemAttachmentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemAttachmentDtoListEnvelope**](ItemAttachmentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItemAttachmentAsync

> EmptyEnvelope UpdateItemAttachmentAsync(ctx, itemAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemAttachmentUpdateDto(itemAttachmentUpdateDto).Execute()

Update an item attachment



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
	itemAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemAttachmentUpdateDto := *openapiclient.NewItemAttachmentUpdateDto() // ItemAttachmentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttachmentsAPI.UpdateItemAttachmentAsync(context.Background(), itemAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemAttachmentUpdateDto(itemAttachmentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttachmentsAPI.UpdateItemAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttachmentsAPI.UpdateItemAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemAttachmentUpdateDto** | [**ItemAttachmentUpdateDto**](ItemAttachmentUpdateDto.md) |  | 

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

