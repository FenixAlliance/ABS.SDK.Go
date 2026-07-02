# \UserOptionsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystemUserOption**](UserOptionsAPI.md#CreateSystemUserOption) | **Post** /api/v2/SystemService/Users/{userId}/Options | Create a new user option (admin)
[**DeleteSystemUserOption**](UserOptionsAPI.md#DeleteSystemUserOption) | **Delete** /api/v2/SystemService/Users/{userId}/Options/{optionId} | Delete a user option (admin)
[**GetSystemUserOptionById**](UserOptionsAPI.md#GetSystemUserOptionById) | **Get** /api/v2/SystemService/Users/{userId}/Options/{optionId} | Retrieve a single user option by its ID (admin)
[**GetSystemUserOptions**](UserOptionsAPI.md#GetSystemUserOptions) | **Get** /api/v2/SystemService/Users/{userId}/Options | Retrieve a list of user options (admin)
[**GetSystemUserOptionsCount**](UserOptionsAPI.md#GetSystemUserOptionsCount) | **Get** /api/v2/SystemService/Users/{userId}/Options/Count | Get the count of user options (admin)
[**PatchSystemUserOption**](UserOptionsAPI.md#PatchSystemUserOption) | **Patch** /api/v2/SystemService/Users/{userId}/Options/{optionId} | Partially update a user option (admin)
[**UpdateSystemUserOption**](UserOptionsAPI.md#UpdateSystemUserOption) | **Put** /api/v2/SystemService/Users/{userId}/Options/{optionId} | Update a user option (admin)



## CreateSystemUserOption

> EmptyEnvelope CreateSystemUserOption(ctx, userId).Key(key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionCreateDto(optionCreateDto).Execute()

Create a new user option (admin)



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
	key := "key_example" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionCreateDto := *openapiclient.NewOptionCreateDto("Key_example", "Value_example") // OptionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOptionsAPI.CreateSystemUserOption(context.Background(), userId).Key(key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionCreateDto(optionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOptionsAPI.CreateSystemUserOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSystemUserOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UserOptionsAPI.CreateSystemUserOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemUserOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 
 **portalId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **optionCreateDto** | [**OptionCreateDto**](OptionCreateDto.md) |  | 

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


## DeleteSystemUserOption

> EmptyEnvelope DeleteSystemUserOption(ctx, userId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a user option (admin)



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOptionsAPI.DeleteSystemUserOption(context.Background(), userId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOptionsAPI.DeleteSystemUserOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSystemUserOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UserOptionsAPI.DeleteSystemUserOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemUserOptionRequest struct via the builder pattern


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


## GetSystemUserOptionById

> OptionDtoEnvelope GetSystemUserOptionById(ctx, userId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single user option by its ID (admin)



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOptionsAPI.GetSystemUserOptionById(context.Background(), userId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOptionsAPI.GetSystemUserOptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemUserOptionById`: OptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UserOptionsAPI.GetSystemUserOptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemUserOptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**OptionDtoEnvelope**](OptionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemUserOptions

> OptionDtoListEnvelope GetSystemUserOptions(ctx, userId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of user options (admin)



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOptionsAPI.GetSystemUserOptions(context.Background(), userId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOptionsAPI.GetSystemUserOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemUserOptions`: OptionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UserOptionsAPI.GetSystemUserOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemUserOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**OptionDtoListEnvelope**](OptionDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemUserOptionsCount

> Int32Envelope GetSystemUserOptionsCount(ctx, userId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of user options (admin)



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOptionsAPI.GetSystemUserOptionsCount(context.Background(), userId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOptionsAPI.GetSystemUserOptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemUserOptionsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `UserOptionsAPI.GetSystemUserOptionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemUserOptionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalId** | **string** |  | 
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


## PatchSystemUserOption

> EmptyEnvelope PatchSystemUserOption(ctx, userId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Partially update a user option (admin)



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOptionsAPI.PatchSystemUserOption(context.Background(), userId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOptionsAPI.PatchSystemUserOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemUserOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UserOptionsAPI.PatchSystemUserOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemUserOptionRequest struct via the builder pattern


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

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSystemUserOption

> EmptyEnvelope UpdateSystemUserOption(ctx, userId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()

Update a user option (admin)



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionUpdateDto := *openapiclient.NewOptionUpdateDto() // OptionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserOptionsAPI.UpdateSystemUserOption(context.Background(), userId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserOptionsAPI.UpdateSystemUserOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSystemUserOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UserOptionsAPI.UpdateSystemUserOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemUserOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **optionUpdateDto** | [**OptionUpdateDto**](OptionUpdateDto.md) |  | 

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

