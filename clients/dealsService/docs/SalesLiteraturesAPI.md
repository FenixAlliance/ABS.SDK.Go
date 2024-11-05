# \SalesLiteraturesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2DealsServiceSalesLiteraturesExtendedGet**](SalesLiteraturesAPI.md#ApiV2DealsServiceSalesLiteraturesExtendedGet) | **Get** /api/v2/DealsService/SalesLiteratures/Extended | 
[**ApiV2DealsServiceSalesLiteraturesGet**](SalesLiteraturesAPI.md#ApiV2DealsServiceSalesLiteraturesGet) | **Get** /api/v2/DealsService/SalesLiteratures | 
[**ApiV2DealsServiceSalesLiteraturesPost**](SalesLiteraturesAPI.md#ApiV2DealsServiceSalesLiteraturesPost) | **Post** /api/v2/DealsService/SalesLiteratures | 
[**ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDelete**](SalesLiteraturesAPI.md#ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDelete) | **Delete** /api/v2/DealsService/SalesLiteratures/{salesLiteratureId} | 
[**ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGet**](SalesLiteraturesAPI.md#ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGet) | **Get** /api/v2/DealsService/SalesLiteratures/{salesLiteratureId} | 
[**ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPut**](SalesLiteraturesAPI.md#ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPut) | **Put** /api/v2/DealsService/SalesLiteratures/{salesLiteratureId} | 



## ApiV2DealsServiceSalesLiteraturesExtendedGet

> ExtendedSalesLiteratureDtoListEnvelope ApiV2DealsServiceSalesLiteraturesExtendedGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesExtendedGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesExtendedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceSalesLiteraturesExtendedGet`: ExtendedSalesLiteratureDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesExtendedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceSalesLiteraturesExtendedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExtendedSalesLiteratureDtoListEnvelope**](ExtendedSalesLiteratureDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceSalesLiteraturesGet

> SalesLiteratureDtoListEnvelope ApiV2DealsServiceSalesLiteraturesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceSalesLiteraturesGet`: SalesLiteratureDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceSalesLiteraturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SalesLiteratureDtoListEnvelope**](SalesLiteratureDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceSalesLiteraturesPost

> EmptyEnvelope ApiV2DealsServiceSalesLiteraturesPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SalesLiteratureCreateDto(salesLiteratureCreateDto).Execute()



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
	salesLiteratureCreateDto := *openapiclient.NewSalesLiteratureCreateDto() // SalesLiteratureCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SalesLiteratureCreateDto(salesLiteratureCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceSalesLiteraturesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceSalesLiteraturesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **salesLiteratureCreateDto** | [**SalesLiteratureCreateDto**](SalesLiteratureCreateDto.md) |  | 

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


## ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDelete

> EmptyEnvelope ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDelete(ctx, salesLiteratureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	salesLiteratureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDelete(context.Background(), salesLiteratureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**salesLiteratureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceSalesLiteraturesSalesLiteratureIdDeleteRequest struct via the builder pattern


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


## ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGet

> SalesLiteratureDtoEnvelope ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGet(ctx, salesLiteratureId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	salesLiteratureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGet(context.Background(), salesLiteratureId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGet`: SalesLiteratureDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**salesLiteratureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceSalesLiteraturesSalesLiteratureIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SalesLiteratureDtoEnvelope**](SalesLiteratureDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPut

> EmptyEnvelope ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPut(ctx, salesLiteratureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SalesLiteratureUpdateDto(salesLiteratureUpdateDto).Execute()



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
	salesLiteratureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	salesLiteratureUpdateDto := *openapiclient.NewSalesLiteratureUpdateDto() // SalesLiteratureUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPut(context.Background(), salesLiteratureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SalesLiteratureUpdateDto(salesLiteratureUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SalesLiteraturesAPI.ApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**salesLiteratureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2DealsServiceSalesLiteraturesSalesLiteratureIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **salesLiteratureUpdateDto** | [**SalesLiteratureUpdateDto**](SalesLiteratureUpdateDto.md) |  | 

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

