# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountCurrentUserFollowersAsync**](UsersAPI.md#CountCurrentUserFollowersAsync) | **Get** /api/v2/Me/Followers/Count | Count the social profiles that follow the current user
[**CountCurrentUserFollowsAsync**](UsersAPI.md#CountCurrentUserFollowsAsync) | **Get** /api/v2/Me/Follows/Count | Count the social profiles that the current user follows
[**CountCurrentUserNotificationsAsync**](UsersAPI.md#CountCurrentUserNotificationsAsync) | **Get** /api/v2/Me/Notifications/Count | Count the notifications for the current user
[**CountCurrentUserTenantsAsync**](UsersAPI.md#CountCurrentUserTenantsAsync) | **Get** /api/v2/Me/Tenants/Count | Count the tenants that the current user is enrolled in
[**GetCurrentUserAddressesAsync**](UsersAPI.md#GetCurrentUserAddressesAsync) | **Get** /api/v2/Me/Addresses | Get the list of addresses for the current user
[**GetCurrentUserAsync**](UsersAPI.md#GetCurrentUserAsync) | **Get** /api/v2/Me | Gets the current user
[**GetCurrentUserAvatarAsync**](UsersAPI.md#GetCurrentUserAvatarAsync) | **Get** /api/v2/Me/Avatar | Get the current user&#39;s avatar
[**GetCurrentUserCartAsync**](UsersAPI.md#GetCurrentUserCartAsync) | **Get** /api/v2/Me/Cart | Get the current user&#39;s cart
[**GetCurrentUserEnrollmentsAsync**](UsersAPI.md#GetCurrentUserEnrollmentsAsync) | **Get** /api/v2/Me/Enrollments | Get the list of enrollments for the current user
[**GetCurrentUserEnrollmentsExtendedAsync**](UsersAPI.md#GetCurrentUserEnrollmentsExtendedAsync) | **Get** /api/v2/Me/Enrollments/Extended | Get the list of enrollments for the current user
[**GetCurrentUserFollowersAsync**](UsersAPI.md#GetCurrentUserFollowersAsync) | **Get** /api/v2/Me/Followers | Get the social profiles that follow the current user
[**GetCurrentUserFollowsAsync**](UsersAPI.md#GetCurrentUserFollowsAsync) | **Get** /api/v2/Me/Follows | Get the social profiles that the current user follows
[**GetCurrentUserInvitationAsync**](UsersAPI.md#GetCurrentUserInvitationAsync) | **Get** /api/v2/Me/Invitations | Get the list of tenant enrollment invitations for the current user
[**GetCurrentUserNotificationsAsync**](UsersAPI.md#GetCurrentUserNotificationsAsync) | **Get** /api/v2/Me/Notifications | Get the list of notifications for the current user
[**GetCurrentUserSettingsAsync**](UsersAPI.md#GetCurrentUserSettingsAsync) | **Get** /api/v2/Me/Settings | Get the settings for the current user
[**GetCurrentUserSocialProfileAsync**](UsersAPI.md#GetCurrentUserSocialProfileAsync) | **Get** /api/v2/Me/SocialProfile | Get the current user&#39;s social profile
[**GetCurrentUserTenantsAsync**](UsersAPI.md#GetCurrentUserTenantsAsync) | **Get** /api/v2/Me/Tenants | Get the tenants that the current user is enrolled in
[**GetCurrentUserTenantsExtendedAsync**](UsersAPI.md#GetCurrentUserTenantsExtendedAsync) | **Get** /api/v2/Me/Tenants/Extended | Get the tenants that the current user is enrolled in
[**GetCurrentUserWalletAsync**](UsersAPI.md#GetCurrentUserWalletAsync) | **Get** /api/v2/Me/Wallet | Get the current user&#39;s billing profile
[**GetEnrollmentAsync**](UsersAPI.md#GetEnrollmentAsync) | **Get** /api/v2/Me/Enrollments/{enrollmentId} | Get a single TenantEnrollment by its ID
[**GetExtendedCurrentUserAsync**](UsersAPI.md#GetExtendedCurrentUserAsync) | **Get** /api/v2/Me/Extended | Get the current user&#39;s extended profile
[**PatchCurrentUserAsync**](UsersAPI.md#PatchCurrentUserAsync) | **Patch** /api/v2/Me | Partially update the current user&#39;s profile
[**UpdateAvatarAsync**](UsersAPI.md#UpdateAvatarAsync) | **Post** /api/v2/Me/Avatar | Update the current user&#39;s avatar
[**UpdateCurrentUserAsync**](UsersAPI.md#UpdateCurrentUserAsync) | **Put** /api/v2/Me | Update the current user&#39;s profile
[**UpdateCurrentUserSettingsAsync**](UsersAPI.md#UpdateCurrentUserSettingsAsync) | **Put** /api/v2/Me/Settings | Update the settings for the current user



## CountCurrentUserFollowersAsync

> Int32Envelope CountCurrentUserFollowersAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count the social profiles that follow the current user



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
	resp, r, err := apiClient.UsersAPI.CountCurrentUserFollowersAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CountCurrentUserFollowersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCurrentUserFollowersAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CountCurrentUserFollowersAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCurrentUserFollowersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountCurrentUserFollowsAsync

> Int32Envelope CountCurrentUserFollowsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count the social profiles that the current user follows



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
	resp, r, err := apiClient.UsersAPI.CountCurrentUserFollowsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CountCurrentUserFollowsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCurrentUserFollowsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CountCurrentUserFollowsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCurrentUserFollowsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountCurrentUserNotificationsAsync

> Int32Envelope CountCurrentUserNotificationsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count the notifications for the current user



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
	resp, r, err := apiClient.UsersAPI.CountCurrentUserNotificationsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CountCurrentUserNotificationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCurrentUserNotificationsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CountCurrentUserNotificationsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCurrentUserNotificationsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountCurrentUserTenantsAsync

> Int32Envelope CountCurrentUserTenantsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count the tenants that the current user is enrolled in



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
	resp, r, err := apiClient.UsersAPI.CountCurrentUserTenantsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CountCurrentUserTenantsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCurrentUserTenantsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CountCurrentUserTenantsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCurrentUserTenantsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserAddressesAsync

> AddressDtoListEnvelope GetCurrentUserAddressesAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the list of addresses for the current user



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserAddressesAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserAddressesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserAddressesAsync`: AddressDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserAddressesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserAddressesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**AddressDtoListEnvelope**](AddressDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserAsync

> UserDtoEnvelope GetCurrentUserAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current user



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserAsync`: UserDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**UserDtoEnvelope**](UserDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserAvatarAsync

> *os.File GetCurrentUserAvatarAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the current user's avatar



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserAvatarAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserAvatarAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserAvatarAsync`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserAvatarAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserAvatarAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserCartAsync

> CartDtoEnvelope GetCurrentUserCartAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the current user's cart



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserCartAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserCartAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserCartAsync`: CartDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserCartAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserCartAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CartDtoEnvelope**](CartDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserEnrollmentsAsync

> TenantEnrollmentDtoListEnvelope GetCurrentUserEnrollmentsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the list of enrollments for the current user



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserEnrollmentsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserEnrollmentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserEnrollmentsAsync`: TenantEnrollmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserEnrollmentsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserEnrollmentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantEnrollmentDtoListEnvelope**](TenantEnrollmentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserEnrollmentsExtendedAsync

> ExtendedTenantEnrollmentDtoListEnvelope GetCurrentUserEnrollmentsExtendedAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the list of enrollments for the current user



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserEnrollmentsExtendedAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserEnrollmentsExtendedAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserEnrollmentsExtendedAsync`: ExtendedTenantEnrollmentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserEnrollmentsExtendedAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserEnrollmentsExtendedAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExtendedTenantEnrollmentDtoListEnvelope**](ExtendedTenantEnrollmentDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserFollowersAsync

> FollowRecordDtoListEnvelope GetCurrentUserFollowersAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the social profiles that follow the current user



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserFollowersAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserFollowersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserFollowersAsync`: FollowRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserFollowersAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserFollowersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FollowRecordDtoListEnvelope**](FollowRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserFollowsAsync

> FollowRecordDtoListEnvelope GetCurrentUserFollowsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the social profiles that the current user follows



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserFollowsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserFollowsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserFollowsAsync`: FollowRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserFollowsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserFollowsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FollowRecordDtoListEnvelope**](FollowRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserInvitationAsync

> TenantInvitationDtoListEnvelope GetCurrentUserInvitationAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the list of tenant enrollment invitations for the current user



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserInvitationAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserInvitationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserInvitationAsync`: TenantInvitationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserInvitationAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserInvitationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantInvitationDtoListEnvelope**](TenantInvitationDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserNotificationsAsync

> NotificationDtoListEnvelope GetCurrentUserNotificationsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the list of notifications for the current user



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserNotificationsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserNotificationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserNotificationsAsync`: NotificationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserNotificationsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserNotificationsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**NotificationDtoListEnvelope**](NotificationDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserSettingsAsync

> UserSettingsDtoEnvelope GetCurrentUserSettingsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the settings for the current user



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserSettingsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserSettingsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserSettingsAsync`: UserSettingsDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserSettingsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserSettingsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**UserSettingsDtoEnvelope**](UserSettingsDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserSocialProfileAsync

> SocialProfileDtoEnvelope GetCurrentUserSocialProfileAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the current user's social profile



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserSocialProfileAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserSocialProfileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserSocialProfileAsync`: SocialProfileDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserSocialProfileAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserSocialProfileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**SocialProfileDtoEnvelope**](SocialProfileDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserTenantsAsync

> TenantDtoListEnvelope GetCurrentUserTenantsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the tenants that the current user is enrolled in



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserTenantsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserTenantsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserTenantsAsync`: TenantDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserTenantsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserTenantsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantDtoListEnvelope**](TenantDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserTenantsExtendedAsync

> ExtendedTenantDtoListEnvelope GetCurrentUserTenantsExtendedAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the tenants that the current user is enrolled in



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserTenantsExtendedAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserTenantsExtendedAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserTenantsExtendedAsync`: ExtendedTenantDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserTenantsExtendedAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserTenantsExtendedAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExtendedTenantDtoListEnvelope**](ExtendedTenantDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUserWalletAsync

> WalletDtoEnvelope GetCurrentUserWalletAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the current user's billing profile



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUserWalletAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserWalletAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUserWalletAsync`: WalletDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserWalletAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserWalletAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WalletDtoEnvelope**](WalletDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnrollmentAsync

> TenantEnrollmentDtoEnvelope GetEnrollmentAsync(ctx, enrollmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get a single TenantEnrollment by its ID



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
	enrollmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetEnrollmentAsync(context.Background(), enrollmentId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetEnrollmentAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnrollmentAsync`: TenantEnrollmentDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetEnrollmentAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enrollmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnrollmentAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TenantEnrollmentDtoEnvelope**](TenantEnrollmentDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtendedCurrentUserAsync

> ExtendedUserDtoEnvelope GetExtendedCurrentUserAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the current user's extended profile



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
	resp, r, err := apiClient.UsersAPI.GetExtendedCurrentUserAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetExtendedCurrentUserAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedCurrentUserAsync`: ExtendedUserDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetExtendedCurrentUserAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedCurrentUserAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExtendedUserDtoEnvelope**](ExtendedUserDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCurrentUserAsync

> EmptyEnvelope PatchCurrentUserAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Partially update the current user's profile



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
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.PatchCurrentUserAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PatchCurrentUserAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCurrentUserAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PatchCurrentUserAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchCurrentUserAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, multipart/form-data
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAvatarAsync

> EmptyEnvelope UpdateAvatarAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Avatar(avatar).Execute()

Update the current user's avatar



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
	resp, r, err := apiClient.UsersAPI.UpdateAvatarAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateAvatarAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAvatarAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateAvatarAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAvatarAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **avatar** | ***os.File** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json, application/xml
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrentUserAsync

> EmptyEnvelope UpdateCurrentUserAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserUpdateDto(userUpdateDto).Execute()

Update the current user's profile



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
	userUpdateDto := *openapiclient.NewUserUpdateDto() // UserUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateCurrentUserAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserUpdateDto(userUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateCurrentUserAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurrentUserAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateCurrentUserAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentUserAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **userUpdateDto** | [**UserUpdateDto**](UserUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, multipart/form-data
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrentUserSettingsAsync

> UserSettingsDtoEnvelope UpdateCurrentUserSettingsAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserSettingsUpdateDto(userSettingsUpdateDto).Execute()

Update the settings for the current user



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
	userSettingsUpdateDto := *openapiclient.NewUserSettingsUpdateDto("DateFormat_example", "CurrencyFormat_example", "DateTimeFormat_example", "SiteTheme_example") // UserSettingsUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateCurrentUserSettingsAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserSettingsUpdateDto(userSettingsUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateCurrentUserSettingsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurrentUserSettingsAsync`: UserSettingsDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateCurrentUserSettingsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentUserSettingsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **userSettingsUpdateDto** | [**UserSettingsUpdateDto**](UserSettingsUpdateDto.md) |  | 

### Return type

[**UserSettingsDtoEnvelope**](UserSettingsDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, multipart/form-data
- **Accept**: application/json, application/xml, multipart/form-data

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

