# \FiscalAuthorityYearsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFiscalYear**](FiscalAuthorityYearsAPI.md#CreateFiscalYear) | **Post** /api/v2/AccountingService/Fiscals/Authorities/FiscalYears | Create a fiscal year
[**DeleteFiscalYear**](FiscalAuthorityYearsAPI.md#DeleteFiscalYear) | **Delete** /api/v2/AccountingService/Fiscals/Authorities/FiscalYears/{fiscalYearId} | Delete a fiscal year
[**GetFiscalYear**](FiscalAuthorityYearsAPI.md#GetFiscalYear) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalYears/{fiscalYearId} | Get fiscal year by ID for an authority
[**GetFiscalYears**](FiscalAuthorityYearsAPI.md#GetFiscalYears) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{authorityId}/FiscalYears | Get fiscal years for an authority
[**GetFiscalYearsCount**](FiscalAuthorityYearsAPI.md#GetFiscalYearsCount) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalYears/Count | Get fiscal years count for an authority
[**UpdateFiscalYear**](FiscalAuthorityYearsAPI.md#UpdateFiscalYear) | **Put** /api/v2/AccountingService/Fiscals/Authorities/FiscalYears/{fiscalYearId} | Update a fiscal year



## CreateFiscalYear

> EmptyEnvelope CreateFiscalYear(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalYearCreateDto(fiscalYearCreateDto).Execute()

Create a fiscal year



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
	fiscalYearCreateDto := *openapiclient.NewFiscalYearCreateDto() // FiscalYearCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalAuthorityYearsAPI.CreateFiscalYear(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalYearCreateDto(fiscalYearCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthorityYearsAPI.CreateFiscalYear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiscalYear`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthorityYearsAPI.CreateFiscalYear`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiscalYearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalYearCreateDto** | [**FiscalYearCreateDto**](FiscalYearCreateDto.md) |  | 

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


## DeleteFiscalYear

> EmptyEnvelope DeleteFiscalYear(ctx, fiscalYearId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a fiscal year



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
	fiscalYearId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalAuthorityYearsAPI.DeleteFiscalYear(context.Background(), fiscalYearId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthorityYearsAPI.DeleteFiscalYear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFiscalYear`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthorityYearsAPI.DeleteFiscalYear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalYearId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiscalYearRequest struct via the builder pattern


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


## GetFiscalYear

> FiscalYearDtoEnvelope GetFiscalYear(ctx, fiscalAuthorityId, fiscalYearId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal year by ID for an authority



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
	resp, r, err := apiClient.FiscalAuthorityYearsAPI.GetFiscalYear(context.Background(), fiscalAuthorityId, fiscalYearId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthorityYearsAPI.GetFiscalYear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalYear`: FiscalYearDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthorityYearsAPI.GetFiscalYear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**fiscalYearId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalYearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalYearDtoEnvelope**](FiscalYearDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalYears

> FiscalYearDtoListEnvelope GetFiscalYears(ctx, authorityId).TenantId(tenantId).FiscalAuthorityId(fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal years for an authority



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
	authorityId := "authorityId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalAuthorityYearsAPI.GetFiscalYears(context.Background(), authorityId).TenantId(tenantId).FiscalAuthorityId(fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthorityYearsAPI.GetFiscalYears``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalYears`: FiscalYearDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthorityYearsAPI.GetFiscalYears`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **fiscalAuthorityId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalYearDtoListEnvelope**](FiscalYearDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalYearsCount

> Int32Envelope GetFiscalYearsCount(ctx, fiscalAuthorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal years count for an authority



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalAuthorityYearsAPI.GetFiscalYearsCount(context.Background(), fiscalAuthorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthorityYearsAPI.GetFiscalYearsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalYearsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthorityYearsAPI.GetFiscalYearsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalYearsCountRequest struct via the builder pattern


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


## UpdateFiscalYear

> EmptyEnvelope UpdateFiscalYear(ctx, fiscalYearId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalYearUpdateDto(fiscalYearUpdateDto).Execute()

Update a fiscal year



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
	fiscalYearId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	fiscalYearUpdateDto := *openapiclient.NewFiscalYearUpdateDto() // FiscalYearUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalAuthorityYearsAPI.UpdateFiscalYear(context.Background(), fiscalYearId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalYearUpdateDto(fiscalYearUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthorityYearsAPI.UpdateFiscalYear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFiscalYear`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthorityYearsAPI.UpdateFiscalYear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalYearId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFiscalYearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalYearUpdateDto** | [**FiscalYearUpdateDto**](FiscalYearUpdateDto.md) |  | 

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

