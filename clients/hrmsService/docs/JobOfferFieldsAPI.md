# \JobOfferFieldsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobOfferFieldAsync**](JobOfferFieldsAPI.md#CreateJobOfferFieldAsync) | **Post** /api/v2/HrmsService/JobOfferFields | Create a job offer field
[**DeleteJobOfferFieldAsync**](JobOfferFieldsAPI.md#DeleteJobOfferFieldAsync) | **Delete** /api/v2/HrmsService/JobOfferFields/{jobOfferFieldId} | Delete a job offer field
[**GetJobOfferFieldByIdAsync**](JobOfferFieldsAPI.md#GetJobOfferFieldByIdAsync) | **Get** /api/v2/HrmsService/JobOfferFields/{jobOfferFieldId} | Get job offer field by ID
[**GetJobOfferFieldsAsync**](JobOfferFieldsAPI.md#GetJobOfferFieldsAsync) | **Get** /api/v2/HrmsService/JobOfferFields | Get job offer fields
[**GetJobOfferFieldsCountAsync**](JobOfferFieldsAPI.md#GetJobOfferFieldsCountAsync) | **Get** /api/v2/HrmsService/JobOfferFields/Count | Count job offer fields
[**PatchJobOfferFieldAsync**](JobOfferFieldsAPI.md#PatchJobOfferFieldAsync) | **Patch** /api/v2/HrmsService/JobOfferFields/{jobOfferFieldId} | Patch a job offer field
[**UpdateJobOfferFieldAsync**](JobOfferFieldsAPI.md#UpdateJobOfferFieldAsync) | **Put** /api/v2/HrmsService/JobOfferFields/{jobOfferFieldId} | Update a job offer field



## CreateJobOfferFieldAsync

> EmptyEnvelope CreateJobOfferFieldAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferFieldRecordCreateDto(jobOfferFieldRecordCreateDto).Execute()

Create a job offer field



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
	jobOfferFieldRecordCreateDto := *openapiclient.NewJobOfferFieldRecordCreateDto("JobFieldId_example") // JobOfferFieldRecordCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOfferFieldsAPI.CreateJobOfferFieldAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferFieldRecordCreateDto(jobOfferFieldRecordCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOfferFieldsAPI.CreateJobOfferFieldAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobOfferFieldAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOfferFieldsAPI.CreateJobOfferFieldAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobOfferFieldAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferFieldRecordCreateDto** | [**JobOfferFieldRecordCreateDto**](JobOfferFieldRecordCreateDto.md) |  | 

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


## DeleteJobOfferFieldAsync

> EmptyEnvelope DeleteJobOfferFieldAsync(ctx, jobOfferFieldId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a job offer field



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
	jobOfferFieldId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOfferFieldsAPI.DeleteJobOfferFieldAsync(context.Background(), jobOfferFieldId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOfferFieldsAPI.DeleteJobOfferFieldAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJobOfferFieldAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOfferFieldsAPI.DeleteJobOfferFieldAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferFieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobOfferFieldAsyncRequest struct via the builder pattern


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


## GetJobOfferFieldByIdAsync

> JobOfferFieldRecordDtoEnvelope GetJobOfferFieldByIdAsync(ctx, jobOfferFieldId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get job offer field by ID



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
	jobOfferFieldId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOfferFieldsAPI.GetJobOfferFieldByIdAsync(context.Background(), jobOfferFieldId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOfferFieldsAPI.GetJobOfferFieldByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobOfferFieldByIdAsync`: JobOfferFieldRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOfferFieldsAPI.GetJobOfferFieldByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferFieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobOfferFieldByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JobOfferFieldRecordDtoEnvelope**](JobOfferFieldRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobOfferFieldsAsync

> JobOfferFieldRecordDtoListEnvelope GetJobOfferFieldsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferFieldRecordDtoCollectionQueryParameters(jobOfferFieldRecordDtoCollectionQueryParameters).Execute()

Get job offer fields



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
	jobOfferFieldRecordDtoCollectionQueryParameters := *openapiclient.NewJobOfferFieldRecordDtoCollectionQueryParameters() // JobOfferFieldRecordDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOfferFieldsAPI.GetJobOfferFieldsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferFieldRecordDtoCollectionQueryParameters(jobOfferFieldRecordDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOfferFieldsAPI.GetJobOfferFieldsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobOfferFieldsAsync`: JobOfferFieldRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOfferFieldsAPI.GetJobOfferFieldsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobOfferFieldsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferFieldRecordDtoCollectionQueryParameters** | [**JobOfferFieldRecordDtoCollectionQueryParameters**](JobOfferFieldRecordDtoCollectionQueryParameters.md) |  | 

### Return type

[**JobOfferFieldRecordDtoListEnvelope**](JobOfferFieldRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobOfferFieldsCountAsync

> Int32Envelope GetJobOfferFieldsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferFieldRecordDtoCollectionQueryParameters(jobOfferFieldRecordDtoCollectionQueryParameters).Execute()

Count job offer fields



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
	jobOfferFieldRecordDtoCollectionQueryParameters := *openapiclient.NewJobOfferFieldRecordDtoCollectionQueryParameters() // JobOfferFieldRecordDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOfferFieldsAPI.GetJobOfferFieldsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferFieldRecordDtoCollectionQueryParameters(jobOfferFieldRecordDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOfferFieldsAPI.GetJobOfferFieldsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobOfferFieldsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `JobOfferFieldsAPI.GetJobOfferFieldsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobOfferFieldsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferFieldRecordDtoCollectionQueryParameters** | [**JobOfferFieldRecordDtoCollectionQueryParameters**](JobOfferFieldRecordDtoCollectionQueryParameters.md) |  | 

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


## PatchJobOfferFieldAsync

> EmptyEnvelope PatchJobOfferFieldAsync(ctx, jobOfferFieldId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a job offer field



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
	jobOfferFieldId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOfferFieldsAPI.PatchJobOfferFieldAsync(context.Background(), jobOfferFieldId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOfferFieldsAPI.PatchJobOfferFieldAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchJobOfferFieldAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOfferFieldsAPI.PatchJobOfferFieldAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferFieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchJobOfferFieldAsyncRequest struct via the builder pattern


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


## UpdateJobOfferFieldAsync

> EmptyEnvelope UpdateJobOfferFieldAsync(ctx, jobOfferFieldId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferFieldRecordUpdateDto(jobOfferFieldRecordUpdateDto).Execute()

Update a job offer field



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
	jobOfferFieldId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	jobOfferFieldRecordUpdateDto := *openapiclient.NewJobOfferFieldRecordUpdateDto() // JobOfferFieldRecordUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOfferFieldsAPI.UpdateJobOfferFieldAsync(context.Background(), jobOfferFieldId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferFieldRecordUpdateDto(jobOfferFieldRecordUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOfferFieldsAPI.UpdateJobOfferFieldAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJobOfferFieldAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOfferFieldsAPI.UpdateJobOfferFieldAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferFieldId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobOfferFieldAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferFieldRecordUpdateDto** | [**JobOfferFieldRecordUpdateDto**](JobOfferFieldRecordUpdateDto.md) |  | 

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

