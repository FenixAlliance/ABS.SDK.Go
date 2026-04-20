# \FinancialBooksAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFinancialBookAsync**](FinancialBooksAPI.md#CreateFinancialBookAsync) | **Post** /api/v2/AccountingService/FinancialBooks | Creates a new financial book
[**DeleteFinancialBookAsync**](FinancialBooksAPI.md#DeleteFinancialBookAsync) | **Delete** /api/v2/AccountingService/FinancialBooks/{financialBookId} | Deletes an existing financial book
[**GetFinancialBookDetailsAsync**](FinancialBooksAPI.md#GetFinancialBookDetailsAsync) | **Get** /api/v2/AccountingService/FinancialBooks/{financialBookId} | Gets the details of a specific financial book
[**GetFinancialBooksAsync**](FinancialBooksAPI.md#GetFinancialBooksAsync) | **Get** /api/v2/AccountingService/FinancialBooks | Get all financial books for a tenant
[**GetFinancialBooksCountAsync**](FinancialBooksAPI.md#GetFinancialBooksCountAsync) | **Get** /api/v2/AccountingService/FinancialBooks/Count | Get the count of financial books
[**UpdateFinancialBookAsync**](FinancialBooksAPI.md#UpdateFinancialBookAsync) | **Put** /api/v2/AccountingService/FinancialBooks/{financialBookId} | Updates an existing financial book



## CreateFinancialBookAsync

> EmptyEnvelope CreateFinancialBookAsync(ctx).TenantId(tenantId).FinancialBookCreateDto(financialBookCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Creates a new financial book



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
	financialBookCreateDto := *openapiclient.NewFinancialBookCreateDto("Name_example") // FinancialBookCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinancialBooksAPI.CreateFinancialBookAsync(context.Background()).TenantId(tenantId).FinancialBookCreateDto(financialBookCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialBooksAPI.CreateFinancialBookAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFinancialBookAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FinancialBooksAPI.CreateFinancialBookAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFinancialBookAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **financialBookCreateDto** | [**FinancialBookCreateDto**](FinancialBookCreateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## DeleteFinancialBookAsync

> EmptyEnvelope DeleteFinancialBookAsync(ctx, financialBookId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes an existing financial book



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
	financialBookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinancialBooksAPI.DeleteFinancialBookAsync(context.Background(), financialBookId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialBooksAPI.DeleteFinancialBookAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFinancialBookAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FinancialBooksAPI.DeleteFinancialBookAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**financialBookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFinancialBookAsyncRequest struct via the builder pattern


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


## GetFinancialBookDetailsAsync

> FinancialBookDtoEnvelope GetFinancialBookDetailsAsync(ctx, financialBookId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the details of a specific financial book



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
	financialBookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinancialBooksAPI.GetFinancialBookDetailsAsync(context.Background(), financialBookId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialBooksAPI.GetFinancialBookDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinancialBookDetailsAsync`: FinancialBookDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FinancialBooksAPI.GetFinancialBookDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**financialBookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialBookDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FinancialBookDtoEnvelope**](FinancialBookDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialBooksAsync

> FinancialBookDtoListEnvelope GetFinancialBooksAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all financial books for a tenant



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
	resp, r, err := apiClient.FinancialBooksAPI.GetFinancialBooksAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialBooksAPI.GetFinancialBooksAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinancialBooksAsync`: FinancialBookDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FinancialBooksAPI.GetFinancialBooksAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialBooksAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FinancialBookDtoListEnvelope**](FinancialBookDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialBooksCountAsync

> Int32Envelope GetFinancialBooksCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of financial books



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
	resp, r, err := apiClient.FinancialBooksAPI.GetFinancialBooksCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialBooksAPI.GetFinancialBooksCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinancialBooksCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `FinancialBooksAPI.GetFinancialBooksCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialBooksCountAsyncRequest struct via the builder pattern


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


## UpdateFinancialBookAsync

> EmptyEnvelope UpdateFinancialBookAsync(ctx, financialBookId).TenantId(tenantId).FinancialBookUpdateDto(financialBookUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Updates an existing financial book



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
	financialBookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	financialBookUpdateDto := *openapiclient.NewFinancialBookUpdateDto() // FinancialBookUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinancialBooksAPI.UpdateFinancialBookAsync(context.Background(), financialBookId).TenantId(tenantId).FinancialBookUpdateDto(financialBookUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinancialBooksAPI.UpdateFinancialBookAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFinancialBookAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FinancialBooksAPI.UpdateFinancialBookAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**financialBookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFinancialBookAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **financialBookUpdateDto** | [**FinancialBookUpdateDto**](FinancialBookUpdateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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

