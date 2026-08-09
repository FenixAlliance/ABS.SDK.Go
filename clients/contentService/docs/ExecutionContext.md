# ExecutionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthenticated** | Pointer to **bool** |  | [optional] 
**CurrentCartId** | Pointer to **NullableString** |  | [optional] [readonly] 
**CurrentUserId** | Pointer to **NullableString** |  | [optional] [readonly] 
**CurrentTenantId** | Pointer to **NullableString** |  | [optional] [readonly] 
**CurrentPortalId** | Pointer to **NullableString** |  | [optional] [readonly] 
**CurrentEnrollmentId** | Pointer to **NullableString** |  | [optional] [readonly] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] [readonly] 
**PageSize** | Pointer to **int32** |  | [optional] 
**DateFormat** | Pointer to **NullableString** |  | [optional] 
**CurrencyFormat** | Pointer to **NullableString** |  | [optional] 
**DateTimeFormat** | Pointer to **NullableString** |  | [optional] 
**ToDateDataSummaries** | Pointer to **time.Time** |  | [optional] 
**FromDateDataSummaries** | Pointer to **time.Time** |  | [optional] 
**Authorization** | Pointer to [**AuthResult**](AuthResult.md) |  | [optional] 
**User** | Pointer to [**ExtendedUserDto**](ExtendedUserDto.md) |  | [optional] 
**CurrentTenant** | Pointer to [**ExtendedTenantDto**](ExtendedTenantDto.md) |  | [optional] 
**CurrentEnrollment** | Pointer to [**TenantEnrollmentDto**](TenantEnrollmentDto.md) |  | [optional] 
**SelectedTenantMappings** | Pointer to [**CrmContext**](CrmContext.md) |  | [optional] 
**PortalOwnerMappings** | Pointer to [**CrmContext**](CrmContext.md) |  | [optional] 
**RootTenantMappings** | Pointer to [**CrmContext**](CrmContext.md) |  | [optional] 
**Cart** | Pointer to [**CartDto**](CartDto.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyDto**](CurrencyDto.md) |  | [optional] 
**ForexRates** | Pointer to [**ForexRatesDto**](ForexRatesDto.md) |  | [optional] 
**ExchangeRate** | Pointer to [**Money**](Money.md) |  | [optional] 
**Country** | Pointer to [**CountryDto**](CountryDto.md) |  | [optional] 
**RootTenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 
**CurrentPortal** | Pointer to [**WebPortalDto**](WebPortalDto.md) |  | [optional] 
**Tenants** | Pointer to [**[]ExtendedTenantDto**](ExtendedTenantDto.md) |  | [optional] 
**Enrollments** | Pointer to [**[]ExtendedTenantEnrollmentDto**](ExtendedTenantEnrollmentDto.md) |  | [optional] 
**AvailablePortals** | Pointer to [**[]WebPortalDto**](WebPortalDto.md) |  | [optional] 
**Invitations** | Pointer to [**[]ExtendedInviteDto**](ExtendedInviteDto.md) |  | [optional] 
**GrantedPermissions** | Pointer to **[]string** |  | [optional] 
**AccessibleFeatures** | Pointer to [**[]SuiteLicenseFeatureDto**](SuiteLicenseFeatureDto.md) |  | [optional] 
**CultureName** | Pointer to **NullableString** |  | [optional] [readonly] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExecutionContext

`func NewExecutionContext() *ExecutionContext`

NewExecutionContext instantiates a new ExecutionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionContextWithDefaults

`func NewExecutionContextWithDefaults() *ExecutionContext`

NewExecutionContextWithDefaults instantiates a new ExecutionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthenticated

`func (o *ExecutionContext) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *ExecutionContext) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *ExecutionContext) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *ExecutionContext) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.

### GetCurrentCartId

`func (o *ExecutionContext) GetCurrentCartId() string`

GetCurrentCartId returns the CurrentCartId field if non-nil, zero value otherwise.

### GetCurrentCartIdOk

`func (o *ExecutionContext) GetCurrentCartIdOk() (*string, bool)`

GetCurrentCartIdOk returns a tuple with the CurrentCartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCartId

`func (o *ExecutionContext) SetCurrentCartId(v string)`

SetCurrentCartId sets CurrentCartId field to given value.

### HasCurrentCartId

`func (o *ExecutionContext) HasCurrentCartId() bool`

HasCurrentCartId returns a boolean if a field has been set.

### SetCurrentCartIdNil

`func (o *ExecutionContext) SetCurrentCartIdNil(b bool)`

 SetCurrentCartIdNil sets the value for CurrentCartId to be an explicit nil

### UnsetCurrentCartId
`func (o *ExecutionContext) UnsetCurrentCartId()`

UnsetCurrentCartId ensures that no value is present for CurrentCartId, not even an explicit nil
### GetCurrentUserId

`func (o *ExecutionContext) GetCurrentUserId() string`

GetCurrentUserId returns the CurrentUserId field if non-nil, zero value otherwise.

### GetCurrentUserIdOk

`func (o *ExecutionContext) GetCurrentUserIdOk() (*string, bool)`

GetCurrentUserIdOk returns a tuple with the CurrentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserId

`func (o *ExecutionContext) SetCurrentUserId(v string)`

SetCurrentUserId sets CurrentUserId field to given value.

### HasCurrentUserId

`func (o *ExecutionContext) HasCurrentUserId() bool`

HasCurrentUserId returns a boolean if a field has been set.

### SetCurrentUserIdNil

`func (o *ExecutionContext) SetCurrentUserIdNil(b bool)`

 SetCurrentUserIdNil sets the value for CurrentUserId to be an explicit nil

### UnsetCurrentUserId
`func (o *ExecutionContext) UnsetCurrentUserId()`

UnsetCurrentUserId ensures that no value is present for CurrentUserId, not even an explicit nil
### GetCurrentTenantId

`func (o *ExecutionContext) GetCurrentTenantId() string`

GetCurrentTenantId returns the CurrentTenantId field if non-nil, zero value otherwise.

### GetCurrentTenantIdOk

`func (o *ExecutionContext) GetCurrentTenantIdOk() (*string, bool)`

GetCurrentTenantIdOk returns a tuple with the CurrentTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTenantId

`func (o *ExecutionContext) SetCurrentTenantId(v string)`

SetCurrentTenantId sets CurrentTenantId field to given value.

### HasCurrentTenantId

`func (o *ExecutionContext) HasCurrentTenantId() bool`

HasCurrentTenantId returns a boolean if a field has been set.

### SetCurrentTenantIdNil

`func (o *ExecutionContext) SetCurrentTenantIdNil(b bool)`

 SetCurrentTenantIdNil sets the value for CurrentTenantId to be an explicit nil

### UnsetCurrentTenantId
`func (o *ExecutionContext) UnsetCurrentTenantId()`

UnsetCurrentTenantId ensures that no value is present for CurrentTenantId, not even an explicit nil
### GetCurrentPortalId

`func (o *ExecutionContext) GetCurrentPortalId() string`

GetCurrentPortalId returns the CurrentPortalId field if non-nil, zero value otherwise.

### GetCurrentPortalIdOk

`func (o *ExecutionContext) GetCurrentPortalIdOk() (*string, bool)`

GetCurrentPortalIdOk returns a tuple with the CurrentPortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPortalId

`func (o *ExecutionContext) SetCurrentPortalId(v string)`

SetCurrentPortalId sets CurrentPortalId field to given value.

### HasCurrentPortalId

`func (o *ExecutionContext) HasCurrentPortalId() bool`

HasCurrentPortalId returns a boolean if a field has been set.

### SetCurrentPortalIdNil

`func (o *ExecutionContext) SetCurrentPortalIdNil(b bool)`

 SetCurrentPortalIdNil sets the value for CurrentPortalId to be an explicit nil

### UnsetCurrentPortalId
`func (o *ExecutionContext) UnsetCurrentPortalId()`

UnsetCurrentPortalId ensures that no value is present for CurrentPortalId, not even an explicit nil
### GetCurrentEnrollmentId

`func (o *ExecutionContext) GetCurrentEnrollmentId() string`

GetCurrentEnrollmentId returns the CurrentEnrollmentId field if non-nil, zero value otherwise.

### GetCurrentEnrollmentIdOk

`func (o *ExecutionContext) GetCurrentEnrollmentIdOk() (*string, bool)`

GetCurrentEnrollmentIdOk returns a tuple with the CurrentEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEnrollmentId

`func (o *ExecutionContext) SetCurrentEnrollmentId(v string)`

SetCurrentEnrollmentId sets CurrentEnrollmentId field to given value.

### HasCurrentEnrollmentId

`func (o *ExecutionContext) HasCurrentEnrollmentId() bool`

HasCurrentEnrollmentId returns a boolean if a field has been set.

### SetCurrentEnrollmentIdNil

`func (o *ExecutionContext) SetCurrentEnrollmentIdNil(b bool)`

 SetCurrentEnrollmentIdNil sets the value for CurrentEnrollmentId to be an explicit nil

### UnsetCurrentEnrollmentId
`func (o *ExecutionContext) UnsetCurrentEnrollmentId()`

UnsetCurrentEnrollmentId ensures that no value is present for CurrentEnrollmentId, not even an explicit nil
### GetCurrencyId

`func (o *ExecutionContext) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ExecutionContext) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ExecutionContext) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ExecutionContext) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ExecutionContext) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ExecutionContext) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetPageSize

`func (o *ExecutionContext) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ExecutionContext) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ExecutionContext) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ExecutionContext) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetDateFormat

`func (o *ExecutionContext) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *ExecutionContext) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *ExecutionContext) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *ExecutionContext) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### SetDateFormatNil

`func (o *ExecutionContext) SetDateFormatNil(b bool)`

 SetDateFormatNil sets the value for DateFormat to be an explicit nil

### UnsetDateFormat
`func (o *ExecutionContext) UnsetDateFormat()`

UnsetDateFormat ensures that no value is present for DateFormat, not even an explicit nil
### GetCurrencyFormat

`func (o *ExecutionContext) GetCurrencyFormat() string`

GetCurrencyFormat returns the CurrencyFormat field if non-nil, zero value otherwise.

### GetCurrencyFormatOk

`func (o *ExecutionContext) GetCurrencyFormatOk() (*string, bool)`

GetCurrencyFormatOk returns a tuple with the CurrencyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyFormat

`func (o *ExecutionContext) SetCurrencyFormat(v string)`

SetCurrencyFormat sets CurrencyFormat field to given value.

### HasCurrencyFormat

`func (o *ExecutionContext) HasCurrencyFormat() bool`

HasCurrencyFormat returns a boolean if a field has been set.

### SetCurrencyFormatNil

`func (o *ExecutionContext) SetCurrencyFormatNil(b bool)`

 SetCurrencyFormatNil sets the value for CurrencyFormat to be an explicit nil

### UnsetCurrencyFormat
`func (o *ExecutionContext) UnsetCurrencyFormat()`

UnsetCurrencyFormat ensures that no value is present for CurrencyFormat, not even an explicit nil
### GetDateTimeFormat

`func (o *ExecutionContext) GetDateTimeFormat() string`

GetDateTimeFormat returns the DateTimeFormat field if non-nil, zero value otherwise.

### GetDateTimeFormatOk

`func (o *ExecutionContext) GetDateTimeFormatOk() (*string, bool)`

GetDateTimeFormatOk returns a tuple with the DateTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFormat

`func (o *ExecutionContext) SetDateTimeFormat(v string)`

SetDateTimeFormat sets DateTimeFormat field to given value.

### HasDateTimeFormat

`func (o *ExecutionContext) HasDateTimeFormat() bool`

HasDateTimeFormat returns a boolean if a field has been set.

### SetDateTimeFormatNil

`func (o *ExecutionContext) SetDateTimeFormatNil(b bool)`

 SetDateTimeFormatNil sets the value for DateTimeFormat to be an explicit nil

### UnsetDateTimeFormat
`func (o *ExecutionContext) UnsetDateTimeFormat()`

UnsetDateTimeFormat ensures that no value is present for DateTimeFormat, not even an explicit nil
### GetToDateDataSummaries

`func (o *ExecutionContext) GetToDateDataSummaries() time.Time`

GetToDateDataSummaries returns the ToDateDataSummaries field if non-nil, zero value otherwise.

### GetToDateDataSummariesOk

`func (o *ExecutionContext) GetToDateDataSummariesOk() (*time.Time, bool)`

GetToDateDataSummariesOk returns a tuple with the ToDateDataSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDateDataSummaries

`func (o *ExecutionContext) SetToDateDataSummaries(v time.Time)`

SetToDateDataSummaries sets ToDateDataSummaries field to given value.

### HasToDateDataSummaries

`func (o *ExecutionContext) HasToDateDataSummaries() bool`

HasToDateDataSummaries returns a boolean if a field has been set.

### GetFromDateDataSummaries

`func (o *ExecutionContext) GetFromDateDataSummaries() time.Time`

GetFromDateDataSummaries returns the FromDateDataSummaries field if non-nil, zero value otherwise.

### GetFromDateDataSummariesOk

`func (o *ExecutionContext) GetFromDateDataSummariesOk() (*time.Time, bool)`

GetFromDateDataSummariesOk returns a tuple with the FromDateDataSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateDataSummaries

`func (o *ExecutionContext) SetFromDateDataSummaries(v time.Time)`

SetFromDateDataSummaries sets FromDateDataSummaries field to given value.

### HasFromDateDataSummaries

`func (o *ExecutionContext) HasFromDateDataSummaries() bool`

HasFromDateDataSummaries returns a boolean if a field has been set.

### GetAuthorization

`func (o *ExecutionContext) GetAuthorization() AuthResult`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *ExecutionContext) GetAuthorizationOk() (*AuthResult, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *ExecutionContext) SetAuthorization(v AuthResult)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *ExecutionContext) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetUser

`func (o *ExecutionContext) GetUser() ExtendedUserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExecutionContext) GetUserOk() (*ExtendedUserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExecutionContext) SetUser(v ExtendedUserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ExecutionContext) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCurrentTenant

`func (o *ExecutionContext) GetCurrentTenant() ExtendedTenantDto`

GetCurrentTenant returns the CurrentTenant field if non-nil, zero value otherwise.

### GetCurrentTenantOk

`func (o *ExecutionContext) GetCurrentTenantOk() (*ExtendedTenantDto, bool)`

GetCurrentTenantOk returns a tuple with the CurrentTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTenant

`func (o *ExecutionContext) SetCurrentTenant(v ExtendedTenantDto)`

SetCurrentTenant sets CurrentTenant field to given value.

### HasCurrentTenant

`func (o *ExecutionContext) HasCurrentTenant() bool`

HasCurrentTenant returns a boolean if a field has been set.

### GetCurrentEnrollment

`func (o *ExecutionContext) GetCurrentEnrollment() TenantEnrollmentDto`

GetCurrentEnrollment returns the CurrentEnrollment field if non-nil, zero value otherwise.

### GetCurrentEnrollmentOk

`func (o *ExecutionContext) GetCurrentEnrollmentOk() (*TenantEnrollmentDto, bool)`

GetCurrentEnrollmentOk returns a tuple with the CurrentEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEnrollment

`func (o *ExecutionContext) SetCurrentEnrollment(v TenantEnrollmentDto)`

SetCurrentEnrollment sets CurrentEnrollment field to given value.

### HasCurrentEnrollment

`func (o *ExecutionContext) HasCurrentEnrollment() bool`

HasCurrentEnrollment returns a boolean if a field has been set.

### GetSelectedTenantMappings

`func (o *ExecutionContext) GetSelectedTenantMappings() CrmContext`

GetSelectedTenantMappings returns the SelectedTenantMappings field if non-nil, zero value otherwise.

### GetSelectedTenantMappingsOk

`func (o *ExecutionContext) GetSelectedTenantMappingsOk() (*CrmContext, bool)`

GetSelectedTenantMappingsOk returns a tuple with the SelectedTenantMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedTenantMappings

`func (o *ExecutionContext) SetSelectedTenantMappings(v CrmContext)`

SetSelectedTenantMappings sets SelectedTenantMappings field to given value.

### HasSelectedTenantMappings

`func (o *ExecutionContext) HasSelectedTenantMappings() bool`

HasSelectedTenantMappings returns a boolean if a field has been set.

### GetPortalOwnerMappings

`func (o *ExecutionContext) GetPortalOwnerMappings() CrmContext`

GetPortalOwnerMappings returns the PortalOwnerMappings field if non-nil, zero value otherwise.

### GetPortalOwnerMappingsOk

`func (o *ExecutionContext) GetPortalOwnerMappingsOk() (*CrmContext, bool)`

GetPortalOwnerMappingsOk returns a tuple with the PortalOwnerMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalOwnerMappings

`func (o *ExecutionContext) SetPortalOwnerMappings(v CrmContext)`

SetPortalOwnerMappings sets PortalOwnerMappings field to given value.

### HasPortalOwnerMappings

`func (o *ExecutionContext) HasPortalOwnerMappings() bool`

HasPortalOwnerMappings returns a boolean if a field has been set.

### GetRootTenantMappings

`func (o *ExecutionContext) GetRootTenantMappings() CrmContext`

GetRootTenantMappings returns the RootTenantMappings field if non-nil, zero value otherwise.

### GetRootTenantMappingsOk

`func (o *ExecutionContext) GetRootTenantMappingsOk() (*CrmContext, bool)`

GetRootTenantMappingsOk returns a tuple with the RootTenantMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTenantMappings

`func (o *ExecutionContext) SetRootTenantMappings(v CrmContext)`

SetRootTenantMappings sets RootTenantMappings field to given value.

### HasRootTenantMappings

`func (o *ExecutionContext) HasRootTenantMappings() bool`

HasRootTenantMappings returns a boolean if a field has been set.

### GetCart

`func (o *ExecutionContext) GetCart() CartDto`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *ExecutionContext) GetCartOk() (*CartDto, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *ExecutionContext) SetCart(v CartDto)`

SetCart sets Cart field to given value.

### HasCart

`func (o *ExecutionContext) HasCart() bool`

HasCart returns a boolean if a field has been set.

### GetCurrency

`func (o *ExecutionContext) GetCurrency() CurrencyDto`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExecutionContext) GetCurrencyOk() (*CurrencyDto, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExecutionContext) SetCurrency(v CurrencyDto)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ExecutionContext) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetForexRates

`func (o *ExecutionContext) GetForexRates() ForexRatesDto`

GetForexRates returns the ForexRates field if non-nil, zero value otherwise.

### GetForexRatesOk

`func (o *ExecutionContext) GetForexRatesOk() (*ForexRatesDto, bool)`

GetForexRatesOk returns a tuple with the ForexRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRates

`func (o *ExecutionContext) SetForexRates(v ForexRatesDto)`

SetForexRates sets ForexRates field to given value.

### HasForexRates

`func (o *ExecutionContext) HasForexRates() bool`

HasForexRates returns a boolean if a field has been set.

### GetExchangeRate

`func (o *ExecutionContext) GetExchangeRate() Money`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *ExecutionContext) GetExchangeRateOk() (*Money, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *ExecutionContext) SetExchangeRate(v Money)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *ExecutionContext) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetCountry

`func (o *ExecutionContext) GetCountry() CountryDto`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ExecutionContext) GetCountryOk() (*CountryDto, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ExecutionContext) SetCountry(v CountryDto)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ExecutionContext) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRootTenant

`func (o *ExecutionContext) GetRootTenant() TenantDto`

GetRootTenant returns the RootTenant field if non-nil, zero value otherwise.

### GetRootTenantOk

`func (o *ExecutionContext) GetRootTenantOk() (*TenantDto, bool)`

GetRootTenantOk returns a tuple with the RootTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTenant

`func (o *ExecutionContext) SetRootTenant(v TenantDto)`

SetRootTenant sets RootTenant field to given value.

### HasRootTenant

`func (o *ExecutionContext) HasRootTenant() bool`

HasRootTenant returns a boolean if a field has been set.

### GetCurrentPortal

`func (o *ExecutionContext) GetCurrentPortal() WebPortalDto`

GetCurrentPortal returns the CurrentPortal field if non-nil, zero value otherwise.

### GetCurrentPortalOk

`func (o *ExecutionContext) GetCurrentPortalOk() (*WebPortalDto, bool)`

GetCurrentPortalOk returns a tuple with the CurrentPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPortal

`func (o *ExecutionContext) SetCurrentPortal(v WebPortalDto)`

SetCurrentPortal sets CurrentPortal field to given value.

### HasCurrentPortal

`func (o *ExecutionContext) HasCurrentPortal() bool`

HasCurrentPortal returns a boolean if a field has been set.

### GetTenants

`func (o *ExecutionContext) GetTenants() []ExtendedTenantDto`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ExecutionContext) GetTenantsOk() (*[]ExtendedTenantDto, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ExecutionContext) SetTenants(v []ExtendedTenantDto)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ExecutionContext) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### SetTenantsNil

`func (o *ExecutionContext) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *ExecutionContext) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil
### GetEnrollments

`func (o *ExecutionContext) GetEnrollments() []ExtendedTenantEnrollmentDto`

GetEnrollments returns the Enrollments field if non-nil, zero value otherwise.

### GetEnrollmentsOk

`func (o *ExecutionContext) GetEnrollmentsOk() (*[]ExtendedTenantEnrollmentDto, bool)`

GetEnrollmentsOk returns a tuple with the Enrollments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollments

`func (o *ExecutionContext) SetEnrollments(v []ExtendedTenantEnrollmentDto)`

SetEnrollments sets Enrollments field to given value.

### HasEnrollments

`func (o *ExecutionContext) HasEnrollments() bool`

HasEnrollments returns a boolean if a field has been set.

### SetEnrollmentsNil

`func (o *ExecutionContext) SetEnrollmentsNil(b bool)`

 SetEnrollmentsNil sets the value for Enrollments to be an explicit nil

### UnsetEnrollments
`func (o *ExecutionContext) UnsetEnrollments()`

UnsetEnrollments ensures that no value is present for Enrollments, not even an explicit nil
### GetAvailablePortals

`func (o *ExecutionContext) GetAvailablePortals() []WebPortalDto`

GetAvailablePortals returns the AvailablePortals field if non-nil, zero value otherwise.

### GetAvailablePortalsOk

`func (o *ExecutionContext) GetAvailablePortalsOk() (*[]WebPortalDto, bool)`

GetAvailablePortalsOk returns a tuple with the AvailablePortals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePortals

`func (o *ExecutionContext) SetAvailablePortals(v []WebPortalDto)`

SetAvailablePortals sets AvailablePortals field to given value.

### HasAvailablePortals

`func (o *ExecutionContext) HasAvailablePortals() bool`

HasAvailablePortals returns a boolean if a field has been set.

### SetAvailablePortalsNil

`func (o *ExecutionContext) SetAvailablePortalsNil(b bool)`

 SetAvailablePortalsNil sets the value for AvailablePortals to be an explicit nil

### UnsetAvailablePortals
`func (o *ExecutionContext) UnsetAvailablePortals()`

UnsetAvailablePortals ensures that no value is present for AvailablePortals, not even an explicit nil
### GetInvitations

`func (o *ExecutionContext) GetInvitations() []ExtendedInviteDto`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *ExecutionContext) GetInvitationsOk() (*[]ExtendedInviteDto, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *ExecutionContext) SetInvitations(v []ExtendedInviteDto)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *ExecutionContext) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### SetInvitationsNil

`func (o *ExecutionContext) SetInvitationsNil(b bool)`

 SetInvitationsNil sets the value for Invitations to be an explicit nil

### UnsetInvitations
`func (o *ExecutionContext) UnsetInvitations()`

UnsetInvitations ensures that no value is present for Invitations, not even an explicit nil
### GetGrantedPermissions

`func (o *ExecutionContext) GetGrantedPermissions() []string`

GetGrantedPermissions returns the GrantedPermissions field if non-nil, zero value otherwise.

### GetGrantedPermissionsOk

`func (o *ExecutionContext) GetGrantedPermissionsOk() (*[]string, bool)`

GetGrantedPermissionsOk returns a tuple with the GrantedPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedPermissions

`func (o *ExecutionContext) SetGrantedPermissions(v []string)`

SetGrantedPermissions sets GrantedPermissions field to given value.

### HasGrantedPermissions

`func (o *ExecutionContext) HasGrantedPermissions() bool`

HasGrantedPermissions returns a boolean if a field has been set.

### SetGrantedPermissionsNil

`func (o *ExecutionContext) SetGrantedPermissionsNil(b bool)`

 SetGrantedPermissionsNil sets the value for GrantedPermissions to be an explicit nil

### UnsetGrantedPermissions
`func (o *ExecutionContext) UnsetGrantedPermissions()`

UnsetGrantedPermissions ensures that no value is present for GrantedPermissions, not even an explicit nil
### GetAccessibleFeatures

`func (o *ExecutionContext) GetAccessibleFeatures() []SuiteLicenseFeatureDto`

GetAccessibleFeatures returns the AccessibleFeatures field if non-nil, zero value otherwise.

### GetAccessibleFeaturesOk

`func (o *ExecutionContext) GetAccessibleFeaturesOk() (*[]SuiteLicenseFeatureDto, bool)`

GetAccessibleFeaturesOk returns a tuple with the AccessibleFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibleFeatures

`func (o *ExecutionContext) SetAccessibleFeatures(v []SuiteLicenseFeatureDto)`

SetAccessibleFeatures sets AccessibleFeatures field to given value.

### HasAccessibleFeatures

`func (o *ExecutionContext) HasAccessibleFeatures() bool`

HasAccessibleFeatures returns a boolean if a field has been set.

### SetAccessibleFeaturesNil

`func (o *ExecutionContext) SetAccessibleFeaturesNil(b bool)`

 SetAccessibleFeaturesNil sets the value for AccessibleFeatures to be an explicit nil

### UnsetAccessibleFeatures
`func (o *ExecutionContext) UnsetAccessibleFeatures()`

UnsetAccessibleFeatures ensures that no value is present for AccessibleFeatures, not even an explicit nil
### GetCultureName

`func (o *ExecutionContext) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *ExecutionContext) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *ExecutionContext) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.

### HasCultureName

`func (o *ExecutionContext) HasCultureName() bool`

HasCultureName returns a boolean if a field has been set.

### SetCultureNameNil

`func (o *ExecutionContext) SetCultureNameNil(b bool)`

 SetCultureNameNil sets the value for CultureName to be an explicit nil

### UnsetCultureName
`func (o *ExecutionContext) UnsetCultureName()`

UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
### GetTimezoneId

`func (o *ExecutionContext) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ExecutionContext) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ExecutionContext) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ExecutionContext) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ExecutionContext) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ExecutionContext) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


