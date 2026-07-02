# \CourseTeamMembershipsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCourseTeamMembershipAsync**](CourseTeamMembershipsAPI.md#CreateCourseTeamMembershipAsync) | **Post** /api/v2/LearningService/CourseTeamMemberships | Create a course team membership
[**DeleteCourseTeamMembershipAsync**](CourseTeamMembershipsAPI.md#DeleteCourseTeamMembershipAsync) | **Delete** /api/v2/LearningService/CourseTeamMemberships/{membershipId} | Delete a course team membership
[**GetCourseTeamMembershipByIdAsync**](CourseTeamMembershipsAPI.md#GetCourseTeamMembershipByIdAsync) | **Get** /api/v2/LearningService/CourseTeamMemberships/{membershipId} | Get course team membership by ID
[**GetCourseTeamMembershipsAsync**](CourseTeamMembershipsAPI.md#GetCourseTeamMembershipsAsync) | **Get** /api/v2/LearningService/CourseTeamMemberships | Get all course team memberships
[**GetCourseTeamMembershipsCountAsync**](CourseTeamMembershipsAPI.md#GetCourseTeamMembershipsCountAsync) | **Get** /api/v2/LearningService/CourseTeamMemberships/Count | Get course team memberships count
[**PatchCourseTeamMembershipAsync**](CourseTeamMembershipsAPI.md#PatchCourseTeamMembershipAsync) | **Patch** /api/v2/LearningService/CourseTeamMemberships/{membershipId} | Patch a course team membership
[**UpdateCourseTeamMembershipAsync**](CourseTeamMembershipsAPI.md#UpdateCourseTeamMembershipAsync) | **Put** /api/v2/LearningService/CourseTeamMemberships/{membershipId} | Update a course team membership



## CreateCourseTeamMembershipAsync

> CreateCourseTeamMembershipAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseTeamMembershipCreateDto(courseTeamMembershipCreateDto).Execute()

Create a course team membership



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
	courseTeamMembershipCreateDto := *openapiclient.NewCourseTeamMembershipCreateDto("CourseId_example", "InstructorProfileId_example") // CourseTeamMembershipCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseTeamMembershipsAPI.CreateCourseTeamMembershipAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseTeamMembershipCreateDto(courseTeamMembershipCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseTeamMembershipsAPI.CreateCourseTeamMembershipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCourseTeamMembershipAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseTeamMembershipCreateDto** | [**CourseTeamMembershipCreateDto**](CourseTeamMembershipCreateDto.md) |  | 

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


## DeleteCourseTeamMembershipAsync

> DeleteCourseTeamMembershipAsync(ctx, membershipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a course team membership



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
	membershipId := "membershipId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseTeamMembershipsAPI.DeleteCourseTeamMembershipAsync(context.Background(), membershipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseTeamMembershipsAPI.DeleteCourseTeamMembershipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCourseTeamMembershipAsyncRequest struct via the builder pattern


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


## GetCourseTeamMembershipByIdAsync

> CourseTeamMembershipDto GetCourseTeamMembershipByIdAsync(ctx, membershipId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course team membership by ID



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
	membershipId := "membershipId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseTeamMembershipsAPI.GetCourseTeamMembershipByIdAsync(context.Background(), membershipId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseTeamMembershipsAPI.GetCourseTeamMembershipByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseTeamMembershipByIdAsync`: CourseTeamMembershipDto
	fmt.Fprintf(os.Stdout, "Response from `CourseTeamMembershipsAPI.GetCourseTeamMembershipByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseTeamMembershipByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CourseTeamMembershipDto**](CourseTeamMembershipDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseTeamMembershipsAsync

> []CourseTeamMembershipDto GetCourseTeamMembershipsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all course team memberships



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
	resp, r, err := apiClient.CourseTeamMembershipsAPI.GetCourseTeamMembershipsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseTeamMembershipsAPI.GetCourseTeamMembershipsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseTeamMembershipsAsync`: []CourseTeamMembershipDto
	fmt.Fprintf(os.Stdout, "Response from `CourseTeamMembershipsAPI.GetCourseTeamMembershipsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseTeamMembershipsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]CourseTeamMembershipDto**](CourseTeamMembershipDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCourseTeamMembershipsCountAsync

> int32 GetCourseTeamMembershipsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get course team memberships count



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
	resp, r, err := apiClient.CourseTeamMembershipsAPI.GetCourseTeamMembershipsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseTeamMembershipsAPI.GetCourseTeamMembershipsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseTeamMembershipsCountAsync`: int32
	fmt.Fprintf(os.Stdout, "Response from `CourseTeamMembershipsAPI.GetCourseTeamMembershipsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseTeamMembershipsCountAsyncRequest struct via the builder pattern


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


## PatchCourseTeamMembershipAsync

> EmptyEnvelope PatchCourseTeamMembershipAsync(ctx, membershipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a course team membership



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
	membershipId := "membershipId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseTeamMembershipsAPI.PatchCourseTeamMembershipAsync(context.Background(), membershipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseTeamMembershipsAPI.PatchCourseTeamMembershipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCourseTeamMembershipAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CourseTeamMembershipsAPI.PatchCourseTeamMembershipAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCourseTeamMembershipAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## UpdateCourseTeamMembershipAsync

> UpdateCourseTeamMembershipAsync(ctx, membershipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseTeamMembershipUpdateDto(courseTeamMembershipUpdateDto).Execute()

Update a course team membership



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
	membershipId := "membershipId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	courseTeamMembershipUpdateDto := *openapiclient.NewCourseTeamMembershipUpdateDto() // CourseTeamMembershipUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CourseTeamMembershipsAPI.UpdateCourseTeamMembershipAsync(context.Background(), membershipId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CourseTeamMembershipUpdateDto(courseTeamMembershipUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseTeamMembershipsAPI.UpdateCourseTeamMembershipAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**membershipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCourseTeamMembershipAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **courseTeamMembershipUpdateDto** | [**CourseTeamMembershipUpdateDto**](CourseTeamMembershipUpdateDto.md) |  | 

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

