# \ItemWarrantyPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountItemWarrantyPoliciesAsync**](ItemWarrantyPoliciesAPI.md#CountItemWarrantyPoliciesAsync) | **Get** /api/v2/CatalogService/ItemWarrantyPolicies/Count | Count item warranty policies
[**GetItemWarrantyPoliciesAsync**](ItemWarrantyPoliciesAPI.md#GetItemWarrantyPoliciesAsync) | **Get** /api/v2/CatalogService/ItemWarrantyPolicies | Get item warranty policies
[**GetItemWarrantyPolicyByIdAsync**](ItemWarrantyPoliciesAPI.md#GetItemWarrantyPolicyByIdAsync) | **Get** /api/v2/CatalogService/ItemWarrantyPolicies/{itemWarrantyPolicyId} | Get item warranty policy by ID
[**RelateItemToWarrantyPolicyAsync**](ItemWarrantyPoliciesAPI.md#RelateItemToWarrantyPolicyAsync) | **Post** /api/v2/CatalogService/ItemWarrantyPolicies | Relate item to warranty policy
[**RemoveWarrantyPolicyFromItemAsync**](ItemWarrantyPoliciesAPI.md#RemoveWarrantyPolicyFromItemAsync) | **Delete** /api/v2/CatalogService/ItemWarrantyPolicies/{itemWarrantyPolicyId} | Remove warranty policy from item



## CountItemWarrantyPoliciesAsync

> Int32Envelope CountItemWarrantyPoliciesAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count item warranty policies



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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemWarrantyPoliciesAPI.CountItemWarrantyPoliciesAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemWarrantyPoliciesAPI.CountItemWarrantyPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountItemWarrantyPoliciesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemWarrantyPoliciesAPI.CountItemWarrantyPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountItemWarrantyPoliciesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** |  | 
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


## GetItemWarrantyPoliciesAsync

> ItemWarrantyPolicyDtoListEnvelope GetItemWarrantyPoliciesAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item warranty policies



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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemWarrantyPoliciesAPI.GetItemWarrantyPoliciesAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemWarrantyPoliciesAPI.GetItemWarrantyPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemWarrantyPoliciesAsync`: ItemWarrantyPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemWarrantyPoliciesAPI.GetItemWarrantyPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemWarrantyPoliciesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemWarrantyPolicyDtoListEnvelope**](ItemWarrantyPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemWarrantyPolicyByIdAsync

> ItemWarrantyPolicyDtoEnvelope GetItemWarrantyPolicyByIdAsync(ctx, itemWarrantyPolicyId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item warranty policy by ID



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
	itemWarrantyPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemWarrantyPoliciesAPI.GetItemWarrantyPolicyByIdAsync(context.Background(), itemWarrantyPolicyId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemWarrantyPoliciesAPI.GetItemWarrantyPolicyByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemWarrantyPolicyByIdAsync`: ItemWarrantyPolicyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemWarrantyPoliciesAPI.GetItemWarrantyPolicyByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemWarrantyPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemWarrantyPolicyByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemWarrantyPolicyDtoEnvelope**](ItemWarrantyPolicyDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelateItemToWarrantyPolicyAsync

> RelateItemToWarrantyPolicyAsync(ctx).TenantId(tenantId).ItemId(itemId).WarrantyPolicyId(warrantyPolicyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Relate item to warranty policy



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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	warrantyPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemWarrantyPoliciesAPI.RelateItemToWarrantyPolicyAsync(context.Background()).TenantId(tenantId).ItemId(itemId).WarrantyPolicyId(warrantyPolicyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemWarrantyPoliciesAPI.RelateItemToWarrantyPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelateItemToWarrantyPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **itemId** | **string** |  | 
 **warrantyPolicyId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWarrantyPolicyFromItemAsync

> RemoveWarrantyPolicyFromItemAsync(ctx, itemWarrantyPolicyId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove warranty policy from item



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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemWarrantyPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemWarrantyPoliciesAPI.RemoveWarrantyPolicyFromItemAsync(context.Background(), itemWarrantyPolicyId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemWarrantyPoliciesAPI.RemoveWarrantyPolicyFromItemAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemWarrantyPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWarrantyPolicyFromItemAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **itemId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

