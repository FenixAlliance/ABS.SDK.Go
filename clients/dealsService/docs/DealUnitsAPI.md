# \DealUnitsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculateDealUnitAsync**](DealUnitsAPI.md#CalculateDealUnitAsync) | **Put** /api/v2/DealsService/DealUnits/{dealUnitId}/Calculate | Calculate a deal unit
[**CalculateDealUnitLineAsync**](DealUnitsAPI.md#CalculateDealUnitLineAsync) | **Put** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/{dealUnitLineId}/Calculate | Calculate a deal unit line
[**CreateDealUnitAsync**](DealUnitsAPI.md#CreateDealUnitAsync) | **Post** /api/v2/DealsService/DealUnits | Create a deal unit
[**CreateGetDealUnitLinesAsync**](DealUnitsAPI.md#CreateGetDealUnitLinesAsync) | **Post** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines | Create a deal unit line
[**DeleteDealUnitAsync**](DealUnitsAPI.md#DeleteDealUnitAsync) | **Delete** /api/v2/DealsService/DealUnits/{dealUnitId} | Delete a deal unit
[**DeleteDealUnitPriceAsync**](DealUnitsAPI.md#DeleteDealUnitPriceAsync) | **Delete** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/{dealUnitLineId} | Delete a deal unit line
[**GetDealUnitAsync**](DealUnitsAPI.md#GetDealUnitAsync) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId} | Get deal unit by ID
[**GetDealUnitLinesAsync**](DealUnitsAPI.md#GetDealUnitLinesAsync) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines | Get deal unit lines
[**GetDealUnitLinesCountAsync**](DealUnitsAPI.md#GetDealUnitLinesCountAsync) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/Count | Get deal unit lines count
[**GetDealUnitPriceAsync**](DealUnitsAPI.md#GetDealUnitPriceAsync) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/{dealUnitLineId} | Get a deal unit line by ID
[**GetDealUnitsAsync**](DealUnitsAPI.md#GetDealUnitsAsync) | **Get** /api/v2/DealsService/DealUnits | Get deal units
[**GetDealUnitsCountAsync**](DealUnitsAPI.md#GetDealUnitsCountAsync) | **Get** /api/v2/DealsService/DealUnits/Count | Get deal units count
[**GetExtendedDealUnitAsync**](DealUnitsAPI.md#GetExtendedDealUnitAsync) | **Get** /api/v2/DealsService/DealUnits/{dealUnitId}/Extended | Get extended deal unit by ID
[**GetExtendedDealUnitsAsync**](DealUnitsAPI.md#GetExtendedDealUnitsAsync) | **Get** /api/v2/DealsService/DealUnits/Extended | Get extended deal units
[**UpdateDealUnitAsync**](DealUnitsAPI.md#UpdateDealUnitAsync) | **Put** /api/v2/DealsService/DealUnits/{dealUnitId} | Update a deal unit
[**UpdateDealUnitPriceAsync**](DealUnitsAPI.md#UpdateDealUnitPriceAsync) | **Put** /api/v2/DealsService/DealUnits/{dealUnitId}/Lines/{dealUnitLineId} | Update a deal unit line



## CalculateDealUnitAsync

> EmptyEnvelope CalculateDealUnitAsync(ctx, dealUnitId).TenantId(tenantId).Execute()

Calculate a deal unit



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.CalculateDealUnitAsync(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.CalculateDealUnitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateDealUnitAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.CalculateDealUnitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalculateDealUnitAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## CalculateDealUnitLineAsync

> EmptyEnvelope CalculateDealUnitLineAsync(ctx, dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()

Calculate a deal unit line



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.CalculateDealUnitLineAsync(context.Background(), dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.CalculateDealUnitLineAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CalculateDealUnitLineAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.CalculateDealUnitLineAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 
**dealUnitLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCalculateDealUnitLineAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## CreateDealUnitAsync

> EmptyEnvelope CreateDealUnitAsync(ctx).TenantId(tenantId).DealUnitCreateDto(dealUnitCreateDto).Execute()

Create a deal unit



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
	dealUnitCreateDto := *openapiclient.NewDealUnitCreateDto() // DealUnitCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.CreateDealUnitAsync(context.Background()).TenantId(tenantId).DealUnitCreateDto(dealUnitCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.CreateDealUnitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDealUnitAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.CreateDealUnitAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDealUnitAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **dealUnitCreateDto** | [**DealUnitCreateDto**](DealUnitCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGetDealUnitLinesAsync

> EmptyEnvelope CreateGetDealUnitLinesAsync(ctx, dealUnitId).TenantId(tenantId).DealUnitLineCreateDto(dealUnitLineCreateDto).Execute()

Create a deal unit line



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineCreateDto := *openapiclient.NewDealUnitLineCreateDto() // DealUnitLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.CreateGetDealUnitLinesAsync(context.Background(), dealUnitId).TenantId(tenantId).DealUnitLineCreateDto(dealUnitLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.CreateGetDealUnitLinesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGetDealUnitLinesAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.CreateGetDealUnitLinesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGetDealUnitLinesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **dealUnitLineCreateDto** | [**DealUnitLineCreateDto**](DealUnitLineCreateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDealUnitAsync

> EmptyEnvelope DeleteDealUnitAsync(ctx, dealUnitId).TenantId(tenantId).Execute()

Delete a deal unit



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.DeleteDealUnitAsync(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.DeleteDealUnitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDealUnitAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.DeleteDealUnitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDealUnitAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## DeleteDealUnitPriceAsync

> EmptyEnvelope DeleteDealUnitPriceAsync(ctx, dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()

Delete a deal unit line



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.DeleteDealUnitPriceAsync(context.Background(), dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.DeleteDealUnitPriceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDealUnitPriceAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.DeleteDealUnitPriceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 
**dealUnitLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDealUnitPriceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## GetDealUnitAsync

> DealUnitDtoEnvelope GetDealUnitAsync(ctx, dealUnitId).TenantId(tenantId).Execute()

Get deal unit by ID



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetDealUnitAsync(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetDealUnitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitAsync`: DealUnitDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetDealUnitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**DealUnitDtoEnvelope**](DealUnitDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDealUnitLinesAsync

> DealUnitLineDtoListEnvelope GetDealUnitLinesAsync(ctx, dealUnitId).TenantId(tenantId).ItemId(itemId).Execute()

Get deal unit lines



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetDealUnitLinesAsync(context.Background(), dealUnitId).TenantId(tenantId).ItemId(itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetDealUnitLinesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitLinesAsync`: DealUnitLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetDealUnitLinesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitLinesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **itemId** | **string** |  | 

### Return type

[**DealUnitLineDtoListEnvelope**](DealUnitLineDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDealUnitLinesCountAsync

> Int32Envelope GetDealUnitLinesCountAsync(ctx, dealUnitId).TenantId(tenantId).Execute()

Get deal unit lines count



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetDealUnitLinesCountAsync(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetDealUnitLinesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitLinesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetDealUnitLinesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitLinesCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## GetDealUnitPriceAsync

> DealUnitLineDtoEnvelope GetDealUnitPriceAsync(ctx, dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()

Get a deal unit line by ID



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetDealUnitPriceAsync(context.Background(), dealUnitId, dealUnitLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetDealUnitPriceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitPriceAsync`: DealUnitLineDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetDealUnitPriceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 
**dealUnitLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitPriceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**DealUnitLineDtoEnvelope**](DealUnitLineDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDealUnitsAsync

> DealUnitDtoListEnvelope GetDealUnitsAsync(ctx).TenantId(tenantId).Execute()

Get deal units



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetDealUnitsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetDealUnitsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitsAsync`: DealUnitDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetDealUnitsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**DealUnitDtoListEnvelope**](DealUnitDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDealUnitsCountAsync

> Int32Envelope GetDealUnitsCountAsync(ctx).TenantId(tenantId).Execute()

Get deal units count



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetDealUnitsCountAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetDealUnitsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDealUnitsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetDealUnitsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDealUnitsCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## GetExtendedDealUnitAsync

> ExtendedDealUnitDtoEnvelope GetExtendedDealUnitAsync(ctx, dealUnitId).TenantId(tenantId).Execute()

Get extended deal unit by ID



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetExtendedDealUnitAsync(context.Background(), dealUnitId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetExtendedDealUnitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedDealUnitAsync`: ExtendedDealUnitDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetExtendedDealUnitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedDealUnitAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**ExtendedDealUnitDtoEnvelope**](ExtendedDealUnitDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtendedDealUnitsAsync

> ExtendedDealUnitDtoListEnvelope GetExtendedDealUnitsAsync(ctx).TenantId(tenantId).Execute()

Get extended deal units



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.GetExtendedDealUnitsAsync(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.GetExtendedDealUnitsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExtendedDealUnitsAsync`: ExtendedDealUnitDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.GetExtendedDealUnitsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtendedDealUnitsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ExtendedDealUnitDtoListEnvelope**](ExtendedDealUnitDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDealUnitAsync

> EmptyEnvelope UpdateDealUnitAsync(ctx, dealUnitId).TenantId(tenantId).DealUnitUpdateDto(dealUnitUpdateDto).Execute()

Update a deal unit



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitUpdateDto := *openapiclient.NewDealUnitUpdateDto() // DealUnitUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.UpdateDealUnitAsync(context.Background(), dealUnitId).TenantId(tenantId).DealUnitUpdateDto(dealUnitUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.UpdateDealUnitAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDealUnitAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.UpdateDealUnitAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDealUnitAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **dealUnitUpdateDto** | [**DealUnitUpdateDto**](DealUnitUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDealUnitPriceAsync

> EmptyEnvelope UpdateDealUnitPriceAsync(ctx, dealUnitId, dealUnitLineId).TenantId(tenantId).DealUnitLineUpdateDto(dealUnitLineUpdateDto).Execute()

Update a deal unit line



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
	dealUnitId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	dealUnitLineUpdateDto := *openapiclient.NewDealUnitLineUpdateDto() // DealUnitLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DealUnitsAPI.UpdateDealUnitPriceAsync(context.Background(), dealUnitId, dealUnitLineId).TenantId(tenantId).DealUnitLineUpdateDto(dealUnitLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DealUnitsAPI.UpdateDealUnitPriceAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDealUnitPriceAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DealUnitsAPI.UpdateDealUnitPriceAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dealUnitId** | **string** |  | 
**dealUnitLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDealUnitPriceAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **dealUnitLineUpdateDto** | [**DealUnitLineUpdateDto**](DealUnitLineUpdateDto.md) |  | 

### Return type

[**EmptyEnvelope**](EmptyEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

