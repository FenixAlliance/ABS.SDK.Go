# CognitiveAgentCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Soul** | Pointer to **NullableString** |  | [optional] 
**ProviderKey** | Pointer to **NullableString** |  | [optional] 
**ModelId** | Pointer to **NullableString** |  | [optional] 
**EngineKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentCreateDto

`func NewCognitiveAgentCreateDto(name string, ) *CognitiveAgentCreateDto`

NewCognitiveAgentCreateDto instantiates a new CognitiveAgentCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentCreateDtoWithDefaults

`func NewCognitiveAgentCreateDtoWithDefaults() *CognitiveAgentCreateDto`

NewCognitiveAgentCreateDtoWithDefaults instantiates a new CognitiveAgentCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveAgentCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveAgentCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveAgentCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveAgentCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CognitiveAgentCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveAgentCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveAgentCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveAgentCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CognitiveAgentCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CognitiveAgentCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CognitiveAgentCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetAvatar

`func (o *CognitiveAgentCreateDto) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CognitiveAgentCreateDto) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CognitiveAgentCreateDto) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CognitiveAgentCreateDto) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *CognitiveAgentCreateDto) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *CognitiveAgentCreateDto) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetDescription

`func (o *CognitiveAgentCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CognitiveAgentCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CognitiveAgentCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CognitiveAgentCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CognitiveAgentCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CognitiveAgentCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSoul

`func (o *CognitiveAgentCreateDto) GetSoul() string`

GetSoul returns the Soul field if non-nil, zero value otherwise.

### GetSoulOk

`func (o *CognitiveAgentCreateDto) GetSoulOk() (*string, bool)`

GetSoulOk returns a tuple with the Soul field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoul

`func (o *CognitiveAgentCreateDto) SetSoul(v string)`

SetSoul sets Soul field to given value.

### HasSoul

`func (o *CognitiveAgentCreateDto) HasSoul() bool`

HasSoul returns a boolean if a field has been set.

### SetSoulNil

`func (o *CognitiveAgentCreateDto) SetSoulNil(b bool)`

 SetSoulNil sets the value for Soul to be an explicit nil

### UnsetSoul
`func (o *CognitiveAgentCreateDto) UnsetSoul()`

UnsetSoul ensures that no value is present for Soul, not even an explicit nil
### GetProviderKey

`func (o *CognitiveAgentCreateDto) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *CognitiveAgentCreateDto) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *CognitiveAgentCreateDto) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *CognitiveAgentCreateDto) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### SetProviderKeyNil

`func (o *CognitiveAgentCreateDto) SetProviderKeyNil(b bool)`

 SetProviderKeyNil sets the value for ProviderKey to be an explicit nil

### UnsetProviderKey
`func (o *CognitiveAgentCreateDto) UnsetProviderKey()`

UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
### GetModelId

`func (o *CognitiveAgentCreateDto) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CognitiveAgentCreateDto) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CognitiveAgentCreateDto) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *CognitiveAgentCreateDto) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### SetModelIdNil

`func (o *CognitiveAgentCreateDto) SetModelIdNil(b bool)`

 SetModelIdNil sets the value for ModelId to be an explicit nil

### UnsetModelId
`func (o *CognitiveAgentCreateDto) UnsetModelId()`

UnsetModelId ensures that no value is present for ModelId, not even an explicit nil
### GetEngineKey

`func (o *CognitiveAgentCreateDto) GetEngineKey() string`

GetEngineKey returns the EngineKey field if non-nil, zero value otherwise.

### GetEngineKeyOk

`func (o *CognitiveAgentCreateDto) GetEngineKeyOk() (*string, bool)`

GetEngineKeyOk returns a tuple with the EngineKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineKey

`func (o *CognitiveAgentCreateDto) SetEngineKey(v string)`

SetEngineKey sets EngineKey field to given value.

### HasEngineKey

`func (o *CognitiveAgentCreateDto) HasEngineKey() bool`

HasEngineKey returns a boolean if a field has been set.

### SetEngineKeyNil

`func (o *CognitiveAgentCreateDto) SetEngineKeyNil(b bool)`

 SetEngineKeyNil sets the value for EngineKey to be an explicit nil

### UnsetEngineKey
`func (o *CognitiveAgentCreateDto) UnsetEngineKey()`

UnsetEngineKey ensures that no value is present for EngineKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


