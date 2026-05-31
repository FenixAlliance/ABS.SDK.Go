# \CourseAssignmentComponentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourseAssignmentComponentAsync**](CourseAssignmentComponentsAPI.md#CreateCourseAssignmentComponentAsync) | **Post** /api/v2/LearningService/CourseAssignmentComponents | Create a course assignment component
[**DeleteCourseAssignmentComponentAsync**](CourseAssignmentComponentsAPI.md#DeleteCourseAssignmentComponentAsync) | **Delete** /api/v2/LearningService/CourseAssignmentComponents/{componentId} | Delete a course assignment component
[**GetCourseAssignmentComponentByIdAsync**](CourseAssignmentComponentsAPI.md#GetCourseAssignmentComponentByIdAsync) | **Get** /api/v2/LearningService/CourseAssignmentComponents/{componentId} | Get course assignment component by ID
[**GetCourseAssignmentComponentsAsync**](CourseAssignmentComponentsAPI.md#GetCourseAssignmentComponentsAsync) | **Get** /api/v2/LearningService/CourseAssignmentComponents | Get all course assignment components
[**GetCourseAssignmentComponentsCountAsync**](CourseAssignmentComponentsAPI.md#GetCourseAssignmentComponentsCountAsync) | **Get** /api/v2/LearningService/CourseAssignmentComponents/Count | Get course assignment components count
[**UpdateCourseAssignmentComponentAsync**](CourseAssignmentComponentsAPI.md#UpdateCourseAssignmentComponentAsync) | **Put** /api/v2/LearningService/CourseAssignmentComponents/{componentId} | Update a course assignment component



## CreateCourseAssignmentComponentAsync

> CreateCourseAssignmentComponentAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseAssignmentComponentCreateDto(courseAssignmentComponentCreateDto).Execute()

Create a course assignment component



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
	courseAssignmentComponentCreateDto := *openapiclient.NewCourseAssignmentComponentCreateDto("Title_example", "CourseAssignmentID_example", "CourseID_example") // CourseAssignmentComponentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseAssignmentComponentsAPI.CreateCourseAssignmentComponentAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseAssignmentComponentCreateDto(courseAssignmentComponentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAssignmentComponentsAPI.CreateCourseAssignmentComponentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseAssignmentComponentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseAssignmentComponentCreateDto** | [**CourseAssignmentComponentCreateDto**](CourseAssignmentComponentCreateDto.md) |  | 

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


## DeleteCourseAssignmentComponentAsync

> DeleteCourseAssignmentComponentAsync(ctx, componentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a course assignment component



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
	componentId := "componentId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseAssignmentComponentsAPI.DeleteCourseAssignmentComponentAsync(context.Background(), componentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAssignmentComponentsAPI.DeleteCourseAssignmentComponentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCourseAssignmentComponentAsyncRequest struct via the builder pattern


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


## GetCourseAssignmentComponentByIdAsync

> CourseAssignmentComponentDto GetCourseAssignmentComponentByIdAsync(ctx, componentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course assignment component by ID



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
	componentId := "componentId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseAssignmentComponentsAPI.GetCourseAssignmentComponentByIdAsync(context.Background(), componentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAssignmentComponentsAPI.GetCourseAssignmentComponentByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseAssignmentComponentByIdAsync`: CourseAssignmentComponentDto
	fmt.Fprintf(os.Stdout, "Response from `CourseAssignmentComponentsAPI.GetCourseAssignmentComponentByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseAssignmentComponentByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CourseAssignmentComponentDto**](CourseAssignmentComponentDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseAssignmentComponentsAsync

> []CourseAssignmentComponentDto GetCourseAssignmentComponentsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all course assignment components



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
	resp, r, err := apiClient.CourseAssignmentComponentsAPI.GetCourseAssignmentComponentsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAssignmentComponentsAPI.GetCourseAssignmentComponentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseAssignmentComponentsAsync`: []CourseAssignmentComponentDto
	fmt.Fprintf(os.Stdout, "Response from `CourseAssignmentComponentsAPI.GetCourseAssignmentComponentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseAssignmentComponentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseAssignmentComponentDto**](CourseAssignmentComponentDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseAssignmentComponentsCountAsync

> int32 GetCourseAssignmentComponentsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course assignment components count



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
	resp, r, err := apiClient.CourseAssignmentComponentsAPI.GetCourseAssignmentComponentsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAssignmentComponentsAPI.GetCourseAssignmentComponentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseAssignmentComponentsCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CourseAssignmentComponentsAPI.GetCourseAssignmentComponentsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseAssignmentComponentsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCourseAssignmentComponentAsync

> UpdateCourseAssignmentComponentAsync(ctx, componentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseAssignmentComponentUpdateDto(courseAssignmentComponentUpdateDto).Execute()

Update a course assignment component



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
	componentId := "componentId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseAssignmentComponentUpdateDto := *openapiclient.NewCourseAssignmentComponentUpdateDto() // CourseAssignmentComponentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseAssignmentComponentsAPI.UpdateCourseAssignmentComponentAsync(context.Background(), componentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseAssignmentComponentUpdateDto(courseAssignmentComponentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAssignmentComponentsAPI.UpdateCourseAssignmentComponentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCourseAssignmentComponentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseAssignmentComponentUpdateDto** | [**CourseAssignmentComponentUpdateDto**](CourseAssignmentComponentUpdateDto.md) |  | 

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

