# \BankingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBank**](BankingAPI.md#CreateBank) | **Post** /api/v2/AccountingService/Banking | Creates a new bank
[**CreateBankAccount**](BankingAPI.md#CreateBankAccount) | **Post** /api/v2/AccountingService/Banking/{bankId}/Accounts | Creates a new bank account
[**CreateBankGuarantee**](BankingAPI.md#CreateBankGuarantee) | **Post** /api/v2/AccountingService/Banking/{bankId}/Guarantees | Creates a new bank guarantee
[**CreateBankTransaction**](BankingAPI.md#CreateBankTransaction) | **Post** /api/v2/AccountingService/Banking/{bankId}/Transactions | Creates a new bank transaction
[**DeleteBank**](BankingAPI.md#DeleteBank) | **Delete** /api/v2/AccountingService/Banking/{bankId} | Deletes a bank
[**DeleteBankAccount**](BankingAPI.md#DeleteBankAccount) | **Delete** /api/v2/AccountingService/Banking/{bankId}/Accounts/{accountId} | Deletes a bank account
[**DeleteBankGuarantee**](BankingAPI.md#DeleteBankGuarantee) | **Delete** /api/v2/AccountingService/Banking/{bankId}/Guarantees/{guaranteeId} | Deletes a bank guarantee
[**DeleteBankTransaction**](BankingAPI.md#DeleteBankTransaction) | **Delete** /api/v2/AccountingService/Banking/{bankId}/Transactions/{transactionId} | Deletes a bank transaction
[**GetBank**](BankingAPI.md#GetBank) | **Get** /api/v2/AccountingService/Banking/{bankId} | Gets the current tenant bank
[**GetBankAccount**](BankingAPI.md#GetBankAccount) | **Get** /api/v2/AccountingService/Banking/{bankId}/Accounts/{accountId} | Gets the current tenant bank account
[**GetBankAccounts**](BankingAPI.md#GetBankAccounts) | **Get** /api/v2/AccountingService/Banking/{bankId}/Accounts | Gets the current tenant bank accounts
[**GetBankAccountsCount**](BankingAPI.md#GetBankAccountsCount) | **Get** /api/v2/AccountingService/Banking/{bankId}/Accounts/Count | Gets the current tenant bank accounts count
[**GetBankGuarantee**](BankingAPI.md#GetBankGuarantee) | **Get** /api/v2/AccountingService/Banking/{bankId}/Guarantees/{guaranteeId} | Gets the current tenant bank guarantee
[**GetBankGuarantees**](BankingAPI.md#GetBankGuarantees) | **Get** /api/v2/AccountingService/Banking/{bankId}/Guarantees | Gets the current tenant bank guarantees
[**GetBankGuaranteesCount**](BankingAPI.md#GetBankGuaranteesCount) | **Get** /api/v2/AccountingService/Banking/{bankId}/Guarantees/Count | Gets the current tenant bank guarantees count
[**GetBankTransaction**](BankingAPI.md#GetBankTransaction) | **Get** /api/v2/AccountingService/Banking/{bankId}/Transactions/{transactionId} | Gets the current tenant bank transaction
[**GetBankTransactions**](BankingAPI.md#GetBankTransactions) | **Get** /api/v2/AccountingService/Banking/{bankId}/Transactions | Gets the current tenant bank transactions
[**GetBankTransactionsCount**](BankingAPI.md#GetBankTransactionsCount) | **Get** /api/v2/AccountingService/Banking/{bankId}/Transactions/Count | Gets the current tenant bank transactions count
[**GetBanks**](BankingAPI.md#GetBanks) | **Get** /api/v2/AccountingService/Banking | Gets the current tenant banks
[**GetBanksCount**](BankingAPI.md#GetBanksCount) | **Get** /api/v2/AccountingService/Banking/Count | Gets the current tenant banks count
[**UpdateBank**](BankingAPI.md#UpdateBank) | **Put** /api/v2/AccountingService/Banking/{bankId} | Updates a bank
[**UpdateBankAccount**](BankingAPI.md#UpdateBankAccount) | **Put** /api/v2/AccountingService/Banking/{bankId}/Accounts/{accountId} | Updates a bank account
[**UpdateBankGuarantee**](BankingAPI.md#UpdateBankGuarantee) | **Put** /api/v2/AccountingService/Banking/{bankId}/Guarantees/{guaranteeId} | Updates a bank guarantee
[**UpdateBankTransaction**](BankingAPI.md#UpdateBankTransaction) | **Put** /api/v2/AccountingService/Banking/{bankId}/Transactions/{transactionId} | Updates a bank transaction



## CreateBank

> BankDtoEnvelope CreateBank(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankCreateDto(bankCreateDto).Execute()

Creates a new bank



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
	bankCreateDto := *openapiclient.NewBankCreateDto() // BankCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.CreateBank(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankCreateDto(bankCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.CreateBank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBank`: BankDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.CreateBank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankCreateDto** | [**BankCreateDto**](BankCreateDto.md) |  | 

### Return type

[**BankDtoEnvelope**](BankDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankAccount

> BankAccountDtoEnvelope CreateBankAccount(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankAccountCreateDto(bankAccountCreateDto).Execute()

Creates a new bank account



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	bankAccountCreateDto := *openapiclient.NewBankAccountCreateDto("Name_example", "CurrencyId_example", "AccountCategory_example") // BankAccountCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.CreateBankAccount(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankAccountCreateDto(bankAccountCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.CreateBankAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankAccount`: BankAccountDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.CreateBankAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankAccountCreateDto** | [**BankAccountCreateDto**](BankAccountCreateDto.md) |  | 

### Return type

[**BankAccountDtoEnvelope**](BankAccountDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankGuarantee

> BankGuaranteeDtoEnvelope CreateBankGuarantee(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankGuaranteeCreateDto(bankGuaranteeCreateDto).Execute()

Creates a new bank guarantee



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	bankGuaranteeCreateDto := *openapiclient.NewBankGuaranteeCreateDto() // BankGuaranteeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.CreateBankGuarantee(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankGuaranteeCreateDto(bankGuaranteeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.CreateBankGuarantee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankGuarantee`: BankGuaranteeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.CreateBankGuarantee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankGuaranteeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankGuaranteeCreateDto** | [**BankGuaranteeCreateDto**](BankGuaranteeCreateDto.md) |  | 

### Return type

[**BankGuaranteeDtoEnvelope**](BankGuaranteeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransaction

> BankTransactionDtoEnvelope CreateBankTransaction(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankTransactionCreateDto(bankTransactionCreateDto).Execute()

Creates a new bank transaction



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	bankTransactionCreateDto := *openapiclient.NewBankTransactionCreateDto() // BankTransactionCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.CreateBankTransaction(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankTransactionCreateDto(bankTransactionCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.CreateBankTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBankTransaction`: BankTransactionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.CreateBankTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBankTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankTransactionCreateDto** | [**BankTransactionCreateDto**](BankTransactionCreateDto.md) |  | 

### Return type

[**BankTransactionDtoEnvelope**](BankTransactionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBank

> DeleteBank(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a bank



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BankingAPI.DeleteBank(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.DeleteBank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBankAccount

> DeleteBankAccount(ctx, bankId, accountId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a bank account



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BankingAPI.DeleteBankAccount(context.Background(), bankId, accountId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.DeleteBankAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBankAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBankGuarantee

> DeleteBankGuarantee(ctx, bankId, guaranteeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a bank guarantee



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	guaranteeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BankingAPI.DeleteBankGuarantee(context.Background(), bankId, guaranteeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.DeleteBankGuarantee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**guaranteeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBankGuaranteeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBankTransaction

> DeleteBankTransaction(ctx, bankId, transactionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Deletes a bank transaction



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BankingAPI.DeleteBankTransaction(context.Background(), bankId, transactionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.DeleteBankTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBankTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBank

> BankDtoEnvelope GetBank(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBank(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBank`: BankDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBank`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BankDtoEnvelope**](BankDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankAccount

> BankAccountDtoEnvelope GetBankAccount(ctx, bankId, accountId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank account



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankAccount(context.Background(), bankId, accountId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankAccount`: BankAccountDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BankAccountDtoEnvelope**](BankAccountDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankAccounts

> BankAccountDtoListEnvelope GetBankAccounts(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank accounts



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankAccounts(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankAccounts`: BankAccountDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BankAccountDtoListEnvelope**](BankAccountDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankAccountsCount

> Int32Envelope GetBankAccountsCount(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank accounts count



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankAccountsCount(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankAccountsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankAccountsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankAccountsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankAccountsCountRequest struct via the builder pattern


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


## GetBankGuarantee

> BankGuaranteeDtoEnvelope GetBankGuarantee(ctx, bankId, guaranteeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank guarantee



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	guaranteeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankGuarantee(context.Background(), bankId, guaranteeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankGuarantee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankGuarantee`: BankGuaranteeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankGuarantee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**guaranteeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankGuaranteeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BankGuaranteeDtoEnvelope**](BankGuaranteeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankGuarantees

> BankGuaranteeDtoListEnvelope GetBankGuarantees(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank guarantees



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankGuarantees(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankGuarantees``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankGuarantees`: BankGuaranteeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankGuarantees`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankGuaranteesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BankGuaranteeDtoListEnvelope**](BankGuaranteeDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankGuaranteesCount

> Int32Envelope GetBankGuaranteesCount(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank guarantees count



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankGuaranteesCount(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankGuaranteesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankGuaranteesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankGuaranteesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankGuaranteesCountRequest struct via the builder pattern


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


## GetBankTransaction

> BankTransactionDtoEnvelope GetBankTransaction(ctx, bankId, transactionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank transaction



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankTransaction(context.Background(), bankId, transactionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransaction`: BankTransactionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BankTransactionDtoEnvelope**](BankTransactionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactions

> BankTransactionDtoListEnvelope GetBankTransactions(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank transactions



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankTransactions(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactions`: BankTransactionDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BankTransactionDtoListEnvelope**](BankTransactionDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionsCount

> Int32Envelope GetBankTransactionsCount(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant bank transactions count



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.GetBankTransactionsCount(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBankTransactionsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactionsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBankTransactionsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionsCountRequest struct via the builder pattern


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


## GetBanks

> BankDtoListEnvelope GetBanks(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant banks



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
	resp, r, err := apiClient.BankingAPI.GetBanks(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBanks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBanks`: BankDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBanks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBanksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**BankDtoListEnvelope**](BankDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBanksCount

> Int32Envelope GetBanksCount(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Gets the current tenant banks count



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
	resp, r, err := apiClient.BankingAPI.GetBanksCount(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.GetBanksCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBanksCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.GetBanksCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBanksCountRequest struct via the builder pattern


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


## UpdateBank

> BankDtoEnvelope UpdateBank(ctx, bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankUpdateDto(bankUpdateDto).Execute()

Updates a bank



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	bankUpdateDto := *openapiclient.NewBankUpdateDto() // BankUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.UpdateBank(context.Background(), bankId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankUpdateDto(bankUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.UpdateBank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBank`: BankDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.UpdateBank`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankUpdateDto** | [**BankUpdateDto**](BankUpdateDto.md) |  | 

### Return type

[**BankDtoEnvelope**](BankDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankAccount

> BankAccountDtoEnvelope UpdateBankAccount(ctx, bankId, accountId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankAccountUpdateDto(bankAccountUpdateDto).Execute()

Updates a bank account



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	bankAccountUpdateDto := *openapiclient.NewBankAccountUpdateDto("Name_example", "CurrencyId_example") // BankAccountUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.UpdateBankAccount(context.Background(), bankId, accountId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankAccountUpdateDto(bankAccountUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.UpdateBankAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBankAccount`: BankAccountDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.UpdateBankAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankAccountUpdateDto** | [**BankAccountUpdateDto**](BankAccountUpdateDto.md) |  | 

### Return type

[**BankAccountDtoEnvelope**](BankAccountDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankGuarantee

> BankGuaranteeDtoEnvelope UpdateBankGuarantee(ctx, bankId, guaranteeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankGuaranteeUpdateDto(bankGuaranteeUpdateDto).Execute()

Updates a bank guarantee



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	guaranteeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	bankGuaranteeUpdateDto := *openapiclient.NewBankGuaranteeUpdateDto() // BankGuaranteeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.UpdateBankGuarantee(context.Background(), bankId, guaranteeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankGuaranteeUpdateDto(bankGuaranteeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.UpdateBankGuarantee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBankGuarantee`: BankGuaranteeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.UpdateBankGuarantee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**guaranteeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankGuaranteeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankGuaranteeUpdateDto** | [**BankGuaranteeUpdateDto**](BankGuaranteeUpdateDto.md) |  | 

### Return type

[**BankGuaranteeDtoEnvelope**](BankGuaranteeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankTransaction

> BankTransactionDtoEnvelope UpdateBankTransaction(ctx, bankId, transactionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankTransactionUpdateDto(bankTransactionUpdateDto).Execute()

Updates a bank transaction



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
	bankId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	transactionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	bankTransactionUpdateDto := *openapiclient.NewBankTransactionUpdateDto() // BankTransactionUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankingAPI.UpdateBankTransaction(context.Background(), bankId, transactionId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).BankTransactionUpdateDto(bankTransactionUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankingAPI.UpdateBankTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBankTransaction`: BankTransactionDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `BankingAPI.UpdateBankTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bankId** | **string** |  | 
**transactionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **bankTransactionUpdateDto** | [**BankTransactionUpdateDto**](BankTransactionUpdateDto.md) |  | 

### Return type

[**BankTransactionDtoEnvelope**](BankTransactionDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

