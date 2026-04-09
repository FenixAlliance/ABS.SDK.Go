# \EmployeeEnrollmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenantEmployeeEnrollment**](EmployeeEnrollmentsAPI.md#CreateTenantEmployeeEnrollment) | **Post** /api/v2/TenantsService/EmployeeEnrollments | Create a new tenant employee enrollment
[**DeleteTenantEmployeeEnrollment**](EmployeeEnrollmentsAPI.md#DeleteTenantEmployeeEnrollment) | **Delete** /api/v2/TenantsService/EmployeeEnrollments/{tenantEmployeeEnrollmentId} | Delete a tenant employee enrollment
[**GetTenantEmployeeEnrollmentById**](EmployeeEnrollmentsAPI.md#GetTenantEmployeeEnrollmentById) | **Get** /api/v2/TenantsService/EmployeeEnrollments/{tenantEmployeeEnrollmentId} | Retrieve a single tenant employee enrollment by its ID
[**GetTenantEmployeeEnrollments**](EmployeeEnrollmentsAPI.md#GetTenantEmployeeEnrollments) | **Get** /api/v2/TenantsService/EmployeeEnrollments | Retrieve a list of tenant employee enrollments
[**GetTenantEmployeeEnrollmentsCount**](EmployeeEnrollmentsAPI.md#GetTenantEmployeeEnrollmentsCount) | **Get** /api/v2/TenantsService/EmployeeEnrollments/Count | Get the count of tenant employee enrollments
[**UpdateTenantEmployeeEnrollment**](EmployeeEnrollmentsAPI.md#UpdateTenantEmployeeEnrollment) | **Put** /api/v2/TenantsService/EmployeeEnrollments/{tenantEmployeeEnrollmentId} | Update a tenant employee enrollment



## CreateTenantEmployeeEnrollment

> EmptyEnvelope CreateTenantEmployeeEnrollment(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamEmployeeEnrollmentCreateDto(tenantTeamEmployeeEnrollmentCreateDto).Execute()

Create a new tenant employee enrollment



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
	tenantTeamEmployeeEnrollmentCreateDto := *openapiclient.NewTenantTeamEmployeeEnrollmentCreateDto("BusinessID_example", "BusinessProfileRecordID_example", "BusinessTeamID_example", "EmployeeProfileID_example") // TenantTeamEmployeeEnrollmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeEnrollmentsAPI.CreateTenantEmployeeEnrollment(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamEmployeeEnrollmentCreateDto(tenantTeamEmployeeEnrollmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeEnrollmentsAPI.CreateTenantEmployeeEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantEmployeeEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeEnrollmentsAPI.CreateTenantEmployeeEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantEmployeeEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantTeamEmployeeEnrollmentCreateDto** | [**TenantTeamEmployeeEnrollmentCreateDto**](TenantTeamEmployeeEnrollmentCreateDto.md) |  | 

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


## DeleteTenantEmployeeEnrollment

> EmptyEnvelope DeleteTenantEmployeeEnrollment(ctx, tenantEmployeeEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a tenant employee enrollment



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
	tenantEmployeeEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeEnrollmentsAPI.DeleteTenantEmployeeEnrollment(context.Background(), tenantEmployeeEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeEnrollmentsAPI.DeleteTenantEmployeeEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantEmployeeEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeEnrollmentsAPI.DeleteTenantEmployeeEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantEmployeeEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantEmployeeEnrollmentRequest struct via the builder pattern


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


## GetTenantEmployeeEnrollmentById

> TenantTeamEmployeeEnrollmentDtoEnvelope GetTenantEmployeeEnrollmentById(ctx, tenantEmployeeEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single tenant employee enrollment by its ID



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
	tenantEmployeeEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollmentById(context.Background(), tenantEmployeeEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantEmployeeEnrollmentById`: TenantTeamEmployeeEnrollmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantEmployeeEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantEmployeeEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantTeamEmployeeEnrollmentDtoEnvelope**](TenantTeamEmployeeEnrollmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantEmployeeEnrollments

> TenantTeamEmployeeEnrollmentDtoListEnvelope GetTenantEmployeeEnrollments(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of tenant employee enrollments



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
	resp, r, err := apiClient.EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollments(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantEmployeeEnrollments`: TenantTeamEmployeeEnrollmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantEmployeeEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantTeamEmployeeEnrollmentDtoListEnvelope**](TenantTeamEmployeeEnrollmentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantEmployeeEnrollmentsCount

> Int32Envelope GetTenantEmployeeEnrollmentsCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of tenant employee enrollments



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
	resp, r, err := apiClient.EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollmentsCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollmentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantEmployeeEnrollmentsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeEnrollmentsAPI.GetTenantEmployeeEnrollmentsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantEmployeeEnrollmentsCountRequest struct via the builder pattern


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


## UpdateTenantEmployeeEnrollment

> EmptyEnvelope UpdateTenantEmployeeEnrollment(ctx, tenantEmployeeEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamEmployeeEnrollmentUpdateDto(tenantTeamEmployeeEnrollmentUpdateDto).Execute()

Update a tenant employee enrollment



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
	tenantEmployeeEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	tenantTeamEmployeeEnrollmentUpdateDto := *openapiclient.NewTenantTeamEmployeeEnrollmentUpdateDto() // TenantTeamEmployeeEnrollmentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmployeeEnrollmentsAPI.UpdateTenantEmployeeEnrollment(context.Background(), tenantEmployeeEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamEmployeeEnrollmentUpdateDto(tenantTeamEmployeeEnrollmentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmployeeEnrollmentsAPI.UpdateTenantEmployeeEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenantEmployeeEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmployeeEnrollmentsAPI.UpdateTenantEmployeeEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantEmployeeEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantEmployeeEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantTeamEmployeeEnrollmentUpdateDto** | [**TenantTeamEmployeeEnrollmentUpdateDto**](TenantTeamEmployeeEnrollmentUpdateDto.md) |  | 

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

