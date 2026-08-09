# \ShippingMethodsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShippingMethodAsync**](ShippingMethodsAPI.md#CreateShippingMethodAsync) | **Post** /api/v2/ShipmentsService/ShippingMethods | Create a shipping method
[**DeleteShippingMethodAsync**](ShippingMethodsAPI.md#DeleteShippingMethodAsync) | **Delete** /api/v2/ShipmentsService/ShippingMethods/{methodId} | Delete a shipping method
[**GetShippingMethodByIdAsync**](ShippingMethodsAPI.md#GetShippingMethodByIdAsync) | **Get** /api/v2/ShipmentsService/ShippingMethods/{methodId} | Get shipping method by ID
[**GetShippingMethodsAsync**](ShippingMethodsAPI.md#GetShippingMethodsAsync) | **Get** /api/v2/ShipmentsService/ShippingMethods | Get all shipping methods
[**GetShippingMethodsCountAsync**](ShippingMethodsAPI.md#GetShippingMethodsCountAsync) | **Get** /api/v2/ShipmentsService/ShippingMethods/Count | Get shipping methods count
[**PatchShippingMethodAsync**](ShippingMethodsAPI.md#PatchShippingMethodAsync) | **Patch** /api/v2/ShipmentsService/ShippingMethods/{methodId} | Patch a shipping method
[**UpdateShippingMethodAsync**](ShippingMethodsAPI.md#UpdateShippingMethodAsync) | **Put** /api/v2/ShipmentsService/ShippingMethods/{methodId} | Update a shipping method



## CreateShippingMethodAsync

> CreateShippingMethodAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ShippingMethodCreateDto(shippingMethodCreateDto).Execute()

Create a shipping method



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
	shippingMethodCreateDto := *openapiclient.NewShippingMethodCreateDto("Name_example") // ShippingMethodCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShippingMethodsAPI.CreateShippingMethodAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ShippingMethodCreateDto(shippingMethodCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsAPI.CreateShippingMethodAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShippingMethodAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **shippingMethodCreateDto** | [**ShippingMethodCreateDto**](ShippingMethodCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShippingMethodAsync

> DeleteShippingMethodAsync(ctx, methodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a shipping method



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
	methodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShippingMethodsAPI.DeleteShippingMethodAsync(context.Background(), methodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsAPI.DeleteShippingMethodAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShippingMethodAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShippingMethodByIdAsync

> ShippingMethodDtoEnvelope GetShippingMethodByIdAsync(ctx, methodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get shipping method by ID



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
	methodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShippingMethodsAPI.GetShippingMethodByIdAsync(context.Background(), methodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsAPI.GetShippingMethodByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShippingMethodByIdAsync`: ShippingMethodDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ShippingMethodsAPI.GetShippingMethodByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShippingMethodByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ShippingMethodDtoEnvelope**](ShippingMethodDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShippingMethodsAsync

> ShippingMethodDtoListEnvelope GetShippingMethodsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ShippingMethodDtoCollectionQueryParameters(shippingMethodDtoCollectionQueryParameters).Execute()

Get all shipping methods



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
	shippingMethodDtoCollectionQueryParameters := *openapiclient.NewShippingMethodDtoCollectionQueryParameters() // ShippingMethodDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShippingMethodsAPI.GetShippingMethodsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ShippingMethodDtoCollectionQueryParameters(shippingMethodDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsAPI.GetShippingMethodsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShippingMethodsAsync`: ShippingMethodDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ShippingMethodsAPI.GetShippingMethodsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShippingMethodsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **shippingMethodDtoCollectionQueryParameters** | [**ShippingMethodDtoCollectionQueryParameters**](ShippingMethodDtoCollectionQueryParameters.md) |  | 

### Return type

[**ShippingMethodDtoListEnvelope**](ShippingMethodDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShippingMethodsCountAsync

> Int32Envelope GetShippingMethodsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ShippingMethodDtoCollectionQueryParameters(shippingMethodDtoCollectionQueryParameters).Execute()

Get shipping methods count



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
	shippingMethodDtoCollectionQueryParameters := *openapiclient.NewShippingMethodDtoCollectionQueryParameters() // ShippingMethodDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShippingMethodsAPI.GetShippingMethodsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ShippingMethodDtoCollectionQueryParameters(shippingMethodDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsAPI.GetShippingMethodsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShippingMethodsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ShippingMethodsAPI.GetShippingMethodsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShippingMethodsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **shippingMethodDtoCollectionQueryParameters** | [**ShippingMethodDtoCollectionQueryParameters**](ShippingMethodDtoCollectionQueryParameters.md) |  | 

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


## PatchShippingMethodAsync

> EmptyEnvelope PatchShippingMethodAsync(ctx, methodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a shipping method



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
	methodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShippingMethodsAPI.PatchShippingMethodAsync(context.Background(), methodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsAPI.PatchShippingMethodAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchShippingMethodAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ShippingMethodsAPI.PatchShippingMethodAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchShippingMethodAsyncRequest struct via the builder pattern


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


## UpdateShippingMethodAsync

> UpdateShippingMethodAsync(ctx, methodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ShippingMethodUpdateDto(shippingMethodUpdateDto).Execute()

Update a shipping method



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
	methodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	shippingMethodUpdateDto := *openapiclient.NewShippingMethodUpdateDto() // ShippingMethodUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShippingMethodsAPI.UpdateShippingMethodAsync(context.Background(), methodId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ShippingMethodUpdateDto(shippingMethodUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodsAPI.UpdateShippingMethodAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**methodId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShippingMethodAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **shippingMethodUpdateDto** | [**ShippingMethodUpdateDto**](ShippingMethodUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

