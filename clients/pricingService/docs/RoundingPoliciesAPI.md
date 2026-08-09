# \RoundingPoliciesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoundingPolicyAsync**](RoundingPoliciesAPI.md#CreateRoundingPolicyAsync) | **Post** /api/v2/PricingService/RoundingPolicies | Creates a rounding policy
[**DeleteRoundingPolicyAsync**](RoundingPoliciesAPI.md#DeleteRoundingPolicyAsync) | **Delete** /api/v2/PricingService/RoundingPolicies/{roundingPolicyId} | Deletes a rounding policy
[**GetRoundingPoliciesAsync**](RoundingPoliciesAPI.md#GetRoundingPoliciesAsync) | **Get** /api/v2/PricingService/RoundingPolicies | Gets all rounding policies
[**GetRoundingPoliciesCountAsync**](RoundingPoliciesAPI.md#GetRoundingPoliciesCountAsync) | **Get** /api/v2/PricingService/RoundingPolicies/Count | Counts rounding policies
[**GetRoundingPolicyByIdAsync**](RoundingPoliciesAPI.md#GetRoundingPolicyByIdAsync) | **Get** /api/v2/PricingService/RoundingPolicies/{roundingPolicyId} | Gets a rounding policy by ID
[**PatchRoundingPolicyAsync**](RoundingPoliciesAPI.md#PatchRoundingPolicyAsync) | **Patch** /api/v2/PricingService/RoundingPolicies/{roundingPolicyId} | Patches a rounding policy
[**UpdateRoundingPolicyAsync**](RoundingPoliciesAPI.md#UpdateRoundingPolicyAsync) | **Put** /api/v2/PricingService/RoundingPolicies/{roundingPolicyId} | Updates a rounding policy



## CreateRoundingPolicyAsync

> EmptyEnvelope CreateRoundingPolicyAsync(ctx).TenantId(tenantId).RoundingPolicyCreateDto(roundingPolicyCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Creates a rounding policy



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
	roundingPolicyCreateDto := *openapiclient.NewRoundingPolicyCreateDto() // RoundingPolicyCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoundingPoliciesAPI.CreateRoundingPolicyAsync(context.Background()).TenantId(tenantId).RoundingPolicyCreateDto(roundingPolicyCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoundingPoliciesAPI.CreateRoundingPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRoundingPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RoundingPoliciesAPI.CreateRoundingPolicyAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoundingPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **roundingPolicyCreateDto** | [**RoundingPolicyCreateDto**](RoundingPolicyCreateDto.md) |  | 
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


## DeleteRoundingPolicyAsync

> EmptyEnvelope DeleteRoundingPolicyAsync(ctx, roundingPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a rounding policy



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
	roundingPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoundingPoliciesAPI.DeleteRoundingPolicyAsync(context.Background(), roundingPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoundingPoliciesAPI.DeleteRoundingPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRoundingPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RoundingPoliciesAPI.DeleteRoundingPolicyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roundingPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoundingPolicyAsyncRequest struct via the builder pattern


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


## GetRoundingPoliciesAsync

> RoundingPolicyDtoListEnvelope GetRoundingPoliciesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RoundingPolicyDtoCollectionQueryParameters(roundingPolicyDtoCollectionQueryParameters).Execute()

Gets all rounding policies



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
	roundingPolicyDtoCollectionQueryParameters := *openapiclient.NewRoundingPolicyDtoCollectionQueryParameters() // RoundingPolicyDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoundingPoliciesAPI.GetRoundingPoliciesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RoundingPolicyDtoCollectionQueryParameters(roundingPolicyDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoundingPoliciesAPI.GetRoundingPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoundingPoliciesAsync`: RoundingPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RoundingPoliciesAPI.GetRoundingPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoundingPoliciesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **roundingPolicyDtoCollectionQueryParameters** | [**RoundingPolicyDtoCollectionQueryParameters**](RoundingPolicyDtoCollectionQueryParameters.md) |  | 

### Return type

[**RoundingPolicyDtoListEnvelope**](RoundingPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoundingPoliciesCountAsync

> Int32Envelope GetRoundingPoliciesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RoundingPolicyDtoCollectionQueryParameters(roundingPolicyDtoCollectionQueryParameters).Execute()

Counts rounding policies



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
	roundingPolicyDtoCollectionQueryParameters := *openapiclient.NewRoundingPolicyDtoCollectionQueryParameters() // RoundingPolicyDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoundingPoliciesAPI.GetRoundingPoliciesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RoundingPolicyDtoCollectionQueryParameters(roundingPolicyDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoundingPoliciesAPI.GetRoundingPoliciesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoundingPoliciesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `RoundingPoliciesAPI.GetRoundingPoliciesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoundingPoliciesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **roundingPolicyDtoCollectionQueryParameters** | [**RoundingPolicyDtoCollectionQueryParameters**](RoundingPolicyDtoCollectionQueryParameters.md) |  | 

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


## GetRoundingPolicyByIdAsync

> RoundingPolicyDtoEnvelope GetRoundingPolicyByIdAsync(ctx, roundingPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets a rounding policy by ID



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
	roundingPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoundingPoliciesAPI.GetRoundingPolicyByIdAsync(context.Background(), roundingPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoundingPoliciesAPI.GetRoundingPolicyByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoundingPolicyByIdAsync`: RoundingPolicyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RoundingPoliciesAPI.GetRoundingPolicyByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roundingPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoundingPolicyByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**RoundingPolicyDtoEnvelope**](RoundingPolicyDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRoundingPolicyAsync

> EmptyEnvelope PatchRoundingPolicyAsync(ctx, roundingPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patches a rounding policy



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
	roundingPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoundingPoliciesAPI.PatchRoundingPolicyAsync(context.Background(), roundingPolicyId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoundingPoliciesAPI.PatchRoundingPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRoundingPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RoundingPoliciesAPI.PatchRoundingPolicyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roundingPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRoundingPolicyAsyncRequest struct via the builder pattern


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


## UpdateRoundingPolicyAsync

> EmptyEnvelope UpdateRoundingPolicyAsync(ctx, roundingPolicyId).TenantId(tenantId).RoundingPolicyUpdateDto(roundingPolicyUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Updates a rounding policy



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
	roundingPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	roundingPolicyUpdateDto := *openapiclient.NewRoundingPolicyUpdateDto() // RoundingPolicyUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoundingPoliciesAPI.UpdateRoundingPolicyAsync(context.Background(), roundingPolicyId).TenantId(tenantId).RoundingPolicyUpdateDto(roundingPolicyUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoundingPoliciesAPI.UpdateRoundingPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoundingPolicyAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RoundingPoliciesAPI.UpdateRoundingPolicyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roundingPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoundingPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **roundingPolicyUpdateDto** | [**RoundingPolicyUpdateDto**](RoundingPolicyUpdateDto.md) |  | 
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

