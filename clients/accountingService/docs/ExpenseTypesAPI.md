# \ExpenseTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExpenseType**](ExpenseTypesAPI.md#CreateExpenseType) | **Post** /api/v2/AccountingService/ExpenseTypes | Create an expense type
[**DeleteExpenseType**](ExpenseTypesAPI.md#DeleteExpenseType) | **Delete** /api/v2/AccountingService/ExpenseTypes/{expenseTypeId} | Delete an expense type
[**GetExpenseType**](ExpenseTypesAPI.md#GetExpenseType) | **Get** /api/v2/AccountingService/ExpenseTypes/{expenseTypeId} | Get an expense type by id
[**GetExpenseTypes**](ExpenseTypesAPI.md#GetExpenseTypes) | **Get** /api/v2/AccountingService/ExpenseTypes | Get all expense types for a tenant
[**GetExpenseTypesCount**](ExpenseTypesAPI.md#GetExpenseTypesCount) | **Get** /api/v2/AccountingService/ExpenseTypes/Count | Get the count of expense types for a tenant
[**PatchExpenseType**](ExpenseTypesAPI.md#PatchExpenseType) | **Patch** /api/v2/AccountingService/ExpenseTypes/{expenseTypeId} | Patch an expense type
[**UpdateExpenseType**](ExpenseTypesAPI.md#UpdateExpenseType) | **Put** /api/v2/AccountingService/ExpenseTypes/{expenseTypeId} | Update an expense type



## CreateExpenseType

> EmptyEnvelope CreateExpenseType(ctx).TenantId(tenantId).ExpenseTypeCreateDto(expenseTypeCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create an expense type



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
	expenseTypeCreateDto := *openapiclient.NewExpenseTypeCreateDto() // ExpenseTypeCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExpenseTypesAPI.CreateExpenseType(context.Background()).TenantId(tenantId).ExpenseTypeCreateDto(expenseTypeCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExpenseTypesAPI.CreateExpenseType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateExpenseType`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExpenseTypesAPI.CreateExpenseType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExpenseTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **expenseTypeCreateDto** | [**ExpenseTypeCreateDto**](ExpenseTypeCreateDto.md) |  | 
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


## DeleteExpenseType

> EmptyEnvelope DeleteExpenseType(ctx, expenseTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an expense type



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
	expenseTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExpenseTypesAPI.DeleteExpenseType(context.Background(), expenseTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExpenseTypesAPI.DeleteExpenseType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteExpenseType`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExpenseTypesAPI.DeleteExpenseType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExpenseTypeRequest struct via the builder pattern


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


## GetExpenseType

> ExpenseTypeDtoEnvelope GetExpenseType(ctx, expenseTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get an expense type by id



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
	expenseTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExpenseTypesAPI.GetExpenseType(context.Background(), expenseTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExpenseTypesAPI.GetExpenseType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenseType`: ExpenseTypeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExpenseTypesAPI.GetExpenseType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExpenseTypeDtoEnvelope**](ExpenseTypeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseTypes

> ExpenseTypeDtoListEnvelope GetExpenseTypes(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all expense types for a tenant



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
	resp, r, err := apiClient.ExpenseTypesAPI.GetExpenseTypes(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExpenseTypesAPI.GetExpenseTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenseTypes`: ExpenseTypeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExpenseTypesAPI.GetExpenseTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExpenseTypeDtoListEnvelope**](ExpenseTypeDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseTypesCount

> Int32Envelope GetExpenseTypesCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of expense types for a tenant



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
	resp, r, err := apiClient.ExpenseTypesAPI.GetExpenseTypesCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExpenseTypesAPI.GetExpenseTypesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpenseTypesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ExpenseTypesAPI.GetExpenseTypesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseTypesCountRequest struct via the builder pattern


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


## PatchExpenseType

> EmptyEnvelope PatchExpenseType(ctx, expenseTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch an expense type



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
	expenseTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExpenseTypesAPI.PatchExpenseType(context.Background(), expenseTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExpenseTypesAPI.PatchExpenseType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExpenseType`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExpenseTypesAPI.PatchExpenseType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExpenseTypeRequest struct via the builder pattern


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


## UpdateExpenseType

> EmptyEnvelope UpdateExpenseType(ctx, expenseTypeId).TenantId(tenantId).ExpenseTypeUpdateDto(expenseTypeUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update an expense type



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
	expenseTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	expenseTypeUpdateDto := *openapiclient.NewExpenseTypeUpdateDto() // ExpenseTypeUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExpenseTypesAPI.UpdateExpenseType(context.Background(), expenseTypeId).TenantId(tenantId).ExpenseTypeUpdateDto(expenseTypeUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExpenseTypesAPI.UpdateExpenseType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateExpenseType`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ExpenseTypesAPI.UpdateExpenseType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expenseTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExpenseTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **expenseTypeUpdateDto** | [**ExpenseTypeUpdateDto**](ExpenseTypeUpdateDto.md) |  | 
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

