# \BillsOfLadingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillOfLadingAsync**](BillsOfLadingAPI.md#CreateBillOfLadingAsync) | **Post** /api/v2/ShipmentsService/BillsOfLading | Create a bill of lading
[**CreateBillOfLadingLineAsync**](BillsOfLadingAPI.md#CreateBillOfLadingLineAsync) | **Post** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId}/Lines | Create a bill of lading line
[**DeleteBillOfLadingAsync**](BillsOfLadingAPI.md#DeleteBillOfLadingAsync) | **Delete** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId} | Delete a bill of lading
[**DeleteBillOfLadingLineAsync**](BillsOfLadingAPI.md#DeleteBillOfLadingLineAsync) | **Delete** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId}/Lines/{lineId} | Delete a bill of lading line
[**GetBillOfLadingByIdAsync**](BillsOfLadingAPI.md#GetBillOfLadingByIdAsync) | **Get** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId} | Get bill of lading by ID
[**GetBillOfLadingLineByIdAsync**](BillsOfLadingAPI.md#GetBillOfLadingLineByIdAsync) | **Get** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId}/Lines/{lineId} | Get bill of lading line by ID
[**GetBillOfLadingLinesAsync**](BillsOfLadingAPI.md#GetBillOfLadingLinesAsync) | **Get** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId}/Lines | Get bill of lading lines
[**GetBillOfLadingLinesCountAsync**](BillsOfLadingAPI.md#GetBillOfLadingLinesCountAsync) | **Get** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId}/Lines/Count | Get bill of lading lines count
[**GetBillsOfLadingAsync**](BillsOfLadingAPI.md#GetBillsOfLadingAsync) | **Get** /api/v2/ShipmentsService/BillsOfLading | Get all bills of lading
[**GetBillsOfLadingCountAsync**](BillsOfLadingAPI.md#GetBillsOfLadingCountAsync) | **Get** /api/v2/ShipmentsService/BillsOfLading/Count | Get bills of lading count
[**PatchBillOfLadingAsync**](BillsOfLadingAPI.md#PatchBillOfLadingAsync) | **Patch** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId} | Patch a bill of lading
[**PatchBillOfLadingLineAsync**](BillsOfLadingAPI.md#PatchBillOfLadingLineAsync) | **Patch** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId}/Lines/{lineId} | Patch a bill of lading line
[**UpdateBillOfLadingAsync**](BillsOfLadingAPI.md#UpdateBillOfLadingAsync) | **Put** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId} | Update a bill of lading
[**UpdateBillOfLadingLineAsync**](BillsOfLadingAPI.md#UpdateBillOfLadingLineAsync) | **Put** /api/v2/ShipmentsService/BillsOfLading/{billOfLadingId}/Lines/{lineId} | Update a bill of lading line



## CreateBillOfLadingAsync

> EmptyEnvelope CreateBillOfLadingAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BillOfLadingCreateDto(billOfLadingCreateDto).Execute()

Create a bill of lading



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
	billOfLadingCreateDto := *openapiclient.NewBillOfLadingCreateDto() // BillOfLadingCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.CreateBillOfLadingAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BillOfLadingCreateDto(billOfLadingCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.CreateBillOfLadingAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillOfLadingAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.CreateBillOfLadingAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillOfLadingAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **billOfLadingCreateDto** | [**BillOfLadingCreateDto**](BillOfLadingCreateDto.md) |  | 

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


## CreateBillOfLadingLineAsync

> EmptyEnvelope CreateBillOfLadingLineAsync(ctx, billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BillOfLadingLineCreateDto(billOfLadingLineCreateDto).Execute()

Create a bill of lading line



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	billOfLadingLineCreateDto := *openapiclient.NewBillOfLadingLineCreateDto() // BillOfLadingLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.CreateBillOfLadingLineAsync(context.Background(), billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BillOfLadingLineCreateDto(billOfLadingLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.CreateBillOfLadingLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillOfLadingLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.CreateBillOfLadingLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillOfLadingLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **billOfLadingLineCreateDto** | [**BillOfLadingLineCreateDto**](BillOfLadingLineCreateDto.md) |  | 

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


## DeleteBillOfLadingAsync

> EmptyEnvelope DeleteBillOfLadingAsync(ctx, billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a bill of lading



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.DeleteBillOfLadingAsync(context.Background(), billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.DeleteBillOfLadingAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBillOfLadingAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.DeleteBillOfLadingAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillOfLadingAsyncRequest struct via the builder pattern


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


## DeleteBillOfLadingLineAsync

> EmptyEnvelope DeleteBillOfLadingLineAsync(ctx, billOfLadingId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a bill of lading line



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.DeleteBillOfLadingLineAsync(context.Background(), billOfLadingId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.DeleteBillOfLadingLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBillOfLadingLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.DeleteBillOfLadingLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillOfLadingLineAsyncRequest struct via the builder pattern


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


## GetBillOfLadingByIdAsync

> BillOfLadingDtoEnvelope GetBillOfLadingByIdAsync(ctx, billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get bill of lading by ID



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.GetBillOfLadingByIdAsync(context.Background(), billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.GetBillOfLadingByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillOfLadingByIdAsync`: BillOfLadingDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.GetBillOfLadingByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillOfLadingByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BillOfLadingDtoEnvelope**](BillOfLadingDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillOfLadingLineByIdAsync

> BillOfLadingLineDtoEnvelope GetBillOfLadingLineByIdAsync(ctx, billOfLadingId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get bill of lading line by ID



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.GetBillOfLadingLineByIdAsync(context.Background(), billOfLadingId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.GetBillOfLadingLineByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillOfLadingLineByIdAsync`: BillOfLadingLineDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.GetBillOfLadingLineByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillOfLadingLineByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BillOfLadingLineDtoEnvelope**](BillOfLadingLineDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillOfLadingLinesAsync

> BillOfLadingLineDtoListEnvelope GetBillOfLadingLinesAsync(ctx, billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get bill of lading lines



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.GetBillOfLadingLinesAsync(context.Background(), billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.GetBillOfLadingLinesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillOfLadingLinesAsync`: BillOfLadingLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.GetBillOfLadingLinesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillOfLadingLinesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BillOfLadingLineDtoListEnvelope**](BillOfLadingLineDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillOfLadingLinesCountAsync

> Int32Envelope GetBillOfLadingLinesCountAsync(ctx, billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get bill of lading lines count



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.GetBillOfLadingLinesCountAsync(context.Background(), billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.GetBillOfLadingLinesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillOfLadingLinesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.GetBillOfLadingLinesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillOfLadingLinesCountAsyncRequest struct via the builder pattern


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


## GetBillsOfLadingAsync

> BillOfLadingDtoListEnvelope GetBillsOfLadingAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all bills of lading



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
	resp, r, err := apiClient.BillsOfLadingAPI.GetBillsOfLadingAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.GetBillsOfLadingAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillsOfLadingAsync`: BillOfLadingDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.GetBillsOfLadingAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillsOfLadingAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BillOfLadingDtoListEnvelope**](BillOfLadingDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillsOfLadingCountAsync

> Int32Envelope GetBillsOfLadingCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get bills of lading count



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
	resp, r, err := apiClient.BillsOfLadingAPI.GetBillsOfLadingCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.GetBillsOfLadingCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillsOfLadingCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.GetBillsOfLadingCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillsOfLadingCountAsyncRequest struct via the builder pattern


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


## PatchBillOfLadingAsync

> EmptyEnvelope PatchBillOfLadingAsync(ctx, billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a bill of lading



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.PatchBillOfLadingAsync(context.Background(), billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.PatchBillOfLadingAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchBillOfLadingAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.PatchBillOfLadingAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBillOfLadingAsyncRequest struct via the builder pattern


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


## PatchBillOfLadingLineAsync

> EmptyEnvelope PatchBillOfLadingLineAsync(ctx, billOfLadingId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a bill of lading line



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.PatchBillOfLadingLineAsync(context.Background(), billOfLadingId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.PatchBillOfLadingLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchBillOfLadingLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.PatchBillOfLadingLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBillOfLadingLineAsyncRequest struct via the builder pattern


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


## UpdateBillOfLadingAsync

> EmptyEnvelope UpdateBillOfLadingAsync(ctx, billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BillOfLadingUpdateDto(billOfLadingUpdateDto).Execute()

Update a bill of lading



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	billOfLadingUpdateDto := *openapiclient.NewBillOfLadingUpdateDto() // BillOfLadingUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.UpdateBillOfLadingAsync(context.Background(), billOfLadingId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BillOfLadingUpdateDto(billOfLadingUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.UpdateBillOfLadingAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBillOfLadingAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.UpdateBillOfLadingAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillOfLadingAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **billOfLadingUpdateDto** | [**BillOfLadingUpdateDto**](BillOfLadingUpdateDto.md) |  | 

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


## UpdateBillOfLadingLineAsync

> EmptyEnvelope UpdateBillOfLadingLineAsync(ctx, billOfLadingId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BillOfLadingLineUpdateDto(billOfLadingLineUpdateDto).Execute()

Update a bill of lading line



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
	billOfLadingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	billOfLadingLineUpdateDto := *openapiclient.NewBillOfLadingLineUpdateDto() // BillOfLadingLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillsOfLadingAPI.UpdateBillOfLadingLineAsync(context.Background(), billOfLadingId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BillOfLadingLineUpdateDto(billOfLadingLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillsOfLadingAPI.UpdateBillOfLadingLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBillOfLadingLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillsOfLadingAPI.UpdateBillOfLadingLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillOfLadingLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **billOfLadingLineUpdateDto** | [**BillOfLadingLineUpdateDto**](BillOfLadingLineUpdateDto.md) |  | 

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

