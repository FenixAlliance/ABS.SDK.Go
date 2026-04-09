# \SupportTicketPrioritiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportTicketPriorityAsync**](SupportTicketPrioritiesAPI.md#CreateSupportTicketPriorityAsync) | **Post** /api/v2/SupportService/SupportTicketPriorities | Create a new support ticket priority
[**DeleteSupportTicketPriorityAsync**](SupportTicketPrioritiesAPI.md#DeleteSupportTicketPriorityAsync) | **Delete** /api/v2/SupportService/SupportTicketPriorities/{supportTicketPriorityId} | Delete a support ticket priority
[**GetSupportTicketPrioritiesAsync**](SupportTicketPrioritiesAPI.md#GetSupportTicketPrioritiesAsync) | **Get** /api/v2/SupportService/SupportTicketPriorities | Retrieve a list of support ticket priorities
[**GetSupportTicketPrioritiesCountAsync**](SupportTicketPrioritiesAPI.md#GetSupportTicketPrioritiesCountAsync) | **Get** /api/v2/SupportService/SupportTicketPriorities/Count | Get the count of support ticket priorities
[**GetSupportTicketPriorityAsync**](SupportTicketPrioritiesAPI.md#GetSupportTicketPriorityAsync) | **Get** /api/v2/SupportService/SupportTicketPriorities/{supportTicketPriorityId} | Retrieve a support ticket priority by ID
[**UpdateSupportTicketPriorityAsync**](SupportTicketPrioritiesAPI.md#UpdateSupportTicketPriorityAsync) | **Put** /api/v2/SupportService/SupportTicketPriorities/{supportTicketPriorityId} | Update a support ticket priority



## CreateSupportTicketPriorityAsync

> EmptyEnvelope CreateSupportTicketPriorityAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketPriorityCreateDto(supportTicketPriorityCreateDto).Execute()

Create a new support ticket priority



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
	supportTicketPriorityCreateDto := *openapiclient.NewSupportTicketPriorityCreateDto() // SupportTicketPriorityCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.CreateSupportTicketPriorityAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketPriorityCreateDto(supportTicketPriorityCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.CreateSupportTicketPriorityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportTicketPriorityAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.CreateSupportTicketPriorityAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportTicketPriorityAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportTicketPriorityCreateDto** | [**SupportTicketPriorityCreateDto**](SupportTicketPriorityCreateDto.md) |  | 

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


## DeleteSupportTicketPriorityAsync

> EmptyEnvelope DeleteSupportTicketPriorityAsync(ctx, supportTicketPriorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a support ticket priority



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
	supportTicketPriorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.DeleteSupportTicketPriorityAsync(context.Background(), supportTicketPriorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.DeleteSupportTicketPriorityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSupportTicketPriorityAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.DeleteSupportTicketPriorityAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketPriorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupportTicketPriorityAsyncRequest struct via the builder pattern


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


## GetSupportTicketPrioritiesAsync

> SupportTicketPriorityDtoListEnvelope GetSupportTicketPrioritiesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of support ticket priorities



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
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.GetSupportTicketPrioritiesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.GetSupportTicketPrioritiesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketPrioritiesAsync`: SupportTicketPriorityDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.GetSupportTicketPrioritiesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketPrioritiesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketPriorityDtoListEnvelope**](SupportTicketPriorityDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportTicketPrioritiesCountAsync

> Int32Envelope GetSupportTicketPrioritiesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of support ticket priorities



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
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.GetSupportTicketPrioritiesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.GetSupportTicketPrioritiesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketPrioritiesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.GetSupportTicketPrioritiesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketPrioritiesCountAsyncRequest struct via the builder pattern


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


## GetSupportTicketPriorityAsync

> SupportTicketPriorityDtoEnvelope GetSupportTicketPriorityAsync(ctx, supportTicketPriorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a support ticket priority by ID



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
	supportTicketPriorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.GetSupportTicketPriorityAsync(context.Background(), supportTicketPriorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.GetSupportTicketPriorityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketPriorityAsync`: SupportTicketPriorityDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.GetSupportTicketPriorityAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketPriorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketPriorityAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketPriorityDtoEnvelope**](SupportTicketPriorityDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportTicketPriorityAsync

> EmptyEnvelope UpdateSupportTicketPriorityAsync(ctx, supportTicketPriorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketPriorityUpdateDto(supportTicketPriorityUpdateDto).Execute()

Update a support ticket priority



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
	supportTicketPriorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	supportTicketPriorityUpdateDto := *openapiclient.NewSupportTicketPriorityUpdateDto() // SupportTicketPriorityUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketPrioritiesAPI.UpdateSupportTicketPriorityAsync(context.Background(), supportTicketPriorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketPriorityUpdateDto(supportTicketPriorityUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketPrioritiesAPI.UpdateSupportTicketPriorityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSupportTicketPriorityAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketPrioritiesAPI.UpdateSupportTicketPriorityAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketPriorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupportTicketPriorityAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportTicketPriorityUpdateDto** | [**SupportTicketPriorityUpdateDto**](SupportTicketPriorityUpdateDto.md) |  | 

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

