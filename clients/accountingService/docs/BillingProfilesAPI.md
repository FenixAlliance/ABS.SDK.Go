# \BillingProfilesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingProfileAsync**](BillingProfilesAPI.md#CreateBillingProfileAsync) | **Post** /api/v2/AccountingService/BillingProfiles | Creates a new billing profile
[**DeleteBillingProfileAsync**](BillingProfilesAPI.md#DeleteBillingProfileAsync) | **Delete** /api/v2/AccountingService/BillingProfiles/{billingProfileId} | Deletes a billing profile
[**GetBillingProfileByIdAsync**](BillingProfilesAPI.md#GetBillingProfileByIdAsync) | **Get** /api/v2/AccountingService/BillingProfiles/{billingProfileId} | Gets a billing profile by id
[**GetBillingProfilesAsync**](BillingProfilesAPI.md#GetBillingProfilesAsync) | **Get** /api/v2/AccountingService/BillingProfiles | Gets all billing profiles
[**GetBillingProfilesCountAsync**](BillingProfilesAPI.md#GetBillingProfilesCountAsync) | **Get** /api/v2/AccountingService/BillingProfiles/Count | Gets the count of billing profiles
[**UpdateBillingProfileAsync**](BillingProfilesAPI.md#UpdateBillingProfileAsync) | **Put** /api/v2/AccountingService/BillingProfiles/{billingProfileId} | Updates an existing billing profile



## CreateBillingProfileAsync

> EmptyEnvelope CreateBillingProfileAsync(ctx).TenantId(tenantId).BillingProfileCreateDto(billingProfileCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Creates a new billing profile



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
	billingProfileCreateDto := *openapiclient.NewBillingProfileCreateDto("TaxId_example", "Phone_example", "Email_example", "Address_example", "PostalCode_example", "BusinessName_example", "CommercialName_example", "CountryId_example", "StateId_example", "CityId_example", "FiscalIdentificationTypeId_example", "FiscalAuthorityId_example", "FiscalRegimeId_example") // BillingProfileCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingProfilesAPI.CreateBillingProfileAsync(context.Background()).TenantId(tenantId).BillingProfileCreateDto(billingProfileCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingProfilesAPI.CreateBillingProfileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBillingProfileAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillingProfilesAPI.CreateBillingProfileAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingProfileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **billingProfileCreateDto** | [**BillingProfileCreateDto**](BillingProfileCreateDto.md) |  | 
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


## DeleteBillingProfileAsync

> EmptyEnvelope DeleteBillingProfileAsync(ctx, billingProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a billing profile



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
	billingProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingProfilesAPI.DeleteBillingProfileAsync(context.Background(), billingProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingProfilesAPI.DeleteBillingProfileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBillingProfileAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillingProfilesAPI.DeleteBillingProfileAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBillingProfileAsyncRequest struct via the builder pattern


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


## GetBillingProfileByIdAsync

> BillingProfileDtoEnvelope GetBillingProfileByIdAsync(ctx, billingProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets a billing profile by id



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
	billingProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingProfilesAPI.GetBillingProfileByIdAsync(context.Background(), billingProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingProfilesAPI.GetBillingProfileByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingProfileByIdAsync`: BillingProfileDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillingProfilesAPI.GetBillingProfileByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingProfileByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BillingProfileDtoEnvelope**](BillingProfileDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingProfilesAsync

> BillingProfileDtoIReadOnlyListEnvelope GetBillingProfilesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets all billing profiles



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
	resp, r, err := apiClient.BillingProfilesAPI.GetBillingProfilesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingProfilesAPI.GetBillingProfilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingProfilesAsync`: BillingProfileDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillingProfilesAPI.GetBillingProfilesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingProfilesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BillingProfileDtoIReadOnlyListEnvelope**](BillingProfileDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingProfilesCountAsync

> Int32Envelope GetBillingProfilesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the count of billing profiles



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
	resp, r, err := apiClient.BillingProfilesAPI.GetBillingProfilesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingProfilesAPI.GetBillingProfilesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillingProfilesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BillingProfilesAPI.GetBillingProfilesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingProfilesCountAsyncRequest struct via the builder pattern


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


## UpdateBillingProfileAsync

> EmptyEnvelope UpdateBillingProfileAsync(ctx, billingProfileId).TenantId(tenantId).BillingProfileUpdateDto(billingProfileUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Updates an existing billing profile



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
	billingProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	billingProfileUpdateDto := *openapiclient.NewBillingProfileUpdateDto() // BillingProfileUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingProfilesAPI.UpdateBillingProfileAsync(context.Background(), billingProfileId).TenantId(tenantId).BillingProfileUpdateDto(billingProfileUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingProfilesAPI.UpdateBillingProfileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBillingProfileAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BillingProfilesAPI.UpdateBillingProfileAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillingProfileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **billingProfileUpdateDto** | [**BillingProfileUpdateDto**](BillingProfileUpdateDto.md) |  | 
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

