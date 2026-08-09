# \ApplicationPrincipalsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableApplicationPrincipalAsync**](ApplicationPrincipalsAPI.md#DisableApplicationPrincipalAsync) | **Post** /api/v2/SecurityService/ApplicationPrincipals/{principalId}/Disable | Disable an application principal
[**EnableApplicationPrincipalAsync**](ApplicationPrincipalsAPI.md#EnableApplicationPrincipalAsync) | **Post** /api/v2/SecurityService/ApplicationPrincipals/{principalId}/Enable | Enable an application principal
[**GetApplicationPrincipalAsync**](ApplicationPrincipalsAPI.md#GetApplicationPrincipalAsync) | **Get** /api/v2/SecurityService/ApplicationPrincipals/{principalId} | Get application principal by ID
[**GetApplicationPrincipalsAsync**](ApplicationPrincipalsAPI.md#GetApplicationPrincipalsAsync) | **Get** /api/v2/SecurityService/ApplicationPrincipals | Get all application principals
[**GetApplicationPrincipalsCountAsync**](ApplicationPrincipalsAPI.md#GetApplicationPrincipalsCountAsync) | **Get** /api/v2/SecurityService/ApplicationPrincipals/Count | Get application principals count
[**GrantPermissionAsync**](ApplicationPrincipalsAPI.md#GrantPermissionAsync) | **Post** /api/v2/SecurityService/ApplicationPrincipals/{principalId}/Permissions | Grant a permission to an application principal
[**ProvisionApplicationPrincipalAsync**](ApplicationPrincipalsAPI.md#ProvisionApplicationPrincipalAsync) | **Post** /api/v2/SecurityService/ApplicationPrincipals/Provision | Provision an application principal
[**RevokePermissionAsync**](ApplicationPrincipalsAPI.md#RevokePermissionAsync) | **Delete** /api/v2/SecurityService/ApplicationPrincipals/{principalId}/Permissions/{permission} | Revoke a permission from an application principal
[**SuspendApplicationPrincipalAsync**](ApplicationPrincipalsAPI.md#SuspendApplicationPrincipalAsync) | **Post** /api/v2/SecurityService/ApplicationPrincipals/{principalId}/Suspend | Suspend an application principal



## DisableApplicationPrincipalAsync

> EmptyEnvelope DisableApplicationPrincipalAsync(ctx, principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Disable an application principal



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.DisableApplicationPrincipalAsync(context.Background(), principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.DisableApplicationPrincipalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableApplicationPrincipalAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.DisableApplicationPrincipalAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableApplicationPrincipalAsyncRequest struct via the builder pattern


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


## EnableApplicationPrincipalAsync

> EmptyEnvelope EnableApplicationPrincipalAsync(ctx, principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Enable an application principal



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.EnableApplicationPrincipalAsync(context.Background(), principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.EnableApplicationPrincipalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableApplicationPrincipalAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.EnableApplicationPrincipalAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableApplicationPrincipalAsyncRequest struct via the builder pattern


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


## GetApplicationPrincipalAsync

> ApplicationPrincipalDetailDtoEnvelope GetApplicationPrincipalAsync(ctx, principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get application principal by ID



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.GetApplicationPrincipalAsync(context.Background(), principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.GetApplicationPrincipalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationPrincipalAsync`: ApplicationPrincipalDetailDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.GetApplicationPrincipalAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPrincipalAsyncRequest struct via the builder pattern


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


## GetApplicationPrincipalsAsync

> ApplicationPrincipalDtoListEnvelope GetApplicationPrincipalsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ApplicationPrincipalDtoCollectionQueryParameters(applicationPrincipalDtoCollectionQueryParameters).Execute()

Get all application principals



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
	applicationPrincipalDtoCollectionQueryParameters := *openapiclient.NewApplicationPrincipalDtoCollectionQueryParameters() // ApplicationPrincipalDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.GetApplicationPrincipalsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ApplicationPrincipalDtoCollectionQueryParameters(applicationPrincipalDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.GetApplicationPrincipalsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationPrincipalsAsync`: ApplicationPrincipalDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.GetApplicationPrincipalsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPrincipalsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **applicationPrincipalDtoCollectionQueryParameters** | [**ApplicationPrincipalDtoCollectionQueryParameters**](ApplicationPrincipalDtoCollectionQueryParameters.md) |  | 

### Return type

[**ApplicationPrincipalDtoListEnvelope**](ApplicationPrincipalDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationPrincipalsCountAsync

> Int32Envelope GetApplicationPrincipalsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ApplicationPrincipalDtoCollectionQueryParameters(applicationPrincipalDtoCollectionQueryParameters).Execute()

Get application principals count



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
	applicationPrincipalDtoCollectionQueryParameters := *openapiclient.NewApplicationPrincipalDtoCollectionQueryParameters() // ApplicationPrincipalDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.GetApplicationPrincipalsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ApplicationPrincipalDtoCollectionQueryParameters(applicationPrincipalDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.GetApplicationPrincipalsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationPrincipalsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.GetApplicationPrincipalsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPrincipalsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
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


## GrantPermissionAsync

> EmptyEnvelope GrantPermissionAsync(ctx, principalId).TenantId(tenantId).ApplicationPrincipalPermissionRequestDto(applicationPrincipalPermissionRequestDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Grant a permission to an application principal



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	applicationPrincipalPermissionRequestDto := *openapiclient.NewApplicationPrincipalPermissionRequestDto("Permission_example") // ApplicationPrincipalPermissionRequestDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.GrantPermissionAsync(context.Background(), principalId).TenantId(tenantId).ApplicationPrincipalPermissionRequestDto(applicationPrincipalPermissionRequestDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.GrantPermissionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GrantPermissionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.GrantPermissionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantPermissionAsyncRequest struct via the builder pattern


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


## ProvisionApplicationPrincipalAsync

> ApplicationPrincipalProvisioningResultDtoEnvelope ProvisionApplicationPrincipalAsync(ctx).TenantId(tenantId).ApplicationPrincipalProvisionRequestDto(applicationPrincipalProvisionRequestDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Provision an application principal



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
	applicationPrincipalProvisionRequestDto := *openapiclient.NewApplicationPrincipalProvisionRequestDto("BusinessApplicationId_example") // ApplicationPrincipalProvisionRequestDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.ProvisionApplicationPrincipalAsync(context.Background()).TenantId(tenantId).ApplicationPrincipalProvisionRequestDto(applicationPrincipalProvisionRequestDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.ProvisionApplicationPrincipalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisionApplicationPrincipalAsync`: ApplicationPrincipalProvisioningResultDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.ProvisionApplicationPrincipalAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisionApplicationPrincipalAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **applicationPrincipalProvisionRequestDto** | [**ApplicationPrincipalProvisionRequestDto**](ApplicationPrincipalProvisionRequestDto.md) |  | 
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


## RevokePermissionAsync

> EmptyEnvelope RevokePermissionAsync(ctx, principalId, permission).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Revoke a permission from an application principal



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	permission := "permission_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.RevokePermissionAsync(context.Background(), principalId, permission).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.RevokePermissionAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokePermissionAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.RevokePermissionAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 
**permission** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionAsyncRequest struct via the builder pattern


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


## SuspendApplicationPrincipalAsync

> EmptyEnvelope SuspendApplicationPrincipalAsync(ctx, principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Suspend an application principal



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
	principalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationPrincipalsAPI.SuspendApplicationPrincipalAsync(context.Background(), principalId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPrincipalsAPI.SuspendApplicationPrincipalAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuspendApplicationPrincipalAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ApplicationPrincipalsAPI.SuspendApplicationPrincipalAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendApplicationPrincipalAsyncRequest struct via the builder pattern


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

