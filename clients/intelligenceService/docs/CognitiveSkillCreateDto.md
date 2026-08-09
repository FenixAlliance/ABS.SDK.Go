# CognitiveSkillCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ToolKey** | Pointer to **NullableString** |  | [optional] 
**ConfigJson** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Tools** | Pointer to [**[]CognitiveSkillToolDto**](CognitiveSkillToolDto.md) |  | [optional] 

## Methods

### NewCognitiveSkillCreateDto

`func NewCognitiveSkillCreateDto(name string, ) *CognitiveSkillCreateDto`

NewCognitiveSkillCreateDto instantiates a new CognitiveSkillCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveSkillCreateDtoWithDefaults

`func NewCognitiveSkillCreateDtoWithDefaults() *CognitiveSkillCreateDto`

NewCognitiveSkillCreateDtoWithDefaults instantiates a new CognitiveSkillCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveSkillCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveSkillCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveSkillCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveSkillCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CognitiveSkillCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveSkillCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveSkillCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveSkillCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *CognitiveSkillCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CognitiveSkillCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CognitiveSkillCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CognitiveSkillCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CognitiveSkillCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CognitiveSkillCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CognitiveSkillCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CognitiveSkillCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CognitiveSkillCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetToolKey

`func (o *CognitiveSkillCreateDto) GetToolKey() string`

GetToolKey returns the ToolKey field if non-nil, zero value otherwise.

### GetToolKeyOk

`func (o *CognitiveSkillCreateDto) GetToolKeyOk() (*string, bool)`

GetToolKeyOk returns a tuple with the ToolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolKey

`func (o *CognitiveSkillCreateDto) SetToolKey(v string)`

SetToolKey sets ToolKey field to given value.

### HasToolKey

`func (o *CognitiveSkillCreateDto) HasToolKey() bool`

HasToolKey returns a boolean if a field has been set.

### SetToolKeyNil

`func (o *CognitiveSkillCreateDto) SetToolKeyNil(b bool)`

 SetToolKeyNil sets the value for ToolKey to be an explicit nil

### UnsetToolKey
`func (o *CognitiveSkillCreateDto) UnsetToolKey()`

UnsetToolKey ensures that no value is present for ToolKey, not even an explicit nil
### GetConfigJson

`func (o *CognitiveSkillCreateDto) GetConfigJson() string`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *CognitiveSkillCreateDto) GetConfigJsonOk() (*string, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *CognitiveSkillCreateDto) SetConfigJson(v string)`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *CognitiveSkillCreateDto) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.

### SetConfigJsonNil

`func (o *CognitiveSkillCreateDto) SetConfigJsonNil(b bool)`

 SetConfigJsonNil sets the value for ConfigJson to be an explicit nil

### UnsetConfigJson
`func (o *CognitiveSkillCreateDto) UnsetConfigJson()`

UnsetConfigJson ensures that no value is present for ConfigJson, not even an explicit nil
### GetEnabled

`func (o *CognitiveSkillCreateDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CognitiveSkillCreateDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CognitiveSkillCreateDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CognitiveSkillCreateDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTools

`func (o *CognitiveSkillCreateDto) GetTools() []CognitiveSkillToolDto`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CognitiveSkillCreateDto) GetToolsOk() (*[]CognitiveSkillToolDto, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CognitiveSkillCreateDto) SetTools(v []CognitiveSkillToolDto)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CognitiveSkillCreateDto) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *CognitiveSkillCreateDto) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *CognitiveSkillCreateDto) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


