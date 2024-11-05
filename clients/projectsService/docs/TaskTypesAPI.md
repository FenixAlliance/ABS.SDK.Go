# \TaskTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsServiceTaskTypesPost**](TaskTypesAPI.md#ApiV2ProjectsServiceTaskTypesPost) | **Post** /api/v2/ProjectsService/TaskTypes | 
[**ApiV2ProjectsServiceTaskTypesTaskTypeIdDelete**](TaskTypesAPI.md#ApiV2ProjectsServiceTaskTypesTaskTypeIdDelete) | **Delete** /api/v2/ProjectsService/TaskTypes/{taskTypeId} | 
[**ApiV2ProjectsServiceTaskTypesTaskTypeIdGet**](TaskTypesAPI.md#ApiV2ProjectsServiceTaskTypesTaskTypeIdGet) | **Get** /api/v2/ProjectsService/TaskTypes/{taskTypeId} | 
[**ApiV2ProjectsServiceTaskTypesTaskTypeIdPut**](TaskTypesAPI.md#ApiV2ProjectsServiceTaskTypesTaskTypeIdPut) | **Put** /api/v2/ProjectsService/TaskTypes/{taskTypeId} | 



## ApiV2ProjectsServiceTaskTypesPost

> TaskTypeDto ApiV2ProjectsServiceTaskTypesPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaskTypeCreateDto(taskTypeCreateDto).Execute()



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
	taskTypeCreateDto := *openapiclient.NewTaskTypeCreateDto() // TaskTypeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskTypesAPI.ApiV2ProjectsServiceTaskTypesPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaskTypeCreateDto(taskTypeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskTypesAPI.ApiV2ProjectsServiceTaskTypesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskTypesPost`: TaskTypeDto
	fmt.Fprintf(os.Stdout, "Response from `TaskTypesAPI.ApiV2ProjectsServiceTaskTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **taskTypeCreateDto** | [**TaskTypeCreateDto**](TaskTypeCreateDto.md) |  | 

### Return type

[**TaskTypeDto**](TaskTypeDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceTaskTypesTaskTypeIdDelete

> TaskTypeDto ApiV2ProjectsServiceTaskTypesTaskTypeIdDelete(ctx, taskTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdDelete(context.Background(), taskTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskTypesTaskTypeIdDelete`: TaskTypeDto
	fmt.Fprintf(os.Stdout, "Response from `TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskTypesTaskTypeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaskTypeDto**](TaskTypeDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceTaskTypesTaskTypeIdGet

> TaskTypeDto ApiV2ProjectsServiceTaskTypesTaskTypeIdGet(ctx, taskTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdGet(context.Background(), taskTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskTypesTaskTypeIdGet`: TaskTypeDto
	fmt.Fprintf(os.Stdout, "Response from `TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskTypesTaskTypeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaskTypeDto**](TaskTypeDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceTaskTypesTaskTypeIdPut

> TaskTypeDto ApiV2ProjectsServiceTaskTypesTaskTypeIdPut(ctx, taskTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaskTypeUpdateDto(taskTypeUpdateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	taskTypeUpdateDto := *openapiclient.NewTaskTypeUpdateDto() // TaskTypeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdPut(context.Background(), taskTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaskTypeUpdateDto(taskTypeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskTypesTaskTypeIdPut`: TaskTypeDto
	fmt.Fprintf(os.Stdout, "Response from `TaskTypesAPI.ApiV2ProjectsServiceTaskTypesTaskTypeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskTypesTaskTypeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **taskTypeUpdateDto** | [**TaskTypeUpdateDto**](TaskTypeUpdateDto.md) |  | 

### Return type

[**TaskTypeDto**](TaskTypeDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

