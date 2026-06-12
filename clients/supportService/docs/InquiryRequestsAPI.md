# \InquiryRequestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInquiryRequestAsync**](InquiryRequestsAPI.md#CreateInquiryRequestAsync) | **Post** /api/v2/SupportService/InquiryRequests | Create an inquiry request
[**DeleteInquiryRequestAsync**](InquiryRequestsAPI.md#DeleteInquiryRequestAsync) | **Delete** /api/v2/SupportService/InquiryRequests/{inquiryRequestId} | Delete an inquiry request
[**GetInquiryRequestAsync**](InquiryRequestsAPI.md#GetInquiryRequestAsync) | **Get** /api/v2/SupportService/InquiryRequests/{inquiryRequestId} | Retrieve an inquiry request by ID
[**GetInquiryRequestsAsync**](InquiryRequestsAPI.md#GetInquiryRequestsAsync) | **Get** /api/v2/SupportService/InquiryRequests | Retrieve inquiry requests
[**GetInquiryRequestsCountAsync**](InquiryRequestsAPI.md#GetInquiryRequestsCountAsync) | **Get** /api/v2/SupportService/InquiryRequests/Count | Get inquiry requests count
[**PatchInquiryRequestAsync**](InquiryRequestsAPI.md#PatchInquiryRequestAsync) | **Patch** /api/v2/SupportService/InquiryRequests/{inquiryRequestId} | Patch an inquiry request
[**UpdateInquiryRequestAsync**](InquiryRequestsAPI.md#UpdateInquiryRequestAsync) | **Put** /api/v2/SupportService/InquiryRequests/{inquiryRequestId} | Update an inquiry request



## CreateInquiryRequestAsync

> EmptyEnvelope CreateInquiryRequestAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InquiryRequestCreateDto(inquiryRequestCreateDto).Execute()

Create an inquiry request

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
	inquiryRequestCreateDto := *openapiclient.NewInquiryRequestCreateDto("Name_example", "Email_example", "Message_example") // InquiryRequestCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InquiryRequestsAPI.CreateInquiryRequestAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InquiryRequestCreateDto(inquiryRequestCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InquiryRequestsAPI.CreateInquiryRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInquiryRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InquiryRequestsAPI.CreateInquiryRequestAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInquiryRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inquiryRequestCreateDto** | [**InquiryRequestCreateDto**](InquiryRequestCreateDto.md) |  | 

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


## DeleteInquiryRequestAsync

> EmptyEnvelope DeleteInquiryRequestAsync(ctx, inquiryRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an inquiry request

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
	inquiryRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InquiryRequestsAPI.DeleteInquiryRequestAsync(context.Background(), inquiryRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InquiryRequestsAPI.DeleteInquiryRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInquiryRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InquiryRequestsAPI.DeleteInquiryRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inquiryRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInquiryRequestAsyncRequest struct via the builder pattern


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


## GetInquiryRequestAsync

> InquiryRequestDtoEnvelope GetInquiryRequestAsync(ctx, inquiryRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve an inquiry request by ID

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
	inquiryRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InquiryRequestsAPI.GetInquiryRequestAsync(context.Background(), inquiryRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InquiryRequestsAPI.GetInquiryRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInquiryRequestAsync`: InquiryRequestDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InquiryRequestsAPI.GetInquiryRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inquiryRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInquiryRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InquiryRequestDtoEnvelope**](InquiryRequestDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInquiryRequestsAsync

> InquiryRequestDtoListEnvelope GetInquiryRequestsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve inquiry requests

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
	resp, r, err := apiClient.InquiryRequestsAPI.GetInquiryRequestsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InquiryRequestsAPI.GetInquiryRequestsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInquiryRequestsAsync`: InquiryRequestDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InquiryRequestsAPI.GetInquiryRequestsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInquiryRequestsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InquiryRequestDtoListEnvelope**](InquiryRequestDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInquiryRequestsCountAsync

> Int32Envelope GetInquiryRequestsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get inquiry requests count

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
	resp, r, err := apiClient.InquiryRequestsAPI.GetInquiryRequestsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InquiryRequestsAPI.GetInquiryRequestsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInquiryRequestsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InquiryRequestsAPI.GetInquiryRequestsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInquiryRequestsCountAsyncRequest struct via the builder pattern


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


## PatchInquiryRequestAsync

> EmptyEnvelope PatchInquiryRequestAsync(ctx, inquiryRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch an inquiry request



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
	inquiryRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InquiryRequestsAPI.PatchInquiryRequestAsync(context.Background(), inquiryRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InquiryRequestsAPI.PatchInquiryRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchInquiryRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InquiryRequestsAPI.PatchInquiryRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inquiryRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchInquiryRequestAsyncRequest struct via the builder pattern


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


## UpdateInquiryRequestAsync

> EmptyEnvelope UpdateInquiryRequestAsync(ctx, inquiryRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InquiryRequestUpdateDto(inquiryRequestUpdateDto).Execute()

Update an inquiry request

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
	inquiryRequestId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	inquiryRequestUpdateDto := *openapiclient.NewInquiryRequestUpdateDto() // InquiryRequestUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InquiryRequestsAPI.UpdateInquiryRequestAsync(context.Background(), inquiryRequestId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InquiryRequestUpdateDto(inquiryRequestUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InquiryRequestsAPI.UpdateInquiryRequestAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInquiryRequestAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InquiryRequestsAPI.UpdateInquiryRequestAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inquiryRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInquiryRequestAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inquiryRequestUpdateDto** | [**InquiryRequestUpdateDto**](InquiryRequestUpdateDto.md) |  | 

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

