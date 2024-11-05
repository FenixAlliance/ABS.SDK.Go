# \SupportEntitlementsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SupportServiceSupportEntitlementsCountGet**](SupportEntitlementsAPI.md#ApiV2SupportServiceSupportEntitlementsCountGet) | **Get** /api/v2/SupportService/SupportEntitlements/Count | 
[**ApiV2SupportServiceSupportEntitlementsGet**](SupportEntitlementsAPI.md#ApiV2SupportServiceSupportEntitlementsGet) | **Get** /api/v2/SupportService/SupportEntitlements | 
[**ApiV2SupportServiceSupportEntitlementsPost**](SupportEntitlementsAPI.md#ApiV2SupportServiceSupportEntitlementsPost) | **Post** /api/v2/SupportService/SupportEntitlements | 
[**ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete**](SupportEntitlementsAPI.md#ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete) | **Delete** /api/v2/SupportService/SupportEntitlements/{supportEntitlementId} | 
[**ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet**](SupportEntitlementsAPI.md#ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet) | **Get** /api/v2/SupportService/SupportEntitlements/{supportEntitlementId} | 
[**ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut**](SupportEntitlementsAPI.md#ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut) | **Put** /api/v2/SupportService/SupportEntitlements/{supportEntitlementId} | 



## ApiV2SupportServiceSupportEntitlementsCountGet

> Int32Envelope ApiV2SupportServiceSupportEntitlementsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportEntitlementsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportEntitlementsCountGetRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportEntitlementsGet

> SupportEntitlementDtoListEnvelope ApiV2SupportServiceSupportEntitlementsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportEntitlementsGet`: SupportEntitlementDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportEntitlementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportEntitlementDtoListEnvelope**](SupportEntitlementDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportEntitlementsPost

> EmptyEnvelope ApiV2SupportServiceSupportEntitlementsPost(ctx).SupportEntitlementCreateDto(supportEntitlementCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportEntitlementCreateDto := *openapiclient.NewSupportEntitlementCreateDto() // SupportEntitlementCreateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsPost(context.Background()).SupportEntitlementCreateDto(supportEntitlementCreateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportEntitlementsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportEntitlementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportEntitlementCreateDto** | [**SupportEntitlementCreateDto**](SupportEntitlementCreateDto.md) |  | 
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


## ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete

> EmptyEnvelope ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete(ctx, supportEntitlementId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportEntitlementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete(context.Background(), supportEntitlementId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportEntitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportEntitlementsSupportEntitlementIdDeleteRequest struct via the builder pattern


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


## ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet

> SupportEntitlementDtoEnvelope ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet(ctx, supportEntitlementId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportEntitlementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet(context.Background(), supportEntitlementId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet`: SupportEntitlementDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportEntitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportEntitlementsSupportEntitlementIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SupportEntitlementDtoEnvelope**](SupportEntitlementDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut

> EmptyEnvelope ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut(ctx, supportEntitlementId).SupportEntitlementUpdateDto(supportEntitlementUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	supportEntitlementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	supportEntitlementUpdateDto := *openapiclient.NewSupportEntitlementUpdateDto() // SupportEntitlementUpdateDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut(context.Background(), supportEntitlementId).SupportEntitlementUpdateDto(supportEntitlementUpdateDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SupportEntitlementsAPI.ApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supportEntitlementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SupportServiceSupportEntitlementsSupportEntitlementIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportEntitlementUpdateDto** | [**SupportEntitlementUpdateDto**](SupportEntitlementUpdateDto.md) |  | 
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

