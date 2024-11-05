# \MigrationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SystemServiceMigrationsGet**](MigrationsAPI.md#ApiV2SystemServiceMigrationsGet) | **Get** /api/v2/SystemService/Migrations | 
[**ApiV2SystemServiceMigrationsMigratePost**](MigrationsAPI.md#ApiV2SystemServiceMigrationsMigratePost) | **Post** /api/v2/SystemService/Migrations/Migrate | 



## ApiV2SystemServiceMigrationsGet

> StringListEnvelope ApiV2SystemServiceMigrationsGet(ctx).Pending(pending).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	pending := true // bool |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationsAPI.ApiV2SystemServiceMigrationsGet(context.Background()).Pending(pending).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationsAPI.ApiV2SystemServiceMigrationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SystemServiceMigrationsGet`: StringListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MigrationsAPI.ApiV2SystemServiceMigrationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SystemServiceMigrationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pending** | **bool** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**StringListEnvelope**](StringListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SystemServiceMigrationsMigratePost

> StringListEnvelope ApiV2SystemServiceMigrationsMigratePost(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationsAPI.ApiV2SystemServiceMigrationsMigratePost(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationsAPI.ApiV2SystemServiceMigrationsMigratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2SystemServiceMigrationsMigratePost`: StringListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `MigrationsAPI.ApiV2SystemServiceMigrationsMigratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SystemServiceMigrationsMigratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**StringListEnvelope**](StringListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

