# \LeaveTypesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLeaveTypeAsync**](LeaveTypesAPI.md#CreateLeaveTypeAsync) | **Post** /api/v2/HrmsService/LeaveTypes | Create a leave type
[**DeleteLeaveTypeAsync**](LeaveTypesAPI.md#DeleteLeaveTypeAsync) | **Delete** /api/v2/HrmsService/LeaveTypes/{leaveTypeId} | Delete a leave type
[**GetLeaveTypeByIdAsync**](LeaveTypesAPI.md#GetLeaveTypeByIdAsync) | **Get** /api/v2/HrmsService/LeaveTypes/{leaveTypeId} | Get leave type by ID
[**GetLeaveTypesAsync**](LeaveTypesAPI.md#GetLeaveTypesAsync) | **Get** /api/v2/HrmsService/LeaveTypes | Get leave types
[**GetLeaveTypesCountAsync**](LeaveTypesAPI.md#GetLeaveTypesCountAsync) | **Get** /api/v2/HrmsService/LeaveTypes/Count | Count leave types
[**UpdateLeaveTypeAsync**](LeaveTypesAPI.md#UpdateLeaveTypeAsync) | **Put** /api/v2/HrmsService/LeaveTypes/{leaveTypeId} | Update a leave type



## CreateLeaveTypeAsync

> EmptyEnvelope CreateLeaveTypeAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveTypeCreateDto(leaveTypeCreateDto).Execute()

Create a leave type



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
	leaveTypeCreateDto := *openapiclient.NewLeaveTypeCreateDto("Title_example") // LeaveTypeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveTypesAPI.CreateLeaveTypeAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveTypeCreateDto(leaveTypeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveTypesAPI.CreateLeaveTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLeaveTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveTypesAPI.CreateLeaveTypeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLeaveTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **leaveTypeCreateDto** | [**LeaveTypeCreateDto**](LeaveTypeCreateDto.md) |  | 

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


## DeleteLeaveTypeAsync

> EmptyEnvelope DeleteLeaveTypeAsync(ctx, leaveTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a leave type



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
	leaveTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveTypesAPI.DeleteLeaveTypeAsync(context.Background(), leaveTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveTypesAPI.DeleteLeaveTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLeaveTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveTypesAPI.DeleteLeaveTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leaveTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLeaveTypeAsyncRequest struct via the builder pattern


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


## GetLeaveTypeByIdAsync

> LeaveTypeDtoEnvelope GetLeaveTypeByIdAsync(ctx, leaveTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get leave type by ID



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
	leaveTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveTypesAPI.GetLeaveTypeByIdAsync(context.Background(), leaveTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveTypesAPI.GetLeaveTypeByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaveTypeByIdAsync`: LeaveTypeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveTypesAPI.GetLeaveTypeByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leaveTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaveTypeByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**LeaveTypeDtoEnvelope**](LeaveTypeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaveTypesAsync

> LeaveTypeDtoListEnvelope GetLeaveTypesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveTypeDtoCollectionQueryParameters(leaveTypeDtoCollectionQueryParameters).Execute()

Get leave types



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
	leaveTypeDtoCollectionQueryParameters := *openapiclient.NewLeaveTypeDtoCollectionQueryParameters() // LeaveTypeDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveTypesAPI.GetLeaveTypesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveTypeDtoCollectionQueryParameters(leaveTypeDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveTypesAPI.GetLeaveTypesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaveTypesAsync`: LeaveTypeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveTypesAPI.GetLeaveTypesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaveTypesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **leaveTypeDtoCollectionQueryParameters** | [**LeaveTypeDtoCollectionQueryParameters**](LeaveTypeDtoCollectionQueryParameters.md) |  | 

### Return type

[**LeaveTypeDtoListEnvelope**](LeaveTypeDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaveTypesCountAsync

> Int32Envelope GetLeaveTypesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveTypeDtoCollectionQueryParameters(leaveTypeDtoCollectionQueryParameters).Execute()

Count leave types



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
	leaveTypeDtoCollectionQueryParameters := *openapiclient.NewLeaveTypeDtoCollectionQueryParameters() // LeaveTypeDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveTypesAPI.GetLeaveTypesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveTypeDtoCollectionQueryParameters(leaveTypeDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveTypesAPI.GetLeaveTypesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaveTypesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveTypesAPI.GetLeaveTypesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaveTypesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **leaveTypeDtoCollectionQueryParameters** | [**LeaveTypeDtoCollectionQueryParameters**](LeaveTypeDtoCollectionQueryParameters.md) |  | 

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


## UpdateLeaveTypeAsync

> EmptyEnvelope UpdateLeaveTypeAsync(ctx, leaveTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveTypeUpdateDto(leaveTypeUpdateDto).Execute()

Update a leave type



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
	leaveTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	leaveTypeUpdateDto := *openapiclient.NewLeaveTypeUpdateDto() // LeaveTypeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveTypesAPI.UpdateLeaveTypeAsync(context.Background(), leaveTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveTypeUpdateDto(leaveTypeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveTypesAPI.UpdateLeaveTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLeaveTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveTypesAPI.UpdateLeaveTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leaveTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLeaveTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **leaveTypeUpdateDto** | [**LeaveTypeUpdateDto**](LeaveTypeUpdateDto.md) |  | 

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

