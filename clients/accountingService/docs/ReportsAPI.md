# \ReportsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrialBalanceAsync**](ReportsAPI.md#GetTrialBalanceAsync) | **Get** /api/v2/AccountingService/Reports/TrialBalance | Trial balance for a fiscal period



## GetTrialBalanceAsync

> TrialBalanceDtoEnvelope GetTrialBalanceAsync(ctx).TenantId(tenantId).FiscalPeriodId(fiscalPeriodId).FinancialBookId(financialBookId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Trial balance for a fiscal period



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
	fiscalPeriodId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	financialBookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	currencyId := "currencyId_example" // string |  (optional) (default to "USD.USA")
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetTrialBalanceAsync(context.Background()).TenantId(tenantId).FiscalPeriodId(fiscalPeriodId).FinancialBookId(financialBookId).CurrencyId(currencyId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetTrialBalanceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrialBalanceAsync`: TrialBalanceDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetTrialBalanceAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialBalanceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **fiscalPeriodId** | **string** |  | 
 **financialBookId** | **string** |  | 
 **currencyId** | **string** |  | [default to &quot;USD.USA&quot;]
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TrialBalanceDtoEnvelope**](TrialBalanceDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

