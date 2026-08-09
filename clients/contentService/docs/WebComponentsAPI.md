# \WebComponentsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountWebComponentsAsync**](WebComponentsAPI.md#CountWebComponentsAsync) | **Get** /api/v2/ContentService/WebComponents/Count | Count web components
[**CreateWebComponentAsync**](WebComponentsAPI.md#CreateWebComponentAsync) | **Post** /api/v2/ContentService/WebComponents | Create a web component
[**DeleteWebComponentAsync**](WebComponentsAPI.md#DeleteWebComponentAsync) | **Delete** /api/v2/ContentService/WebComponents/{webComponentId} | Delete a web component
[**GetWebComponentByIdAsync**](WebComponentsAPI.md#GetWebComponentByIdAsync) | **Get** /api/v2/ContentService/WebComponents/{webComponentId} | Get web component by ID
[**GetWebComponentsAsync**](WebComponentsAPI.md#GetWebComponentsAsync) | **Get** /api/v2/ContentService/WebComponents | Get web components
[**UpdateWebComponentAsync**](WebComponentsAPI.md#UpdateWebComponentAsync) | **Put** /api/v2/ContentService/WebComponents/{webComponentId} | Update a web component



## CountWebComponentsAsync

> Int32Envelope CountWebComponentsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebComponentDtoCollectionQueryParameters(webComponentDtoCollectionQueryParameters).Execute()

Count web components



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
	webComponentDtoCollectionQueryParameters := *openapiclient.NewWebComponentDtoCollectionQueryParameters() // WebComponentDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebComponentsAPI.CountWebComponentsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebComponentDtoCollectionQueryParameters(webComponentDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebComponentsAPI.CountWebComponentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWebComponentsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WebComponentsAPI.CountWebComponentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountWebComponentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webComponentDtoCollectionQueryParameters** | [**WebComponentDtoCollectionQueryParameters**](WebComponentDtoCollectionQueryParameters.md) |  | 

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


## CreateWebComponentAsync

> EmptyEnvelope CreateWebComponentAsync(ctx).TenantId(tenantId).WebComponentCreateDto(webComponentCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create a web component



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
	webComponentCreateDto := *openapiclient.NewWebComponentCreateDto("Name_example") // WebComponentCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebComponentsAPI.CreateWebComponentAsync(context.Background()).TenantId(tenantId).WebComponentCreateDto(webComponentCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebComponentsAPI.CreateWebComponentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebComponentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebComponentsAPI.CreateWebComponentAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebComponentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **webComponentCreateDto** | [**WebComponentCreateDto**](WebComponentCreateDto.md) |  | 
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


## DeleteWebComponentAsync

> EmptyEnvelope DeleteWebComponentAsync(ctx, webComponentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a web component



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
	webComponentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebComponentsAPI.DeleteWebComponentAsync(context.Background(), webComponentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebComponentsAPI.DeleteWebComponentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebComponentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebComponentsAPI.DeleteWebComponentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webComponentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebComponentAsyncRequest struct via the builder pattern


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


## GetWebComponentByIdAsync

> WebComponentDtoEnvelope GetWebComponentByIdAsync(ctx, webComponentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get web component by ID



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
	webComponentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebComponentsAPI.GetWebComponentByIdAsync(context.Background(), webComponentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebComponentsAPI.GetWebComponentByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebComponentByIdAsync`: WebComponentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebComponentsAPI.GetWebComponentByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webComponentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebComponentByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WebComponentDtoEnvelope**](WebComponentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebComponentsAsync

> WebComponentDtoListEnvelope GetWebComponentsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebComponentDtoCollectionQueryParameters(webComponentDtoCollectionQueryParameters).Execute()

Get web components



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
	webComponentDtoCollectionQueryParameters := *openapiclient.NewWebComponentDtoCollectionQueryParameters() // WebComponentDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebComponentsAPI.GetWebComponentsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebComponentDtoCollectionQueryParameters(webComponentDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebComponentsAPI.GetWebComponentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebComponentsAsync`: WebComponentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebComponentsAPI.GetWebComponentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebComponentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webComponentDtoCollectionQueryParameters** | [**WebComponentDtoCollectionQueryParameters**](WebComponentDtoCollectionQueryParameters.md) |  | 

### Return type

[**WebComponentDtoListEnvelope**](WebComponentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebComponentAsync

> EmptyEnvelope UpdateWebComponentAsync(ctx, webComponentId).TenantId(tenantId).WebComponentUpdateDto(webComponentUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update a web component



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
	webComponentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	webComponentUpdateDto := *openapiclient.NewWebComponentUpdateDto() // WebComponentUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebComponentsAPI.UpdateWebComponentAsync(context.Background(), webComponentId).TenantId(tenantId).WebComponentUpdateDto(webComponentUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebComponentsAPI.UpdateWebComponentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebComponentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebComponentsAPI.UpdateWebComponentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webComponentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebComponentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **webComponentUpdateDto** | [**WebComponentUpdateDto**](WebComponentUpdateDto.md) |  | 
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

