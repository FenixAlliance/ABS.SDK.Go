# \ItemAttributeOptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemAttributeOptionAsync**](ItemAttributeOptionsAPI.md#CreateItemAttributeOptionAsync) | **Post** /api/v2/CatalogService/ItemAttributeOptions | Create a new item attribute option
[**DeleteItemAttributeOptionAsync**](ItemAttributeOptionsAPI.md#DeleteItemAttributeOptionAsync) | **Delete** /api/v2/CatalogService/ItemAttributeOptions/{itemAttributeOptionId} | Delete an item attribute option
[**GetItemAttributeOptionByIdAsync**](ItemAttributeOptionsAPI.md#GetItemAttributeOptionByIdAsync) | **Get** /api/v2/CatalogService/ItemAttributeOptions/{itemAttributeOptionId} | Get item attribute option by ID
[**GetItemAttributeOptionsAsync**](ItemAttributeOptionsAPI.md#GetItemAttributeOptionsAsync) | **Get** /api/v2/CatalogService/ItemAttributeOptions | Get all item attribute options
[**GetItemAttributeOptionsCountAsync**](ItemAttributeOptionsAPI.md#GetItemAttributeOptionsCountAsync) | **Get** /api/v2/CatalogService/ItemAttributeOptions/Count | Get item attribute options count
[**PatchItemAttributeOptionAsync**](ItemAttributeOptionsAPI.md#PatchItemAttributeOptionAsync) | **Patch** /api/v2/CatalogService/ItemAttributeOptions/{itemAttributeOptionId} | Patch an item attribute option
[**UpdateItemAttributeOptionAsync**](ItemAttributeOptionsAPI.md#UpdateItemAttributeOptionAsync) | **Put** /api/v2/CatalogService/ItemAttributeOptions/{itemAttributeOptionId} | Update an item attribute option



## CreateItemAttributeOptionAsync

> ItemAttributeOptionDtoEnvelope CreateItemAttributeOptionAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemAttributeOptionCreateDto(itemAttributeOptionCreateDto).Execute()

Create a new item attribute option



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
	itemAttributeOptionCreateDto := *openapiclient.NewItemAttributeOptionCreateDto("Name_example", "ItemAttributeId_example") // ItemAttributeOptionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttributeOptionsAPI.CreateItemAttributeOptionAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemAttributeOptionCreateDto(itemAttributeOptionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttributeOptionsAPI.CreateItemAttributeOptionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemAttributeOptionAsync`: ItemAttributeOptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttributeOptionsAPI.CreateItemAttributeOptionAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemAttributeOptionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemAttributeOptionCreateDto** | [**ItemAttributeOptionCreateDto**](ItemAttributeOptionCreateDto.md) |  | 

### Return type

[**ItemAttributeOptionDtoEnvelope**](ItemAttributeOptionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItemAttributeOptionAsync

> DeleteItemAttributeOptionAsync(ctx, itemAttributeOptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item attribute option



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
	itemAttributeOptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemAttributeOptionsAPI.DeleteItemAttributeOptionAsync(context.Background(), itemAttributeOptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttributeOptionsAPI.DeleteItemAttributeOptionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemAttributeOptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemAttributeOptionAsyncRequest struct via the builder pattern


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


## GetItemAttributeOptionByIdAsync

> ItemAttributeOptionDtoEnvelope GetItemAttributeOptionByIdAsync(ctx, itemAttributeOptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item attribute option by ID



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
	itemAttributeOptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttributeOptionsAPI.GetItemAttributeOptionByIdAsync(context.Background(), itemAttributeOptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttributeOptionsAPI.GetItemAttributeOptionByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemAttributeOptionByIdAsync`: ItemAttributeOptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttributeOptionsAPI.GetItemAttributeOptionByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemAttributeOptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemAttributeOptionByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemAttributeOptionDtoEnvelope**](ItemAttributeOptionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemAttributeOptionsAsync

> ItemAttributeOptionDtoListEnvelope GetItemAttributeOptionsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all item attribute options



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttributeOptionsAPI.GetItemAttributeOptionsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttributeOptionsAPI.GetItemAttributeOptionsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemAttributeOptionsAsync`: ItemAttributeOptionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttributeOptionsAPI.GetItemAttributeOptionsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemAttributeOptionsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemAttributeOptionDtoListEnvelope**](ItemAttributeOptionDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemAttributeOptionsCountAsync

> Int32Envelope GetItemAttributeOptionsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item attribute options count



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttributeOptionsAPI.GetItemAttributeOptionsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttributeOptionsAPI.GetItemAttributeOptionsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemAttributeOptionsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttributeOptionsAPI.GetItemAttributeOptionsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemAttributeOptionsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## PatchItemAttributeOptionAsync

> EmptyEnvelope PatchItemAttributeOptionAsync(ctx, itemAttributeOptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch an item attribute option



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
	itemAttributeOptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttributeOptionsAPI.PatchItemAttributeOptionAsync(context.Background(), itemAttributeOptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttributeOptionsAPI.PatchItemAttributeOptionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchItemAttributeOptionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttributeOptionsAPI.PatchItemAttributeOptionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemAttributeOptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchItemAttributeOptionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## UpdateItemAttributeOptionAsync

> ItemAttributeOptionDtoEnvelope UpdateItemAttributeOptionAsync(ctx, itemAttributeOptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemAttributeOptionUpdateDto(itemAttributeOptionUpdateDto).Execute()

Update an item attribute option



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
	itemAttributeOptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemAttributeOptionUpdateDto := *openapiclient.NewItemAttributeOptionUpdateDto("Name_example") // ItemAttributeOptionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemAttributeOptionsAPI.UpdateItemAttributeOptionAsync(context.Background(), itemAttributeOptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemAttributeOptionUpdateDto(itemAttributeOptionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemAttributeOptionsAPI.UpdateItemAttributeOptionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemAttributeOptionAsync`: ItemAttributeOptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemAttributeOptionsAPI.UpdateItemAttributeOptionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemAttributeOptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemAttributeOptionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemAttributeOptionUpdateDto** | [**ItemAttributeOptionUpdateDto**](ItemAttributeOptionUpdateDto.md) |  | 

### Return type

[**ItemAttributeOptionDtoEnvelope**](ItemAttributeOptionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

