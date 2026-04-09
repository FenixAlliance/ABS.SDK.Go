# \InvoicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateInvoiceDiscounts**](InvoicesAPI.md#AggregateInvoiceDiscounts) | **Post** /api/v2/InvoicingService/Invoices/DiscountsAggregate | Aggregate invoice discounts.
[**AggregateInvoiceGlobalSurcharges**](InvoicesAPI.md#AggregateInvoiceGlobalSurcharges) | **Post** /api/v2/InvoicingService/Invoices/GlobalSurchargesAggregate | Aggregate invoice global surcharges.
[**AggregateInvoiceTaxBases**](InvoicesAPI.md#AggregateInvoiceTaxBases) | **Post** /api/v2/InvoicingService/Invoices/TaxBasesAggregate | Aggregate invoice tax bases.
[**AggregateInvoiceTaxes**](InvoicesAPI.md#AggregateInvoiceTaxes) | **Post** /api/v2/InvoicingService/Invoices/TaxesAggregate | Aggregate invoice taxes.
[**AggregateInvoiceTotals**](InvoicesAPI.md#AggregateInvoiceTotals) | **Post** /api/v2/InvoicingService/Invoices/TotalsAggregate | Aggregate invoice totals.
[**CalculateInvoice**](InvoicesAPI.md#CalculateInvoice) | **Put** /api/v2/InvoicingService/Invoices/{invoiceId}/Calculate | Calculate an invoice.
[**CalculateInvoiceLine**](InvoicesAPI.md#CalculateInvoiceLine) | **Put** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId}/Calculate | Calculate an invoice line.
[**CreateInvoice**](InvoicesAPI.md#CreateInvoice) | **Post** /api/v2/InvoicingService/Invoices | Create a new invoice.
[**CreateInvoiceAdjustment**](InvoicesAPI.md#CreateInvoiceAdjustment) | **Post** /api/v2/InvoicingService/Invoices/{invoiceId}/Adjustments | Create a new invoice adjustment.
[**CreateInvoiceLine**](InvoicesAPI.md#CreateInvoiceLine) | **Post** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines | Create a new invoice line.
[**CreateInvoiceLineTax**](InvoicesAPI.md#CreateInvoiceLineTax) | **Post** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId}/Taxes | Create a new tax for an invoice line.
[**CreateInvoiceReference**](InvoicesAPI.md#CreateInvoiceReference) | **Post** /api/v2/InvoicingService/Invoices/{invoiceId}/References | Create a new invoice reference.
[**DeleteInvoice**](InvoicesAPI.md#DeleteInvoice) | **Delete** /api/v2/InvoicingService/Invoices/{invoiceId} | Delete an invoice.
[**DeleteInvoiceAdjustment**](InvoicesAPI.md#DeleteInvoiceAdjustment) | **Delete** /api/v2/InvoicingService/Invoices/{invoiceId}/Adjustments/{invoiceAdjustmentId} | Delete an invoice adjustment.
[**DeleteInvoiceLine**](InvoicesAPI.md#DeleteInvoiceLine) | **Delete** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId} | Delete an invoice line.
[**DeleteInvoiceLineTax**](InvoicesAPI.md#DeleteInvoiceLineTax) | **Delete** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId}/Taxes/{invoiceLineTaxId} | Delete a tax from an invoice line.
[**DeleteInvoiceReference**](InvoicesAPI.md#DeleteInvoiceReference) | **Delete** /api/v2/InvoicingService/Invoices/{invoiceId}/References/{invoiceReferenceId} | Delete an invoice reference.
[**GetExtendedInvoice**](InvoicesAPI.md#GetExtendedInvoice) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Extended | Get an extended invoice by ID.
[**GetExtendedInvoices**](InvoicesAPI.md#GetExtendedInvoices) | **Get** /api/v2/InvoicingService/Invoices/Extended | Get a list of extended invoices.
[**GetExtendedInvoicesCount**](InvoicesAPI.md#GetExtendedInvoicesCount) | **Get** /api/v2/InvoicingService/Invoices/Extended/Count | Get the count of extended invoices.
[**GetInvoice**](InvoicesAPI.md#GetInvoice) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId} | Get an invoice by ID.
[**GetInvoiceAdjustment**](InvoicesAPI.md#GetInvoiceAdjustment) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Adjustments/{invoiceAdjustmentId} | Get an invoice adjustment by ID.
[**GetInvoiceAdjustments**](InvoicesAPI.md#GetInvoiceAdjustments) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Adjustments | Get invoice adjustments.
[**GetInvoiceAdjustmentsCount**](InvoicesAPI.md#GetInvoiceAdjustmentsCount) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Adjustments/Count | Get the count of invoice adjustments.
[**GetInvoiceLine**](InvoicesAPI.md#GetInvoiceLine) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId} | Get an invoice line by ID.
[**GetInvoiceLineTaxes**](InvoicesAPI.md#GetInvoiceLineTaxes) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId}/Taxes | Get taxes for an invoice line.
[**GetInvoiceLineTaxesCount**](InvoicesAPI.md#GetInvoiceLineTaxesCount) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId}/Taxes/Count | Get the count of taxes for an invoice line.
[**GetInvoiceLines**](InvoicesAPI.md#GetInvoiceLines) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines | Get invoice lines.
[**GetInvoiceLinesCount**](InvoicesAPI.md#GetInvoiceLinesCount) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/Count | Get the count of invoice lines.
[**GetInvoicePayments**](InvoicesAPI.md#GetInvoicePayments) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Payments | Get payments for an invoice.
[**GetInvoicePaymentsCount**](InvoicesAPI.md#GetInvoicePaymentsCount) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/Payments/Count | Get the count of payments for an invoice.
[**GetInvoiceReference**](InvoicesAPI.md#GetInvoiceReference) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/References/{invoiceReferenceId} | Get an invoice reference by ID.
[**GetInvoiceReferences**](InvoicesAPI.md#GetInvoiceReferences) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/References | Get invoice references.
[**GetInvoiceReferencesCount**](InvoicesAPI.md#GetInvoiceReferencesCount) | **Get** /api/v2/InvoicingService/Invoices/{invoiceId}/References/Count | Get the count of invoice references.
[**GetInvoices**](InvoicesAPI.md#GetInvoices) | **Get** /api/v2/InvoicingService/Invoices | Get a list of invoices.
[**GetInvoicesCount**](InvoicesAPI.md#GetInvoicesCount) | **Get** /api/v2/InvoicingService/Invoices/Count | Get the count of invoices.
[**PreviewInvoiceEmail**](InvoicesAPI.md#PreviewInvoiceEmail) | **Post** /api/v2/InvoicingService/Invoices/{invoiceId}/Emails/Preview | Preview the rendered email for an invoice.
[**SendInvoiceEmail**](InvoicesAPI.md#SendInvoiceEmail) | **Post** /api/v2/InvoicingService/Invoices/{invoiceId}/Emails/Send | Send an invoice transactional email to recipients.
[**UpdateInvoice**](InvoicesAPI.md#UpdateInvoice) | **Put** /api/v2/InvoicingService/Invoices/{invoiceId} | Update an invoice.
[**UpdateInvoiceAdjustment**](InvoicesAPI.md#UpdateInvoiceAdjustment) | **Put** /api/v2/InvoicingService/Invoices/{invoiceId}/Adjustments/{invoiceAdjustmentId} | Update an invoice adjustment.
[**UpdateInvoiceLine**](InvoicesAPI.md#UpdateInvoiceLine) | **Put** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId} | Update an invoice line.
[**UpdateInvoiceLineTax**](InvoicesAPI.md#UpdateInvoiceLineTax) | **Put** /api/v2/InvoicingService/Invoices/{invoiceId}/Lines/{invoiceLineId}/Taxes/{invoiceLineTaxId} | Update a tax for an invoice line.
[**UpdateInvoiceReference**](InvoicesAPI.md#UpdateInvoiceReference) | **Put** /api/v2/InvoicingService/Invoices/{invoiceId}/References/{invoiceReferenceId} | Update an invoice reference.



## AggregateInvoiceDiscounts

> MoneyEnvelope AggregateInvoiceDiscounts(ctx).RequestBody(requestBody).CurrencyId(currencyId).Execute()

Aggregate invoice discounts.



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
	requestBody := []string{"Property_example"} // []string | 
	currencyId := "currencyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.AggregateInvoiceDiscounts(context.Background()).RequestBody(requestBody).CurrencyId(currencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.AggregateInvoiceDiscounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateInvoiceDiscounts`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.AggregateInvoiceDiscounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateInvoiceDiscountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **currencyId** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateInvoiceGlobalSurcharges

> MoneyEnvelope AggregateInvoiceGlobalSurcharges(ctx).RequestBody(requestBody).CurrencyId(currencyId).Execute()

Aggregate invoice global surcharges.



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
	requestBody := []string{"Property_example"} // []string | 
	currencyId := "currencyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.AggregateInvoiceGlobalSurcharges(context.Background()).RequestBody(requestBody).CurrencyId(currencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.AggregateInvoiceGlobalSurcharges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateInvoiceGlobalSurcharges`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.AggregateInvoiceGlobalSurcharges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateInvoiceGlobalSurchargesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **currencyId** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateInvoiceTaxBases

> MoneyEnvelope AggregateInvoiceTaxBases(ctx).RequestBody(requestBody).CurrencyId(currencyId).Execute()

Aggregate invoice tax bases.



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
	requestBody := []string{"Property_example"} // []string | 
	currencyId := "currencyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.AggregateInvoiceTaxBases(context.Background()).RequestBody(requestBody).CurrencyId(currencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.AggregateInvoiceTaxBases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateInvoiceTaxBases`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.AggregateInvoiceTaxBases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateInvoiceTaxBasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **currencyId** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateInvoiceTaxes

> MoneyEnvelope AggregateInvoiceTaxes(ctx).RequestBody(requestBody).CurrencyId(currencyId).Execute()

Aggregate invoice taxes.



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
	requestBody := []string{"Property_example"} // []string | 
	currencyId := "currencyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.AggregateInvoiceTaxes(context.Background()).RequestBody(requestBody).CurrencyId(currencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.AggregateInvoiceTaxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateInvoiceTaxes`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.AggregateInvoiceTaxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateInvoiceTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **currencyId** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AggregateInvoiceTotals

> MoneyEnvelope AggregateInvoiceTotals(ctx).RequestBody(requestBody).CurrencyId(currencyId).Execute()

Aggregate invoice totals.



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
	requestBody := []string{"Property_example"} // []string | 
	currencyId := "currencyId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.AggregateInvoiceTotals(context.Background()).RequestBody(requestBody).CurrencyId(currencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.AggregateInvoiceTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AggregateInvoiceTotals`: MoneyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.AggregateInvoiceTotals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAggregateInvoiceTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** |  | 
 **currencyId** | **string** |  | 

### Return type

[**MoneyEnvelope**](MoneyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CalculateInvoice

> EmptyEnvelope CalculateInvoice(ctx, invoiceId).TenantId(tenantId).Execute()

Calculate an invoice.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.CalculateInvoice(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CalculateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateInvoice`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CalculateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalculateInvoiceRequest struct via the builder pattern


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


## CalculateInvoiceLine

> EmptyEnvelope CalculateInvoiceLine(ctx, invoiceId, invoiceLineId).TenantId(tenantId).Execute()

Calculate an invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.CalculateInvoiceLine(context.Background(), invoiceId, invoiceLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CalculateInvoiceLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateInvoiceLine`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CalculateInvoiceLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalculateInvoiceLineRequest struct via the builder pattern


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


## CreateInvoice

> EmptyEnvelope CreateInvoice(ctx).TenantId(tenantId).InvoiceCreateDto(invoiceCreateDto).Execute()

Create a new invoice.



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
	invoiceCreateDto := *openapiclient.NewInvoiceCreateDto() // InvoiceCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.CreateInvoice(context.Background()).TenantId(tenantId).InvoiceCreateDto(invoiceCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoice`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **invoiceCreateDto** | [**InvoiceCreateDto**](InvoiceCreateDto.md) |  | 

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


## CreateInvoiceAdjustment

> EmptyEnvelope CreateInvoiceAdjustment(ctx, invoiceId).TenantId(tenantId).InvoiceAdjustmentCreateDto(invoiceAdjustmentCreateDto).Execute()

Create a new invoice adjustment.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceAdjustmentCreateDto := *openapiclient.NewInvoiceAdjustmentCreateDto() // InvoiceAdjustmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.CreateInvoiceAdjustment(context.Background(), invoiceId).TenantId(tenantId).InvoiceAdjustmentCreateDto(invoiceAdjustmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateInvoiceAdjustment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceAdjustment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateInvoiceAdjustment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceAdjustmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **invoiceAdjustmentCreateDto** | [**InvoiceAdjustmentCreateDto**](InvoiceAdjustmentCreateDto.md) |  | 

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


## CreateInvoiceLine

> InvoiceLineDtoIReadOnlyListEnvelope CreateInvoiceLine(ctx, invoiceId).TenantId(tenantId).InvoiceLineCreateDto(invoiceLineCreateDto).Execute()

Create a new invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineCreateDto := *openapiclient.NewInvoiceLineCreateDto() // InvoiceLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.CreateInvoiceLine(context.Background(), invoiceId).TenantId(tenantId).InvoiceLineCreateDto(invoiceLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateInvoiceLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceLine`: InvoiceLineDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateInvoiceLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **invoiceLineCreateDto** | [**InvoiceLineCreateDto**](InvoiceLineCreateDto.md) |  | 

### Return type

[**InvoiceLineDtoIReadOnlyListEnvelope**](InvoiceLineDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceLineTax

> EmptyEnvelope CreateInvoiceLineTax(ctx, invoiceId, invoiceLineId).TenantId(tenantId).InvoiceLineAppliedTaxCreateDto(invoiceLineAppliedTaxCreateDto).Execute()

Create a new tax for an invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineAppliedTaxCreateDto := *openapiclient.NewInvoiceLineAppliedTaxCreateDto() // InvoiceLineAppliedTaxCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.CreateInvoiceLineTax(context.Background(), invoiceId, invoiceLineId).TenantId(tenantId).InvoiceLineAppliedTaxCreateDto(invoiceLineAppliedTaxCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateInvoiceLineTax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceLineTax`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateInvoiceLineTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceLineTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **invoiceLineAppliedTaxCreateDto** | [**InvoiceLineAppliedTaxCreateDto**](InvoiceLineAppliedTaxCreateDto.md) |  | 

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


## CreateInvoiceReference

> EmptyEnvelope CreateInvoiceReference(ctx, invoiceId).TenantId(tenantId).InvoiceReferenceCreateDto(invoiceReferenceCreateDto).Execute()

Create a new invoice reference.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceReferenceCreateDto := *openapiclient.NewInvoiceReferenceCreateDto() // InvoiceReferenceCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.CreateInvoiceReference(context.Background(), invoiceId).TenantId(tenantId).InvoiceReferenceCreateDto(invoiceReferenceCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateInvoiceReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvoiceReference`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateInvoiceReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **invoiceReferenceCreateDto** | [**InvoiceReferenceCreateDto**](InvoiceReferenceCreateDto.md) |  | 

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


## DeleteInvoice

> EmptyEnvelope DeleteInvoice(ctx, invoiceId).TenantId(tenantId).Execute()

Delete an invoice.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.DeleteInvoice(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.DeleteInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInvoice`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.DeleteInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceRequest struct via the builder pattern


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


## DeleteInvoiceAdjustment

> EmptyEnvelope DeleteInvoiceAdjustment(ctx, invoiceId, invoiceAdjustmentId).TenantId(tenantId).Execute()

Delete an invoice adjustment.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceAdjustmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.DeleteInvoiceAdjustment(context.Background(), invoiceId, invoiceAdjustmentId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.DeleteInvoiceAdjustment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInvoiceAdjustment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.DeleteInvoiceAdjustment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceAdjustmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceAdjustmentRequest struct via the builder pattern


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


## DeleteInvoiceLine

> EmptyEnvelope DeleteInvoiceLine(ctx, invoiceId, invoiceLineId).TenantId(tenantId).Execute()

Delete an invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.DeleteInvoiceLine(context.Background(), invoiceId, invoiceLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.DeleteInvoiceLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInvoiceLine`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.DeleteInvoiceLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceLineRequest struct via the builder pattern


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


## DeleteInvoiceLineTax

> EmptyEnvelope DeleteInvoiceLineTax(ctx, invoiceId, invoiceLineId, invoiceLineTaxId).TenantId(tenantId).Execute()

Delete a tax from an invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineTaxId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.DeleteInvoiceLineTax(context.Background(), invoiceId, invoiceLineId, invoiceLineTaxId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.DeleteInvoiceLineTax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInvoiceLineTax`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.DeleteInvoiceLineTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 
**invoiceLineTaxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceLineTaxRequest struct via the builder pattern


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


## DeleteInvoiceReference

> EmptyEnvelope DeleteInvoiceReference(ctx, invoiceId, invoiceReferenceId).TenantId(tenantId).Execute()

Delete an invoice reference.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceReferenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.DeleteInvoiceReference(context.Background(), invoiceId, invoiceReferenceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.DeleteInvoiceReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInvoiceReference`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.DeleteInvoiceReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceReferenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceReferenceRequest struct via the builder pattern


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


## GetExtendedInvoice

> InvoiceDtoEnvelope GetExtendedInvoice(ctx, invoiceId).TenantId(tenantId).Execute()

Get an extended invoice by ID.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetExtendedInvoice(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetExtendedInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedInvoice`: InvoiceDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetExtendedInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**InvoiceDtoEnvelope**](InvoiceDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtendedInvoices

> ExtendedInvoiceDtoListEnvelope GetExtendedInvoices(ctx).TenantId(tenantId).Execute()

Get a list of extended invoices.



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
	resp, r, err := apiClient.InvoicesAPI.GetExtendedInvoices(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetExtendedInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedInvoices`: ExtendedInvoiceDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetExtendedInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ExtendedInvoiceDtoListEnvelope**](ExtendedInvoiceDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtendedInvoicesCount

> Int32Envelope GetExtendedInvoicesCount(ctx).TenantId(tenantId).Execute()

Get the count of extended invoices.



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
	resp, r, err := apiClient.InvoicesAPI.GetExtendedInvoicesCount(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetExtendedInvoicesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedInvoicesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetExtendedInvoicesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedInvoicesCountRequest struct via the builder pattern


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


## GetInvoice

> InvoiceDtoEnvelope GetInvoice(ctx, invoiceId).TenantId(tenantId).Execute()

Get an invoice by ID.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoice(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoice`: InvoiceDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**InvoiceDtoEnvelope**](InvoiceDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAdjustment

> InvoiceAdjustmentDtoEnvelope GetInvoiceAdjustment(ctx, invoiceId, invoiceAdjustmentId).TenantId(tenantId).Execute()

Get an invoice adjustment by ID.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceAdjustmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceAdjustment(context.Background(), invoiceId, invoiceAdjustmentId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceAdjustment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceAdjustment`: InvoiceAdjustmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceAdjustment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceAdjustmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceAdjustmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**InvoiceAdjustmentDtoEnvelope**](InvoiceAdjustmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAdjustments

> InvoiceAdjustmentDtoIReadOnlyListEnvelope GetInvoiceAdjustments(ctx, invoiceId).TenantId(tenantId).Execute()

Get invoice adjustments.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceAdjustments(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceAdjustments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceAdjustments`: InvoiceAdjustmentDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceAdjustments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceAdjustmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**InvoiceAdjustmentDtoIReadOnlyListEnvelope**](InvoiceAdjustmentDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAdjustmentsCount

> Int32Envelope GetInvoiceAdjustmentsCount(ctx, invoiceId).TenantId(tenantId).Execute()

Get the count of invoice adjustments.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceAdjustmentsCount(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceAdjustmentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceAdjustmentsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceAdjustmentsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceAdjustmentsCountRequest struct via the builder pattern


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


## GetInvoiceLine

> InvoiceLineDtoEnvelope GetInvoiceLine(ctx, invoiceId, invoiceLineId).TenantId(tenantId).Execute()

Get an invoice line by ID.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceLine(context.Background(), invoiceId, invoiceLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceLine`: InvoiceLineDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**InvoiceLineDtoEnvelope**](InvoiceLineDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceLineTaxes

> InvoiceLineAppliedTaxDtoIReadOnlyListEnvelope GetInvoiceLineTaxes(ctx, invoiceId, invoiceLineId).TenantId(tenantId).Execute()

Get taxes for an invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceLineTaxes(context.Background(), invoiceId, invoiceLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceLineTaxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceLineTaxes`: InvoiceLineAppliedTaxDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceLineTaxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceLineTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**InvoiceLineAppliedTaxDtoIReadOnlyListEnvelope**](InvoiceLineAppliedTaxDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceLineTaxesCount

> Int32Envelope GetInvoiceLineTaxesCount(ctx, invoiceId, invoiceLineId).TenantId(tenantId).Execute()

Get the count of taxes for an invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceLineTaxesCount(context.Background(), invoiceId, invoiceLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceLineTaxesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceLineTaxesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceLineTaxesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceLineTaxesCountRequest struct via the builder pattern


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


## GetInvoiceLines

> InvoiceLineDtoListEnvelope GetInvoiceLines(ctx, invoiceId).TenantId(tenantId).ItemId(itemId).Execute()

Get invoice lines.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceLines(context.Background(), invoiceId).TenantId(tenantId).ItemId(itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceLines`: InvoiceLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **itemId** | **string** |  | 

### Return type

[**InvoiceLineDtoListEnvelope**](InvoiceLineDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceLinesCount

> Int32Envelope GetInvoiceLinesCount(ctx, invoiceId).TenantId(tenantId).Execute()

Get the count of invoice lines.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceLinesCount(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceLinesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceLinesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceLinesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceLinesCountRequest struct via the builder pattern


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


## GetInvoicePayments

> PaymentDtoIReadOnlyListEnvelope GetInvoicePayments(ctx, invoiceId).Execute()

Get payments for an invoice.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoicePayments(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoicePayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoicePayments`: PaymentDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoicePayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentDtoIReadOnlyListEnvelope**](PaymentDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoicePaymentsCount

> Int32Envelope GetInvoicePaymentsCount(ctx, invoiceId).Execute()

Get the count of payments for an invoice.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoicePaymentsCount(context.Background(), invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoicePaymentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoicePaymentsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoicePaymentsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicePaymentsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetInvoiceReference

> InvoiceReferenceDtoEnvelope GetInvoiceReference(ctx, invoiceId, invoiceReferenceId).TenantId(tenantId).Execute()

Get an invoice reference by ID.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceReferenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceReference(context.Background(), invoiceId, invoiceReferenceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceReference`: InvoiceReferenceDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceReferenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**InvoiceReferenceDtoEnvelope**](InvoiceReferenceDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceReferences

> InvoiceReferenceDtoIReadOnlyListEnvelope GetInvoiceReferences(ctx, invoiceId).TenantId(tenantId).Execute()

Get invoice references.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceReferences(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceReferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceReferences`: InvoiceReferenceDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**InvoiceReferenceDtoIReadOnlyListEnvelope**](InvoiceReferenceDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceReferencesCount

> Int32Envelope GetInvoiceReferencesCount(ctx, invoiceId).TenantId(tenantId).Execute()

Get the count of invoice references.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoiceReferencesCount(context.Background(), invoiceId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceReferencesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoiceReferencesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceReferencesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceReferencesCountRequest struct via the builder pattern


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


## GetInvoices

> InvoiceDtoListEnvelope GetInvoices(ctx).TenantId(tenantId).Execute()

Get a list of invoices.



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
	resp, r, err := apiClient.InvoicesAPI.GetInvoices(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoices`: InvoiceDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**InvoiceDtoListEnvelope**](InvoiceDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoicesCount

> Int32Envelope GetInvoicesCount(ctx).TenantId(tenantId).Execute()

Get the count of invoices.



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
	resp, r, err := apiClient.InvoicesAPI.GetInvoicesCount(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoicesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoicesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoicesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesCountRequest struct via the builder pattern


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


## PreviewInvoiceEmail

> PreviewInvoiceEmail(ctx, invoiceId).TenantId(tenantId).EmailDispatchRequest(emailDispatchRequest).Execute()

Preview the rendered email for an invoice.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	emailDispatchRequest := *openapiclient.NewEmailDispatchRequest("Title_example", "Message_example", "Culture_example", "UiCulture_example", []string{"Recipients_example"}) // EmailDispatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoicesAPI.PreviewInvoiceEmail(context.Background(), invoiceId).TenantId(tenantId).EmailDispatchRequest(emailDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.PreviewInvoiceEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewInvoiceEmailRequest struct via the builder pattern


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


## SendInvoiceEmail

> Envelope SendInvoiceEmail(ctx, invoiceId).TenantId(tenantId).EmailDispatchRequest(emailDispatchRequest).Execute()

Send an invoice transactional email to recipients.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	emailDispatchRequest := *openapiclient.NewEmailDispatchRequest("Title_example", "Message_example", "Culture_example", "UiCulture_example", []string{"Recipients_example"}) // EmailDispatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.SendInvoiceEmail(context.Background(), invoiceId).TenantId(tenantId).EmailDispatchRequest(emailDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.SendInvoiceEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendInvoiceEmail`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.SendInvoiceEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendInvoiceEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **emailDispatchRequest** | [**EmailDispatchRequest**](EmailDispatchRequest.md) |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> EmptyEnvelope UpdateInvoice(ctx, invoiceId).TenantId(tenantId).InvoiceUpdateDto(invoiceUpdateDto).Execute()

Update an invoice.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceUpdateDto := *openapiclient.NewInvoiceUpdateDto() // InvoiceUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.UpdateInvoice(context.Background(), invoiceId).TenantId(tenantId).InvoiceUpdateDto(invoiceUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.UpdateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoice`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.UpdateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **invoiceUpdateDto** | [**InvoiceUpdateDto**](InvoiceUpdateDto.md) |  | 

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


## UpdateInvoiceAdjustment

> EmptyEnvelope UpdateInvoiceAdjustment(ctx, invoiceId, invoiceAdjustmentId).TenantId(tenantId).InvoiceAdjustmentUpdateDto(invoiceAdjustmentUpdateDto).Execute()

Update an invoice adjustment.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceAdjustmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceAdjustmentUpdateDto := *openapiclient.NewInvoiceAdjustmentUpdateDto() // InvoiceAdjustmentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.UpdateInvoiceAdjustment(context.Background(), invoiceId, invoiceAdjustmentId).TenantId(tenantId).InvoiceAdjustmentUpdateDto(invoiceAdjustmentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.UpdateInvoiceAdjustment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoiceAdjustment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.UpdateInvoiceAdjustment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceAdjustmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceAdjustmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **invoiceAdjustmentUpdateDto** | [**InvoiceAdjustmentUpdateDto**](InvoiceAdjustmentUpdateDto.md) |  | 

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


## UpdateInvoiceLine

> InvoiceLineDtoEnvelope UpdateInvoiceLine(ctx, invoiceId, invoiceLineId).TenantId(tenantId).InvoiceLineUpdateDto(invoiceLineUpdateDto).Execute()

Update an invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineUpdateDto := *openapiclient.NewInvoiceLineUpdateDto() // InvoiceLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.UpdateInvoiceLine(context.Background(), invoiceId, invoiceLineId).TenantId(tenantId).InvoiceLineUpdateDto(invoiceLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.UpdateInvoiceLine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoiceLine`: InvoiceLineDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.UpdateInvoiceLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **invoiceLineUpdateDto** | [**InvoiceLineUpdateDto**](InvoiceLineUpdateDto.md) |  | 

### Return type

[**InvoiceLineDtoEnvelope**](InvoiceLineDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoiceLineTax

> EmptyEnvelope UpdateInvoiceLineTax(ctx, invoiceId, invoiceLineId, invoiceLineTaxId).TenantId(tenantId).InvoiceLineAppliedTaxUpdateDto(invoiceLineAppliedTaxUpdateDto).Execute()

Update a tax for an invoice line.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineTaxId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceLineAppliedTaxUpdateDto := *openapiclient.NewInvoiceLineAppliedTaxUpdateDto() // InvoiceLineAppliedTaxUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.UpdateInvoiceLineTax(context.Background(), invoiceId, invoiceLineId, invoiceLineTaxId).TenantId(tenantId).InvoiceLineAppliedTaxUpdateDto(invoiceLineAppliedTaxUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.UpdateInvoiceLineTax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoiceLineTax`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.UpdateInvoiceLineTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceLineId** | **string** |  | 
**invoiceLineTaxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceLineTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



 **invoiceLineAppliedTaxUpdateDto** | [**InvoiceLineAppliedTaxUpdateDto**](InvoiceLineAppliedTaxUpdateDto.md) |  | 

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


## UpdateInvoiceReference

> EmptyEnvelope UpdateInvoiceReference(ctx, invoiceId, invoiceReferenceId).TenantId(tenantId).InvoiceReferenceUpdateDto(invoiceReferenceUpdateDto).Execute()

Update an invoice reference.



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceReferenceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoiceReferenceUpdateDto := *openapiclient.NewInvoiceReferenceUpdateDto() // InvoiceReferenceUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.UpdateInvoiceReference(context.Background(), invoiceId, invoiceReferenceId).TenantId(tenantId).InvoiceReferenceUpdateDto(invoiceReferenceUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.UpdateInvoiceReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoiceReference`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.UpdateInvoiceReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 
**invoiceReferenceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **invoiceReferenceUpdateDto** | [**InvoiceReferenceUpdateDto**](InvoiceReferenceUpdateDto.md) |  | 

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

