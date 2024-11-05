# ContactCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**TenantId** | **string** |  | 
**Type** | **int32** |  | 
**FirstName** | **string** |  | 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Email** | **string** |  | 
**TaxId** | Pointer to **NullableString** |  | [optional] 
**PrimaryContactId** | Pointer to **NullableString** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] [readonly] 
**About** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**MobilePhone** | Pointer to **NullableString** |  | [optional] 
**BusinessPhone** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**Duns** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**WebUrl** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**LanguageId** | Pointer to **NullableString** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**Birthday** | Pointer to **NullableTime** |  | [optional] 
**StreetLine1** | Pointer to **NullableString** |  | [optional] 
**StreetLine2** | Pointer to **NullableString** |  | [optional] 
**GitHubUrl** | Pointer to **NullableString** |  | [optional] 
**TwitchUrl** | Pointer to **NullableString** |  | [optional] 
**RedditUrl** | Pointer to **NullableString** |  | [optional] 
**TikTokUrl** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**TwitterUrl** | Pointer to **NullableString** |  | [optional] 
**FacebookUrl** | Pointer to **NullableString** |  | [optional] 
**YouTubeUrl** | Pointer to **NullableString** |  | [optional] 
**LinkedInUrl** | Pointer to **NullableString** |  | [optional] 
**InstagramUrl** | Pointer to **NullableString** |  | [optional] 
**GithubUsername** | Pointer to **NullableString** |  | [optional] 
**InstagramUsername** | Pointer to **interface{}** |  | [optional] 
**TikTokUsername** | Pointer to **interface{}** |  | [optional] 
**StackExchangeUrl** | Pointer to **interface{}** |  | [optional] 
**StackOverflowUrl** | Pointer to **interface{}** |  | [optional] 
**ParentContactId** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewContactCreateDto

`func NewContactCreateDto(tenantId string, type_ int32, firstName string, email string, ) *ContactCreateDto`

NewContactCreateDto instantiates a new ContactCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCreateDtoWithDefaults

`func NewContactCreateDtoWithDefaults() *ContactCreateDto`

NewContactCreateDtoWithDefaults instantiates a new ContactCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ContactCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContactCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContactCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContactCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *ContactCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContactCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContactCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetType

`func (o *ContactCreateDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactCreateDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactCreateDto) SetType(v int32)`

SetType sets Type field to given value.


### GetFirstName

`func (o *ContactCreateDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactCreateDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactCreateDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ContactCreateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactCreateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactCreateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ContactCreateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ContactCreateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ContactCreateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmail

`func (o *ContactCreateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactCreateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactCreateDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTaxId

`func (o *ContactCreateDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *ContactCreateDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *ContactCreateDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *ContactCreateDto) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *ContactCreateDto) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *ContactCreateDto) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetPrimaryContactId

`func (o *ContactCreateDto) GetPrimaryContactId() string`

GetPrimaryContactId returns the PrimaryContactId field if non-nil, zero value otherwise.

### GetPrimaryContactIdOk

`func (o *ContactCreateDto) GetPrimaryContactIdOk() (*string, bool)`

GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactId

`func (o *ContactCreateDto) SetPrimaryContactId(v string)`

SetPrimaryContactId sets PrimaryContactId field to given value.

### HasPrimaryContactId

`func (o *ContactCreateDto) HasPrimaryContactId() bool`

HasPrimaryContactId returns a boolean if a field has been set.

### SetPrimaryContactIdNil

`func (o *ContactCreateDto) SetPrimaryContactIdNil(b bool)`

 SetPrimaryContactIdNil sets the value for PrimaryContactId to be an explicit nil

### UnsetPrimaryContactId
`func (o *ContactCreateDto) UnsetPrimaryContactId()`

UnsetPrimaryContactId ensures that no value is present for PrimaryContactId, not even an explicit nil
### GetQualifiedName

`func (o *ContactCreateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *ContactCreateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *ContactCreateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *ContactCreateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *ContactCreateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *ContactCreateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetAbout

`func (o *ContactCreateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ContactCreateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ContactCreateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *ContactCreateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *ContactCreateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *ContactCreateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetCountryId

`func (o *ContactCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ContactCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ContactCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ContactCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ContactCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ContactCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *ContactCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *ContactCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *ContactCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *ContactCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *ContactCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *ContactCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *ContactCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ContactCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ContactCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ContactCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ContactCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ContactCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetMobilePhone

`func (o *ContactCreateDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ContactCreateDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ContactCreateDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ContactCreateDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *ContactCreateDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *ContactCreateDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetBusinessPhone

`func (o *ContactCreateDto) GetBusinessPhone() string`

GetBusinessPhone returns the BusinessPhone field if non-nil, zero value otherwise.

### GetBusinessPhoneOk

`func (o *ContactCreateDto) GetBusinessPhoneOk() (*string, bool)`

GetBusinessPhoneOk returns a tuple with the BusinessPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhone

`func (o *ContactCreateDto) SetBusinessPhone(v string)`

SetBusinessPhone sets BusinessPhone field to given value.

### HasBusinessPhone

`func (o *ContactCreateDto) HasBusinessPhone() bool`

HasBusinessPhone returns a boolean if a field has been set.

### SetBusinessPhoneNil

`func (o *ContactCreateDto) SetBusinessPhoneNil(b bool)`

 SetBusinessPhoneNil sets the value for BusinessPhone to be an explicit nil

### UnsetBusinessPhone
`func (o *ContactCreateDto) UnsetBusinessPhone()`

UnsetBusinessPhone ensures that no value is present for BusinessPhone, not even an explicit nil
### GetPostalCode

`func (o *ContactCreateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ContactCreateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ContactCreateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ContactCreateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ContactCreateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ContactCreateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetDuns

`func (o *ContactCreateDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *ContactCreateDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *ContactCreateDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *ContactCreateDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *ContactCreateDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *ContactCreateDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetJobTitle

`func (o *ContactCreateDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *ContactCreateDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *ContactCreateDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *ContactCreateDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *ContactCreateDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *ContactCreateDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetWebUrl

`func (o *ContactCreateDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ContactCreateDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ContactCreateDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *ContactCreateDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *ContactCreateDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *ContactCreateDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetCurrencyId

`func (o *ContactCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ContactCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ContactCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ContactCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ContactCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ContactCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetLanguageId

`func (o *ContactCreateDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *ContactCreateDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *ContactCreateDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *ContactCreateDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *ContactCreateDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *ContactCreateDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetTimezoneId

`func (o *ContactCreateDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ContactCreateDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ContactCreateDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ContactCreateDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ContactCreateDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ContactCreateDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetBirthday

`func (o *ContactCreateDto) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *ContactCreateDto) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *ContactCreateDto) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *ContactCreateDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *ContactCreateDto) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *ContactCreateDto) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetStreetLine1

`func (o *ContactCreateDto) GetStreetLine1() string`

GetStreetLine1 returns the StreetLine1 field if non-nil, zero value otherwise.

### GetStreetLine1Ok

`func (o *ContactCreateDto) GetStreetLine1Ok() (*string, bool)`

GetStreetLine1Ok returns a tuple with the StreetLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine1

`func (o *ContactCreateDto) SetStreetLine1(v string)`

SetStreetLine1 sets StreetLine1 field to given value.

### HasStreetLine1

`func (o *ContactCreateDto) HasStreetLine1() bool`

HasStreetLine1 returns a boolean if a field has been set.

### SetStreetLine1Nil

`func (o *ContactCreateDto) SetStreetLine1Nil(b bool)`

 SetStreetLine1Nil sets the value for StreetLine1 to be an explicit nil

### UnsetStreetLine1
`func (o *ContactCreateDto) UnsetStreetLine1()`

UnsetStreetLine1 ensures that no value is present for StreetLine1, not even an explicit nil
### GetStreetLine2

`func (o *ContactCreateDto) GetStreetLine2() string`

GetStreetLine2 returns the StreetLine2 field if non-nil, zero value otherwise.

### GetStreetLine2Ok

`func (o *ContactCreateDto) GetStreetLine2Ok() (*string, bool)`

GetStreetLine2Ok returns a tuple with the StreetLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine2

`func (o *ContactCreateDto) SetStreetLine2(v string)`

SetStreetLine2 sets StreetLine2 field to given value.

### HasStreetLine2

`func (o *ContactCreateDto) HasStreetLine2() bool`

HasStreetLine2 returns a boolean if a field has been set.

### SetStreetLine2Nil

`func (o *ContactCreateDto) SetStreetLine2Nil(b bool)`

 SetStreetLine2Nil sets the value for StreetLine2 to be an explicit nil

### UnsetStreetLine2
`func (o *ContactCreateDto) UnsetStreetLine2()`

UnsetStreetLine2 ensures that no value is present for StreetLine2, not even an explicit nil
### GetGitHubUrl

`func (o *ContactCreateDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *ContactCreateDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *ContactCreateDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *ContactCreateDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *ContactCreateDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *ContactCreateDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetTwitchUrl

`func (o *ContactCreateDto) GetTwitchUrl() string`

GetTwitchUrl returns the TwitchUrl field if non-nil, zero value otherwise.

### GetTwitchUrlOk

`func (o *ContactCreateDto) GetTwitchUrlOk() (*string, bool)`

GetTwitchUrlOk returns a tuple with the TwitchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchUrl

`func (o *ContactCreateDto) SetTwitchUrl(v string)`

SetTwitchUrl sets TwitchUrl field to given value.

### HasTwitchUrl

`func (o *ContactCreateDto) HasTwitchUrl() bool`

HasTwitchUrl returns a boolean if a field has been set.

### SetTwitchUrlNil

`func (o *ContactCreateDto) SetTwitchUrlNil(b bool)`

 SetTwitchUrlNil sets the value for TwitchUrl to be an explicit nil

### UnsetTwitchUrl
`func (o *ContactCreateDto) UnsetTwitchUrl()`

UnsetTwitchUrl ensures that no value is present for TwitchUrl, not even an explicit nil
### GetRedditUrl

`func (o *ContactCreateDto) GetRedditUrl() string`

GetRedditUrl returns the RedditUrl field if non-nil, zero value otherwise.

### GetRedditUrlOk

`func (o *ContactCreateDto) GetRedditUrlOk() (*string, bool)`

GetRedditUrlOk returns a tuple with the RedditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedditUrl

`func (o *ContactCreateDto) SetRedditUrl(v string)`

SetRedditUrl sets RedditUrl field to given value.

### HasRedditUrl

`func (o *ContactCreateDto) HasRedditUrl() bool`

HasRedditUrl returns a boolean if a field has been set.

### SetRedditUrlNil

`func (o *ContactCreateDto) SetRedditUrlNil(b bool)`

 SetRedditUrlNil sets the value for RedditUrl to be an explicit nil

### UnsetRedditUrl
`func (o *ContactCreateDto) UnsetRedditUrl()`

UnsetRedditUrl ensures that no value is present for RedditUrl, not even an explicit nil
### GetTikTokUrl

`func (o *ContactCreateDto) GetTikTokUrl() string`

GetTikTokUrl returns the TikTokUrl field if non-nil, zero value otherwise.

### GetTikTokUrlOk

`func (o *ContactCreateDto) GetTikTokUrlOk() (*string, bool)`

GetTikTokUrlOk returns a tuple with the TikTokUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTokUrl

`func (o *ContactCreateDto) SetTikTokUrl(v string)`

SetTikTokUrl sets TikTokUrl field to given value.

### HasTikTokUrl

`func (o *ContactCreateDto) HasTikTokUrl() bool`

HasTikTokUrl returns a boolean if a field has been set.

### SetTikTokUrlNil

`func (o *ContactCreateDto) SetTikTokUrlNil(b bool)`

 SetTikTokUrlNil sets the value for TikTokUrl to be an explicit nil

### UnsetTikTokUrl
`func (o *ContactCreateDto) UnsetTikTokUrl()`

UnsetTikTokUrl ensures that no value is present for TikTokUrl, not even an explicit nil
### GetWebsiteUrl

`func (o *ContactCreateDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ContactCreateDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ContactCreateDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ContactCreateDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *ContactCreateDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *ContactCreateDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetTwitterUrl

`func (o *ContactCreateDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *ContactCreateDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *ContactCreateDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *ContactCreateDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *ContactCreateDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *ContactCreateDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetFacebookUrl

`func (o *ContactCreateDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *ContactCreateDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *ContactCreateDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *ContactCreateDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *ContactCreateDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *ContactCreateDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *ContactCreateDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *ContactCreateDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *ContactCreateDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *ContactCreateDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *ContactCreateDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *ContactCreateDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *ContactCreateDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *ContactCreateDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *ContactCreateDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *ContactCreateDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *ContactCreateDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *ContactCreateDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *ContactCreateDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *ContactCreateDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *ContactCreateDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *ContactCreateDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *ContactCreateDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *ContactCreateDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetGithubUsername

`func (o *ContactCreateDto) GetGithubUsername() string`

GetGithubUsername returns the GithubUsername field if non-nil, zero value otherwise.

### GetGithubUsernameOk

`func (o *ContactCreateDto) GetGithubUsernameOk() (*string, bool)`

GetGithubUsernameOk returns a tuple with the GithubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUsername

`func (o *ContactCreateDto) SetGithubUsername(v string)`

SetGithubUsername sets GithubUsername field to given value.

### HasGithubUsername

`func (o *ContactCreateDto) HasGithubUsername() bool`

HasGithubUsername returns a boolean if a field has been set.

### SetGithubUsernameNil

`func (o *ContactCreateDto) SetGithubUsernameNil(b bool)`

 SetGithubUsernameNil sets the value for GithubUsername to be an explicit nil

### UnsetGithubUsername
`func (o *ContactCreateDto) UnsetGithubUsername()`

UnsetGithubUsername ensures that no value is present for GithubUsername, not even an explicit nil
### GetInstagramUsername

`func (o *ContactCreateDto) GetInstagramUsername() interface{}`

GetInstagramUsername returns the InstagramUsername field if non-nil, zero value otherwise.

### GetInstagramUsernameOk

`func (o *ContactCreateDto) GetInstagramUsernameOk() (*interface{}, bool)`

GetInstagramUsernameOk returns a tuple with the InstagramUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUsername

`func (o *ContactCreateDto) SetInstagramUsername(v interface{})`

SetInstagramUsername sets InstagramUsername field to given value.

### HasInstagramUsername

`func (o *ContactCreateDto) HasInstagramUsername() bool`

HasInstagramUsername returns a boolean if a field has been set.

### SetInstagramUsernameNil

`func (o *ContactCreateDto) SetInstagramUsernameNil(b bool)`

 SetInstagramUsernameNil sets the value for InstagramUsername to be an explicit nil

### UnsetInstagramUsername
`func (o *ContactCreateDto) UnsetInstagramUsername()`

UnsetInstagramUsername ensures that no value is present for InstagramUsername, not even an explicit nil
### GetTikTokUsername

`func (o *ContactCreateDto) GetTikTokUsername() interface{}`

GetTikTokUsername returns the TikTokUsername field if non-nil, zero value otherwise.

### GetTikTokUsernameOk

`func (o *ContactCreateDto) GetTikTokUsernameOk() (*interface{}, bool)`

GetTikTokUsernameOk returns a tuple with the TikTokUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTokUsername

`func (o *ContactCreateDto) SetTikTokUsername(v interface{})`

SetTikTokUsername sets TikTokUsername field to given value.

### HasTikTokUsername

`func (o *ContactCreateDto) HasTikTokUsername() bool`

HasTikTokUsername returns a boolean if a field has been set.

### SetTikTokUsernameNil

`func (o *ContactCreateDto) SetTikTokUsernameNil(b bool)`

 SetTikTokUsernameNil sets the value for TikTokUsername to be an explicit nil

### UnsetTikTokUsername
`func (o *ContactCreateDto) UnsetTikTokUsername()`

UnsetTikTokUsername ensures that no value is present for TikTokUsername, not even an explicit nil
### GetStackExchangeUrl

`func (o *ContactCreateDto) GetStackExchangeUrl() interface{}`

GetStackExchangeUrl returns the StackExchangeUrl field if non-nil, zero value otherwise.

### GetStackExchangeUrlOk

`func (o *ContactCreateDto) GetStackExchangeUrlOk() (*interface{}, bool)`

GetStackExchangeUrlOk returns a tuple with the StackExchangeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackExchangeUrl

`func (o *ContactCreateDto) SetStackExchangeUrl(v interface{})`

SetStackExchangeUrl sets StackExchangeUrl field to given value.

### HasStackExchangeUrl

`func (o *ContactCreateDto) HasStackExchangeUrl() bool`

HasStackExchangeUrl returns a boolean if a field has been set.

### SetStackExchangeUrlNil

`func (o *ContactCreateDto) SetStackExchangeUrlNil(b bool)`

 SetStackExchangeUrlNil sets the value for StackExchangeUrl to be an explicit nil

### UnsetStackExchangeUrl
`func (o *ContactCreateDto) UnsetStackExchangeUrl()`

UnsetStackExchangeUrl ensures that no value is present for StackExchangeUrl, not even an explicit nil
### GetStackOverflowUrl

`func (o *ContactCreateDto) GetStackOverflowUrl() interface{}`

GetStackOverflowUrl returns the StackOverflowUrl field if non-nil, zero value otherwise.

### GetStackOverflowUrlOk

`func (o *ContactCreateDto) GetStackOverflowUrlOk() (*interface{}, bool)`

GetStackOverflowUrlOk returns a tuple with the StackOverflowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackOverflowUrl

`func (o *ContactCreateDto) SetStackOverflowUrl(v interface{})`

SetStackOverflowUrl sets StackOverflowUrl field to given value.

### HasStackOverflowUrl

`func (o *ContactCreateDto) HasStackOverflowUrl() bool`

HasStackOverflowUrl returns a boolean if a field has been set.

### SetStackOverflowUrlNil

`func (o *ContactCreateDto) SetStackOverflowUrlNil(b bool)`

 SetStackOverflowUrlNil sets the value for StackOverflowUrl to be an explicit nil

### UnsetStackOverflowUrl
`func (o *ContactCreateDto) UnsetStackOverflowUrl()`

UnsetStackOverflowUrl ensures that no value is present for StackOverflowUrl, not even an explicit nil
### GetParentContactId

`func (o *ContactCreateDto) GetParentContactId() interface{}`

GetParentContactId returns the ParentContactId field if non-nil, zero value otherwise.

### GetParentContactIdOk

`func (o *ContactCreateDto) GetParentContactIdOk() (*interface{}, bool)`

GetParentContactIdOk returns a tuple with the ParentContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentContactId

`func (o *ContactCreateDto) SetParentContactId(v interface{})`

SetParentContactId sets ParentContactId field to given value.

### HasParentContactId

`func (o *ContactCreateDto) HasParentContactId() bool`

HasParentContactId returns a boolean if a field has been set.

### SetParentContactIdNil

`func (o *ContactCreateDto) SetParentContactIdNil(b bool)`

 SetParentContactIdNil sets the value for ParentContactId to be an explicit nil

### UnsetParentContactId
`func (o *ContactCreateDto) UnsetParentContactId()`

UnsetParentContactId ensures that no value is present for ParentContactId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


