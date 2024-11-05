# ContactDto

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
**Duns** | Pointer to **NullableString** |  | [optional] 
**TaxId** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**Street** | Pointer to **NullableString** |  | [optional] 
**CartId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**ZipCode** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**WalletId** | Pointer to **NullableString** |  | [optional] 
**FaxNumber** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**StreetLine1** | Pointer to **NullableString** |  | [optional] 
**StreetLine2** | Pointer to **NullableString** |  | [optional] 
**TerritoryId** | Pointer to **NullableString** |  | [optional] 
**MobilePhone** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**AnnualRevenue** | Pointer to **NullableString** |  | [optional] 
**RelatedUserId** | Pointer to **NullableString** |  | [optional] 
**BusinessPhone** | Pointer to **NullableString** |  | [optional] 
**OwnerContactId** | Pointer to **NullableString** |  | [optional] 
**RelatedTenantId** | Pointer to **NullableString** |  | [optional] 
**ActivityFeedId** | Pointer to **NullableString** |  | [optional] 
**ParentContactId** | Pointer to **NullableString** |  | [optional] 
**IdentityProvider** | Pointer to **NullableString** |  | [optional] 
**PartnerProfileId** | Pointer to **NullableString** |  | [optional] 
**PrimaryContactId** | Pointer to **NullableString** |  | [optional] 
**ActiveDirectoryId** | Pointer to **NullableString** |  | [optional] 
**IdentityProviderAccessToken** | Pointer to **NullableString** |  | [optional] 
**Birthday** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewContactDto

`func NewContactDto() *ContactDto`

NewContactDto instantiates a new ContactDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactDtoWithDefaults

`func NewContactDtoWithDefaults() *ContactDto`

NewContactDtoWithDefaults instantiates a new ContactDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ContactDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ContactDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ContactDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContactDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContactDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContactDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ContactDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ContactDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQualifiedName

`func (o *ContactDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *ContactDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *ContactDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *ContactDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *ContactDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *ContactDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetTenantId

`func (o *ContactDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContactDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContactDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ContactDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ContactDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ContactDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *ContactDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactDto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ContactDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPublicName

`func (o *ContactDto) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *ContactDto) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *ContactDto) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *ContactDto) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *ContactDto) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *ContactDto) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetFirstName

`func (o *ContactDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ContactDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ContactDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ContactDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ContactDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ContactDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ContactDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ContactDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetJobTitle

`func (o *ContactDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *ContactDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *ContactDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *ContactDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *ContactDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *ContactDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetCoverUrl

`func (o *ContactDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *ContactDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *ContactDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *ContactDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *ContactDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *ContactDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetAvatarUrl

`func (o *ContactDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ContactDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ContactDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ContactDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *ContactDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *ContactDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetCountryId

`func (o *ContactDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ContactDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ContactDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ContactDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ContactDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ContactDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTimezoneId

`func (o *ContactDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ContactDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ContactDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ContactDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ContactDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ContactDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetLanguageId

`func (o *ContactDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *ContactDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *ContactDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *ContactDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *ContactDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *ContactDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetSocialProfileId

`func (o *ContactDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ContactDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ContactDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ContactDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ContactDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ContactDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetWebUrl

`func (o *ContactDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ContactDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ContactDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *ContactDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *ContactDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *ContactDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetGitHubUrl

`func (o *ContactDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *ContactDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *ContactDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *ContactDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *ContactDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *ContactDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetTwitchUrl

`func (o *ContactDto) GetTwitchUrl() string`

GetTwitchUrl returns the TwitchUrl field if non-nil, zero value otherwise.

### GetTwitchUrlOk

`func (o *ContactDto) GetTwitchUrlOk() (*string, bool)`

GetTwitchUrlOk returns a tuple with the TwitchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchUrl

`func (o *ContactDto) SetTwitchUrl(v string)`

SetTwitchUrl sets TwitchUrl field to given value.

### HasTwitchUrl

`func (o *ContactDto) HasTwitchUrl() bool`

HasTwitchUrl returns a boolean if a field has been set.

### SetTwitchUrlNil

`func (o *ContactDto) SetTwitchUrlNil(b bool)`

 SetTwitchUrlNil sets the value for TwitchUrl to be an explicit nil

### UnsetTwitchUrl
`func (o *ContactDto) UnsetTwitchUrl()`

UnsetTwitchUrl ensures that no value is present for TwitchUrl, not even an explicit nil
### GetRedditUrl

`func (o *ContactDto) GetRedditUrl() string`

GetRedditUrl returns the RedditUrl field if non-nil, zero value otherwise.

### GetRedditUrlOk

`func (o *ContactDto) GetRedditUrlOk() (*string, bool)`

GetRedditUrlOk returns a tuple with the RedditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedditUrl

`func (o *ContactDto) SetRedditUrl(v string)`

SetRedditUrl sets RedditUrl field to given value.

### HasRedditUrl

`func (o *ContactDto) HasRedditUrl() bool`

HasRedditUrl returns a boolean if a field has been set.

### SetRedditUrlNil

`func (o *ContactDto) SetRedditUrlNil(b bool)`

 SetRedditUrlNil sets the value for RedditUrl to be an explicit nil

### UnsetRedditUrl
`func (o *ContactDto) UnsetRedditUrl()`

UnsetRedditUrl ensures that no value is present for RedditUrl, not even an explicit nil
### GetTikTokUrl

`func (o *ContactDto) GetTikTokUrl() string`

GetTikTokUrl returns the TikTokUrl field if non-nil, zero value otherwise.

### GetTikTokUrlOk

`func (o *ContactDto) GetTikTokUrlOk() (*string, bool)`

GetTikTokUrlOk returns a tuple with the TikTokUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTokUrl

`func (o *ContactDto) SetTikTokUrl(v string)`

SetTikTokUrl sets TikTokUrl field to given value.

### HasTikTokUrl

`func (o *ContactDto) HasTikTokUrl() bool`

HasTikTokUrl returns a boolean if a field has been set.

### SetTikTokUrlNil

`func (o *ContactDto) SetTikTokUrlNil(b bool)`

 SetTikTokUrlNil sets the value for TikTokUrl to be an explicit nil

### UnsetTikTokUrl
`func (o *ContactDto) UnsetTikTokUrl()`

UnsetTikTokUrl ensures that no value is present for TikTokUrl, not even an explicit nil
### GetWebsiteUrl

`func (o *ContactDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ContactDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ContactDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ContactDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *ContactDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *ContactDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetTwitterUrl

`func (o *ContactDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *ContactDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *ContactDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *ContactDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *ContactDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *ContactDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetFacebookUrl

`func (o *ContactDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *ContactDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *ContactDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *ContactDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *ContactDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *ContactDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *ContactDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *ContactDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *ContactDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *ContactDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *ContactDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *ContactDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *ContactDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *ContactDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *ContactDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *ContactDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *ContactDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *ContactDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *ContactDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *ContactDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *ContactDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *ContactDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *ContactDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *ContactDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetGithubUsername

`func (o *ContactDto) GetGithubUsername() string`

GetGithubUsername returns the GithubUsername field if non-nil, zero value otherwise.

### GetGithubUsernameOk

`func (o *ContactDto) GetGithubUsernameOk() (*string, bool)`

GetGithubUsernameOk returns a tuple with the GithubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUsername

`func (o *ContactDto) SetGithubUsername(v string)`

SetGithubUsername sets GithubUsername field to given value.

### HasGithubUsername

`func (o *ContactDto) HasGithubUsername() bool`

HasGithubUsername returns a boolean if a field has been set.

### SetGithubUsernameNil

`func (o *ContactDto) SetGithubUsernameNil(b bool)`

 SetGithubUsernameNil sets the value for GithubUsername to be an explicit nil

### UnsetGithubUsername
`func (o *ContactDto) UnsetGithubUsername()`

UnsetGithubUsername ensures that no value is present for GithubUsername, not even an explicit nil
### GetDuns

`func (o *ContactDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *ContactDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *ContactDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *ContactDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *ContactDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *ContactDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetTaxId

`func (o *ContactDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *ContactDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *ContactDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *ContactDto) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *ContactDto) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *ContactDto) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetEmail

`func (o *ContactDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ContactDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ContactDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAbout

`func (o *ContactDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ContactDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ContactDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *ContactDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *ContactDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *ContactDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetStreet

`func (o *ContactDto) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ContactDto) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ContactDto) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ContactDto) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *ContactDto) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *ContactDto) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetCartId

`func (o *ContactDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *ContactDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *ContactDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *ContactDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *ContactDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *ContactDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetCityId

`func (o *ContactDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ContactDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ContactDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ContactDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ContactDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ContactDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetZipCode

`func (o *ContactDto) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ContactDto) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ContactDto) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ContactDto) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *ContactDto) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *ContactDto) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetStateId

`func (o *ContactDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *ContactDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *ContactDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *ContactDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *ContactDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *ContactDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetWalletId

`func (o *ContactDto) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ContactDto) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ContactDto) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *ContactDto) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *ContactDto) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *ContactDto) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetFaxNumber

`func (o *ContactDto) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *ContactDto) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *ContactDto) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *ContactDto) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### SetFaxNumberNil

`func (o *ContactDto) SetFaxNumberNil(b bool)`

 SetFaxNumberNil sets the value for FaxNumber to be an explicit nil

### UnsetFaxNumber
`func (o *ContactDto) UnsetFaxNumber()`

UnsetFaxNumber ensures that no value is present for FaxNumber, not even an explicit nil
### GetPostalCode

`func (o *ContactDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ContactDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ContactDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ContactDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ContactDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ContactDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCurrencyId

`func (o *ContactDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ContactDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ContactDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ContactDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ContactDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ContactDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetStreetLine1

`func (o *ContactDto) GetStreetLine1() string`

GetStreetLine1 returns the StreetLine1 field if non-nil, zero value otherwise.

### GetStreetLine1Ok

`func (o *ContactDto) GetStreetLine1Ok() (*string, bool)`

GetStreetLine1Ok returns a tuple with the StreetLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine1

`func (o *ContactDto) SetStreetLine1(v string)`

SetStreetLine1 sets StreetLine1 field to given value.

### HasStreetLine1

`func (o *ContactDto) HasStreetLine1() bool`

HasStreetLine1 returns a boolean if a field has been set.

### SetStreetLine1Nil

`func (o *ContactDto) SetStreetLine1Nil(b bool)`

 SetStreetLine1Nil sets the value for StreetLine1 to be an explicit nil

### UnsetStreetLine1
`func (o *ContactDto) UnsetStreetLine1()`

UnsetStreetLine1 ensures that no value is present for StreetLine1, not even an explicit nil
### GetStreetLine2

`func (o *ContactDto) GetStreetLine2() string`

GetStreetLine2 returns the StreetLine2 field if non-nil, zero value otherwise.

### GetStreetLine2Ok

`func (o *ContactDto) GetStreetLine2Ok() (*string, bool)`

GetStreetLine2Ok returns a tuple with the StreetLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine2

`func (o *ContactDto) SetStreetLine2(v string)`

SetStreetLine2 sets StreetLine2 field to given value.

### HasStreetLine2

`func (o *ContactDto) HasStreetLine2() bool`

HasStreetLine2 returns a boolean if a field has been set.

### SetStreetLine2Nil

`func (o *ContactDto) SetStreetLine2Nil(b bool)`

 SetStreetLine2Nil sets the value for StreetLine2 to be an explicit nil

### UnsetStreetLine2
`func (o *ContactDto) UnsetStreetLine2()`

UnsetStreetLine2 ensures that no value is present for StreetLine2, not even an explicit nil
### GetTerritoryId

`func (o *ContactDto) GetTerritoryId() string`

GetTerritoryId returns the TerritoryId field if non-nil, zero value otherwise.

### GetTerritoryIdOk

`func (o *ContactDto) GetTerritoryIdOk() (*string, bool)`

GetTerritoryIdOk returns a tuple with the TerritoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritoryId

`func (o *ContactDto) SetTerritoryId(v string)`

SetTerritoryId sets TerritoryId field to given value.

### HasTerritoryId

`func (o *ContactDto) HasTerritoryId() bool`

HasTerritoryId returns a boolean if a field has been set.

### SetTerritoryIdNil

`func (o *ContactDto) SetTerritoryIdNil(b bool)`

 SetTerritoryIdNil sets the value for TerritoryId to be an explicit nil

### UnsetTerritoryId
`func (o *ContactDto) UnsetTerritoryId()`

UnsetTerritoryId ensures that no value is present for TerritoryId, not even an explicit nil
### GetMobilePhone

`func (o *ContactDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ContactDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ContactDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ContactDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *ContactDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *ContactDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetEnrollmentId

`func (o *ContactDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ContactDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ContactDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ContactDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ContactDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ContactDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAnnualRevenue

`func (o *ContactDto) GetAnnualRevenue() string`

GetAnnualRevenue returns the AnnualRevenue field if non-nil, zero value otherwise.

### GetAnnualRevenueOk

`func (o *ContactDto) GetAnnualRevenueOk() (*string, bool)`

GetAnnualRevenueOk returns a tuple with the AnnualRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualRevenue

`func (o *ContactDto) SetAnnualRevenue(v string)`

SetAnnualRevenue sets AnnualRevenue field to given value.

### HasAnnualRevenue

`func (o *ContactDto) HasAnnualRevenue() bool`

HasAnnualRevenue returns a boolean if a field has been set.

### SetAnnualRevenueNil

`func (o *ContactDto) SetAnnualRevenueNil(b bool)`

 SetAnnualRevenueNil sets the value for AnnualRevenue to be an explicit nil

### UnsetAnnualRevenue
`func (o *ContactDto) UnsetAnnualRevenue()`

UnsetAnnualRevenue ensures that no value is present for AnnualRevenue, not even an explicit nil
### GetRelatedUserId

`func (o *ContactDto) GetRelatedUserId() string`

GetRelatedUserId returns the RelatedUserId field if non-nil, zero value otherwise.

### GetRelatedUserIdOk

`func (o *ContactDto) GetRelatedUserIdOk() (*string, bool)`

GetRelatedUserIdOk returns a tuple with the RelatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedUserId

`func (o *ContactDto) SetRelatedUserId(v string)`

SetRelatedUserId sets RelatedUserId field to given value.

### HasRelatedUserId

`func (o *ContactDto) HasRelatedUserId() bool`

HasRelatedUserId returns a boolean if a field has been set.

### SetRelatedUserIdNil

`func (o *ContactDto) SetRelatedUserIdNil(b bool)`

 SetRelatedUserIdNil sets the value for RelatedUserId to be an explicit nil

### UnsetRelatedUserId
`func (o *ContactDto) UnsetRelatedUserId()`

UnsetRelatedUserId ensures that no value is present for RelatedUserId, not even an explicit nil
### GetBusinessPhone

`func (o *ContactDto) GetBusinessPhone() string`

GetBusinessPhone returns the BusinessPhone field if non-nil, zero value otherwise.

### GetBusinessPhoneOk

`func (o *ContactDto) GetBusinessPhoneOk() (*string, bool)`

GetBusinessPhoneOk returns a tuple with the BusinessPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhone

`func (o *ContactDto) SetBusinessPhone(v string)`

SetBusinessPhone sets BusinessPhone field to given value.

### HasBusinessPhone

`func (o *ContactDto) HasBusinessPhone() bool`

HasBusinessPhone returns a boolean if a field has been set.

### SetBusinessPhoneNil

`func (o *ContactDto) SetBusinessPhoneNil(b bool)`

 SetBusinessPhoneNil sets the value for BusinessPhone to be an explicit nil

### UnsetBusinessPhone
`func (o *ContactDto) UnsetBusinessPhone()`

UnsetBusinessPhone ensures that no value is present for BusinessPhone, not even an explicit nil
### GetOwnerContactId

`func (o *ContactDto) GetOwnerContactId() string`

GetOwnerContactId returns the OwnerContactId field if non-nil, zero value otherwise.

### GetOwnerContactIdOk

`func (o *ContactDto) GetOwnerContactIdOk() (*string, bool)`

GetOwnerContactIdOk returns a tuple with the OwnerContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContactId

`func (o *ContactDto) SetOwnerContactId(v string)`

SetOwnerContactId sets OwnerContactId field to given value.

### HasOwnerContactId

`func (o *ContactDto) HasOwnerContactId() bool`

HasOwnerContactId returns a boolean if a field has been set.

### SetOwnerContactIdNil

`func (o *ContactDto) SetOwnerContactIdNil(b bool)`

 SetOwnerContactIdNil sets the value for OwnerContactId to be an explicit nil

### UnsetOwnerContactId
`func (o *ContactDto) UnsetOwnerContactId()`

UnsetOwnerContactId ensures that no value is present for OwnerContactId, not even an explicit nil
### GetRelatedTenantId

`func (o *ContactDto) GetRelatedTenantId() string`

GetRelatedTenantId returns the RelatedTenantId field if non-nil, zero value otherwise.

### GetRelatedTenantIdOk

`func (o *ContactDto) GetRelatedTenantIdOk() (*string, bool)`

GetRelatedTenantIdOk returns a tuple with the RelatedTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTenantId

`func (o *ContactDto) SetRelatedTenantId(v string)`

SetRelatedTenantId sets RelatedTenantId field to given value.

### HasRelatedTenantId

`func (o *ContactDto) HasRelatedTenantId() bool`

HasRelatedTenantId returns a boolean if a field has been set.

### SetRelatedTenantIdNil

`func (o *ContactDto) SetRelatedTenantIdNil(b bool)`

 SetRelatedTenantIdNil sets the value for RelatedTenantId to be an explicit nil

### UnsetRelatedTenantId
`func (o *ContactDto) UnsetRelatedTenantId()`

UnsetRelatedTenantId ensures that no value is present for RelatedTenantId, not even an explicit nil
### GetActivityFeedId

`func (o *ContactDto) GetActivityFeedId() string`

GetActivityFeedId returns the ActivityFeedId field if non-nil, zero value otherwise.

### GetActivityFeedIdOk

`func (o *ContactDto) GetActivityFeedIdOk() (*string, bool)`

GetActivityFeedIdOk returns a tuple with the ActivityFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityFeedId

`func (o *ContactDto) SetActivityFeedId(v string)`

SetActivityFeedId sets ActivityFeedId field to given value.

### HasActivityFeedId

`func (o *ContactDto) HasActivityFeedId() bool`

HasActivityFeedId returns a boolean if a field has been set.

### SetActivityFeedIdNil

`func (o *ContactDto) SetActivityFeedIdNil(b bool)`

 SetActivityFeedIdNil sets the value for ActivityFeedId to be an explicit nil

### UnsetActivityFeedId
`func (o *ContactDto) UnsetActivityFeedId()`

UnsetActivityFeedId ensures that no value is present for ActivityFeedId, not even an explicit nil
### GetParentContactId

`func (o *ContactDto) GetParentContactId() string`

GetParentContactId returns the ParentContactId field if non-nil, zero value otherwise.

### GetParentContactIdOk

`func (o *ContactDto) GetParentContactIdOk() (*string, bool)`

GetParentContactIdOk returns a tuple with the ParentContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentContactId

`func (o *ContactDto) SetParentContactId(v string)`

SetParentContactId sets ParentContactId field to given value.

### HasParentContactId

`func (o *ContactDto) HasParentContactId() bool`

HasParentContactId returns a boolean if a field has been set.

### SetParentContactIdNil

`func (o *ContactDto) SetParentContactIdNil(b bool)`

 SetParentContactIdNil sets the value for ParentContactId to be an explicit nil

### UnsetParentContactId
`func (o *ContactDto) UnsetParentContactId()`

UnsetParentContactId ensures that no value is present for ParentContactId, not even an explicit nil
### GetIdentityProvider

`func (o *ContactDto) GetIdentityProvider() string`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *ContactDto) GetIdentityProviderOk() (*string, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *ContactDto) SetIdentityProvider(v string)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *ContactDto) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### SetIdentityProviderNil

`func (o *ContactDto) SetIdentityProviderNil(b bool)`

 SetIdentityProviderNil sets the value for IdentityProvider to be an explicit nil

### UnsetIdentityProvider
`func (o *ContactDto) UnsetIdentityProvider()`

UnsetIdentityProvider ensures that no value is present for IdentityProvider, not even an explicit nil
### GetPartnerProfileId

`func (o *ContactDto) GetPartnerProfileId() string`

GetPartnerProfileId returns the PartnerProfileId field if non-nil, zero value otherwise.

### GetPartnerProfileIdOk

`func (o *ContactDto) GetPartnerProfileIdOk() (*string, bool)`

GetPartnerProfileIdOk returns a tuple with the PartnerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProfileId

`func (o *ContactDto) SetPartnerProfileId(v string)`

SetPartnerProfileId sets PartnerProfileId field to given value.

### HasPartnerProfileId

`func (o *ContactDto) HasPartnerProfileId() bool`

HasPartnerProfileId returns a boolean if a field has been set.

### SetPartnerProfileIdNil

`func (o *ContactDto) SetPartnerProfileIdNil(b bool)`

 SetPartnerProfileIdNil sets the value for PartnerProfileId to be an explicit nil

### UnsetPartnerProfileId
`func (o *ContactDto) UnsetPartnerProfileId()`

UnsetPartnerProfileId ensures that no value is present for PartnerProfileId, not even an explicit nil
### GetPrimaryContactId

`func (o *ContactDto) GetPrimaryContactId() string`

GetPrimaryContactId returns the PrimaryContactId field if non-nil, zero value otherwise.

### GetPrimaryContactIdOk

`func (o *ContactDto) GetPrimaryContactIdOk() (*string, bool)`

GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactId

`func (o *ContactDto) SetPrimaryContactId(v string)`

SetPrimaryContactId sets PrimaryContactId field to given value.

### HasPrimaryContactId

`func (o *ContactDto) HasPrimaryContactId() bool`

HasPrimaryContactId returns a boolean if a field has been set.

### SetPrimaryContactIdNil

`func (o *ContactDto) SetPrimaryContactIdNil(b bool)`

 SetPrimaryContactIdNil sets the value for PrimaryContactId to be an explicit nil

### UnsetPrimaryContactId
`func (o *ContactDto) UnsetPrimaryContactId()`

UnsetPrimaryContactId ensures that no value is present for PrimaryContactId, not even an explicit nil
### GetActiveDirectoryId

`func (o *ContactDto) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *ContactDto) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *ContactDto) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.

### HasActiveDirectoryId

`func (o *ContactDto) HasActiveDirectoryId() bool`

HasActiveDirectoryId returns a boolean if a field has been set.

### SetActiveDirectoryIdNil

`func (o *ContactDto) SetActiveDirectoryIdNil(b bool)`

 SetActiveDirectoryIdNil sets the value for ActiveDirectoryId to be an explicit nil

### UnsetActiveDirectoryId
`func (o *ContactDto) UnsetActiveDirectoryId()`

UnsetActiveDirectoryId ensures that no value is present for ActiveDirectoryId, not even an explicit nil
### GetIdentityProviderAccessToken

`func (o *ContactDto) GetIdentityProviderAccessToken() string`

GetIdentityProviderAccessToken returns the IdentityProviderAccessToken field if non-nil, zero value otherwise.

### GetIdentityProviderAccessTokenOk

`func (o *ContactDto) GetIdentityProviderAccessTokenOk() (*string, bool)`

GetIdentityProviderAccessTokenOk returns a tuple with the IdentityProviderAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderAccessToken

`func (o *ContactDto) SetIdentityProviderAccessToken(v string)`

SetIdentityProviderAccessToken sets IdentityProviderAccessToken field to given value.

### HasIdentityProviderAccessToken

`func (o *ContactDto) HasIdentityProviderAccessToken() bool`

HasIdentityProviderAccessToken returns a boolean if a field has been set.

### SetIdentityProviderAccessTokenNil

`func (o *ContactDto) SetIdentityProviderAccessTokenNil(b bool)`

 SetIdentityProviderAccessTokenNil sets the value for IdentityProviderAccessToken to be an explicit nil

### UnsetIdentityProviderAccessToken
`func (o *ContactDto) UnsetIdentityProviderAccessToken()`

UnsetIdentityProviderAccessToken ensures that no value is present for IdentityProviderAccessToken, not even an explicit nil
### GetBirthday

`func (o *ContactDto) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *ContactDto) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *ContactDto) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *ContactDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *ContactDto) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *ContactDto) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


