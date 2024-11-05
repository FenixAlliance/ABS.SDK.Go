# SupportTicketConversationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**SupportTicketID** | Pointer to **NullableString** |  | [optional] 
**Topic** | Pointer to **NullableString** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**ClosedTimestamp** | Pointer to **time.Time** |  | [optional] 
**SocialProfileID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportTicketConversationDto

`func NewSupportTicketConversationDto() *SupportTicketConversationDto`

NewSupportTicketConversationDto instantiates a new SupportTicketConversationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketConversationDtoWithDefaults

`func NewSupportTicketConversationDtoWithDefaults() *SupportTicketConversationDto`

NewSupportTicketConversationDtoWithDefaults instantiates a new SupportTicketConversationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportTicketConversationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicketConversationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicketConversationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicketConversationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SupportTicketConversationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SupportTicketConversationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SupportTicketConversationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportTicketConversationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportTicketConversationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportTicketConversationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SupportTicketConversationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SupportTicketConversationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSupportTicketID

`func (o *SupportTicketConversationDto) GetSupportTicketID() string`

GetSupportTicketID returns the SupportTicketID field if non-nil, zero value otherwise.

### GetSupportTicketIDOk

`func (o *SupportTicketConversationDto) GetSupportTicketIDOk() (*string, bool)`

GetSupportTicketIDOk returns a tuple with the SupportTicketID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTicketID

`func (o *SupportTicketConversationDto) SetSupportTicketID(v string)`

SetSupportTicketID sets SupportTicketID field to given value.

### HasSupportTicketID

`func (o *SupportTicketConversationDto) HasSupportTicketID() bool`

HasSupportTicketID returns a boolean if a field has been set.

### SetSupportTicketIDNil

`func (o *SupportTicketConversationDto) SetSupportTicketIDNil(b bool)`

 SetSupportTicketIDNil sets the value for SupportTicketID to be an explicit nil

### UnsetSupportTicketID
`func (o *SupportTicketConversationDto) UnsetSupportTicketID()`

UnsetSupportTicketID ensures that no value is present for SupportTicketID, not even an explicit nil
### GetTopic

`func (o *SupportTicketConversationDto) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *SupportTicketConversationDto) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *SupportTicketConversationDto) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *SupportTicketConversationDto) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *SupportTicketConversationDto) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *SupportTicketConversationDto) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetClosed

`func (o *SupportTicketConversationDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *SupportTicketConversationDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *SupportTicketConversationDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *SupportTicketConversationDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClosedTimestamp

`func (o *SupportTicketConversationDto) GetClosedTimestamp() time.Time`

GetClosedTimestamp returns the ClosedTimestamp field if non-nil, zero value otherwise.

### GetClosedTimestampOk

`func (o *SupportTicketConversationDto) GetClosedTimestampOk() (*time.Time, bool)`

GetClosedTimestampOk returns a tuple with the ClosedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedTimestamp

`func (o *SupportTicketConversationDto) SetClosedTimestamp(v time.Time)`

SetClosedTimestamp sets ClosedTimestamp field to given value.

### HasClosedTimestamp

`func (o *SupportTicketConversationDto) HasClosedTimestamp() bool`

HasClosedTimestamp returns a boolean if a field has been set.

### GetSocialProfileID

`func (o *SupportTicketConversationDto) GetSocialProfileID() string`

GetSocialProfileID returns the SocialProfileID field if non-nil, zero value otherwise.

### GetSocialProfileIDOk

`func (o *SupportTicketConversationDto) GetSocialProfileIDOk() (*string, bool)`

GetSocialProfileIDOk returns a tuple with the SocialProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileID

`func (o *SupportTicketConversationDto) SetSocialProfileID(v string)`

SetSocialProfileID sets SocialProfileID field to given value.

### HasSocialProfileID

`func (o *SupportTicketConversationDto) HasSocialProfileID() bool`

HasSocialProfileID returns a boolean if a field has been set.

### SetSocialProfileIDNil

`func (o *SupportTicketConversationDto) SetSocialProfileIDNil(b bool)`

 SetSocialProfileIDNil sets the value for SocialProfileID to be an explicit nil

### UnsetSocialProfileID
`func (o *SupportTicketConversationDto) UnsetSocialProfileID()`

UnsetSocialProfileID ensures that no value is present for SocialProfileID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


