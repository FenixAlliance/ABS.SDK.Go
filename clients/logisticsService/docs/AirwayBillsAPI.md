# \AirwayBillsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAirwayBillLineAsync**](AirwayBillsAPI.md#AddAirwayBillLineAsync) | **Post** /api/v2/LogisticsService/AirwayBills/{billId}/Lines | Add a line to airway bill
[**CancelAirwayBillAsync**](AirwayBillsAPI.md#CancelAirwayBillAsync) | **Post** /api/v2/LogisticsService/AirwayBills/{billId}/Cancel | Cancel an airway bill
[**CreateAirwayBillAsync**](AirwayBillsAPI.md#CreateAirwayBillAsync) | **Post** /api/v2/LogisticsService/AirwayBills | Create an airway bill
[**DeleteAirwayBillAsync**](AirwayBillsAPI.md#DeleteAirwayBillAsync) | **Delete** /api/v2/LogisticsService/AirwayBills/{billId} | Delete an airway bill
[**GetAirwayBillByIdAsync**](AirwayBillsAPI.md#GetAirwayBillByIdAsync) | **Get** /api/v2/LogisticsService/AirwayBills/{billId} | Get airway bill by ID
[**GetAirwayBillLinesAsync**](AirwayBillsAPI.md#GetAirwayBillLinesAsync) | **Get** /api/v2/LogisticsService/AirwayBills/{billId}/Lines | Get airway bill lines
[**GetAirwayBillLinesCountAsync**](AirwayBillsAPI.md#GetAirwayBillLinesCountAsync) | **Get** /api/v2/LogisticsService/AirwayBills/{billId}/Lines/Count | Get airway bill lines count
[**GetAirwayBillsAsync**](AirwayBillsAPI.md#GetAirwayBillsAsync) | **Get** /api/v2/LogisticsService/AirwayBills | Get all airway bills
[**GetAirwayBillsCountAsync**](AirwayBillsAPI.md#GetAirwayBillsCountAsync) | **Get** /api/v2/LogisticsService/AirwayBills/Count | Get airway bills count
[**IssueAirwayBillAsync**](AirwayBillsAPI.md#IssueAirwayBillAsync) | **Post** /api/v2/LogisticsService/AirwayBills/{billId}/Issue | Issue an airway bill
[**MarkAirwayBillArrivedAsync**](AirwayBillsAPI.md#MarkAirwayBillArrivedAsync) | **Post** /api/v2/LogisticsService/AirwayBills/{billId}/MarkArrived | Mark airway bill arrived
[**MarkAirwayBillDeliveredAsync**](AirwayBillsAPI.md#MarkAirwayBillDeliveredAsync) | **Post** /api/v2/LogisticsService/AirwayBills/{billId}/MarkDelivered | Mark airway bill delivered
[**MarkAirwayBillInTransitAsync**](AirwayBillsAPI.md#MarkAirwayBillInTransitAsync) | **Post** /api/v2/LogisticsService/AirwayBills/{billId}/MarkInTransit | Mark airway bill in transit
[**PatchAirwayBillAsync**](AirwayBillsAPI.md#PatchAirwayBillAsync) | **Patch** /api/v2/LogisticsService/AirwayBills/{billId} | Patch an airway bill
[**PatchAirwayBillLineAsync**](AirwayBillsAPI.md#PatchAirwayBillLineAsync) | **Patch** /api/v2/LogisticsService/AirwayBills/{billId}/Lines/{lineId} | Patch an airway bill line
[**RemoveAirwayBillLineAsync**](AirwayBillsAPI.md#RemoveAirwayBillLineAsync) | **Delete** /api/v2/LogisticsService/AirwayBills/{billId}/Lines/{lineId} | Remove an airway bill line
[**UpdateAirwayBillAsync**](AirwayBillsAPI.md#UpdateAirwayBillAsync) | **Put** /api/v2/LogisticsService/AirwayBills/{billId} | Update an airway bill
[**UpdateAirwayBillLineAsync**](AirwayBillsAPI.md#UpdateAirwayBillLineAsync) | **Put** /api/v2/LogisticsService/AirwayBills/{billId}/Lines/{lineId} | Update an airway bill line



## AddAirwayBillLineAsync

> EmptyEnvelope AddAirwayBillLineAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineCreateDto(waybillLineCreateDto).Execute()

Add a line to airway bill



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
	resp, r, err := apiClient.AirwayBillsAPI.AddAirwayBillLineAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineCreateDto(waybillLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.AddAirwayBillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAirwayBillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.AddAirwayBillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAirwayBillLineAsyncRequest struct via the builder pattern


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


## CancelAirwayBillAsync

> EmptyEnvelope CancelAirwayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Cancel an airway bill



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
	resp, r, err := apiClient.AirwayBillsAPI.CancelAirwayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.CancelAirwayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelAirwayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.CancelAirwayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelAirwayBillAsyncRequest struct via the builder pattern


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


## CreateAirwayBillAsync

> EmptyEnvelope CreateAirwayBillAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AirwayBillCreateDto(airwayBillCreateDto).Execute()

Create an airway bill



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
	airwayBillCreateDto := *openapiclient.NewAirwayBillCreateDto() // AirwayBillCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AirwayBillsAPI.CreateAirwayBillAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AirwayBillCreateDto(airwayBillCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.CreateAirwayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAirwayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.CreateAirwayBillAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAirwayBillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **airwayBillCreateDto** | [**AirwayBillCreateDto**](AirwayBillCreateDto.md) |  | 

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


## DeleteAirwayBillAsync

> EmptyEnvelope DeleteAirwayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an airway bill



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
	resp, r, err := apiClient.AirwayBillsAPI.DeleteAirwayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.DeleteAirwayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAirwayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.DeleteAirwayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAirwayBillAsyncRequest struct via the builder pattern


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


## GetAirwayBillByIdAsync

> AirwayBillDtoEnvelope GetAirwayBillByIdAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get airway bill by ID



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
	resp, r, err := apiClient.AirwayBillsAPI.GetAirwayBillByIdAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.GetAirwayBillByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAirwayBillByIdAsync`: AirwayBillDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.GetAirwayBillByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAirwayBillByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AirwayBillDtoEnvelope**](AirwayBillDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAirwayBillLinesAsync

> WaybillLineDtoListEnvelope GetAirwayBillLinesAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get airway bill lines



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
	resp, r, err := apiClient.AirwayBillsAPI.GetAirwayBillLinesAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.GetAirwayBillLinesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAirwayBillLinesAsync`: WaybillLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.GetAirwayBillLinesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAirwayBillLinesAsyncRequest struct via the builder pattern


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


## GetAirwayBillLinesCountAsync

> Int32Envelope GetAirwayBillLinesCountAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get airway bill lines count



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
	resp, r, err := apiClient.AirwayBillsAPI.GetAirwayBillLinesCountAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.GetAirwayBillLinesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAirwayBillLinesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.GetAirwayBillLinesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAirwayBillLinesCountAsyncRequest struct via the builder pattern


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


## GetAirwayBillsAsync

> AirwayBillDtoListEnvelope GetAirwayBillsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all airway bills



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
	resp, r, err := apiClient.AirwayBillsAPI.GetAirwayBillsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.GetAirwayBillsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAirwayBillsAsync`: AirwayBillDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.GetAirwayBillsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAirwayBillsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AirwayBillDtoListEnvelope**](AirwayBillDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAirwayBillsCountAsync

> Int32Envelope GetAirwayBillsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get airway bills count



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
	resp, r, err := apiClient.AirwayBillsAPI.GetAirwayBillsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.GetAirwayBillsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAirwayBillsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.GetAirwayBillsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAirwayBillsCountAsyncRequest struct via the builder pattern


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


## IssueAirwayBillAsync

> EmptyEnvelope IssueAirwayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Issue an airway bill



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
	resp, r, err := apiClient.AirwayBillsAPI.IssueAirwayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.IssueAirwayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueAirwayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.IssueAirwayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueAirwayBillAsyncRequest struct via the builder pattern


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


## MarkAirwayBillArrivedAsync

> EmptyEnvelope MarkAirwayBillArrivedAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Mark airway bill arrived



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
	resp, r, err := apiClient.AirwayBillsAPI.MarkAirwayBillArrivedAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.MarkAirwayBillArrivedAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkAirwayBillArrivedAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.MarkAirwayBillArrivedAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkAirwayBillArrivedAsyncRequest struct via the builder pattern


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


## MarkAirwayBillDeliveredAsync

> EmptyEnvelope MarkAirwayBillDeliveredAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Mark airway bill delivered



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
	resp, r, err := apiClient.AirwayBillsAPI.MarkAirwayBillDeliveredAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.MarkAirwayBillDeliveredAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkAirwayBillDeliveredAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.MarkAirwayBillDeliveredAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkAirwayBillDeliveredAsyncRequest struct via the builder pattern


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


## MarkAirwayBillInTransitAsync

> EmptyEnvelope MarkAirwayBillInTransitAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Mark airway bill in transit



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
	resp, r, err := apiClient.AirwayBillsAPI.MarkAirwayBillInTransitAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.MarkAirwayBillInTransitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkAirwayBillInTransitAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.MarkAirwayBillInTransitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkAirwayBillInTransitAsyncRequest struct via the builder pattern


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


## PatchAirwayBillAsync

> EmptyEnvelope PatchAirwayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch an airway bill



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
	resp, r, err := apiClient.AirwayBillsAPI.PatchAirwayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.PatchAirwayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAirwayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.PatchAirwayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAirwayBillAsyncRequest struct via the builder pattern


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


## PatchAirwayBillLineAsync

> EmptyEnvelope PatchAirwayBillLineAsync(ctx, billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch an airway bill line



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
	resp, r, err := apiClient.AirwayBillsAPI.PatchAirwayBillLineAsync(context.Background(), billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.PatchAirwayBillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAirwayBillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.PatchAirwayBillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAirwayBillLineAsyncRequest struct via the builder pattern


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


## RemoveAirwayBillLineAsync

> EmptyEnvelope RemoveAirwayBillLineAsync(ctx, billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove an airway bill line



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
	resp, r, err := apiClient.AirwayBillsAPI.RemoveAirwayBillLineAsync(context.Background(), billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.RemoveAirwayBillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAirwayBillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.RemoveAirwayBillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAirwayBillLineAsyncRequest struct via the builder pattern


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


## UpdateAirwayBillAsync

> EmptyEnvelope UpdateAirwayBillAsync(ctx, billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AirwayBillUpdateDto(airwayBillUpdateDto).Execute()

Update an airway bill



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
	airwayBillUpdateDto := *openapiclient.NewAirwayBillUpdateDto() // AirwayBillUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AirwayBillsAPI.UpdateAirwayBillAsync(context.Background(), billId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AirwayBillUpdateDto(airwayBillUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.UpdateAirwayBillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAirwayBillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.UpdateAirwayBillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAirwayBillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **airwayBillUpdateDto** | [**AirwayBillUpdateDto**](AirwayBillUpdateDto.md) |  | 

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


## UpdateAirwayBillLineAsync

> EmptyEnvelope UpdateAirwayBillLineAsync(ctx, billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineUpdateDto(waybillLineUpdateDto).Execute()

Update an airway bill line



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
	resp, r, err := apiClient.AirwayBillsAPI.UpdateAirwayBillLineAsync(context.Background(), billId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineUpdateDto(waybillLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AirwayBillsAPI.UpdateAirwayBillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAirwayBillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AirwayBillsAPI.UpdateAirwayBillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAirwayBillLineAsyncRequest struct via the builder pattern


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

