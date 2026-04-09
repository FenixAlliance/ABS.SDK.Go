# TenantCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**LegalName** | Pointer to **NullableString** |  | [optional] 
**Email** | **string** |  | 
**Phone** | Pointer to **NullableString** |  | [optional] 
**WebUrl** | Pointer to **NullableString** |  | [optional] 
**Handler** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**Slogan** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | **string** |  | 
**Duns** | Pointer to **NullableString** |  | [optional] 
**TaxId** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**CountryId** | **string** |  | 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**LanguageId** | Pointer to **NullableString** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**BusinessTypeId** | Pointer to **NullableString** |  | [optional] 
**BusinessSegmentId** | Pointer to **NullableString** |  | [optional] 
**BusinessIndustryId** | Pointer to **NullableString** |  | [optional] 
**BusinessSizeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantCreateDto

`func NewTenantCreateDto(name string, email string, currencyId string, countryId string, ) *TenantCreateDto`

NewTenantCreateDto instantiates a new TenantCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantCreateDtoWithDefaults

`func NewTenantCreateDtoWithDefaults() *TenantCreateDto`

NewTenantCreateDtoWithDefaults instantiates a new TenantCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *TenantCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetLegalName

`func (o *TenantCreateDto) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *TenantCreateDto) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *TenantCreateDto) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *TenantCreateDto) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *TenantCreateDto) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *TenantCreateDto) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetEmail

`func (o *TenantCreateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TenantCreateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TenantCreateDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhone

`func (o *TenantCreateDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TenantCreateDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TenantCreateDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TenantCreateDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *TenantCreateDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *TenantCreateDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetWebUrl

`func (o *TenantCreateDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *TenantCreateDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *TenantCreateDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *TenantCreateDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *TenantCreateDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *TenantCreateDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetHandler

`func (o *TenantCreateDto) GetHandler() string`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *TenantCreateDto) GetHandlerOk() (*string, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *TenantCreateDto) SetHandler(v string)`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *TenantCreateDto) HasHandler() bool`

HasHandler returns a boolean if a field has been set.

### SetHandlerNil

`func (o *TenantCreateDto) SetHandlerNil(b bool)`

 SetHandlerNil sets the value for Handler to be an explicit nil

### UnsetHandler
`func (o *TenantCreateDto) UnsetHandler()`

UnsetHandler ensures that no value is present for Handler, not even an explicit nil
### GetAbout

`func (o *TenantCreateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *TenantCreateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *TenantCreateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *TenantCreateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *TenantCreateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *TenantCreateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetSlogan

`func (o *TenantCreateDto) GetSlogan() string`

GetSlogan returns the Slogan field if non-nil, zero value otherwise.

### GetSloganOk

`func (o *TenantCreateDto) GetSloganOk() (*string, bool)`

GetSloganOk returns a tuple with the Slogan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlogan

`func (o *TenantCreateDto) SetSlogan(v string)`

SetSlogan sets Slogan field to given value.

### HasSlogan

`func (o *TenantCreateDto) HasSlogan() bool`

HasSlogan returns a boolean if a field has been set.

### SetSloganNil

`func (o *TenantCreateDto) SetSloganNil(b bool)`

 SetSloganNil sets the value for Slogan to be an explicit nil

### UnsetSlogan
`func (o *TenantCreateDto) UnsetSlogan()`

UnsetSlogan ensures that no value is present for Slogan, not even an explicit nil
### GetCurrencyId

`func (o *TenantCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *TenantCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *TenantCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.


### GetDuns

`func (o *TenantCreateDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *TenantCreateDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *TenantCreateDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *TenantCreateDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *TenantCreateDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *TenantCreateDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetTaxId

`func (o *TenantCreateDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *TenantCreateDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *TenantCreateDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *TenantCreateDto) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *TenantCreateDto) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *TenantCreateDto) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetAvatarUrl

`func (o *TenantCreateDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *TenantCreateDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *TenantCreateDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *TenantCreateDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *TenantCreateDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *TenantCreateDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetCountryId

`func (o *TenantCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *TenantCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *TenantCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.


### GetStateId

`func (o *TenantCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *TenantCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *TenantCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *TenantCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *TenantCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *TenantCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *TenantCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *TenantCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *TenantCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *TenantCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *TenantCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *TenantCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetLanguageId

`func (o *TenantCreateDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *TenantCreateDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *TenantCreateDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *TenantCreateDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *TenantCreateDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *TenantCreateDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetTimezoneId

`func (o *TenantCreateDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *TenantCreateDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *TenantCreateDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *TenantCreateDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *TenantCreateDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *TenantCreateDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetBusinessTypeId

`func (o *TenantCreateDto) GetBusinessTypeId() string`

GetBusinessTypeId returns the BusinessTypeId field if non-nil, zero value otherwise.

### GetBusinessTypeIdOk

`func (o *TenantCreateDto) GetBusinessTypeIdOk() (*string, bool)`

GetBusinessTypeIdOk returns a tuple with the BusinessTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTypeId

`func (o *TenantCreateDto) SetBusinessTypeId(v string)`

SetBusinessTypeId sets BusinessTypeId field to given value.

### HasBusinessTypeId

`func (o *TenantCreateDto) HasBusinessTypeId() bool`

HasBusinessTypeId returns a boolean if a field has been set.

### SetBusinessTypeIdNil

`func (o *TenantCreateDto) SetBusinessTypeIdNil(b bool)`

 SetBusinessTypeIdNil sets the value for BusinessTypeId to be an explicit nil

### UnsetBusinessTypeId
`func (o *TenantCreateDto) UnsetBusinessTypeId()`

UnsetBusinessTypeId ensures that no value is present for BusinessTypeId, not even an explicit nil
### GetBusinessSegmentId

`func (o *TenantCreateDto) GetBusinessSegmentId() string`

GetBusinessSegmentId returns the BusinessSegmentId field if non-nil, zero value otherwise.

### GetBusinessSegmentIdOk

`func (o *TenantCreateDto) GetBusinessSegmentIdOk() (*string, bool)`

GetBusinessSegmentIdOk returns a tuple with the BusinessSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessSegmentId

`func (o *TenantCreateDto) SetBusinessSegmentId(v string)`

SetBusinessSegmentId sets BusinessSegmentId field to given value.

### HasBusinessSegmentId

`func (o *TenantCreateDto) HasBusinessSegmentId() bool`

HasBusinessSegmentId returns a boolean if a field has been set.

### SetBusinessSegmentIdNil

`func (o *TenantCreateDto) SetBusinessSegmentIdNil(b bool)`

 SetBusinessSegmentIdNil sets the value for BusinessSegmentId to be an explicit nil

### UnsetBusinessSegmentId
`func (o *TenantCreateDto) UnsetBusinessSegmentId()`

UnsetBusinessSegmentId ensures that no value is present for BusinessSegmentId, not even an explicit nil
### GetBusinessIndustryId

`func (o *TenantCreateDto) GetBusinessIndustryId() string`

GetBusinessIndustryId returns the BusinessIndustryId field if non-nil, zero value otherwise.

### GetBusinessIndustryIdOk

`func (o *TenantCreateDto) GetBusinessIndustryIdOk() (*string, bool)`

GetBusinessIndustryIdOk returns a tuple with the BusinessIndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessIndustryId

`func (o *TenantCreateDto) SetBusinessIndustryId(v string)`

SetBusinessIndustryId sets BusinessIndustryId field to given value.

### HasBusinessIndustryId

`func (o *TenantCreateDto) HasBusinessIndustryId() bool`

HasBusinessIndustryId returns a boolean if a field has been set.

### SetBusinessIndustryIdNil

`func (o *TenantCreateDto) SetBusinessIndustryIdNil(b bool)`

 SetBusinessIndustryIdNil sets the value for BusinessIndustryId to be an explicit nil

### UnsetBusinessIndustryId
`func (o *TenantCreateDto) UnsetBusinessIndustryId()`

UnsetBusinessIndustryId ensures that no value is present for BusinessIndustryId, not even an explicit nil
### GetBusinessSizeId

`func (o *TenantCreateDto) GetBusinessSizeId() string`

GetBusinessSizeId returns the BusinessSizeId field if non-nil, zero value otherwise.

### GetBusinessSizeIdOk

`func (o *TenantCreateDto) GetBusinessSizeIdOk() (*string, bool)`

GetBusinessSizeIdOk returns a tuple with the BusinessSizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessSizeId

`func (o *TenantCreateDto) SetBusinessSizeId(v string)`

SetBusinessSizeId sets BusinessSizeId field to given value.

### HasBusinessSizeId

`func (o *TenantCreateDto) HasBusinessSizeId() bool`

HasBusinessSizeId returns a boolean if a field has been set.

### SetBusinessSizeIdNil

`func (o *TenantCreateDto) SetBusinessSizeIdNil(b bool)`

 SetBusinessSizeIdNil sets the value for BusinessSizeId to be an explicit nil

### UnsetBusinessSizeId
`func (o *TenantCreateDto) UnsetBusinessSizeId()`

UnsetBusinessSizeId ensures that no value is present for BusinessSizeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


