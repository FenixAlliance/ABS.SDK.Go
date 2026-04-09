# \BudgetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBudgetAccountEntryAsync**](BudgetsAPI.md#CreateBudgetAccountEntryAsync) | **Post** /api/v2/AccountingService/Budgets/{budgetId}/AccountEntries | Creates a budget account entry
[**CreateBudgetAsync**](BudgetsAPI.md#CreateBudgetAsync) | **Post** /api/v2/AccountingService/Budgets | Creates a budget
[**DeleteBudgetAccountEntryAsync**](BudgetsAPI.md#DeleteBudgetAccountEntryAsync) | **Delete** /api/v2/AccountingService/Budgets/{budgetId}/AccountEntries/{entryId} | Deletes a budget account entry
[**DeleteBudgetAsync**](BudgetsAPI.md#DeleteBudgetAsync) | **Delete** /api/v2/AccountingService/Budgets/{budgetId} | Deletes a budget
[**GetBudgetAccountEntriesCollectionAsync**](BudgetsAPI.md#GetBudgetAccountEntriesCollectionAsync) | **Get** /api/v2/AccountingService/Budgets/{budgetId}/AccountEntries | Gets all budget account entries
[**GetBudgetAccountEntryAsync**](BudgetsAPI.md#GetBudgetAccountEntryAsync) | **Get** /api/v2/AccountingService/Budgets/{budgetId}/AccountEntries/{entryId} | Gets a budget account entry by id
[**GetBudgetDetailsAsync**](BudgetsAPI.md#GetBudgetDetailsAsync) | **Get** /api/v2/AccountingService/Budgets/{budgetId} | Gets a budget by id
[**GetBudgetsAsync**](BudgetsAPI.md#GetBudgetsAsync) | **Get** /api/v2/AccountingService/Budgets | Gets all budgets
[**UpdateBudgetAccountEntryAsync**](BudgetsAPI.md#UpdateBudgetAccountEntryAsync) | **Put** /api/v2/AccountingService/Budgets/{budgetId}/AccountEntries/{entryId} | Updates a budget account entry
[**UpdateBudgetAsync**](BudgetsAPI.md#UpdateBudgetAsync) | **Put** /api/v2/AccountingService/Budgets/{budgetId} | Updates a budget



## CreateBudgetAccountEntryAsync

> EmptyEnvelope CreateBudgetAccountEntryAsync(ctx, budgetId).TenantId(tenantId).BudgetAccountEntryCreateDto(budgetAccountEntryCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Creates a budget account entry



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
	budgetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	budgetAccountEntryCreateDto := *openapiclient.NewBudgetAccountEntryCreateDto("Description_example", "CurrencyId_example") // BudgetAccountEntryCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.CreateBudgetAccountEntryAsync(context.Background(), budgetId).TenantId(tenantId).BudgetAccountEntryCreateDto(budgetAccountEntryCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.CreateBudgetAccountEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBudgetAccountEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.CreateBudgetAccountEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBudgetAccountEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **budgetAccountEntryCreateDto** | [**BudgetAccountEntryCreateDto**](BudgetAccountEntryCreateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## CreateBudgetAsync

> EmptyEnvelope CreateBudgetAsync(ctx).TenantId(tenantId).BudgetCreateDto(budgetCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Creates a budget



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
	budgetCreateDto := *openapiclient.NewBudgetCreateDto() // BudgetCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.CreateBudgetAsync(context.Background()).TenantId(tenantId).BudgetCreateDto(budgetCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.CreateBudgetAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBudgetAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.CreateBudgetAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBudgetAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **budgetCreateDto** | [**BudgetCreateDto**](BudgetCreateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## DeleteBudgetAccountEntryAsync

> EmptyEnvelope DeleteBudgetAccountEntryAsync(ctx, budgetId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a budget account entry



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
	budgetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.DeleteBudgetAccountEntryAsync(context.Background(), budgetId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.DeleteBudgetAccountEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBudgetAccountEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.DeleteBudgetAccountEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBudgetAccountEntryAsyncRequest struct via the builder pattern


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


## DeleteBudgetAsync

> EmptyEnvelope DeleteBudgetAsync(ctx, budgetId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a budget



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
	budgetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.DeleteBudgetAsync(context.Background(), budgetId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.DeleteBudgetAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBudgetAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.DeleteBudgetAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBudgetAsyncRequest struct via the builder pattern


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


## GetBudgetAccountEntriesCollectionAsync

> BudgetAccountEntryDtoIReadOnlyListEnvelope GetBudgetAccountEntriesCollectionAsync(ctx, budgetId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets all budget account entries



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
	budgetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.GetBudgetAccountEntriesCollectionAsync(context.Background(), budgetId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.GetBudgetAccountEntriesCollectionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetAccountEntriesCollectionAsync`: BudgetAccountEntryDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.GetBudgetAccountEntriesCollectionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetAccountEntriesCollectionAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BudgetAccountEntryDtoIReadOnlyListEnvelope**](BudgetAccountEntryDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetAccountEntryAsync

> BudgetAccountEntryDtoEnvelope GetBudgetAccountEntryAsync(ctx, budgetId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets a budget account entry by id



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
	budgetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.GetBudgetAccountEntryAsync(context.Background(), budgetId, entryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.GetBudgetAccountEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetAccountEntryAsync`: BudgetAccountEntryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.GetBudgetAccountEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetAccountEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BudgetAccountEntryDtoEnvelope**](BudgetAccountEntryDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetDetailsAsync

> BudgetDtoEnvelope GetBudgetDetailsAsync(ctx, budgetId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets a budget by id



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
	budgetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.GetBudgetDetailsAsync(context.Background(), budgetId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.GetBudgetDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetDetailsAsync`: BudgetDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.GetBudgetDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BudgetDtoEnvelope**](BudgetDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetsAsync

> BudgetDtoIReadOnlyListEnvelope GetBudgetsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets all budgets



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
	resp, r, err := apiClient.BudgetsAPI.GetBudgetsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.GetBudgetsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBudgetsAsync`: BudgetDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.GetBudgetsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BudgetDtoIReadOnlyListEnvelope**](BudgetDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBudgetAccountEntryAsync

> EmptyEnvelope UpdateBudgetAccountEntryAsync(ctx, budgetId, entryId).TenantId(tenantId).BudgetAccountEntryUpdateDto(budgetAccountEntryUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Updates a budget account entry



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
	budgetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	entryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	budgetAccountEntryUpdateDto := *openapiclient.NewBudgetAccountEntryUpdateDto() // BudgetAccountEntryUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.UpdateBudgetAccountEntryAsync(context.Background(), budgetId, entryId).TenantId(tenantId).BudgetAccountEntryUpdateDto(budgetAccountEntryUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UpdateBudgetAccountEntryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBudgetAccountEntryAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UpdateBudgetAccountEntryAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 
**entryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBudgetAccountEntryAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **budgetAccountEntryUpdateDto** | [**BudgetAccountEntryUpdateDto**](BudgetAccountEntryUpdateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## UpdateBudgetAsync

> EmptyEnvelope UpdateBudgetAsync(ctx, budgetId).TenantId(tenantId).BudgetUpdateDto(budgetUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Updates a budget



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
	budgetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	budgetUpdateDto := *openapiclient.NewBudgetUpdateDto() // BudgetUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BudgetsAPI.UpdateBudgetAsync(context.Background(), budgetId).TenantId(tenantId).BudgetUpdateDto(budgetUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BudgetsAPI.UpdateBudgetAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBudgetAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BudgetsAPI.UpdateBudgetAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**budgetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBudgetAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **budgetUpdateDto** | [**BudgetUpdateDto**](BudgetUpdateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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

