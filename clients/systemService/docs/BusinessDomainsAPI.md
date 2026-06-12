# \BusinessDomainsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemBusinessDomain**](BusinessDomainsAPI.md#DeleteSystemBusinessDomain) | **Delete** /api/v2/SystemService/BusinessDomains/{businessDomainId} | Delete a business domain
[**GetSystemBusinessDomainById**](BusinessDomainsAPI.md#GetSystemBusinessDomainById) | **Get** /api/v2/SystemService/BusinessDomains/{businessDomainId} | Retrieve a business domain by its ID
[**GetSystemBusinessDomains**](BusinessDomainsAPI.md#GetSystemBusinessDomains) | **Get** /api/v2/SystemService/BusinessDomains | Retrieve all business domains in the system
[**GetSystemBusinessDomainsCount**](BusinessDomainsAPI.md#GetSystemBusinessDomainsCount) | **Get** /api/v2/SystemService/BusinessDomains/Count | Get the count of all business domains in the system
[**VerifySystemBusinessDomain**](BusinessDomainsAPI.md#VerifySystemBusinessDomain) | **Post** /api/v2/SystemService/BusinessDomains/{businessDomainId}/Verify | Verify a business domain



## DeleteSystemBusinessDomain

> EmptyEnvelope DeleteSystemBusinessDomain(ctx, businessDomainId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

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
	businessDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessDomainsAPI.DeleteSystemBusinessDomain(context.Background(), businessDomainId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.DeleteSystemBusinessDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSystemBusinessDomain`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.DeleteSystemBusinessDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemBusinessDomainRequest struct via the builder pattern


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


## GetSystemBusinessDomainById

> BusinessDomainDtoEnvelope GetSystemBusinessDomainById(ctx, businessDomainId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a business domain by its ID



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
	businessDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessDomainsAPI.GetSystemBusinessDomainById(context.Background(), businessDomainId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.GetSystemBusinessDomainById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemBusinessDomainById`: BusinessDomainDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.GetSystemBusinessDomainById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemBusinessDomainByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetSystemBusinessDomains

> BusinessDomainDtoListEnvelope GetSystemBusinessDomains(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve all business domains in the system



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
	resp, r, err := apiClient.BusinessDomainsAPI.GetSystemBusinessDomains(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.GetSystemBusinessDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemBusinessDomains`: BusinessDomainDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.GetSystemBusinessDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemBusinessDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## GetSystemBusinessDomainsCount

> Int32Envelope GetSystemBusinessDomainsCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of all business domains in the system



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
	resp, r, err := apiClient.BusinessDomainsAPI.GetSystemBusinessDomainsCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.GetSystemBusinessDomainsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemBusinessDomainsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.GetSystemBusinessDomainsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemBusinessDomainsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## VerifySystemBusinessDomain

> EmptyEnvelope VerifySystemBusinessDomain(ctx, businessDomainId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

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
	businessDomainId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BusinessDomainsAPI.VerifySystemBusinessDomain(context.Background(), businessDomainId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BusinessDomainsAPI.VerifySystemBusinessDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifySystemBusinessDomain`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BusinessDomainsAPI.VerifySystemBusinessDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessDomainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifySystemBusinessDomainRequest struct via the builder pattern


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

