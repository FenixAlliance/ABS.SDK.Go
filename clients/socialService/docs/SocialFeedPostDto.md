# SocialFeedPostDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 
**SocialProfileName** | Pointer to **NullableString** |  | [optional] 
**SocialProfileAvatarUrl** | Pointer to **NullableString** |  | [optional] 
**CommentsCount** | Pointer to **int32** |  | [optional] 
**ReactionsCount** | Pointer to **int32** |  | [optional] 
**SocialProfileType** | Pointer to **NullableString** |  | [optional] 
**BodyHtml** | Pointer to **NullableString** |  | [optional] 
**BodyFormat** | Pointer to **NullableString** |  | [optional] 
**BackgroundStyle** | Pointer to **NullableString** |  | [optional] 
**SocialFeedId** | Pointer to **NullableString** |  | [optional] 
**Facepile** | Pointer to [**[]SocialPostReactionFacepileDto**](SocialPostReactionFacepileDto.md) |  | [optional] 
**Attachments** | Pointer to [**[]SocialPostAttachmentRefDto**](SocialPostAttachmentRefDto.md) |  | [optional] 
**MyReaction** | Pointer to **NullableString** |  | [optional] 
**MyReactionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSocialFeedPostDto

`func NewSocialFeedPostDto() *SocialFeedPostDto`

NewSocialFeedPostDto instantiates a new SocialFeedPostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialFeedPostDtoWithDefaults

`func NewSocialFeedPostDtoWithDefaults() *SocialFeedPostDto`

NewSocialFeedPostDtoWithDefaults instantiates a new SocialFeedPostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialFeedPostDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialFeedPostDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialFeedPostDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialFeedPostDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SocialFeedPostDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SocialFeedPostDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SocialFeedPostDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SocialFeedPostDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SocialFeedPostDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SocialFeedPostDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SocialFeedPostDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SocialFeedPostDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *SocialFeedPostDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SocialFeedPostDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SocialFeedPostDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SocialFeedPostDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SocialFeedPostDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SocialFeedPostDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetMessage

`func (o *SocialFeedPostDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SocialFeedPostDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SocialFeedPostDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SocialFeedPostDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SocialFeedPostDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SocialFeedPostDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSocialProfileId

`func (o *SocialFeedPostDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SocialFeedPostDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SocialFeedPostDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SocialFeedPostDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SocialFeedPostDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SocialFeedPostDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil
### GetSocialProfileName

`func (o *SocialFeedPostDto) GetSocialProfileName() string`

GetSocialProfileName returns the SocialProfileName field if non-nil, zero value otherwise.

### GetSocialProfileNameOk

`func (o *SocialFeedPostDto) GetSocialProfileNameOk() (*string, bool)`

GetSocialProfileNameOk returns a tuple with the SocialProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileName

`func (o *SocialFeedPostDto) SetSocialProfileName(v string)`

SetSocialProfileName sets SocialProfileName field to given value.

### HasSocialProfileName

`func (o *SocialFeedPostDto) HasSocialProfileName() bool`

HasSocialProfileName returns a boolean if a field has been set.

### SetSocialProfileNameNil

`func (o *SocialFeedPostDto) SetSocialProfileNameNil(b bool)`

 SetSocialProfileNameNil sets the value for SocialProfileName to be an explicit nil

### UnsetSocialProfileName
`func (o *SocialFeedPostDto) UnsetSocialProfileName()`

UnsetSocialProfileName ensures that no value is present for SocialProfileName, not even an explicit nil
### GetSocialProfileAvatarUrl

`func (o *SocialFeedPostDto) GetSocialProfileAvatarUrl() string`

GetSocialProfileAvatarUrl returns the SocialProfileAvatarUrl field if non-nil, zero value otherwise.

### GetSocialProfileAvatarUrlOk

`func (o *SocialFeedPostDto) GetSocialProfileAvatarUrlOk() (*string, bool)`

GetSocialProfileAvatarUrlOk returns a tuple with the SocialProfileAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileAvatarUrl

`func (o *SocialFeedPostDto) SetSocialProfileAvatarUrl(v string)`

SetSocialProfileAvatarUrl sets SocialProfileAvatarUrl field to given value.

### HasSocialProfileAvatarUrl

`func (o *SocialFeedPostDto) HasSocialProfileAvatarUrl() bool`

HasSocialProfileAvatarUrl returns a boolean if a field has been set.

### SetSocialProfileAvatarUrlNil

`func (o *SocialFeedPostDto) SetSocialProfileAvatarUrlNil(b bool)`

 SetSocialProfileAvatarUrlNil sets the value for SocialProfileAvatarUrl to be an explicit nil

### UnsetSocialProfileAvatarUrl
`func (o *SocialFeedPostDto) UnsetSocialProfileAvatarUrl()`

UnsetSocialProfileAvatarUrl ensures that no value is present for SocialProfileAvatarUrl, not even an explicit nil
### GetCommentsCount

`func (o *SocialFeedPostDto) GetCommentsCount() int32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *SocialFeedPostDto) GetCommentsCountOk() (*int32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *SocialFeedPostDto) SetCommentsCount(v int32)`

SetCommentsCount sets CommentsCount field to given value.

### HasCommentsCount

`func (o *SocialFeedPostDto) HasCommentsCount() bool`

HasCommentsCount returns a boolean if a field has been set.

### GetReactionsCount

`func (o *SocialFeedPostDto) GetReactionsCount() int32`

GetReactionsCount returns the ReactionsCount field if non-nil, zero value otherwise.

### GetReactionsCountOk

`func (o *SocialFeedPostDto) GetReactionsCountOk() (*int32, bool)`

GetReactionsCountOk returns a tuple with the ReactionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionsCount

`func (o *SocialFeedPostDto) SetReactionsCount(v int32)`

SetReactionsCount sets ReactionsCount field to given value.

### HasReactionsCount

`func (o *SocialFeedPostDto) HasReactionsCount() bool`

HasReactionsCount returns a boolean if a field has been set.

### GetSocialProfileType

`func (o *SocialFeedPostDto) GetSocialProfileType() string`

GetSocialProfileType returns the SocialProfileType field if non-nil, zero value otherwise.

### GetSocialProfileTypeOk

`func (o *SocialFeedPostDto) GetSocialProfileTypeOk() (*string, bool)`

GetSocialProfileTypeOk returns a tuple with the SocialProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileType

`func (o *SocialFeedPostDto) SetSocialProfileType(v string)`

SetSocialProfileType sets SocialProfileType field to given value.

### HasSocialProfileType

`func (o *SocialFeedPostDto) HasSocialProfileType() bool`

HasSocialProfileType returns a boolean if a field has been set.

### SetSocialProfileTypeNil

`func (o *SocialFeedPostDto) SetSocialProfileTypeNil(b bool)`

 SetSocialProfileTypeNil sets the value for SocialProfileType to be an explicit nil

### UnsetSocialProfileType
`func (o *SocialFeedPostDto) UnsetSocialProfileType()`

UnsetSocialProfileType ensures that no value is present for SocialProfileType, not even an explicit nil
### GetBodyHtml

`func (o *SocialFeedPostDto) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *SocialFeedPostDto) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *SocialFeedPostDto) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *SocialFeedPostDto) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### SetBodyHtmlNil

`func (o *SocialFeedPostDto) SetBodyHtmlNil(b bool)`

 SetBodyHtmlNil sets the value for BodyHtml to be an explicit nil

### UnsetBodyHtml
`func (o *SocialFeedPostDto) UnsetBodyHtml()`

UnsetBodyHtml ensures that no value is present for BodyHtml, not even an explicit nil
### GetBodyFormat

`func (o *SocialFeedPostDto) GetBodyFormat() string`

GetBodyFormat returns the BodyFormat field if non-nil, zero value otherwise.

### GetBodyFormatOk

`func (o *SocialFeedPostDto) GetBodyFormatOk() (*string, bool)`

GetBodyFormatOk returns a tuple with the BodyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyFormat

`func (o *SocialFeedPostDto) SetBodyFormat(v string)`

SetBodyFormat sets BodyFormat field to given value.

### HasBodyFormat

`func (o *SocialFeedPostDto) HasBodyFormat() bool`

HasBodyFormat returns a boolean if a field has been set.

### SetBodyFormatNil

`func (o *SocialFeedPostDto) SetBodyFormatNil(b bool)`

 SetBodyFormatNil sets the value for BodyFormat to be an explicit nil

### UnsetBodyFormat
`func (o *SocialFeedPostDto) UnsetBodyFormat()`

UnsetBodyFormat ensures that no value is present for BodyFormat, not even an explicit nil
### GetBackgroundStyle

`func (o *SocialFeedPostDto) GetBackgroundStyle() string`

GetBackgroundStyle returns the BackgroundStyle field if non-nil, zero value otherwise.

### GetBackgroundStyleOk

`func (o *SocialFeedPostDto) GetBackgroundStyleOk() (*string, bool)`

GetBackgroundStyleOk returns a tuple with the BackgroundStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundStyle

`func (o *SocialFeedPostDto) SetBackgroundStyle(v string)`

SetBackgroundStyle sets BackgroundStyle field to given value.

### HasBackgroundStyle

`func (o *SocialFeedPostDto) HasBackgroundStyle() bool`

HasBackgroundStyle returns a boolean if a field has been set.

### SetBackgroundStyleNil

`func (o *SocialFeedPostDto) SetBackgroundStyleNil(b bool)`

 SetBackgroundStyleNil sets the value for BackgroundStyle to be an explicit nil

### UnsetBackgroundStyle
`func (o *SocialFeedPostDto) UnsetBackgroundStyle()`

UnsetBackgroundStyle ensures that no value is present for BackgroundStyle, not even an explicit nil
### GetSocialFeedId

`func (o *SocialFeedPostDto) GetSocialFeedId() string`

GetSocialFeedId returns the SocialFeedId field if non-nil, zero value otherwise.

### GetSocialFeedIdOk

`func (o *SocialFeedPostDto) GetSocialFeedIdOk() (*string, bool)`

GetSocialFeedIdOk returns a tuple with the SocialFeedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialFeedId

`func (o *SocialFeedPostDto) SetSocialFeedId(v string)`

SetSocialFeedId sets SocialFeedId field to given value.

### HasSocialFeedId

`func (o *SocialFeedPostDto) HasSocialFeedId() bool`

HasSocialFeedId returns a boolean if a field has been set.

### SetSocialFeedIdNil

`func (o *SocialFeedPostDto) SetSocialFeedIdNil(b bool)`

 SetSocialFeedIdNil sets the value for SocialFeedId to be an explicit nil

### UnsetSocialFeedId
`func (o *SocialFeedPostDto) UnsetSocialFeedId()`

UnsetSocialFeedId ensures that no value is present for SocialFeedId, not even an explicit nil
### GetFacepile

`func (o *SocialFeedPostDto) GetFacepile() []SocialPostReactionFacepileDto`

GetFacepile returns the Facepile field if non-nil, zero value otherwise.

### GetFacepileOk

`func (o *SocialFeedPostDto) GetFacepileOk() (*[]SocialPostReactionFacepileDto, bool)`

GetFacepileOk returns a tuple with the Facepile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacepile

`func (o *SocialFeedPostDto) SetFacepile(v []SocialPostReactionFacepileDto)`

SetFacepile sets Facepile field to given value.

### HasFacepile

`func (o *SocialFeedPostDto) HasFacepile() bool`

HasFacepile returns a boolean if a field has been set.

### SetFacepileNil

`func (o *SocialFeedPostDto) SetFacepileNil(b bool)`

 SetFacepileNil sets the value for Facepile to be an explicit nil

### UnsetFacepile
`func (o *SocialFeedPostDto) UnsetFacepile()`

UnsetFacepile ensures that no value is present for Facepile, not even an explicit nil
### GetAttachments

`func (o *SocialFeedPostDto) GetAttachments() []SocialPostAttachmentRefDto`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SocialFeedPostDto) GetAttachmentsOk() (*[]SocialPostAttachmentRefDto, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SocialFeedPostDto) SetAttachments(v []SocialPostAttachmentRefDto)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SocialFeedPostDto) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *SocialFeedPostDto) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *SocialFeedPostDto) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetMyReaction

`func (o *SocialFeedPostDto) GetMyReaction() string`

GetMyReaction returns the MyReaction field if non-nil, zero value otherwise.

### GetMyReactionOk

`func (o *SocialFeedPostDto) GetMyReactionOk() (*string, bool)`

GetMyReactionOk returns a tuple with the MyReaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyReaction

`func (o *SocialFeedPostDto) SetMyReaction(v string)`

SetMyReaction sets MyReaction field to given value.

### HasMyReaction

`func (o *SocialFeedPostDto) HasMyReaction() bool`

HasMyReaction returns a boolean if a field has been set.

### SetMyReactionNil

`func (o *SocialFeedPostDto) SetMyReactionNil(b bool)`

 SetMyReactionNil sets the value for MyReaction to be an explicit nil

### UnsetMyReaction
`func (o *SocialFeedPostDto) UnsetMyReaction()`

UnsetMyReaction ensures that no value is present for MyReaction, not even an explicit nil
### GetMyReactionId

`func (o *SocialFeedPostDto) GetMyReactionId() string`

GetMyReactionId returns the MyReactionId field if non-nil, zero value otherwise.

### GetMyReactionIdOk

`func (o *SocialFeedPostDto) GetMyReactionIdOk() (*string, bool)`

GetMyReactionIdOk returns a tuple with the MyReactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyReactionId

`func (o *SocialFeedPostDto) SetMyReactionId(v string)`

SetMyReactionId sets MyReactionId field to given value.

### HasMyReactionId

`func (o *SocialFeedPostDto) HasMyReactionId() bool`

HasMyReactionId returns a boolean if a field has been set.

### SetMyReactionIdNil

`func (o *SocialFeedPostDto) SetMyReactionIdNil(b bool)`

 SetMyReactionIdNil sets the value for MyReactionId to be an explicit nil

### UnsetMyReactionId
`func (o *SocialFeedPostDto) UnsetMyReactionId()`

UnsetMyReactionId ensures that no value is present for MyReactionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


