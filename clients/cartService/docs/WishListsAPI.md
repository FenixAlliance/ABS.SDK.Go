# \WishListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProductToWishList**](WishListsAPI.md#AddProductToWishList) | **Post** /api/v2/CartService/WishLists/Records | Add a product to a wish list
[**CreateWishList**](WishListsAPI.md#CreateWishList) | **Post** /api/v2/CartService/WishLists | Create a wish list
[**DeleteWishList**](WishListsAPI.md#DeleteWishList) | **Delete** /api/v2/CartService/WishLists/{wishListId} | Delete a wish list
[**DeleteWishListRecord**](WishListsAPI.md#DeleteWishListRecord) | **Delete** /api/v2/CartService/WishLists/Records/{recordId} | Delete a wish list record
[**GetCartWishListDetailsAsync**](WishListsAPI.md#GetCartWishListDetailsAsync) | **Get** /api/v2/CartService/WishLists/{wishListId}/Details | Get wish list details
[**GetCartWishListItemsAsync**](WishListsAPI.md#GetCartWishListItemsAsync) | **Get** /api/v2/CartService/WishLists/{wishListId}/Records | Get wish list item records
[**GetWishListAsync**](WishListsAPI.md#GetWishListAsync) | **Get** /api/v2/CartService/WishLists/{cartId} | Get wish lists for a cart
[**IsProductInWishLists**](WishListsAPI.md#IsProductInWishLists) | **Get** /api/v2/CartService/WishLists/Contains | Check if a product is in any wish list
[**UpdateProductToWishList**](WishListsAPI.md#UpdateProductToWishList) | **Put** /api/v2/CartService/WishLists/{wishListId} | Update a wish list
[**WishListExists**](WishListsAPI.md#WishListExists) | **Get** /api/v2/CartService/WishLists/Exists | Check if a wish list exists
[**WishListExistsHeadAsync**](WishListsAPI.md#WishListExistsHeadAsync) | **Head** /api/v2/CartService/WishLists/Exists | Check if a wish list exists (HEAD)



## AddProductToWishList

> EmptyEnvelope AddProductToWishList(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProductToWishListRequest(productToWishListRequest).Execute()

Add a product to a wish list



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
	productToWishListRequest := *openapiclient.NewProductToWishListRequest() // ProductToWishListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishListsAPI.AddProductToWishList(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProductToWishListRequest(productToWishListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.AddProductToWishList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProductToWishList`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.AddProductToWishList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddProductToWishListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **productToWishListRequest** | [**ProductToWishListRequest**](ProductToWishListRequest.md) |  | 

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


## CreateWishList

> EmptyEnvelope CreateWishList(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).NewWishListRequest(newWishListRequest).Execute()

Create a wish list



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
	newWishListRequest := *openapiclient.NewNewWishListRequest() // NewWishListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishListsAPI.CreateWishList(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).NewWishListRequest(newWishListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.CreateWishList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWishList`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.CreateWishList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWishListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **newWishListRequest** | [**NewWishListRequest**](NewWishListRequest.md) |  | 

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


## DeleteWishList

> EmptyEnvelope DeleteWishList(ctx, wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a wish list



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
	wishListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishListsAPI.DeleteWishList(context.Background(), wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.DeleteWishList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWishList`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.DeleteWishList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wishListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWishListRequest struct via the builder pattern


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


## DeleteWishListRecord

> DeleteWishListRecord(ctx, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a wish list record



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
	r, err := apiClient.WishListsAPI.DeleteWishListRecord(context.Background(), recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.DeleteWishListRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWishListRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCartWishListDetailsAsync

> WishListDto GetCartWishListDetailsAsync(ctx, wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get wish list details



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
	wishListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishListsAPI.GetCartWishListDetailsAsync(context.Background(), wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.GetCartWishListDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCartWishListDetailsAsync`: WishListDto
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.GetCartWishListDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wishListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCartWishListDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WishListDto**](WishListDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCartWishListItemsAsync

> []WishListItemRecordDto GetCartWishListItemsAsync(ctx, wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get wish list item records



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
	wishListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishListsAPI.GetCartWishListItemsAsync(context.Background(), wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.GetCartWishListItemsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCartWishListItemsAsync`: []WishListItemRecordDto
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.GetCartWishListItemsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wishListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCartWishListItemsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]WishListItemRecordDto**](WishListItemRecordDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWishListAsync

> []WishListDto GetWishListAsync(ctx, cartId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get wish lists for a cart



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
	resp, r, err := apiClient.WishListsAPI.GetWishListAsync(context.Background(), cartId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.GetWishListAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWishListAsync`: []WishListDto
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.GetWishListAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWishListAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**[]WishListDto**](WishListDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsProductInWishLists

> BooleanEnvelope IsProductInWishLists(ctx).CartId(cartId).ProductId(productId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Check if a product is in any wish list



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
	resp, r, err := apiClient.WishListsAPI.IsProductInWishLists(context.Background()).CartId(cartId).ProductId(productId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.IsProductInWishLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IsProductInWishLists`: BooleanEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.IsProductInWishLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIsProductInWishListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cartId** | **string** |  | 
 **productId** | **string** |  | 
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


## UpdateProductToWishList

> EmptyEnvelope UpdateProductToWishList(ctx, wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WishListUpdateDto(wishListUpdateDto).Execute()

Update a wish list



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
	wishListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	wishListUpdateDto := *openapiclient.NewWishListUpdateDto("Title_example") // WishListUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishListsAPI.UpdateProductToWishList(context.Background(), wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WishListUpdateDto(wishListUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.UpdateProductToWishList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProductToWishList`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.UpdateProductToWishList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wishListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductToWishListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **wishListUpdateDto** | [**WishListUpdateDto**](WishListUpdateDto.md) |  | 

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


## WishListExists

> BooleanEnvelope WishListExists(ctx).WishListId(wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Check if a wish list exists



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
	wishListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishListsAPI.WishListExists(context.Background()).WishListId(wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.WishListExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishListExists`: BooleanEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.WishListExists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishListExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wishListId** | **string** |  | 
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


## WishListExistsHeadAsync

> EmptyEnvelope WishListExistsHeadAsync(ctx).WishListId(wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Check if a wish list exists (HEAD)



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
	wishListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishListsAPI.WishListExistsHeadAsync(context.Background()).WishListId(wishListId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishListsAPI.WishListExistsHeadAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishListExistsHeadAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WishListsAPI.WishListExistsHeadAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishListExistsHeadAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wishListId** | **string** |  | 
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

