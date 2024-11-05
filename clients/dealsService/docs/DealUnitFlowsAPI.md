# \DealUnitFlowsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2DealsServiceDealUnitFlowsCountGet**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsCountGet) | **Get** /api/v2/DealsService/DealUnitFlows/Count | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDelete**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDelete) | **Delete** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId} | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGet**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGet) | **Get** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId} | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPut**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPut) | **Put** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId} | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGet**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGet) | **Get** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages/Count | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDelete**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDelete) | **Delete** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages/{dealUnitFlowStageId} | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGet**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGet) | **Get** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages/{dealUnitFlowStageId} | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPut**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPut) | **Put** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages/{dealUnitFlowStageId} | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGet**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGet) | **Get** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages | 
[**ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPost**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPost) | **Post** /api/v2/DealsService/DealUnitFlows/{dealUnitFlowId}/Stages | 
[**ApiV2DealsServiceDealUnitFlowsGet**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsGet) | **Get** /api/v2/DealsService/DealUnitFlows | 
[**ApiV2DealsServiceDealUnitFlowsPost**](DealUnitFlowsAPI.md#ApiV2DealsServiceDealUnitFlowsPost) | **Post** /api/v2/DealsService/DealUnitFlows | 



## ApiV2DealsServiceDealUnitFlowsCountGet

> Int32Envelope ApiV2DealsServiceDealUnitFlowsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsCountGetRequest struct via the builder pattern


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


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDelete

> EmptyEnvelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDelete(ctx, dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDelete(context.Background(), dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdDeleteRequest struct via the builder pattern


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


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGet

> DealUnitFlowDtoEnvelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGet(ctx, dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGet(context.Background(), dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGet`: DealUnitFlowDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DealUnitFlowDtoEnvelope**](DealUnitFlowDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPut

> EmptyEnvelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPut(ctx, dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DealUnitFlowUpdateDto(dealUnitFlowUpdateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	dealUnitFlowUpdateDto := *openapiclient.NewDealUnitFlowUpdateDto() // DealUnitFlowUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPut(context.Background(), dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DealUnitFlowUpdateDto(dealUnitFlowUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **dealUnitFlowUpdateDto** | [**DealUnitFlowUpdateDto**](DealUnitFlowUpdateDto.md) |  | 

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


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGet

> Int32Envelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGet(ctx, dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGet(context.Background(), dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesCountGetRequest struct via the builder pattern


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


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDelete

> EmptyEnvelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDelete(ctx, dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDelete(context.Background(), dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 
**dealUnitFlowStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdDeleteRequest struct via the builder pattern


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


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGet

> DealUnitFlowStageDtoEnvelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGet(ctx, dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGet(context.Background(), dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGet`: DealUnitFlowStageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 
**dealUnitFlowStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DealUnitFlowStageDtoEnvelope**](DealUnitFlowStageDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPut

> EmptyEnvelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPut(ctx, dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DealUnitFlowStageUpdateDto(dealUnitFlowStageUpdateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	dealUnitFlowStageUpdateDto := *openapiclient.NewDealUnitFlowStageUpdateDto() // DealUnitFlowStageUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPut(context.Background(), dealUnitFlowId, dealUnitFlowStageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DealUnitFlowStageUpdateDto(dealUnitFlowStageUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 
**dealUnitFlowStageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesDealUnitFlowStageIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **dealUnitFlowStageUpdateDto** | [**DealUnitFlowStageUpdateDto**](DealUnitFlowStageUpdateDto.md) |  | 

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


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGet

> DealUnitFlowStageDtoListEnvelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGet(ctx, dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGet(context.Background(), dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGet`: DealUnitFlowStageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DealUnitFlowStageDtoListEnvelope**](DealUnitFlowStageDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPost

> EmptyEnvelope ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPost(ctx, dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DealUnitFlowStageCreateDto(dealUnitFlowStageCreateDto).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	dealUnitFlowStageCreateDto := *openapiclient.NewDealUnitFlowStageCreateDto() // DealUnitFlowStageCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPost(context.Background(), dealUnitFlowId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DealUnitFlowStageCreateDto(dealUnitFlowStageCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitFlowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsDealUnitFlowIdStagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **dealUnitFlowStageCreateDto** | [**DealUnitFlowStageCreateDto**](DealUnitFlowStageCreateDto.md) |  | 

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


## ApiV2DealsServiceDealUnitFlowsGet

> DealUnitFlowDtoListEnvelope ApiV2DealsServiceDealUnitFlowsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsGet`: DealUnitFlowDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DealUnitFlowDtoListEnvelope**](DealUnitFlowDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitFlowsPost

> EmptyEnvelope ApiV2DealsServiceDealUnitFlowsPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DealUnitFlowCreateDto(dealUnitFlowCreateDto).Execute()



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
	dealUnitFlowCreateDto := *openapiclient.NewDealUnitFlowCreateDto() // DealUnitFlowCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DealUnitFlowCreateDto(dealUnitFlowCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitFlowsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitFlowsAPI.ApiV2DealsServiceDealUnitFlowsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitFlowsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **dealUnitFlowCreateDto** | [**DealUnitFlowCreateDto**](DealUnitFlowCreateDto.md) |  | 

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

