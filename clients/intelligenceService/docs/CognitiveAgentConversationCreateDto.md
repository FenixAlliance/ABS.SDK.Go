# CognitiveAgentConversationCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ChannelID** | Pointer to **NullableString** |  | [optional] 
**ConversationID** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**Current** | Pointer to **bool** |  | [optional] 
**AccountHolderId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**ReceiverBusinessId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentConversationCreateDto

`func NewCognitiveAgentConversationCreateDto() *CognitiveAgentConversationCreateDto`

NewCognitiveAgentConversationCreateDto instantiates a new CognitiveAgentConversationCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentConversationCreateDtoWithDefaults

`func NewCognitiveAgentConversationCreateDtoWithDefaults() *CognitiveAgentConversationCreateDto`

NewCognitiveAgentConversationCreateDtoWithDefaults instantiates a new CognitiveAgentConversationCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveAgentConversationCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveAgentConversationCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveAgentConversationCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveAgentConversationCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CognitiveAgentConversationCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveAgentConversationCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveAgentConversationCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveAgentConversationCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetChannelID

`func (o *CognitiveAgentConversationCreateDto) GetChannelID() string`

GetChannelID returns the ChannelID field if non-nil, zero value otherwise.

### GetChannelIDOk

`func (o *CognitiveAgentConversationCreateDto) GetChannelIDOk() (*string, bool)`

GetChannelIDOk returns a tuple with the ChannelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelID

`func (o *CognitiveAgentConversationCreateDto) SetChannelID(v string)`

SetChannelID sets ChannelID field to given value.

### HasChannelID

`func (o *CognitiveAgentConversationCreateDto) HasChannelID() bool`

HasChannelID returns a boolean if a field has been set.

### SetChannelIDNil

`func (o *CognitiveAgentConversationCreateDto) SetChannelIDNil(b bool)`

 SetChannelIDNil sets the value for ChannelID to be an explicit nil

### UnsetChannelID
`func (o *CognitiveAgentConversationCreateDto) UnsetChannelID()`

UnsetChannelID ensures that no value is present for ChannelID, not even an explicit nil
### GetConversationID

`func (o *CognitiveAgentConversationCreateDto) GetConversationID() string`

GetConversationID returns the ConversationID field if non-nil, zero value otherwise.

### GetConversationIDOk

`func (o *CognitiveAgentConversationCreateDto) GetConversationIDOk() (*string, bool)`

GetConversationIDOk returns a tuple with the ConversationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationID

`func (o *CognitiveAgentConversationCreateDto) SetConversationID(v string)`

SetConversationID sets ConversationID field to given value.

### HasConversationID

`func (o *CognitiveAgentConversationCreateDto) HasConversationID() bool`

HasConversationID returns a boolean if a field has been set.

### SetConversationIDNil

`func (o *CognitiveAgentConversationCreateDto) SetConversationIDNil(b bool)`

 SetConversationIDNil sets the value for ConversationID to be an explicit nil

### UnsetConversationID
`func (o *CognitiveAgentConversationCreateDto) UnsetConversationID()`

UnsetConversationID ensures that no value is present for ConversationID, not even an explicit nil
### GetState

`func (o *CognitiveAgentConversationCreateDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CognitiveAgentConversationCreateDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CognitiveAgentConversationCreateDto) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CognitiveAgentConversationCreateDto) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CognitiveAgentConversationCreateDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CognitiveAgentConversationCreateDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCurrent

`func (o *CognitiveAgentConversationCreateDto) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CognitiveAgentConversationCreateDto) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CognitiveAgentConversationCreateDto) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *CognitiveAgentConversationCreateDto) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetAccountHolderId

`func (o *CognitiveAgentConversationCreateDto) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *CognitiveAgentConversationCreateDto) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *CognitiveAgentConversationCreateDto) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.

### HasAccountHolderId

`func (o *CognitiveAgentConversationCreateDto) HasAccountHolderId() bool`

HasAccountHolderId returns a boolean if a field has been set.

### SetAccountHolderIdNil

`func (o *CognitiveAgentConversationCreateDto) SetAccountHolderIdNil(b bool)`

 SetAccountHolderIdNil sets the value for AccountHolderId to be an explicit nil

### UnsetAccountHolderId
`func (o *CognitiveAgentConversationCreateDto) UnsetAccountHolderId()`

UnsetAccountHolderId ensures that no value is present for AccountHolderId, not even an explicit nil
### GetIndividualId

`func (o *CognitiveAgentConversationCreateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *CognitiveAgentConversationCreateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *CognitiveAgentConversationCreateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *CognitiveAgentConversationCreateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *CognitiveAgentConversationCreateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *CognitiveAgentConversationCreateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *CognitiveAgentConversationCreateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CognitiveAgentConversationCreateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CognitiveAgentConversationCreateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CognitiveAgentConversationCreateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CognitiveAgentConversationCreateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CognitiveAgentConversationCreateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverBusinessId

`func (o *CognitiveAgentConversationCreateDto) GetReceiverBusinessId() string`

GetReceiverBusinessId returns the ReceiverBusinessId field if non-nil, zero value otherwise.

### GetReceiverBusinessIdOk

`func (o *CognitiveAgentConversationCreateDto) GetReceiverBusinessIdOk() (*string, bool)`

GetReceiverBusinessIdOk returns a tuple with the ReceiverBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBusinessId

`func (o *CognitiveAgentConversationCreateDto) SetReceiverBusinessId(v string)`

SetReceiverBusinessId sets ReceiverBusinessId field to given value.

### HasReceiverBusinessId

`func (o *CognitiveAgentConversationCreateDto) HasReceiverBusinessId() bool`

HasReceiverBusinessId returns a boolean if a field has been set.

### SetReceiverBusinessIdNil

`func (o *CognitiveAgentConversationCreateDto) SetReceiverBusinessIdNil(b bool)`

 SetReceiverBusinessIdNil sets the value for ReceiverBusinessId to be an explicit nil

### UnsetReceiverBusinessId
`func (o *CognitiveAgentConversationCreateDto) UnsetReceiverBusinessId()`

UnsetReceiverBusinessId ensures that no value is present for ReceiverBusinessId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


