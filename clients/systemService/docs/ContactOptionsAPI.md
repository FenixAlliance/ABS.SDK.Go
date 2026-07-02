# \ContactOptionsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystemContactOption**](ContactOptionsAPI.md#CreateSystemContactOption) | **Post** /api/v2/SystemService/Contacts/{contactId}/Options | Create a new contact option (admin)
[**DeleteSystemContactOption**](ContactOptionsAPI.md#DeleteSystemContactOption) | **Delete** /api/v2/SystemService/Contacts/{contactId}/Options/{optionId} | Delete a contact option (admin)
[**GetSystemContactOptionById**](ContactOptionsAPI.md#GetSystemContactOptionById) | **Get** /api/v2/SystemService/Contacts/{contactId}/Options/{optionId} | Retrieve a single contact option by its ID (admin)
[**GetSystemContactOptions**](ContactOptionsAPI.md#GetSystemContactOptions) | **Get** /api/v2/SystemService/Contacts/{contactId}/Options | Retrieve a list of contact options (admin)
[**GetSystemContactOptionsCount**](ContactOptionsAPI.md#GetSystemContactOptionsCount) | **Get** /api/v2/SystemService/Contacts/{contactId}/Options/Count | Get the count of contact options (admin)
[**PatchSystemContactOption**](ContactOptionsAPI.md#PatchSystemContactOption) | **Patch** /api/v2/SystemService/Contacts/{contactId}/Options/{optionId} | Partially update a contact option (admin)
[**UpdateSystemContactOption**](ContactOptionsAPI.md#UpdateSystemContactOption) | **Put** /api/v2/SystemService/Contacts/{contactId}/Options/{optionId} | Update a contact option (admin)



## CreateSystemContactOption

> EmptyEnvelope CreateSystemContactOption(ctx, contactId).Key(key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionCreateDto(optionCreateDto).Execute()

Create a new contact option (admin)



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
	key := "key_example" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionCreateDto := *openapiclient.NewOptionCreateDto("Key_example", "Value_example") // OptionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactOptionsAPI.CreateSystemContactOption(context.Background(), contactId).Key(key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionCreateDto(optionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactOptionsAPI.CreateSystemContactOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSystemContactOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ContactOptionsAPI.CreateSystemContactOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemContactOptionRequest struct via the builder pattern


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


## DeleteSystemContactOption

> EmptyEnvelope DeleteSystemContactOption(ctx, contactId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a contact option (admin)



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactOptionsAPI.DeleteSystemContactOption(context.Background(), contactId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactOptionsAPI.DeleteSystemContactOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSystemContactOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ContactOptionsAPI.DeleteSystemContactOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemContactOptionRequest struct via the builder pattern


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


## GetSystemContactOptionById

> OptionDtoEnvelope GetSystemContactOptionById(ctx, contactId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single contact option by its ID (admin)



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactOptionsAPI.GetSystemContactOptionById(context.Background(), contactId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactOptionsAPI.GetSystemContactOptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemContactOptionById`: OptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ContactOptionsAPI.GetSystemContactOptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemContactOptionByIdRequest struct via the builder pattern


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


## GetSystemContactOptions

> OptionDtoListEnvelope GetSystemContactOptions(ctx, contactId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of contact options (admin)



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactOptionsAPI.GetSystemContactOptions(context.Background(), contactId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactOptionsAPI.GetSystemContactOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemContactOptions`: OptionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ContactOptionsAPI.GetSystemContactOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemContactOptionsRequest struct via the builder pattern


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


## GetSystemContactOptionsCount

> Int32Envelope GetSystemContactOptionsCount(ctx, contactId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of contact options (admin)



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactOptionsAPI.GetSystemContactOptionsCount(context.Background(), contactId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactOptionsAPI.GetSystemContactOptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemContactOptionsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ContactOptionsAPI.GetSystemContactOptionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemContactOptionsCountRequest struct via the builder pattern


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


## PatchSystemContactOption

> EmptyEnvelope PatchSystemContactOption(ctx, contactId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Partially update a contact option (admin)



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactOptionsAPI.PatchSystemContactOption(context.Background(), contactId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactOptionsAPI.PatchSystemContactOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSystemContactOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ContactOptionsAPI.PatchSystemContactOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSystemContactOptionRequest struct via the builder pattern


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


## UpdateSystemContactOption

> EmptyEnvelope UpdateSystemContactOption(ctx, contactId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()

Update a contact option (admin)



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionUpdateDto := *openapiclient.NewOptionUpdateDto() // OptionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactOptionsAPI.UpdateSystemContactOption(context.Background(), contactId, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactOptionsAPI.UpdateSystemContactOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSystemContactOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ContactOptionsAPI.UpdateSystemContactOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemContactOptionRequest struct via the builder pattern


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

