# CognitiveAgentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Soul** | Pointer to **NullableString** |  | [optional] 
**ProviderKey** | Pointer to **NullableString** |  | [optional] 
**ModelId** | Pointer to **NullableString** |  | [optional] 
**EngineKey** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentDto

`func NewCognitiveAgentDto() *CognitiveAgentDto`

NewCognitiveAgentDto instantiates a new CognitiveAgentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentDtoWithDefaults

`func NewCognitiveAgentDtoWithDefaults() *CognitiveAgentDto`

NewCognitiveAgentDtoWithDefaults instantiates a new CognitiveAgentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveAgentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveAgentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveAgentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveAgentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CognitiveAgentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CognitiveAgentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CognitiveAgentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveAgentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveAgentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveAgentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CognitiveAgentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CognitiveAgentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *CognitiveAgentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CognitiveAgentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CognitiveAgentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CognitiveAgentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CognitiveAgentDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CognitiveAgentDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAvatar

`func (o *CognitiveAgentDto) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CognitiveAgentDto) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CognitiveAgentDto) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CognitiveAgentDto) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *CognitiveAgentDto) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *CognitiveAgentDto) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetDescription

`func (o *CognitiveAgentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CognitiveAgentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CognitiveAgentDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CognitiveAgentDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CognitiveAgentDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CognitiveAgentDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSoul

`func (o *CognitiveAgentDto) GetSoul() string`

GetSoul returns the Soul field if non-nil, zero value otherwise.

### GetSoulOk

`func (o *CognitiveAgentDto) GetSoulOk() (*string, bool)`

GetSoulOk returns a tuple with the Soul field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoul

`func (o *CognitiveAgentDto) SetSoul(v string)`

SetSoul sets Soul field to given value.

### HasSoul

`func (o *CognitiveAgentDto) HasSoul() bool`

HasSoul returns a boolean if a field has been set.

### SetSoulNil

`func (o *CognitiveAgentDto) SetSoulNil(b bool)`

 SetSoulNil sets the value for Soul to be an explicit nil

### UnsetSoul
`func (o *CognitiveAgentDto) UnsetSoul()`

UnsetSoul ensures that no value is present for Soul, not even an explicit nil
### GetProviderKey

`func (o *CognitiveAgentDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *CognitiveAgentDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *CognitiveAgentDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *CognitiveAgentDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *CognitiveAgentDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *CognitiveAgentDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetModelId

`func (o *CognitiveAgentDto) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CognitiveAgentDto) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CognitiveAgentDto) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *CognitiveAgentDto) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### SetModelIdNil

`func (o *CognitiveAgentDto) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *CognitiveAgentDto) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetEngineKey

`func (o *CognitiveAgentDto) GetEngineKey() string`

GetEngineKey returns the EngineKey field if non-nil, zero value otherwise.

### GetEngineKeyOk

`func (o *CognitiveAgentDto) GetEngineKeyOk() (*string, bool)`

GetEngineKeyOk returns a tuple with the EngineKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineKey

`func (o *CognitiveAgentDto) SetEngineKey(v string)`

SetEngineKey sets EngineKey field to given value.

### HasEngineKey

`func (o *CognitiveAgentDto) HasEngineKey() bool`

HasEngineKey returns a boolean if a field has been set.

### SetEngineKeyNil

`func (o *CognitiveAgentDto) SetEngineKeyNil(b bool)`

 SetEngineKeyNil sets the value for EngineKey to be an explicit nil

### UnsetEngineKey
`func (o *CognitiveAgentDto) UnsetEngineKey()`

UnsetEngineKey ensures that no value is present for EngineKey, not even an explicit nil
### GetTenantId

`func (o *CognitiveAgentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CognitiveAgentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CognitiveAgentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CognitiveAgentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CognitiveAgentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CognitiveAgentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CognitiveAgentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CognitiveAgentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CognitiveAgentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CognitiveAgentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CognitiveAgentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CognitiveAgentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


