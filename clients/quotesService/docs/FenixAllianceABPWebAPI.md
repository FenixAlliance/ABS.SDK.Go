# \FenixAllianceABPWebAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForgotPasswordPost**](FenixAllianceABPWebAPI.md#ForgotPasswordPost) | **Post** /forgotPassword | 
[**HealthGet**](FenixAllianceABPWebAPI.md#HealthGet) | **Get** /health | 
[**HelloGet**](FenixAllianceABPWebAPI.md#HelloGet) | **Get** /hello | 
[**LoginPost**](FenixAllianceABPWebAPI.md#LoginPost) | **Post** /login | 
[**Manage2faPost**](FenixAllianceABPWebAPI.md#Manage2faPost) | **Post** /manage/2fa | 
[**ManageInfoGet**](FenixAllianceABPWebAPI.md#ManageInfoGet) | **Get** /manage/info | 
[**ManageInfoPost**](FenixAllianceABPWebAPI.md#ManageInfoPost) | **Post** /manage/info | 
[**MapIdentityApiConfirmEmail**](FenixAllianceABPWebAPI.md#MapIdentityApiConfirmEmail) | **Get** /confirmEmail | 
[**RefreshPost**](FenixAllianceABPWebAPI.md#RefreshPost) | **Post** /refresh | 
[**RegisterPost**](FenixAllianceABPWebAPI.md#RegisterPost) | **Post** /register | 
[**ResendConfirmationEmailPost**](FenixAllianceABPWebAPI.md#ResendConfirmationEmailPost) | **Post** /resendConfirmationEmail | 
[**ResetPasswordPost**](FenixAllianceABPWebAPI.md#ResetPasswordPost) | **Post** /resetPassword | 
[**VersionGet**](FenixAllianceABPWebAPI.md#VersionGet) | **Get** /version | 



## ForgotPasswordPost

> ForgotPasswordPost(ctx).ForgotPasswordRequest(forgotPasswordRequest).Execute()



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
	forgotPasswordRequest := *openapiclient.NewForgotPasswordRequest("Email_example") // ForgotPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FenixAllianceABPWebAPI.ForgotPasswordPost(context.Background()).ForgotPasswordRequest(forgotPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.ForgotPasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordRequest** | [**ForgotPasswordRequest**](ForgotPasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthGet

> HealthGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FenixAllianceABPWebAPI.HealthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.HealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HelloGet

> HelloGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FenixAllianceABPWebAPI.HelloGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.HelloGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHelloGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginPost

> AccessTokenResponse LoginPost(ctx).LoginRequest(loginRequest).UseCookies(useCookies).UseSessionCookies(useSessionCookies).Execute()



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
	loginRequest := *openapiclient.NewLoginRequest("Email_example", "Password_example") // LoginRequest | 
	useCookies := true // bool |  (optional)
	useSessionCookies := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FenixAllianceABPWebAPI.LoginPost(context.Background()).LoginRequest(loginRequest).UseCookies(useCookies).UseSessionCookies(useSessionCookies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.LoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginPost`: AccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `FenixAllianceABPWebAPI.LoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginRequest** | [**LoginRequest**](LoginRequest.md) |  | 
 **useCookies** | **bool** |  | 
 **useSessionCookies** | **bool** |  | 

### Return type

[**AccessTokenResponse**](AccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Manage2faPost

> TwoFactorResponse Manage2faPost(ctx).TwoFactorRequest(twoFactorRequest).Execute()



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
	twoFactorRequest := *openapiclient.NewTwoFactorRequest() // TwoFactorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FenixAllianceABPWebAPI.Manage2faPost(context.Background()).TwoFactorRequest(twoFactorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.Manage2faPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Manage2faPost`: TwoFactorResponse
	fmt.Fprintf(os.Stdout, "Response from `FenixAllianceABPWebAPI.Manage2faPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManage2faPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorRequest** | [**TwoFactorRequest**](TwoFactorRequest.md) |  | 

### Return type

[**TwoFactorResponse**](TwoFactorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageInfoGet

> InfoResponse ManageInfoGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FenixAllianceABPWebAPI.ManageInfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.ManageInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManageInfoGet`: InfoResponse
	fmt.Fprintf(os.Stdout, "Response from `FenixAllianceABPWebAPI.ManageInfoGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiManageInfoGetRequest struct via the builder pattern


### Return type

[**InfoResponse**](InfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageInfoPost

> InfoResponse ManageInfoPost(ctx).InfoRequest(infoRequest).Execute()



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
	infoRequest := *openapiclient.NewInfoRequest() // InfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FenixAllianceABPWebAPI.ManageInfoPost(context.Background()).InfoRequest(infoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.ManageInfoPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ManageInfoPost`: InfoResponse
	fmt.Fprintf(os.Stdout, "Response from `FenixAllianceABPWebAPI.ManageInfoPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManageInfoPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **infoRequest** | [**InfoRequest**](InfoRequest.md) |  | 

### Return type

[**InfoResponse**](InfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MapIdentityApiConfirmEmail

> MapIdentityApiConfirmEmail(ctx).UserId(userId).Code(code).ChangedEmail(changedEmail).Execute()



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
	userId := "userId_example" // string | 
	code := "code_example" // string | 
	changedEmail := "changedEmail_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FenixAllianceABPWebAPI.MapIdentityApiConfirmEmail(context.Background()).UserId(userId).Code(code).ChangedEmail(changedEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.MapIdentityApiConfirmEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMapIdentityApiConfirmEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **code** | **string** |  | 
 **changedEmail** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshPost

> AccessTokenResponse RefreshPost(ctx).RefreshRequest(refreshRequest).Execute()



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
	refreshRequest := *openapiclient.NewRefreshRequest("RefreshToken_example") // RefreshRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FenixAllianceABPWebAPI.RefreshPost(context.Background()).RefreshRequest(refreshRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.RefreshPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshPost`: AccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `FenixAllianceABPWebAPI.RefreshPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshRequest** | [**RefreshRequest**](RefreshRequest.md) |  | 

### Return type

[**AccessTokenResponse**](AccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterPost

> RegisterPost(ctx).RegisterRequest(registerRequest).Execute()



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
	registerRequest := *openapiclient.NewRegisterRequest("Email_example", "Password_example") // RegisterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FenixAllianceABPWebAPI.RegisterPost(context.Background()).RegisterRequest(registerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.RegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerRequest** | [**RegisterRequest**](RegisterRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendConfirmationEmailPost

> ResendConfirmationEmailPost(ctx).ResendConfirmationEmailRequest(resendConfirmationEmailRequest).Execute()



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
	resendConfirmationEmailRequest := *openapiclient.NewResendConfirmationEmailRequest("Email_example") // ResendConfirmationEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FenixAllianceABPWebAPI.ResendConfirmationEmailPost(context.Background()).ResendConfirmationEmailRequest(resendConfirmationEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.ResendConfirmationEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendConfirmationEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resendConfirmationEmailRequest** | [**ResendConfirmationEmailRequest**](ResendConfirmationEmailRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPasswordPost

> ResetPasswordPost(ctx).ResetPasswordRequest(resetPasswordRequest).Execute()



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
	resetPasswordRequest := *openapiclient.NewResetPasswordRequest("Email_example", "ResetCode_example", "NewPassword_example") // ResetPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FenixAllianceABPWebAPI.ResetPasswordPost(context.Background()).ResetPasswordRequest(resetPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.ResetPasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionGet

> VersionGet(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FenixAllianceABPWebAPI.VersionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FenixAllianceABPWebAPI.VersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

