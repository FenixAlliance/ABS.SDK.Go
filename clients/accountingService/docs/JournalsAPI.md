# \JournalsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountJournalsAsync**](JournalsAPI.md#CountJournalsAsync) | **Get** /api/v2/AccountingService/Journals/Count | Count journals
[**CreateJournalAsync**](JournalsAPI.md#CreateJournalAsync) | **Post** /api/v2/AccountingService/Journals | Create journal
[**CreateJournalEntryAsync**](JournalsAPI.md#CreateJournalEntryAsync) | **Post** /api/v2/AccountingService/Journals/{journalId}/Entries | Create journal entry
[**DeleteJournalAsync**](JournalsAPI.md#DeleteJournalAsync) | **Delete** /api/v2/AccountingService/Journals/{journalId} | Delete journal
[**DeleteJournalEntryAsync**](JournalsAPI.md#DeleteJournalEntryAsync) | **Delete** /api/v2/AccountingService/Journals/{journalId}/Entries/{entryId} | Delete journal entry
[**GetJournalDetailsAsync**](JournalsAPI.md#GetJournalDetailsAsync) | **Get** /api/v2/AccountingService/Journals/{journalId} | Get journal by ID
[**GetJournalEntriesAsync**](JournalsAPI.md#GetJournalEntriesAsync) | **Get** /api/v2/AccountingService/Journals/{journalId}/Entries | Get journal entries
[**GetJournalEntriesCountAsync**](JournalsAPI.md#GetJournalEntriesCountAsync) | **Get** /api/v2/AccountingService/Journals/{journalId}/Entries/Count | Count journal entries
[**GetJournalsAsync**](JournalsAPI.md#GetJournalsAsync) | **Get** /api/v2/AccountingService/Journals | Get all journals
[**UpdateJournalAsync**](JournalsAPI.md#UpdateJournalAsync) | **Put** /api/v2/AccountingService/Journals/{journalId} | Update journal
[**UpdateJournalEntryAsync**](JournalsAPI.md#UpdateJournalEntryAsync) | **Put** /api/v2/AccountingService/Journals/{journalId}/Entries/{entryId} | Update journal entry



## CountJournalsAsync

> Int32Envelope CountJournalsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count journals



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
	resp, r, err := apiClient.JournalsAPI.CountJournalsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.CountJournalsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountJournalsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.CountJournalsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountJournalsAsyncRequest struct via the builder pattern


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


## CreateJournalAsync

> EmptyEnvelope CreateJournalAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalCreateDto(journalCreateDto).Execute()

Create journal



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
	journalCreateDto := *openapiclient.NewJournalCreateDto("Name_example") // JournalCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.CreateJournalAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalCreateDto(journalCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.CreateJournalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJournalAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.CreateJournalAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJournalAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **journalCreateDto** | [**JournalCreateDto**](JournalCreateDto.md) |  | 

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


## CreateJournalEntryAsync

> EmptyEnvelope CreateJournalEntryAsync(ctx, journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalEntryCreateDto(journalEntryCreateDto).Execute()

Create journal entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	journalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	journalEntryCreateDto := *openapiclient.NewJournalEntryCreateDto("Description_example", time.Now(), "JournalId_example", "CurrencyId_example", "DebitAccountId_example", "CreditAccountId_example") // JournalEntryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.CreateJournalEntryAsync(context.Background(), journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalEntryCreateDto(journalEntryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.CreateJournalEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJournalEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.CreateJournalEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJournalEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **journalEntryCreateDto** | [**JournalEntryCreateDto**](JournalEntryCreateDto.md) |  | 

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


## DeleteJournalAsync

> EmptyEnvelope DeleteJournalAsync(ctx, journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete journal



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
	journalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.DeleteJournalAsync(context.Background(), journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.DeleteJournalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJournalAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.DeleteJournalAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJournalAsyncRequest struct via the builder pattern


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


## DeleteJournalEntryAsync

> EmptyEnvelope DeleteJournalEntryAsync(ctx, journalId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete journal entry



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
	journalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.DeleteJournalEntryAsync(context.Background(), journalId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.DeleteJournalEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJournalEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.DeleteJournalEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJournalEntryAsyncRequest struct via the builder pattern


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


## GetJournalDetailsAsync

> JournalDtoEnvelope GetJournalDetailsAsync(ctx, journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get journal by ID



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
	journalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.GetJournalDetailsAsync(context.Background(), journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.GetJournalDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournalDetailsAsync`: JournalDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.GetJournalDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JournalDtoEnvelope**](JournalDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournalEntriesAsync

> JournalEntryDtoIReadOnlyListEnvelope GetJournalEntriesAsync(ctx, journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get journal entries



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
	journalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.GetJournalEntriesAsync(context.Background(), journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.GetJournalEntriesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournalEntriesAsync`: JournalEntryDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.GetJournalEntriesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalEntriesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JournalEntryDtoIReadOnlyListEnvelope**](JournalEntryDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournalEntriesCountAsync

> Int32Envelope GetJournalEntriesCountAsync(ctx, journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count journal entries



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
	journalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.GetJournalEntriesCountAsync(context.Background(), journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.GetJournalEntriesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournalEntriesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.GetJournalEntriesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalEntriesCountAsyncRequest struct via the builder pattern


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


## GetJournalsAsync

> JournalDtoIReadOnlyListEnvelope GetJournalsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all journals



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
	resp, r, err := apiClient.JournalsAPI.GetJournalsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.GetJournalsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJournalsAsync`: JournalDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.GetJournalsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJournalsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JournalDtoIReadOnlyListEnvelope**](JournalDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJournalAsync

> EmptyEnvelope UpdateJournalAsync(ctx, journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalUpdateDto(journalUpdateDto).Execute()

Update journal



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
	journalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	journalUpdateDto := *openapiclient.NewJournalUpdateDto() // JournalUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.UpdateJournalAsync(context.Background(), journalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalUpdateDto(journalUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.UpdateJournalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJournalAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.UpdateJournalAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJournalAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **journalUpdateDto** | [**JournalUpdateDto**](JournalUpdateDto.md) |  | 

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


## UpdateJournalEntryAsync

> EmptyEnvelope UpdateJournalEntryAsync(ctx, journalId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalEntryUpdateDto(journalEntryUpdateDto).Execute()

Update journal entry



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	journalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	journalEntryUpdateDto := *openapiclient.NewJournalEntryUpdateDto("Description_example", time.Now(), "JournalId_example", "CurrencyId_example", "DebitAccountId_example", "CreditAccountId_example") // JournalEntryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JournalsAPI.UpdateJournalEntryAsync(context.Background(), journalId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JournalEntryUpdateDto(journalEntryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JournalsAPI.UpdateJournalEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJournalEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JournalsAPI.UpdateJournalEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**journalId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJournalEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **journalEntryUpdateDto** | [**JournalEntryUpdateDto**](JournalEntryUpdateDto.md) |  | 

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

