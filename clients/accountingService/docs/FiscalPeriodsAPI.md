# \FiscalPeriodsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFiscalPeriod**](FiscalPeriodsAPI.md#CreateFiscalPeriod) | **Post** /api/v2/AccountingService/Fiscals/Authorities/FiscalPeriods | Create a fiscal period
[**DeleteFiscalPeriod**](FiscalPeriodsAPI.md#DeleteFiscalPeriod) | **Delete** /api/v2/AccountingService/Fiscals/Authorities/FiscalPeriods/{fiscalPeriodId} | Delete a fiscal period
[**GetFiscalPeriod**](FiscalPeriodsAPI.md#GetFiscalPeriod) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalYears/{fiscalYearId}/FiscalPeriods/{fiscalPeriodId} | Get fiscal period by ID
[**GetFiscalPeriods**](FiscalPeriodsAPI.md#GetFiscalPeriods) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{authorityId}/FiscalYears/{fiscalYearId}/FiscalPeriods | Get fiscal periods for a fiscal year
[**GetFiscalPeriodsCount**](FiscalPeriodsAPI.md#GetFiscalPeriodsCount) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalYears/{fiscalYearId}/FiscalPeriods/Count | Get fiscal periods count
[**UpdateFiscalPeriod**](FiscalPeriodsAPI.md#UpdateFiscalPeriod) | **Put** /api/v2/AccountingService/Fiscals/Authorities/FiscalPeriods/{fiscalPeriodId} | Update a fiscal period



## CreateFiscalPeriod

> EmptyEnvelope CreateFiscalPeriod(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalPeriodCreateDto(fiscalPeriodCreateDto).Execute()

Create a fiscal period



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
	fiscalPeriodCreateDto := *openapiclient.NewFiscalPeriodCreateDto() // FiscalPeriodCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalPeriodsAPI.CreateFiscalPeriod(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalPeriodCreateDto(fiscalPeriodCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalPeriodsAPI.CreateFiscalPeriod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiscalPeriod`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalPeriodsAPI.CreateFiscalPeriod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiscalPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalPeriodCreateDto** | [**FiscalPeriodCreateDto**](FiscalPeriodCreateDto.md) |  | 

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


## DeleteFiscalPeriod

> EmptyEnvelope DeleteFiscalPeriod(ctx, fiscalPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a fiscal period



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
	fiscalPeriodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalPeriodsAPI.DeleteFiscalPeriod(context.Background(), fiscalPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalPeriodsAPI.DeleteFiscalPeriod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFiscalPeriod`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalPeriodsAPI.DeleteFiscalPeriod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalPeriodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiscalPeriodRequest struct via the builder pattern


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


## GetFiscalPeriod

> FiscalPeriodDtoEnvelope GetFiscalPeriod(ctx, fiscalAuthorityId, fiscalYearId, fiscalPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal period by ID



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
	fiscalYearId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	fiscalPeriodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalPeriodsAPI.GetFiscalPeriod(context.Background(), fiscalAuthorityId, fiscalYearId, fiscalPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalPeriodsAPI.GetFiscalPeriod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalPeriod`: FiscalPeriodDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalPeriodsAPI.GetFiscalPeriod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**fiscalYearId** | **string** |  | 
**fiscalPeriodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 



 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalPeriodDtoEnvelope**](FiscalPeriodDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalPeriods

> FiscalPeriodDtoListEnvelope GetFiscalPeriods(ctx, fiscalYearId, authorityId).TenantId(tenantId).FiscalAuthorityId(fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal periods for a fiscal year



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
	fiscalYearId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	authorityId := "authorityId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalPeriodsAPI.GetFiscalPeriods(context.Background(), fiscalYearId, authorityId).TenantId(tenantId).FiscalAuthorityId(fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalPeriodsAPI.GetFiscalPeriods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalPeriods`: FiscalPeriodDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalPeriodsAPI.GetFiscalPeriods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalYearId** | **string** |  | 
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalPeriodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fiscalAuthorityId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalPeriodDtoListEnvelope**](FiscalPeriodDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalPeriodsCount

> Int32Envelope GetFiscalPeriodsCount(ctx, fiscalAuthorityId, fiscalYearId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal periods count



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
	fiscalYearId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalPeriodsAPI.GetFiscalPeriodsCount(context.Background(), fiscalAuthorityId, fiscalYearId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalPeriodsAPI.GetFiscalPeriodsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalPeriodsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalPeriodsAPI.GetFiscalPeriodsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**fiscalYearId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalPeriodsCountRequest struct via the builder pattern


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


## UpdateFiscalPeriod

> EmptyEnvelope UpdateFiscalPeriod(ctx, fiscalPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalPeriodUpdateDto(fiscalPeriodUpdateDto).Execute()

Update a fiscal period



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
	fiscalPeriodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	fiscalPeriodUpdateDto := *openapiclient.NewFiscalPeriodUpdateDto() // FiscalPeriodUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalPeriodsAPI.UpdateFiscalPeriod(context.Background(), fiscalPeriodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalPeriodUpdateDto(fiscalPeriodUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalPeriodsAPI.UpdateFiscalPeriod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFiscalPeriod`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalPeriodsAPI.UpdateFiscalPeriod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalPeriodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFiscalPeriodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalPeriodUpdateDto** | [**FiscalPeriodUpdateDto**](FiscalPeriodUpdateDto.md) |  | 

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

