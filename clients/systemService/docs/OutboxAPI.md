# \OutboxAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOutboxMessage**](OutboxAPI.md#CancelOutboxMessage) | **Post** /api/v2/SystemService/Outbox/Messages/{id}/Cancel | Cancel an outbox message
[**DeadLetterOutboxMessage**](OutboxAPI.md#DeadLetterOutboxMessage) | **Post** /api/v2/SystemService/Outbox/Messages/{id}/DeadLetter | Manually dead-letter an outbox message
[**ExpediteOutboxMessage**](OutboxAPI.md#ExpediteOutboxMessage) | **Post** /api/v2/SystemService/Outbox/Messages/{id}/Expedite | Expedite a failed (retry-eligible) outbox message
[**GetOutboxCorrelationChain**](OutboxAPI.md#GetOutboxCorrelationChain) | **Get** /api/v2/SystemService/Outbox/Correlations/{correlationId} | Get an outbox correlation chain
[**GetOutboxHealth**](OutboxAPI.md#GetOutboxHealth) | **Get** /api/v2/SystemService/Outbox/Health | Get durable-outbox relay health
[**GetOutboxMessage**](OutboxAPI.md#GetOutboxMessage) | **Get** /api/v2/SystemService/Outbox/Messages/{id} | Get one outbox message
[**GetOutboxMessages**](OutboxAPI.md#GetOutboxMessages) | **Get** /api/v2/SystemService/Outbox/Messages | List outbox messages
[**GetOutboxMessagesCount**](OutboxAPI.md#GetOutboxMessagesCount) | **Get** /api/v2/SystemService/Outbox/Messages/Count | Count outbox messages
[**ReleaseOutboxMessageLease**](OutboxAPI.md#ReleaseOutboxMessageLease) | **Post** /api/v2/SystemService/Outbox/Messages/{id}/ReleaseLease | Release a stuck outbox lease
[**ReplayOutboxMessage**](OutboxAPI.md#ReplayOutboxMessage) | **Post** /api/v2/SystemService/Outbox/Messages/{id}/Replay | Replay a dead-lettered or failed outbox message



## CancelOutboxMessage

> EmptyEnvelope CancelOutboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).OutboxAdminReasonDto(outboxAdminReasonDto).Execute()

Cancel an outbox message



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
	outboxAdminReasonDto := *openapiclient.NewOutboxAdminReasonDto("Reason_example") // OutboxAdminReasonDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.CancelOutboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).OutboxAdminReasonDto(outboxAdminReasonDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.CancelOutboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOutboxMessage`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.CancelOutboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOutboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **outboxAdminReasonDto** | [**OutboxAdminReasonDto**](OutboxAdminReasonDto.md) |  | 

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


## DeadLetterOutboxMessage

> EmptyEnvelope DeadLetterOutboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).OutboxAdminReasonDto(outboxAdminReasonDto).Execute()

Manually dead-letter an outbox message



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
	outboxAdminReasonDto := *openapiclient.NewOutboxAdminReasonDto("Reason_example") // OutboxAdminReasonDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.DeadLetterOutboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).OutboxAdminReasonDto(outboxAdminReasonDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.DeadLetterOutboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeadLetterOutboxMessage`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.DeadLetterOutboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeadLetterOutboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **outboxAdminReasonDto** | [**OutboxAdminReasonDto**](OutboxAdminReasonDto.md) |  | 

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


## ExpediteOutboxMessage

> EmptyEnvelope ExpediteOutboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Expedite a failed (retry-eligible) outbox message



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
	resp, r, err := apiClient.OutboxAPI.ExpediteOutboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.ExpediteOutboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpediteOutboxMessage`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.ExpediteOutboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpediteOutboxMessageRequest struct via the builder pattern


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


## GetOutboxCorrelationChain

> OutboxMessageDtoIReadOnlyListEnvelope GetOutboxCorrelationChain(ctx, correlationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get an outbox correlation chain



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
	resp, r, err := apiClient.OutboxAPI.GetOutboxCorrelationChain(context.Background(), correlationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.GetOutboxCorrelationChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboxCorrelationChain`: OutboxMessageDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.GetOutboxCorrelationChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**correlationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboxCorrelationChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**OutboxMessageDtoIReadOnlyListEnvelope**](OutboxMessageDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutboxHealth

> OutboxHealthDtoEnvelope GetOutboxHealth(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get durable-outbox relay health



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
	resp, r, err := apiClient.OutboxAPI.GetOutboxHealth(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.GetOutboxHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboxHealth`: OutboxHealthDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.GetOutboxHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboxHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**OutboxHealthDtoEnvelope**](OutboxHealthDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutboxMessage

> OutboxMessageDtoEnvelope GetOutboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get one outbox message



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
	resp, r, err := apiClient.OutboxAPI.GetOutboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.GetOutboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboxMessage`: OutboxMessageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.GetOutboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboxMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**OutboxMessageDtoEnvelope**](OutboxMessageDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutboxMessages

> OutboxMessageDtoIReadOnlyListEnvelope GetOutboxMessages(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).OutboxMessageDtoCollectionQueryParameters(outboxMessageDtoCollectionQueryParameters).Execute()

List outbox messages



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
	outboxMessageDtoCollectionQueryParameters := *openapiclient.NewOutboxMessageDtoCollectionQueryParameters() // OutboxMessageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.GetOutboxMessages(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).OutboxMessageDtoCollectionQueryParameters(outboxMessageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.GetOutboxMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboxMessages`: OutboxMessageDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.GetOutboxMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboxMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **outboxMessageDtoCollectionQueryParameters** | [**OutboxMessageDtoCollectionQueryParameters**](OutboxMessageDtoCollectionQueryParameters.md) |  | 

### Return type

[**OutboxMessageDtoIReadOnlyListEnvelope**](OutboxMessageDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutboxMessagesCount

> Int32Envelope GetOutboxMessagesCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).OutboxMessageDtoCollectionQueryParameters(outboxMessageDtoCollectionQueryParameters).Execute()

Count outbox messages



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
	outboxMessageDtoCollectionQueryParameters := *openapiclient.NewOutboxMessageDtoCollectionQueryParameters() // OutboxMessageDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutboxAPI.GetOutboxMessagesCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).OutboxMessageDtoCollectionQueryParameters(outboxMessageDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.GetOutboxMessagesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboxMessagesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.GetOutboxMessagesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboxMessagesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **outboxMessageDtoCollectionQueryParameters** | [**OutboxMessageDtoCollectionQueryParameters**](OutboxMessageDtoCollectionQueryParameters.md) |  | 

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


## ReleaseOutboxMessageLease

> EmptyEnvelope ReleaseOutboxMessageLease(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Release a stuck outbox lease



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
	resp, r, err := apiClient.OutboxAPI.ReleaseOutboxMessageLease(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.ReleaseOutboxMessageLease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseOutboxMessageLease`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.ReleaseOutboxMessageLease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseOutboxMessageLeaseRequest struct via the builder pattern


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


## ReplayOutboxMessage

> EmptyEnvelope ReplayOutboxMessage(ctx, id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Replay a dead-lettered or failed outbox message



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
	resp, r, err := apiClient.OutboxAPI.ReplayOutboxMessage(context.Background(), id).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboxAPI.ReplayOutboxMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayOutboxMessage`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OutboxAPI.ReplayOutboxMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayOutboxMessageRequest struct via the builder pattern


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

