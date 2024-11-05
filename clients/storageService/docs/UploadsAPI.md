# \UploadsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2StorageServiceUploadsPost**](UploadsAPI.md#ApiV2StorageServiceUploadsPost) | **Post** /api/v2/StorageService/Uploads | 



## ApiV2StorageServiceUploadsPost

> EmptyEnvelope ApiV2StorageServiceUploadsPost(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Notes(notes).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadId(parentFileUploadId).FilePath(filePath).File(file).ID(iD).Timestamp(timestamp).Execute()



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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)
	iD := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	timestamp := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadsAPI.ApiV2StorageServiceUploadsPost(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Notes(notes).Title(title).Author(author).IsFolder(isFolder).FileName(fileName).Abstract(abstract).KeyWords(keyWords).ValidResponse(validResponse).ParentFileUploadId(parentFileUploadId).FilePath(filePath).File(file).ID(iD).Timestamp(timestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadsAPI.ApiV2StorageServiceUploadsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2StorageServiceUploadsPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UploadsAPI.ApiV2StorageServiceUploadsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2StorageServiceUploadsPostRequest struct via the builder pattern


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
 **file** | ***os.File** |  | 
 **iD** | **string** |  | 
 **timestamp** | **time.Time** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

