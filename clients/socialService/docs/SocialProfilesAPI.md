# \SocialProfilesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountConversationsAsync**](SocialProfilesAPI.md#CountConversationsAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Conversations/Count | Count Conversations
[**CountFollowedProfilesAsync**](SocialProfilesAPI.md#CountFollowedProfilesAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Follows/Profiles/Count | Count Followed Profiles
[**CountFollowerProfilesAsync**](SocialProfilesAPI.md#CountFollowerProfilesAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Followers/Profiles/Count | Count Follower Profiles
[**CountFollowersAsync**](SocialProfilesAPI.md#CountFollowersAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Followers/Count | Count Followers
[**CountFollowsAsync**](SocialProfilesAPI.md#CountFollowsAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Follows/Count | Count Follows
[**CountMessagesAsync**](SocialProfilesAPI.md#CountMessagesAsync) | **Get** /api/v2/SocialService/SocialProfiles/{conversationId}/Messages/Count | Count Messages
[**CountNotificationsAsync**](SocialProfilesAPI.md#CountNotificationsAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Notifications/Count | Count Notifications
[**CountSocialProfilesAsync**](SocialProfilesAPI.md#CountSocialProfilesAsync) | **Get** /api/v2/SocialService/SocialProfiles/Count | Count Social Profiles
[**CreateConversationAsync**](SocialProfilesAPI.md#CreateConversationAsync) | **Post** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Conversations | Create Conversation
[**CreateMessageAsync**](SocialProfilesAPI.md#CreateMessageAsync) | **Post** /api/v2/SocialService/SocialProfiles/{conversationId}/Messages | Create Message
[**DeleteMessageAsync**](SocialProfilesAPI.md#DeleteMessageAsync) | **Delete** /api/v2/SocialService/SocialProfiles/{conversationId}/Messages/{messageId} | Delete Message
[**FollowAsync**](SocialProfilesAPI.md#FollowAsync) | **Post** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Follows/{followedSocialProfileId} | Follow
[**FollowExistsAsync**](SocialProfilesAPI.md#FollowExistsAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Follows/{followedSocialProfileId} | Check if Follow Exists
[**GetConversationsAsync**](SocialProfilesAPI.md#GetConversationsAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Conversations | Get Conversations
[**GetFollowedProfilesAsync**](SocialProfilesAPI.md#GetFollowedProfilesAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Follows/Profiles | Get Followed Profiles
[**GetFollowerProfilesAsync**](SocialProfilesAPI.md#GetFollowerProfilesAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Followers/Profiles | Get Follower Profiles
[**GetFollowersAsync**](SocialProfilesAPI.md#GetFollowersAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Followers | Get Followers
[**GetFollowsAsync**](SocialProfilesAPI.md#GetFollowsAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Follows | Get Follows
[**GetMessagesAsync**](SocialProfilesAPI.md#GetMessagesAsync) | **Get** /api/v2/SocialService/SocialProfiles/{conversationId}/Messages | Get Messages
[**GetNotificationsAsync**](SocialProfilesAPI.md#GetNotificationsAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Notifications | Get Notifications
[**GetSocialProfileAsync**](SocialProfilesAPI.md#GetSocialProfileAsync) | **Get** /api/v2/SocialService/SocialProfiles/{socialProfileId} | Get Social Profile
[**GetSocialProfilesAsync**](SocialProfilesAPI.md#GetSocialProfilesAsync) | **Get** /api/v2/SocialService/SocialProfiles | Get Social Profiles
[**UnfollowAsync**](SocialProfilesAPI.md#UnfollowAsync) | **Delete** /api/v2/SocialService/SocialProfiles/{socialProfileId}/Follows/{followedSocialProfileId} | Unfollow
[**UpdateMessageAsync**](SocialProfilesAPI.md#UpdateMessageAsync) | **Put** /api/v2/SocialService/SocialProfiles/{conversationId}/Messages/{messageId} | Update Message



## CountConversationsAsync

> Int32Envelope CountConversationsAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Conversations



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
	resp, r, err := apiClient.SocialProfilesAPI.CountConversationsAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CountConversationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountConversationsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CountConversationsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountConversationsAsyncRequest struct via the builder pattern


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


## CountFollowedProfilesAsync

> Int32Envelope CountFollowedProfilesAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Followed Profiles



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
	resp, r, err := apiClient.SocialProfilesAPI.CountFollowedProfilesAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CountFollowedProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountFollowedProfilesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CountFollowedProfilesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountFollowedProfilesAsyncRequest struct via the builder pattern


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


## CountFollowerProfilesAsync

> Int32Envelope CountFollowerProfilesAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Follower Profiles



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
	resp, r, err := apiClient.SocialProfilesAPI.CountFollowerProfilesAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CountFollowerProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountFollowerProfilesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CountFollowerProfilesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountFollowerProfilesAsyncRequest struct via the builder pattern


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


## CountFollowersAsync

> Int32Envelope CountFollowersAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Followers



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
	resp, r, err := apiClient.SocialProfilesAPI.CountFollowersAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CountFollowersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountFollowersAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CountFollowersAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountFollowersAsyncRequest struct via the builder pattern


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


## CountFollowsAsync

> Int32Envelope CountFollowsAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Follows



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
	resp, r, err := apiClient.SocialProfilesAPI.CountFollowsAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CountFollowsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountFollowsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CountFollowsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountFollowsAsyncRequest struct via the builder pattern


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


## CountMessagesAsync

> Int32Envelope CountMessagesAsync(ctx, conversationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Messages



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.CountMessagesAsync(context.Background(), conversationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CountMessagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountMessagesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CountMessagesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountMessagesAsyncRequest struct via the builder pattern


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


## CountNotificationsAsync

> Int32Envelope CountNotificationsAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Notifications



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
	resp, r, err := apiClient.SocialProfilesAPI.CountNotificationsAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CountNotificationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountNotificationsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CountNotificationsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountNotificationsAsyncRequest struct via the builder pattern


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


## CountSocialProfilesAsync

> Int32Envelope CountSocialProfilesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Social Profiles



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
	resp, r, err := apiClient.SocialProfilesAPI.CountSocialProfilesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CountSocialProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSocialProfilesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CountSocialProfilesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountSocialProfilesAsyncRequest struct via the builder pattern


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


## CreateConversationAsync

> EmptyEnvelope CreateConversationAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ConversationCreateDto(conversationCreateDto).Execute()

Create Conversation



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
	conversationCreateDto := *openapiclient.NewConversationCreateDto() // ConversationCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.CreateConversationAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ConversationCreateDto(conversationCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CreateConversationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConversationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CreateConversationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConversationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **conversationCreateDto** | [**ConversationCreateDto**](ConversationCreateDto.md) |  | 

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


## CreateMessageAsync

> EmptyEnvelope CreateMessageAsync(ctx, conversationId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PrivateMessageCreateDto(privateMessageCreateDto).Execute()

Create Message



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	privateMessageCreateDto := *openapiclient.NewPrivateMessageCreateDto() // PrivateMessageCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.CreateMessageAsync(context.Background(), conversationId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PrivateMessageCreateDto(privateMessageCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.CreateMessageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMessageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.CreateMessageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMessageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **privateMessageCreateDto** | [**PrivateMessageCreateDto**](PrivateMessageCreateDto.md) |  | 

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


## DeleteMessageAsync

> EmptyEnvelope DeleteMessageAsync(ctx, conversationId, messageId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete Message



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.DeleteMessageAsync(context.Background(), conversationId, messageId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.DeleteMessageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMessageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.DeleteMessageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageAsyncRequest struct via the builder pattern


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


## FollowAsync

> EmptyEnvelope FollowAsync(ctx, socialProfileId, followedSocialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Follow



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
	followedSocialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.FollowAsync(context.Background(), socialProfileId, followedSocialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.FollowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.FollowAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 
**followedSocialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFollowAsyncRequest struct via the builder pattern


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


## FollowExistsAsync

> BooleanEnvelope FollowExistsAsync(ctx, socialProfileId, followedSocialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Check if Follow Exists



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
	followedSocialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.FollowExistsAsync(context.Background(), socialProfileId, followedSocialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.FollowExistsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowExistsAsync`: BooleanEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.FollowExistsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 
**followedSocialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFollowExistsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BooleanEnvelope**](BooleanEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConversationsAsync

> ConversationDtoListEnvelope GetConversationsAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Conversations



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
	resp, r, err := apiClient.SocialProfilesAPI.GetConversationsAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetConversationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConversationsAsync`: ConversationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetConversationsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConversationsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ConversationDtoListEnvelope**](ConversationDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowedProfilesAsync

> SocialProfileDtoListEnvelope GetFollowedProfilesAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Followed Profiles



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
	resp, r, err := apiClient.SocialProfilesAPI.GetFollowedProfilesAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetFollowedProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowedProfilesAsync`: SocialProfileDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetFollowedProfilesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowedProfilesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialProfileDtoListEnvelope**](SocialProfileDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowerProfilesAsync

> SocialProfileDtoListEnvelope GetFollowerProfilesAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Follower Profiles



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
	resp, r, err := apiClient.SocialProfilesAPI.GetFollowerProfilesAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetFollowerProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowerProfilesAsync`: SocialProfileDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetFollowerProfilesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowerProfilesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialProfileDtoListEnvelope**](SocialProfileDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowersAsync

> FollowRecordDtoListEnvelope GetFollowersAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Followers



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
	resp, r, err := apiClient.SocialProfilesAPI.GetFollowersAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetFollowersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowersAsync`: FollowRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetFollowersAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FollowRecordDtoListEnvelope**](FollowRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowsAsync

> FollowRecordDtoListEnvelope GetFollowsAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Follows



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
	resp, r, err := apiClient.SocialProfilesAPI.GetFollowsAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetFollowsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowsAsync`: FollowRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetFollowsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FollowRecordDtoListEnvelope**](FollowRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagesAsync

> PrivateMessageDtoListEnvelope GetMessagesAsync(ctx, conversationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Messages



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.GetMessagesAsync(context.Background(), conversationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetMessagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessagesAsync`: PrivateMessageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetMessagesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PrivateMessageDtoListEnvelope**](PrivateMessageDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsAsync

> NotificationDtoListEnvelope GetNotificationsAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Notifications



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
	resp, r, err := apiClient.SocialProfilesAPI.GetNotificationsAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetNotificationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsAsync`: NotificationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetNotificationsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**NotificationDtoListEnvelope**](NotificationDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialProfileAsync

> SocialProfileDtoEnvelope GetSocialProfileAsync(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Social Profile



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
	resp, r, err := apiClient.SocialProfilesAPI.GetSocialProfileAsync(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetSocialProfileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialProfileAsync`: SocialProfileDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetSocialProfileAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialProfileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialProfileDtoEnvelope**](SocialProfileDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialProfilesAsync

> SocialProfileDtoListEnvelope GetSocialProfilesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Social Profiles



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
	resp, r, err := apiClient.SocialProfilesAPI.GetSocialProfilesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.GetSocialProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialProfilesAsync`: SocialProfileDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.GetSocialProfilesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialProfilesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialProfileDtoListEnvelope**](SocialProfileDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfollowAsync

> EmptyEnvelope UnfollowAsync(ctx, socialProfileId, followedSocialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Unfollow



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
	followedSocialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.UnfollowAsync(context.Background(), socialProfileId, followedSocialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.UnfollowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnfollowAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.UnfollowAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 
**followedSocialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowAsyncRequest struct via the builder pattern


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


## UpdateMessageAsync

> EmptyEnvelope UpdateMessageAsync(ctx, conversationId, messageId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PrivateMessageUpdateDto(privateMessageUpdateDto).Execute()

Update Message



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	privateMessageUpdateDto := *openapiclient.NewPrivateMessageUpdateDto() // PrivateMessageUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialProfilesAPI.UpdateMessageAsync(context.Background(), conversationId, messageId).SocialProfileId(socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PrivateMessageUpdateDto(privateMessageUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialProfilesAPI.UpdateMessageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMessageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialProfilesAPI.UpdateMessageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **privateMessageUpdateDto** | [**PrivateMessageUpdateDto**](PrivateMessageUpdateDto.md) |  | 

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

