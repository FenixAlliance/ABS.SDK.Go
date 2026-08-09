# \PointOfSalesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountPointOfSalesAsync**](PointOfSalesAPI.md#CountPointOfSalesAsync) | **Get** /api/v2/SalesService/PointOfSales/Count | Get point of sales count
[**CreatePointOfSaleAsync**](PointOfSalesAPI.md#CreatePointOfSaleAsync) | **Post** /api/v2/SalesService/PointOfSales | Create a point of sale
[**DeletePointOfSaleAsync**](PointOfSalesAPI.md#DeletePointOfSaleAsync) | **Delete** /api/v2/SalesService/PointOfSales/{pointOfSaleId} | Delete a point of sale
[**GetPointOfSaleAsync**](PointOfSalesAPI.md#GetPointOfSaleAsync) | **Get** /api/v2/SalesService/PointOfSales/{pointOfSaleId} | Get point of sale by ID
[**GetPointOfSalesAsync**](PointOfSalesAPI.md#GetPointOfSalesAsync) | **Get** /api/v2/SalesService/PointOfSales | Get point of sales
[**PatchPointOfSaleAsync**](PointOfSalesAPI.md#PatchPointOfSaleAsync) | **Patch** /api/v2/SalesService/PointOfSales/{pointOfSaleId} | Patch a point of sale
[**UpdatePointOfSaleAsync**](PointOfSalesAPI.md#UpdatePointOfSaleAsync) | **Put** /api/v2/SalesService/PointOfSales/{pointOfSaleId} | Update a point of sale



## CountPointOfSalesAsync

> Int32Envelope CountPointOfSalesAsync(ctx).TenantId(tenantId).PointOfSaleDtoCollectionQueryParameters(pointOfSaleDtoCollectionQueryParameters).Execute()

Get point of sales count



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
	pointOfSaleDtoCollectionQueryParameters := *openapiclient.NewPointOfSaleDtoCollectionQueryParameters() // PointOfSaleDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointOfSalesAPI.CountPointOfSalesAsync(context.Background()).TenantId(tenantId).PointOfSaleDtoCollectionQueryParameters(pointOfSaleDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointOfSalesAPI.CountPointOfSalesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountPointOfSalesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `PointOfSalesAPI.CountPointOfSalesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountPointOfSalesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **pointOfSaleDtoCollectionQueryParameters** | [**PointOfSaleDtoCollectionQueryParameters**](PointOfSaleDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePointOfSaleAsync

> EmptyEnvelope CreatePointOfSaleAsync(ctx).TenantId(tenantId).PointOfSaleCreateDto(pointOfSaleCreateDto).Execute()

Create a point of sale



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
	pointOfSaleCreateDto := *openapiclient.NewPointOfSaleCreateDto("Title_example") // PointOfSaleCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointOfSalesAPI.CreatePointOfSaleAsync(context.Background()).TenantId(tenantId).PointOfSaleCreateDto(pointOfSaleCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointOfSalesAPI.CreatePointOfSaleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePointOfSaleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PointOfSalesAPI.CreatePointOfSaleAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePointOfSaleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **pointOfSaleCreateDto** | [**PointOfSaleCreateDto**](PointOfSaleCreateDto.md) |  | 

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


## DeletePointOfSaleAsync

> EmptyEnvelope DeletePointOfSaleAsync(ctx, pointOfSaleId).TenantId(tenantId).Execute()

Delete a point of sale



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
	pointOfSaleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointOfSalesAPI.DeletePointOfSaleAsync(context.Background(), pointOfSaleId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointOfSalesAPI.DeletePointOfSaleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePointOfSaleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PointOfSalesAPI.DeletePointOfSaleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pointOfSaleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePointOfSaleAsyncRequest struct via the builder pattern


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


## GetPointOfSaleAsync

> PointOfSaleDtoEnvelope GetPointOfSaleAsync(ctx, pointOfSaleId).TenantId(tenantId).Execute()

Get point of sale by ID



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
	pointOfSaleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointOfSalesAPI.GetPointOfSaleAsync(context.Background(), pointOfSaleId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointOfSalesAPI.GetPointOfSaleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPointOfSaleAsync`: PointOfSaleDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PointOfSalesAPI.GetPointOfSaleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pointOfSaleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointOfSaleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**PointOfSaleDtoEnvelope**](PointOfSaleDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPointOfSalesAsync

> PointOfSaleDtoListEnvelope GetPointOfSalesAsync(ctx).TenantId(tenantId).PointOfSaleDtoCollectionQueryParameters(pointOfSaleDtoCollectionQueryParameters).Execute()

Get point of sales



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
	pointOfSaleDtoCollectionQueryParameters := *openapiclient.NewPointOfSaleDtoCollectionQueryParameters() // PointOfSaleDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointOfSalesAPI.GetPointOfSalesAsync(context.Background()).TenantId(tenantId).PointOfSaleDtoCollectionQueryParameters(pointOfSaleDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointOfSalesAPI.GetPointOfSalesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPointOfSalesAsync`: PointOfSaleDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PointOfSalesAPI.GetPointOfSalesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPointOfSalesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **pointOfSaleDtoCollectionQueryParameters** | [**PointOfSaleDtoCollectionQueryParameters**](PointOfSaleDtoCollectionQueryParameters.md) |  | 

### Return type

[**PointOfSaleDtoListEnvelope**](PointOfSaleDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPointOfSaleAsync

> EmptyEnvelope PatchPointOfSaleAsync(ctx, pointOfSaleId).TenantId(tenantId).PatchOperation(patchOperation).Execute()

Patch a point of sale



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
	pointOfSaleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointOfSalesAPI.PatchPointOfSaleAsync(context.Background(), pointOfSaleId).TenantId(tenantId).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointOfSalesAPI.PatchPointOfSaleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPointOfSaleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PointOfSalesAPI.PatchPointOfSaleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pointOfSaleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPointOfSaleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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


## UpdatePointOfSaleAsync

> EmptyEnvelope UpdatePointOfSaleAsync(ctx, pointOfSaleId).TenantId(tenantId).PointOfSaleUpdateDto(pointOfSaleUpdateDto).Execute()

Update a point of sale



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
	pointOfSaleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pointOfSaleUpdateDto := *openapiclient.NewPointOfSaleUpdateDto() // PointOfSaleUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PointOfSalesAPI.UpdatePointOfSaleAsync(context.Background(), pointOfSaleId).TenantId(tenantId).PointOfSaleUpdateDto(pointOfSaleUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PointOfSalesAPI.UpdatePointOfSaleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePointOfSaleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PointOfSalesAPI.UpdatePointOfSaleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pointOfSaleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePointOfSaleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **pointOfSaleUpdateDto** | [**PointOfSaleUpdateDto**](PointOfSaleUpdateDto.md) |  | 

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

