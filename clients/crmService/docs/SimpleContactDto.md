# SimpleContactDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**QualifiedName** | Pointer to **NullableString** |  | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**PublicName** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**JobTitle** | Pointer to **NullableString** |  | [optional] 
**CoverUrl** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**TimezoneId** | Pointer to **NullableString** |  | [optional] 
**LanguageId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**WebUrl** | Pointer to **NullableString** |  | [optional] 
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

## Methods

### NewSimpleContactDto

`func NewSimpleContactDto() *SimpleContactDto`

NewSimpleContactDto instantiates a new SimpleContactDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleContactDtoWithDefaults

`func NewSimpleContactDtoWithDefaults() *SimpleContactDto`

NewSimpleContactDtoWithDefaults instantiates a new SimpleContactDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleContactDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleContactDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleContactDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleContactDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SimpleContactDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SimpleContactDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SimpleContactDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SimpleContactDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SimpleContactDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SimpleContactDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SimpleContactDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SimpleContactDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQualifiedName

`func (o *SimpleContactDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *SimpleContactDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *SimpleContactDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *SimpleContactDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *SimpleContactDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *SimpleContactDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetTenantId

`func (o *SimpleContactDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SimpleContactDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SimpleContactDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SimpleContactDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SimpleContactDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SimpleContactDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *SimpleContactDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleContactDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleContactDto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleContactDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPublicName

`func (o *SimpleContactDto) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *SimpleContactDto) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *SimpleContactDto) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *SimpleContactDto) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *SimpleContactDto) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *SimpleContactDto) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetFirstName

`func (o *SimpleContactDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SimpleContactDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SimpleContactDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SimpleContactDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *SimpleContactDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *SimpleContactDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *SimpleContactDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SimpleContactDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SimpleContactDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SimpleContactDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *SimpleContactDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *SimpleContactDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetJobTitle

`func (o *SimpleContactDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *SimpleContactDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *SimpleContactDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *SimpleContactDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *SimpleContactDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *SimpleContactDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetCoverUrl

`func (o *SimpleContactDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *SimpleContactDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *SimpleContactDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *SimpleContactDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *SimpleContactDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *SimpleContactDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetAvatarUrl

`func (o *SimpleContactDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *SimpleContactDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *SimpleContactDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *SimpleContactDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *SimpleContactDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *SimpleContactDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetCountryId

`func (o *SimpleContactDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *SimpleContactDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *SimpleContactDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *SimpleContactDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *SimpleContactDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *SimpleContactDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTimezoneId

`func (o *SimpleContactDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *SimpleContactDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *SimpleContactDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *SimpleContactDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *SimpleContactDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *SimpleContactDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetLanguageId

`func (o *SimpleContactDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *SimpleContactDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *SimpleContactDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *SimpleContactDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *SimpleContactDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *SimpleContactDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetSocialProfileId

`func (o *SimpleContactDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SimpleContactDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SimpleContactDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SimpleContactDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SimpleContactDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SimpleContactDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetWebUrl

`func (o *SimpleContactDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *SimpleContactDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *SimpleContactDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *SimpleContactDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *SimpleContactDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *SimpleContactDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetGitHubUrl

`func (o *SimpleContactDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *SimpleContactDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *SimpleContactDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *SimpleContactDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *SimpleContactDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *SimpleContactDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetTwitchUrl

`func (o *SimpleContactDto) GetTwitchUrl() string`

GetTwitchUrl returns the TwitchUrl field if non-nil, zero value otherwise.

### GetTwitchUrlOk

`func (o *SimpleContactDto) GetTwitchUrlOk() (*string, bool)`

GetTwitchUrlOk returns a tuple with the TwitchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchUrl

`func (o *SimpleContactDto) SetTwitchUrl(v string)`

SetTwitchUrl sets TwitchUrl field to given value.

### HasTwitchUrl

`func (o *SimpleContactDto) HasTwitchUrl() bool`

HasTwitchUrl returns a boolean if a field has been set.

### SetTwitchUrlNil

`func (o *SimpleContactDto) SetTwitchUrlNil(b bool)`

 SetTwitchUrlNil sets the value for TwitchUrl to be an explicit nil

### UnsetTwitchUrl
`func (o *SimpleContactDto) UnsetTwitchUrl()`

UnsetTwitchUrl ensures that no value is present for TwitchUrl, not even an explicit nil
### GetRedditUrl

`func (o *SimpleContactDto) GetRedditUrl() string`

GetRedditUrl returns the RedditUrl field if non-nil, zero value otherwise.

### GetRedditUrlOk

`func (o *SimpleContactDto) GetRedditUrlOk() (*string, bool)`

GetRedditUrlOk returns a tuple with the RedditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedditUrl

`func (o *SimpleContactDto) SetRedditUrl(v string)`

SetRedditUrl sets RedditUrl field to given value.

### HasRedditUrl

`func (o *SimpleContactDto) HasRedditUrl() bool`

HasRedditUrl returns a boolean if a field has been set.

### SetRedditUrlNil

`func (o *SimpleContactDto) SetRedditUrlNil(b bool)`

 SetRedditUrlNil sets the value for RedditUrl to be an explicit nil

### UnsetRedditUrl
`func (o *SimpleContactDto) UnsetRedditUrl()`

UnsetRedditUrl ensures that no value is present for RedditUrl, not even an explicit nil
### GetTikTokUrl

`func (o *SimpleContactDto) GetTikTokUrl() string`

GetTikTokUrl returns the TikTokUrl field if non-nil, zero value otherwise.

### GetTikTokUrlOk

`func (o *SimpleContactDto) GetTikTokUrlOk() (*string, bool)`

GetTikTokUrlOk returns a tuple with the TikTokUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTokUrl

`func (o *SimpleContactDto) SetTikTokUrl(v string)`

SetTikTokUrl sets TikTokUrl field to given value.

### HasTikTokUrl

`func (o *SimpleContactDto) HasTikTokUrl() bool`

HasTikTokUrl returns a boolean if a field has been set.

### SetTikTokUrlNil

`func (o *SimpleContactDto) SetTikTokUrlNil(b bool)`

 SetTikTokUrlNil sets the value for TikTokUrl to be an explicit nil

### UnsetTikTokUrl
`func (o *SimpleContactDto) UnsetTikTokUrl()`

UnsetTikTokUrl ensures that no value is present for TikTokUrl, not even an explicit nil
### GetWebsiteUrl

`func (o *SimpleContactDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *SimpleContactDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *SimpleContactDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *SimpleContactDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *SimpleContactDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *SimpleContactDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetTwitterUrl

`func (o *SimpleContactDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *SimpleContactDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *SimpleContactDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *SimpleContactDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *SimpleContactDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *SimpleContactDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetFacebookUrl

`func (o *SimpleContactDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *SimpleContactDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *SimpleContactDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *SimpleContactDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *SimpleContactDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *SimpleContactDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *SimpleContactDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *SimpleContactDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *SimpleContactDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *SimpleContactDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *SimpleContactDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *SimpleContactDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *SimpleContactDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *SimpleContactDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *SimpleContactDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *SimpleContactDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *SimpleContactDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *SimpleContactDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *SimpleContactDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *SimpleContactDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *SimpleContactDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *SimpleContactDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *SimpleContactDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *SimpleContactDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetGithubUsername

`func (o *SimpleContactDto) GetGithubUsername() string`

GetGithubUsername returns the GithubUsername field if non-nil, zero value otherwise.

### GetGithubUsernameOk

`func (o *SimpleContactDto) GetGithubUsernameOk() (*string, bool)`

GetGithubUsernameOk returns a tuple with the GithubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUsername

`func (o *SimpleContactDto) SetGithubUsername(v string)`

SetGithubUsername sets GithubUsername field to given value.

### HasGithubUsername

`func (o *SimpleContactDto) HasGithubUsername() bool`

HasGithubUsername returns a boolean if a field has been set.

### SetGithubUsernameNil

`func (o *SimpleContactDto) SetGithubUsernameNil(b bool)`

 SetGithubUsernameNil sets the value for GithubUsername to be an explicit nil

### UnsetGithubUsername
`func (o *SimpleContactDto) UnsetGithubUsername()`

UnsetGithubUsername ensures that no value is present for GithubUsername, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


