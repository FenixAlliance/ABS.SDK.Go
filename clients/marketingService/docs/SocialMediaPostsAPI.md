# \SocialMediaPostsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSocialMediaPostAsync**](SocialMediaPostsAPI.md#CreateSocialMediaPostAsync) | **Post** /api/v2/MarketingService/SocialMediaPosts | Create a social media post
[**DeleteSocialMediaPostAsync**](SocialMediaPostsAPI.md#DeleteSocialMediaPostAsync) | **Delete** /api/v2/MarketingService/SocialMediaPosts/{socialmediapostId} | Delete a social media post
[**GetSocialMediaPostDetailsAsync**](SocialMediaPostsAPI.md#GetSocialMediaPostDetailsAsync) | **Get** /api/v2/MarketingService/SocialMediaPosts/{socialmediapostId} | Get social media post by ID
[**GetSocialMediaPostsCountAsync**](SocialMediaPostsAPI.md#GetSocialMediaPostsCountAsync) | **Get** /api/v2/MarketingService/SocialMediaPosts/Count | Get social media posts count
[**GetSocialMediaPostsODataAsync**](SocialMediaPostsAPI.md#GetSocialMediaPostsODataAsync) | **Get** /api/v2/MarketingService/SocialMediaPosts | Get social media posts
[**UpdateSocialMediaPostAsync**](SocialMediaPostsAPI.md#UpdateSocialMediaPostAsync) | **Put** /api/v2/MarketingService/SocialMediaPosts/{socialmediapostId} | Update a social media post



## CreateSocialMediaPostAsync

> EmptyEnvelope CreateSocialMediaPostAsync(ctx).TenantId(tenantId).SocialMediaPostCreateDto(socialMediaPostCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create a social media post



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
	socialMediaPostCreateDto := *openapiclient.NewSocialMediaPostCreateDto() // SocialMediaPostCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialMediaPostsAPI.CreateSocialMediaPostAsync(context.Background()).TenantId(tenantId).SocialMediaPostCreateDto(socialMediaPostCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.CreateSocialMediaPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSocialMediaPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.CreateSocialMediaPostAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSocialMediaPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **socialMediaPostCreateDto** | [**SocialMediaPostCreateDto**](SocialMediaPostCreateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## DeleteSocialMediaPostAsync

> EmptyEnvelope DeleteSocialMediaPostAsync(ctx, socialmediapostId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a social media post



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
	socialmediapostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialMediaPostsAPI.DeleteSocialMediaPostAsync(context.Background(), socialmediapostId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.DeleteSocialMediaPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSocialMediaPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.DeleteSocialMediaPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialmediapostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSocialMediaPostAsyncRequest struct via the builder pattern


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


## GetSocialMediaPostDetailsAsync

> SocialMediaPostDtoEnvelope GetSocialMediaPostDetailsAsync(ctx, socialmediapostId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social media post by ID



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
	socialmediapostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialMediaPostsAPI.GetSocialMediaPostDetailsAsync(context.Background(), socialmediapostId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.GetSocialMediaPostDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialMediaPostDetailsAsync`: SocialMediaPostDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.GetSocialMediaPostDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialmediapostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialMediaPostDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialMediaPostDtoEnvelope**](SocialMediaPostDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialMediaPostsCountAsync

> Int32Envelope GetSocialMediaPostsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social media posts count



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
	resp, r, err := apiClient.SocialMediaPostsAPI.GetSocialMediaPostsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.GetSocialMediaPostsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialMediaPostsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.GetSocialMediaPostsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialMediaPostsCountAsyncRequest struct via the builder pattern


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


## GetSocialMediaPostsODataAsync

> SocialMediaPostDtoListEnvelope GetSocialMediaPostsODataAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social media posts



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
	resp, r, err := apiClient.SocialMediaPostsAPI.GetSocialMediaPostsODataAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.GetSocialMediaPostsODataAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialMediaPostsODataAsync`: SocialMediaPostDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.GetSocialMediaPostsODataAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialMediaPostsODataAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialMediaPostDtoListEnvelope**](SocialMediaPostDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSocialMediaPostAsync

> EmptyEnvelope UpdateSocialMediaPostAsync(ctx, socialmediapostId).TenantId(tenantId).SocialMediaPostUpdateDto(socialMediaPostUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update a social media post



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
	socialmediapostId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialMediaPostUpdateDto := *openapiclient.NewSocialMediaPostUpdateDto() // SocialMediaPostUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialMediaPostsAPI.UpdateSocialMediaPostAsync(context.Background(), socialmediapostId).TenantId(tenantId).SocialMediaPostUpdateDto(socialMediaPostUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.UpdateSocialMediaPostAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSocialMediaPostAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.UpdateSocialMediaPostAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialmediapostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSocialMediaPostAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **socialMediaPostUpdateDto** | [**SocialMediaPostUpdateDto**](SocialMediaPostUpdateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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

