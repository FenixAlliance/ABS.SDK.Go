# \SupportTicketsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SupportServiceSupportTicketsCountGet**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsCountGet) | **Get** /api/v2/SupportService/SupportTickets/Count | 
[**ApiV2SupportServiceSupportTicketsGet**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsGet) | **Get** /api/v2/SupportService/SupportTickets | 
[**ApiV2SupportServiceSupportTicketsPost**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsPost) | **Post** /api/v2/SupportService/SupportTickets | 
[**ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGet**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGet) | **Get** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations | 
[**ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPost**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPost) | **Post** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations | 
[**ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDelete**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDelete) | **Delete** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations/{supportTicketConversationId} | 
[**ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGet**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGet) | **Get** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations/{supportTicketConversationId} | 
[**ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGet**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGet) | **Get** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations/{supportTicketConversationId}/Messages | 
[**ApiV2SupportServiceSupportTicketsSupportTicketIdDelete**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsSupportTicketIdDelete) | **Delete** /api/v2/SupportService/SupportTickets/{supportTicketId} | 
[**ApiV2SupportServiceSupportTicketsSupportTicketIdGet**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsSupportTicketIdGet) | **Get** /api/v2/SupportService/SupportTickets/{supportTicketId} | 
[**ApiV2SupportServiceSupportTicketsSupportTicketIdPut**](SupportTicketsAPI.md#ApiV2SupportServiceSupportTicketsSupportTicketIdPut) | **Put** /api/v2/SupportService/SupportTickets/{supportTicketId} | 



## ApiV2SupportServiceSupportTicketsCountGet

> Int32Envelope ApiV2SupportServiceSupportTicketsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsGet

> SupportTicketDtoListEnvelope ApiV2SupportServiceSupportTicketsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsGet`: SupportTicketDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketDtoListEnvelope**](SupportTicketDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsPost

> EmptyEnvelope ApiV2SupportServiceSupportTicketsPost(ctx).SupportTicketCreateDto(supportTicketCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketCreateDto := *openapiclient.NewSupportTicketCreateDto() // SupportTicketCreateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsPost(context.Background()).SupportTicketCreateDto(supportTicketCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportTicketCreateDto** | [**SupportTicketCreateDto**](SupportTicketCreateDto.md) |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGet

> SupportTicketConversationDtoListEnvelope ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGet(ctx, supportTicketId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGet(context.Background(), supportTicketId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGet`: SupportTicketConversationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsSupportTicketIdConversationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketConversationDtoListEnvelope**](SupportTicketConversationDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPost

> EmptyEnvelope ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPost(ctx, supportTicketId).SupportTicketConversationCreateDto(supportTicketConversationCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketConversationCreateDto := *openapiclient.NewSupportTicketConversationCreateDto() // SupportTicketConversationCreateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPost(context.Background(), supportTicketId).SupportTicketConversationCreateDto(supportTicketConversationCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsSupportTicketIdConversationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportTicketConversationCreateDto** | [**SupportTicketConversationCreateDto**](SupportTicketConversationCreateDto.md) |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDelete

> EmptyEnvelope ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDelete(ctx, supportTicketId, supportTicketConversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketConversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDelete(context.Background(), supportTicketId, supportTicketConversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 
**supportTicketConversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGet

> SupportTicketConversationDtoEnvelope ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGet(ctx, supportTicketId, supportTicketConversationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketConversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGet(context.Background(), supportTicketId, supportTicketConversationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGet`: SupportTicketConversationDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 
**supportTicketConversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketConversationDtoEnvelope**](SupportTicketConversationDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGet

> PrivateMessageDtoListEnvelope ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGet(ctx, supportTicketConversationId, supportTicketId).PageNumber(pageNumber).PageSize(pageSize).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketConversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketId := "supportTicketId_example" // string | 
	pageNumber := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGet(context.Background(), supportTicketConversationId, supportTicketId).PageNumber(pageNumber).PageSize(pageSize).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGet`: PrivateMessageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketConversationId** | **string** |  | 
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsSupportTicketIdConversationsSupportTicketConversationIdMessagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageNumber** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PrivateMessageDtoListEnvelope**](PrivateMessageDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsSupportTicketIdDelete

> EmptyEnvelope ApiV2SupportServiceSupportTicketsSupportTicketIdDelete(ctx, supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdDelete(context.Background(), supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsSupportTicketIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsSupportTicketIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsSupportTicketIdGet

> SupportTicketDtoEnvelope ApiV2SupportServiceSupportTicketsSupportTicketIdGet(ctx, supportTicketId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdGet(context.Background(), supportTicketId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsSupportTicketIdGet`: SupportTicketDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsSupportTicketIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketDtoEnvelope**](SupportTicketDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketsSupportTicketIdPut

> EmptyEnvelope ApiV2SupportServiceSupportTicketsSupportTicketIdPut(ctx, supportTicketId).SupportTicketUpdateDto(supportTicketUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketUpdateDto := *openapiclient.NewSupportTicketUpdateDto() // SupportTicketUpdateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdPut(context.Background(), supportTicketId).SupportTicketUpdateDto(supportTicketUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketsSupportTicketIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.ApiV2SupportServiceSupportTicketsSupportTicketIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketsSupportTicketIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportTicketUpdateDto** | [**SupportTicketUpdateDto**](SupportTicketUpdateDto.md) |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

