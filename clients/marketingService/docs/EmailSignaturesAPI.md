# \EmailSignaturesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailSignatureAsync**](EmailSignaturesAPI.md#CreateEmailSignatureAsync) | **Post** /api/v2/MarketingService/EmailSignatures | Create an email signature
[**DeleteEmailSignatureAsync**](EmailSignaturesAPI.md#DeleteEmailSignatureAsync) | **Delete** /api/v2/MarketingService/EmailSignatures/{emailsignatureId} | Delete an email signature
[**GetEmailSignatureDetailsAsync**](EmailSignaturesAPI.md#GetEmailSignatureDetailsAsync) | **Get** /api/v2/MarketingService/EmailSignatures/{emailsignatureId} | Get email signature by ID
[**GetEmailSignaturesCountAsync**](EmailSignaturesAPI.md#GetEmailSignaturesCountAsync) | **Get** /api/v2/MarketingService/EmailSignatures/Count | Get email signatures count
[**GetEmailSignaturesODataAsync**](EmailSignaturesAPI.md#GetEmailSignaturesODataAsync) | **Get** /api/v2/MarketingService/EmailSignatures | Get email signatures
[**UpdateEmailSignatureAsync**](EmailSignaturesAPI.md#UpdateEmailSignatureAsync) | **Put** /api/v2/MarketingService/EmailSignatures/{emailsignatureId} | Update an email signature



## CreateEmailSignatureAsync

> EmptyEnvelope CreateEmailSignatureAsync(ctx).TenantId(tenantId).EmailSignatureCreateDto(emailSignatureCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create an email signature



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
	resp, r, err := apiClient.EmailSignaturesAPI.CreateEmailSignatureAsync(context.Background()).TenantId(tenantId).EmailSignatureCreateDto(emailSignatureCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.CreateEmailSignatureAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmailSignatureAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.CreateEmailSignatureAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailSignatureAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **emailSignatureCreateDto** | [**EmailSignatureCreateDto**](EmailSignatureCreateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailSignatureAsync

> EmptyEnvelope DeleteEmailSignatureAsync(ctx, emailsignatureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an email signature



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
	resp, r, err := apiClient.EmailSignaturesAPI.DeleteEmailSignatureAsync(context.Background(), emailsignatureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.DeleteEmailSignatureAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEmailSignatureAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.DeleteEmailSignatureAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailsignatureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailSignatureAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSignatureDetailsAsync

> EmailSignatureDtoEnvelope GetEmailSignatureDetailsAsync(ctx, emailsignatureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get email signature by ID



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
	resp, r, err := apiClient.EmailSignaturesAPI.GetEmailSignatureDetailsAsync(context.Background(), emailsignatureId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.GetEmailSignatureDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailSignatureDetailsAsync`: EmailSignatureDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.GetEmailSignatureDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailsignatureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSignatureDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmailSignatureDtoEnvelope**](EmailSignatureDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSignaturesCountAsync

> Int32Envelope GetEmailSignaturesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get email signatures count



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
	resp, r, err := apiClient.EmailSignaturesAPI.GetEmailSignaturesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.GetEmailSignaturesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailSignaturesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.GetEmailSignaturesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSignaturesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
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


## GetEmailSignaturesODataAsync

> EmailSignatureDtoListEnvelope GetEmailSignaturesODataAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get email signatures



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
	resp, r, err := apiClient.EmailSignaturesAPI.GetEmailSignaturesODataAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.GetEmailSignaturesODataAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailSignaturesODataAsync`: EmailSignatureDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.GetEmailSignaturesODataAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSignaturesODataAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmailSignatureDtoListEnvelope**](EmailSignatureDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailSignatureAsync

> EmptyEnvelope UpdateEmailSignatureAsync(ctx, emailsignatureId).TenantId(tenantId).EmailSignatureUpdateDto(emailSignatureUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update an email signature



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
	resp, r, err := apiClient.EmailSignaturesAPI.UpdateEmailSignatureAsync(context.Background(), emailsignatureId).TenantId(tenantId).EmailSignatureUpdateDto(emailSignatureUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSignaturesAPI.UpdateEmailSignatureAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEmailSignatureAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailSignaturesAPI.UpdateEmailSignatureAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailsignatureId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailSignatureAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **emailSignatureUpdateDto** | [**EmailSignatureUpdateDto**](EmailSignatureUpdateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

