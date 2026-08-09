# CognitiveSkillToolDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolKey** | **string** |  | 
**ConfigJson** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCognitiveSkillToolDto

`func NewCognitiveSkillToolDto(toolKey string, ) *CognitiveSkillToolDto`

NewCognitiveSkillToolDto instantiates a new CognitiveSkillToolDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveSkillToolDtoWithDefaults

`func NewCognitiveSkillToolDtoWithDefaults() *CognitiveSkillToolDto`

NewCognitiveSkillToolDtoWithDefaults instantiates a new CognitiveSkillToolDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolKey

`func (o *CognitiveSkillToolDto) GetToolKey() string`

GetToolKey returns the ToolKey field if non-nil, zero value otherwise.

### GetToolKeyOk

`func (o *CognitiveSkillToolDto) GetToolKeyOk() (*string, bool)`

GetToolKeyOk returns a tuple with the ToolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolKey

`func (o *CognitiveSkillToolDto) SetToolKey(v string)`

SetToolKey sets ToolKey field to given value.


### GetConfigJson

`func (o *CognitiveSkillToolDto) GetConfigJson() string`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *CognitiveSkillToolDto) GetConfigJsonOk() (*string, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *CognitiveSkillToolDto) SetConfigJson(v string)`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *CognitiveSkillToolDto) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.

### SetConfigJsonNil

`func (o *CognitiveSkillToolDto) SetConfigJsonNil(b bool)`

 SetConfigJsonNil sets the value for ConfigJson to be an explicit nil

### UnsetConfigJson
`func (o *CognitiveSkillToolDto) UnsetConfigJson()`

UnsetConfigJson ensures that no value is present for ConfigJson, not even an explicit nil
### GetEnabled

`func (o *CognitiveSkillToolDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CognitiveSkillToolDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CognitiveSkillToolDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CognitiveSkillToolDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


