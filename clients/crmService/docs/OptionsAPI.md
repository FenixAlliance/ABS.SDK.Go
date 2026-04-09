# \OptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContactOption**](OptionsAPI.md#CreateContactOption) | **Post** /api/v2/CrmService/Contacts/{contactId}/Options | Create a new contact option
[**DeleteContactOption**](OptionsAPI.md#DeleteContactOption) | **Delete** /api/v2/CrmService/Contacts/{contactId}/Options/{optionId} | Delete a contact option
[**GetContactOptionById**](OptionsAPI.md#GetContactOptionById) | **Get** /api/v2/CrmService/Contacts/{contactId}/Options/{optionId} | Retrieve a single contact option by its ID
[**GetContactOptionByKey**](OptionsAPI.md#GetContactOptionByKey) | **Get** /api/v2/CrmService/Contacts/{contactId}/Options/Key/{key} | Retrieve a single contact option by its key
[**GetContactOptions**](OptionsAPI.md#GetContactOptions) | **Get** /api/v2/CrmService/Contacts/{contactId}/Options | Retrieve a list of contact options
[**GetContactOptionsCount**](OptionsAPI.md#GetContactOptionsCount) | **Get** /api/v2/CrmService/Contacts/{contactId}/Options/Count | Get the count of contact options
[**UpdateContactOption**](OptionsAPI.md#UpdateContactOption) | **Put** /api/v2/CrmService/Contacts/{contactId}/Options/{optionId} | Update a contact option
[**UpsertContactOption**](OptionsAPI.md#UpsertContactOption) | **Put** /api/v2/CrmService/Contacts/{contactId}/Options/Upsert/{key} | Create or update a contact option by key



## CreateContactOption

> EmptyEnvelope CreateContactOption(ctx, contactId).TenantId(tenantId).Key(key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionCreateDto(optionCreateDto).Execute()

Create a new contact option



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	key := "key_example" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionCreateDto := *openapiclient.NewOptionCreateDto("Key_example", "Value_example") // OptionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.CreateContactOption(context.Background(), contactId).TenantId(tenantId).Key(key).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionCreateDto(optionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.CreateContactOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.CreateContactOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## DeleteContactOption

> EmptyEnvelope DeleteContactOption(ctx, contactId, optionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a contact option



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.DeleteContactOption(context.Background(), contactId, optionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.DeleteContactOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContactOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.DeleteContactOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactOptionRequest struct via the builder pattern


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


## GetContactOptionById

> OptionDtoEnvelope GetContactOptionById(ctx, contactId, optionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single contact option by its ID



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetContactOptionById(context.Background(), contactId, optionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetContactOptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactOptionById`: OptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetContactOptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactOptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## GetContactOptionByKey

> OptionDtoEnvelope GetContactOptionByKey(ctx, contactId, key).TenantId(tenantId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single contact option by its key



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	key := "key_example" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetContactOptionByKey(context.Background(), contactId, key).TenantId(tenantId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetContactOptionByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactOptionByKey`: OptionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetContactOptionByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactOptionByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## GetContactOptions

> OptionDtoListEnvelope GetContactOptions(ctx, contactId).TenantId(tenantId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of contact options



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetContactOptions(context.Background(), contactId).TenantId(tenantId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetContactOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactOptions`: OptionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetContactOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## GetContactOptionsCount

> Int32Envelope GetContactOptionsCount(ctx, contactId).TenantId(tenantId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of contact options



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.GetContactOptionsCount(context.Background(), contactId).TenantId(tenantId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.GetContactOptionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactOptionsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.GetContactOptionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactOptionsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## UpdateContactOption

> EmptyEnvelope UpdateContactOption(ctx, contactId, optionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()

Update a contact option



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	optionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionUpdateDto := *openapiclient.NewOptionUpdateDto() // OptionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.UpdateContactOption(context.Background(), contactId, optionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.UpdateContactOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContactOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.UpdateContactOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## UpsertContactOption

> EmptyEnvelope UpsertContactOption(ctx, contactId, key).TenantId(tenantId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()

Create or update a contact option by key



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
	contactId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	key := "key_example" // string | 
	portalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	optionUpdateDto := *openapiclient.NewOptionUpdateDto() // OptionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionsAPI.UpsertContactOption(context.Background(), contactId, key).TenantId(tenantId).PortalId(portalId).ApiVersion(apiVersion).XApiVersion(xApiVersion).OptionUpdateDto(optionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionsAPI.UpsertContactOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertContactOption`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `OptionsAPI.UpsertContactOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertContactOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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

