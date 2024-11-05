# \SocialPostBucketsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceSocialPostBucketsCountGet**](SocialPostBucketsAPI.md#ApiV2MarketingServiceSocialPostBucketsCountGet) | **Get** /api/v2/MarketingService/SocialPostBuckets/Count | 
[**ApiV2MarketingServiceSocialPostBucketsGet**](SocialPostBucketsAPI.md#ApiV2MarketingServiceSocialPostBucketsGet) | **Get** /api/v2/MarketingService/SocialPostBuckets | 
[**ApiV2MarketingServiceSocialPostBucketsPost**](SocialPostBucketsAPI.md#ApiV2MarketingServiceSocialPostBucketsPost) | **Post** /api/v2/MarketingService/SocialPostBuckets | 
[**ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDelete**](SocialPostBucketsAPI.md#ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDelete) | **Delete** /api/v2/MarketingService/SocialPostBuckets/{socialpostbucketId} | 
[**ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGet**](SocialPostBucketsAPI.md#ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGet) | **Get** /api/v2/MarketingService/SocialPostBuckets/{socialpostbucketId} | 
[**ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPut**](SocialPostBucketsAPI.md#ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPut) | **Put** /api/v2/MarketingService/SocialPostBuckets/{socialpostbucketId} | 



## ApiV2MarketingServiceSocialPostBucketsCountGet

> Int32Envelope ApiV2MarketingServiceSocialPostBucketsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialPostBucketsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialPostBucketsCountGetRequest struct via the builder pattern


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


## ApiV2MarketingServiceSocialPostBucketsGet

> SocialPostBucketDtoListEnvelope ApiV2MarketingServiceSocialPostBucketsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialPostBucketsGet`: SocialPostBucketDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialPostBucketsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostBucketDtoListEnvelope**](SocialPostBucketDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceSocialPostBucketsPost

> EmptyEnvelope ApiV2MarketingServiceSocialPostBucketsPost(ctx).TenantId(tenantId).SocialPostBucketCreateDto(socialPostBucketCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsPost(context.Background()).TenantId(tenantId).SocialPostBucketCreateDto(socialPostBucketCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialPostBucketsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialPostBucketsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **socialPostBucketCreateDto** | [**SocialPostBucketCreateDto**](SocialPostBucketCreateDto.md) |  | 
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


## ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDelete

> EmptyEnvelope ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDelete(ctx, socialpostbucketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDelete(context.Background(), socialpostbucketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialpostbucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdDeleteRequest struct via the builder pattern


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


## ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGet

> SocialPostBucketDtoEnvelope ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGet(ctx, socialpostbucketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGet(context.Background(), socialpostbucketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGet`: SocialPostBucketDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialpostbucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialPostBucketDtoEnvelope**](SocialPostBucketDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPut

> EmptyEnvelope ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPut(ctx, socialpostbucketId).TenantId(tenantId).SocialPostBucketUpdateDto(socialPostBucketUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPut(context.Background(), socialpostbucketId).TenantId(tenantId).SocialPostBucketUpdateDto(socialPostBucketUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SocialPostBucketsAPI.ApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialpostbucketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceSocialPostBucketsSocialpostbucketIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **socialPostBucketUpdateDto** | [**SocialPostBucketUpdateDto**](SocialPostBucketUpdateDto.md) |  | 
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

