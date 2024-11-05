# \SupportTicketPrioritiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SupportServiceSupportTicketPrioritiesCountGet**](SupportTicketPrioritiesAPI.md#ApiV2SupportServiceSupportTicketPrioritiesCountGet) | **Get** /api/v2/SupportService/SupportTicketPriorities/Count | 
[**ApiV2SupportServiceSupportTicketPrioritiesGet**](SupportTicketPrioritiesAPI.md#ApiV2SupportServiceSupportTicketPrioritiesGet) | **Get** /api/v2/SupportService/SupportTicketPriorities | 
[**ApiV2SupportServiceSupportTicketPrioritiesPost**](SupportTicketPrioritiesAPI.md#ApiV2SupportServiceSupportTicketPrioritiesPost) | **Post** /api/v2/SupportService/SupportTicketPriorities | 
[**ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDelete**](SupportTicketPrioritiesAPI.md#ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDelete) | **Delete** /api/v2/SupportService/SupportTicketPriorities/{supportTicketPriorityId} | 
[**ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGet**](SupportTicketPrioritiesAPI.md#ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGet) | **Get** /api/v2/SupportService/SupportTicketPriorities/{supportTicketPriorityId} | 
[**ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPut**](SupportTicketPrioritiesAPI.md#ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPut) | **Put** /api/v2/SupportService/SupportTicketPriorities/{supportTicketPriorityId} | 



## ApiV2SupportServiceSupportTicketPrioritiesCountGet

> Int32Envelope ApiV2SupportServiceSupportTicketPrioritiesCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketPrioritiesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketPrioritiesCountGetRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportTicketPrioritiesGet

> SupportTicketPriorityDtoListEnvelope ApiV2SupportServiceSupportTicketPrioritiesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketPrioritiesGet`: SupportTicketPriorityDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketPrioritiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketPriorityDtoListEnvelope**](SupportTicketPriorityDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketPrioritiesPost

> EmptyEnvelope ApiV2SupportServiceSupportTicketPrioritiesPost(ctx).SupportTicketPriorityCreateDto(supportTicketPriorityCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketPriorityCreateDto := *openapiclient.NewSupportTicketPriorityCreateDto() // SupportTicketPriorityCreateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesPost(context.Background()).SupportTicketPriorityCreateDto(supportTicketPriorityCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketPrioritiesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketPrioritiesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportTicketPriorityCreateDto** | [**SupportTicketPriorityCreateDto**](SupportTicketPriorityCreateDto.md) |  | 
 **tenantId** | **string** |  | 
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


## ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDelete

> EmptyEnvelope ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDelete(ctx, supportTicketPriorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketPriorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDelete(context.Background(), supportTicketPriorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketPriorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdDeleteRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGet

> SupportTicketPriorityDtoEnvelope ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGet(ctx, supportTicketPriorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketPriorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGet(context.Background(), supportTicketPriorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGet`: SupportTicketPriorityDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketPriorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketPriorityDtoEnvelope**](SupportTicketPriorityDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPut

> EmptyEnvelope ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPut(ctx, supportTicketPriorityId).SupportTicketPriorityUpdateDto(supportTicketPriorityUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketPriorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketPriorityUpdateDto := *openapiclient.NewSupportTicketPriorityUpdateDto() // SupportTicketPriorityUpdateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPut(context.Background(), supportTicketPriorityId).SupportTicketPriorityUpdateDto(supportTicketPriorityUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.ApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketPriorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketPrioritiesSupportTicketPriorityIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportTicketPriorityUpdateDto** | [**SupportTicketPriorityUpdateDto**](SupportTicketPriorityUpdateDto.md) |  | 
 **tenantId** | **string** |  | 
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

