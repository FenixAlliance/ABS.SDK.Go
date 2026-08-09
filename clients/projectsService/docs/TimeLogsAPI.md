# \TimeLogsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountProjectPeriodTimeLogsAsync**](TimeLogsAPI.md#CountProjectPeriodTimeLogsAsync) | **Get** /api/v2/ProjectsService/TimeLogs/Count | Get the count of project period time logs
[**CreateProjectTimeLogAsync**](TimeLogsAPI.md#CreateProjectTimeLogAsync) | **Post** /api/v2/ProjectsService/TimeLogs | Create a new project time log
[**DeleteProjectTimeLogAsync**](TimeLogsAPI.md#DeleteProjectTimeLogAsync) | **Delete** /api/v2/ProjectsService/TimeLogs/{timeLogId} | Delete a project time log
[**GetProjectPeriodTimeLogsAsync**](TimeLogsAPI.md#GetProjectPeriodTimeLogsAsync) | **Get** /api/v2/ProjectsService/TimeLogs | Retrieve project period time logs
[**GetProjectTimeLogByIdAsync**](TimeLogsAPI.md#GetProjectTimeLogByIdAsync) | **Get** /api/v2/ProjectsService/TimeLogs/{timeLogId} | Retrieve a project time log by ID
[**GetProjectTimeLogsAsync**](TimeLogsAPI.md#GetProjectTimeLogsAsync) | **Get** /api/v2/ProjectsService/TimeLogs/ForProject/{projectId} | Retrieve time logs for a project
[**GetProjectTimeLogsByResponsibleContactAsync**](TimeLogsAPI.md#GetProjectTimeLogsByResponsibleContactAsync) | **Get** /api/v2/ProjectsService/TimeLogs/ByResponsibleContact | Retrieve time logs by responsible contact
[**GetProjectTimeLogsCreatedByContactAsync**](TimeLogsAPI.md#GetProjectTimeLogsCreatedByContactAsync) | **Get** /api/v2/ProjectsService/TimeLogs/CreatedByContact | Retrieve time logs created by a contact
[**PatchProjectTimeLogAsync**](TimeLogsAPI.md#PatchProjectTimeLogAsync) | **Patch** /api/v2/ProjectsService/TimeLogs/{timeLogId} | Patch a project time log
[**UpdateProjectTimeLogAsync**](TimeLogsAPI.md#UpdateProjectTimeLogAsync) | **Put** /api/v2/ProjectsService/TimeLogs/{timeLogId} | Update a project time log



## CountProjectPeriodTimeLogsAsync

> Int32Envelope CountProjectPeriodTimeLogsAsync(ctx).TenantId(tenantId).ProjectPeriodId(projectPeriodId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogDtoCollectionQueryParameters(projectTimeLogDtoCollectionQueryParameters).Execute()

Get the count of project period time logs



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
	projectPeriodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectTimeLogDtoCollectionQueryParameters := *openapiclient.NewProjectTimeLogDtoCollectionQueryParameters() // ProjectTimeLogDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeLogsAPI.CountProjectPeriodTimeLogsAsync(context.Background()).TenantId(tenantId).ProjectPeriodId(projectPeriodId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogDtoCollectionQueryParameters(projectTimeLogDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.CountProjectPeriodTimeLogsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountProjectPeriodTimeLogsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TimeLogsAPI.CountProjectPeriodTimeLogsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountProjectPeriodTimeLogsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **projectPeriodId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectTimeLogDtoCollectionQueryParameters** | [**ProjectTimeLogDtoCollectionQueryParameters**](ProjectTimeLogDtoCollectionQueryParameters.md) |  | 

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


## CreateProjectTimeLogAsync

> CreateProjectTimeLogAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogCreateDto(projectTimeLogCreateDto).Execute()

Create a new project time log



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
	projectTimeLogCreateDto := *openapiclient.NewProjectTimeLogCreateDto("ProjectTaskId_example", "ProjectPeriodId_example") // ProjectTimeLogCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeLogsAPI.CreateProjectTimeLogAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogCreateDto(projectTimeLogCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.CreateProjectTimeLogAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectTimeLogAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectTimeLogCreateDto** | [**ProjectTimeLogCreateDto**](ProjectTimeLogCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectTimeLogAsync

> DeleteProjectTimeLogAsync(ctx, timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a project time log



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
	timeLogId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeLogsAPI.DeleteProjectTimeLogAsync(context.Background(), timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.DeleteProjectTimeLogAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeLogId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectTimeLogAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectPeriodTimeLogsAsync

> ProjectTimeLogDtoListEnvelope GetProjectPeriodTimeLogsAsync(ctx).TenantId(tenantId).ProjectPeriodId(projectPeriodId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogDtoCollectionQueryParameters(projectTimeLogDtoCollectionQueryParameters).Execute()

Retrieve project period time logs



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
	projectPeriodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectTimeLogDtoCollectionQueryParameters := *openapiclient.NewProjectTimeLogDtoCollectionQueryParameters() // ProjectTimeLogDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeLogsAPI.GetProjectPeriodTimeLogsAsync(context.Background()).TenantId(tenantId).ProjectPeriodId(projectPeriodId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogDtoCollectionQueryParameters(projectTimeLogDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.GetProjectPeriodTimeLogsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectPeriodTimeLogsAsync`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TimeLogsAPI.GetProjectPeriodTimeLogsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectPeriodTimeLogsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **projectPeriodId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectTimeLogDtoCollectionQueryParameters** | [**ProjectTimeLogDtoCollectionQueryParameters**](ProjectTimeLogDtoCollectionQueryParameters.md) |  | 

### Return type

[**ProjectTimeLogDtoListEnvelope**](ProjectTimeLogDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTimeLogByIdAsync

> ProjectTimeLogDtoEnvelope GetProjectTimeLogByIdAsync(ctx, timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a project time log by ID



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
	timeLogId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeLogsAPI.GetProjectTimeLogByIdAsync(context.Background(), timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.GetProjectTimeLogByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTimeLogByIdAsync`: ProjectTimeLogDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TimeLogsAPI.GetProjectTimeLogByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeLogId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTimeLogByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProjectTimeLogDtoEnvelope**](ProjectTimeLogDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTimeLogsAsync

> ProjectTimeLogDtoListEnvelope GetProjectTimeLogsAsync(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve time logs for a project



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
	resp, r, err := apiClient.TimeLogsAPI.GetProjectTimeLogsAsync(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.GetProjectTimeLogsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTimeLogsAsync`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TimeLogsAPI.GetProjectTimeLogsAsync`: %v\n", resp)
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
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## GetProjectTimeLogsByResponsibleContactAsync

> ProjectTimeLogDtoListEnvelope GetProjectTimeLogsByResponsibleContactAsync(ctx).ContactId(contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve time logs by responsible contact



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeLogsAPI.GetProjectTimeLogsByResponsibleContactAsync(context.Background()).ContactId(contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.GetProjectTimeLogsByResponsibleContactAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTimeLogsByResponsibleContactAsync`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TimeLogsAPI.GetProjectTimeLogsByResponsibleContactAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTimeLogsByResponsibleContactAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string** |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## GetProjectTimeLogsCreatedByContactAsync

> ProjectTimeLogDtoListEnvelope GetProjectTimeLogsCreatedByContactAsync(ctx).ContactId(contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve time logs created by a contact



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeLogsAPI.GetProjectTimeLogsCreatedByContactAsync(context.Background()).ContactId(contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.GetProjectTimeLogsCreatedByContactAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTimeLogsCreatedByContactAsync`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TimeLogsAPI.GetProjectTimeLogsCreatedByContactAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTimeLogsCreatedByContactAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string** |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## PatchProjectTimeLogAsync

> PatchProjectTimeLogAsync(ctx, timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a project time log



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
	timeLogId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeLogsAPI.PatchProjectTimeLogAsync(context.Background(), timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.PatchProjectTimeLogAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeLogId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectTimeLogAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectTimeLogAsync

> UpdateProjectTimeLogAsync(ctx, timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogUpdateDto(projectTimeLogUpdateDto).Execute()

Update a project time log



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
	timeLogId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectTimeLogUpdateDto := *openapiclient.NewProjectTimeLogUpdateDto() // ProjectTimeLogUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeLogsAPI.UpdateProjectTimeLogAsync(context.Background(), timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogUpdateDto(projectTimeLogUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogsAPI.UpdateProjectTimeLogAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeLogId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectTimeLogAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectTimeLogUpdateDto** | [**ProjectTimeLogUpdateDto**](ProjectTimeLogUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

