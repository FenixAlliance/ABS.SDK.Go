# \AssetTransfersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssetTransferAsync**](AssetTransfersAPI.md#CreateAssetTransferAsync) | **Post** /api/v2/AssetsService/AssetTransfers | Creates a new asset transfer
[**DeleteAssetTransferAsync**](AssetTransfersAPI.md#DeleteAssetTransferAsync) | **Delete** /api/v2/AssetsService/AssetTransfers/{transferId} | Deletes an asset transfer
[**GetAssetTransferAsync**](AssetTransfersAPI.md#GetAssetTransferAsync) | **Get** /api/v2/AssetsService/AssetTransfers/{transferId} | Gets a single asset transfer by ID
[**GetAssetTransfersAsync**](AssetTransfersAPI.md#GetAssetTransfersAsync) | **Get** /api/v2/AssetsService/AssetTransfers | Gets a list of asset transfers
[**GetAssetTransfersCountAsync**](AssetTransfersAPI.md#GetAssetTransfersCountAsync) | **Get** /api/v2/AssetsService/AssetTransfers/Count | Gets the count of asset transfers
[**UpdateAssetTransferAsync**](AssetTransfersAPI.md#UpdateAssetTransferAsync) | **Put** /api/v2/AssetsService/AssetTransfers/{transferId} | Updates an existing asset transfer



## CreateAssetTransferAsync

> EmptyEnvelope CreateAssetTransferAsync(ctx).TenantId(tenantId).AssetTransferCreateDto(assetTransferCreateDto).Execute()

Creates a new asset transfer



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
	assetTransferCreateDto := *openapiclient.NewAssetTransferCreateDto() // AssetTransferCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTransfersAPI.CreateAssetTransferAsync(context.Background()).TenantId(tenantId).AssetTransferCreateDto(assetTransferCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTransfersAPI.CreateAssetTransferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssetTransferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AssetTransfersAPI.CreateAssetTransferAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetTransferAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **assetTransferCreateDto** | [**AssetTransferCreateDto**](AssetTransferCreateDto.md) |  | 

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


## DeleteAssetTransferAsync

> EmptyEnvelope DeleteAssetTransferAsync(ctx, transferId).TenantId(tenantId).Execute()

Deletes an asset transfer



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
	transferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTransfersAPI.DeleteAssetTransferAsync(context.Background(), transferId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTransfersAPI.DeleteAssetTransferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAssetTransferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AssetTransfersAPI.DeleteAssetTransferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssetTransferAsyncRequest struct via the builder pattern


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


## GetAssetTransferAsync

> AssetTransferDtoEnvelope GetAssetTransferAsync(ctx, transferId).TenantId(tenantId).Execute()

Gets a single asset transfer by ID



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
	transferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTransfersAPI.GetAssetTransferAsync(context.Background(), transferId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTransfersAPI.GetAssetTransferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTransferAsync`: AssetTransferDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AssetTransfersAPI.GetAssetTransferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTransferAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**AssetTransferDtoEnvelope**](AssetTransferDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTransfersAsync

> AssetTransferDtoListEnvelope GetAssetTransfersAsync(ctx).TenantId(tenantId).Execute()

Gets a list of asset transfers



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
	resp, r, err := apiClient.AssetTransfersAPI.GetAssetTransfersAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTransfersAPI.GetAssetTransfersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTransfersAsync`: AssetTransferDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AssetTransfersAPI.GetAssetTransfersAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTransfersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**AssetTransferDtoListEnvelope**](AssetTransferDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetTransfersCountAsync

> Int32Envelope GetAssetTransfersCountAsync(ctx).TenantId(tenantId).Execute()

Gets the count of asset transfers



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
	resp, r, err := apiClient.AssetTransfersAPI.GetAssetTransfersCountAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTransfersAPI.GetAssetTransfersCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssetTransfersCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `AssetTransfersAPI.GetAssetTransfersCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetTransfersCountAsyncRequest struct via the builder pattern


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


## UpdateAssetTransferAsync

> EmptyEnvelope UpdateAssetTransferAsync(ctx, transferId).TenantId(tenantId).AssetTransferUpdateDto(assetTransferUpdateDto).Execute()

Updates an existing asset transfer



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
	transferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	assetTransferUpdateDto := *openapiclient.NewAssetTransferUpdateDto() // AssetTransferUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetTransfersAPI.UpdateAssetTransferAsync(context.Background(), transferId).TenantId(tenantId).AssetTransferUpdateDto(assetTransferUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetTransfersAPI.UpdateAssetTransferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAssetTransferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AssetTransfersAPI.UpdateAssetTransferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetTransferAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **assetTransferUpdateDto** | [**AssetTransferUpdateDto**](AssetTransferUpdateDto.md) |  | 

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

