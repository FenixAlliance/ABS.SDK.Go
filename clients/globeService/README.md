# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.1.4089
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://fenixalliance.com.co/Support](https://fenixalliance.com.co/Support)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CountriesAPI* | [**ApiV2GlobeServiceCountriesCountryIdCallingCodesGet**](docs/CountriesAPI.md#apiv2globeservicecountriescountryidcallingcodesget) | **Get** /api/v2/GlobeService/Countries/{countryId}/CallingCodes | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesCountryIdCurrenciesGet**](docs/CountriesAPI.md#apiv2globeservicecountriescountryidcurrenciesget) | **Get** /api/v2/GlobeService/Countries/{countryId}/Currencies | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesCountryIdGet**](docs/CountriesAPI.md#apiv2globeservicecountriescountryidget) | **Get** /api/v2/GlobeService/Countries/{countryId} | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdCitiesGet**](docs/CountriesAPI.md#apiv2globeservicecountriescountryidstatescountrystateidcitiesget) | **Get** /api/v2/GlobeService/Countries/{countryId}/States/{countryStateId}/Cities | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesCountryIdStatesCountryStateIdGet**](docs/CountriesAPI.md#apiv2globeservicecountriescountryidstatescountrystateidget) | **Get** /api/v2/GlobeService/Countries/{countryId}/States/{countryStateId} | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesCountryIdStatesGet**](docs/CountriesAPI.md#apiv2globeservicecountriescountryidstatesget) | **Get** /api/v2/GlobeService/Countries/{countryId}/States | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesCountryIdTimezonesGet**](docs/CountriesAPI.md#apiv2globeservicecountriescountryidtimezonesget) | **Get** /api/v2/GlobeService/Countries/{countryId}/Timezones | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesCountryIdTopLevelDomainsGet**](docs/CountriesAPI.md#apiv2globeservicecountriescountryidtopleveldomainsget) | **Get** /api/v2/GlobeService/Countries/{countryId}/TopLevelDomains | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesGet**](docs/CountriesAPI.md#apiv2globeservicecountriesget) | **Get** /api/v2/GlobeService/Countries | 
*CountriesAPI* | [**ApiV2GlobeServiceCountriesSearchGet**](docs/CountriesAPI.md#apiv2globeservicecountriessearchget) | **Get** /api/v2/GlobeService/Countries/Search | 
*CurrenciesAPI* | [**ApiV2GlobeServiceCurrenciesCurrencyIdGet**](docs/CurrenciesAPI.md#apiv2globeservicecurrenciescurrencyidget) | **Get** /api/v2/GlobeService/Currencies/{currencyId} | 
*CurrenciesAPI* | [**ApiV2GlobeServiceCurrenciesGet**](docs/CurrenciesAPI.md#apiv2globeservicecurrenciesget) | **Get** /api/v2/GlobeService/Currencies | 
*LanguagesAPI* | [**ApiV2GlobeServiceLanguagesGet**](docs/LanguagesAPI.md#apiv2globeservicelanguagesget) | **Get** /api/v2/GlobeService/Languages | 
*LanguagesAPI* | [**ApiV2GlobeServiceLanguagesLanguageIdGet**](docs/LanguagesAPI.md#apiv2globeservicelanguageslanguageidget) | **Get** /api/v2/GlobeService/Languages/{languageId} | 
*MigrationsAPI* | [**ApiV2GlobalSystemMigratePost**](docs/MigrationsAPI.md#apiv2globalsystemmigratepost) | **Post** /api/v2/Global/System/Migrate | 
*TimezonesAPI* | [**ApiV2GlobeServiceTimezonesGet**](docs/TimezonesAPI.md#apiv2globeservicetimezonesget) | **Get** /api/v2/GlobeService/Timezones | 
*TimezonesAPI* | [**ApiV2GlobeServiceTimezonesTimeZoneIdGet**](docs/TimezonesAPI.md#apiv2globeservicetimezonestimezoneidget) | **Get** /api/v2/GlobeService/Timezones/{timeZoneId} | 


## Documentation For Models

 - [CityDto](docs/CityDto.md)
 - [CityDtoListEnvelope](docs/CityDtoListEnvelope.md)
 - [CountryCallingCodeDto](docs/CountryCallingCodeDto.md)
 - [CountryCallingCodeDtoListEnvelope](docs/CountryCallingCodeDtoListEnvelope.md)
 - [CountryDto](docs/CountryDto.md)
 - [CountryDtoEnvelope](docs/CountryDtoEnvelope.md)
 - [CountryDtoListEnvelope](docs/CountryDtoListEnvelope.md)
 - [CountryLanguageDto](docs/CountryLanguageDto.md)
 - [CountryLanguageDtoEnvelope](docs/CountryLanguageDtoEnvelope.md)
 - [CountryLanguageDtoListEnvelope](docs/CountryLanguageDtoListEnvelope.md)
 - [CountryStateDto](docs/CountryStateDto.md)
 - [CountryStateDtoEnvelope](docs/CountryStateDtoEnvelope.md)
 - [CountryStateDtoListEnvelope](docs/CountryStateDtoListEnvelope.md)
 - [CountryTopLevelDomainDto](docs/CountryTopLevelDomainDto.md)
 - [CountryTopLevelDomainDtoListEnvelope](docs/CountryTopLevelDomainDtoListEnvelope.md)
 - [CurrencyDto](docs/CurrencyDto.md)
 - [CurrencyDtoEnvelope](docs/CurrencyDtoEnvelope.md)
 - [CurrencyDtoListEnvelope](docs/CurrencyDtoListEnvelope.md)
 - [Error](docs/Error.md)
 - [ErrorEnvelope](docs/ErrorEnvelope.md)
 - [PaymentResponse](docs/PaymentResponse.md)
 - [ResponseStatus](docs/ResponseStatus.md)
 - [TimezoneDto](docs/TimezoneDto.md)
 - [TimezoneDtoEnvelope](docs/TimezoneDtoEnvelope.md)
 - [TimezoneDtoListEnvelope](docs/TimezoneDtoListEnvelope.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Bearer

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Bearer and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		openapi.ContextAPIKeys,
		map[string]openapi.APIKey{
			"Bearer": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@fenix-alliance.com

