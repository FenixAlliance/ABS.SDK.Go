# NotificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Read** | Pointer to **bool** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 
**ReadTimestamp** | Pointer to **time.Time** |  | [optional] 
**IssuedTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNotificationDto

`func NewNotificationDto() *NotificationDto`

NewNotificationDto instantiates a new NotificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDtoWithDefaults

`func NewNotificationDtoWithDefaults() *NotificationDto`

NewNotificationDtoWithDefaults instantiates a new NotificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NotificationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NotificationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *NotificationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NotificationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NotificationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NotificationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *NotificationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *NotificationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRead

`func (o *NotificationDto) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *NotificationDto) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *NotificationDto) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *NotificationDto) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetIcon

`func (o *NotificationDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *NotificationDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *NotificationDto) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *NotificationDto) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *NotificationDto) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *NotificationDto) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetMessage

`func (o *NotificationDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotificationDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *NotificationDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *NotificationDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRedirectUrl

`func (o *NotificationDto) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *NotificationDto) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *NotificationDto) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *NotificationDto) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *NotificationDto) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *NotificationDto) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetSocialProfileID

`func (o *NotificationDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *NotificationDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *NotificationDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *NotificationDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *NotificationDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *NotificationDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil
### GetReadTimestamp

`func (o *NotificationDto) GetReadTimestamp() time.Time`

GetReadTimestamp returns the ReadTimestamp field if non-nil, zero value otherwise.

### GetReadTimestampOk

`func (o *NotificationDto) GetReadTimestampOk() (*time.Time, bool)`

GetReadTimestampOk returns a tuple with the ReadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimestamp

`func (o *NotificationDto) SetReadTimestamp(v time.Time)`

SetReadTimestamp sets ReadTimestamp field to given value.

### HasReadTimestamp

`func (o *NotificationDto) HasReadTimestamp() bool`

HasReadTimestamp returns a boolean if a field has been set.

### GetIssuedTimestamp

`func (o *NotificationDto) GetIssuedTimestamp() time.Time`

GetIssuedTimestamp returns the IssuedTimestamp field if non-nil, zero value otherwise.

### GetIssuedTimestampOk

`func (o *NotificationDto) GetIssuedTimestampOk() (*time.Time, bool)`

GetIssuedTimestampOk returns a tuple with the IssuedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTimestamp

`func (o *NotificationDto) SetIssuedTimestamp(v time.Time)`

SetIssuedTimestamp sets IssuedTimestamp field to given value.

### HasIssuedTimestamp

`func (o *NotificationDto) HasIssuedTimestamp() bool`

HasIssuedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


