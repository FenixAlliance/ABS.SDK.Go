# \LicensesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiLicensingLicensesGeneratePost**](LicensesAPI.md#ApiLicensingLicensesGeneratePost) | **Post** /api/Licensing/Licenses/Generate | 
[**ApiLicensingLicensesValidateAttributesGet**](LicensesAPI.md#ApiLicensingLicensesValidateAttributesGet) | **Get** /api/Licensing/Licenses/Validate/Attributes | 
[**ApiLicensingLicensesValidateErrorsGet**](LicensesAPI.md#ApiLicensingLicensesValidateErrorsGet) | **Get** /api/Licensing/Licenses/Validate/Errors | 
[**ApiLicensingLicensesValidateGet**](LicensesAPI.md#ApiLicensingLicensesValidateGet) | **Get** /api/Licensing/Licenses/Validate | 



## ApiLicensingLicensesGeneratePost

> StringEnvelope ApiLicensingLicensesGeneratePost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LicenseKeyRequest(licenseKeyRequest).Execute()



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
	licenseKeyRequest := *openapiclient.NewLicenseKeyRequest() // LicenseKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.ApiLicensingLicensesGeneratePost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LicenseKeyRequest(licenseKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.ApiLicensingLicensesGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiLicensingLicensesGeneratePost`: StringEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.ApiLicensingLicensesGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiLicensingLicensesGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **licenseKeyRequest** | [**LicenseKeyRequest**](LicenseKeyRequest.md) |  | 

### Return type

[**StringEnvelope**](StringEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiLicensingLicensesValidateAttributesGet

> LicenseAttributesListEnvelope ApiLicensingLicensesValidateAttributesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LicenseKey(licenseKey).Execute()



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
	licenseKey := *openapiclient.NewLicenseKey() // LicenseKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.ApiLicensingLicensesValidateAttributesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LicenseKey(licenseKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.ApiLicensingLicensesValidateAttributesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiLicensingLicensesValidateAttributesGet`: LicenseAttributesListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.ApiLicensingLicensesValidateAttributesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiLicensingLicensesValidateAttributesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **licenseKey** | [**LicenseKey**](LicenseKey.md) |  | 

### Return type

[**LicenseAttributesListEnvelope**](LicenseAttributesListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiLicensingLicensesValidateErrorsGet

> LicenseValidationErrorListEnvelope ApiLicensingLicensesValidateErrorsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LicenseKey(licenseKey).Execute()



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
	licenseKey := *openapiclient.NewLicenseKey() // LicenseKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.ApiLicensingLicensesValidateErrorsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LicenseKey(licenseKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.ApiLicensingLicensesValidateErrorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiLicensingLicensesValidateErrorsGet`: LicenseValidationErrorListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.ApiLicensingLicensesValidateErrorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiLicensingLicensesValidateErrorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **licenseKey** | [**LicenseKey**](LicenseKey.md) |  | 

### Return type

[**LicenseValidationErrorListEnvelope**](LicenseValidationErrorListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiLicensingLicensesValidateGet

> BooleanEnvelope ApiLicensingLicensesValidateGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LicenseKey(licenseKey).Execute()



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
	licenseKey := *openapiclient.NewLicenseKey() // LicenseKey |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensesAPI.ApiLicensingLicensesValidateGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LicenseKey(licenseKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensesAPI.ApiLicensingLicensesValidateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiLicensingLicensesValidateGet`: BooleanEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LicensesAPI.ApiLicensingLicensesValidateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiLicensingLicensesValidateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **licenseKey** | [**LicenseKey**](LicenseKey.md) |  | 

### Return type

[**BooleanEnvelope**](BooleanEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

