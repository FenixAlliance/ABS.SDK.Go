# \TrainingProgramCoursesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrainingProgramCourseAsync**](TrainingProgramCoursesAPI.md#CreateTrainingProgramCourseAsync) | **Post** /api/v2/HrmsService/TrainingProgramCourses | Create a training program course
[**DeleteTrainingProgramCourseAsync**](TrainingProgramCoursesAPI.md#DeleteTrainingProgramCourseAsync) | **Delete** /api/v2/HrmsService/TrainingProgramCourses/{courseId} | Delete a training program course
[**GetTrainingProgramCourseByIdAsync**](TrainingProgramCoursesAPI.md#GetTrainingProgramCourseByIdAsync) | **Get** /api/v2/HrmsService/TrainingProgramCourses/{courseId} | Get training program course by ID
[**GetTrainingProgramCoursesAsync**](TrainingProgramCoursesAPI.md#GetTrainingProgramCoursesAsync) | **Get** /api/v2/HrmsService/TrainingProgramCourses | Get training program courses
[**GetTrainingProgramCoursesCountAsync**](TrainingProgramCoursesAPI.md#GetTrainingProgramCoursesCountAsync) | **Get** /api/v2/HrmsService/TrainingProgramCourses/Count | Count training program courses
[**UpdateTrainingProgramCourseAsync**](TrainingProgramCoursesAPI.md#UpdateTrainingProgramCourseAsync) | **Put** /api/v2/HrmsService/TrainingProgramCourses/{courseId} | Update a training program course



## CreateTrainingProgramCourseAsync

> EmptyEnvelope CreateTrainingProgramCourseAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TrainingProgramCourseCreateDto(trainingProgramCourseCreateDto).Execute()

Create a training program course



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
	trainingProgramCourseCreateDto := *openapiclient.NewTrainingProgramCourseCreateDto("TrainingProgramId_example", "CourseId_example") // TrainingProgramCourseCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainingProgramCoursesAPI.CreateTrainingProgramCourseAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TrainingProgramCourseCreateDto(trainingProgramCourseCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingProgramCoursesAPI.CreateTrainingProgramCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrainingProgramCourseAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingProgramCoursesAPI.CreateTrainingProgramCourseAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrainingProgramCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **trainingProgramCourseCreateDto** | [**TrainingProgramCourseCreateDto**](TrainingProgramCourseCreateDto.md) |  | 

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


## DeleteTrainingProgramCourseAsync

> EmptyEnvelope DeleteTrainingProgramCourseAsync(ctx, courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a training program course



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
	courseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainingProgramCoursesAPI.DeleteTrainingProgramCourseAsync(context.Background(), courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingProgramCoursesAPI.DeleteTrainingProgramCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrainingProgramCourseAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingProgramCoursesAPI.DeleteTrainingProgramCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrainingProgramCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## GetTrainingProgramCourseByIdAsync

> TrainingProgramCourseDtoEnvelope GetTrainingProgramCourseByIdAsync(ctx, courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get training program course by ID



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
	courseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainingProgramCoursesAPI.GetTrainingProgramCourseByIdAsync(context.Background(), courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingProgramCoursesAPI.GetTrainingProgramCourseByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainingProgramCourseByIdAsync`: TrainingProgramCourseDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingProgramCoursesAPI.GetTrainingProgramCourseByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainingProgramCourseByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TrainingProgramCourseDtoEnvelope**](TrainingProgramCourseDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainingProgramCoursesAsync

> TrainingProgramCourseDtoListEnvelope GetTrainingProgramCoursesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get training program courses



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
	resp, r, err := apiClient.TrainingProgramCoursesAPI.GetTrainingProgramCoursesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingProgramCoursesAPI.GetTrainingProgramCoursesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainingProgramCoursesAsync`: TrainingProgramCourseDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingProgramCoursesAPI.GetTrainingProgramCoursesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainingProgramCoursesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TrainingProgramCourseDtoListEnvelope**](TrainingProgramCourseDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrainingProgramCoursesCountAsync

> Int32Envelope GetTrainingProgramCoursesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count training program courses



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
	resp, r, err := apiClient.TrainingProgramCoursesAPI.GetTrainingProgramCoursesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingProgramCoursesAPI.GetTrainingProgramCoursesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrainingProgramCoursesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingProgramCoursesAPI.GetTrainingProgramCoursesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrainingProgramCoursesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## UpdateTrainingProgramCourseAsync

> EmptyEnvelope UpdateTrainingProgramCourseAsync(ctx, courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TrainingProgramCourseUpdateDto(trainingProgramCourseUpdateDto).Execute()

Update a training program course



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
	courseId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	trainingProgramCourseUpdateDto := *openapiclient.NewTrainingProgramCourseUpdateDto() // TrainingProgramCourseUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainingProgramCoursesAPI.UpdateTrainingProgramCourseAsync(context.Background(), courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TrainingProgramCourseUpdateDto(trainingProgramCourseUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingProgramCoursesAPI.UpdateTrainingProgramCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTrainingProgramCourseAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingProgramCoursesAPI.UpdateTrainingProgramCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTrainingProgramCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **trainingProgramCourseUpdateDto** | [**TrainingProgramCourseUpdateDto**](TrainingProgramCourseUpdateDto.md) |  | 

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

