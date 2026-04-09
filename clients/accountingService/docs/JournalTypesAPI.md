# \JournalTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJournalTypeAsync**](JournalTypesAPI.md#CreateJournalTypeAsync) | **Post** /api/v2/AccountingService/JournalTypes | Creates a new journal type
[**DeleteJournalTypeAsync**](JournalTypesAPI.md#DeleteJournalTypeAsync) | **Delete** /api/v2/AccountingService/JournalTypes/{journalTypeId} | Deletes a journal type
[**GetJournalTypeDetailsAsync**](JournalTypesAPI.md#GetJournalTypeDetailsAsync) | **Get** /api/v2/AccountingService/JournalTypes/{journalTypeId} | Retrieves a journal type by ID
[**GetJournalTypesAsync**](JournalTypesAPI.md#GetJournalTypesAsync) | **Get** /api/v2/AccountingService/JournalTypes | Retrieves all journal types
[**GetJournalTypesCountAsync**](JournalTypesAPI.md#GetJournalTypesCountAsync) | **Get** /api/v2/AccountingService/JournalTypes/Count | Counts journal types
[**UpdateJournalTypeAsync**](JournalTypesAPI.md#UpdateJournalTypeAsync) | **Put** /api/v2/AccountingService/JournalTypes/{journalTypeId} | Updates an existing journal type



## CreateJournalTypeAsync

> EmptyEnvelope CreateJournalTypeAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalTypeCreateDto(journalTypeCreateDto).Execute()

Creates a new journal type



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
	journalTypeCreateDto := *openapiclient.NewJournalTypeCreateDto() // JournalTypeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalTypesAPI.CreateJournalTypeAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalTypeCreateDto(journalTypeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalTypesAPI.CreateJournalTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJournalTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalTypesAPI.CreateJournalTypeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJournalTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **journalTypeCreateDto** | [**JournalTypeCreateDto**](JournalTypeCreateDto.md) |  | 

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


## DeleteJournalTypeAsync

> EmptyEnvelope DeleteJournalTypeAsync(ctx, journalTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a journal type



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
	journalTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalTypesAPI.DeleteJournalTypeAsync(context.Background(), journalTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalTypesAPI.DeleteJournalTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJournalTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalTypesAPI.DeleteJournalTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJournalTypeAsyncRequest struct via the builder pattern


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


## GetJournalTypeDetailsAsync

> JournalTypeDtoEnvelope GetJournalTypeDetailsAsync(ctx, journalTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieves a journal type by ID



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
	journalTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalTypesAPI.GetJournalTypeDetailsAsync(context.Background(), journalTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalTypesAPI.GetJournalTypeDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournalTypeDetailsAsync`: JournalTypeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalTypesAPI.GetJournalTypeDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalTypeDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JournalTypeDtoEnvelope**](JournalTypeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournalTypesAsync

> JournalTypeDtoIReadOnlyListEnvelope GetJournalTypesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieves all journal types



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
	resp, r, err := apiClient.JournalTypesAPI.GetJournalTypesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalTypesAPI.GetJournalTypesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournalTypesAsync`: JournalTypeDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalTypesAPI.GetJournalTypesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalTypesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JournalTypeDtoIReadOnlyListEnvelope**](JournalTypeDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournalTypesCountAsync

> Int32Envelope GetJournalTypesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Counts journal types



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
	resp, r, err := apiClient.JournalTypesAPI.GetJournalTypesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalTypesAPI.GetJournalTypesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournalTypesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `JournalTypesAPI.GetJournalTypesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalTypesCountAsyncRequest struct via the builder pattern


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


## UpdateJournalTypeAsync

> EmptyEnvelope UpdateJournalTypeAsync(ctx, journalTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalTypeUpdateDto(journalTypeUpdateDto).Execute()

Updates an existing journal type



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
	journalTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	journalTypeUpdateDto := *openapiclient.NewJournalTypeUpdateDto() // JournalTypeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalTypesAPI.UpdateJournalTypeAsync(context.Background(), journalTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalTypeUpdateDto(journalTypeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalTypesAPI.UpdateJournalTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJournalTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalTypesAPI.UpdateJournalTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJournalTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **journalTypeUpdateDto** | [**JournalTypeUpdateDto**](JournalTypeUpdateDto.md) |  | 

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

