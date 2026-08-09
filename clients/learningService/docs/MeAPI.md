# \MeAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyAverageScoreAsync**](MeAPI.md#GetMyAverageScoreAsync) | **Get** /api/v2/LearningService/Me/AverageScore | Get current user&#39;s average score
[**GetMyCertificatesAsync**](MeAPI.md#GetMyCertificatesAsync) | **Get** /api/v2/LearningService/Me/Certificates | Get current user&#39;s completion certificates
[**GetMyCertificatesCountAsync**](MeAPI.md#GetMyCertificatesCountAsync) | **Get** /api/v2/LearningService/Me/Certificates/Count | Get current user&#39;s certificates count
[**GetMyEnrollmentsAsync**](MeAPI.md#GetMyEnrollmentsAsync) | **Get** /api/v2/LearningService/Me/Enrollments | Get current user&#39;s course enrollments
[**GetMyEnrollmentsCountAsync**](MeAPI.md#GetMyEnrollmentsCountAsync) | **Get** /api/v2/LearningService/Me/Enrollments/Count | Get current user&#39;s enrollment count
[**GetMyHoursCompletedAsync**](MeAPI.md#GetMyHoursCompletedAsync) | **Get** /api/v2/LearningService/Me/HoursCompleted | Get current user&#39;s completed hours
[**GetMyInstructorCoursesAsync**](MeAPI.md#GetMyInstructorCoursesAsync) | **Get** /api/v2/LearningService/Me/InstructorCourses | Get current user&#39;s instructor courses
[**GetMyInstructorCoursesCountAsync**](MeAPI.md#GetMyInstructorCoursesCountAsync) | **Get** /api/v2/LearningService/Me/InstructorCourses/Count | Get current user&#39;s instructor courses count
[**GetMyInstructorProfilesAsync**](MeAPI.md#GetMyInstructorProfilesAsync) | **Get** /api/v2/LearningService/Me/InstructorProfiles | Get current user&#39;s instructor profiles
[**GetMyInstructorProfilesCountAsync**](MeAPI.md#GetMyInstructorProfilesCountAsync) | **Get** /api/v2/LearningService/Me/InstructorProfiles/Count | Get current user&#39;s instructor profiles count
[**GetMyPendingTaskCountAsync**](MeAPI.md#GetMyPendingTaskCountAsync) | **Get** /api/v2/LearningService/Me/PendingTasks | Get current user&#39;s pending task count
[**GetMyStudentCoursesAsync**](MeAPI.md#GetMyStudentCoursesAsync) | **Get** /api/v2/LearningService/Me/Courses | Get current user&#39;s enrolled courses
[**GetMyStudentCoursesCountAsync**](MeAPI.md#GetMyStudentCoursesCountAsync) | **Get** /api/v2/LearningService/Me/Courses/Count | Get current user&#39;s enrolled courses count
[**GetMyStudentProfilesAsync**](MeAPI.md#GetMyStudentProfilesAsync) | **Get** /api/v2/LearningService/Me/StudentProfiles | Get current user&#39;s student profiles
[**GetMyStudentProfilesCountAsync**](MeAPI.md#GetMyStudentProfilesCountAsync) | **Get** /api/v2/LearningService/Me/StudentProfiles/Count | Get current user&#39;s student profiles count



## GetMyAverageScoreAsync

> AverageDtoEnvelope GetMyAverageScoreAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get current user's average score

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyAverageScoreAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyAverageScoreAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyAverageScoreAsync`: AverageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyAverageScoreAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyAverageScoreAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AverageDtoEnvelope**](AverageDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyCertificatesAsync

> CourseCompletionCertificateDtoIReadOnlyListEnvelope GetMyCertificatesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCompletionCertificateDtoCollectionQueryParameters(courseCompletionCertificateDtoCollectionQueryParameters).Execute()

Get current user's completion certificates

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseCompletionCertificateDtoCollectionQueryParameters := *openapiclient.NewCourseCompletionCertificateDtoCollectionQueryParameters() // CourseCompletionCertificateDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyCertificatesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCompletionCertificateDtoCollectionQueryParameters(courseCompletionCertificateDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyCertificatesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyCertificatesAsync`: CourseCompletionCertificateDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyCertificatesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyCertificatesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseCompletionCertificateDtoCollectionQueryParameters** | [**CourseCompletionCertificateDtoCollectionQueryParameters**](CourseCompletionCertificateDtoCollectionQueryParameters.md) |  | 

### Return type

[**CourseCompletionCertificateDtoIReadOnlyListEnvelope**](CourseCompletionCertificateDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyCertificatesCountAsync

> int32 GetMyCertificatesCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCompletionCertificateDtoCollectionQueryParameters(courseCompletionCertificateDtoCollectionQueryParameters).Execute()

Get current user's certificates count

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseCompletionCertificateDtoCollectionQueryParameters := *openapiclient.NewCourseCompletionCertificateDtoCollectionQueryParameters() // CourseCompletionCertificateDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyCertificatesCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCompletionCertificateDtoCollectionQueryParameters(courseCompletionCertificateDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyCertificatesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyCertificatesCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyCertificatesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyCertificatesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseCompletionCertificateDtoCollectionQueryParameters** | [**CourseCompletionCertificateDtoCollectionQueryParameters**](CourseCompletionCertificateDtoCollectionQueryParameters.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyEnrollmentsAsync

> CourseEnrollmentDtoIReadOnlyListEnvelope GetMyEnrollmentsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseEnrollmentDtoCollectionQueryParameters(courseEnrollmentDtoCollectionQueryParameters).Execute()

Get current user's course enrollments

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseEnrollmentDtoCollectionQueryParameters := *openapiclient.NewCourseEnrollmentDtoCollectionQueryParameters() // CourseEnrollmentDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyEnrollmentsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseEnrollmentDtoCollectionQueryParameters(courseEnrollmentDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyEnrollmentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyEnrollmentsAsync`: CourseEnrollmentDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyEnrollmentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyEnrollmentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseEnrollmentDtoCollectionQueryParameters** | [**CourseEnrollmentDtoCollectionQueryParameters**](CourseEnrollmentDtoCollectionQueryParameters.md) |  | 

### Return type

[**CourseEnrollmentDtoIReadOnlyListEnvelope**](CourseEnrollmentDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyEnrollmentsCountAsync

> int32 GetMyEnrollmentsCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseEnrollmentDtoCollectionQueryParameters(courseEnrollmentDtoCollectionQueryParameters).Execute()

Get current user's enrollment count

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseEnrollmentDtoCollectionQueryParameters := *openapiclient.NewCourseEnrollmentDtoCollectionQueryParameters() // CourseEnrollmentDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyEnrollmentsCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseEnrollmentDtoCollectionQueryParameters(courseEnrollmentDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyEnrollmentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyEnrollmentsCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyEnrollmentsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyEnrollmentsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseEnrollmentDtoCollectionQueryParameters** | [**CourseEnrollmentDtoCollectionQueryParameters**](CourseEnrollmentDtoCollectionQueryParameters.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyHoursCompletedAsync

> CountDtoEnvelope GetMyHoursCompletedAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get current user's completed hours

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyHoursCompletedAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyHoursCompletedAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyHoursCompletedAsync`: CountDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyHoursCompletedAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyHoursCompletedAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountDtoEnvelope**](CountDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyInstructorCoursesAsync

> CourseDtoIReadOnlyListEnvelope GetMyInstructorCoursesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseDtoCollectionQueryParameters(courseDtoCollectionQueryParameters).Execute()

Get current user's instructor courses

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseDtoCollectionQueryParameters := *openapiclient.NewCourseDtoCollectionQueryParameters() // CourseDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyInstructorCoursesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseDtoCollectionQueryParameters(courseDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyInstructorCoursesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyInstructorCoursesAsync`: CourseDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyInstructorCoursesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyInstructorCoursesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseDtoCollectionQueryParameters** | [**CourseDtoCollectionQueryParameters**](CourseDtoCollectionQueryParameters.md) |  | 

### Return type

[**CourseDtoIReadOnlyListEnvelope**](CourseDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyInstructorCoursesCountAsync

> int32 GetMyInstructorCoursesCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseDtoCollectionQueryParameters(courseDtoCollectionQueryParameters).Execute()

Get current user's instructor courses count

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseDtoCollectionQueryParameters := *openapiclient.NewCourseDtoCollectionQueryParameters() // CourseDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyInstructorCoursesCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseDtoCollectionQueryParameters(courseDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyInstructorCoursesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyInstructorCoursesCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyInstructorCoursesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyInstructorCoursesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseDtoCollectionQueryParameters** | [**CourseDtoCollectionQueryParameters**](CourseDtoCollectionQueryParameters.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyInstructorProfilesAsync

> InstructorProfileDtoIReadOnlyListEnvelope GetMyInstructorProfilesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).InstructorProfileDtoCollectionQueryParameters(instructorProfileDtoCollectionQueryParameters).Execute()

Get current user's instructor profiles

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	instructorProfileDtoCollectionQueryParameters := *openapiclient.NewInstructorProfileDtoCollectionQueryParameters() // InstructorProfileDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyInstructorProfilesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).InstructorProfileDtoCollectionQueryParameters(instructorProfileDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyInstructorProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyInstructorProfilesAsync`: InstructorProfileDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyInstructorProfilesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyInstructorProfilesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **instructorProfileDtoCollectionQueryParameters** | [**InstructorProfileDtoCollectionQueryParameters**](InstructorProfileDtoCollectionQueryParameters.md) |  | 

### Return type

[**InstructorProfileDtoIReadOnlyListEnvelope**](InstructorProfileDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyInstructorProfilesCountAsync

> int32 GetMyInstructorProfilesCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).InstructorProfileDtoCollectionQueryParameters(instructorProfileDtoCollectionQueryParameters).Execute()

Get current user's instructor profiles count

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	instructorProfileDtoCollectionQueryParameters := *openapiclient.NewInstructorProfileDtoCollectionQueryParameters() // InstructorProfileDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyInstructorProfilesCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).InstructorProfileDtoCollectionQueryParameters(instructorProfileDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyInstructorProfilesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyInstructorProfilesCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyInstructorProfilesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyInstructorProfilesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **instructorProfileDtoCollectionQueryParameters** | [**InstructorProfileDtoCollectionQueryParameters**](InstructorProfileDtoCollectionQueryParameters.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPendingTaskCountAsync

> CountDtoEnvelope GetMyPendingTaskCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get current user's pending task count

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyPendingTaskCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyPendingTaskCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPendingTaskCountAsync`: CountDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyPendingTaskCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPendingTaskCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountDtoEnvelope**](CountDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyStudentCoursesAsync

> CourseDtoIReadOnlyListEnvelope GetMyStudentCoursesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseDtoCollectionQueryParameters(courseDtoCollectionQueryParameters).Execute()

Get current user's enrolled courses

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseDtoCollectionQueryParameters := *openapiclient.NewCourseDtoCollectionQueryParameters() // CourseDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyStudentCoursesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseDtoCollectionQueryParameters(courseDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyStudentCoursesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyStudentCoursesAsync`: CourseDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyStudentCoursesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyStudentCoursesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseDtoCollectionQueryParameters** | [**CourseDtoCollectionQueryParameters**](CourseDtoCollectionQueryParameters.md) |  | 

### Return type

[**CourseDtoIReadOnlyListEnvelope**](CourseDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyStudentCoursesCountAsync

> int32 GetMyStudentCoursesCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseDtoCollectionQueryParameters(courseDtoCollectionQueryParameters).Execute()

Get current user's enrolled courses count

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseDtoCollectionQueryParameters := *openapiclient.NewCourseDtoCollectionQueryParameters() // CourseDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyStudentCoursesCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseDtoCollectionQueryParameters(courseDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyStudentCoursesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyStudentCoursesCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyStudentCoursesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyStudentCoursesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseDtoCollectionQueryParameters** | [**CourseDtoCollectionQueryParameters**](CourseDtoCollectionQueryParameters.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyStudentProfilesAsync

> StudentProfileDtoIReadOnlyListEnvelope GetMyStudentProfilesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).StudentProfileDtoCollectionQueryParameters(studentProfileDtoCollectionQueryParameters).Execute()

Get current user's student profiles

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	studentProfileDtoCollectionQueryParameters := *openapiclient.NewStudentProfileDtoCollectionQueryParameters() // StudentProfileDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyStudentProfilesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).StudentProfileDtoCollectionQueryParameters(studentProfileDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyStudentProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyStudentProfilesAsync`: StudentProfileDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyStudentProfilesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyStudentProfilesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **studentProfileDtoCollectionQueryParameters** | [**StudentProfileDtoCollectionQueryParameters**](StudentProfileDtoCollectionQueryParameters.md) |  | 

### Return type

[**StudentProfileDtoIReadOnlyListEnvelope**](StudentProfileDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyStudentProfilesCountAsync

> int32 GetMyStudentProfilesCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).StudentProfileDtoCollectionQueryParameters(studentProfileDtoCollectionQueryParameters).Execute()

Get current user's student profiles count

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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	studentProfileDtoCollectionQueryParameters := *openapiclient.NewStudentProfileDtoCollectionQueryParameters() // StudentProfileDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MeAPI.GetMyStudentProfilesCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).StudentProfileDtoCollectionQueryParameters(studentProfileDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MeAPI.GetMyStudentProfilesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyStudentProfilesCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `MeAPI.GetMyStudentProfilesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyStudentProfilesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **studentProfileDtoCollectionQueryParameters** | [**StudentProfileDtoCollectionQueryParameters**](StudentProfileDtoCollectionQueryParameters.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

