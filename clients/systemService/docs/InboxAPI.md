# \InboxAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelInboxMessageRetry**](InboxAPI.md#CancelInboxMessageRetry) | **Post** /api/v2/SystemService/Inbox/Messages/{id}/CancelRetry | Cancel a scheduled inbox retry
[**DeadLetterInboxMessage**](InboxAPI.md#DeadLetterInboxMessage) | **Post** /api/v2/SystemService/Inbox/Messages/{id}/DeadLetter | Manually dead-letter an inbox message
[**ExpediteInboxMessage**](InboxAPI.md#ExpediteInboxMessage) | **Post** /api/v2/SystemService/Inbox/Messages/{id}/Expedite | Expedite a retry-scheduled inbox message
[**GetDuplicateInboxMessages**](InboxAPI.md#GetDuplicateInboxMessages) | **Get** /api/v2/SystemService/Inbox/Duplicates | List duplicate-bearing inbox messages
[**GetDuplicateInboxMessagesCount**](InboxAPI.md#GetDuplicateInboxMessagesCount) | **Get** /api/v2/SystemService/Inbox/Duplicates/Count | Count duplicate-bearing inbox messages
[**GetInboxCorrelationChain**](InboxAPI.md#GetInboxCorrelationChain) | **Get** /api/v2/SystemService/Inbox/Correlations/{correlationId} | Get an inbox correlation chain
[**GetInboxHealth**](InboxAPI.md#GetInboxHealth) | **Get** /api/v2/SystemService/Inbox/Health | Get durable-inbox processor health
[**GetInboxMessage**](InboxAPI.md#GetInboxMessage) | **Get** /api/v2/SystemService/Inbox/Messages/{id} | Get one inbox message
[**GetInboxMessages**](InboxAPI.md#GetInboxMessages) | **Get** /api/v2/SystemService/Inbox/Messages | List inbox messages
[**GetInboxMessagesCount**](InboxAPI.md#GetInboxMessagesCount) | **Get** /api/v2/SystemService/Inbox/Messages/Count | Count inbox messages
[**QuarantineInboxMessage**](InboxAPI.md#QuarantineInboxMessage) | **Post** /api/v2/SystemService/Inbox/Messages/{id}/Quarantine | Manually quarantine an inbox message
[**ReleaseInboxMessageLease**](InboxAPI.md#ReleaseInboxMessageLease) | **Post** /api/v2/SystemService/Inbox/Messages/{id}/ReleaseLease | Release a stuck inbox lease
[**ReplayInboxMessage**](InboxAPI.md#ReplayInboxMessage) | **Post** /api/v2/SystemService/Inbox/Messages/{id}/Replay | Replay a terminal inbox message as a new generation



## CancelInboxMessageRetry

> EmptyEnvelope CancelInboxMessageRetry(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxAdminReasonDto(inboxAdminReasonDto).Execute()

Cancel a scheduled inbox retry



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	inboxAdminReasonDto := *openapiclient.NewInboxAdminReasonDto("Reason_example") // InboxAdminReasonDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.CancelInboxMessageRetry(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxAdminReasonDto(inboxAdminReasonDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.CancelInboxMessageRetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelInboxMessageRetry`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.CancelInboxMessageRetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelInboxMessageRetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inboxAdminReasonDto** | [**InboxAdminReasonDto**](InboxAdminReasonDto.md) |  | 

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


## DeadLetterInboxMessage

> EmptyEnvelope DeadLetterInboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxAdminReasonDto(inboxAdminReasonDto).Execute()

Manually dead-letter an inbox message



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	inboxAdminReasonDto := *openapiclient.NewInboxAdminReasonDto("Reason_example") // InboxAdminReasonDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.DeadLetterInboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxAdminReasonDto(inboxAdminReasonDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.DeadLetterInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeadLetterInboxMessage`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.DeadLetterInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeadLetterInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inboxAdminReasonDto** | [**InboxAdminReasonDto**](InboxAdminReasonDto.md) |  | 

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


## ExpediteInboxMessage

> EmptyEnvelope ExpediteInboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Expedite a retry-scheduled inbox message



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.ExpediteInboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.ExpediteInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpediteInboxMessage`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.ExpediteInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpediteInboxMessageRequest struct via the builder pattern


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


## GetDuplicateInboxMessages

> InboxMessageDtoIReadOnlyListEnvelope GetDuplicateInboxMessages(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxMessageDtoCollectionQueryParameters(inboxMessageDtoCollectionQueryParameters).Execute()

List duplicate-bearing inbox messages



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
	inboxMessageDtoCollectionQueryParameters := *openapiclient.NewInboxMessageDtoCollectionQueryParameters() // InboxMessageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.GetDuplicateInboxMessages(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxMessageDtoCollectionQueryParameters(inboxMessageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.GetDuplicateInboxMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDuplicateInboxMessages`: InboxMessageDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.GetDuplicateInboxMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDuplicateInboxMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inboxMessageDtoCollectionQueryParameters** | [**InboxMessageDtoCollectionQueryParameters**](InboxMessageDtoCollectionQueryParameters.md) |  | 

### Return type

[**InboxMessageDtoIReadOnlyListEnvelope**](InboxMessageDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDuplicateInboxMessagesCount

> Int32Envelope GetDuplicateInboxMessagesCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxMessageDtoCollectionQueryParameters(inboxMessageDtoCollectionQueryParameters).Execute()

Count duplicate-bearing inbox messages



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
	inboxMessageDtoCollectionQueryParameters := *openapiclient.NewInboxMessageDtoCollectionQueryParameters() // InboxMessageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.GetDuplicateInboxMessagesCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxMessageDtoCollectionQueryParameters(inboxMessageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.GetDuplicateInboxMessagesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDuplicateInboxMessagesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.GetDuplicateInboxMessagesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDuplicateInboxMessagesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inboxMessageDtoCollectionQueryParameters** | [**InboxMessageDtoCollectionQueryParameters**](InboxMessageDtoCollectionQueryParameters.md) |  | 

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


## GetInboxCorrelationChain

> InboxMessageDtoIReadOnlyListEnvelope GetInboxCorrelationChain(ctx, correlationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get an inbox correlation chain



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
	correlationId := "correlationId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.GetInboxCorrelationChain(context.Background(), correlationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.GetInboxCorrelationChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxCorrelationChain`: InboxMessageDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.GetInboxCorrelationChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**correlationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxCorrelationChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InboxMessageDtoIReadOnlyListEnvelope**](InboxMessageDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxHealth

> InboxHealthDtoEnvelope GetInboxHealth(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get durable-inbox processor health



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.GetInboxHealth(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.GetInboxHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxHealth`: InboxHealthDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.GetInboxHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InboxHealthDtoEnvelope**](InboxHealthDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxMessage

> InboxMessageDtoEnvelope GetInboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get one inbox message



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.GetInboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.GetInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxMessage`: InboxMessageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.GetInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InboxMessageDtoEnvelope**](InboxMessageDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxMessages

> InboxMessageDtoIReadOnlyListEnvelope GetInboxMessages(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxMessageDtoCollectionQueryParameters(inboxMessageDtoCollectionQueryParameters).Execute()

List inbox messages



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
	inboxMessageDtoCollectionQueryParameters := *openapiclient.NewInboxMessageDtoCollectionQueryParameters() // InboxMessageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.GetInboxMessages(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxMessageDtoCollectionQueryParameters(inboxMessageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.GetInboxMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxMessages`: InboxMessageDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.GetInboxMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inboxMessageDtoCollectionQueryParameters** | [**InboxMessageDtoCollectionQueryParameters**](InboxMessageDtoCollectionQueryParameters.md) |  | 

### Return type

[**InboxMessageDtoIReadOnlyListEnvelope**](InboxMessageDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboxMessagesCount

> Int32Envelope GetInboxMessagesCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxMessageDtoCollectionQueryParameters(inboxMessageDtoCollectionQueryParameters).Execute()

Count inbox messages



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
	inboxMessageDtoCollectionQueryParameters := *openapiclient.NewInboxMessageDtoCollectionQueryParameters() // InboxMessageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.GetInboxMessagesCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxMessageDtoCollectionQueryParameters(inboxMessageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.GetInboxMessagesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboxMessagesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.GetInboxMessagesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxMessagesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inboxMessageDtoCollectionQueryParameters** | [**InboxMessageDtoCollectionQueryParameters**](InboxMessageDtoCollectionQueryParameters.md) |  | 

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


## QuarantineInboxMessage

> EmptyEnvelope QuarantineInboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxAdminReasonDto(inboxAdminReasonDto).Execute()

Manually quarantine an inbox message



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	inboxAdminReasonDto := *openapiclient.NewInboxAdminReasonDto("Reason_example") // InboxAdminReasonDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.QuarantineInboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxAdminReasonDto(inboxAdminReasonDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.QuarantineInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuarantineInboxMessage`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.QuarantineInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuarantineInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inboxAdminReasonDto** | [**InboxAdminReasonDto**](InboxAdminReasonDto.md) |  | 

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


## ReleaseInboxMessageLease

> EmptyEnvelope ReleaseInboxMessageLease(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Release a stuck inbox lease



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.ReleaseInboxMessageLease(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.ReleaseInboxMessageLease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseInboxMessageLease`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.ReleaseInboxMessageLease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseInboxMessageLeaseRequest struct via the builder pattern


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


## ReplayInboxMessage

> InboxReplayResultDtoEnvelope ReplayInboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxAdminReasonDto(inboxAdminReasonDto).Execute()

Replay a terminal inbox message as a new generation



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	inboxAdminReasonDto := *openapiclient.NewInboxAdminReasonDto("Reason_example") // InboxAdminReasonDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboxAPI.ReplayInboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).InboxAdminReasonDto(inboxAdminReasonDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboxAPI.ReplayInboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayInboxMessage`: InboxReplayResultDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `InboxAPI.ReplayInboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayInboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **inboxAdminReasonDto** | [**InboxAdminReasonDto**](InboxAdminReasonDto.md) |  | 

### Return type

[**InboxReplayResultDtoEnvelope**](InboxReplayResultDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

