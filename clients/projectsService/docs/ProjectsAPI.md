# \ProjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectAsync**](ProjectsAPI.md#CreateProjectAsync) | **Post** /api/v2/ProjectsService/Projects | Creates a new project
[**CreateProjectPeriodAsync**](ProjectsAPI.md#CreateProjectPeriodAsync) | **Post** /api/v2/ProjectsService/Projects/{projectId}/Periods | Creates a project period
[**CreateProjectTaskAsync**](ProjectsAPI.md#CreateProjectTaskAsync) | **Post** /api/v2/ProjectsService/Projects/{projectId}/Tasks | Creates a project task
[**DeleteProjectAsync**](ProjectsAPI.md#DeleteProjectAsync) | **Delete** /api/v2/ProjectsService/Projects/{projectId} | Deletes a project
[**DeleteProjectPeriodAsync**](ProjectsAPI.md#DeleteProjectPeriodAsync) | **Delete** /api/v2/ProjectsService/Projects/{projectId}/Periods/{projectPeriodId} | Deletes a project period
[**DeleteProjectTaskAsync**](ProjectsAPI.md#DeleteProjectTaskAsync) | **Delete** /api/v2/ProjectsService/Projects/{projectId}/Tasks/{projectTaskId} | Deletes a project task
[**GetProjectByIdAsync**](ProjectsAPI.md#GetProjectByIdAsync) | **Get** /api/v2/ProjectsService/Projects/{projectId} | Gets a project by ID
[**GetProjectPeriodsAsync**](ProjectsAPI.md#GetProjectPeriodsAsync) | **Get** /api/v2/ProjectsService/Projects/{projectId}/Periods | Retrieves project periods
[**GetProjectTaskCategoriesAsync**](ProjectsAPI.md#GetProjectTaskCategoriesAsync) | **Get** /api/v2/ProjectsService/Projects/{projectId}/TaskCategories | Retrieves project task categories
[**GetProjectTaskCategoriesCountAsync**](ProjectsAPI.md#GetProjectTaskCategoriesCountAsync) | **Get** /api/v2/ProjectsService/Projects/{projectId}/TaskCategories/Count | Counts project task categories
[**GetProjectTasksAsync**](ProjectsAPI.md#GetProjectTasksAsync) | **Get** /api/v2/ProjectsService/Projects/{projectId}/Tasks | Retrieves project tasks
[**GetProjectTasksCountAsync**](ProjectsAPI.md#GetProjectTasksCountAsync) | **Get** /api/v2/ProjectsService/Projects/{projectId}/Tasks/Count | Counts project tasks
[**GetProjectTimeLogsAsync**](ProjectsAPI.md#GetProjectTimeLogsAsync) | **Get** /api/v2/ProjectsService/Projects/{projectId}/TimeLogs | Retrieves project time logs
[**GetProjectTimeLogsCountAsync**](ProjectsAPI.md#GetProjectTimeLogsCountAsync) | **Get** /api/v2/ProjectsService/Projects/{projectId}/TimeLogs/Count | Counts project time logs
[**GetProjectsByTenantIdAsync**](ProjectsAPI.md#GetProjectsByTenantIdAsync) | **Get** /api/v2/ProjectsService/Projects | Retrieves all projects
[**GetProjectsCountByTenantIdAsync**](ProjectsAPI.md#GetProjectsCountByTenantIdAsync) | **Get** /api/v2/ProjectsService/Projects/Count | Counts projects
[**UpdateProjectAsync**](ProjectsAPI.md#UpdateProjectAsync) | **Put** /api/v2/ProjectsService/Projects/{projectId} | Updates a project
[**UpdateProjectPeriodAsync**](ProjectsAPI.md#UpdateProjectPeriodAsync) | **Put** /api/v2/ProjectsService/Projects/{projectId}/Periods/{projectPeriodId} | Updates a project period
[**UpdateProjectTaskAsync**](ProjectsAPI.md#UpdateProjectTaskAsync) | **Put** /api/v2/ProjectsService/Projects/{projectId}/Tasks/{projectTaskId} | Updates a project task



## CreateProjectAsync

> EmptyEnvelope CreateProjectAsync(ctx).TenantId(tenantId).ProjectCreateDto(projectCreateDto).Execute()

Creates a new project



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
	projectCreateDto := *openapiclient.NewProjectCreateDto() // ProjectCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.CreateProjectAsync(context.Background()).TenantId(tenantId).ProjectCreateDto(projectCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CreateProjectAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.CreateProjectAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **projectCreateDto** | [**ProjectCreateDto**](ProjectCreateDto.md) |  | 

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


## CreateProjectPeriodAsync

> EmptyEnvelope CreateProjectPeriodAsync(ctx, projectId).TenantId(tenantId).ProjectPeriodCreateDto(projectPeriodCreateDto).Execute()

Creates a project period



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectPeriodCreateDto := *openapiclient.NewProjectPeriodCreateDto() // ProjectPeriodCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.CreateProjectPeriodAsync(context.Background(), projectId).TenantId(tenantId).ProjectPeriodCreateDto(projectPeriodCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CreateProjectPeriodAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectPeriodAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.CreateProjectPeriodAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectPeriodAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **projectPeriodCreateDto** | [**ProjectPeriodCreateDto**](ProjectPeriodCreateDto.md) |  | 

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


## CreateProjectTaskAsync

> EmptyEnvelope CreateProjectTaskAsync(ctx, projectId).TenantId(tenantId).ProjectTaskCreateDto(projectTaskCreateDto).Execute()

Creates a project task



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectTaskCreateDto := *openapiclient.NewProjectTaskCreateDto() // ProjectTaskCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.CreateProjectTaskAsync(context.Background(), projectId).TenantId(tenantId).ProjectTaskCreateDto(projectTaskCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.CreateProjectTaskAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectTaskAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.CreateProjectTaskAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectTaskAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **projectTaskCreateDto** | [**ProjectTaskCreateDto**](ProjectTaskCreateDto.md) |  | 

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


## DeleteProjectAsync

> EmptyEnvelope DeleteProjectAsync(ctx, projectId).TenantId(tenantId).Execute()

Deletes a project



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.DeleteProjectAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProjectAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.DeleteProjectAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

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


## DeleteProjectPeriodAsync

> EmptyEnvelope DeleteProjectPeriodAsync(ctx, projectId, projectPeriodId).TenantId(tenantId).Execute()

Deletes a project period



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectPeriodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.DeleteProjectPeriodAsync(context.Background(), projectId, projectPeriodId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProjectPeriodAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectPeriodAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.DeleteProjectPeriodAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectPeriodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectPeriodAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **string** |  | 

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


## DeleteProjectTaskAsync

> EmptyEnvelope DeleteProjectTaskAsync(ctx, projectId, projectTaskId).TenantId(tenantId).Execute()

Deletes a project task



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectTaskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.DeleteProjectTaskAsync(context.Background(), projectId, projectTaskId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProjectTaskAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectTaskAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.DeleteProjectTaskAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectTaskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectTaskAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## GetProjectByIdAsync

> ProjectDtoEnvelope GetProjectByIdAsync(ctx, projectId).TenantId(tenantId).Execute()

Gets a project by ID



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectByIdAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectByIdAsync`: ProjectDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**ProjectDtoEnvelope**](ProjectDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectPeriodsAsync

> ProjectPeriodDtoListEnvelope GetProjectPeriodsAsync(ctx, projectId).TenantId(tenantId).Execute()

Retrieves project periods



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectPeriodsAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectPeriodsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectPeriodsAsync`: ProjectPeriodDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectPeriodsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectPeriodsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**ProjectPeriodDtoListEnvelope**](ProjectPeriodDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTaskCategoriesAsync

> TaskCategoryDtoListEnvelope GetProjectTaskCategoriesAsync(ctx, projectId).TenantId(tenantId).Execute()

Retrieves project task categories



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectTaskCategoriesAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectTaskCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTaskCategoriesAsync`: TaskCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectTaskCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTaskCategoriesAsyncRequest struct via the builder pattern


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


## GetProjectTaskCategoriesCountAsync

> Int32Envelope GetProjectTaskCategoriesCountAsync(ctx, projectId).TenantId(tenantId).Execute()

Counts project task categories



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectTaskCategoriesCountAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectTaskCategoriesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTaskCategoriesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectTaskCategoriesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTaskCategoriesCountAsyncRequest struct via the builder pattern


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


## GetProjectTasksAsync

> ProjectTaskDtoListEnvelope GetProjectTasksAsync(ctx, projectId).TenantId(tenantId).Execute()

Retrieves project tasks



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectTasksAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectTasksAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTasksAsync`: ProjectTaskDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectTasksAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTasksAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**ProjectTaskDtoListEnvelope**](ProjectTaskDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTasksCountAsync

> Int32Envelope GetProjectTasksCountAsync(ctx, projectId).TenantId(tenantId).Execute()

Counts project tasks



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectTasksCountAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectTasksCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTasksCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectTasksCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTasksCountAsyncRequest struct via the builder pattern


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


## GetProjectTimeLogsAsync

> ProjectTimeLogDtoListEnvelope GetProjectTimeLogsAsync(ctx, projectId).TenantId(tenantId).Execute()

Retrieves project time logs



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectTimeLogsAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectTimeLogsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTimeLogsAsync`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectTimeLogsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTimeLogsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

### Return type

[**ProjectTimeLogDtoListEnvelope**](ProjectTimeLogDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTimeLogsCountAsync

> Int32Envelope GetProjectTimeLogsCountAsync(ctx, projectId).TenantId(tenantId).Execute()

Counts project time logs



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectTimeLogsCountAsync(context.Background(), projectId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectTimeLogsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTimeLogsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectTimeLogsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTimeLogsCountAsyncRequest struct via the builder pattern


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


## GetProjectsByTenantIdAsync

> ProjectDtoListEnvelope GetProjectsByTenantIdAsync(ctx).TenantId(tenantId).Execute()

Retrieves all projects



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
	resp, r, err := apiClient.ProjectsAPI.GetProjectsByTenantIdAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsByTenantIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsByTenantIdAsync`: ProjectDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsByTenantIdAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsByTenantIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ProjectDtoListEnvelope**](ProjectDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsCountByTenantIdAsync

> Int32Envelope GetProjectsCountByTenantIdAsync(ctx).TenantId(tenantId).Execute()

Counts projects



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
	resp, r, err := apiClient.ProjectsAPI.GetProjectsCountByTenantIdAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsCountByTenantIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsCountByTenantIdAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsCountByTenantIdAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsCountByTenantIdAsyncRequest struct via the builder pattern


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


## UpdateProjectAsync

> EmptyEnvelope UpdateProjectAsync(ctx, projectId).TenantId(tenantId).ProjectUpdateDto(projectUpdateDto).Execute()

Updates a project



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectUpdateDto := *openapiclient.NewProjectUpdateDto() // ProjectUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.UpdateProjectAsync(context.Background(), projectId).TenantId(tenantId).ProjectUpdateDto(projectUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProjectAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateProjectAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **projectUpdateDto** | [**ProjectUpdateDto**](ProjectUpdateDto.md) |  | 

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


## UpdateProjectPeriodAsync

> EmptyEnvelope UpdateProjectPeriodAsync(ctx, projectId, projectPeriodId).TenantId(tenantId).ProjectPeriodUpdateDto(projectPeriodUpdateDto).Execute()

Updates a project period



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectPeriodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectPeriodUpdateDto := *openapiclient.NewProjectPeriodUpdateDto() // ProjectPeriodUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.UpdateProjectPeriodAsync(context.Background(), projectId, projectPeriodId).TenantId(tenantId).ProjectPeriodUpdateDto(projectPeriodUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProjectPeriodAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectPeriodAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateProjectPeriodAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectPeriodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectPeriodAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **string** |  | 
 **projectPeriodUpdateDto** | [**ProjectPeriodUpdateDto**](ProjectPeriodUpdateDto.md) |  | 

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


## UpdateProjectTaskAsync

> EmptyEnvelope UpdateProjectTaskAsync(ctx, projectId, projectTaskId).TenantId(tenantId).ProjectTaskUpdateDto(projectTaskUpdateDto).Execute()

Updates a project task



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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectTaskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	projectTaskUpdateDto := *openapiclient.NewProjectTaskUpdateDto() // ProjectTaskUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.UpdateProjectTaskAsync(context.Background(), projectId, projectTaskId).TenantId(tenantId).ProjectTaskUpdateDto(projectTaskUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProjectTaskAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectTaskAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateProjectTaskAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectTaskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectTaskAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **string** |  | 
 **projectTaskUpdateDto** | [**ProjectTaskUpdateDto**](ProjectTaskUpdateDto.md) |  | 

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

