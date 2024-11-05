# \WalletsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWalletLocationAsync**](WalletsAPI.md#CreateWalletLocationAsync) | **Post** /api/v2/WalletsService/Wallets/{walletId}/Locations | Create Wallet Location
[**DeleteWalletLocationAsync**](WalletsAPI.md#DeleteWalletLocationAsync) | **Delete** /api/v2/WalletsService/Wallets/{walletId}/Locations/{locationId} | Delete Wallet Location
[**GetIncomingPaymentsAsync**](WalletsAPI.md#GetIncomingPaymentsAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Incoming | Get Incoming Payments
[**GetIncomingPaymentsCountAsync**](WalletsAPI.md#GetIncomingPaymentsCountAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Incoming/Count | Get Incoming Payments Count
[**GetIncomingWalletInvoicesAsync**](WalletsAPI.md#GetIncomingWalletInvoicesAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Incoming | Get Incoming Wallet Invoices
[**GetIncomingWalletInvoicesCountAsync**](WalletsAPI.md#GetIncomingWalletInvoicesCountAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Incoming/Count | Get Incoming Wallet Invoices Count
[**GetOutgoingPaymentsAsync**](WalletsAPI.md#GetOutgoingPaymentsAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Outgoing | Get Outgoing Payments
[**GetOutgoingPaymentsCountAsync**](WalletsAPI.md#GetOutgoingPaymentsCountAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Outgoing/Count | Get Outgoing Payments Count
[**GetOutgoingWalletInvoicesAsync**](WalletsAPI.md#GetOutgoingWalletInvoicesAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Outgoing | Get Outgoing Wallet Invoices
[**GetOutgoingWalletInvoicesCountAsync**](WalletsAPI.md#GetOutgoingWalletInvoicesCountAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Outgoing/Count | Get Outgoing Wallet Invoices Count
[**GetWalletDetailsAsync**](WalletsAPI.md#GetWalletDetailsAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId} | Get Wallet Details
[**GetWalletExtendedOrdersAsync**](WalletsAPI.md#GetWalletExtendedOrdersAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Orders/Extended | Get Wallet Extended Orders
[**GetWalletInvoicesAsync**](WalletsAPI.md#GetWalletInvoicesAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices | Get Wallet Invoices
[**GetWalletInvoicesCountAsync**](WalletsAPI.md#GetWalletInvoicesCountAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Count | Get Wallet Invoices Count
[**GetWalletLocationAsync**](WalletsAPI.md#GetWalletLocationAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Locations/{locationId} | Get Wallet Location
[**GetWalletLocationsAsync**](WalletsAPI.md#GetWalletLocationsAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Locations | Get Wallet Locations
[**GetWalletLocationsCountAsync**](WalletsAPI.md#GetWalletLocationsCountAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Locations/Count | Get Wallet Locations Count
[**GetWalletOrdersAsync**](WalletsAPI.md#GetWalletOrdersAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Orders | Get Wallet Orders
[**GetWalletOrdersCountAsync**](WalletsAPI.md#GetWalletOrdersCountAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Orders/Count | Get Wallet Orders Count
[**GetWalletPaymentsAsync**](WalletsAPI.md#GetWalletPaymentsAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments | Get Wallet Payments
[**GetWalletPaymentsCountAsync**](WalletsAPI.md#GetWalletPaymentsCountAsync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Count | Get Wallet Payments Count
[**UpdateWalletLocationAsync**](WalletsAPI.md#UpdateWalletLocationAsync) | **Put** /api/v2/WalletsService/Wallets/{walletId}/Locations/{locationId} | Update Wallet Location



## CreateWalletLocationAsync

> EmptyEnvelope CreateWalletLocationAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LocationCreateDto(locationCreateDto).Execute()

Create Wallet Location



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	locationCreateDto := *openapiclient.NewLocationCreateDto() // LocationCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.CreateWalletLocationAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LocationCreateDto(locationCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CreateWalletLocationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWalletLocationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CreateWalletLocationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletLocationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **locationCreateDto** | [**LocationCreateDto**](LocationCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWalletLocationAsync

> EmptyEnvelope DeleteWalletLocationAsync(ctx, walletId, locationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete Wallet Location



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	locationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.DeleteWalletLocationAsync(context.Background(), walletId, locationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.DeleteWalletLocationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWalletLocationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.DeleteWalletLocationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 
**locationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWalletLocationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomingPaymentsAsync

> PaymentDtoListEnvelope GetIncomingPaymentsAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Incoming Payments



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetIncomingPaymentsAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetIncomingPaymentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomingPaymentsAsync`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetIncomingPaymentsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomingPaymentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomingPaymentsCountAsync

> Int32Envelope GetIncomingPaymentsCountAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Incoming Payments Count



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetIncomingPaymentsCountAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetIncomingPaymentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomingPaymentsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetIncomingPaymentsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomingPaymentsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomingWalletInvoicesAsync

> InvoiceDtoListEnvelope GetIncomingWalletInvoicesAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Incoming Wallet Invoices



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetIncomingWalletInvoicesAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetIncomingWalletInvoicesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomingWalletInvoicesAsync`: InvoiceDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetIncomingWalletInvoicesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomingWalletInvoicesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InvoiceDtoListEnvelope**](InvoiceDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncomingWalletInvoicesCountAsync

> Int32Envelope GetIncomingWalletInvoicesCountAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Incoming Wallet Invoices Count



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetIncomingWalletInvoicesCountAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetIncomingWalletInvoicesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIncomingWalletInvoicesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetIncomingWalletInvoicesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncomingWalletInvoicesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutgoingPaymentsAsync

> PaymentDtoListEnvelope GetOutgoingPaymentsAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Outgoing Payments



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetOutgoingPaymentsAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetOutgoingPaymentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutgoingPaymentsAsync`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetOutgoingPaymentsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutgoingPaymentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutgoingPaymentsCountAsync

> Int32Envelope GetOutgoingPaymentsCountAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Outgoing Payments Count



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetOutgoingPaymentsCountAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetOutgoingPaymentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutgoingPaymentsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetOutgoingPaymentsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutgoingPaymentsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutgoingWalletInvoicesAsync

> InvoiceDtoListEnvelope GetOutgoingWalletInvoicesAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Outgoing Wallet Invoices



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetOutgoingWalletInvoicesAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetOutgoingWalletInvoicesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutgoingWalletInvoicesAsync`: InvoiceDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetOutgoingWalletInvoicesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutgoingWalletInvoicesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InvoiceDtoListEnvelope**](InvoiceDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutgoingWalletInvoicesCountAsync

> Int32Envelope GetOutgoingWalletInvoicesCountAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Outgoing Wallet Invoices Count



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetOutgoingWalletInvoicesCountAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetOutgoingWalletInvoicesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutgoingWalletInvoicesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetOutgoingWalletInvoicesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutgoingWalletInvoicesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletDetailsAsync

> WalletDtoEnvelope GetWalletDetailsAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Details



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletDetailsAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletDetailsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletDetailsAsync`: WalletDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletDetailsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletDetailsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WalletDtoEnvelope**](WalletDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletExtendedOrdersAsync

> ExtendedOrderDtoListEnvelope GetWalletExtendedOrdersAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Extended Orders



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletExtendedOrdersAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletExtendedOrdersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletExtendedOrdersAsync`: ExtendedOrderDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletExtendedOrdersAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletExtendedOrdersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ExtendedOrderDtoListEnvelope**](ExtendedOrderDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletInvoicesAsync

> InvoiceDtoListEnvelope GetWalletInvoicesAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Invoices



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletInvoicesAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletInvoicesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletInvoicesAsync`: InvoiceDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletInvoicesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletInvoicesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**InvoiceDtoListEnvelope**](InvoiceDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletInvoicesCountAsync

> Int32Envelope GetWalletInvoicesCountAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Invoices Count



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletInvoicesCountAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletInvoicesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletInvoicesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletInvoicesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletInvoicesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletLocationAsync

> LocationDtoEnvelope GetWalletLocationAsync(ctx, walletId, locationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Location



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	locationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletLocationAsync(context.Background(), walletId, locationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletLocationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletLocationAsync`: LocationDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletLocationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 
**locationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletLocationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**LocationDtoEnvelope**](LocationDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletLocationsAsync

> LocationDtoListEnvelope GetWalletLocationsAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Locations



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletLocationsAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletLocationsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletLocationsAsync`: LocationDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletLocationsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletLocationsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**LocationDtoListEnvelope**](LocationDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletLocationsCountAsync

> Int32Envelope GetWalletLocationsCountAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Locations Count



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletLocationsCountAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletLocationsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletLocationsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletLocationsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletLocationsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletOrdersAsync

> OrderDtoListEnvelope GetWalletOrdersAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Orders



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletOrdersAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletOrdersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletOrdersAsync`: OrderDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletOrdersAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletOrdersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**OrderDtoListEnvelope**](OrderDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletOrdersCountAsync

> Int32Envelope GetWalletOrdersCountAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Orders Count



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletOrdersCountAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletOrdersCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletOrdersCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletOrdersCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletOrdersCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletPaymentsAsync

> PaymentDtoListEnvelope GetWalletPaymentsAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Payments



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletPaymentsAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletPaymentsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletPaymentsAsync`: PaymentDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletPaymentsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletPaymentsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**PaymentDtoListEnvelope**](PaymentDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletPaymentsCountAsync

> Int32Envelope GetWalletPaymentsCountAsync(ctx, walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Wallet Payments Count



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.GetWalletPaymentsCountAsync(context.Background(), walletId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletPaymentsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletPaymentsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletPaymentsCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletPaymentsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWalletLocationAsync

> EmptyEnvelope UpdateWalletLocationAsync(ctx, walletId, locationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LocationUpdateDto(locationUpdateDto).Execute()

Update Wallet Location



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
	walletId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	locationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	locationUpdateDto := *openapiclient.NewLocationUpdateDto() // LocationUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.UpdateWalletLocationAsync(context.Background(), walletId, locationId).ApiVersion(apiVersion).XApiVersion(xApiVersion).LocationUpdateDto(locationUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.UpdateWalletLocationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWalletLocationAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.UpdateWalletLocationAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**walletId** | **string** |  | 
**locationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletLocationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **locationUpdateDto** | [**LocationUpdateDto**](LocationUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

