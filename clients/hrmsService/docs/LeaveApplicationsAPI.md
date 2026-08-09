# \LeaveApplicationsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLeaveApplicationAsync**](LeaveApplicationsAPI.md#CreateLeaveApplicationAsync) | **Post** /api/v2/HrmsService/LeaveApplications | Create a leave application
[**DeleteLeaveApplicationAsync**](LeaveApplicationsAPI.md#DeleteLeaveApplicationAsync) | **Delete** /api/v2/HrmsService/LeaveApplications/{leaveApplicationId} | Delete a leave application
[**GetLeaveApplicationByIdAsync**](LeaveApplicationsAPI.md#GetLeaveApplicationByIdAsync) | **Get** /api/v2/HrmsService/LeaveApplications/{leaveApplicationId} | Get leave application by ID
[**GetLeaveApplicationsAsync**](LeaveApplicationsAPI.md#GetLeaveApplicationsAsync) | **Get** /api/v2/HrmsService/LeaveApplications | Get leave applications
[**GetLeaveApplicationsCountAsync**](LeaveApplicationsAPI.md#GetLeaveApplicationsCountAsync) | **Get** /api/v2/HrmsService/LeaveApplications/Count | Count leave applications
[**PatchLeaveApplicationAsync**](LeaveApplicationsAPI.md#PatchLeaveApplicationAsync) | **Patch** /api/v2/HrmsService/LeaveApplications/{leaveApplicationId} | Patch a leave application
[**UpdateLeaveApplicationAsync**](LeaveApplicationsAPI.md#UpdateLeaveApplicationAsync) | **Put** /api/v2/HrmsService/LeaveApplications/{leaveApplicationId} | Update a leave application



## CreateLeaveApplicationAsync

> EmptyEnvelope CreateLeaveApplicationAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveApplicationCreateDto(leaveApplicationCreateDto).Execute()

Create a leave application



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
	leaveApplicationCreateDto := *openapiclient.NewLeaveApplicationCreateDto("LeaveTypeId_example", "EmployeeProfileId_example") // LeaveApplicationCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveApplicationsAPI.CreateLeaveApplicationAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveApplicationCreateDto(leaveApplicationCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveApplicationsAPI.CreateLeaveApplicationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLeaveApplicationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveApplicationsAPI.CreateLeaveApplicationAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLeaveApplicationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **leaveApplicationCreateDto** | [**LeaveApplicationCreateDto**](LeaveApplicationCreateDto.md) |  | 

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


## DeleteLeaveApplicationAsync

> EmptyEnvelope DeleteLeaveApplicationAsync(ctx, leaveApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a leave application



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
	leaveApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveApplicationsAPI.DeleteLeaveApplicationAsync(context.Background(), leaveApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveApplicationsAPI.DeleteLeaveApplicationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLeaveApplicationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveApplicationsAPI.DeleteLeaveApplicationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leaveApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLeaveApplicationAsyncRequest struct via the builder pattern


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


## GetLeaveApplicationByIdAsync

> LeaveApplicationDtoEnvelope GetLeaveApplicationByIdAsync(ctx, leaveApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get leave application by ID



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
	leaveApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveApplicationsAPI.GetLeaveApplicationByIdAsync(context.Background(), leaveApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveApplicationsAPI.GetLeaveApplicationByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaveApplicationByIdAsync`: LeaveApplicationDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveApplicationsAPI.GetLeaveApplicationByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leaveApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaveApplicationByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**LeaveApplicationDtoEnvelope**](LeaveApplicationDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaveApplicationsAsync

> LeaveApplicationDtoListEnvelope GetLeaveApplicationsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveApplicationDtoCollectionQueryParameters(leaveApplicationDtoCollectionQueryParameters).Execute()

Get leave applications



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
	leaveApplicationDtoCollectionQueryParameters := *openapiclient.NewLeaveApplicationDtoCollectionQueryParameters() // LeaveApplicationDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveApplicationsAPI.GetLeaveApplicationsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveApplicationDtoCollectionQueryParameters(leaveApplicationDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveApplicationsAPI.GetLeaveApplicationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaveApplicationsAsync`: LeaveApplicationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveApplicationsAPI.GetLeaveApplicationsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaveApplicationsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **leaveApplicationDtoCollectionQueryParameters** | [**LeaveApplicationDtoCollectionQueryParameters**](LeaveApplicationDtoCollectionQueryParameters.md) |  | 

### Return type

[**LeaveApplicationDtoListEnvelope**](LeaveApplicationDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeaveApplicationsCountAsync

> Int32Envelope GetLeaveApplicationsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveApplicationDtoCollectionQueryParameters(leaveApplicationDtoCollectionQueryParameters).Execute()

Count leave applications



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
	leaveApplicationDtoCollectionQueryParameters := *openapiclient.NewLeaveApplicationDtoCollectionQueryParameters() // LeaveApplicationDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveApplicationsAPI.GetLeaveApplicationsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveApplicationDtoCollectionQueryParameters(leaveApplicationDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveApplicationsAPI.GetLeaveApplicationsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLeaveApplicationsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveApplicationsAPI.GetLeaveApplicationsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaveApplicationsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **leaveApplicationDtoCollectionQueryParameters** | [**LeaveApplicationDtoCollectionQueryParameters**](LeaveApplicationDtoCollectionQueryParameters.md) |  | 

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


## PatchLeaveApplicationAsync

> EmptyEnvelope PatchLeaveApplicationAsync(ctx, leaveApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a leave application



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
	leaveApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveApplicationsAPI.PatchLeaveApplicationAsync(context.Background(), leaveApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveApplicationsAPI.PatchLeaveApplicationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchLeaveApplicationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveApplicationsAPI.PatchLeaveApplicationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leaveApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLeaveApplicationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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


## UpdateLeaveApplicationAsync

> EmptyEnvelope UpdateLeaveApplicationAsync(ctx, leaveApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveApplicationUpdateDto(leaveApplicationUpdateDto).Execute()

Update a leave application



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
	leaveApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	leaveApplicationUpdateDto := *openapiclient.NewLeaveApplicationUpdateDto() // LeaveApplicationUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LeaveApplicationsAPI.UpdateLeaveApplicationAsync(context.Background(), leaveApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LeaveApplicationUpdateDto(leaveApplicationUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LeaveApplicationsAPI.UpdateLeaveApplicationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLeaveApplicationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LeaveApplicationsAPI.UpdateLeaveApplicationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**leaveApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLeaveApplicationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **leaveApplicationUpdateDto** | [**LeaveApplicationUpdateDto**](LeaveApplicationUpdateDto.md) |  | 

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

