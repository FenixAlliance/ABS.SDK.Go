# \WarrantyRequestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWarrantyRequestAsync**](WarrantyRequestsAPI.md#CreateWarrantyRequestAsync) | **Post** /api/v2/SupportService/WarrantyRequests | Create a warranty request
[**DeleteWarrantyRequestAsync**](WarrantyRequestsAPI.md#DeleteWarrantyRequestAsync) | **Delete** /api/v2/SupportService/WarrantyRequests/{warrantyRequestId} | Delete a warranty request
[**GetWarrantyRequestAsync**](WarrantyRequestsAPI.md#GetWarrantyRequestAsync) | **Get** /api/v2/SupportService/WarrantyRequests/{warrantyRequestId} | Retrieve a warranty request by ID
[**GetWarrantyRequestsAsync**](WarrantyRequestsAPI.md#GetWarrantyRequestsAsync) | **Get** /api/v2/SupportService/WarrantyRequests | Retrieve warranty requests
[**GetWarrantyRequestsCountAsync**](WarrantyRequestsAPI.md#GetWarrantyRequestsCountAsync) | **Get** /api/v2/SupportService/WarrantyRequests/Count | Get warranty requests count
[**PatchWarrantyRequestAsync**](WarrantyRequestsAPI.md#PatchWarrantyRequestAsync) | **Patch** /api/v2/SupportService/WarrantyRequests/{warrantyRequestId} | Patch a warranty request
[**UpdateWarrantyRequestAsync**](WarrantyRequestsAPI.md#UpdateWarrantyRequestAsync) | **Put** /api/v2/SupportService/WarrantyRequests/{warrantyRequestId} | Update a warranty request



## CreateWarrantyRequestAsync

> EmptyEnvelope CreateWarrantyRequestAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WarrantyRequestCreateDto(warrantyRequestCreateDto).Execute()

Create a warranty request

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
	warrantyRequestCreateDto := *openapiclient.NewWarrantyRequestCreateDto("Title_example") // WarrantyRequestCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyRequestsAPI.CreateWarrantyRequestAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WarrantyRequestCreateDto(warrantyRequestCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyRequestsAPI.CreateWarrantyRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWarrantyRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyRequestsAPI.CreateWarrantyRequestAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarrantyRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **warrantyRequestCreateDto** | [**WarrantyRequestCreateDto**](WarrantyRequestCreateDto.md) |  | 

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


## DeleteWarrantyRequestAsync

> EmptyEnvelope DeleteWarrantyRequestAsync(ctx, warrantyRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a warranty request

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
	warrantyRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyRequestsAPI.DeleteWarrantyRequestAsync(context.Background(), warrantyRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyRequestsAPI.DeleteWarrantyRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWarrantyRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyRequestsAPI.DeleteWarrantyRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantyRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWarrantyRequestAsyncRequest struct via the builder pattern


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


## GetWarrantyRequestAsync

> WarrantyRequestDtoEnvelope GetWarrantyRequestAsync(ctx, warrantyRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a warranty request by ID

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
	warrantyRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyRequestsAPI.GetWarrantyRequestAsync(context.Background(), warrantyRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyRequestsAPI.GetWarrantyRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarrantyRequestAsync`: WarrantyRequestDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyRequestsAPI.GetWarrantyRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantyRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarrantyRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WarrantyRequestDtoEnvelope**](WarrantyRequestDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarrantyRequestsAsync

> WarrantyRequestDtoListEnvelope GetWarrantyRequestsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve warranty requests

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
	resp, r, err := apiClient.WarrantyRequestsAPI.GetWarrantyRequestsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyRequestsAPI.GetWarrantyRequestsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarrantyRequestsAsync`: WarrantyRequestDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyRequestsAPI.GetWarrantyRequestsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWarrantyRequestsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WarrantyRequestDtoListEnvelope**](WarrantyRequestDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarrantyRequestsCountAsync

> Int32Envelope GetWarrantyRequestsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get warranty requests count

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
	resp, r, err := apiClient.WarrantyRequestsAPI.GetWarrantyRequestsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyRequestsAPI.GetWarrantyRequestsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarrantyRequestsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyRequestsAPI.GetWarrantyRequestsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWarrantyRequestsCountAsyncRequest struct via the builder pattern


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


## PatchWarrantyRequestAsync

> EmptyEnvelope PatchWarrantyRequestAsync(ctx, warrantyRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a warranty request



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
	warrantyRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyRequestsAPI.PatchWarrantyRequestAsync(context.Background(), warrantyRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyRequestsAPI.PatchWarrantyRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWarrantyRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyRequestsAPI.PatchWarrantyRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantyRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWarrantyRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## UpdateWarrantyRequestAsync

> EmptyEnvelope UpdateWarrantyRequestAsync(ctx, warrantyRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WarrantyRequestUpdateDto(warrantyRequestUpdateDto).Execute()

Update a warranty request

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
	warrantyRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	warrantyRequestUpdateDto := *openapiclient.NewWarrantyRequestUpdateDto() // WarrantyRequestUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyRequestsAPI.UpdateWarrantyRequestAsync(context.Background(), warrantyRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WarrantyRequestUpdateDto(warrantyRequestUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyRequestsAPI.UpdateWarrantyRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWarrantyRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyRequestsAPI.UpdateWarrantyRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantyRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWarrantyRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **warrantyRequestUpdateDto** | [**WarrantyRequestUpdateDto**](WarrantyRequestUpdateDto.md) |  | 

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

