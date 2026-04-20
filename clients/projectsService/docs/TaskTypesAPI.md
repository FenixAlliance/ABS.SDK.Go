# \TaskTypesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskTypeAsync**](TaskTypesAPI.md#CreateTaskTypeAsync) | **Post** /api/v2/ProjectsService/TaskTypes | Creates a new task type
[**DeleteTaskTypeAsync**](TaskTypesAPI.md#DeleteTaskTypeAsync) | **Delete** /api/v2/ProjectsService/TaskTypes/{taskTypeId} | Deletes a task type
[**GetTaskTypeByIdAsync**](TaskTypesAPI.md#GetTaskTypeByIdAsync) | **Get** /api/v2/ProjectsService/TaskTypes/{taskTypeId} | Gets a task type by ID
[**UpdateTaskTypeAsync**](TaskTypesAPI.md#UpdateTaskTypeAsync) | **Put** /api/v2/ProjectsService/TaskTypes/{taskTypeId} | Updates a task type



## CreateTaskTypeAsync

> TaskTypeDto CreateTaskTypeAsync(ctx).TenantId(tenantId).TaskTypeCreateDto(taskTypeCreateDto).Execute()

Creates a new task type



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
	taskTypeCreateDto := *openapiclient.NewTaskTypeCreateDto() // TaskTypeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskTypesAPI.CreateTaskTypeAsync(context.Background()).TenantId(tenantId).TaskTypeCreateDto(taskTypeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskTypesAPI.CreateTaskTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskTypeAsync`: TaskTypeDto
	fmt.Fprintf(os.Stdout, "Response from `TaskTypesAPI.CreateTaskTypeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **taskTypeCreateDto** | [**TaskTypeCreateDto**](TaskTypeCreateDto.md) |  | 

### Return type

[**TaskTypeDto**](TaskTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskTypeAsync

> TaskTypeDto DeleteTaskTypeAsync(ctx, taskTypeId).TenantId(tenantId).Execute()

Deletes a task type



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
	taskTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskTypesAPI.DeleteTaskTypeAsync(context.Background(), taskTypeId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskTypesAPI.DeleteTaskTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTaskTypeAsync`: TaskTypeDto
	fmt.Fprintf(os.Stdout, "Response from `TaskTypesAPI.DeleteTaskTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**TaskTypeDto**](TaskTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskTypeByIdAsync

> TaskTypeDto GetTaskTypeByIdAsync(ctx, taskTypeId).TenantId(tenantId).Execute()

Gets a task type by ID



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
	taskTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskTypesAPI.GetTaskTypeByIdAsync(context.Background(), taskTypeId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskTypesAPI.GetTaskTypeByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskTypeByIdAsync`: TaskTypeDto
	fmt.Fprintf(os.Stdout, "Response from `TaskTypesAPI.GetTaskTypeByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskTypeByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**TaskTypeDto**](TaskTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskTypeAsync

> TaskTypeDto UpdateTaskTypeAsync(ctx, taskTypeId).TenantId(tenantId).TaskTypeUpdateDto(taskTypeUpdateDto).Execute()

Updates a task type



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
	taskTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	taskTypeUpdateDto := *openapiclient.NewTaskTypeUpdateDto() // TaskTypeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskTypesAPI.UpdateTaskTypeAsync(context.Background(), taskTypeId).TenantId(tenantId).TaskTypeUpdateDto(taskTypeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskTypesAPI.UpdateTaskTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTaskTypeAsync`: TaskTypeDto
	fmt.Fprintf(os.Stdout, "Response from `TaskTypesAPI.UpdateTaskTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **taskTypeUpdateDto** | [**TaskTypeUpdateDto**](TaskTypeUpdateDto.md) |  | 

### Return type

[**TaskTypeDto**](TaskTypeDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

