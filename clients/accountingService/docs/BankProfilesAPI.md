# \BankProfilesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBankProfiles**](BankProfilesAPI.md#GetBankProfiles) | **Get** /api/v2/AccountingService/BankProfiles | Get all bank profiles for a tenant
[**GetBankProfilesCount**](BankProfilesAPI.md#GetBankProfilesCount) | **Get** /api/v2/AccountingService/BankProfiles/Count | Get bank profiles count



## GetBankProfiles

> BankProfileDtoListEnvelope GetBankProfiles(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankProfileDtoCollectionQueryParameters(bankProfileDtoCollectionQueryParameters).Execute()

Get all bank profiles for a tenant



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
	bankProfileDtoCollectionQueryParameters := *openapiclient.NewBankProfileDtoCollectionQueryParameters() // BankProfileDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankProfilesAPI.GetBankProfiles(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankProfileDtoCollectionQueryParameters(bankProfileDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankProfilesAPI.GetBankProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankProfiles`: BankProfileDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankProfilesAPI.GetBankProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBankProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankProfileDtoCollectionQueryParameters** | [**BankProfileDtoCollectionQueryParameters**](BankProfileDtoCollectionQueryParameters.md) |  | 

### Return type

[**BankProfileDtoListEnvelope**](BankProfileDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankProfilesCount

> Int32Envelope GetBankProfilesCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankProfileDtoCollectionQueryParameters(bankProfileDtoCollectionQueryParameters).Execute()

Get bank profiles count



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
	bankProfileDtoCollectionQueryParameters := *openapiclient.NewBankProfileDtoCollectionQueryParameters() // BankProfileDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankProfilesAPI.GetBankProfilesCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankProfileDtoCollectionQueryParameters(bankProfileDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankProfilesAPI.GetBankProfilesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankProfilesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BankProfilesAPI.GetBankProfilesCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBankProfilesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankProfileDtoCollectionQueryParameters** | [**BankProfileDtoCollectionQueryParameters**](BankProfileDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

