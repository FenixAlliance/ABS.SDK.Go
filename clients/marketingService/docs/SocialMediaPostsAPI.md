# \SocialMediaPostsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceSocialMediaPostsCountGet**](SocialMediaPostsAPI.md#ApiV2MarketingServiceSocialMediaPostsCountGet) | **Get** /api/v2/MarketingService/SocialMediaPosts/Count | 
[**ApiV2MarketingServiceSocialMediaPostsGet**](SocialMediaPostsAPI.md#ApiV2MarketingServiceSocialMediaPostsGet) | **Get** /api/v2/MarketingService/SocialMediaPosts | 
[**ApiV2MarketingServiceSocialMediaPostsPost**](SocialMediaPostsAPI.md#ApiV2MarketingServiceSocialMediaPostsPost) | **Post** /api/v2/MarketingService/SocialMediaPosts | 
[**ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDelete**](SocialMediaPostsAPI.md#ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDelete) | **Delete** /api/v2/MarketingService/SocialMediaPosts/{socialmediapostId} | 
[**ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGet**](SocialMediaPostsAPI.md#ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGet) | **Get** /api/v2/MarketingService/SocialMediaPosts/{socialmediapostId} | 
[**ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPut**](SocialMediaPostsAPI.md#ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPut) | **Put** /api/v2/MarketingService/SocialMediaPosts/{socialmediapostId} | 



## ApiV2MarketingServiceSocialMediaPostsCountGet

> Int32Envelope ApiV2MarketingServiceSocialMediaPostsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialMediaPostsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialMediaPostsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceSocialMediaPostsGet

> SocialMediaPostDtoListEnvelope ApiV2MarketingServiceSocialMediaPostsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialMediaPostsGet`: SocialMediaPostDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialMediaPostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialMediaPostDtoListEnvelope**](SocialMediaPostDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceSocialMediaPostsPost

> EmptyEnvelope ApiV2MarketingServiceSocialMediaPostsPost(ctx).TenantId(tenantId).SocialMediaPostCreateDto(socialMediaPostCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsPost(context.Background()).TenantId(tenantId).SocialMediaPostCreateDto(socialMediaPostCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialMediaPostsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialMediaPostsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **socialMediaPostCreateDto** | [**SocialMediaPostCreateDto**](SocialMediaPostCreateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDelete

> EmptyEnvelope ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDelete(ctx, socialmediapostId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDelete(context.Background(), socialmediapostId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialmediapostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialMediaPostsSocialmediapostIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGet

> SocialMediaPostDtoEnvelope ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGet(ctx, socialmediapostId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGet(context.Background(), socialmediapostId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGet`: SocialMediaPostDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialmediapostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialMediaPostsSocialmediapostIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialMediaPostDtoEnvelope**](SocialMediaPostDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPut

> EmptyEnvelope ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPut(ctx, socialmediapostId).TenantId(tenantId).SocialMediaPostUpdateDto(socialMediaPostUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPut(context.Background(), socialmediapostId).TenantId(tenantId).SocialMediaPostUpdateDto(socialMediaPostUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialMediaPostsAPI.ApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialmediapostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialMediaPostsSocialmediapostIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **socialMediaPostUpdateDto** | [**SocialMediaPostUpdateDto**](SocialMediaPostUpdateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

