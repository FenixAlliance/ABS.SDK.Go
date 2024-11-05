# \CountriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV2GlobeServiceCountriesCountryIdCallingCodesGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesCountryIdCallingCodesGet) | **Get** /api/v2/GlobeService/Countries/{countryId}/CallingCodes | 
[**ApiV2GlobeServiceCountriesCountryIdCurrenciesGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesCountryIdCurrenciesGet) | **Get** /api/v2/GlobeService/Countries/{countryId}/Currencies | 
[**ApiV2GlobeServiceCountriesCountryIdGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesCountryIdGet) | **Get** /api/v2/GlobeService/Countries/{countryId} | 
[**ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet) | **Get** /api/v2/GlobeService/Countries/{countryId}/States/{countryStateId}/Cities | 
[**ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet) | **Get** /api/v2/GlobeService/Countries/{countryId}/States/{countryStateId} | 
[**ApiV2GlobeServiceCountriesCountryIdStatesGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesCountryIdStatesGet) | **Get** /api/v2/GlobeService/Countries/{countryId}/States | 
[**ApiV2GlobeServiceCountriesCountryIdTimezonesGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesCountryIdTimezonesGet) | **Get** /api/v2/GlobeService/Countries/{countryId}/Timezones | 
[**ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet) | **Get** /api/v2/GlobeService/Countries/{countryId}/TopLevelDomains | 
[**ApiV2GlobeServiceCountriesGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesGet) | **Get** /api/v2/GlobeService/Countries | 
[**ApiV2GlobeServiceCountriesSearchGet**](CountriesAPI.md#ApiV2GlobeServiceCountriesSearchGet) | **Get** /api/v2/GlobeService/Countries/Search | 



## ApiV2GlobeServiceCountriesCountryIdCallingCodesGet

> CountryCallingCodeDtoListEnvelope ApiV2GlobeServiceCountriesCountryIdCallingCodesGet(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryId := "countryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesCountryIdCallingCodesGet(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdCallingCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesCountryIdCallingCodesGet`: CountryCallingCodeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdCallingCodesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesCountryIdCallingCodesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryCallingCodeDtoListEnvelope**](CountryCallingCodeDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesCountryIdCurrenciesGet

> CurrencyDtoListEnvelope ApiV2GlobeServiceCountriesCountryIdCurrenciesGet(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryId := "countryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesCountryIdCurrenciesGet(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdCurrenciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesCountryIdCurrenciesGet`: CurrencyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdCurrenciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesCountryIdCurrenciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CurrencyDtoListEnvelope**](CurrencyDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesCountryIdGet

> CountryDtoEnvelope ApiV2GlobeServiceCountriesCountryIdGet(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryId := "countryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesCountryIdGet(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesCountryIdGet`: CountryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesCountryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryDtoEnvelope**](CountryDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet

> CityDtoListEnvelope ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet(ctx, countryStateId, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryStateId := "countryStateId_example" // string | 
	countryId := "countryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet(context.Background(), countryStateId, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet`: CityDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryStateId** | **string** |  | 
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CityDtoListEnvelope**](CityDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet

> CountryStateDtoEnvelope ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet(ctx, countryStateId, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryStateId := "countryStateId_example" // string | 
	countryId := "countryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet(context.Background(), countryStateId, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet`: CountryStateDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryStateId** | **string** |  | 
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryStateDtoEnvelope**](CountryStateDtoEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesCountryIdStatesGet

> CountryStateDtoListEnvelope ApiV2GlobeServiceCountriesCountryIdStatesGet(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryId := "countryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesGet(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesCountryIdStatesGet`: CountryStateDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdStatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesCountryIdStatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryStateDtoListEnvelope**](CountryStateDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesCountryIdTimezonesGet

> TimezoneDtoListEnvelope ApiV2GlobeServiceCountriesCountryIdTimezonesGet(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryId := "countryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesCountryIdTimezonesGet(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdTimezonesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesCountryIdTimezonesGet`: TimezoneDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdTimezonesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesCountryIdTimezonesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TimezoneDtoListEnvelope**](TimezoneDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet

> CountryTopLevelDomainDtoListEnvelope ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryId := "countryId_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet`: CountryTopLevelDomainDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryTopLevelDomainDtoListEnvelope**](CountryTopLevelDomainDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesGet

> CountryDtoListEnvelope ApiV2GlobeServiceCountriesGet(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesGet(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesGet`: CountryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryDtoListEnvelope**](CountryDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV2GlobeServiceCountriesSearchGet

> CountryDtoListEnvelope ApiV2GlobeServiceCountriesSearchGet(ctx).CountryName(countryName).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()



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
	countryName := "countryName_example" // string | 
	apiVersion := "apiVersion_example" // string |  (optional)
	xApiVersion := "xApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountriesAPI.ApiV2GlobeServiceCountriesSearchGet(context.Background()).CountryName(countryName).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.ApiV2GlobeServiceCountriesSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV2GlobeServiceCountriesSearchGet`: CountryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.ApiV2GlobeServiceCountriesSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV2GlobeServiceCountriesSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryName** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryDtoListEnvelope**](CountryDtoListEnvelope.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

