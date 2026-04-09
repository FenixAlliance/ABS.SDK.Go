# \TerritoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenantTerritory**](TerritoriesAPI.md#CreateTenantTerritory) | **Post** /api/v2/TenantsService/Territories | Create a new tenant territory
[**DeleteTenantTerritory**](TerritoriesAPI.md#DeleteTenantTerritory) | **Delete** /api/v2/TenantsService/Territories/{tenantTerritoryId} | Delete a tenant territory
[**GetTenantTerritories**](TerritoriesAPI.md#GetTenantTerritories) | **Get** /api/v2/TenantsService/Territories | Retrieve a list of tenant territories
[**GetTenantTerritoriesCount**](TerritoriesAPI.md#GetTenantTerritoriesCount) | **Get** /api/v2/TenantsService/Territories/Count | Get the count of tenant territories
[**GetTenantTerritoryById**](TerritoriesAPI.md#GetTenantTerritoryById) | **Get** /api/v2/TenantsService/Territories/{tenantTerritoryId} | Retrieve a single tenant territory by its ID
[**UpdateTenantTerritory**](TerritoriesAPI.md#UpdateTenantTerritory) | **Put** /api/v2/TenantsService/Territories/{tenantTerritoryId} | Update a tenant territory



## CreateTenantTerritory

> EmptyEnvelope CreateTenantTerritory(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTerritoryCreateDto(tenantTerritoryCreateDto).Execute()

Create a new tenant territory



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
	tenantTerritoryCreateDto := *openapiclient.NewTenantTerritoryCreateDto() // TenantTerritoryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerritoriesAPI.CreateTenantTerritory(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTerritoryCreateDto(tenantTerritoryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerritoriesAPI.CreateTenantTerritory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantTerritory`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TerritoriesAPI.CreateTenantTerritory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantTerritoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantTerritoryCreateDto** | [**TenantTerritoryCreateDto**](TenantTerritoryCreateDto.md) |  | 

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


## DeleteTenantTerritory

> EmptyEnvelope DeleteTenantTerritory(ctx, tenantTerritoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a tenant territory



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
	tenantTerritoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerritoriesAPI.DeleteTenantTerritory(context.Background(), tenantTerritoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerritoriesAPI.DeleteTenantTerritory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantTerritory`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TerritoriesAPI.DeleteTenantTerritory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTerritoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantTerritoryRequest struct via the builder pattern


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


## GetTenantTerritories

> TenantTerritoryDtoListEnvelope GetTenantTerritories(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of tenant territories



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
	resp, r, err := apiClient.TerritoriesAPI.GetTenantTerritories(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerritoriesAPI.GetTenantTerritories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTerritories`: TenantTerritoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TerritoriesAPI.GetTenantTerritories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTerritoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantTerritoryDtoListEnvelope**](TenantTerritoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantTerritoriesCount

> Int32Envelope GetTenantTerritoriesCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of tenant territories



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
	resp, r, err := apiClient.TerritoriesAPI.GetTenantTerritoriesCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerritoriesAPI.GetTenantTerritoriesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTerritoriesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TerritoriesAPI.GetTenantTerritoriesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTerritoriesCountRequest struct via the builder pattern


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


## GetTenantTerritoryById

> TenantTerritoryDtoEnvelope GetTenantTerritoryById(ctx, tenantTerritoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single tenant territory by its ID



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
	tenantTerritoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerritoriesAPI.GetTenantTerritoryById(context.Background(), tenantTerritoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerritoriesAPI.GetTenantTerritoryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTerritoryById`: TenantTerritoryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TerritoriesAPI.GetTenantTerritoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTerritoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTerritoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantTerritoryDtoEnvelope**](TenantTerritoryDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantTerritory

> EmptyEnvelope UpdateTenantTerritory(ctx, tenantTerritoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTerritoryUpdateDto(tenantTerritoryUpdateDto).Execute()

Update a tenant territory



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
	tenantTerritoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	tenantTerritoryUpdateDto := *openapiclient.NewTenantTerritoryUpdateDto() // TenantTerritoryUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TerritoriesAPI.UpdateTenantTerritory(context.Background(), tenantTerritoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTerritoryUpdateDto(tenantTerritoryUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TerritoriesAPI.UpdateTenantTerritory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenantTerritory`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TerritoriesAPI.UpdateTenantTerritory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTerritoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantTerritoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantTerritoryUpdateDto** | [**TenantTerritoryUpdateDto**](TenantTerritoryUpdateDto.md) |  | 

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

