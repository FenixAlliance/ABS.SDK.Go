# \ApplicationPrincipalsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableGlobalApplicationPrincipal**](ApplicationPrincipalsAPI.md#DisableGlobalApplicationPrincipal) | **Post** /api/v2/SystemService/ApplicationPrincipals/{principalId}/Disable | Disable an application principal (global)
[**EnableGlobalApplicationPrincipal**](ApplicationPrincipalsAPI.md#EnableGlobalApplicationPrincipal) | **Post** /api/v2/SystemService/ApplicationPrincipals/{principalId}/Enable | Enable an application principal (global)
[**GetGlobalApplicationPrincipal**](ApplicationPrincipalsAPI.md#GetGlobalApplicationPrincipal) | **Get** /api/v2/SystemService/ApplicationPrincipals/{principalId} | Get one application principal (any tenant)
[**GetGlobalApplicationPrincipals**](ApplicationPrincipalsAPI.md#GetGlobalApplicationPrincipals) | **Get** /api/v2/SystemService/ApplicationPrincipals | List application principals across all tenants
[**GetGlobalApplicationPrincipalsCount**](ApplicationPrincipalsAPI.md#GetGlobalApplicationPrincipalsCount) | **Get** /api/v2/SystemService/ApplicationPrincipals/Count | Count application principals across all tenants
[**GrantGlobalApplicationPrincipalPermission**](ApplicationPrincipalsAPI.md#GrantGlobalApplicationPrincipalPermission) | **Post** /api/v2/SystemService/ApplicationPrincipals/{principalId}/Permissions | Grant a permission to an application principal (any tenant)
[**ProvisionGlobalApplicationPrincipal**](ApplicationPrincipalsAPI.md#ProvisionGlobalApplicationPrincipal) | **Post** /api/v2/SystemService/ApplicationPrincipals/Provision | Provision an application principal (any tenant, incl. system-locked)
[**ProvisionPaymentsConnector**](ApplicationPrincipalsAPI.md#ProvisionPaymentsConnector) | **Post** /api/v2/SystemService/ApplicationPrincipals/PaymentsConnector | Provision the platform payments-connector identity
[**RevokeGlobalApplicationPrincipalPermission**](ApplicationPrincipalsAPI.md#RevokeGlobalApplicationPrincipalPermission) | **Delete** /api/v2/SystemService/ApplicationPrincipals/{principalId}/Permissions/{permission} | Revoke a permission from an application principal (any tenant)
[**SuspendGlobalApplicationPrincipal**](ApplicationPrincipalsAPI.md#SuspendGlobalApplicationPrincipal) | **Post** /api/v2/SystemService/ApplicationPrincipals/{principalId}/Suspend | Suspend an application principal (global)



## DisableGlobalApplicationPrincipal

> EmptyEnvelope DisableGlobalApplicationPrincipal(ctx, principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Disable an application principal (global)



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.DisableGlobalApplicationPrincipal(context.Background(), principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.DisableGlobalApplicationPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableGlobalApplicationPrincipal`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.DisableGlobalApplicationPrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableGlobalApplicationPrincipalRequest struct via the builder pattern


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


## EnableGlobalApplicationPrincipal

> EmptyEnvelope EnableGlobalApplicationPrincipal(ctx, principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Enable an application principal (global)



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.EnableGlobalApplicationPrincipal(context.Background(), principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.EnableGlobalApplicationPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableGlobalApplicationPrincipal`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.EnableGlobalApplicationPrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableGlobalApplicationPrincipalRequest struct via the builder pattern


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


## GetGlobalApplicationPrincipal

> ApplicationPrincipalDetailDtoEnvelope GetGlobalApplicationPrincipal(ctx, principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get one application principal (any tenant)



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.GetGlobalApplicationPrincipal(context.Background(), principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.GetGlobalApplicationPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalApplicationPrincipal`: ApplicationPrincipalDetailDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.GetGlobalApplicationPrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalApplicationPrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ApplicationPrincipalDetailDtoEnvelope**](ApplicationPrincipalDetailDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalApplicationPrincipals

> ApplicationPrincipalDtoIReadOnlyListEnvelope GetGlobalApplicationPrincipals(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).ApplicationPrincipalDtoCollectionQueryParameters(applicationPrincipalDtoCollectionQueryParameters).Execute()

List application principals across all tenants



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	applicationPrincipalDtoCollectionQueryParameters := *openapiclient.NewApplicationPrincipalDtoCollectionQueryParameters() // ApplicationPrincipalDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.GetGlobalApplicationPrincipals(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).ApplicationPrincipalDtoCollectionQueryParameters(applicationPrincipalDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.GetGlobalApplicationPrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalApplicationPrincipals`: ApplicationPrincipalDtoIReadOnlyListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.GetGlobalApplicationPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalApplicationPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **applicationPrincipalDtoCollectionQueryParameters** | [**ApplicationPrincipalDtoCollectionQueryParameters**](ApplicationPrincipalDtoCollectionQueryParameters.md) |  | 

### Return type

[**ApplicationPrincipalDtoIReadOnlyListEnvelope**](ApplicationPrincipalDtoIReadOnlyListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalApplicationPrincipalsCount

> Int32Envelope GetGlobalApplicationPrincipalsCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).ApplicationPrincipalDtoCollectionQueryParameters(applicationPrincipalDtoCollectionQueryParameters).Execute()

Count application principals across all tenants



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	applicationPrincipalDtoCollectionQueryParameters := *openapiclient.NewApplicationPrincipalDtoCollectionQueryParameters() // ApplicationPrincipalDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.GetGlobalApplicationPrincipalsCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).ApplicationPrincipalDtoCollectionQueryParameters(applicationPrincipalDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.GetGlobalApplicationPrincipalsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalApplicationPrincipalsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.GetGlobalApplicationPrincipalsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalApplicationPrincipalsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **applicationPrincipalDtoCollectionQueryParameters** | [**ApplicationPrincipalDtoCollectionQueryParameters**](ApplicationPrincipalDtoCollectionQueryParameters.md) |  | 

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


## GrantGlobalApplicationPrincipalPermission

> EmptyEnvelope GrantGlobalApplicationPrincipalPermission(ctx, principalId).TenantId(tenantId).ApplicationPrincipalPermissionRequestDto(applicationPrincipalPermissionRequestDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Grant a permission to an application principal (any tenant)



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	applicationPrincipalPermissionRequestDto := *openapiclient.NewApplicationPrincipalPermissionRequestDto("Permission_example") // ApplicationPrincipalPermissionRequestDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.GrantGlobalApplicationPrincipalPermission(context.Background(), principalId).TenantId(tenantId).ApplicationPrincipalPermissionRequestDto(applicationPrincipalPermissionRequestDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.GrantGlobalApplicationPrincipalPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GrantGlobalApplicationPrincipalPermission`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.GrantGlobalApplicationPrincipalPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantGlobalApplicationPrincipalPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **applicationPrincipalPermissionRequestDto** | [**ApplicationPrincipalPermissionRequestDto**](ApplicationPrincipalPermissionRequestDto.md) |  | 
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


## ProvisionGlobalApplicationPrincipal

> ApplicationPrincipalProvisioningResultDtoEnvelope ProvisionGlobalApplicationPrincipal(ctx).ApplicationPrincipalProvisionRequestDto(applicationPrincipalProvisionRequestDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Provision an application principal (any tenant, incl. system-locked)



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
	applicationPrincipalProvisionRequestDto := *openapiclient.NewApplicationPrincipalProvisionRequestDto("BusinessApplicationId_example") // ApplicationPrincipalProvisionRequestDto | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.ProvisionGlobalApplicationPrincipal(context.Background()).ApplicationPrincipalProvisionRequestDto(applicationPrincipalProvisionRequestDto).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.ProvisionGlobalApplicationPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisionGlobalApplicationPrincipal`: ApplicationPrincipalProvisioningResultDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.ProvisionGlobalApplicationPrincipal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionGlobalApplicationPrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationPrincipalProvisionRequestDto** | [**ApplicationPrincipalProvisionRequestDto**](ApplicationPrincipalProvisionRequestDto.md) |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ApplicationPrincipalProvisioningResultDtoEnvelope**](ApplicationPrincipalProvisioningResultDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionPaymentsConnector

> ApplicationPrincipalProvisioningResultDtoEnvelope ProvisionPaymentsConnector(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Provision the platform payments-connector identity



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.ProvisionPaymentsConnector(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.ProvisionPaymentsConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisionPaymentsConnector`: ApplicationPrincipalProvisioningResultDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.ProvisionPaymentsConnector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionPaymentsConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ApplicationPrincipalProvisioningResultDtoEnvelope**](ApplicationPrincipalProvisioningResultDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeGlobalApplicationPrincipalPermission

> EmptyEnvelope RevokeGlobalApplicationPrincipalPermission(ctx, principalId, permission).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Revoke a permission from an application principal (any tenant)



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	permission := "permission_example" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.RevokeGlobalApplicationPrincipalPermission(context.Background(), principalId, permission).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.RevokeGlobalApplicationPrincipalPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeGlobalApplicationPrincipalPermission`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.RevokeGlobalApplicationPrincipalPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 
**permission** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeGlobalApplicationPrincipalPermissionRequest struct via the builder pattern


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


## SuspendGlobalApplicationPrincipal

> EmptyEnvelope SuspendGlobalApplicationPrincipal(ctx, principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Suspend an application principal (global)



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.SuspendGlobalApplicationPrincipal(context.Background(), principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.SuspendGlobalApplicationPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuspendGlobalApplicationPrincipal`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.SuspendGlobalApplicationPrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendGlobalApplicationPrincipalRequest struct via the builder pattern


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

