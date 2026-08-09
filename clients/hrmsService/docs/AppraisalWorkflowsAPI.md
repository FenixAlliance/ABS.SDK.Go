# \AppraisalWorkflowsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppraisalWorkflowAsync**](AppraisalWorkflowsAPI.md#CreateAppraisalWorkflowAsync) | **Post** /api/v2/HrmsService/AppraisalWorkflows | Create an appraisal workflow
[**DeleteAppraisalWorkflowAsync**](AppraisalWorkflowsAPI.md#DeleteAppraisalWorkflowAsync) | **Delete** /api/v2/HrmsService/AppraisalWorkflows/{workflowId} | Delete an appraisal workflow
[**GetAppraisalWorkflowByIdAsync**](AppraisalWorkflowsAPI.md#GetAppraisalWorkflowByIdAsync) | **Get** /api/v2/HrmsService/AppraisalWorkflows/{workflowId} | Get appraisal workflow by ID
[**GetAppraisalWorkflowsAsync**](AppraisalWorkflowsAPI.md#GetAppraisalWorkflowsAsync) | **Get** /api/v2/HrmsService/AppraisalWorkflows | Get appraisal workflows
[**GetAppraisalWorkflowsCountAsync**](AppraisalWorkflowsAPI.md#GetAppraisalWorkflowsCountAsync) | **Get** /api/v2/HrmsService/AppraisalWorkflows/Count | Count appraisal workflows
[**UpdateAppraisalWorkflowAsync**](AppraisalWorkflowsAPI.md#UpdateAppraisalWorkflowAsync) | **Put** /api/v2/HrmsService/AppraisalWorkflows/{workflowId} | Update an appraisal workflow



## CreateAppraisalWorkflowAsync

> EmptyEnvelope CreateAppraisalWorkflowAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalWorkflowCreateDto(appraisalWorkflowCreateDto).Execute()

Create an appraisal workflow



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
	appraisalWorkflowCreateDto := *openapiclient.NewAppraisalWorkflowCreateDto("Name_example") // AppraisalWorkflowCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalWorkflowsAPI.CreateAppraisalWorkflowAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalWorkflowCreateDto(appraisalWorkflowCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalWorkflowsAPI.CreateAppraisalWorkflowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppraisalWorkflowAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalWorkflowsAPI.CreateAppraisalWorkflowAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppraisalWorkflowAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appraisalWorkflowCreateDto** | [**AppraisalWorkflowCreateDto**](AppraisalWorkflowCreateDto.md) |  | 

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


## DeleteAppraisalWorkflowAsync

> EmptyEnvelope DeleteAppraisalWorkflowAsync(ctx, workflowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an appraisal workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalWorkflowsAPI.DeleteAppraisalWorkflowAsync(context.Background(), workflowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalWorkflowsAPI.DeleteAppraisalWorkflowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppraisalWorkflowAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalWorkflowsAPI.DeleteAppraisalWorkflowAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppraisalWorkflowAsyncRequest struct via the builder pattern


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


## GetAppraisalWorkflowByIdAsync

> AppraisalWorkflowDtoEnvelope GetAppraisalWorkflowByIdAsync(ctx, workflowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get appraisal workflow by ID



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalWorkflowsAPI.GetAppraisalWorkflowByIdAsync(context.Background(), workflowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalWorkflowsAPI.GetAppraisalWorkflowByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppraisalWorkflowByIdAsync`: AppraisalWorkflowDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalWorkflowsAPI.GetAppraisalWorkflowByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppraisalWorkflowByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AppraisalWorkflowDtoEnvelope**](AppraisalWorkflowDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppraisalWorkflowsAsync

> AppraisalWorkflowDtoListEnvelope GetAppraisalWorkflowsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalWorkflowDtoCollectionQueryParameters(appraisalWorkflowDtoCollectionQueryParameters).Execute()

Get appraisal workflows



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
	appraisalWorkflowDtoCollectionQueryParameters := *openapiclient.NewAppraisalWorkflowDtoCollectionQueryParameters() // AppraisalWorkflowDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalWorkflowsAPI.GetAppraisalWorkflowsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalWorkflowDtoCollectionQueryParameters(appraisalWorkflowDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalWorkflowsAPI.GetAppraisalWorkflowsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppraisalWorkflowsAsync`: AppraisalWorkflowDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalWorkflowsAPI.GetAppraisalWorkflowsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppraisalWorkflowsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appraisalWorkflowDtoCollectionQueryParameters** | [**AppraisalWorkflowDtoCollectionQueryParameters**](AppraisalWorkflowDtoCollectionQueryParameters.md) |  | 

### Return type

[**AppraisalWorkflowDtoListEnvelope**](AppraisalWorkflowDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppraisalWorkflowsCountAsync

> Int32Envelope GetAppraisalWorkflowsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalWorkflowDtoCollectionQueryParameters(appraisalWorkflowDtoCollectionQueryParameters).Execute()

Count appraisal workflows



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
	appraisalWorkflowDtoCollectionQueryParameters := *openapiclient.NewAppraisalWorkflowDtoCollectionQueryParameters() // AppraisalWorkflowDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalWorkflowsAPI.GetAppraisalWorkflowsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalWorkflowDtoCollectionQueryParameters(appraisalWorkflowDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalWorkflowsAPI.GetAppraisalWorkflowsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppraisalWorkflowsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalWorkflowsAPI.GetAppraisalWorkflowsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppraisalWorkflowsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appraisalWorkflowDtoCollectionQueryParameters** | [**AppraisalWorkflowDtoCollectionQueryParameters**](AppraisalWorkflowDtoCollectionQueryParameters.md) |  | 

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


## UpdateAppraisalWorkflowAsync

> EmptyEnvelope UpdateAppraisalWorkflowAsync(ctx, workflowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalWorkflowUpdateDto(appraisalWorkflowUpdateDto).Execute()

Update an appraisal workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	appraisalWorkflowUpdateDto := *openapiclient.NewAppraisalWorkflowUpdateDto() // AppraisalWorkflowUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalWorkflowsAPI.UpdateAppraisalWorkflowAsync(context.Background(), workflowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalWorkflowUpdateDto(appraisalWorkflowUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalWorkflowsAPI.UpdateAppraisalWorkflowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppraisalWorkflowAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalWorkflowsAPI.UpdateAppraisalWorkflowAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppraisalWorkflowAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appraisalWorkflowUpdateDto** | [**AppraisalWorkflowUpdateDto**](AppraisalWorkflowUpdateDto.md) |  | 

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

