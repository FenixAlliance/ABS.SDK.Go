# \SupportTicketsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportTicketAsync**](SupportTicketsAPI.md#CreateSupportTicketAsync) | **Post** /api/v2/SupportService/SupportTickets | Create a new support ticket
[**DeleteSupportTicketAsync**](SupportTicketsAPI.md#DeleteSupportTicketAsync) | **Delete** /api/v2/SupportService/SupportTickets/{supportTicketId} | Delete a support ticket
[**DeleteSupportTicketConversationAsync**](SupportTicketsAPI.md#DeleteSupportTicketConversationAsync) | **Delete** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations/{supportTicketConversationId} | Delete a conversation from a support ticket
[**GetSupportTicketAsync**](SupportTicketsAPI.md#GetSupportTicketAsync) | **Get** /api/v2/SupportService/SupportTickets/{supportTicketId} | Retrieve a support ticket by ID
[**GetSupportTicketConversationAsync**](SupportTicketsAPI.md#GetSupportTicketConversationAsync) | **Get** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations/{supportTicketConversationId} | Retrieve a specific conversation for a support ticket
[**GetSupportTicketConversationMessagesAsync**](SupportTicketsAPI.md#GetSupportTicketConversationMessagesAsync) | **Get** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations/{supportTicketConversationId}/Messages | Retrieve messages for a support ticket conversation
[**GetSupportTicketConversationsAsync**](SupportTicketsAPI.md#GetSupportTicketConversationsAsync) | **Get** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations | Retrieve conversations for a support ticket
[**GetSupportTicketsAsync**](SupportTicketsAPI.md#GetSupportTicketsAsync) | **Get** /api/v2/SupportService/SupportTickets | Retrieve a list of support tickets
[**GetSupportTicketsCountAsync**](SupportTicketsAPI.md#GetSupportTicketsCountAsync) | **Get** /api/v2/SupportService/SupportTickets/Count | Get the count of support tickets
[**RelateSupportTicketToConversationAsync**](SupportTicketsAPI.md#RelateSupportTicketToConversationAsync) | **Post** /api/v2/SupportService/SupportTickets/{supportTicketId}/Conversations | Create a conversation for a support ticket
[**UpdateSupportTicketAsync**](SupportTicketsAPI.md#UpdateSupportTicketAsync) | **Put** /api/v2/SupportService/SupportTickets/{supportTicketId} | Update a support ticket



## CreateSupportTicketAsync

> EmptyEnvelope CreateSupportTicketAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketCreateDto(supportTicketCreateDto).Execute()

Create a new support ticket



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
	supportTicketCreateDto := *openapiclient.NewSupportTicketCreateDto() // SupportTicketCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.CreateSupportTicketAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketCreateDto(supportTicketCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.CreateSupportTicketAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSupportTicketAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.CreateSupportTicketAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportTicketAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportTicketCreateDto** | [**SupportTicketCreateDto**](SupportTicketCreateDto.md) |  | 

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


## DeleteSupportTicketAsync

> EmptyEnvelope DeleteSupportTicketAsync(ctx, supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a support ticket



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.DeleteSupportTicketAsync(context.Background(), supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.DeleteSupportTicketAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSupportTicketAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.DeleteSupportTicketAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupportTicketAsyncRequest struct via the builder pattern


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


## DeleteSupportTicketConversationAsync

> EmptyEnvelope DeleteSupportTicketConversationAsync(ctx, supportTicketId, supportTicketConversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a conversation from a support ticket



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketConversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.DeleteSupportTicketConversationAsync(context.Background(), supportTicketId, supportTicketConversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.DeleteSupportTicketConversationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSupportTicketConversationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.DeleteSupportTicketConversationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 
**supportTicketConversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSupportTicketConversationAsyncRequest struct via the builder pattern


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


## GetSupportTicketAsync

> SupportTicketDtoEnvelope GetSupportTicketAsync(ctx, supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a support ticket by ID



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.GetSupportTicketAsync(context.Background(), supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.GetSupportTicketAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketAsync`: SupportTicketDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.GetSupportTicketAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketDtoEnvelope**](SupportTicketDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportTicketConversationAsync

> SupportTicketConversationDtoEnvelope GetSupportTicketConversationAsync(ctx, supportTicketId, supportTicketConversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a specific conversation for a support ticket



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketConversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.GetSupportTicketConversationAsync(context.Background(), supportTicketId, supportTicketConversationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.GetSupportTicketConversationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketConversationAsync`: SupportTicketConversationDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.GetSupportTicketConversationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 
**supportTicketConversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketConversationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketConversationDtoEnvelope**](SupportTicketConversationDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportTicketConversationMessagesAsync

> PrivateMessageDtoListEnvelope GetSupportTicketConversationMessagesAsync(ctx, supportTicketId, supportTicketConversationId).TenantId(tenantId).PageNumber(pageNumber).PageSize(pageSize).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve messages for a support ticket conversation



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketConversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	pageNumber := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.GetSupportTicketConversationMessagesAsync(context.Background(), supportTicketId, supportTicketConversationId).TenantId(tenantId).PageNumber(pageNumber).PageSize(pageSize).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.GetSupportTicketConversationMessagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketConversationMessagesAsync`: PrivateMessageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.GetSupportTicketConversationMessagesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 
**supportTicketConversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketConversationMessagesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **pageNumber** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PrivateMessageDtoListEnvelope**](PrivateMessageDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportTicketConversationsAsync

> SupportTicketConversationDtoListEnvelope GetSupportTicketConversationsAsync(ctx, supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve conversations for a support ticket



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.GetSupportTicketConversationsAsync(context.Background(), supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.GetSupportTicketConversationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketConversationsAsync`: SupportTicketConversationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.GetSupportTicketConversationsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketConversationsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketConversationDtoListEnvelope**](SupportTicketConversationDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportTicketsAsync

> SupportTicketDtoListEnvelope GetSupportTicketsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of support tickets



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.GetSupportTicketsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.GetSupportTicketsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketsAsync`: SupportTicketDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.GetSupportTicketsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketDtoListEnvelope**](SupportTicketDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportTicketsCountAsync

> Int32Envelope GetSupportTicketsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of support tickets



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.GetSupportTicketsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.GetSupportTicketsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSupportTicketsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.GetSupportTicketsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportTicketsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelateSupportTicketToConversationAsync

> EmptyEnvelope RelateSupportTicketToConversationAsync(ctx, supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketConversationCreateDto(supportTicketConversationCreateDto).Execute()

Create a conversation for a support ticket



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	supportTicketConversationCreateDto := *openapiclient.NewSupportTicketConversationCreateDto() // SupportTicketConversationCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.RelateSupportTicketToConversationAsync(context.Background(), supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketConversationCreateDto(supportTicketConversationCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.RelateSupportTicketToConversationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelateSupportTicketToConversationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.RelateSupportTicketToConversationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelateSupportTicketToConversationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportTicketConversationCreateDto** | [**SupportTicketConversationCreateDto**](SupportTicketConversationCreateDto.md) |  | 

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


## UpdateSupportTicketAsync

> EmptyEnvelope UpdateSupportTicketAsync(ctx, supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketUpdateDto(supportTicketUpdateDto).Execute()

Update a support ticket



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
	supportTicketId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	supportTicketUpdateDto := *openapiclient.NewSupportTicketUpdateDto() // SupportTicketUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketsAPI.UpdateSupportTicketAsync(context.Background(), supportTicketId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SupportTicketUpdateDto(supportTicketUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketsAPI.UpdateSupportTicketAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSupportTicketAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketsAPI.UpdateSupportTicketAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSupportTicketAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **supportTicketUpdateDto** | [**SupportTicketUpdateDto**](SupportTicketUpdateDto.md) |  | 

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

