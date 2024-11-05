# \RadzenEditorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2StorageServiceRadzenEditorUploadsIdPost**](RadzenEditorAPI.md#ApiV2StorageServiceRadzenEditorUploadsIdPost) | **Post** /api/v2/StorageService/RadzenEditor/Uploads/{id} | 
[**ApiV2StorageServiceRadzenEditorUploadsImagePost**](RadzenEditorAPI.md#ApiV2StorageServiceRadzenEditorUploadsImagePost) | **Post** /api/v2/StorageService/RadzenEditor/Uploads/Image | 
[**ApiV2StorageServiceRadzenEditorUploadsMultiplePost**](RadzenEditorAPI.md#ApiV2StorageServiceRadzenEditorUploadsMultiplePost) | **Post** /api/v2/StorageService/RadzenEditor/Uploads/Multiple | 
[**ApiV2StorageServiceRadzenEditorUploadsSinglePost**](RadzenEditorAPI.md#ApiV2StorageServiceRadzenEditorUploadsSinglePost) | **Post** /api/v2/StorageService/RadzenEditor/Uploads/Single | 
[**ApiV2StorageServiceRadzenEditorUploadsSpecificPost**](RadzenEditorAPI.md#ApiV2StorageServiceRadzenEditorUploadsSpecificPost) | **Post** /api/v2/StorageService/RadzenEditor/Uploads/Specific | 



## ApiV2StorageServiceRadzenEditorUploadsIdPost

> ApiV2StorageServiceRadzenEditorUploadsIdPost(ctx, id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Files(files).Execute()



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
	id := int32(56) // int32 | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	files := []*os.File{"TODO"} // []*os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsIdPost(context.Background(), id).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Files(files).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2StorageServiceRadzenEditorUploadsIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **files** | **[]*os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2StorageServiceRadzenEditorUploadsImagePost

> ApiV2StorageServiceRadzenEditorUploadsImagePost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsImagePost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsImagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2StorageServiceRadzenEditorUploadsImagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2StorageServiceRadzenEditorUploadsMultiplePost

> ApiV2StorageServiceRadzenEditorUploadsMultiplePost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Files(files).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	files := []*os.File{"TODO"} // []*os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsMultiplePost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Files(files).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsMultiplePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2StorageServiceRadzenEditorUploadsMultiplePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **files** | **[]*os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2StorageServiceRadzenEditorUploadsSinglePost

> ApiV2StorageServiceRadzenEditorUploadsSinglePost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsSinglePost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsSinglePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2StorageServiceRadzenEditorUploadsSinglePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2StorageServiceRadzenEditorUploadsSpecificPost

> ApiV2StorageServiceRadzenEditorUploadsSpecificPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsSpecificPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadzenEditorAPI.ApiV2StorageServiceRadzenEditorUploadsSpecificPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2StorageServiceRadzenEditorUploadsSpecificPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

