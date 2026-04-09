# \OrdersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateOrder**](OrdersAPI.md#CalculateOrder) | **Put** /api/v2/OrdersService/Orders/{orderId}/Calculate | Calculates totals for an order.
[**CalculateOrderLine**](OrdersAPI.md#CalculateOrderLine) | **Put** /api/v2/OrdersService/Orders/{orderId}/Lines/{orderLineId}/Calculate | Calculates totals for an order line.
[**CreateOrder**](OrdersAPI.md#CreateOrder) | **Post** /api/v2/OrdersService/Orders | Creates a new order.
[**CreateOrderLine**](OrdersAPI.md#CreateOrderLine) | **Post** /api/v2/OrdersService/Orders/{orderId}/Lines | Creates a new order line.
[**DeleteOrder**](OrdersAPI.md#DeleteOrder) | **Delete** /api/v2/OrdersService/Orders/{orderId} | Deletes an order.
[**DeleteOrderLine**](OrdersAPI.md#DeleteOrderLine) | **Delete** /api/v2/OrdersService/Orders/{orderId}/Lines/{orderLineId} | Deletes an order line.
[**GetExtendedOrders**](OrdersAPI.md#GetExtendedOrders) | **Get** /api/v2/OrdersService/Orders/Extended | Gets a list of extended orders for a tenant.
[**GetOrder**](OrdersAPI.md#GetOrder) | **Get** /api/v2/OrdersService/Orders/{orderId} | Gets a specific order by ID.
[**GetOrderLine**](OrdersAPI.md#GetOrderLine) | **Get** /api/v2/OrdersService/Orders/{orderId}/Lines/{orderLineId} | Gets a specific order line.
[**GetOrderLines**](OrdersAPI.md#GetOrderLines) | **Get** /api/v2/OrdersService/Orders/{orderId}/Lines | Gets order lines for an order.
[**GetOrderLinesCount**](OrdersAPI.md#GetOrderLinesCount) | **Get** /api/v2/OrdersService/Orders/{orderId}/Lines/Count | Gets the count of order lines for an order.
[**GetOrders**](OrdersAPI.md#GetOrders) | **Get** /api/v2/OrdersService/Orders | Gets a list of orders for a tenant.
[**GetOrdersCount**](OrdersAPI.md#GetOrdersCount) | **Get** /api/v2/OrdersService/Orders/Count | Gets the count of orders for a tenant.
[**PreviewOrderEmailTemplate**](OrdersAPI.md#PreviewOrderEmailTemplate) | **Post** /api/v2/OrdersService/Orders/{orderId}/Emails/Preview | Preview the rendered email for an Order.
[**SendOrderEmail**](OrdersAPI.md#SendOrderEmail) | **Post** /api/v2/OrdersService/Orders/{orderId}/Emails/Send | Send a transactional email for an order.
[**SubmitCart**](OrdersAPI.md#SubmitCart) | **Post** /api/v2/OrdersService/Orders/SubmitCart | Submits a cart and creates an order.
[**UpdateOrder**](OrdersAPI.md#UpdateOrder) | **Put** /api/v2/OrdersService/Orders/{orderId} | Updates an existing order.
[**UpdateOrderLine**](OrdersAPI.md#UpdateOrderLine) | **Put** /api/v2/OrdersService/Orders/{orderId}/Lines/{orderLineId} | Updates an order line.



## CalculateOrder

> EmptyEnvelope CalculateOrder(ctx, orderId).TenantId(tenantId).Execute()

Calculates totals for an order.



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
	resp, r, err := apiClient.OrdersAPI.CalculateOrder(context.Background(), orderId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CalculateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateOrder`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CalculateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalculateOrderRequest struct via the builder pattern


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


## CalculateOrderLine

> EmptyEnvelope CalculateOrderLine(ctx, orderId, orderLineId).TenantId(tenantId).Execute()

Calculates totals for an order line.



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
	resp, r, err := apiClient.OrdersAPI.CalculateOrderLine(context.Background(), orderId, orderLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CalculateOrderLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateOrderLine`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CalculateOrderLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**orderLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalculateOrderLineRequest struct via the builder pattern


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


## CreateOrder

> EmptyEnvelope CreateOrder(ctx).TenantId(tenantId).OrderCreateDto(orderCreateDto).Execute()

Creates a new order.



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
	resp, r, err := apiClient.OrdersAPI.CreateOrder(context.Background()).TenantId(tenantId).OrderCreateDto(orderCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **orderCreateDto** | [**OrderCreateDto**](OrderCreateDto.md) |  | 

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


## CreateOrderLine

> EmptyEnvelope CreateOrderLine(ctx, orderId).TenantId(tenantId).OrderLineCreateDto(orderLineCreateDto).Execute()

Creates a new order line.



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
	resp, r, err := apiClient.OrdersAPI.CreateOrderLine(context.Background(), orderId).TenantId(tenantId).OrderLineCreateDto(orderLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CreateOrderLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrderLine`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CreateOrderLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **orderLineCreateDto** | [**OrderLineCreateDto**](OrderLineCreateDto.md) |  | 

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


## DeleteOrder

> EmptyEnvelope DeleteOrder(ctx, orderId).TenantId(tenantId).Execute()

Deletes an order.



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
	resp, r, err := apiClient.OrdersAPI.DeleteOrder(context.Background(), orderId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.DeleteOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrder`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.DeleteOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderRequest struct via the builder pattern


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


## DeleteOrderLine

> EmptyEnvelope DeleteOrderLine(ctx, orderId, orderLineId).TenantId(tenantId).Execute()

Deletes an order line.



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
	resp, r, err := apiClient.OrdersAPI.DeleteOrderLine(context.Background(), orderId, orderLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.DeleteOrderLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrderLine`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.DeleteOrderLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**orderLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderLineRequest struct via the builder pattern


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


## GetExtendedOrders

> ExtendedOrderDtoListEnvelope GetExtendedOrders(ctx).TenantId(tenantId).Execute()

Gets a list of extended orders for a tenant.



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
	resp, r, err := apiClient.OrdersAPI.GetExtendedOrders(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetExtendedOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedOrders`: ExtendedOrderDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetExtendedOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ExtendedOrderDtoListEnvelope**](ExtendedOrderDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> OrderDtoEnvelope GetOrder(ctx, orderId).TenantId(tenantId).Execute()

Gets a specific order by ID.



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
	resp, r, err := apiClient.OrdersAPI.GetOrder(context.Background(), orderId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrder`: OrderDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**OrderDtoEnvelope**](OrderDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderLine

> OrderLineDtoEnvelope GetOrderLine(ctx, orderId, orderLineId).TenantId(tenantId).Execute()

Gets a specific order line.



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
	resp, r, err := apiClient.OrdersAPI.GetOrderLine(context.Background(), orderId, orderLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrderLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderLine`: OrderLineDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrderLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**orderLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**OrderLineDtoEnvelope**](OrderLineDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderLines

> OrderLineDtoListEnvelope GetOrderLines(ctx, orderId).TenantId(tenantId).ItemId(itemId).Execute()

Gets order lines for an order.



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
	resp, r, err := apiClient.OrdersAPI.GetOrderLines(context.Background(), orderId).TenantId(tenantId).ItemId(itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrderLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderLines`: OrderLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrderLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **itemId** | **string** |  | 

### Return type

[**OrderLineDtoListEnvelope**](OrderLineDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderLinesCount

> Int32Envelope GetOrderLinesCount(ctx, orderId).TenantId(tenantId).Execute()

Gets the count of order lines for an order.



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
	resp, r, err := apiClient.OrdersAPI.GetOrderLinesCount(context.Background(), orderId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrderLinesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderLinesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrderLinesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderLinesCountRequest struct via the builder pattern


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


## GetOrders

> OrderDtoListEnvelope GetOrders(ctx).TenantId(tenantId).Execute()

Gets a list of orders for a tenant.



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
	resp, r, err := apiClient.OrdersAPI.GetOrders(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrders`: OrderDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**OrderDtoListEnvelope**](OrderDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersCount

> Int32Envelope GetOrdersCount(ctx).TenantId(tenantId).Execute()

Gets the count of orders for a tenant.



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
	resp, r, err := apiClient.OrdersAPI.GetOrdersCount(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrdersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrdersCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersCountRequest struct via the builder pattern


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


## PreviewOrderEmailTemplate

> PreviewOrderEmailTemplate(ctx, orderId).TenantId(tenantId).EmailDispatchRequest(emailDispatchRequest).Execute()

Preview the rendered email for an Order.



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
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	emailDispatchRequest := *openapiclient.NewEmailDispatchRequest("Title_example", "Message_example", "Culture_example", "UiCulture_example", []string{"Recipients_example"}) // EmailDispatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersAPI.PreviewOrderEmailTemplate(context.Background(), orderId).TenantId(tenantId).EmailDispatchRequest(emailDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.PreviewOrderEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewOrderEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **emailDispatchRequest** | [**EmailDispatchRequest**](EmailDispatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendOrderEmail

> EmptyEnvelope SendOrderEmail(ctx, orderId).TenantId(tenantId).EmailDispatchRequest(emailDispatchRequest).Execute()

Send a transactional email for an order.



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
	emailDispatchRequest := *openapiclient.NewEmailDispatchRequest("Title_example", "Message_example", "Culture_example", "UiCulture_example", []string{"Recipients_example"}) // EmailDispatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.SendOrderEmail(context.Background(), orderId).TenantId(tenantId).EmailDispatchRequest(emailDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.SendOrderEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendOrderEmail`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.SendOrderEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendOrderEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **emailDispatchRequest** | [**EmailDispatchRequest**](EmailDispatchRequest.md) |  | 

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


## SubmitCart

> OrderDtoEnvelope SubmitCart(ctx).CartId(cartId).Execute()

Submits a cart and creates an order.



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
	resp, r, err := apiClient.OrdersAPI.SubmitCart(context.Background()).CartId(cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.SubmitCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCart`: OrderDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.SubmitCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartId** | **string** |  | 

### Return type

[**OrderDtoEnvelope**](OrderDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrder

> EmptyEnvelope UpdateOrder(ctx, orderId).TenantId(tenantId).OrderUpdateDto(orderUpdateDto).Execute()

Updates an existing order.



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
	resp, r, err := apiClient.OrdersAPI.UpdateOrder(context.Background(), orderId).TenantId(tenantId).OrderUpdateDto(orderUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.UpdateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrder`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.UpdateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **orderUpdateDto** | [**OrderUpdateDto**](OrderUpdateDto.md) |  | 

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


## UpdateOrderLine

> EmptyEnvelope UpdateOrderLine(ctx, orderId, orderLineId).TenantId(tenantId).OrderLineUpdateDto(orderLineUpdateDto).Execute()

Updates an order line.



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
	resp, r, err := apiClient.OrdersAPI.UpdateOrderLine(context.Background(), orderId, orderLineId).TenantId(tenantId).OrderLineUpdateDto(orderLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.UpdateOrderLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrderLine`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.UpdateOrderLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** |  | 
**orderLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **orderLineUpdateDto** | [**OrderLineUpdateDto**](OrderLineUpdateDto.md) |  | 

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

