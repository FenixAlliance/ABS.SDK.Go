# \PermissionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SecurityServicePermissionsGet**](PermissionsAPI.md#ApiV2SecurityServicePermissionsGet) | **Get** /api/v2/SecurityService/Permissions | 
[**ApiV2SecurityServicePermissionsPost**](PermissionsAPI.md#ApiV2SecurityServicePermissionsPost) | **Post** /api/v2/SecurityService/Permissions | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDelete**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDelete) | **Delete** /api/v2/SecurityService/Permissions/{securityPermissionId}/Applications/{applicationId} | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPost**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPost) | **Post** /api/v2/SecurityService/Permissions/{securityPermissionId}/Applications/{applicationId} | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdDelete**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdDelete) | **Delete** /api/v2/SecurityService/Permissions/{securityPermissionId} | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDelete**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDelete) | **Delete** /api/v2/SecurityService/Permissions/{securityPermissionId}/Enrollments/{enrollmentId} | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPost**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPost) | **Post** /api/v2/SecurityService/Permissions/{securityPermissionId}/Enrollments/{enrollmentId} | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGet**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGet) | **Get** /api/v2/SecurityService/Permissions/{securityPermissionId}/Enrollments | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdGet**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdGet) | **Get** /api/v2/SecurityService/Permissions/{securityPermissionId} | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdPut**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdPut) | **Put** /api/v2/SecurityService/Permissions/{securityPermissionId} | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDelete**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDelete) | **Delete** /api/v2/SecurityService/Permissions/{securityPermissionId}/Roles/{securityRoleId} | 
[**ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPost**](PermissionsAPI.md#ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPost) | **Post** /api/v2/SecurityService/Permissions/{securityPermissionId}/Roles/{securityRoleId} | 



## ApiV2SecurityServicePermissionsGet

> SecurityRoleDtoListEnvelope ApiV2SecurityServicePermissionsGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsGet`: SecurityRoleDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SecurityRoleDtoListEnvelope**](SecurityRoleDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SecurityServicePermissionsPost

> EmptyEnvelope ApiV2SecurityServicePermissionsPost(ctx).TenantId(tenantId).SecurityPermissionCreateDto(securityPermissionCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionCreateDto := *openapiclient.NewSecurityPermissionCreateDto("Name_example", "TenantId_example") // SecurityPermissionCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsPost(context.Background()).TenantId(tenantId).SecurityPermissionCreateDto(securityPermissionCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **securityPermissionCreateDto** | [**SecurityPermissionCreateDto**](SecurityPermissionCreateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDelete

> EmptyEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDelete(ctx, securityPermissionId, applicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	applicationId := "applicationId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDelete(context.Background(), securityPermissionId, applicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPost

> EmptyEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPost(ctx, securityPermissionId, applicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	applicationId := "applicationId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPost(context.Background(), securityPermissionId, applicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdApplicationsApplicationIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SecurityServicePermissionsSecurityPermissionIdDelete

> EmptyEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdDelete(ctx, securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdDelete(context.Background(), securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDelete

> EmptyEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDelete(ctx, securityPermissionId, enrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	enrollmentId := "enrollmentId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDelete(context.Background(), securityPermissionId, enrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 
**enrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPost

> EmptyEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPost(ctx, securityPermissionId, enrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	enrollmentId := "enrollmentId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPost(context.Background(), securityPermissionId, enrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 
**enrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsEnrollmentIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGet

> TenantEnrolmentDtoListEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGet(ctx, securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGet(context.Background(), securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGet`: TenantEnrolmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdEnrollmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantEnrolmentDtoListEnvelope**](TenantEnrolmentDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SecurityServicePermissionsSecurityPermissionIdGet

> SecurityRoleDtoListEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdGet(ctx, securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdGet(context.Background(), securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdGet`: SecurityRoleDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SecurityRoleDtoListEnvelope**](SecurityRoleDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SecurityServicePermissionsSecurityPermissionIdPut

> EmptyEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdPut(ctx, securityPermissionId).TenantId(tenantId).SecurityPermissionUpdateDto(securityPermissionUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	securityPermissionUpdateDto := *openapiclient.NewSecurityPermissionUpdateDto("Name_example") // SecurityPermissionUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdPut(context.Background(), securityPermissionId).TenantId(tenantId).SecurityPermissionUpdateDto(securityPermissionUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **securityPermissionUpdateDto** | [**SecurityPermissionUpdateDto**](SecurityPermissionUpdateDto.md) |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDelete

> EmptyEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDelete(ctx, securityPermissionId, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	securityRoleId := "securityRoleId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDelete(context.Background(), securityPermissionId, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 
**securityRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPost

> EmptyEnvelope ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPost(ctx, securityPermissionId, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityPermissionId := "securityPermissionId_example" // string | 
	securityRoleId := "securityRoleId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPost(context.Background(), securityPermissionId, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.ApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityPermissionId** | **string** |  | 
**securityRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServicePermissionsSecurityPermissionIdRolesSecurityRoleIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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

