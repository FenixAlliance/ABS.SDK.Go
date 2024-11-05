# \TenantsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](TenantsAPI.md#CreateTenant) | **Post** /api/v2/SystemService/Tenants | Create a new tenant.
[**DeleteTenant**](TenantsAPI.md#DeleteTenant) | **Delete** /api/v2/SystemService/Tenants/{tenantId} | Delete a specific tenant by ID.
[**GetAllExtendedTenants**](TenantsAPI.md#GetAllExtendedTenants) | **Get** /api/v2/SystemService/Tenants/Extended | Get all extended tenants available on this suite server instance.
[**GetAllTenants**](TenantsAPI.md#GetAllTenants) | **Get** /api/v2/SystemService/Tenants | Get all tenants available on this suite server instance.
[**GetExtendedTenantsCount**](TenantsAPI.md#GetExtendedTenantsCount) | **Get** /api/v2/SystemService/Tenants/Extended/Count | Get the total count of extended tenants available on this suite server instance.
[**GetTenant**](TenantsAPI.md#GetTenant) | **Get** /api/v2/SystemService/Tenants/{tenantId} | Get a specific tenant by ID.
[**GetTenantsCount**](TenantsAPI.md#GetTenantsCount) | **Get** /api/v2/SystemService/Tenants/Count | Get the total count of tenants available on this suite server instance.
[**UpdateTenant**](TenantsAPI.md#UpdateTenant) | **Put** /api/v2/SystemService/Tenants/{tenantId} | Update a specific tenant by ID.



## CreateTenant

> TenantDtoEnvelope CreateTenant(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantCreateDto(tenantCreateDto).Execute()

Create a new tenant.



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
	tenantCreateDto := *openapiclient.NewTenantCreateDto() // TenantCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.CreateTenant(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantCreateDto(tenantCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.CreateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenant`: TenantDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantCreateDto** | [**TenantCreateDto**](TenantCreateDto.md) |  | 

### Return type

[**TenantDtoEnvelope**](TenantDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenant

> EmptyEnvelope DeleteTenant(ctx, tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a specific tenant by ID.



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
	resp, r, err := apiClient.TenantsAPI.DeleteTenant(context.Background(), tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.DeleteTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenant`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.DeleteTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllExtendedTenants

> ExtendedTenantDtoListEnvelope GetAllExtendedTenants(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all extended tenants available on this suite server instance.



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
	resp, r, err := apiClient.TenantsAPI.GetAllExtendedTenants(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetAllExtendedTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllExtendedTenants`: ExtendedTenantDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetAllExtendedTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllExtendedTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExtendedTenantDtoListEnvelope**](ExtendedTenantDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTenants

> TenantDtoListEnvelope GetAllTenants(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all tenants available on this suite server instance.



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
	resp, r, err := apiClient.TenantsAPI.GetAllTenants(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetAllTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTenants`: TenantDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetAllTenants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTenantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantDtoListEnvelope**](TenantDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtendedTenantsCount

> Int32Envelope GetExtendedTenantsCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the total count of extended tenants available on this suite server instance.



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
	resp, r, err := apiClient.TenantsAPI.GetExtendedTenantsCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetExtendedTenantsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedTenantsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetExtendedTenantsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedTenantsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenant

> TenantDtoEnvelope GetTenant(ctx, tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get a specific tenant by ID.



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
	resp, r, err := apiClient.TenantsAPI.GetTenant(context.Background(), tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenant`: TenantDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantDtoEnvelope**](TenantDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantsCount

> Int32Envelope GetTenantsCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the total count of tenants available on this suite server instance.



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
	resp, r, err := apiClient.TenantsAPI.GetTenantsCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenantsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenantsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenant

> TenantDtoEnvelope UpdateTenant(ctx, tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantUpdateDto(tenantUpdateDto).Execute()

Update a specific tenant by ID.



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
	tenantUpdateDto := *openapiclient.NewTenantUpdateDto() // TenantUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.UpdateTenant(context.Background(), tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantUpdateDto(tenantUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.UpdateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenant`: TenantDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantUpdateDto** | [**TenantUpdateDto**](TenantUpdateDto.md) |  | 

### Return type

[**TenantDtoEnvelope**](TenantDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

