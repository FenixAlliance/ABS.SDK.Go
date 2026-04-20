# \SocialPostsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSocialPostAsync**](SocialPostsAPI.md#CreateSocialPostAsync) | **Post** /api/v2/SocialService/SocialPosts | Create a social post
[**CreateSocialPostAttachmentAsync**](SocialPostsAPI.md#CreateSocialPostAttachmentAsync) | **Post** /api/v2/SocialService/SocialPosts/{socialPostId}/Attachments | Create a social post attachment
[**CreateSocialPostCommentAsync**](SocialPostsAPI.md#CreateSocialPostCommentAsync) | **Post** /api/v2/SocialService/SocialPosts/{socialPostId}/Comments | Create a social post comment
[**CreateSocialPostReactionAsync**](SocialPostsAPI.md#CreateSocialPostReactionAsync) | **Post** /api/v2/SocialService/SocialPosts/{socialPostId}/Reactions | Create a social post reaction
[**DeleteSocialPostAsync**](SocialPostsAPI.md#DeleteSocialPostAsync) | **Delete** /api/v2/SocialService/SocialPosts/{socialPostId} | Delete a social post
[**DeleteSocialPostAttachmentAsync**](SocialPostsAPI.md#DeleteSocialPostAttachmentAsync) | **Delete** /api/v2/SocialService/SocialPosts/{socialPostId}/Attachments/{attachmentId} | Delete a social post attachment
[**DeleteSocialPostCommentAsync**](SocialPostsAPI.md#DeleteSocialPostCommentAsync) | **Delete** /api/v2/SocialService/SocialPosts/{socialPostId}/Comments/{commentId} | Delete a social post comment
[**DeleteSocialPostReactionAsync**](SocialPostsAPI.md#DeleteSocialPostReactionAsync) | **Delete** /api/v2/SocialService/SocialPosts/{socialPostId}/Reactions/{reactionId} | Delete a social post reaction
[**GetSocialPostAsync**](SocialPostsAPI.md#GetSocialPostAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId} | Get social post by ID
[**GetSocialPostAttachmentAsync**](SocialPostsAPI.md#GetSocialPostAttachmentAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Attachments/{attachmentId} | Get social post attachment by ID
[**GetSocialPostAttachmentsAsync**](SocialPostsAPI.md#GetSocialPostAttachmentsAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Attachments | Get social post attachments
[**GetSocialPostAttachmentsCountAsync**](SocialPostsAPI.md#GetSocialPostAttachmentsCountAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Attachments/Count | Count social post attachments
[**GetSocialPostCommentAsync**](SocialPostsAPI.md#GetSocialPostCommentAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Comments/{commentId} | Get social post comment by ID
[**GetSocialPostCommentsAsync**](SocialPostsAPI.md#GetSocialPostCommentsAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Comments | Get social post comments
[**GetSocialPostCommentsCountAsync**](SocialPostsAPI.md#GetSocialPostCommentsCountAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Comments/Count | Count social post comments
[**GetSocialPostReactionAsync**](SocialPostsAPI.md#GetSocialPostReactionAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Reactions/{reactionId} | Get social post reaction by ID
[**GetSocialPostReactionsAsync**](SocialPostsAPI.md#GetSocialPostReactionsAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Reactions | Get social post reactions
[**GetSocialPostReactionsCountAsync**](SocialPostsAPI.md#GetSocialPostReactionsCountAsync) | **Get** /api/v2/SocialService/SocialPosts/{socialPostId}/Reactions/Count | Count social post reactions
[**GetSocialPostsAsync**](SocialPostsAPI.md#GetSocialPostsAsync) | **Get** /api/v2/SocialService/SocialPosts | Get social posts
[**GetSocialPostsCountAsync**](SocialPostsAPI.md#GetSocialPostsCountAsync) | **Get** /api/v2/SocialService/SocialPosts/Count | Count social posts
[**UpdateSocialPostAsync**](SocialPostsAPI.md#UpdateSocialPostAsync) | **Put** /api/v2/SocialService/SocialPosts/{socialPostId} | Update a social post
[**UpdateSocialPostAttachmentAsync**](SocialPostsAPI.md#UpdateSocialPostAttachmentAsync) | **Put** /api/v2/SocialService/SocialPosts/{socialPostId}/Attachments/{attachmentId} | Update a social post attachment
[**UpdateSocialPostCommentAsync**](SocialPostsAPI.md#UpdateSocialPostCommentAsync) | **Put** /api/v2/SocialService/SocialPosts/{socialPostId}/Comments/{commentId} | Update a social post comment
[**UpdateSocialPostReactionAsync**](SocialPostsAPI.md#UpdateSocialPostReactionAsync) | **Put** /api/v2/SocialService/SocialPosts/{socialPostId}/Reactions/{reactionId} | Update a social post reaction



## CreateSocialPostAsync

> SocialPostDtoEnvelope CreateSocialPostAsync(ctx).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostCreateDto(socialPostCreateDto).Execute()

Create a social post



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialPostCreateDto := *openapiclient.NewSocialPostCreateDto() // SocialPostCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.CreateSocialPostAsync(context.Background()).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostCreateDto(socialPostCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.CreateSocialPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSocialPostAsync`: SocialPostDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.CreateSocialPostAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSocialPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialPostCreateDto** | [**SocialPostCreateDto**](SocialPostCreateDto.md) |  | 

### Return type

[**SocialPostDtoEnvelope**](SocialPostDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSocialPostAttachmentAsync

> SocialPostAttachmentDtoEnvelope CreateSocialPostAttachmentAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostAttachmentCreateDto(socialPostAttachmentCreateDto).Execute()

Create a social post attachment



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
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialPostAttachmentCreateDto := *openapiclient.NewSocialPostAttachmentCreateDto() // SocialPostAttachmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.CreateSocialPostAttachmentAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostAttachmentCreateDto(socialPostAttachmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.CreateSocialPostAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSocialPostAttachmentAsync`: SocialPostAttachmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.CreateSocialPostAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSocialPostAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **socialProfileId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialPostAttachmentCreateDto** | [**SocialPostAttachmentCreateDto**](SocialPostAttachmentCreateDto.md) |  | 

### Return type

[**SocialPostAttachmentDtoEnvelope**](SocialPostAttachmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSocialPostCommentAsync

> EmptyEnvelope CreateSocialPostCommentAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostCommentCreateDto(socialPostCommentCreateDto).Execute()

Create a social post comment



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialPostCommentCreateDto := *openapiclient.NewSocialPostCommentCreateDto("Message_example") // SocialPostCommentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.CreateSocialPostCommentAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostCommentCreateDto(socialPostCommentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.CreateSocialPostCommentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSocialPostCommentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.CreateSocialPostCommentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSocialPostCommentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialPostCommentCreateDto** | [**SocialPostCommentCreateDto**](SocialPostCommentCreateDto.md) |  | 

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


## CreateSocialPostReactionAsync

> SocialReactionDtoEnvelope CreateSocialPostReactionAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialReactionCreateDto(socialReactionCreateDto).Execute()

Create a social post reaction



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
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialReactionCreateDto := *openapiclient.NewSocialReactionCreateDto() // SocialReactionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.CreateSocialPostReactionAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialReactionCreateDto(socialReactionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.CreateSocialPostReactionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSocialPostReactionAsync`: SocialReactionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.CreateSocialPostReactionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSocialPostReactionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **socialProfileId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialReactionCreateDto** | [**SocialReactionCreateDto**](SocialReactionCreateDto.md) |  | 

### Return type

[**SocialReactionDtoEnvelope**](SocialReactionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSocialPostAsync

> EmptyEnvelope DeleteSocialPostAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a social post



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.DeleteSocialPostAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.DeleteSocialPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSocialPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.DeleteSocialPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSocialPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

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


## DeleteSocialPostAttachmentAsync

> EmptyEnvelope DeleteSocialPostAttachmentAsync(ctx, socialPostId, attachmentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a social post attachment



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	attachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.DeleteSocialPostAttachmentAsync(context.Background(), socialPostId, attachmentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.DeleteSocialPostAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSocialPostAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.DeleteSocialPostAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSocialPostAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


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


## DeleteSocialPostCommentAsync

> EmptyEnvelope DeleteSocialPostCommentAsync(ctx, socialPostId, commentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a social post comment



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	commentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.DeleteSocialPostCommentAsync(context.Background(), socialPostId, commentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.DeleteSocialPostCommentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSocialPostCommentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.DeleteSocialPostCommentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSocialPostCommentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


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


## DeleteSocialPostReactionAsync

> EmptyEnvelope DeleteSocialPostReactionAsync(ctx, socialPostId, reactionId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a social post reaction



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	reactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.DeleteSocialPostReactionAsync(context.Background(), socialPostId, reactionId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.DeleteSocialPostReactionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSocialPostReactionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.DeleteSocialPostReactionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**reactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSocialPostReactionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


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


## GetSocialPostAsync

> SocialPostDtoEnvelope GetSocialPostAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post by ID



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostAsync`: SocialPostDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostDtoEnvelope**](SocialPostDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostAttachmentAsync

> EmptyEnvelope GetSocialPostAttachmentAsync(ctx, socialPostId, attachmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post attachment by ID



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
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	attachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostAttachmentAsync(context.Background(), socialPostId, attachmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSocialPostAttachmentsAsync

> SocialPostAttachmentDtoListEnvelope GetSocialPostAttachmentsAsync(ctx, socialPostId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post attachments



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
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostAttachmentsAsync(context.Background(), socialPostId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostAttachmentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostAttachmentsAsync`: SocialPostAttachmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostAttachmentsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostAttachmentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostAttachmentDtoListEnvelope**](SocialPostAttachmentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostAttachmentsCountAsync

> Int32Envelope GetSocialPostAttachmentsCountAsync(ctx, socialPostId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count social post attachments



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
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostAttachmentsCountAsync(context.Background(), socialPostId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostAttachmentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostAttachmentsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostAttachmentsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostAttachmentsCountAsyncRequest struct via the builder pattern


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


## GetSocialPostCommentAsync

> SocialPostCommentDtoEnvelope GetSocialPostCommentAsync(ctx, socialPostId, commentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post comment by ID



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	commentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostCommentAsync(context.Background(), socialPostId, commentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostCommentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostCommentAsync`: SocialPostCommentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostCommentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostCommentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostCommentDtoEnvelope**](SocialPostCommentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostCommentsAsync

> SocialPostCommentDtoListEnvelope GetSocialPostCommentsAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post comments



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostCommentsAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostCommentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostCommentsAsync`: SocialPostCommentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostCommentsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostCommentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostCommentDtoListEnvelope**](SocialPostCommentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostCommentsCountAsync

> Int32Envelope GetSocialPostCommentsCountAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count social post comments



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostCommentsCountAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostCommentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostCommentsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostCommentsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostCommentsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

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


## GetSocialPostReactionAsync

> SocialReactionDtoEnvelope GetSocialPostReactionAsync(ctx, socialPostId, reactionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post reaction by ID



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
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	reactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostReactionAsync(context.Background(), socialPostId, reactionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostReactionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostReactionAsync`: SocialReactionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostReactionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**reactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostReactionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialReactionDtoEnvelope**](SocialReactionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostReactionsAsync

> SocialReactionDtoListEnvelope GetSocialPostReactionsAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post reactions



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
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostReactionsAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostReactionsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostReactionsAsync`: SocialReactionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostReactionsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostReactionsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **socialProfileId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialReactionDtoListEnvelope**](SocialReactionDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostReactionsCountAsync

> Int32Envelope GetSocialPostReactionsCountAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count social post reactions



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
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostReactionsCountAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostReactionsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostReactionsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostReactionsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostReactionsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **socialProfileId** | **string** |  | 
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


## GetSocialPostsAsync

> SocialPostDtoListEnvelope GetSocialPostsAsync(ctx).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social posts



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostsAsync(context.Background()).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostsAsync`: SocialPostDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostDtoListEnvelope**](SocialPostDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostsCountAsync

> Int32Envelope GetSocialPostsCountAsync(ctx).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count social posts



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.GetSocialPostsCountAsync(context.Background()).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.GetSocialPostsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.GetSocialPostsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 
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


## UpdateSocialPostAsync

> EmptyEnvelope UpdateSocialPostAsync(ctx, socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostUpdateDto(socialPostUpdateDto).Execute()

Update a social post



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialPostUpdateDto := *openapiclient.NewSocialPostUpdateDto() // SocialPostUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.UpdateSocialPostAsync(context.Background(), socialPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostUpdateDto(socialPostUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.UpdateSocialPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSocialPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.UpdateSocialPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSocialPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialPostUpdateDto** | [**SocialPostUpdateDto**](SocialPostUpdateDto.md) |  | 

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


## UpdateSocialPostAttachmentAsync

> EmptyEnvelope UpdateSocialPostAttachmentAsync(ctx, socialPostId, attachmentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostAttachmentUpdateDto(socialPostAttachmentUpdateDto).Execute()

Update a social post attachment



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	attachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialPostAttachmentUpdateDto := *openapiclient.NewSocialPostAttachmentUpdateDto() // SocialPostAttachmentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.UpdateSocialPostAttachmentAsync(context.Background(), socialPostId, attachmentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostAttachmentUpdateDto(socialPostAttachmentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.UpdateSocialPostAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSocialPostAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.UpdateSocialPostAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSocialPostAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialPostAttachmentUpdateDto** | [**SocialPostAttachmentUpdateDto**](SocialPostAttachmentUpdateDto.md) |  | 

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


## UpdateSocialPostCommentAsync

> EmptyEnvelope UpdateSocialPostCommentAsync(ctx, socialPostId, commentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostCommentUpdateDto(socialPostCommentUpdateDto).Execute()

Update a social post comment



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	commentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialPostCommentUpdateDto := *openapiclient.NewSocialPostCommentUpdateDto() // SocialPostCommentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.UpdateSocialPostCommentAsync(context.Background(), socialPostId, commentId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialPostCommentUpdateDto(socialPostCommentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.UpdateSocialPostCommentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSocialPostCommentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.UpdateSocialPostCommentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSocialPostCommentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialPostCommentUpdateDto** | [**SocialPostCommentUpdateDto**](SocialPostCommentUpdateDto.md) |  | 

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


## UpdateSocialPostReactionAsync

> EmptyEnvelope UpdateSocialPostReactionAsync(ctx, socialPostId, reactionId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialReactionUpdateDto(socialReactionUpdateDto).Execute()

Update a social post reaction



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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	reactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialReactionUpdateDto := *openapiclient.NewSocialReactionUpdateDto() // SocialReactionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostsAPI.UpdateSocialPostReactionAsync(context.Background(), socialPostId, reactionId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialReactionUpdateDto(socialReactionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostsAPI.UpdateSocialPostReactionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSocialPostReactionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostsAPI.UpdateSocialPostReactionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialPostId** | **string** |  | 
**reactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSocialPostReactionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialReactionUpdateDto** | [**SocialReactionUpdateDto**](SocialReactionUpdateDto.md) |  | 

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

