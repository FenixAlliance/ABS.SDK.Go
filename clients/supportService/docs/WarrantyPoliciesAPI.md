# \WarrantyPoliciesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWarrantyPolicyAsync**](WarrantyPoliciesAPI.md#CreateWarrantyPolicyAsync) | **Post** /api/v2/SupportService/WarrantyPolicies | Create a new warranty policy
[**DeleteWarrantyPolicyAsync**](WarrantyPoliciesAPI.md#DeleteWarrantyPolicyAsync) | **Delete** /api/v2/SupportService/WarrantyPolicies/{warrantyPolicyId} | Delete a warranty policy
[**GetWarrantyPoliciesAsync**](WarrantyPoliciesAPI.md#GetWarrantyPoliciesAsync) | **Get** /api/v2/SupportService/WarrantyPolicies | Retrieve a list of warranty policies
[**GetWarrantyPoliciesCountAsync**](WarrantyPoliciesAPI.md#GetWarrantyPoliciesCountAsync) | **Get** /api/v2/SupportService/WarrantyPolicies/Count | Get the count of warranty policies
[**GetWarrantyPolicyAsync**](WarrantyPoliciesAPI.md#GetWarrantyPolicyAsync) | **Get** /api/v2/SupportService/WarrantyPolicies/{warrantyPolicyId} | Retrieve a warranty policy by ID
[**PatchWarrantyPolicyAsync**](WarrantyPoliciesAPI.md#PatchWarrantyPolicyAsync) | **Patch** /api/v2/SupportService/WarrantyPolicies/{warrantyPolicyId} | Patch a warranty policy
[**UpdateWarrantyPolicyAsync**](WarrantyPoliciesAPI.md#UpdateWarrantyPolicyAsync) | **Put** /api/v2/SupportService/WarrantyPolicies/{warrantyPolicyId} | Update a warranty policy



## CreateWarrantyPolicyAsync

> EmptyEnvelope CreateWarrantyPolicyAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemWarrantyPolicyCreateDto(itemWarrantyPolicyCreateDto).Execute()

Create a new warranty policy

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
	itemWarrantyPolicyCreateDto := *openapiclient.NewItemWarrantyPolicyCreateDto("Title_example") // ItemWarrantyPolicyCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyPoliciesAPI.CreateWarrantyPolicyAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemWarrantyPolicyCreateDto(itemWarrantyPolicyCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyPoliciesAPI.CreateWarrantyPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWarrantyPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyPoliciesAPI.CreateWarrantyPolicyAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWarrantyPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemWarrantyPolicyCreateDto** | [**ItemWarrantyPolicyCreateDto**](ItemWarrantyPolicyCreateDto.md) |  | 

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


## DeleteWarrantyPolicyAsync

> EmptyEnvelope DeleteWarrantyPolicyAsync(ctx, warrantyPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a warranty policy

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
	warrantyPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyPoliciesAPI.DeleteWarrantyPolicyAsync(context.Background(), warrantyPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyPoliciesAPI.DeleteWarrantyPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWarrantyPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyPoliciesAPI.DeleteWarrantyPolicyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantyPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWarrantyPolicyAsyncRequest struct via the builder pattern


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


## GetWarrantyPoliciesAsync

> ItemWarrantyPolicyDtoListEnvelope GetWarrantyPoliciesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemWarrantyPolicyDtoCollectionQueryParameters(itemWarrantyPolicyDtoCollectionQueryParameters).Execute()

Retrieve a list of warranty policies

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
	itemWarrantyPolicyDtoCollectionQueryParameters := *openapiclient.NewItemWarrantyPolicyDtoCollectionQueryParameters() // ItemWarrantyPolicyDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyPoliciesAPI.GetWarrantyPoliciesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemWarrantyPolicyDtoCollectionQueryParameters(itemWarrantyPolicyDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyPoliciesAPI.GetWarrantyPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarrantyPoliciesAsync`: ItemWarrantyPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyPoliciesAPI.GetWarrantyPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWarrantyPoliciesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemWarrantyPolicyDtoCollectionQueryParameters** | [**ItemWarrantyPolicyDtoCollectionQueryParameters**](ItemWarrantyPolicyDtoCollectionQueryParameters.md) |  | 

### Return type

[**ItemWarrantyPolicyDtoListEnvelope**](ItemWarrantyPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWarrantyPoliciesCountAsync

> Int32Envelope GetWarrantyPoliciesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemWarrantyPolicyDtoCollectionQueryParameters(itemWarrantyPolicyDtoCollectionQueryParameters).Execute()

Get the count of warranty policies

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
	itemWarrantyPolicyDtoCollectionQueryParameters := *openapiclient.NewItemWarrantyPolicyDtoCollectionQueryParameters() // ItemWarrantyPolicyDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyPoliciesAPI.GetWarrantyPoliciesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemWarrantyPolicyDtoCollectionQueryParameters(itemWarrantyPolicyDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyPoliciesAPI.GetWarrantyPoliciesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarrantyPoliciesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyPoliciesAPI.GetWarrantyPoliciesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWarrantyPoliciesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemWarrantyPolicyDtoCollectionQueryParameters** | [**ItemWarrantyPolicyDtoCollectionQueryParameters**](ItemWarrantyPolicyDtoCollectionQueryParameters.md) |  | 

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


## GetWarrantyPolicyAsync

> ItemWarrantyPolicyDtoEnvelope GetWarrantyPolicyAsync(ctx, warrantyPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a warranty policy by ID

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
	warrantyPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyPoliciesAPI.GetWarrantyPolicyAsync(context.Background(), warrantyPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyPoliciesAPI.GetWarrantyPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWarrantyPolicyAsync`: ItemWarrantyPolicyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyPoliciesAPI.GetWarrantyPolicyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantyPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWarrantyPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemWarrantyPolicyDtoEnvelope**](ItemWarrantyPolicyDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWarrantyPolicyAsync

> EmptyEnvelope PatchWarrantyPolicyAsync(ctx, warrantyPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a warranty policy



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
	warrantyPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyPoliciesAPI.PatchWarrantyPolicyAsync(context.Background(), warrantyPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyPoliciesAPI.PatchWarrantyPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWarrantyPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyPoliciesAPI.PatchWarrantyPolicyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantyPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWarrantyPolicyAsyncRequest struct via the builder pattern


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


## UpdateWarrantyPolicyAsync

> EmptyEnvelope UpdateWarrantyPolicyAsync(ctx, warrantyPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemWarrantyPolicyUpdateDto(itemWarrantyPolicyUpdateDto).Execute()

Update a warranty policy

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
	warrantyPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemWarrantyPolicyUpdateDto := *openapiclient.NewItemWarrantyPolicyUpdateDto() // ItemWarrantyPolicyUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarrantyPoliciesAPI.UpdateWarrantyPolicyAsync(context.Background(), warrantyPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemWarrantyPolicyUpdateDto(itemWarrantyPolicyUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarrantyPoliciesAPI.UpdateWarrantyPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWarrantyPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WarrantyPoliciesAPI.UpdateWarrantyPolicyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warrantyPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWarrantyPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemWarrantyPolicyUpdateDto** | [**ItemWarrantyPolicyUpdateDto**](ItemWarrantyPolicyUpdateDto.md) |  | 

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

