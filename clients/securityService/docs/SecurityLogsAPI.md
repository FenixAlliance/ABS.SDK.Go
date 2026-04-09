# \SecurityLogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityLogsAsync**](SecurityLogsAPI.md#GetSecurityLogsAsync) | **Get** /api/v2/SecurityService/SecurityLogs | Get business security logs
[**GetSecurityLogsCountAsync**](SecurityLogsAPI.md#GetSecurityLogsCountAsync) | **Get** /api/v2/SecurityService/SecurityLogs/Count | Get business security logs count



## GetSecurityLogsAsync

> BusinessSecurityLogDtoListEnvelope GetSecurityLogsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get business security logs



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityLogsAPI.GetSecurityLogsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityLogsAPI.GetSecurityLogsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityLogsAsync`: BusinessSecurityLogDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `SecurityLogsAPI.GetSecurityLogsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityLogsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BusinessSecurityLogDtoListEnvelope**](BusinessSecurityLogDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityLogsCountAsync

> Int32Envelope GetSecurityLogsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get business security logs count



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityLogsAPI.GetSecurityLogsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityLogsAPI.GetSecurityLogsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityLogsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `SecurityLogsAPI.GetSecurityLogsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityLogsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

