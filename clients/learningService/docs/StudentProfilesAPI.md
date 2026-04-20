# \StudentProfilesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2LearningServiceStudentProfilesCountGet**](StudentProfilesAPI.md#ApiV2LearningServiceStudentProfilesCountGet) | **Get** /api/v2/LearningService/StudentProfiles/Count | 
[**ApiV2LearningServiceStudentProfilesGet**](StudentProfilesAPI.md#ApiV2LearningServiceStudentProfilesGet) | **Get** /api/v2/LearningService/StudentProfiles | 
[**ApiV2LearningServiceStudentProfilesPost**](StudentProfilesAPI.md#ApiV2LearningServiceStudentProfilesPost) | **Post** /api/v2/LearningService/StudentProfiles | 
[**ApiV2LearningServiceStudentProfilesStudentProfileIdAverageGet**](StudentProfilesAPI.md#ApiV2LearningServiceStudentProfilesStudentProfileIdAverageGet) | **Get** /api/v2/LearningService/StudentProfiles/{studentProfileId}/Average | 
[**ApiV2LearningServiceStudentProfilesStudentProfileIdDelete**](StudentProfilesAPI.md#ApiV2LearningServiceStudentProfilesStudentProfileIdDelete) | **Delete** /api/v2/LearningService/StudentProfiles/{studentProfileId} | 
[**ApiV2LearningServiceStudentProfilesStudentProfileIdGet**](StudentProfilesAPI.md#ApiV2LearningServiceStudentProfilesStudentProfileIdGet) | **Get** /api/v2/LearningService/StudentProfiles/{studentProfileId} | 
[**ApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGet**](StudentProfilesAPI.md#ApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGet) | **Get** /api/v2/LearningService/StudentProfiles/{studentProfileId}/HoursCompleted | 
[**ApiV2LearningServiceStudentProfilesStudentProfileIdPut**](StudentProfilesAPI.md#ApiV2LearningServiceStudentProfilesStudentProfileIdPut) | **Put** /api/v2/LearningService/StudentProfiles/{studentProfileId} | 



## ApiV2LearningServiceStudentProfilesCountGet

> int32 ApiV2LearningServiceStudentProfilesCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.StudentProfilesAPI.ApiV2LearningServiceStudentProfilesCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2LearningServiceStudentProfilesCountGet`: int32
	fmt.Fprintf(os.Stdout, "Response from `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceStudentProfilesCountGetRequest struct via the builder pattern


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


## ApiV2LearningServiceStudentProfilesGet

> []StudentProfileDto ApiV2LearningServiceStudentProfilesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.StudentProfilesAPI.ApiV2LearningServiceStudentProfilesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2LearningServiceStudentProfilesGet`: []StudentProfileDto
	fmt.Fprintf(os.Stdout, "Response from `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceStudentProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
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


## ApiV2LearningServiceStudentProfilesPost

> ApiV2LearningServiceStudentProfilesPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).StudentProfileCreateDto(studentProfileCreateDto).Execute()



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
	studentProfileCreateDto := *openapiclient.NewStudentProfileCreateDto() // StudentProfileCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudentProfilesAPI.ApiV2LearningServiceStudentProfilesPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).StudentProfileCreateDto(studentProfileCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceStudentProfilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **studentProfileCreateDto** | [**StudentProfileCreateDto**](StudentProfileCreateDto.md) |  | 

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


## ApiV2LearningServiceStudentProfilesStudentProfileIdAverageGet

> AverageDto ApiV2LearningServiceStudentProfilesStudentProfileIdAverageGet(ctx, studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdAverageGet(context.Background(), studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdAverageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2LearningServiceStudentProfilesStudentProfileIdAverageGet`: AverageDto
	fmt.Fprintf(os.Stdout, "Response from `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdAverageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studentProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceStudentProfilesStudentProfileIdAverageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AverageDto**](AverageDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2LearningServiceStudentProfilesStudentProfileIdDelete

> ApiV2LearningServiceStudentProfilesStudentProfileIdDelete(ctx, studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	r, err := apiClient.StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdDelete(context.Background(), studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studentProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceStudentProfilesStudentProfileIdDeleteRequest struct via the builder pattern


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


## ApiV2LearningServiceStudentProfilesStudentProfileIdGet

> StudentProfileDto ApiV2LearningServiceStudentProfilesStudentProfileIdGet(ctx, studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdGet(context.Background(), studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2LearningServiceStudentProfilesStudentProfileIdGet`: StudentProfileDto
	fmt.Fprintf(os.Stdout, "Response from `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studentProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceStudentProfilesStudentProfileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**StudentProfileDto**](StudentProfileDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGet

> CountDto ApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGet(ctx, studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGet(context.Background(), studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGet`: CountDto
	fmt.Fprintf(os.Stdout, "Response from `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studentProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceStudentProfilesStudentProfileIdHoursCompletedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountDto**](CountDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2LearningServiceStudentProfilesStudentProfileIdPut

> ApiV2LearningServiceStudentProfilesStudentProfileIdPut(ctx, studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).StudentProfileUpdateDto(studentProfileUpdateDto).Execute()



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
	studentProfileUpdateDto := *openapiclient.NewStudentProfileUpdateDto() // StudentProfileUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdPut(context.Background(), studentProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).StudentProfileUpdateDto(studentProfileUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StudentProfilesAPI.ApiV2LearningServiceStudentProfilesStudentProfileIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studentProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceStudentProfilesStudentProfileIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **studentProfileUpdateDto** | [**StudentProfileUpdateDto**](StudentProfileUpdateDto.md) |  | 

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

