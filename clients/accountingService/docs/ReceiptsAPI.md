# \ReceiptsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReceiptAsync**](ReceiptsAPI.md#CreateReceiptAsync) | **Post** /api/v2/AccountingService/Receipts | Creates a new receipt
[**DeleteReceiptAsync**](ReceiptsAPI.md#DeleteReceiptAsync) | **Delete** /api/v2/AccountingService/Receipts/{receiptId} | Deletes a receipt
[**GetReceiptDetailsAsync**](ReceiptsAPI.md#GetReceiptDetailsAsync) | **Get** /api/v2/AccountingService/Receipts/{receiptId} | Gets details of a receipt
[**GetReceiptsAsync**](ReceiptsAPI.md#GetReceiptsAsync) | **Get** /api/v2/AccountingService/Receipts | Retrieves tenant receipts
[**GetReceiptsCountAsync**](ReceiptsAPI.md#GetReceiptsCountAsync) | **Get** /api/v2/AccountingService/Receipts/Count | Gets count of tenant receipts
[**UpdateReceiptAsync**](ReceiptsAPI.md#UpdateReceiptAsync) | **Put** /api/v2/AccountingService/Receipts/{receiptId} | Updates a receipt



## CreateReceiptAsync

> EmptyEnvelope CreateReceiptAsync(ctx).TenantId(tenantId).ReceiptCreateDto(receiptCreateDto).Execute()

Creates a new receipt



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
	receiptCreateDto := *openapiclient.NewReceiptCreateDto() // ReceiptCreateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiptsAPI.CreateReceiptAsync(context.Background()).TenantId(tenantId).ReceiptCreateDto(receiptCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.CreateReceiptAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReceiptAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.CreateReceiptAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **receiptCreateDto** | [**ReceiptCreateDto**](ReceiptCreateDto.md) |  | 

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


## DeleteReceiptAsync

> EmptyEnvelope DeleteReceiptAsync(ctx, receiptId).TenantId(tenantId).Execute()

Deletes a receipt



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
	receiptId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiptsAPI.DeleteReceiptAsync(context.Background(), receiptId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.DeleteReceiptAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteReceiptAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.DeleteReceiptAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReceiptAsyncRequest struct via the builder pattern


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


## GetReceiptDetailsAsync

> ReceiptDtoEnvelope GetReceiptDetailsAsync(ctx, receiptId).TenantId(tenantId).Execute()

Gets details of a receipt



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
	receiptId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiptsAPI.GetReceiptDetailsAsync(context.Background(), receiptId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.GetReceiptDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceiptDetailsAsync`: ReceiptDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.GetReceiptDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**ReceiptDtoEnvelope**](ReceiptDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptsAsync

> ReceiptDtoIReadOnlyListEnvelope GetReceiptsAsync(ctx).TenantId(tenantId).Execute()

Retrieves tenant receipts



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
	resp, r, err := apiClient.ReceiptsAPI.GetReceiptsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.GetReceiptsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceiptsAsync`: ReceiptDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.GetReceiptsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ReceiptDtoIReadOnlyListEnvelope**](ReceiptDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptsCountAsync

> Int32Envelope GetReceiptsCountAsync(ctx).TenantId(tenantId).Execute()

Gets count of tenant receipts



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
	resp, r, err := apiClient.ReceiptsAPI.GetReceiptsCountAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.GetReceiptsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReceiptsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.GetReceiptsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReceiptsCountAsyncRequest struct via the builder pattern


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


## UpdateReceiptAsync

> EmptyEnvelope UpdateReceiptAsync(ctx, receiptId).TenantId(tenantId).ReceiptUpdateDto(receiptUpdateDto).Execute()

Updates a receipt



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
	receiptId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	receiptUpdateDto := *openapiclient.NewReceiptUpdateDto() // ReceiptUpdateDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReceiptsAPI.UpdateReceiptAsync(context.Background(), receiptId).TenantId(tenantId).ReceiptUpdateDto(receiptUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReceiptsAPI.UpdateReceiptAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReceiptAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ReceiptsAPI.UpdateReceiptAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReceiptAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **receiptUpdateDto** | [**ReceiptUpdateDto**](ReceiptUpdateDto.md) |  | 

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

