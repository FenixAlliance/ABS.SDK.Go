# \MarketingListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceMarketingListsCountGet**](MarketingListsAPI.md#ApiV2MarketingServiceMarketingListsCountGet) | **Get** /api/v2/MarketingService/MarketingLists/Count | 
[**ApiV2MarketingServiceMarketingListsGet**](MarketingListsAPI.md#ApiV2MarketingServiceMarketingListsGet) | **Get** /api/v2/MarketingService/MarketingLists | 
[**ApiV2MarketingServiceMarketingListsMarketinglistIdDelete**](MarketingListsAPI.md#ApiV2MarketingServiceMarketingListsMarketinglistIdDelete) | **Delete** /api/v2/MarketingService/MarketingLists/{marketinglistId} | 
[**ApiV2MarketingServiceMarketingListsMarketinglistIdGet**](MarketingListsAPI.md#ApiV2MarketingServiceMarketingListsMarketinglistIdGet) | **Get** /api/v2/MarketingService/MarketingLists/{marketinglistId} | 
[**ApiV2MarketingServiceMarketingListsMarketinglistIdPut**](MarketingListsAPI.md#ApiV2MarketingServiceMarketingListsMarketinglistIdPut) | **Put** /api/v2/MarketingService/MarketingLists/{marketinglistId} | 
[**ApiV2MarketingServiceMarketingListsPost**](MarketingListsAPI.md#ApiV2MarketingServiceMarketingListsPost) | **Post** /api/v2/MarketingService/MarketingLists | 



## ApiV2MarketingServiceMarketingListsCountGet

> Int32Envelope ApiV2MarketingServiceMarketingListsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.MarketingListsAPI.ApiV2MarketingServiceMarketingListsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingListsAPI.ApiV2MarketingServiceMarketingListsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingListsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingListsAPI.ApiV2MarketingServiceMarketingListsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingListsCountGetRequest struct via the builder pattern


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


## ApiV2MarketingServiceMarketingListsGet

> MarketingListDtoListEnvelope ApiV2MarketingServiceMarketingListsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.MarketingListsAPI.ApiV2MarketingServiceMarketingListsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingListsAPI.ApiV2MarketingServiceMarketingListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingListsGet`: MarketingListDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingListsAPI.ApiV2MarketingServiceMarketingListsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingListsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**MarketingListDtoListEnvelope**](MarketingListDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceMarketingListsMarketinglistIdDelete

> EmptyEnvelope ApiV2MarketingServiceMarketingListsMarketinglistIdDelete(ctx, marketinglistId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	marketinglistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdDelete(context.Background(), marketinglistId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingListsMarketinglistIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketinglistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingListsMarketinglistIdDeleteRequest struct via the builder pattern


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


## ApiV2MarketingServiceMarketingListsMarketinglistIdGet

> MarketingListDtoEnvelope ApiV2MarketingServiceMarketingListsMarketinglistIdGet(ctx, marketinglistId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	marketinglistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdGet(context.Background(), marketinglistId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingListsMarketinglistIdGet`: MarketingListDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketinglistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingListsMarketinglistIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**MarketingListDtoEnvelope**](MarketingListDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceMarketingListsMarketinglistIdPut

> EmptyEnvelope ApiV2MarketingServiceMarketingListsMarketinglistIdPut(ctx, marketinglistId).TenantId(tenantId).MarketingListUpdateDto(marketingListUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	marketinglistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	marketingListUpdateDto := *openapiclient.NewMarketingListUpdateDto() // MarketingListUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdPut(context.Background(), marketinglistId).TenantId(tenantId).MarketingListUpdateDto(marketingListUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingListsMarketinglistIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingListsAPI.ApiV2MarketingServiceMarketingListsMarketinglistIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketinglistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingListsMarketinglistIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **marketingListUpdateDto** | [**MarketingListUpdateDto**](MarketingListUpdateDto.md) |  | 
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


## ApiV2MarketingServiceMarketingListsPost

> EmptyEnvelope ApiV2MarketingServiceMarketingListsPost(ctx).TenantId(tenantId).MarketingListCreateDto(marketingListCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	marketingListCreateDto := *openapiclient.NewMarketingListCreateDto() // MarketingListCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingListsAPI.ApiV2MarketingServiceMarketingListsPost(context.Background()).TenantId(tenantId).MarketingListCreateDto(marketingListCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingListsAPI.ApiV2MarketingServiceMarketingListsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceMarketingListsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingListsAPI.ApiV2MarketingServiceMarketingListsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceMarketingListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **marketingListCreateDto** | [**MarketingListCreateDto**](MarketingListCreateDto.md) |  | 
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

