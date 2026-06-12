# \ItemPickListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemPickListAsync**](ItemPickListsAPI.md#CreateItemPickListAsync) | **Post** /api/v2/LogisticsService/ItemPickLists | Create an item pick list
[**CreateItemPickListEntryAsync**](ItemPickListsAPI.md#CreateItemPickListEntryAsync) | **Post** /api/v2/LogisticsService/ItemPickLists/{pickListId}/Entries | Create a pick list entry
[**DeleteItemPickListAsync**](ItemPickListsAPI.md#DeleteItemPickListAsync) | **Delete** /api/v2/LogisticsService/ItemPickLists/{pickListId} | Delete an item pick list
[**DeleteItemPickListEntryAsync**](ItemPickListsAPI.md#DeleteItemPickListEntryAsync) | **Delete** /api/v2/LogisticsService/ItemPickLists/{pickListId}/Entries/{entryId} | Delete a pick list entry
[**GetItemPickListByIdAsync**](ItemPickListsAPI.md#GetItemPickListByIdAsync) | **Get** /api/v2/LogisticsService/ItemPickLists/{pickListId} | Get item pick list by ID
[**GetItemPickListEntriesAsync**](ItemPickListsAPI.md#GetItemPickListEntriesAsync) | **Get** /api/v2/LogisticsService/ItemPickLists/{pickListId}/Entries | Get pick list entries
[**GetItemPickListEntriesCountAsync**](ItemPickListsAPI.md#GetItemPickListEntriesCountAsync) | **Get** /api/v2/LogisticsService/ItemPickLists/{pickListId}/Entries/Count | Get pick list entries count
[**GetItemPickListEntryByIdAsync**](ItemPickListsAPI.md#GetItemPickListEntryByIdAsync) | **Get** /api/v2/LogisticsService/ItemPickLists/{pickListId}/Entries/{entryId} | Get pick list entry by ID
[**GetItemPickListsAsync**](ItemPickListsAPI.md#GetItemPickListsAsync) | **Get** /api/v2/LogisticsService/ItemPickLists | Get all item pick lists
[**GetItemPickListsCountAsync**](ItemPickListsAPI.md#GetItemPickListsCountAsync) | **Get** /api/v2/LogisticsService/ItemPickLists/Count | Get item pick lists count
[**PatchItemPickListAsync**](ItemPickListsAPI.md#PatchItemPickListAsync) | **Patch** /api/v2/LogisticsService/ItemPickLists/{pickListId} | Patch an item pick list
[**PatchItemPickListEntryAsync**](ItemPickListsAPI.md#PatchItemPickListEntryAsync) | **Patch** /api/v2/LogisticsService/ItemPickLists/{pickListId}/Entries/{entryId} | Patch a pick list entry
[**UpdateItemPickListAsync**](ItemPickListsAPI.md#UpdateItemPickListAsync) | **Put** /api/v2/LogisticsService/ItemPickLists/{pickListId} | Update an item pick list
[**UpdateItemPickListEntryAsync**](ItemPickListsAPI.md#UpdateItemPickListEntryAsync) | **Put** /api/v2/LogisticsService/ItemPickLists/{pickListId}/Entries/{entryId} | Update a pick list entry



## CreateItemPickListAsync

> EmptyEnvelope CreateItemPickListAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPickListCreateDto(itemPickListCreateDto).Execute()

Create an item pick list



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
	itemPickListCreateDto := *openapiclient.NewItemPickListCreateDto("Name_example") // ItemPickListCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.CreateItemPickListAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPickListCreateDto(itemPickListCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.CreateItemPickListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemPickListAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.CreateItemPickListAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemPickListAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPickListCreateDto** | [**ItemPickListCreateDto**](ItemPickListCreateDto.md) |  | 

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


## CreateItemPickListEntryAsync

> EmptyEnvelope CreateItemPickListEntryAsync(ctx, pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPickListEntryCreateDto(itemPickListEntryCreateDto).Execute()

Create a pick list entry



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemPickListEntryCreateDto := *openapiclient.NewItemPickListEntryCreateDto("ItemId_example", "WarehouseId_example", "ItemPickListId_example") // ItemPickListEntryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.CreateItemPickListEntryAsync(context.Background(), pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPickListEntryCreateDto(itemPickListEntryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.CreateItemPickListEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemPickListEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.CreateItemPickListEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemPickListEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPickListEntryCreateDto** | [**ItemPickListEntryCreateDto**](ItemPickListEntryCreateDto.md) |  | 

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


## DeleteItemPickListAsync

> EmptyEnvelope DeleteItemPickListAsync(ctx, pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item pick list



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.DeleteItemPickListAsync(context.Background(), pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.DeleteItemPickListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteItemPickListAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.DeleteItemPickListAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemPickListAsyncRequest struct via the builder pattern


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


## DeleteItemPickListEntryAsync

> EmptyEnvelope DeleteItemPickListEntryAsync(ctx, pickListId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a pick list entry



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.DeleteItemPickListEntryAsync(context.Background(), pickListId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.DeleteItemPickListEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteItemPickListEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.DeleteItemPickListEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemPickListEntryAsyncRequest struct via the builder pattern


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


## GetItemPickListByIdAsync

> ItemPickListDtoEnvelope GetItemPickListByIdAsync(ctx, pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item pick list by ID



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.GetItemPickListByIdAsync(context.Background(), pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.GetItemPickListByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPickListByIdAsync`: ItemPickListDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.GetItemPickListByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPickListByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPickListDtoEnvelope**](ItemPickListDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemPickListEntriesAsync

> ItemPickListEntryDtoListEnvelope GetItemPickListEntriesAsync(ctx, pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get pick list entries



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.GetItemPickListEntriesAsync(context.Background(), pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.GetItemPickListEntriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPickListEntriesAsync`: ItemPickListEntryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.GetItemPickListEntriesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPickListEntriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPickListEntryDtoListEnvelope**](ItemPickListEntryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemPickListEntriesCountAsync

> Int32Envelope GetItemPickListEntriesCountAsync(ctx, pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get pick list entries count



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.GetItemPickListEntriesCountAsync(context.Background(), pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.GetItemPickListEntriesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPickListEntriesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.GetItemPickListEntriesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPickListEntriesCountAsyncRequest struct via the builder pattern


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


## GetItemPickListEntryByIdAsync

> ItemPickListEntryDtoEnvelope GetItemPickListEntryByIdAsync(ctx, pickListId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get pick list entry by ID



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.GetItemPickListEntryByIdAsync(context.Background(), pickListId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.GetItemPickListEntryByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPickListEntryByIdAsync`: ItemPickListEntryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.GetItemPickListEntryByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPickListEntryByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPickListEntryDtoEnvelope**](ItemPickListEntryDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemPickListsAsync

> ItemPickListDtoListEnvelope GetItemPickListsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all item pick lists



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
	resp, r, err := apiClient.ItemPickListsAPI.GetItemPickListsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.GetItemPickListsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPickListsAsync`: ItemPickListDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.GetItemPickListsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPickListsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemPickListDtoListEnvelope**](ItemPickListDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemPickListsCountAsync

> Int32Envelope GetItemPickListsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item pick lists count



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
	resp, r, err := apiClient.ItemPickListsAPI.GetItemPickListsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.GetItemPickListsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemPickListsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.GetItemPickListsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemPickListsCountAsyncRequest struct via the builder pattern


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


## PatchItemPickListAsync

> EmptyEnvelope PatchItemPickListAsync(ctx, pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch an item pick list



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.PatchItemPickListAsync(context.Background(), pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.PatchItemPickListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchItemPickListAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.PatchItemPickListAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchItemPickListAsyncRequest struct via the builder pattern


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


## PatchItemPickListEntryAsync

> EmptyEnvelope PatchItemPickListEntryAsync(ctx, pickListId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a pick list entry



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.PatchItemPickListEntryAsync(context.Background(), pickListId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.PatchItemPickListEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchItemPickListEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.PatchItemPickListEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchItemPickListEntryAsyncRequest struct via the builder pattern


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


## UpdateItemPickListAsync

> EmptyEnvelope UpdateItemPickListAsync(ctx, pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPickListUpdateDto(itemPickListUpdateDto).Execute()

Update an item pick list



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemPickListUpdateDto := *openapiclient.NewItemPickListUpdateDto() // ItemPickListUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.UpdateItemPickListAsync(context.Background(), pickListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPickListUpdateDto(itemPickListUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.UpdateItemPickListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemPickListAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.UpdateItemPickListAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemPickListAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPickListUpdateDto** | [**ItemPickListUpdateDto**](ItemPickListUpdateDto.md) |  | 

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


## UpdateItemPickListEntryAsync

> EmptyEnvelope UpdateItemPickListEntryAsync(ctx, pickListId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPickListEntryUpdateDto(itemPickListEntryUpdateDto).Execute()

Update a pick list entry



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
	pickListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemPickListEntryUpdateDto := *openapiclient.NewItemPickListEntryUpdateDto() // ItemPickListEntryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemPickListsAPI.UpdateItemPickListEntryAsync(context.Background(), pickListId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemPickListEntryUpdateDto(itemPickListEntryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemPickListsAPI.UpdateItemPickListEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemPickListEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemPickListsAPI.UpdateItemPickListEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickListId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemPickListEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemPickListEntryUpdateDto** | [**ItemPickListEntryUpdateDto**](ItemPickListEntryUpdateDto.md) |  | 

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

