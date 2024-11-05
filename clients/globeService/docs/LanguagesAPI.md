# \LanguagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2GlobeServiceLanguagesGet**](LanguagesAPI.md#ApiV2GlobeServiceLanguagesGet) | **Get** /api/v2/GlobeService/Languages | 
[**ApiV2GlobeServiceLanguagesLanguageIdGet**](LanguagesAPI.md#ApiV2GlobeServiceLanguagesLanguageIdGet) | **Get** /api/v2/GlobeService/Languages/{languageId} | 



## ApiV2GlobeServiceLanguagesGet

> CountryLanguageDtoListEnvelope ApiV2GlobeServiceLanguagesGet(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.LanguagesAPI.ApiV2GlobeServiceLanguagesGet(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguagesAPI.ApiV2GlobeServiceLanguagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceLanguagesGet`: CountryLanguageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LanguagesAPI.ApiV2GlobeServiceLanguagesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceLanguagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryLanguageDtoListEnvelope**](CountryLanguageDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceLanguagesLanguageIdGet

> CountryLanguageDtoEnvelope ApiV2GlobeServiceLanguagesLanguageIdGet(ctx, languageId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	languageId := "languageId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguagesAPI.ApiV2GlobeServiceLanguagesLanguageIdGet(context.Background(), languageId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguagesAPI.ApiV2GlobeServiceLanguagesLanguageIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceLanguagesLanguageIdGet`: CountryLanguageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `LanguagesAPI.ApiV2GlobeServiceLanguagesLanguageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**languageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceLanguagesLanguageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryLanguageDtoEnvelope**](CountryLanguageDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

