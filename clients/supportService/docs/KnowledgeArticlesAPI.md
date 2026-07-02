# \KnowledgeArticlesAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKnowledgeArticleAsync**](KnowledgeArticlesAPI.md#CreateKnowledgeArticleAsync) | **Post** /api/v2/SupportService/KnowledgeArticles | Create a knowledge article
[**DeleteKnowledgeArticleAsync**](KnowledgeArticlesAPI.md#DeleteKnowledgeArticleAsync) | **Delete** /api/v2/SupportService/KnowledgeArticles/{knowledgeArticleId} | Delete a knowledge article
[**GetKnowledgeArticleAsync**](KnowledgeArticlesAPI.md#GetKnowledgeArticleAsync) | **Get** /api/v2/SupportService/KnowledgeArticles/{knowledgeArticleId} | Retrieve a knowledge article by ID
[**GetKnowledgeArticlesAsync**](KnowledgeArticlesAPI.md#GetKnowledgeArticlesAsync) | **Get** /api/v2/SupportService/KnowledgeArticles | Retrieve knowledge articles
[**GetKnowledgeArticlesCountAsync**](KnowledgeArticlesAPI.md#GetKnowledgeArticlesCountAsync) | **Get** /api/v2/SupportService/KnowledgeArticles/Count | Get knowledge articles count
[**PatchKnowledgeArticleAsync**](KnowledgeArticlesAPI.md#PatchKnowledgeArticleAsync) | **Patch** /api/v2/SupportService/KnowledgeArticles/{knowledgeArticleId} | Patch a knowledge article
[**UpdateKnowledgeArticleAsync**](KnowledgeArticlesAPI.md#UpdateKnowledgeArticleAsync) | **Put** /api/v2/SupportService/KnowledgeArticles/{knowledgeArticleId} | Update a knowledge article



## CreateKnowledgeArticleAsync

> EmptyEnvelope CreateKnowledgeArticleAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).KnowledgeArticleCreateDto(knowledgeArticleCreateDto).Execute()

Create a knowledge article

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
	knowledgeArticleCreateDto := *openapiclient.NewKnowledgeArticleCreateDto("Title_example") // KnowledgeArticleCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeArticlesAPI.CreateKnowledgeArticleAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).KnowledgeArticleCreateDto(knowledgeArticleCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeArticlesAPI.CreateKnowledgeArticleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKnowledgeArticleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeArticlesAPI.CreateKnowledgeArticleAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKnowledgeArticleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **knowledgeArticleCreateDto** | [**KnowledgeArticleCreateDto**](KnowledgeArticleCreateDto.md) |  | 

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


## DeleteKnowledgeArticleAsync

> EmptyEnvelope DeleteKnowledgeArticleAsync(ctx, knowledgeArticleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a knowledge article

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
	knowledgeArticleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeArticlesAPI.DeleteKnowledgeArticleAsync(context.Background(), knowledgeArticleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeArticlesAPI.DeleteKnowledgeArticleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKnowledgeArticleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeArticlesAPI.DeleteKnowledgeArticleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**knowledgeArticleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKnowledgeArticleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

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


## GetKnowledgeArticleAsync

> KnowledgeArticleDtoEnvelope GetKnowledgeArticleAsync(ctx, knowledgeArticleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve a knowledge article by ID

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
	knowledgeArticleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeArticlesAPI.GetKnowledgeArticleAsync(context.Background(), knowledgeArticleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeArticlesAPI.GetKnowledgeArticleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeArticleAsync`: KnowledgeArticleDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeArticlesAPI.GetKnowledgeArticleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**knowledgeArticleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeArticleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**KnowledgeArticleDtoEnvelope**](KnowledgeArticleDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeArticlesAsync

> KnowledgeArticleDtoListEnvelope GetKnowledgeArticlesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Retrieve knowledge articles

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
	resp, r, err := apiClient.KnowledgeArticlesAPI.GetKnowledgeArticlesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeArticlesAPI.GetKnowledgeArticlesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeArticlesAsync`: KnowledgeArticleDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeArticlesAPI.GetKnowledgeArticlesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeArticlesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**KnowledgeArticleDtoListEnvelope**](KnowledgeArticleDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKnowledgeArticlesCountAsync

> Int32Envelope GetKnowledgeArticlesCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get knowledge articles count

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
	resp, r, err := apiClient.KnowledgeArticlesAPI.GetKnowledgeArticlesCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeArticlesAPI.GetKnowledgeArticlesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKnowledgeArticlesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeArticlesAPI.GetKnowledgeArticlesCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKnowledgeArticlesCountAsyncRequest struct via the builder pattern


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


## PatchKnowledgeArticleAsync

> EmptyEnvelope PatchKnowledgeArticleAsync(ctx, knowledgeArticleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch a knowledge article



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
	knowledgeArticleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeArticlesAPI.PatchKnowledgeArticleAsync(context.Background(), knowledgeArticleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeArticlesAPI.PatchKnowledgeArticleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchKnowledgeArticleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeArticlesAPI.PatchKnowledgeArticleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**knowledgeArticleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchKnowledgeArticleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **operation** | [**[]Operation**](Operation.md) |  | 

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


## UpdateKnowledgeArticleAsync

> EmptyEnvelope UpdateKnowledgeArticleAsync(ctx, knowledgeArticleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).KnowledgeArticleUpdateDto(knowledgeArticleUpdateDto).Execute()

Update a knowledge article

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
	knowledgeArticleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	knowledgeArticleUpdateDto := *openapiclient.NewKnowledgeArticleUpdateDto() // KnowledgeArticleUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KnowledgeArticlesAPI.UpdateKnowledgeArticleAsync(context.Background(), knowledgeArticleId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).KnowledgeArticleUpdateDto(knowledgeArticleUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KnowledgeArticlesAPI.UpdateKnowledgeArticleAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateKnowledgeArticleAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `KnowledgeArticlesAPI.UpdateKnowledgeArticleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**knowledgeArticleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKnowledgeArticleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **knowledgeArticleUpdateDto** | [**KnowledgeArticleUpdateDto**](KnowledgeArticleUpdateDto.md) |  | 

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

