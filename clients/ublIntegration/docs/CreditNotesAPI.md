# \CreditNotesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2UblServiceCreditNotesInvoiceIdGet**](CreditNotesAPI.md#ApiV2UblServiceCreditNotesInvoiceIdGet) | **Get** /api/v2/UblService/CreditNotes/{invoiceId} | 



## ApiV2UblServiceCreditNotesInvoiceIdGet

> ApiV2UblServiceCreditNotesInvoiceIdGet(ctx, invoiceId).TenantId(tenantId).Profile(profile).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	profile := "profile_example" // string |  (optional) (default to "Generic")
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CreditNotesAPI.ApiV2UblServiceCreditNotesInvoiceIdGet(context.Background(), invoiceId).TenantId(tenantId).Profile(profile).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.ApiV2UblServiceCreditNotesInvoiceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2UblServiceCreditNotesInvoiceIdGetRequest struct via the builder pattern


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

