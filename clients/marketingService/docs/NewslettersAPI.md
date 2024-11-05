# \NewslettersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceNewslettersCountGet**](NewslettersAPI.md#ApiV2MarketingServiceNewslettersCountGet) | **Get** /api/v2/MarketingService/Newsletters/Count | 
[**ApiV2MarketingServiceNewslettersGet**](NewslettersAPI.md#ApiV2MarketingServiceNewslettersGet) | **Get** /api/v2/MarketingService/Newsletters | 
[**ApiV2MarketingServiceNewslettersNewsletterIdDelete**](NewslettersAPI.md#ApiV2MarketingServiceNewslettersNewsletterIdDelete) | **Delete** /api/v2/MarketingService/Newsletters/{newsletterId} | 
[**ApiV2MarketingServiceNewslettersNewsletterIdGet**](NewslettersAPI.md#ApiV2MarketingServiceNewslettersNewsletterIdGet) | **Get** /api/v2/MarketingService/Newsletters/{newsletterId} | 
[**ApiV2MarketingServiceNewslettersNewsletterIdPut**](NewslettersAPI.md#ApiV2MarketingServiceNewslettersNewsletterIdPut) | **Put** /api/v2/MarketingService/Newsletters/{newsletterId} | 
[**ApiV2MarketingServiceNewslettersPost**](NewslettersAPI.md#ApiV2MarketingServiceNewslettersPost) | **Post** /api/v2/MarketingService/Newsletters | 



## ApiV2MarketingServiceNewslettersCountGet

> Int32Envelope ApiV2MarketingServiceNewslettersCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewslettersAPI.ApiV2MarketingServiceNewslettersCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceNewslettersCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `NewslettersAPI.ApiV2MarketingServiceNewslettersCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceNewslettersCountGetRequest struct via the builder pattern


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


## ApiV2MarketingServiceNewslettersGet

> ApiV2MarketingServiceNewslettersGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	r, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewslettersAPI.ApiV2MarketingServiceNewslettersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceNewslettersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceNewslettersNewsletterIdDelete

> EmptyEnvelope ApiV2MarketingServiceNewslettersNewsletterIdDelete(ctx, newsletterId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	newsletterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdDelete(context.Background(), newsletterId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceNewslettersNewsletterIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newsletterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceNewslettersNewsletterIdDeleteRequest struct via the builder pattern


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


## ApiV2MarketingServiceNewslettersNewsletterIdGet

> NewsletterDtoEnvelope ApiV2MarketingServiceNewslettersNewsletterIdGet(ctx, newsletterId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	newsletterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdGet(context.Background(), newsletterId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceNewslettersNewsletterIdGet`: NewsletterDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newsletterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceNewslettersNewsletterIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**NewsletterDtoEnvelope**](NewsletterDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceNewslettersNewsletterIdPut

> EmptyEnvelope ApiV2MarketingServiceNewslettersNewsletterIdPut(ctx, newsletterId).TenantId(tenantId).NewsletterUpdateDto(newsletterUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	newsletterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	newsletterUpdateDto := *openapiclient.NewNewsletterUpdateDto() // NewsletterUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdPut(context.Background(), newsletterId).TenantId(tenantId).NewsletterUpdateDto(newsletterUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceNewslettersNewsletterIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewslettersAPI.ApiV2MarketingServiceNewslettersNewsletterIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newsletterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceNewslettersNewsletterIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **newsletterUpdateDto** | [**NewsletterUpdateDto**](NewsletterUpdateDto.md) |  | 
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


## ApiV2MarketingServiceNewslettersPost

> EmptyEnvelope ApiV2MarketingServiceNewslettersPost(ctx).TenantId(tenantId).NewsletterCreateDto(newsletterCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	newsletterCreateDto := *openapiclient.NewNewsletterCreateDto() // NewsletterCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewslettersAPI.ApiV2MarketingServiceNewslettersPost(context.Background()).TenantId(tenantId).NewsletterCreateDto(newsletterCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewslettersAPI.ApiV2MarketingServiceNewslettersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceNewslettersPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewslettersAPI.ApiV2MarketingServiceNewslettersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceNewslettersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **newsletterCreateDto** | [**NewsletterCreateDto**](NewsletterCreateDto.md) |  | 
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

