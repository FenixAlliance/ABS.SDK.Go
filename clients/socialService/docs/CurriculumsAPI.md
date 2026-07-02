# \CurriculumsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCurriculumAsync**](CurriculumsAPI.md#CreateCurriculumAsync) | **Post** /api/v2/SocialService/Curriculums | Create a curriculum
[**DeleteCurriculumAsync**](CurriculumsAPI.md#DeleteCurriculumAsync) | **Delete** /api/v2/SocialService/Curriculums/{curriculumId} | Delete a curriculum
[**GetCurriculumAsync**](CurriculumsAPI.md#GetCurriculumAsync) | **Get** /api/v2/SocialService/Curriculums/{curriculumId} | Get curriculum by ID
[**GetCurriculumsAsync**](CurriculumsAPI.md#GetCurriculumsAsync) | **Get** /api/v2/SocialService/Curriculums | Get curricula
[**GetCurriculumsCountAsync**](CurriculumsAPI.md#GetCurriculumsCountAsync) | **Get** /api/v2/SocialService/Curriculums/Count | Count curricula
[**PatchCurriculumAsync**](CurriculumsAPI.md#PatchCurriculumAsync) | **Patch** /api/v2/SocialService/Curriculums/{curriculumId} | Patch a curriculum
[**UpdateCurriculumAsync**](CurriculumsAPI.md#UpdateCurriculumAsync) | **Put** /api/v2/SocialService/Curriculums/{curriculumId} | Update a curriculum



## CreateCurriculumAsync

> EmptyEnvelope CreateCurriculumAsync(ctx).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CurriculumCreateDto(curriculumCreateDto).Execute()

Create a curriculum



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	curriculumCreateDto := *openapiclient.NewCurriculumCreateDto() // CurriculumCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumsAPI.CreateCurriculumAsync(context.Background()).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CurriculumCreateDto(curriculumCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumsAPI.CreateCurriculumAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCurriculumAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumsAPI.CreateCurriculumAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCurriculumAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **curriculumCreateDto** | [**CurriculumCreateDto**](CurriculumCreateDto.md) |  | 

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


## DeleteCurriculumAsync

> EmptyEnvelope DeleteCurriculumAsync(ctx, curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a curriculum



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumsAPI.DeleteCurriculumAsync(context.Background(), curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumsAPI.DeleteCurriculumAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCurriculumAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumsAPI.DeleteCurriculumAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCurriculumAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

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


## GetCurriculumAsync

> CurriculumDtoEnvelope GetCurriculumAsync(ctx, curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get curriculum by ID



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumsAPI.GetCurriculumAsync(context.Background(), curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumsAPI.GetCurriculumAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurriculumAsync`: CurriculumDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumsAPI.GetCurriculumAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurriculumAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CurriculumDtoEnvelope**](CurriculumDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurriculumsAsync

> CurriculumDtoListEnvelope GetCurriculumsAsync(ctx).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get curricula



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumsAPI.GetCurriculumsAsync(context.Background()).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumsAPI.GetCurriculumsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurriculumsAsync`: CurriculumDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumsAPI.GetCurriculumsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurriculumsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CurriculumDtoListEnvelope**](CurriculumDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurriculumsCountAsync

> Int32Envelope GetCurriculumsCountAsync(ctx).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count curricula



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumsAPI.GetCurriculumsCountAsync(context.Background()).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumsAPI.GetCurriculumsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurriculumsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumsAPI.GetCurriculumsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurriculumsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 
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


## PatchCurriculumAsync

> EmptyEnvelope PatchCurriculumAsync(ctx, curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a curriculum



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumsAPI.PatchCurriculumAsync(context.Background(), curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumsAPI.PatchCurriculumAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCurriculumAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumsAPI.PatchCurriculumAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCurriculumAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **tenantId** | **string** |  | 
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


## UpdateCurriculumAsync

> EmptyEnvelope UpdateCurriculumAsync(ctx, curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CurriculumUpdateDto(curriculumUpdateDto).Execute()

Update a curriculum



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	curriculumUpdateDto := *openapiclient.NewCurriculumUpdateDto() // CurriculumUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumsAPI.UpdateCurriculumAsync(context.Background(), curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CurriculumUpdateDto(curriculumUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumsAPI.UpdateCurriculumAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurriculumAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumsAPI.UpdateCurriculumAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurriculumAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialProfileId** | **string** |  | 

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **curriculumUpdateDto** | [**CurriculumUpdateDto**](CurriculumUpdateDto.md) |  | 

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

