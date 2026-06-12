# \ActivityFeedsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountActivityTypesAsync**](ActivityFeedsAPI.md#CountActivityTypesAsync) | **Get** /api/v2/ActivitiesService/ActivityTypes/Count | Count Activity Types
[**CreateActivityAsync**](ActivityFeedsAPI.md#CreateActivityAsync) | **Post** /api/v2/ActivitiesService/ActivityFeeds/{activityFeedId}/Activities | Create an activity
[**CreateActivityTypeAsync**](ActivityFeedsAPI.md#CreateActivityTypeAsync) | **Post** /api/v2/ActivitiesService/ActivityTypes | Create Activity Type
[**DeleteActivityAsync**](ActivityFeedsAPI.md#DeleteActivityAsync) | **Delete** /api/v2/ActivitiesService/ActivityFeeds/{activityFeedId}/Activities/{activityId} | Delete an activity
[**DeleteActivityTypeAsync**](ActivityFeedsAPI.md#DeleteActivityTypeAsync) | **Delete** /api/v2/ActivitiesService/ActivityTypes/{activityTypeId} | Delete Activity Type
[**GetActivitiesAsync**](ActivityFeedsAPI.md#GetActivitiesAsync) | **Get** /api/v2/ActivitiesService/ActivityFeeds/{activityFeedId}/Activities | Get activities
[**GetActivitiesCountAsync**](ActivityFeedsAPI.md#GetActivitiesCountAsync) | **Get** /api/v2/ActivitiesService/ActivityFeeds/{activityFeedId}/Activities/Count | Count activities
[**GetActivityAsync**](ActivityFeedsAPI.md#GetActivityAsync) | **Get** /api/v2/ActivitiesService/ActivityFeeds/{activityFeedId}/Activities/{activityId} | Get activity by ID
[**GetActivityFeedAsync**](ActivityFeedsAPI.md#GetActivityFeedAsync) | **Get** /api/v2/ActivitiesService/ActivityFeeds/{activityFeedId} | Get activity feed by ID
[**GetActivityFeedsAsync**](ActivityFeedsAPI.md#GetActivityFeedsAsync) | **Get** /api/v2/ActivitiesService/ActivityFeeds | Get activity feeds
[**GetActivityFeedsCountAsync**](ActivityFeedsAPI.md#GetActivityFeedsCountAsync) | **Get** /api/v2/ActivitiesService/ActivityFeeds/Count | Count activity feeds
[**GetActivityTypeByIdAsync**](ActivityFeedsAPI.md#GetActivityTypeByIdAsync) | **Get** /api/v2/ActivitiesService/ActivityTypes/{activityTypeId} | Get Activity Type
[**GetActivityTypesAsync**](ActivityFeedsAPI.md#GetActivityTypesAsync) | **Get** /api/v2/ActivitiesService/ActivityTypes | Get Activity Types
[**PatchActivityAsync**](ActivityFeedsAPI.md#PatchActivityAsync) | **Patch** /api/v2/ActivitiesService/ActivityFeeds/{activityFeedId}/Activities/{activityId} | Patch an activity
[**PatchActivityTypeAsync**](ActivityFeedsAPI.md#PatchActivityTypeAsync) | **Patch** /api/v2/ActivitiesService/ActivityTypes/{activityTypeId} | Patch Activity Type
[**UpdateActivityAsync**](ActivityFeedsAPI.md#UpdateActivityAsync) | **Put** /api/v2/ActivitiesService/ActivityFeeds/{activityFeedId}/Activities/{activityId} | Update an activity
[**UpdateActivityTypeAsync**](ActivityFeedsAPI.md#UpdateActivityTypeAsync) | **Put** /api/v2/ActivitiesService/ActivityTypes/{activityTypeId} | Update Activity Type



## CountActivityTypesAsync

> Int32Envelope CountActivityTypesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count Activity Types



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
	resp, r, err := apiClient.ActivityFeedsAPI.CountActivityTypesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.CountActivityTypesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountActivityTypesAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.CountActivityTypesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountActivityTypesAsyncRequest struct via the builder pattern


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


## CreateActivityAsync

> ActivityRecordDtoEnvelope CreateActivityAsync(ctx, activityFeedId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ActivityRecordCreateDto(activityRecordCreateDto).Execute()

Create an activity



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
	activityFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	activityRecordCreateDto := *openapiclient.NewActivityRecordCreateDto() // ActivityRecordCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.CreateActivityAsync(context.Background(), activityFeedId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ActivityRecordCreateDto(activityRecordCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.CreateActivityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActivityAsync`: ActivityRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.CreateActivityAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityFeedId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateActivityAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **activityRecordCreateDto** | [**ActivityRecordCreateDto**](ActivityRecordCreateDto.md) |  | 

### Return type

[**ActivityRecordDtoEnvelope**](ActivityRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateActivityTypeAsync

> Envelope CreateActivityTypeAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ActivityTypeCreateDto(activityTypeCreateDto).Execute()

Create Activity Type



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
	activityTypeCreateDto := *openapiclient.NewActivityTypeCreateDto() // ActivityTypeCreateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.CreateActivityTypeAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ActivityTypeCreateDto(activityTypeCreateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.CreateActivityTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateActivityTypeAsync`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.CreateActivityTypeAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActivityTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **activityTypeCreateDto** | [**ActivityTypeCreateDto**](ActivityTypeCreateDto.md) |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActivityAsync

> EmptyEnvelope DeleteActivityAsync(ctx, activityFeedId, activityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete an activity



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
	activityFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	activityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.DeleteActivityAsync(context.Background(), activityFeedId, activityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.DeleteActivityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteActivityAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.DeleteActivityAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityFeedId** | **string** |  | 
**activityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActivityAsyncRequest struct via the builder pattern


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


## DeleteActivityTypeAsync

> Envelope DeleteActivityTypeAsync(ctx, activityTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Delete Activity Type



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
	activityTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.DeleteActivityTypeAsync(context.Background(), activityTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.DeleteActivityTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteActivityTypeAsync`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.DeleteActivityTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActivityTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivitiesAsync

> ActivityRecordDtoListEnvelope GetActivitiesAsync(ctx, activityFeedId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get activities



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
	activityFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.GetActivitiesAsync(context.Background(), activityFeedId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.GetActivitiesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivitiesAsync`: ActivityRecordDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.GetActivitiesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityFeedId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivitiesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ActivityRecordDtoListEnvelope**](ActivityRecordDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivitiesCountAsync

> Int32Envelope GetActivitiesCountAsync(ctx, activityFeedId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count activities



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
	activityFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.GetActivitiesCountAsync(context.Background(), activityFeedId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.GetActivitiesCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivitiesCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.GetActivitiesCountAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityFeedId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivitiesCountAsyncRequest struct via the builder pattern


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


## GetActivityAsync

> ActivityRecordDtoEnvelope GetActivityAsync(ctx, activityFeedId, activityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get activity by ID



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
	activityFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	activityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.GetActivityAsync(context.Background(), activityFeedId, activityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.GetActivityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityAsync`: ActivityRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.GetActivityAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityFeedId** | **string** |  | 
**activityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ActivityRecordDtoEnvelope**](ActivityRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityFeedAsync

> ActivityFeedDtoEnvelope GetActivityFeedAsync(ctx, activityFeedId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get activity feed by ID



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
	activityFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.GetActivityFeedAsync(context.Background(), activityFeedId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.GetActivityFeedAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityFeedAsync`: ActivityFeedDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.GetActivityFeedAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityFeedId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityFeedAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ActivityFeedDtoEnvelope**](ActivityFeedDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityFeedsAsync

> ActivityFeedDtoListEnvelope GetActivityFeedsAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get activity feeds



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
	resp, r, err := apiClient.ActivityFeedsAPI.GetActivityFeedsAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.GetActivityFeedsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityFeedsAsync`: ActivityFeedDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.GetActivityFeedsAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityFeedsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ActivityFeedDtoListEnvelope**](ActivityFeedDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityFeedsCountAsync

> Int32Envelope GetActivityFeedsCountAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count activity feeds



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
	resp, r, err := apiClient.ActivityFeedsAPI.GetActivityFeedsCountAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.GetActivityFeedsCountAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityFeedsCountAsync`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.GetActivityFeedsCountAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityFeedsCountAsyncRequest struct via the builder pattern


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


## GetActivityTypeByIdAsync

> ActivityTypeDtoEnvelope GetActivityTypeByIdAsync(ctx, activityTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Activity Type



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
	activityTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.GetActivityTypeByIdAsync(context.Background(), activityTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.GetActivityTypeByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityTypeByIdAsync`: ActivityTypeDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.GetActivityTypeByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityTypeByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ActivityTypeDtoEnvelope**](ActivityTypeDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivityTypesAsync

> ActivityTypeDtoListEnvelope GetActivityTypesAsync(ctx).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get Activity Types



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
	resp, r, err := apiClient.ActivityFeedsAPI.GetActivityTypesAsync(context.Background()).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.GetActivityTypesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityTypesAsync`: ActivityTypeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.GetActivityTypesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityTypesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**ActivityTypeDtoListEnvelope**](ActivityTypeDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchActivityAsync

> EmptyEnvelope PatchActivityAsync(ctx, activityFeedId, activityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch an activity



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
	activityFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	activityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.PatchActivityAsync(context.Background(), activityFeedId, activityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.PatchActivityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchActivityAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.PatchActivityAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityFeedId** | **string** |  | 
**activityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchActivityAsyncRequest struct via the builder pattern


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


## PatchActivityTypeAsync

> EmptyEnvelope PatchActivityTypeAsync(ctx, activityTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()

Patch Activity Type



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
	activityTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	operation := []openapiclient.Operation{*openapiclient.NewOperation()} // []Operation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.PatchActivityTypeAsync(context.Background(), activityTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Operation(operation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.PatchActivityTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchActivityTypeAsync`: EmptyEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.PatchActivityTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchActivityTypeAsyncRequest struct via the builder pattern


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


## UpdateActivityAsync

> ActivityRecordDtoEnvelope UpdateActivityAsync(ctx, activityFeedId, activityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ActivityRecordUpdateDto(activityRecordUpdateDto).Execute()

Update an activity



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
	activityFeedId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	activityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	activityRecordUpdateDto := *openapiclient.NewActivityRecordUpdateDto() // ActivityRecordUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.UpdateActivityAsync(context.Background(), activityFeedId, activityId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ActivityRecordUpdateDto(activityRecordUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.UpdateActivityAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActivityAsync`: ActivityRecordDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.UpdateActivityAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityFeedId** | **string** |  | 
**activityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActivityAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **activityRecordUpdateDto** | [**ActivityRecordUpdateDto**](ActivityRecordUpdateDto.md) |  | 

### Return type

[**ActivityRecordDtoEnvelope**](ActivityRecordDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActivityTypeAsync

> Envelope UpdateActivityTypeAsync(ctx, activityTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ActivityTypeUpdateDto(activityTypeUpdateDto).Execute()

Update Activity Type



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
	activityTypeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	activityTypeUpdateDto := *openapiclient.NewActivityTypeUpdateDto() // ActivityTypeUpdateDto |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityFeedsAPI.UpdateActivityTypeAsync(context.Background(), activityTypeId).TenantId(tenantId).ApiVersion(apiVersion).XApiVersion(xApiVersion).ActivityTypeUpdateDto(activityTypeUpdateDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityFeedsAPI.UpdateActivityTypeAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActivityTypeAsync`: Envelope
	fmt.Fprintf(os.Stdout, "Response from `ActivityFeedsAPI.UpdateActivityTypeAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActivityTypeAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** |  | 

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **activityTypeUpdateDto** | [**ActivityTypeUpdateDto**](ActivityTypeUpdateDto.md) |  | 

### Return type

[**Envelope**](Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

