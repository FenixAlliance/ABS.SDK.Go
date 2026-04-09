# \ItemRefundPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountItemRefundPoliciesAsync**](ItemRefundPoliciesAPI.md#CountItemRefundPoliciesAsync) | **Get** /api/v2/CatalogService/ItemRefundPolicies/Count | Count item refund policies
[**GetItemRefundPoliciesAsync**](ItemRefundPoliciesAPI.md#GetItemRefundPoliciesAsync) | **Get** /api/v2/CatalogService/ItemRefundPolicies | Get item refund policies
[**GetItemRefundPolicyByIdAsync**](ItemRefundPoliciesAPI.md#GetItemRefundPolicyByIdAsync) | **Get** /api/v2/CatalogService/ItemRefundPolicies/{itemRefundPolicyId} | Get item refund policy by ID
[**RelateItemToRefundPolicyAsync**](ItemRefundPoliciesAPI.md#RelateItemToRefundPolicyAsync) | **Post** /api/v2/CatalogService/ItemRefundPolicies | Relate item to refund policy
[**RemoveRefundPolicyFromItemAsync**](ItemRefundPoliciesAPI.md#RemoveRefundPolicyFromItemAsync) | **Delete** /api/v2/CatalogService/ItemRefundPolicies/{itemRefundPolicyId} | Remove refund policy from item



## CountItemRefundPoliciesAsync

> Int32Envelope CountItemRefundPoliciesAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count item refund policies



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
	resp, r, err := apiClient.ItemRefundPoliciesAPI.CountItemRefundPoliciesAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRefundPoliciesAPI.CountItemRefundPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountItemRefundPoliciesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRefundPoliciesAPI.CountItemRefundPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountItemRefundPoliciesAsyncRequest struct via the builder pattern


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


## GetItemRefundPoliciesAsync

> ItemRefundPolicyDtoListEnvelope GetItemRefundPoliciesAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item refund policies



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
	resp, r, err := apiClient.ItemRefundPoliciesAPI.GetItemRefundPoliciesAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRefundPoliciesAPI.GetItemRefundPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemRefundPoliciesAsync`: ItemRefundPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRefundPoliciesAPI.GetItemRefundPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRefundPoliciesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemRefundPolicyDtoListEnvelope**](ItemRefundPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemRefundPolicyByIdAsync

> ItemRefundPolicyDtoEnvelope GetItemRefundPolicyByIdAsync(ctx, itemRefundPolicyId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item refund policy by ID



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
	itemRefundPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemRefundPoliciesAPI.GetItemRefundPolicyByIdAsync(context.Background(), itemRefundPolicyId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRefundPoliciesAPI.GetItemRefundPolicyByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemRefundPolicyByIdAsync`: ItemRefundPolicyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemRefundPoliciesAPI.GetItemRefundPolicyByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemRefundPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRefundPolicyByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemRefundPolicyDtoEnvelope**](ItemRefundPolicyDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelateItemToRefundPolicyAsync

> RelateItemToRefundPolicyAsync(ctx).TenantId(tenantId).ItemId(itemId).RefundPolicyId(refundPolicyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Relate item to refund policy



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
	refundPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemRefundPoliciesAPI.RelateItemToRefundPolicyAsync(context.Background()).TenantId(tenantId).ItemId(itemId).RefundPolicyId(refundPolicyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRefundPoliciesAPI.RelateItemToRefundPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelateItemToRefundPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **itemId** | **string** |  | 
 **refundPolicyId** | **string** |  | 
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


## RemoveRefundPolicyFromItemAsync

> RemoveRefundPolicyFromItemAsync(ctx, itemRefundPolicyId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove refund policy from item



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
	itemRefundPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemRefundPoliciesAPI.RemoveRefundPolicyFromItemAsync(context.Background(), itemRefundPolicyId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRefundPoliciesAPI.RemoveRefundPolicyFromItemAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemRefundPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRefundPolicyFromItemAsyncRequest struct via the builder pattern


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

