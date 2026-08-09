# \CognitiveAgentMessagesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCognitiveAgentMessageByIdAsync**](CognitiveAgentMessagesAPI.md#GetCognitiveAgentMessageByIdAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Conversations/{conversationId}/Messages/{id} | Get a cognitive agent conversation message by ID
[**GetCognitiveAgentMessagesAsync**](CognitiveAgentMessagesAPI.md#GetCognitiveAgentMessagesAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Conversations/{conversationId}/Messages | Get all messages for a cognitive agent conversation
[**GetCognitiveAgentMessagesCountAsync**](CognitiveAgentMessagesAPI.md#GetCognitiveAgentMessagesCountAsync) | **Get** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Conversations/{conversationId}/Messages/Count | Get message count for a cognitive agent conversation



## GetCognitiveAgentMessageByIdAsync

> CognitiveAgentMessageDtoEnvelope GetCognitiveAgentMessageByIdAsync(ctx, agentId, conversationId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get a cognitive agent conversation message by ID



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentMessagesAPI.GetCognitiveAgentMessageByIdAsync(context.Background(), agentId, conversationId, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentMessagesAPI.GetCognitiveAgentMessageByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentMessageByIdAsync`: CognitiveAgentMessageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentMessagesAPI.GetCognitiveAgentMessageByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**conversationId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentMessageByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CognitiveAgentMessageDtoEnvelope**](CognitiveAgentMessageDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCognitiveAgentMessagesAsync

> CognitiveAgentMessageDtoListEnvelope GetCognitiveAgentMessagesAsync(ctx, agentId, conversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentMessageDtoCollectionQueryParameters(cognitiveAgentMessageDtoCollectionQueryParameters).Execute()

Get all messages for a cognitive agent conversation



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	cognitiveAgentMessageDtoCollectionQueryParameters := *openapiclient.NewCognitiveAgentMessageDtoCollectionQueryParameters() // CognitiveAgentMessageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentMessagesAPI.GetCognitiveAgentMessagesAsync(context.Background(), agentId, conversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentMessageDtoCollectionQueryParameters(cognitiveAgentMessageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentMessagesAPI.GetCognitiveAgentMessagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentMessagesAsync`: CognitiveAgentMessageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentMessagesAPI.GetCognitiveAgentMessagesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentMessagesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentMessageDtoCollectionQueryParameters** | [**CognitiveAgentMessageDtoCollectionQueryParameters**](CognitiveAgentMessageDtoCollectionQueryParameters.md) |  | 

### Return type

[**CognitiveAgentMessageDtoListEnvelope**](CognitiveAgentMessageDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCognitiveAgentMessagesCountAsync

> Int32Envelope GetCognitiveAgentMessagesCountAsync(ctx, agentId, conversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentMessageDtoCollectionQueryParameters(cognitiveAgentMessageDtoCollectionQueryParameters).Execute()

Get message count for a cognitive agent conversation



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	cognitiveAgentMessageDtoCollectionQueryParameters := *openapiclient.NewCognitiveAgentMessageDtoCollectionQueryParameters() // CognitiveAgentMessageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentMessagesAPI.GetCognitiveAgentMessagesCountAsync(context.Background(), agentId, conversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CognitiveAgentMessageDtoCollectionQueryParameters(cognitiveAgentMessageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentMessagesAPI.GetCognitiveAgentMessagesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCognitiveAgentMessagesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentMessagesAPI.GetCognitiveAgentMessagesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCognitiveAgentMessagesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **cognitiveAgentMessageDtoCollectionQueryParameters** | [**CognitiveAgentMessageDtoCollectionQueryParameters**](CognitiveAgentMessageDtoCollectionQueryParameters.md) |  | 

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

