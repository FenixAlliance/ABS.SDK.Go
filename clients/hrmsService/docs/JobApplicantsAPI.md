# \JobApplicantsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobApplicantAsync**](JobApplicantsAPI.md#CreateJobApplicantAsync) | **Post** /api/v2/HrmsService/JobApplicants | Create a job applicant
[**DeleteJobApplicantAsync**](JobApplicantsAPI.md#DeleteJobApplicantAsync) | **Delete** /api/v2/HrmsService/JobApplicants/{jobApplicantId} | Delete a job applicant
[**GetJobApplicantByIdAsync**](JobApplicantsAPI.md#GetJobApplicantByIdAsync) | **Get** /api/v2/HrmsService/JobApplicants/{jobApplicantId} | Get job applicant by ID
[**GetJobApplicantsAsync**](JobApplicantsAPI.md#GetJobApplicantsAsync) | **Get** /api/v2/HrmsService/JobApplicants | Get job applicants
[**GetJobApplicantsCountAsync**](JobApplicantsAPI.md#GetJobApplicantsCountAsync) | **Get** /api/v2/HrmsService/JobApplicants/Count | Count job applicants
[**PatchJobApplicantAsync**](JobApplicantsAPI.md#PatchJobApplicantAsync) | **Patch** /api/v2/HrmsService/JobApplicants/{jobApplicantId} | Patch a job applicant
[**UpdateJobApplicantAsync**](JobApplicantsAPI.md#UpdateJobApplicantAsync) | **Put** /api/v2/HrmsService/JobApplicants/{jobApplicantId} | Update a job applicant



## CreateJobApplicantAsync

> EmptyEnvelope CreateJobApplicantAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobApplicantProfileCreateDto(jobApplicantProfileCreateDto).Execute()

Create a job applicant



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
	jobApplicantProfileCreateDto := *openapiclient.NewJobApplicantProfileCreateDto() // JobApplicantProfileCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicantsAPI.CreateJobApplicantAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobApplicantProfileCreateDto(jobApplicantProfileCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicantsAPI.CreateJobApplicantAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobApplicantAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicantsAPI.CreateJobApplicantAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobApplicantAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobApplicantProfileCreateDto** | [**JobApplicantProfileCreateDto**](JobApplicantProfileCreateDto.md) |  | 

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


## DeleteJobApplicantAsync

> EmptyEnvelope DeleteJobApplicantAsync(ctx, jobApplicantId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a job applicant



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
	jobApplicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicantsAPI.DeleteJobApplicantAsync(context.Background(), jobApplicantId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicantsAPI.DeleteJobApplicantAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJobApplicantAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicantsAPI.DeleteJobApplicantAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobApplicantAsyncRequest struct via the builder pattern


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


## GetJobApplicantByIdAsync

> JobApplicantProfileDtoEnvelope GetJobApplicantByIdAsync(ctx, jobApplicantId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get job applicant by ID



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
	jobApplicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicantsAPI.GetJobApplicantByIdAsync(context.Background(), jobApplicantId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicantsAPI.GetJobApplicantByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobApplicantByIdAsync`: JobApplicantProfileDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicantsAPI.GetJobApplicantByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobApplicantByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JobApplicantProfileDtoEnvelope**](JobApplicantProfileDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobApplicantsAsync

> JobApplicantProfileDtoListEnvelope GetJobApplicantsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobApplicantProfileDtoCollectionQueryParameters(jobApplicantProfileDtoCollectionQueryParameters).Execute()

Get job applicants



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
	jobApplicantProfileDtoCollectionQueryParameters := *openapiclient.NewJobApplicantProfileDtoCollectionQueryParameters() // JobApplicantProfileDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicantsAPI.GetJobApplicantsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobApplicantProfileDtoCollectionQueryParameters(jobApplicantProfileDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicantsAPI.GetJobApplicantsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobApplicantsAsync`: JobApplicantProfileDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicantsAPI.GetJobApplicantsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobApplicantsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobApplicantProfileDtoCollectionQueryParameters** | [**JobApplicantProfileDtoCollectionQueryParameters**](JobApplicantProfileDtoCollectionQueryParameters.md) |  | 

### Return type

[**JobApplicantProfileDtoListEnvelope**](JobApplicantProfileDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobApplicantsCountAsync

> Int32Envelope GetJobApplicantsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobApplicantProfileDtoCollectionQueryParameters(jobApplicantProfileDtoCollectionQueryParameters).Execute()

Count job applicants



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
	jobApplicantProfileDtoCollectionQueryParameters := *openapiclient.NewJobApplicantProfileDtoCollectionQueryParameters() // JobApplicantProfileDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicantsAPI.GetJobApplicantsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobApplicantProfileDtoCollectionQueryParameters(jobApplicantProfileDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicantsAPI.GetJobApplicantsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobApplicantsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicantsAPI.GetJobApplicantsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobApplicantsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobApplicantProfileDtoCollectionQueryParameters** | [**JobApplicantProfileDtoCollectionQueryParameters**](JobApplicantProfileDtoCollectionQueryParameters.md) |  | 

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


## PatchJobApplicantAsync

> EmptyEnvelope PatchJobApplicantAsync(ctx, jobApplicantId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a job applicant



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
	jobApplicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicantsAPI.PatchJobApplicantAsync(context.Background(), jobApplicantId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicantsAPI.PatchJobApplicantAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchJobApplicantAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicantsAPI.PatchJobApplicantAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchJobApplicantAsyncRequest struct via the builder pattern


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


## UpdateJobApplicantAsync

> EmptyEnvelope UpdateJobApplicantAsync(ctx, jobApplicantId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobApplicantProfileUpdateDto(jobApplicantProfileUpdateDto).Execute()

Update a job applicant



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
	jobApplicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	jobApplicantProfileUpdateDto := *openapiclient.NewJobApplicantProfileUpdateDto() // JobApplicantProfileUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobApplicantsAPI.UpdateJobApplicantAsync(context.Background(), jobApplicantId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobApplicantProfileUpdateDto(jobApplicantProfileUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobApplicantsAPI.UpdateJobApplicantAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJobApplicantAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobApplicantsAPI.UpdateJobApplicantAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobApplicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobApplicantAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobApplicantProfileUpdateDto** | [**JobApplicantProfileUpdateDto**](JobApplicantProfileUpdateDto.md) |  | 

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

