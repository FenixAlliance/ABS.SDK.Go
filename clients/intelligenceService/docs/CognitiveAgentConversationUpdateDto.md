# CognitiveAgentConversationUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelID** | Pointer to **NullableString** |  | [optional] 
**ConversationID** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**Current** | Pointer to **bool** |  | [optional] 
**Claimed** | Pointer to **bool** |  | [optional] 
**AccountHolderId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**ReceiverBusinessId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentConversationUpdateDto

`func NewCognitiveAgentConversationUpdateDto() *CognitiveAgentConversationUpdateDto`

NewCognitiveAgentConversationUpdateDto instantiates a new CognitiveAgentConversationUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentConversationUpdateDtoWithDefaults

`func NewCognitiveAgentConversationUpdateDtoWithDefaults() *CognitiveAgentConversationUpdateDto`

NewCognitiveAgentConversationUpdateDtoWithDefaults instantiates a new CognitiveAgentConversationUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelID

`func (o *CognitiveAgentConversationUpdateDto) GetChannelID() string`

GetChannelID returns the ChannelID field if non-nil, zero value otherwise.

### GetChannelIDOk

`func (o *CognitiveAgentConversationUpdateDto) GetChannelIDOk() (*string, bool)`

GetChannelIDOk returns a tuple with the ChannelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelID

`func (o *CognitiveAgentConversationUpdateDto) SetChannelID(v string)`

SetChannelID sets ChannelID field to given value.

### HasChannelID

`func (o *CognitiveAgentConversationUpdateDto) HasChannelID() bool`

HasChannelID returns a boolean if a field has been set.

### SetChannelIDNil

`func (o *CognitiveAgentConversationUpdateDto) SetChannelIDNil(b bool)`

 SetChannelIDNil sets the value for ChannelID to be an explicit nil

### UnsetChannelID
`func (o *CognitiveAgentConversationUpdateDto) UnsetChannelID()`

UnsetChannelID ensures that no value is present for ChannelID, not even an explicit nil
### GetConversationID

`func (o *CognitiveAgentConversationUpdateDto) GetConversationID() string`

GetConversationID returns the ConversationID field if non-nil, zero value otherwise.

### GetConversationIDOk

`func (o *CognitiveAgentConversationUpdateDto) GetConversationIDOk() (*string, bool)`

GetConversationIDOk returns a tuple with the ConversationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationID

`func (o *CognitiveAgentConversationUpdateDto) SetConversationID(v string)`

SetConversationID sets ConversationID field to given value.

### HasConversationID

`func (o *CognitiveAgentConversationUpdateDto) HasConversationID() bool`

HasConversationID returns a boolean if a field has been set.

### SetConversationIDNil

`func (o *CognitiveAgentConversationUpdateDto) SetConversationIDNil(b bool)`

 SetConversationIDNil sets the value for ConversationID to be an explicit nil

### UnsetConversationID
`func (o *CognitiveAgentConversationUpdateDto) UnsetConversationID()`

UnsetConversationID ensures that no value is present for ConversationID, not even an explicit nil
### GetState

`func (o *CognitiveAgentConversationUpdateDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CognitiveAgentConversationUpdateDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CognitiveAgentConversationUpdateDto) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CognitiveAgentConversationUpdateDto) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CognitiveAgentConversationUpdateDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CognitiveAgentConversationUpdateDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCurrent

`func (o *CognitiveAgentConversationUpdateDto) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CognitiveAgentConversationUpdateDto) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CognitiveAgentConversationUpdateDto) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *CognitiveAgentConversationUpdateDto) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetClaimed

`func (o *CognitiveAgentConversationUpdateDto) GetClaimed() bool`

GetClaimed returns the Claimed field if non-nil, zero value otherwise.

### GetClaimedOk

`func (o *CognitiveAgentConversationUpdateDto) GetClaimedOk() (*bool, bool)`

GetClaimedOk returns a tuple with the Claimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimed

`func (o *CognitiveAgentConversationUpdateDto) SetClaimed(v bool)`

SetClaimed sets Claimed field to given value.

### HasClaimed

`func (o *CognitiveAgentConversationUpdateDto) HasClaimed() bool`

HasClaimed returns a boolean if a field has been set.

### GetAccountHolderId

`func (o *CognitiveAgentConversationUpdateDto) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *CognitiveAgentConversationUpdateDto) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *CognitiveAgentConversationUpdateDto) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.

### HasAccountHolderId

`func (o *CognitiveAgentConversationUpdateDto) HasAccountHolderId() bool`

HasAccountHolderId returns a boolean if a field has been set.

### SetAccountHolderIdNil

`func (o *CognitiveAgentConversationUpdateDto) SetAccountHolderIdNil(b bool)`

 SetAccountHolderIdNil sets the value for AccountHolderId to be an explicit nil

### UnsetAccountHolderId
`func (o *CognitiveAgentConversationUpdateDto) UnsetAccountHolderId()`

UnsetAccountHolderId ensures that no value is present for AccountHolderId, not even an explicit nil
### GetIndividualId

`func (o *CognitiveAgentConversationUpdateDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *CognitiveAgentConversationUpdateDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *CognitiveAgentConversationUpdateDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *CognitiveAgentConversationUpdateDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *CognitiveAgentConversationUpdateDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *CognitiveAgentConversationUpdateDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *CognitiveAgentConversationUpdateDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CognitiveAgentConversationUpdateDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CognitiveAgentConversationUpdateDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CognitiveAgentConversationUpdateDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CognitiveAgentConversationUpdateDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CognitiveAgentConversationUpdateDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverBusinessId

`func (o *CognitiveAgentConversationUpdateDto) GetReceiverBusinessId() string`

GetReceiverBusinessId returns the ReceiverBusinessId field if non-nil, zero value otherwise.

### GetReceiverBusinessIdOk

`func (o *CognitiveAgentConversationUpdateDto) GetReceiverBusinessIdOk() (*string, bool)`

GetReceiverBusinessIdOk returns a tuple with the ReceiverBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBusinessId

`func (o *CognitiveAgentConversationUpdateDto) SetReceiverBusinessId(v string)`

SetReceiverBusinessId sets ReceiverBusinessId field to given value.

### HasReceiverBusinessId

`func (o *CognitiveAgentConversationUpdateDto) HasReceiverBusinessId() bool`

HasReceiverBusinessId returns a boolean if a field has been set.

### SetReceiverBusinessIdNil

`func (o *CognitiveAgentConversationUpdateDto) SetReceiverBusinessIdNil(b bool)`

 SetReceiverBusinessIdNil sets the value for ReceiverBusinessId to be an explicit nil

### UnsetReceiverBusinessId
`func (o *CognitiveAgentConversationUpdateDto) UnsetReceiverBusinessId()`

UnsetReceiverBusinessId ensures that no value is present for ReceiverBusinessId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


