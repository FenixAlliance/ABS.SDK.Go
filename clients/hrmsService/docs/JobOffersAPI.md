# \JobOffersAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseJobOfferAsync**](JobOffersAPI.md#CloseJobOfferAsync) | **Post** /api/v2/HrmsService/JobOffers/{jobOfferId}/Close | Close a job offer
[**CreateJobOfferAsync**](JobOffersAPI.md#CreateJobOfferAsync) | **Post** /api/v2/HrmsService/JobOffers | Create a job offer
[**DeleteJobOfferAsync**](JobOffersAPI.md#DeleteJobOfferAsync) | **Delete** /api/v2/HrmsService/JobOffers/{jobOfferId} | Delete a job offer
[**FillJobOfferAsync**](JobOffersAPI.md#FillJobOfferAsync) | **Post** /api/v2/HrmsService/JobOffers/{jobOfferId}/Fill | Mark a job offer filled
[**GetJobOfferByIdAsync**](JobOffersAPI.md#GetJobOfferByIdAsync) | **Get** /api/v2/HrmsService/JobOffers/{jobOfferId} | Get job offer by ID
[**GetJobOffersAsync**](JobOffersAPI.md#GetJobOffersAsync) | **Get** /api/v2/HrmsService/JobOffers | Get job offers
[**GetJobOffersCountAsync**](JobOffersAPI.md#GetJobOffersCountAsync) | **Get** /api/v2/HrmsService/JobOffers/Count | Count job offers
[**GetPublicJobOfferByIdAsync**](JobOffersAPI.md#GetPublicJobOfferByIdAsync) | **Get** /api/v2/HrmsService/JobOffers/Public/{jobOfferId} | Get public job offer by ID
[**GetPublicJobOffersAsync**](JobOffersAPI.md#GetPublicJobOffersAsync) | **Get** /api/v2/HrmsService/JobOffers/Public | Get public job offers
[**GetPublicJobOffersCountAsync**](JobOffersAPI.md#GetPublicJobOffersCountAsync) | **Get** /api/v2/HrmsService/JobOffers/Public/Count | Count public job offers
[**PatchJobOfferAsync**](JobOffersAPI.md#PatchJobOfferAsync) | **Patch** /api/v2/HrmsService/JobOffers/{jobOfferId} | Patch a job offer
[**PublishJobOfferAsync**](JobOffersAPI.md#PublishJobOfferAsync) | **Post** /api/v2/HrmsService/JobOffers/{jobOfferId}/Publish | Publish a job offer
[**UpdateJobOfferAsync**](JobOffersAPI.md#UpdateJobOfferAsync) | **Put** /api/v2/HrmsService/JobOffers/{jobOfferId} | Update a job offer



## CloseJobOfferAsync

> EmptyEnvelope CloseJobOfferAsync(ctx, jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Close a job offer



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
	jobOfferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.CloseJobOfferAsync(context.Background(), jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.CloseJobOfferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloseJobOfferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.CloseJobOfferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseJobOfferAsyncRequest struct via the builder pattern


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


## CreateJobOfferAsync

> EmptyEnvelope CreateJobOfferAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferCreateDto(jobOfferCreateDto).Execute()

Create a job offer



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
	jobOfferCreateDto := *openapiclient.NewJobOfferCreateDto() // JobOfferCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.CreateJobOfferAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferCreateDto(jobOfferCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.CreateJobOfferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJobOfferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.CreateJobOfferAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobOfferAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferCreateDto** | [**JobOfferCreateDto**](JobOfferCreateDto.md) |  | 

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


## DeleteJobOfferAsync

> EmptyEnvelope DeleteJobOfferAsync(ctx, jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete a job offer



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
	jobOfferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.DeleteJobOfferAsync(context.Background(), jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.DeleteJobOfferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJobOfferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.DeleteJobOfferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobOfferAsyncRequest struct via the builder pattern


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


## FillJobOfferAsync

> EmptyEnvelope FillJobOfferAsync(ctx, jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Mark a job offer filled



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
	jobOfferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.FillJobOfferAsync(context.Background(), jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.FillJobOfferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FillJobOfferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.FillJobOfferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFillJobOfferAsyncRequest struct via the builder pattern


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


## GetJobOfferByIdAsync

> JobOfferDtoEnvelope GetJobOfferByIdAsync(ctx, jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get job offer by ID



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
	jobOfferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.GetJobOfferByIdAsync(context.Background(), jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.GetJobOfferByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobOfferByIdAsync`: JobOfferDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.GetJobOfferByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobOfferByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JobOfferDtoEnvelope**](JobOfferDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobOffersAsync

> JobOfferDtoListEnvelope GetJobOffersAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferDtoCollectionQueryParameters(jobOfferDtoCollectionQueryParameters).Execute()

Get job offers



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
	jobOfferDtoCollectionQueryParameters := *openapiclient.NewJobOfferDtoCollectionQueryParameters() // JobOfferDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.GetJobOffersAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferDtoCollectionQueryParameters(jobOfferDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.GetJobOffersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobOffersAsync`: JobOfferDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.GetJobOffersAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobOffersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferDtoCollectionQueryParameters** | [**JobOfferDtoCollectionQueryParameters**](JobOfferDtoCollectionQueryParameters.md) |  | 

### Return type

[**JobOfferDtoListEnvelope**](JobOfferDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobOffersCountAsync

> Int32Envelope GetJobOffersCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferDtoCollectionQueryParameters(jobOfferDtoCollectionQueryParameters).Execute()

Count job offers



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
	jobOfferDtoCollectionQueryParameters := *openapiclient.NewJobOfferDtoCollectionQueryParameters() // JobOfferDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.GetJobOffersCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferDtoCollectionQueryParameters(jobOfferDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.GetJobOffersCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobOffersCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.GetJobOffersCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobOffersCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferDtoCollectionQueryParameters** | [**JobOfferDtoCollectionQueryParameters**](JobOfferDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicJobOfferByIdAsync

> JobOfferDtoEnvelope GetPublicJobOfferByIdAsync(ctx, jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get public job offer by ID



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
	jobOfferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.GetPublicJobOfferByIdAsync(context.Background(), jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.GetPublicJobOfferByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicJobOfferByIdAsync`: JobOfferDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.GetPublicJobOfferByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicJobOfferByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**JobOfferDtoEnvelope**](JobOfferDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicJobOffersAsync

> JobOfferDtoListEnvelope GetPublicJobOffersAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferDtoCollectionQueryParameters(jobOfferDtoCollectionQueryParameters).Execute()

Get public job offers



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	jobOfferDtoCollectionQueryParameters := *openapiclient.NewJobOfferDtoCollectionQueryParameters() // JobOfferDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.GetPublicJobOffersAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferDtoCollectionQueryParameters(jobOfferDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.GetPublicJobOffersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicJobOffersAsync`: JobOfferDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.GetPublicJobOffersAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicJobOffersAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferDtoCollectionQueryParameters** | [**JobOfferDtoCollectionQueryParameters**](JobOfferDtoCollectionQueryParameters.md) |  | 

### Return type

[**JobOfferDtoListEnvelope**](JobOfferDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicJobOffersCountAsync

> Int32Envelope GetPublicJobOffersCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferDtoCollectionQueryParameters(jobOfferDtoCollectionQueryParameters).Execute()

Count public job offers



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
	tenantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	jobOfferDtoCollectionQueryParameters := *openapiclient.NewJobOfferDtoCollectionQueryParameters() // JobOfferDtoCollectionQueryParameters |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.GetPublicJobOffersCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferDtoCollectionQueryParameters(jobOfferDtoCollectionQueryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.GetPublicJobOffersCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicJobOffersCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.GetPublicJobOffersCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicJobOffersCountAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferDtoCollectionQueryParameters** | [**JobOfferDtoCollectionQueryParameters**](JobOfferDtoCollectionQueryParameters.md) |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchJobOfferAsync

> EmptyEnvelope PatchJobOfferAsync(ctx, jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()

Patch a job offer



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
	jobOfferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	patchOperation := []openapiclient.PatchOperation{*openapiclient.NewPatchOperation()} // []PatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.PatchJobOfferAsync(context.Background(), jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).PatchOperation(patchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.PatchJobOfferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchJobOfferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.PatchJobOfferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchJobOfferAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **patchOperation** | [**[]PatchOperation**](PatchOperation.md) |  | 

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


## PublishJobOfferAsync

> EmptyEnvelope PublishJobOfferAsync(ctx, jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Publish a job offer



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
	jobOfferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.PublishJobOfferAsync(context.Background(), jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.PublishJobOfferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishJobOfferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.PublishJobOfferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishJobOfferAsyncRequest struct via the builder pattern


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


## UpdateJobOfferAsync

> EmptyEnvelope UpdateJobOfferAsync(ctx, jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferUpdateDto(jobOfferUpdateDto).Execute()

Update a job offer



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
	jobOfferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	jobOfferUpdateDto := *openapiclient.NewJobOfferUpdateDto() // JobOfferUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobOffersAPI.UpdateJobOfferAsync(context.Background(), jobOfferId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).JobOfferUpdateDto(jobOfferUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobOffersAPI.UpdateJobOfferAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateJobOfferAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `JobOffersAPI.UpdateJobOfferAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobOfferId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJobOfferAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **jobOfferUpdateDto** | [**JobOfferUpdateDto**](JobOfferUpdateDto.md) |  | 

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

