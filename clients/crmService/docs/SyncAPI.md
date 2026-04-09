# \SyncAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncCurrentHolderToCurrentTenantCrm**](SyncAPI.md#SyncCurrentHolderToCurrentTenantCrm) | **Post** /api/v2/CrmService/Sync | Sync the current user into the current tenant&#39;s contact list
[**SyncCurrentHolderToTenantCrm**](SyncAPI.md#SyncCurrentHolderToTenantCrm) | **Post** /api/v2/CrmService/Sync/Me | Sync the current user into a tenant&#39;s contact list
[**SyncHolderToTenantCrmAsync**](SyncAPI.md#SyncHolderToTenantCrmAsync) | **Post** /api/v2/CrmService/Sync/User | Sync a user into a tenant&#39;s contact list
[**SyncTenantToTenantCrm**](SyncAPI.md#SyncTenantToTenantCrm) | **Post** /api/v2/CrmService/Sync/Tenant | Sync a tenant into another tenant&#39;s contact list



## SyncCurrentHolderToCurrentTenantCrm

> Envelope SyncCurrentHolderToCurrentTenantCrm(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Sync the current user into the current tenant's contact list



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
	resp, r, err := apiClient.SyncAPI.SyncCurrentHolderToCurrentTenantCrm(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.SyncCurrentHolderToCurrentTenantCrm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncCurrentHolderToCurrentTenantCrm`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.SyncCurrentHolderToCurrentTenantCrm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncCurrentHolderToCurrentTenantCrmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncCurrentHolderToTenantCrm

> Envelope SyncCurrentHolderToTenantCrm(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Sync the current user into a tenant's contact list



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
	resp, r, err := apiClient.SyncAPI.SyncCurrentHolderToTenantCrm(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.SyncCurrentHolderToTenantCrm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncCurrentHolderToTenantCrm`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.SyncCurrentHolderToTenantCrm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncCurrentHolderToTenantCrmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncHolderToTenantCrmAsync

> Envelope SyncHolderToTenantCrmAsync(ctx).TenantId(tenantId).RelatedUserId(relatedUserId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Sync a user into a tenant's contact list



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
	relatedUserId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.SyncHolderToTenantCrmAsync(context.Background()).TenantId(tenantId).RelatedUserId(relatedUserId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.SyncHolderToTenantCrmAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncHolderToTenantCrmAsync`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.SyncHolderToTenantCrmAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncHolderToTenantCrmAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **relatedUserId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncTenantToTenantCrm

> EmptyEnvelope SyncTenantToTenantCrm(ctx).TenantId(tenantId).RelatedTenantId(relatedTenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Sync a tenant into another tenant's contact list



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
	relatedTenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyncAPI.SyncTenantToTenantCrm(context.Background()).TenantId(tenantId).RelatedTenantId(relatedTenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncAPI.SyncTenantToTenantCrm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncTenantToTenantCrm`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SyncAPI.SyncTenantToTenantCrm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncTenantToTenantCrmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **relatedTenantId** | **string** |  | 
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

