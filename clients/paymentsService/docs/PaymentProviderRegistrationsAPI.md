# \PaymentProviderRegistrationsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAsync**](PaymentProviderRegistrationsAPI.md#CreateAsync) | **Post** /api/v2/PaymentsService/PaymentProviderRegistrations | Provisions a provider webhook registration
[**GetAsync**](PaymentProviderRegistrationsAPI.md#GetAsync) | **Get** /api/v2/PaymentsService/PaymentProviderRegistrations | Lists the tenant&#39;s provider registrations
[**GetCountAsync**](PaymentProviderRegistrationsAPI.md#GetCountAsync) | **Get** /api/v2/PaymentsService/PaymentProviderRegistrations/Count | Counts the tenant&#39;s provider registrations
[**RotateKeyAsync**](PaymentProviderRegistrationsAPI.md#RotateKeyAsync) | **Post** /api/v2/PaymentsService/PaymentProviderRegistrations/{registrationId}/RotateKey | Rotates a registration&#39;s webhook key



## CreateAsync

> ProviderWebhookRegistrationCreatedDtoEnvelope CreateAsync(ctx).TenantId(tenantId).CreateProviderWebhookRegistrationRequest(createProviderWebhookRegistrationRequest).Execute()

Provisions a provider webhook registration



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
	createProviderWebhookRegistrationRequest := *openapiclient.NewCreateProviderWebhookRegistrationRequest() // CreateProviderWebhookRegistrationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentProviderRegistrationsAPI.CreateAsync(context.Background()).TenantId(tenantId).CreateProviderWebhookRegistrationRequest(createProviderWebhookRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentProviderRegistrationsAPI.CreateAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAsync`: ProviderWebhookRegistrationCreatedDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentProviderRegistrationsAPI.CreateAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **createProviderWebhookRegistrationRequest** | [**CreateProviderWebhookRegistrationRequest**](CreateProviderWebhookRegistrationRequest.md) |  | 

### Return type

[**ProviderWebhookRegistrationCreatedDtoEnvelope**](ProviderWebhookRegistrationCreatedDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsync

> PaymentProviderRegistrationDtoListEnvelope GetAsync(ctx).TenantId(tenantId).PaymentProviderRegistrationDtoCollectionQueryParameters(paymentProviderRegistrationDtoCollectionQueryParameters).Execute()

Lists the tenant's provider registrations



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
	paymentProviderRegistrationDtoCollectionQueryParameters := *openapiclient.NewPaymentProviderRegistrationDtoCollectionQueryParameters() // PaymentProviderRegistrationDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentProviderRegistrationsAPI.GetAsync(context.Background()).TenantId(tenantId).PaymentProviderRegistrationDtoCollectionQueryParameters(paymentProviderRegistrationDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentProviderRegistrationsAPI.GetAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAsync`: PaymentProviderRegistrationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentProviderRegistrationsAPI.GetAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **paymentProviderRegistrationDtoCollectionQueryParameters** | [**PaymentProviderRegistrationDtoCollectionQueryParameters**](PaymentProviderRegistrationDtoCollectionQueryParameters.md) |  | 

### Return type

[**PaymentProviderRegistrationDtoListEnvelope**](PaymentProviderRegistrationDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountAsync

> Int32Envelope GetCountAsync(ctx).TenantId(tenantId).PaymentProviderRegistrationDtoCollectionQueryParameters(paymentProviderRegistrationDtoCollectionQueryParameters).Execute()

Counts the tenant's provider registrations



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
	paymentProviderRegistrationDtoCollectionQueryParameters := *openapiclient.NewPaymentProviderRegistrationDtoCollectionQueryParameters() // PaymentProviderRegistrationDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentProviderRegistrationsAPI.GetCountAsync(context.Background()).TenantId(tenantId).PaymentProviderRegistrationDtoCollectionQueryParameters(paymentProviderRegistrationDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentProviderRegistrationsAPI.GetCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentProviderRegistrationsAPI.GetCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **paymentProviderRegistrationDtoCollectionQueryParameters** | [**PaymentProviderRegistrationDtoCollectionQueryParameters**](PaymentProviderRegistrationDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateKeyAsync

> ProviderWebhookRegistrationCreatedDtoEnvelope RotateKeyAsync(ctx, registrationId).TenantId(tenantId).Execute()

Rotates a registration's webhook key



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
	registrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentProviderRegistrationsAPI.RotateKeyAsync(context.Background(), registrationId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentProviderRegistrationsAPI.RotateKeyAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateKeyAsync`: ProviderWebhookRegistrationCreatedDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PaymentProviderRegistrationsAPI.RotateKeyAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateKeyAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**ProviderWebhookRegistrationCreatedDtoEnvelope**](ProviderWebhookRegistrationCreatedDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

