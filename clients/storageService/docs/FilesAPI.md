# \FilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFileAsync**](FilesAPI.md#CreateFileAsync) | **Post** /api/v2/StorageService/Files | 
[**DeleteFileAsync**](FilesAPI.md#DeleteFileAsync) | **Delete** /api/v2/StorageService/Files/{fileId} | 
[**DownloadFileAsync**](FilesAPI.md#DownloadFileAsync) | **Get** /api/v2/StorageService/Files/{fileId}/Raw | 
[**GetFileAsync**](FilesAPI.md#GetFileAsync) | **Get** /api/v2/StorageService/Files/{fileId} | 
[**GetFilesAsync**](FilesAPI.md#GetFilesAsync) | **Get** /api/v2/StorageService/Files | 
[**UpdateFileAsync**](FilesAPI.md#UpdateFileAsync) | **Put** /api/v2/StorageService/Files/{fileId} | 



## CreateFileAsync

> EmptyEnvelope CreateFileAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Id(id).Timestamp(timestamp).Notes(notes).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadId(parentFileUploadId).FilePath(filePath).File(file).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	timestamp := time.Now() // time.Time |  (optional)
	notes := "notes_example" // string |  (optional)
	title := "title_example" // string |  (optional)
	author := "author_example" // string |  (optional)
	isFolder := true // bool |  (optional)
	fileName := "fileName_example" // string |  (optional)
	abstract := "abstract_example" // string |  (optional)
	keyWords := "keyWords_example" // string |  (optional)
	validResponse := true // bool |  (optional)
	parentFileUploadId := "parentFileUploadId_example" // string |  (optional)
	filePath := "filePath_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.CreateFileAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Id(id).Timestamp(timestamp).Notes(notes).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadId(parentFileUploadId).FilePath(filePath).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.CreateFileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.CreateFileAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **id** | **string** |  | 
 **timestamp** | **time.Time** |  | 
 **notes** | **string** |  | 
 **title** | **string** |  | 
 **author** | **string** |  | 
 **isFolder** | **bool** |  | 
 **fileName** | **string** |  | 
 **abstract** | **string** |  | 
 **keyWords** | **string** |  | 
 **validResponse** | **bool** |  | 
 **parentFileUploadId** | **string** |  | 
 **filePath** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileAsync

> FileUploadDtoEnvelope DeleteFileAsync(ctx, fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.DeleteFileAsync(context.Background(), fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFileAsync`: FileUploadDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteFileAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FileUploadDtoEnvelope**](FileUploadDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFileAsync

> *os.File DownloadFileAsync(ctx, fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.DownloadFileAsync(context.Background(), fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DownloadFileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadFileAsync`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DownloadFileAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileAsync

> FileUploadDtoEnvelope GetFileAsync(ctx, fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileAsync(context.Background(), fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileAsync`: FileUploadDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FileUploadDtoEnvelope**](FileUploadDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesAsync

> FileUploadDtoEnvelope GetFilesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFilesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFilesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesAsync`: FileUploadDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFilesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**FileUploadDtoEnvelope**](FileUploadDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileAsync

> FileUploadDtoEnvelope UpdateFileAsync(ctx, fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Notes(notes).Metadata(metadata).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadID(parentFileUploadID).FilePath(filePath).File(file).Execute()



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	notes := "notes_example" // string |  (optional)
	metadata := "metadata_example" // string |  (optional)
	title := "title_example" // string |  (optional)
	author := "author_example" // string |  (optional)
	isFolder := true // bool |  (optional)
	fileName := "fileName_example" // string |  (optional)
	abstract := "abstract_example" // string |  (optional)
	keyWords := "keyWords_example" // string |  (optional)
	validResponse := true // bool |  (optional)
	parentFileUploadID := "parentFileUploadID_example" // string |  (optional)
	filePath := "filePath_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UpdateFileAsync(context.Background(), fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Notes(notes).Metadata(metadata).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadID(parentFileUploadID).FilePath(filePath).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UpdateFileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileAsync`: FileUploadDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UpdateFileAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **notes** | **string** |  | 
 **metadata** | **string** |  | 
 **title** | **string** |  | 
 **author** | **string** |  | 
 **isFolder** | **bool** |  | 
 **fileName** | **string** |  | 
 **abstract** | **string** |  | 
 **keyWords** | **string** |  | 
 **validResponse** | **bool** |  | 
 **parentFileUploadID** | **string** |  | 
 **filePath** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

[**FileUploadDtoEnvelope**](FileUploadDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

