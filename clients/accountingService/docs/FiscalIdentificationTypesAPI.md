# \FiscalIdentificationTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFiscalIdentificationType**](FiscalIdentificationTypesAPI.md#CreateFiscalIdentificationType) | **Post** /api/v2/AccountingService/Fiscals/Authorities/IdentificationTypes | Create a fiscal identification type
[**DeleteFiscalIdentificationType**](FiscalIdentificationTypesAPI.md#DeleteFiscalIdentificationType) | **Delete** /api/v2/AccountingService/Fiscals/Authorities/IdentificationTypes/{identificationTypeId} | Delete a fiscal identification type
[**GetFiscalIdentificationType**](FiscalIdentificationTypesAPI.md#GetFiscalIdentificationType) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{fiscalAuthorityId}/IdentificationTypes/{identificationTypeId} | Get fiscal identification type by ID
[**GetFiscalIdentificationTypes**](FiscalIdentificationTypesAPI.md#GetFiscalIdentificationTypes) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{authorityId}/IdentificationTypes | Get fiscal identification types for an authority
[**GetFiscalIdentificationTypesCount**](FiscalIdentificationTypesAPI.md#GetFiscalIdentificationTypesCount) | **Get** /api/v2/AccountingService/Fiscals/Authorities/{authorityId}/IdentificationTypes/Count | Get fiscal identification types count
[**UpdateFiscalIdentificationType**](FiscalIdentificationTypesAPI.md#UpdateFiscalIdentificationType) | **Put** /api/v2/AccountingService/Fiscals/Authorities/IdentificationTypes/{identificationTypeId} | Update a fiscal identification type



## CreateFiscalIdentificationType

> EmptyEnvelope CreateFiscalIdentificationType(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalIdentificationTypeCreateDto(fiscalIdentificationTypeCreateDto).Execute()

Create a fiscal identification type



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
	fiscalIdentificationTypeCreateDto := *openapiclient.NewFiscalIdentificationTypeCreateDto() // FiscalIdentificationTypeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalIdentificationTypesAPI.CreateFiscalIdentificationType(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalIdentificationTypeCreateDto(fiscalIdentificationTypeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalIdentificationTypesAPI.CreateFiscalIdentificationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFiscalIdentificationType`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalIdentificationTypesAPI.CreateFiscalIdentificationType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFiscalIdentificationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalIdentificationTypeCreateDto** | [**FiscalIdentificationTypeCreateDto**](FiscalIdentificationTypeCreateDto.md) |  | 

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


## DeleteFiscalIdentificationType

> EmptyEnvelope DeleteFiscalIdentificationType(ctx, identificationTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a fiscal identification type



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
	identificationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalIdentificationTypesAPI.DeleteFiscalIdentificationType(context.Background(), identificationTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalIdentificationTypesAPI.DeleteFiscalIdentificationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFiscalIdentificationType`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalIdentificationTypesAPI.DeleteFiscalIdentificationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identificationTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFiscalIdentificationTypeRequest struct via the builder pattern


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


## GetFiscalIdentificationType

> FiscalIdentificationTypeDtoEnvelope GetFiscalIdentificationType(ctx, fiscalAuthorityId, identificationTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal identification type by ID



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
	identificationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalIdentificationTypesAPI.GetFiscalIdentificationType(context.Background(), fiscalAuthorityId, identificationTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalIdentificationTypesAPI.GetFiscalIdentificationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalIdentificationType`: FiscalIdentificationTypeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalIdentificationTypesAPI.GetFiscalIdentificationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fiscalAuthorityId** | **string** |  | 
**identificationTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalIdentificationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalIdentificationTypeDtoEnvelope**](FiscalIdentificationTypeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalIdentificationTypes

> FiscalIdentificationTypeDtoListEnvelope GetFiscalIdentificationTypes(ctx, authorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal identification types for an authority



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
	authorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalIdentificationTypesAPI.GetFiscalIdentificationTypes(context.Background(), authorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalIdentificationTypesAPI.GetFiscalIdentificationTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalIdentificationTypes`: FiscalIdentificationTypeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalIdentificationTypesAPI.GetFiscalIdentificationTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalIdentificationTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FiscalIdentificationTypeDtoListEnvelope**](FiscalIdentificationTypeDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiscalIdentificationTypesCount

> Int32Envelope GetFiscalIdentificationTypesCount(ctx, authorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get fiscal identification types count



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
	authorityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalIdentificationTypesAPI.GetFiscalIdentificationTypesCount(context.Background(), authorityId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalIdentificationTypesAPI.GetFiscalIdentificationTypesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiscalIdentificationTypesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalIdentificationTypesAPI.GetFiscalIdentificationTypesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiscalIdentificationTypesCountRequest struct via the builder pattern


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


## UpdateFiscalIdentificationType

> EmptyEnvelope UpdateFiscalIdentificationType(ctx, identificationTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalIdentificationTypeUpdateDto(fiscalIdentificationTypeUpdateDto).Execute()

Update a fiscal identification type



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
	identificationTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	fiscalIdentificationTypeUpdateDto := *openapiclient.NewFiscalIdentificationTypeUpdateDto() // FiscalIdentificationTypeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FiscalIdentificationTypesAPI.UpdateFiscalIdentificationType(context.Background(), identificationTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).FiscalIdentificationTypeUpdateDto(fiscalIdentificationTypeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FiscalIdentificationTypesAPI.UpdateFiscalIdentificationType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFiscalIdentificationType`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FiscalIdentificationTypesAPI.UpdateFiscalIdentificationType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identificationTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFiscalIdentificationTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **fiscalIdentificationTypeUpdateDto** | [**FiscalIdentificationTypeUpdateDto**](FiscalIdentificationTypeUpdateDto.md) |  | 

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

