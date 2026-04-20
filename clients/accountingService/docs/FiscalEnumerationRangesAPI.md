# \FiscalEnumerationRangesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoiceEnumerationRange**](FiscalEnumerationRangesAPI.md#CreateInvoiceEnumerationRange) | **Post** /api/v2/AccountingService/Fiscals/Authorities/EnumerationRanges | Create an invoice enumeration range
[**DeleteInvoiceEnumerationRange**](FiscalEnumerationRangesAPI.md#DeleteInvoiceEnumerationRange) | **Delete** /api/v2/AccountingService/Fiscals/Authorities/EnumerationRanges/{enumerationRangeId} | Delete an invoice enumeration range
[**GetInvoiceEnumerationRange**](FiscalEnumerationRangesAPI.md#GetInvoiceEnumerationRange) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/EnumerationRanges/{enumerationRangeId} | Get invoice enumeration range by ID
[**GetInvoiceEnumerationRanges**](FiscalEnumerationRangesAPI.md#GetInvoiceEnumerationRanges) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{authorityId}/EnumerationRanges | Get invoice enumeration ranges for an authority
[**GetInvoiceEnumerationRangesCount**](FiscalEnumerationRangesAPI.md#GetInvoiceEnumerationRangesCount) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/EnumerationRanges/Count | Get invoice enumeration ranges count
[**UpdateInvoiceEnumerationRange**](FiscalEnumerationRangesAPI.md#UpdateInvoiceEnumerationRange) | **Put** /api/v2/AccountingService/Fiscals/Authorities/EnumerationRanges/{enumerationRangeId} | Update an invoice enumeration range



## CreateInvoiceEnumerationRange

> EmptyEnvelope CreateInvoiceEnumerationRange(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InvoiceEnumerationRangeCreateDto(invoiceEnumerationRangeCreateDto).Execute()

Create an invoice enumeration range



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
	tenantId := map[string]interface{}{ ... } // map[string]interface{} | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	invoiceEnumerationRangeCreateDto := *openapiclient.NewInvoiceEnumerationRangeCreateDto(time.Now(), time.Now()) // InvoiceEnumerationRangeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalEnumerationRangesAPI.CreateInvoiceEnumerationRange(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InvoiceEnumerationRangeCreateDto(invoiceEnumerationRangeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalEnumerationRangesAPI.CreateInvoiceEnumerationRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceEnumerationRange`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalEnumerationRangesAPI.CreateInvoiceEnumerationRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceEnumerationRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
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


## DeleteInvoiceEnumerationRange

> EmptyEnvelope DeleteInvoiceEnumerationRange(ctx, enumerationRangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

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
	tenantId := map[string]interface{}{ ... } // map[string]interface{} | 
	enumerationRangeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalEnumerationRangesAPI.DeleteInvoiceEnumerationRange(context.Background(), enumerationRangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalEnumerationRangesAPI.DeleteInvoiceEnumerationRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInvoiceEnumerationRange`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalEnumerationRangesAPI.DeleteInvoiceEnumerationRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enumerationRangeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceEnumerationRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

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


## GetInvoiceEnumerationRange

> InvoiceEnumerationRangeDtoEnvelope GetInvoiceEnumerationRange(ctx, fiscalAuthorityId, enumerationRangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

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
	tenantId := map[string]interface{}{ ... } // map[string]interface{} | 
	fiscalAuthorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	enumerationRangeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalEnumerationRangesAPI.GetInvoiceEnumerationRange(context.Background(), fiscalAuthorityId, enumerationRangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalEnumerationRangesAPI.GetInvoiceEnumerationRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceEnumerationRange`: InvoiceEnumerationRangeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalEnumerationRangesAPI.GetInvoiceEnumerationRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**enumerationRangeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceEnumerationRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 


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


## GetInvoiceEnumerationRanges

> InvoiceEnumerationRangeDtoListEnvelope GetInvoiceEnumerationRanges(ctx, authorityId).FiscalAuthorityId(fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get invoice enumeration ranges for an authority



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
	fiscalAuthorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	authorityId := "authorityId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalEnumerationRangesAPI.GetInvoiceEnumerationRanges(context.Background(), authorityId).FiscalAuthorityId(fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalEnumerationRangesAPI.GetInvoiceEnumerationRanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceEnumerationRanges`: InvoiceEnumerationRangeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalEnumerationRangesAPI.GetInvoiceEnumerationRanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceEnumerationRangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fiscalAuthorityId** | **string** |  | 

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


## GetInvoiceEnumerationRangesCount

> Int32Envelope GetInvoiceEnumerationRangesCount(ctx, fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get invoice enumeration ranges count



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
	fiscalAuthorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalEnumerationRangesAPI.GetInvoiceEnumerationRangesCount(context.Background(), fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalEnumerationRangesAPI.GetInvoiceEnumerationRangesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceEnumerationRangesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalEnumerationRangesAPI.GetInvoiceEnumerationRangesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceEnumerationRangesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## UpdateInvoiceEnumerationRange

> EmptyEnvelope UpdateInvoiceEnumerationRange(ctx, enumerationRangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InvoiceEnumerationRangeUpdateDto(invoiceEnumerationRangeUpdateDto).Execute()

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
	tenantId := map[string]interface{}{ ... } // map[string]interface{} | 
	enumerationRangeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	invoiceEnumerationRangeUpdateDto := *openapiclient.NewInvoiceEnumerationRangeUpdateDto() // InvoiceEnumerationRangeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalEnumerationRangesAPI.UpdateInvoiceEnumerationRange(context.Background(), enumerationRangeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).InvoiceEnumerationRangeUpdateDto(invoiceEnumerationRangeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalEnumerationRangesAPI.UpdateInvoiceEnumerationRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoiceEnumerationRange`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalEnumerationRangesAPI.UpdateInvoiceEnumerationRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enumerationRangeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceEnumerationRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

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

