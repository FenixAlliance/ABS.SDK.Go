# \CognitiveAgentVariablesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCognitiveAgentVariableAsync**](CognitiveAgentVariablesAPI.md#CreateCognitiveAgentVariableAsync) | **Post** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Variables | Add a variable to a cognitive agent
[**DeleteCognitiveAgentVariableAsync**](CognitiveAgentVariablesAPI.md#DeleteCognitiveAgentVariableAsync) | **Delete** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Variables/{id} | Remove a variable from a cognitive agent
[**GetCognitiveAgentVariableByIdAsync**](CognitiveAgentVariablesAPI.md#GetCognitiveAgentVariableByIdAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Variables/{id} | Get a cognitive agent variable by ID
[**GetCognitiveAgentVariablesAsync**](CognitiveAgentVariablesAPI.md#GetCognitiveAgentVariablesAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Variables | Get all variables for a cognitive agent
[**GetCognitiveAgentVariablesCountAsync**](CognitiveAgentVariablesAPI.md#GetCognitiveAgentVariablesCountAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Variables/Count | Get variable count for a cognitive agent
[**UpdateCognitiveAgentVariableAsync**](CognitiveAgentVariablesAPI.md#UpdateCognitiveAgentVariableAsync) | **Put** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Variables/{id} | Update a cognitive agent variable



## CreateCognitiveAgentVariableAsync

> CreateCognitiveAgentVariableAsync(ctx, agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentVariableCreateDto(cognitiveAgentVariableCreateDto).Execute()

Add a variable to a cognitive agent



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	cognitiveAgentVariableCreateDto := *openapiclient.NewCognitiveAgentVariableCreateDto("Key_example") // CognitiveAgentVariableCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CognitiveAgentVariablesAPI.CreateCognitiveAgentVariableAsync(context.Background(), agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentVariableCreateDto(cognitiveAgentVariableCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentVariablesAPI.CreateCognitiveAgentVariableAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCognitiveAgentVariableAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentVariableCreateDto** | [**CognitiveAgentVariableCreateDto**](CognitiveAgentVariableCreateDto.md) |  | 

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


## DeleteCognitiveAgentVariableAsync

> DeleteCognitiveAgentVariableAsync(ctx, agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove a variable from a cognitive agent



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CognitiveAgentVariablesAPI.DeleteCognitiveAgentVariableAsync(context.Background(), agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentVariablesAPI.DeleteCognitiveAgentVariableAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCognitiveAgentVariableAsyncRequest struct via the builder pattern


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


## GetCognitiveAgentVariableByIdAsync

> CognitiveAgentVariableDtoEnvelope GetCognitiveAgentVariableByIdAsync(ctx, agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get a cognitive agent variable by ID



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentVariablesAPI.GetCognitiveAgentVariableByIdAsync(context.Background(), agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentVariablesAPI.GetCognitiveAgentVariableByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentVariableByIdAsync`: CognitiveAgentVariableDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentVariablesAPI.GetCognitiveAgentVariableByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentVariableByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CognitiveAgentVariableDtoEnvelope**](CognitiveAgentVariableDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCognitiveAgentVariablesAsync

> CognitiveAgentVariableDtoListEnvelope GetCognitiveAgentVariablesAsync(ctx, agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentVariableDtoCollectionQueryParameters(cognitiveAgentVariableDtoCollectionQueryParameters).Execute()

Get all variables for a cognitive agent



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	cognitiveAgentVariableDtoCollectionQueryParameters := *openapiclient.NewCognitiveAgentVariableDtoCollectionQueryParameters() // CognitiveAgentVariableDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentVariablesAPI.GetCognitiveAgentVariablesAsync(context.Background(), agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentVariableDtoCollectionQueryParameters(cognitiveAgentVariableDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentVariablesAPI.GetCognitiveAgentVariablesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentVariablesAsync`: CognitiveAgentVariableDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentVariablesAPI.GetCognitiveAgentVariablesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentVariablesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentVariableDtoCollectionQueryParameters** | [**CognitiveAgentVariableDtoCollectionQueryParameters**](CognitiveAgentVariableDtoCollectionQueryParameters.md) |  | 

### Return type

[**CognitiveAgentVariableDtoListEnvelope**](CognitiveAgentVariableDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCognitiveAgentVariablesCountAsync

> Int32Envelope GetCognitiveAgentVariablesCountAsync(ctx, agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentVariableDtoCollectionQueryParameters(cognitiveAgentVariableDtoCollectionQueryParameters).Execute()

Get variable count for a cognitive agent



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	cognitiveAgentVariableDtoCollectionQueryParameters := *openapiclient.NewCognitiveAgentVariableDtoCollectionQueryParameters() // CognitiveAgentVariableDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentVariablesAPI.GetCognitiveAgentVariablesCountAsync(context.Background(), agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentVariableDtoCollectionQueryParameters(cognitiveAgentVariableDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentVariablesAPI.GetCognitiveAgentVariablesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentVariablesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentVariablesAPI.GetCognitiveAgentVariablesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentVariablesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentVariableDtoCollectionQueryParameters** | [**CognitiveAgentVariableDtoCollectionQueryParameters**](CognitiveAgentVariableDtoCollectionQueryParameters.md) |  | 

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


## UpdateCognitiveAgentVariableAsync

> UpdateCognitiveAgentVariableAsync(ctx, agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentVariableUpdateDto(cognitiveAgentVariableUpdateDto).Execute()

Update a cognitive agent variable



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
	agentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	cognitiveAgentVariableUpdateDto := *openapiclient.NewCognitiveAgentVariableUpdateDto() // CognitiveAgentVariableUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CognitiveAgentVariablesAPI.UpdateCognitiveAgentVariableAsync(context.Background(), agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentVariableUpdateDto(cognitiveAgentVariableUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentVariablesAPI.UpdateCognitiveAgentVariableAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCognitiveAgentVariableAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentVariableUpdateDto** | [**CognitiveAgentVariableUpdateDto**](CognitiveAgentVariableUpdateDto.md) |  | 

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

