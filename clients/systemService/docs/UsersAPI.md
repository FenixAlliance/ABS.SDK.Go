# \UsersAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminPreviewUserEmailTemplate**](UsersAPI.md#AdminPreviewUserEmailTemplate) | **Post** /api/v2/SystemService/Users/{userId}/Emails/Preview | Preview the rendered email for a user.
[**AdminSendUserEmail**](UsersAPI.md#AdminSendUserEmail) | **Post** /api/v2/SystemService/Users/{userId}/Emails/Send | Send an email to a user.
[**CreateAccountHolderAsync**](UsersAPI.md#CreateAccountHolderAsync) | **Post** /api/v2/SystemService/Users | Create a new user
[**DeleteAccountHolderAsync**](UsersAPI.md#DeleteAccountHolderAsync) | **Delete** /api/v2/SystemService/Users/{userId} | Delete a user
[**GetExtendedAccountHolderAsync**](UsersAPI.md#GetExtendedAccountHolderAsync) | **Get** /api/v2/SystemService/Users/{userId}/Extended | Retrieve an extended user by ID
[**GetExtendedUsersAsync**](UsersAPI.md#GetExtendedUsersAsync) | **Get** /api/v2/SystemService/Users/Extended | Retrieve a list of extended users
[**GetExtendedUsersCountAsync**](UsersAPI.md#GetExtendedUsersCountAsync) | **Get** /api/v2/SystemService/Users/Extended/Count | Get the count of extended users
[**GetUserAdminDetailAsync**](UsersAPI.md#GetUserAdminDetailAsync) | **Get** /api/v2/SystemService/Users/{userId}/AdminDetail | Retrieve the admin detail aggregate for a user
[**GetUserAsync**](UsersAPI.md#GetUserAsync) | **Get** /api/v2/SystemService/Users/{userId} | Retrieve a user by ID
[**GetUsersAsync**](UsersAPI.md#GetUsersAsync) | **Get** /api/v2/SystemService/Users | Retrieve a list of users
[**GetUsersCountAsync**](UsersAPI.md#GetUsersCountAsync) | **Get** /api/v2/SystemService/Users/Count | Get the count of users
[**PatchAccountHolderAsync**](UsersAPI.md#PatchAccountHolderAsync) | **Patch** /api/v2/SystemService/Users/{userId} | Partially update a user
[**SetUserPasswordAsync**](UsersAPI.md#SetUserPasswordAsync) | **Post** /api/v2/SystemService/Users/{userId}/Password | Set a user&#39;s password
[**UpdateAccountHolderAdminProfileAsync**](UsersAPI.md#UpdateAccountHolderAdminProfileAsync) | **Put** /api/v2/SystemService/Users/{userId}/AdminProfile | Update a user&#39;s admin-managed profile
[**UpdateAccountHolderAsync**](UsersAPI.md#UpdateAccountHolderAsync) | **Put** /api/v2/SystemService/Users/{userId} | Update a user



## AdminPreviewUserEmailTemplate

> AdminPreviewUserEmailTemplate(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).EmailDispatchRequest(emailDispatchRequest).Execute()

Preview the rendered email for a user.



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
	emailDispatchRequest := *openapiclient.NewEmailDispatchRequest("Title_example", "Message_example", "Culture_example", "UiCulture_example", []string{"Recipients_example"}) // EmailDispatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminPreviewUserEmailTemplate(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).EmailDispatchRequest(emailDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminPreviewUserEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminPreviewUserEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **emailDispatchRequest** | [**EmailDispatchRequest**](EmailDispatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSendUserEmail

> EmptyEnvelope AdminSendUserEmail(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).EmailDispatchRequest(emailDispatchRequest).Execute()

Send an email to a user.



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
	emailDispatchRequest := *openapiclient.NewEmailDispatchRequest("Title_example", "Message_example", "Culture_example", "UiCulture_example", []string{"Recipients_example"}) // EmailDispatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminSendUserEmail(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).EmailDispatchRequest(emailDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminSendUserEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSendUserEmail`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminSendUserEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSendUserEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **emailDispatchRequest** | [**EmailDispatchRequest**](EmailDispatchRequest.md) |  | 

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


## CreateAccountHolderAsync

> EmptyEnvelope CreateAccountHolderAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserCreateDto(userCreateDto).Execute()

Create a new user



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
	userCreateDto := *openapiclient.NewUserCreateDto() // UserCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CreateAccountHolderAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserCreateDto(userCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateAccountHolderAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountHolderAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateAccountHolderAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountHolderAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **userCreateDto** | [**UserCreateDto**](UserCreateDto.md) |  | 

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


## DeleteAccountHolderAsync

> EmptyEnvelope DeleteAccountHolderAsync(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a user



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
	resp, r, err := apiClient.UsersAPI.DeleteAccountHolderAsync(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteAccountHolderAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountHolderAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteAccountHolderAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountHolderAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## GetExtendedAccountHolderAsync

> ExtendedUserDtoEnvelope GetExtendedAccountHolderAsync(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve an extended user by ID



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
	userId := "userId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetExtendedAccountHolderAsync(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetExtendedAccountHolderAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedAccountHolderAsync`: ExtendedUserDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetExtendedAccountHolderAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedAccountHolderAsyncRequest struct via the builder pattern


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
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtendedUsersAsync

> ExtendedUserDtoListEnvelope GetExtendedUsersAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).ExtendedUserDtoCollectionQueryParameters(extendedUserDtoCollectionQueryParameters).Execute()

Retrieve a list of extended users



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
	extendedUserDtoCollectionQueryParameters := *openapiclient.NewExtendedUserDtoCollectionQueryParameters() // ExtendedUserDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetExtendedUsersAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).ExtendedUserDtoCollectionQueryParameters(extendedUserDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetExtendedUsersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedUsersAsync`: ExtendedUserDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetExtendedUsersAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedUsersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **extendedUserDtoCollectionQueryParameters** | [**ExtendedUserDtoCollectionQueryParameters**](ExtendedUserDtoCollectionQueryParameters.md) |  | 

### Return type

[**ExtendedUserDtoListEnvelope**](ExtendedUserDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtendedUsersCountAsync

> Int32Envelope GetExtendedUsersCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).ExtendedUserDtoCollectionQueryParameters(extendedUserDtoCollectionQueryParameters).Execute()

Get the count of extended users



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
	extendedUserDtoCollectionQueryParameters := *openapiclient.NewExtendedUserDtoCollectionQueryParameters() // ExtendedUserDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetExtendedUsersCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).ExtendedUserDtoCollectionQueryParameters(extendedUserDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetExtendedUsersCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedUsersCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetExtendedUsersCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedUsersCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **extendedUserDtoCollectionQueryParameters** | [**ExtendedUserDtoCollectionQueryParameters**](ExtendedUserDtoCollectionQueryParameters.md) |  | 

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


## GetUserAdminDetailAsync

> UserAdminDetailDtoEnvelope GetUserAdminDetailAsync(ctx, userId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve the admin detail aggregate for a user



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
	userId := "userId_example" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserAdminDetailAsync(context.Background(), userId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserAdminDetailAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAdminDetailAsync`: UserAdminDetailDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserAdminDetailAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAdminDetailAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**UserAdminDetailDtoEnvelope**](UserAdminDetailDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAsync

> UserDtoEnvelope GetUserAsync(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a user by ID



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
	userId := "userId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserAsync(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAsync`: UserDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAsyncRequest struct via the builder pattern


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
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersAsync

> UserDtoListEnvelope GetUsersAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserDtoCollectionQueryParameters(userDtoCollectionQueryParameters).Execute()

Retrieve a list of users



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
	userDtoCollectionQueryParameters := *openapiclient.NewUserDtoCollectionQueryParameters() // UserDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsersAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserDtoCollectionQueryParameters(userDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersAsync`: UserDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **userDtoCollectionQueryParameters** | [**UserDtoCollectionQueryParameters**](UserDtoCollectionQueryParameters.md) |  | 

### Return type

[**UserDtoListEnvelope**](UserDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersCountAsync

> Int32Envelope GetUsersCountAsync(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserDtoCollectionQueryParameters(userDtoCollectionQueryParameters).Execute()

Get the count of users



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
	userDtoCollectionQueryParameters := *openapiclient.NewUserDtoCollectionQueryParameters() // UserDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsersCountAsync(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserDtoCollectionQueryParameters(userDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **userDtoCollectionQueryParameters** | [**UserDtoCollectionQueryParameters**](UserDtoCollectionQueryParameters.md) |  | 

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


## PatchAccountHolderAsync

> EmptyEnvelope PatchAccountHolderAsync(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Partially update a user



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
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.PatchAccountHolderAsync(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PatchAccountHolderAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAccountHolderAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PatchAccountHolderAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAccountHolderAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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


## SetUserPasswordAsync

> EmptyEnvelope SetUserPasswordAsync(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SetUserPasswordDto(setUserPasswordDto).Execute()

Set a user's password



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
	setUserPasswordDto := *openapiclient.NewSetUserPasswordDto("NewPassword_example") // SetUserPasswordDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SetUserPasswordAsync(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).SetUserPasswordDto(setUserPasswordDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SetUserPasswordAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetUserPasswordAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SetUserPasswordAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetUserPasswordAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **setUserPasswordDto** | [**SetUserPasswordDto**](SetUserPasswordDto.md) |  | 

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


## UpdateAccountHolderAdminProfileAsync

> EmptyEnvelope UpdateAccountHolderAdminProfileAsync(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserAdminUpdateDto(userAdminUpdateDto).Execute()

Update a user's admin-managed profile



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
	userAdminUpdateDto := *openapiclient.NewUserAdminUpdateDto() // UserAdminUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateAccountHolderAdminProfileAsync(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserAdminUpdateDto(userAdminUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateAccountHolderAdminProfileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountHolderAdminProfileAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateAccountHolderAdminProfileAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountHolderAdminProfileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **userAdminUpdateDto** | [**UserAdminUpdateDto**](UserAdminUpdateDto.md) |  | 

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


## UpdateAccountHolderAsync

> EmptyEnvelope UpdateAccountHolderAsync(ctx, userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserUpdateDto(userUpdateDto).Execute()

Update a user



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
	userUpdateDto := *openapiclient.NewUserUpdateDto() // UserUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateAccountHolderAsync(context.Background(), userId).ApiVersion(apiVersion).XApiVersion(xApiVersion).UserUpdateDto(userUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateAccountHolderAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountHolderAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateAccountHolderAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountHolderAsyncRequest struct via the builder pattern


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

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

