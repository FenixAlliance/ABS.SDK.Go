# \FilesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFileAsync**](FilesAPI.md#CreateFileAsync) | **Post** /api/v2/StorageService/Files | 
[**DeleteFileAsync**](FilesAPI.md#DeleteFileAsync) | **Delete** /api/v2/StorageService/Files/{fileId} | 
[**DownloadFileAsync**](FilesAPI.md#DownloadFileAsync) | **Get** /api/v2/StorageService/Files/{fileId}/Raw | 
[**GetFileAsync**](FilesAPI.md#GetFileAsync) | **Get** /api/v2/StorageService/Files/{fileId} | 
[**GetFileThumbnailAsync**](FilesAPI.md#GetFileThumbnailAsync) | **Get** /api/v2/StorageService/Files/{fileId}/Thumbnail | 
[**GetFilesAsync**](FilesAPI.md#GetFilesAsync) | **Get** /api/v2/StorageService/Files | 
[**GetFilesCountAsync**](FilesAPI.md#GetFilesCountAsync) | **Get** /api/v2/StorageService/Files/Count | 
[**UpdateFileAsync**](FilesAPI.md#UpdateFileAsync) | **Put** /api/v2/StorageService/Files/{fileId} | 



## CreateFileAsync

> EmptyEnvelope CreateFileAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Notes(notes).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadId(parentFileUploadId).FilePath(filePath).PublicAccessType(publicAccessType).Purpose(purpose).SocialProfileIdValue(socialProfileIdValue).AppFileContent(appFileContent).AppFileSha256(appFileSha256).AppFileCreatedAtUtc(appFileCreatedAtUtc).AppFileUserIdValue(appFileUserIdValue).AppFileTenantIdValue(appFileTenantIdValue).AppFileEnrollmentIdValue(appFileEnrollmentIdValue).AppFileSource(appFileSource).AppFileLength(appFileLength).AppFileName(appFileName).AppFileFileName(appFileFileName).AppFileLastModified(appFileLastModified).AppFileSize(appFileSize).AppFileContentType(appFileContentType).AppFileContentDisposition(appFileContentDisposition).AppFileHeaders(appFileHeaders).Id(id).Timestamp(timestamp).Execute()



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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
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
	publicAccessType := "publicAccessType_example" // string |  (optional)
	purpose := "purpose_example" // string |  (optional)
	socialProfileIdValue := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	appFileContent := string(BYTE_ARRAY_DATA_HERE) // string |  (optional)
	appFileSha256 := "appFileSha256_example" // string |  (optional)
	appFileCreatedAtUtc := time.Now() // time.Time |  (optional)
	appFileUserIdValue := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	appFileTenantIdValue := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	appFileEnrollmentIdValue := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	appFileSource := "appFileSource_example" // string |  (optional)
	appFileLength := int64(789) // int64 |  (optional)
	appFileName := "appFileName_example" // string |  (optional)
	appFileFileName := "appFileFileName_example" // string |  (optional)
	appFileLastModified := time.Now() // time.Time |  (optional)
	appFileSize := int64(789) // int64 |  (optional)
	appFileContentType := "appFileContentType_example" // string |  (optional)
	appFileContentDisposition := "appFileContentDisposition_example" // string |  (optional)
	appFileHeaders := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	timestamp := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.CreateFileAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Notes(notes).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadId(parentFileUploadId).FilePath(filePath).PublicAccessType(publicAccessType).Purpose(purpose).SocialProfileIdValue(socialProfileIdValue).AppFileContent(appFileContent).AppFileSha256(appFileSha256).AppFileCreatedAtUtc(appFileCreatedAtUtc).AppFileUserIdValue(appFileUserIdValue).AppFileTenantIdValue(appFileTenantIdValue).AppFileEnrollmentIdValue(appFileEnrollmentIdValue).AppFileSource(appFileSource).AppFileLength(appFileLength).AppFileName(appFileName).AppFileFileName(appFileFileName).AppFileLastModified(appFileLastModified).AppFileSize(appFileSize).AppFileContentType(appFileContentType).AppFileContentDisposition(appFileContentDisposition).AppFileHeaders(appFileHeaders).Id(id).Timestamp(timestamp).Execute()
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
 **file** | ***os.File** |  | 
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
 **publicAccessType** | **string** |  | 
 **purpose** | **string** |  | 
 **socialProfileIdValue** | **string** |  | 
 **appFileContent** | **string** |  | 
 **appFileSha256** | **string** |  | 
 **appFileCreatedAtUtc** | **time.Time** |  | 
 **appFileUserIdValue** | **string** |  | 
 **appFileTenantIdValue** | **string** |  | 
 **appFileEnrollmentIdValue** | **string** |  | 
 **appFileSource** | **string** |  | 
 **appFileLength** | **int64** |  | 
 **appFileName** | **string** |  | 
 **appFileFileName** | **string** |  | 
 **appFileLastModified** | **time.Time** |  | 
 **appFileSize** | **int64** |  | 
 **appFileContentType** | **string** |  | 
 **appFileContentDisposition** | **string** |  | 
 **appFileHeaders** | **map[string]string** |  | 
 **id** | **string** |  | 
 **timestamp** | **time.Time** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileAsync

> EmptyEnvelope DeleteFileAsync(ctx, fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	// response from `DeleteFileAsync`: EmptyEnvelope
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

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileThumbnailAsync

> *os.File GetFileThumbnailAsync(ctx, fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.FilesAPI.GetFileThumbnailAsync(context.Background(), fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileThumbnailAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileThumbnailAsync`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileThumbnailAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileThumbnailAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFilesCountAsync

> int64 GetFilesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.FilesAPI.GetFilesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFilesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFilesCountAsync`: int64
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFilesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFileAsync

> EmptyEnvelope UpdateFileAsync(ctx, fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Notes(notes).Metadata(metadata).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadID(parentFileUploadID).FilePath(filePath).AppFileContent(appFileContent).AppFileSha256(appFileSha256).AppFileCreatedAtUtc(appFileCreatedAtUtc).AppFileUserIdValue(appFileUserIdValue).AppFileTenantIdValue(appFileTenantIdValue).AppFileEnrollmentIdValue(appFileEnrollmentIdValue).AppFileSource(appFileSource).AppFileLength(appFileLength).AppFileName(appFileName).AppFileFileName(appFileFileName).AppFileLastModified(appFileLastModified).AppFileSize(appFileSize).AppFileContentType(appFileContentType).AppFileContentDisposition(appFileContentDisposition).AppFileHeaders(appFileHeaders).Execute()



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
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
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
	appFileContent := string(BYTE_ARRAY_DATA_HERE) // string |  (optional)
	appFileSha256 := "appFileSha256_example" // string |  (optional)
	appFileCreatedAtUtc := time.Now() // time.Time |  (optional)
	appFileUserIdValue := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	appFileTenantIdValue := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	appFileEnrollmentIdValue := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	appFileSource := "appFileSource_example" // string |  (optional)
	appFileLength := int64(789) // int64 |  (optional)
	appFileName := "appFileName_example" // string |  (optional)
	appFileFileName := "appFileFileName_example" // string |  (optional)
	appFileLastModified := time.Now() // time.Time |  (optional)
	appFileSize := int64(789) // int64 |  (optional)
	appFileContentType := "appFileContentType_example" // string |  (optional)
	appFileContentDisposition := "appFileContentDisposition_example" // string |  (optional)
	appFileHeaders := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UpdateFileAsync(context.Background(), fileId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).File(file).Notes(notes).Metadata(metadata).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadID(parentFileUploadID).FilePath(filePath).AppFileContent(appFileContent).AppFileSha256(appFileSha256).AppFileCreatedAtUtc(appFileCreatedAtUtc).AppFileUserIdValue(appFileUserIdValue).AppFileTenantIdValue(appFileTenantIdValue).AppFileEnrollmentIdValue(appFileEnrollmentIdValue).AppFileSource(appFileSource).AppFileLength(appFileLength).AppFileName(appFileName).AppFileFileName(appFileFileName).AppFileLastModified(appFileLastModified).AppFileSize(appFileSize).AppFileContentType(appFileContentType).AppFileContentDisposition(appFileContentDisposition).AppFileHeaders(appFileHeaders).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UpdateFileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFileAsync`: EmptyEnvelope
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
 **file** | ***os.File** |  | 
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
 **appFileContent** | **string** |  | 
 **appFileSha256** | **string** |  | 
 **appFileCreatedAtUtc** | **time.Time** |  | 
 **appFileUserIdValue** | **string** |  | 
 **appFileTenantIdValue** | **string** |  | 
 **appFileEnrollmentIdValue** | **string** |  | 
 **appFileSource** | **string** |  | 
 **appFileLength** | **int64** |  | 
 **appFileName** | **string** |  | 
 **appFileFileName** | **string** |  | 
 **appFileLastModified** | **time.Time** |  | 
 **appFileSize** | **int64** |  | 
 **appFileContentType** | **string** |  | 
 **appFileContentDisposition** | **string** |  | 
 **appFileHeaders** | **map[string]string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

