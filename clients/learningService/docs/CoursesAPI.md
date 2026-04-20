# \CoursesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourseAsync**](CoursesAPI.md#CreateCourseAsync) | **Post** /api/v2/LearningService/Courses | Create a new course
[**DeleteCourseAsync**](CoursesAPI.md#DeleteCourseAsync) | **Delete** /api/v2/LearningService/Courses/{courseId} | Delete a course
[**GetCourseArticlesByCourseWikiAsync**](CoursesAPI.md#GetCourseArticlesByCourseWikiAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Articles/{wikiId} | Get course articles by course wiki
[**GetCourseArticlesByCourseWikiCountAsync**](CoursesAPI.md#GetCourseArticlesByCourseWikiCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Articles/{wikiId}/Count | Get course articles by course wiki count
[**GetCourseAssignmentsByCourseAsync**](CoursesAPI.md#GetCourseAssignmentsByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Assignments | Get course assignments by course
[**GetCourseAssignmentsByCourseCountAsync**](CoursesAPI.md#GetCourseAssignmentsByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Assignments/Count | Get course assignments by course count
[**GetCourseByIdAsync**](CoursesAPI.md#GetCourseByIdAsync) | **Get** /api/v2/LearningService/Courses/{courseId} | Get course by ID
[**GetCourseCategoriesByCourseAsync**](CoursesAPI.md#GetCourseCategoriesByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Categories | Get course categories by course
[**GetCourseCategoriesByCourseCountAsync**](CoursesAPI.md#GetCourseCategoriesByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Categories/Count | Get course categories by course count
[**GetCourseCohortsByCourseAsync**](CoursesAPI.md#GetCourseCohortsByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Cohorts | Get course cohorts by course
[**GetCourseCohortsByCourseCountAsync**](CoursesAPI.md#GetCourseCohortsByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Cohorts/Count | Get course cohorts by course count
[**GetCourseEnrollmentsByCourseAsync**](CoursesAPI.md#GetCourseEnrollmentsByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Enrollments | Get enrollments by course
[**GetCourseFilesByCourseAsync**](CoursesAPI.md#GetCourseFilesByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Files | Get course files by course
[**GetCourseFilesByCourseCountAsync**](CoursesAPI.md#GetCourseFilesByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Files/Count | Get course files by course count
[**GetCourseForumsByCourseAsync**](CoursesAPI.md#GetCourseForumsByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Forums | Get course forums by course
[**GetCourseForumsByCourseCountAsync**](CoursesAPI.md#GetCourseForumsByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Forums/Count | Get course forums by course count
[**GetCourseHandoutsByCourseAsync**](CoursesAPI.md#GetCourseHandoutsByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Handouts | Get course handouts by course
[**GetCourseHandoutsByCourseCountAsync**](CoursesAPI.md#GetCourseHandoutsByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Handouts/Count | Get course handouts by course count
[**GetCourseLibrariesByCourseAsync**](CoursesAPI.md#GetCourseLibrariesByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Libraries | Get course libraries by course
[**GetCourseLibrariesByCourseCountAsync**](CoursesAPI.md#GetCourseLibrariesByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Libraries/Count | Get course libraries by course count
[**GetCoursePagesByCourseAsync**](CoursesAPI.md#GetCoursePagesByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Pages | Get course pages by course
[**GetCoursePagesByCourseCountAsync**](CoursesAPI.md#GetCoursePagesByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Pages/Count | Get course pages by course count
[**GetCourseProblemSetsByCourseAsync**](CoursesAPI.md#GetCourseProblemSetsByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/ProblemSets | Get course problem sets by course
[**GetCourseProblemSetsByCourseCountAsync**](CoursesAPI.md#GetCourseProblemSetsByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/ProblemSets/Count | Get course problem sets by course count
[**GetCourseSectionsByCourseAsync**](CoursesAPI.md#GetCourseSectionsByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Sections | Get course sections by course
[**GetCourseSectionsByCourseCountAsync**](CoursesAPI.md#GetCourseSectionsByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Sections/Count | Get course sections by course count
[**GetCourseUnitComponentsByCourseAsync**](CoursesAPI.md#GetCourseUnitComponentsByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/UnitComponents | Get course unit components by course
[**GetCourseUnitComponentsByCourseCountAsync**](CoursesAPI.md#GetCourseUnitComponentsByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/UnitComponents/Count | Get course unit components by course count
[**GetCourseUnitsBySectionAsync**](CoursesAPI.md#GetCourseUnitsBySectionAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Units/{sectionId} | Get course units by section
[**GetCourseUnitsBySectionCountAsync**](CoursesAPI.md#GetCourseUnitsBySectionCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Units/{sectionId}/Count | Get course units by section count
[**GetCourseUpdatesByCourseAsync**](CoursesAPI.md#GetCourseUpdatesByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Updates | Get course updates by course
[**GetCourseUpdatesByCourseCountAsync**](CoursesAPI.md#GetCourseUpdatesByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Updates/Count | Get course updates by course count
[**GetCourseWikisByCourseAsync**](CoursesAPI.md#GetCourseWikisByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Wikis | Get course wikis by course
[**GetCourseWikisByCourseCountAsync**](CoursesAPI.md#GetCourseWikisByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Wikis/Count | Get course wikis by course count
[**GetCoursesAsync**](CoursesAPI.md#GetCoursesAsync) | **Get** /api/v2/LearningService/Courses | Get courses
[**GetCoursesCountAsync**](CoursesAPI.md#GetCoursesCountAsync) | **Get** /api/v2/LearningService/Courses/Count | Get courses count
[**GetInstructorProfilesByCourseAsync**](CoursesAPI.md#GetInstructorProfilesByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Instructors | Get instructor profiles by course
[**GetInstructorProfilesByCourseCountAsync**](CoursesAPI.md#GetInstructorProfilesByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Instructors/Count | Get instructor profiles by course count
[**GetStudentProfilesByCourseAsync**](CoursesAPI.md#GetStudentProfilesByCourseAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Students | Get student profiles by course
[**GetStudentProfilesByCourseCountAsync**](CoursesAPI.md#GetStudentProfilesByCourseCountAsync) | **Get** /api/v2/LearningService/Courses/{courseId}/Students/Count | Get student profiles by course count
[**UpdateCourseAsync**](CoursesAPI.md#UpdateCourseAsync) | **Put** /api/v2/LearningService/Courses/{courseId} | Update a course



## CreateCourseAsync

> CreateCourseAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCreateDto(courseCreateDto).Execute()

Create a new course



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
	courseCreateDto := *openapiclient.NewCourseCreateDto() // CourseCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoursesAPI.CreateCourseAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseCreateDto(courseCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.CreateCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseCreateDto** | [**CourseCreateDto**](CourseCreateDto.md) |  | 

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


## DeleteCourseAsync

> DeleteCourseAsync(ctx, courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a course



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
	r, err := apiClient.CoursesAPI.DeleteCourseAsync(context.Background(), courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.DeleteCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCourseAsyncRequest struct via the builder pattern


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


## GetCourseArticlesByCourseWikiAsync

> []CourseArticleDto GetCourseArticlesByCourseWikiAsync(ctx, courseId, wikiId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course articles by course wiki



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
	courseId := "courseId_example" // string | 
	wikiId := "wikiId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseArticlesByCourseWikiAsync(context.Background(), courseId, wikiId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseArticlesByCourseWikiAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseArticlesByCourseWikiAsync`: []CourseArticleDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseArticlesByCourseWikiAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 
**wikiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseArticlesByCourseWikiAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseArticleDto**](CourseArticleDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseArticlesByCourseWikiCountAsync

> int32 GetCourseArticlesByCourseWikiCountAsync(ctx, courseId, wikiId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course articles by course wiki count



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
	courseId := "courseId_example" // string | 
	wikiId := "wikiId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseArticlesByCourseWikiCountAsync(context.Background(), courseId, wikiId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseArticlesByCourseWikiCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseArticlesByCourseWikiCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseArticlesByCourseWikiCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 
**wikiId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseArticlesByCourseWikiCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetCourseAssignmentsByCourseAsync

> []CourseAssignmentDto GetCourseAssignmentsByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course assignments by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseAssignmentsByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseAssignmentsByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseAssignmentsByCourseAsync`: []CourseAssignmentDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseAssignmentsByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseAssignmentsByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseAssignmentDto**](CourseAssignmentDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseAssignmentsByCourseCountAsync

> int32 GetCourseAssignmentsByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course assignments by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseAssignmentsByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseAssignmentsByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseAssignmentsByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseAssignmentsByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseAssignmentsByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseByIdAsync

> CourseDto GetCourseByIdAsync(ctx, courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course by ID



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
	resp, r, err := apiClient.CoursesAPI.GetCourseByIdAsync(context.Background(), courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseByIdAsync`: CourseDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CourseDto**](CourseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseCategoriesByCourseAsync

> []CourseCategoryDto GetCourseCategoriesByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course categories by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseCategoriesByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseCategoriesByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseCategoriesByCourseAsync`: []CourseCategoryDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseCategoriesByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseCategoriesByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseCategoriesByCourseCountAsync

> int32 GetCourseCategoriesByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course categories by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseCategoriesByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseCategoriesByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseCategoriesByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseCategoriesByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseCategoriesByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseCohortsByCourseAsync

> []CourseCohortDto GetCourseCohortsByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course cohorts by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseCohortsByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseCohortsByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseCohortsByCourseAsync`: []CourseCohortDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseCohortsByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseCohortsByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseCohortDto**](CourseCohortDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseCohortsByCourseCountAsync

> int32 GetCourseCohortsByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course cohorts by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseCohortsByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseCohortsByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseCohortsByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseCohortsByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseCohortsByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseEnrollmentsByCourseAsync

> []CourseEnrollmentDto GetCourseEnrollmentsByCourseAsync(ctx, courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get enrollments by course



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
	resp, r, err := apiClient.CoursesAPI.GetCourseEnrollmentsByCourseAsync(context.Background(), courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseEnrollmentsByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseEnrollmentsByCourseAsync`: []CourseEnrollmentDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseEnrollmentsByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseEnrollmentsByCourseAsyncRequest struct via the builder pattern


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


## GetCourseFilesByCourseAsync

> []CourseFileDto GetCourseFilesByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course files by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseFilesByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseFilesByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseFilesByCourseAsync`: []CourseFileDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseFilesByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseFilesByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseFileDto**](CourseFileDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseFilesByCourseCountAsync

> int32 GetCourseFilesByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course files by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseFilesByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseFilesByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseFilesByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseFilesByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseFilesByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseForumsByCourseAsync

> []CourseForumDto GetCourseForumsByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course forums by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseForumsByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseForumsByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseForumsByCourseAsync`: []CourseForumDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseForumsByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseForumsByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseForumDto**](CourseForumDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseForumsByCourseCountAsync

> int32 GetCourseForumsByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course forums by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseForumsByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseForumsByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseForumsByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseForumsByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseForumsByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseHandoutsByCourseAsync

> []CourseHandoutDto GetCourseHandoutsByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course handouts by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseHandoutsByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseHandoutsByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseHandoutsByCourseAsync`: []CourseHandoutDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseHandoutsByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseHandoutsByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseHandoutDto**](CourseHandoutDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseHandoutsByCourseCountAsync

> int32 GetCourseHandoutsByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course handouts by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseHandoutsByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseHandoutsByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseHandoutsByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseHandoutsByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseHandoutsByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseLibrariesByCourseAsync

> []CourseLibraryDto GetCourseLibrariesByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course libraries by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseLibrariesByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseLibrariesByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseLibrariesByCourseAsync`: []CourseLibraryDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseLibrariesByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseLibrariesByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseLibraryDto**](CourseLibraryDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseLibrariesByCourseCountAsync

> int32 GetCourseLibrariesByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course libraries by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseLibrariesByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseLibrariesByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseLibrariesByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseLibrariesByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseLibrariesByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCoursePagesByCourseAsync

> []CoursePageDto GetCoursePagesByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course pages by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCoursePagesByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCoursePagesByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoursePagesByCourseAsync`: []CoursePageDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCoursePagesByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCoursePagesByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CoursePageDto**](CoursePageDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoursePagesByCourseCountAsync

> int32 GetCoursePagesByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course pages by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCoursePagesByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCoursePagesByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoursePagesByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCoursePagesByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCoursePagesByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseProblemSetsByCourseAsync

> []CourseProblemSetDto GetCourseProblemSetsByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course problem sets by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseProblemSetsByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseProblemSetsByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseProblemSetsByCourseAsync`: []CourseProblemSetDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseProblemSetsByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseProblemSetsByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseProblemSetDto**](CourseProblemSetDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseProblemSetsByCourseCountAsync

> int32 GetCourseProblemSetsByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course problem sets by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseProblemSetsByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseProblemSetsByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseProblemSetsByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseProblemSetsByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseProblemSetsByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseSectionsByCourseAsync

> []CourseSectionDto GetCourseSectionsByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course sections by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseSectionsByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseSectionsByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseSectionsByCourseAsync`: []CourseSectionDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseSectionsByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseSectionsByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseSectionDto**](CourseSectionDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseSectionsByCourseCountAsync

> int32 GetCourseSectionsByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course sections by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseSectionsByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseSectionsByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseSectionsByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseSectionsByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseSectionsByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseUnitComponentsByCourseAsync

> []CourseUnitComponentDto GetCourseUnitComponentsByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course unit components by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseUnitComponentsByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseUnitComponentsByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseUnitComponentsByCourseAsync`: []CourseUnitComponentDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseUnitComponentsByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseUnitComponentsByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseUnitComponentDto**](CourseUnitComponentDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseUnitComponentsByCourseCountAsync

> int32 GetCourseUnitComponentsByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course unit components by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseUnitComponentsByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseUnitComponentsByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseUnitComponentsByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseUnitComponentsByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseUnitComponentsByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseUnitsBySectionAsync

> []CourseUnitDto GetCourseUnitsBySectionAsync(ctx, courseId, sectionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course units by section



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
	courseId := "courseId_example" // string | 
	sectionId := "sectionId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseUnitsBySectionAsync(context.Background(), courseId, sectionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseUnitsBySectionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseUnitsBySectionAsync`: []CourseUnitDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseUnitsBySectionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 
**sectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseUnitsBySectionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseUnitDto**](CourseUnitDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseUnitsBySectionCountAsync

> int32 GetCourseUnitsBySectionCountAsync(ctx, courseId, sectionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course units by section count



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
	courseId := "courseId_example" // string | 
	sectionId := "sectionId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseUnitsBySectionCountAsync(context.Background(), courseId, sectionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseUnitsBySectionCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseUnitsBySectionCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseUnitsBySectionCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 
**sectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseUnitsBySectionCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetCourseUpdatesByCourseAsync

> []CourseNewsDto GetCourseUpdatesByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course updates by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseUpdatesByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseUpdatesByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseUpdatesByCourseAsync`: []CourseNewsDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseUpdatesByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseUpdatesByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseNewsDto**](CourseNewsDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseUpdatesByCourseCountAsync

> int32 GetCourseUpdatesByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course updates by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseUpdatesByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseUpdatesByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseUpdatesByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseUpdatesByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseUpdatesByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCourseWikisByCourseAsync

> []CourseWikiDto GetCourseWikisByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course wikis by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseWikisByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseWikisByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseWikisByCourseAsync`: []CourseWikiDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseWikisByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseWikisByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseWikiDto**](CourseWikiDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseWikisByCourseCountAsync

> int32 GetCourseWikisByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course wikis by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetCourseWikisByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCourseWikisByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseWikisByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCourseWikisByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseWikisByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetCoursesAsync

> []CourseDto GetCoursesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get courses



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
	resp, r, err := apiClient.CoursesAPI.GetCoursesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCoursesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoursesAsync`: []CourseDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCoursesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCoursesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseDto**](CourseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoursesCountAsync

> int32 GetCoursesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get courses count



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
	resp, r, err := apiClient.CoursesAPI.GetCoursesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetCoursesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoursesCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetCoursesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCoursesCountAsyncRequest struct via the builder pattern


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


## GetInstructorProfilesByCourseAsync

> []InstructorProfileDto GetInstructorProfilesByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get instructor profiles by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetInstructorProfilesByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetInstructorProfilesByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstructorProfilesByCourseAsync`: []InstructorProfileDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetInstructorProfilesByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstructorProfilesByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]InstructorProfileDto**](InstructorProfileDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstructorProfilesByCourseCountAsync

> int32 GetInstructorProfilesByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get instructor profiles by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetInstructorProfilesByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetInstructorProfilesByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstructorProfilesByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetInstructorProfilesByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstructorProfilesByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetStudentProfilesByCourseAsync

> []StudentProfileDto GetStudentProfilesByCourseAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get student profiles by course



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetStudentProfilesByCourseAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetStudentProfilesByCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudentProfilesByCourseAsync`: []StudentProfileDto
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetStudentProfilesByCourseAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudentProfilesByCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]StudentProfileDto**](StudentProfileDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudentProfilesByCourseCountAsync

> int32 GetStudentProfilesByCourseCountAsync(ctx, courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get student profiles by course count



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
	courseId := "courseId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoursesAPI.GetStudentProfilesByCourseCountAsync(context.Background(), courseId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.GetStudentProfilesByCourseCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudentProfilesByCourseCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CoursesAPI.GetStudentProfilesByCourseCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudentProfilesByCourseCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## UpdateCourseAsync

> UpdateCourseAsync(ctx, courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseUpdateDto(courseUpdateDto).Execute()

Update a course



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
	courseUpdateDto := *openapiclient.NewCourseUpdateDto() // CourseUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoursesAPI.UpdateCourseAsync(context.Background(), courseId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseUpdateDto(courseUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoursesAPI.UpdateCourseAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCourseAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseUpdateDto** | [**CourseUpdateDto**](CourseUpdateDto.md) |  | 

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

