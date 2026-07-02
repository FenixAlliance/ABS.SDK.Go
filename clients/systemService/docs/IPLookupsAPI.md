# \IPLookupsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemIPLookup**](IPLookupsAPI.md#DeleteSystemIPLookup) | **Delete** /api/v2/SystemService/IPLookups/{ipLookupId} | Delete a system IP lookup
[**GetSystemIPLookupById**](IPLookupsAPI.md#GetSystemIPLookupById) | **Get** /api/v2/SystemService/IPLookups/{ipLookupId} | Retrieve a single system IP lookup by its ID
[**GetSystemIPLookups**](IPLookupsAPI.md#GetSystemIPLookups) | **Get** /api/v2/SystemService/IPLookups | Retrieve a list of system IP lookups
[**GetSystemIPLookupsCount**](IPLookupsAPI.md#GetSystemIPLookupsCount) | **Get** /api/v2/SystemService/IPLookups/Count | Get the count of system IP lookups



## DeleteSystemIPLookup

> EmptyEnvelope DeleteSystemIPLookup(ctx, ipLookupId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a system IP lookup



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
	ipLookupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPLookupsAPI.DeleteSystemIPLookup(context.Background(), ipLookupId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPLookupsAPI.DeleteSystemIPLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSystemIPLookup`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `IPLookupsAPI.DeleteSystemIPLookup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipLookupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemIPLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemIPLookupById

> IPLookupDtoEnvelope GetSystemIPLookupById(ctx, ipLookupId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a single system IP lookup by its ID



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
	ipLookupId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPLookupsAPI.GetSystemIPLookupById(context.Background(), ipLookupId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPLookupsAPI.GetSystemIPLookupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemIPLookupById`: IPLookupDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `IPLookupsAPI.GetSystemIPLookupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipLookupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemIPLookupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**IPLookupDtoEnvelope**](IPLookupDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemIPLookups

> IPLookupDtoListEnvelope GetSystemIPLookups(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a list of system IP lookups



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
	resp, r, err := apiClient.IPLookupsAPI.GetSystemIPLookups(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPLookupsAPI.GetSystemIPLookups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemIPLookups`: IPLookupDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `IPLookupsAPI.GetSystemIPLookups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemIPLookupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**IPLookupDtoListEnvelope**](IPLookupDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemIPLookupsCount

> Int32Envelope GetSystemIPLookupsCount(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get the count of system IP lookups



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
	resp, r, err := apiClient.IPLookupsAPI.GetSystemIPLookupsCount(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPLookupsAPI.GetSystemIPLookupsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemIPLookupsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `IPLookupsAPI.GetSystemIPLookupsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemIPLookupsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

