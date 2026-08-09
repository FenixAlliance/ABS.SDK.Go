# \MarketingLeadsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMarketingLeadAsync**](MarketingLeadsAPI.md#CreateMarketingLeadAsync) | **Post** /api/v2/MarketingService/MarketingLeads | Create a marketing lead
[**DeleteMarketingLeadAsync**](MarketingLeadsAPI.md#DeleteMarketingLeadAsync) | **Delete** /api/v2/MarketingService/MarketingLeads/{marketingLeadId} | Delete a marketing lead
[**GetMarketingLeadDetailsAsync**](MarketingLeadsAPI.md#GetMarketingLeadDetailsAsync) | **Get** /api/v2/MarketingService/MarketingLeads/{marketingLeadId} | Get marketing lead by ID
[**GetMarketingLeadsCountAsync**](MarketingLeadsAPI.md#GetMarketingLeadsCountAsync) | **Get** /api/v2/MarketingService/MarketingLeads/Count | Get marketing leads count
[**GetMarketingLeadsODataAsync**](MarketingLeadsAPI.md#GetMarketingLeadsODataAsync) | **Get** /api/v2/MarketingService/MarketingLeads | Get marketing leads
[**PatchMarketingLeadAsync**](MarketingLeadsAPI.md#PatchMarketingLeadAsync) | **Patch** /api/v2/MarketingService/MarketingLeads/{marketingLeadId} | Patch a marketing lead
[**UpdateMarketingLeadAsync**](MarketingLeadsAPI.md#UpdateMarketingLeadAsync) | **Put** /api/v2/MarketingService/MarketingLeads/{marketingLeadId} | Update a marketing lead



## CreateMarketingLeadAsync

> EmptyEnvelope CreateMarketingLeadAsync(ctx).TenantId(tenantId).MarketingLeadCreateDto(marketingLeadCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create a marketing lead

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
	marketingLeadCreateDto := *openapiclient.NewMarketingLeadCreateDto() // MarketingLeadCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingLeadsAPI.CreateMarketingLeadAsync(context.Background()).TenantId(tenantId).MarketingLeadCreateDto(marketingLeadCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingLeadsAPI.CreateMarketingLeadAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMarketingLeadAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingLeadsAPI.CreateMarketingLeadAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMarketingLeadAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **marketingLeadCreateDto** | [**MarketingLeadCreateDto**](MarketingLeadCreateDto.md) |  | 
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


## DeleteMarketingLeadAsync

> EmptyEnvelope DeleteMarketingLeadAsync(ctx, marketingLeadId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a marketing lead

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
	marketingLeadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingLeadsAPI.DeleteMarketingLeadAsync(context.Background(), marketingLeadId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingLeadsAPI.DeleteMarketingLeadAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMarketingLeadAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingLeadsAPI.DeleteMarketingLeadAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingLeadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingLeadAsyncRequest struct via the builder pattern


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


## GetMarketingLeadDetailsAsync

> MarketingLeadDtoEnvelope GetMarketingLeadDetailsAsync(ctx, marketingLeadId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get marketing lead by ID

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
	marketingLeadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingLeadsAPI.GetMarketingLeadDetailsAsync(context.Background(), marketingLeadId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingLeadsAPI.GetMarketingLeadDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingLeadDetailsAsync`: MarketingLeadDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingLeadsAPI.GetMarketingLeadDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingLeadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingLeadDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**MarketingLeadDtoEnvelope**](MarketingLeadDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingLeadsCountAsync

> Int32Envelope GetMarketingLeadsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).MarketingLeadDtoCollectionQueryParameters(marketingLeadDtoCollectionQueryParameters).Execute()

Get marketing leads count

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
	marketingLeadDtoCollectionQueryParameters := *openapiclient.NewMarketingLeadDtoCollectionQueryParameters() // MarketingLeadDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingLeadsAPI.GetMarketingLeadsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).MarketingLeadDtoCollectionQueryParameters(marketingLeadDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingLeadsAPI.GetMarketingLeadsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingLeadsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingLeadsAPI.GetMarketingLeadsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingLeadsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **marketingLeadDtoCollectionQueryParameters** | [**MarketingLeadDtoCollectionQueryParameters**](MarketingLeadDtoCollectionQueryParameters.md) |  | 

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


## GetMarketingLeadsODataAsync

> MarketingLeadDtoListEnvelope GetMarketingLeadsODataAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).MarketingLeadDtoCollectionQueryParameters(marketingLeadDtoCollectionQueryParameters).Execute()

Get marketing leads



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
	marketingLeadDtoCollectionQueryParameters := *openapiclient.NewMarketingLeadDtoCollectionQueryParameters() // MarketingLeadDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingLeadsAPI.GetMarketingLeadsODataAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).MarketingLeadDtoCollectionQueryParameters(marketingLeadDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingLeadsAPI.GetMarketingLeadsODataAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingLeadsODataAsync`: MarketingLeadDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingLeadsAPI.GetMarketingLeadsODataAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingLeadsODataAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **marketingLeadDtoCollectionQueryParameters** | [**MarketingLeadDtoCollectionQueryParameters**](MarketingLeadDtoCollectionQueryParameters.md) |  | 

### Return type

[**MarketingLeadDtoListEnvelope**](MarketingLeadDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingLeadAsync

> EmptyEnvelope PatchMarketingLeadAsync(ctx, marketingLeadId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a marketing lead



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
	marketingLeadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingLeadsAPI.PatchMarketingLeadAsync(context.Background(), marketingLeadId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingLeadsAPI.PatchMarketingLeadAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingLeadAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingLeadsAPI.PatchMarketingLeadAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingLeadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingLeadAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
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


## UpdateMarketingLeadAsync

> EmptyEnvelope UpdateMarketingLeadAsync(ctx, marketingLeadId).TenantId(tenantId).MarketingLeadUpdateDto(marketingLeadUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update a marketing lead

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
	marketingLeadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	marketingLeadUpdateDto := *openapiclient.NewMarketingLeadUpdateDto() // MarketingLeadUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingLeadsAPI.UpdateMarketingLeadAsync(context.Background(), marketingLeadId).TenantId(tenantId).MarketingLeadUpdateDto(marketingLeadUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingLeadsAPI.UpdateMarketingLeadAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMarketingLeadAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MarketingLeadsAPI.UpdateMarketingLeadAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketingLeadId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMarketingLeadAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **marketingLeadUpdateDto** | [**MarketingLeadUpdateDto**](MarketingLeadUpdateDto.md) |  | 
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

