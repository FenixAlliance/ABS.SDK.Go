# \TeamProjectEnrollmentsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenantTeamProjectEnrollment**](TeamProjectEnrollmentsAPI.md#CreateTenantTeamProjectEnrollment) | **Post** /api/v2/TenantsService/TeamProjectEnrollments | Create a new tenant team project enrollment
[**DeleteTenantTeamProjectEnrollment**](TeamProjectEnrollmentsAPI.md#DeleteTenantTeamProjectEnrollment) | **Delete** /api/v2/TenantsService/TeamProjectEnrollments/{tenantTeamProjectEnrollmentId} | Delete a tenant team project enrollment
[**GetTenantTeamProjectEnrollmentById**](TeamProjectEnrollmentsAPI.md#GetTenantTeamProjectEnrollmentById) | **Get** /api/v2/TenantsService/TeamProjectEnrollments/{tenantTeamProjectEnrollmentId} | Retrieve a single tenant team project enrollment by its ID
[**GetTenantTeamProjectEnrollments**](TeamProjectEnrollmentsAPI.md#GetTenantTeamProjectEnrollments) | **Get** /api/v2/TenantsService/TeamProjectEnrollments | Retrieve a list of tenant team project enrollments
[**GetTenantTeamProjectEnrollmentsCount**](TeamProjectEnrollmentsAPI.md#GetTenantTeamProjectEnrollmentsCount) | **Get** /api/v2/TenantsService/TeamProjectEnrollments/Count | Get the count of tenant team project enrollments
[**UpdateTenantTeamProjectEnrollment**](TeamProjectEnrollmentsAPI.md#UpdateTenantTeamProjectEnrollment) | **Put** /api/v2/TenantsService/TeamProjectEnrollments/{tenantTeamProjectEnrollmentId} | Update a tenant team project enrollment



## CreateTenantTeamProjectEnrollment

> EmptyEnvelope CreateTenantTeamProjectEnrollment(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamProjectEnrollmentCreateDto(tenantTeamProjectEnrollmentCreateDto).Execute()

Create a new tenant team project enrollment



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
	tenantTeamProjectEnrollmentCreateDto := *openapiclient.NewTenantTeamProjectEnrollmentCreateDto("BusinessTeamID_example", "ProjectID_example") // TenantTeamProjectEnrollmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamProjectEnrollmentsAPI.CreateTenantTeamProjectEnrollment(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamProjectEnrollmentCreateDto(tenantTeamProjectEnrollmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamProjectEnrollmentsAPI.CreateTenantTeamProjectEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantTeamProjectEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamProjectEnrollmentsAPI.CreateTenantTeamProjectEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantTeamProjectEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantTeamProjectEnrollmentCreateDto** | [**TenantTeamProjectEnrollmentCreateDto**](TenantTeamProjectEnrollmentCreateDto.md) |  | 

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


## DeleteTenantTeamProjectEnrollment

> EmptyEnvelope DeleteTenantTeamProjectEnrollment(ctx, tenantTeamProjectEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a tenant team project enrollment



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
	tenantTeamProjectEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamProjectEnrollmentsAPI.DeleteTenantTeamProjectEnrollment(context.Background(), tenantTeamProjectEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamProjectEnrollmentsAPI.DeleteTenantTeamProjectEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantTeamProjectEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamProjectEnrollmentsAPI.DeleteTenantTeamProjectEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTeamProjectEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantTeamProjectEnrollmentRequest struct via the builder pattern


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


## GetTenantTeamProjectEnrollmentById

> TenantTeamProjectEnrollmentDtoEnvelope GetTenantTeamProjectEnrollmentById(ctx, tenantTeamProjectEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single tenant team project enrollment by its ID



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
	tenantTeamProjectEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollmentById(context.Background(), tenantTeamProjectEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTeamProjectEnrollmentById`: TenantTeamProjectEnrollmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTeamProjectEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTeamProjectEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantTeamProjectEnrollmentDtoEnvelope**](TenantTeamProjectEnrollmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantTeamProjectEnrollments

> TenantTeamProjectEnrollmentDtoListEnvelope GetTenantTeamProjectEnrollments(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of tenant team project enrollments



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
	resp, r, err := apiClient.TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollments(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTeamProjectEnrollments`: TenantTeamProjectEnrollmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTeamProjectEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantTeamProjectEnrollmentDtoListEnvelope**](TenantTeamProjectEnrollmentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantTeamProjectEnrollmentsCount

> Int32Envelope GetTenantTeamProjectEnrollmentsCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of tenant team project enrollments



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
	resp, r, err := apiClient.TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollmentsCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollmentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTeamProjectEnrollmentsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TeamProjectEnrollmentsAPI.GetTenantTeamProjectEnrollmentsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTeamProjectEnrollmentsCountRequest struct via the builder pattern


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


## UpdateTenantTeamProjectEnrollment

> EmptyEnvelope UpdateTenantTeamProjectEnrollment(ctx, tenantTeamProjectEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamProjectEnrollmentUpdateDto(tenantTeamProjectEnrollmentUpdateDto).Execute()

Update a tenant team project enrollment



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
	tenantTeamProjectEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	tenantTeamProjectEnrollmentUpdateDto := *openapiclient.NewTenantTeamProjectEnrollmentUpdateDto() // TenantTeamProjectEnrollmentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamProjectEnrollmentsAPI.UpdateTenantTeamProjectEnrollment(context.Background(), tenantTeamProjectEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamProjectEnrollmentUpdateDto(tenantTeamProjectEnrollmentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamProjectEnrollmentsAPI.UpdateTenantTeamProjectEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenantTeamProjectEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamProjectEnrollmentsAPI.UpdateTenantTeamProjectEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTeamProjectEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantTeamProjectEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantTeamProjectEnrollmentUpdateDto** | [**TenantTeamProjectEnrollmentUpdateDto**](TenantTeamProjectEnrollmentUpdateDto.md) |  | 

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

