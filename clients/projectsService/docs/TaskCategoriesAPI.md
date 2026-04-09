# \TaskCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountTenantTaskCategoriesAsync**](TaskCategoriesAPI.md#CountTenantTaskCategoriesAsync) | **Get** /api/v2/ProjectsService/TaskCategories/Count | Counts task categories
[**CreateTaskCategoryAsync**](TaskCategoriesAPI.md#CreateTaskCategoryAsync) | **Post** /api/v2/ProjectsService/TaskCategories | Creates a new task category
[**DeleteTaskCategoryAsync**](TaskCategoriesAPI.md#DeleteTaskCategoryAsync) | **Delete** /api/v2/ProjectsService/TaskCategories/{taskCategoryId} | Deletes a task category
[**GetTaskCategoryByIdAsync**](TaskCategoriesAPI.md#GetTaskCategoryByIdAsync) | **Get** /api/v2/ProjectsService/TaskCategories/{taskCategoryId} | Gets a task category by ID
[**GetTaskCategoryTaskTypesAsync**](TaskCategoriesAPI.md#GetTaskCategoryTaskTypesAsync) | **Get** /api/v2/ProjectsService/TaskCategories/{taskCategoryId}/Types | Retrieves task types for a category
[**GetTenantTaskCategoriesAsync**](TaskCategoriesAPI.md#GetTenantTaskCategoriesAsync) | **Get** /api/v2/ProjectsService/TaskCategories | Retrieves all task categories
[**UpdateTaskCategoryAsync**](TaskCategoriesAPI.md#UpdateTaskCategoryAsync) | **Put** /api/v2/ProjectsService/TaskCategories/{taskCategoryId} | Updates a task category



## CountTenantTaskCategoriesAsync

> Int32Envelope CountTenantTaskCategoriesAsync(ctx).TenantId(tenantId).Execute()

Counts task categories



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.CountTenantTaskCategoriesAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.CountTenantTaskCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountTenantTaskCategoriesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.CountTenantTaskCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountTenantTaskCategoriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## CreateTaskCategoryAsync

> TaskCategoryDto CreateTaskCategoryAsync(ctx).TenantId(tenantId).TaskCategoryCreateDto(taskCategoryCreateDto).Execute()

Creates a new task category



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
	taskCategoryCreateDto := *openapiclient.NewTaskCategoryCreateDto() // TaskCategoryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.CreateTaskCategoryAsync(context.Background()).TenantId(tenantId).TaskCategoryCreateDto(taskCategoryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.CreateTaskCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaskCategoryAsync`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.CreateTaskCategoryAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskCategoryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **taskCategoryCreateDto** | [**TaskCategoryCreateDto**](TaskCategoryCreateDto.md) |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskCategoryAsync

> TaskCategoryDto DeleteTaskCategoryAsync(ctx, taskCategoryId).TenantId(tenantId).Execute()

Deletes a task category



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
	taskCategoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.DeleteTaskCategoryAsync(context.Background(), taskCategoryId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.DeleteTaskCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTaskCategoryAsync`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.DeleteTaskCategoryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskCategoryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskCategoryByIdAsync

> TaskCategoryDto GetTaskCategoryByIdAsync(ctx, taskCategoryId).TenantId(tenantId).Execute()

Gets a task category by ID



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
	taskCategoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.GetTaskCategoryByIdAsync(context.Background(), taskCategoryId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.GetTaskCategoryByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskCategoryByIdAsync`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.GetTaskCategoryByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskCategoryByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskCategoryTaskTypesAsync

> TaskCategoryDto GetTaskCategoryTaskTypesAsync(ctx, taskCategoryId).TenantId(tenantId).Execute()

Retrieves task types for a category



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
	taskCategoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.GetTaskCategoryTaskTypesAsync(context.Background(), taskCategoryId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.GetTaskCategoryTaskTypesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskCategoryTaskTypesAsync`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.GetTaskCategoryTaskTypesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskCategoryTaskTypesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantTaskCategoriesAsync

> TaskCategoryDtoListEnvelope GetTenantTaskCategoriesAsync(ctx).TenantId(tenantId).Execute()

Retrieves all task categories



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.GetTenantTaskCategoriesAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.GetTenantTaskCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTaskCategoriesAsync`: TaskCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.GetTenantTaskCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTaskCategoriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**TaskCategoryDtoListEnvelope**](TaskCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskCategoryAsync

> TaskCategoryDto UpdateTaskCategoryAsync(ctx, taskCategoryId).TenantId(tenantId).TaskCategoryUpdateDto(taskCategoryUpdateDto).Execute()

Updates a task category



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
	taskCategoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	taskCategoryUpdateDto := *openapiclient.NewTaskCategoryUpdateDto() // TaskCategoryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.UpdateTaskCategoryAsync(context.Background(), taskCategoryId).TenantId(tenantId).TaskCategoryUpdateDto(taskCategoryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.UpdateTaskCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTaskCategoryAsync`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.UpdateTaskCategoryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskCategoryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **taskCategoryUpdateDto** | [**TaskCategoryUpdateDto**](TaskCategoryUpdateDto.md) |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

