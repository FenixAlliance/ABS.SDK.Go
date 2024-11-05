# \RolesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SecurityServiceRolesGet**](RolesAPI.md#ApiV2SecurityServiceRolesGet) | **Get** /api/v2/SecurityService/Roles | 
[**ApiV2SecurityServiceRolesPost**](RolesAPI.md#ApiV2SecurityServiceRolesPost) | **Post** /api/v2/SecurityService/Roles | 
[**ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDelete**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDelete) | **Delete** /api/v2/SecurityService/Roles/{securityRoleId}/Applications/{applicationId} | 
[**ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPost**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPost) | **Post** /api/v2/SecurityService/Roles/{securityRoleId}/Applications/{applicationId} | 
[**ApiV2SecurityServiceRolesSecurityRoleIdDelete**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdDelete) | **Delete** /api/v2/SecurityService/Roles/{securityRoleId} | 
[**ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDelete**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDelete) | **Delete** /api/v2/SecurityService/Roles/{securityRoleId}/Enrollments/{enrollmentId} | 
[**ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPost**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPost) | **Post** /api/v2/SecurityService/Roles/{securityRoleId}/Enrollments/{enrollmentId} | 
[**ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGet**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGet) | **Get** /api/v2/SecurityService/Roles/{securityRoleId}/Enrollments | 
[**ApiV2SecurityServiceRolesSecurityRoleIdGet**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdGet) | **Get** /api/v2/SecurityService/Roles/{securityRoleId} | 
[**ApiV2SecurityServiceRolesSecurityRoleIdPermissionsGet**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdPermissionsGet) | **Get** /api/v2/SecurityService/Roles/{securityRoleId}/Permissions | 
[**ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDelete**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDelete) | **Delete** /api/v2/SecurityService/Roles/{securityRoleId}/Permissions/{securityPermissionId} | 
[**ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPost**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPost) | **Post** /api/v2/SecurityService/Roles/{securityRoleId}/Permissions/{securityPermissionId} | 
[**ApiV2SecurityServiceRolesSecurityRoleIdPut**](RolesAPI.md#ApiV2SecurityServiceRolesSecurityRoleIdPut) | **Put** /api/v2/SecurityService/Roles/{securityRoleId} | 



## ApiV2SecurityServiceRolesGet

> SecurityRoleDtoListEnvelope ApiV2SecurityServiceRolesGet(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesGet(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesGet`: SecurityRoleDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesGetRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesPost

> EmptyEnvelope ApiV2SecurityServiceRolesPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SecurityRoleCreateDto(securityRoleCreateDto).Execute()



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
	securityRoleCreateDto := *openapiclient.NewSecurityRoleCreateDto("Name_example", "TenantId_example") // SecurityRoleCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SecurityRoleCreateDto(securityRoleCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **securityRoleCreateDto** | [**SecurityRoleCreateDto**](SecurityRoleCreateDto.md) |  | 

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


## ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDelete

> EmptyEnvelope ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDelete(ctx, securityRoleId, applicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDelete(context.Background(), securityRoleId, applicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdDeleteRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPost

> EmptyEnvelope ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPost(ctx, securityRoleId, applicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPost(context.Background(), securityRoleId, applicationId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdApplicationsApplicationIdPostRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdDelete

> EmptyEnvelope ApiV2SecurityServiceRolesSecurityRoleIdDelete(ctx, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdDelete(context.Background(), securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdDeleteRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDelete

> EmptyEnvelope ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDelete(ctx, securityRoleId, enrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	enrollmentId := "enrollmentId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDelete(context.Background(), securityRoleId, enrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 
**enrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdDeleteRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPost

> EmptyEnvelope ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPost(ctx, securityRoleId, enrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	enrollmentId := "enrollmentId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPost(context.Background(), securityRoleId, enrollmentId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 
**enrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsEnrollmentIdPostRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGet

> TenantEnrolmentDtoListEnvelope ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGet(ctx, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGet(context.Background(), securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGet`: TenantEnrolmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdEnrollmentsGetRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdGet

> SecurityRoleDtoListEnvelope ApiV2SecurityServiceRolesSecurityRoleIdGet(ctx, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdGet(context.Background(), securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdGet`: SecurityRoleDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdGetRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdPermissionsGet

> SecurityPermissionDtoListEnvelope ApiV2SecurityServiceRolesSecurityRoleIdPermissionsGet(ctx, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsGet(context.Background(), securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdPermissionsGet`: SecurityPermissionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SecurityPermissionDtoListEnvelope**](SecurityPermissionDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDelete

> EmptyEnvelope ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDelete(ctx, securityRoleId, securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	securityPermissionId := "securityPermissionId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDelete(context.Background(), securityRoleId, securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 
**securityPermissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdDeleteRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPost

> EmptyEnvelope ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPost(ctx, securityRoleId, securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	securityPermissionId := "securityPermissionId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPost(context.Background(), securityRoleId, securityPermissionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 
**securityPermissionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdPermissionsSecurityPermissionIdPostRequest struct via the builder pattern


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


## ApiV2SecurityServiceRolesSecurityRoleIdPut

> EmptyEnvelope ApiV2SecurityServiceRolesSecurityRoleIdPut(ctx, securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SecurityRoleUpdateDto(securityRoleUpdateDto).Execute()



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
	securityRoleId := "securityRoleId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	securityRoleUpdateDto := *openapiclient.NewSecurityRoleUpdateDto("Name_example") // SecurityRoleUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPut(context.Background(), securityRoleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SecurityRoleUpdateDto(securityRoleUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SecurityServiceRolesSecurityRoleIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.ApiV2SecurityServiceRolesSecurityRoleIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityRoleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SecurityServiceRolesSecurityRoleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **securityRoleUpdateDto** | [**SecurityRoleUpdateDto**](SecurityRoleUpdateDto.md) |  | 

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

