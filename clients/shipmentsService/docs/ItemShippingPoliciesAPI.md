# \ItemShippingPoliciesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemShippingPolicyAsync**](ItemShippingPoliciesAPI.md#CreateItemShippingPolicyAsync) | **Post** /api/v2/ShipmentsService/ItemShippingPolicies | Create an item shipping policy
[**DeleteItemShippingPolicyAsync**](ItemShippingPoliciesAPI.md#DeleteItemShippingPolicyAsync) | **Delete** /api/v2/ShipmentsService/ItemShippingPolicies/{policyId} | Delete an item shipping policy
[**GetItemShippingPoliciesAsync**](ItemShippingPoliciesAPI.md#GetItemShippingPoliciesAsync) | **Get** /api/v2/ShipmentsService/ItemShippingPolicies | Get all item shipping policies
[**GetItemShippingPoliciesCountAsync**](ItemShippingPoliciesAPI.md#GetItemShippingPoliciesCountAsync) | **Get** /api/v2/ShipmentsService/ItemShippingPolicies/Count | Get item shipping policies count
[**GetItemShippingPolicyByIdAsync**](ItemShippingPoliciesAPI.md#GetItemShippingPolicyByIdAsync) | **Get** /api/v2/ShipmentsService/ItemShippingPolicies/{policyId} | Get item shipping policy by ID
[**PatchItemShippingPolicyAsync**](ItemShippingPoliciesAPI.md#PatchItemShippingPolicyAsync) | **Patch** /api/v2/ShipmentsService/ItemShippingPolicies/{policyId} | Patch an item shipping policy
[**UpdateItemShippingPolicyAsync**](ItemShippingPoliciesAPI.md#UpdateItemShippingPolicyAsync) | **Put** /api/v2/ShipmentsService/ItemShippingPolicies/{policyId} | Update an item shipping policy



## CreateItemShippingPolicyAsync

> CreateItemShippingPolicyAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemShippingPolicyCreateDto(itemShippingPolicyCreateDto).Execute()

Create an item shipping policy



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
	itemShippingPolicyCreateDto := *openapiclient.NewItemShippingPolicyCreateDto("Title_example", "Type_example", "Code_example", "CurrencyID_example", "ShippingCourierID_example") // ItemShippingPolicyCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemShippingPoliciesAPI.CreateItemShippingPolicyAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemShippingPolicyCreateDto(itemShippingPolicyCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemShippingPoliciesAPI.CreateItemShippingPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemShippingPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemShippingPolicyCreateDto** | [**ItemShippingPolicyCreateDto**](ItemShippingPolicyCreateDto.md) |  | 

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


## DeleteItemShippingPolicyAsync

> DeleteItemShippingPolicyAsync(ctx, policyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an item shipping policy



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
	policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemShippingPoliciesAPI.DeleteItemShippingPolicyAsync(context.Background(), policyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemShippingPoliciesAPI.DeleteItemShippingPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemShippingPolicyAsyncRequest struct via the builder pattern


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


## GetItemShippingPoliciesAsync

> ItemShippingPolicyDtoListEnvelope GetItemShippingPoliciesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemShippingPolicyDtoCollectionQueryParameters(itemShippingPolicyDtoCollectionQueryParameters).Execute()

Get all item shipping policies



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
	itemShippingPolicyDtoCollectionQueryParameters := *openapiclient.NewItemShippingPolicyDtoCollectionQueryParameters() // ItemShippingPolicyDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemShippingPoliciesAPI.GetItemShippingPoliciesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemShippingPolicyDtoCollectionQueryParameters(itemShippingPolicyDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemShippingPoliciesAPI.GetItemShippingPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemShippingPoliciesAsync`: ItemShippingPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemShippingPoliciesAPI.GetItemShippingPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemShippingPoliciesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemShippingPolicyDtoCollectionQueryParameters** | [**ItemShippingPolicyDtoCollectionQueryParameters**](ItemShippingPolicyDtoCollectionQueryParameters.md) |  | 

### Return type

[**ItemShippingPolicyDtoListEnvelope**](ItemShippingPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemShippingPoliciesCountAsync

> Int32Envelope GetItemShippingPoliciesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemShippingPolicyDtoCollectionQueryParameters(itemShippingPolicyDtoCollectionQueryParameters).Execute()

Get item shipping policies count



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
	itemShippingPolicyDtoCollectionQueryParameters := *openapiclient.NewItemShippingPolicyDtoCollectionQueryParameters() // ItemShippingPolicyDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemShippingPoliciesAPI.GetItemShippingPoliciesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemShippingPolicyDtoCollectionQueryParameters(itemShippingPolicyDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemShippingPoliciesAPI.GetItemShippingPoliciesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemShippingPoliciesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemShippingPoliciesAPI.GetItemShippingPoliciesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemShippingPoliciesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemShippingPolicyDtoCollectionQueryParameters** | [**ItemShippingPolicyDtoCollectionQueryParameters**](ItemShippingPolicyDtoCollectionQueryParameters.md) |  | 

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


## GetItemShippingPolicyByIdAsync

> ItemShippingPolicyDtoEnvelope GetItemShippingPolicyByIdAsync(ctx, policyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item shipping policy by ID



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
	policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemShippingPoliciesAPI.GetItemShippingPolicyByIdAsync(context.Background(), policyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemShippingPoliciesAPI.GetItemShippingPolicyByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemShippingPolicyByIdAsync`: ItemShippingPolicyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemShippingPoliciesAPI.GetItemShippingPolicyByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemShippingPolicyByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemShippingPolicyDtoEnvelope**](ItemShippingPolicyDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchItemShippingPolicyAsync

> EmptyEnvelope PatchItemShippingPolicyAsync(ctx, policyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch an item shipping policy



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
	policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemShippingPoliciesAPI.PatchItemShippingPolicyAsync(context.Background(), policyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemShippingPoliciesAPI.PatchItemShippingPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchItemShippingPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemShippingPoliciesAPI.PatchItemShippingPolicyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchItemShippingPolicyAsyncRequest struct via the builder pattern


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


## UpdateItemShippingPolicyAsync

> UpdateItemShippingPolicyAsync(ctx, policyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemShippingPolicyUpdateDto(itemShippingPolicyUpdateDto).Execute()

Update an item shipping policy



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
	policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	itemShippingPolicyUpdateDto := *openapiclient.NewItemShippingPolicyUpdateDto("Title_example", "Type_example", "Code_example", "CurrencyID_example", "ShippingCourierID_example") // ItemShippingPolicyUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemShippingPoliciesAPI.UpdateItemShippingPolicyAsync(context.Background(), policyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ItemShippingPolicyUpdateDto(itemShippingPolicyUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemShippingPoliciesAPI.UpdateItemShippingPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemShippingPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **itemShippingPolicyUpdateDto** | [**ItemShippingPolicyUpdateDto**](ItemShippingPolicyUpdateDto.md) |  | 

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

