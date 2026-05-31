# \UploadsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SaveFileAsync**](UploadsAPI.md#SaveFileAsync) | **Post** /api/v2/StorageService/Uploads | Upload a file



## SaveFileAsync

> EmptyEnvelope SaveFileAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Notes(notes).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadId(parentFileUploadId).FilePath(filePath).AppFileContent(appFileContent).AppFileSha256(appFileSha256).AppFileCreatedAtUtc(appFileCreatedAtUtc).AppFileUserIdValue(appFileUserIdValue).AppFileTenantIdValue(appFileTenantIdValue).AppFileEnrollmentIdValue(appFileEnrollmentIdValue).AppFileSource(appFileSource).AppFileLength(appFileLength).AppFileName(appFileName).AppFileFileName(appFileFileName).AppFileLastModified(appFileLastModified).AppFileSize(appFileSize).AppFileContentType(appFileContentType).AppFileContentDisposition(appFileContentDisposition).AppFileHeaders(appFileHeaders).Id(id).Timestamp(timestamp).Execute()

Upload a file



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
	resp, r, err := apiClient.UploadsAPI.SaveFileAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Notes(notes).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadId(parentFileUploadId).FilePath(filePath).AppFileContent(appFileContent).AppFileSha256(appFileSha256).AppFileCreatedAtUtc(appFileCreatedAtUtc).AppFileUserIdValue(appFileUserIdValue).AppFileTenantIdValue(appFileTenantIdValue).AppFileEnrollmentIdValue(appFileEnrollmentIdValue).AppFileSource(appFileSource).AppFileLength(appFileLength).AppFileName(appFileName).AppFileFileName(appFileFileName).AppFileLastModified(appFileLastModified).AppFileSize(appFileSize).AppFileContentType(appFileContentType).AppFileContentDisposition(appFileContentDisposition).AppFileHeaders(appFileHeaders).Id(id).Timestamp(timestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadsAPI.SaveFileAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveFileAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UploadsAPI.SaveFileAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveFileAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
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
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

