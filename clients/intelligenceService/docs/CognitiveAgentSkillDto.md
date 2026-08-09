# CognitiveAgentSkillDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CognitiveAgentId** | Pointer to **NullableString** |  | [optional] 
**CognitiveSkillId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfigJson** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentSkillDto

`func NewCognitiveAgentSkillDto() *CognitiveAgentSkillDto`

NewCognitiveAgentSkillDto instantiates a new CognitiveAgentSkillDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentSkillDtoWithDefaults

`func NewCognitiveAgentSkillDtoWithDefaults() *CognitiveAgentSkillDto`

NewCognitiveAgentSkillDtoWithDefaults instantiates a new CognitiveAgentSkillDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveAgentSkillDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveAgentSkillDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveAgentSkillDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveAgentSkillDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CognitiveAgentSkillDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CognitiveAgentSkillDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CognitiveAgentSkillDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveAgentSkillDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveAgentSkillDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveAgentSkillDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CognitiveAgentSkillDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CognitiveAgentSkillDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCognitiveAgentId

`func (o *CognitiveAgentSkillDto) GetCognitiveAgentId() string`

GetCognitiveAgentId returns the CognitiveAgentId field if non-nil, zero value otherwise.

### GetCognitiveAgentIdOk

`func (o *CognitiveAgentSkillDto) GetCognitiveAgentIdOk() (*string, bool)`

GetCognitiveAgentIdOk returns a tuple with the CognitiveAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitiveAgentId

`func (o *CognitiveAgentSkillDto) SetCognitiveAgentId(v string)`

SetCognitiveAgentId sets CognitiveAgentId field to given value.

### HasCognitiveAgentId

`func (o *CognitiveAgentSkillDto) HasCognitiveAgentId() bool`

HasCognitiveAgentId returns a boolean if a field has been set.

### SetCognitiveAgentIdNil

`func (o *CognitiveAgentSkillDto) SetCognitiveAgentIdNil(b bool)`

 SetCognitiveAgentIdNil sets the value for CognitiveAgentId to be an explicit nil

### UnsetCognitiveAgentId
`func (o *CognitiveAgentSkillDto) UnsetCognitiveAgentId()`

UnsetCognitiveAgentId ensures that no value is present for CognitiveAgentId, not even an explicit nil
### GetCognitiveSkillId

`func (o *CognitiveAgentSkillDto) GetCognitiveSkillId() string`

GetCognitiveSkillId returns the CognitiveSkillId field if non-nil, zero value otherwise.

### GetCognitiveSkillIdOk

`func (o *CognitiveAgentSkillDto) GetCognitiveSkillIdOk() (*string, bool)`

GetCognitiveSkillIdOk returns a tuple with the CognitiveSkillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitiveSkillId

`func (o *CognitiveAgentSkillDto) SetCognitiveSkillId(v string)`

SetCognitiveSkillId sets CognitiveSkillId field to given value.

### HasCognitiveSkillId

`func (o *CognitiveAgentSkillDto) HasCognitiveSkillId() bool`

HasCognitiveSkillId returns a boolean if a field has been set.

### SetCognitiveSkillIdNil

`func (o *CognitiveAgentSkillDto) SetCognitiveSkillIdNil(b bool)`

 SetCognitiveSkillIdNil sets the value for CognitiveSkillId to be an explicit nil

### UnsetCognitiveSkillId
`func (o *CognitiveAgentSkillDto) UnsetCognitiveSkillId()`

UnsetCognitiveSkillId ensures that no value is present for CognitiveSkillId, not even an explicit nil
### GetEnabled

`func (o *CognitiveAgentSkillDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CognitiveAgentSkillDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CognitiveAgentSkillDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CognitiveAgentSkillDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfigJson

`func (o *CognitiveAgentSkillDto) GetConfigJson() string`

GetConfigJson returns the ConfigJson field if non-nil, zero value otherwise.

### GetConfigJsonOk

`func (o *CognitiveAgentSkillDto) GetConfigJsonOk() (*string, bool)`

GetConfigJsonOk returns a tuple with the ConfigJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJson

`func (o *CognitiveAgentSkillDto) SetConfigJson(v string)`

SetConfigJson sets ConfigJson field to given value.

### HasConfigJson

`func (o *CognitiveAgentSkillDto) HasConfigJson() bool`

HasConfigJson returns a boolean if a field has been set.

### SetConfigJsonNil

`func (o *CognitiveAgentSkillDto) SetConfigJsonNil(b bool)`

 SetConfigJsonNil sets the value for ConfigJson to be an explicit nil

### UnsetConfigJson
`func (o *CognitiveAgentSkillDto) UnsetConfigJson()`

UnsetConfigJson ensures that no value is present for ConfigJson, not even an explicit nil
### GetTenantId

`func (o *CognitiveAgentSkillDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CognitiveAgentSkillDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CognitiveAgentSkillDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CognitiveAgentSkillDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CognitiveAgentSkillDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CognitiveAgentSkillDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CognitiveAgentSkillDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CognitiveAgentSkillDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CognitiveAgentSkillDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CognitiveAgentSkillDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CognitiveAgentSkillDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CognitiveAgentSkillDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


