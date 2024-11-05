# \EmailTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceEmailTemplatesCountGet**](EmailTemplatesAPI.md#ApiV2MarketingServiceEmailTemplatesCountGet) | **Get** /api/v2/MarketingService/EmailTemplates/Count | 
[**ApiV2MarketingServiceEmailTemplatesEmailTemplateIdDelete**](EmailTemplatesAPI.md#ApiV2MarketingServiceEmailTemplatesEmailTemplateIdDelete) | **Delete** /api/v2/MarketingService/EmailTemplates/{emailTemplateId} | 
[**ApiV2MarketingServiceEmailTemplatesEmailTemplateIdGet**](EmailTemplatesAPI.md#ApiV2MarketingServiceEmailTemplatesEmailTemplateIdGet) | **Get** /api/v2/MarketingService/EmailTemplates/{emailTemplateId} | 
[**ApiV2MarketingServiceEmailTemplatesEmailTemplateIdPut**](EmailTemplatesAPI.md#ApiV2MarketingServiceEmailTemplatesEmailTemplateIdPut) | **Put** /api/v2/MarketingService/EmailTemplates/{emailTemplateId} | 
[**ApiV2MarketingServiceEmailTemplatesGet**](EmailTemplatesAPI.md#ApiV2MarketingServiceEmailTemplatesGet) | **Get** /api/v2/MarketingService/EmailTemplates | 
[**ApiV2MarketingServiceEmailTemplatesPost**](EmailTemplatesAPI.md#ApiV2MarketingServiceEmailTemplatesPost) | **Post** /api/v2/MarketingService/EmailTemplates | 



## ApiV2MarketingServiceEmailTemplatesCountGet

> Int32Envelope ApiV2MarketingServiceEmailTemplatesCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailTemplatesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailTemplatesCountGetRequest struct via the builder pattern


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


## ApiV2MarketingServiceEmailTemplatesEmailTemplateIdDelete

> EmptyEnvelope ApiV2MarketingServiceEmailTemplatesEmailTemplateIdDelete(ctx, emailTemplateId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailTemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdDelete(context.Background(), emailTemplateId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailTemplatesEmailTemplateIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailTemplatesEmailTemplateIdDeleteRequest struct via the builder pattern


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


## ApiV2MarketingServiceEmailTemplatesEmailTemplateIdGet

> EmailTemplateDtoEnvelope ApiV2MarketingServiceEmailTemplatesEmailTemplateIdGet(ctx, emailTemplateId).TenantId(tenantId).EmailTemplatesId(emailTemplatesId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailTemplatesId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	emailTemplateId := "emailTemplateId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdGet(context.Background(), emailTemplateId).TenantId(tenantId).EmailTemplatesId(emailTemplatesId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailTemplatesEmailTemplateIdGet`: EmailTemplateDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailTemplatesEmailTemplateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **emailTemplatesId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmailTemplateDtoEnvelope**](EmailTemplateDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceEmailTemplatesEmailTemplateIdPut

> EmptyEnvelope ApiV2MarketingServiceEmailTemplatesEmailTemplateIdPut(ctx, emailTemplateId).TenantId(tenantId).EmailTemplateUpdateDto(emailTemplateUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailTemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	emailTemplateUpdateDto := *openapiclient.NewEmailTemplateUpdateDto() // EmailTemplateUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdPut(context.Background(), emailTemplateId).TenantId(tenantId).EmailTemplateUpdateDto(emailTemplateUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailTemplatesEmailTemplateIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesEmailTemplateIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailTemplatesEmailTemplateIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **emailTemplateUpdateDto** | [**EmailTemplateUpdateDto**](EmailTemplateUpdateDto.md) |  | 
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


## ApiV2MarketingServiceEmailTemplatesGet

> EmailTemplateDtoListEnvelope ApiV2MarketingServiceEmailTemplatesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailTemplatesGet`: EmailTemplateDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailTemplatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmailTemplateDtoListEnvelope**](EmailTemplateDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceEmailTemplatesPost

> EmptyEnvelope ApiV2MarketingServiceEmailTemplatesPost(ctx).TenantId(tenantId).EmailTemplateCreateDto(emailTemplateCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailTemplateCreateDto := *openapiclient.NewEmailTemplateCreateDto() // EmailTemplateCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesPost(context.Background()).TenantId(tenantId).EmailTemplateCreateDto(emailTemplateCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailTemplatesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailTemplatesAPI.ApiV2MarketingServiceEmailTemplatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **emailTemplateCreateDto** | [**EmailTemplateCreateDto**](EmailTemplateCreateDto.md) |  | 
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

