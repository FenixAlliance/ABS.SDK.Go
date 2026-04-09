# \FiscalResponsibilityRecordsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFiscalResponsibilityRecord**](FiscalResponsibilityRecordsAPI.md#CreateFiscalResponsibilityRecord) | **Post** /api/v2/AccountingService/Fiscals/Authorities/FiscalResponsibilityRecords | Create a fiscal responsibility record
[**DeleteFiscalResponsibilityRecord**](FiscalResponsibilityRecordsAPI.md#DeleteFiscalResponsibilityRecord) | **Delete** /api/v2/AccountingService/Fiscals/Authorities/FiscalResponsibilityRecords/{fiscalResponsibilityRecordId} | Delete a fiscal responsibility record
[**GetFiscalResponsibilityRecord**](FiscalResponsibilityRecordsAPI.md#GetFiscalResponsibilityRecord) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalResponsibilities/{fiscalResponsibilityId}/FiscalResponsibilityRecords/{fiscalResponsibilityRecordId} | Get fiscal responsibility record by ID
[**GetFiscalResponsibilityRecords**](FiscalResponsibilityRecordsAPI.md#GetFiscalResponsibilityRecords) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalResponsibilities/{fiscalResponsibilityId}/FiscalResponsibilityRecords | Get fiscal responsibility records
[**GetFiscalResponsibilityRecordsCount**](FiscalResponsibilityRecordsAPI.md#GetFiscalResponsibilityRecordsCount) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalResponsibilities/{fiscalResponsibilityId}/FiscalResponsibilityRecords/Count | Get fiscal responsibility records count
[**UpdateFiscalResponsibilityRecord**](FiscalResponsibilityRecordsAPI.md#UpdateFiscalResponsibilityRecord) | **Put** /api/v2/AccountingService/Fiscals/Authorities/FiscalResponsibilityRecords/{fiscalResponsibilityRecordId} | Update a fiscal responsibility record



## CreateFiscalResponsibilityRecord

> EmptyEnvelope CreateFiscalResponsibilityRecord(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalResponsibilityRecordCreateDto(fiscalResponsibilityRecordCreateDto).Execute()

Create a fiscal responsibility record



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	fiscalResponsibilityRecordCreateDto := *openapiclient.NewFiscalResponsibilityRecordCreateDto() // FiscalResponsibilityRecordCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalResponsibilityRecordsAPI.CreateFiscalResponsibilityRecord(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalResponsibilityRecordCreateDto(fiscalResponsibilityRecordCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalResponsibilityRecordsAPI.CreateFiscalResponsibilityRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiscalResponsibilityRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalResponsibilityRecordsAPI.CreateFiscalResponsibilityRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiscalResponsibilityRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalResponsibilityRecordCreateDto** | [**FiscalResponsibilityRecordCreateDto**](FiscalResponsibilityRecordCreateDto.md) |  | 

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


## DeleteFiscalResponsibilityRecord

> EmptyEnvelope DeleteFiscalResponsibilityRecord(ctx, fiscalResponsibilityRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a fiscal responsibility record



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
	fiscalResponsibilityRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalResponsibilityRecordsAPI.DeleteFiscalResponsibilityRecord(context.Background(), fiscalResponsibilityRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalResponsibilityRecordsAPI.DeleteFiscalResponsibilityRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFiscalResponsibilityRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalResponsibilityRecordsAPI.DeleteFiscalResponsibilityRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalResponsibilityRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiscalResponsibilityRecordRequest struct via the builder pattern


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


## GetFiscalResponsibilityRecord

> FiscalResponsibilityRecordDtoEnvelope GetFiscalResponsibilityRecord(ctx, fiscalAuthorityId, fiscalResponsibilityId, fiscalResponsibilityRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal responsibility record by ID



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
	fiscalResponsibilityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fiscalResponsibilityRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecord(context.Background(), fiscalAuthorityId, fiscalResponsibilityId, fiscalResponsibilityRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalResponsibilityRecord`: FiscalResponsibilityRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**fiscalResponsibilityId** | **string** |  | 
**fiscalResponsibilityRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalResponsibilityRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 



 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalResponsibilityRecordDtoEnvelope**](FiscalResponsibilityRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalResponsibilityRecords

> FiscalResponsibilityRecordDtoListEnvelope GetFiscalResponsibilityRecords(ctx, fiscalAuthorityId, fiscalResponsibilityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal responsibility records



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
	fiscalResponsibilityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecords(context.Background(), fiscalAuthorityId, fiscalResponsibilityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalResponsibilityRecords`: FiscalResponsibilityRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**fiscalResponsibilityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalResponsibilityRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalResponsibilityRecordDtoListEnvelope**](FiscalResponsibilityRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalResponsibilityRecordsCount

> Int32Envelope GetFiscalResponsibilityRecordsCount(ctx, fiscalAuthorityId, fiscalResponsibilityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal responsibility records count



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
	fiscalResponsibilityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecordsCount(context.Background(), fiscalAuthorityId, fiscalResponsibilityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecordsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalResponsibilityRecordsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalResponsibilityRecordsAPI.GetFiscalResponsibilityRecordsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**fiscalResponsibilityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalResponsibilityRecordsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 


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


## UpdateFiscalResponsibilityRecord

> EmptyEnvelope UpdateFiscalResponsibilityRecord(ctx, fiscalResponsibilityRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalResponsibilityRecordUpdateDto(fiscalResponsibilityRecordUpdateDto).Execute()

Update a fiscal responsibility record



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
	fiscalResponsibilityRecordId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	fiscalResponsibilityRecordUpdateDto := *openapiclient.NewFiscalResponsibilityRecordUpdateDto() // FiscalResponsibilityRecordUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalResponsibilityRecordsAPI.UpdateFiscalResponsibilityRecord(context.Background(), fiscalResponsibilityRecordId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalResponsibilityRecordUpdateDto(fiscalResponsibilityRecordUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalResponsibilityRecordsAPI.UpdateFiscalResponsibilityRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFiscalResponsibilityRecord`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalResponsibilityRecordsAPI.UpdateFiscalResponsibilityRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalResponsibilityRecordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFiscalResponsibilityRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalResponsibilityRecordUpdateDto** | [**FiscalResponsibilityRecordUpdateDto**](FiscalResponsibilityRecordUpdateDto.md) |  | 

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

