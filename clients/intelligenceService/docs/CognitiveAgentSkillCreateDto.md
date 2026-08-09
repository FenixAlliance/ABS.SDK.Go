# CognitiveAgentSkillCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CognitiveSkillId** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigJson** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentSkillCreateDto

`func NewCognitiveAgentSkillCreateDto(cognitiveSkillId string, ) *CognitiveAgentSkillCreateDto`

NewCognitiveAgentSkillCreateDto instantiates a new CognitiveAgentSkillCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentSkillCreateDtoWithDefaults

`func NewCognitiveAgentSkillCreateDtoWithDefaults() *CognitiveAgentSkillCreateDto`

NewCognitiveAgentSkillCreateDtoWithDefaults instantiates a new CognitiveAgentSkillCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveAgentSkillCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveAgentSkillCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveAgentSkillCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveAgentSkillCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CognitiveAgentSkillCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveAgentSkillCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveAgentSkillCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveAgentSkillCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCognitiveSkillId

`func (o *CognitiveAgentSkillCreateDto) GetCognitiveSkillId() string`

GetCognitiveSkillId returns the CognitiveSkillId field if non-nil, zero value otherwise.

### GetCognitiveSkillIdOk

`func (o *CognitiveAgentSkillCreateDto) GetCognitiveSkillIdOk() (*string, bool)`

GetCognitiveSkillIdOk returns a tuple with the CognitiveSkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitiveSkillId

`func (o *CognitiveAgentSkillCreateDto) SetCognitiveSkillId(v string)`

SetCognitiveSkillId sets CognitiveSkillId field to given value.


### GetEnabled

`func (o *CognitiveAgentSkillCreateDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CognitiveAgentSkillCreateDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CognitiveAgentSkillCreateDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CognitiveAgentSkillCreateDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigJson

`func (o *CognitiveAgentSkillCreateDto) GetConfigJson() string`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *CognitiveAgentSkillCreateDto) GetConfigJsonOk() (*string, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *CognitiveAgentSkillCreateDto) SetConfigJson(v string)`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *CognitiveAgentSkillCreateDto) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.

### SetConfigJsonNil

`func (o *CognitiveAgentSkillCreateDto) SetConfigJsonNil(b bool)`

 SetConfigJsonNil sets the value for ConfigJson to be an explicit nil

### UnsetConfigJson
`func (o *CognitiveAgentSkillCreateDto) UnsetConfigJson()`

UnsetConfigJson ensures that no value is present for ConfigJson, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


