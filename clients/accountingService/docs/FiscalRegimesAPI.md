# \FiscalRegimesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFiscalRegime**](FiscalRegimesAPI.md#CreateFiscalRegime) | **Post** /api/v2/AccountingService/Fiscals/Authorities/FiscalRegimes | Create a fiscal regime
[**DeleteFiscalRegime**](FiscalRegimesAPI.md#DeleteFiscalRegime) | **Delete** /api/v2/AccountingService/Fiscals/Authorities/FiscalRegimes/{regimeId} | Delete a fiscal regime
[**GetFiscalRegime**](FiscalRegimesAPI.md#GetFiscalRegime) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalRegimes/{regimeId} | Get fiscal regime by ID
[**GetFiscalRegimes**](FiscalRegimesAPI.md#GetFiscalRegimes) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{authorityId}/FiscalRegimes | Get fiscal regimes for an authority
[**GetFiscalRegimesCount**](FiscalRegimesAPI.md#GetFiscalRegimesCount) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/FiscalRegimes/Count | Get fiscal regimes count
[**UpdateFiscalRegime**](FiscalRegimesAPI.md#UpdateFiscalRegime) | **Put** /api/v2/AccountingService/Fiscals/Authorities/FiscalRegimes/{regimeId} | Update a fiscal regime



## CreateFiscalRegime

> EmptyEnvelope CreateFiscalRegime(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalRegimeCreateDto(fiscalRegimeCreateDto).Execute()

Create a fiscal regime



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
	fiscalRegimeCreateDto := *openapiclient.NewFiscalRegimeCreateDto() // FiscalRegimeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalRegimesAPI.CreateFiscalRegime(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalRegimeCreateDto(fiscalRegimeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalRegimesAPI.CreateFiscalRegime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiscalRegime`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalRegimesAPI.CreateFiscalRegime`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiscalRegimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalRegimeCreateDto** | [**FiscalRegimeCreateDto**](FiscalRegimeCreateDto.md) |  | 

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


## DeleteFiscalRegime

> EmptyEnvelope DeleteFiscalRegime(ctx, regimeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a fiscal regime



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
	regimeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalRegimesAPI.DeleteFiscalRegime(context.Background(), regimeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalRegimesAPI.DeleteFiscalRegime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFiscalRegime`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalRegimesAPI.DeleteFiscalRegime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regimeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiscalRegimeRequest struct via the builder pattern


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


## GetFiscalRegime

> FiscalRegimeDtoEnvelope GetFiscalRegime(ctx, fiscalAuthorityId, regimeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal regime by ID



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
	regimeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalRegimesAPI.GetFiscalRegime(context.Background(), fiscalAuthorityId, regimeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalRegimesAPI.GetFiscalRegime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalRegime`: FiscalRegimeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalRegimesAPI.GetFiscalRegime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**regimeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalRegimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalRegimeDtoEnvelope**](FiscalRegimeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalRegimes

> FiscalRegimeDtoListEnvelope GetFiscalRegimes(ctx, authorityId).FiscalAuthorityId(fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal regimes for an authority



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
	resp, r, err := apiClient.FiscalRegimesAPI.GetFiscalRegimes(context.Background(), authorityId).FiscalAuthorityId(fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalRegimesAPI.GetFiscalRegimes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalRegimes`: FiscalRegimeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalRegimesAPI.GetFiscalRegimes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalRegimesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fiscalAuthorityId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalRegimeDtoListEnvelope**](FiscalRegimeDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalRegimesCount

> Int32Envelope GetFiscalRegimesCount(ctx, fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal regimes count



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
	resp, r, err := apiClient.FiscalRegimesAPI.GetFiscalRegimesCount(context.Background(), fiscalAuthorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalRegimesAPI.GetFiscalRegimesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalRegimesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalRegimesAPI.GetFiscalRegimesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalRegimesCountRequest struct via the builder pattern


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


## UpdateFiscalRegime

> EmptyEnvelope UpdateFiscalRegime(ctx, regimeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalRegimeUpdateDto(fiscalRegimeUpdateDto).Execute()

Update a fiscal regime



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
	regimeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	fiscalRegimeUpdateDto := *openapiclient.NewFiscalRegimeUpdateDto() // FiscalRegimeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalRegimesAPI.UpdateFiscalRegime(context.Background(), regimeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalRegimeUpdateDto(fiscalRegimeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalRegimesAPI.UpdateFiscalRegime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFiscalRegime`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalRegimesAPI.UpdateFiscalRegime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**regimeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFiscalRegimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalRegimeUpdateDto** | [**FiscalRegimeUpdateDto**](FiscalRegimeUpdateDto.md) |  | 

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

