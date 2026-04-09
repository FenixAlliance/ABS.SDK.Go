# \TimeLogApprovalsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestProjectHoursApprovalAsync**](TimeLogApprovalsAPI.md#RequestProjectHoursApprovalAsync) | **Post** /api/v2/TimeTrackerService/TimeLogApprovals | Request project hours approval
[**UpdateProjectHoursApprovalApproverAsync**](TimeLogApprovalsAPI.md#UpdateProjectHoursApprovalApproverAsync) | **Put** /api/v2/TimeTrackerService/TimeLogApprovals/{approvalId}/Approver | Update approval approver
[**UpdateProjectHoursApprovalStatusAsync**](TimeLogApprovalsAPI.md#UpdateProjectHoursApprovalStatusAsync) | **Put** /api/v2/TimeTrackerService/TimeLogApprovals/{approvalId}/Status | Update approval status



## RequestProjectHoursApprovalAsync

> RequestProjectHoursApprovalAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalCreateDto(projectHoursApprovalCreateDto).Execute()

Request project hours approval



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
	r, err := apiClient.TimeLogApprovalsAPI.RequestProjectHoursApprovalAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalCreateDto(projectHoursApprovalCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogApprovalsAPI.RequestProjectHoursApprovalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestProjectHoursApprovalAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectHoursApprovalCreateDto** | [**ProjectHoursApprovalCreateDto**](ProjectHoursApprovalCreateDto.md) |  | 

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


## UpdateProjectHoursApprovalApproverAsync

> UpdateProjectHoursApprovalApproverAsync(ctx, approvalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalApproverUpdateDto(projectHoursApprovalApproverUpdateDto).Execute()

Update approval approver



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
	r, err := apiClient.TimeLogApprovalsAPI.UpdateProjectHoursApprovalApproverAsync(context.Background(), approvalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalApproverUpdateDto(projectHoursApprovalApproverUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogApprovalsAPI.UpdateProjectHoursApprovalApproverAsync``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateProjectHoursApprovalApproverAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectHoursApprovalApproverUpdateDto** | [**ProjectHoursApprovalApproverUpdateDto**](ProjectHoursApprovalApproverUpdateDto.md) |  | 

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


## UpdateProjectHoursApprovalStatusAsync

> UpdateProjectHoursApprovalStatusAsync(ctx, approvalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalStatusUpdateDto(projectHoursApprovalStatusUpdateDto).Execute()

Update approval status



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
	r, err := apiClient.TimeLogApprovalsAPI.UpdateProjectHoursApprovalStatusAsync(context.Background(), approvalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProjectHoursApprovalStatusUpdateDto(projectHoursApprovalStatusUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeLogApprovalsAPI.UpdateProjectHoursApprovalStatusAsync``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateProjectHoursApprovalStatusAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **projectHoursApprovalStatusUpdateDto** | [**ProjectHoursApprovalStatusUpdateDto**](ProjectHoursApprovalStatusUpdateDto.md) |  | 

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

