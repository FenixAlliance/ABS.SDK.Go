# \RadzenEditorAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RadzenUploadImage**](RadzenEditorAPI.md#RadzenUploadImage) | **Post** /api/v2/fs/radzen/tenants/{tenantId}/upload/image | Upload an editor image to tenant storage.
[**RadzenUploadImageScoped**](RadzenEditorAPI.md#RadzenUploadImageScoped) | **Post** /api/v2/fs/radzen/tenants/{tenantId}/{recordType}/{recordId}/upload/image | Upload an editor image scoped to a record.
[**RadzenUploadSingle**](RadzenEditorAPI.md#RadzenUploadSingle) | **Post** /api/v2/fs/radzen/tenants/{tenantId}/upload/single | Upload a single editor file to tenant storage.
[**RadzenUploadSingleScoped**](RadzenEditorAPI.md#RadzenUploadSingleScoped) | **Post** /api/v2/fs/radzen/tenants/{tenantId}/{recordType}/{recordId}/upload/single | Upload a single editor file scoped to a record.
[**RadzenUploadStream**](RadzenEditorAPI.md#RadzenUploadStream) | **Put** /api/v2/fs/radzen/tenants/{tenantId}/upload/stream | Chunked editor upload (not implemented).
[**RadzenUploadStreamScoped**](RadzenEditorAPI.md#RadzenUploadStreamScoped) | **Put** /api/v2/fs/radzen/tenants/{tenantId}/{recordType}/{recordId}/upload/stream | Chunked editor upload scoped to a record (not implemented).
[**RadzenUploadUserImage**](RadzenEditorAPI.md#RadzenUploadUserImage) | **Post** /api/v2/fs/radzen/users/upload/image | Upload an editor image to user storage.
[**RadzenUploadUserImageScoped**](RadzenEditorAPI.md#RadzenUploadUserImageScoped) | **Post** /api/v2/fs/radzen/users/{recordType}/{recordId}/upload/image | Upload a user editor image scoped to a record.



## RadzenUploadImage

> RadzenUploadImage(ctx, tenantId).Visibility(visibility).SocialProfileId(socialProfileId).Purpose(purpose).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()

Upload an editor image to tenant storage.

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
	visibility := "visibility_example" // string |  (optional)
	socialProfileId := "socialProfileId_example" // string |  (optional)
	purpose := "purpose_example" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.RadzenUploadImage(context.Background(), tenantId).Visibility(visibility).SocialProfileId(socialProfileId).Purpose(purpose).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.RadzenUploadImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadzenUploadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibility** | **string** |  | 
 **socialProfileId** | **string** |  | 
 **purpose** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadzenUploadImageScoped

> RadzenUploadImageScoped(ctx, tenantId, recordType, recordId).Visibility(visibility).SocialProfileId(socialProfileId).Purpose(purpose).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()

Upload an editor image scoped to a record.

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
	recordType := "recordType_example" // string | 
	recordId := "recordId_example" // string | 
	visibility := "visibility_example" // string |  (optional)
	socialProfileId := "socialProfileId_example" // string |  (optional)
	purpose := "purpose_example" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.RadzenUploadImageScoped(context.Background(), tenantId, recordType, recordId).Visibility(visibility).SocialProfileId(socialProfileId).Purpose(purpose).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.RadzenUploadImageScoped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**recordType** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadzenUploadImageScopedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **visibility** | **string** |  | 
 **socialProfileId** | **string** |  | 
 **purpose** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadzenUploadSingle

> RadzenUploadSingle(ctx, tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()

Upload a single editor file to tenant storage.

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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.RadzenUploadSingle(context.Background(), tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.RadzenUploadSingle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadzenUploadSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadzenUploadSingleScoped

> RadzenUploadSingleScoped(ctx, tenantId, recordType, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()

Upload a single editor file scoped to a record.

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
	recordType := "recordType_example" // string | 
	recordId := "recordId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.RadzenUploadSingleScoped(context.Background(), tenantId, recordType, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.RadzenUploadSingleScoped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**recordType** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadzenUploadSingleScopedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadzenUploadStream

> RadzenUploadStream(ctx, tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Chunked editor upload (not implemented).

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
	tenantId := "tenantId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.RadzenUploadStream(context.Background(), tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.RadzenUploadStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadzenUploadStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadzenUploadStreamScoped

> RadzenUploadStreamScoped(ctx, tenantId, recordType, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Chunked editor upload scoped to a record (not implemented).

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
	tenantId := "tenantId_example" // string | 
	recordType := "recordType_example" // string | 
	recordId := "recordId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.RadzenUploadStreamScoped(context.Background(), tenantId, recordType, recordId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.RadzenUploadStreamScoped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**recordType** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadzenUploadStreamScopedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadzenUploadUserImage

> RadzenUploadUserImage(ctx).Visibility(visibility).SocialProfileId(socialProfileId).Purpose(purpose).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()

Upload an editor image to user storage.

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
	visibility := "visibility_example" // string |  (optional)
	socialProfileId := "socialProfileId_example" // string |  (optional)
	purpose := "purpose_example" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.RadzenUploadUserImage(context.Background()).Visibility(visibility).SocialProfileId(socialProfileId).Purpose(purpose).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.RadzenUploadUserImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRadzenUploadUserImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visibility** | **string** |  | 
 **socialProfileId** | **string** |  | 
 **purpose** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadzenUploadUserImageScoped

> RadzenUploadUserImageScoped(ctx, recordType, recordId).Visibility(visibility).SocialProfileId(socialProfileId).Purpose(purpose).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()

Upload a user editor image scoped to a record.

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
	recordType := "recordType_example" // string | 
	recordId := "recordId_example" // string | 
	visibility := "visibility_example" // string |  (optional)
	socialProfileId := "socialProfileId_example" // string |  (optional)
	purpose := "purpose_example" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.RadzenUploadUserImageScoped(context.Background(), recordType, recordId).Visibility(visibility).SocialProfileId(socialProfileId).Purpose(purpose).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.RadzenUploadUserImageScoped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordType** | **string** |  | 
**recordId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadzenUploadUserImageScopedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **visibility** | **string** |  | 
 **socialProfileId** | **string** |  | 
 **purpose** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

