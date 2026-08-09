# CognitiveSkillDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ToolKey** | Pointer to **NullableString** |  | [optional] 
**ConfigJson** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Tools** | Pointer to [**[]CognitiveSkillToolDto**](CognitiveSkillToolDto.md) |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveSkillDto

`func NewCognitiveSkillDto() *CognitiveSkillDto`

NewCognitiveSkillDto instantiates a new CognitiveSkillDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveSkillDtoWithDefaults

`func NewCognitiveSkillDtoWithDefaults() *CognitiveSkillDto`

NewCognitiveSkillDtoWithDefaults instantiates a new CognitiveSkillDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveSkillDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveSkillDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveSkillDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveSkillDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CognitiveSkillDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CognitiveSkillDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CognitiveSkillDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveSkillDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveSkillDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveSkillDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CognitiveSkillDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CognitiveSkillDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *CognitiveSkillDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CognitiveSkillDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CognitiveSkillDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CognitiveSkillDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CognitiveSkillDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CognitiveSkillDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CognitiveSkillDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CognitiveSkillDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CognitiveSkillDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CognitiveSkillDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CognitiveSkillDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CognitiveSkillDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetToolKey

`func (o *CognitiveSkillDto) GetToolKey() string`

GetToolKey returns the ToolKey field if non-nil, zero value otherwise.

### GetToolKeyOk

`func (o *CognitiveSkillDto) GetToolKeyOk() (*string, bool)`

GetToolKeyOk returns a tuple with the ToolKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolKey

`func (o *CognitiveSkillDto) SetToolKey(v string)`

SetToolKey sets ToolKey field to given value.

### HasToolKey

`func (o *CognitiveSkillDto) HasToolKey() bool`

HasToolKey returns a boolean if a field has been set.

### SetToolKeyNil

`func (o *CognitiveSkillDto) SetToolKeyNil(b bool)`

 SetToolKeyNil sets the value for ToolKey to be an explicit nil

### UnsetToolKey
`func (o *CognitiveSkillDto) UnsetToolKey()`

UnsetToolKey ensures that no value is present for ToolKey, not even an explicit nil
### GetConfigJson

`func (o *CognitiveSkillDto) GetConfigJson() string`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *CognitiveSkillDto) GetConfigJsonOk() (*string, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *CognitiveSkillDto) SetConfigJson(v string)`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *CognitiveSkillDto) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.

### SetConfigJsonNil

`func (o *CognitiveSkillDto) SetConfigJsonNil(b bool)`

 SetConfigJsonNil sets the value for ConfigJson to be an explicit nil

### UnsetConfigJson
`func (o *CognitiveSkillDto) UnsetConfigJson()`

UnsetConfigJson ensures that no value is present for ConfigJson, not even an explicit nil
### GetEnabled

`func (o *CognitiveSkillDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CognitiveSkillDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CognitiveSkillDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CognitiveSkillDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTools

`func (o *CognitiveSkillDto) GetTools() []CognitiveSkillToolDto`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CognitiveSkillDto) GetToolsOk() (*[]CognitiveSkillToolDto, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CognitiveSkillDto) SetTools(v []CognitiveSkillToolDto)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CognitiveSkillDto) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *CognitiveSkillDto) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *CognitiveSkillDto) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetTenantId

`func (o *CognitiveSkillDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CognitiveSkillDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CognitiveSkillDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CognitiveSkillDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CognitiveSkillDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CognitiveSkillDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CognitiveSkillDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CognitiveSkillDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CognitiveSkillDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CognitiveSkillDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CognitiveSkillDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CognitiveSkillDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


