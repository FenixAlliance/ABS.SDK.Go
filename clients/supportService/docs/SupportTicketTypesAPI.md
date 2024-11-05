# \SupportTicketTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SupportServiceSupportTicketTypesCountGet**](SupportTicketTypesAPI.md#ApiV2SupportServiceSupportTicketTypesCountGet) | **Get** /api/v2/SupportService/SupportTicketTypes/Count | 
[**ApiV2SupportServiceSupportTicketTypesGet**](SupportTicketTypesAPI.md#ApiV2SupportServiceSupportTicketTypesGet) | **Get** /api/v2/SupportService/SupportTicketTypes | 
[**ApiV2SupportServiceSupportTicketTypesPost**](SupportTicketTypesAPI.md#ApiV2SupportServiceSupportTicketTypesPost) | **Post** /api/v2/SupportService/SupportTicketTypes | 
[**ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDelete**](SupportTicketTypesAPI.md#ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDelete) | **Delete** /api/v2/SupportService/SupportTicketTypes/{supportTicketTypeId} | 
[**ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGet**](SupportTicketTypesAPI.md#ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGet) | **Get** /api/v2/SupportService/SupportTicketTypes/{supportTicketTypeId} | 
[**ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPut**](SupportTicketTypesAPI.md#ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPut) | **Put** /api/v2/SupportService/SupportTicketTypes/{supportTicketTypeId} | 



## ApiV2SupportServiceSupportTicketTypesCountGet

> Int32Envelope ApiV2SupportServiceSupportTicketTypesCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketTypesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketTypesCountGetRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportTicketTypesGet

> SupportTicketTypeDtoListEnvelope ApiV2SupportServiceSupportTicketTypesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketTypesGet`: SupportTicketTypeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketTypeDtoListEnvelope**](SupportTicketTypeDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketTypesPost

> EmptyEnvelope ApiV2SupportServiceSupportTicketTypesPost(ctx).SupportTicketTypeCreateDto(supportTicketTypeCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketTypeCreateDto := *openapiclient.NewSupportTicketTypeCreateDto() // SupportTicketTypeCreateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesPost(context.Background()).SupportTicketTypeCreateDto(supportTicketTypeCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketTypesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportTicketTypeCreateDto** | [**SupportTicketTypeCreateDto**](SupportTicketTypeCreateDto.md) |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDelete

> EmptyEnvelope ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDelete(ctx, supportTicketTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDelete(context.Background(), supportTicketTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdDeleteRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGet

> SupportTicketTypeDtoEnvelope ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGet(ctx, supportTicketTypeId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGet(context.Background(), supportTicketTypeId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGet`: SupportTicketTypeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportTicketTypeDtoEnvelope**](SupportTicketTypeDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPut

> EmptyEnvelope ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPut(ctx, supportTicketTypeId).SupportTicketTypeUpdateDto(supportTicketTypeUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportTicketTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportTicketTypeUpdateDto := *openapiclient.NewSupportTicketTypeUpdateDto() // SupportTicketTypeUpdateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPut(context.Background(), supportTicketTypeId).SupportTicketTypeUpdateDto(supportTicketTypeUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportTicketTypesAPI.ApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportTicketTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportTicketTypesSupportTicketTypeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportTicketTypeUpdateDto** | [**SupportTicketTypeUpdateDto**](SupportTicketTypeUpdateDto.md) |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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

