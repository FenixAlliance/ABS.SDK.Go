# \SupportRequestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SupportServiceSupportRequestsCountGet**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsCountGet) | **Get** /api/v2/SupportService/SupportRequests/Count | 
[**ApiV2SupportServiceSupportRequestsGet**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsGet) | **Get** /api/v2/SupportService/SupportRequests | 
[**ApiV2SupportServiceSupportRequestsPost**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsPost) | **Post** /api/v2/SupportService/SupportRequests | 
[**ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGet**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGet) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId}/Attachments/{attachmentId} | 
[**ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGet**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGet) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId}/Attachments/Count | 
[**ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGet**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGet) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId}/Attachments | 
[**ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPost**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPost) | **Post** /api/v2/SupportService/SupportRequests/{supportRequestId}/Attachments | 
[**ApiV2SupportServiceSupportRequestsSupportRequestIdDelete**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsSupportRequestIdDelete) | **Delete** /api/v2/SupportService/SupportRequests/{supportRequestId} | 
[**ApiV2SupportServiceSupportRequestsSupportRequestIdGet**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsSupportRequestIdGet) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId} | 
[**ApiV2SupportServiceSupportRequestsSupportRequestIdPut**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsSupportRequestIdPut) | **Put** /api/v2/SupportService/SupportRequests/{supportRequestId} | 
[**ApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGet**](SupportRequestsAPI.md#ApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGet) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId}/Tickets | 



## ApiV2SupportServiceSupportRequestsCountGet

> Int32Envelope ApiV2SupportServiceSupportRequestsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsCountGetRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportRequestsGet

> SupportRequestDtoListEnvelope ApiV2SupportServiceSupportRequestsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsGet`: SupportRequestDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportRequestDtoListEnvelope**](SupportRequestDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportRequestsPost

> EmptyEnvelope ApiV2SupportServiceSupportRequestsPost(ctx).SupportRequestCreateDto(supportRequestCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestCreateDto := *openapiclient.NewSupportRequestCreateDto("Title_example") // SupportRequestCreateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsPost(context.Background()).SupportRequestCreateDto(supportRequestCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportRequestCreateDto** | [**SupportRequestCreateDto**](SupportRequestCreateDto.md) |  | 
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


## ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGet

> SupportRequestAttachmentDtoEnvelope ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGet(ctx, supportRequestId, attachmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	attachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGet(context.Background(), supportRequestId, attachmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGet`: SupportRequestAttachmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsAttachmentIdGetRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGet

> Int32Envelope ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGet(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGet(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsCountGetRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGet

> SupportRequestAttachmentDtoListEnvelope ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGet(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGet(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGet`: SupportRequestAttachmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsGetRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPost

> EmptyEnvelope ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPost(ctx, supportRequestId).SupportRequestAttachmentCreateDto(supportRequestAttachmentCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportRequestAttachmentCreateDto := *openapiclient.NewSupportRequestAttachmentCreateDto() // SupportRequestAttachmentCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPost(context.Background(), supportRequestId).SupportRequestAttachmentCreateDto(supportRequestAttachmentCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsSupportRequestIdAttachmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportRequestAttachmentCreateDto** | [**SupportRequestAttachmentCreateDto**](SupportRequestAttachmentCreateDto.md) |  | 
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


## ApiV2SupportServiceSupportRequestsSupportRequestIdDelete

> EmptyEnvelope ApiV2SupportServiceSupportRequestsSupportRequestIdDelete(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdDelete(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsSupportRequestIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsSupportRequestIdDeleteRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportRequestsSupportRequestIdGet

> SupportRequestDtoEnvelope ApiV2SupportServiceSupportRequestsSupportRequestIdGet(ctx, supportRequestId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdGet(context.Background(), supportRequestId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsSupportRequestIdGet`: SupportRequestDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsSupportRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportRequestDtoEnvelope**](SupportRequestDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportRequestsSupportRequestIdPut

> EmptyEnvelope ApiV2SupportServiceSupportRequestsSupportRequestIdPut(ctx, supportRequestId).SupportRequestUpdateDto(supportRequestUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportRequestUpdateDto := *openapiclient.NewSupportRequestUpdateDto() // SupportRequestUpdateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdPut(context.Background(), supportRequestId).SupportRequestUpdateDto(supportRequestUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsSupportRequestIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsSupportRequestIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportRequestUpdateDto** | [**SupportRequestUpdateDto**](SupportRequestUpdateDto.md) |  | 
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


## ApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGet

> SupportTicketDtoListEnvelope ApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGet(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGet(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGet`: SupportTicketDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.ApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportRequestsSupportRequestIdTicketsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketDtoListEnvelope**](SupportTicketDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

