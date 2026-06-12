# SupportTicketConversationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Topic** | Pointer to **NullableString** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**ClosedTimestamp** | Pointer to **time.Time** |  | [optional] 
**SocialProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSupportTicketConversationCreateDto

`func NewSupportTicketConversationCreateDto() *SupportTicketConversationCreateDto`

NewSupportTicketConversationCreateDto instantiates a new SupportTicketConversationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportTicketConversationCreateDtoWithDefaults

`func NewSupportTicketConversationCreateDtoWithDefaults() *SupportTicketConversationCreateDto`

NewSupportTicketConversationCreateDtoWithDefaults instantiates a new SupportTicketConversationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportTicketConversationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportTicketConversationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportTicketConversationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupportTicketConversationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SupportTicketConversationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SupportTicketConversationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SupportTicketConversationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SupportTicketConversationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTopic

`func (o *SupportTicketConversationCreateDto) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *SupportTicketConversationCreateDto) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *SupportTicketConversationCreateDto) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *SupportTicketConversationCreateDto) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *SupportTicketConversationCreateDto) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *SupportTicketConversationCreateDto) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetClosed

`func (o *SupportTicketConversationCreateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *SupportTicketConversationCreateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *SupportTicketConversationCreateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *SupportTicketConversationCreateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetClosedTimestamp

`func (o *SupportTicketConversationCreateDto) GetClosedTimestamp() time.Time`

GetClosedTimestamp returns the ClosedTimestamp field if non-nil, zero value otherwise.

### GetClosedTimestampOk

`func (o *SupportTicketConversationCreateDto) GetClosedTimestampOk() (*time.Time, bool)`

GetClosedTimestampOk returns a tuple with the ClosedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedTimestamp

`func (o *SupportTicketConversationCreateDto) SetClosedTimestamp(v time.Time)`

SetClosedTimestamp sets ClosedTimestamp field to given value.

### HasClosedTimestamp

`func (o *SupportTicketConversationCreateDto) HasClosedTimestamp() bool`

HasClosedTimestamp returns a boolean if a field has been set.

### GetSocialProfileId

`func (o *SupportTicketConversationCreateDto) GetSocialProfileId() string`

GetSocialProfileId returns the SocialProfileId field if non-nil, zero value otherwise.

### GetSocialProfileIdOk

`func (o *SupportTicketConversationCreateDto) GetSocialProfileIdOk() (*string, bool)`

GetSocialProfileIdOk returns a tuple with the SocialProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfileId

`func (o *SupportTicketConversationCreateDto) SetSocialProfileId(v string)`

SetSocialProfileId sets SocialProfileId field to given value.

### HasSocialProfileId

`func (o *SupportTicketConversationCreateDto) HasSocialProfileId() bool`

HasSocialProfileId returns a boolean if a field has been set.

### SetSocialProfileIdNil

`func (o *SupportTicketConversationCreateDto) SetSocialProfileIdNil(b bool)`

 SetSocialProfileIdNil sets the value for SocialProfileId to be an explicit nil

### UnsetSocialProfileId
`func (o *SupportTicketConversationCreateDto) UnsetSocialProfileId()`

UnsetSocialProfileId ensures that no value is present for SocialProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


