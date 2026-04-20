# \SocialFeedsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedPostAsync**](SocialFeedsAPI.md#CreateFeedPostAsync) | **Post** /api/v2/SocialService/SocialFeeds/{socialFeedId}/Posts | Create a social feed post
[**DeleteFeedPostAsync**](SocialFeedsAPI.md#DeleteFeedPostAsync) | **Delete** /api/v2/SocialService/SocialFeeds/{socialFeedId}/Posts/{feedPostId} | Delete a social feed post
[**GetFeedNotifications**](SocialFeedsAPI.md#GetFeedNotifications) | **Get** /api/v2/SocialService/SocialFeeds | Get social feeds
[**GetFeedPostAsync**](SocialFeedsAPI.md#GetFeedPostAsync) | **Get** /api/v2/SocialService/SocialFeeds/{socialFeedId}/Posts/{feedPostId} | Get social feed post by ID
[**GetFeedPostsAsync**](SocialFeedsAPI.md#GetFeedPostsAsync) | **Get** /api/v2/SocialService/SocialFeeds/{socialFeedId}/Posts | Get social feed posts
[**GetFeedPostsCountAsync**](SocialFeedsAPI.md#GetFeedPostsCountAsync) | **Get** /api/v2/SocialService/SocialFeeds/{socialFeedId}/Posts/Count | Count social feed posts
[**GetNotificationAsync**](SocialFeedsAPI.md#GetNotificationAsync) | **Get** /api/v2/SocialService/SocialFeeds/{socialFeedId} | Get social feed by ID
[**GetNotificationsCountAsync**](SocialFeedsAPI.md#GetNotificationsCountAsync) | **Get** /api/v2/SocialService/SocialFeeds/Count | Count social feeds
[**UpdateFeedPostAsync**](SocialFeedsAPI.md#UpdateFeedPostAsync) | **Put** /api/v2/SocialService/SocialFeeds/{socialFeedId}/Posts/{feedPostId} | Update a social feed post



## CreateFeedPostAsync

> SocialFeedPostDtoEnvelope CreateFeedPostAsync(ctx, socialFeedId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialFeedPostCreateDto(socialFeedPostCreateDto).Execute()

Create a social feed post



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
	socialFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialFeedPostCreateDto := *openapiclient.NewSocialFeedPostCreateDto() // SocialFeedPostCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialFeedsAPI.CreateFeedPostAsync(context.Background(), socialFeedId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialFeedPostCreateDto(socialFeedPostCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.CreateFeedPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFeedPostAsync`: SocialFeedPostDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.CreateFeedPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialFeedId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialFeedPostCreateDto** | [**SocialFeedPostCreateDto**](SocialFeedPostCreateDto.md) |  | 

### Return type

[**SocialFeedPostDtoEnvelope**](SocialFeedPostDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeedPostAsync

> EmptyEnvelope DeleteFeedPostAsync(ctx, socialFeedId, feedPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a social feed post



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
	socialFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	feedPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialFeedsAPI.DeleteFeedPostAsync(context.Background(), socialFeedId, feedPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.DeleteFeedPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFeedPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.DeleteFeedPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialFeedId** | **string** |  | 
**feedPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeedPostAsyncRequest struct via the builder pattern


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


## GetFeedNotifications

> SocialFeedDtoListEnvelope GetFeedNotifications(ctx).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social feeds



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
	resp, r, err := apiClient.SocialFeedsAPI.GetFeedNotifications(context.Background()).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.GetFeedNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedNotifications`: SocialFeedDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.GetFeedNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialFeedDtoListEnvelope**](SocialFeedDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedPostAsync

> SocialFeedPostDtoEnvelope GetFeedPostAsync(ctx, socialFeedId, feedPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social feed post by ID



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
	socialFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	feedPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialFeedsAPI.GetFeedPostAsync(context.Background(), socialFeedId, feedPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.GetFeedPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedPostAsync`: SocialFeedPostDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.GetFeedPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialFeedId** | **string** |  | 
**feedPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialFeedPostDtoEnvelope**](SocialFeedPostDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedPostsAsync

> SocialFeedPostDtoListEnvelope GetFeedPostsAsync(ctx, socialFeedId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social feed posts



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
	socialFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialFeedsAPI.GetFeedPostsAsync(context.Background(), socialFeedId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.GetFeedPostsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedPostsAsync`: SocialFeedPostDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.GetFeedPostsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialFeedId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedPostsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialFeedPostDtoListEnvelope**](SocialFeedPostDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeedPostsCountAsync

> Int32Envelope GetFeedPostsCountAsync(ctx, socialFeedId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count social feed posts



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
	socialFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialFeedsAPI.GetFeedPostsCountAsync(context.Background(), socialFeedId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.GetFeedPostsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeedPostsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.GetFeedPostsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialFeedId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedPostsCountAsyncRequest struct via the builder pattern


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


## GetNotificationAsync

> SocialFeedDtoEnvelope GetNotificationAsync(ctx, socialFeedId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social feed by ID



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
	socialFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialFeedsAPI.GetNotificationAsync(context.Background(), socialFeedId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.GetNotificationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationAsync`: SocialFeedDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.GetNotificationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialFeedId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialFeedDtoEnvelope**](SocialFeedDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsCountAsync

> Int32Envelope GetNotificationsCountAsync(ctx).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count social feeds



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
	resp, r, err := apiClient.SocialFeedsAPI.GetNotificationsCountAsync(context.Background()).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.GetNotificationsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.GetNotificationsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsCountAsyncRequest struct via the builder pattern


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


## UpdateFeedPostAsync

> SocialFeedPostDtoEnvelope UpdateFeedPostAsync(ctx, socialFeedId, feedPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialFeedPostUpdateDto(socialFeedPostUpdateDto).Execute()

Update a social feed post



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
	socialFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	feedPostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	socialFeedPostUpdateDto := *openapiclient.NewSocialFeedPostUpdateDto() // SocialFeedPostUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialFeedsAPI.UpdateFeedPostAsync(context.Background(), socialFeedId, feedPostId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SocialFeedPostUpdateDto(socialFeedPostUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialFeedsAPI.UpdateFeedPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFeedPostAsync`: SocialFeedPostDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialFeedsAPI.UpdateFeedPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialFeedId** | **string** |  | 
**feedPostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeedPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **socialFeedPostUpdateDto** | [**SocialFeedPostUpdateDto**](SocialFeedPostUpdateDto.md) |  | 

### Return type

[**SocialFeedPostDtoEnvelope**](SocialFeedPostDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

