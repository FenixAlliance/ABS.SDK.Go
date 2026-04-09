# \ItemTaxPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountItemTaxPoliciesAsync**](ItemTaxPoliciesAPI.md#CountItemTaxPoliciesAsync) | **Get** /api/v2/CatalogService/ItemTaxPolicies/Count | Count item tax policies
[**GetItemTaxPoliciesAsync**](ItemTaxPoliciesAPI.md#GetItemTaxPoliciesAsync) | **Get** /api/v2/CatalogService/ItemTaxPolicies | Get item tax policies
[**GetItemTaxPolicyByIdAsync**](ItemTaxPoliciesAPI.md#GetItemTaxPolicyByIdAsync) | **Get** /api/v2/CatalogService/ItemTaxPolicies/{itemTaxPolicyId} | Get item tax policy by ID
[**RelateItemToTaxPolicyAsync**](ItemTaxPoliciesAPI.md#RelateItemToTaxPolicyAsync) | **Post** /api/v2/CatalogService/ItemTaxPolicies | Relate item to tax policy
[**RemoveTaxPolicyFromItemAsync**](ItemTaxPoliciesAPI.md#RemoveTaxPolicyFromItemAsync) | **Delete** /api/v2/CatalogService/ItemTaxPolicies/{itemTaxPolicyId} | Remove tax policy from item



## CountItemTaxPoliciesAsync

> Int32Envelope CountItemTaxPoliciesAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count item tax policies



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
	resp, r, err := apiClient.ItemTaxPoliciesAPI.CountItemTaxPoliciesAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTaxPoliciesAPI.CountItemTaxPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountItemTaxPoliciesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ItemTaxPoliciesAPI.CountItemTaxPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountItemTaxPoliciesAsyncRequest struct via the builder pattern


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


## GetItemTaxPoliciesAsync

> ItemTaxPolicyDtoListEnvelope GetItemTaxPoliciesAsync(ctx).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item tax policies



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
	resp, r, err := apiClient.ItemTaxPoliciesAPI.GetItemTaxPoliciesAsync(context.Background()).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTaxPoliciesAPI.GetItemTaxPoliciesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemTaxPoliciesAsync`: ItemTaxPolicyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemTaxPoliciesAPI.GetItemTaxPoliciesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemTaxPoliciesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemTaxPolicyDtoListEnvelope**](ItemTaxPolicyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemTaxPolicyByIdAsync

> ItemTaxPolicyDtoEnvelope GetItemTaxPolicyByIdAsync(ctx, itemTaxPolicyId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get item tax policy by ID



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
	itemTaxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemTaxPoliciesAPI.GetItemTaxPolicyByIdAsync(context.Background(), itemTaxPolicyId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTaxPoliciesAPI.GetItemTaxPolicyByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemTaxPolicyByIdAsync`: ItemTaxPolicyDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ItemTaxPoliciesAPI.GetItemTaxPolicyByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemTaxPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemTaxPolicyByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ItemTaxPolicyDtoEnvelope**](ItemTaxPolicyDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelateItemToTaxPolicyAsync

> RelateItemToTaxPolicyAsync(ctx).TenantId(tenantId).ItemId(itemId).TaxPolicyId(taxPolicyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Relate item to tax policy



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
	taxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemTaxPoliciesAPI.RelateItemToTaxPolicyAsync(context.Background()).TenantId(tenantId).ItemId(itemId).TaxPolicyId(taxPolicyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTaxPoliciesAPI.RelateItemToTaxPolicyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelateItemToTaxPolicyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **itemId** | **string** |  | 
 **taxPolicyId** | **string** |  | 
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


## RemoveTaxPolicyFromItemAsync

> RemoveTaxPolicyFromItemAsync(ctx, itemTaxPolicyId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Remove tax policy from item



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
	itemTaxPolicyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemTaxPoliciesAPI.RemoveTaxPolicyFromItemAsync(context.Background(), itemTaxPolicyId).TenantId(tenantId).ItemId(itemId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemTaxPoliciesAPI.RemoveTaxPolicyFromItemAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemTaxPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTaxPolicyFromItemAsyncRequest struct via the builder pattern


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

