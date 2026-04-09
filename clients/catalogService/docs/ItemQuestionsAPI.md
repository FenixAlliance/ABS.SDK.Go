# \ItemQuestionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemQuestionAsync**](ItemQuestionsAPI.md#CreateItemQuestionAsync) | **Post** /api/v2/CatalogService/ItemQuestions | Create a new item question
[**DeleteItemQuestionAsync**](ItemQuestionsAPI.md#DeleteItemQuestionAsync) | **Delete** /api/v2/CatalogService/ItemQuestions/{itemQuestionId} | Delete an item question
[**GetItemQuestionByIdAsync**](ItemQuestionsAPI.md#GetItemQuestionByIdAsync) | **Get** /api/v2/CatalogService/ItemQuestions/{itemQuestionId} | Get item question by ID
[**GetItemQuestionsAsync**](ItemQuestionsAPI.md#GetItemQuestionsAsync) | **Get** /api/v2/CatalogService/ItemQuestions | Get all item questions
[**UpdateItemQuestionAsync**](ItemQuestionsAPI.md#UpdateItemQuestionAsync) | **Put** /api/v2/CatalogService/ItemQuestions/{itemQuestionId} | Update an item question



## CreateItemQuestionAsync

> ItemQuestionDtoEnvelope CreateItemQuestionAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemQuestionCreateDto(itemQuestionCreateDto).Execute()

Create a new item question



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
	itemQuestionCreateDto := *openapiclient.NewItemQuestionCreateDto("Title_example", false, "Question_example", "BusinessID_example", "ItemID_example") // ItemQuestionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemQuestionsAPI.CreateItemQuestionAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemQuestionCreateDto(itemQuestionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemQuestionsAPI.CreateItemQuestionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemQuestionAsync`: ItemQuestionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemQuestionsAPI.CreateItemQuestionAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemQuestionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemQuestionCreateDto** | [**ItemQuestionCreateDto**](ItemQuestionCreateDto.md) |  | 

### Return type

[**ItemQuestionDtoEnvelope**](ItemQuestionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItemQuestionAsync

> DeleteItemQuestionAsync(ctx, itemQuestionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item question



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
	itemQuestionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemQuestionsAPI.DeleteItemQuestionAsync(context.Background(), itemQuestionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemQuestionsAPI.DeleteItemQuestionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemQuestionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemQuestionAsyncRequest struct via the builder pattern


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


## GetItemQuestionByIdAsync

> ItemQuestionDtoEnvelope GetItemQuestionByIdAsync(ctx, itemQuestionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item question by ID



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
	itemQuestionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemQuestionsAPI.GetItemQuestionByIdAsync(context.Background(), itemQuestionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemQuestionsAPI.GetItemQuestionByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemQuestionByIdAsync`: ItemQuestionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemQuestionsAPI.GetItemQuestionByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemQuestionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemQuestionByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemQuestionDtoEnvelope**](ItemQuestionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemQuestionsAsync

> ItemQuestionDtoListEnvelope GetItemQuestionsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all item questions



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
	resp, r, err := apiClient.ItemQuestionsAPI.GetItemQuestionsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemQuestionsAPI.GetItemQuestionsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemQuestionsAsync`: ItemQuestionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemQuestionsAPI.GetItemQuestionsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemQuestionsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemQuestionDtoListEnvelope**](ItemQuestionDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItemQuestionAsync

> UpdateItemQuestionAsync(ctx, itemQuestionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemQuestionUpdateDto(itemQuestionUpdateDto).Execute()

Update an item question



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
	itemQuestionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemQuestionUpdateDto := *openapiclient.NewItemQuestionUpdateDto(false) // ItemQuestionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemQuestionsAPI.UpdateItemQuestionAsync(context.Background(), itemQuestionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemQuestionUpdateDto(itemQuestionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemQuestionsAPI.UpdateItemQuestionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemQuestionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemQuestionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemQuestionUpdateDto** | [**ItemQuestionUpdateDto**](ItemQuestionUpdateDto.md) |  | 

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

