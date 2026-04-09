# \TeamContactEnrollmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenantTeamContactEnrollment**](TeamContactEnrollmentsAPI.md#CreateTenantTeamContactEnrollment) | **Post** /api/v2/TenantsService/TeamContactEnrollments | Create a new tenant team contact enrollment
[**DeleteTenantTeamContactEnrollment**](TeamContactEnrollmentsAPI.md#DeleteTenantTeamContactEnrollment) | **Delete** /api/v2/TenantsService/TeamContactEnrollments/{tenantTeamContactEnrollmentId} | Delete a tenant team contact enrollment
[**GetTenantTeamContactEnrollmentById**](TeamContactEnrollmentsAPI.md#GetTenantTeamContactEnrollmentById) | **Get** /api/v2/TenantsService/TeamContactEnrollments/{tenantTeamContactEnrollmentId} | Retrieve a single tenant team contact enrollment by its ID
[**GetTenantTeamContactEnrollments**](TeamContactEnrollmentsAPI.md#GetTenantTeamContactEnrollments) | **Get** /api/v2/TenantsService/TeamContactEnrollments | Retrieve a list of tenant team contact enrollments
[**GetTenantTeamContactEnrollmentsCount**](TeamContactEnrollmentsAPI.md#GetTenantTeamContactEnrollmentsCount) | **Get** /api/v2/TenantsService/TeamContactEnrollments/Count | Get the count of tenant team contact enrollments
[**UpdateTenantTeamContactEnrollment**](TeamContactEnrollmentsAPI.md#UpdateTenantTeamContactEnrollment) | **Put** /api/v2/TenantsService/TeamContactEnrollments/{tenantTeamContactEnrollmentId} | Update a tenant team contact enrollment



## CreateTenantTeamContactEnrollment

> EmptyEnvelope CreateTenantTeamContactEnrollment(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamContactEnrollmentCreateDto(tenantTeamContactEnrollmentCreateDto).Execute()

Create a new tenant team contact enrollment



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
	tenantTeamContactEnrollmentCreateDto := *openapiclient.NewTenantTeamContactEnrollmentCreateDto("BusinessID_example", "BusinessProfileRecordID_example", "BusinessTeamID_example", "ContactID_example") // TenantTeamContactEnrollmentCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamContactEnrollmentsAPI.CreateTenantTeamContactEnrollment(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamContactEnrollmentCreateDto(tenantTeamContactEnrollmentCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamContactEnrollmentsAPI.CreateTenantTeamContactEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantTeamContactEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamContactEnrollmentsAPI.CreateTenantTeamContactEnrollment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantTeamContactEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantTeamContactEnrollmentCreateDto** | [**TenantTeamContactEnrollmentCreateDto**](TenantTeamContactEnrollmentCreateDto.md) |  | 

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


## DeleteTenantTeamContactEnrollment

> EmptyEnvelope DeleteTenantTeamContactEnrollment(ctx, tenantTeamContactEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a tenant team contact enrollment



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
	tenantTeamContactEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamContactEnrollmentsAPI.DeleteTenantTeamContactEnrollment(context.Background(), tenantTeamContactEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamContactEnrollmentsAPI.DeleteTenantTeamContactEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantTeamContactEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamContactEnrollmentsAPI.DeleteTenantTeamContactEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTeamContactEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantTeamContactEnrollmentRequest struct via the builder pattern


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


## GetTenantTeamContactEnrollmentById

> TenantTeamContactEnrollmentDtoEnvelope GetTenantTeamContactEnrollmentById(ctx, tenantTeamContactEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single tenant team contact enrollment by its ID



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
	tenantTeamContactEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollmentById(context.Background(), tenantTeamContactEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTeamContactEnrollmentById`: TenantTeamContactEnrollmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollmentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTeamContactEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTeamContactEnrollmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantTeamContactEnrollmentDtoEnvelope**](TenantTeamContactEnrollmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantTeamContactEnrollments

> TenantTeamContactEnrollmentDtoListEnvelope GetTenantTeamContactEnrollments(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of tenant team contact enrollments



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
	resp, r, err := apiClient.TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollments(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTeamContactEnrollments`: TenantTeamContactEnrollmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTeamContactEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantTeamContactEnrollmentDtoListEnvelope**](TenantTeamContactEnrollmentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantTeamContactEnrollmentsCount

> Int32Envelope GetTenantTeamContactEnrollmentsCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of tenant team contact enrollments



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
	resp, r, err := apiClient.TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollmentsCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollmentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantTeamContactEnrollmentsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `TeamContactEnrollmentsAPI.GetTenantTeamContactEnrollmentsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantTeamContactEnrollmentsCountRequest struct via the builder pattern


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


## UpdateTenantTeamContactEnrollment

> EmptyEnvelope UpdateTenantTeamContactEnrollment(ctx, tenantTeamContactEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamContactEnrollmentUpdateDto(tenantTeamContactEnrollmentUpdateDto).Execute()

Update a tenant team contact enrollment



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
	tenantTeamContactEnrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	tenantTeamContactEnrollmentUpdateDto := *openapiclient.NewTenantTeamContactEnrollmentUpdateDto() // TenantTeamContactEnrollmentUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamContactEnrollmentsAPI.UpdateTenantTeamContactEnrollment(context.Background(), tenantTeamContactEnrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).TenantTeamContactEnrollmentUpdateDto(tenantTeamContactEnrollmentUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamContactEnrollmentsAPI.UpdateTenantTeamContactEnrollment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenantTeamContactEnrollment`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TeamContactEnrollmentsAPI.UpdateTenantTeamContactEnrollment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantTeamContactEnrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantTeamContactEnrollmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **tenantTeamContactEnrollmentUpdateDto** | [**TenantTeamContactEnrollmentUpdateDto**](TenantTeamContactEnrollmentUpdateDto.md) |  | 

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

