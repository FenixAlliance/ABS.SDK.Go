# \RequiredSkillsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequiredSkillAsync**](RequiredSkillsAPI.md#CreateRequiredSkillAsync) | **Post** /api/v2/HrmsService/RequiredSkills | Create a required skill
[**DeleteRequiredSkillAsync**](RequiredSkillsAPI.md#DeleteRequiredSkillAsync) | **Delete** /api/v2/HrmsService/RequiredSkills/{requiredSkillId} | Delete a required skill
[**GetRequiredSkillByIdAsync**](RequiredSkillsAPI.md#GetRequiredSkillByIdAsync) | **Get** /api/v2/HrmsService/RequiredSkills/{requiredSkillId} | Get required skill by ID
[**GetRequiredSkillsAsync**](RequiredSkillsAPI.md#GetRequiredSkillsAsync) | **Get** /api/v2/HrmsService/RequiredSkills | Get required skills
[**GetRequiredSkillsCountAsync**](RequiredSkillsAPI.md#GetRequiredSkillsCountAsync) | **Get** /api/v2/HrmsService/RequiredSkills/Count | Count required skills
[**PatchRequiredSkillAsync**](RequiredSkillsAPI.md#PatchRequiredSkillAsync) | **Patch** /api/v2/HrmsService/RequiredSkills/{requiredSkillId} | Patch a required skill
[**UpdateRequiredSkillAsync**](RequiredSkillsAPI.md#UpdateRequiredSkillAsync) | **Put** /api/v2/HrmsService/RequiredSkills/{requiredSkillId} | Update a required skill



## CreateRequiredSkillAsync

> EmptyEnvelope CreateRequiredSkillAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RequiredSkillRecordCreateDto(requiredSkillRecordCreateDto).Execute()

Create a required skill



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
	requiredSkillRecordCreateDto := *openapiclient.NewRequiredSkillRecordCreateDto("SkillId_example") // RequiredSkillRecordCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequiredSkillsAPI.CreateRequiredSkillAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RequiredSkillRecordCreateDto(requiredSkillRecordCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequiredSkillsAPI.CreateRequiredSkillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequiredSkillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RequiredSkillsAPI.CreateRequiredSkillAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequiredSkillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **requiredSkillRecordCreateDto** | [**RequiredSkillRecordCreateDto**](RequiredSkillRecordCreateDto.md) |  | 

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


## DeleteRequiredSkillAsync

> EmptyEnvelope DeleteRequiredSkillAsync(ctx, requiredSkillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a required skill



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
	requiredSkillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequiredSkillsAPI.DeleteRequiredSkillAsync(context.Background(), requiredSkillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequiredSkillsAPI.DeleteRequiredSkillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRequiredSkillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RequiredSkillsAPI.DeleteRequiredSkillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requiredSkillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequiredSkillAsyncRequest struct via the builder pattern


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


## GetRequiredSkillByIdAsync

> RequiredSkillRecordDtoEnvelope GetRequiredSkillByIdAsync(ctx, requiredSkillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get required skill by ID



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
	requiredSkillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequiredSkillsAPI.GetRequiredSkillByIdAsync(context.Background(), requiredSkillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequiredSkillsAPI.GetRequiredSkillByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequiredSkillByIdAsync`: RequiredSkillRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RequiredSkillsAPI.GetRequiredSkillByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requiredSkillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequiredSkillByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**RequiredSkillRecordDtoEnvelope**](RequiredSkillRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequiredSkillsAsync

> RequiredSkillRecordDtoListEnvelope GetRequiredSkillsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get required skills



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
	resp, r, err := apiClient.RequiredSkillsAPI.GetRequiredSkillsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequiredSkillsAPI.GetRequiredSkillsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequiredSkillsAsync`: RequiredSkillRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RequiredSkillsAPI.GetRequiredSkillsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequiredSkillsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**RequiredSkillRecordDtoListEnvelope**](RequiredSkillRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequiredSkillsCountAsync

> Int32Envelope GetRequiredSkillsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count required skills



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
	resp, r, err := apiClient.RequiredSkillsAPI.GetRequiredSkillsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequiredSkillsAPI.GetRequiredSkillsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequiredSkillsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `RequiredSkillsAPI.GetRequiredSkillsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequiredSkillsCountAsyncRequest struct via the builder pattern


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


## PatchRequiredSkillAsync

> EmptyEnvelope PatchRequiredSkillAsync(ctx, requiredSkillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a required skill



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
	requiredSkillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequiredSkillsAPI.PatchRequiredSkillAsync(context.Background(), requiredSkillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequiredSkillsAPI.PatchRequiredSkillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRequiredSkillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RequiredSkillsAPI.PatchRequiredSkillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requiredSkillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRequiredSkillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## UpdateRequiredSkillAsync

> EmptyEnvelope UpdateRequiredSkillAsync(ctx, requiredSkillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RequiredSkillRecordUpdateDto(requiredSkillRecordUpdateDto).Execute()

Update a required skill



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
	requiredSkillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	requiredSkillRecordUpdateDto := *openapiclient.NewRequiredSkillRecordUpdateDto() // RequiredSkillRecordUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequiredSkillsAPI.UpdateRequiredSkillAsync(context.Background(), requiredSkillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RequiredSkillRecordUpdateDto(requiredSkillRecordUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequiredSkillsAPI.UpdateRequiredSkillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRequiredSkillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RequiredSkillsAPI.UpdateRequiredSkillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requiredSkillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequiredSkillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **requiredSkillRecordUpdateDto** | [**RequiredSkillRecordUpdateDto**](RequiredSkillRecordUpdateDto.md) |  | 

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

