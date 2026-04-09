# \BlogPostsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlogPostAsync**](BlogPostsAPI.md#CreateBlogPostAsync) | **Post** /api/v2/ContentService/BlogPosts | Create a new blog post
[**CreateCategoryForBlogPostAsync**](BlogPostsAPI.md#CreateCategoryForBlogPostAsync) | **Post** /api/v2/ContentService/BlogPosts/{blogPostId}/Categories | Create a category for a blog post
[**CreateCommentForBlogPostAsync**](BlogPostsAPI.md#CreateCommentForBlogPostAsync) | **Post** /api/v2/ContentService/BlogPosts/{blogPostId}/Comments | Create a comment for a blog post
[**CreateTagForBlogPostAsync**](BlogPostsAPI.md#CreateTagForBlogPostAsync) | **Post** /api/v2/ContentService/BlogPosts/{blogPostId}/Tags | Create a tag for a blog post
[**DeleteBlogPostAsync**](BlogPostsAPI.md#DeleteBlogPostAsync) | **Delete** /api/v2/ContentService/BlogPosts/{blogPostId} | Delete a blog post
[**DeleteCommentFromBlogPostAsync**](BlogPostsAPI.md#DeleteCommentFromBlogPostAsync) | **Delete** /api/v2/ContentService/BlogPosts/{blogPostId}/Comments/{commentId} | Delete a blog post comment
[**GetBlogPostByIdAsync**](BlogPostsAPI.md#GetBlogPostByIdAsync) | **Get** /api/v2/ContentService/BlogPosts/{blogPostId} | Get a blog post by ID
[**GetBlogPostsAsync**](BlogPostsAPI.md#GetBlogPostsAsync) | **Get** /api/v2/ContentService/BlogPosts | Retrieve a list of blog posts
[**GetBlogPostsCountAsync**](BlogPostsAPI.md#GetBlogPostsCountAsync) | **Get** /api/v2/ContentService/BlogPosts/Count | Get the count of blog posts
[**GetCategoriesForBlogPostAsync**](BlogPostsAPI.md#GetCategoriesForBlogPostAsync) | **Get** /api/v2/ContentService/BlogPosts/{blogPostId}/Categories | Get categories for a blog post
[**GetCommentsForBlogPostAsync**](BlogPostsAPI.md#GetCommentsForBlogPostAsync) | **Get** /api/v2/ContentService/BlogPosts/{blogPostId}/Comments | Get comments for a blog post
[**GetRepliesForCommentAsync**](BlogPostsAPI.md#GetRepliesForCommentAsync) | **Get** /api/v2/ContentService/BlogPosts/{blogPostId}/Comments/{commentId}/Replies | Get replies for a comment
[**GetTagsForBlogPostAsync**](BlogPostsAPI.md#GetTagsForBlogPostAsync) | **Get** /api/v2/ContentService/BlogPosts/{blogPostId}/Tags | Get tags for a blog post
[**RelateCategoryToBlogPostAsync**](BlogPostsAPI.md#RelateCategoryToBlogPostAsync) | **Post** /api/v2/ContentService/BlogPosts/{blogPostId}/Categories/{categoryId} | Relate an existing category to a blog post
[**RelateTagToBlogPostAsync**](BlogPostsAPI.md#RelateTagToBlogPostAsync) | **Post** /api/v2/ContentService/BlogPosts/{blogPostId}/Tags/{tagId} | Relate an existing tag to a blog post
[**ReplyToCommentAsync**](BlogPostsAPI.md#ReplyToCommentAsync) | **Post** /api/v2/ContentService/BlogPosts/{blogPostId}/Comments/{commentId}/Reply | Reply to a blog post comment
[**UnrelateCategoryFromBlogPostAsync**](BlogPostsAPI.md#UnrelateCategoryFromBlogPostAsync) | **Delete** /api/v2/ContentService/BlogPosts/{blogPostId}/Categories/{categoryId} | Remove a category from a blog post
[**UnrelateTagFromBlogPostAsync**](BlogPostsAPI.md#UnrelateTagFromBlogPostAsync) | **Delete** /api/v2/ContentService/BlogPosts/{blogPostId}/Tags/{tagId} | Remove a tag from a blog post
[**UpdateBlogPostAsync**](BlogPostsAPI.md#UpdateBlogPostAsync) | **Put** /api/v2/ContentService/BlogPosts/{blogPostId} | Update a blog post



## CreateBlogPostAsync

> EmptyEnvelope CreateBlogPostAsync(ctx).TenantId(tenantId).BlogPostCreateDto(blogPostCreateDto).Execute()

Create a new blog post



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
	blogPostCreateDto := *openapiclient.NewBlogPostCreateDto() // BlogPostCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.CreateBlogPostAsync(context.Background()).TenantId(tenantId).BlogPostCreateDto(blogPostCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.CreateBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.CreateBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **blogPostCreateDto** | [**BlogPostCreateDto**](BlogPostCreateDto.md) |  | 

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


## CreateCategoryForBlogPostAsync

> EmptyEnvelope CreateCategoryForBlogPostAsync(ctx, blogPostId).TenantId(tenantId).BlogPostCategoryCreateDto(blogPostCategoryCreateDto).Execute()

Create a category for a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	blogPostCategoryCreateDto := *openapiclient.NewBlogPostCategoryCreateDto() // BlogPostCategoryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.CreateCategoryForBlogPostAsync(context.Background(), blogPostId).TenantId(tenantId).BlogPostCategoryCreateDto(blogPostCategoryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.CreateCategoryForBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCategoryForBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.CreateCategoryForBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCategoryForBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## CreateCommentForBlogPostAsync

> EmptyEnvelope CreateCommentForBlogPostAsync(ctx, blogPostId).TenantId(tenantId).BlogPostCommentCreateDto(blogPostCommentCreateDto).Execute()

Create a comment for a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	blogPostCommentCreateDto := *openapiclient.NewBlogPostCommentCreateDto("Message_example") // BlogPostCommentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.CreateCommentForBlogPostAsync(context.Background(), blogPostId).TenantId(tenantId).BlogPostCommentCreateDto(blogPostCommentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.CreateCommentForBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommentForBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.CreateCommentForBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentForBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **blogPostCommentCreateDto** | [**BlogPostCommentCreateDto**](BlogPostCommentCreateDto.md) |  | 

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


## CreateTagForBlogPostAsync

> EmptyEnvelope CreateTagForBlogPostAsync(ctx, blogPostId).TenantId(tenantId).BlogPostTagCreateDto(blogPostTagCreateDto).Execute()

Create a tag for a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	blogPostTagCreateDto := *openapiclient.NewBlogPostTagCreateDto() // BlogPostTagCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.CreateTagForBlogPostAsync(context.Background(), blogPostId).TenantId(tenantId).BlogPostTagCreateDto(blogPostTagCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.CreateTagForBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTagForBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.CreateTagForBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagForBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **blogPostTagCreateDto** | [**BlogPostTagCreateDto**](BlogPostTagCreateDto.md) |  | 

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


## DeleteBlogPostAsync

> EmptyEnvelope DeleteBlogPostAsync(ctx, blogPostId).TenantId(tenantId).Execute()

Delete a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.DeleteBlogPostAsync(context.Background(), blogPostId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.DeleteBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.DeleteBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## DeleteCommentFromBlogPostAsync

> EmptyEnvelope DeleteCommentFromBlogPostAsync(ctx, blogPostId, commentId).TenantId(tenantId).Execute()

Delete a blog post comment



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	commentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.DeleteCommentFromBlogPostAsync(context.Background(), blogPostId, commentId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.DeleteCommentFromBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCommentFromBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.DeleteCommentFromBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentFromBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## GetBlogPostByIdAsync

> BlogPostDtoEnvelope GetBlogPostByIdAsync(ctx, blogPostId).Execute()

Get a blog post by ID



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetBlogPostByIdAsync(context.Background(), blogPostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetBlogPostByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostByIdAsync`: BlogPostDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetBlogPostByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlogPostDtoEnvelope**](BlogPostDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlogPostsAsync

> BlogPostDtoListEnvelope GetBlogPostsAsync(ctx).TenantId(tenantId).Execute()

Retrieve a list of blog posts



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetBlogPostsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetBlogPostsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostsAsync`: BlogPostDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetBlogPostsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## GetBlogPostsCountAsync

> Int32Envelope GetBlogPostsCountAsync(ctx).TenantId(tenantId).Execute()

Get the count of blog posts



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetBlogPostsCountAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetBlogPostsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlogPostsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetBlogPostsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlogPostsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## GetCategoriesForBlogPostAsync

> BlogPostCategoryDtoListEnvelope GetCategoriesForBlogPostAsync(ctx, blogPostId).Execute()

Get categories for a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetCategoriesForBlogPostAsync(context.Background(), blogPostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetCategoriesForBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoriesForBlogPostAsync`: BlogPostCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetCategoriesForBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesForBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetCommentsForBlogPostAsync

> BlogPostCommentDtoListEnvelope GetCommentsForBlogPostAsync(ctx, blogPostId).Execute()

Get comments for a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetCommentsForBlogPostAsync(context.Background(), blogPostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetCommentsForBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentsForBlogPostAsync`: BlogPostCommentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetCommentsForBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsForBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlogPostCommentDtoListEnvelope**](BlogPostCommentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepliesForCommentAsync

> BlogPostCommentDtoListEnvelope GetRepliesForCommentAsync(ctx, commentId, blogPostId).Execute()

Get replies for a comment



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
	commentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	blogPostId := "blogPostId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetRepliesForCommentAsync(context.Background(), commentId, blogPostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetRepliesForCommentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepliesForCommentAsync`: BlogPostCommentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetRepliesForCommentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepliesForCommentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPostCommentDtoListEnvelope**](BlogPostCommentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsForBlogPostAsync

> BlogPostTagDtoListEnvelope GetTagsForBlogPostAsync(ctx, blogPostId).Execute()

Get tags for a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetTagsForBlogPostAsync(context.Background(), blogPostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetTagsForBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagsForBlogPostAsync`: BlogPostTagDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetTagsForBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsForBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlogPostTagDtoListEnvelope**](BlogPostTagDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelateCategoryToBlogPostAsync

> EmptyEnvelope RelateCategoryToBlogPostAsync(ctx, blogPostId, categoryId).TenantId(tenantId).Execute()

Relate an existing category to a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	categoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.RelateCategoryToBlogPostAsync(context.Background(), blogPostId, categoryId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.RelateCategoryToBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelateCategoryToBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.RelateCategoryToBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelateCategoryToBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## RelateTagToBlogPostAsync

> EmptyEnvelope RelateTagToBlogPostAsync(ctx, blogPostId, tagId).TenantId(tenantId).Execute()

Relate an existing tag to a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.RelateTagToBlogPostAsync(context.Background(), blogPostId, tagId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.RelateTagToBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelateTagToBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.RelateTagToBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelateTagToBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## ReplyToCommentAsync

> EmptyEnvelope ReplyToCommentAsync(ctx, blogPostId, commentId).TenantId(tenantId).BlogPostCommentCreateDto(blogPostCommentCreateDto).Execute()

Reply to a blog post comment



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	commentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	blogPostCommentCreateDto := *openapiclient.NewBlogPostCommentCreateDto("Message_example") // BlogPostCommentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.ReplyToCommentAsync(context.Background(), blogPostId, commentId).TenantId(tenantId).BlogPostCommentCreateDto(blogPostCommentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.ReplyToCommentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplyToCommentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.ReplyToCommentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplyToCommentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **blogPostCommentCreateDto** | [**BlogPostCommentCreateDto**](BlogPostCommentCreateDto.md) |  | 

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


## UnrelateCategoryFromBlogPostAsync

> EmptyEnvelope UnrelateCategoryFromBlogPostAsync(ctx, blogPostId, categoryId).TenantId(tenantId).Execute()

Remove a category from a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	categoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.UnrelateCategoryFromBlogPostAsync(context.Background(), blogPostId, categoryId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.UnrelateCategoryFromBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnrelateCategoryFromBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.UnrelateCategoryFromBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnrelateCategoryFromBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## UnrelateTagFromBlogPostAsync

> EmptyEnvelope UnrelateTagFromBlogPostAsync(ctx, blogPostId, tagId).TenantId(tenantId).Execute()

Remove a tag from a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.UnrelateTagFromBlogPostAsync(context.Background(), blogPostId, tagId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.UnrelateTagFromBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnrelateTagFromBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.UnrelateTagFromBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnrelateTagFromBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## UpdateBlogPostAsync

> EmptyEnvelope UpdateBlogPostAsync(ctx, blogPostId).TenantId(tenantId).BlogPostUpdateDto(blogPostUpdateDto).Execute()

Update a blog post



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
	blogPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	blogPostUpdateDto := *openapiclient.NewBlogPostUpdateDto() // BlogPostUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.UpdateBlogPostAsync(context.Background(), blogPostId).TenantId(tenantId).BlogPostUpdateDto(blogPostUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.UpdateBlogPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlogPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.UpdateBlogPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blogPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlogPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **blogPostUpdateDto** | [**BlogPostUpdateDto**](BlogPostUpdateDto.md) |  | 

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

