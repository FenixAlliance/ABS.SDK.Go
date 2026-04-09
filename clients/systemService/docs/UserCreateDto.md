# UserCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**Birthday** | Pointer to **time.Time** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**PublicName** | Pointer to **NullableString** |  | [optional] 
**IdProvider** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**GitHubUrl** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**TwitterUrl** | Pointer to **NullableString** |  | [optional] 
**FacebookUrl** | Pointer to **NullableString** |  | [optional] 
**YouTubeUrl** | Pointer to **NullableString** |  | [optional] 
**LinkedInUrl** | Pointer to **NullableString** |  | [optional] 
**InstagramUrl** | Pointer to **NullableString** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**LanguageId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserCreateDto

`func NewUserCreateDto() *UserCreateDto`

NewUserCreateDto instantiates a new UserCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateDtoWithDefaults

`func NewUserCreateDtoWithDefaults() *UserCreateDto`

NewUserCreateDtoWithDefaults instantiates a new UserCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *UserCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UserCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetQualifiedName

`func (o *UserCreateDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *UserCreateDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *UserCreateDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *UserCreateDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *UserCreateDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *UserCreateDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetBirthday

`func (o *UserCreateDto) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UserCreateDto) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UserCreateDto) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UserCreateDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetFirstName

`func (o *UserCreateDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCreateDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCreateDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCreateDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UserCreateDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserCreateDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UserCreateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCreateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCreateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCreateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserCreateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserCreateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPublicName

`func (o *UserCreateDto) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *UserCreateDto) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *UserCreateDto) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *UserCreateDto) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *UserCreateDto) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *UserCreateDto) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetIdProvider

`func (o *UserCreateDto) GetIdProvider() string`

GetIdProvider returns the IdProvider field if non-nil, zero value otherwise.

### GetIdProviderOk

`func (o *UserCreateDto) GetIdProviderOk() (*string, bool)`

GetIdProviderOk returns a tuple with the IdProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProvider

`func (o *UserCreateDto) SetIdProvider(v string)`

SetIdProvider sets IdProvider field to given value.

### HasIdProvider

`func (o *UserCreateDto) HasIdProvider() bool`

HasIdProvider returns a boolean if a field has been set.

### SetIdProviderNil

`func (o *UserCreateDto) SetIdProviderNil(b bool)`

 SetIdProviderNil sets the value for IdProvider to be an explicit nil

### UnsetIdProvider
`func (o *UserCreateDto) UnsetIdProvider()`

UnsetIdProvider ensures that no value is present for IdProvider, not even an explicit nil
### GetGender

`func (o *UserCreateDto) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UserCreateDto) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UserCreateDto) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UserCreateDto) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetEmail

`func (o *UserCreateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreateDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserCreateDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserCreateDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAbout

`func (o *UserCreateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *UserCreateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *UserCreateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *UserCreateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *UserCreateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *UserCreateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetStatus

`func (o *UserCreateDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserCreateDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserCreateDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserCreateDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UserCreateDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UserCreateDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetJobTitle

`func (o *UserCreateDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UserCreateDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UserCreateDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UserCreateDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *UserCreateDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *UserCreateDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetGitHubUrl

`func (o *UserCreateDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *UserCreateDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *UserCreateDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *UserCreateDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *UserCreateDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *UserCreateDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetWebsiteUrl

`func (o *UserCreateDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *UserCreateDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *UserCreateDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *UserCreateDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *UserCreateDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *UserCreateDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetTwitterUrl

`func (o *UserCreateDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *UserCreateDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *UserCreateDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *UserCreateDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *UserCreateDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *UserCreateDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetFacebookUrl

`func (o *UserCreateDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *UserCreateDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *UserCreateDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *UserCreateDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *UserCreateDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *UserCreateDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *UserCreateDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *UserCreateDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *UserCreateDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *UserCreateDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *UserCreateDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *UserCreateDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *UserCreateDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *UserCreateDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *UserCreateDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *UserCreateDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *UserCreateDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *UserCreateDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *UserCreateDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *UserCreateDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *UserCreateDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *UserCreateDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *UserCreateDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *UserCreateDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetTimezoneId

`func (o *UserCreateDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *UserCreateDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *UserCreateDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *UserCreateDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *UserCreateDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *UserCreateDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetLanguageId

`func (o *UserCreateDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *UserCreateDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *UserCreateDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *UserCreateDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *UserCreateDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *UserCreateDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetCurrencyId

`func (o *UserCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *UserCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *UserCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *UserCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *UserCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *UserCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetCountryId

`func (o *UserCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *UserCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *UserCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *UserCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *UserCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *UserCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *UserCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *UserCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *UserCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *UserCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *UserCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *UserCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *UserCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *UserCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *UserCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *UserCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *UserCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *UserCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetPassword

`func (o *UserCreateDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreateDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreateDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreateDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UserCreateDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UserCreateDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


