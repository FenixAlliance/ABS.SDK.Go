# \RailWaybillsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRailWaybillLineAsync**](RailWaybillsAPI.md#AddRailWaybillLineAsync) | **Post** /api/v2/LogisticsService/RailWaybills/{waybillId}/Lines | Add a line to rail waybill
[**CancelRailWaybillAsync**](RailWaybillsAPI.md#CancelRailWaybillAsync) | **Post** /api/v2/LogisticsService/RailWaybills/{waybillId}/Cancel | Cancel a rail waybill
[**CreateRailWaybillAsync**](RailWaybillsAPI.md#CreateRailWaybillAsync) | **Post** /api/v2/LogisticsService/RailWaybills | Create a rail waybill
[**DeleteRailWaybillAsync**](RailWaybillsAPI.md#DeleteRailWaybillAsync) | **Delete** /api/v2/LogisticsService/RailWaybills/{waybillId} | Delete a rail waybill
[**GetRailWaybillByIdAsync**](RailWaybillsAPI.md#GetRailWaybillByIdAsync) | **Get** /api/v2/LogisticsService/RailWaybills/{waybillId} | Get rail waybill by ID
[**GetRailWaybillLinesAsync**](RailWaybillsAPI.md#GetRailWaybillLinesAsync) | **Get** /api/v2/LogisticsService/RailWaybills/{waybillId}/Lines | Get rail waybill lines
[**GetRailWaybillLinesCountAsync**](RailWaybillsAPI.md#GetRailWaybillLinesCountAsync) | **Get** /api/v2/LogisticsService/RailWaybills/{waybillId}/Lines/Count | Get rail waybill lines count
[**GetRailWaybillsAsync**](RailWaybillsAPI.md#GetRailWaybillsAsync) | **Get** /api/v2/LogisticsService/RailWaybills | Get all rail waybills
[**GetRailWaybillsCountAsync**](RailWaybillsAPI.md#GetRailWaybillsCountAsync) | **Get** /api/v2/LogisticsService/RailWaybills/Count | Get rail waybills count
[**IssueRailWaybillAsync**](RailWaybillsAPI.md#IssueRailWaybillAsync) | **Post** /api/v2/LogisticsService/RailWaybills/{waybillId}/Issue | Issue a rail waybill
[**MarkRailWaybillDeliveredAsync**](RailWaybillsAPI.md#MarkRailWaybillDeliveredAsync) | **Post** /api/v2/LogisticsService/RailWaybills/{waybillId}/MarkDelivered | Mark rail waybill delivered
[**MarkRailWaybillInTransitAsync**](RailWaybillsAPI.md#MarkRailWaybillInTransitAsync) | **Post** /api/v2/LogisticsService/RailWaybills/{waybillId}/MarkInTransit | Mark rail waybill in transit
[**PatchRailWaybillAsync**](RailWaybillsAPI.md#PatchRailWaybillAsync) | **Patch** /api/v2/LogisticsService/RailWaybills/{waybillId} | Patch a rail waybill
[**PatchRailWaybillLineAsync**](RailWaybillsAPI.md#PatchRailWaybillLineAsync) | **Patch** /api/v2/LogisticsService/RailWaybills/{waybillId}/Lines/{lineId} | Patch a rail waybill line
[**RemoveRailWaybillLineAsync**](RailWaybillsAPI.md#RemoveRailWaybillLineAsync) | **Delete** /api/v2/LogisticsService/RailWaybills/{waybillId}/Lines/{lineId} | Remove a rail waybill line
[**UpdateRailWaybillAsync**](RailWaybillsAPI.md#UpdateRailWaybillAsync) | **Put** /api/v2/LogisticsService/RailWaybills/{waybillId} | Update a rail waybill
[**UpdateRailWaybillLineAsync**](RailWaybillsAPI.md#UpdateRailWaybillLineAsync) | **Put** /api/v2/LogisticsService/RailWaybills/{waybillId}/Lines/{lineId} | Update a rail waybill line



## AddRailWaybillLineAsync

> EmptyEnvelope AddRailWaybillLineAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineCreateDto(waybillLineCreateDto).Execute()

Add a line to rail waybill



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	waybillLineCreateDto := *openapiclient.NewWaybillLineCreateDto() // WaybillLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.AddRailWaybillLineAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineCreateDto(waybillLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.AddRailWaybillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRailWaybillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.AddRailWaybillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRailWaybillLineAsyncRequest struct via the builder pattern


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


## CancelRailWaybillAsync

> EmptyEnvelope CancelRailWaybillAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Cancel a rail waybill



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.CancelRailWaybillAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.CancelRailWaybillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelRailWaybillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.CancelRailWaybillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRailWaybillAsyncRequest struct via the builder pattern


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


## CreateRailWaybillAsync

> EmptyEnvelope CreateRailWaybillAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RailWaybillCreateDto(railWaybillCreateDto).Execute()

Create a rail waybill



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
	railWaybillCreateDto := *openapiclient.NewRailWaybillCreateDto() // RailWaybillCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.CreateRailWaybillAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RailWaybillCreateDto(railWaybillCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.CreateRailWaybillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRailWaybillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.CreateRailWaybillAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRailWaybillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **railWaybillCreateDto** | [**RailWaybillCreateDto**](RailWaybillCreateDto.md) |  | 

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


## DeleteRailWaybillAsync

> EmptyEnvelope DeleteRailWaybillAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a rail waybill



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.DeleteRailWaybillAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.DeleteRailWaybillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRailWaybillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.DeleteRailWaybillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRailWaybillAsyncRequest struct via the builder pattern


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


## GetRailWaybillByIdAsync

> RailWaybillDtoEnvelope GetRailWaybillByIdAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get rail waybill by ID



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.GetRailWaybillByIdAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.GetRailWaybillByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRailWaybillByIdAsync`: RailWaybillDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.GetRailWaybillByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRailWaybillByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**RailWaybillDtoEnvelope**](RailWaybillDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRailWaybillLinesAsync

> WaybillLineDtoListEnvelope GetRailWaybillLinesAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineDtoCollectionQueryParameters(waybillLineDtoCollectionQueryParameters).Execute()

Get rail waybill lines



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	waybillLineDtoCollectionQueryParameters := *openapiclient.NewWaybillLineDtoCollectionQueryParameters() // WaybillLineDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.GetRailWaybillLinesAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineDtoCollectionQueryParameters(waybillLineDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.GetRailWaybillLinesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRailWaybillLinesAsync`: WaybillLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.GetRailWaybillLinesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRailWaybillLinesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **waybillLineDtoCollectionQueryParameters** | [**WaybillLineDtoCollectionQueryParameters**](WaybillLineDtoCollectionQueryParameters.md) |  | 

### Return type

[**WaybillLineDtoListEnvelope**](WaybillLineDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRailWaybillLinesCountAsync

> Int32Envelope GetRailWaybillLinesCountAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineDtoCollectionQueryParameters(waybillLineDtoCollectionQueryParameters).Execute()

Get rail waybill lines count



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	waybillLineDtoCollectionQueryParameters := *openapiclient.NewWaybillLineDtoCollectionQueryParameters() // WaybillLineDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.GetRailWaybillLinesCountAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineDtoCollectionQueryParameters(waybillLineDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.GetRailWaybillLinesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRailWaybillLinesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.GetRailWaybillLinesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRailWaybillLinesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **waybillLineDtoCollectionQueryParameters** | [**WaybillLineDtoCollectionQueryParameters**](WaybillLineDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRailWaybillsAsync

> RailWaybillDtoListEnvelope GetRailWaybillsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RailWaybillDtoCollectionQueryParameters(railWaybillDtoCollectionQueryParameters).Execute()

Get all rail waybills



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
	railWaybillDtoCollectionQueryParameters := *openapiclient.NewRailWaybillDtoCollectionQueryParameters() // RailWaybillDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.GetRailWaybillsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RailWaybillDtoCollectionQueryParameters(railWaybillDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.GetRailWaybillsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRailWaybillsAsync`: RailWaybillDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.GetRailWaybillsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRailWaybillsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **railWaybillDtoCollectionQueryParameters** | [**RailWaybillDtoCollectionQueryParameters**](RailWaybillDtoCollectionQueryParameters.md) |  | 

### Return type

[**RailWaybillDtoListEnvelope**](RailWaybillDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRailWaybillsCountAsync

> Int32Envelope GetRailWaybillsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RailWaybillDtoCollectionQueryParameters(railWaybillDtoCollectionQueryParameters).Execute()

Get rail waybills count



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
	railWaybillDtoCollectionQueryParameters := *openapiclient.NewRailWaybillDtoCollectionQueryParameters() // RailWaybillDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.GetRailWaybillsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RailWaybillDtoCollectionQueryParameters(railWaybillDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.GetRailWaybillsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRailWaybillsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.GetRailWaybillsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRailWaybillsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **railWaybillDtoCollectionQueryParameters** | [**RailWaybillDtoCollectionQueryParameters**](RailWaybillDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueRailWaybillAsync

> EmptyEnvelope IssueRailWaybillAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Issue a rail waybill



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.IssueRailWaybillAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.IssueRailWaybillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueRailWaybillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.IssueRailWaybillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueRailWaybillAsyncRequest struct via the builder pattern


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


## MarkRailWaybillDeliveredAsync

> EmptyEnvelope MarkRailWaybillDeliveredAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Mark rail waybill delivered



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.MarkRailWaybillDeliveredAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.MarkRailWaybillDeliveredAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkRailWaybillDeliveredAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.MarkRailWaybillDeliveredAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkRailWaybillDeliveredAsyncRequest struct via the builder pattern


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


## MarkRailWaybillInTransitAsync

> EmptyEnvelope MarkRailWaybillInTransitAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Mark rail waybill in transit



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.MarkRailWaybillInTransitAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.MarkRailWaybillInTransitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkRailWaybillInTransitAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.MarkRailWaybillInTransitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkRailWaybillInTransitAsyncRequest struct via the builder pattern


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


## PatchRailWaybillAsync

> EmptyEnvelope PatchRailWaybillAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a rail waybill



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.PatchRailWaybillAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.PatchRailWaybillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRailWaybillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.PatchRailWaybillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRailWaybillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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


## PatchRailWaybillLineAsync

> EmptyEnvelope PatchRailWaybillLineAsync(ctx, waybillId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a rail waybill line



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.PatchRailWaybillLineAsync(context.Background(), waybillId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.PatchRailWaybillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRailWaybillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.PatchRailWaybillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRailWaybillLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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


## RemoveRailWaybillLineAsync

> EmptyEnvelope RemoveRailWaybillLineAsync(ctx, waybillId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove a rail waybill line



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.RemoveRailWaybillLineAsync(context.Background(), waybillId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.RemoveRailWaybillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveRailWaybillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.RemoveRailWaybillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRailWaybillLineAsyncRequest struct via the builder pattern


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


## UpdateRailWaybillAsync

> EmptyEnvelope UpdateRailWaybillAsync(ctx, waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RailWaybillUpdateDto(railWaybillUpdateDto).Execute()

Update a rail waybill



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	railWaybillUpdateDto := *openapiclient.NewRailWaybillUpdateDto() // RailWaybillUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.UpdateRailWaybillAsync(context.Background(), waybillId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RailWaybillUpdateDto(railWaybillUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.UpdateRailWaybillAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRailWaybillAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.UpdateRailWaybillAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRailWaybillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **railWaybillUpdateDto** | [**RailWaybillUpdateDto**](RailWaybillUpdateDto.md) |  | 

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


## UpdateRailWaybillLineAsync

> EmptyEnvelope UpdateRailWaybillLineAsync(ctx, waybillId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineUpdateDto(waybillLineUpdateDto).Execute()

Update a rail waybill line



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
	waybillId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	waybillLineUpdateDto := *openapiclient.NewWaybillLineUpdateDto() // WaybillLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RailWaybillsAPI.UpdateRailWaybillLineAsync(context.Background(), waybillId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WaybillLineUpdateDto(waybillLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RailWaybillsAPI.UpdateRailWaybillLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRailWaybillLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RailWaybillsAPI.UpdateRailWaybillLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**waybillId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRailWaybillLineAsyncRequest struct via the builder pattern


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

