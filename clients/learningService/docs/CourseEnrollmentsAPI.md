# \CourseEnrollmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourseEnrollmentAsync**](CourseEnrollmentsAPI.md#CreateCourseEnrollmentAsync) | **Post** /api/v2/LearningService/CourseEnrollments | Create a new course enrollment
[**DeleteCourseEnrollmentAsync**](CourseEnrollmentsAPI.md#DeleteCourseEnrollmentAsync) | **Delete** /api/v2/LearningService/CourseEnrollments/{courseEnrollmentId} | Delete a course enrollment
[**GetCourseEnrollmentAsync**](CourseEnrollmentsAPI.md#GetCourseEnrollmentAsync) | **Get** /api/v2/LearningService/CourseEnrollments/{courseEnrollmentId} | Get course enrollment by ID
[**GetEnrollmentsAsync**](CourseEnrollmentsAPI.md#GetEnrollmentsAsync) | **Get** /api/v2/LearningService/CourseEnrollments | Get all course enrollments
[**GetEnrollmentsCountAsync**](CourseEnrollmentsAPI.md#GetEnrollmentsCountAsync) | **Get** /api/v2/LearningService/CourseEnrollments/Count | Get course enrollments count
[**GetStudentCourseEnrollmentsAsync**](CourseEnrollmentsAPI.md#GetStudentCourseEnrollmentsAsync) | **Get** /api/v2/LearningService/CourseEnrollments/Student/{studentProfileId} | Get enrollments by student
[**UpdateCourseEnrollmentAsync**](CourseEnrollmentsAPI.md#UpdateCourseEnrollmentAsync) | **Put** /api/v2/LearningService/CourseEnrollments/{courseEnrollmentId} | Update a course enrollment



## CreateCourseEnrollmentAsync

> CreateCourseEnrollmentAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseEnrollmentCreateDto(courseEnrollmentCreateDto).Execute()

Create a new course enrollment



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
	courseEnrollmentCreateDto := *openapiclient.NewCourseEnrollmentCreateDto() // CourseEnrollmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseEnrollmentsAPI.CreateCourseEnrollmentAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseEnrollmentCreateDto(courseEnrollmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseEnrollmentsAPI.CreateCourseEnrollmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseEnrollmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseEnrollmentCreateDto** | [**CourseEnrollmentCreateDto**](CourseEnrollmentCreateDto.md) |  | 

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


## DeleteCourseEnrollmentAsync

> DeleteCourseEnrollmentAsync(ctx, courseEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a course enrollment



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
	courseEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseEnrollmentsAPI.DeleteCourseEnrollmentAsync(context.Background(), courseEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseEnrollmentsAPI.DeleteCourseEnrollmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCourseEnrollmentAsyncRequest struct via the builder pattern


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


## GetCourseEnrollmentAsync

> CourseEnrollmentDto GetCourseEnrollmentAsync(ctx, courseEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course enrollment by ID



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
	courseEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseEnrollmentsAPI.GetCourseEnrollmentAsync(context.Background(), courseEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseEnrollmentsAPI.GetCourseEnrollmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseEnrollmentAsync`: CourseEnrollmentDto
	fmt.Fprintf(os.Stdout, "Response from `CourseEnrollmentsAPI.GetCourseEnrollmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseEnrollmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CourseEnrollmentDto**](CourseEnrollmentDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentsAsync

> []CourseEnrollmentDto GetEnrollmentsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all course enrollments



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
	resp, r, err := apiClient.CourseEnrollmentsAPI.GetEnrollmentsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseEnrollmentsAPI.GetEnrollmentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnrollmentsAsync`: []CourseEnrollmentDto
	fmt.Fprintf(os.Stdout, "Response from `CourseEnrollmentsAPI.GetEnrollmentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseEnrollmentDto**](CourseEnrollmentDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentsCountAsync

> int32 GetEnrollmentsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course enrollments count



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
	resp, r, err := apiClient.CourseEnrollmentsAPI.GetEnrollmentsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseEnrollmentsAPI.GetEnrollmentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnrollmentsCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CourseEnrollmentsAPI.GetEnrollmentsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentsCountAsyncRequest struct via the builder pattern


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


## GetStudentCourseEnrollmentsAsync

> []CourseEnrollmentDto GetStudentCourseEnrollmentsAsync(ctx, studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get enrollments by student



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
	studentProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseEnrollmentsAPI.GetStudentCourseEnrollmentsAsync(context.Background(), studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseEnrollmentsAPI.GetStudentCourseEnrollmentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudentCourseEnrollmentsAsync`: []CourseEnrollmentDto
	fmt.Fprintf(os.Stdout, "Response from `CourseEnrollmentsAPI.GetStudentCourseEnrollmentsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studentProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudentCourseEnrollmentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseEnrollmentDto**](CourseEnrollmentDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCourseEnrollmentAsync

> UpdateCourseEnrollmentAsync(ctx, courseEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseEnrollmentUpdateDto(courseEnrollmentUpdateDto).Execute()

Update a course enrollment



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
	courseEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseEnrollmentUpdateDto := *openapiclient.NewCourseEnrollmentUpdateDto() // CourseEnrollmentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseEnrollmentsAPI.UpdateCourseEnrollmentAsync(context.Background(), courseEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseEnrollmentUpdateDto(courseEnrollmentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseEnrollmentsAPI.UpdateCourseEnrollmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCourseEnrollmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseEnrollmentUpdateDto** | [**CourseEnrollmentUpdateDto**](CourseEnrollmentUpdateDto.md) |  | 

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

