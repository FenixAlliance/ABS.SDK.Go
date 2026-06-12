# \LocalizationStringsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountLocalizationStringsAsync**](LocalizationStringsAPI.md#CountLocalizationStringsAsync) | **Get** /api/v2/ContentService/LocalizationStrings/Count | Count localization strings
[**CreateLocalizationStringAsync**](LocalizationStringsAPI.md#CreateLocalizationStringAsync) | **Post** /api/v2/ContentService/LocalizationStrings | Create a localization string
[**DeleteLocalizationStringAsync**](LocalizationStringsAPI.md#DeleteLocalizationStringAsync) | **Delete** /api/v2/ContentService/LocalizationStrings/{localizationStringId} | Delete a localization string
[**GetLocalizationStringByIdAsync**](LocalizationStringsAPI.md#GetLocalizationStringByIdAsync) | **Get** /api/v2/ContentService/LocalizationStrings/{localizationStringId} | Get localization string by ID
[**GetLocalizationStringsAsync**](LocalizationStringsAPI.md#GetLocalizationStringsAsync) | **Get** /api/v2/ContentService/LocalizationStrings | Get localization strings
[**UpdateLocalizationStringAsync**](LocalizationStringsAPI.md#UpdateLocalizationStringAsync) | **Put** /api/v2/ContentService/LocalizationStrings/{localizationStringId} | Update a localization string



## CountLocalizationStringsAsync

> Int32Envelope CountLocalizationStringsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count localization strings



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
	resp, r, err := apiClient.LocalizationStringsAPI.CountLocalizationStringsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationStringsAPI.CountLocalizationStringsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountLocalizationStringsAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `LocalizationStringsAPI.CountLocalizationStringsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountLocalizationStringsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
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


## CreateLocalizationStringAsync

> EmptyEnvelope CreateLocalizationStringAsync(ctx).TenantId(tenantId).LocalizationStringCreateDto(localizationStringCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Create a localization string



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
	localizationStringCreateDto := *openapiclient.NewLocalizationStringCreateDto("Base_example") // LocalizationStringCreateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalizationStringsAPI.CreateLocalizationStringAsync(context.Background()).TenantId(tenantId).LocalizationStringCreateDto(localizationStringCreateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationStringsAPI.CreateLocalizationStringAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLocalizationStringAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LocalizationStringsAPI.CreateLocalizationStringAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocalizationStringAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **localizationStringCreateDto** | [**LocalizationStringCreateDto**](LocalizationStringCreateDto.md) |  | 
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


## DeleteLocalizationStringAsync

> EmptyEnvelope DeleteLocalizationStringAsync(ctx, localizationStringId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a localization string



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
	localizationStringId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalizationStringsAPI.DeleteLocalizationStringAsync(context.Background(), localizationStringId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationStringsAPI.DeleteLocalizationStringAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLocalizationStringAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LocalizationStringsAPI.DeleteLocalizationStringAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localizationStringId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalizationStringAsyncRequest struct via the builder pattern


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


## GetLocalizationStringByIdAsync

> LocalizationStringDtoEnvelope GetLocalizationStringByIdAsync(ctx, localizationStringId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get localization string by ID



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
	localizationStringId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalizationStringsAPI.GetLocalizationStringByIdAsync(context.Background(), localizationStringId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationStringsAPI.GetLocalizationStringByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationStringByIdAsync`: LocalizationStringDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LocalizationStringsAPI.GetLocalizationStringByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localizationStringId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationStringByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**LocalizationStringDtoEnvelope**](LocalizationStringDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationStringsAsync

> LocalizationStringDtoListEnvelope GetLocalizationStringsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get localization strings



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
	resp, r, err := apiClient.LocalizationStringsAPI.GetLocalizationStringsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationStringsAPI.GetLocalizationStringsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationStringsAsync`: LocalizationStringDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LocalizationStringsAPI.GetLocalizationStringsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationStringsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**LocalizationStringDtoListEnvelope**](LocalizationStringDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocalizationStringAsync

> EmptyEnvelope UpdateLocalizationStringAsync(ctx, localizationStringId).TenantId(tenantId).LocalizationStringUpdateDto(localizationStringUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Update a localization string



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
	localizationStringId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	localizationStringUpdateDto := *openapiclient.NewLocalizationStringUpdateDto() // LocalizationStringUpdateDto | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalizationStringsAPI.UpdateLocalizationStringAsync(context.Background(), localizationStringId).TenantId(tenantId).LocalizationStringUpdateDto(localizationStringUpdateDto).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationStringsAPI.UpdateLocalizationStringAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLocalizationStringAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LocalizationStringsAPI.UpdateLocalizationStringAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**localizationStringId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocalizationStringAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **localizationStringUpdateDto** | [**LocalizationStringUpdateDto**](LocalizationStringUpdateDto.md) |  | 
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

