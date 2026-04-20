# \TrackingPixelsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrackingPixelAsync**](TrackingPixelsAPI.md#GetTrackingPixelAsync) | **Get** /api/v2/MarketingService/TrackingPixels/{pixelId} | Get a tracking pixel



## GetTrackingPixelAsync

> OrderDtoEnvelope GetTrackingPixelAsync(ctx, pixelId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get a tracking pixel



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
	pixelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackingPixelsAPI.GetTrackingPixelAsync(context.Background(), pixelId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingPixelsAPI.GetTrackingPixelAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrackingPixelAsync`: OrderDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrackingPixelsAPI.GetTrackingPixelAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pixelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackingPixelAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**OrderDtoEnvelope**](OrderDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

