# \ProjectTimeLogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGet**](ProjectTimeLogsAPI.md#ApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGet) | **Get** /api/v2/TimeTrackerService/ProjectTimeLogs/ByResponsibleContact | 
[**ApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGet**](ProjectTimeLogsAPI.md#ApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGet) | **Get** /api/v2/TimeTrackerService/ProjectTimeLogs/CreatedByContact | 
[**ApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGet**](ProjectTimeLogsAPI.md#ApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGet) | **Get** /api/v2/TimeTrackerService/ProjectTimeLogs/ForProject/{projectId} | 
[**ApiV2TimeTrackerServiceProjectTimeLogsGet**](ProjectTimeLogsAPI.md#ApiV2TimeTrackerServiceProjectTimeLogsGet) | **Get** /api/v2/TimeTrackerService/ProjectTimeLogs | 
[**ApiV2TimeTrackerServiceProjectTimeLogsPost**](ProjectTimeLogsAPI.md#ApiV2TimeTrackerServiceProjectTimeLogsPost) | **Post** /api/v2/TimeTrackerService/ProjectTimeLogs | 
[**ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdDelete**](ProjectTimeLogsAPI.md#ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdDelete) | **Delete** /api/v2/TimeTrackerService/ProjectTimeLogs/{timeLogId} | 
[**ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGet**](ProjectTimeLogsAPI.md#ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGet) | **Get** /api/v2/TimeTrackerService/ProjectTimeLogs/{timeLogId} | 
[**ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdPut**](ProjectTimeLogsAPI.md#ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdPut) | **Put** /api/v2/TimeTrackerService/ProjectTimeLogs/{timeLogId} | 



## ApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGet

> ProjectTimeLogDtoListEnvelope ApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGet(ctx).ContactId(contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGet(context.Background()).ContactId(contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGet`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceProjectTimeLogsByResponsibleContactGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string** |  | 
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


## ApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGet

> ProjectTimeLogDtoListEnvelope ApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGet(ctx).ContactId(contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGet(context.Background()).ContactId(contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGet`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceProjectTimeLogsCreatedByContactGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **string** |  | 
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


## ApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGet

> ProjectTimeLogDtoListEnvelope ApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGet(ctx, projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGet(context.Background(), projectId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGet`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceProjectTimeLogsForProjectProjectIdGetRequest struct via the builder pattern


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


## ApiV2TimeTrackerServiceProjectTimeLogsGet

> ProjectTimeLogDtoListEnvelope ApiV2TimeTrackerServiceProjectTimeLogsGet(ctx).TenantId(tenantId).ProjectPeriodId(projectPeriodId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsGet(context.Background()).TenantId(tenantId).ProjectPeriodId(projectPeriodId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TimeTrackerServiceProjectTimeLogsGet`: ProjectTimeLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceProjectTimeLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **projectPeriodId** | **string** |  | 
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


## ApiV2TimeTrackerServiceProjectTimeLogsPost

> ApiV2TimeTrackerServiceProjectTimeLogsPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogCreateDto(projectTimeLogCreateDto).Execute()



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
	projectTimeLogCreateDto := *openapiclient.NewProjectTimeLogCreateDto("ProjectTaskID_example", "ProjectPeriodID_example") // ProjectTimeLogCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogCreateDto(projectTimeLogCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceProjectTimeLogsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectTimeLogCreateDto** | [**ProjectTimeLogCreateDto**](ProjectTimeLogCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdDelete

> ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdDelete(ctx, timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	r, err := apiClient.ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdDelete(context.Background(), timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGet

> ProjectTimeLogDtoEnvelope ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGet(ctx, timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGet(context.Background(), timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGet`: ProjectTimeLogDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeLogId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProjectTimeLogDtoEnvelope**](ProjectTimeLogDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdPut

> ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdPut(ctx, timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogUpdateDto(projectTimeLogUpdateDto).Execute()



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
	r, err := apiClient.ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdPut(context.Background(), timeLogId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectTimeLogUpdateDto(projectTimeLogUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTimeLogsAPI.ApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceProjectTimeLogsTimeLogIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectTimeLogUpdateDto** | [**ProjectTimeLogUpdateDto**](ProjectTimeLogUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

