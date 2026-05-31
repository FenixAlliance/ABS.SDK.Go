# \CompareAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddItemToCompareTableAsync**](CompareAPI.md#AddItemToCompareTableAsync) | **Post** /api/v2/CartService/Compare | Add an item to the compare table
[**GetItemToCompareRecord**](CompareAPI.md#GetItemToCompareRecord) | **Get** /api/v2/CartService/Compare/{recordId}/Details | Get compare record details
[**GetItemToCompareRecords**](CompareAPI.md#GetItemToCompareRecords) | **Get** /api/v2/CartService/Compare/{cartId} | Get items to compare in a cart
[**RemoveItemFromCompareTable**](CompareAPI.md#RemoveItemFromCompareTable) | **Delete** /api/v2/CartService/Compare/{recordId} | Remove an item from the compare table



## AddItemToCompareTableAsync

> ItemCartRecordDto AddItemToCompareTableAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).AddProductToCompareRequest(addProductToCompareRequest).Execute()

Add an item to the compare table



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
	addProductToCompareRequest := *openapiclient.NewAddProductToCompareRequest() // AddProductToCompareRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompareAPI.AddItemToCompareTableAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).AddProductToCompareRequest(addProductToCompareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompareAPI.AddItemToCompareTableAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddItemToCompareTableAsync`: ItemCartRecordDto
	fmt.Fprintf(os.Stdout, "Response from `CompareAPI.AddItemToCompareTableAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddItemToCompareTableAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **addProductToCompareRequest** | [**AddProductToCompareRequest**](AddProductToCompareRequest.md) |  | 

### Return type

[**ItemCartRecordDto**](ItemCartRecordDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemToCompareRecord

> ItemToCompareCartRecordDtoEnvelope GetItemToCompareRecord(ctx, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get compare record details



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
	resp, r, err := apiClient.CompareAPI.GetItemToCompareRecord(context.Background(), recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompareAPI.GetItemToCompareRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemToCompareRecord`: ItemToCompareCartRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CompareAPI.GetItemToCompareRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemToCompareRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemToCompareCartRecordDtoEnvelope**](ItemToCompareCartRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemToCompareRecords

> ItemToCompareCartRecordDtoListEnvelope GetItemToCompareRecords(ctx, cartId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get items to compare in a cart



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
	resp, r, err := apiClient.CompareAPI.GetItemToCompareRecords(context.Background(), cartId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompareAPI.GetItemToCompareRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemToCompareRecords`: ItemToCompareCartRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CompareAPI.GetItemToCompareRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemToCompareRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemToCompareCartRecordDtoListEnvelope**](ItemToCompareCartRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveItemFromCompareTable

> ItemToCompareCartRecordDto RemoveItemFromCompareTable(ctx, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove an item from the compare table



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
	resp, r, err := apiClient.CompareAPI.RemoveItemFromCompareTable(context.Background(), recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompareAPI.RemoveItemFromCompareTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveItemFromCompareTable`: ItemToCompareCartRecordDto
	fmt.Fprintf(os.Stdout, "Response from `CompareAPI.RemoveItemFromCompareTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveItemFromCompareTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemToCompareCartRecordDto**](ItemToCompareCartRecordDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

