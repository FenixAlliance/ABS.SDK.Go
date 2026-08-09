# \CognitiveAgentSkillsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCognitiveAgentSkillAsync**](CognitiveAgentSkillsAPI.md#CreateCognitiveAgentSkillAsync) | **Post** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Skills | Assign a skill to a cognitive agent
[**DeleteCognitiveAgentSkillAsync**](CognitiveAgentSkillsAPI.md#DeleteCognitiveAgentSkillAsync) | **Delete** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Skills/{id} | Remove a skill assignment from a cognitive agent
[**GetCognitiveAgentSkillByIdAsync**](CognitiveAgentSkillsAPI.md#GetCognitiveAgentSkillByIdAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Skills/{id} | Get a cognitive agent skill assignment by ID
[**GetCognitiveAgentSkillsAsync**](CognitiveAgentSkillsAPI.md#GetCognitiveAgentSkillsAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Skills | Get all skill assignments for a cognitive agent
[**GetCognitiveAgentSkillsCountAsync**](CognitiveAgentSkillsAPI.md#GetCognitiveAgentSkillsCountAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Skills/Count | Get skill assignment count for a cognitive agent
[**UpdateCognitiveAgentSkillAsync**](CognitiveAgentSkillsAPI.md#UpdateCognitiveAgentSkillAsync) | **Put** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Skills/{id} | Update a cognitive agent skill assignment



## CreateCognitiveAgentSkillAsync

> CreateCognitiveAgentSkillAsync(ctx, agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentSkillCreateDto(cognitiveAgentSkillCreateDto).Execute()

Assign a skill to a cognitive agent



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
	cognitiveAgentSkillCreateDto := *openapiclient.NewCognitiveAgentSkillCreateDto("CognitiveSkillId_example") // CognitiveAgentSkillCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CognitiveAgentSkillsAPI.CreateCognitiveAgentSkillAsync(context.Background(), agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentSkillCreateDto(cognitiveAgentSkillCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentSkillsAPI.CreateCognitiveAgentSkillAsync``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateCognitiveAgentSkillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentSkillCreateDto** | [**CognitiveAgentSkillCreateDto**](CognitiveAgentSkillCreateDto.md) |  | 

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


## DeleteCognitiveAgentSkillAsync

> DeleteCognitiveAgentSkillAsync(ctx, agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove a skill assignment from a cognitive agent



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
	r, err := apiClient.CognitiveAgentSkillsAPI.DeleteCognitiveAgentSkillAsync(context.Background(), agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentSkillsAPI.DeleteCognitiveAgentSkillAsync``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCognitiveAgentSkillAsyncRequest struct via the builder pattern


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


## GetCognitiveAgentSkillByIdAsync

> CognitiveAgentSkillDtoEnvelope GetCognitiveAgentSkillByIdAsync(ctx, agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get a cognitive agent skill assignment by ID



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
	resp, r, err := apiClient.CognitiveAgentSkillsAPI.GetCognitiveAgentSkillByIdAsync(context.Background(), agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentSkillsAPI.GetCognitiveAgentSkillByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentSkillByIdAsync`: CognitiveAgentSkillDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentSkillsAPI.GetCognitiveAgentSkillByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentSkillByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CognitiveAgentSkillDtoEnvelope**](CognitiveAgentSkillDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCognitiveAgentSkillsAsync

> CognitiveAgentSkillDtoListEnvelope GetCognitiveAgentSkillsAsync(ctx, agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentSkillDtoCollectionQueryParameters(cognitiveAgentSkillDtoCollectionQueryParameters).Execute()

Get all skill assignments for a cognitive agent



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
	cognitiveAgentSkillDtoCollectionQueryParameters := *openapiclient.NewCognitiveAgentSkillDtoCollectionQueryParameters() // CognitiveAgentSkillDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentSkillsAPI.GetCognitiveAgentSkillsAsync(context.Background(), agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentSkillDtoCollectionQueryParameters(cognitiveAgentSkillDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentSkillsAPI.GetCognitiveAgentSkillsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentSkillsAsync`: CognitiveAgentSkillDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentSkillsAPI.GetCognitiveAgentSkillsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentSkillsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentSkillDtoCollectionQueryParameters** | [**CognitiveAgentSkillDtoCollectionQueryParameters**](CognitiveAgentSkillDtoCollectionQueryParameters.md) |  | 

### Return type

[**CognitiveAgentSkillDtoListEnvelope**](CognitiveAgentSkillDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCognitiveAgentSkillsCountAsync

> Int32Envelope GetCognitiveAgentSkillsCountAsync(ctx, agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentSkillDtoCollectionQueryParameters(cognitiveAgentSkillDtoCollectionQueryParameters).Execute()

Get skill assignment count for a cognitive agent



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
	cognitiveAgentSkillDtoCollectionQueryParameters := *openapiclient.NewCognitiveAgentSkillDtoCollectionQueryParameters() // CognitiveAgentSkillDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentSkillsAPI.GetCognitiveAgentSkillsCountAsync(context.Background(), agentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentSkillDtoCollectionQueryParameters(cognitiveAgentSkillDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentSkillsAPI.GetCognitiveAgentSkillsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentSkillsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentSkillsAPI.GetCognitiveAgentSkillsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentSkillsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentSkillDtoCollectionQueryParameters** | [**CognitiveAgentSkillDtoCollectionQueryParameters**](CognitiveAgentSkillDtoCollectionQueryParameters.md) |  | 

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


## UpdateCognitiveAgentSkillAsync

> UpdateCognitiveAgentSkillAsync(ctx, agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentSkillUpdateDto(cognitiveAgentSkillUpdateDto).Execute()

Update a cognitive agent skill assignment



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
	cognitiveAgentSkillUpdateDto := *openapiclient.NewCognitiveAgentSkillUpdateDto() // CognitiveAgentSkillUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CognitiveAgentSkillsAPI.UpdateCognitiveAgentSkillAsync(context.Background(), agentId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentSkillUpdateDto(cognitiveAgentSkillUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentSkillsAPI.UpdateCognitiveAgentSkillAsync``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateCognitiveAgentSkillAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentSkillUpdateDto** | [**CognitiveAgentSkillUpdateDto**](CognitiveAgentSkillUpdateDto.md) |  | 

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

