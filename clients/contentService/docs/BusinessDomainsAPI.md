# \BusinessDomainsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBusinessDomainAsync**](BusinessDomainsAPI.md#CreateBusinessDomainAsync) | **Post** /api/v2/ContentService/BusinessDomains | Register a business domain
[**DeleteBusinessDomainAsync**](BusinessDomainsAPI.md#DeleteBusinessDomainAsync) | **Delete** /api/v2/ContentService/BusinessDomains/{businessDomainId} | Delete a business domain
[**GetBusinessDomainByIdAsync**](BusinessDomainsAPI.md#GetBusinessDomainByIdAsync) | **Get** /api/v2/ContentService/BusinessDomains/{businessDomainId} | Get business domain by ID
[**GetBusinessDomainsAsync**](BusinessDomainsAPI.md#GetBusinessDomainsAsync) | **Get** /api/v2/ContentService/BusinessDomains | Get business domains
[**GetBusinessDomainsCountAsync**](BusinessDomainsAPI.md#GetBusinessDomainsCountAsync) | **Get** /api/v2/ContentService/BusinessDomains/Count | Get business domains count
[**UpdateBusinessDomainAsync**](BusinessDomainsAPI.md#UpdateBusinessDomainAsync) | **Put** /api/v2/ContentService/BusinessDomains/{businessDomainId} | Update a business domain
[**VerifyBusinessDomainAsync**](BusinessDomainsAPI.md#VerifyBusinessDomainAsync) | **Post** /api/v2/ContentService/BusinessDomains/{businessDomainId}/Verify | Verify a business domain



## CreateBusinessDomainAsync

> EmptyEnvelope CreateBusinessDomainAsync(ctx).TenantId(tenantId).BusinessDomainCreateDto(businessDomainCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Register a business domain



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
	businessDomainCreateDto := *openapiclient.NewBusinessDomainCreateDto("Domain_example") // BusinessDomainCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessDomainsAPI.CreateBusinessDomainAsync(context.Background()).TenantId(tenantId).BusinessDomainCreateDto(businessDomainCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.CreateBusinessDomainAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBusinessDomainAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.CreateBusinessDomainAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBusinessDomainAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **businessDomainCreateDto** | [**BusinessDomainCreateDto**](BusinessDomainCreateDto.md) |  | 
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


## DeleteBusinessDomainAsync

> EmptyEnvelope DeleteBusinessDomainAsync(ctx, businessDomainId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a business domain



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
	businessDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessDomainsAPI.DeleteBusinessDomainAsync(context.Background(), businessDomainId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.DeleteBusinessDomainAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBusinessDomainAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.DeleteBusinessDomainAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessDomainAsyncRequest struct via the builder pattern


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


## GetBusinessDomainByIdAsync

> BusinessDomainDtoEnvelope GetBusinessDomainByIdAsync(ctx, businessDomainId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get business domain by ID



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
	businessDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessDomainsAPI.GetBusinessDomainByIdAsync(context.Background(), businessDomainId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.GetBusinessDomainByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBusinessDomainByIdAsync`: BusinessDomainDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.GetBusinessDomainByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessDomainByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BusinessDomainDtoEnvelope**](BusinessDomainDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessDomainsAsync

> BusinessDomainDtoListEnvelope GetBusinessDomainsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get business domains



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
	resp, r, err := apiClient.BusinessDomainsAPI.GetBusinessDomainsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.GetBusinessDomainsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBusinessDomainsAsync`: BusinessDomainDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.GetBusinessDomainsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessDomainsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BusinessDomainDtoListEnvelope**](BusinessDomainDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBusinessDomainsCountAsync

> Int32Envelope GetBusinessDomainsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get business domains count



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
	resp, r, err := apiClient.BusinessDomainsAPI.GetBusinessDomainsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.GetBusinessDomainsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBusinessDomainsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.GetBusinessDomainsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessDomainsCountAsyncRequest struct via the builder pattern


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


## UpdateBusinessDomainAsync

> EmptyEnvelope UpdateBusinessDomainAsync(ctx, businessDomainId).TenantId(tenantId).BusinessDomainUpdateDto(businessDomainUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update a business domain



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
	businessDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	businessDomainUpdateDto := *openapiclient.NewBusinessDomainUpdateDto() // BusinessDomainUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessDomainsAPI.UpdateBusinessDomainAsync(context.Background(), businessDomainId).TenantId(tenantId).BusinessDomainUpdateDto(businessDomainUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.UpdateBusinessDomainAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBusinessDomainAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.UpdateBusinessDomainAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessDomainAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **businessDomainUpdateDto** | [**BusinessDomainUpdateDto**](BusinessDomainUpdateDto.md) |  | 
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


## VerifyBusinessDomainAsync

> EmptyEnvelope VerifyBusinessDomainAsync(ctx, businessDomainId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Verify a business domain



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
	businessDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessDomainsAPI.VerifyBusinessDomainAsync(context.Background(), businessDomainId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.VerifyBusinessDomainAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyBusinessDomainAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.VerifyBusinessDomainAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyBusinessDomainAsyncRequest struct via the builder pattern


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

