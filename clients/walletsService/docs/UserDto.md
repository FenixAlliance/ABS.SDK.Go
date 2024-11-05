# UserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] [readonly] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**PublicName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**CoverUrl** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**GitHubUrl** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**TwitterUrl** | Pointer to **NullableString** |  | [optional] 
**YouTubeUrl** | Pointer to **NullableString** |  | [optional] 
**LinkedInUrl** | Pointer to **NullableString** |  | [optional] 
**FacebookUrl** | Pointer to **NullableString** |  | [optional] 
**InstagramUrl** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**Birthday** | Pointer to **NullableTime** |  | [optional] 
**IdProvider** | Pointer to **NullableString** |  | [optional] 
**LanguageId** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableInt32** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**SocialFeedId** | Pointer to **NullableString** |  | [optional] 
**CurrentTenantId** | Pointer to **NullableString** |  | [optional] 
**CurrentEnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**WalletId** | Pointer to **NullableString** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 
**PublicIdentifier** | Pointer to **NullableString** |  | [optional] 
**IdentityProvider** | Pointer to **NullableString** |  | [optional] 
**PhoneNumberConfirmed** | Pointer to **bool** |  | [optional] 
**EmailConfirmed** | Pointer to **bool** |  | [optional] 
**Availability** | Pointer to **NullableInt32** |  | [optional] 
**LockoutEnabled** | Pointer to **bool** |  | [optional] 
**LockoutEnd** | Pointer to **NullableTime** |  | [optional] 
**EnrollmentsCount** | Pointer to **NullableInt32** |  | [optional] 
**SiteTheme** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUserDto

`func NewUserDto() *UserDto`

NewUserDto instantiates a new UserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDtoWithDefaults

`func NewUserDtoWithDefaults() *UserDto`

NewUserDtoWithDefaults instantiates a new UserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *UserDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UserDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *UserDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *UserDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetFullName

`func (o *UserDto) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UserDto) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UserDto) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UserDto) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *UserDto) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *UserDto) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetQualifiedName

`func (o *UserDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *UserDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *UserDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *UserDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *UserDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *UserDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetPublicName

`func (o *UserDto) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *UserDto) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *UserDto) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *UserDto) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *UserDto) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *UserDto) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetLastName

`func (o *UserDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetFirstName

`func (o *UserDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UserDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetCoverUrl

`func (o *UserDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *UserDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *UserDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *UserDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *UserDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *UserDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetAvatarUrl

`func (o *UserDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *UserDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *UserDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetGitHubUrl

`func (o *UserDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *UserDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *UserDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *UserDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *UserDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *UserDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetCountryId

`func (o *UserDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *UserDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *UserDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *UserDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *UserDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *UserDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTimezoneId

`func (o *UserDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *UserDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *UserDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *UserDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *UserDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *UserDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetWebsiteUrl

`func (o *UserDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *UserDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *UserDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *UserDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *UserDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *UserDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetTwitterUrl

`func (o *UserDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *UserDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *UserDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *UserDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *UserDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *UserDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *UserDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *UserDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *UserDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *UserDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *UserDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *UserDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *UserDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *UserDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *UserDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *UserDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *UserDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *UserDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetFacebookUrl

`func (o *UserDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *UserDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *UserDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *UserDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *UserDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *UserDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetInstagramUrl

`func (o *UserDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *UserDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *UserDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *UserDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *UserDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *UserDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetSocialProfileId

`func (o *UserDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *UserDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *UserDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *UserDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *UserDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *UserDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetBirthday

`func (o *UserDto) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UserDto) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UserDto) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UserDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *UserDto) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *UserDto) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetIdProvider

`func (o *UserDto) GetIdProvider() string`

GetIdProvider returns the IdProvider field if non-nil, zero value otherwise.

### GetIdProviderOk

`func (o *UserDto) GetIdProviderOk() (*string, bool)`

GetIdProviderOk returns a tuple with the IdProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProvider

`func (o *UserDto) SetIdProvider(v string)`

SetIdProvider sets IdProvider field to given value.

### HasIdProvider

`func (o *UserDto) HasIdProvider() bool`

HasIdProvider returns a boolean if a field has been set.

### SetIdProviderNil

`func (o *UserDto) SetIdProviderNil(b bool)`

 SetIdProviderNil sets the value for IdProvider to be an explicit nil

### UnsetIdProvider
`func (o *UserDto) UnsetIdProvider()`

UnsetIdProvider ensures that no value is present for IdProvider, not even an explicit nil
### GetLanguageId

`func (o *UserDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *UserDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *UserDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *UserDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *UserDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *UserDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetGender

`func (o *UserDto) GetGender() int32`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UserDto) GetGenderOk() (*int32, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UserDto) SetGender(v int32)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UserDto) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *UserDto) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *UserDto) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetCityId

`func (o *UserDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *UserDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *UserDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *UserDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *UserDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *UserDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetStateId

`func (o *UserDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *UserDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *UserDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *UserDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *UserDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *UserDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetEmail

`func (o *UserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAbout

`func (o *UserDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *UserDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *UserDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *UserDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *UserDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *UserDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetJobTitle

`func (o *UserDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UserDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UserDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UserDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *UserDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *UserDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetSocialFeedId

`func (o *UserDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *UserDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *UserDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *UserDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *UserDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *UserDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil
### GetCurrentTenantId

`func (o *UserDto) GetCurrentTenantId() string`

GetCurrentTenantId returns the CurrentTenantId field if non-nil, zero value otherwise.

### GetCurrentTenantIdOk

`func (o *UserDto) GetCurrentTenantIdOk() (*string, bool)`

GetCurrentTenantIdOk returns a tuple with the CurrentTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTenantId

`func (o *UserDto) SetCurrentTenantId(v string)`

SetCurrentTenantId sets CurrentTenantId field to given value.

### HasCurrentTenantId

`func (o *UserDto) HasCurrentTenantId() bool`

HasCurrentTenantId returns a boolean if a field has been set.

### SetCurrentTenantIdNil

`func (o *UserDto) SetCurrentTenantIdNil(b bool)`

 SetCurrentTenantIdNil sets the value for CurrentTenantId to be an explicit nil

### UnsetCurrentTenantId
`func (o *UserDto) UnsetCurrentTenantId()`

UnsetCurrentTenantId ensures that no value is present for CurrentTenantId, not even an explicit nil
### GetCurrentEnrollmentId

`func (o *UserDto) GetCurrentEnrollmentId() string`

GetCurrentEnrollmentId returns the CurrentEnrollmentId field if non-nil, zero value otherwise.

### GetCurrentEnrollmentIdOk

`func (o *UserDto) GetCurrentEnrollmentIdOk() (*string, bool)`

GetCurrentEnrollmentIdOk returns a tuple with the CurrentEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEnrollmentId

`func (o *UserDto) SetCurrentEnrollmentId(v string)`

SetCurrentEnrollmentId sets CurrentEnrollmentId field to given value.

### HasCurrentEnrollmentId

`func (o *UserDto) HasCurrentEnrollmentId() bool`

HasCurrentEnrollmentId returns a boolean if a field has been set.

### SetCurrentEnrollmentIdNil

`func (o *UserDto) SetCurrentEnrollmentIdNil(b bool)`

 SetCurrentEnrollmentIdNil sets the value for CurrentEnrollmentId to be an explicit nil

### UnsetCurrentEnrollmentId
`func (o *UserDto) UnsetCurrentEnrollmentId()`

UnsetCurrentEnrollmentId ensures that no value is present for CurrentEnrollmentId, not even an explicit nil
### GetStatus

`func (o *UserDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UserDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UserDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCartId

`func (o *UserDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *UserDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *UserDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *UserDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *UserDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *UserDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetWalletId

`func (o *UserDto) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *UserDto) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *UserDto) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *UserDto) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *UserDto) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *UserDto) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetUserName

`func (o *UserDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UserDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetCurrencyId

`func (o *UserDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *UserDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *UserDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *UserDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *UserDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *UserDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetPhoneNumber

`func (o *UserDto) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserDto) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserDto) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserDto) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *UserDto) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *UserDto) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPublicIdentifier

`func (o *UserDto) GetPublicIdentifier() string`

GetPublicIdentifier returns the PublicIdentifier field if non-nil, zero value otherwise.

### GetPublicIdentifierOk

`func (o *UserDto) GetPublicIdentifierOk() (*string, bool)`

GetPublicIdentifierOk returns a tuple with the PublicIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdentifier

`func (o *UserDto) SetPublicIdentifier(v string)`

SetPublicIdentifier sets PublicIdentifier field to given value.

### HasPublicIdentifier

`func (o *UserDto) HasPublicIdentifier() bool`

HasPublicIdentifier returns a boolean if a field has been set.

### SetPublicIdentifierNil

`func (o *UserDto) SetPublicIdentifierNil(b bool)`

 SetPublicIdentifierNil sets the value for PublicIdentifier to be an explicit nil

### UnsetPublicIdentifier
`func (o *UserDto) UnsetPublicIdentifier()`

UnsetPublicIdentifier ensures that no value is present for PublicIdentifier, not even an explicit nil
### GetIdentityProvider

`func (o *UserDto) GetIdentityProvider() string`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *UserDto) GetIdentityProviderOk() (*string, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *UserDto) SetIdentityProvider(v string)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *UserDto) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### SetIdentityProviderNil

`func (o *UserDto) SetIdentityProviderNil(b bool)`

 SetIdentityProviderNil sets the value for IdentityProvider to be an explicit nil

### UnsetIdentityProvider
`func (o *UserDto) UnsetIdentityProvider()`

UnsetIdentityProvider ensures that no value is present for IdentityProvider, not even an explicit nil
### GetPhoneNumberConfirmed

`func (o *UserDto) GetPhoneNumberConfirmed() bool`

GetPhoneNumberConfirmed returns the PhoneNumberConfirmed field if non-nil, zero value otherwise.

### GetPhoneNumberConfirmedOk

`func (o *UserDto) GetPhoneNumberConfirmedOk() (*bool, bool)`

GetPhoneNumberConfirmedOk returns a tuple with the PhoneNumberConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberConfirmed

`func (o *UserDto) SetPhoneNumberConfirmed(v bool)`

SetPhoneNumberConfirmed sets PhoneNumberConfirmed field to given value.

### HasPhoneNumberConfirmed

`func (o *UserDto) HasPhoneNumberConfirmed() bool`

HasPhoneNumberConfirmed returns a boolean if a field has been set.

### GetEmailConfirmed

`func (o *UserDto) GetEmailConfirmed() bool`

GetEmailConfirmed returns the EmailConfirmed field if non-nil, zero value otherwise.

### GetEmailConfirmedOk

`func (o *UserDto) GetEmailConfirmedOk() (*bool, bool)`

GetEmailConfirmedOk returns a tuple with the EmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmed

`func (o *UserDto) SetEmailConfirmed(v bool)`

SetEmailConfirmed sets EmailConfirmed field to given value.

### HasEmailConfirmed

`func (o *UserDto) HasEmailConfirmed() bool`

HasEmailConfirmed returns a boolean if a field has been set.

### GetAvailability

`func (o *UserDto) GetAvailability() int32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *UserDto) GetAvailabilityOk() (*int32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *UserDto) SetAvailability(v int32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *UserDto) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *UserDto) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *UserDto) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetLockoutEnabled

`func (o *UserDto) GetLockoutEnabled() bool`

GetLockoutEnabled returns the LockoutEnabled field if non-nil, zero value otherwise.

### GetLockoutEnabledOk

`func (o *UserDto) GetLockoutEnabledOk() (*bool, bool)`

GetLockoutEnabledOk returns a tuple with the LockoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnabled

`func (o *UserDto) SetLockoutEnabled(v bool)`

SetLockoutEnabled sets LockoutEnabled field to given value.

### HasLockoutEnabled

`func (o *UserDto) HasLockoutEnabled() bool`

HasLockoutEnabled returns a boolean if a field has been set.

### GetLockoutEnd

`func (o *UserDto) GetLockoutEnd() time.Time`

GetLockoutEnd returns the LockoutEnd field if non-nil, zero value otherwise.

### GetLockoutEndOk

`func (o *UserDto) GetLockoutEndOk() (*time.Time, bool)`

GetLockoutEndOk returns a tuple with the LockoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnd

`func (o *UserDto) SetLockoutEnd(v time.Time)`

SetLockoutEnd sets LockoutEnd field to given value.

### HasLockoutEnd

`func (o *UserDto) HasLockoutEnd() bool`

HasLockoutEnd returns a boolean if a field has been set.

### SetLockoutEndNil

`func (o *UserDto) SetLockoutEndNil(b bool)`

 SetLockoutEndNil sets the value for LockoutEnd to be an explicit nil

### UnsetLockoutEnd
`func (o *UserDto) UnsetLockoutEnd()`

UnsetLockoutEnd ensures that no value is present for LockoutEnd, not even an explicit nil
### GetEnrollmentsCount

`func (o *UserDto) GetEnrollmentsCount() int32`

GetEnrollmentsCount returns the EnrollmentsCount field if non-nil, zero value otherwise.

### GetEnrollmentsCountOk

`func (o *UserDto) GetEnrollmentsCountOk() (*int32, bool)`

GetEnrollmentsCountOk returns a tuple with the EnrollmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentsCount

`func (o *UserDto) SetEnrollmentsCount(v int32)`

SetEnrollmentsCount sets EnrollmentsCount field to given value.

### HasEnrollmentsCount

`func (o *UserDto) HasEnrollmentsCount() bool`

HasEnrollmentsCount returns a boolean if a field has been set.

### SetEnrollmentsCountNil

`func (o *UserDto) SetEnrollmentsCountNil(b bool)`

 SetEnrollmentsCountNil sets the value for EnrollmentsCount to be an explicit nil

### UnsetEnrollmentsCount
`func (o *UserDto) UnsetEnrollmentsCount()`

UnsetEnrollmentsCount ensures that no value is present for EnrollmentsCount, not even an explicit nil
### GetSiteTheme

`func (o *UserDto) GetSiteTheme() int32`

GetSiteTheme returns the SiteTheme field if non-nil, zero value otherwise.

### GetSiteThemeOk

`func (o *UserDto) GetSiteThemeOk() (*int32, bool)`

GetSiteThemeOk returns a tuple with the SiteTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTheme

`func (o *UserDto) SetSiteTheme(v int32)`

SetSiteTheme sets SiteTheme field to given value.

### HasSiteTheme

`func (o *UserDto) HasSiteTheme() bool`

HasSiteTheme returns a boolean if a field has been set.

### SetSiteThemeNil

`func (o *UserDto) SetSiteThemeNil(b bool)`

 SetSiteThemeNil sets the value for SiteTheme to be an explicit nil

### UnsetSiteTheme
`func (o *UserDto) UnsetSiteTheme()`

UnsetSiteTheme ensures that no value is present for SiteTheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


