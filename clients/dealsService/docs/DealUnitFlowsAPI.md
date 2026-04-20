# \DealUnitFlowsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDealUnitFlowAsync**](DealUnitFlowsAPI.md#CreateDealUnitFlowAsync) | **Post** /api/v2/DealsService/DealUnitFlows | Create a deal unit flow
[**CreateDealUnitFlowStageAsync**](DealUnitFlowsAPI.md#CreateDealUnitFlowStageAsync) | **Post** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages | Create a deal unit flow stage
[**DeleteDealUnitFlowAsync**](DealUnitFlowsAPI.md#DeleteDealUnitFlowAsync) | **Delete** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId} | Delete a deal unit flow
[**DeleteDealUnitFlowStageAsync**](DealUnitFlowsAPI.md#DeleteDealUnitFlowStageAsync) | **Delete** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages/{dealUnitFlowStageId} | Delete a deal unit flow stage
[**GetDealUnitFlowAsync**](DealUnitFlowsAPI.md#GetDealUnitFlowAsync) | **Get** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId} | Get deal unit flow by ID
[**GetDealUnitFlowStageAsync**](DealUnitFlowsAPI.md#GetDealUnitFlowStageAsync) | **Get** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages/{dealUnitFlowStageId} | Get a deal unit flow stage by ID
[**GetDealUnitFlowStagesAsync**](DealUnitFlowsAPI.md#GetDealUnitFlowStagesAsync) | **Get** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages | Get stages for a deal unit flow
[**GetDealUnitFlowStagesCountAsync**](DealUnitFlowsAPI.md#GetDealUnitFlowStagesCountAsync) | **Get** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages/Count | Get stages count for a deal unit flow
[**GetDealUnitFlowsAsync**](DealUnitFlowsAPI.md#GetDealUnitFlowsAsync) | **Get** /api/v2/DealsService/DealUnitFlows | Get deal unit flows
[**GetDealUnitFlowsCountAsync**](DealUnitFlowsAPI.md#GetDealUnitFlowsCountAsync) | **Get** /api/v2/DealsService/DealUnitFlows/Count | Get deal unit flows count
[**UpdateDealUnitFlowAsync**](DealUnitFlowsAPI.md#UpdateDealUnitFlowAsync) | **Put** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId} | Update a deal unit flow
[**UpdateDealUnitFlowStageAsync**](DealUnitFlowsAPI.md#UpdateDealUnitFlowStageAsync) | **Put** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages/{dealUnitFlowStageId} | Update a deal unit flow stage



## CreateDealUnitFlowAsync

> EmptyEnvelope CreateDealUnitFlowAsync(ctx).TenantId(tenantId).DealUnitFlowCreateDto(dealUnitFlowCreateDto).Execute()

Create a deal unit flow



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
	dealUnitFlowCreateDto := *openapiclient.NewDealUnitFlowCreateDto() // DealUnitFlowCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.CreateDealUnitFlowAsync(context.Background()).TenantId(tenantId).DealUnitFlowCreateDto(dealUnitFlowCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.CreateDealUnitFlowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDealUnitFlowAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.CreateDealUnitFlowAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDealUnitFlowAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **dealUnitFlowCreateDto** | [**DealUnitFlowCreateDto**](DealUnitFlowCreateDto.md) |  | 

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


## CreateDealUnitFlowStageAsync

> EmptyEnvelope CreateDealUnitFlowStageAsync(ctx, dealUnitFlowId).TenantId(tenantId).DealUnitFlowStageCreateDto(dealUnitFlowStageCreateDto).Execute()

Create a deal unit flow stage



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitFlowStageCreateDto := *openapiclient.NewDealUnitFlowStageCreateDto() // DealUnitFlowStageCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.CreateDealUnitFlowStageAsync(context.Background(), dealUnitFlowId).TenantId(tenantId).DealUnitFlowStageCreateDto(dealUnitFlowStageCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.CreateDealUnitFlowStageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDealUnitFlowStageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.CreateDealUnitFlowStageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDealUnitFlowStageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **dealUnitFlowStageCreateDto** | [**DealUnitFlowStageCreateDto**](DealUnitFlowStageCreateDto.md) |  | 

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


## DeleteDealUnitFlowAsync

> EmptyEnvelope DeleteDealUnitFlowAsync(ctx, dealUnitFlowId).TenantId(tenantId).Execute()

Delete a deal unit flow



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.DeleteDealUnitFlowAsync(context.Background(), dealUnitFlowId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.DeleteDealUnitFlowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDealUnitFlowAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.DeleteDealUnitFlowAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDealUnitFlowAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## DeleteDealUnitFlowStageAsync

> EmptyEnvelope DeleteDealUnitFlowStageAsync(ctx, dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).Execute()

Delete a deal unit flow stage



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitFlowStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.DeleteDealUnitFlowStageAsync(context.Background(), dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.DeleteDealUnitFlowStageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDealUnitFlowStageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.DeleteDealUnitFlowStageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 
**dealUnitFlowStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDealUnitFlowStageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## GetDealUnitFlowAsync

> DealUnitFlowDtoEnvelope GetDealUnitFlowAsync(ctx, dealUnitFlowId).TenantId(tenantId).Execute()

Get deal unit flow by ID



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.GetDealUnitFlowAsync(context.Background(), dealUnitFlowId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.GetDealUnitFlowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitFlowAsync`: DealUnitFlowDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.GetDealUnitFlowAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitFlowAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**DealUnitFlowDtoEnvelope**](DealUnitFlowDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDealUnitFlowStageAsync

> DealUnitFlowStageDtoEnvelope GetDealUnitFlowStageAsync(ctx, dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).Execute()

Get a deal unit flow stage by ID



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitFlowStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.GetDealUnitFlowStageAsync(context.Background(), dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.GetDealUnitFlowStageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitFlowStageAsync`: DealUnitFlowStageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.GetDealUnitFlowStageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 
**dealUnitFlowStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitFlowStageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**DealUnitFlowStageDtoEnvelope**](DealUnitFlowStageDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDealUnitFlowStagesAsync

> DealUnitFlowStageDtoListEnvelope GetDealUnitFlowStagesAsync(ctx, dealUnitFlowId).TenantId(tenantId).Execute()

Get stages for a deal unit flow



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.GetDealUnitFlowStagesAsync(context.Background(), dealUnitFlowId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.GetDealUnitFlowStagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitFlowStagesAsync`: DealUnitFlowStageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.GetDealUnitFlowStagesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitFlowStagesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**DealUnitFlowStageDtoListEnvelope**](DealUnitFlowStageDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDealUnitFlowStagesCountAsync

> Int32Envelope GetDealUnitFlowStagesCountAsync(ctx, dealUnitFlowId).TenantId(tenantId).Execute()

Get stages count for a deal unit flow



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.GetDealUnitFlowStagesCountAsync(context.Background(), dealUnitFlowId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.GetDealUnitFlowStagesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitFlowStagesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.GetDealUnitFlowStagesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitFlowStagesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## GetDealUnitFlowsAsync

> DealUnitFlowDtoListEnvelope GetDealUnitFlowsAsync(ctx).TenantId(tenantId).Execute()

Get deal unit flows



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.GetDealUnitFlowsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.GetDealUnitFlowsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitFlowsAsync`: DealUnitFlowDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.GetDealUnitFlowsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitFlowsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**DealUnitFlowDtoListEnvelope**](DealUnitFlowDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDealUnitFlowsCountAsync

> Int32Envelope GetDealUnitFlowsCountAsync(ctx).TenantId(tenantId).Execute()

Get deal unit flows count



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.GetDealUnitFlowsCountAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.GetDealUnitFlowsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitFlowsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.GetDealUnitFlowsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitFlowsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## UpdateDealUnitFlowAsync

> EmptyEnvelope UpdateDealUnitFlowAsync(ctx, dealUnitFlowId).TenantId(tenantId).DealUnitFlowUpdateDto(dealUnitFlowUpdateDto).Execute()

Update a deal unit flow



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitFlowUpdateDto := *openapiclient.NewDealUnitFlowUpdateDto() // DealUnitFlowUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.UpdateDealUnitFlowAsync(context.Background(), dealUnitFlowId).TenantId(tenantId).DealUnitFlowUpdateDto(dealUnitFlowUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.UpdateDealUnitFlowAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDealUnitFlowAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.UpdateDealUnitFlowAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDealUnitFlowAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **dealUnitFlowUpdateDto** | [**DealUnitFlowUpdateDto**](DealUnitFlowUpdateDto.md) |  | 

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


## UpdateDealUnitFlowStageAsync

> EmptyEnvelope UpdateDealUnitFlowStageAsync(ctx, dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).DealUnitFlowStageUpdateDto(dealUnitFlowStageUpdateDto).Execute()

Update a deal unit flow stage



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
	dealUnitFlowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitFlowStageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitFlowStageUpdateDto := *openapiclient.NewDealUnitFlowStageUpdateDto() // DealUnitFlowStageUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.UpdateDealUnitFlowStageAsync(context.Background(), dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).DealUnitFlowStageUpdateDto(dealUnitFlowStageUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.UpdateDealUnitFlowStageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDealUnitFlowStageAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.UpdateDealUnitFlowStageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 
**dealUnitFlowStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDealUnitFlowStageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **dealUnitFlowStageUpdateDto** | [**DealUnitFlowStageUpdateDto**](DealUnitFlowStageUpdateDto.md) |  | 

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

