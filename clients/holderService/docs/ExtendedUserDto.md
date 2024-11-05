# ExtendedUserDto

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
**Cart** | Pointer to [**CartDto**](CartDto.md) |  | [optional] 
**Wallet** | Pointer to [**WalletDto**](WalletDto.md) |  | [optional] 
**SocialProfile** | Pointer to [**SocialProfileDto**](SocialProfileDto.md) |  | [optional] 
**Settings** | Pointer to [**UserSettingsDto**](UserSettingsDto.md) |  | [optional] 

## Methods

### NewExtendedUserDto

`func NewExtendedUserDto() *ExtendedUserDto`

NewExtendedUserDto instantiates a new ExtendedUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedUserDtoWithDefaults

`func NewExtendedUserDtoWithDefaults() *ExtendedUserDto`

NewExtendedUserDtoWithDefaults instantiates a new ExtendedUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedUserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedUserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedUserDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedUserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedUserDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedUserDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ExtendedUserDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedUserDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedUserDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtendedUserDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ExtendedUserDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ExtendedUserDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetFullName

`func (o *ExtendedUserDto) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ExtendedUserDto) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ExtendedUserDto) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ExtendedUserDto) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *ExtendedUserDto) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *ExtendedUserDto) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetQualifiedName

`func (o *ExtendedUserDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *ExtendedUserDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *ExtendedUserDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *ExtendedUserDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *ExtendedUserDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *ExtendedUserDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetPublicName

`func (o *ExtendedUserDto) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *ExtendedUserDto) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *ExtendedUserDto) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *ExtendedUserDto) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *ExtendedUserDto) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *ExtendedUserDto) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetLastName

`func (o *ExtendedUserDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ExtendedUserDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ExtendedUserDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ExtendedUserDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ExtendedUserDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ExtendedUserDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetFirstName

`func (o *ExtendedUserDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ExtendedUserDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ExtendedUserDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ExtendedUserDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ExtendedUserDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ExtendedUserDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetCoverUrl

`func (o *ExtendedUserDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *ExtendedUserDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *ExtendedUserDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *ExtendedUserDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *ExtendedUserDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *ExtendedUserDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetAvatarUrl

`func (o *ExtendedUserDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ExtendedUserDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ExtendedUserDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ExtendedUserDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *ExtendedUserDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *ExtendedUserDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetGitHubUrl

`func (o *ExtendedUserDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *ExtendedUserDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *ExtendedUserDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *ExtendedUserDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *ExtendedUserDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *ExtendedUserDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetCountryId

`func (o *ExtendedUserDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ExtendedUserDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ExtendedUserDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ExtendedUserDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ExtendedUserDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ExtendedUserDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTimezoneId

`func (o *ExtendedUserDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ExtendedUserDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ExtendedUserDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ExtendedUserDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ExtendedUserDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ExtendedUserDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetWebsiteUrl

`func (o *ExtendedUserDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ExtendedUserDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ExtendedUserDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ExtendedUserDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *ExtendedUserDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *ExtendedUserDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetTwitterUrl

`func (o *ExtendedUserDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *ExtendedUserDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *ExtendedUserDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *ExtendedUserDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *ExtendedUserDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *ExtendedUserDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *ExtendedUserDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *ExtendedUserDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *ExtendedUserDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *ExtendedUserDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *ExtendedUserDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *ExtendedUserDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *ExtendedUserDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *ExtendedUserDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *ExtendedUserDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *ExtendedUserDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *ExtendedUserDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *ExtendedUserDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetFacebookUrl

`func (o *ExtendedUserDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *ExtendedUserDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *ExtendedUserDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *ExtendedUserDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *ExtendedUserDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *ExtendedUserDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetInstagramUrl

`func (o *ExtendedUserDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *ExtendedUserDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *ExtendedUserDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *ExtendedUserDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *ExtendedUserDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *ExtendedUserDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetSocialProfileId

`func (o *ExtendedUserDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ExtendedUserDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ExtendedUserDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ExtendedUserDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ExtendedUserDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ExtendedUserDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetBirthday

`func (o *ExtendedUserDto) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *ExtendedUserDto) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *ExtendedUserDto) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *ExtendedUserDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *ExtendedUserDto) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *ExtendedUserDto) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetIdProvider

`func (o *ExtendedUserDto) GetIdProvider() string`

GetIdProvider returns the IdProvider field if non-nil, zero value otherwise.

### GetIdProviderOk

`func (o *ExtendedUserDto) GetIdProviderOk() (*string, bool)`

GetIdProviderOk returns a tuple with the IdProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProvider

`func (o *ExtendedUserDto) SetIdProvider(v string)`

SetIdProvider sets IdProvider field to given value.

### HasIdProvider

`func (o *ExtendedUserDto) HasIdProvider() bool`

HasIdProvider returns a boolean if a field has been set.

### SetIdProviderNil

`func (o *ExtendedUserDto) SetIdProviderNil(b bool)`

 SetIdProviderNil sets the value for IdProvider to be an explicit nil

### UnsetIdProvider
`func (o *ExtendedUserDto) UnsetIdProvider()`

UnsetIdProvider ensures that no value is present for IdProvider, not even an explicit nil
### GetLanguageId

`func (o *ExtendedUserDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *ExtendedUserDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *ExtendedUserDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *ExtendedUserDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *ExtendedUserDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *ExtendedUserDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetGender

`func (o *ExtendedUserDto) GetGender() int32`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ExtendedUserDto) GetGenderOk() (*int32, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ExtendedUserDto) SetGender(v int32)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ExtendedUserDto) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ExtendedUserDto) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ExtendedUserDto) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetCityId

`func (o *ExtendedUserDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ExtendedUserDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ExtendedUserDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ExtendedUserDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ExtendedUserDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ExtendedUserDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetStateId

`func (o *ExtendedUserDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *ExtendedUserDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *ExtendedUserDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *ExtendedUserDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *ExtendedUserDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *ExtendedUserDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetEmail

`func (o *ExtendedUserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ExtendedUserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ExtendedUserDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ExtendedUserDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ExtendedUserDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ExtendedUserDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAbout

`func (o *ExtendedUserDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ExtendedUserDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ExtendedUserDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *ExtendedUserDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *ExtendedUserDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *ExtendedUserDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetJobTitle

`func (o *ExtendedUserDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *ExtendedUserDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *ExtendedUserDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *ExtendedUserDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *ExtendedUserDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *ExtendedUserDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetSocialFeedId

`func (o *ExtendedUserDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *ExtendedUserDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *ExtendedUserDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *ExtendedUserDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *ExtendedUserDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *ExtendedUserDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil
### GetCurrentTenantId

`func (o *ExtendedUserDto) GetCurrentTenantId() string`

GetCurrentTenantId returns the CurrentTenantId field if non-nil, zero value otherwise.

### GetCurrentTenantIdOk

`func (o *ExtendedUserDto) GetCurrentTenantIdOk() (*string, bool)`

GetCurrentTenantIdOk returns a tuple with the CurrentTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTenantId

`func (o *ExtendedUserDto) SetCurrentTenantId(v string)`

SetCurrentTenantId sets CurrentTenantId field to given value.

### HasCurrentTenantId

`func (o *ExtendedUserDto) HasCurrentTenantId() bool`

HasCurrentTenantId returns a boolean if a field has been set.

### SetCurrentTenantIdNil

`func (o *ExtendedUserDto) SetCurrentTenantIdNil(b bool)`

 SetCurrentTenantIdNil sets the value for CurrentTenantId to be an explicit nil

### UnsetCurrentTenantId
`func (o *ExtendedUserDto) UnsetCurrentTenantId()`

UnsetCurrentTenantId ensures that no value is present for CurrentTenantId, not even an explicit nil
### GetCurrentEnrollmentId

`func (o *ExtendedUserDto) GetCurrentEnrollmentId() string`

GetCurrentEnrollmentId returns the CurrentEnrollmentId field if non-nil, zero value otherwise.

### GetCurrentEnrollmentIdOk

`func (o *ExtendedUserDto) GetCurrentEnrollmentIdOk() (*string, bool)`

GetCurrentEnrollmentIdOk returns a tuple with the CurrentEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEnrollmentId

`func (o *ExtendedUserDto) SetCurrentEnrollmentId(v string)`

SetCurrentEnrollmentId sets CurrentEnrollmentId field to given value.

### HasCurrentEnrollmentId

`func (o *ExtendedUserDto) HasCurrentEnrollmentId() bool`

HasCurrentEnrollmentId returns a boolean if a field has been set.

### SetCurrentEnrollmentIdNil

`func (o *ExtendedUserDto) SetCurrentEnrollmentIdNil(b bool)`

 SetCurrentEnrollmentIdNil sets the value for CurrentEnrollmentId to be an explicit nil

### UnsetCurrentEnrollmentId
`func (o *ExtendedUserDto) UnsetCurrentEnrollmentId()`

UnsetCurrentEnrollmentId ensures that no value is present for CurrentEnrollmentId, not even an explicit nil
### GetStatus

`func (o *ExtendedUserDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtendedUserDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtendedUserDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExtendedUserDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ExtendedUserDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ExtendedUserDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCartId

`func (o *ExtendedUserDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *ExtendedUserDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *ExtendedUserDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *ExtendedUserDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *ExtendedUserDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *ExtendedUserDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetWalletId

`func (o *ExtendedUserDto) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExtendedUserDto) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExtendedUserDto) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *ExtendedUserDto) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *ExtendedUserDto) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *ExtendedUserDto) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetUserName

`func (o *ExtendedUserDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ExtendedUserDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ExtendedUserDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ExtendedUserDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ExtendedUserDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ExtendedUserDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetCurrencyId

`func (o *ExtendedUserDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ExtendedUserDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ExtendedUserDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ExtendedUserDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ExtendedUserDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ExtendedUserDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetPhoneNumber

`func (o *ExtendedUserDto) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ExtendedUserDto) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ExtendedUserDto) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ExtendedUserDto) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ExtendedUserDto) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ExtendedUserDto) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetPublicIdentifier

`func (o *ExtendedUserDto) GetPublicIdentifier() string`

GetPublicIdentifier returns the PublicIdentifier field if non-nil, zero value otherwise.

### GetPublicIdentifierOk

`func (o *ExtendedUserDto) GetPublicIdentifierOk() (*string, bool)`

GetPublicIdentifierOk returns a tuple with the PublicIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdentifier

`func (o *ExtendedUserDto) SetPublicIdentifier(v string)`

SetPublicIdentifier sets PublicIdentifier field to given value.

### HasPublicIdentifier

`func (o *ExtendedUserDto) HasPublicIdentifier() bool`

HasPublicIdentifier returns a boolean if a field has been set.

### SetPublicIdentifierNil

`func (o *ExtendedUserDto) SetPublicIdentifierNil(b bool)`

 SetPublicIdentifierNil sets the value for PublicIdentifier to be an explicit nil

### UnsetPublicIdentifier
`func (o *ExtendedUserDto) UnsetPublicIdentifier()`

UnsetPublicIdentifier ensures that no value is present for PublicIdentifier, not even an explicit nil
### GetIdentityProvider

`func (o *ExtendedUserDto) GetIdentityProvider() string`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *ExtendedUserDto) GetIdentityProviderOk() (*string, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *ExtendedUserDto) SetIdentityProvider(v string)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *ExtendedUserDto) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### SetIdentityProviderNil

`func (o *ExtendedUserDto) SetIdentityProviderNil(b bool)`

 SetIdentityProviderNil sets the value for IdentityProvider to be an explicit nil

### UnsetIdentityProvider
`func (o *ExtendedUserDto) UnsetIdentityProvider()`

UnsetIdentityProvider ensures that no value is present for IdentityProvider, not even an explicit nil
### GetPhoneNumberConfirmed

`func (o *ExtendedUserDto) GetPhoneNumberConfirmed() bool`

GetPhoneNumberConfirmed returns the PhoneNumberConfirmed field if non-nil, zero value otherwise.

### GetPhoneNumberConfirmedOk

`func (o *ExtendedUserDto) GetPhoneNumberConfirmedOk() (*bool, bool)`

GetPhoneNumberConfirmedOk returns a tuple with the PhoneNumberConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberConfirmed

`func (o *ExtendedUserDto) SetPhoneNumberConfirmed(v bool)`

SetPhoneNumberConfirmed sets PhoneNumberConfirmed field to given value.

### HasPhoneNumberConfirmed

`func (o *ExtendedUserDto) HasPhoneNumberConfirmed() bool`

HasPhoneNumberConfirmed returns a boolean if a field has been set.

### GetEmailConfirmed

`func (o *ExtendedUserDto) GetEmailConfirmed() bool`

GetEmailConfirmed returns the EmailConfirmed field if non-nil, zero value otherwise.

### GetEmailConfirmedOk

`func (o *ExtendedUserDto) GetEmailConfirmedOk() (*bool, bool)`

GetEmailConfirmedOk returns a tuple with the EmailConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConfirmed

`func (o *ExtendedUserDto) SetEmailConfirmed(v bool)`

SetEmailConfirmed sets EmailConfirmed field to given value.

### HasEmailConfirmed

`func (o *ExtendedUserDto) HasEmailConfirmed() bool`

HasEmailConfirmed returns a boolean if a field has been set.

### GetAvailability

`func (o *ExtendedUserDto) GetAvailability() int32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ExtendedUserDto) GetAvailabilityOk() (*int32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ExtendedUserDto) SetAvailability(v int32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ExtendedUserDto) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *ExtendedUserDto) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *ExtendedUserDto) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil
### GetLockoutEnabled

`func (o *ExtendedUserDto) GetLockoutEnabled() bool`

GetLockoutEnabled returns the LockoutEnabled field if non-nil, zero value otherwise.

### GetLockoutEnabledOk

`func (o *ExtendedUserDto) GetLockoutEnabledOk() (*bool, bool)`

GetLockoutEnabledOk returns a tuple with the LockoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnabled

`func (o *ExtendedUserDto) SetLockoutEnabled(v bool)`

SetLockoutEnabled sets LockoutEnabled field to given value.

### HasLockoutEnabled

`func (o *ExtendedUserDto) HasLockoutEnabled() bool`

HasLockoutEnabled returns a boolean if a field has been set.

### GetLockoutEnd

`func (o *ExtendedUserDto) GetLockoutEnd() time.Time`

GetLockoutEnd returns the LockoutEnd field if non-nil, zero value otherwise.

### GetLockoutEndOk

`func (o *ExtendedUserDto) GetLockoutEndOk() (*time.Time, bool)`

GetLockoutEndOk returns a tuple with the LockoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnd

`func (o *ExtendedUserDto) SetLockoutEnd(v time.Time)`

SetLockoutEnd sets LockoutEnd field to given value.

### HasLockoutEnd

`func (o *ExtendedUserDto) HasLockoutEnd() bool`

HasLockoutEnd returns a boolean if a field has been set.

### SetLockoutEndNil

`func (o *ExtendedUserDto) SetLockoutEndNil(b bool)`

 SetLockoutEndNil sets the value for LockoutEnd to be an explicit nil

### UnsetLockoutEnd
`func (o *ExtendedUserDto) UnsetLockoutEnd()`

UnsetLockoutEnd ensures that no value is present for LockoutEnd, not even an explicit nil
### GetEnrollmentsCount

`func (o *ExtendedUserDto) GetEnrollmentsCount() int32`

GetEnrollmentsCount returns the EnrollmentsCount field if non-nil, zero value otherwise.

### GetEnrollmentsCountOk

`func (o *ExtendedUserDto) GetEnrollmentsCountOk() (*int32, bool)`

GetEnrollmentsCountOk returns a tuple with the EnrollmentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentsCount

`func (o *ExtendedUserDto) SetEnrollmentsCount(v int32)`

SetEnrollmentsCount sets EnrollmentsCount field to given value.

### HasEnrollmentsCount

`func (o *ExtendedUserDto) HasEnrollmentsCount() bool`

HasEnrollmentsCount returns a boolean if a field has been set.

### SetEnrollmentsCountNil

`func (o *ExtendedUserDto) SetEnrollmentsCountNil(b bool)`

 SetEnrollmentsCountNil sets the value for EnrollmentsCount to be an explicit nil

### UnsetEnrollmentsCount
`func (o *ExtendedUserDto) UnsetEnrollmentsCount()`

UnsetEnrollmentsCount ensures that no value is present for EnrollmentsCount, not even an explicit nil
### GetSiteTheme

`func (o *ExtendedUserDto) GetSiteTheme() int32`

GetSiteTheme returns the SiteTheme field if non-nil, zero value otherwise.

### GetSiteThemeOk

`func (o *ExtendedUserDto) GetSiteThemeOk() (*int32, bool)`

GetSiteThemeOk returns a tuple with the SiteTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTheme

`func (o *ExtendedUserDto) SetSiteTheme(v int32)`

SetSiteTheme sets SiteTheme field to given value.

### HasSiteTheme

`func (o *ExtendedUserDto) HasSiteTheme() bool`

HasSiteTheme returns a boolean if a field has been set.

### SetSiteThemeNil

`func (o *ExtendedUserDto) SetSiteThemeNil(b bool)`

 SetSiteThemeNil sets the value for SiteTheme to be an explicit nil

### UnsetSiteTheme
`func (o *ExtendedUserDto) UnsetSiteTheme()`

UnsetSiteTheme ensures that no value is present for SiteTheme, not even an explicit nil
### GetCart

`func (o *ExtendedUserDto) GetCart() CartDto`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *ExtendedUserDto) GetCartOk() (*CartDto, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *ExtendedUserDto) SetCart(v CartDto)`

SetCart sets Cart field to given value.

### HasCart

`func (o *ExtendedUserDto) HasCart() bool`

HasCart returns a boolean if a field has been set.

### GetWallet

`func (o *ExtendedUserDto) GetWallet() WalletDto`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *ExtendedUserDto) GetWalletOk() (*WalletDto, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *ExtendedUserDto) SetWallet(v WalletDto)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *ExtendedUserDto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetSocialProfile

`func (o *ExtendedUserDto) GetSocialProfile() SocialProfileDto`

GetSocialProfile returns the SocialProfile field if non-nil, zero value otherwise.

### GetSocialProfileOk

`func (o *ExtendedUserDto) GetSocialProfileOk() (*SocialProfileDto, bool)`

GetSocialProfileOk returns a tuple with the SocialProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfile

`func (o *ExtendedUserDto) SetSocialProfile(v SocialProfileDto)`

SetSocialProfile sets SocialProfile field to given value.

### HasSocialProfile

`func (o *ExtendedUserDto) HasSocialProfile() bool`

HasSocialProfile returns a boolean if a field has been set.

### GetSettings

`func (o *ExtendedUserDto) GetSettings() UserSettingsDto`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ExtendedUserDto) GetSettingsOk() (*UserSettingsDto, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ExtendedUserDto) SetSettings(v UserSettingsDto)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ExtendedUserDto) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


