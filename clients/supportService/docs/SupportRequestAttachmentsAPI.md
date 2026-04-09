# \SupportRequestAttachmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportRequestAttachmentAsync**](SupportRequestAttachmentsAPI.md#CreateSupportRequestAttachmentAsync) | **Post** /api/v2/SupportService/SupportRequestAttachments | Create a new support request attachment
[**DeleteSupportRequestAttachmentAsync**](SupportRequestAttachmentsAPI.md#DeleteSupportRequestAttachmentAsync) | **Delete** /api/v2/SupportService/SupportRequestAttachments/{supportRequestAttachmentId} | Delete a support request attachment
[**GetSupportRequestAttachmentAsync**](SupportRequestAttachmentsAPI.md#GetSupportRequestAttachmentAsync) | **Get** /api/v2/SupportService/SupportRequestAttachments/{supportRequestAttachmentId} | Retrieve a support request attachment by ID
[**GetSupportRequestAttachmentsAsync**](SupportRequestAttachmentsAPI.md#GetSupportRequestAttachmentsAsync) | **Get** /api/v2/SupportService/SupportRequestAttachments | Retrieve a list of support request attachments
[**GetSupportRequestAttachmentsCountAsync**](SupportRequestAttachmentsAPI.md#GetSupportRequestAttachmentsCountAsync) | **Get** /api/v2/SupportService/SupportRequestAttachments/Count | Get the count of support request attachments
[**UpdateSupportRequestAttachmentAsync**](SupportRequestAttachmentsAPI.md#UpdateSupportRequestAttachmentAsync) | **Put** /api/v2/SupportService/SupportRequestAttachments/{supportRequestAttachmentId} | Update a support request attachment



## CreateSupportRequestAttachmentAsync

> EmptyEnvelope CreateSupportRequestAttachmentAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestAttachmentCreateDto(supportRequestAttachmentCreateDto).Execute()

Create a new support request attachment



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
	supportRequestAttachmentCreateDto := *openapiclient.NewSupportRequestAttachmentCreateDto() // SupportRequestAttachmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.CreateSupportRequestAttachmentAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestAttachmentCreateDto(supportRequestAttachmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.CreateSupportRequestAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportRequestAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.CreateSupportRequestAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportRequestAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportRequestAttachmentCreateDto** | [**SupportRequestAttachmentCreateDto**](SupportRequestAttachmentCreateDto.md) |  | 

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


## DeleteSupportRequestAttachmentAsync

> EmptyEnvelope DeleteSupportRequestAttachmentAsync(ctx, supportRequestAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a support request attachment



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
	supportRequestAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.DeleteSupportRequestAttachmentAsync(context.Background(), supportRequestAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.DeleteSupportRequestAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSupportRequestAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.DeleteSupportRequestAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupportRequestAttachmentAsyncRequest struct via the builder pattern


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


## GetSupportRequestAttachmentAsync

> SupportRequestAttachmentDtoEnvelope GetSupportRequestAttachmentAsync(ctx, supportRequestAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a support request attachment by ID



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
	supportRequestAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.GetSupportRequestAttachmentAsync(context.Background(), supportRequestAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.GetSupportRequestAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestAttachmentAsync`: SupportRequestAttachmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.GetSupportRequestAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportRequestAttachmentDtoEnvelope**](SupportRequestAttachmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportRequestAttachmentsAsync

> SupportRequestAttachmentDtoListEnvelope GetSupportRequestAttachmentsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of support request attachments



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
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.GetSupportRequestAttachmentsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.GetSupportRequestAttachmentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestAttachmentsAsync`: SupportRequestAttachmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.GetSupportRequestAttachmentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestAttachmentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportRequestAttachmentDtoListEnvelope**](SupportRequestAttachmentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportRequestAttachmentsCountAsync

> Int32Envelope GetSupportRequestAttachmentsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of support request attachments



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
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.GetSupportRequestAttachmentsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.GetSupportRequestAttachmentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestAttachmentsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.GetSupportRequestAttachmentsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestAttachmentsCountAsyncRequest struct via the builder pattern


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


## UpdateSupportRequestAttachmentAsync

> EmptyEnvelope UpdateSupportRequestAttachmentAsync(ctx, supportRequestAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestAttachmentUpdateDto(supportRequestAttachmentUpdateDto).Execute()

Update a support request attachment



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
	supportRequestAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	supportRequestAttachmentUpdateDto := *openapiclient.NewSupportRequestAttachmentUpdateDto() // SupportRequestAttachmentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.UpdateSupportRequestAttachmentAsync(context.Background(), supportRequestAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestAttachmentUpdateDto(supportRequestAttachmentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.UpdateSupportRequestAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSupportRequestAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.UpdateSupportRequestAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupportRequestAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportRequestAttachmentUpdateDto** | [**SupportRequestAttachmentUpdateDto**](SupportRequestAttachmentUpdateDto.md) |  | 

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

