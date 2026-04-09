# \FiscalAuthoritiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFiscalAuthority**](FiscalAuthoritiesAPI.md#CreateFiscalAuthority) | **Post** /api/v2/AccountingService/Fiscals/Authorities | Create a fiscal authority
[**DeleteFiscalAuthority**](FiscalAuthoritiesAPI.md#DeleteFiscalAuthority) | **Delete** /api/v2/AccountingService/Fiscals/Authorities/{authorityId} | Delete a fiscal authority
[**GetFiscalAuthorities**](FiscalAuthoritiesAPI.md#GetFiscalAuthorities) | **Get** /api/v2/AccountingService/Fiscals/Authorities | Get fiscal authorities
[**GetFiscalAuthoritiesCount**](FiscalAuthoritiesAPI.md#GetFiscalAuthoritiesCount) | **Get** /api/v2/AccountingService/Fiscals/Authorities/Count | Get fiscal authorities count
[**GetFiscalAuthority**](FiscalAuthoritiesAPI.md#GetFiscalAuthority) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{authorityId} | Get fiscal authority by ID
[**UpdateFiscalAuthority**](FiscalAuthoritiesAPI.md#UpdateFiscalAuthority) | **Put** /api/v2/AccountingService/Fiscals/Authorities/{authorityId} | Update a fiscal authority



## CreateFiscalAuthority

> EmptyEnvelope CreateFiscalAuthority(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalAuthorityCreateDto(fiscalAuthorityCreateDto).Execute()

Create a fiscal authority



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
	fiscalAuthorityCreateDto := *openapiclient.NewFiscalAuthorityCreateDto() // FiscalAuthorityCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalAuthoritiesAPI.CreateFiscalAuthority(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalAuthorityCreateDto(fiscalAuthorityCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthoritiesAPI.CreateFiscalAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiscalAuthority`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthoritiesAPI.CreateFiscalAuthority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiscalAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalAuthorityCreateDto** | [**FiscalAuthorityCreateDto**](FiscalAuthorityCreateDto.md) |  | 

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


## DeleteFiscalAuthority

> EmptyEnvelope DeleteFiscalAuthority(ctx, authorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a fiscal authority



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
	resp, r, err := apiClient.FiscalAuthoritiesAPI.DeleteFiscalAuthority(context.Background(), authorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthoritiesAPI.DeleteFiscalAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFiscalAuthority`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthoritiesAPI.DeleteFiscalAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiscalAuthorityRequest struct via the builder pattern


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


## GetFiscalAuthorities

> FiscalAuthorityDtoListEnvelope GetFiscalAuthorities(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal authorities



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
	resp, r, err := apiClient.FiscalAuthoritiesAPI.GetFiscalAuthorities(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthoritiesAPI.GetFiscalAuthorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalAuthorities`: FiscalAuthorityDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthoritiesAPI.GetFiscalAuthorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalAuthoritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalAuthorityDtoListEnvelope**](FiscalAuthorityDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalAuthoritiesCount

> Int32Envelope GetFiscalAuthoritiesCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal authorities count



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
	resp, r, err := apiClient.FiscalAuthoritiesAPI.GetFiscalAuthoritiesCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthoritiesAPI.GetFiscalAuthoritiesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalAuthoritiesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthoritiesAPI.GetFiscalAuthoritiesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalAuthoritiesCountRequest struct via the builder pattern


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


## GetFiscalAuthority

> FiscalAuthorityDtoEnvelope GetFiscalAuthority(ctx, authorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal authority by ID



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
	resp, r, err := apiClient.FiscalAuthoritiesAPI.GetFiscalAuthority(context.Background(), authorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthoritiesAPI.GetFiscalAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalAuthority`: FiscalAuthorityDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthoritiesAPI.GetFiscalAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalAuthorityDtoEnvelope**](FiscalAuthorityDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFiscalAuthority

> EmptyEnvelope UpdateFiscalAuthority(ctx, authorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalAuthorityUpdateDto(fiscalAuthorityUpdateDto).Execute()

Update a fiscal authority



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
	fiscalAuthorityUpdateDto := *openapiclient.NewFiscalAuthorityUpdateDto() // FiscalAuthorityUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalAuthoritiesAPI.UpdateFiscalAuthority(context.Background(), authorityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalAuthorityUpdateDto(fiscalAuthorityUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalAuthoritiesAPI.UpdateFiscalAuthority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFiscalAuthority`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalAuthoritiesAPI.UpdateFiscalAuthority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFiscalAuthorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalAuthorityUpdateDto** | [**FiscalAuthorityUpdateDto**](FiscalAuthorityUpdateDto.md) |  | 

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

