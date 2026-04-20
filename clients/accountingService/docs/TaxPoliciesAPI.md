# \TaxPoliciesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppliedTaxPolicyRecord**](TaxPoliciesAPI.md#CreateAppliedTaxPolicyRecord) | **Post** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/AppliedTaxPolicyRecords | Create an applied tax policy record
[**CreateItemTaxPolicyRecord**](TaxPoliciesAPI.md#CreateItemTaxPolicyRecord) | **Post** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/ItemTaxPolicyRecords | Create an item tax policy record
[**CreateTaxPolicy**](TaxPoliciesAPI.md#CreateTaxPolicy) | **Post** /api/v2/AccountingService/TaxPolicies | Create a tax policy
[**DeleteAppliedTaxPolicyRecord**](TaxPoliciesAPI.md#DeleteAppliedTaxPolicyRecord) | **Delete** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/AppliedTaxPolicyRecords/{appliedTaxPolicyRecordId} | Delete an applied tax policy record
[**DeleteItemTaxPolicyRecord**](TaxPoliciesAPI.md#DeleteItemTaxPolicyRecord) | **Delete** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/ItemTaxPolicyRecords/{itemTaxPolicyRecordId} | Delete an item tax policy record
[**DeleteTaxPolicy**](TaxPoliciesAPI.md#DeleteTaxPolicy) | **Delete** /api/v2/AccountingService/TaxPolicies/{id} | Delete a tax policy
[**GetAppliedTaxPolicyRecord**](TaxPoliciesAPI.md#GetAppliedTaxPolicyRecord) | **Get** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/AppliedTaxPolicyRecords/{appliedTaxPolicyRecordId} | Get applied tax policy record by ID
[**GetAppliedTaxPolicyRecords**](TaxPoliciesAPI.md#GetAppliedTaxPolicyRecords) | **Get** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/AppliedTaxPolicyRecords | Get applied tax policy records
[**GetAppliedTaxPolicyRecordsCount**](TaxPoliciesAPI.md#GetAppliedTaxPolicyRecordsCount) | **Get** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/AppliedTaxPolicyRecords/Count | Get applied tax policy records count
[**GetItemTaxPolicyRecord**](TaxPoliciesAPI.md#GetItemTaxPolicyRecord) | **Get** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/ItemTaxPolicyRecords/{itemTaxPolicyRecordId} | Get item tax policy record by ID
[**GetItemTaxPolicyRecords**](TaxPoliciesAPI.md#GetItemTaxPolicyRecords) | **Get** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/ItemTaxPolicyRecords | Get item tax policy records
[**GetTaxPolicies**](TaxPoliciesAPI.md#GetTaxPolicies) | **Get** /api/v2/AccountingService/TaxPolicies | Get all tax policies for a tenant
[**GetTaxPoliciesByAuthority**](TaxPoliciesAPI.md#GetTaxPoliciesByAuthority) | **Get** /api/v2/AccountingService/TaxPolicies/ByAuthority/{authorityId} | Get tax policies by fiscal authority
[**GetTaxPoliciesCount**](TaxPoliciesAPI.md#GetTaxPoliciesCount) | **Get** /api/v2/AccountingService/TaxPolicies/Count | Get tax policies count
[**GetTaxPolicy**](TaxPoliciesAPI.md#GetTaxPolicy) | **Get** /api/v2/AccountingService/TaxPolicies/{id} | Get tax policy by ID
[**UpdateAppliedTaxPolicyRecord**](TaxPoliciesAPI.md#UpdateAppliedTaxPolicyRecord) | **Put** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/AppliedTaxPolicyRecords/{appliedTaxPolicyRecordId} | Update an applied tax policy record
[**UpdateItemTaxPolicyRecord**](TaxPoliciesAPI.md#UpdateItemTaxPolicyRecord) | **Put** /api/v2/AccountingService/TaxPolicies/{taxPolicyId}/ItemTaxPolicyRecords/{itemTaxPolicyRecordId} | Update an item tax policy record
[**UpdateTaxPolicy**](TaxPoliciesAPI.md#UpdateTaxPolicy) | **Put** /api/v2/AccountingService/TaxPolicies/{id} | Update a tax policy



## CreateAppliedTaxPolicyRecord

> EmptyEnvelope CreateAppliedTaxPolicyRecord(ctx, taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppliedTaxPolicyRecordCreateDto(appliedTaxPolicyRecordCreateDto).Execute()

Create an applied tax policy record



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	appliedTaxPolicyRecordCreateDto := *openapiclient.NewAppliedTaxPolicyRecordCreateDto() // AppliedTaxPolicyRecordCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.CreateAppliedTaxPolicyRecord(context.Background(), taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppliedTaxPolicyRecordCreateDto(appliedTaxPolicyRecordCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.CreateAppliedTaxPolicyRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAppliedTaxPolicyRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.CreateAppliedTaxPolicyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppliedTaxPolicyRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appliedTaxPolicyRecordCreateDto** | [**AppliedTaxPolicyRecordCreateDto**](AppliedTaxPolicyRecordCreateDto.md) |  | 

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


## CreateItemTaxPolicyRecord

> EmptyEnvelope CreateItemTaxPolicyRecord(ctx, taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemTaxPolicyRecordCreateDto(itemTaxPolicyRecordCreateDto).Execute()

Create an item tax policy record



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemTaxPolicyRecordCreateDto := *openapiclient.NewItemTaxPolicyRecordCreateDto() // ItemTaxPolicyRecordCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.CreateItemTaxPolicyRecord(context.Background(), taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemTaxPolicyRecordCreateDto(itemTaxPolicyRecordCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.CreateItemTaxPolicyRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateItemTaxPolicyRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.CreateItemTaxPolicyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemTaxPolicyRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemTaxPolicyRecordCreateDto** | [**ItemTaxPolicyRecordCreateDto**](ItemTaxPolicyRecordCreateDto.md) |  | 

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


## CreateTaxPolicy

> EmptyEnvelope CreateTaxPolicy(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaxPolicyCreateDto(taxPolicyCreateDto).Execute()

Create a tax policy



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
	taxPolicyCreateDto := *openapiclient.NewTaxPolicyCreateDto() // TaxPolicyCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.CreateTaxPolicy(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaxPolicyCreateDto(taxPolicyCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.CreateTaxPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTaxPolicy`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.CreateTaxPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaxPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **taxPolicyCreateDto** | [**TaxPolicyCreateDto**](TaxPolicyCreateDto.md) |  | 

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


## DeleteAppliedTaxPolicyRecord

> EmptyEnvelope DeleteAppliedTaxPolicyRecord(ctx, taxPolicyId, appliedTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an applied tax policy record



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	appliedTaxPolicyRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.DeleteAppliedTaxPolicyRecord(context.Background(), taxPolicyId, appliedTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.DeleteAppliedTaxPolicyRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAppliedTaxPolicyRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.DeleteAppliedTaxPolicyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 
**appliedTaxPolicyRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppliedTaxPolicyRecordRequest struct via the builder pattern


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


## DeleteItemTaxPolicyRecord

> EmptyEnvelope DeleteItemTaxPolicyRecord(ctx, taxPolicyId, itemTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item tax policy record



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemTaxPolicyRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.DeleteItemTaxPolicyRecord(context.Background(), taxPolicyId, itemTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.DeleteItemTaxPolicyRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteItemTaxPolicyRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.DeleteItemTaxPolicyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 
**itemTaxPolicyRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemTaxPolicyRecordRequest struct via the builder pattern


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


## DeleteTaxPolicy

> EmptyEnvelope DeleteTaxPolicy(ctx, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a tax policy



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.DeleteTaxPolicy(context.Background(), id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.DeleteTaxPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTaxPolicy`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.DeleteTaxPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaxPolicyRequest struct via the builder pattern


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


## GetAppliedTaxPolicyRecord

> AppliedTaxPolicyRecordDtoEnvelope GetAppliedTaxPolicyRecord(ctx, taxPolicyId, appliedTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get applied tax policy record by ID



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	appliedTaxPolicyRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.GetAppliedTaxPolicyRecord(context.Background(), taxPolicyId, appliedTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetAppliedTaxPolicyRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppliedTaxPolicyRecord`: AppliedTaxPolicyRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetAppliedTaxPolicyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 
**appliedTaxPolicyRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedTaxPolicyRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AppliedTaxPolicyRecordDtoEnvelope**](AppliedTaxPolicyRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliedTaxPolicyRecords

> AppliedTaxPolicyRecordDtoListEnvelope GetAppliedTaxPolicyRecords(ctx, taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get applied tax policy records



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.GetAppliedTaxPolicyRecords(context.Background(), taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetAppliedTaxPolicyRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppliedTaxPolicyRecords`: AppliedTaxPolicyRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetAppliedTaxPolicyRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedTaxPolicyRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AppliedTaxPolicyRecordDtoListEnvelope**](AppliedTaxPolicyRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliedTaxPolicyRecordsCount

> Int32Envelope GetAppliedTaxPolicyRecordsCount(ctx, taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get applied tax policy records count



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.GetAppliedTaxPolicyRecordsCount(context.Background(), taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetAppliedTaxPolicyRecordsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppliedTaxPolicyRecordsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetAppliedTaxPolicyRecordsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedTaxPolicyRecordsCountRequest struct via the builder pattern


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


## GetItemTaxPolicyRecord

> ItemTaxPolicyRecordDtoEnvelope GetItemTaxPolicyRecord(ctx, taxPolicyId, itemTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item tax policy record by ID



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemTaxPolicyRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.GetItemTaxPolicyRecord(context.Background(), taxPolicyId, itemTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetItemTaxPolicyRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemTaxPolicyRecord`: ItemTaxPolicyRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetItemTaxPolicyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 
**itemTaxPolicyRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemTaxPolicyRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemTaxPolicyRecordDtoEnvelope**](ItemTaxPolicyRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemTaxPolicyRecords

> ItemTaxPolicyRecordDtoListEnvelope GetItemTaxPolicyRecords(ctx, taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item tax policy records



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.GetItemTaxPolicyRecords(context.Background(), taxPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetItemTaxPolicyRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemTaxPolicyRecords`: ItemTaxPolicyRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetItemTaxPolicyRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemTaxPolicyRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemTaxPolicyRecordDtoListEnvelope**](ItemTaxPolicyRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxPolicies

> TaxPolicyDtoListEnvelope GetTaxPolicies(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all tax policies for a tenant



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
	resp, r, err := apiClient.TaxPoliciesAPI.GetTaxPolicies(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetTaxPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxPolicies`: TaxPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetTaxPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaxPolicyDtoListEnvelope**](TaxPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxPoliciesByAuthority

> TaxPolicyDtoListEnvelope GetTaxPoliciesByAuthority(ctx, authorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get tax policies by fiscal authority



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
	authorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.GetTaxPoliciesByAuthority(context.Background(), authorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetTaxPoliciesByAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxPoliciesByAuthority`: TaxPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetTaxPoliciesByAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxPoliciesByAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaxPolicyDtoListEnvelope**](TaxPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxPoliciesCount

> Int32Envelope GetTaxPoliciesCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get tax policies count



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
	resp, r, err := apiClient.TaxPoliciesAPI.GetTaxPoliciesCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetTaxPoliciesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxPoliciesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetTaxPoliciesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxPoliciesCountRequest struct via the builder pattern


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


## GetTaxPolicy

> TaxPolicyDtoEnvelope GetTaxPolicy(ctx, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get tax policy by ID



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.GetTaxPolicy(context.Background(), id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.GetTaxPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaxPolicy`: TaxPolicyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.GetTaxPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TaxPolicyDtoEnvelope**](TaxPolicyDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppliedTaxPolicyRecord

> EmptyEnvelope UpdateAppliedTaxPolicyRecord(ctx, taxPolicyId, appliedTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppliedTaxPolicyRecordUpdateDto(appliedTaxPolicyRecordUpdateDto).Execute()

Update an applied tax policy record



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	appliedTaxPolicyRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	appliedTaxPolicyRecordUpdateDto := *openapiclient.NewAppliedTaxPolicyRecordUpdateDto() // AppliedTaxPolicyRecordUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.UpdateAppliedTaxPolicyRecord(context.Background(), taxPolicyId, appliedTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).AppliedTaxPolicyRecordUpdateDto(appliedTaxPolicyRecordUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.UpdateAppliedTaxPolicyRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAppliedTaxPolicyRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.UpdateAppliedTaxPolicyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 
**appliedTaxPolicyRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppliedTaxPolicyRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **appliedTaxPolicyRecordUpdateDto** | [**AppliedTaxPolicyRecordUpdateDto**](AppliedTaxPolicyRecordUpdateDto.md) |  | 

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


## UpdateItemTaxPolicyRecord

> EmptyEnvelope UpdateItemTaxPolicyRecord(ctx, taxPolicyId, itemTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemTaxPolicyRecordUpdateDto(itemTaxPolicyRecordUpdateDto).Execute()

Update an item tax policy record



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemTaxPolicyRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemTaxPolicyRecordUpdateDto := *openapiclient.NewItemTaxPolicyRecordUpdateDto() // ItemTaxPolicyRecordUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.UpdateItemTaxPolicyRecord(context.Background(), taxPolicyId, itemTaxPolicyRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemTaxPolicyRecordUpdateDto(itemTaxPolicyRecordUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.UpdateItemTaxPolicyRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemTaxPolicyRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.UpdateItemTaxPolicyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxPolicyId** | **string** |  | 
**itemTaxPolicyRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemTaxPolicyRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemTaxPolicyRecordUpdateDto** | [**ItemTaxPolicyRecordUpdateDto**](ItemTaxPolicyRecordUpdateDto.md) |  | 

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


## UpdateTaxPolicy

> EmptyEnvelope UpdateTaxPolicy(ctx, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaxPolicyUpdateDto(taxPolicyUpdateDto).Execute()

Update a tax policy



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	taxPolicyUpdateDto := *openapiclient.NewTaxPolicyUpdateDto() // TaxPolicyUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxPoliciesAPI.UpdateTaxPolicy(context.Background(), id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TaxPolicyUpdateDto(taxPolicyUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxPoliciesAPI.UpdateTaxPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTaxPolicy`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TaxPoliciesAPI.UpdateTaxPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaxPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **taxPolicyUpdateDto** | [**TaxPolicyUpdateDto**](TaxPolicyUpdateDto.md) |  | 

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

