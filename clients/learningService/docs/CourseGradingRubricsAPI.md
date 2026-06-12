# \CourseGradingRubricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourseGradingRubricAsync**](CourseGradingRubricsAPI.md#CreateCourseGradingRubricAsync) | **Post** /api/v2/LearningService/CourseGradingRubrics | Create a course grading rubric
[**DeleteCourseGradingRubricAsync**](CourseGradingRubricsAPI.md#DeleteCourseGradingRubricAsync) | **Delete** /api/v2/LearningService/CourseGradingRubrics/{rubricId} | Delete a course grading rubric
[**GetCourseGradingRubricByIdAsync**](CourseGradingRubricsAPI.md#GetCourseGradingRubricByIdAsync) | **Get** /api/v2/LearningService/CourseGradingRubrics/{rubricId} | Get course grading rubric by ID
[**GetCourseGradingRubricsAsync**](CourseGradingRubricsAPI.md#GetCourseGradingRubricsAsync) | **Get** /api/v2/LearningService/CourseGradingRubrics | Get all course grading rubrics
[**GetCourseGradingRubricsCountAsync**](CourseGradingRubricsAPI.md#GetCourseGradingRubricsCountAsync) | **Get** /api/v2/LearningService/CourseGradingRubrics/Count | Get course grading rubrics count
[**PatchCourseGradingRubricAsync**](CourseGradingRubricsAPI.md#PatchCourseGradingRubricAsync) | **Patch** /api/v2/LearningService/CourseGradingRubrics/{rubricId} | Patch a course grading rubric
[**UpdateCourseGradingRubricAsync**](CourseGradingRubricsAPI.md#UpdateCourseGradingRubricAsync) | **Put** /api/v2/LearningService/CourseGradingRubrics/{rubricId} | Update a course grading rubric



## CreateCourseGradingRubricAsync

> CreateCourseGradingRubricAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseGradingRubricCreateDto(courseGradingRubricCreateDto).Execute()

Create a course grading rubric



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
	courseGradingRubricCreateDto := *openapiclient.NewCourseGradingRubricCreateDto("Title_example", "CourseId_example") // CourseGradingRubricCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseGradingRubricsAPI.CreateCourseGradingRubricAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseGradingRubricCreateDto(courseGradingRubricCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseGradingRubricsAPI.CreateCourseGradingRubricAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseGradingRubricAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseGradingRubricCreateDto** | [**CourseGradingRubricCreateDto**](CourseGradingRubricCreateDto.md) |  | 

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


## DeleteCourseGradingRubricAsync

> DeleteCourseGradingRubricAsync(ctx, rubricId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a course grading rubric



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
	rubricId := "rubricId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseGradingRubricsAPI.DeleteCourseGradingRubricAsync(context.Background(), rubricId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseGradingRubricsAPI.DeleteCourseGradingRubricAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rubricId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCourseGradingRubricAsyncRequest struct via the builder pattern


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


## GetCourseGradingRubricByIdAsync

> CourseGradingRubricDto GetCourseGradingRubricByIdAsync(ctx, rubricId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course grading rubric by ID



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
	rubricId := "rubricId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseGradingRubricsAPI.GetCourseGradingRubricByIdAsync(context.Background(), rubricId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseGradingRubricsAPI.GetCourseGradingRubricByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseGradingRubricByIdAsync`: CourseGradingRubricDto
	fmt.Fprintf(os.Stdout, "Response from `CourseGradingRubricsAPI.GetCourseGradingRubricByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rubricId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseGradingRubricByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CourseGradingRubricDto**](CourseGradingRubricDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseGradingRubricsAsync

> []CourseGradingRubricDto GetCourseGradingRubricsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all course grading rubrics



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
	resp, r, err := apiClient.CourseGradingRubricsAPI.GetCourseGradingRubricsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseGradingRubricsAPI.GetCourseGradingRubricsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseGradingRubricsAsync`: []CourseGradingRubricDto
	fmt.Fprintf(os.Stdout, "Response from `CourseGradingRubricsAPI.GetCourseGradingRubricsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseGradingRubricsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseGradingRubricDto**](CourseGradingRubricDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseGradingRubricsCountAsync

> int32 GetCourseGradingRubricsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course grading rubrics count



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
	resp, r, err := apiClient.CourseGradingRubricsAPI.GetCourseGradingRubricsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseGradingRubricsAPI.GetCourseGradingRubricsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseGradingRubricsCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CourseGradingRubricsAPI.GetCourseGradingRubricsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseGradingRubricsCountAsyncRequest struct via the builder pattern


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


## PatchCourseGradingRubricAsync

> PatchCourseGradingRubricAsync(ctx, rubricId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a course grading rubric



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
	rubricId := "rubricId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseGradingRubricsAPI.PatchCourseGradingRubricAsync(context.Background(), rubricId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseGradingRubricsAPI.PatchCourseGradingRubricAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rubricId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCourseGradingRubricAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## UpdateCourseGradingRubricAsync

> UpdateCourseGradingRubricAsync(ctx, rubricId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseGradingRubricUpdateDto(courseGradingRubricUpdateDto).Execute()

Update a course grading rubric



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
	rubricId := "rubricId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseGradingRubricUpdateDto := *openapiclient.NewCourseGradingRubricUpdateDto() // CourseGradingRubricUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseGradingRubricsAPI.UpdateCourseGradingRubricAsync(context.Background(), rubricId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseGradingRubricUpdateDto(courseGradingRubricUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseGradingRubricsAPI.UpdateCourseGradingRubricAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rubricId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCourseGradingRubricAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseGradingRubricUpdateDto** | [**CourseGradingRubricUpdateDto**](CourseGradingRubricUpdateDto.md) |  | 

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

