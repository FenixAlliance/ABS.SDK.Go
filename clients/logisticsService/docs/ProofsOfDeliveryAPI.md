# \ProofsOfDeliveryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProofOfDeliveryLineAsync**](ProofsOfDeliveryAPI.md#AddProofOfDeliveryLineAsync) | **Post** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Lines | Add a line to proof of delivery
[**AttachDeliveryNoteAsync**](ProofsOfDeliveryAPI.md#AttachDeliveryNoteAsync) | **Post** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/DeliveryNotes/{noteId} | Attach a delivery note
[**CreateProofOfDeliveryAsync**](ProofsOfDeliveryAPI.md#CreateProofOfDeliveryAsync) | **Post** /api/v2/LogisticsService/ProofsOfDelivery | Create a proof of delivery
[**DeleteProofOfDeliveryAsync**](ProofsOfDeliveryAPI.md#DeleteProofOfDeliveryAsync) | **Delete** /api/v2/LogisticsService/ProofsOfDelivery/{podId} | Delete a proof of delivery
[**DetachDeliveryNoteAsync**](ProofsOfDeliveryAPI.md#DetachDeliveryNoteAsync) | **Delete** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/DeliveryNotes/{noteId} | Detach a delivery note
[**DisputeProofOfDeliveryAsync**](ProofsOfDeliveryAPI.md#DisputeProofOfDeliveryAsync) | **Post** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Dispute | Dispute a proof of delivery
[**GetProofOfDeliveryByIdAsync**](ProofsOfDeliveryAPI.md#GetProofOfDeliveryByIdAsync) | **Get** /api/v2/LogisticsService/ProofsOfDelivery/{podId} | Get proof of delivery by ID
[**GetProofOfDeliveryDeliveryNotesAsync**](ProofsOfDeliveryAPI.md#GetProofOfDeliveryDeliveryNotesAsync) | **Get** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/DeliveryNotes | Get attached delivery notes
[**GetProofOfDeliveryDeliveryNotesCountAsync**](ProofsOfDeliveryAPI.md#GetProofOfDeliveryDeliveryNotesCountAsync) | **Get** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/DeliveryNotes/Count | Get delivery notes count
[**GetProofOfDeliveryLinesAsync**](ProofsOfDeliveryAPI.md#GetProofOfDeliveryLinesAsync) | **Get** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Lines | Get proof of delivery lines
[**GetProofOfDeliveryLinesCountAsync**](ProofsOfDeliveryAPI.md#GetProofOfDeliveryLinesCountAsync) | **Get** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Lines/Count | Get proof of delivery lines count
[**GetProofsOfDeliveryAsync**](ProofsOfDeliveryAPI.md#GetProofsOfDeliveryAsync) | **Get** /api/v2/LogisticsService/ProofsOfDelivery | Get all proofs of delivery
[**GetProofsOfDeliveryCountAsync**](ProofsOfDeliveryAPI.md#GetProofsOfDeliveryCountAsync) | **Get** /api/v2/LogisticsService/ProofsOfDelivery/Count | Get proofs of delivery count
[**PatchProofOfDeliveryAsync**](ProofsOfDeliveryAPI.md#PatchProofOfDeliveryAsync) | **Patch** /api/v2/LogisticsService/ProofsOfDelivery/{podId} | Patch a proof of delivery
[**PatchProofOfDeliveryLineAsync**](ProofsOfDeliveryAPI.md#PatchProofOfDeliveryLineAsync) | **Patch** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Lines/{lineId} | Patch a proof of delivery line
[**RejectProofOfDeliveryAsync**](ProofsOfDeliveryAPI.md#RejectProofOfDeliveryAsync) | **Post** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Reject | Reject a proof of delivery
[**RemoveProofOfDeliveryLineAsync**](ProofsOfDeliveryAPI.md#RemoveProofOfDeliveryLineAsync) | **Delete** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Lines/{lineId} | Remove a proof of delivery line
[**SignProofOfDeliveryAsync**](ProofsOfDeliveryAPI.md#SignProofOfDeliveryAsync) | **Post** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Sign | Sign a proof of delivery
[**UpdateProofOfDeliveryAsync**](ProofsOfDeliveryAPI.md#UpdateProofOfDeliveryAsync) | **Put** /api/v2/LogisticsService/ProofsOfDelivery/{podId} | Update a proof of delivery
[**UpdateProofOfDeliveryLineAsync**](ProofsOfDeliveryAPI.md#UpdateProofOfDeliveryLineAsync) | **Put** /api/v2/LogisticsService/ProofsOfDelivery/{podId}/Lines/{lineId} | Update a proof of delivery line



## AddProofOfDeliveryLineAsync

> EmptyEnvelope AddProofOfDeliveryLineAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProofOfDeliveryLineCreateDto(proofOfDeliveryLineCreateDto).Execute()

Add a line to proof of delivery



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	proofOfDeliveryLineCreateDto := *openapiclient.NewProofOfDeliveryLineCreateDto() // ProofOfDeliveryLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.AddProofOfDeliveryLineAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProofOfDeliveryLineCreateDto(proofOfDeliveryLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.AddProofOfDeliveryLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProofOfDeliveryLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.AddProofOfDeliveryLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProofOfDeliveryLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **proofOfDeliveryLineCreateDto** | [**ProofOfDeliveryLineCreateDto**](ProofOfDeliveryLineCreateDto.md) |  | 

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


## AttachDeliveryNoteAsync

> EmptyEnvelope AttachDeliveryNoteAsync(ctx, podId, noteId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Attach a delivery note



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	noteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.AttachDeliveryNoteAsync(context.Background(), podId, noteId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.AttachDeliveryNoteAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachDeliveryNoteAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.AttachDeliveryNoteAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachDeliveryNoteAsyncRequest struct via the builder pattern


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


## CreateProofOfDeliveryAsync

> EmptyEnvelope CreateProofOfDeliveryAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProofOfDeliveryCreateDto(proofOfDeliveryCreateDto).Execute()

Create a proof of delivery



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
	proofOfDeliveryCreateDto := *openapiclient.NewProofOfDeliveryCreateDto() // ProofOfDeliveryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.CreateProofOfDeliveryAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProofOfDeliveryCreateDto(proofOfDeliveryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.CreateProofOfDeliveryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProofOfDeliveryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.CreateProofOfDeliveryAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProofOfDeliveryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **proofOfDeliveryCreateDto** | [**ProofOfDeliveryCreateDto**](ProofOfDeliveryCreateDto.md) |  | 

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


## DeleteProofOfDeliveryAsync

> EmptyEnvelope DeleteProofOfDeliveryAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a proof of delivery



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.DeleteProofOfDeliveryAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.DeleteProofOfDeliveryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProofOfDeliveryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.DeleteProofOfDeliveryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProofOfDeliveryAsyncRequest struct via the builder pattern


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


## DetachDeliveryNoteAsync

> EmptyEnvelope DetachDeliveryNoteAsync(ctx, podId, noteId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Detach a delivery note



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	noteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.DetachDeliveryNoteAsync(context.Background(), podId, noteId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.DetachDeliveryNoteAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachDeliveryNoteAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.DetachDeliveryNoteAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 
**noteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachDeliveryNoteAsyncRequest struct via the builder pattern


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


## DisputeProofOfDeliveryAsync

> EmptyEnvelope DisputeProofOfDeliveryAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DisputeProofOfDeliveryRequest(disputeProofOfDeliveryRequest).Execute()

Dispute a proof of delivery



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	disputeProofOfDeliveryRequest := *openapiclient.NewDisputeProofOfDeliveryRequest() // DisputeProofOfDeliveryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.DisputeProofOfDeliveryAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DisputeProofOfDeliveryRequest(disputeProofOfDeliveryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.DisputeProofOfDeliveryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeProofOfDeliveryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.DisputeProofOfDeliveryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisputeProofOfDeliveryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **disputeProofOfDeliveryRequest** | [**DisputeProofOfDeliveryRequest**](DisputeProofOfDeliveryRequest.md) |  | 

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


## GetProofOfDeliveryByIdAsync

> ProofOfDeliveryDtoEnvelope GetProofOfDeliveryByIdAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get proof of delivery by ID



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.GetProofOfDeliveryByIdAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.GetProofOfDeliveryByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProofOfDeliveryByIdAsync`: ProofOfDeliveryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.GetProofOfDeliveryByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProofOfDeliveryByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProofOfDeliveryDtoEnvelope**](ProofOfDeliveryDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProofOfDeliveryDeliveryNotesAsync

> DeliveryNoteDtoListEnvelope GetProofOfDeliveryDeliveryNotesAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get attached delivery notes



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.GetProofOfDeliveryDeliveryNotesAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.GetProofOfDeliveryDeliveryNotesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProofOfDeliveryDeliveryNotesAsync`: DeliveryNoteDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.GetProofOfDeliveryDeliveryNotesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProofOfDeliveryDeliveryNotesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DeliveryNoteDtoListEnvelope**](DeliveryNoteDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProofOfDeliveryDeliveryNotesCountAsync

> Int32Envelope GetProofOfDeliveryDeliveryNotesCountAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get delivery notes count



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.GetProofOfDeliveryDeliveryNotesCountAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.GetProofOfDeliveryDeliveryNotesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProofOfDeliveryDeliveryNotesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.GetProofOfDeliveryDeliveryNotesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProofOfDeliveryDeliveryNotesCountAsyncRequest struct via the builder pattern


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


## GetProofOfDeliveryLinesAsync

> ProofOfDeliveryLineDtoListEnvelope GetProofOfDeliveryLinesAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get proof of delivery lines



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.GetProofOfDeliveryLinesAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.GetProofOfDeliveryLinesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProofOfDeliveryLinesAsync`: ProofOfDeliveryLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.GetProofOfDeliveryLinesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProofOfDeliveryLinesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProofOfDeliveryLineDtoListEnvelope**](ProofOfDeliveryLineDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProofOfDeliveryLinesCountAsync

> Int32Envelope GetProofOfDeliveryLinesCountAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get proof of delivery lines count



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.GetProofOfDeliveryLinesCountAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.GetProofOfDeliveryLinesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProofOfDeliveryLinesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.GetProofOfDeliveryLinesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProofOfDeliveryLinesCountAsyncRequest struct via the builder pattern


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


## GetProofsOfDeliveryAsync

> ProofOfDeliveryDtoListEnvelope GetProofsOfDeliveryAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all proofs of delivery



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
	resp, r, err := apiClient.ProofsOfDeliveryAPI.GetProofsOfDeliveryAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.GetProofsOfDeliveryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProofsOfDeliveryAsync`: ProofOfDeliveryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.GetProofsOfDeliveryAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProofsOfDeliveryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ProofOfDeliveryDtoListEnvelope**](ProofOfDeliveryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProofsOfDeliveryCountAsync

> Int32Envelope GetProofsOfDeliveryCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get proofs of delivery count



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
	resp, r, err := apiClient.ProofsOfDeliveryAPI.GetProofsOfDeliveryCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.GetProofsOfDeliveryCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProofsOfDeliveryCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.GetProofsOfDeliveryCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProofsOfDeliveryCountAsyncRequest struct via the builder pattern


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


## PatchProofOfDeliveryAsync

> EmptyEnvelope PatchProofOfDeliveryAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a proof of delivery



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.PatchProofOfDeliveryAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.PatchProofOfDeliveryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProofOfDeliveryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.PatchProofOfDeliveryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProofOfDeliveryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## PatchProofOfDeliveryLineAsync

> EmptyEnvelope PatchProofOfDeliveryLineAsync(ctx, podId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a proof of delivery line



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.PatchProofOfDeliveryLineAsync(context.Background(), podId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.PatchProofOfDeliveryLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProofOfDeliveryLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.PatchProofOfDeliveryLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProofOfDeliveryLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## RejectProofOfDeliveryAsync

> EmptyEnvelope RejectProofOfDeliveryAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RejectProofOfDeliveryRequest(rejectProofOfDeliveryRequest).Execute()

Reject a proof of delivery



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	rejectProofOfDeliveryRequest := *openapiclient.NewRejectProofOfDeliveryRequest() // RejectProofOfDeliveryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.RejectProofOfDeliveryAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).RejectProofOfDeliveryRequest(rejectProofOfDeliveryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.RejectProofOfDeliveryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectProofOfDeliveryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.RejectProofOfDeliveryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectProofOfDeliveryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **rejectProofOfDeliveryRequest** | [**RejectProofOfDeliveryRequest**](RejectProofOfDeliveryRequest.md) |  | 

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


## RemoveProofOfDeliveryLineAsync

> EmptyEnvelope RemoveProofOfDeliveryLineAsync(ctx, podId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove a proof of delivery line



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.RemoveProofOfDeliveryLineAsync(context.Background(), podId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.RemoveProofOfDeliveryLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveProofOfDeliveryLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.RemoveProofOfDeliveryLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProofOfDeliveryLineAsyncRequest struct via the builder pattern


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


## SignProofOfDeliveryAsync

> EmptyEnvelope SignProofOfDeliveryAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SignProofOfDeliveryRequest(signProofOfDeliveryRequest).Execute()

Sign a proof of delivery



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	signProofOfDeliveryRequest := *openapiclient.NewSignProofOfDeliveryRequest() // SignProofOfDeliveryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.SignProofOfDeliveryAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SignProofOfDeliveryRequest(signProofOfDeliveryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.SignProofOfDeliveryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignProofOfDeliveryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.SignProofOfDeliveryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignProofOfDeliveryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **signProofOfDeliveryRequest** | [**SignProofOfDeliveryRequest**](SignProofOfDeliveryRequest.md) |  | 

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


## UpdateProofOfDeliveryAsync

> EmptyEnvelope UpdateProofOfDeliveryAsync(ctx, podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProofOfDeliveryUpdateDto(proofOfDeliveryUpdateDto).Execute()

Update a proof of delivery



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	proofOfDeliveryUpdateDto := *openapiclient.NewProofOfDeliveryUpdateDto() // ProofOfDeliveryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.UpdateProofOfDeliveryAsync(context.Background(), podId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProofOfDeliveryUpdateDto(proofOfDeliveryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.UpdateProofOfDeliveryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProofOfDeliveryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.UpdateProofOfDeliveryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProofOfDeliveryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **proofOfDeliveryUpdateDto** | [**ProofOfDeliveryUpdateDto**](ProofOfDeliveryUpdateDto.md) |  | 

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


## UpdateProofOfDeliveryLineAsync

> EmptyEnvelope UpdateProofOfDeliveryLineAsync(ctx, podId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProofOfDeliveryLineUpdateDto(proofOfDeliveryLineUpdateDto).Execute()

Update a proof of delivery line



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
	podId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	lineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	proofOfDeliveryLineUpdateDto := *openapiclient.NewProofOfDeliveryLineUpdateDto() // ProofOfDeliveryLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProofsOfDeliveryAPI.UpdateProofOfDeliveryLineAsync(context.Background(), podId, lineId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ProofOfDeliveryLineUpdateDto(proofOfDeliveryLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProofsOfDeliveryAPI.UpdateProofOfDeliveryLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProofOfDeliveryLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ProofsOfDeliveryAPI.UpdateProofOfDeliveryLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**podId** | **string** |  | 
**lineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProofOfDeliveryLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **proofOfDeliveryLineUpdateDto** | [**ProofOfDeliveryLineUpdateDto**](ProofOfDeliveryLineUpdateDto.md) |  | 

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

