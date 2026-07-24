# \DespatchAdvicesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2UblServiceDespatchAdvicesShipmentIdGet**](DespatchAdvicesAPI.md#ApiV2UblServiceDespatchAdvicesShipmentIdGet) | **Get** /api/v2/UblService/DespatchAdvices/{shipmentId} | 



## ApiV2UblServiceDespatchAdvicesShipmentIdGet

> ApiV2UblServiceDespatchAdvicesShipmentIdGet(ctx, shipmentId).TenantId(tenantId).Profile(profile).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	shipmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	profile := "profile_example" // string |  (optional) (default to "Generic")
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DespatchAdvicesAPI.ApiV2UblServiceDespatchAdvicesShipmentIdGet(context.Background(), shipmentId).TenantId(tenantId).Profile(profile).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DespatchAdvicesAPI.ApiV2UblServiceDespatchAdvicesShipmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2UblServiceDespatchAdvicesShipmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **profile** | **string** |  | [default to &quot;Generic&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

