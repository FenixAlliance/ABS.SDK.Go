# \DiscountListsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDiscountList**](DiscountListsAPI.md#CreateDiscountList) | **Post** /api/v2/PricingService/DiscountLists | Creates a new discount list
[**CreateDiscountListEntry**](DiscountListsAPI.md#CreateDiscountListEntry) | **Post** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts | Creates a discount list entry
[**DeleteDiscountList**](DiscountListsAPI.md#DeleteDiscountList) | **Delete** /api/v2/PricingService/DiscountLists/{discountListId} | Deletes a discount list
[**DeleteDiscountListEntry**](DiscountListsAPI.md#DeleteDiscountListEntry) | **Delete** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts/{discountListEntryId} | Deletes a discount list entry
[**GetDiscountList**](DiscountListsAPI.md#GetDiscountList) | **Get** /api/v2/PricingService/DiscountLists/{discountListId} | Gets a discount list by ID
[**GetDiscountListEntries**](DiscountListsAPI.md#GetDiscountListEntries) | **Get** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts | Retrieves discounts in a discount list
[**GetDiscountListEntriesCount**](DiscountListsAPI.md#GetDiscountListEntriesCount) | **Get** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts/Count | Counts discounts in a discount list
[**GetDiscountListEntry**](DiscountListsAPI.md#GetDiscountListEntry) | **Get** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts/{discountListEntryId} | Gets a discount list entry by ID
[**GetDiscountLists**](DiscountListsAPI.md#GetDiscountLists) | **Get** /api/v2/PricingService/DiscountLists | Retrieves all discount lists
[**GetDiscountListsCount**](DiscountListsAPI.md#GetDiscountListsCount) | **Get** /api/v2/PricingService/DiscountLists/Count | Counts discount lists
[**UpdateDiscountList**](DiscountListsAPI.md#UpdateDiscountList) | **Put** /api/v2/PricingService/DiscountLists/{discountListId} | Updates a discount list
[**UpdateDiscountListEntry**](DiscountListsAPI.md#UpdateDiscountListEntry) | **Put** /api/v2/PricingService/DiscountLists/{discountListId}/Discounts/{discountListEntryId} | Updates a discount list entry



## CreateDiscountList

> EmptyEnvelope CreateDiscountList(ctx).TenantId(tenantId).DiscountListCreateDto(discountListCreateDto).Execute()

Creates a new discount list



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
	discountListCreateDto := *openapiclient.NewDiscountListCreateDto() // DiscountListCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.CreateDiscountList(context.Background()).TenantId(tenantId).DiscountListCreateDto(discountListCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.CreateDiscountList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDiscountList`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.CreateDiscountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **discountListCreateDto** | [**DiscountListCreateDto**](DiscountListCreateDto.md) |  | 

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


## CreateDiscountListEntry

> EmptyEnvelope CreateDiscountListEntry(ctx, discountListId).TenantId(tenantId).DiscountCreateDto(discountCreateDto).Execute()

Creates a discount list entry



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountCreateDto := *openapiclient.NewDiscountCreateDto() // DiscountCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.CreateDiscountListEntry(context.Background(), discountListId).TenantId(tenantId).DiscountCreateDto(discountCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.CreateDiscountListEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDiscountListEntry`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.CreateDiscountListEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscountListEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **discountCreateDto** | [**DiscountCreateDto**](DiscountCreateDto.md) |  | 

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


## DeleteDiscountList

> EmptyEnvelope DeleteDiscountList(ctx, discountListId).TenantId(tenantId).Execute()

Deletes a discount list



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.DeleteDiscountList(context.Background(), discountListId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.DeleteDiscountList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDiscountList`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.DeleteDiscountList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiscountListRequest struct via the builder pattern


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


## DeleteDiscountListEntry

> EmptyEnvelope DeleteDiscountListEntry(ctx, discountListId, discountListEntryId).TenantId(tenantId).Execute()

Deletes a discount list entry



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountListEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.DeleteDiscountListEntry(context.Background(), discountListId, discountListEntryId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.DeleteDiscountListEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDiscountListEntry`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.DeleteDiscountListEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 
**discountListEntryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiscountListEntryRequest struct via the builder pattern


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


## GetDiscountList

> DiscountListDtoEnvelope GetDiscountList(ctx, discountListId).TenantId(tenantId).Execute()

Gets a discount list by ID



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.GetDiscountList(context.Background(), discountListId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.GetDiscountList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscountList`: DiscountListDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.GetDiscountList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**DiscountListDtoEnvelope**](DiscountListDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscountListEntries

> DiscountDtoListEnvelope GetDiscountListEntries(ctx, discountListId).TenantId(tenantId).Execute()

Retrieves discounts in a discount list



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.GetDiscountListEntries(context.Background(), discountListId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.GetDiscountListEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscountListEntries`: DiscountDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.GetDiscountListEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountListEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


### Return type

[**DiscountDtoListEnvelope**](DiscountDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscountListEntriesCount

> Int32Envelope GetDiscountListEntriesCount(ctx, discountListId).TenantId(tenantId).Execute()

Counts discounts in a discount list



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.GetDiscountListEntriesCount(context.Background(), discountListId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.GetDiscountListEntriesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscountListEntriesCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.GetDiscountListEntriesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountListEntriesCountRequest struct via the builder pattern


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


## GetDiscountListEntry

> DiscountDtoEnvelope GetDiscountListEntry(ctx, discountListId, discountListEntryId).TenantId(tenantId).Execute()

Gets a discount list entry by ID



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountListEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.GetDiscountListEntry(context.Background(), discountListId, discountListEntryId).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.GetDiscountListEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscountListEntry`: DiscountDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.GetDiscountListEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 
**discountListEntryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountListEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 



### Return type

[**DiscountDtoEnvelope**](DiscountDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscountLists

> DiscountListDtoListEnvelope GetDiscountLists(ctx).TenantId(tenantId).Execute()

Retrieves all discount lists



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
	resp, r, err := apiClient.DiscountListsAPI.GetDiscountLists(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.GetDiscountLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscountLists`: DiscountListDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.GetDiscountLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

### Return type

[**DiscountListDtoListEnvelope**](DiscountListDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscountListsCount

> Int32Envelope GetDiscountListsCount(ctx).TenantId(tenantId).Execute()

Counts discount lists



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
	resp, r, err := apiClient.DiscountListsAPI.GetDiscountListsCount(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.GetDiscountListsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDiscountListsCount`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.GetDiscountListsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountListsCountRequest struct via the builder pattern


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


## UpdateDiscountList

> EmptyEnvelope UpdateDiscountList(ctx, discountListId).TenantId(tenantId).DiscountListUpdateDto(discountListUpdateDto).Execute()

Updates a discount list



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountListUpdateDto := *openapiclient.NewDiscountListUpdateDto() // DiscountListUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.UpdateDiscountList(context.Background(), discountListId).TenantId(tenantId).DiscountListUpdateDto(discountListUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.UpdateDiscountList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDiscountList`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.UpdateDiscountList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiscountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **discountListUpdateDto** | [**DiscountListUpdateDto**](DiscountListUpdateDto.md) |  | 

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


## UpdateDiscountListEntry

> EmptyEnvelope UpdateDiscountListEntry(ctx, discountListId, discountListEntryId).TenantId(tenantId).DiscountUpdateDto(discountUpdateDto).Execute()

Updates a discount list entry



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
	discountListId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountListEntryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	discountUpdateDto := *openapiclient.NewDiscountUpdateDto() // DiscountUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DiscountListsAPI.UpdateDiscountListEntry(context.Background(), discountListId, discountListEntryId).TenantId(tenantId).DiscountUpdateDto(discountUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscountListsAPI.UpdateDiscountListEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDiscountListEntry`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DiscountListsAPI.UpdateDiscountListEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**discountListId** | **string** |  | 
**discountListEntryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiscountListEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **discountUpdateDto** | [**DiscountUpdateDto**](DiscountUpdateDto.md) |  | 

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

