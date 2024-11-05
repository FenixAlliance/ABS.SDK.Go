# \TimeLogApprovalsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPut**](TimeLogApprovalsAPI.md#ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPut) | **Put** /api/v2/TimeTrackerService/TimeLogApprovals/{approvalId}/Approver | 
[**ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPut**](TimeLogApprovalsAPI.md#ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPut) | **Put** /api/v2/TimeTrackerService/TimeLogApprovals/{approvalId}/Status | 
[**ApiV2TimeTrackerServiceTimeLogApprovalsPost**](TimeLogApprovalsAPI.md#ApiV2TimeTrackerServiceTimeLogApprovalsPost) | **Post** /api/v2/TimeTrackerService/TimeLogApprovals | 



## ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPut

> ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPut(ctx, approvalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalApproverUpdateDto(projectHoursApprovalApproverUpdateDto).Execute()



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
	approvalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectHoursApprovalApproverUpdateDto := *openapiclient.NewProjectHoursApprovalApproverUpdateDto() // ProjectHoursApprovalApproverUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPut(context.Background(), approvalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalApproverUpdateDto(projectHoursApprovalApproverUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdApproverPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectHoursApprovalApproverUpdateDto** | [**ProjectHoursApprovalApproverUpdateDto**](ProjectHoursApprovalApproverUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPut

> ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPut(ctx, approvalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalStatusUpdateDto(projectHoursApprovalStatusUpdateDto).Execute()



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
	approvalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	projectHoursApprovalStatusUpdateDto := *openapiclient.NewProjectHoursApprovalStatusUpdateDto() // ProjectHoursApprovalStatusUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPut(context.Background(), approvalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalStatusUpdateDto(projectHoursApprovalStatusUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceTimeLogApprovalsApprovalIdStatusPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectHoursApprovalStatusUpdateDto** | [**ProjectHoursApprovalStatusUpdateDto**](ProjectHoursApprovalStatusUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2TimeTrackerServiceTimeLogApprovalsPost

> ApiV2TimeTrackerServiceTimeLogApprovalsPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalCreateDto(projectHoursApprovalCreateDto).Execute()



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
	projectHoursApprovalCreateDto := *openapiclient.NewProjectHoursApprovalCreateDto() // ProjectHoursApprovalCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalCreateDto(projectHoursApprovalCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogApprovalsAPI.ApiV2TimeTrackerServiceTimeLogApprovalsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2TimeTrackerServiceTimeLogApprovalsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectHoursApprovalCreateDto** | [**ProjectHoursApprovalCreateDto**](ProjectHoursApprovalCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

