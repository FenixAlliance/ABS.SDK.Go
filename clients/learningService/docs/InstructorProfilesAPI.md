# \InstructorProfilesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2LearningServiceInstructorProfilesCountGet**](InstructorProfilesAPI.md#ApiV2LearningServiceInstructorProfilesCountGet) | **Get** /api/v2/LearningService/InstructorProfiles/Count | 
[**ApiV2LearningServiceInstructorProfilesGet**](InstructorProfilesAPI.md#ApiV2LearningServiceInstructorProfilesGet) | **Get** /api/v2/LearningService/InstructorProfiles | 
[**ApiV2LearningServiceInstructorProfilesInstructorProfileIdDelete**](InstructorProfilesAPI.md#ApiV2LearningServiceInstructorProfilesInstructorProfileIdDelete) | **Delete** /api/v2/LearningService/InstructorProfiles/{instructorProfileId} | 
[**ApiV2LearningServiceInstructorProfilesInstructorProfileIdGet**](InstructorProfilesAPI.md#ApiV2LearningServiceInstructorProfilesInstructorProfileIdGet) | **Get** /api/v2/LearningService/InstructorProfiles/{instructorProfileId} | 
[**ApiV2LearningServiceInstructorProfilesInstructorProfileIdPut**](InstructorProfilesAPI.md#ApiV2LearningServiceInstructorProfilesInstructorProfileIdPut) | **Put** /api/v2/LearningService/InstructorProfiles/{instructorProfileId} | 
[**ApiV2LearningServiceInstructorProfilesPost**](InstructorProfilesAPI.md#ApiV2LearningServiceInstructorProfilesPost) | **Post** /api/v2/LearningService/InstructorProfiles | 



## ApiV2LearningServiceInstructorProfilesCountGet

> int32 ApiV2LearningServiceInstructorProfilesCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2LearningServiceInstructorProfilesCountGet`: int32
	fmt.Fprintf(os.Stdout, "Response from `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceInstructorProfilesCountGetRequest struct via the builder pattern


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


## ApiV2LearningServiceInstructorProfilesGet

> []InstructorProfileDto ApiV2LearningServiceInstructorProfilesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2LearningServiceInstructorProfilesGet`: []InstructorProfileDto
	fmt.Fprintf(os.Stdout, "Response from `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceInstructorProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
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


## ApiV2LearningServiceInstructorProfilesInstructorProfileIdDelete

> ApiV2LearningServiceInstructorProfilesInstructorProfileIdDelete(ctx, instructorProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	instructorProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesInstructorProfileIdDelete(context.Background(), instructorProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesInstructorProfileIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instructorProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceInstructorProfilesInstructorProfileIdDeleteRequest struct via the builder pattern


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


## ApiV2LearningServiceInstructorProfilesInstructorProfileIdGet

> InstructorProfileDto ApiV2LearningServiceInstructorProfilesInstructorProfileIdGet(ctx, instructorProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	instructorProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesInstructorProfileIdGet(context.Background(), instructorProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesInstructorProfileIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2LearningServiceInstructorProfilesInstructorProfileIdGet`: InstructorProfileDto
	fmt.Fprintf(os.Stdout, "Response from `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesInstructorProfileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instructorProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceInstructorProfilesInstructorProfileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InstructorProfileDto**](InstructorProfileDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2LearningServiceInstructorProfilesInstructorProfileIdPut

> ApiV2LearningServiceInstructorProfilesInstructorProfileIdPut(ctx, instructorProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InstructorProfileUpdateDto(instructorProfileUpdateDto).Execute()



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
	instructorProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	instructorProfileUpdateDto := *openapiclient.NewInstructorProfileUpdateDto() // InstructorProfileUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesInstructorProfileIdPut(context.Background(), instructorProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InstructorProfileUpdateDto(instructorProfileUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesInstructorProfileIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instructorProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceInstructorProfilesInstructorProfileIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **instructorProfileUpdateDto** | [**InstructorProfileUpdateDto**](InstructorProfileUpdateDto.md) |  | 

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


## ApiV2LearningServiceInstructorProfilesPost

> ApiV2LearningServiceInstructorProfilesPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InstructorProfileCreateDto(instructorProfileCreateDto).Execute()



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
	instructorProfileCreateDto := *openapiclient.NewInstructorProfileCreateDto() // InstructorProfileCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InstructorProfileCreateDto(instructorProfileCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstructorProfilesAPI.ApiV2LearningServiceInstructorProfilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2LearningServiceInstructorProfilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **instructorProfileCreateDto** | [**InstructorProfileCreateDto**](InstructorProfileCreateDto.md) |  | 

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

