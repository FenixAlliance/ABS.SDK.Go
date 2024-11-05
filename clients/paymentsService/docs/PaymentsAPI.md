# \PaymentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2PaymentsServicePaymentsGet**](PaymentsAPI.md#ApiV2PaymentsServicePaymentsGet) | **Get** /api/v2/PaymentsService/Payments | 
[**ApiV2PaymentsServicePaymentsPaymentIdDelete**](PaymentsAPI.md#ApiV2PaymentsServicePaymentsPaymentIdDelete) | **Delete** /api/v2/PaymentsService/Payments/{paymentId} | 
[**ApiV2PaymentsServicePaymentsPaymentIdDetailsGet**](PaymentsAPI.md#ApiV2PaymentsServicePaymentsPaymentIdDetailsGet) | **Get** /api/v2/PaymentsService/Payments/{paymentId}/Details | 
[**ApiV2PaymentsServicePaymentsPaymentIdGet**](PaymentsAPI.md#ApiV2PaymentsServicePaymentsPaymentIdGet) | **Get** /api/v2/PaymentsService/Payments/{paymentId} | 
[**ApiV2PaymentsServicePaymentsPaymentIdPut**](PaymentsAPI.md#ApiV2PaymentsServicePaymentsPaymentIdPut) | **Put** /api/v2/PaymentsService/Payments/{paymentId} | 
[**ApiV2PaymentsServicePaymentsPost**](PaymentsAPI.md#ApiV2PaymentsServicePaymentsPost) | **Post** /api/v2/PaymentsService/Payments | 



## ApiV2PaymentsServicePaymentsGet

> PaymentDtoListEnvelope ApiV2PaymentsServicePaymentsGet(ctx).TenantId(tenantId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ApiV2PaymentsServicePaymentsGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ApiV2PaymentsServicePaymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PaymentsServicePaymentsGet`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ApiV2PaymentsServicePaymentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PaymentsServicePaymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PaymentsServicePaymentsPaymentIdDelete

> EmptyEnvelope ApiV2PaymentsServicePaymentsPaymentIdDelete(ctx, paymentId).TenantId(tenantId).Execute()



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
	paymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdDelete(context.Background(), paymentId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PaymentsServicePaymentsPaymentIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PaymentsServicePaymentsPaymentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## ApiV2PaymentsServicePaymentsPaymentIdDetailsGet

> PaymentDtoListEnvelope ApiV2PaymentsServicePaymentsPaymentIdDetailsGet(ctx, paymentId).Execute()



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
	paymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdDetailsGet(context.Background(), paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PaymentsServicePaymentsPaymentIdDetailsGet`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PaymentsServicePaymentsPaymentIdDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PaymentsServicePaymentsPaymentIdGet

> PaymentDtoListEnvelope ApiV2PaymentsServicePaymentsPaymentIdGet(ctx, paymentId).Execute()



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
	paymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdGet(context.Background(), paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PaymentsServicePaymentsPaymentIdGet`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PaymentsServicePaymentsPaymentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2PaymentsServicePaymentsPaymentIdPut

> EmptyEnvelope ApiV2PaymentsServicePaymentsPaymentIdPut(ctx, paymentId).TenantId(tenantId).PaymentUpdateDto(paymentUpdateDto).Execute()



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
	paymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	paymentUpdateDto := *openapiclient.NewPaymentUpdateDto() // PaymentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdPut(context.Background(), paymentId).TenantId(tenantId).PaymentUpdateDto(paymentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PaymentsServicePaymentsPaymentIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ApiV2PaymentsServicePaymentsPaymentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PaymentsServicePaymentsPaymentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **paymentUpdateDto** | [**PaymentUpdateDto**](PaymentUpdateDto.md) |  | 

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


## ApiV2PaymentsServicePaymentsPost

> EmptyEnvelope ApiV2PaymentsServicePaymentsPost(ctx).TenantId(tenantId).PaymentCreateDto(paymentCreateDto).Execute()



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
	paymentCreateDto := *openapiclient.NewPaymentCreateDto() // PaymentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ApiV2PaymentsServicePaymentsPost(context.Background()).TenantId(tenantId).PaymentCreateDto(paymentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ApiV2PaymentsServicePaymentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2PaymentsServicePaymentsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ApiV2PaymentsServicePaymentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2PaymentsServicePaymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **paymentCreateDto** | [**PaymentCreateDto**](PaymentCreateDto.md) |  | 

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

