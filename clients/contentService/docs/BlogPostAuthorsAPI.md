# \BlogPostAuthorsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountBlogPostsByAuthorAsync**](BlogPostAuthorsAPI.md#CountBlogPostsByAuthorAsync) | **Get** /api/v2/ContentService/BlogPostAuthors/{authorId}/BlogPosts/Count | Count blog posts by author
[**GetBlogAuthorByIdAsync**](BlogPostAuthorsAPI.md#GetBlogAuthorByIdAsync) | **Get** /api/v2/ContentService/BlogPostAuthors/{authorId} | Get blog author by ID
[**GetBlogAuthorsAsync**](BlogPostAuthorsAPI.md#GetBlogAuthorsAsync) | **Get** /api/v2/ContentService/BlogPostAuthors | Get blog authors
[**GetBlogPostsByAuthorAsync**](BlogPostAuthorsAPI.md#GetBlogPostsByAuthorAsync) | **Get** /api/v2/ContentService/BlogPostAuthors/{authorId}/BlogPosts | Get blog posts by author



## CountBlogPostsByAuthorAsync

> Int32Envelope CountBlogPostsByAuthorAsync(ctx, authorId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count blog posts by author



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
	authorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAuthorsAPI.CountBlogPostsByAuthorAsync(context.Background(), authorId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAuthorsAPI.CountBlogPostsByAuthorAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountBlogPostsByAuthorAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAuthorsAPI.CountBlogPostsByAuthorAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountBlogPostsByAuthorAsyncRequest struct via the builder pattern


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


## GetBlogAuthorByIdAsync

> BlogAuthorDtoEnvelope GetBlogAuthorByIdAsync(ctx, authorId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get blog author by ID



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
	authorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAuthorsAPI.GetBlogAuthorByIdAsync(context.Background(), authorId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAuthorsAPI.GetBlogAuthorByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogAuthorByIdAsync`: BlogAuthorDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAuthorsAPI.GetBlogAuthorByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogAuthorByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BlogAuthorDtoEnvelope**](BlogAuthorDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogAuthorsAsync

> BlogAuthorDtoListEnvelope GetBlogAuthorsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get blog authors



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
	resp, r, err := apiClient.BlogPostAuthorsAPI.GetBlogAuthorsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAuthorsAPI.GetBlogAuthorsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogAuthorsAsync`: BlogAuthorDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAuthorsAPI.GetBlogAuthorsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogAuthorsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BlogAuthorDtoListEnvelope**](BlogAuthorDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostsByAuthorAsync

> BlogPostDtoListEnvelope GetBlogPostsByAuthorAsync(ctx, authorId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get blog posts by author



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
	authorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostAuthorsAPI.GetBlogPostsByAuthorAsync(context.Background(), authorId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostAuthorsAPI.GetBlogPostsByAuthorAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostsByAuthorAsync`: BlogPostDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostAuthorsAPI.GetBlogPostsByAuthorAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostsByAuthorAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BlogPostDtoListEnvelope**](BlogPostDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

