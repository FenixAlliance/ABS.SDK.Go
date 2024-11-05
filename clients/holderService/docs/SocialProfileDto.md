# SocialProfileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**Cover** | Pointer to **NullableString** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**CountryName** | Pointer to **NullableString** |  | [optional] 
**IdentityId** | Pointer to **NullableString** |  | [optional] 
**FollowsCount** | Pointer to **NullableInt32** |  | [optional] 
**MessagesCount** | Pointer to **NullableInt32** |  | [optional] 
**FollowersCount** | Pointer to **NullableInt32** |  | [optional] 
**NotificationsCount** | Pointer to **NullableInt32** |  | [optional] 
**UnreadNotificationsCount** | Pointer to **NullableInt32** |  | [optional] 
**UnreadMessagesCount** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **NullableInt32** |  | [optional] 
**SocialFeedId** | Pointer to **NullableString** |  | [optional] 
**TwitterUrl** | Pointer to **NullableString** |  | [optional] 
**FacebookURL** | Pointer to **NullableString** |  | [optional] 
**LinkedInURL** | Pointer to **NullableString** |  | [optional] 
**YoutubeURL** | Pointer to **NullableString** |  | [optional] 
**GithubURL** | Pointer to **NullableString** |  | [optional] 
**PinterestURL** | Pointer to **NullableString** |  | [optional] 
**DribbleURL** | Pointer to **NullableString** |  | [optional] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialProfileDto

`func NewSocialProfileDto() *SocialProfileDto`

NewSocialProfileDto instantiates a new SocialProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialProfileDtoWithDefaults

`func NewSocialProfileDtoWithDefaults() *SocialProfileDto`

NewSocialProfileDtoWithDefaults instantiates a new SocialProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialProfileDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialProfileDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialProfileDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialProfileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialProfileDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialProfileDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialProfileDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialProfileDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialProfileDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialProfileDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialProfileDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialProfileDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *SocialProfileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SocialProfileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SocialProfileDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SocialProfileDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SocialProfileDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SocialProfileDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAbout

`func (o *SocialProfileDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *SocialProfileDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *SocialProfileDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *SocialProfileDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *SocialProfileDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *SocialProfileDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetCover

`func (o *SocialProfileDto) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *SocialProfileDto) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *SocialProfileDto) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *SocialProfileDto) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *SocialProfileDto) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *SocialProfileDto) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil
### GetAvatar

`func (o *SocialProfileDto) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *SocialProfileDto) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *SocialProfileDto) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *SocialProfileDto) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *SocialProfileDto) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *SocialProfileDto) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetCountryId

`func (o *SocialProfileDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *SocialProfileDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *SocialProfileDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *SocialProfileDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *SocialProfileDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *SocialProfileDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetCountryName

`func (o *SocialProfileDto) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *SocialProfileDto) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *SocialProfileDto) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *SocialProfileDto) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### SetCountryNameNil

`func (o *SocialProfileDto) SetCountryNameNil(b bool)`

 SetCountryNameNil sets the value for CountryName to be an explicit nil

### UnsetCountryName
`func (o *SocialProfileDto) UnsetCountryName()`

UnsetCountryName ensures that no value is present for CountryName, not even an explicit nil
### GetIdentityId

`func (o *SocialProfileDto) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *SocialProfileDto) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *SocialProfileDto) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *SocialProfileDto) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### SetIdentityIdNil

`func (o *SocialProfileDto) SetIdentityIdNil(b bool)`

 SetIdentityIdNil sets the value for IdentityId to be an explicit nil

### UnsetIdentityId
`func (o *SocialProfileDto) UnsetIdentityId()`

UnsetIdentityId ensures that no value is present for IdentityId, not even an explicit nil
### GetFollowsCount

`func (o *SocialProfileDto) GetFollowsCount() int32`

GetFollowsCount returns the FollowsCount field if non-nil, zero value otherwise.

### GetFollowsCountOk

`func (o *SocialProfileDto) GetFollowsCountOk() (*int32, bool)`

GetFollowsCountOk returns a tuple with the FollowsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowsCount

`func (o *SocialProfileDto) SetFollowsCount(v int32)`

SetFollowsCount sets FollowsCount field to given value.

### HasFollowsCount

`func (o *SocialProfileDto) HasFollowsCount() bool`

HasFollowsCount returns a boolean if a field has been set.

### SetFollowsCountNil

`func (o *SocialProfileDto) SetFollowsCountNil(b bool)`

 SetFollowsCountNil sets the value for FollowsCount to be an explicit nil

### UnsetFollowsCount
`func (o *SocialProfileDto) UnsetFollowsCount()`

UnsetFollowsCount ensures that no value is present for FollowsCount, not even an explicit nil
### GetMessagesCount

`func (o *SocialProfileDto) GetMessagesCount() int32`

GetMessagesCount returns the MessagesCount field if non-nil, zero value otherwise.

### GetMessagesCountOk

`func (o *SocialProfileDto) GetMessagesCountOk() (*int32, bool)`

GetMessagesCountOk returns a tuple with the MessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesCount

`func (o *SocialProfileDto) SetMessagesCount(v int32)`

SetMessagesCount sets MessagesCount field to given value.

### HasMessagesCount

`func (o *SocialProfileDto) HasMessagesCount() bool`

HasMessagesCount returns a boolean if a field has been set.

### SetMessagesCountNil

`func (o *SocialProfileDto) SetMessagesCountNil(b bool)`

 SetMessagesCountNil sets the value for MessagesCount to be an explicit nil

### UnsetMessagesCount
`func (o *SocialProfileDto) UnsetMessagesCount()`

UnsetMessagesCount ensures that no value is present for MessagesCount, not even an explicit nil
### GetFollowersCount

`func (o *SocialProfileDto) GetFollowersCount() int32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *SocialProfileDto) GetFollowersCountOk() (*int32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *SocialProfileDto) SetFollowersCount(v int32)`

SetFollowersCount sets FollowersCount field to given value.

### HasFollowersCount

`func (o *SocialProfileDto) HasFollowersCount() bool`

HasFollowersCount returns a boolean if a field has been set.

### SetFollowersCountNil

`func (o *SocialProfileDto) SetFollowersCountNil(b bool)`

 SetFollowersCountNil sets the value for FollowersCount to be an explicit nil

### UnsetFollowersCount
`func (o *SocialProfileDto) UnsetFollowersCount()`

UnsetFollowersCount ensures that no value is present for FollowersCount, not even an explicit nil
### GetNotificationsCount

`func (o *SocialProfileDto) GetNotificationsCount() int32`

GetNotificationsCount returns the NotificationsCount field if non-nil, zero value otherwise.

### GetNotificationsCountOk

`func (o *SocialProfileDto) GetNotificationsCountOk() (*int32, bool)`

GetNotificationsCountOk returns a tuple with the NotificationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsCount

`func (o *SocialProfileDto) SetNotificationsCount(v int32)`

SetNotificationsCount sets NotificationsCount field to given value.

### HasNotificationsCount

`func (o *SocialProfileDto) HasNotificationsCount() bool`

HasNotificationsCount returns a boolean if a field has been set.

### SetNotificationsCountNil

`func (o *SocialProfileDto) SetNotificationsCountNil(b bool)`

 SetNotificationsCountNil sets the value for NotificationsCount to be an explicit nil

### UnsetNotificationsCount
`func (o *SocialProfileDto) UnsetNotificationsCount()`

UnsetNotificationsCount ensures that no value is present for NotificationsCount, not even an explicit nil
### GetUnreadNotificationsCount

`func (o *SocialProfileDto) GetUnreadNotificationsCount() int32`

GetUnreadNotificationsCount returns the UnreadNotificationsCount field if non-nil, zero value otherwise.

### GetUnreadNotificationsCountOk

`func (o *SocialProfileDto) GetUnreadNotificationsCountOk() (*int32, bool)`

GetUnreadNotificationsCountOk returns a tuple with the UnreadNotificationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadNotificationsCount

`func (o *SocialProfileDto) SetUnreadNotificationsCount(v int32)`

SetUnreadNotificationsCount sets UnreadNotificationsCount field to given value.

### HasUnreadNotificationsCount

`func (o *SocialProfileDto) HasUnreadNotificationsCount() bool`

HasUnreadNotificationsCount returns a boolean if a field has been set.

### SetUnreadNotificationsCountNil

`func (o *SocialProfileDto) SetUnreadNotificationsCountNil(b bool)`

 SetUnreadNotificationsCountNil sets the value for UnreadNotificationsCount to be an explicit nil

### UnsetUnreadNotificationsCount
`func (o *SocialProfileDto) UnsetUnreadNotificationsCount()`

UnsetUnreadNotificationsCount ensures that no value is present for UnreadNotificationsCount, not even an explicit nil
### GetUnreadMessagesCount

`func (o *SocialProfileDto) GetUnreadMessagesCount() int32`

GetUnreadMessagesCount returns the UnreadMessagesCount field if non-nil, zero value otherwise.

### GetUnreadMessagesCountOk

`func (o *SocialProfileDto) GetUnreadMessagesCountOk() (*int32, bool)`

GetUnreadMessagesCountOk returns a tuple with the UnreadMessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMessagesCount

`func (o *SocialProfileDto) SetUnreadMessagesCount(v int32)`

SetUnreadMessagesCount sets UnreadMessagesCount field to given value.

### HasUnreadMessagesCount

`func (o *SocialProfileDto) HasUnreadMessagesCount() bool`

HasUnreadMessagesCount returns a boolean if a field has been set.

### SetUnreadMessagesCountNil

`func (o *SocialProfileDto) SetUnreadMessagesCountNil(b bool)`

 SetUnreadMessagesCountNil sets the value for UnreadMessagesCount to be an explicit nil

### UnsetUnreadMessagesCount
`func (o *SocialProfileDto) UnsetUnreadMessagesCount()`

UnsetUnreadMessagesCount ensures that no value is present for UnreadMessagesCount, not even an explicit nil
### GetType

`func (o *SocialProfileDto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SocialProfileDto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SocialProfileDto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SocialProfileDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SocialProfileDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SocialProfileDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSocialFeedId

`func (o *SocialProfileDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *SocialProfileDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *SocialProfileDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *SocialProfileDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *SocialProfileDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *SocialProfileDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil
### GetTwitterUrl

`func (o *SocialProfileDto) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *SocialProfileDto) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *SocialProfileDto) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *SocialProfileDto) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### SetTwitterUrlNil

`func (o *SocialProfileDto) SetTwitterUrlNil(b bool)`

 SetTwitterUrlNil sets the value for TwitterUrl to be an explicit nil

### UnsetTwitterUrl
`func (o *SocialProfileDto) UnsetTwitterUrl()`

UnsetTwitterUrl ensures that no value is present for TwitterUrl, not even an explicit nil
### GetFacebookURL

`func (o *SocialProfileDto) GetFacebookURL() string`

GetFacebookURL returns the FacebookURL field if non-nil, zero value otherwise.

### GetFacebookURLOk

`func (o *SocialProfileDto) GetFacebookURLOk() (*string, bool)`

GetFacebookURLOk returns a tuple with the FacebookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookURL

`func (o *SocialProfileDto) SetFacebookURL(v string)`

SetFacebookURL sets FacebookURL field to given value.

### HasFacebookURL

`func (o *SocialProfileDto) HasFacebookURL() bool`

HasFacebookURL returns a boolean if a field has been set.

### SetFacebookURLNil

`func (o *SocialProfileDto) SetFacebookURLNil(b bool)`

 SetFacebookURLNil sets the value for FacebookURL to be an explicit nil

### UnsetFacebookURL
`func (o *SocialProfileDto) UnsetFacebookURL()`

UnsetFacebookURL ensures that no value is present for FacebookURL, not even an explicit nil
### GetLinkedInURL

`func (o *SocialProfileDto) GetLinkedInURL() string`

GetLinkedInURL returns the LinkedInURL field if non-nil, zero value otherwise.

### GetLinkedInURLOk

`func (o *SocialProfileDto) GetLinkedInURLOk() (*string, bool)`

GetLinkedInURLOk returns a tuple with the LinkedInURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInURL

`func (o *SocialProfileDto) SetLinkedInURL(v string)`

SetLinkedInURL sets LinkedInURL field to given value.

### HasLinkedInURL

`func (o *SocialProfileDto) HasLinkedInURL() bool`

HasLinkedInURL returns a boolean if a field has been set.

### SetLinkedInURLNil

`func (o *SocialProfileDto) SetLinkedInURLNil(b bool)`

 SetLinkedInURLNil sets the value for LinkedInURL to be an explicit nil

### UnsetLinkedInURL
`func (o *SocialProfileDto) UnsetLinkedInURL()`

UnsetLinkedInURL ensures that no value is present for LinkedInURL, not even an explicit nil
### GetYoutubeURL

`func (o *SocialProfileDto) GetYoutubeURL() string`

GetYoutubeURL returns the YoutubeURL field if non-nil, zero value otherwise.

### GetYoutubeURLOk

`func (o *SocialProfileDto) GetYoutubeURLOk() (*string, bool)`

GetYoutubeURLOk returns a tuple with the YoutubeURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutubeURL

`func (o *SocialProfileDto) SetYoutubeURL(v string)`

SetYoutubeURL sets YoutubeURL field to given value.

### HasYoutubeURL

`func (o *SocialProfileDto) HasYoutubeURL() bool`

HasYoutubeURL returns a boolean if a field has been set.

### SetYoutubeURLNil

`func (o *SocialProfileDto) SetYoutubeURLNil(b bool)`

 SetYoutubeURLNil sets the value for YoutubeURL to be an explicit nil

### UnsetYoutubeURL
`func (o *SocialProfileDto) UnsetYoutubeURL()`

UnsetYoutubeURL ensures that no value is present for YoutubeURL, not even an explicit nil
### GetGithubURL

`func (o *SocialProfileDto) GetGithubURL() string`

GetGithubURL returns the GithubURL field if non-nil, zero value otherwise.

### GetGithubURLOk

`func (o *SocialProfileDto) GetGithubURLOk() (*string, bool)`

GetGithubURLOk returns a tuple with the GithubURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubURL

`func (o *SocialProfileDto) SetGithubURL(v string)`

SetGithubURL sets GithubURL field to given value.

### HasGithubURL

`func (o *SocialProfileDto) HasGithubURL() bool`

HasGithubURL returns a boolean if a field has been set.

### SetGithubURLNil

`func (o *SocialProfileDto) SetGithubURLNil(b bool)`

 SetGithubURLNil sets the value for GithubURL to be an explicit nil

### UnsetGithubURL
`func (o *SocialProfileDto) UnsetGithubURL()`

UnsetGithubURL ensures that no value is present for GithubURL, not even an explicit nil
### GetPinterestURL

`func (o *SocialProfileDto) GetPinterestURL() string`

GetPinterestURL returns the PinterestURL field if non-nil, zero value otherwise.

### GetPinterestURLOk

`func (o *SocialProfileDto) GetPinterestURLOk() (*string, bool)`

GetPinterestURLOk returns a tuple with the PinterestURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinterestURL

`func (o *SocialProfileDto) SetPinterestURL(v string)`

SetPinterestURL sets PinterestURL field to given value.

### HasPinterestURL

`func (o *SocialProfileDto) HasPinterestURL() bool`

HasPinterestURL returns a boolean if a field has been set.

### SetPinterestURLNil

`func (o *SocialProfileDto) SetPinterestURLNil(b bool)`

 SetPinterestURLNil sets the value for PinterestURL to be an explicit nil

### UnsetPinterestURL
`func (o *SocialProfileDto) UnsetPinterestURL()`

UnsetPinterestURL ensures that no value is present for PinterestURL, not even an explicit nil
### GetDribbleURL

`func (o *SocialProfileDto) GetDribbleURL() string`

GetDribbleURL returns the DribbleURL field if non-nil, zero value otherwise.

### GetDribbleURLOk

`func (o *SocialProfileDto) GetDribbleURLOk() (*string, bool)`

GetDribbleURLOk returns a tuple with the DribbleURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDribbleURL

`func (o *SocialProfileDto) SetDribbleURL(v string)`

SetDribbleURL sets DribbleURL field to given value.

### HasDribbleURL

`func (o *SocialProfileDto) HasDribbleURL() bool`

HasDribbleURL returns a boolean if a field has been set.

### SetDribbleURLNil

`func (o *SocialProfileDto) SetDribbleURLNil(b bool)`

 SetDribbleURLNil sets the value for DribbleURL to be an explicit nil

### UnsetDribbleURL
`func (o *SocialProfileDto) UnsetDribbleURL()`

UnsetDribbleURL ensures that no value is present for DribbleURL, not even an explicit nil
### GetDomain

`func (o *SocialProfileDto) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SocialProfileDto) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SocialProfileDto) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SocialProfileDto) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *SocialProfileDto) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *SocialProfileDto) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetNotes

`func (o *SocialProfileDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SocialProfileDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SocialProfileDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SocialProfileDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *SocialProfileDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *SocialProfileDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


