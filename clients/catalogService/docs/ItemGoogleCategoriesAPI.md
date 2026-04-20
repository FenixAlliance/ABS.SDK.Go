# \ItemGoogleCategoriesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllItemGoogleCategoriesAsync**](ItemGoogleCategoriesAPI.md#GetAllItemGoogleCategoriesAsync) | **Get** /api/v2/CatalogService/ItemGoogleCategories/All | Get all Google item categories (all)
[**GetChildrenItemGoogleCategoriesByIdAsync**](ItemGoogleCategoriesAPI.md#GetChildrenItemGoogleCategoriesByIdAsync) | **Get** /api/v2/CatalogService/ItemGoogleCategories/{itemCategoryId}/Children | Get children Google item categories by ID
[**GetItemGoogleCategoriesAsync**](ItemGoogleCategoriesAPI.md#GetItemGoogleCategoriesAsync) | **Get** /api/v2/CatalogService/ItemGoogleCategories | Get all Google item categories
[**GetItemGoogleCategoriesCountAsync**](ItemGoogleCategoriesAPI.md#GetItemGoogleCategoriesCountAsync) | **Get** /api/v2/CatalogService/ItemGoogleCategories/Count | Get Google item categories count
[**GetItemGoogleCategoriesTreeAsync**](ItemGoogleCategoriesAPI.md#GetItemGoogleCategoriesTreeAsync) | **Get** /api/v2/CatalogService/ItemGoogleCategories/tree | Get Google item categories tree
[**GetItemGoogleCategoryByIdAsync**](ItemGoogleCategoriesAPI.md#GetItemGoogleCategoryByIdAsync) | **Get** /api/v2/CatalogService/ItemGoogleCategories/{itemCategoryId} | Get Google item category by ID
[**GetRootItemGoogleCategoriesAsync**](ItemGoogleCategoriesAPI.md#GetRootItemGoogleCategoriesAsync) | **Get** /api/v2/CatalogService/ItemGoogleCategories/Primary | Get root Google item categories
[**MapItemGoogleCategoriesTreeAsync**](ItemGoogleCategoriesAPI.md#MapItemGoogleCategoriesTreeAsync) | **Post** /api/v2/CatalogService/ItemGoogleCategories/tree | Map Google item categories tree



## GetAllItemGoogleCategoriesAsync

> ItemGoogleCategoryDtoListEnvelope GetAllItemGoogleCategoriesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all Google item categories (all)



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemGoogleCategoriesAPI.GetAllItemGoogleCategoriesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemGoogleCategoriesAPI.GetAllItemGoogleCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllItemGoogleCategoriesAsync`: ItemGoogleCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemGoogleCategoriesAPI.GetAllItemGoogleCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllItemGoogleCategoriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemGoogleCategoryDtoListEnvelope**](ItemGoogleCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChildrenItemGoogleCategoriesByIdAsync

> ItemGoogleCategoryDtoListEnvelope GetChildrenItemGoogleCategoriesByIdAsync(ctx, itemCategoryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get children Google item categories by ID



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
	itemCategoryId := "itemCategoryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemGoogleCategoriesAPI.GetChildrenItemGoogleCategoriesByIdAsync(context.Background(), itemCategoryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemGoogleCategoriesAPI.GetChildrenItemGoogleCategoriesByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChildrenItemGoogleCategoriesByIdAsync`: ItemGoogleCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemGoogleCategoriesAPI.GetChildrenItemGoogleCategoriesByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChildrenItemGoogleCategoriesByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemGoogleCategoryDtoListEnvelope**](ItemGoogleCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemGoogleCategoriesAsync

> ItemGoogleCategoryDtoListEnvelope GetItemGoogleCategoriesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all Google item categories



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemGoogleCategoriesAPI.GetItemGoogleCategoriesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemGoogleCategoriesAPI.GetItemGoogleCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemGoogleCategoriesAsync`: ItemGoogleCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemGoogleCategoriesAPI.GetItemGoogleCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemGoogleCategoriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemGoogleCategoryDtoListEnvelope**](ItemGoogleCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemGoogleCategoriesCountAsync

> Int32Envelope GetItemGoogleCategoriesCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Google item categories count



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemGoogleCategoriesAPI.GetItemGoogleCategoriesCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemGoogleCategoriesAPI.GetItemGoogleCategoriesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemGoogleCategoriesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemGoogleCategoriesAPI.GetItemGoogleCategoriesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemGoogleCategoriesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## GetItemGoogleCategoriesTreeAsync

> ItemGoogleCategoryDtoListEnvelope GetItemGoogleCategoriesTreeAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Google item categories tree



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemGoogleCategoriesAPI.GetItemGoogleCategoriesTreeAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemGoogleCategoriesAPI.GetItemGoogleCategoriesTreeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemGoogleCategoriesTreeAsync`: ItemGoogleCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemGoogleCategoriesAPI.GetItemGoogleCategoriesTreeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemGoogleCategoriesTreeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemGoogleCategoryDtoListEnvelope**](ItemGoogleCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemGoogleCategoryByIdAsync

> ItemGoogleCategoryDtoEnvelope GetItemGoogleCategoryByIdAsync(ctx, itemCategoryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Google item category by ID



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
	itemCategoryId := "itemCategoryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemGoogleCategoriesAPI.GetItemGoogleCategoryByIdAsync(context.Background(), itemCategoryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemGoogleCategoriesAPI.GetItemGoogleCategoryByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemGoogleCategoryByIdAsync`: ItemGoogleCategoryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemGoogleCategoriesAPI.GetItemGoogleCategoryByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemGoogleCategoryByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemGoogleCategoryDtoEnvelope**](ItemGoogleCategoryDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootItemGoogleCategoriesAsync

> ItemGoogleCategoryDtoListEnvelope GetRootItemGoogleCategoriesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get root Google item categories



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemGoogleCategoriesAPI.GetRootItemGoogleCategoriesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemGoogleCategoriesAPI.GetRootItemGoogleCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRootItemGoogleCategoriesAsync`: ItemGoogleCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemGoogleCategoriesAPI.GetRootItemGoogleCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRootItemGoogleCategoriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemGoogleCategoryDtoListEnvelope**](ItemGoogleCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MapItemGoogleCategoriesTreeAsync

> ItemGoogleCategoryDtoListEnvelope MapItemGoogleCategoriesTreeAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Map Google item categories tree



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
	resp, r, err := apiClient.ItemGoogleCategoriesAPI.MapItemGoogleCategoriesTreeAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemGoogleCategoriesAPI.MapItemGoogleCategoriesTreeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MapItemGoogleCategoriesTreeAsync`: ItemGoogleCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemGoogleCategoriesAPI.MapItemGoogleCategoriesTreeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMapItemGoogleCategoriesTreeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemGoogleCategoryDtoListEnvelope**](ItemGoogleCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

