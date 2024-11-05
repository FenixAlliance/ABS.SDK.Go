# \SupportRequestAttachmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SupportServiceSupportRequestAttachmentsCountGet**](SupportRequestAttachmentsAPI.md#ApiV2SupportServiceSupportRequestAttachmentsCountGet) | **Get** /api/v2/SupportService/SupportRequestAttachments/Count | 
[**ApiV2SupportServiceSupportRequestAttachmentsGet**](SupportRequestAttachmentsAPI.md#ApiV2SupportServiceSupportRequestAttachmentsGet) | **Get** /api/v2/SupportService/SupportRequestAttachments | 
[**ApiV2SupportServiceSupportRequestAttachmentsPost**](SupportRequestAttachmentsAPI.md#ApiV2SupportServiceSupportRequestAttachmentsPost) | **Post** /api/v2/SupportService/SupportRequestAttachments | 
[**ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete**](SupportRequestAttachmentsAPI.md#ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete) | **Delete** /api/v2/SupportService/SupportRequestAttachments/{supportRequestAttachmentId} | 
[**ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet**](SupportRequestAttachmentsAPI.md#ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet) | **Get** /api/v2/SupportService/SupportRequestAttachments/{supportRequestAttachmentId} | 
[**ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut**](SupportRequestAttachmentsAPI.md#ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut) | **Put** /api/v2/SupportService/SupportRequestAttachments/{supportRequestAttachmentId} | 



## ApiV2SupportServiceSupportRequestAttachmentsCountGet

> Int32Envelope ApiV2SupportServiceSupportRequestAttachmentsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestAttachmentsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestAttachmentsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportRequestAttachmentsGet

> SupportRequestAttachmentDtoListEnvelope ApiV2SupportServiceSupportRequestAttachmentsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestAttachmentsGet`: SupportRequestAttachmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestAttachmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportRequestAttachmentDtoListEnvelope**](SupportRequestAttachmentDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportRequestAttachmentsPost

> EmptyEnvelope ApiV2SupportServiceSupportRequestAttachmentsPost(ctx).SupportRequestAttachmentCreateDto(supportRequestAttachmentCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestAttachmentCreateDto := *openapiclient.NewSupportRequestAttachmentCreateDto() // SupportRequestAttachmentCreateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsPost(context.Background()).SupportRequestAttachmentCreateDto(supportRequestAttachmentCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestAttachmentsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestAttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportRequestAttachmentCreateDto** | [**SupportRequestAttachmentCreateDto**](SupportRequestAttachmentCreateDto.md) |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete

> EmptyEnvelope ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete(ctx, supportRequestAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete(context.Background(), supportRequestAttachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet

> SupportRequestAttachmentDtoEnvelope ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet(ctx, supportRequestAttachmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet(context.Background(), supportRequestAttachmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet`: SupportRequestAttachmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportRequestAttachmentDtoEnvelope**](SupportRequestAttachmentDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut

> EmptyEnvelope ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut(ctx, supportRequestAttachmentId).SupportRequestAttachmentUpdateDto(supportRequestAttachmentUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestAttachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportRequestAttachmentUpdateDto := *openapiclient.NewSupportRequestAttachmentUpdateDto() // SupportRequestAttachmentUpdateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut(context.Background(), supportRequestAttachmentId).SupportRequestAttachmentUpdateDto(supportRequestAttachmentUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestAttachmentsAPI.ApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestAttachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestAttachmentsSupportRequestAttachmentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportRequestAttachmentUpdateDto** | [**SupportRequestAttachmentUpdateDto**](SupportRequestAttachmentUpdateDto.md) |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

