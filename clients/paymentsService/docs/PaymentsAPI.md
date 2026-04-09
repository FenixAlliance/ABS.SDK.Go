# \PaymentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentAsync**](PaymentsAPI.md#CreatePaymentAsync) | **Post** /api/v2/PaymentsService/Payments | Creates a new payment
[**DeletePaymentAsync**](PaymentsAPI.md#DeletePaymentAsync) | **Delete** /api/v2/PaymentsService/Payments/{paymentId} | Deletes a payment
[**GetPaymentAsync**](PaymentsAPI.md#GetPaymentAsync) | **Get** /api/v2/PaymentsService/Payments/{paymentId}/Details | Gets a payment by ID (deprecated)
[**GetPaymentAsyncV2**](PaymentsAPI.md#GetPaymentAsyncV2) | **Get** /api/v2/PaymentsService/Payments/{paymentId} | Gets a payment by ID
[**GetPaymentsAsync**](PaymentsAPI.md#GetPaymentsAsync) | **Get** /api/v2/PaymentsService/Payments | Retrieves all payments
[**UpdatePaymentAsync**](PaymentsAPI.md#UpdatePaymentAsync) | **Put** /api/v2/PaymentsService/Payments/{paymentId} | Updates a payment



## CreatePaymentAsync

> EmptyEnvelope CreatePaymentAsync(ctx).TenantId(tenantId).PaymentCreateDto(paymentCreateDto).Execute()

Creates a new payment



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
	resp, r, err := apiClient.PaymentsAPI.CreatePaymentAsync(context.Background()).TenantId(tenantId).PaymentCreateDto(paymentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.CreatePaymentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.CreatePaymentAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **paymentCreateDto** | [**PaymentCreateDto**](PaymentCreateDto.md) |  | 

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


## DeletePaymentAsync

> EmptyEnvelope DeletePaymentAsync(ctx, paymentId).TenantId(tenantId).Execute()

Deletes a payment



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
	resp, r, err := apiClient.PaymentsAPI.DeletePaymentAsync(context.Background(), paymentId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.DeletePaymentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePaymentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.DeletePaymentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## GetPaymentAsync

> PaymentDtoListEnvelope GetPaymentAsync(ctx, paymentId).Execute()

Gets a payment by ID (deprecated)



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
	resp, r, err := apiClient.PaymentsAPI.GetPaymentAsync(context.Background(), paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPaymentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentAsync`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPaymentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentAsyncV2

> PaymentDtoListEnvelope GetPaymentAsyncV2(ctx, paymentId).Execute()

Gets a payment by ID



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
	resp, r, err := apiClient.PaymentsAPI.GetPaymentAsyncV2(context.Background(), paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPaymentAsyncV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentAsyncV2`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPaymentAsyncV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentAsyncV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentsAsync

> PaymentDtoListEnvelope GetPaymentsAsync(ctx).TenantId(tenantId).Execute()

Retrieves all payments



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
	resp, r, err := apiClient.PaymentsAPI.GetPaymentsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPaymentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentsAsync`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPaymentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentAsync

> EmptyEnvelope UpdatePaymentAsync(ctx, paymentId).TenantId(tenantId).PaymentUpdateDto(paymentUpdateDto).Execute()

Updates a payment



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
	resp, r, err := apiClient.PaymentsAPI.UpdatePaymentAsync(context.Background(), paymentId).TenantId(tenantId).PaymentUpdateDto(paymentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.UpdatePaymentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePaymentAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.UpdatePaymentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **paymentUpdateDto** | [**PaymentUpdateDto**](PaymentUpdateDto.md) |  | 

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

