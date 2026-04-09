# MerchantDto

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
**ProductsCount** | Pointer to **int32** |  | [optional] 
**MerchantRating** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewMerchantDto

`func NewMerchantDto() *MerchantDto`

NewMerchantDto instantiates a new MerchantDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantDtoWithDefaults

`func NewMerchantDtoWithDefaults() *MerchantDto`

NewMerchantDtoWithDefaults instantiates a new MerchantDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MerchantDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MerchantDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MerchantDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MerchantDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MerchantDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MerchantDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *MerchantDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MerchantDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MerchantDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MerchantDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *MerchantDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *MerchantDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQualifiedName

`func (o *MerchantDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *MerchantDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *MerchantDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *MerchantDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *MerchantDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *MerchantDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetTaxId

`func (o *MerchantDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *MerchantDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *MerchantDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *MerchantDto) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *MerchantDto) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *MerchantDto) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetAbout

`func (o *MerchantDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *MerchantDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *MerchantDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *MerchantDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *MerchantDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *MerchantDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetWalletId

`func (o *MerchantDto) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MerchantDto) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MerchantDto) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *MerchantDto) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *MerchantDto) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *MerchantDto) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetSocialFeedId

`func (o *MerchantDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *MerchantDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *MerchantDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *MerchantDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *MerchantDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *MerchantDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil
### GetBusinessIndustryId

`func (o *MerchantDto) GetBusinessIndustryId() string`

GetBusinessIndustryId returns the BusinessIndustryId field if non-nil, zero value otherwise.

### GetBusinessIndustryIdOk

`func (o *MerchantDto) GetBusinessIndustryIdOk() (*string, bool)`

GetBusinessIndustryIdOk returns a tuple with the BusinessIndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessIndustryId

`func (o *MerchantDto) SetBusinessIndustryId(v string)`

SetBusinessIndustryId sets BusinessIndustryId field to given value.

### HasBusinessIndustryId

`func (o *MerchantDto) HasBusinessIndustryId() bool`

HasBusinessIndustryId returns a boolean if a field has been set.

### SetBusinessIndustryIdNil

`func (o *MerchantDto) SetBusinessIndustryIdNil(b bool)`

 SetBusinessIndustryIdNil sets the value for BusinessIndustryId to be an explicit nil

### UnsetBusinessIndustryId
`func (o *MerchantDto) UnsetBusinessIndustryId()`

UnsetBusinessIndustryId ensures that no value is present for BusinessIndustryId, not even an explicit nil
### GetBusinessSegmentId

`func (o *MerchantDto) GetBusinessSegmentId() string`

GetBusinessSegmentId returns the BusinessSegmentId field if non-nil, zero value otherwise.

### GetBusinessSegmentIdOk

`func (o *MerchantDto) GetBusinessSegmentIdOk() (*string, bool)`

GetBusinessSegmentIdOk returns a tuple with the BusinessSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessSegmentId

`func (o *MerchantDto) SetBusinessSegmentId(v string)`

SetBusinessSegmentId sets BusinessSegmentId field to given value.

### HasBusinessSegmentId

`func (o *MerchantDto) HasBusinessSegmentId() bool`

HasBusinessSegmentId returns a boolean if a field has been set.

### SetBusinessSegmentIdNil

`func (o *MerchantDto) SetBusinessSegmentIdNil(b bool)`

 SetBusinessSegmentIdNil sets the value for BusinessSegmentId to be an explicit nil

### UnsetBusinessSegmentId
`func (o *MerchantDto) UnsetBusinessSegmentId()`

UnsetBusinessSegmentId ensures that no value is present for BusinessSegmentId, not even an explicit nil
### GetSocialProfileId

`func (o *MerchantDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *MerchantDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *MerchantDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *MerchantDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *MerchantDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *MerchantDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetLanguageId

`func (o *MerchantDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *MerchantDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *MerchantDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *MerchantDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *MerchantDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *MerchantDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetName

`func (o *MerchantDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MerchantDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MerchantDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MerchantDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MerchantDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MerchantDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDuns

`func (o *MerchantDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *MerchantDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *MerchantDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *MerchantDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *MerchantDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *MerchantDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetSlogan

`func (o *MerchantDto) GetSlogan() string`

GetSlogan returns the Slogan field if non-nil, zero value otherwise.

### GetSloganOk

`func (o *MerchantDto) GetSloganOk() (*string, bool)`

GetSloganOk returns a tuple with the Slogan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlogan

`func (o *MerchantDto) SetSlogan(v string)`

SetSlogan sets Slogan field to given value.

### HasSlogan

`func (o *MerchantDto) HasSlogan() bool`

HasSlogan returns a boolean if a field has been set.

### SetSloganNil

`func (o *MerchantDto) SetSloganNil(b bool)`

 SetSloganNil sets the value for Slogan to be an explicit nil

### UnsetSlogan
`func (o *MerchantDto) UnsetSlogan()`

UnsetSlogan ensures that no value is present for Slogan, not even an explicit nil
### GetLegalName

`func (o *MerchantDto) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *MerchantDto) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *MerchantDto) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *MerchantDto) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *MerchantDto) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *MerchantDto) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetCoverUrl

`func (o *MerchantDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *MerchantDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *MerchantDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *MerchantDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *MerchantDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *MerchantDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetAvatarUrl

`func (o *MerchantDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *MerchantDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *MerchantDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *MerchantDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *MerchantDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *MerchantDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetCartId

`func (o *MerchantDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *MerchantDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *MerchantDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *MerchantDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *MerchantDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *MerchantDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetCurrencyId

`func (o *MerchantDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *MerchantDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *MerchantDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *MerchantDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *MerchantDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *MerchantDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTimezoneId

`func (o *MerchantDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *MerchantDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *MerchantDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *MerchantDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *MerchantDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *MerchantDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetCountryId

`func (o *MerchantDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *MerchantDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *MerchantDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *MerchantDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *MerchantDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *MerchantDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *MerchantDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *MerchantDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *MerchantDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *MerchantDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *MerchantDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *MerchantDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *MerchantDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *MerchantDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *MerchantDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *MerchantDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *MerchantDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *MerchantDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetEmail

`func (o *MerchantDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MerchantDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MerchantDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MerchantDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MerchantDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MerchantDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *MerchantDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MerchantDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MerchantDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MerchantDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MerchantDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MerchantDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetWebUrl

`func (o *MerchantDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MerchantDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MerchantDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MerchantDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MerchantDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MerchantDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetFacebookUrl

`func (o *MerchantDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *MerchantDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *MerchantDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *MerchantDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *MerchantDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *MerchantDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetTwitterUrl

`func (o *MerchantDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *MerchantDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *MerchantDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *MerchantDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *MerchantDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *MerchantDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetGitHubUrl

`func (o *MerchantDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *MerchantDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *MerchantDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *MerchantDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *MerchantDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *MerchantDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *MerchantDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *MerchantDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *MerchantDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *MerchantDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *MerchantDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *MerchantDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *MerchantDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *MerchantDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *MerchantDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *MerchantDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *MerchantDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *MerchantDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *MerchantDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *MerchantDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *MerchantDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *MerchantDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *MerchantDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *MerchantDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetWhatsAppNumber

`func (o *MerchantDto) GetWhatsAppNumber() string`

GetWhatsAppNumber returns the WhatsAppNumber field if non-nil, zero value otherwise.

### GetWhatsAppNumberOk

`func (o *MerchantDto) GetWhatsAppNumberOk() (*string, bool)`

GetWhatsAppNumberOk returns a tuple with the WhatsAppNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsAppNumber

`func (o *MerchantDto) SetWhatsAppNumber(v string)`

SetWhatsAppNumber sets WhatsAppNumber field to given value.

### HasWhatsAppNumber

`func (o *MerchantDto) HasWhatsAppNumber() bool`

HasWhatsAppNumber returns a boolean if a field has been set.

### SetWhatsAppNumberNil

`func (o *MerchantDto) SetWhatsAppNumberNil(b bool)`

 SetWhatsAppNumberNil sets the value for WhatsAppNumber to be an explicit nil

### UnsetWhatsAppNumber
`func (o *MerchantDto) UnsetWhatsAppNumber()`

UnsetWhatsAppNumber ensures that no value is present for WhatsAppNumber, not even an explicit nil
### GetSupportPhoneNumber

`func (o *MerchantDto) GetSupportPhoneNumber() string`

GetSupportPhoneNumber returns the SupportPhoneNumber field if non-nil, zero value otherwise.

### GetSupportPhoneNumberOk

`func (o *MerchantDto) GetSupportPhoneNumberOk() (*string, bool)`

GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhoneNumber

`func (o *MerchantDto) SetSupportPhoneNumber(v string)`

SetSupportPhoneNumber sets SupportPhoneNumber field to given value.

### HasSupportPhoneNumber

`func (o *MerchantDto) HasSupportPhoneNumber() bool`

HasSupportPhoneNumber returns a boolean if a field has been set.

### SetSupportPhoneNumberNil

`func (o *MerchantDto) SetSupportPhoneNumberNil(b bool)`

 SetSupportPhoneNumberNil sets the value for SupportPhoneNumber to be an explicit nil

### UnsetSupportPhoneNumber
`func (o *MerchantDto) UnsetSupportPhoneNumber()`

UnsetSupportPhoneNumber ensures that no value is present for SupportPhoneNumber, not even an explicit nil
### GetVerified

`func (o *MerchantDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *MerchantDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *MerchantDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *MerchantDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetBusinessName

`func (o *MerchantDto) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *MerchantDto) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *MerchantDto) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *MerchantDto) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### SetBusinessNameNil

`func (o *MerchantDto) SetBusinessNameNil(b bool)`

 SetBusinessNameNil sets the value for BusinessName to be an explicit nil

### UnsetBusinessName
`func (o *MerchantDto) UnsetBusinessName()`

UnsetBusinessName ensures that no value is present for BusinessName, not even an explicit nil
### GetBusinessLegalName

`func (o *MerchantDto) GetBusinessLegalName() string`

GetBusinessLegalName returns the BusinessLegalName field if non-nil, zero value otherwise.

### GetBusinessLegalNameOk

`func (o *MerchantDto) GetBusinessLegalNameOk() (*string, bool)`

GetBusinessLegalNameOk returns a tuple with the BusinessLegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessLegalName

`func (o *MerchantDto) SetBusinessLegalName(v string)`

SetBusinessLegalName sets BusinessLegalName field to given value.

### HasBusinessLegalName

`func (o *MerchantDto) HasBusinessLegalName() bool`

HasBusinessLegalName returns a boolean if a field has been set.

### SetBusinessLegalNameNil

`func (o *MerchantDto) SetBusinessLegalNameNil(b bool)`

 SetBusinessLegalNameNil sets the value for BusinessLegalName to be an explicit nil

### UnsetBusinessLegalName
`func (o *MerchantDto) UnsetBusinessLegalName()`

UnsetBusinessLegalName ensures that no value is present for BusinessLegalName, not even an explicit nil
### GetTwitterUsername

`func (o *MerchantDto) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *MerchantDto) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *MerchantDto) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *MerchantDto) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *MerchantDto) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *MerchantDto) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetProductsCount

`func (o *MerchantDto) GetProductsCount() int32`

GetProductsCount returns the ProductsCount field if non-nil, zero value otherwise.

### GetProductsCountOk

`func (o *MerchantDto) GetProductsCountOk() (*int32, bool)`

GetProductsCountOk returns a tuple with the ProductsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsCount

`func (o *MerchantDto) SetProductsCount(v int32)`

SetProductsCount sets ProductsCount field to given value.

### HasProductsCount

`func (o *MerchantDto) HasProductsCount() bool`

HasProductsCount returns a boolean if a field has been set.

### GetMerchantRating

`func (o *MerchantDto) GetMerchantRating() float64`

GetMerchantRating returns the MerchantRating field if non-nil, zero value otherwise.

### GetMerchantRatingOk

`func (o *MerchantDto) GetMerchantRatingOk() (*float64, bool)`

GetMerchantRatingOk returns a tuple with the MerchantRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRating

`func (o *MerchantDto) SetMerchantRating(v float64)`

SetMerchantRating sets MerchantRating field to given value.

### HasMerchantRating

`func (o *MerchantDto) HasMerchantRating() bool`

HasMerchantRating returns a boolean if a field has been set.

### SetMerchantRatingNil

`func (o *MerchantDto) SetMerchantRatingNil(b bool)`

 SetMerchantRatingNil sets the value for MerchantRating to be an explicit nil

### UnsetMerchantRating
`func (o *MerchantDto) UnsetMerchantRating()`

UnsetMerchantRating ensures that no value is present for MerchantRating, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


