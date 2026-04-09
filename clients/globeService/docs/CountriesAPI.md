# \CountriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountCountries**](CountriesAPI.md#CountCountries) | **Get** /api/v2/GlobeService/Countries/Count | Count countries
[**GetAllCountries**](CountriesAPI.md#GetAllCountries) | **Get** /api/v2/GlobeService/Countries | Get all countries
[**GetCallingCodesByCountryIdAsync**](CountriesAPI.md#GetCallingCodesByCountryIdAsync) | **Get** /api/v2/GlobeService/Countries/{countryId}/CallingCodes | Get calling codes for a country
[**GetCitiesByCountryStateIdAsync**](CountriesAPI.md#GetCitiesByCountryStateIdAsync) | **Get** /api/v2/GlobeService/Countries/{countryId}/States/{countryStateId}/Cities | Get cities for a state
[**GetCountryById**](CountriesAPI.md#GetCountryById) | **Get** /api/v2/GlobeService/Countries/{countryId} | Get country by ID
[**GetCountryStateByIdAsync**](CountriesAPI.md#GetCountryStateByIdAsync) | **Get** /api/v2/GlobeService/Countries/{countryId}/States/{countryStateId} | Get state by ID
[**GetCountryStatesAsync**](CountriesAPI.md#GetCountryStatesAsync) | **Get** /api/v2/GlobeService/Countries/{countryId}/States | Get states for a country
[**GetEnabledCurrenciesByCountryIdAsync**](CountriesAPI.md#GetEnabledCurrenciesByCountryIdAsync) | **Get** /api/v2/GlobeService/Countries/{countryId}/Currencies | Get currencies for a country
[**GetTimeZonesByCountryIdAsync**](CountriesAPI.md#GetTimeZonesByCountryIdAsync) | **Get** /api/v2/GlobeService/Countries/{countryId}/Timezones | Get timezones for a country
[**GetTopLevelDomainsByCountryIdAsync**](CountriesAPI.md#GetTopLevelDomainsByCountryIdAsync) | **Get** /api/v2/GlobeService/Countries/{countryId}/TopLevelDomains | Get top-level domains for a country
[**SearchCountriesByNameAsync**](CountriesAPI.md#SearchCountriesByNameAsync) | **Get** /api/v2/GlobeService/Countries/Search | Search countries by name



## CountCountries

> Int32Envelope CountCountries(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Count countries



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
	resp, r, err := apiClient.CountriesAPI.CountCountries(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.CountCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountCountries`: Int32Envelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.CountCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**Int32Envelope**](Int32Envelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllCountries

> CountryDtoListEnvelope GetAllCountries(ctx).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get all countries



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
	resp, r, err := apiClient.CountriesAPI.GetAllCountries(context.Background()).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetAllCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCountries`: CountryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetAllCountries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCountriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryDtoListEnvelope**](CountryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallingCodesByCountryIdAsync

> CountryCallingCodeDtoListEnvelope GetCallingCodesByCountryIdAsync(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get calling codes for a country



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
	resp, r, err := apiClient.CountriesAPI.GetCallingCodesByCountryIdAsync(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCallingCodesByCountryIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallingCodesByCountryIdAsync`: CountryCallingCodeDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCallingCodesByCountryIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallingCodesByCountryIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryCallingCodeDtoListEnvelope**](CountryCallingCodeDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCitiesByCountryStateIdAsync

> CityDtoListEnvelope GetCitiesByCountryStateIdAsync(ctx, countryStateId, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get cities for a state



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
	resp, r, err := apiClient.CountriesAPI.GetCitiesByCountryStateIdAsync(context.Background(), countryStateId, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCitiesByCountryStateIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCitiesByCountryStateIdAsync`: CityDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCitiesByCountryStateIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryStateId** | **string** |  | 
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCitiesByCountryStateIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CityDtoListEnvelope**](CityDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryById

> CountryDtoEnvelope GetCountryById(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get country by ID



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
	resp, r, err := apiClient.CountriesAPI.GetCountryById(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryById`: CountryDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryDtoEnvelope**](CountryDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryStateByIdAsync

> CountryStateDtoEnvelope GetCountryStateByIdAsync(ctx, countryStateId, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get state by ID



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
	resp, r, err := apiClient.CountriesAPI.GetCountryStateByIdAsync(context.Background(), countryStateId, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryStateByIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryStateByIdAsync`: CountryStateDtoEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryStateByIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryStateId** | **string** |  | 
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryStateByIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryStateDtoEnvelope**](CountryStateDtoEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountryStatesAsync

> CountryStateDtoListEnvelope GetCountryStatesAsync(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get states for a country



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
	resp, r, err := apiClient.CountriesAPI.GetCountryStatesAsync(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetCountryStatesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountryStatesAsync`: CountryStateDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetCountryStatesAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountryStatesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryStateDtoListEnvelope**](CountryStateDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnabledCurrenciesByCountryIdAsync

> CurrencyDtoListEnvelope GetEnabledCurrenciesByCountryIdAsync(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get currencies for a country



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
	resp, r, err := apiClient.CountriesAPI.GetEnabledCurrenciesByCountryIdAsync(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetEnabledCurrenciesByCountryIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnabledCurrenciesByCountryIdAsync`: CurrencyDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetEnabledCurrenciesByCountryIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledCurrenciesByCountryIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CurrencyDtoListEnvelope**](CurrencyDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimeZonesByCountryIdAsync

> TimezoneDtoListEnvelope GetTimeZonesByCountryIdAsync(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get timezones for a country



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
	resp, r, err := apiClient.CountriesAPI.GetTimeZonesByCountryIdAsync(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetTimeZonesByCountryIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimeZonesByCountryIdAsync`: TimezoneDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetTimeZonesByCountryIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimeZonesByCountryIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**TimezoneDtoListEnvelope**](TimezoneDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopLevelDomainsByCountryIdAsync

> CountryTopLevelDomainDtoListEnvelope GetTopLevelDomainsByCountryIdAsync(ctx, countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Get top-level domains for a country



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
	resp, r, err := apiClient.CountriesAPI.GetTopLevelDomainsByCountryIdAsync(context.Background(), countryId).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.GetTopLevelDomainsByCountryIdAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTopLevelDomainsByCountryIdAsync`: CountryTopLevelDomainDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.GetTopLevelDomainsByCountryIdAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopLevelDomainsByCountryIdAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryTopLevelDomainDtoListEnvelope**](CountryTopLevelDomainDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCountriesByNameAsync

> CountryDtoListEnvelope SearchCountriesByNameAsync(ctx).CountryName(countryName).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()

Search countries by name



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
	resp, r, err := apiClient.CountriesAPI.SearchCountriesByNameAsync(context.Background()).CountryName(countryName).ApiVersion(apiVersion).XApiVersion(xApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountriesAPI.SearchCountriesByNameAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCountriesByNameAsync`: CountryDtoListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `CountriesAPI.SearchCountriesByNameAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCountriesByNameAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryName** | **string** |  | 
 **apiVersion** | **string** |  | 
 **xApiVersion** | **string** |  | 

### Return type

[**CountryDtoListEnvelope**](CountryDtoListEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

