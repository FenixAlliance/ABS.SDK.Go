# \PortalsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystemPortal**](PortalsAPI.md#CreateSystemPortal) | **Post** /api/v2/SystemService/Portals | Create a new system portal
[**DeleteSystemPortal**](PortalsAPI.md#DeleteSystemPortal) | **Delete** /api/v2/SystemService/Portals/{portalId} | Delete a system portal
[**GetSystemPortalById**](PortalsAPI.md#GetSystemPortalById) | **Get** /api/v2/SystemService/Portals/{portalId} | Retrieve a single system portal by its ID
[**GetSystemPortals**](PortalsAPI.md#GetSystemPortals) | **Get** /api/v2/SystemService/Portals | Retrieve a list of system portals
[**GetSystemPortalsCount**](PortalsAPI.md#GetSystemPortalsCount) | **Get** /api/v2/SystemService/Portals/Count | Get the count of system portals
[**PatchSystemPortal**](PortalsAPI.md#PatchSystemPortal) | **Patch** /api/v2/SystemService/Portals/{portalId} | Partially update a system portal
[**UpdateSystemPortal**](PortalsAPI.md#UpdateSystemPortal) | **Put** /api/v2/SystemService/Portals/{portalId} | Update a system portal



## CreateSystemPortal

> EmptyEnvelope CreateSystemPortal(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPortalCreateDto(webPortalCreateDto).Execute()

Create a new system portal



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	webPortalCreateDto := *openapiclient.NewWebPortalCreateDto() // WebPortalCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalsAPI.CreateSystemPortal(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPortalCreateDto(webPortalCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalsAPI.CreateSystemPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSystemPortal`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PortalsAPI.CreateSystemPortal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webPortalCreateDto** | [**WebPortalCreateDto**](WebPortalCreateDto.md) |  | 

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


## DeleteSystemPortal

> EmptyEnvelope DeleteSystemPortal(ctx, portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a system portal



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalsAPI.DeleteSystemPortal(context.Background(), portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalsAPI.DeleteSystemPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSystemPortal`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PortalsAPI.DeleteSystemPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetSystemPortalById

> WebPortalDtoEnvelope GetSystemPortalById(ctx, portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single system portal by its ID



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalsAPI.GetSystemPortalById(context.Background(), portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalsAPI.GetSystemPortalById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemPortalById`: WebPortalDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PortalsAPI.GetSystemPortalById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemPortalByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WebPortalDtoEnvelope**](WebPortalDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemPortals

> WebPortalDtoListEnvelope GetSystemPortals(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPortalDtoCollectionQueryParameters(webPortalDtoCollectionQueryParameters).Execute()

Retrieve a list of system portals



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	webPortalDtoCollectionQueryParameters := *openapiclient.NewWebPortalDtoCollectionQueryParameters() // WebPortalDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalsAPI.GetSystemPortals(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPortalDtoCollectionQueryParameters(webPortalDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalsAPI.GetSystemPortals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemPortals`: WebPortalDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PortalsAPI.GetSystemPortals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemPortalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webPortalDtoCollectionQueryParameters** | [**WebPortalDtoCollectionQueryParameters**](WebPortalDtoCollectionQueryParameters.md) |  | 

### Return type

[**WebPortalDtoListEnvelope**](WebPortalDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemPortalsCount

> Int32Envelope GetSystemPortalsCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPortalDtoCollectionQueryParameters(webPortalDtoCollectionQueryParameters).Execute()

Get the count of system portals



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	webPortalDtoCollectionQueryParameters := *openapiclient.NewWebPortalDtoCollectionQueryParameters() // WebPortalDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalsAPI.GetSystemPortalsCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPortalDtoCollectionQueryParameters(webPortalDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalsAPI.GetSystemPortalsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemPortalsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `PortalsAPI.GetSystemPortalsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemPortalsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webPortalDtoCollectionQueryParameters** | [**WebPortalDtoCollectionQueryParameters**](WebPortalDtoCollectionQueryParameters.md) |  | 

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


## PatchSystemPortal

> EmptyEnvelope PatchSystemPortal(ctx, portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Partially update a system portal



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalsAPI.PatchSystemPortal(context.Background(), portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalsAPI.PatchSystemPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemPortal`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PortalsAPI.PatchSystemPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## UpdateSystemPortal

> EmptyEnvelope UpdateSystemPortal(ctx, portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPortalUpdateDto(webPortalUpdateDto).Execute()

Update a system portal



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	webPortalUpdateDto := *openapiclient.NewWebPortalUpdateDto() // WebPortalUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortalsAPI.UpdateSystemPortal(context.Background(), portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPortalUpdateDto(webPortalUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortalsAPI.UpdateSystemPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSystemPortal`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PortalsAPI.UpdateSystemPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webPortalUpdateDto** | [**WebPortalUpdateDto**](WebPortalUpdateDto.md) |  | 

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

