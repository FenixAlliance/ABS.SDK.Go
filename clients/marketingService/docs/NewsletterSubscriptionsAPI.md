# \NewsletterSubscriptionsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewsletterSubscriptionAsync**](NewsletterSubscriptionsAPI.md#CreateNewsletterSubscriptionAsync) | **Post** /api/v2/MarketingService/NewsletterSubscriptions | Create a newsletter subscription
[**DeleteNewsletterSubscriptionAsync**](NewsletterSubscriptionsAPI.md#DeleteNewsletterSubscriptionAsync) | **Delete** /api/v2/MarketingService/NewsletterSubscriptions/{newsletterSubscriptionId} | Delete a newsletter subscription
[**GetNewsletterSubscriptionByIdAsync**](NewsletterSubscriptionsAPI.md#GetNewsletterSubscriptionByIdAsync) | **Get** /api/v2/MarketingService/NewsletterSubscriptions/{newsletterSubscriptionId} | Get newsletter subscription by ID
[**GetNewsletterSubscriptionsAsync**](NewsletterSubscriptionsAPI.md#GetNewsletterSubscriptionsAsync) | **Get** /api/v2/MarketingService/NewsletterSubscriptions | Get newsletter subscriptions
[**GetNewsletterSubscriptionsCountAsync**](NewsletterSubscriptionsAPI.md#GetNewsletterSubscriptionsCountAsync) | **Get** /api/v2/MarketingService/NewsletterSubscriptions/Count | Get newsletter subscriptions count
[**UpdateNewsletterSubscriptionAsync**](NewsletterSubscriptionsAPI.md#UpdateNewsletterSubscriptionAsync) | **Put** /api/v2/MarketingService/NewsletterSubscriptions/{newsletterSubscriptionId} | Update a newsletter subscription



## CreateNewsletterSubscriptionAsync

> EmptyEnvelope CreateNewsletterSubscriptionAsync(ctx).TenantId(tenantId).NewsletterSubscriptionCreateDto(newsletterSubscriptionCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create a newsletter subscription



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
	newsletterSubscriptionCreateDto := *openapiclient.NewNewsletterSubscriptionCreateDto() // NewsletterSubscriptionCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsletterSubscriptionsAPI.CreateNewsletterSubscriptionAsync(context.Background()).TenantId(tenantId).NewsletterSubscriptionCreateDto(newsletterSubscriptionCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsletterSubscriptionsAPI.CreateNewsletterSubscriptionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewsletterSubscriptionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewsletterSubscriptionsAPI.CreateNewsletterSubscriptionAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewsletterSubscriptionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **newsletterSubscriptionCreateDto** | [**NewsletterSubscriptionCreateDto**](NewsletterSubscriptionCreateDto.md) |  | 
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


## DeleteNewsletterSubscriptionAsync

> EmptyEnvelope DeleteNewsletterSubscriptionAsync(ctx, newsletterSubscriptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a newsletter subscription



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
	newsletterSubscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsletterSubscriptionsAPI.DeleteNewsletterSubscriptionAsync(context.Background(), newsletterSubscriptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsletterSubscriptionsAPI.DeleteNewsletterSubscriptionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNewsletterSubscriptionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewsletterSubscriptionsAPI.DeleteNewsletterSubscriptionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newsletterSubscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNewsletterSubscriptionAsyncRequest struct via the builder pattern


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


## GetNewsletterSubscriptionByIdAsync

> NewsletterSubscriptionDtoEnvelope GetNewsletterSubscriptionByIdAsync(ctx, newsletterSubscriptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get newsletter subscription by ID



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
	newsletterSubscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsletterSubscriptionsAPI.GetNewsletterSubscriptionByIdAsync(context.Background(), newsletterSubscriptionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsletterSubscriptionsAPI.GetNewsletterSubscriptionByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewsletterSubscriptionByIdAsync`: NewsletterSubscriptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewsletterSubscriptionsAPI.GetNewsletterSubscriptionByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newsletterSubscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNewsletterSubscriptionByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**NewsletterSubscriptionDtoEnvelope**](NewsletterSubscriptionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewsletterSubscriptionsAsync

> NewsletterSubscriptionDtoListEnvelope GetNewsletterSubscriptionsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).NewsletterSubscriptionDtoCollectionQueryParameters(newsletterSubscriptionDtoCollectionQueryParameters).Execute()

Get newsletter subscriptions



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
	newsletterSubscriptionDtoCollectionQueryParameters := *openapiclient.NewNewsletterSubscriptionDtoCollectionQueryParameters() // NewsletterSubscriptionDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsletterSubscriptionsAPI.GetNewsletterSubscriptionsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).NewsletterSubscriptionDtoCollectionQueryParameters(newsletterSubscriptionDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsletterSubscriptionsAPI.GetNewsletterSubscriptionsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewsletterSubscriptionsAsync`: NewsletterSubscriptionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewsletterSubscriptionsAPI.GetNewsletterSubscriptionsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNewsletterSubscriptionsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **newsletterSubscriptionDtoCollectionQueryParameters** | [**NewsletterSubscriptionDtoCollectionQueryParameters**](NewsletterSubscriptionDtoCollectionQueryParameters.md) |  | 

### Return type

[**NewsletterSubscriptionDtoListEnvelope**](NewsletterSubscriptionDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNewsletterSubscriptionsCountAsync

> Int32Envelope GetNewsletterSubscriptionsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).NewsletterSubscriptionDtoCollectionQueryParameters(newsletterSubscriptionDtoCollectionQueryParameters).Execute()

Get newsletter subscriptions count



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
	newsletterSubscriptionDtoCollectionQueryParameters := *openapiclient.NewNewsletterSubscriptionDtoCollectionQueryParameters() // NewsletterSubscriptionDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsletterSubscriptionsAPI.GetNewsletterSubscriptionsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).NewsletterSubscriptionDtoCollectionQueryParameters(newsletterSubscriptionDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsletterSubscriptionsAPI.GetNewsletterSubscriptionsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNewsletterSubscriptionsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `NewsletterSubscriptionsAPI.GetNewsletterSubscriptionsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNewsletterSubscriptionsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **newsletterSubscriptionDtoCollectionQueryParameters** | [**NewsletterSubscriptionDtoCollectionQueryParameters**](NewsletterSubscriptionDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNewsletterSubscriptionAsync

> EmptyEnvelope UpdateNewsletterSubscriptionAsync(ctx, newsletterSubscriptionId).TenantId(tenantId).NewsletterSubscriptionUpdateDto(newsletterSubscriptionUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update a newsletter subscription



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
	newsletterSubscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	newsletterSubscriptionUpdateDto := *openapiclient.NewNewsletterSubscriptionUpdateDto() // NewsletterSubscriptionUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsletterSubscriptionsAPI.UpdateNewsletterSubscriptionAsync(context.Background(), newsletterSubscriptionId).TenantId(tenantId).NewsletterSubscriptionUpdateDto(newsletterSubscriptionUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsletterSubscriptionsAPI.UpdateNewsletterSubscriptionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNewsletterSubscriptionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `NewsletterSubscriptionsAPI.UpdateNewsletterSubscriptionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**newsletterSubscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNewsletterSubscriptionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **newsletterSubscriptionUpdateDto** | [**NewsletterSubscriptionUpdateDto**](NewsletterSubscriptionUpdateDto.md) |  | 
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

