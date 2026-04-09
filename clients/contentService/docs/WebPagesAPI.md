# \WebPagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountWebPagesAsync**](WebPagesAPI.md#CountWebPagesAsync) | **Get** /api/v2/ContentService/WebPages/Count | Count web pages
[**CreateWebPageAsync**](WebPagesAPI.md#CreateWebPageAsync) | **Post** /api/v2/ContentService/WebPages | Create a web page
[**CreateWebPageCategoryRelationAsync**](WebPagesAPI.md#CreateWebPageCategoryRelationAsync) | **Post** /api/v2/ContentService/WebPages/{webPageId}/Categories | Create a web page category relation
[**CreateWebPageTagRelationAsync**](WebPagesAPI.md#CreateWebPageTagRelationAsync) | **Post** /api/v2/ContentService/WebPages/{webPageId}/Tags | Create a web page tag relation
[**DeleteWebPageAsync**](WebPagesAPI.md#DeleteWebPageAsync) | **Delete** /api/v2/ContentService/WebPages/{webPageId} | Delete a web page
[**GetCategoriesByWebPageAsync**](WebPagesAPI.md#GetCategoriesByWebPageAsync) | **Get** /api/v2/ContentService/WebPages/{webPageId}/Categories | Get categories by web page
[**GetTagsByWebPageAsync**](WebPagesAPI.md#GetTagsByWebPageAsync) | **Get** /api/v2/ContentService/WebPages/{webPageId}/Tags | Get tags by web page
[**GetWebPageByIdAsync**](WebPagesAPI.md#GetWebPageByIdAsync) | **Get** /api/v2/ContentService/WebPages/{webPageId} | Get web page by ID
[**GetWebPagesAsync**](WebPagesAPI.md#GetWebPagesAsync) | **Get** /api/v2/ContentService/WebPages | Get web pages
[**RelateWebPageToCategoryAsync**](WebPagesAPI.md#RelateWebPageToCategoryAsync) | **Post** /api/v2/ContentService/WebPages/{webPageId}/Categories/{categoryId} | Relate web page to category
[**RelateWebPageToTagAsync**](WebPagesAPI.md#RelateWebPageToTagAsync) | **Post** /api/v2/ContentService/WebPages/{webPageId}/Tags/{tagId} | Relate web page to tag
[**UnrelateWebPageCategoryAsync**](WebPagesAPI.md#UnrelateWebPageCategoryAsync) | **Delete** /api/v2/ContentService/WebPages/{webPageId}/Categories/{categoryId} | Unrelate web page from category
[**UnrelateWebPageTagAsync**](WebPagesAPI.md#UnrelateWebPageTagAsync) | **Delete** /api/v2/ContentService/WebPages/{webPageId}/Tags/{tagId} | Unrelate web page from tag
[**UpdateWebPageAsync**](WebPagesAPI.md#UpdateWebPageAsync) | **Put** /api/v2/ContentService/WebPages/{webPageId} | Update a web page



## CountWebPagesAsync

> Int32Envelope CountWebPagesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count web pages



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
	resp, r, err := apiClient.WebPagesAPI.CountWebPagesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.CountWebPagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountWebPagesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `WebPagesAPI.CountWebPagesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountWebPagesAsyncRequest struct via the builder pattern


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


## CreateWebPageAsync

> CreateWebPageAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPageCreateDto(webPageCreateDto).Execute()

Create a web page



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
	webPageCreateDto := *openapiclient.NewWebPageCreateDto() // WebPageCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.CreateWebPageAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPageCreateDto(webPageCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.CreateWebPageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebPageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webPageCreateDto** | [**WebPageCreateDto**](WebPageCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebPageCategoryRelationAsync

> CreateWebPageCategoryRelationAsync(ctx, webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPageCategoryCreateDto(webPageCategoryCreateDto).Execute()

Create a web page category relation



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	webPageCategoryCreateDto := *openapiclient.NewWebPageCategoryCreateDto() // WebPageCategoryCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.CreateWebPageCategoryRelationAsync(context.Background(), webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPageCategoryCreateDto(webPageCategoryCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.CreateWebPageCategoryRelationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebPageCategoryRelationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webPageCategoryCreateDto** | [**WebPageCategoryCreateDto**](WebPageCategoryCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebPageTagRelationAsync

> CreateWebPageTagRelationAsync(ctx, webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPageTagCreateDto(webPageTagCreateDto).Execute()

Create a web page tag relation



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	webPageTagCreateDto := *openapiclient.NewWebPageTagCreateDto() // WebPageTagCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.CreateWebPageTagRelationAsync(context.Background(), webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPageTagCreateDto(webPageTagCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.CreateWebPageTagRelationAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebPageTagRelationAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webPageTagCreateDto** | [**WebPageTagCreateDto**](WebPageTagCreateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebPageAsync

> DeleteWebPageAsync(ctx, webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a web page



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.DeleteWebPageAsync(context.Background(), webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.DeleteWebPageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebPageAsyncRequest struct via the builder pattern


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


## GetCategoriesByWebPageAsync

> WebPageCategoryDtoListEnvelope GetCategoriesByWebPageAsync(ctx, webPageId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get categories by web page



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebPagesAPI.GetCategoriesByWebPageAsync(context.Background(), webPageId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.GetCategoriesByWebPageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategoriesByWebPageAsync`: WebPageCategoryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebPagesAPI.GetCategoriesByWebPageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesByWebPageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WebPageCategoryDtoListEnvelope**](WebPageCategoryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsByWebPageAsync

> WebPageTagDtoListEnvelope GetTagsByWebPageAsync(ctx, webPageId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get tags by web page



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebPagesAPI.GetTagsByWebPageAsync(context.Background(), webPageId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.GetTagsByWebPageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagsByWebPageAsync`: WebPageTagDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebPagesAPI.GetTagsByWebPageAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsByWebPageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WebPageTagDtoListEnvelope**](WebPageTagDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebPageByIdAsync

> WebPageDtoEnvelope GetWebPageByIdAsync(ctx, webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get web page by ID



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebPagesAPI.GetWebPageByIdAsync(context.Background(), webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.GetWebPageByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebPageByIdAsync`: WebPageDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebPagesAPI.GetWebPageByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebPageByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WebPageDtoEnvelope**](WebPageDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebPagesAsync

> WebPageDtoListEnvelope GetWebPagesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get web pages



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
	resp, r, err := apiClient.WebPagesAPI.GetWebPagesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.GetWebPagesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebPagesAsync`: WebPageDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `WebPagesAPI.GetWebPagesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebPagesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**WebPageDtoListEnvelope**](WebPageDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelateWebPageToCategoryAsync

> RelateWebPageToCategoryAsync(ctx, webPageId, categoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Relate web page to category



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	categoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.RelateWebPageToCategoryAsync(context.Background(), webPageId, categoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.RelateWebPageToCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelateWebPageToCategoryAsyncRequest struct via the builder pattern


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


## RelateWebPageToTagAsync

> RelateWebPageToTagAsync(ctx, webPageId, tagId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Relate web page to tag



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.RelateWebPageToTagAsync(context.Background(), webPageId, tagId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.RelateWebPageToTagAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelateWebPageToTagAsyncRequest struct via the builder pattern


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


## UnrelateWebPageCategoryAsync

> UnrelateWebPageCategoryAsync(ctx, webPageId, categoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Unrelate web page from category



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	categoryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.UnrelateWebPageCategoryAsync(context.Background(), webPageId, categoryId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.UnrelateWebPageCategoryAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnrelateWebPageCategoryAsyncRequest struct via the builder pattern


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


## UnrelateWebPageTagAsync

> UnrelateWebPageTagAsync(ctx, webPageId, tagId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Unrelate web page from tag



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.UnrelateWebPageTagAsync(context.Background(), webPageId, tagId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.UnrelateWebPageTagAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnrelateWebPageTagAsyncRequest struct via the builder pattern


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


## UpdateWebPageAsync

> UpdateWebPageAsync(ctx, webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPageUpdateDto(webPageUpdateDto).Execute()

Update a web page



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
	webPageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	webPageUpdateDto := *openapiclient.NewWebPageUpdateDto() // WebPageUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebPagesAPI.UpdateWebPageAsync(context.Background(), webPageId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).WebPageUpdateDto(webPageUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebPagesAPI.UpdateWebPageAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webPageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebPageAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **webPageUpdateDto** | [**WebPageUpdateDto**](WebPageUpdateDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

