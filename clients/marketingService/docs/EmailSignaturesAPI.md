# \EmailSignaturesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceEmailSignaturesCountGet**](EmailSignaturesAPI.md#ApiV2MarketingServiceEmailSignaturesCountGet) | **Get** /api/v2/MarketingService/EmailSignatures/Count | 
[**ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete**](EmailSignaturesAPI.md#ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete) | **Delete** /api/v2/MarketingService/EmailSignatures/{emailsignatureId} | 
[**ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet**](EmailSignaturesAPI.md#ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet) | **Get** /api/v2/MarketingService/EmailSignatures/{emailsignatureId} | 
[**ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut**](EmailSignaturesAPI.md#ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut) | **Put** /api/v2/MarketingService/EmailSignatures/{emailsignatureId} | 
[**ApiV2MarketingServiceEmailSignaturesGet**](EmailSignaturesAPI.md#ApiV2MarketingServiceEmailSignaturesGet) | **Get** /api/v2/MarketingService/EmailSignatures | 
[**ApiV2MarketingServiceEmailSignaturesPost**](EmailSignaturesAPI.md#ApiV2MarketingServiceEmailSignaturesPost) | **Post** /api/v2/MarketingService/EmailSignatures | 



## ApiV2MarketingServiceEmailSignaturesCountGet

> Int32Envelope ApiV2MarketingServiceEmailSignaturesCountGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesCountGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailSignaturesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailSignaturesCountGetRequest struct via the builder pattern


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


## ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete

> EmptyEnvelope ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete(ctx, emailsignatureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailsignatureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete(context.Background(), emailsignatureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailsignatureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailSignaturesEmailsignatureIdDeleteRequest struct via the builder pattern


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


## ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet

> EmailSignatureDtoEnvelope ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet(ctx, emailsignatureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailsignatureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet(context.Background(), emailsignatureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet`: EmailSignatureDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailsignatureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailSignaturesEmailsignatureIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmailSignatureDtoEnvelope**](EmailSignatureDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut

> EmptyEnvelope ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut(ctx, emailsignatureId).TenantId(tenantId).EmailSignatureUpdateDto(emailSignatureUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailsignatureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	emailSignatureUpdateDto := *openapiclient.NewEmailSignatureUpdateDto() // EmailSignatureUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut(context.Background(), emailsignatureId).TenantId(tenantId).EmailSignatureUpdateDto(emailSignatureUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesEmailsignatureIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailsignatureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailSignaturesEmailsignatureIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **emailSignatureUpdateDto** | [**EmailSignatureUpdateDto**](EmailSignatureUpdateDto.md) |  | 
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


## ApiV2MarketingServiceEmailSignaturesGet

> EmailSignatureDtoListEnvelope ApiV2MarketingServiceEmailSignaturesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailSignaturesGet`: EmailSignatureDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailSignaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmailSignatureDtoListEnvelope**](EmailSignatureDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2MarketingServiceEmailSignaturesPost

> EmptyEnvelope ApiV2MarketingServiceEmailSignaturesPost(ctx).TenantId(tenantId).EmailSignatureCreateDto(emailSignatureCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	emailSignatureCreateDto := *openapiclient.NewEmailSignatureCreateDto() // EmailSignatureCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesPost(context.Background()).TenantId(tenantId).EmailSignatureCreateDto(emailSignatureCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceEmailSignaturesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.ApiV2MarketingServiceEmailSignaturesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceEmailSignaturesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **emailSignatureCreateDto** | [**EmailSignatureCreateDto**](EmailSignatureCreateDto.md) |  | 
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

