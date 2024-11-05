# \AvatarsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvatar**](AvatarsAPI.md#GetAvatar) | **Get** /api/v2/StorageService/Avatars/{socialProfileId} | Get the avatar for a given social profile.
[**GetContactAvatar**](AvatarsAPI.md#GetContactAvatar) | **Get** /api/v2/StorageService/Avatars/Contact/{contactId} | Get the avatar for a given contact.
[**GetCurrentUserAvatar**](AvatarsAPI.md#GetCurrentUserAvatar) | **Get** /api/v2/StorageService/Avatars/User | Get the avatar for the current user.
[**GetTenantAvatar**](AvatarsAPI.md#GetTenantAvatar) | **Get** /api/v2/StorageService/Avatars/Tenant/{tenantId} | Get the avatar for a given tenant.
[**GetUserAvatar**](AvatarsAPI.md#GetUserAvatar) | **Get** /api/v2/StorageService/Avatars/User/{userId} | Get the avatar for a given user.
[**UpdateContactAvatar**](AvatarsAPI.md#UpdateContactAvatar) | **Post** /api/v2/StorageService/Avatars/Contacts/{contactId} | Update the avatar for a given contact.
[**UpdateTenantAvatar**](AvatarsAPI.md#UpdateTenantAvatar) | **Post** /api/v2/StorageService/Avatars/Tenant/{tenantId} | Update the avatar for a given tenant.
[**UpdateUserAvatar**](AvatarsAPI.md#UpdateUserAvatar) | **Post** /api/v2/StorageService/Avatars/User | Update the avatar for the current user.



## GetAvatar

> EmptyEnvelope GetAvatar(ctx, socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the avatar for a given social profile.

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
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.GetAvatar(context.Background(), socialProfileId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvatar`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**socialProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactAvatar

> EmptyEnvelope GetContactAvatar(ctx, contactId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the avatar for a given contact.

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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.GetContactAvatar(context.Background(), contactId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetContactAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactAvatar`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetContactAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserAvatar

> EmptyEnvelope GetCurrentUserAvatar(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the avatar for the current user.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.GetCurrentUserAvatar(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetCurrentUserAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserAvatar`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetCurrentUserAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantAvatar

> EmptyEnvelope GetTenantAvatar(ctx, tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the avatar for a given tenant.

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
	resp, r, err := apiClient.AvatarsAPI.GetTenantAvatar(context.Background(), tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetTenantAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantAvatar`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetTenantAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAvatar

> EmptyEnvelope GetUserAvatar(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the avatar for a given user.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.GetUserAvatar(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetUserAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAvatar`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetUserAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactAvatar

> EmptyEnvelope UpdateContactAvatar(ctx, contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Avatar(avatar).Execute()

Update the avatar for a given contact.

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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	avatar := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.UpdateContactAvatar(context.Background(), contactId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.UpdateContactAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContactAvatar`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.UpdateContactAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **avatar** | ***os.File** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenantAvatar

> EmptyEnvelope UpdateTenantAvatar(ctx, tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Avatar(avatar).Execute()

Update the avatar for a given tenant.

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
	avatar := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.UpdateTenantAvatar(context.Background(), tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.UpdateTenantAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTenantAvatar`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.UpdateTenantAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **avatar** | ***os.File** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserAvatar

> EmptyEnvelope UpdateUserAvatar(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Avatar(avatar).Execute()

Update the avatar for the current user.

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
	avatar := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.UpdateUserAvatar(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.UpdateUserAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserAvatar`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.UpdateUserAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **avatar** | ***os.File** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

