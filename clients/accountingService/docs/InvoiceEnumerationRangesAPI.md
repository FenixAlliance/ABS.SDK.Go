# \InvoiceEnumerationRangesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoiceEnumerationRangeAsync**](InvoiceEnumerationRangesAPI.md#CreateInvoiceEnumerationRangeAsync) | **Post** /api/v2/AccountingService/InvoiceEnumerationRanges | Create a new invoice enumeration range
[**DeleteInvoiceEnumerationRangeAsync**](InvoiceEnumerationRangesAPI.md#DeleteInvoiceEnumerationRangeAsync) | **Delete** /api/v2/AccountingService/InvoiceEnumerationRanges/{rangeId} | Delete an invoice enumeration range
[**GetInvoiceEnumerationRangeDetailsAsync**](InvoiceEnumerationRangesAPI.md#GetInvoiceEnumerationRangeDetailsAsync) | **Get** /api/v2/AccountingService/InvoiceEnumerationRanges/{rangeId} | Get invoice enumeration range by ID
[**GetInvoiceEnumerationRangesAsync**](InvoiceEnumerationRangesAPI.md#GetInvoiceEnumerationRangesAsync) | **Get** /api/v2/AccountingService/InvoiceEnumerationRanges | Get all invoice enumeration ranges
[**UpdateInvoiceEnumerationRangeAsync**](InvoiceEnumerationRangesAPI.md#UpdateInvoiceEnumerationRangeAsync) | **Put** /api/v2/AccountingService/InvoiceEnumerationRanges/{rangeId} | Update an invoice enumeration range



## CreateInvoiceEnumerationRangeAsync

> EmptyEnvelope CreateInvoiceEnumerationRangeAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InvoiceEnumerationRangeCreateDto(invoiceEnumerationRangeCreateDto).Execute()

Create a new invoice enumeration range



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	invoiceEnumerationRangeCreateDto := *openapiclient.NewInvoiceEnumerationRangeCreateDto(time.Now(), time.Now()) // InvoiceEnumerationRangeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceEnumerationRangesAPI.CreateInvoiceEnumerationRangeAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InvoiceEnumerationRangeCreateDto(invoiceEnumerationRangeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEnumerationRangesAPI.CreateInvoiceEnumerationRangeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceEnumerationRangeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEnumerationRangesAPI.CreateInvoiceEnumerationRangeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceEnumerationRangeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **invoiceEnumerationRangeCreateDto** | [**InvoiceEnumerationRangeCreateDto**](InvoiceEnumerationRangeCreateDto.md) |  | 

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


## DeleteInvoiceEnumerationRangeAsync

> EmptyEnvelope DeleteInvoiceEnumerationRangeAsync(ctx, rangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an invoice enumeration range



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
	rangeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceEnumerationRangesAPI.DeleteInvoiceEnumerationRangeAsync(context.Background(), rangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEnumerationRangesAPI.DeleteInvoiceEnumerationRangeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInvoiceEnumerationRangeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEnumerationRangesAPI.DeleteInvoiceEnumerationRangeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rangeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceEnumerationRangeAsyncRequest struct via the builder pattern


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


## GetInvoiceEnumerationRangeDetailsAsync

> InvoiceEnumerationRangeDtoEnvelope GetInvoiceEnumerationRangeDetailsAsync(ctx, rangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get invoice enumeration range by ID



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
	rangeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceEnumerationRangesAPI.GetInvoiceEnumerationRangeDetailsAsync(context.Background(), rangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEnumerationRangesAPI.GetInvoiceEnumerationRangeDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceEnumerationRangeDetailsAsync`: InvoiceEnumerationRangeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEnumerationRangesAPI.GetInvoiceEnumerationRangeDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rangeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceEnumerationRangeDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InvoiceEnumerationRangeDtoEnvelope**](InvoiceEnumerationRangeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceEnumerationRangesAsync

> InvoiceEnumerationRangeDtoListEnvelope GetInvoiceEnumerationRangesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all invoice enumeration ranges



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
	resp, r, err := apiClient.InvoiceEnumerationRangesAPI.GetInvoiceEnumerationRangesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEnumerationRangesAPI.GetInvoiceEnumerationRangesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceEnumerationRangesAsync`: InvoiceEnumerationRangeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEnumerationRangesAPI.GetInvoiceEnumerationRangesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceEnumerationRangesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InvoiceEnumerationRangeDtoListEnvelope**](InvoiceEnumerationRangeDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoiceEnumerationRangeAsync

> EmptyEnvelope UpdateInvoiceEnumerationRangeAsync(ctx, rangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InvoiceEnumerationRangeUpdateDto(invoiceEnumerationRangeUpdateDto).Execute()

Update an invoice enumeration range



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
	rangeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	invoiceEnumerationRangeUpdateDto := *openapiclient.NewInvoiceEnumerationRangeUpdateDto() // InvoiceEnumerationRangeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceEnumerationRangesAPI.UpdateInvoiceEnumerationRangeAsync(context.Background(), rangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InvoiceEnumerationRangeUpdateDto(invoiceEnumerationRangeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceEnumerationRangesAPI.UpdateInvoiceEnumerationRangeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoiceEnumerationRangeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoiceEnumerationRangesAPI.UpdateInvoiceEnumerationRangeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rangeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceEnumerationRangeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **invoiceEnumerationRangeUpdateDto** | [**InvoiceEnumerationRangeUpdateDto**](InvoiceEnumerationRangeUpdateDto.md) |  | 

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

