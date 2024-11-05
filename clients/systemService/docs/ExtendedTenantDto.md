# ExtendedTenantDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] [readonly] 
**TaxId** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**WalletId** | Pointer to **NullableString** |  | [optional] 
**SocialFeedId** | Pointer to **NullableString** |  | [optional] 
**BusinessIndustryId** | Pointer to **NullableString** |  | [optional] 
**BusinessSegmentId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**LanguageId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Duns** | Pointer to **NullableString** |  | [optional] 
**Slogan** | Pointer to **NullableString** |  | [optional] 
**LegalName** | Pointer to **NullableString** |  | [optional] 
**CoverUrl** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**WebUrl** | Pointer to **NullableString** |  | [optional] 
**FacebookUrl** | Pointer to **NullableString** |  | [optional] 
**TwitterUrl** | Pointer to **NullableString** |  | [optional] 
**GitHubUrl** | Pointer to **NullableString** |  | [optional] 
**LinkedInUrl** | Pointer to **NullableString** |  | [optional] 
**InstagramUrl** | Pointer to **NullableString** |  | [optional] 
**YouTubeUrl** | Pointer to **NullableString** |  | [optional] 
**WhatsAppNumber** | Pointer to **NullableString** |  | [optional] 
**SupportPhoneNumber** | Pointer to **NullableString** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**BusinessName** | Pointer to **NullableString** |  | [optional] [readonly] 
**BusinessLegalName** | Pointer to **NullableString** |  | [optional] [readonly] 
**TwitterUsername** | Pointer to **NullableString** |  | [optional] 
**Cart** | Pointer to [**CartDto**](CartDto.md) |  | [optional] 
**Wallet** | Pointer to [**WalletDto**](WalletDto.md) |  | [optional] 
**SocialProfile** | Pointer to [**SocialProfileDto**](SocialProfileDto.md) |  | [optional] 

## Methods

### NewExtendedTenantDto

`func NewExtendedTenantDto() *ExtendedTenantDto`

NewExtendedTenantDto instantiates a new ExtendedTenantDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedTenantDtoWithDefaults

`func NewExtendedTenantDtoWithDefaults() *ExtendedTenantDto`

NewExtendedTenantDtoWithDefaults instantiates a new ExtendedTenantDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedTenantDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedTenantDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedTenantDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedTenantDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedTenantDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedTenantDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ExtendedTenantDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedTenantDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedTenantDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtendedTenantDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ExtendedTenantDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ExtendedTenantDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQualifiedName

`func (o *ExtendedTenantDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *ExtendedTenantDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *ExtendedTenantDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *ExtendedTenantDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *ExtendedTenantDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *ExtendedTenantDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetTaxId

`func (o *ExtendedTenantDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *ExtendedTenantDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *ExtendedTenantDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *ExtendedTenantDto) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *ExtendedTenantDto) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *ExtendedTenantDto) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetAbout

`func (o *ExtendedTenantDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ExtendedTenantDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ExtendedTenantDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *ExtendedTenantDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *ExtendedTenantDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *ExtendedTenantDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetWalletId

`func (o *ExtendedTenantDto) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExtendedTenantDto) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExtendedTenantDto) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *ExtendedTenantDto) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *ExtendedTenantDto) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *ExtendedTenantDto) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetSocialFeedId

`func (o *ExtendedTenantDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *ExtendedTenantDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *ExtendedTenantDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *ExtendedTenantDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *ExtendedTenantDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *ExtendedTenantDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil
### GetBusinessIndustryId

`func (o *ExtendedTenantDto) GetBusinessIndustryId() string`

GetBusinessIndustryId returns the BusinessIndustryId field if non-nil, zero value otherwise.

### GetBusinessIndustryIdOk

`func (o *ExtendedTenantDto) GetBusinessIndustryIdOk() (*string, bool)`

GetBusinessIndustryIdOk returns a tuple with the BusinessIndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessIndustryId

`func (o *ExtendedTenantDto) SetBusinessIndustryId(v string)`

SetBusinessIndustryId sets BusinessIndustryId field to given value.

### HasBusinessIndustryId

`func (o *ExtendedTenantDto) HasBusinessIndustryId() bool`

HasBusinessIndustryId returns a boolean if a field has been set.

### SetBusinessIndustryIdNil

`func (o *ExtendedTenantDto) SetBusinessIndustryIdNil(b bool)`

 SetBusinessIndustryIdNil sets the value for BusinessIndustryId to be an explicit nil

### UnsetBusinessIndustryId
`func (o *ExtendedTenantDto) UnsetBusinessIndustryId()`

UnsetBusinessIndustryId ensures that no value is present for BusinessIndustryId, not even an explicit nil
### GetBusinessSegmentId

`func (o *ExtendedTenantDto) GetBusinessSegmentId() string`

GetBusinessSegmentId returns the BusinessSegmentId field if non-nil, zero value otherwise.

### GetBusinessSegmentIdOk

`func (o *ExtendedTenantDto) GetBusinessSegmentIdOk() (*string, bool)`

GetBusinessSegmentIdOk returns a tuple with the BusinessSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessSegmentId

`func (o *ExtendedTenantDto) SetBusinessSegmentId(v string)`

SetBusinessSegmentId sets BusinessSegmentId field to given value.

### HasBusinessSegmentId

`func (o *ExtendedTenantDto) HasBusinessSegmentId() bool`

HasBusinessSegmentId returns a boolean if a field has been set.

### SetBusinessSegmentIdNil

`func (o *ExtendedTenantDto) SetBusinessSegmentIdNil(b bool)`

 SetBusinessSegmentIdNil sets the value for BusinessSegmentId to be an explicit nil

### UnsetBusinessSegmentId
`func (o *ExtendedTenantDto) UnsetBusinessSegmentId()`

UnsetBusinessSegmentId ensures that no value is present for BusinessSegmentId, not even an explicit nil
### GetSocialProfileId

`func (o *ExtendedTenantDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ExtendedTenantDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ExtendedTenantDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ExtendedTenantDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ExtendedTenantDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ExtendedTenantDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetLanguageId

`func (o *ExtendedTenantDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *ExtendedTenantDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *ExtendedTenantDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *ExtendedTenantDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *ExtendedTenantDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *ExtendedTenantDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetName

`func (o *ExtendedTenantDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedTenantDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedTenantDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtendedTenantDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExtendedTenantDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExtendedTenantDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDuns

`func (o *ExtendedTenantDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *ExtendedTenantDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *ExtendedTenantDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *ExtendedTenantDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *ExtendedTenantDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *ExtendedTenantDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetSlogan

`func (o *ExtendedTenantDto) GetSlogan() string`

GetSlogan returns the Slogan field if non-nil, zero value otherwise.

### GetSloganOk

`func (o *ExtendedTenantDto) GetSloganOk() (*string, bool)`

GetSloganOk returns a tuple with the Slogan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlogan

`func (o *ExtendedTenantDto) SetSlogan(v string)`

SetSlogan sets Slogan field to given value.

### HasSlogan

`func (o *ExtendedTenantDto) HasSlogan() bool`

HasSlogan returns a boolean if a field has been set.

### SetSloganNil

`func (o *ExtendedTenantDto) SetSloganNil(b bool)`

 SetSloganNil sets the value for Slogan to be an explicit nil

### UnsetSlogan
`func (o *ExtendedTenantDto) UnsetSlogan()`

UnsetSlogan ensures that no value is present for Slogan, not even an explicit nil
### GetLegalName

`func (o *ExtendedTenantDto) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *ExtendedTenantDto) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *ExtendedTenantDto) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *ExtendedTenantDto) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *ExtendedTenantDto) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *ExtendedTenantDto) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetCoverUrl

`func (o *ExtendedTenantDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *ExtendedTenantDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *ExtendedTenantDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *ExtendedTenantDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *ExtendedTenantDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *ExtendedTenantDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetAvatarUrl

`func (o *ExtendedTenantDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ExtendedTenantDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ExtendedTenantDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ExtendedTenantDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *ExtendedTenantDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *ExtendedTenantDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetCartId

`func (o *ExtendedTenantDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *ExtendedTenantDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *ExtendedTenantDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *ExtendedTenantDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *ExtendedTenantDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *ExtendedTenantDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetCurrencyId

`func (o *ExtendedTenantDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ExtendedTenantDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ExtendedTenantDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ExtendedTenantDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ExtendedTenantDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ExtendedTenantDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTimezoneId

`func (o *ExtendedTenantDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ExtendedTenantDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ExtendedTenantDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ExtendedTenantDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ExtendedTenantDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ExtendedTenantDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetCountryId

`func (o *ExtendedTenantDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ExtendedTenantDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ExtendedTenantDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ExtendedTenantDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ExtendedTenantDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ExtendedTenantDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *ExtendedTenantDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *ExtendedTenantDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *ExtendedTenantDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *ExtendedTenantDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *ExtendedTenantDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *ExtendedTenantDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *ExtendedTenantDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ExtendedTenantDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ExtendedTenantDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ExtendedTenantDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ExtendedTenantDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ExtendedTenantDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetEmail

`func (o *ExtendedTenantDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ExtendedTenantDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ExtendedTenantDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ExtendedTenantDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ExtendedTenantDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ExtendedTenantDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *ExtendedTenantDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ExtendedTenantDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ExtendedTenantDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ExtendedTenantDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *ExtendedTenantDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *ExtendedTenantDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetWebUrl

`func (o *ExtendedTenantDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ExtendedTenantDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ExtendedTenantDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *ExtendedTenantDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *ExtendedTenantDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *ExtendedTenantDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetFacebookUrl

`func (o *ExtendedTenantDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *ExtendedTenantDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *ExtendedTenantDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *ExtendedTenantDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *ExtendedTenantDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *ExtendedTenantDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetTwitterUrl

`func (o *ExtendedTenantDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *ExtendedTenantDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *ExtendedTenantDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *ExtendedTenantDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *ExtendedTenantDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *ExtendedTenantDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetGitHubUrl

`func (o *ExtendedTenantDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *ExtendedTenantDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *ExtendedTenantDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *ExtendedTenantDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *ExtendedTenantDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *ExtendedTenantDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *ExtendedTenantDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *ExtendedTenantDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *ExtendedTenantDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *ExtendedTenantDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *ExtendedTenantDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *ExtendedTenantDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *ExtendedTenantDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *ExtendedTenantDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *ExtendedTenantDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *ExtendedTenantDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *ExtendedTenantDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *ExtendedTenantDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *ExtendedTenantDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *ExtendedTenantDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *ExtendedTenantDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *ExtendedTenantDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *ExtendedTenantDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *ExtendedTenantDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetWhatsAppNumber

`func (o *ExtendedTenantDto) GetWhatsAppNumber() string`

GetWhatsAppNumber returns the WhatsAppNumber field if non-nil, zero value otherwise.

### GetWhatsAppNumberOk

`func (o *ExtendedTenantDto) GetWhatsAppNumberOk() (*string, bool)`

GetWhatsAppNumberOk returns a tuple with the WhatsAppNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsAppNumber

`func (o *ExtendedTenantDto) SetWhatsAppNumber(v string)`

SetWhatsAppNumber sets WhatsAppNumber field to given value.

### HasWhatsAppNumber

`func (o *ExtendedTenantDto) HasWhatsAppNumber() bool`

HasWhatsAppNumber returns a boolean if a field has been set.

### SetWhatsAppNumberNil

`func (o *ExtendedTenantDto) SetWhatsAppNumberNil(b bool)`

 SetWhatsAppNumberNil sets the value for WhatsAppNumber to be an explicit nil

### UnsetWhatsAppNumber
`func (o *ExtendedTenantDto) UnsetWhatsAppNumber()`

UnsetWhatsAppNumber ensures that no value is present for WhatsAppNumber, not even an explicit nil
### GetSupportPhoneNumber

`func (o *ExtendedTenantDto) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *ExtendedTenantDto) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *ExtendedTenantDto) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.

### HasSupportPhoneNumber

`func (o *ExtendedTenantDto) HasSupportPhoneNumber() bool`

HasSupportPhoneNumber returns a boolean if a field has been set.

### SetSupportPhoneNumberNil

`func (o *ExtendedTenantDto) SetSupportPhoneNumberNil(b bool)`

 SetSupportPhoneNumberNil sets the value for SupportPhoneNumber to be an explicit nil

### UnsetSupportPhoneNumber
`func (o *ExtendedTenantDto) UnsetSupportPhoneNumber()`

UnsetSupportPhoneNumber ensures that no value is present for SupportPhoneNumber, not even an explicit nil
### GetVerified

`func (o *ExtendedTenantDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ExtendedTenantDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ExtendedTenantDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ExtendedTenantDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetBusinessName

`func (o *ExtendedTenantDto) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *ExtendedTenantDto) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *ExtendedTenantDto) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *ExtendedTenantDto) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### SetBusinessNameNil

`func (o *ExtendedTenantDto) SetBusinessNameNil(b bool)`

 SetBusinessNameNil sets the value for BusinessName to be an explicit nil

### UnsetBusinessName
`func (o *ExtendedTenantDto) UnsetBusinessName()`

UnsetBusinessName ensures that no value is present for BusinessName, not even an explicit nil
### GetBusinessLegalName

`func (o *ExtendedTenantDto) GetBusinessLegalName() string`

GetBusinessLegalName returns the BusinessLegalName field if non-nil, zero value otherwise.

### GetBusinessLegalNameOk

`func (o *ExtendedTenantDto) GetBusinessLegalNameOk() (*string, bool)`

GetBusinessLegalNameOk returns a tuple with the BusinessLegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLegalName

`func (o *ExtendedTenantDto) SetBusinessLegalName(v string)`

SetBusinessLegalName sets BusinessLegalName field to given value.

### HasBusinessLegalName

`func (o *ExtendedTenantDto) HasBusinessLegalName() bool`

HasBusinessLegalName returns a boolean if a field has been set.

### SetBusinessLegalNameNil

`func (o *ExtendedTenantDto) SetBusinessLegalNameNil(b bool)`

 SetBusinessLegalNameNil sets the value for BusinessLegalName to be an explicit nil

### UnsetBusinessLegalName
`func (o *ExtendedTenantDto) UnsetBusinessLegalName()`

UnsetBusinessLegalName ensures that no value is present for BusinessLegalName, not even an explicit nil
### GetTwitterUsername

`func (o *ExtendedTenantDto) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *ExtendedTenantDto) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *ExtendedTenantDto) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *ExtendedTenantDto) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *ExtendedTenantDto) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *ExtendedTenantDto) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetCart

`func (o *ExtendedTenantDto) GetCart() CartDto`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *ExtendedTenantDto) GetCartOk() (*CartDto, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *ExtendedTenantDto) SetCart(v CartDto)`

SetCart sets Cart field to given value.

### HasCart

`func (o *ExtendedTenantDto) HasCart() bool`

HasCart returns a boolean if a field has been set.

### GetWallet

`func (o *ExtendedTenantDto) GetWallet() WalletDto`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *ExtendedTenantDto) GetWalletOk() (*WalletDto, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *ExtendedTenantDto) SetWallet(v WalletDto)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *ExtendedTenantDto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetSocialProfile

`func (o *ExtendedTenantDto) GetSocialProfile() SocialProfileDto`

GetSocialProfile returns the SocialProfile field if non-nil, zero value otherwise.

### GetSocialProfileOk

`func (o *ExtendedTenantDto) GetSocialProfileOk() (*SocialProfileDto, bool)`

GetSocialProfileOk returns a tuple with the SocialProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfile

`func (o *ExtendedTenantDto) SetSocialProfile(v SocialProfileDto)`

SetSocialProfile sets SocialProfile field to given value.

### HasSocialProfile

`func (o *ExtendedTenantDto) HasSocialProfile() bool`

HasSocialProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


