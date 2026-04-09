# BlogAuthorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] 
**Birthday** | Pointer to **time.Time** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**PublicName** | Pointer to **NullableString** |  | [optional] 
**LanguageId** | Pointer to **NullableString** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**CoverUrl** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**GitHubUrl** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**TwitterUrl** | Pointer to **NullableString** |  | [optional] 
**FacebookUrl** | Pointer to **NullableString** |  | [optional] 
**YouTubeUrl** | Pointer to **NullableString** |  | [optional] 
**LinkedInUrl** | Pointer to **NullableString** |  | [optional] 
**InstagramUrl** | Pointer to **NullableString** |  | [optional] 
**SocialFeedId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**WalletId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**PublicIdentifier** | Pointer to **NullableString** |  | [optional] 
**PhoneNumberConfirmed** | Pointer to **bool** |  | [optional] 
**Availability** | Pointer to **string** |  | [optional] 

## Methods

### NewBlogAuthorDto

`func NewBlogAuthorDto() *BlogAuthorDto`

NewBlogAuthorDto instantiates a new BlogAuthorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlogAuthorDtoWithDefaults

`func NewBlogAuthorDtoWithDefaults() *BlogAuthorDto`

NewBlogAuthorDtoWithDefaults instantiates a new BlogAuthorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlogAuthorDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlogAuthorDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlogAuthorDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlogAuthorDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BlogAuthorDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BlogAuthorDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BlogAuthorDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BlogAuthorDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BlogAuthorDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BlogAuthorDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BlogAuthorDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BlogAuthorDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQualifiedName

`func (o *BlogAuthorDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *BlogAuthorDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *BlogAuthorDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *BlogAuthorDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *BlogAuthorDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *BlogAuthorDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetBirthday

`func (o *BlogAuthorDto) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *BlogAuthorDto) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *BlogAuthorDto) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *BlogAuthorDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetFirstName

`func (o *BlogAuthorDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BlogAuthorDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BlogAuthorDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BlogAuthorDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *BlogAuthorDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *BlogAuthorDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *BlogAuthorDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BlogAuthorDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BlogAuthorDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BlogAuthorDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *BlogAuthorDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *BlogAuthorDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPublicName

`func (o *BlogAuthorDto) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *BlogAuthorDto) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *BlogAuthorDto) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *BlogAuthorDto) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *BlogAuthorDto) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *BlogAuthorDto) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetLanguageId

`func (o *BlogAuthorDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *BlogAuthorDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *BlogAuthorDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *BlogAuthorDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *BlogAuthorDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *BlogAuthorDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetTimezoneId

`func (o *BlogAuthorDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *BlogAuthorDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *BlogAuthorDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *BlogAuthorDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *BlogAuthorDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *BlogAuthorDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetCountryId

`func (o *BlogAuthorDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *BlogAuthorDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *BlogAuthorDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *BlogAuthorDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *BlogAuthorDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *BlogAuthorDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCityId

`func (o *BlogAuthorDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *BlogAuthorDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *BlogAuthorDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *BlogAuthorDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *BlogAuthorDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *BlogAuthorDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetStateId

`func (o *BlogAuthorDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *BlogAuthorDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *BlogAuthorDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *BlogAuthorDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *BlogAuthorDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *BlogAuthorDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetAbout

`func (o *BlogAuthorDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *BlogAuthorDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *BlogAuthorDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *BlogAuthorDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *BlogAuthorDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *BlogAuthorDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetJobTitle

`func (o *BlogAuthorDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *BlogAuthorDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *BlogAuthorDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *BlogAuthorDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *BlogAuthorDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *BlogAuthorDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetCoverUrl

`func (o *BlogAuthorDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *BlogAuthorDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *BlogAuthorDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *BlogAuthorDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *BlogAuthorDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *BlogAuthorDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetAvatarUrl

`func (o *BlogAuthorDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *BlogAuthorDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *BlogAuthorDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *BlogAuthorDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *BlogAuthorDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *BlogAuthorDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetGitHubUrl

`func (o *BlogAuthorDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *BlogAuthorDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *BlogAuthorDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *BlogAuthorDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *BlogAuthorDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *BlogAuthorDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetWebsiteUrl

`func (o *BlogAuthorDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *BlogAuthorDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *BlogAuthorDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *BlogAuthorDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *BlogAuthorDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *BlogAuthorDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetTwitterUrl

`func (o *BlogAuthorDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *BlogAuthorDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *BlogAuthorDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *BlogAuthorDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *BlogAuthorDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *BlogAuthorDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetFacebookUrl

`func (o *BlogAuthorDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *BlogAuthorDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *BlogAuthorDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *BlogAuthorDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *BlogAuthorDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *BlogAuthorDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *BlogAuthorDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *BlogAuthorDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *BlogAuthorDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *BlogAuthorDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *BlogAuthorDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *BlogAuthorDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *BlogAuthorDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *BlogAuthorDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *BlogAuthorDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *BlogAuthorDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *BlogAuthorDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *BlogAuthorDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *BlogAuthorDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *BlogAuthorDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *BlogAuthorDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *BlogAuthorDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *BlogAuthorDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *BlogAuthorDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetSocialFeedId

`func (o *BlogAuthorDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *BlogAuthorDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *BlogAuthorDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *BlogAuthorDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *BlogAuthorDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *BlogAuthorDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil
### GetSocialProfileId

`func (o *BlogAuthorDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *BlogAuthorDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *BlogAuthorDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *BlogAuthorDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *BlogAuthorDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *BlogAuthorDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetGender

`func (o *BlogAuthorDto) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *BlogAuthorDto) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *BlogAuthorDto) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *BlogAuthorDto) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetWalletId

`func (o *BlogAuthorDto) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *BlogAuthorDto) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *BlogAuthorDto) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *BlogAuthorDto) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *BlogAuthorDto) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *BlogAuthorDto) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetStatus

`func (o *BlogAuthorDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlogAuthorDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlogAuthorDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlogAuthorDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BlogAuthorDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BlogAuthorDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUserName

`func (o *BlogAuthorDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *BlogAuthorDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *BlogAuthorDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *BlogAuthorDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *BlogAuthorDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *BlogAuthorDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetPublicIdentifier

`func (o *BlogAuthorDto) GetPublicIdentifier() string`

GetPublicIdentifier returns the PublicIdentifier field if non-nil, zero value otherwise.

### GetPublicIdentifierOk

`func (o *BlogAuthorDto) GetPublicIdentifierOk() (*string, bool)`

GetPublicIdentifierOk returns a tuple with the PublicIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdentifier

`func (o *BlogAuthorDto) SetPublicIdentifier(v string)`

SetPublicIdentifier sets PublicIdentifier field to given value.

### HasPublicIdentifier

`func (o *BlogAuthorDto) HasPublicIdentifier() bool`

HasPublicIdentifier returns a boolean if a field has been set.

### SetPublicIdentifierNil

`func (o *BlogAuthorDto) SetPublicIdentifierNil(b bool)`

 SetPublicIdentifierNil sets the value for PublicIdentifier to be an explicit nil

### UnsetPublicIdentifier
`func (o *BlogAuthorDto) UnsetPublicIdentifier()`

UnsetPublicIdentifier ensures that no value is present for PublicIdentifier, not even an explicit nil
### GetPhoneNumberConfirmed

`func (o *BlogAuthorDto) GetPhoneNumberConfirmed() bool`

GetPhoneNumberConfirmed returns the PhoneNumberConfirmed field if non-nil, zero value otherwise.

### GetPhoneNumberConfirmedOk

`func (o *BlogAuthorDto) GetPhoneNumberConfirmedOk() (*bool, bool)`

GetPhoneNumberConfirmedOk returns a tuple with the PhoneNumberConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberConfirmed

`func (o *BlogAuthorDto) SetPhoneNumberConfirmed(v bool)`

SetPhoneNumberConfirmed sets PhoneNumberConfirmed field to given value.

### HasPhoneNumberConfirmed

`func (o *BlogAuthorDto) HasPhoneNumberConfirmed() bool`

HasPhoneNumberConfirmed returns a boolean if a field has been set.

### GetAvailability

`func (o *BlogAuthorDto) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *BlogAuthorDto) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *BlogAuthorDto) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *BlogAuthorDto) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


