# \ItemRetainSamplesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemRetainSampleAsync**](ItemRetainSamplesAPI.md#CreateItemRetainSampleAsync) | **Post** /api/v2/LogisticsService/ItemRetainSamples | Create an item retain sample
[**DeleteItemRetainSampleAsync**](ItemRetainSamplesAPI.md#DeleteItemRetainSampleAsync) | **Delete** /api/v2/LogisticsService/ItemRetainSamples/{retainSampleId} | Delete an item retain sample
[**GetItemRetainSampleByIdAsync**](ItemRetainSamplesAPI.md#GetItemRetainSampleByIdAsync) | **Get** /api/v2/LogisticsService/ItemRetainSamples/{retainSampleId} | Get item retain sample by ID
[**GetItemRetainSamplesAsync**](ItemRetainSamplesAPI.md#GetItemRetainSamplesAsync) | **Get** /api/v2/LogisticsService/ItemRetainSamples | Get all item retain samples
[**GetItemRetainSamplesCountAsync**](ItemRetainSamplesAPI.md#GetItemRetainSamplesCountAsync) | **Get** /api/v2/LogisticsService/ItemRetainSamples/Count | Get item retain samples count
[**PatchItemRetainSampleAsync**](ItemRetainSamplesAPI.md#PatchItemRetainSampleAsync) | **Patch** /api/v2/LogisticsService/ItemRetainSamples/{retainSampleId} | Patch an item retain sample
[**UpdateItemRetainSampleAsync**](ItemRetainSamplesAPI.md#UpdateItemRetainSampleAsync) | **Put** /api/v2/LogisticsService/ItemRetainSamples/{retainSampleId} | Update an item retain sample



## CreateItemRetainSampleAsync

> EmptyEnvelope CreateItemRetainSampleAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemRetainSampleCreateDto(itemRetainSampleCreateDto).Execute()

Create an item retain sample



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
	itemRetainSampleCreateDto := *openapiclient.NewItemRetainSampleCreateDto("WarehouseId_example", "ItemId_example") // ItemRetainSampleCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemRetainSamplesAPI.CreateItemRetainSampleAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemRetainSampleCreateDto(itemRetainSampleCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRetainSamplesAPI.CreateItemRetainSampleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemRetainSampleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRetainSamplesAPI.CreateItemRetainSampleAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemRetainSampleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemRetainSampleCreateDto** | [**ItemRetainSampleCreateDto**](ItemRetainSampleCreateDto.md) |  | 

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


## DeleteItemRetainSampleAsync

> EmptyEnvelope DeleteItemRetainSampleAsync(ctx, retainSampleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item retain sample



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
	retainSampleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemRetainSamplesAPI.DeleteItemRetainSampleAsync(context.Background(), retainSampleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRetainSamplesAPI.DeleteItemRetainSampleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteItemRetainSampleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRetainSamplesAPI.DeleteItemRetainSampleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retainSampleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemRetainSampleAsyncRequest struct via the builder pattern


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


## GetItemRetainSampleByIdAsync

> ItemRetainSampleDtoEnvelope GetItemRetainSampleByIdAsync(ctx, retainSampleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item retain sample by ID



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
	retainSampleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemRetainSamplesAPI.GetItemRetainSampleByIdAsync(context.Background(), retainSampleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRetainSamplesAPI.GetItemRetainSampleByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemRetainSampleByIdAsync`: ItemRetainSampleDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRetainSamplesAPI.GetItemRetainSampleByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retainSampleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRetainSampleByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemRetainSampleDtoEnvelope**](ItemRetainSampleDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemRetainSamplesAsync

> ItemRetainSampleDtoListEnvelope GetItemRetainSamplesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemRetainSampleDtoCollectionQueryParameters(itemRetainSampleDtoCollectionQueryParameters).Execute()

Get all item retain samples



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
	itemRetainSampleDtoCollectionQueryParameters := *openapiclient.NewItemRetainSampleDtoCollectionQueryParameters() // ItemRetainSampleDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemRetainSamplesAPI.GetItemRetainSamplesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemRetainSampleDtoCollectionQueryParameters(itemRetainSampleDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRetainSamplesAPI.GetItemRetainSamplesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemRetainSamplesAsync`: ItemRetainSampleDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRetainSamplesAPI.GetItemRetainSamplesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRetainSamplesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemRetainSampleDtoCollectionQueryParameters** | [**ItemRetainSampleDtoCollectionQueryParameters**](ItemRetainSampleDtoCollectionQueryParameters.md) |  | 

### Return type

[**ItemRetainSampleDtoListEnvelope**](ItemRetainSampleDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemRetainSamplesCountAsync

> Int32Envelope GetItemRetainSamplesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemRetainSampleDtoCollectionQueryParameters(itemRetainSampleDtoCollectionQueryParameters).Execute()

Get item retain samples count



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
	itemRetainSampleDtoCollectionQueryParameters := *openapiclient.NewItemRetainSampleDtoCollectionQueryParameters() // ItemRetainSampleDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemRetainSamplesAPI.GetItemRetainSamplesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemRetainSampleDtoCollectionQueryParameters(itemRetainSampleDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRetainSamplesAPI.GetItemRetainSamplesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemRetainSamplesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRetainSamplesAPI.GetItemRetainSamplesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRetainSamplesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemRetainSampleDtoCollectionQueryParameters** | [**ItemRetainSampleDtoCollectionQueryParameters**](ItemRetainSampleDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchItemRetainSampleAsync

> EmptyEnvelope PatchItemRetainSampleAsync(ctx, retainSampleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch an item retain sample



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
	retainSampleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemRetainSamplesAPI.PatchItemRetainSampleAsync(context.Background(), retainSampleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRetainSamplesAPI.PatchItemRetainSampleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchItemRetainSampleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRetainSamplesAPI.PatchItemRetainSampleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retainSampleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchItemRetainSampleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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


## UpdateItemRetainSampleAsync

> EmptyEnvelope UpdateItemRetainSampleAsync(ctx, retainSampleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemRetainSampleUpdateDto(itemRetainSampleUpdateDto).Execute()

Update an item retain sample



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
	retainSampleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemRetainSampleUpdateDto := *openapiclient.NewItemRetainSampleUpdateDto() // ItemRetainSampleUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemRetainSamplesAPI.UpdateItemRetainSampleAsync(context.Background(), retainSampleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemRetainSampleUpdateDto(itemRetainSampleUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRetainSamplesAPI.UpdateItemRetainSampleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemRetainSampleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRetainSamplesAPI.UpdateItemRetainSampleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**retainSampleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemRetainSampleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemRetainSampleUpdateDto** | [**ItemRetainSampleUpdateDto**](ItemRetainSampleUpdateDto.md) |  | 

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

