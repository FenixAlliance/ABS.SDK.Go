# CognitiveAgentMessageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CognitiveAgentConversationId** | Pointer to **NullableString** |  | [optional] 
**CognitiveAgentId** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**Sequence** | Pointer to **int32** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**MetadataJson** | Pointer to **NullableString** |  | [optional] 
**AiRunId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentMessageDto

`func NewCognitiveAgentMessageDto() *CognitiveAgentMessageDto`

NewCognitiveAgentMessageDto instantiates a new CognitiveAgentMessageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentMessageDtoWithDefaults

`func NewCognitiveAgentMessageDtoWithDefaults() *CognitiveAgentMessageDto`

NewCognitiveAgentMessageDtoWithDefaults instantiates a new CognitiveAgentMessageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveAgentMessageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveAgentMessageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveAgentMessageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveAgentMessageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CognitiveAgentMessageDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CognitiveAgentMessageDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CognitiveAgentMessageDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveAgentMessageDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveAgentMessageDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveAgentMessageDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CognitiveAgentMessageDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CognitiveAgentMessageDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCognitiveAgentConversationId

`func (o *CognitiveAgentMessageDto) GetCognitiveAgentConversationId() string`

GetCognitiveAgentConversationId returns the CognitiveAgentConversationId field if non-nil, zero value otherwise.

### GetCognitiveAgentConversationIdOk

`func (o *CognitiveAgentMessageDto) GetCognitiveAgentConversationIdOk() (*string, bool)`

GetCognitiveAgentConversationIdOk returns a tuple with the CognitiveAgentConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitiveAgentConversationId

`func (o *CognitiveAgentMessageDto) SetCognitiveAgentConversationId(v string)`

SetCognitiveAgentConversationId sets CognitiveAgentConversationId field to given value.

### HasCognitiveAgentConversationId

`func (o *CognitiveAgentMessageDto) HasCognitiveAgentConversationId() bool`

HasCognitiveAgentConversationId returns a boolean if a field has been set.

### SetCognitiveAgentConversationIdNil

`func (o *CognitiveAgentMessageDto) SetCognitiveAgentConversationIdNil(b bool)`

 SetCognitiveAgentConversationIdNil sets the value for CognitiveAgentConversationId to be an explicit nil

### UnsetCognitiveAgentConversationId
`func (o *CognitiveAgentMessageDto) UnsetCognitiveAgentConversationId()`

UnsetCognitiveAgentConversationId ensures that no value is present for CognitiveAgentConversationId, not even an explicit nil
### GetCognitiveAgentId

`func (o *CognitiveAgentMessageDto) GetCognitiveAgentId() string`

GetCognitiveAgentId returns the CognitiveAgentId field if non-nil, zero value otherwise.

### GetCognitiveAgentIdOk

`func (o *CognitiveAgentMessageDto) GetCognitiveAgentIdOk() (*string, bool)`

GetCognitiveAgentIdOk returns a tuple with the CognitiveAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitiveAgentId

`func (o *CognitiveAgentMessageDto) SetCognitiveAgentId(v string)`

SetCognitiveAgentId sets CognitiveAgentId field to given value.

### HasCognitiveAgentId

`func (o *CognitiveAgentMessageDto) HasCognitiveAgentId() bool`

HasCognitiveAgentId returns a boolean if a field has been set.

### SetCognitiveAgentIdNil

`func (o *CognitiveAgentMessageDto) SetCognitiveAgentIdNil(b bool)`

 SetCognitiveAgentIdNil sets the value for CognitiveAgentId to be an explicit nil

### UnsetCognitiveAgentId
`func (o *CognitiveAgentMessageDto) UnsetCognitiveAgentId()`

UnsetCognitiveAgentId ensures that no value is present for CognitiveAgentId, not even an explicit nil
### GetRole

`func (o *CognitiveAgentMessageDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CognitiveAgentMessageDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CognitiveAgentMessageDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CognitiveAgentMessageDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *CognitiveAgentMessageDto) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *CognitiveAgentMessageDto) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetContent

`func (o *CognitiveAgentMessageDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CognitiveAgentMessageDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CognitiveAgentMessageDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CognitiveAgentMessageDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CognitiveAgentMessageDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CognitiveAgentMessageDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetSequence

`func (o *CognitiveAgentMessageDto) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *CognitiveAgentMessageDto) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *CognitiveAgentMessageDto) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *CognitiveAgentMessageDto) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetCreationDate

`func (o *CognitiveAgentMessageDto) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CognitiveAgentMessageDto) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CognitiveAgentMessageDto) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *CognitiveAgentMessageDto) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetMetadataJson

`func (o *CognitiveAgentMessageDto) GetMetadataJson() string`

GetMetadataJson returns the MetadataJson field if non-nil, zero value otherwise.

### GetMetadataJsonOk

`func (o *CognitiveAgentMessageDto) GetMetadataJsonOk() (*string, bool)`

GetMetadataJsonOk returns a tuple with the MetadataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataJson

`func (o *CognitiveAgentMessageDto) SetMetadataJson(v string)`

SetMetadataJson sets MetadataJson field to given value.

### HasMetadataJson

`func (o *CognitiveAgentMessageDto) HasMetadataJson() bool`

HasMetadataJson returns a boolean if a field has been set.

### SetMetadataJsonNil

`func (o *CognitiveAgentMessageDto) SetMetadataJsonNil(b bool)`

 SetMetadataJsonNil sets the value for MetadataJson to be an explicit nil

### UnsetMetadataJson
`func (o *CognitiveAgentMessageDto) UnsetMetadataJson()`

UnsetMetadataJson ensures that no value is present for MetadataJson, not even an explicit nil
### GetAiRunId

`func (o *CognitiveAgentMessageDto) GetAiRunId() string`

GetAiRunId returns the AiRunId field if non-nil, zero value otherwise.

### GetAiRunIdOk

`func (o *CognitiveAgentMessageDto) GetAiRunIdOk() (*string, bool)`

GetAiRunIdOk returns a tuple with the AiRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiRunId

`func (o *CognitiveAgentMessageDto) SetAiRunId(v string)`

SetAiRunId sets AiRunId field to given value.

### HasAiRunId

`func (o *CognitiveAgentMessageDto) HasAiRunId() bool`

HasAiRunId returns a boolean if a field has been set.

### SetAiRunIdNil

`func (o *CognitiveAgentMessageDto) SetAiRunIdNil(b bool)`

 SetAiRunIdNil sets the value for AiRunId to be an explicit nil

### UnsetAiRunId
`func (o *CognitiveAgentMessageDto) UnsetAiRunId()`

UnsetAiRunId ensures that no value is present for AiRunId, not even an explicit nil
### GetTenantId

`func (o *CognitiveAgentMessageDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CognitiveAgentMessageDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CognitiveAgentMessageDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CognitiveAgentMessageDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CognitiveAgentMessageDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CognitiveAgentMessageDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CognitiveAgentMessageDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CognitiveAgentMessageDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CognitiveAgentMessageDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CognitiveAgentMessageDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CognitiveAgentMessageDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CognitiveAgentMessageDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


