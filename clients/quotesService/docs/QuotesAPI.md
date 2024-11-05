# \QuotesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2QuotesServiceQuotesCountGet**](QuotesAPI.md#ApiV2QuotesServiceQuotesCountGet) | **Get** /api/v2/QuotesService/Quotes/Count | 
[**ApiV2QuotesServiceQuotesExtendedGet**](QuotesAPI.md#ApiV2QuotesServiceQuotesExtendedGet) | **Get** /api/v2/QuotesService/Quotes/Extended | 
[**ApiV2QuotesServiceQuotesGet**](QuotesAPI.md#ApiV2QuotesServiceQuotesGet) | **Get** /api/v2/QuotesService/Quotes | 
[**ApiV2QuotesServiceQuotesPost**](QuotesAPI.md#ApiV2QuotesServiceQuotesPost) | **Post** /api/v2/QuotesService/Quotes | 
[**ApiV2QuotesServiceQuotesQuoteIdCalculatePut**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdCalculatePut) | **Put** /api/v2/QuotesService/Quotes/{quoteId}/Calculate | 
[**ApiV2QuotesServiceQuotesQuoteIdDelete**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdDelete) | **Delete** /api/v2/QuotesService/Quotes/{quoteId} | 
[**ApiV2QuotesServiceQuotesQuoteIdLinesCountGet**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdLinesCountGet) | **Get** /api/v2/QuotesService/Quotes/{quoteId}/Lines/Count | 
[**ApiV2QuotesServiceQuotesQuoteIdLinesGet**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdLinesGet) | **Get** /api/v2/QuotesService/Quotes/{quoteId}/Lines | 
[**ApiV2QuotesServiceQuotesQuoteIdLinesPost**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdLinesPost) | **Post** /api/v2/QuotesService/Quotes/{quoteId}/Lines | 
[**ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePut**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePut) | **Put** /api/v2/QuotesService/Quotes/{quoteId}/Lines/{quoteLineId}/Calculate | 
[**ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDelete**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDelete) | **Delete** /api/v2/QuotesService/Quotes/{quoteId}/Lines/{quoteLineId} | 
[**ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGet**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGet) | **Get** /api/v2/QuotesService/Quotes/{quoteId}/Lines/{quoteLineId} | 
[**ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPut**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPut) | **Put** /api/v2/QuotesService/Quotes/{quoteId}/Lines/{quoteLineId} | 
[**ApiV2QuotesServiceQuotesQuoteIdPut**](QuotesAPI.md#ApiV2QuotesServiceQuotesQuoteIdPut) | **Put** /api/v2/QuotesService/Quotes/{quoteId} | 
[**GetQuoteAsync**](QuotesAPI.md#GetQuoteAsync) | **Get** /api/v2/QuotesService/Quotes/{quoteId} | 



## ApiV2QuotesServiceQuotesCountGet

> Int32Envelope ApiV2QuotesServiceQuotesCountGet(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesCountGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

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


## ApiV2QuotesServiceQuotesExtendedGet

> ExtendedQuoteDtoListEnvelope ApiV2QuotesServiceQuotesExtendedGet(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesExtendedGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesExtendedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesExtendedGet`: ExtendedQuoteDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesExtendedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesExtendedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**ExtendedQuoteDtoListEnvelope**](ExtendedQuoteDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2QuotesServiceQuotesGet

> QuoteDtoListEnvelope ApiV2QuotesServiceQuotesGet(ctx).TenantId(tenantId).Execute()



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
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesGet(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesGet`: QuoteDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**QuoteDtoListEnvelope**](QuoteDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2QuotesServiceQuotesPost

> EmptyEnvelope ApiV2QuotesServiceQuotesPost(ctx).TenantId(tenantId).QuoteCreateDto(quoteCreateDto).Execute()



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
	quoteCreateDto := *openapiclient.NewQuoteCreateDto() // QuoteCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesPost(context.Background()).TenantId(tenantId).QuoteCreateDto(quoteCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **quoteCreateDto** | [**QuoteCreateDto**](QuoteCreateDto.md) |  | 

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


## ApiV2QuotesServiceQuotesQuoteIdCalculatePut

> EmptyEnvelope ApiV2QuotesServiceQuotesQuoteIdCalculatePut(ctx, quoteId).TenantId(tenantId).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdCalculatePut(context.Background(), quoteId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdCalculatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdCalculatePut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdCalculatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdCalculatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## ApiV2QuotesServiceQuotesQuoteIdDelete

> EmptyEnvelope ApiV2QuotesServiceQuotesQuoteIdDelete(ctx, quoteId).TenantId(tenantId).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdDelete(context.Background(), quoteId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 

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


## ApiV2QuotesServiceQuotesQuoteIdLinesCountGet

> Int32Envelope ApiV2QuotesServiceQuotesQuoteIdLinesCountGet(ctx, quoteId).TenantId(tenantId).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesCountGet(context.Background(), quoteId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdLinesCountGet`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdLinesCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


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


## ApiV2QuotesServiceQuotesQuoteIdLinesGet

> QuoteLineDtoListEnvelope ApiV2QuotesServiceQuotesQuoteIdLinesGet(ctx, quoteId).TenantId(tenantId).ItemId(itemId).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesGet(context.Background(), quoteId).TenantId(tenantId).ItemId(itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdLinesGet`: QuoteLineDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdLinesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **itemId** | **string** |  | 

### Return type

[**QuoteLineDtoListEnvelope**](QuoteLineDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2QuotesServiceQuotesQuoteIdLinesPost

> EmptyEnvelope ApiV2QuotesServiceQuotesQuoteIdLinesPost(ctx, quoteId).TenantId(tenantId).QuoteLineCreateDto(quoteLineCreateDto).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteLineCreateDto := *openapiclient.NewQuoteLineCreateDto() // QuoteLineCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesPost(context.Background(), quoteId).TenantId(tenantId).QuoteLineCreateDto(quoteLineCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdLinesPost`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdLinesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **quoteLineCreateDto** | [**QuoteLineCreateDto**](QuoteLineCreateDto.md) |  | 

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


## ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePut

> EmptyEnvelope ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePut(ctx, quoteId, quoteLineId).TenantId(tenantId).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePut(context.Background(), quoteId, quoteLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**quoteLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdCalculatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDelete

> EmptyEnvelope ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDelete(ctx, quoteId, quoteLineId).TenantId(tenantId).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDelete(context.Background(), quoteId, quoteLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDelete`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**quoteLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



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


## ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGet

> QuoteLineDtoEnvelope ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGet(ctx, quoteId, quoteLineId).TenantId(tenantId).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGet(context.Background(), quoteId, quoteLineId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGet`: QuoteLineDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**quoteLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**QuoteLineDtoEnvelope**](QuoteLineDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPut

> EmptyEnvelope ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPut(ctx, quoteId, quoteLineId).TenantId(tenantId).QuoteLineUpdateDto(quoteLineUpdateDto).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteLineId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteLineUpdateDto := *openapiclient.NewQuoteLineUpdateDto() // QuoteLineUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPut(context.Background(), quoteId, quoteLineId).TenantId(tenantId).QuoteLineUpdateDto(quoteLineUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 
**quoteLineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdLinesQuoteLineIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **quoteLineUpdateDto** | [**QuoteLineUpdateDto**](QuoteLineUpdateDto.md) |  | 

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


## ApiV2QuotesServiceQuotesQuoteIdPut

> EmptyEnvelope ApiV2QuotesServiceQuotesQuoteIdPut(ctx, quoteId).TenantId(tenantId).QuoteUpdateDto(quoteUpdateDto).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	quoteUpdateDto := *openapiclient.NewQuoteUpdateDto() // QuoteUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdPut(context.Background(), quoteId).TenantId(tenantId).QuoteUpdateDto(quoteUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2QuotesServiceQuotesQuoteIdPut`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.ApiV2QuotesServiceQuotesQuoteIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2QuotesServiceQuotesQuoteIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **quoteUpdateDto** | [**QuoteUpdateDto**](QuoteUpdateDto.md) |  | 

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


## GetQuoteAsync

> QuoteDtoEnvelope GetQuoteAsync(ctx, quoteId).TenantId(tenantId).Execute()



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
	quoteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetQuoteAsync(context.Background(), quoteId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetQuoteAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuoteAsync`: QuoteDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetQuoteAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**QuoteDtoEnvelope**](QuoteDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

