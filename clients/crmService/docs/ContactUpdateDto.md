# ContactUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Birthday** | Pointer to **NullableTime** |  | [optional] 
**Duns** | Pointer to **NullableString** |  | [optional] 
**TaxId** | Pointer to **NullableString** |  | [optional] 
**Email** | **string** |  | 
**FirstName** | **string** |  | 
**LastName** | Pointer to **NullableString** |  | [optional] 
**PrimaryContactId** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] [readonly] 
**About** | Pointer to **NullableString** |  | [optional] 
**MobilePhone** | Pointer to **NullableString** |  | [optional] 
**BusinessPhone** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**ParentContactId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**StreetLine1** | Pointer to **NullableString** |  | [optional] 
**StreetLine2** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**LanguageId** | Pointer to **NullableString** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**CoverUrl** | Pointer to **NullableString** |  | [optional] 
**GithubUsername** | Pointer to **NullableString** |  | [optional] 
**InstagramUsername** | Pointer to **NullableString** |  | [optional] 
**TwitchUrl** | Pointer to **NullableString** |  | [optional] 
**RedditUrl** | Pointer to **NullableString** |  | [optional] 
**GitHubUrl** | Pointer to **NullableString** |  | [optional] 
**GithubUrl** | Pointer to **NullableString** |  | [optional] 
**TikTokUrl** | Pointer to **NullableString** |  | [optional] 
**TwitterUrl** | Pointer to **NullableString** |  | [optional] 
**YouTubeUrl** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**FacebookUrl** | Pointer to **NullableString** |  | [optional] 
**LinkedInUrl** | Pointer to **NullableString** |  | [optional] 
**InstagramUrl** | Pointer to **NullableString** |  | [optional] 
**TikTokUsername** | Pointer to **NullableString** |  | [optional] 
**StackExchangeUrl** | Pointer to **NullableString** |  | [optional] 
**StackOverflowUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactUpdateDto

`func NewContactUpdateDto(type_ int32, email string, firstName string, ) *ContactUpdateDto`

NewContactUpdateDto instantiates a new ContactUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactUpdateDtoWithDefaults

`func NewContactUpdateDtoWithDefaults() *ContactUpdateDto`

NewContactUpdateDtoWithDefaults instantiates a new ContactUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContactUpdateDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactUpdateDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactUpdateDto) SetType(v int32)`

SetType sets Type field to given value.


### GetBirthday

`func (o *ContactUpdateDto) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *ContactUpdateDto) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *ContactUpdateDto) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *ContactUpdateDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *ContactUpdateDto) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *ContactUpdateDto) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetDuns

`func (o *ContactUpdateDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *ContactUpdateDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *ContactUpdateDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *ContactUpdateDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *ContactUpdateDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *ContactUpdateDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetTaxId

`func (o *ContactUpdateDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *ContactUpdateDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *ContactUpdateDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *ContactUpdateDto) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *ContactUpdateDto) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *ContactUpdateDto) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetEmail

`func (o *ContactUpdateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactUpdateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactUpdateDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *ContactUpdateDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactUpdateDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactUpdateDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ContactUpdateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactUpdateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactUpdateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ContactUpdateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ContactUpdateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ContactUpdateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPrimaryContactId

`func (o *ContactUpdateDto) GetPrimaryContactId() string`

GetPrimaryContactId returns the PrimaryContactId field if non-nil, zero value otherwise.

### GetPrimaryContactIdOk

`func (o *ContactUpdateDto) GetPrimaryContactIdOk() (*string, bool)`

GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactId

`func (o *ContactUpdateDto) SetPrimaryContactId(v string)`

SetPrimaryContactId sets PrimaryContactId field to given value.

### HasPrimaryContactId

`func (o *ContactUpdateDto) HasPrimaryContactId() bool`

HasPrimaryContactId returns a boolean if a field has been set.

### SetPrimaryContactIdNil

`func (o *ContactUpdateDto) SetPrimaryContactIdNil(b bool)`

 SetPrimaryContactIdNil sets the value for PrimaryContactId to be an explicit nil

### UnsetPrimaryContactId
`func (o *ContactUpdateDto) UnsetPrimaryContactId()`

UnsetPrimaryContactId ensures that no value is present for PrimaryContactId, not even an explicit nil
### GetQualifiedName

`func (o *ContactUpdateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *ContactUpdateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *ContactUpdateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *ContactUpdateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *ContactUpdateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *ContactUpdateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetAbout

`func (o *ContactUpdateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ContactUpdateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ContactUpdateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *ContactUpdateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *ContactUpdateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *ContactUpdateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetMobilePhone

`func (o *ContactUpdateDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ContactUpdateDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ContactUpdateDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ContactUpdateDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *ContactUpdateDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *ContactUpdateDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetBusinessPhone

`func (o *ContactUpdateDto) GetBusinessPhone() string`

GetBusinessPhone returns the BusinessPhone field if non-nil, zero value otherwise.

### GetBusinessPhoneOk

`func (o *ContactUpdateDto) GetBusinessPhoneOk() (*string, bool)`

GetBusinessPhoneOk returns a tuple with the BusinessPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhone

`func (o *ContactUpdateDto) SetBusinessPhone(v string)`

SetBusinessPhone sets BusinessPhone field to given value.

### HasBusinessPhone

`func (o *ContactUpdateDto) HasBusinessPhone() bool`

HasBusinessPhone returns a boolean if a field has been set.

### SetBusinessPhoneNil

`func (o *ContactUpdateDto) SetBusinessPhoneNil(b bool)`

 SetBusinessPhoneNil sets the value for BusinessPhone to be an explicit nil

### UnsetBusinessPhone
`func (o *ContactUpdateDto) UnsetBusinessPhone()`

UnsetBusinessPhone ensures that no value is present for BusinessPhone, not even an explicit nil
### GetJobTitle

`func (o *ContactUpdateDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *ContactUpdateDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *ContactUpdateDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *ContactUpdateDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *ContactUpdateDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *ContactUpdateDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetCountryId

`func (o *ContactUpdateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ContactUpdateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ContactUpdateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ContactUpdateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ContactUpdateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ContactUpdateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetParentContactId

`func (o *ContactUpdateDto) GetParentContactId() string`

GetParentContactId returns the ParentContactId field if non-nil, zero value otherwise.

### GetParentContactIdOk

`func (o *ContactUpdateDto) GetParentContactIdOk() (*string, bool)`

GetParentContactIdOk returns a tuple with the ParentContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentContactId

`func (o *ContactUpdateDto) SetParentContactId(v string)`

SetParentContactId sets ParentContactId field to given value.

### HasParentContactId

`func (o *ContactUpdateDto) HasParentContactId() bool`

HasParentContactId returns a boolean if a field has been set.

### SetParentContactIdNil

`func (o *ContactUpdateDto) SetParentContactIdNil(b bool)`

 SetParentContactIdNil sets the value for ParentContactId to be an explicit nil

### UnsetParentContactId
`func (o *ContactUpdateDto) UnsetParentContactId()`

UnsetParentContactId ensures that no value is present for ParentContactId, not even an explicit nil
### GetStateId

`func (o *ContactUpdateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *ContactUpdateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *ContactUpdateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *ContactUpdateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *ContactUpdateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *ContactUpdateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *ContactUpdateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ContactUpdateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ContactUpdateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ContactUpdateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ContactUpdateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ContactUpdateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetPostalCode

`func (o *ContactUpdateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ContactUpdateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ContactUpdateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ContactUpdateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ContactUpdateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ContactUpdateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetStreetLine1

`func (o *ContactUpdateDto) GetStreetLine1() string`

GetStreetLine1 returns the StreetLine1 field if non-nil, zero value otherwise.

### GetStreetLine1Ok

`func (o *ContactUpdateDto) GetStreetLine1Ok() (*string, bool)`

GetStreetLine1Ok returns a tuple with the StreetLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine1

`func (o *ContactUpdateDto) SetStreetLine1(v string)`

SetStreetLine1 sets StreetLine1 field to given value.

### HasStreetLine1

`func (o *ContactUpdateDto) HasStreetLine1() bool`

HasStreetLine1 returns a boolean if a field has been set.

### SetStreetLine1Nil

`func (o *ContactUpdateDto) SetStreetLine1Nil(b bool)`

 SetStreetLine1Nil sets the value for StreetLine1 to be an explicit nil

### UnsetStreetLine1
`func (o *ContactUpdateDto) UnsetStreetLine1()`

UnsetStreetLine1 ensures that no value is present for StreetLine1, not even an explicit nil
### GetStreetLine2

`func (o *ContactUpdateDto) GetStreetLine2() string`

GetStreetLine2 returns the StreetLine2 field if non-nil, zero value otherwise.

### GetStreetLine2Ok

`func (o *ContactUpdateDto) GetStreetLine2Ok() (*string, bool)`

GetStreetLine2Ok returns a tuple with the StreetLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine2

`func (o *ContactUpdateDto) SetStreetLine2(v string)`

SetStreetLine2 sets StreetLine2 field to given value.

### HasStreetLine2

`func (o *ContactUpdateDto) HasStreetLine2() bool`

HasStreetLine2 returns a boolean if a field has been set.

### SetStreetLine2Nil

`func (o *ContactUpdateDto) SetStreetLine2Nil(b bool)`

 SetStreetLine2Nil sets the value for StreetLine2 to be an explicit nil

### UnsetStreetLine2
`func (o *ContactUpdateDto) UnsetStreetLine2()`

UnsetStreetLine2 ensures that no value is present for StreetLine2, not even an explicit nil
### GetCurrencyId

`func (o *ContactUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ContactUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ContactUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ContactUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ContactUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ContactUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetLanguageId

`func (o *ContactUpdateDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *ContactUpdateDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *ContactUpdateDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *ContactUpdateDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *ContactUpdateDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *ContactUpdateDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetTimezoneId

`func (o *ContactUpdateDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ContactUpdateDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ContactUpdateDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ContactUpdateDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ContactUpdateDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ContactUpdateDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetCoverUrl

`func (o *ContactUpdateDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *ContactUpdateDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *ContactUpdateDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *ContactUpdateDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *ContactUpdateDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *ContactUpdateDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetGithubUsername

`func (o *ContactUpdateDto) GetGithubUsername() string`

GetGithubUsername returns the GithubUsername field if non-nil, zero value otherwise.

### GetGithubUsernameOk

`func (o *ContactUpdateDto) GetGithubUsernameOk() (*string, bool)`

GetGithubUsernameOk returns a tuple with the GithubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUsername

`func (o *ContactUpdateDto) SetGithubUsername(v string)`

SetGithubUsername sets GithubUsername field to given value.

### HasGithubUsername

`func (o *ContactUpdateDto) HasGithubUsername() bool`

HasGithubUsername returns a boolean if a field has been set.

### SetGithubUsernameNil

`func (o *ContactUpdateDto) SetGithubUsernameNil(b bool)`

 SetGithubUsernameNil sets the value for GithubUsername to be an explicit nil

### UnsetGithubUsername
`func (o *ContactUpdateDto) UnsetGithubUsername()`

UnsetGithubUsername ensures that no value is present for GithubUsername, not even an explicit nil
### GetInstagramUsername

`func (o *ContactUpdateDto) GetInstagramUsername() string`

GetInstagramUsername returns the InstagramUsername field if non-nil, zero value otherwise.

### GetInstagramUsernameOk

`func (o *ContactUpdateDto) GetInstagramUsernameOk() (*string, bool)`

GetInstagramUsernameOk returns a tuple with the InstagramUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUsername

`func (o *ContactUpdateDto) SetInstagramUsername(v string)`

SetInstagramUsername sets InstagramUsername field to given value.

### HasInstagramUsername

`func (o *ContactUpdateDto) HasInstagramUsername() bool`

HasInstagramUsername returns a boolean if a field has been set.

### SetInstagramUsernameNil

`func (o *ContactUpdateDto) SetInstagramUsernameNil(b bool)`

 SetInstagramUsernameNil sets the value for InstagramUsername to be an explicit nil

### UnsetInstagramUsername
`func (o *ContactUpdateDto) UnsetInstagramUsername()`

UnsetInstagramUsername ensures that no value is present for InstagramUsername, not even an explicit nil
### GetTwitchUrl

`func (o *ContactUpdateDto) GetTwitchUrl() string`

GetTwitchUrl returns the TwitchUrl field if non-nil, zero value otherwise.

### GetTwitchUrlOk

`func (o *ContactUpdateDto) GetTwitchUrlOk() (*string, bool)`

GetTwitchUrlOk returns a tuple with the TwitchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchUrl

`func (o *ContactUpdateDto) SetTwitchUrl(v string)`

SetTwitchUrl sets TwitchUrl field to given value.

### HasTwitchUrl

`func (o *ContactUpdateDto) HasTwitchUrl() bool`

HasTwitchUrl returns a boolean if a field has been set.

### SetTwitchUrlNil

`func (o *ContactUpdateDto) SetTwitchUrlNil(b bool)`

 SetTwitchUrlNil sets the value for TwitchUrl to be an explicit nil

### UnsetTwitchUrl
`func (o *ContactUpdateDto) UnsetTwitchUrl()`

UnsetTwitchUrl ensures that no value is present for TwitchUrl, not even an explicit nil
### GetRedditUrl

`func (o *ContactUpdateDto) GetRedditUrl() string`

GetRedditUrl returns the RedditUrl field if non-nil, zero value otherwise.

### GetRedditUrlOk

`func (o *ContactUpdateDto) GetRedditUrlOk() (*string, bool)`

GetRedditUrlOk returns a tuple with the RedditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedditUrl

`func (o *ContactUpdateDto) SetRedditUrl(v string)`

SetRedditUrl sets RedditUrl field to given value.

### HasRedditUrl

`func (o *ContactUpdateDto) HasRedditUrl() bool`

HasRedditUrl returns a boolean if a field has been set.

### SetRedditUrlNil

`func (o *ContactUpdateDto) SetRedditUrlNil(b bool)`

 SetRedditUrlNil sets the value for RedditUrl to be an explicit nil

### UnsetRedditUrl
`func (o *ContactUpdateDto) UnsetRedditUrl()`

UnsetRedditUrl ensures that no value is present for RedditUrl, not even an explicit nil
### GetGitHubUrl

`func (o *ContactUpdateDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *ContactUpdateDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *ContactUpdateDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *ContactUpdateDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *ContactUpdateDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *ContactUpdateDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetGithubUrl

`func (o *ContactUpdateDto) GetGithubUrl() string`

GetGithubUrl returns the GithubUrl field if non-nil, zero value otherwise.

### GetGithubUrlOk

`func (o *ContactUpdateDto) GetGithubUrlOk() (*string, bool)`

GetGithubUrlOk returns a tuple with the GithubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUrl

`func (o *ContactUpdateDto) SetGithubUrl(v string)`

SetGithubUrl sets GithubUrl field to given value.

### HasGithubUrl

`func (o *ContactUpdateDto) HasGithubUrl() bool`

HasGithubUrl returns a boolean if a field has been set.

### SetGithubUrlNil

`func (o *ContactUpdateDto) SetGithubUrlNil(b bool)`

 SetGithubUrlNil sets the value for GithubUrl to be an explicit nil

### UnsetGithubUrl
`func (o *ContactUpdateDto) UnsetGithubUrl()`

UnsetGithubUrl ensures that no value is present for GithubUrl, not even an explicit nil
### GetTikTokUrl

`func (o *ContactUpdateDto) GetTikTokUrl() string`

GetTikTokUrl returns the TikTokUrl field if non-nil, zero value otherwise.

### GetTikTokUrlOk

`func (o *ContactUpdateDto) GetTikTokUrlOk() (*string, bool)`

GetTikTokUrlOk returns a tuple with the TikTokUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTokUrl

`func (o *ContactUpdateDto) SetTikTokUrl(v string)`

SetTikTokUrl sets TikTokUrl field to given value.

### HasTikTokUrl

`func (o *ContactUpdateDto) HasTikTokUrl() bool`

HasTikTokUrl returns a boolean if a field has been set.

### SetTikTokUrlNil

`func (o *ContactUpdateDto) SetTikTokUrlNil(b bool)`

 SetTikTokUrlNil sets the value for TikTokUrl to be an explicit nil

### UnsetTikTokUrl
`func (o *ContactUpdateDto) UnsetTikTokUrl()`

UnsetTikTokUrl ensures that no value is present for TikTokUrl, not even an explicit nil
### GetTwitterUrl

`func (o *ContactUpdateDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *ContactUpdateDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *ContactUpdateDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *ContactUpdateDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *ContactUpdateDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *ContactUpdateDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *ContactUpdateDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *ContactUpdateDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *ContactUpdateDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *ContactUpdateDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *ContactUpdateDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *ContactUpdateDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetWebsiteUrl

`func (o *ContactUpdateDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ContactUpdateDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ContactUpdateDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ContactUpdateDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *ContactUpdateDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *ContactUpdateDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetFacebookUrl

`func (o *ContactUpdateDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *ContactUpdateDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *ContactUpdateDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *ContactUpdateDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *ContactUpdateDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *ContactUpdateDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *ContactUpdateDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *ContactUpdateDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *ContactUpdateDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *ContactUpdateDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *ContactUpdateDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *ContactUpdateDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *ContactUpdateDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *ContactUpdateDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *ContactUpdateDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *ContactUpdateDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *ContactUpdateDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *ContactUpdateDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetTikTokUsername

`func (o *ContactUpdateDto) GetTikTokUsername() string`

GetTikTokUsername returns the TikTokUsername field if non-nil, zero value otherwise.

### GetTikTokUsernameOk

`func (o *ContactUpdateDto) GetTikTokUsernameOk() (*string, bool)`

GetTikTokUsernameOk returns a tuple with the TikTokUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTokUsername

`func (o *ContactUpdateDto) SetTikTokUsername(v string)`

SetTikTokUsername sets TikTokUsername field to given value.

### HasTikTokUsername

`func (o *ContactUpdateDto) HasTikTokUsername() bool`

HasTikTokUsername returns a boolean if a field has been set.

### SetTikTokUsernameNil

`func (o *ContactUpdateDto) SetTikTokUsernameNil(b bool)`

 SetTikTokUsernameNil sets the value for TikTokUsername to be an explicit nil

### UnsetTikTokUsername
`func (o *ContactUpdateDto) UnsetTikTokUsername()`

UnsetTikTokUsername ensures that no value is present for TikTokUsername, not even an explicit nil
### GetStackExchangeUrl

`func (o *ContactUpdateDto) GetStackExchangeUrl() string`

GetStackExchangeUrl returns the StackExchangeUrl field if non-nil, zero value otherwise.

### GetStackExchangeUrlOk

`func (o *ContactUpdateDto) GetStackExchangeUrlOk() (*string, bool)`

GetStackExchangeUrlOk returns a tuple with the StackExchangeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackExchangeUrl

`func (o *ContactUpdateDto) SetStackExchangeUrl(v string)`

SetStackExchangeUrl sets StackExchangeUrl field to given value.

### HasStackExchangeUrl

`func (o *ContactUpdateDto) HasStackExchangeUrl() bool`

HasStackExchangeUrl returns a boolean if a field has been set.

### SetStackExchangeUrlNil

`func (o *ContactUpdateDto) SetStackExchangeUrlNil(b bool)`

 SetStackExchangeUrlNil sets the value for StackExchangeUrl to be an explicit nil

### UnsetStackExchangeUrl
`func (o *ContactUpdateDto) UnsetStackExchangeUrl()`

UnsetStackExchangeUrl ensures that no value is present for StackExchangeUrl, not even an explicit nil
### GetStackOverflowUrl

`func (o *ContactUpdateDto) GetStackOverflowUrl() string`

GetStackOverflowUrl returns the StackOverflowUrl field if non-nil, zero value otherwise.

### GetStackOverflowUrlOk

`func (o *ContactUpdateDto) GetStackOverflowUrlOk() (*string, bool)`

GetStackOverflowUrlOk returns a tuple with the StackOverflowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackOverflowUrl

`func (o *ContactUpdateDto) SetStackOverflowUrl(v string)`

SetStackOverflowUrl sets StackOverflowUrl field to given value.

### HasStackOverflowUrl

`func (o *ContactUpdateDto) HasStackOverflowUrl() bool`

HasStackOverflowUrl returns a boolean if a field has been set.

### SetStackOverflowUrlNil

`func (o *ContactUpdateDto) SetStackOverflowUrlNil(b bool)`

 SetStackOverflowUrlNil sets the value for StackOverflowUrl to be an explicit nil

### UnsetStackOverflowUrl
`func (o *ContactUpdateDto) UnsetStackOverflowUrl()`

UnsetStackOverflowUrl ensures that no value is present for StackOverflowUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


