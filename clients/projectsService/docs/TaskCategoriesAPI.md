# \TaskCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2ProjectsServiceTaskCategoriesGet**](TaskCategoriesAPI.md#ApiV2ProjectsServiceTaskCategoriesGet) | **Get** /api/v2/ProjectsService/TaskCategories | 
[**ApiV2ProjectsServiceTaskCategoriesPost**](TaskCategoriesAPI.md#ApiV2ProjectsServiceTaskCategoriesPost) | **Post** /api/v2/ProjectsService/TaskCategories | 
[**ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDelete**](TaskCategoriesAPI.md#ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDelete) | **Delete** /api/v2/ProjectsService/TaskCategories/{taskCategoryId} | 
[**ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGet**](TaskCategoriesAPI.md#ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGet) | **Get** /api/v2/ProjectsService/TaskCategories/{taskCategoryId} | 
[**ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPut**](TaskCategoriesAPI.md#ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPut) | **Put** /api/v2/ProjectsService/TaskCategories/{taskCategoryId} | 
[**ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGet**](TaskCategoriesAPI.md#ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGet) | **Get** /api/v2/ProjectsService/TaskCategories/{taskCategoryId}/Types | 



## ApiV2ProjectsServiceTaskCategoriesGet

> TaskCategoryDtoListEnvelope ApiV2ProjectsServiceTaskCategoriesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskCategoriesGet`: TaskCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskCategoriesGetRequest struct via the builder pattern


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


## ApiV2ProjectsServiceTaskCategoriesPost

> TaskCategoryDto ApiV2ProjectsServiceTaskCategoriesPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaskCategoryCreateDto(taskCategoryCreateDto).Execute()



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
	taskCategoryCreateDto := *openapiclient.NewTaskCategoryCreateDto() // TaskCategoryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaskCategoryCreateDto(taskCategoryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskCategoriesPost`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **taskCategoryCreateDto** | [**TaskCategoryCreateDto**](TaskCategoryCreateDto.md) |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDelete

> TaskCategoryDto ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDelete(ctx, taskCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDelete(context.Background(), taskCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDelete`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskCategoriesTaskCategoryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGet

> TaskCategoryDto ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGet(ctx, taskCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGet(context.Background(), taskCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGet`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskCategoriesTaskCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPut

> TaskCategoryDto ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPut(ctx, taskCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaskCategoryUpdateDto(taskCategoryUpdateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	taskCategoryUpdateDto := *openapiclient.NewTaskCategoryUpdateDto() // TaskCategoryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPut(context.Background(), taskCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaskCategoryUpdateDto(taskCategoryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPut`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskCategoriesTaskCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **taskCategoryUpdateDto** | [**TaskCategoryUpdateDto**](TaskCategoryUpdateDto.md) |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGet

> TaskCategoryDto ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGet(ctx, taskCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGet(context.Background(), taskCategoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGet`: TaskCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `TaskCategoriesAPI.ApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskCategoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2ProjectsServiceTaskCategoriesTaskCategoryIdTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaskCategoryDto**](TaskCategoryDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

