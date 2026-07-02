# \CurriculumExperiencesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCurriculumExperienceAsync**](CurriculumExperiencesAPI.md#CreateCurriculumExperienceAsync) | **Post** /api/v2/SocialService/Curriculums/{curriculumId}/Experiences | Create a curriculum experience
[**DeleteCurriculumExperienceAsync**](CurriculumExperiencesAPI.md#DeleteCurriculumExperienceAsync) | **Delete** /api/v2/SocialService/Curriculums/{curriculumId}/Experiences/{experienceId} | Delete a curriculum experience
[**GetCurriculumExperienceAsync**](CurriculumExperiencesAPI.md#GetCurriculumExperienceAsync) | **Get** /api/v2/SocialService/Curriculums/{curriculumId}/Experiences/{experienceId} | Get curriculum experience by ID
[**GetCurriculumExperiencesAsync**](CurriculumExperiencesAPI.md#GetCurriculumExperiencesAsync) | **Get** /api/v2/SocialService/Curriculums/{curriculumId}/Experiences | Get curriculum experiences
[**GetCurriculumExperiencesCountAsync**](CurriculumExperiencesAPI.md#GetCurriculumExperiencesCountAsync) | **Get** /api/v2/SocialService/Curriculums/{curriculumId}/Experiences/Count | Count curriculum experiences
[**PatchCurriculumExperienceAsync**](CurriculumExperiencesAPI.md#PatchCurriculumExperienceAsync) | **Patch** /api/v2/SocialService/Curriculums/{curriculumId}/Experiences/{experienceId} | Patch a curriculum experience
[**UpdateCurriculumExperienceAsync**](CurriculumExperiencesAPI.md#UpdateCurriculumExperienceAsync) | **Put** /api/v2/SocialService/Curriculums/{curriculumId}/Experiences/{experienceId} | Update a curriculum experience



## CreateCurriculumExperienceAsync

> EmptyEnvelope CreateCurriculumExperienceAsync(ctx, curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CurriculumExperienceCreateDto(curriculumExperienceCreateDto).Execute()

Create a curriculum experience



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	curriculumExperienceCreateDto := *openapiclient.NewCurriculumExperienceCreateDto() // CurriculumExperienceCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumExperiencesAPI.CreateCurriculumExperienceAsync(context.Background(), curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CurriculumExperienceCreateDto(curriculumExperienceCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumExperiencesAPI.CreateCurriculumExperienceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCurriculumExperienceAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumExperiencesAPI.CreateCurriculumExperienceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCurriculumExperienceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **socialProfileId** | **string** |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **curriculumExperienceCreateDto** | [**CurriculumExperienceCreateDto**](CurriculumExperienceCreateDto.md) |  | 

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


## DeleteCurriculumExperienceAsync

> EmptyEnvelope DeleteCurriculumExperienceAsync(ctx, curriculumId, experienceId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a curriculum experience



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumExperiencesAPI.DeleteCurriculumExperienceAsync(context.Background(), curriculumId, experienceId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumExperiencesAPI.DeleteCurriculumExperienceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCurriculumExperienceAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumExperiencesAPI.DeleteCurriculumExperienceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCurriculumExperienceAsyncRequest struct via the builder pattern


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


## GetCurriculumExperienceAsync

> CurriculumExperienceDtoEnvelope GetCurriculumExperienceAsync(ctx, curriculumId, experienceId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get curriculum experience by ID



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumExperiencesAPI.GetCurriculumExperienceAsync(context.Background(), curriculumId, experienceId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumExperiencesAPI.GetCurriculumExperienceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurriculumExperienceAsync`: CurriculumExperienceDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumExperiencesAPI.GetCurriculumExperienceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurriculumExperienceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **socialProfileId** | **string** |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CurriculumExperienceDtoEnvelope**](CurriculumExperienceDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurriculumExperiencesAsync

> CurriculumExperienceDtoListEnvelope GetCurriculumExperiencesAsync(ctx, curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get curriculum experiences



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumExperiencesAPI.GetCurriculumExperiencesAsync(context.Background(), curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumExperiencesAPI.GetCurriculumExperiencesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurriculumExperiencesAsync`: CurriculumExperienceDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumExperiencesAPI.GetCurriculumExperiencesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurriculumExperiencesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **socialProfileId** | **string** |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CurriculumExperienceDtoListEnvelope**](CurriculumExperienceDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurriculumExperiencesCountAsync

> Int32Envelope GetCurriculumExperiencesCountAsync(ctx, curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count curriculum experiences



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumExperiencesAPI.GetCurriculumExperiencesCountAsync(context.Background(), curriculumId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumExperiencesAPI.GetCurriculumExperiencesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurriculumExperiencesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumExperiencesAPI.GetCurriculumExperiencesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurriculumExperiencesCountAsyncRequest struct via the builder pattern


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


## PatchCurriculumExperienceAsync

> EmptyEnvelope PatchCurriculumExperienceAsync(ctx, curriculumId, experienceId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a curriculum experience



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumExperiencesAPI.PatchCurriculumExperienceAsync(context.Background(), curriculumId, experienceId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumExperiencesAPI.PatchCurriculumExperienceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCurriculumExperienceAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumExperiencesAPI.PatchCurriculumExperienceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCurriculumExperienceAsyncRequest struct via the builder pattern


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


## UpdateCurriculumExperienceAsync

> EmptyEnvelope UpdateCurriculumExperienceAsync(ctx, curriculumId, experienceId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CurriculumExperienceUpdateDto(curriculumExperienceUpdateDto).Execute()

Update a curriculum experience



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
	curriculumId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	experienceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	socialProfileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	curriculumExperienceUpdateDto := *openapiclient.NewCurriculumExperienceUpdateDto() // CurriculumExperienceUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CurriculumExperiencesAPI.UpdateCurriculumExperienceAsync(context.Background(), curriculumId, experienceId).SocialProfileId(socialProfileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).CurriculumExperienceUpdateDto(curriculumExperienceUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CurriculumExperiencesAPI.UpdateCurriculumExperienceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurriculumExperienceAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CurriculumExperiencesAPI.UpdateCurriculumExperienceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**curriculumId** | **string** |  | 
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurriculumExperienceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **socialProfileId** | **string** |  | 
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **curriculumExperienceUpdateDto** | [**CurriculumExperienceUpdateDto**](CurriculumExperienceUpdateDto.md) |  | 

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

