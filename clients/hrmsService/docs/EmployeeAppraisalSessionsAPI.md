# \EmployeeAppraisalSessionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmployeeAppraisalSessionAsync**](EmployeeAppraisalSessionsAPI.md#CreateEmployeeAppraisalSessionAsync) | **Post** /api/v2/HrmsService/EmployeeAppraisalSessions | Create an employee appraisal session
[**DeleteEmployeeAppraisalSessionAsync**](EmployeeAppraisalSessionsAPI.md#DeleteEmployeeAppraisalSessionAsync) | **Delete** /api/v2/HrmsService/EmployeeAppraisalSessions/{sessionId} | Delete an employee appraisal session
[**GetEmployeeAppraisalSessionByIdAsync**](EmployeeAppraisalSessionsAPI.md#GetEmployeeAppraisalSessionByIdAsync) | **Get** /api/v2/HrmsService/EmployeeAppraisalSessions/{sessionId} | Get employee appraisal session by ID
[**GetEmployeeAppraisalSessionsAsync**](EmployeeAppraisalSessionsAPI.md#GetEmployeeAppraisalSessionsAsync) | **Get** /api/v2/HrmsService/EmployeeAppraisalSessions | Get employee appraisal sessions
[**GetEmployeeAppraisalSessionsCountAsync**](EmployeeAppraisalSessionsAPI.md#GetEmployeeAppraisalSessionsCountAsync) | **Get** /api/v2/HrmsService/EmployeeAppraisalSessions/Count | Count employee appraisal sessions
[**UpdateEmployeeAppraisalSessionAsync**](EmployeeAppraisalSessionsAPI.md#UpdateEmployeeAppraisalSessionAsync) | **Put** /api/v2/HrmsService/EmployeeAppraisalSessions/{sessionId} | Update an employee appraisal session



## CreateEmployeeAppraisalSessionAsync

> EmptyEnvelope CreateEmployeeAppraisalSessionAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).EmployeeAppraisalSessionCreateDto(employeeAppraisalSessionCreateDto).Execute()

Create an employee appraisal session



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
	employeeAppraisalSessionCreateDto := *openapiclient.NewEmployeeAppraisalSessionCreateDto("EmployeeProfileId_example", "AppraisalWorkflowId_example") // EmployeeAppraisalSessionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeAppraisalSessionsAPI.CreateEmployeeAppraisalSessionAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).EmployeeAppraisalSessionCreateDto(employeeAppraisalSessionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAppraisalSessionsAPI.CreateEmployeeAppraisalSessionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmployeeAppraisalSessionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeAppraisalSessionsAPI.CreateEmployeeAppraisalSessionAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmployeeAppraisalSessionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **employeeAppraisalSessionCreateDto** | [**EmployeeAppraisalSessionCreateDto**](EmployeeAppraisalSessionCreateDto.md) |  | 

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


## DeleteEmployeeAppraisalSessionAsync

> EmptyEnvelope DeleteEmployeeAppraisalSessionAsync(ctx, sessionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an employee appraisal session



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeAppraisalSessionsAPI.DeleteEmployeeAppraisalSessionAsync(context.Background(), sessionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAppraisalSessionsAPI.DeleteEmployeeAppraisalSessionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEmployeeAppraisalSessionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeAppraisalSessionsAPI.DeleteEmployeeAppraisalSessionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmployeeAppraisalSessionAsyncRequest struct via the builder pattern


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


## GetEmployeeAppraisalSessionByIdAsync

> EmployeeAppraisalSessionDtoEnvelope GetEmployeeAppraisalSessionByIdAsync(ctx, sessionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get employee appraisal session by ID



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionByIdAsync(context.Background(), sessionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployeeAppraisalSessionByIdAsync`: EmployeeAppraisalSessionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployeeAppraisalSessionByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmployeeAppraisalSessionDtoEnvelope**](EmployeeAppraisalSessionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployeeAppraisalSessionsAsync

> EmployeeAppraisalSessionDtoListEnvelope GetEmployeeAppraisalSessionsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get employee appraisal sessions



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
	resp, r, err := apiClient.EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployeeAppraisalSessionsAsync`: EmployeeAppraisalSessionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployeeAppraisalSessionsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmployeeAppraisalSessionDtoListEnvelope**](EmployeeAppraisalSessionDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployeeAppraisalSessionsCountAsync

> Int32Envelope GetEmployeeAppraisalSessionsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count employee appraisal sessions



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
	resp, r, err := apiClient.EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmployeeAppraisalSessionsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeAppraisalSessionsAPI.GetEmployeeAppraisalSessionsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmployeeAppraisalSessionsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmployeeAppraisalSessionAsync

> EmptyEnvelope UpdateEmployeeAppraisalSessionAsync(ctx, sessionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).EmployeeAppraisalSessionUpdateDto(employeeAppraisalSessionUpdateDto).Execute()

Update an employee appraisal session



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	employeeAppraisalSessionUpdateDto := *openapiclient.NewEmployeeAppraisalSessionUpdateDto() // EmployeeAppraisalSessionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeAppraisalSessionsAPI.UpdateEmployeeAppraisalSessionAsync(context.Background(), sessionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).EmployeeAppraisalSessionUpdateDto(employeeAppraisalSessionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAppraisalSessionsAPI.UpdateEmployeeAppraisalSessionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmployeeAppraisalSessionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeAppraisalSessionsAPI.UpdateEmployeeAppraisalSessionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmployeeAppraisalSessionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **employeeAppraisalSessionUpdateDto** | [**EmployeeAppraisalSessionUpdateDto**](EmployeeAppraisalSessionUpdateDto.md) |  | 

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

