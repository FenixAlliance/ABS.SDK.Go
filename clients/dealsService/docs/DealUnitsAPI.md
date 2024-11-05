# \DealUnitsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2DealsServiceDealUnitsCountGet**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsCountGet) | **Get** /api/v2/DealsService/DealUnits/Count | 
[**ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut) | **Put** /api/v2/DealsService/DealUnits/{dealUnitId}/Calculate | 
[**ApiV2DealsServiceDealUnitsDealUnitIdDelete**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdDelete) | **Delete** /api/v2/DealsService/DealUnits/{dealUnitId} | 
[**ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId}/Extended | 
[**ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/Count | 
[**ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut) | **Put** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/{dealUnitLineId}/Calculate | 
[**ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete) | **Delete** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/{dealUnitLineId} | 
[**ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/{dealUnitLineId} | 
[**ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut) | **Put** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/{dealUnitLineId} | 
[**ApiV2DealsServiceDealUnitsDealUnitIdLinesGet**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdLinesGet) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines | 
[**ApiV2DealsServiceDealUnitsDealUnitIdLinesPost**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdLinesPost) | **Post** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines | 
[**ApiV2DealsServiceDealUnitsDealUnitIdPut**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsDealUnitIdPut) | **Put** /api/v2/DealsService/DealUnits/{dealUnitId} | 
[**ApiV2DealsServiceDealUnitsExtendedGet**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsExtendedGet) | **Get** /api/v2/DealsService/DealUnits/Extended | 
[**ApiV2DealsServiceDealUnitsGet**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsGet) | **Get** /api/v2/DealsService/DealUnits | 
[**ApiV2DealsServiceDealUnitsPost**](DealUnitsAPI.md#ApiV2DealsServiceDealUnitsPost) | **Post** /api/v2/DealsService/DealUnits | 
[**GetDealUnitAsync**](DealUnitsAPI.md#GetDealUnitAsync) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId} | 



## ApiV2DealsServiceDealUnitsCountGet

> Int32Envelope ApiV2DealsServiceDealUnitsCountGet(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsCountGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut

> EmptyEnvelope ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut(ctx, dealUnitId).TenantId(tenantId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdCalculatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdCalculatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## ApiV2DealsServiceDealUnitsDealUnitIdDelete

> EmptyEnvelope ApiV2DealsServiceDealUnitsDealUnitIdDelete(ctx, dealUnitId).TenantId(tenantId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdDelete(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet

> ExtendedDealUnitDtoEnvelope ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet(ctx, dealUnitId).TenantId(tenantId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet`: ExtendedDealUnitDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdExtendedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdExtendedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**ExtendedDealUnitDtoEnvelope**](ExtendedDealUnitDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet

> Int32Envelope ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet(ctx, dealUnitId).TenantId(tenantId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdLinesCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut

> EmptyEnvelope ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut(ctx, dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut(context.Background(), dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 
**dealUnitLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdCalculatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete

> EmptyEnvelope ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete(ctx, dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete(context.Background(), dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 
**dealUnitLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet

> DealUnitLineDtoEnvelope ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet(ctx, dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet(context.Background(), dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet`: DealUnitLineDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 
**dealUnitLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**DealUnitLineDtoEnvelope**](DealUnitLineDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut

> EmptyEnvelope ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut(ctx, dealUnitId, dealUnitLineId).TenantId(tenantId).DealUnitLineUpdateDto(dealUnitLineUpdateDto).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineUpdateDto := *openapiclient.NewDealUnitLineUpdateDto() // DealUnitLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut(context.Background(), dealUnitId, dealUnitLineId).TenantId(tenantId).DealUnitLineUpdateDto(dealUnitLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 
**dealUnitLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdLinesDealUnitLineIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **dealUnitLineUpdateDto** | [**DealUnitLineUpdateDto**](DealUnitLineUpdateDto.md) |  | 

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


## ApiV2DealsServiceDealUnitsDealUnitIdLinesGet

> DealUnitLineDtoListEnvelope ApiV2DealsServiceDealUnitsDealUnitIdLinesGet(ctx, dealUnitId).TenantId(tenantId).ItemId(itemId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesGet(context.Background(), dealUnitId).TenantId(tenantId).ItemId(itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdLinesGet`: DealUnitLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdLinesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **itemId** | **string** |  | 

### Return type

[**DealUnitLineDtoListEnvelope**](DealUnitLineDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitsDealUnitIdLinesPost

> EmptyEnvelope ApiV2DealsServiceDealUnitsDealUnitIdLinesPost(ctx, dealUnitId).TenantId(tenantId).DealUnitLineCreateDto(dealUnitLineCreateDto).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineCreateDto := *openapiclient.NewDealUnitLineCreateDto() // DealUnitLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesPost(context.Background(), dealUnitId).TenantId(tenantId).DealUnitLineCreateDto(dealUnitLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdLinesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdLinesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdLinesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **dealUnitLineCreateDto** | [**DealUnitLineCreateDto**](DealUnitLineCreateDto.md) |  | 

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


## ApiV2DealsServiceDealUnitsDealUnitIdPut

> EmptyEnvelope ApiV2DealsServiceDealUnitsDealUnitIdPut(ctx, dealUnitId).TenantId(tenantId).DealUnitUpdateDto(dealUnitUpdateDto).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitUpdateDto := *openapiclient.NewDealUnitUpdateDto() // DealUnitUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdPut(context.Background(), dealUnitId).TenantId(tenantId).DealUnitUpdateDto(dealUnitUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsDealUnitIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsDealUnitIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsDealUnitIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **dealUnitUpdateDto** | [**DealUnitUpdateDto**](DealUnitUpdateDto.md) |  | 

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


## ApiV2DealsServiceDealUnitsExtendedGet

> ExtendedDealUnitDtoListEnvelope ApiV2DealsServiceDealUnitsExtendedGet(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsExtendedGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsExtendedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsExtendedGet`: ExtendedDealUnitDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsExtendedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsExtendedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ExtendedDealUnitDtoListEnvelope**](ExtendedDealUnitDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitsGet

> DealUnitDtoListEnvelope ApiV2DealsServiceDealUnitsGet(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsGet`: DealUnitDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**DealUnitDtoListEnvelope**](DealUnitDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceDealUnitsPost

> EmptyEnvelope ApiV2DealsServiceDealUnitsPost(ctx).TenantId(tenantId).DealUnitCreateDto(dealUnitCreateDto).Execute()



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
	dealUnitCreateDto := *openapiclient.NewDealUnitCreateDto() // DealUnitCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.ApiV2DealsServiceDealUnitsPost(context.Background()).TenantId(tenantId).DealUnitCreateDto(dealUnitCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.ApiV2DealsServiceDealUnitsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceDealUnitsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.ApiV2DealsServiceDealUnitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceDealUnitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **dealUnitCreateDto** | [**DealUnitCreateDto**](DealUnitCreateDto.md) |  | 

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


## GetDealUnitAsync

> DealUnitDtoEnvelope GetDealUnitAsync(ctx, dealUnitId).TenantId(tenantId).Execute()



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetDealUnitAsync(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetDealUnitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitAsync`: DealUnitDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetDealUnitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**DealUnitDtoEnvelope**](DealUnitDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

