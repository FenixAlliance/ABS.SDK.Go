# \EmailGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceEmailGroupsCountGet**](EmailGroupsAPI.md#ApiV2MarketingServiceEmailGroupsCountGet) | **Get** /api/v2/MarketingService/EmailGroups/Count | 
[**ApiV2MarketingServiceEmailGroupsEmailgroupIdDelete**](EmailGroupsAPI.md#ApiV2MarketingServiceEmailGroupsEmailgroupIdDelete) | **Delete** /api/v2/MarketingService/EmailGroups/{emailgroupId} | 
[**ApiV2MarketingServiceEmailGroupsEmailgroupIdGet**](EmailGroupsAPI.md#ApiV2MarketingServiceEmailGroupsEmailgroupIdGet) | **Get** /api/v2/MarketingService/EmailGroups/{emailgroupId} | 
[**ApiV2MarketingServiceEmailGroupsEmailgroupIdPut**](EmailGroupsAPI.md#ApiV2MarketingServiceEmailGroupsEmailgroupIdPut) | **Put** /api/v2/MarketingService/EmailGroups/{emailgroupId} | 
[**ApiV2MarketingServiceEmailGroupsGet**](EmailGroupsAPI.md#ApiV2MarketingServiceEmailGroupsGet) | **Get** /api/v2/MarketingService/EmailGroups | 
[**ApiV2MarketingServiceEmailGroupsPost**](EmailGroupsAPI.md#ApiV2MarketingServiceEmailGroupsPost) | **Post** /api/v2/MarketingService/EmailGroups | 



## ApiV2MarketingServiceEmailGroupsCountGet

> Int32Envelope ApiV2MarketingServiceEmailGroupsCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailGroupsCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailGroupsCountGetRequest struct via the builder pattern


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


## ApiV2MarketingServiceEmailGroupsEmailgroupIdDelete

> EmptyEnvelope ApiV2MarketingServiceEmailGroupsEmailgroupIdDelete(ctx, emailgroupId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailgroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdDelete(context.Background(), emailgroupId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailGroupsEmailgroupIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailgroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailGroupsEmailgroupIdDeleteRequest struct via the builder pattern


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


## ApiV2MarketingServiceEmailGroupsEmailgroupIdGet

> EmailGroupDtoEnvelope ApiV2MarketingServiceEmailGroupsEmailgroupIdGet(ctx, emailgroupId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailgroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdGet(context.Background(), emailgroupId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailGroupsEmailgroupIdGet`: EmailGroupDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailgroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailGroupsEmailgroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmailGroupDtoEnvelope**](EmailGroupDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceEmailGroupsEmailgroupIdPut

> EmptyEnvelope ApiV2MarketingServiceEmailGroupsEmailgroupIdPut(ctx, emailgroupId).TenantId(tenantId).EmailGroupUpdateDto(emailGroupUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailgroupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	emailGroupUpdateDto := *openapiclient.NewEmailGroupUpdateDto() // EmailGroupUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdPut(context.Background(), emailgroupId).TenantId(tenantId).EmailGroupUpdateDto(emailGroupUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailGroupsEmailgroupIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsEmailgroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailgroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailGroupsEmailgroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **emailGroupUpdateDto** | [**EmailGroupUpdateDto**](EmailGroupUpdateDto.md) |  | 
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


## ApiV2MarketingServiceEmailGroupsGet

> EmailGroupDtoListEnvelope ApiV2MarketingServiceEmailGroupsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailGroupsGet`: EmailGroupDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmailGroupDtoListEnvelope**](EmailGroupDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceEmailGroupsPost

> EmptyEnvelope ApiV2MarketingServiceEmailGroupsPost(ctx).TenantId(tenantId).EmailGroupCreateDto(emailGroupCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailGroupCreateDto := *openapiclient.NewEmailGroupCreateDto() // EmailGroupCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsPost(context.Background()).TenantId(tenantId).EmailGroupCreateDto(emailGroupCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailGroupsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailGroupsAPI.ApiV2MarketingServiceEmailGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **emailGroupCreateDto** | [**EmailGroupCreateDto**](EmailGroupCreateDto.md) |  | 
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

