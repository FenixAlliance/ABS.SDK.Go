# \LedgerTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLedgerTypeAsync**](LedgerTypesAPI.md#CreateLedgerTypeAsync) | **Post** /api/v2/AccountingService/LedgerTypes | Creates a new ledger type
[**DeleteLedgerTypeAsync**](LedgerTypesAPI.md#DeleteLedgerTypeAsync) | **Delete** /api/v2/AccountingService/LedgerTypes/{ledgerTypeId} | Deletes a ledger type
[**GetLedgerTypeDetailsAsync**](LedgerTypesAPI.md#GetLedgerTypeDetailsAsync) | **Get** /api/v2/AccountingService/LedgerTypes/{ledgerTypeId} | Gets a ledger type by ID
[**GetLedgerTypesAsync**](LedgerTypesAPI.md#GetLedgerTypesAsync) | **Get** /api/v2/AccountingService/LedgerTypes | Retrieves all ledger types
[**GetLedgerTypesCountAsync**](LedgerTypesAPI.md#GetLedgerTypesCountAsync) | **Get** /api/v2/AccountingService/LedgerTypes/Count | Counts ledger types
[**UpdateLedgerTypeAsync**](LedgerTypesAPI.md#UpdateLedgerTypeAsync) | **Put** /api/v2/AccountingService/LedgerTypes/{ledgerTypeId} | Updates a ledger type



## CreateLedgerTypeAsync

> EmptyEnvelope CreateLedgerTypeAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LedgerTypeCreateDto(ledgerTypeCreateDto).Execute()

Creates a new ledger type



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
	ledgerTypeCreateDto := *openapiclient.NewLedgerTypeCreateDto("Name_example") // LedgerTypeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerTypesAPI.CreateLedgerTypeAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LedgerTypeCreateDto(ledgerTypeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerTypesAPI.CreateLedgerTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLedgerTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LedgerTypesAPI.CreateLedgerTypeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLedgerTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **ledgerTypeCreateDto** | [**LedgerTypeCreateDto**](LedgerTypeCreateDto.md) |  | 

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


## DeleteLedgerTypeAsync

> EmptyEnvelope DeleteLedgerTypeAsync(ctx, ledgerTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a ledger type



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
	ledgerTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerTypesAPI.DeleteLedgerTypeAsync(context.Background(), ledgerTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerTypesAPI.DeleteLedgerTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLedgerTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LedgerTypesAPI.DeleteLedgerTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledgerTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLedgerTypeAsyncRequest struct via the builder pattern


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


## GetLedgerTypeDetailsAsync

> LedgerTypeDtoEnvelope GetLedgerTypeDetailsAsync(ctx, ledgerTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets a ledger type by ID



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
	ledgerTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerTypesAPI.GetLedgerTypeDetailsAsync(context.Background(), ledgerTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerTypesAPI.GetLedgerTypeDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerTypeDetailsAsync`: LedgerTypeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LedgerTypesAPI.GetLedgerTypeDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledgerTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerTypeDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**LedgerTypeDtoEnvelope**](LedgerTypeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLedgerTypesAsync

> LedgerTypeDtoIReadOnlyListEnvelope GetLedgerTypesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieves all ledger types



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
	resp, r, err := apiClient.LedgerTypesAPI.GetLedgerTypesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerTypesAPI.GetLedgerTypesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerTypesAsync`: LedgerTypeDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LedgerTypesAPI.GetLedgerTypesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerTypesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**LedgerTypeDtoIReadOnlyListEnvelope**](LedgerTypeDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLedgerTypesCountAsync

> Int32Envelope GetLedgerTypesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Counts ledger types



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
	resp, r, err := apiClient.LedgerTypesAPI.GetLedgerTypesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerTypesAPI.GetLedgerTypesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedgerTypesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `LedgerTypesAPI.GetLedgerTypesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerTypesCountAsyncRequest struct via the builder pattern


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


## UpdateLedgerTypeAsync

> EmptyEnvelope UpdateLedgerTypeAsync(ctx, ledgerTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LedgerTypeUpdateDto(ledgerTypeUpdateDto).Execute()

Updates a ledger type



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
	ledgerTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	ledgerTypeUpdateDto := *openapiclient.NewLedgerTypeUpdateDto() // LedgerTypeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LedgerTypesAPI.UpdateLedgerTypeAsync(context.Background(), ledgerTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LedgerTypeUpdateDto(ledgerTypeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LedgerTypesAPI.UpdateLedgerTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLedgerTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LedgerTypesAPI.UpdateLedgerTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ledgerTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLedgerTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **ledgerTypeUpdateDto** | [**LedgerTypeUpdateDto**](LedgerTypeUpdateDto.md) |  | 

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

