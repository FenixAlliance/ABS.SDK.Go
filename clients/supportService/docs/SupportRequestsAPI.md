# \SupportRequestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportRequestAsync**](SupportRequestsAPI.md#CreateSupportRequestAsync) | **Post** /api/v2/SupportService/SupportRequests | Create a new support request
[**DeleteSupportRequestAsync**](SupportRequestsAPI.md#DeleteSupportRequestAsync) | **Delete** /api/v2/SupportService/SupportRequests/{supportRequestId} | Delete a support request
[**GetSupportRequestAsync**](SupportRequestsAPI.md#GetSupportRequestAsync) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId} | Retrieve a support request by ID
[**GetSupportRequestAttachmentByRequest**](SupportRequestsAPI.md#GetSupportRequestAttachmentByRequest) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId}/Attachments/{attachmentId} | Retrieve a specific attachment for a support request
[**GetSupportRequestAttachmentsByRequest**](SupportRequestsAPI.md#GetSupportRequestAttachmentsByRequest) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId}/Attachments | Retrieve attachments for a support request
[**GetSupportRequestAttachmentsCountByRequest**](SupportRequestsAPI.md#GetSupportRequestAttachmentsCountByRequest) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId}/Attachments/Count | Get the count of attachments for a support request
[**GetSupportRequestTicketsAsync**](SupportRequestsAPI.md#GetSupportRequestTicketsAsync) | **Get** /api/v2/SupportService/SupportRequests/{supportRequestId}/Tickets | Retrieve tickets for a support request
[**GetSupportRequestsAsync**](SupportRequestsAPI.md#GetSupportRequestsAsync) | **Get** /api/v2/SupportService/SupportRequests | Retrieve a list of support requests
[**GetSupportRequestsCountAsync**](SupportRequestsAPI.md#GetSupportRequestsCountAsync) | **Get** /api/v2/SupportService/SupportRequests/Count | Get the count of support requests
[**RelateSupportRequestToAttachmentAsync**](SupportRequestsAPI.md#RelateSupportRequestToAttachmentAsync) | **Post** /api/v2/SupportService/SupportRequests/{supportRequestId}/Attachments | Add an attachment to a support request
[**UpdateSupportRequestAsync**](SupportRequestsAPI.md#UpdateSupportRequestAsync) | **Put** /api/v2/SupportService/SupportRequests/{supportRequestId} | Update a support request



## CreateSupportRequestAsync

> EmptyEnvelope CreateSupportRequestAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestCreateDto(supportRequestCreateDto).Execute()

Create a new support request



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
	supportRequestCreateDto := *openapiclient.NewSupportRequestCreateDto("Title_example") // SupportRequestCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.CreateSupportRequestAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestCreateDto(supportRequestCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.CreateSupportRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.CreateSupportRequestAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportRequestCreateDto** | [**SupportRequestCreateDto**](SupportRequestCreateDto.md) |  | 

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


## DeleteSupportRequestAsync

> EmptyEnvelope DeleteSupportRequestAsync(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a support request



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.DeleteSupportRequestAsync(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.DeleteSupportRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSupportRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.DeleteSupportRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupportRequestAsyncRequest struct via the builder pattern


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


## GetSupportRequestAsync

> SupportRequestDtoEnvelope GetSupportRequestAsync(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a support request by ID



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.GetSupportRequestAsync(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.GetSupportRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestAsync`: SupportRequestDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.GetSupportRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportRequestDtoEnvelope**](SupportRequestDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportRequestAttachmentByRequest

> SupportRequestAttachmentDtoEnvelope GetSupportRequestAttachmentByRequest(ctx, supportRequestId, attachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a specific attachment for a support request



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	attachmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.GetSupportRequestAttachmentByRequest(context.Background(), supportRequestId, attachmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.GetSupportRequestAttachmentByRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestAttachmentByRequest`: SupportRequestAttachmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.GetSupportRequestAttachmentByRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestAttachmentByRequestRequest struct via the builder pattern


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


## GetSupportRequestAttachmentsByRequest

> SupportRequestAttachmentDtoListEnvelope GetSupportRequestAttachmentsByRequest(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve attachments for a support request



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.GetSupportRequestAttachmentsByRequest(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.GetSupportRequestAttachmentsByRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestAttachmentsByRequest`: SupportRequestAttachmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.GetSupportRequestAttachmentsByRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestAttachmentsByRequestRequest struct via the builder pattern


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


## GetSupportRequestAttachmentsCountByRequest

> Int32Envelope GetSupportRequestAttachmentsCountByRequest(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of attachments for a support request



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.GetSupportRequestAttachmentsCountByRequest(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.GetSupportRequestAttachmentsCountByRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestAttachmentsCountByRequest`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.GetSupportRequestAttachmentsCountByRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestAttachmentsCountByRequestRequest struct via the builder pattern


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


## GetSupportRequestTicketsAsync

> SupportTicketDtoListEnvelope GetSupportRequestTicketsAsync(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve tickets for a support request



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.GetSupportRequestTicketsAsync(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.GetSupportRequestTicketsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestTicketsAsync`: SupportTicketDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.GetSupportRequestTicketsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestTicketsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketDtoListEnvelope**](SupportTicketDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportRequestsAsync

> SupportRequestDtoListEnvelope GetSupportRequestsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of support requests



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
	resp, r, err := apiClient.SupportRequestsAPI.GetSupportRequestsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.GetSupportRequestsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestsAsync`: SupportRequestDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.GetSupportRequestsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportRequestDtoListEnvelope**](SupportRequestDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportRequestsCountAsync

> Int32Envelope GetSupportRequestsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of support requests



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
	resp, r, err := apiClient.SupportRequestsAPI.GetSupportRequestsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.GetSupportRequestsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportRequestsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.GetSupportRequestsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportRequestsCountAsyncRequest struct via the builder pattern


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


## RelateSupportRequestToAttachmentAsync

> EmptyEnvelope RelateSupportRequestToAttachmentAsync(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestAttachmentCreateDto(supportRequestAttachmentCreateDto).Execute()

Add an attachment to a support request



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	supportRequestAttachmentCreateDto := *openapiclient.NewSupportRequestAttachmentCreateDto() // SupportRequestAttachmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.RelateSupportRequestToAttachmentAsync(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestAttachmentCreateDto(supportRequestAttachmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.RelateSupportRequestToAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelateSupportRequestToAttachmentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.RelateSupportRequestToAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelateSupportRequestToAttachmentAsyncRequest struct via the builder pattern


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


## UpdateSupportRequestAsync

> EmptyEnvelope UpdateSupportRequestAsync(ctx, supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestUpdateDto(supportRequestUpdateDto).Execute()

Update a support request



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
	supportRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	supportRequestUpdateDto := *openapiclient.NewSupportRequestUpdateDto() // SupportRequestUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportRequestsAPI.UpdateSupportRequestAsync(context.Background(), supportRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportRequestUpdateDto(supportRequestUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportRequestsAPI.UpdateSupportRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSupportRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportRequestsAPI.UpdateSupportRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupportRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportRequestUpdateDto** | [**SupportRequestUpdateDto**](SupportRequestUpdateDto.md) |  | 

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

