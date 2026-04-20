# \EmailsAPI

All URIs are relative to *https://absuite.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminPreviewBasicEmailTemplate**](EmailsAPI.md#AdminPreviewBasicEmailTemplate) | **Post** /api/v2/SystemService/Emails/Preview | Preview a rendered basic email template.
[**AdminSendBasicEmail**](EmailsAPI.md#AdminSendBasicEmail) | **Post** /api/v2/SystemService/Emails/SendBasic | Send a basic transactional email to recipients.



## AdminPreviewBasicEmailTemplate

> AdminPreviewBasicEmailTemplate(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).ObjectEmailDispatchRequest(objectEmailDispatchRequest).Execute()

Preview a rendered basic email template.



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	objectEmailDispatchRequest := *openapiclient.NewObjectEmailDispatchRequest("Title_example", "Message_example", "Culture_example", "UiCulture_example", []string{"Recipients_example"}) // ObjectEmailDispatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailsAPI.AdminPreviewBasicEmailTemplate(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).ObjectEmailDispatchRequest(objectEmailDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.AdminPreviewBasicEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminPreviewBasicEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **objectEmailDispatchRequest** | [**ObjectEmailDispatchRequest**](ObjectEmailDispatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSendBasicEmail

> TenantDtoListEnvelope AdminSendBasicEmail(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).ObjectEmailDispatchRequest(objectEmailDispatchRequest).Execute()

Send a basic transactional email to recipients.



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
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)
	objectEmailDispatchRequest := *openapiclient.NewObjectEmailDispatchRequest("Title_example", "Message_example", "Culture_example", "UiCulture_example", []string{"Recipients_example"}) // ObjectEmailDispatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.AdminSendBasicEmail(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).ObjectEmailDispatchRequest(objectEmailDispatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.AdminSendBasicEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminSendBasicEmail`: TenantDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.AdminSendBasicEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSendBasicEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 
 **objectEmailDispatchRequest** | [**ObjectEmailDispatchRequest**](ObjectEmailDispatchRequest.md) |  | 

### Return type

[**TenantDtoListEnvelope**](TenantDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

