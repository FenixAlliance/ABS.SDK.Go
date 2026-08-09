# \SigningLogsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSigningLogByIdAsync**](SigningLogsAPI.md#GetSigningLogByIdAsync) | **Get** /api/v2/TrustService/SigningLogs/{id} | Get signing log by ID
[**GetSigningLogsAsync**](SigningLogsAPI.md#GetSigningLogsAsync) | **Get** /api/v2/TrustService/SigningLogs | Get all signing logs
[**GetSigningLogsCountAsync**](SigningLogsAPI.md#GetSigningLogsCountAsync) | **Get** /api/v2/TrustService/SigningLogs/Count | Get signing logs count



## GetSigningLogByIdAsync

> SigningLogDto GetSigningLogByIdAsync(ctx, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get signing log by ID



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigningLogsAPI.GetSigningLogByIdAsync(context.Background(), id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigningLogsAPI.GetSigningLogByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSigningLogByIdAsync`: SigningLogDto
	fmt.Fprintf(os.Stdout, "Response from `SigningLogsAPI.GetSigningLogByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSigningLogByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SigningLogDto**](SigningLogDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSigningLogsAsync

> SigningLogDtoListEnvelope GetSigningLogsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SigningLogDtoCollectionQueryParameters(signingLogDtoCollectionQueryParameters).Execute()

Get all signing logs



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
	signingLogDtoCollectionQueryParameters := *openapiclient.NewSigningLogDtoCollectionQueryParameters() // SigningLogDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigningLogsAPI.GetSigningLogsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SigningLogDtoCollectionQueryParameters(signingLogDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigningLogsAPI.GetSigningLogsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSigningLogsAsync`: SigningLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SigningLogsAPI.GetSigningLogsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSigningLogsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **signingLogDtoCollectionQueryParameters** | [**SigningLogDtoCollectionQueryParameters**](SigningLogDtoCollectionQueryParameters.md) |  | 

### Return type

[**SigningLogDtoListEnvelope**](SigningLogDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSigningLogsCountAsync

> Int32Envelope GetSigningLogsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SigningLogDtoCollectionQueryParameters(signingLogDtoCollectionQueryParameters).Execute()

Get signing logs count



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
	signingLogDtoCollectionQueryParameters := *openapiclient.NewSigningLogDtoCollectionQueryParameters() // SigningLogDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SigningLogsAPI.GetSigningLogsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SigningLogDtoCollectionQueryParameters(signingLogDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SigningLogsAPI.GetSigningLogsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSigningLogsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SigningLogsAPI.GetSigningLogsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSigningLogsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **signingLogDtoCollectionQueryParameters** | [**SigningLogDtoCollectionQueryParameters**](SigningLogDtoCollectionQueryParameters.md) |  | 

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

