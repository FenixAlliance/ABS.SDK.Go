# \CognitiveAgentConversationAttachmentsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadCognitiveAgentConversationAttachmentAsync**](CognitiveAgentConversationAttachmentsAPI.md#UploadCognitiveAgentConversationAttachmentAsync) | **Post** /api/v2/IntelligenceService/CognitiveAgents/{agentId}/Conversations/{conversationId}/Attachments | Upload an attachment to a cognitive agent conversation



## UploadCognitiveAgentConversationAttachmentAsync

> ConversationAttachmentUploadResultDtoEnvelope UploadCognitiveAgentConversationAttachmentAsync(ctx, agentId, conversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()

Upload an attachment to a cognitive agent conversation



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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CognitiveAgentConversationAttachmentsAPI.UploadCognitiveAgentConversationAttachmentAsync(context.Background(), agentId, conversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CognitiveAgentConversationAttachmentsAPI.UploadCognitiveAgentConversationAttachmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCognitiveAgentConversationAttachmentAsync`: ConversationAttachmentUploadResultDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CognitiveAgentConversationAttachmentsAPI.UploadCognitiveAgentConversationAttachmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agentId** | **string** |  | 
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCognitiveAgentConversationAttachmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

[**ConversationAttachmentUploadResultDtoEnvelope**](ConversationAttachmentUploadResultDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

