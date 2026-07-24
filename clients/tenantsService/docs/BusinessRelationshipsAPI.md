# \BusinessRelationshipsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBusinessRelationshipAsync**](BusinessRelationshipsAPI.md#CreateBusinessRelationshipAsync) | **Post** /api/v2/TenantsService/BusinessRelationships | Create a business relationship
[**DeleteBusinessRelationshipAsync**](BusinessRelationshipsAPI.md#DeleteBusinessRelationshipAsync) | **Delete** /api/v2/TenantsService/BusinessRelationships/{businessRelationshipId} | Delete a business relationship
[**GetBusinessRelationshipByIdAsync**](BusinessRelationshipsAPI.md#GetBusinessRelationshipByIdAsync) | **Get** /api/v2/TenantsService/BusinessRelationships/{businessRelationshipId} | Get business relationship by ID
[**GetBusinessRelationshipsAsync**](BusinessRelationshipsAPI.md#GetBusinessRelationshipsAsync) | **Get** /api/v2/TenantsService/BusinessRelationships | Get business relationships
[**GetBusinessRelationshipsCountAsync**](BusinessRelationshipsAPI.md#GetBusinessRelationshipsCountAsync) | **Get** /api/v2/TenantsService/BusinessRelationships/Count | Get business relationships count
[**UpdateBusinessRelationshipAsync**](BusinessRelationshipsAPI.md#UpdateBusinessRelationshipAsync) | **Put** /api/v2/TenantsService/BusinessRelationships/{businessRelationshipId} | Update a business relationship



## CreateBusinessRelationshipAsync

> EmptyEnvelope CreateBusinessRelationshipAsync(ctx).TenantId(tenantId).BusinessRelationshipCreateDto(businessRelationshipCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create a business relationship



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
	businessRelationshipCreateDto := *openapiclient.NewBusinessRelationshipCreateDto() // BusinessRelationshipCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessRelationshipsAPI.CreateBusinessRelationshipAsync(context.Background()).TenantId(tenantId).BusinessRelationshipCreateDto(businessRelationshipCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessRelationshipsAPI.CreateBusinessRelationshipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBusinessRelationshipAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessRelationshipsAPI.CreateBusinessRelationshipAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBusinessRelationshipAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **businessRelationshipCreateDto** | [**BusinessRelationshipCreateDto**](BusinessRelationshipCreateDto.md) |  | 
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


## DeleteBusinessRelationshipAsync

> EmptyEnvelope DeleteBusinessRelationshipAsync(ctx, businessRelationshipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a business relationship



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
	businessRelationshipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessRelationshipsAPI.DeleteBusinessRelationshipAsync(context.Background(), businessRelationshipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessRelationshipsAPI.DeleteBusinessRelationshipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBusinessRelationshipAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessRelationshipsAPI.DeleteBusinessRelationshipAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessRelationshipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessRelationshipAsyncRequest struct via the builder pattern


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


## GetBusinessRelationshipByIdAsync

> BusinessRelationshipDtoEnvelope GetBusinessRelationshipByIdAsync(ctx, businessRelationshipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get business relationship by ID



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
	businessRelationshipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessRelationshipsAPI.GetBusinessRelationshipByIdAsync(context.Background(), businessRelationshipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessRelationshipsAPI.GetBusinessRelationshipByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBusinessRelationshipByIdAsync`: BusinessRelationshipDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessRelationshipsAPI.GetBusinessRelationshipByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessRelationshipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessRelationshipByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BusinessRelationshipDtoEnvelope**](BusinessRelationshipDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessRelationshipsAsync

> BusinessRelationshipDtoListEnvelope GetBusinessRelationshipsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get business relationships



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
	resp, r, err := apiClient.BusinessRelationshipsAPI.GetBusinessRelationshipsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessRelationshipsAPI.GetBusinessRelationshipsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBusinessRelationshipsAsync`: BusinessRelationshipDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessRelationshipsAPI.GetBusinessRelationshipsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessRelationshipsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BusinessRelationshipDtoListEnvelope**](BusinessRelationshipDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessRelationshipsCountAsync

> Int32Envelope GetBusinessRelationshipsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get business relationships count



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
	resp, r, err := apiClient.BusinessRelationshipsAPI.GetBusinessRelationshipsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessRelationshipsAPI.GetBusinessRelationshipsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBusinessRelationshipsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessRelationshipsAPI.GetBusinessRelationshipsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessRelationshipsCountAsyncRequest struct via the builder pattern


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


## UpdateBusinessRelationshipAsync

> EmptyEnvelope UpdateBusinessRelationshipAsync(ctx, businessRelationshipId).TenantId(tenantId).BusinessRelationshipUpdateDto(businessRelationshipUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update a business relationship



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
	businessRelationshipId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	businessRelationshipUpdateDto := *openapiclient.NewBusinessRelationshipUpdateDto() // BusinessRelationshipUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessRelationshipsAPI.UpdateBusinessRelationshipAsync(context.Background(), businessRelationshipId).TenantId(tenantId).BusinessRelationshipUpdateDto(businessRelationshipUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessRelationshipsAPI.UpdateBusinessRelationshipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBusinessRelationshipAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessRelationshipsAPI.UpdateBusinessRelationshipAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessRelationshipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessRelationshipAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **businessRelationshipUpdateDto** | [**BusinessRelationshipUpdateDto**](BusinessRelationshipUpdateDto.md) |  | 
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

