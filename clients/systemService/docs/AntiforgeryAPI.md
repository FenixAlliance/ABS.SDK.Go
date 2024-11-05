# \AntiforgeryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2SystemServiceAntiforgeryGetAndStoreTokensGet**](AntiforgeryAPI.md#ApiV2SystemServiceAntiforgeryGetAndStoreTokensGet) | **Get** /api/v2/SystemService/Antiforgery/GetAndStoreTokens | 
[**ApiV2SystemServiceAntiforgeryIsRequestValidGet**](AntiforgeryAPI.md#ApiV2SystemServiceAntiforgeryIsRequestValidGet) | **Get** /api/v2/SystemService/Antiforgery/IsRequestValid | 



## ApiV2SystemServiceAntiforgeryGetAndStoreTokensGet

> ApiV2SystemServiceAntiforgeryGetAndStoreTokensGet(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	r, err := apiClient.AntiforgeryAPI.ApiV2SystemServiceAntiforgeryGetAndStoreTokensGet(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntiforgeryAPI.ApiV2SystemServiceAntiforgeryGetAndStoreTokensGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SystemServiceAntiforgeryGetAndStoreTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2SystemServiceAntiforgeryIsRequestValidGet

> ApiV2SystemServiceAntiforgeryIsRequestValidGet(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	r, err := apiClient.AntiforgeryAPI.ApiV2SystemServiceAntiforgeryIsRequestValidGet(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AntiforgeryAPI.ApiV2SystemServiceAntiforgeryIsRequestValidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2SystemServiceAntiforgeryIsRequestValidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

