# \BillableLineTaxesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillableLineTax**](BillableLineTaxesAPI.md#CreateBillableLineTax) | **Post** /api/v2/AccountingService/BillableLines/{billableLineId}/Taxes | Create a new tax for a billable line.
[**DeleteBillableLineTax**](BillableLineTaxesAPI.md#DeleteBillableLineTax) | **Delete** /api/v2/AccountingService/BillableLines/{billableLineId}/Taxes/{taxId} | Delete a tax from a billable line.
[**GetBillableLineTaxes**](BillableLineTaxesAPI.md#GetBillableLineTaxes) | **Get** /api/v2/AccountingService/BillableLines/{billableLineId}/Taxes | Get taxes for a billable line.
[**GetBillableLineTaxesCount**](BillableLineTaxesAPI.md#GetBillableLineTaxesCount) | **Get** /api/v2/AccountingService/BillableLines/{billableLineId}/Taxes/Count | Get the count of taxes for a billable line.
[**PatchBillableLineTaxAsync**](BillableLineTaxesAPI.md#PatchBillableLineTaxAsync) | **Patch** /api/v2/AccountingService/BillableLines/{billableLineId}/Taxes/{taxId} | Patch a billable line tax
[**UpdateBillableLineTax**](BillableLineTaxesAPI.md#UpdateBillableLineTax) | **Put** /api/v2/AccountingService/BillableLines/{billableLineId}/Taxes/{taxId} | Update a tax for a billable line.



## CreateBillableLineTax

> EmptyEnvelope CreateBillableLineTax(ctx, billableLineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppliedItemTaxRecordCreateDto(appliedItemTaxRecordCreateDto).Execute()

Create a new tax for a billable line.



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
	billableLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	appliedItemTaxRecordCreateDto := *openapiclient.NewAppliedItemTaxRecordCreateDto() // AppliedItemTaxRecordCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillableLineTaxesAPI.CreateBillableLineTax(context.Background(), billableLineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppliedItemTaxRecordCreateDto(appliedItemTaxRecordCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillableLineTaxesAPI.CreateBillableLineTax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillableLineTax`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillableLineTaxesAPI.CreateBillableLineTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billableLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillableLineTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appliedItemTaxRecordCreateDto** | [**AppliedItemTaxRecordCreateDto**](AppliedItemTaxRecordCreateDto.md) |  | 

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


## DeleteBillableLineTax

> EmptyEnvelope DeleteBillableLineTax(ctx, billableLineId, taxId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a tax from a billable line.



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
	billableLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	taxId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillableLineTaxesAPI.DeleteBillableLineTax(context.Background(), billableLineId, taxId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillableLineTaxesAPI.DeleteBillableLineTax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBillableLineTax`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillableLineTaxesAPI.DeleteBillableLineTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billableLineId** | **string** |  | 
**taxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillableLineTaxRequest struct via the builder pattern


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


## GetBillableLineTaxes

> AppliedItemTaxRecordDtoIReadOnlyListEnvelope GetBillableLineTaxes(ctx, billableLineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get taxes for a billable line.



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
	billableLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillableLineTaxesAPI.GetBillableLineTaxes(context.Background(), billableLineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillableLineTaxesAPI.GetBillableLineTaxes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillableLineTaxes`: AppliedItemTaxRecordDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillableLineTaxesAPI.GetBillableLineTaxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billableLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillableLineTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AppliedItemTaxRecordDtoIReadOnlyListEnvelope**](AppliedItemTaxRecordDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillableLineTaxesCount

> Int32Envelope GetBillableLineTaxesCount(ctx, billableLineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of taxes for a billable line.



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
	billableLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillableLineTaxesAPI.GetBillableLineTaxesCount(context.Background(), billableLineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillableLineTaxesAPI.GetBillableLineTaxesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillableLineTaxesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BillableLineTaxesAPI.GetBillableLineTaxesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billableLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillableLineTaxesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## PatchBillableLineTaxAsync

> EmptyEnvelope PatchBillableLineTaxAsync(ctx, billableLineId, taxId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a billable line tax



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
	billableLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	taxId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillableLineTaxesAPI.PatchBillableLineTaxAsync(context.Background(), billableLineId, taxId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillableLineTaxesAPI.PatchBillableLineTaxAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchBillableLineTaxAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillableLineTaxesAPI.PatchBillableLineTaxAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billableLineId** | **string** |  | 
**taxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBillableLineTaxAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## UpdateBillableLineTax

> EmptyEnvelope UpdateBillableLineTax(ctx, billableLineId, taxId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppliedItemTaxRecordUpdateDto(appliedItemTaxRecordUpdateDto).Execute()

Update a tax for a billable line.



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
	billableLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	taxId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	appliedItemTaxRecordUpdateDto := *openapiclient.NewAppliedItemTaxRecordUpdateDto() // AppliedItemTaxRecordUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillableLineTaxesAPI.UpdateBillableLineTax(context.Background(), billableLineId, taxId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppliedItemTaxRecordUpdateDto(appliedItemTaxRecordUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillableLineTaxesAPI.UpdateBillableLineTax``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBillableLineTax`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillableLineTaxesAPI.UpdateBillableLineTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billableLineId** | **string** |  | 
**taxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillableLineTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appliedItemTaxRecordUpdateDto** | [**AppliedItemTaxRecordUpdateDto**](AppliedItemTaxRecordUpdateDto.md) |  | 

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

