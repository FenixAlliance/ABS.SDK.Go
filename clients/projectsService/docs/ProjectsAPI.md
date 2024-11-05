# \ProjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsServiceProjectsCountGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsCountGet) | **Get** /api/v2/ProjectsService/Projects/Count | 
[**ApiV2ProjectsServiceProjectsGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsGet) | **Get** /api/v2/ProjectsService/Projects | 
[**ApiV2ProjectsServiceProjectsPost**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsPost) | **Post** /api/v2/ProjectsService/Projects | 
[**ApiV2ProjectsServiceProjectsProjectIdDelete**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdDelete) | **Delete** /api/v2/ProjectsService/Projects/{projectId} | 
[**ApiV2ProjectsServiceProjectsProjectIdGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdGet) | **Get** /api/v2/ProjectsService/Projects/{projectId} | 
[**ApiV2ProjectsServiceProjectsProjectIdPeriodsGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdPeriodsGet) | **Get** /api/v2/ProjectsService/Projects/{projectId}/Periods | 
[**ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDelete**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDelete) | **Delete** /api/v2/ProjectsService/Projects/{projectId}/Periods/{projectPeriodId} | 
[**ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPut**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPut) | **Put** /api/v2/ProjectsService/Projects/{projectId}/Periods/{projectPeriodId} | 
[**ApiV2ProjectsServiceProjectsProjectIdPut**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdPut) | **Put** /api/v2/ProjectsService/Projects/{projectId} | 
[**ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGet) | **Get** /api/v2/ProjectsService/Projects/{projectId}/TaskCategories/Count | 
[**ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGet) | **Get** /api/v2/ProjectsService/Projects/{projectId}/TaskCategories | 
[**ApiV2ProjectsServiceProjectsProjectIdTasksCountGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTasksCountGet) | **Get** /api/v2/ProjectsService/Projects/{projectId}/Tasks/Count | 
[**ApiV2ProjectsServiceProjectsProjectIdTasksGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTasksGet) | **Get** /api/v2/ProjectsService/Projects/{projectId}/Tasks | 
[**ApiV2ProjectsServiceProjectsProjectIdTasksPost**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTasksPost) | **Post** /api/v2/ProjectsService/Projects/{projectId}/Tasks | 
[**ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDelete**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDelete) | **Delete** /api/v2/ProjectsService/Projects/{projectId}/Tasks/{projectTaskId} | 
[**ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPut**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPut) | **Put** /api/v2/ProjectsService/Projects/{projectId}/Tasks/{projectTaskId} | 
[**ApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGet) | **Get** /api/v2/ProjectsService/Projects/{projectId}/TimeLogs/Count | 
[**ApiV2ProjectsServiceProjectsProjectIdTimeLogsGet**](ProjectsAPI.md#ApiV2ProjectsServiceProjectsProjectIdTimeLogsGet) | **Get** /api/v2/ProjectsService/Projects/{projectId}/TimeLogs | 



## ApiV2ProjectsServiceProjectsCountGet

> Int32Envelope ApiV2ProjectsServiceProjectsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsCountGetRequest struct via the builder pattern


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


## ApiV2ProjectsServiceProjectsGet

> ProjectDtoListEnvelope ApiV2ProjectsServiceProjectsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsGet`: ProjectDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProjectDtoListEnvelope**](ProjectDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceProjectsPost

> EmptyEnvelope ApiV2ProjectsServiceProjectsPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectCreateDto(projectCreateDto).Execute()



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
	projectCreateDto := *openapiclient.NewProjectCreateDto() // ProjectCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectCreateDto(projectCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectCreateDto** | [**ProjectCreateDto**](ProjectCreateDto.md) |  | 

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


## ApiV2ProjectsServiceProjectsProjectIdDelete

> EmptyEnvelope ApiV2ProjectsServiceProjectsProjectIdDelete(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdDelete(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdDeleteRequest struct via the builder pattern


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


## ApiV2ProjectsServiceProjectsProjectIdGet

> ProjectDtoEnvelope ApiV2ProjectsServiceProjectsProjectIdGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdGet`: ProjectDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProjectDtoEnvelope**](ProjectDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceProjectsProjectIdPeriodsGet

> EmptyEnvelope ApiV2ProjectsServiceProjectsProjectIdPeriodsGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectPeriodCreateDto(projectPeriodCreateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectPeriodCreateDto := *openapiclient.NewProjectPeriodCreateDto() // ProjectPeriodCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectPeriodCreateDto(projectPeriodCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdPeriodsGet`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdPeriodsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectPeriodCreateDto** | [**ProjectPeriodCreateDto**](ProjectPeriodCreateDto.md) |  | 

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


## ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDelete

> EmptyEnvelope ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDelete(ctx, projectId, projectPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDelete(context.Background(), projectId, projectPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectPeriodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdDeleteRequest struct via the builder pattern


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


## ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPut

> EmptyEnvelope ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPut(ctx, projectId, projectPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectPeriodUpdateDto(projectPeriodUpdateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectPeriodUpdateDto := *openapiclient.NewProjectPeriodUpdateDto() // ProjectPeriodUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPut(context.Background(), projectId, projectPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectPeriodUpdateDto(projectPeriodUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectPeriodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdPeriodsProjectPeriodIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectPeriodUpdateDto** | [**ProjectPeriodUpdateDto**](ProjectPeriodUpdateDto.md) |  | 

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


## ApiV2ProjectsServiceProjectsProjectIdPut

> EmptyEnvelope ApiV2ProjectsServiceProjectsProjectIdPut(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectUpdateDto(projectUpdateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectUpdateDto := *openapiclient.NewProjectUpdateDto() // ProjectUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPut(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectUpdateDto(projectUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectUpdateDto** | [**ProjectUpdateDto**](ProjectUpdateDto.md) |  | 

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


## ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGet

> Int32Envelope ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTaskCategoriesCountGetRequest struct via the builder pattern


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


## ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGet

> TaskCategoryDtoListEnvelope ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGet`: TaskCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTaskCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaskCategoryDtoListEnvelope**](TaskCategoryDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceProjectsProjectIdTasksCountGet

> Int32Envelope ApiV2ProjectsServiceProjectsProjectIdTasksCountGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksCountGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTasksCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTasksCountGetRequest struct via the builder pattern


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


## ApiV2ProjectsServiceProjectsProjectIdTasksGet

> ProjectTaskDtoListEnvelope ApiV2ProjectsServiceProjectsProjectIdTasksGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTasksGet`: ProjectTaskDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProjectTaskDtoListEnvelope**](ProjectTaskDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceProjectsProjectIdTasksPost

> EmptyEnvelope ApiV2ProjectsServiceProjectsProjectIdTasksPost(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTaskCreateDto(projectTaskCreateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectTaskCreateDto := *openapiclient.NewProjectTaskCreateDto() // ProjectTaskCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksPost(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTaskCreateDto(projectTaskCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTasksPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectTaskCreateDto** | [**ProjectTaskCreateDto**](ProjectTaskCreateDto.md) |  | 

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


## ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDelete

> EmptyEnvelope ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDelete(ctx, projectId, projectTaskId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDelete(context.Background(), projectId, projectTaskId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectTaskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdDeleteRequest struct via the builder pattern


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


## ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPut

> EmptyEnvelope ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPut(ctx, projectId, projectTaskId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTaskUpdateDto(projectTaskUpdateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectTaskUpdateDto := *openapiclient.NewProjectTaskUpdateDto() // ProjectTaskUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPut(context.Background(), projectId, projectTaskId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTaskUpdateDto(projectTaskUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**projectTaskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTasksProjectTaskIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectTaskUpdateDto** | [**ProjectTaskUpdateDto**](ProjectTaskUpdateDto.md) |  | 

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


## ApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGet

> Int32Envelope ApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTimeLogsCountGetRequest struct via the builder pattern


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


## ApiV2ProjectsServiceProjectsProjectIdTimeLogsGet

> ProjectTimeLogDtoListEnvelope ApiV2ProjectsServiceProjectsProjectIdTimeLogsGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTimeLogsGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTimeLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceProjectsProjectIdTimeLogsGet`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ApiV2ProjectsServiceProjectsProjectIdTimeLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceProjectsProjectIdTimeLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProjectTimeLogDtoListEnvelope**](ProjectTimeLogDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

