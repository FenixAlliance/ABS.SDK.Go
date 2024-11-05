# \DiscountListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2PricingServiceDiscountListsCountGet**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsCountGet) | **Get** /api/v2/PricingService/DiscountLists/Count | 
[**ApiV2PricingServiceDiscountListsDiscountListIdDelete**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsDiscountListIdDelete) | **Delete** /api/v2/PricingService/DiscountLists/{discountListId} | 
[**ApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGet**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGet) | **Get** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts/Count | 
[**ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDelete**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDelete) | **Delete** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts/{discountListEntryId} | 
[**ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPut**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPut) | **Put** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts/{discountListEntryId} | 
[**ApiV2PricingServiceDiscountListsDiscountListIdDiscountsGet**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsDiscountListIdDiscountsGet) | **Get** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts | 
[**ApiV2PricingServiceDiscountListsDiscountListIdDiscountsPost**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsDiscountListIdDiscountsPost) | **Post** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts | 
[**ApiV2PricingServiceDiscountListsDiscountListIdGet**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsDiscountListIdGet) | **Get** /api/v2/PricingService/DiscountLists/{discountListId} | 
[**ApiV2PricingServiceDiscountListsDiscountListIdPut**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsDiscountListIdPut) | **Put** /api/v2/PricingService/DiscountLists/{discountListId} | 
[**ApiV2PricingServiceDiscountListsGet**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsGet) | **Get** /api/v2/PricingService/DiscountLists | 
[**ApiV2PricingServiceDiscountListsPost**](DiscountListsAPI.md#ApiV2PricingServiceDiscountListsPost) | **Post** /api/v2/PricingService/DiscountLists | 
[**GetDiscountListEntry**](DiscountListsAPI.md#GetDiscountListEntry) | **Get** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts/{discountListEntryId} | 



## ApiV2PricingServiceDiscountListsCountGet

> Int32Envelope ApiV2PricingServiceDiscountListsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
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


## ApiV2PricingServiceDiscountListsDiscountListIdDelete

> EmptyEnvelope ApiV2PricingServiceDiscountListsDiscountListIdDelete(ctx, discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDelete(context.Background(), discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsDiscountListIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsDiscountListIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## ApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGet

> Int32Envelope ApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGet(ctx, discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGet(context.Background(), discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsDiscountListIdDiscountsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDelete

> EmptyEnvelope ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDelete(ctx, discountListId, discountListEntryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountListEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDelete(context.Background(), discountListId, discountListEntryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 
**discountListEntryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPut

> EmptyEnvelope ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPut(ctx, discountListId, discountListEntryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DiscountUpdateDto(discountUpdateDto).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountListEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	discountUpdateDto := *openapiclient.NewDiscountUpdateDto() // DiscountUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPut(context.Background(), discountListId, discountListEntryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DiscountUpdateDto(discountUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 
**discountListEntryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsDiscountListIdDiscountsDiscountListEntryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **discountUpdateDto** | [**DiscountUpdateDto**](DiscountUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServiceDiscountListsDiscountListIdDiscountsGet

> DiscountDtoListEnvelope ApiV2PricingServiceDiscountListsDiscountListIdDiscountsGet(ctx, discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsGet(context.Background(), discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsDiscountListIdDiscountsGet`: DiscountDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsDiscountListIdDiscountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DiscountDtoListEnvelope**](DiscountDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServiceDiscountListsDiscountListIdDiscountsPost

> EmptyEnvelope ApiV2PricingServiceDiscountListsDiscountListIdDiscountsPost(ctx, discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DiscountCreateDto(discountCreateDto).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	discountCreateDto := *openapiclient.NewDiscountCreateDto() // DiscountCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsPost(context.Background(), discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DiscountCreateDto(discountCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsDiscountListIdDiscountsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdDiscountsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsDiscountListIdDiscountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **discountCreateDto** | [**DiscountCreateDto**](DiscountCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServiceDiscountListsDiscountListIdGet

> DiscountListDtoEnvelope ApiV2PricingServiceDiscountListsDiscountListIdGet(ctx, discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdGet(context.Background(), discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsDiscountListIdGet`: DiscountListDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsDiscountListIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DiscountListDtoEnvelope**](DiscountListDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServiceDiscountListsDiscountListIdPut

> EmptyEnvelope ApiV2PricingServiceDiscountListsDiscountListIdPut(ctx, discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DiscountListUpdateDto(discountListUpdateDto).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	discountListUpdateDto := *openapiclient.NewDiscountListUpdateDto() // DiscountListUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdPut(context.Background(), discountListId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DiscountListUpdateDto(discountListUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsDiscountListIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsDiscountListIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsDiscountListIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **discountListUpdateDto** | [**DiscountListUpdateDto**](DiscountListUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServiceDiscountListsGet

> DiscountListDtoListEnvelope ApiV2PricingServiceDiscountListsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsGet`: DiscountListDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DiscountListDtoListEnvelope**](DiscountListDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PricingServiceDiscountListsPost

> EmptyEnvelope ApiV2PricingServiceDiscountListsPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DiscountListCreateDto(discountListCreateDto).Execute()



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
	discountListCreateDto := *openapiclient.NewDiscountListCreateDto() // DiscountListCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.ApiV2PricingServiceDiscountListsPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).DiscountListCreateDto(discountListCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.ApiV2PricingServiceDiscountListsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PricingServiceDiscountListsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.ApiV2PricingServiceDiscountListsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PricingServiceDiscountListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **discountListCreateDto** | [**DiscountListCreateDto**](DiscountListCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscountListEntry

> DiscountDtoEnvelope GetDiscountListEntry(ctx, discountListId, discountListEntryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountListEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.GetDiscountListEntry(context.Background(), discountListId, discountListEntryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.GetDiscountListEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscountListEntry`: DiscountDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.GetDiscountListEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 
**discountListEntryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountListEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**DiscountDtoEnvelope**](DiscountDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

