# \RecordsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddItemToCart**](RecordsAPI.md#AddItemToCart) | **Post** /api/v2/CartService/Records/AddItem | Add an item to a cart
[**AddProductToCartAsync**](RecordsAPI.md#AddProductToCartAsync) | **Post** /api/v2/CartService/Records | Add a product to a cart with tracking
[**ClearCartAsync**](RecordsAPI.md#ClearCartAsync) | **Post** /api/v2/CartService/Records/ClearCart | Clear all items from a cart
[**DecreaseItemCartRecord**](RecordsAPI.md#DecreaseItemCartRecord) | **Put** /api/v2/CartService/Records/{recordId}/Decrease | Decrease cart record quantity
[**GetItemCartRecord**](RecordsAPI.md#GetItemCartRecord) | **Get** /api/v2/CartService/Records/{recordId}/Details | Get a cart record by ID
[**GetItemsInCartAsync**](RecordsAPI.md#GetItemsInCartAsync) | **Get** /api/v2/CartService/Records/{cartId} | Get all items in a cart
[**IncreaseItemCartRecord**](RecordsAPI.md#IncreaseItemCartRecord) | **Put** /api/v2/CartService/Records/{recordId}/Increase | Increase cart record quantity
[**IsItemAlreadyInCart**](RecordsAPI.md#IsItemAlreadyInCart) | **Get** /api/v2/CartService/Records/IsInCart | Check if an item is in a cart
[**RemoveProductFromCartByParams**](RecordsAPI.md#RemoveProductFromCartByParams) | **Delete** /api/v2/CartService/Records | Remove a product from a cart
[**RemoveProductFromCartByRecordId**](RecordsAPI.md#RemoveProductFromCartByRecordId) | **Delete** /api/v2/CartService/Records/{recordId} | Remove a product from a cart by record ID
[**UpdateItemCartRecord**](RecordsAPI.md#UpdateItemCartRecord) | **Put** /api/v2/CartService/Records/{recordId} | Update a cart record



## AddItemToCart

> EmptyEnvelope AddItemToCart(ctx).CartId(cartId).ItemId(itemId).Quantity(quantity).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Add an item to a cart



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
	cartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	quantity := int32(56) // int32 |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.AddItemToCart(context.Background()).CartId(cartId).ItemId(itemId).Quantity(quantity).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.AddItemToCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddItemToCart`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.AddItemToCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddItemToCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartId** | **string** |  | 
 **itemId** | **string** |  | 
 **quantity** | **int32** |  | 
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


## AddProductToCartAsync

> EmptyEnvelope AddProductToCartAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemCartRecordCreateDto(itemCartRecordCreateDto).Execute()

Add a product to a cart with tracking



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemCartRecordCreateDto := *openapiclient.NewItemCartRecordCreateDto() // ItemCartRecordCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.AddProductToCartAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemCartRecordCreateDto(itemCartRecordCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.AddProductToCartAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProductToCartAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.AddProductToCartAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProductToCartAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemCartRecordCreateDto** | [**ItemCartRecordCreateDto**](ItemCartRecordCreateDto.md) |  | 

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


## ClearCartAsync

> EmptyEnvelope ClearCartAsync(ctx).CartID(cartID).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Clear all items from a cart



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
	cartID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.ClearCartAsync(context.Background()).CartID(cartID).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.ClearCartAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearCartAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.ClearCartAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearCartAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartID** | **string** |  | 
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


## DecreaseItemCartRecord

> EmptyEnvelope DecreaseItemCartRecord(ctx, recordId).Quantity(quantity).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Decrease cart record quantity



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
	recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quantity := float64(1.2) // float64 |  (optional) (default to 1)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.DecreaseItemCartRecord(context.Background(), recordId).Quantity(quantity).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.DecreaseItemCartRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DecreaseItemCartRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.DecreaseItemCartRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecreaseItemCartRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quantity** | **float64** |  | [default to 1]
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


## GetItemCartRecord

> EmptyEnvelope GetItemCartRecord(ctx, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get a cart record by ID



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
	recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetItemCartRecord(context.Background(), recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetItemCartRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemCartRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetItemCartRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemCartRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetItemsInCartAsync

> ItemCartRecordDtoListEnvelope GetItemsInCartAsync(ctx, cartId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all items in a cart



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.GetItemsInCartAsync(context.Background(), cartId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetItemsInCartAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsInCartAsync`: ItemCartRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetItemsInCartAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsInCartAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemCartRecordDtoListEnvelope**](ItemCartRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncreaseItemCartRecord

> EmptyEnvelope IncreaseItemCartRecord(ctx, recordId).Quantity(quantity).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Increase cart record quantity



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
	recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quantity := float64(1.2) // float64 |  (optional) (default to 1)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.IncreaseItemCartRecord(context.Background(), recordId).Quantity(quantity).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.IncreaseItemCartRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IncreaseItemCartRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.IncreaseItemCartRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncreaseItemCartRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **quantity** | **float64** |  | [default to 1]
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


## IsItemAlreadyInCart

> BooleanEnvelope IsItemAlreadyInCart(ctx).ItemID(itemID).CartID(cartID).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Check if an item is in a cart



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
	itemID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	cartID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.IsItemAlreadyInCart(context.Background()).ItemID(itemID).CartID(cartID).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.IsItemAlreadyInCart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsItemAlreadyInCart`: BooleanEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.IsItemAlreadyInCart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsItemAlreadyInCartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemID** | **string** |  | 
 **cartID** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BooleanEnvelope**](BooleanEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProductFromCartByParams

> EmptyEnvelope RemoveProductFromCartByParams(ctx).CartId(cartId).ProductId(productId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove a product from a cart



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
	cartId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.RemoveProductFromCartByParams(context.Background()).CartId(cartId).ProductId(productId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.RemoveProductFromCartByParams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveProductFromCartByParams`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.RemoveProductFromCartByParams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProductFromCartByParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartId** | **string** |  | 
 **productId** | **string** |  | 
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


## RemoveProductFromCartByRecordId

> EmptyEnvelope RemoveProductFromCartByRecordId(ctx, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove a product from a cart by record ID



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
	recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.RemoveProductFromCartByRecordId(context.Background(), recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.RemoveProductFromCartByRecordId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveProductFromCartByRecordId`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.RemoveProductFromCartByRecordId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProductFromCartByRecordIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## UpdateItemCartRecord

> EmptyEnvelope UpdateItemCartRecord(ctx, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemCartRecordUpdateDto(itemCartRecordUpdateDto).Execute()

Update a cart record



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
	recordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemCartRecordUpdateDto := *openapiclient.NewItemCartRecordUpdateDto() // ItemCartRecordUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordsAPI.UpdateItemCartRecord(context.Background(), recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemCartRecordUpdateDto(itemCartRecordUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.UpdateItemCartRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemCartRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.UpdateItemCartRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemCartRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemCartRecordUpdateDto** | [**ItemCartRecordUpdateDto**](ItemCartRecordUpdateDto.md) |  | 

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

