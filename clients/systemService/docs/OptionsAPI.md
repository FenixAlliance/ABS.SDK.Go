# \OptionsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystemOption**](OptionsAPI.md#CreateSystemOption) | **Post** /api/v2/SystemService/Options | Create a new system option
[**DeleteSystemOption**](OptionsAPI.md#DeleteSystemOption) | **Delete** /api/v2/SystemService/Options/{optionId} | Delete a system option
[**GetSystemOptionById**](OptionsAPI.md#GetSystemOptionById) | **Get** /api/v2/SystemService/Options/{optionId} | Retrieve a single system option by its ID
[**GetSystemOptionByKey**](OptionsAPI.md#GetSystemOptionByKey) | **Get** /api/v2/SystemService/Options/Key/{key} | Retrieve a single system option by its key
[**GetSystemOptions**](OptionsAPI.md#GetSystemOptions) | **Get** /api/v2/SystemService/Options | Retrieve a list of system options
[**GetSystemOptionsCount**](OptionsAPI.md#GetSystemOptionsCount) | **Get** /api/v2/SystemService/Options/Count | Get the count of system options
[**UpdateSystemOption**](OptionsAPI.md#UpdateSystemOption) | **Put** /api/v2/SystemService/Options/{optionId} | Update a system option
[**UpsertSystemOption**](OptionsAPI.md#UpsertSystemOption) | **Put** /api/v2/SystemService/Options/Upsert/{key} | Create or update a system option by key



## CreateSystemOption

> EmptyEnvelope CreateSystemOption(ctx).Key(key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionCreateDto(optionCreateDto).Execute()

Create a new system option



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
	key := "key_example" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionCreateDto := *openapiclient.NewOptionCreateDto("Key_example", "Value_example") // OptionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateSystemOption(context.Background()).Key(key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionCreateDto(optionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateSystemOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSystemOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateSystemOption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemOptionRequest struct via the builder pattern


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


## DeleteSystemOption

> EmptyEnvelope DeleteSystemOption(ctx, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a system option



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.DeleteSystemOption(context.Background(), optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.DeleteSystemOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSystemOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.DeleteSystemOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemOptionRequest struct via the builder pattern


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


## GetSystemOptionById

> OptionDtoEnvelope GetSystemOptionById(ctx, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single system option by its ID



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetSystemOptionById(context.Background(), optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetSystemOptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOptionById`: OptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetSystemOptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOptionByIdRequest struct via the builder pattern


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


## GetSystemOptionByKey

> OptionDtoEnvelope GetSystemOptionByKey(ctx, key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single system option by its key



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	key := "key_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetSystemOptionByKey(context.Background(), key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetSystemOptionByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOptionByKey`: OptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetSystemOptionByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOptionByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portalId** | **string** |  | 

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


## GetSystemOptions

> OptionDtoListEnvelope GetSystemOptions(ctx).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of system options



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetSystemOptions(context.Background()).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetSystemOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOptions`: OptionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetSystemOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOptionsRequest struct via the builder pattern


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


## GetSystemOptionsCount

> Int32Envelope GetSystemOptionsCount(ctx).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of system options



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
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetSystemOptionsCount(context.Background()).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetSystemOptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemOptionsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetSystemOptionsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemOptionsCountRequest struct via the builder pattern


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


## UpdateSystemOption

> EmptyEnvelope UpdateSystemOption(ctx, optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()

Update a system option



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
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionUpdateDto := *openapiclient.NewOptionUpdateDto() // OptionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.UpdateSystemOption(context.Background(), optionId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.UpdateSystemOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSystemOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.UpdateSystemOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemOptionRequest struct via the builder pattern


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


## UpsertSystemOption

> EmptyEnvelope UpsertSystemOption(ctx, key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()

Create or update a system option by key



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
	key := "key_example" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionUpdateDto := *openapiclient.NewOptionUpdateDto() // OptionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.UpsertSystemOption(context.Background(), key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.UpsertSystemOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertSystemOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.UpsertSystemOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertSystemOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalId** | **string** |  | 
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

