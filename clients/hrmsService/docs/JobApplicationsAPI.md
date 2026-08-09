# \JobApplicationsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeJobApplicationStatusAsync**](JobApplicationsAPI.md#ChangeJobApplicationStatusAsync) | **Post** /api/v2/HrmsService/JobApplications/{jobApplicationId}/Status | Change job application status
[**CreateJobApplicationAsync**](JobApplicationsAPI.md#CreateJobApplicationAsync) | **Post** /api/v2/HrmsService/JobApplications | Create a job application
[**DeleteJobApplicationAsync**](JobApplicationsAPI.md#DeleteJobApplicationAsync) | **Delete** /api/v2/HrmsService/JobApplications/{jobApplicationId} | Delete a job application
[**GetJobApplicationByIdAsync**](JobApplicationsAPI.md#GetJobApplicationByIdAsync) | **Get** /api/v2/HrmsService/JobApplications/{jobApplicationId} | Get job application by ID
[**GetJobApplicationsAsync**](JobApplicationsAPI.md#GetJobApplicationsAsync) | **Get** /api/v2/HrmsService/JobApplications | Get job applications
[**GetJobApplicationsCountAsync**](JobApplicationsAPI.md#GetJobApplicationsCountAsync) | **Get** /api/v2/HrmsService/JobApplications/Count | Count job applications
[**PatchJobApplicationAsync**](JobApplicationsAPI.md#PatchJobApplicationAsync) | **Patch** /api/v2/HrmsService/JobApplications/{jobApplicationId} | Patch a job application
[**UpdateJobApplicationAsync**](JobApplicationsAPI.md#UpdateJobApplicationAsync) | **Put** /api/v2/HrmsService/JobApplications/{jobApplicationId} | Update a job application



## ChangeJobApplicationStatusAsync

> EmptyEnvelope ChangeJobApplicationStatusAsync(ctx, jobApplicationId).TenantId(tenantId).Status(status).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Change job application status



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
	jobApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	status := "status_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicationsAPI.ChangeJobApplicationStatusAsync(context.Background(), jobApplicationId).TenantId(tenantId).Status(status).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicationsAPI.ChangeJobApplicationStatusAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChangeJobApplicationStatusAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicationsAPI.ChangeJobApplicationStatusAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeJobApplicationStatusAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **status** | **string** |  | 
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


## CreateJobApplicationAsync

> EmptyEnvelope CreateJobApplicationAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferApplicationCreateDto(jobOfferApplicationCreateDto).Execute()

Create a job application



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
	jobOfferApplicationCreateDto := *openapiclient.NewJobOfferApplicationCreateDto() // JobOfferApplicationCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicationsAPI.CreateJobApplicationAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferApplicationCreateDto(jobOfferApplicationCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicationsAPI.CreateJobApplicationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobApplicationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicationsAPI.CreateJobApplicationAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobApplicationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferApplicationCreateDto** | [**JobOfferApplicationCreateDto**](JobOfferApplicationCreateDto.md) |  | 

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


## DeleteJobApplicationAsync

> EmptyEnvelope DeleteJobApplicationAsync(ctx, jobApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a job application



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
	jobApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicationsAPI.DeleteJobApplicationAsync(context.Background(), jobApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicationsAPI.DeleteJobApplicationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJobApplicationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicationsAPI.DeleteJobApplicationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobApplicationAsyncRequest struct via the builder pattern


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


## GetJobApplicationByIdAsync

> JobOfferApplicationDtoEnvelope GetJobApplicationByIdAsync(ctx, jobApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get job application by ID



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
	jobApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicationsAPI.GetJobApplicationByIdAsync(context.Background(), jobApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicationsAPI.GetJobApplicationByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobApplicationByIdAsync`: JobOfferApplicationDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicationsAPI.GetJobApplicationByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobApplicationByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JobOfferApplicationDtoEnvelope**](JobOfferApplicationDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobApplicationsAsync

> JobOfferApplicationDtoListEnvelope GetJobApplicationsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferApplicationDtoCollectionQueryParameters(jobOfferApplicationDtoCollectionQueryParameters).Execute()

Get job applications



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
	jobOfferApplicationDtoCollectionQueryParameters := *openapiclient.NewJobOfferApplicationDtoCollectionQueryParameters() // JobOfferApplicationDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicationsAPI.GetJobApplicationsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferApplicationDtoCollectionQueryParameters(jobOfferApplicationDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicationsAPI.GetJobApplicationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobApplicationsAsync`: JobOfferApplicationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicationsAPI.GetJobApplicationsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobApplicationsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferApplicationDtoCollectionQueryParameters** | [**JobOfferApplicationDtoCollectionQueryParameters**](JobOfferApplicationDtoCollectionQueryParameters.md) |  | 

### Return type

[**JobOfferApplicationDtoListEnvelope**](JobOfferApplicationDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobApplicationsCountAsync

> Int32Envelope GetJobApplicationsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferApplicationDtoCollectionQueryParameters(jobOfferApplicationDtoCollectionQueryParameters).Execute()

Count job applications



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
	jobOfferApplicationDtoCollectionQueryParameters := *openapiclient.NewJobOfferApplicationDtoCollectionQueryParameters() // JobOfferApplicationDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicationsAPI.GetJobApplicationsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferApplicationDtoCollectionQueryParameters(jobOfferApplicationDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicationsAPI.GetJobApplicationsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobApplicationsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicationsAPI.GetJobApplicationsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobApplicationsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferApplicationDtoCollectionQueryParameters** | [**JobOfferApplicationDtoCollectionQueryParameters**](JobOfferApplicationDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchJobApplicationAsync

> EmptyEnvelope PatchJobApplicationAsync(ctx, jobApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a job application



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
	jobApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicationsAPI.PatchJobApplicationAsync(context.Background(), jobApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicationsAPI.PatchJobApplicationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchJobApplicationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicationsAPI.PatchJobApplicationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchJobApplicationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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


## UpdateJobApplicationAsync

> EmptyEnvelope UpdateJobApplicationAsync(ctx, jobApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferApplicationUpdateDto(jobOfferApplicationUpdateDto).Execute()

Update a job application



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
	jobApplicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	jobOfferApplicationUpdateDto := *openapiclient.NewJobOfferApplicationUpdateDto() // JobOfferApplicationUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicationsAPI.UpdateJobApplicationAsync(context.Background(), jobApplicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferApplicationUpdateDto(jobOfferApplicationUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicationsAPI.UpdateJobApplicationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJobApplicationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicationsAPI.UpdateJobApplicationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobApplicationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferApplicationUpdateDto** | [**JobOfferApplicationUpdateDto**](JobOfferApplicationUpdateDto.md) |  | 

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

