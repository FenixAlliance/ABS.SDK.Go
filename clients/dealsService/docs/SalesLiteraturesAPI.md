# \SalesLiteraturesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSalesLiteraturesAsync**](SalesLiteraturesAPI.md#CountSalesLiteraturesAsync) | **Get** /api/v2/DealsService/SalesLiteratures/Count | Get sales literatures count
[**CreateSalesLiteratureAsync**](SalesLiteraturesAPI.md#CreateSalesLiteratureAsync) | **Post** /api/v2/DealsService/SalesLiteratures | Create a sales literature
[**DeleteSalesLiteratureAsync**](SalesLiteraturesAPI.md#DeleteSalesLiteratureAsync) | **Delete** /api/v2/DealsService/SalesLiteratures/{salesLiteratureId} | Delete a sales literature
[**GetExtendedSalesLiteraturesAsync**](SalesLiteraturesAPI.md#GetExtendedSalesLiteraturesAsync) | **Get** /api/v2/DealsService/SalesLiteratures/Extended | Get extended sales literatures
[**GetSalesLiteratureAsync**](SalesLiteraturesAPI.md#GetSalesLiteratureAsync) | **Get** /api/v2/DealsService/SalesLiteratures/{salesLiteratureId} | Get sales literature by ID
[**GetSalesLiteraturesAsync**](SalesLiteraturesAPI.md#GetSalesLiteraturesAsync) | **Get** /api/v2/DealsService/SalesLiteratures | Get sales literatures
[**UpdateSalesLiteratureAsync**](SalesLiteraturesAPI.md#UpdateSalesLiteratureAsync) | **Put** /api/v2/DealsService/SalesLiteratures/{salesLiteratureId} | Update a sales literature



## CountSalesLiteraturesAsync

> Int32Envelope CountSalesLiteraturesAsync(ctx).TenantId(tenantId).Execute()

Get sales literatures count



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.CountSalesLiteraturesAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.CountSalesLiteraturesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSalesLiteraturesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.CountSalesLiteraturesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountSalesLiteraturesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## CreateSalesLiteratureAsync

> EmptyEnvelope CreateSalesLiteratureAsync(ctx).TenantId(tenantId).SalesLiteratureCreateDto(salesLiteratureCreateDto).Execute()

Create a sales literature



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
	salesLiteratureCreateDto := *openapiclient.NewSalesLiteratureCreateDto() // SalesLiteratureCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.CreateSalesLiteratureAsync(context.Background()).TenantId(tenantId).SalesLiteratureCreateDto(salesLiteratureCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.CreateSalesLiteratureAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSalesLiteratureAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.CreateSalesLiteratureAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSalesLiteratureAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **salesLiteratureCreateDto** | [**SalesLiteratureCreateDto**](SalesLiteratureCreateDto.md) |  | 

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


## DeleteSalesLiteratureAsync

> EmptyEnvelope DeleteSalesLiteratureAsync(ctx, salesLiteratureId).TenantId(tenantId).Execute()

Delete a sales literature



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
	salesLiteratureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.DeleteSalesLiteratureAsync(context.Background(), salesLiteratureId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.DeleteSalesLiteratureAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSalesLiteratureAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.DeleteSalesLiteratureAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**salesLiteratureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSalesLiteratureAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## GetExtendedSalesLiteraturesAsync

> ExtendedSalesLiteratureDtoListEnvelope GetExtendedSalesLiteraturesAsync(ctx).TenantId(tenantId).Execute()

Get extended sales literatures



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.GetExtendedSalesLiteraturesAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.GetExtendedSalesLiteraturesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedSalesLiteraturesAsync`: ExtendedSalesLiteratureDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.GetExtendedSalesLiteraturesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedSalesLiteraturesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ExtendedSalesLiteratureDtoListEnvelope**](ExtendedSalesLiteratureDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesLiteratureAsync

> SalesLiteratureDtoEnvelope GetSalesLiteratureAsync(ctx, salesLiteratureId).TenantId(tenantId).Execute()

Get sales literature by ID



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
	salesLiteratureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.GetSalesLiteratureAsync(context.Background(), salesLiteratureId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.GetSalesLiteratureAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesLiteratureAsync`: SalesLiteratureDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.GetSalesLiteratureAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**salesLiteratureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesLiteratureAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**SalesLiteratureDtoEnvelope**](SalesLiteratureDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSalesLiteraturesAsync

> SalesLiteratureDtoListEnvelope GetSalesLiteraturesAsync(ctx).TenantId(tenantId).Execute()

Get sales literatures



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.GetSalesLiteraturesAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.GetSalesLiteraturesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSalesLiteraturesAsync`: SalesLiteratureDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.GetSalesLiteraturesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesLiteraturesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**SalesLiteratureDtoListEnvelope**](SalesLiteratureDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSalesLiteratureAsync

> EmptyEnvelope UpdateSalesLiteratureAsync(ctx, salesLiteratureId).TenantId(tenantId).SalesLiteratureUpdateDto(salesLiteratureUpdateDto).Execute()

Update a sales literature



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
	salesLiteratureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	salesLiteratureUpdateDto := *openapiclient.NewSalesLiteratureUpdateDto() // SalesLiteratureUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.UpdateSalesLiteratureAsync(context.Background(), salesLiteratureId).TenantId(tenantId).SalesLiteratureUpdateDto(salesLiteratureUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.UpdateSalesLiteratureAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSalesLiteratureAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.UpdateSalesLiteratureAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**salesLiteratureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSalesLiteratureAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **salesLiteratureUpdateDto** | [**SalesLiteratureUpdateDto**](SalesLiteratureUpdateDto.md) |  | 

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

