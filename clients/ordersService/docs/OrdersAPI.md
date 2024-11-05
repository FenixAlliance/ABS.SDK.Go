# \OrdersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2OrdersServiceOrdersCountGet**](OrdersAPI.md#ApiV2OrdersServiceOrdersCountGet) | **Get** /api/v2/OrdersService/Orders/Count | 
[**ApiV2OrdersServiceOrdersExtendedGet**](OrdersAPI.md#ApiV2OrdersServiceOrdersExtendedGet) | **Get** /api/v2/OrdersService/Orders/Extended | 
[**ApiV2OrdersServiceOrdersOrderIdCalculatePut**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdCalculatePut) | **Put** /api/v2/OrdersService/Orders/{orderId}/Calculate | 
[**ApiV2OrdersServiceOrdersOrderIdDelete**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdDelete) | **Delete** /api/v2/OrdersService/Orders/{orderId} | 
[**ApiV2OrdersServiceOrdersOrderIdLinesCountGet**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdLinesCountGet) | **Get** /api/v2/OrdersService/Orders/{orderId}/Lines/Count | 
[**ApiV2OrdersServiceOrdersOrderIdLinesGet**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdLinesGet) | **Get** /api/v2/OrdersService/Orders/{orderId}/Lines | 
[**ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePut**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePut) | **Put** /api/v2/OrdersService/Orders/{orderId}/Lines/{orderLineId}/Calculate | 
[**ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDelete**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDelete) | **Delete** /api/v2/OrdersService/Orders/{orderId}/Lines/{orderLineId} | 
[**ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGet**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGet) | **Get** /api/v2/OrdersService/Orders/{orderId}/Lines/{orderLineId} | 
[**ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPut**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPut) | **Put** /api/v2/OrdersService/Orders/{orderId}/Lines/{orderLineId} | 
[**ApiV2OrdersServiceOrdersOrderIdLinesPost**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdLinesPost) | **Post** /api/v2/OrdersService/Orders/{orderId}/Lines | 
[**ApiV2OrdersServiceOrdersOrderIdPut**](OrdersAPI.md#ApiV2OrdersServiceOrdersOrderIdPut) | **Put** /api/v2/OrdersService/Orders/{orderId} | 
[**ApiV2OrdersServiceOrdersPost**](OrdersAPI.md#ApiV2OrdersServiceOrdersPost) | **Post** /api/v2/OrdersService/Orders | 
[**ApiV2OrdersServiceOrdersSubmitCartPost**](OrdersAPI.md#ApiV2OrdersServiceOrdersSubmitCartPost) | **Post** /api/v2/OrdersService/Orders/SubmitCart | 
[**GetOrderAsync**](OrdersAPI.md#GetOrderAsync) | **Get** /api/v2/OrdersService/Orders/{orderId} | 
[**GetOrdersAsync**](OrdersAPI.md#GetOrdersAsync) | **Get** /api/v2/OrdersService/Orders | 



## ApiV2OrdersServiceOrdersCountGet

> Int32Envelope ApiV2OrdersServiceOrdersCountGet(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersCountGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersCountGetRequest struct via the builder pattern


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


## ApiV2OrdersServiceOrdersExtendedGet

> ExtendedOrderDtoListEnvelope ApiV2OrdersServiceOrdersExtendedGet(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersExtendedGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersExtendedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersExtendedGet`: ExtendedOrderDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersExtendedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersExtendedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ExtendedOrderDtoListEnvelope**](ExtendedOrderDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2OrdersServiceOrdersOrderIdCalculatePut

> EmptyEnvelope ApiV2OrdersServiceOrdersOrderIdCalculatePut(ctx, orderId).TenantId(tenantId).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdCalculatePut(context.Background(), orderId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdCalculatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdCalculatePut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdCalculatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdCalculatePutRequest struct via the builder pattern


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


## ApiV2OrdersServiceOrdersOrderIdDelete

> EmptyEnvelope ApiV2OrdersServiceOrdersOrderIdDelete(ctx, orderId).TenantId(tenantId).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdDelete(context.Background(), orderId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdDeleteRequest struct via the builder pattern


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


## ApiV2OrdersServiceOrdersOrderIdLinesCountGet

> Int32Envelope ApiV2OrdersServiceOrdersOrderIdLinesCountGet(ctx, orderId).TenantId(tenantId).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesCountGet(context.Background(), orderId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdLinesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdLinesCountGetRequest struct via the builder pattern


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


## ApiV2OrdersServiceOrdersOrderIdLinesGet

> OrderLineDtoListEnvelope ApiV2OrdersServiceOrdersOrderIdLinesGet(ctx, orderId).TenantId(tenantId).ItemId(itemId).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesGet(context.Background(), orderId).TenantId(tenantId).ItemId(itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdLinesGet`: OrderLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdLinesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **itemId** | **string** |  | 

### Return type

[**OrderLineDtoListEnvelope**](OrderLineDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePut

> EmptyEnvelope ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePut(ctx, orderId, orderLineId).TenantId(tenantId).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orderLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePut(context.Background(), orderId, orderLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**orderLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdCalculatePutRequest struct via the builder pattern


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


## ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDelete

> EmptyEnvelope ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDelete(ctx, orderId, orderLineId).TenantId(tenantId).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orderLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDelete(context.Background(), orderId, orderLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**orderLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdDeleteRequest struct via the builder pattern


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


## ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGet

> OrderLineDtoEnvelope ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGet(ctx, orderId, orderLineId).TenantId(tenantId).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orderLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGet(context.Background(), orderId, orderLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGet`: OrderLineDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**orderLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**OrderLineDtoEnvelope**](OrderLineDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPut

> EmptyEnvelope ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPut(ctx, orderId, orderLineId).TenantId(tenantId).OrderLineUpdateDto(orderLineUpdateDto).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orderLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orderLineUpdateDto := *openapiclient.NewOrderLineUpdateDto() // OrderLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPut(context.Background(), orderId, orderLineId).TenantId(tenantId).OrderLineUpdateDto(orderLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**orderLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdLinesOrderLineIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **orderLineUpdateDto** | [**OrderLineUpdateDto**](OrderLineUpdateDto.md) |  | 

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


## ApiV2OrdersServiceOrdersOrderIdLinesPost

> EmptyEnvelope ApiV2OrdersServiceOrdersOrderIdLinesPost(ctx, orderId).TenantId(tenantId).OrderLineCreateDto(orderLineCreateDto).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orderLineCreateDto := *openapiclient.NewOrderLineCreateDto() // OrderLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesPost(context.Background(), orderId).TenantId(tenantId).OrderLineCreateDto(orderLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdLinesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdLinesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdLinesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **orderLineCreateDto** | [**OrderLineCreateDto**](OrderLineCreateDto.md) |  | 

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


## ApiV2OrdersServiceOrdersOrderIdPut

> EmptyEnvelope ApiV2OrdersServiceOrdersOrderIdPut(ctx, orderId).TenantId(tenantId).OrderUpdateDto(orderUpdateDto).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	orderUpdateDto := *openapiclient.NewOrderUpdateDto() // OrderUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersOrderIdPut(context.Background(), orderId).TenantId(tenantId).OrderUpdateDto(orderUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersOrderIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersOrderIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersOrderIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **orderUpdateDto** | [**OrderUpdateDto**](OrderUpdateDto.md) |  | 

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


## ApiV2OrdersServiceOrdersPost

> EmptyEnvelope ApiV2OrdersServiceOrdersPost(ctx).TenantId(tenantId).OrderCreateDto(orderCreateDto).Execute()



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
	orderCreateDto := *openapiclient.NewOrderCreateDto() // OrderCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersPost(context.Background()).TenantId(tenantId).OrderCreateDto(orderCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **orderCreateDto** | [**OrderCreateDto**](OrderCreateDto.md) |  | 

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


## ApiV2OrdersServiceOrdersSubmitCartPost

> OrderDtoEnvelope ApiV2OrdersServiceOrdersSubmitCartPost(ctx).CartId(cartId).Execute()



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
	cartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.ApiV2OrdersServiceOrdersSubmitCartPost(context.Background()).CartId(cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ApiV2OrdersServiceOrdersSubmitCartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2OrdersServiceOrdersSubmitCartPost`: OrderDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.ApiV2OrdersServiceOrdersSubmitCartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2OrdersServiceOrdersSubmitCartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartId** | **string** |  | 

### Return type

[**OrderDtoEnvelope**](OrderDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderAsync

> OrderDtoEnvelope GetOrderAsync(ctx, orderId).TenantId(tenantId).Execute()



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetOrderAsync(context.Background(), orderId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrderAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderAsync`: OrderDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrderAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**OrderDtoEnvelope**](OrderDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersAsync

> OrderDtoListEnvelope GetOrdersAsync(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.OrdersAPI.GetOrdersAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrdersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersAsync`: OrderDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrdersAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**OrderDtoListEnvelope**](OrderDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

