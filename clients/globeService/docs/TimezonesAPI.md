# \TimezonesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2GlobeServiceTimezonesGet**](TimezonesAPI.md#ApiV2GlobeServiceTimezonesGet) | **Get** /api/v2/GlobeService/Timezones | 
[**ApiV2GlobeServiceTimezonesTimeZoneIdGet**](TimezonesAPI.md#ApiV2GlobeServiceTimezonesTimeZoneIdGet) | **Get** /api/v2/GlobeService/Timezones/{timeZoneId} | 



## ApiV2GlobeServiceTimezonesGet

> TimezoneDtoListEnvelope ApiV2GlobeServiceTimezonesGet(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.TimezonesAPI.ApiV2GlobeServiceTimezonesGet(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezonesAPI.ApiV2GlobeServiceTimezonesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceTimezonesGet`: TimezoneDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TimezonesAPI.ApiV2GlobeServiceTimezonesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceTimezonesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TimezoneDtoListEnvelope**](TimezoneDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceTimezonesTimeZoneIdGet

> TimezoneDtoEnvelope ApiV2GlobeServiceTimezonesTimeZoneIdGet(ctx, timeZoneId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	timeZoneId := "timeZoneId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimezonesAPI.ApiV2GlobeServiceTimezonesTimeZoneIdGet(context.Background(), timeZoneId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimezonesAPI.ApiV2GlobeServiceTimezonesTimeZoneIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceTimezonesTimeZoneIdGet`: TimezoneDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TimezonesAPI.ApiV2GlobeServiceTimezonesTimeZoneIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeZoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceTimezonesTimeZoneIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TimezoneDtoEnvelope**](TimezoneDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

