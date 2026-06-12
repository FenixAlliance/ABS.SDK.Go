# \SeawayBillsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSeawayBillLineAsync**](SeawayBillsAPI.md#AddSeawayBillLineAsync) | **Post** /api/v2/LogisticsService/SeawayBills/{billId}/Lines | Add a line to seaway bill
[**CancelSeawayBillAsync**](SeawayBillsAPI.md#CancelSeawayBillAsync) | **Post** /api/v2/LogisticsService/SeawayBills/{billId}/Cancel | Cancel a seaway bill
[**CreateSeawayBillAsync**](SeawayBillsAPI.md#CreateSeawayBillAsync) | **Post** /api/v2/LogisticsService/SeawayBills | Create a seaway bill
[**DeleteSeawayBillAsync**](SeawayBillsAPI.md#DeleteSeawayBillAsync) | **Delete** /api/v2/LogisticsService/SeawayBills/{billId} | Delete a seaway bill
[**GetSeawayBillByIdAsync**](SeawayBillsAPI.md#GetSeawayBillByIdAsync) | **Get** /api/v2/LogisticsService/SeawayBills/{billId} | Get seaway bill by ID
[**GetSeawayBillLinesAsync**](SeawayBillsAPI.md#GetSeawayBillLinesAsync) | **Get** /api/v2/LogisticsService/SeawayBills/{billId}/Lines | Get seaway bill lines
[**GetSeawayBillLinesCountAsync**](SeawayBillsAPI.md#GetSeawayBillLinesCountAsync) | **Get** /api/v2/LogisticsService/SeawayBills/{billId}/Lines/Count | Get seaway bill lines count
[**GetSeawayBillsAsync**](SeawayBillsAPI.md#GetSeawayBillsAsync) | **Get** /api/v2/LogisticsService/SeawayBills | Get all seaway bills
[**GetSeawayBillsCountAsync**](SeawayBillsAPI.md#GetSeawayBillsCountAsync) | **Get** /api/v2/LogisticsService/SeawayBills/Count | Get seaway bills count
[**IssueSeawayBillAsync**](SeawayBillsAPI.md#IssueSeawayBillAsync) | **Post** /api/v2/LogisticsService/SeawayBills/{billId}/Issue | Issue a seaway bill
[**MarkSeawayBillArrivedAsync**](SeawayBillsAPI.md#MarkSeawayBillArrivedAsync) | **Post** /api/v2/LogisticsService/SeawayBills/{billId}/MarkArrived | Mark seaway bill arrived
[**MarkSeawayBillInTransitAsync**](SeawayBillsAPI.md#MarkSeawayBillInTransitAsync) | **Post** /api/v2/LogisticsService/SeawayBills/{billId}/MarkInTransit | Mark seaway bill in transit
[**PatchSeawayBillAsync**](SeawayBillsAPI.md#PatchSeawayBillAsync) | **Patch** /api/v2/LogisticsService/SeawayBills/{billId} | Patch a seaway bill
[**PatchSeawayBillLineAsync**](SeawayBillsAPI.md#PatchSeawayBillLineAsync) | **Patch** /api/v2/LogisticsService/SeawayBills/{billId}/Lines/{lineId} | Patch a seaway bill line
[**ReleaseSeawayBillAsync**](SeawayBillsAPI.md#ReleaseSeawayBillAsync) | **Post** /api/v2/LogisticsService/SeawayBills/{billId}/Release | Release a seaway bill
[**RemoveSeawayBillLineAsync**](SeawayBillsAPI.md#RemoveSeawayBillLineAsync) | **Delete** /api/v2/LogisticsService/SeawayBills/{billId}/Lines/{lineId} | Remove a seaway bill line
[**UpdateSeawayBillAsync**](SeawayBillsAPI.md#UpdateSeawayBillAsync) | **Put** /api/v2/LogisticsService/SeawayBills/{billId} | Update a seaway bill
[**UpdateSeawayBillLineAsync**](SeawayBillsAPI.md#UpdateSeawayBillLineAsync) | **Put** /api/v2/LogisticsService/SeawayBills/{billId}/Lines/{lineId} | Update a seaway bill line



## AddSeawayBillLineAsync

> EmptyEnvelope AddSeawayBillLineAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineCreateDto(waybillLineCreateDto).Execute()

Add a line to seaway bill



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	waybillLineCreateDto := *openapiclient.NewWaybillLineCreateDto() // WaybillLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.AddSeawayBillLineAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineCreateDto(waybillLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.AddSeawayBillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSeawayBillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.AddSeawayBillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSeawayBillLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **waybillLineCreateDto** | [**WaybillLineCreateDto**](WaybillLineCreateDto.md) |  | 

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


## CancelSeawayBillAsync

> EmptyEnvelope CancelSeawayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Cancel a seaway bill



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.CancelSeawayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.CancelSeawayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSeawayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.CancelSeawayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSeawayBillAsyncRequest struct via the builder pattern


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


## CreateSeawayBillAsync

> EmptyEnvelope CreateSeawayBillAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SeawayBillCreateDto(seawayBillCreateDto).Execute()

Create a seaway bill



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
	seawayBillCreateDto := *openapiclient.NewSeawayBillCreateDto() // SeawayBillCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.CreateSeawayBillAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SeawayBillCreateDto(seawayBillCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.CreateSeawayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSeawayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.CreateSeawayBillAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSeawayBillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **seawayBillCreateDto** | [**SeawayBillCreateDto**](SeawayBillCreateDto.md) |  | 

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


## DeleteSeawayBillAsync

> EmptyEnvelope DeleteSeawayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a seaway bill



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.DeleteSeawayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.DeleteSeawayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSeawayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.DeleteSeawayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSeawayBillAsyncRequest struct via the builder pattern


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


## GetSeawayBillByIdAsync

> SeawayBillDtoEnvelope GetSeawayBillByIdAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get seaway bill by ID



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.GetSeawayBillByIdAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.GetSeawayBillByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeawayBillByIdAsync`: SeawayBillDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.GetSeawayBillByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeawayBillByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SeawayBillDtoEnvelope**](SeawayBillDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeawayBillLinesAsync

> WaybillLineDtoListEnvelope GetSeawayBillLinesAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get seaway bill lines



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.GetSeawayBillLinesAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.GetSeawayBillLinesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeawayBillLinesAsync`: WaybillLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.GetSeawayBillLinesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeawayBillLinesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WaybillLineDtoListEnvelope**](WaybillLineDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeawayBillLinesCountAsync

> Int32Envelope GetSeawayBillLinesCountAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get seaway bill lines count



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.GetSeawayBillLinesCountAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.GetSeawayBillLinesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeawayBillLinesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.GetSeawayBillLinesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeawayBillLinesCountAsyncRequest struct via the builder pattern


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


## GetSeawayBillsAsync

> SeawayBillDtoListEnvelope GetSeawayBillsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all seaway bills



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
	resp, r, err := apiClient.SeawayBillsAPI.GetSeawayBillsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.GetSeawayBillsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeawayBillsAsync`: SeawayBillDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.GetSeawayBillsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeawayBillsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SeawayBillDtoListEnvelope**](SeawayBillDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeawayBillsCountAsync

> Int32Envelope GetSeawayBillsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get seaway bills count



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
	resp, r, err := apiClient.SeawayBillsAPI.GetSeawayBillsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.GetSeawayBillsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeawayBillsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.GetSeawayBillsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeawayBillsCountAsyncRequest struct via the builder pattern


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


## IssueSeawayBillAsync

> EmptyEnvelope IssueSeawayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Issue a seaway bill



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.IssueSeawayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.IssueSeawayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueSeawayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.IssueSeawayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueSeawayBillAsyncRequest struct via the builder pattern


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


## MarkSeawayBillArrivedAsync

> EmptyEnvelope MarkSeawayBillArrivedAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Mark seaway bill arrived



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.MarkSeawayBillArrivedAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.MarkSeawayBillArrivedAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkSeawayBillArrivedAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.MarkSeawayBillArrivedAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkSeawayBillArrivedAsyncRequest struct via the builder pattern


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


## MarkSeawayBillInTransitAsync

> EmptyEnvelope MarkSeawayBillInTransitAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Mark seaway bill in transit



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.MarkSeawayBillInTransitAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.MarkSeawayBillInTransitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkSeawayBillInTransitAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.MarkSeawayBillInTransitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkSeawayBillInTransitAsyncRequest struct via the builder pattern


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


## PatchSeawayBillAsync

> EmptyEnvelope PatchSeawayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a seaway bill



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.PatchSeawayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.PatchSeawayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSeawayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.PatchSeawayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSeawayBillAsyncRequest struct via the builder pattern


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


## PatchSeawayBillLineAsync

> EmptyEnvelope PatchSeawayBillLineAsync(ctx, billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a seaway bill line



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.PatchSeawayBillLineAsync(context.Background(), billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.PatchSeawayBillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSeawayBillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.PatchSeawayBillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSeawayBillLineAsyncRequest struct via the builder pattern


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


## ReleaseSeawayBillAsync

> EmptyEnvelope ReleaseSeawayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Release a seaway bill



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.ReleaseSeawayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.ReleaseSeawayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseSeawayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.ReleaseSeawayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseSeawayBillAsyncRequest struct via the builder pattern


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


## RemoveSeawayBillLineAsync

> EmptyEnvelope RemoveSeawayBillLineAsync(ctx, billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove a seaway bill line



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.RemoveSeawayBillLineAsync(context.Background(), billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.RemoveSeawayBillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSeawayBillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.RemoveSeawayBillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSeawayBillLineAsyncRequest struct via the builder pattern


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


## UpdateSeawayBillAsync

> EmptyEnvelope UpdateSeawayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SeawayBillUpdateDto(seawayBillUpdateDto).Execute()

Update a seaway bill



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	seawayBillUpdateDto := *openapiclient.NewSeawayBillUpdateDto() // SeawayBillUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.UpdateSeawayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SeawayBillUpdateDto(seawayBillUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.UpdateSeawayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSeawayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.UpdateSeawayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSeawayBillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **seawayBillUpdateDto** | [**SeawayBillUpdateDto**](SeawayBillUpdateDto.md) |  | 

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


## UpdateSeawayBillLineAsync

> EmptyEnvelope UpdateSeawayBillLineAsync(ctx, billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineUpdateDto(waybillLineUpdateDto).Execute()

Update a seaway bill line



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
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	waybillLineUpdateDto := *openapiclient.NewWaybillLineUpdateDto() // WaybillLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeawayBillsAPI.UpdateSeawayBillLineAsync(context.Background(), billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineUpdateDto(waybillLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeawayBillsAPI.UpdateSeawayBillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSeawayBillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SeawayBillsAPI.UpdateSeawayBillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSeawayBillLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **waybillLineUpdateDto** | [**WaybillLineUpdateDto**](WaybillLineUpdateDto.md) |  | 

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

