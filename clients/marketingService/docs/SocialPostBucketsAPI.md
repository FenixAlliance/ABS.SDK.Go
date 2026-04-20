# \SocialPostBucketsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSocialPostBucketAsync**](SocialPostBucketsAPI.md#CreateSocialPostBucketAsync) | **Post** /api/v2/MarketingService/SocialPostBuckets | Create a social post bucket
[**DeleteSocialPostBucketAsync**](SocialPostBucketsAPI.md#DeleteSocialPostBucketAsync) | **Delete** /api/v2/MarketingService/SocialPostBuckets/{socialpostbucketId} | Delete a social post bucket
[**GetSocialPostBucketDetailsAsync**](SocialPostBucketsAPI.md#GetSocialPostBucketDetailsAsync) | **Get** /api/v2/MarketingService/SocialPostBuckets/{socialpostbucketId} | Get social post bucket by ID
[**GetSocialPostBucketsCountAsync**](SocialPostBucketsAPI.md#GetSocialPostBucketsCountAsync) | **Get** /api/v2/MarketingService/SocialPostBuckets/Count | Get social post buckets count
[**GetSocialPostBucketsODataAsync**](SocialPostBucketsAPI.md#GetSocialPostBucketsODataAsync) | **Get** /api/v2/MarketingService/SocialPostBuckets | Get social post buckets
[**UpdateSocialPostBucketAsync**](SocialPostBucketsAPI.md#UpdateSocialPostBucketAsync) | **Put** /api/v2/MarketingService/SocialPostBuckets/{socialpostbucketId} | Update a social post bucket



## CreateSocialPostBucketAsync

> EmptyEnvelope CreateSocialPostBucketAsync(ctx).TenantId(tenantId).SocialPostBucketCreateDto(socialPostBucketCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create a social post bucket



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
	socialPostBucketCreateDto := *openapiclient.NewSocialPostBucketCreateDto() // SocialPostBucketCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostBucketsAPI.CreateSocialPostBucketAsync(context.Background()).TenantId(tenantId).SocialPostBucketCreateDto(socialPostBucketCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.CreateSocialPostBucketAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSocialPostBucketAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.CreateSocialPostBucketAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSocialPostBucketAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **socialPostBucketCreateDto** | [**SocialPostBucketCreateDto**](SocialPostBucketCreateDto.md) |  | 
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


## DeleteSocialPostBucketAsync

> EmptyEnvelope DeleteSocialPostBucketAsync(ctx, socialpostbucketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a social post bucket



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
	socialpostbucketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostBucketsAPI.DeleteSocialPostBucketAsync(context.Background(), socialpostbucketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.DeleteSocialPostBucketAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSocialPostBucketAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.DeleteSocialPostBucketAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialpostbucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSocialPostBucketAsyncRequest struct via the builder pattern


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


## GetSocialPostBucketDetailsAsync

> SocialPostBucketDtoEnvelope GetSocialPostBucketDetailsAsync(ctx, socialpostbucketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post bucket by ID



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
	socialpostbucketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostBucketsAPI.GetSocialPostBucketDetailsAsync(context.Background(), socialpostbucketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.GetSocialPostBucketDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostBucketDetailsAsync`: SocialPostBucketDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.GetSocialPostBucketDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialpostbucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostBucketDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostBucketDtoEnvelope**](SocialPostBucketDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostBucketsCountAsync

> Int32Envelope GetSocialPostBucketsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post buckets count



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
	resp, r, err := apiClient.SocialPostBucketsAPI.GetSocialPostBucketsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.GetSocialPostBucketsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostBucketsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.GetSocialPostBucketsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostBucketsCountAsyncRequest struct via the builder pattern


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


## GetSocialPostBucketsODataAsync

> SocialPostBucketDtoListEnvelope GetSocialPostBucketsODataAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get social post buckets



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
	resp, r, err := apiClient.SocialPostBucketsAPI.GetSocialPostBucketsODataAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.GetSocialPostBucketsODataAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostBucketsODataAsync`: SocialPostBucketDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.GetSocialPostBucketsODataAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostBucketsODataAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostBucketDtoListEnvelope**](SocialPostBucketDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSocialPostBucketAsync

> EmptyEnvelope UpdateSocialPostBucketAsync(ctx, socialpostbucketId).TenantId(tenantId).SocialPostBucketUpdateDto(socialPostBucketUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update a social post bucket



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
	socialpostbucketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialPostBucketUpdateDto := *openapiclient.NewSocialPostBucketUpdateDto() // SocialPostBucketUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialPostBucketsAPI.UpdateSocialPostBucketAsync(context.Background(), socialpostbucketId).TenantId(tenantId).SocialPostBucketUpdateDto(socialPostBucketUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.UpdateSocialPostBucketAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSocialPostBucketAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.UpdateSocialPostBucketAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialpostbucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSocialPostBucketAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **socialPostBucketUpdateDto** | [**SocialPostBucketUpdateDto**](SocialPostBucketUpdateDto.md) |  | 
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

