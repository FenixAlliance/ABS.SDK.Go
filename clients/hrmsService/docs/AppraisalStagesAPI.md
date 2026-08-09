# \AppraisalStagesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppraisalStageAsync**](AppraisalStagesAPI.md#CreateAppraisalStageAsync) | **Post** /api/v2/HrmsService/AppraisalStages | Create an appraisal stage
[**DeleteAppraisalStageAsync**](AppraisalStagesAPI.md#DeleteAppraisalStageAsync) | **Delete** /api/v2/HrmsService/AppraisalStages/{stageId} | Delete an appraisal stage
[**GetAppraisalStageByIdAsync**](AppraisalStagesAPI.md#GetAppraisalStageByIdAsync) | **Get** /api/v2/HrmsService/AppraisalStages/{stageId} | Get appraisal stage by ID
[**GetAppraisalStagesAsync**](AppraisalStagesAPI.md#GetAppraisalStagesAsync) | **Get** /api/v2/HrmsService/AppraisalStages | Get appraisal stages
[**GetAppraisalStagesCountAsync**](AppraisalStagesAPI.md#GetAppraisalStagesCountAsync) | **Get** /api/v2/HrmsService/AppraisalStages/Count | Count appraisal stages
[**UpdateAppraisalStageAsync**](AppraisalStagesAPI.md#UpdateAppraisalStageAsync) | **Put** /api/v2/HrmsService/AppraisalStages/{stageId} | Update an appraisal stage



## CreateAppraisalStageAsync

> EmptyEnvelope CreateAppraisalStageAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalStageCreateDto(appraisalStageCreateDto).Execute()

Create an appraisal stage



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
	appraisalStageCreateDto := *openapiclient.NewAppraisalStageCreateDto("Name_example", "AppraisalWorkflowId_example", int32(123)) // AppraisalStageCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalStagesAPI.CreateAppraisalStageAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalStageCreateDto(appraisalStageCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalStagesAPI.CreateAppraisalStageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppraisalStageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalStagesAPI.CreateAppraisalStageAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppraisalStageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appraisalStageCreateDto** | [**AppraisalStageCreateDto**](AppraisalStageCreateDto.md) |  | 

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


## DeleteAppraisalStageAsync

> EmptyEnvelope DeleteAppraisalStageAsync(ctx, stageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an appraisal stage



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
	stageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalStagesAPI.DeleteAppraisalStageAsync(context.Background(), stageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalStagesAPI.DeleteAppraisalStageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppraisalStageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalStagesAPI.DeleteAppraisalStageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppraisalStageAsyncRequest struct via the builder pattern


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


## GetAppraisalStageByIdAsync

> AppraisalStageDtoEnvelope GetAppraisalStageByIdAsync(ctx, stageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get appraisal stage by ID



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
	stageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalStagesAPI.GetAppraisalStageByIdAsync(context.Background(), stageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalStagesAPI.GetAppraisalStageByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppraisalStageByIdAsync`: AppraisalStageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalStagesAPI.GetAppraisalStageByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppraisalStageByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AppraisalStageDtoEnvelope**](AppraisalStageDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppraisalStagesAsync

> AppraisalStageDtoListEnvelope GetAppraisalStagesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalStageDtoCollectionQueryParameters(appraisalStageDtoCollectionQueryParameters).Execute()

Get appraisal stages



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
	appraisalStageDtoCollectionQueryParameters := *openapiclient.NewAppraisalStageDtoCollectionQueryParameters() // AppraisalStageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalStagesAPI.GetAppraisalStagesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalStageDtoCollectionQueryParameters(appraisalStageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalStagesAPI.GetAppraisalStagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppraisalStagesAsync`: AppraisalStageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalStagesAPI.GetAppraisalStagesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppraisalStagesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appraisalStageDtoCollectionQueryParameters** | [**AppraisalStageDtoCollectionQueryParameters**](AppraisalStageDtoCollectionQueryParameters.md) |  | 

### Return type

[**AppraisalStageDtoListEnvelope**](AppraisalStageDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppraisalStagesCountAsync

> Int32Envelope GetAppraisalStagesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalStageDtoCollectionQueryParameters(appraisalStageDtoCollectionQueryParameters).Execute()

Count appraisal stages



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
	appraisalStageDtoCollectionQueryParameters := *openapiclient.NewAppraisalStageDtoCollectionQueryParameters() // AppraisalStageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalStagesAPI.GetAppraisalStagesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalStageDtoCollectionQueryParameters(appraisalStageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalStagesAPI.GetAppraisalStagesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppraisalStagesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalStagesAPI.GetAppraisalStagesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppraisalStagesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appraisalStageDtoCollectionQueryParameters** | [**AppraisalStageDtoCollectionQueryParameters**](AppraisalStageDtoCollectionQueryParameters.md) |  | 

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


## UpdateAppraisalStageAsync

> EmptyEnvelope UpdateAppraisalStageAsync(ctx, stageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalStageUpdateDto(appraisalStageUpdateDto).Execute()

Update an appraisal stage



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
	stageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	appraisalStageUpdateDto := *openapiclient.NewAppraisalStageUpdateDto() // AppraisalStageUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppraisalStagesAPI.UpdateAppraisalStageAsync(context.Background(), stageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppraisalStageUpdateDto(appraisalStageUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppraisalStagesAPI.UpdateAppraisalStageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppraisalStageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AppraisalStagesAPI.UpdateAppraisalStageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppraisalStageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appraisalStageUpdateDto** | [**AppraisalStageUpdateDto**](AppraisalStageUpdateDto.md) |  | 

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

