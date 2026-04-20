# \ServiceLevelAgreementsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceLevelAgreementAsync**](ServiceLevelAgreementsAPI.md#CreateServiceLevelAgreementAsync) | **Post** /api/v2/ServicesService/ServiceLevelAgreements | Create a service level agreement
[**DeleteServiceLevelAgreementAsync**](ServiceLevelAgreementsAPI.md#DeleteServiceLevelAgreementAsync) | **Delete** /api/v2/ServicesService/ServiceLevelAgreements/{serviceLevelAgreementId} | Delete a service level agreement
[**GetServiceLevelAgreementByIdAsync**](ServiceLevelAgreementsAPI.md#GetServiceLevelAgreementByIdAsync) | **Get** /api/v2/ServicesService/ServiceLevelAgreements/{serviceLevelAgreementId} | Get a service level agreement by ID
[**GetServiceLevelAgreementsAsync**](ServiceLevelAgreementsAPI.md#GetServiceLevelAgreementsAsync) | **Get** /api/v2/ServicesService/ServiceLevelAgreements | Get all service level agreements
[**GetServiceLevelAgreementsCountAsync**](ServiceLevelAgreementsAPI.md#GetServiceLevelAgreementsCountAsync) | **Get** /api/v2/ServicesService/ServiceLevelAgreements/Count | Get service level agreements count
[**UpdateServiceLevelAgreementAsync**](ServiceLevelAgreementsAPI.md#UpdateServiceLevelAgreementAsync) | **Put** /api/v2/ServicesService/ServiceLevelAgreements/{serviceLevelAgreementId} | Update a service level agreement



## CreateServiceLevelAgreementAsync

> Envelope CreateServiceLevelAgreementAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ServiceLevelAgreementCreateDto(serviceLevelAgreementCreateDto).Execute()

Create a service level agreement



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
	serviceLevelAgreementCreateDto := *openapiclient.NewServiceLevelAgreementCreateDto() // ServiceLevelAgreementCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelAgreementsAPI.CreateServiceLevelAgreementAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ServiceLevelAgreementCreateDto(serviceLevelAgreementCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelAgreementsAPI.CreateServiceLevelAgreementAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceLevelAgreementAsync`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelAgreementsAPI.CreateServiceLevelAgreementAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceLevelAgreementAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **serviceLevelAgreementCreateDto** | [**ServiceLevelAgreementCreateDto**](ServiceLevelAgreementCreateDto.md) |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceLevelAgreementAsync

> Envelope DeleteServiceLevelAgreementAsync(ctx, serviceLevelAgreementId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a service level agreement



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
	serviceLevelAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelAgreementsAPI.DeleteServiceLevelAgreementAsync(context.Background(), serviceLevelAgreementId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelAgreementsAPI.DeleteServiceLevelAgreementAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceLevelAgreementAsync`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelAgreementsAPI.DeleteServiceLevelAgreementAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceLevelAgreementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceLevelAgreementAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceLevelAgreementByIdAsync

> ServiceLevelAgreementDtoEnvelope GetServiceLevelAgreementByIdAsync(ctx, serviceLevelAgreementId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get a service level agreement by ID



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
	serviceLevelAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelAgreementsAPI.GetServiceLevelAgreementByIdAsync(context.Background(), serviceLevelAgreementId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelAgreementsAPI.GetServiceLevelAgreementByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceLevelAgreementByIdAsync`: ServiceLevelAgreementDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelAgreementsAPI.GetServiceLevelAgreementByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceLevelAgreementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceLevelAgreementByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ServiceLevelAgreementDtoEnvelope**](ServiceLevelAgreementDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceLevelAgreementsAsync

> ServiceLevelAgreementDtoIReadOnlyListEnvelope GetServiceLevelAgreementsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all service level agreements



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
	resp, r, err := apiClient.ServiceLevelAgreementsAPI.GetServiceLevelAgreementsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelAgreementsAPI.GetServiceLevelAgreementsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceLevelAgreementsAsync`: ServiceLevelAgreementDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelAgreementsAPI.GetServiceLevelAgreementsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceLevelAgreementsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ServiceLevelAgreementDtoIReadOnlyListEnvelope**](ServiceLevelAgreementDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceLevelAgreementsCountAsync

> Int32Envelope GetServiceLevelAgreementsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get service level agreements count



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
	resp, r, err := apiClient.ServiceLevelAgreementsAPI.GetServiceLevelAgreementsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelAgreementsAPI.GetServiceLevelAgreementsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServiceLevelAgreementsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelAgreementsAPI.GetServiceLevelAgreementsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceLevelAgreementsCountAsyncRequest struct via the builder pattern


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


## UpdateServiceLevelAgreementAsync

> Envelope UpdateServiceLevelAgreementAsync(ctx, serviceLevelAgreementId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ServiceLevelAgreementUpdateDto(serviceLevelAgreementUpdateDto).Execute()

Update a service level agreement



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
	serviceLevelAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	serviceLevelAgreementUpdateDto := *openapiclient.NewServiceLevelAgreementUpdateDto() // ServiceLevelAgreementUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceLevelAgreementsAPI.UpdateServiceLevelAgreementAsync(context.Background(), serviceLevelAgreementId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ServiceLevelAgreementUpdateDto(serviceLevelAgreementUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelAgreementsAPI.UpdateServiceLevelAgreementAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServiceLevelAgreementAsync`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `ServiceLevelAgreementsAPI.UpdateServiceLevelAgreementAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceLevelAgreementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceLevelAgreementAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **serviceLevelAgreementUpdateDto** | [**ServiceLevelAgreementUpdateDto**](ServiceLevelAgreementUpdateDto.md) |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

