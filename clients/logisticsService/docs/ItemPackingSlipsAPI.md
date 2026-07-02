# \ItemPackingSlipsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemPackingSlipAsync**](ItemPackingSlipsAPI.md#CreateItemPackingSlipAsync) | **Post** /api/v2/LogisticsService/ItemPackingSlips | Create an item packing slip
[**CreateItemPackingSlipEntryAsync**](ItemPackingSlipsAPI.md#CreateItemPackingSlipEntryAsync) | **Post** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId}/Entries | Create a packing slip entry
[**DeleteItemPackingSlipAsync**](ItemPackingSlipsAPI.md#DeleteItemPackingSlipAsync) | **Delete** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId} | Delete an item packing slip
[**DeleteItemPackingSlipEntryAsync**](ItemPackingSlipsAPI.md#DeleteItemPackingSlipEntryAsync) | **Delete** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId}/Entries/{entryId} | Delete a packing slip entry
[**GetItemPackingSlipByIdAsync**](ItemPackingSlipsAPI.md#GetItemPackingSlipByIdAsync) | **Get** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId} | Get item packing slip by ID
[**GetItemPackingSlipEntriesAsync**](ItemPackingSlipsAPI.md#GetItemPackingSlipEntriesAsync) | **Get** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId}/Entries | Get packing slip entries
[**GetItemPackingSlipEntriesCountAsync**](ItemPackingSlipsAPI.md#GetItemPackingSlipEntriesCountAsync) | **Get** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId}/Entries/Count | Get packing slip entries count
[**GetItemPackingSlipEntryByIdAsync**](ItemPackingSlipsAPI.md#GetItemPackingSlipEntryByIdAsync) | **Get** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId}/Entries/{entryId} | Get packing slip entry by ID
[**GetItemPackingSlipsAsync**](ItemPackingSlipsAPI.md#GetItemPackingSlipsAsync) | **Get** /api/v2/LogisticsService/ItemPackingSlips | Get all item packing slips
[**GetItemPackingSlipsCountAsync**](ItemPackingSlipsAPI.md#GetItemPackingSlipsCountAsync) | **Get** /api/v2/LogisticsService/ItemPackingSlips/Count | Get item packing slips count
[**PatchItemPackingSlipAsync**](ItemPackingSlipsAPI.md#PatchItemPackingSlipAsync) | **Patch** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId} | Patch an item packing slip
[**PatchItemPackingSlipEntryAsync**](ItemPackingSlipsAPI.md#PatchItemPackingSlipEntryAsync) | **Patch** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId}/Entries/{entryId} | Patch a packing slip entry
[**UpdateItemPackingSlipAsync**](ItemPackingSlipsAPI.md#UpdateItemPackingSlipAsync) | **Put** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId} | Update an item packing slip
[**UpdateItemPackingSlipEntryAsync**](ItemPackingSlipsAPI.md#UpdateItemPackingSlipEntryAsync) | **Put** /api/v2/LogisticsService/ItemPackingSlips/{packingSlipId}/Entries/{entryId} | Update a packing slip entry



## CreateItemPackingSlipAsync

> EmptyEnvelope CreateItemPackingSlipAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPackingSlipCreateDto(itemPackingSlipCreateDto).Execute()

Create an item packing slip



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
	itemPackingSlipCreateDto := *openapiclient.NewItemPackingSlipCreateDto() // ItemPackingSlipCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.CreateItemPackingSlipAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPackingSlipCreateDto(itemPackingSlipCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.CreateItemPackingSlipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemPackingSlipAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.CreateItemPackingSlipAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemPackingSlipAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPackingSlipCreateDto** | [**ItemPackingSlipCreateDto**](ItemPackingSlipCreateDto.md) |  | 

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


## CreateItemPackingSlipEntryAsync

> EmptyEnvelope CreateItemPackingSlipEntryAsync(ctx, packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPackingSlipEntryCreateDto(itemPackingSlipEntryCreateDto).Execute()

Create a packing slip entry



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemPackingSlipEntryCreateDto := *openapiclient.NewItemPackingSlipEntryCreateDto("ItemId_example", "ItemPackingSlipId_example", "Quantity_example") // ItemPackingSlipEntryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.CreateItemPackingSlipEntryAsync(context.Background(), packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPackingSlipEntryCreateDto(itemPackingSlipEntryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.CreateItemPackingSlipEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemPackingSlipEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.CreateItemPackingSlipEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemPackingSlipEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPackingSlipEntryCreateDto** | [**ItemPackingSlipEntryCreateDto**](ItemPackingSlipEntryCreateDto.md) |  | 

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


## DeleteItemPackingSlipAsync

> EmptyEnvelope DeleteItemPackingSlipAsync(ctx, packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item packing slip



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.DeleteItemPackingSlipAsync(context.Background(), packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.DeleteItemPackingSlipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteItemPackingSlipAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.DeleteItemPackingSlipAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemPackingSlipAsyncRequest struct via the builder pattern


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


## DeleteItemPackingSlipEntryAsync

> EmptyEnvelope DeleteItemPackingSlipEntryAsync(ctx, packingSlipId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a packing slip entry



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.DeleteItemPackingSlipEntryAsync(context.Background(), packingSlipId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.DeleteItemPackingSlipEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteItemPackingSlipEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.DeleteItemPackingSlipEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemPackingSlipEntryAsyncRequest struct via the builder pattern


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


## GetItemPackingSlipByIdAsync

> ItemPackingSlipDtoEnvelope GetItemPackingSlipByIdAsync(ctx, packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item packing slip by ID



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.GetItemPackingSlipByIdAsync(context.Background(), packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.GetItemPackingSlipByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPackingSlipByIdAsync`: ItemPackingSlipDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.GetItemPackingSlipByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPackingSlipByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPackingSlipDtoEnvelope**](ItemPackingSlipDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemPackingSlipEntriesAsync

> ItemPackingSlipEntryDtoListEnvelope GetItemPackingSlipEntriesAsync(ctx, packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get packing slip entries



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.GetItemPackingSlipEntriesAsync(context.Background(), packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.GetItemPackingSlipEntriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPackingSlipEntriesAsync`: ItemPackingSlipEntryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.GetItemPackingSlipEntriesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPackingSlipEntriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPackingSlipEntryDtoListEnvelope**](ItemPackingSlipEntryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemPackingSlipEntriesCountAsync

> Int32Envelope GetItemPackingSlipEntriesCountAsync(ctx, packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get packing slip entries count



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.GetItemPackingSlipEntriesCountAsync(context.Background(), packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.GetItemPackingSlipEntriesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPackingSlipEntriesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.GetItemPackingSlipEntriesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPackingSlipEntriesCountAsyncRequest struct via the builder pattern


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


## GetItemPackingSlipEntryByIdAsync

> ItemPackingSlipEntryDtoEnvelope GetItemPackingSlipEntryByIdAsync(ctx, packingSlipId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get packing slip entry by ID



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.GetItemPackingSlipEntryByIdAsync(context.Background(), packingSlipId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.GetItemPackingSlipEntryByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPackingSlipEntryByIdAsync`: ItemPackingSlipEntryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.GetItemPackingSlipEntryByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPackingSlipEntryByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPackingSlipEntryDtoEnvelope**](ItemPackingSlipEntryDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemPackingSlipsAsync

> ItemPackingSlipDtoListEnvelope GetItemPackingSlipsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all item packing slips



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
	resp, r, err := apiClient.ItemPackingSlipsAPI.GetItemPackingSlipsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.GetItemPackingSlipsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPackingSlipsAsync`: ItemPackingSlipDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.GetItemPackingSlipsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPackingSlipsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPackingSlipDtoListEnvelope**](ItemPackingSlipDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemPackingSlipsCountAsync

> Int32Envelope GetItemPackingSlipsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item packing slips count



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
	resp, r, err := apiClient.ItemPackingSlipsAPI.GetItemPackingSlipsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.GetItemPackingSlipsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPackingSlipsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.GetItemPackingSlipsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPackingSlipsCountAsyncRequest struct via the builder pattern


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


## PatchItemPackingSlipAsync

> EmptyEnvelope PatchItemPackingSlipAsync(ctx, packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch an item packing slip



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.PatchItemPackingSlipAsync(context.Background(), packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.PatchItemPackingSlipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchItemPackingSlipAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.PatchItemPackingSlipAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchItemPackingSlipAsyncRequest struct via the builder pattern


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


## PatchItemPackingSlipEntryAsync

> EmptyEnvelope PatchItemPackingSlipEntryAsync(ctx, packingSlipId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a packing slip entry



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.PatchItemPackingSlipEntryAsync(context.Background(), packingSlipId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.PatchItemPackingSlipEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchItemPackingSlipEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.PatchItemPackingSlipEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchItemPackingSlipEntryAsyncRequest struct via the builder pattern


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


## UpdateItemPackingSlipAsync

> EmptyEnvelope UpdateItemPackingSlipAsync(ctx, packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPackingSlipUpdateDto(itemPackingSlipUpdateDto).Execute()

Update an item packing slip



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemPackingSlipUpdateDto := *openapiclient.NewItemPackingSlipUpdateDto() // ItemPackingSlipUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.UpdateItemPackingSlipAsync(context.Background(), packingSlipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPackingSlipUpdateDto(itemPackingSlipUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.UpdateItemPackingSlipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemPackingSlipAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.UpdateItemPackingSlipAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemPackingSlipAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPackingSlipUpdateDto** | [**ItemPackingSlipUpdateDto**](ItemPackingSlipUpdateDto.md) |  | 

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


## UpdateItemPackingSlipEntryAsync

> EmptyEnvelope UpdateItemPackingSlipEntryAsync(ctx, packingSlipId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPackingSlipEntryUpdateDto(itemPackingSlipEntryUpdateDto).Execute()

Update a packing slip entry



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
	packingSlipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemPackingSlipEntryUpdateDto := *openapiclient.NewItemPackingSlipEntryUpdateDto() // ItemPackingSlipEntryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPackingSlipsAPI.UpdateItemPackingSlipEntryAsync(context.Background(), packingSlipId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPackingSlipEntryUpdateDto(itemPackingSlipEntryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPackingSlipsAPI.UpdateItemPackingSlipEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemPackingSlipEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPackingSlipsAPI.UpdateItemPackingSlipEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packingSlipId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemPackingSlipEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPackingSlipEntryUpdateDto** | [**ItemPackingSlipEntryUpdateDto**](ItemPackingSlipEntryUpdateDto.md) |  | 

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

