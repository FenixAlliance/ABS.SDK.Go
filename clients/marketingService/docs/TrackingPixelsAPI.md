# \TrackingPixelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2MarketingServiceTrackingPixelsPixelIdGet**](TrackingPixelsAPI.md#ApiV2MarketingServiceTrackingPixelsPixelIdGet) | **Get** /api/v2/MarketingService/TrackingPixels/{pixelId} | 



## ApiV2MarketingServiceTrackingPixelsPixelIdGet

> OrderDtoEnvelope ApiV2MarketingServiceTrackingPixelsPixelIdGet(ctx, pixelId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	resp, r, err := apiClient.TrackingPixelsAPI.ApiV2MarketingServiceTrackingPixelsPixelIdGet(context.Background(), pixelId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackingPixelsAPI.ApiV2MarketingServiceTrackingPixelsPixelIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2MarketingServiceTrackingPixelsPixelIdGet`: OrderDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrackingPixelsAPI.ApiV2MarketingServiceTrackingPixelsPixelIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pixelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2MarketingServiceTrackingPixelsPixelIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**OrderDtoEnvelope**](OrderDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

