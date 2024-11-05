# \MarketingCampaignsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceMarketingCampaignsCountGet**](MarketingCampaignsAPI.md#ApiV2MarketingServiceMarketingCampaignsCountGet) | **Get** /api/v2/MarketingService/MarketingCampaigns/Count | 
[**ApiV2MarketingServiceMarketingCampaignsGet**](MarketingCampaignsAPI.md#ApiV2MarketingServiceMarketingCampaignsGet) | **Get** /api/v2/MarketingService/MarketingCampaigns | 
[**ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDelete**](MarketingCampaignsAPI.md#ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDelete) | **Delete** /api/v2/MarketingService/MarketingCampaigns/{marketingcampaignId} | 
[**ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGet**](MarketingCampaignsAPI.md#ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGet) | **Get** /api/v2/MarketingService/MarketingCampaigns/{marketingcampaignId} | 
[**ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPut**](MarketingCampaignsAPI.md#ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPut) | **Put** /api/v2/MarketingService/MarketingCampaigns/{marketingcampaignId} | 
[**ApiV2MarketingServiceMarketingCampaignsPost**](MarketingCampaignsAPI.md#ApiV2MarketingServiceMarketingCampaignsPost) | **Post** /api/v2/MarketingService/MarketingCampaigns | 



## ApiV2MarketingServiceMarketingCampaignsCountGet

> Int32Envelope ApiV2MarketingServiceMarketingCampaignsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingCampaignsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingCampaignsCountGetRequest struct via the builder pattern


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


## ApiV2MarketingServiceMarketingCampaignsGet

> ApiV2MarketingServiceMarketingCampaignsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	r, err := apiClient.MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingCampaignsGetRequest struct via the builder pattern


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


## ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDelete

> EmptyEnvelope ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDelete(ctx, marketingcampaignId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	marketingcampaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDelete(context.Background(), marketingcampaignId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingcampaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdDeleteRequest struct via the builder pattern


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


## ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGet

> MarketingCampaignDtoEnvelope ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGet(ctx, marketingcampaignId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	marketingcampaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGet(context.Background(), marketingcampaignId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGet`: MarketingCampaignDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingcampaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**MarketingCampaignDtoEnvelope**](MarketingCampaignDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPut

> EmptyEnvelope ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPut(ctx, marketingcampaignId).TenantId(tenantId).MarketingCampaignUpdateDto(marketingCampaignUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	marketingcampaignId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	marketingCampaignUpdateDto := *openapiclient.NewMarketingCampaignUpdateDto() // MarketingCampaignUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPut(context.Background(), marketingcampaignId).TenantId(tenantId).MarketingCampaignUpdateDto(marketingCampaignUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingcampaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingCampaignsMarketingcampaignIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **marketingCampaignUpdateDto** | [**MarketingCampaignUpdateDto**](MarketingCampaignUpdateDto.md) |  | 
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


## ApiV2MarketingServiceMarketingCampaignsPost

> EmptyEnvelope ApiV2MarketingServiceMarketingCampaignsPost(ctx).TenantId(tenantId).MarketingCampaignCreateDto(marketingCampaignCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	marketingCampaignCreateDto := *openapiclient.NewMarketingCampaignCreateDto() // MarketingCampaignCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsPost(context.Background()).TenantId(tenantId).MarketingCampaignCreateDto(marketingCampaignCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingCampaignsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingCampaignsAPI.ApiV2MarketingServiceMarketingCampaignsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingCampaignsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **marketingCampaignCreateDto** | [**MarketingCampaignCreateDto**](MarketingCampaignCreateDto.md) |  | 
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

