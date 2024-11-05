# ExtendedContactDto

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
**Cart** | Pointer to [**CartDto**](CartDto.md) |  | [optional] 
**Wallet** | Pointer to [**WalletDto**](WalletDto.md) |  | [optional] 
**SocialProfile** | Pointer to [**SocialProfileDto**](SocialProfileDto.md) |  | [optional] 
**ParentContact** | Pointer to [**SimpleContactDto**](SimpleContactDto.md) |  | [optional] 
**PrimaryContact** | Pointer to [**SimpleContactDto**](SimpleContactDto.md) |  | [optional] 

## Methods

### NewExtendedContactDto

`func NewExtendedContactDto() *ExtendedContactDto`

NewExtendedContactDto instantiates a new ExtendedContactDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedContactDtoWithDefaults

`func NewExtendedContactDtoWithDefaults() *ExtendedContactDto`

NewExtendedContactDtoWithDefaults instantiates a new ExtendedContactDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedContactDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedContactDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedContactDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedContactDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedContactDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedContactDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ExtendedContactDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedContactDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedContactDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtendedContactDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ExtendedContactDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ExtendedContactDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQualifiedName

`func (o *ExtendedContactDto) GetQualifiedName() string`

GetQualifiedName returns the QualifiedName field if non-nil, zero value otherwise.

### GetQualifiedNameOk

`func (o *ExtendedContactDto) GetQualifiedNameOk() (*string, bool)`

GetQualifiedNameOk returns a tuple with the QualifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedName

`func (o *ExtendedContactDto) SetQualifiedName(v string)`

SetQualifiedName sets QualifiedName field to given value.

### HasQualifiedName

`func (o *ExtendedContactDto) HasQualifiedName() bool`

HasQualifiedName returns a boolean if a field has been set.

### SetQualifiedNameNil

`func (o *ExtendedContactDto) SetQualifiedNameNil(b bool)`

 SetQualifiedNameNil sets the value for QualifiedName to be an explicit nil

### UnsetQualifiedName
`func (o *ExtendedContactDto) UnsetQualifiedName()`

UnsetQualifiedName ensures that no value is present for QualifiedName, not even an explicit nil
### GetTenantId

`func (o *ExtendedContactDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtendedContactDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtendedContactDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ExtendedContactDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ExtendedContactDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ExtendedContactDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetType

`func (o *ExtendedContactDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtendedContactDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtendedContactDto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ExtendedContactDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPublicName

`func (o *ExtendedContactDto) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *ExtendedContactDto) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *ExtendedContactDto) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *ExtendedContactDto) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *ExtendedContactDto) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *ExtendedContactDto) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetFirstName

`func (o *ExtendedContactDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ExtendedContactDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ExtendedContactDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ExtendedContactDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ExtendedContactDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ExtendedContactDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *ExtendedContactDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ExtendedContactDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ExtendedContactDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ExtendedContactDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ExtendedContactDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ExtendedContactDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetJobTitle

`func (o *ExtendedContactDto) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *ExtendedContactDto) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *ExtendedContactDto) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *ExtendedContactDto) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### SetJobTitleNil

`func (o *ExtendedContactDto) SetJobTitleNil(b bool)`

 SetJobTitleNil sets the value for JobTitle to be an explicit nil

### UnsetJobTitle
`func (o *ExtendedContactDto) UnsetJobTitle()`

UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
### GetCoverUrl

`func (o *ExtendedContactDto) GetCoverUrl() string`

GetCoverUrl returns the CoverUrl field if non-nil, zero value otherwise.

### GetCoverUrlOk

`func (o *ExtendedContactDto) GetCoverUrlOk() (*string, bool)`

GetCoverUrlOk returns a tuple with the CoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverUrl

`func (o *ExtendedContactDto) SetCoverUrl(v string)`

SetCoverUrl sets CoverUrl field to given value.

### HasCoverUrl

`func (o *ExtendedContactDto) HasCoverUrl() bool`

HasCoverUrl returns a boolean if a field has been set.

### SetCoverUrlNil

`func (o *ExtendedContactDto) SetCoverUrlNil(b bool)`

 SetCoverUrlNil sets the value for CoverUrl to be an explicit nil

### UnsetCoverUrl
`func (o *ExtendedContactDto) UnsetCoverUrl()`

UnsetCoverUrl ensures that no value is present for CoverUrl, not even an explicit nil
### GetAvatarUrl

`func (o *ExtendedContactDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ExtendedContactDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ExtendedContactDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ExtendedContactDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *ExtendedContactDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *ExtendedContactDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetCountryId

`func (o *ExtendedContactDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *ExtendedContactDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *ExtendedContactDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *ExtendedContactDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *ExtendedContactDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *ExtendedContactDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetTimezoneId

`func (o *ExtendedContactDto) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *ExtendedContactDto) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *ExtendedContactDto) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *ExtendedContactDto) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### SetTimezoneIdNil

`func (o *ExtendedContactDto) SetTimezoneIdNil(b bool)`

 SetTimezoneIdNil sets the value for TimezoneId to be an explicit nil

### UnsetTimezoneId
`func (o *ExtendedContactDto) UnsetTimezoneId()`

UnsetTimezoneId ensures that no value is present for TimezoneId, not even an explicit nil
### GetLanguageId

`func (o *ExtendedContactDto) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *ExtendedContactDto) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *ExtendedContactDto) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *ExtendedContactDto) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### SetLanguageIdNil

`func (o *ExtendedContactDto) SetLanguageIdNil(b bool)`

 SetLanguageIdNil sets the value for LanguageId to be an explicit nil

### UnsetLanguageId
`func (o *ExtendedContactDto) UnsetLanguageId()`

UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
### GetSocialProfileId

`func (o *ExtendedContactDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *ExtendedContactDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *ExtendedContactDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *ExtendedContactDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *ExtendedContactDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *ExtendedContactDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetWebUrl

`func (o *ExtendedContactDto) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *ExtendedContactDto) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *ExtendedContactDto) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *ExtendedContactDto) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *ExtendedContactDto) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *ExtendedContactDto) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
### GetGitHubUrl

`func (o *ExtendedContactDto) GetGitHubUrl() string`

GetGitHubUrl returns the GitHubUrl field if non-nil, zero value otherwise.

### GetGitHubUrlOk

`func (o *ExtendedContactDto) GetGitHubUrlOk() (*string, bool)`

GetGitHubUrlOk returns a tuple with the GitHubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHubUrl

`func (o *ExtendedContactDto) SetGitHubUrl(v string)`

SetGitHubUrl sets GitHubUrl field to given value.

### HasGitHubUrl

`func (o *ExtendedContactDto) HasGitHubUrl() bool`

HasGitHubUrl returns a boolean if a field has been set.

### SetGitHubUrlNil

`func (o *ExtendedContactDto) SetGitHubUrlNil(b bool)`

 SetGitHubUrlNil sets the value for GitHubUrl to be an explicit nil

### UnsetGitHubUrl
`func (o *ExtendedContactDto) UnsetGitHubUrl()`

UnsetGitHubUrl ensures that no value is present for GitHubUrl, not even an explicit nil
### GetTwitchUrl

`func (o *ExtendedContactDto) GetTwitchUrl() string`

GetTwitchUrl returns the TwitchUrl field if non-nil, zero value otherwise.

### GetTwitchUrlOk

`func (o *ExtendedContactDto) GetTwitchUrlOk() (*string, bool)`

GetTwitchUrlOk returns a tuple with the TwitchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitchUrl

`func (o *ExtendedContactDto) SetTwitchUrl(v string)`

SetTwitchUrl sets TwitchUrl field to given value.

### HasTwitchUrl

`func (o *ExtendedContactDto) HasTwitchUrl() bool`

HasTwitchUrl returns a boolean if a field has been set.

### SetTwitchUrlNil

`func (o *ExtendedContactDto) SetTwitchUrlNil(b bool)`

 SetTwitchUrlNil sets the value for TwitchUrl to be an explicit nil

### UnsetTwitchUrl
`func (o *ExtendedContactDto) UnsetTwitchUrl()`

UnsetTwitchUrl ensures that no value is present for TwitchUrl, not even an explicit nil
### GetRedditUrl

`func (o *ExtendedContactDto) GetRedditUrl() string`

GetRedditUrl returns the RedditUrl field if non-nil, zero value otherwise.

### GetRedditUrlOk

`func (o *ExtendedContactDto) GetRedditUrlOk() (*string, bool)`

GetRedditUrlOk returns a tuple with the RedditUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedditUrl

`func (o *ExtendedContactDto) SetRedditUrl(v string)`

SetRedditUrl sets RedditUrl field to given value.

### HasRedditUrl

`func (o *ExtendedContactDto) HasRedditUrl() bool`

HasRedditUrl returns a boolean if a field has been set.

### SetRedditUrlNil

`func (o *ExtendedContactDto) SetRedditUrlNil(b bool)`

 SetRedditUrlNil sets the value for RedditUrl to be an explicit nil

### UnsetRedditUrl
`func (o *ExtendedContactDto) UnsetRedditUrl()`

UnsetRedditUrl ensures that no value is present for RedditUrl, not even an explicit nil
### GetTikTokUrl

`func (o *ExtendedContactDto) GetTikTokUrl() string`

GetTikTokUrl returns the TikTokUrl field if non-nil, zero value otherwise.

### GetTikTokUrlOk

`func (o *ExtendedContactDto) GetTikTokUrlOk() (*string, bool)`

GetTikTokUrlOk returns a tuple with the TikTokUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTikTokUrl

`func (o *ExtendedContactDto) SetTikTokUrl(v string)`

SetTikTokUrl sets TikTokUrl field to given value.

### HasTikTokUrl

`func (o *ExtendedContactDto) HasTikTokUrl() bool`

HasTikTokUrl returns a boolean if a field has been set.

### SetTikTokUrlNil

`func (o *ExtendedContactDto) SetTikTokUrlNil(b bool)`

 SetTikTokUrlNil sets the value for TikTokUrl to be an explicit nil

### UnsetTikTokUrl
`func (o *ExtendedContactDto) UnsetTikTokUrl()`

UnsetTikTokUrl ensures that no value is present for TikTokUrl, not even an explicit nil
### GetWebsiteUrl

`func (o *ExtendedContactDto) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ExtendedContactDto) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ExtendedContactDto) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ExtendedContactDto) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *ExtendedContactDto) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *ExtendedContactDto) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetTwitterUrl

`func (o *ExtendedContactDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *ExtendedContactDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *ExtendedContactDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *ExtendedContactDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *ExtendedContactDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *ExtendedContactDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetFacebookUrl

`func (o *ExtendedContactDto) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *ExtendedContactDto) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *ExtendedContactDto) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *ExtendedContactDto) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### SetFacebookUrlNil

`func (o *ExtendedContactDto) SetFacebookUrlNil(b bool)`

 SetFacebookUrlNil sets the value for FacebookUrl to be an explicit nil

### UnsetFacebookUrl
`func (o *ExtendedContactDto) UnsetFacebookUrl()`

UnsetFacebookUrl ensures that no value is present for FacebookUrl, not even an explicit nil
### GetYouTubeUrl

`func (o *ExtendedContactDto) GetYouTubeUrl() string`

GetYouTubeUrl returns the YouTubeUrl field if non-nil, zero value otherwise.

### GetYouTubeUrlOk

`func (o *ExtendedContactDto) GetYouTubeUrlOk() (*string, bool)`

GetYouTubeUrlOk returns a tuple with the YouTubeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYouTubeUrl

`func (o *ExtendedContactDto) SetYouTubeUrl(v string)`

SetYouTubeUrl sets YouTubeUrl field to given value.

### HasYouTubeUrl

`func (o *ExtendedContactDto) HasYouTubeUrl() bool`

HasYouTubeUrl returns a boolean if a field has been set.

### SetYouTubeUrlNil

`func (o *ExtendedContactDto) SetYouTubeUrlNil(b bool)`

 SetYouTubeUrlNil sets the value for YouTubeUrl to be an explicit nil

### UnsetYouTubeUrl
`func (o *ExtendedContactDto) UnsetYouTubeUrl()`

UnsetYouTubeUrl ensures that no value is present for YouTubeUrl, not even an explicit nil
### GetLinkedInUrl

`func (o *ExtendedContactDto) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *ExtendedContactDto) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *ExtendedContactDto) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *ExtendedContactDto) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### SetLinkedInUrlNil

`func (o *ExtendedContactDto) SetLinkedInUrlNil(b bool)`

 SetLinkedInUrlNil sets the value for LinkedInUrl to be an explicit nil

### UnsetLinkedInUrl
`func (o *ExtendedContactDto) UnsetLinkedInUrl()`

UnsetLinkedInUrl ensures that no value is present for LinkedInUrl, not even an explicit nil
### GetInstagramUrl

`func (o *ExtendedContactDto) GetInstagramUrl() string`

GetInstagramUrl returns the InstagramUrl field if non-nil, zero value otherwise.

### GetInstagramUrlOk

`func (o *ExtendedContactDto) GetInstagramUrlOk() (*string, bool)`

GetInstagramUrlOk returns a tuple with the InstagramUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUrl

`func (o *ExtendedContactDto) SetInstagramUrl(v string)`

SetInstagramUrl sets InstagramUrl field to given value.

### HasInstagramUrl

`func (o *ExtendedContactDto) HasInstagramUrl() bool`

HasInstagramUrl returns a boolean if a field has been set.

### SetInstagramUrlNil

`func (o *ExtendedContactDto) SetInstagramUrlNil(b bool)`

 SetInstagramUrlNil sets the value for InstagramUrl to be an explicit nil

### UnsetInstagramUrl
`func (o *ExtendedContactDto) UnsetInstagramUrl()`

UnsetInstagramUrl ensures that no value is present for InstagramUrl, not even an explicit nil
### GetGithubUsername

`func (o *ExtendedContactDto) GetGithubUsername() string`

GetGithubUsername returns the GithubUsername field if non-nil, zero value otherwise.

### GetGithubUsernameOk

`func (o *ExtendedContactDto) GetGithubUsernameOk() (*string, bool)`

GetGithubUsernameOk returns a tuple with the GithubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUsername

`func (o *ExtendedContactDto) SetGithubUsername(v string)`

SetGithubUsername sets GithubUsername field to given value.

### HasGithubUsername

`func (o *ExtendedContactDto) HasGithubUsername() bool`

HasGithubUsername returns a boolean if a field has been set.

### SetGithubUsernameNil

`func (o *ExtendedContactDto) SetGithubUsernameNil(b bool)`

 SetGithubUsernameNil sets the value for GithubUsername to be an explicit nil

### UnsetGithubUsername
`func (o *ExtendedContactDto) UnsetGithubUsername()`

UnsetGithubUsername ensures that no value is present for GithubUsername, not even an explicit nil
### GetDuns

`func (o *ExtendedContactDto) GetDuns() string`

GetDuns returns the Duns field if non-nil, zero value otherwise.

### GetDunsOk

`func (o *ExtendedContactDto) GetDunsOk() (*string, bool)`

GetDunsOk returns a tuple with the Duns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuns

`func (o *ExtendedContactDto) SetDuns(v string)`

SetDuns sets Duns field to given value.

### HasDuns

`func (o *ExtendedContactDto) HasDuns() bool`

HasDuns returns a boolean if a field has been set.

### SetDunsNil

`func (o *ExtendedContactDto) SetDunsNil(b bool)`

 SetDunsNil sets the value for Duns to be an explicit nil

### UnsetDuns
`func (o *ExtendedContactDto) UnsetDuns()`

UnsetDuns ensures that no value is present for Duns, not even an explicit nil
### GetTaxId

`func (o *ExtendedContactDto) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *ExtendedContactDto) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *ExtendedContactDto) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *ExtendedContactDto) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *ExtendedContactDto) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *ExtendedContactDto) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetEmail

`func (o *ExtendedContactDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ExtendedContactDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ExtendedContactDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ExtendedContactDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ExtendedContactDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ExtendedContactDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAbout

`func (o *ExtendedContactDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *ExtendedContactDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *ExtendedContactDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *ExtendedContactDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *ExtendedContactDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *ExtendedContactDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetStreet

`func (o *ExtendedContactDto) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ExtendedContactDto) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ExtendedContactDto) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ExtendedContactDto) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *ExtendedContactDto) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *ExtendedContactDto) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetCartId

`func (o *ExtendedContactDto) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *ExtendedContactDto) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *ExtendedContactDto) SetCartId(v string)`

SetCartId sets CartId field to given value.

### HasCartId

`func (o *ExtendedContactDto) HasCartId() bool`

HasCartId returns a boolean if a field has been set.

### SetCartIdNil

`func (o *ExtendedContactDto) SetCartIdNil(b bool)`

 SetCartIdNil sets the value for CartId to be an explicit nil

### UnsetCartId
`func (o *ExtendedContactDto) UnsetCartId()`

UnsetCartId ensures that no value is present for CartId, not even an explicit nil
### GetCityId

`func (o *ExtendedContactDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *ExtendedContactDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *ExtendedContactDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *ExtendedContactDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *ExtendedContactDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *ExtendedContactDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetZipCode

`func (o *ExtendedContactDto) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ExtendedContactDto) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ExtendedContactDto) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ExtendedContactDto) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *ExtendedContactDto) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *ExtendedContactDto) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetStateId

`func (o *ExtendedContactDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *ExtendedContactDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *ExtendedContactDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *ExtendedContactDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *ExtendedContactDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *ExtendedContactDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetWalletId

`func (o *ExtendedContactDto) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *ExtendedContactDto) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *ExtendedContactDto) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *ExtendedContactDto) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### SetWalletIdNil

`func (o *ExtendedContactDto) SetWalletIdNil(b bool)`

 SetWalletIdNil sets the value for WalletId to be an explicit nil

### UnsetWalletId
`func (o *ExtendedContactDto) UnsetWalletId()`

UnsetWalletId ensures that no value is present for WalletId, not even an explicit nil
### GetFaxNumber

`func (o *ExtendedContactDto) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *ExtendedContactDto) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *ExtendedContactDto) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *ExtendedContactDto) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### SetFaxNumberNil

`func (o *ExtendedContactDto) SetFaxNumberNil(b bool)`

 SetFaxNumberNil sets the value for FaxNumber to be an explicit nil

### UnsetFaxNumber
`func (o *ExtendedContactDto) UnsetFaxNumber()`

UnsetFaxNumber ensures that no value is present for FaxNumber, not even an explicit nil
### GetPostalCode

`func (o *ExtendedContactDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ExtendedContactDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ExtendedContactDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ExtendedContactDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ExtendedContactDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ExtendedContactDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCurrencyId

`func (o *ExtendedContactDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *ExtendedContactDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *ExtendedContactDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *ExtendedContactDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *ExtendedContactDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *ExtendedContactDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetStreetLine1

`func (o *ExtendedContactDto) GetStreetLine1() string`

GetStreetLine1 returns the StreetLine1 field if non-nil, zero value otherwise.

### GetStreetLine1Ok

`func (o *ExtendedContactDto) GetStreetLine1Ok() (*string, bool)`

GetStreetLine1Ok returns a tuple with the StreetLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine1

`func (o *ExtendedContactDto) SetStreetLine1(v string)`

SetStreetLine1 sets StreetLine1 field to given value.

### HasStreetLine1

`func (o *ExtendedContactDto) HasStreetLine1() bool`

HasStreetLine1 returns a boolean if a field has been set.

### SetStreetLine1Nil

`func (o *ExtendedContactDto) SetStreetLine1Nil(b bool)`

 SetStreetLine1Nil sets the value for StreetLine1 to be an explicit nil

### UnsetStreetLine1
`func (o *ExtendedContactDto) UnsetStreetLine1()`

UnsetStreetLine1 ensures that no value is present for StreetLine1, not even an explicit nil
### GetStreetLine2

`func (o *ExtendedContactDto) GetStreetLine2() string`

GetStreetLine2 returns the StreetLine2 field if non-nil, zero value otherwise.

### GetStreetLine2Ok

`func (o *ExtendedContactDto) GetStreetLine2Ok() (*string, bool)`

GetStreetLine2Ok returns a tuple with the StreetLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetLine2

`func (o *ExtendedContactDto) SetStreetLine2(v string)`

SetStreetLine2 sets StreetLine2 field to given value.

### HasStreetLine2

`func (o *ExtendedContactDto) HasStreetLine2() bool`

HasStreetLine2 returns a boolean if a field has been set.

### SetStreetLine2Nil

`func (o *ExtendedContactDto) SetStreetLine2Nil(b bool)`

 SetStreetLine2Nil sets the value for StreetLine2 to be an explicit nil

### UnsetStreetLine2
`func (o *ExtendedContactDto) UnsetStreetLine2()`

UnsetStreetLine2 ensures that no value is present for StreetLine2, not even an explicit nil
### GetTerritoryId

`func (o *ExtendedContactDto) GetTerritoryId() string`

GetTerritoryId returns the TerritoryId field if non-nil, zero value otherwise.

### GetTerritoryIdOk

`func (o *ExtendedContactDto) GetTerritoryIdOk() (*string, bool)`

GetTerritoryIdOk returns a tuple with the TerritoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritoryId

`func (o *ExtendedContactDto) SetTerritoryId(v string)`

SetTerritoryId sets TerritoryId field to given value.

### HasTerritoryId

`func (o *ExtendedContactDto) HasTerritoryId() bool`

HasTerritoryId returns a boolean if a field has been set.

### SetTerritoryIdNil

`func (o *ExtendedContactDto) SetTerritoryIdNil(b bool)`

 SetTerritoryIdNil sets the value for TerritoryId to be an explicit nil

### UnsetTerritoryId
`func (o *ExtendedContactDto) UnsetTerritoryId()`

UnsetTerritoryId ensures that no value is present for TerritoryId, not even an explicit nil
### GetMobilePhone

`func (o *ExtendedContactDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ExtendedContactDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ExtendedContactDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ExtendedContactDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *ExtendedContactDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *ExtendedContactDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetEnrollmentId

`func (o *ExtendedContactDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ExtendedContactDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ExtendedContactDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ExtendedContactDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ExtendedContactDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ExtendedContactDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAnnualRevenue

`func (o *ExtendedContactDto) GetAnnualRevenue() string`

GetAnnualRevenue returns the AnnualRevenue field if non-nil, zero value otherwise.

### GetAnnualRevenueOk

`func (o *ExtendedContactDto) GetAnnualRevenueOk() (*string, bool)`

GetAnnualRevenueOk returns a tuple with the AnnualRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualRevenue

`func (o *ExtendedContactDto) SetAnnualRevenue(v string)`

SetAnnualRevenue sets AnnualRevenue field to given value.

### HasAnnualRevenue

`func (o *ExtendedContactDto) HasAnnualRevenue() bool`

HasAnnualRevenue returns a boolean if a field has been set.

### SetAnnualRevenueNil

`func (o *ExtendedContactDto) SetAnnualRevenueNil(b bool)`

 SetAnnualRevenueNil sets the value for AnnualRevenue to be an explicit nil

### UnsetAnnualRevenue
`func (o *ExtendedContactDto) UnsetAnnualRevenue()`

UnsetAnnualRevenue ensures that no value is present for AnnualRevenue, not even an explicit nil
### GetRelatedUserId

`func (o *ExtendedContactDto) GetRelatedUserId() string`

GetRelatedUserId returns the RelatedUserId field if non-nil, zero value otherwise.

### GetRelatedUserIdOk

`func (o *ExtendedContactDto) GetRelatedUserIdOk() (*string, bool)`

GetRelatedUserIdOk returns a tuple with the RelatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedUserId

`func (o *ExtendedContactDto) SetRelatedUserId(v string)`

SetRelatedUserId sets RelatedUserId field to given value.

### HasRelatedUserId

`func (o *ExtendedContactDto) HasRelatedUserId() bool`

HasRelatedUserId returns a boolean if a field has been set.

### SetRelatedUserIdNil

`func (o *ExtendedContactDto) SetRelatedUserIdNil(b bool)`

 SetRelatedUserIdNil sets the value for RelatedUserId to be an explicit nil

### UnsetRelatedUserId
`func (o *ExtendedContactDto) UnsetRelatedUserId()`

UnsetRelatedUserId ensures that no value is present for RelatedUserId, not even an explicit nil
### GetBusinessPhone

`func (o *ExtendedContactDto) GetBusinessPhone() string`

GetBusinessPhone returns the BusinessPhone field if non-nil, zero value otherwise.

### GetBusinessPhoneOk

`func (o *ExtendedContactDto) GetBusinessPhoneOk() (*string, bool)`

GetBusinessPhoneOk returns a tuple with the BusinessPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhone

`func (o *ExtendedContactDto) SetBusinessPhone(v string)`

SetBusinessPhone sets BusinessPhone field to given value.

### HasBusinessPhone

`func (o *ExtendedContactDto) HasBusinessPhone() bool`

HasBusinessPhone returns a boolean if a field has been set.

### SetBusinessPhoneNil

`func (o *ExtendedContactDto) SetBusinessPhoneNil(b bool)`

 SetBusinessPhoneNil sets the value for BusinessPhone to be an explicit nil

### UnsetBusinessPhone
`func (o *ExtendedContactDto) UnsetBusinessPhone()`

UnsetBusinessPhone ensures that no value is present for BusinessPhone, not even an explicit nil
### GetOwnerContactId

`func (o *ExtendedContactDto) GetOwnerContactId() string`

GetOwnerContactId returns the OwnerContactId field if non-nil, zero value otherwise.

### GetOwnerContactIdOk

`func (o *ExtendedContactDto) GetOwnerContactIdOk() (*string, bool)`

GetOwnerContactIdOk returns a tuple with the OwnerContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContactId

`func (o *ExtendedContactDto) SetOwnerContactId(v string)`

SetOwnerContactId sets OwnerContactId field to given value.

### HasOwnerContactId

`func (o *ExtendedContactDto) HasOwnerContactId() bool`

HasOwnerContactId returns a boolean if a field has been set.

### SetOwnerContactIdNil

`func (o *ExtendedContactDto) SetOwnerContactIdNil(b bool)`

 SetOwnerContactIdNil sets the value for OwnerContactId to be an explicit nil

### UnsetOwnerContactId
`func (o *ExtendedContactDto) UnsetOwnerContactId()`

UnsetOwnerContactId ensures that no value is present for OwnerContactId, not even an explicit nil
### GetRelatedTenantId

`func (o *ExtendedContactDto) GetRelatedTenantId() string`

GetRelatedTenantId returns the RelatedTenantId field if non-nil, zero value otherwise.

### GetRelatedTenantIdOk

`func (o *ExtendedContactDto) GetRelatedTenantIdOk() (*string, bool)`

GetRelatedTenantIdOk returns a tuple with the RelatedTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTenantId

`func (o *ExtendedContactDto) SetRelatedTenantId(v string)`

SetRelatedTenantId sets RelatedTenantId field to given value.

### HasRelatedTenantId

`func (o *ExtendedContactDto) HasRelatedTenantId() bool`

HasRelatedTenantId returns a boolean if a field has been set.

### SetRelatedTenantIdNil

`func (o *ExtendedContactDto) SetRelatedTenantIdNil(b bool)`

 SetRelatedTenantIdNil sets the value for RelatedTenantId to be an explicit nil

### UnsetRelatedTenantId
`func (o *ExtendedContactDto) UnsetRelatedTenantId()`

UnsetRelatedTenantId ensures that no value is present for RelatedTenantId, not even an explicit nil
### GetActivityFeedId

`func (o *ExtendedContactDto) GetActivityFeedId() string`

GetActivityFeedId returns the ActivityFeedId field if non-nil, zero value otherwise.

### GetActivityFeedIdOk

`func (o *ExtendedContactDto) GetActivityFeedIdOk() (*string, bool)`

GetActivityFeedIdOk returns a tuple with the ActivityFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityFeedId

`func (o *ExtendedContactDto) SetActivityFeedId(v string)`

SetActivityFeedId sets ActivityFeedId field to given value.

### HasActivityFeedId

`func (o *ExtendedContactDto) HasActivityFeedId() bool`

HasActivityFeedId returns a boolean if a field has been set.

### SetActivityFeedIdNil

`func (o *ExtendedContactDto) SetActivityFeedIdNil(b bool)`

 SetActivityFeedIdNil sets the value for ActivityFeedId to be an explicit nil

### UnsetActivityFeedId
`func (o *ExtendedContactDto) UnsetActivityFeedId()`

UnsetActivityFeedId ensures that no value is present for ActivityFeedId, not even an explicit nil
### GetParentContactId

`func (o *ExtendedContactDto) GetParentContactId() string`

GetParentContactId returns the ParentContactId field if non-nil, zero value otherwise.

### GetParentContactIdOk

`func (o *ExtendedContactDto) GetParentContactIdOk() (*string, bool)`

GetParentContactIdOk returns a tuple with the ParentContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentContactId

`func (o *ExtendedContactDto) SetParentContactId(v string)`

SetParentContactId sets ParentContactId field to given value.

### HasParentContactId

`func (o *ExtendedContactDto) HasParentContactId() bool`

HasParentContactId returns a boolean if a field has been set.

### SetParentContactIdNil

`func (o *ExtendedContactDto) SetParentContactIdNil(b bool)`

 SetParentContactIdNil sets the value for ParentContactId to be an explicit nil

### UnsetParentContactId
`func (o *ExtendedContactDto) UnsetParentContactId()`

UnsetParentContactId ensures that no value is present for ParentContactId, not even an explicit nil
### GetIdentityProvider

`func (o *ExtendedContactDto) GetIdentityProvider() string`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *ExtendedContactDto) GetIdentityProviderOk() (*string, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *ExtendedContactDto) SetIdentityProvider(v string)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *ExtendedContactDto) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### SetIdentityProviderNil

`func (o *ExtendedContactDto) SetIdentityProviderNil(b bool)`

 SetIdentityProviderNil sets the value for IdentityProvider to be an explicit nil

### UnsetIdentityProvider
`func (o *ExtendedContactDto) UnsetIdentityProvider()`

UnsetIdentityProvider ensures that no value is present for IdentityProvider, not even an explicit nil
### GetPartnerProfileId

`func (o *ExtendedContactDto) GetPartnerProfileId() string`

GetPartnerProfileId returns the PartnerProfileId field if non-nil, zero value otherwise.

### GetPartnerProfileIdOk

`func (o *ExtendedContactDto) GetPartnerProfileIdOk() (*string, bool)`

GetPartnerProfileIdOk returns a tuple with the PartnerProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProfileId

`func (o *ExtendedContactDto) SetPartnerProfileId(v string)`

SetPartnerProfileId sets PartnerProfileId field to given value.

### HasPartnerProfileId

`func (o *ExtendedContactDto) HasPartnerProfileId() bool`

HasPartnerProfileId returns a boolean if a field has been set.

### SetPartnerProfileIdNil

`func (o *ExtendedContactDto) SetPartnerProfileIdNil(b bool)`

 SetPartnerProfileIdNil sets the value for PartnerProfileId to be an explicit nil

### UnsetPartnerProfileId
`func (o *ExtendedContactDto) UnsetPartnerProfileId()`

UnsetPartnerProfileId ensures that no value is present for PartnerProfileId, not even an explicit nil
### GetPrimaryContactId

`func (o *ExtendedContactDto) GetPrimaryContactId() string`

GetPrimaryContactId returns the PrimaryContactId field if non-nil, zero value otherwise.

### GetPrimaryContactIdOk

`func (o *ExtendedContactDto) GetPrimaryContactIdOk() (*string, bool)`

GetPrimaryContactIdOk returns a tuple with the PrimaryContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactId

`func (o *ExtendedContactDto) SetPrimaryContactId(v string)`

SetPrimaryContactId sets PrimaryContactId field to given value.

### HasPrimaryContactId

`func (o *ExtendedContactDto) HasPrimaryContactId() bool`

HasPrimaryContactId returns a boolean if a field has been set.

### SetPrimaryContactIdNil

`func (o *ExtendedContactDto) SetPrimaryContactIdNil(b bool)`

 SetPrimaryContactIdNil sets the value for PrimaryContactId to be an explicit nil

### UnsetPrimaryContactId
`func (o *ExtendedContactDto) UnsetPrimaryContactId()`

UnsetPrimaryContactId ensures that no value is present for PrimaryContactId, not even an explicit nil
### GetActiveDirectoryId

`func (o *ExtendedContactDto) GetActiveDirectoryId() string`

GetActiveDirectoryId returns the ActiveDirectoryId field if non-nil, zero value otherwise.

### GetActiveDirectoryIdOk

`func (o *ExtendedContactDto) GetActiveDirectoryIdOk() (*string, bool)`

GetActiveDirectoryIdOk returns a tuple with the ActiveDirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryId

`func (o *ExtendedContactDto) SetActiveDirectoryId(v string)`

SetActiveDirectoryId sets ActiveDirectoryId field to given value.

### HasActiveDirectoryId

`func (o *ExtendedContactDto) HasActiveDirectoryId() bool`

HasActiveDirectoryId returns a boolean if a field has been set.

### SetActiveDirectoryIdNil

`func (o *ExtendedContactDto) SetActiveDirectoryIdNil(b bool)`

 SetActiveDirectoryIdNil sets the value for ActiveDirectoryId to be an explicit nil

### UnsetActiveDirectoryId
`func (o *ExtendedContactDto) UnsetActiveDirectoryId()`

UnsetActiveDirectoryId ensures that no value is present for ActiveDirectoryId, not even an explicit nil
### GetIdentityProviderAccessToken

`func (o *ExtendedContactDto) GetIdentityProviderAccessToken() string`

GetIdentityProviderAccessToken returns the IdentityProviderAccessToken field if non-nil, zero value otherwise.

### GetIdentityProviderAccessTokenOk

`func (o *ExtendedContactDto) GetIdentityProviderAccessTokenOk() (*string, bool)`

GetIdentityProviderAccessTokenOk returns a tuple with the IdentityProviderAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderAccessToken

`func (o *ExtendedContactDto) SetIdentityProviderAccessToken(v string)`

SetIdentityProviderAccessToken sets IdentityProviderAccessToken field to given value.

### HasIdentityProviderAccessToken

`func (o *ExtendedContactDto) HasIdentityProviderAccessToken() bool`

HasIdentityProviderAccessToken returns a boolean if a field has been set.

### SetIdentityProviderAccessTokenNil

`func (o *ExtendedContactDto) SetIdentityProviderAccessTokenNil(b bool)`

 SetIdentityProviderAccessTokenNil sets the value for IdentityProviderAccessToken to be an explicit nil

### UnsetIdentityProviderAccessToken
`func (o *ExtendedContactDto) UnsetIdentityProviderAccessToken()`

UnsetIdentityProviderAccessToken ensures that no value is present for IdentityProviderAccessToken, not even an explicit nil
### GetBirthday

`func (o *ExtendedContactDto) GetBirthday() time.Time`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *ExtendedContactDto) GetBirthdayOk() (*time.Time, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *ExtendedContactDto) SetBirthday(v time.Time)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *ExtendedContactDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *ExtendedContactDto) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *ExtendedContactDto) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetCart

`func (o *ExtendedContactDto) GetCart() CartDto`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *ExtendedContactDto) GetCartOk() (*CartDto, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *ExtendedContactDto) SetCart(v CartDto)`

SetCart sets Cart field to given value.

### HasCart

`func (o *ExtendedContactDto) HasCart() bool`

HasCart returns a boolean if a field has been set.

### GetWallet

`func (o *ExtendedContactDto) GetWallet() WalletDto`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *ExtendedContactDto) GetWalletOk() (*WalletDto, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *ExtendedContactDto) SetWallet(v WalletDto)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *ExtendedContactDto) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetSocialProfile

`func (o *ExtendedContactDto) GetSocialProfile() SocialProfileDto`

GetSocialProfile returns the SocialProfile field if non-nil, zero value otherwise.

### GetSocialProfileOk

`func (o *ExtendedContactDto) GetSocialProfileOk() (*SocialProfileDto, bool)`

GetSocialProfileOk returns a tuple with the SocialProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfile

`func (o *ExtendedContactDto) SetSocialProfile(v SocialProfileDto)`

SetSocialProfile sets SocialProfile field to given value.

### HasSocialProfile

`func (o *ExtendedContactDto) HasSocialProfile() bool`

HasSocialProfile returns a boolean if a field has been set.

### GetParentContact

`func (o *ExtendedContactDto) GetParentContact() SimpleContactDto`

GetParentContact returns the ParentContact field if non-nil, zero value otherwise.

### GetParentContactOk

`func (o *ExtendedContactDto) GetParentContactOk() (*SimpleContactDto, bool)`

GetParentContactOk returns a tuple with the ParentContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentContact

`func (o *ExtendedContactDto) SetParentContact(v SimpleContactDto)`

SetParentContact sets ParentContact field to given value.

### HasParentContact

`func (o *ExtendedContactDto) HasParentContact() bool`

HasParentContact returns a boolean if a field has been set.

### GetPrimaryContact

`func (o *ExtendedContactDto) GetPrimaryContact() SimpleContactDto`

GetPrimaryContact returns the PrimaryContact field if non-nil, zero value otherwise.

### GetPrimaryContactOk

`func (o *ExtendedContactDto) GetPrimaryContactOk() (*SimpleContactDto, bool)`

GetPrimaryContactOk returns a tuple with the PrimaryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContact

`func (o *ExtendedContactDto) SetPrimaryContact(v SimpleContactDto)`

SetPrimaryContact sets PrimaryContact field to given value.

### HasPrimaryContact

`func (o *ExtendedContactDto) HasPrimaryContact() bool`

HasPrimaryContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


