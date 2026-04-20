# \BlogPostCategoriesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountBlogPostCategoriesAsync**](BlogPostCategoriesAPI.md#CountBlogPostCategoriesAsync) | **Get** /api/v2/ContentService/BlogPostCategories/Count | Count blog post categories
[**CreateBlogPostCategoryAsync**](BlogPostCategoriesAPI.md#CreateBlogPostCategoryAsync) | **Post** /api/v2/ContentService/BlogPostCategories | Create a blog post category
[**DeleteBlogPostCategoryAsync**](BlogPostCategoriesAPI.md#DeleteBlogPostCategoryAsync) | **Delete** /api/v2/ContentService/BlogPostCategories/{blogPostCategoryId} | Delete a blog post category
[**GetBlogPostCategoriesAsync**](BlogPostCategoriesAPI.md#GetBlogPostCategoriesAsync) | **Get** /api/v2/ContentService/BlogPostCategories | Get blog post categories
[**GetBlogPostCategoryByIdAsync**](BlogPostCategoriesAPI.md#GetBlogPostCategoryByIdAsync) | **Get** /api/v2/ContentService/BlogPostCategories/{blogPostCategoryId} | Get blog post category by ID
[**UpdateBlogPostCategoryAsync**](BlogPostCategoriesAPI.md#UpdateBlogPostCategoryAsync) | **Put** /api/v2/ContentService/BlogPostCategories/{blogPostCategoryId} | Update a blog post category



## CountBlogPostCategoriesAsync

> Int32Envelope CountBlogPostCategoriesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count blog post categories



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
	resp, r, err := apiClient.BlogPostCategoriesAPI.CountBlogPostCategoriesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostCategoriesAPI.CountBlogPostCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountBlogPostCategoriesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostCategoriesAPI.CountBlogPostCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountBlogPostCategoriesAsyncRequest struct via the builder pattern


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


## CreateBlogPostCategoryAsync

> EmptyEnvelope CreateBlogPostCategoryAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BlogPostCategoryCreateDto(blogPostCategoryCreateDto).Execute()

Create a blog post category



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
	blogPostCategoryCreateDto := *openapiclient.NewBlogPostCategoryCreateDto() // BlogPostCategoryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostCategoriesAPI.CreateBlogPostCategoryAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BlogPostCategoryCreateDto(blogPostCategoryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostCategoriesAPI.CreateBlogPostCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlogPostCategoryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostCategoriesAPI.CreateBlogPostCategoryAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlogPostCategoryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **blogPostCategoryCreateDto** | [**BlogPostCategoryCreateDto**](BlogPostCategoryCreateDto.md) |  | 

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


## DeleteBlogPostCategoryAsync

> EmptyEnvelope DeleteBlogPostCategoryAsync(ctx, blogPostCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a blog post category



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
	blogPostCategoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostCategoriesAPI.DeleteBlogPostCategoryAsync(context.Background(), blogPostCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostCategoriesAPI.DeleteBlogPostCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlogPostCategoryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostCategoriesAPI.DeleteBlogPostCategoryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlogPostCategoryAsyncRequest struct via the builder pattern


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


## GetBlogPostCategoriesAsync

> BlogPostCategoryDtoListEnvelope GetBlogPostCategoriesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get blog post categories



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
	resp, r, err := apiClient.BlogPostCategoriesAPI.GetBlogPostCategoriesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostCategoriesAPI.GetBlogPostCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostCategoriesAsync`: BlogPostCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostCategoriesAPI.GetBlogPostCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostCategoriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BlogPostCategoryDtoListEnvelope**](BlogPostCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostCategoryByIdAsync

> BlogPostCategoryDtoEnvelope GetBlogPostCategoryByIdAsync(ctx, blogPostCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get blog post category by ID



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
	blogPostCategoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostCategoriesAPI.GetBlogPostCategoryByIdAsync(context.Background(), blogPostCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostCategoriesAPI.GetBlogPostCategoryByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostCategoryByIdAsync`: BlogPostCategoryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostCategoriesAPI.GetBlogPostCategoryByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostCategoryByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BlogPostCategoryDtoEnvelope**](BlogPostCategoryDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlogPostCategoryAsync

> EmptyEnvelope UpdateBlogPostCategoryAsync(ctx, blogPostCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BlogPostCategoryUpdateDto(blogPostCategoryUpdateDto).Execute()

Update a blog post category



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
	blogPostCategoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	blogPostCategoryUpdateDto := *openapiclient.NewBlogPostCategoryUpdateDto() // BlogPostCategoryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostCategoriesAPI.UpdateBlogPostCategoryAsync(context.Background(), blogPostCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BlogPostCategoryUpdateDto(blogPostCategoryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostCategoriesAPI.UpdateBlogPostCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlogPostCategoryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostCategoriesAPI.UpdateBlogPostCategoryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlogPostCategoryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **blogPostCategoryUpdateDto** | [**BlogPostCategoryUpdateDto**](BlogPostCategoryUpdateDto.md) |  | 

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

