# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.1.4089
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://fenixalliance.com.co/Support](https://fenixalliance.com.co/Support)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*WalletsAPI* | [**CreateWalletLocationAsync**](docs/WalletsAPI.md#createwalletlocationasync) | **Post** /api/v2/WalletsService/Wallets/{walletId}/Locations | Create Wallet Location
*WalletsAPI* | [**DeleteWalletLocationAsync**](docs/WalletsAPI.md#deletewalletlocationasync) | **Delete** /api/v2/WalletsService/Wallets/{walletId}/Locations/{locationId} | Delete Wallet Location
*WalletsAPI* | [**GetIncomingPaymentsAsync**](docs/WalletsAPI.md#getincomingpaymentsasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Incoming | Get Incoming Payments
*WalletsAPI* | [**GetIncomingPaymentsCountAsync**](docs/WalletsAPI.md#getincomingpaymentscountasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Incoming/Count | Get Incoming Payments Count
*WalletsAPI* | [**GetIncomingWalletInvoicesAsync**](docs/WalletsAPI.md#getincomingwalletinvoicesasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Incoming | Get Incoming Wallet Invoices
*WalletsAPI* | [**GetIncomingWalletInvoicesCountAsync**](docs/WalletsAPI.md#getincomingwalletinvoicescountasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Incoming/Count | Get Incoming Wallet Invoices Count
*WalletsAPI* | [**GetOutgoingPaymentsAsync**](docs/WalletsAPI.md#getoutgoingpaymentsasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Outgoing | Get Outgoing Payments
*WalletsAPI* | [**GetOutgoingPaymentsCountAsync**](docs/WalletsAPI.md#getoutgoingpaymentscountasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Outgoing/Count | Get Outgoing Payments Count
*WalletsAPI* | [**GetOutgoingWalletInvoicesAsync**](docs/WalletsAPI.md#getoutgoingwalletinvoicesasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Outgoing | Get Outgoing Wallet Invoices
*WalletsAPI* | [**GetOutgoingWalletInvoicesCountAsync**](docs/WalletsAPI.md#getoutgoingwalletinvoicescountasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Outgoing/Count | Get Outgoing Wallet Invoices Count
*WalletsAPI* | [**GetWalletDetailsAsync**](docs/WalletsAPI.md#getwalletdetailsasync) | **Get** /api/v2/WalletsService/Wallets/{walletId} | Get Wallet Details
*WalletsAPI* | [**GetWalletExtendedOrdersAsync**](docs/WalletsAPI.md#getwalletextendedordersasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Orders/Extended | Get Wallet Extended Orders
*WalletsAPI* | [**GetWalletInvoicesAsync**](docs/WalletsAPI.md#getwalletinvoicesasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices | Get Wallet Invoices
*WalletsAPI* | [**GetWalletInvoicesCountAsync**](docs/WalletsAPI.md#getwalletinvoicescountasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Invoices/Count | Get Wallet Invoices Count
*WalletsAPI* | [**GetWalletLocationAsync**](docs/WalletsAPI.md#getwalletlocationasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Locations/{locationId} | Get Wallet Location
*WalletsAPI* | [**GetWalletLocationsAsync**](docs/WalletsAPI.md#getwalletlocationsasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Locations | Get Wallet Locations
*WalletsAPI* | [**GetWalletLocationsCountAsync**](docs/WalletsAPI.md#getwalletlocationscountasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Locations/Count | Get Wallet Locations Count
*WalletsAPI* | [**GetWalletOrdersAsync**](docs/WalletsAPI.md#getwalletordersasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Orders | Get Wallet Orders
*WalletsAPI* | [**GetWalletOrdersCountAsync**](docs/WalletsAPI.md#getwalletorderscountasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Orders/Count | Get Wallet Orders Count
*WalletsAPI* | [**GetWalletPaymentsAsync**](docs/WalletsAPI.md#getwalletpaymentsasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments | Get Wallet Payments
*WalletsAPI* | [**GetWalletPaymentsCountAsync**](docs/WalletsAPI.md#getwalletpaymentscountasync) | **Get** /api/v2/WalletsService/Wallets/{walletId}/Payments/Count | Get Wallet Payments Count
*WalletsAPI* | [**UpdateWalletLocationAsync**](docs/WalletsAPI.md#updatewalletlocationasync) | **Put** /api/v2/WalletsService/Wallets/{walletId}/Locations/{locationId} | Update Wallet Location


## Documentation For Models

 - [ContactDto](docs/ContactDto.md)
 - [Currency](docs/Currency.md)
 - [EmptyEnvelope](docs/EmptyEnvelope.md)
 - [ErrorEnvelope](docs/ErrorEnvelope.md)
 - [ExtendedOrderDto](docs/ExtendedOrderDto.md)
 - [ExtendedOrderDtoListEnvelope](docs/ExtendedOrderDtoListEnvelope.md)
 - [Int32Envelope](docs/Int32Envelope.md)
 - [InvoiceDto](docs/InvoiceDto.md)
 - [InvoiceDtoListEnvelope](docs/InvoiceDtoListEnvelope.md)
 - [LocationCreateDto](docs/LocationCreateDto.md)
 - [LocationDto](docs/LocationDto.md)
 - [LocationDtoEnvelope](docs/LocationDtoEnvelope.md)
 - [LocationDtoListEnvelope](docs/LocationDtoListEnvelope.md)
 - [LocationUpdateDto](docs/LocationUpdateDto.md)
 - [Money](docs/Money.md)
 - [OrderDto](docs/OrderDto.md)
 - [OrderDtoListEnvelope](docs/OrderDtoListEnvelope.md)
 - [PaymentDto](docs/PaymentDto.md)
 - [PaymentDtoListEnvelope](docs/PaymentDtoListEnvelope.md)
 - [TenantDto](docs/TenantDto.md)
 - [TenantEnrolmentDto](docs/TenantEnrolmentDto.md)
 - [UserDto](docs/UserDto.md)
 - [WalletDto](docs/WalletDto.md)
 - [WalletDtoEnvelope](docs/WalletDtoEnvelope.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Bearer

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Bearer and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Bearer": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@fenix-alliance.com

