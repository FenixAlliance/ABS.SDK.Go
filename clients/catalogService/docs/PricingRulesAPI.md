# \PricingRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePricingRule**](PricingRulesAPI.md#CreatePricingRule) | **Post** /api/v2/CatalogService/PricingRules | Create a new pricing rule
[**DeletePricingRule**](PricingRulesAPI.md#DeletePricingRule) | **Delete** /api/v2/CatalogService/PricingRules/{pricingRuleId} | Delete a pricing rule
[**GetPricingRuleById**](PricingRulesAPI.md#GetPricingRuleById) | **Get** /api/v2/CatalogService/PricingRules/{pricingRuleId} | Get pricing rule by ID
[**GetPricingRules**](PricingRulesAPI.md#GetPricingRules) | **Get** /api/v2/CatalogService/PricingRules | Get all pricing rules
[**UpdatePricingRule**](PricingRulesAPI.md#UpdatePricingRule) | **Put** /api/v2/CatalogService/PricingRules/Update | Update a pricing rule



## CreatePricingRule

> PricingRuleDtoEnvelope CreatePricingRule(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PricingRuleCreateDto(pricingRuleCreateDto).Execute()

Create a new pricing rule



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
	pricingRuleCreateDto := *openapiclient.NewPricingRuleCreateDto() // PricingRuleCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingRulesAPI.CreatePricingRule(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PricingRuleCreateDto(pricingRuleCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingRulesAPI.CreatePricingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePricingRule`: PricingRuleDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricingRulesAPI.CreatePricingRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePricingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **pricingRuleCreateDto** | [**PricingRuleCreateDto**](PricingRuleCreateDto.md) |  | 

### Return type

[**PricingRuleDtoEnvelope**](PricingRuleDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePricingRule

> DeletePricingRule(ctx, pricingRuleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a pricing rule



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
	pricingRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PricingRulesAPI.DeletePricingRule(context.Background(), pricingRuleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingRulesAPI.DeletePricingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePricingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## GetPricingRuleById

> PricingRuleDtoEnvelope GetPricingRuleById(ctx, pricingRuleId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get pricing rule by ID



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
	pricingRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingRulesAPI.GetPricingRuleById(context.Background(), pricingRuleId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingRulesAPI.GetPricingRuleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingRuleById`: PricingRuleDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricingRulesAPI.GetPricingRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PricingRuleDtoEnvelope**](PricingRuleDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingRules

> PricingRuleDtoListEnvelope GetPricingRules(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all pricing rules



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
	resp, r, err := apiClient.PricingRulesAPI.GetPricingRules(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingRulesAPI.GetPricingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingRules`: PricingRuleDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PricingRulesAPI.GetPricingRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PricingRuleDtoListEnvelope**](PricingRuleDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePricingRule

> UpdatePricingRule(ctx).TenantId(tenantId).PricingRuleId(pricingRuleId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PricingRuleUpdateDto(pricingRuleUpdateDto).Execute()

Update a pricing rule



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
	pricingRuleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	pricingRuleUpdateDto := *openapiclient.NewPricingRuleUpdateDto() // PricingRuleUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PricingRulesAPI.UpdatePricingRule(context.Background()).TenantId(tenantId).PricingRuleId(pricingRuleId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PricingRuleUpdateDto(pricingRuleUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingRulesAPI.UpdatePricingRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePricingRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **pricingRuleId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **pricingRuleUpdateDto** | [**PricingRuleUpdateDto**](PricingRuleUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

