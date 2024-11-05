# \BlobsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBlobAsync**](BlobsAPI.md#GetBlobAsync) | **Get** /api/v2/StorageService/Blobs/Single | 
[**GetBlobsAsync**](BlobsAPI.md#GetBlobsAsync) | **Get** /api/v2/StorageService/Blobs | 



## GetBlobAsync

> BlobEnvelope GetBlobAsync(ctx).TenantId(tenantId).FilePath(filePath).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	filePath := "filePath_example" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobsAPI.GetBlobAsync(context.Background()).TenantId(tenantId).FilePath(filePath).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobsAPI.GetBlobAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobAsync`: BlobEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlobsAPI.GetBlobAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **filePath** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BlobEnvelope**](BlobEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobsAsync

> BlobEnvelope GetBlobsAsync(ctx).TenantId(tenantId).FolderPath(folderPath).BrowseFilter(browseFilter).FilePrefix(filePrefix).Recurse(recurse).MaxResults(maxResults).IncludeAttributes(includeAttributes).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	folderPath := "folderPath_example" // string |  (optional)
	browseFilter := "browseFilter_example" // string |  (optional)
	filePrefix := "filePrefix_example" // string |  (optional)
	recurse := true // bool |  (optional)
	maxResults := int32(56) // int32 |  (optional)
	includeAttributes := true // bool |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobsAPI.GetBlobsAsync(context.Background()).TenantId(tenantId).FolderPath(folderPath).BrowseFilter(browseFilter).FilePrefix(filePrefix).Recurse(recurse).MaxResults(maxResults).IncludeAttributes(includeAttributes).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobsAPI.GetBlobsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobsAsync`: BlobEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BlobsAPI.GetBlobsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **folderPath** | **string** |  | 
 **browseFilter** | **string** |  | 
 **filePrefix** | **string** |  | 
 **recurse** | **bool** |  | 
 **maxResults** | **int32** |  | 
 **includeAttributes** | **bool** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BlobEnvelope**](BlobEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

