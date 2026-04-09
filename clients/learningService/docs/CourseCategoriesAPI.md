# \CourseCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourseCategoryAsync**](CourseCategoriesAPI.md#CreateCourseCategoryAsync) | **Post** /api/v2/LearningService/CourseCategories | Create a new course category
[**DeleteCourseCategoryAsync**](CourseCategoriesAPI.md#DeleteCourseCategoryAsync) | **Delete** /api/v2/LearningService/CourseCategories/{categoryId} | Delete a course category
[**GetCourseCategoriesAsync**](CourseCategoriesAPI.md#GetCourseCategoriesAsync) | **Get** /api/v2/LearningService/CourseCategories | Get all course categories
[**GetCourseCategoriesCountAsync**](CourseCategoriesAPI.md#GetCourseCategoriesCountAsync) | **Get** /api/v2/LearningService/CourseCategories/Count | Get course categories count
[**GetCourseCategoryByIdAsync**](CourseCategoriesAPI.md#GetCourseCategoryByIdAsync) | **Get** /api/v2/LearningService/CourseCategories/{categoryId} | Get course category by ID
[**UpdateCourseCategoryAsync**](CourseCategoriesAPI.md#UpdateCourseCategoryAsync) | **Put** /api/v2/LearningService/CourseCategories/{categoryId} | Update a course category



## CreateCourseCategoryAsync

> CreateCourseCategoryAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCategoryCreateDto(courseCategoryCreateDto).Execute()

Create a new course category



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
	courseCategoryCreateDto := *openapiclient.NewCourseCategoryCreateDto("Title_example", "BusinessID_example") // CourseCategoryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseCategoriesAPI.CreateCourseCategoryAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCategoryCreateDto(courseCategoryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseCategoriesAPI.CreateCourseCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseCategoryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseCategoryCreateDto** | [**CourseCategoryCreateDto**](CourseCategoryCreateDto.md) |  | 

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


## DeleteCourseCategoryAsync

> DeleteCourseCategoryAsync(ctx, categoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a course category



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
	categoryId := "categoryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseCategoriesAPI.DeleteCourseCategoryAsync(context.Background(), categoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseCategoriesAPI.DeleteCourseCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCourseCategoryAsyncRequest struct via the builder pattern


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


## GetCourseCategoriesAsync

> []CourseCategoryDto GetCourseCategoriesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all course categories



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
	resp, r, err := apiClient.CourseCategoriesAPI.GetCourseCategoriesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseCategoriesAPI.GetCourseCategoriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseCategoriesAsync`: []CourseCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `CourseCategoriesAPI.GetCourseCategoriesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseCategoriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseCategoryDto**](CourseCategoryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseCategoriesCountAsync

> int32 GetCourseCategoriesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course categories count



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
	resp, r, err := apiClient.CourseCategoriesAPI.GetCourseCategoriesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseCategoriesAPI.GetCourseCategoriesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseCategoriesCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CourseCategoriesAPI.GetCourseCategoriesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseCategoriesCountAsyncRequest struct via the builder pattern


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


## GetCourseCategoryByIdAsync

> CourseCategoryDto GetCourseCategoryByIdAsync(ctx, categoryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course category by ID



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
	categoryId := "categoryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseCategoriesAPI.GetCourseCategoryByIdAsync(context.Background(), categoryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseCategoriesAPI.GetCourseCategoryByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseCategoryByIdAsync`: CourseCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `CourseCategoriesAPI.GetCourseCategoryByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseCategoryByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CourseCategoryDto**](CourseCategoryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCourseCategoryAsync

> UpdateCourseCategoryAsync(ctx, categoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCategoryUpdateDto(courseCategoryUpdateDto).Execute()

Update a course category



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
	categoryId := "categoryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseCategoryUpdateDto := *openapiclient.NewCourseCategoryUpdateDto() // CourseCategoryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseCategoriesAPI.UpdateCourseCategoryAsync(context.Background(), categoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCategoryUpdateDto(courseCategoryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseCategoriesAPI.UpdateCourseCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCourseCategoryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseCategoryUpdateDto** | [**CourseCategoryUpdateDto**](CourseCategoryUpdateDto.md) |  | 

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

