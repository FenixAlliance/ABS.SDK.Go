# CognitiveAgentConversationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CognitiveAgentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Claimed** | Pointer to **bool** |  | [optional] 
**Current** | Pointer to **bool** |  | [optional] 
**ChannelID** | Pointer to **NullableString** |  | [optional] 
**ConversationID** | Pointer to **NullableString** |  | [optional] 
**ActivityID** | Pointer to **NullableString** |  | [optional] 
**ActivityFrom** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**LastActivity** | Pointer to **time.Time** |  | [optional] 
**AccountHolderId** | Pointer to **NullableString** |  | [optional] 
**IndividualId** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **NullableString** |  | [optional] 
**ReceiverBusinessId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentConversationDto

`func NewCognitiveAgentConversationDto() *CognitiveAgentConversationDto`

NewCognitiveAgentConversationDto instantiates a new CognitiveAgentConversationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentConversationDtoWithDefaults

`func NewCognitiveAgentConversationDtoWithDefaults() *CognitiveAgentConversationDto`

NewCognitiveAgentConversationDtoWithDefaults instantiates a new CognitiveAgentConversationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveAgentConversationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveAgentConversationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveAgentConversationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveAgentConversationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CognitiveAgentConversationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CognitiveAgentConversationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CognitiveAgentConversationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveAgentConversationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveAgentConversationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveAgentConversationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CognitiveAgentConversationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CognitiveAgentConversationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCognitiveAgentId

`func (o *CognitiveAgentConversationDto) GetCognitiveAgentId() string`

GetCognitiveAgentId returns the CognitiveAgentId field if non-nil, zero value otherwise.

### GetCognitiveAgentIdOk

`func (o *CognitiveAgentConversationDto) GetCognitiveAgentIdOk() (*string, bool)`

GetCognitiveAgentIdOk returns a tuple with the CognitiveAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitiveAgentId

`func (o *CognitiveAgentConversationDto) SetCognitiveAgentId(v string)`

SetCognitiveAgentId sets CognitiveAgentId field to given value.

### HasCognitiveAgentId

`func (o *CognitiveAgentConversationDto) HasCognitiveAgentId() bool`

HasCognitiveAgentId returns a boolean if a field has been set.

### SetCognitiveAgentIdNil

`func (o *CognitiveAgentConversationDto) SetCognitiveAgentIdNil(b bool)`

 SetCognitiveAgentIdNil sets the value for CognitiveAgentId to be an explicit nil

### UnsetCognitiveAgentId
`func (o *CognitiveAgentConversationDto) UnsetCognitiveAgentId()`

UnsetCognitiveAgentId ensures that no value is present for CognitiveAgentId, not even an explicit nil
### GetTenantId

`func (o *CognitiveAgentConversationDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CognitiveAgentConversationDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CognitiveAgentConversationDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CognitiveAgentConversationDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CognitiveAgentConversationDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CognitiveAgentConversationDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CognitiveAgentConversationDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CognitiveAgentConversationDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CognitiveAgentConversationDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CognitiveAgentConversationDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CognitiveAgentConversationDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CognitiveAgentConversationDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTitle

`func (o *CognitiveAgentConversationDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CognitiveAgentConversationDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CognitiveAgentConversationDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CognitiveAgentConversationDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CognitiveAgentConversationDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CognitiveAgentConversationDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetClaimed

`func (o *CognitiveAgentConversationDto) GetClaimed() bool`

GetClaimed returns the Claimed field if non-nil, zero value otherwise.

### GetClaimedOk

`func (o *CognitiveAgentConversationDto) GetClaimedOk() (*bool, bool)`

GetClaimedOk returns a tuple with the Claimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimed

`func (o *CognitiveAgentConversationDto) SetClaimed(v bool)`

SetClaimed sets Claimed field to given value.

### HasClaimed

`func (o *CognitiveAgentConversationDto) HasClaimed() bool`

HasClaimed returns a boolean if a field has been set.

### GetCurrent

`func (o *CognitiveAgentConversationDto) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *CognitiveAgentConversationDto) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *CognitiveAgentConversationDto) SetCurrent(v bool)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *CognitiveAgentConversationDto) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetChannelID

`func (o *CognitiveAgentConversationDto) GetChannelID() string`

GetChannelID returns the ChannelID field if non-nil, zero value otherwise.

### GetChannelIDOk

`func (o *CognitiveAgentConversationDto) GetChannelIDOk() (*string, bool)`

GetChannelIDOk returns a tuple with the ChannelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelID

`func (o *CognitiveAgentConversationDto) SetChannelID(v string)`

SetChannelID sets ChannelID field to given value.

### HasChannelID

`func (o *CognitiveAgentConversationDto) HasChannelID() bool`

HasChannelID returns a boolean if a field has been set.

### SetChannelIDNil

`func (o *CognitiveAgentConversationDto) SetChannelIDNil(b bool)`

 SetChannelIDNil sets the value for ChannelID to be an explicit nil

### UnsetChannelID
`func (o *CognitiveAgentConversationDto) UnsetChannelID()`

UnsetChannelID ensures that no value is present for ChannelID, not even an explicit nil
### GetConversationID

`func (o *CognitiveAgentConversationDto) GetConversationID() string`

GetConversationID returns the ConversationID field if non-nil, zero value otherwise.

### GetConversationIDOk

`func (o *CognitiveAgentConversationDto) GetConversationIDOk() (*string, bool)`

GetConversationIDOk returns a tuple with the ConversationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationID

`func (o *CognitiveAgentConversationDto) SetConversationID(v string)`

SetConversationID sets ConversationID field to given value.

### HasConversationID

`func (o *CognitiveAgentConversationDto) HasConversationID() bool`

HasConversationID returns a boolean if a field has been set.

### SetConversationIDNil

`func (o *CognitiveAgentConversationDto) SetConversationIDNil(b bool)`

 SetConversationIDNil sets the value for ConversationID to be an explicit nil

### UnsetConversationID
`func (o *CognitiveAgentConversationDto) UnsetConversationID()`

UnsetConversationID ensures that no value is present for ConversationID, not even an explicit nil
### GetActivityID

`func (o *CognitiveAgentConversationDto) GetActivityID() string`

GetActivityID returns the ActivityID field if non-nil, zero value otherwise.

### GetActivityIDOk

`func (o *CognitiveAgentConversationDto) GetActivityIDOk() (*string, bool)`

GetActivityIDOk returns a tuple with the ActivityID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityID

`func (o *CognitiveAgentConversationDto) SetActivityID(v string)`

SetActivityID sets ActivityID field to given value.

### HasActivityID

`func (o *CognitiveAgentConversationDto) HasActivityID() bool`

HasActivityID returns a boolean if a field has been set.

### SetActivityIDNil

`func (o *CognitiveAgentConversationDto) SetActivityIDNil(b bool)`

 SetActivityIDNil sets the value for ActivityID to be an explicit nil

### UnsetActivityID
`func (o *CognitiveAgentConversationDto) UnsetActivityID()`

UnsetActivityID ensures that no value is present for ActivityID, not even an explicit nil
### GetActivityFrom

`func (o *CognitiveAgentConversationDto) GetActivityFrom() string`

GetActivityFrom returns the ActivityFrom field if non-nil, zero value otherwise.

### GetActivityFromOk

`func (o *CognitiveAgentConversationDto) GetActivityFromOk() (*string, bool)`

GetActivityFromOk returns a tuple with the ActivityFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityFrom

`func (o *CognitiveAgentConversationDto) SetActivityFrom(v string)`

SetActivityFrom sets ActivityFrom field to given value.

### HasActivityFrom

`func (o *CognitiveAgentConversationDto) HasActivityFrom() bool`

HasActivityFrom returns a boolean if a field has been set.

### SetActivityFromNil

`func (o *CognitiveAgentConversationDto) SetActivityFromNil(b bool)`

 SetActivityFromNil sets the value for ActivityFrom to be an explicit nil

### UnsetActivityFrom
`func (o *CognitiveAgentConversationDto) UnsetActivityFrom()`

UnsetActivityFrom ensures that no value is present for ActivityFrom, not even an explicit nil
### GetState

`func (o *CognitiveAgentConversationDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CognitiveAgentConversationDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CognitiveAgentConversationDto) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CognitiveAgentConversationDto) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CognitiveAgentConversationDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CognitiveAgentConversationDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetLastActivity

`func (o *CognitiveAgentConversationDto) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *CognitiveAgentConversationDto) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *CognitiveAgentConversationDto) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *CognitiveAgentConversationDto) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetAccountHolderId

`func (o *CognitiveAgentConversationDto) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *CognitiveAgentConversationDto) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *CognitiveAgentConversationDto) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.

### HasAccountHolderId

`func (o *CognitiveAgentConversationDto) HasAccountHolderId() bool`

HasAccountHolderId returns a boolean if a field has been set.

### SetAccountHolderIdNil

`func (o *CognitiveAgentConversationDto) SetAccountHolderIdNil(b bool)`

 SetAccountHolderIdNil sets the value for AccountHolderId to be an explicit nil

### UnsetAccountHolderId
`func (o *CognitiveAgentConversationDto) UnsetAccountHolderId()`

UnsetAccountHolderId ensures that no value is present for AccountHolderId, not even an explicit nil
### GetIndividualId

`func (o *CognitiveAgentConversationDto) GetIndividualId() string`

GetIndividualId returns the IndividualId field if non-nil, zero value otherwise.

### GetIndividualIdOk

`func (o *CognitiveAgentConversationDto) GetIndividualIdOk() (*string, bool)`

GetIndividualIdOk returns a tuple with the IndividualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualId

`func (o *CognitiveAgentConversationDto) SetIndividualId(v string)`

SetIndividualId sets IndividualId field to given value.

### HasIndividualId

`func (o *CognitiveAgentConversationDto) HasIndividualId() bool`

HasIndividualId returns a boolean if a field has been set.

### SetIndividualIdNil

`func (o *CognitiveAgentConversationDto) SetIndividualIdNil(b bool)`

 SetIndividualIdNil sets the value for IndividualId to be an explicit nil

### UnsetIndividualId
`func (o *CognitiveAgentConversationDto) UnsetIndividualId()`

UnsetIndividualId ensures that no value is present for IndividualId, not even an explicit nil
### GetOrganizationId

`func (o *CognitiveAgentConversationDto) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CognitiveAgentConversationDto) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CognitiveAgentConversationDto) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CognitiveAgentConversationDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CognitiveAgentConversationDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CognitiveAgentConversationDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetReceiverBusinessId

`func (o *CognitiveAgentConversationDto) GetReceiverBusinessId() string`

GetReceiverBusinessId returns the ReceiverBusinessId field if non-nil, zero value otherwise.

### GetReceiverBusinessIdOk

`func (o *CognitiveAgentConversationDto) GetReceiverBusinessIdOk() (*string, bool)`

GetReceiverBusinessIdOk returns a tuple with the ReceiverBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverBusinessId

`func (o *CognitiveAgentConversationDto) SetReceiverBusinessId(v string)`

SetReceiverBusinessId sets ReceiverBusinessId field to given value.

### HasReceiverBusinessId

`func (o *CognitiveAgentConversationDto) HasReceiverBusinessId() bool`

HasReceiverBusinessId returns a boolean if a field has been set.

### SetReceiverBusinessIdNil

`func (o *CognitiveAgentConversationDto) SetReceiverBusinessIdNil(b bool)`

 SetReceiverBusinessIdNil sets the value for ReceiverBusinessId to be an explicit nil

### UnsetReceiverBusinessId
`func (o *CognitiveAgentConversationDto) UnsetReceiverBusinessId()`

UnsetReceiverBusinessId ensures that no value is present for ReceiverBusinessId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


