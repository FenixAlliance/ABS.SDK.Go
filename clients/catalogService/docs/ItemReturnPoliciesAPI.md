# \ItemReturnPoliciesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountItemReturnPoliciesAsync**](ItemReturnPoliciesAPI.md#CountItemReturnPoliciesAsync) | **Get** /api/v2/CatalogService/ItemReturnPolicies/Count | Count item return policies
[**GetItemReturnPoliciesAsync**](ItemReturnPoliciesAPI.md#GetItemReturnPoliciesAsync) | **Get** /api/v2/CatalogService/ItemReturnPolicies | Get item return policies
[**GetItemReturnPolicyByIdAsync**](ItemReturnPoliciesAPI.md#GetItemReturnPolicyByIdAsync) | **Get** /api/v2/CatalogService/ItemReturnPolicies/{itemReturnPolicyId} | Get item return policy by ID
[**RelateItemToReturnPolicyAsync**](ItemReturnPoliciesAPI.md#RelateItemToReturnPolicyAsync) | **Post** /api/v2/CatalogService/ItemReturnPolicies | Relate item to return policy
[**RemoveReturnPolicyFromItemAsync**](ItemReturnPoliciesAPI.md#RemoveReturnPolicyFromItemAsync) | **Delete** /api/v2/CatalogService/ItemReturnPolicies/{itemReturnPolicyId} | Remove return policy from item



## CountItemReturnPoliciesAsync

> Int32Envelope CountItemReturnPoliciesAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count item return policies



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
	resp, r, err := apiClient.ItemReturnPoliciesAPI.CountItemReturnPoliciesAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReturnPoliciesAPI.CountItemReturnPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountItemReturnPoliciesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemReturnPoliciesAPI.CountItemReturnPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountItemReturnPoliciesAsyncRequest struct via the builder pattern


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


## GetItemReturnPoliciesAsync

> ItemReturnPolicyDtoListEnvelope GetItemReturnPoliciesAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item return policies



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
	resp, r, err := apiClient.ItemReturnPoliciesAPI.GetItemReturnPoliciesAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReturnPoliciesAPI.GetItemReturnPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemReturnPoliciesAsync`: ItemReturnPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemReturnPoliciesAPI.GetItemReturnPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemReturnPoliciesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemReturnPolicyDtoListEnvelope**](ItemReturnPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemReturnPolicyByIdAsync

> ItemReturnPolicyDtoEnvelope GetItemReturnPolicyByIdAsync(ctx, itemReturnPolicyId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item return policy by ID



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
	itemReturnPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemReturnPoliciesAPI.GetItemReturnPolicyByIdAsync(context.Background(), itemReturnPolicyId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReturnPoliciesAPI.GetItemReturnPolicyByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemReturnPolicyByIdAsync`: ItemReturnPolicyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemReturnPoliciesAPI.GetItemReturnPolicyByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemReturnPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemReturnPolicyByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemReturnPolicyDtoEnvelope**](ItemReturnPolicyDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelateItemToReturnPolicyAsync

> RelateItemToReturnPolicyAsync(ctx).TenantId(tenantId).ItemId(itemId).ReturnPolicyId(returnPolicyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Relate item to return policy



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
	returnPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemReturnPoliciesAPI.RelateItemToReturnPolicyAsync(context.Background()).TenantId(tenantId).ItemId(itemId).ReturnPolicyId(returnPolicyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReturnPoliciesAPI.RelateItemToReturnPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelateItemToReturnPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **itemId** | **string** |  | 
 **returnPolicyId** | **string** |  | 
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


## RemoveReturnPolicyFromItemAsync

> RemoveReturnPolicyFromItemAsync(ctx, itemReturnPolicyId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove return policy from item



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
	itemReturnPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemReturnPoliciesAPI.RemoveReturnPolicyFromItemAsync(context.Background(), itemReturnPolicyId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemReturnPoliciesAPI.RemoveReturnPolicyFromItemAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemReturnPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReturnPolicyFromItemAsyncRequest struct via the builder pattern


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

