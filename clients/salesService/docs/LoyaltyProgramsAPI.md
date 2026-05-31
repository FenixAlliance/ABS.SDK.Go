# \LoyaltyProgramsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountLoyaltyProgramsAsync**](LoyaltyProgramsAPI.md#CountLoyaltyProgramsAsync) | **Get** /api/v2/SalesService/LoyaltyPrograms/Count | Get loyalty programs count
[**CreateLoyaltyProgramAsync**](LoyaltyProgramsAPI.md#CreateLoyaltyProgramAsync) | **Post** /api/v2/SalesService/LoyaltyPrograms | Create a loyalty program
[**DeleteLoyaltyProgramAsync**](LoyaltyProgramsAPI.md#DeleteLoyaltyProgramAsync) | **Delete** /api/v2/SalesService/LoyaltyPrograms/{loyaltyProgramId} | Delete a loyalty program
[**GetLoyaltyProgramAsync**](LoyaltyProgramsAPI.md#GetLoyaltyProgramAsync) | **Get** /api/v2/SalesService/LoyaltyPrograms/{loyaltyProgramId} | Get loyalty program by ID
[**GetLoyaltyProgramsAsync**](LoyaltyProgramsAPI.md#GetLoyaltyProgramsAsync) | **Get** /api/v2/SalesService/LoyaltyPrograms | Get loyalty programs
[**UpdateLoyaltyProgramAsync**](LoyaltyProgramsAPI.md#UpdateLoyaltyProgramAsync) | **Put** /api/v2/SalesService/LoyaltyPrograms/{loyaltyProgramId} | Update a loyalty program



## CountLoyaltyProgramsAsync

> Int32Envelope CountLoyaltyProgramsAsync(ctx).TenantId(tenantId).Execute()

Get loyalty programs count



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoyaltyProgramsAPI.CountLoyaltyProgramsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoyaltyProgramsAPI.CountLoyaltyProgramsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountLoyaltyProgramsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `LoyaltyProgramsAPI.CountLoyaltyProgramsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountLoyaltyProgramsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## CreateLoyaltyProgramAsync

> EmptyEnvelope CreateLoyaltyProgramAsync(ctx).TenantId(tenantId).LoyaltyProgramCreateDto(loyaltyProgramCreateDto).Execute()

Create a loyalty program



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
	loyaltyProgramCreateDto := *openapiclient.NewLoyaltyProgramCreateDto("Title_example") // LoyaltyProgramCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoyaltyProgramsAPI.CreateLoyaltyProgramAsync(context.Background()).TenantId(tenantId).LoyaltyProgramCreateDto(loyaltyProgramCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoyaltyProgramsAPI.CreateLoyaltyProgramAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoyaltyProgramAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LoyaltyProgramsAPI.CreateLoyaltyProgramAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoyaltyProgramAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **loyaltyProgramCreateDto** | [**LoyaltyProgramCreateDto**](LoyaltyProgramCreateDto.md) |  | 

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


## DeleteLoyaltyProgramAsync

> EmptyEnvelope DeleteLoyaltyProgramAsync(ctx, loyaltyProgramId).TenantId(tenantId).Execute()

Delete a loyalty program



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
	loyaltyProgramId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoyaltyProgramsAPI.DeleteLoyaltyProgramAsync(context.Background(), loyaltyProgramId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoyaltyProgramsAPI.DeleteLoyaltyProgramAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLoyaltyProgramAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LoyaltyProgramsAPI.DeleteLoyaltyProgramAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoyaltyProgramAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## GetLoyaltyProgramAsync

> LoyaltyProgramDtoEnvelope GetLoyaltyProgramAsync(ctx, loyaltyProgramId).TenantId(tenantId).Execute()

Get loyalty program by ID



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
	loyaltyProgramId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoyaltyProgramsAPI.GetLoyaltyProgramAsync(context.Background(), loyaltyProgramId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoyaltyProgramsAPI.GetLoyaltyProgramAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyProgramAsync`: LoyaltyProgramDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LoyaltyProgramsAPI.GetLoyaltyProgramAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**LoyaltyProgramDtoEnvelope**](LoyaltyProgramDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramsAsync

> LoyaltyProgramDtoListEnvelope GetLoyaltyProgramsAsync(ctx).TenantId(tenantId).Execute()

Get loyalty programs



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoyaltyProgramsAPI.GetLoyaltyProgramsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoyaltyProgramsAPI.GetLoyaltyProgramsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyProgramsAsync`: LoyaltyProgramDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LoyaltyProgramsAPI.GetLoyaltyProgramsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**LoyaltyProgramDtoListEnvelope**](LoyaltyProgramDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoyaltyProgramAsync

> EmptyEnvelope UpdateLoyaltyProgramAsync(ctx, loyaltyProgramId).TenantId(tenantId).LoyaltyProgramUpdateDto(loyaltyProgramUpdateDto).Execute()

Update a loyalty program



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
	loyaltyProgramId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	loyaltyProgramUpdateDto := *openapiclient.NewLoyaltyProgramUpdateDto() // LoyaltyProgramUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoyaltyProgramsAPI.UpdateLoyaltyProgramAsync(context.Background(), loyaltyProgramId).TenantId(tenantId).LoyaltyProgramUpdateDto(loyaltyProgramUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoyaltyProgramsAPI.UpdateLoyaltyProgramAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoyaltyProgramAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LoyaltyProgramsAPI.UpdateLoyaltyProgramAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoyaltyProgramAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **loyaltyProgramUpdateDto** | [**LoyaltyProgramUpdateDto**](LoyaltyProgramUpdateDto.md) |  | 

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

