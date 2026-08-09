# CognitiveSkillUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ToolKey** | Pointer to **NullableString** |  | [optional] 
**ConfigJson** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Tools** | Pointer to [**[]CognitiveSkillToolDto**](CognitiveSkillToolDto.md) |  | [optional] 

## Methods

### NewCognitiveSkillUpdateDto

`func NewCognitiveSkillUpdateDto() *CognitiveSkillUpdateDto`

NewCognitiveSkillUpdateDto instantiates a new CognitiveSkillUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveSkillUpdateDtoWithDefaults

`func NewCognitiveSkillUpdateDtoWithDefaults() *CognitiveSkillUpdateDto`

NewCognitiveSkillUpdateDtoWithDefaults instantiates a new CognitiveSkillUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CognitiveSkillUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CognitiveSkillUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CognitiveSkillUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CognitiveSkillUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CognitiveSkillUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CognitiveSkillUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CognitiveSkillUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CognitiveSkillUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CognitiveSkillUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CognitiveSkillUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CognitiveSkillUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CognitiveSkillUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetToolKey

`func (o *CognitiveSkillUpdateDto) GetToolKey() string`

GetToolKey returns the ToolKey field if non-nil, zero value otherwise.

### GetToolKeyOk

`func (o *CognitiveSkillUpdateDto) GetToolKeyOk() (*string, bool)`

GetToolKeyOk returns a tuple with the ToolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolKey

`func (o *CognitiveSkillUpdateDto) SetToolKey(v string)`

SetToolKey sets ToolKey field to given value.

### HasToolKey

`func (o *CognitiveSkillUpdateDto) HasToolKey() bool`

HasToolKey returns a boolean if a field has been set.

### SetToolKeyNil

`func (o *CognitiveSkillUpdateDto) SetToolKeyNil(b bool)`

 SetToolKeyNil sets the value for ToolKey to be an explicit nil

### UnsetToolKey
`func (o *CognitiveSkillUpdateDto) UnsetToolKey()`

UnsetToolKey ensures that no value is present for ToolKey, not even an explicit nil
### GetConfigJson

`func (o *CognitiveSkillUpdateDto) GetConfigJson() string`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *CognitiveSkillUpdateDto) GetConfigJsonOk() (*string, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *CognitiveSkillUpdateDto) SetConfigJson(v string)`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *CognitiveSkillUpdateDto) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.

### SetConfigJsonNil

`func (o *CognitiveSkillUpdateDto) SetConfigJsonNil(b bool)`

 SetConfigJsonNil sets the value for ConfigJson to be an explicit nil

### UnsetConfigJson
`func (o *CognitiveSkillUpdateDto) UnsetConfigJson()`

UnsetConfigJson ensures that no value is present for ConfigJson, not even an explicit nil
### GetEnabled

`func (o *CognitiveSkillUpdateDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CognitiveSkillUpdateDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CognitiveSkillUpdateDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CognitiveSkillUpdateDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTools

`func (o *CognitiveSkillUpdateDto) GetTools() []CognitiveSkillToolDto`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CognitiveSkillUpdateDto) GetToolsOk() (*[]CognitiveSkillToolDto, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CognitiveSkillUpdateDto) SetTools(v []CognitiveSkillToolDto)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CognitiveSkillUpdateDto) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *CognitiveSkillUpdateDto) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *CognitiveSkillUpdateDto) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


