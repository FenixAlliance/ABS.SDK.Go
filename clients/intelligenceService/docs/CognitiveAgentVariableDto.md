# CognitiveAgentVariableDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**CognitiveAgentId** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCognitiveAgentVariableDto

`func NewCognitiveAgentVariableDto() *CognitiveAgentVariableDto`

NewCognitiveAgentVariableDto instantiates a new CognitiveAgentVariableDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCognitiveAgentVariableDtoWithDefaults

`func NewCognitiveAgentVariableDtoWithDefaults() *CognitiveAgentVariableDto`

NewCognitiveAgentVariableDtoWithDefaults instantiates a new CognitiveAgentVariableDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CognitiveAgentVariableDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CognitiveAgentVariableDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CognitiveAgentVariableDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CognitiveAgentVariableDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CognitiveAgentVariableDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CognitiveAgentVariableDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CognitiveAgentVariableDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CognitiveAgentVariableDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CognitiveAgentVariableDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CognitiveAgentVariableDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CognitiveAgentVariableDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CognitiveAgentVariableDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCognitiveAgentId

`func (o *CognitiveAgentVariableDto) GetCognitiveAgentId() string`

GetCognitiveAgentId returns the CognitiveAgentId field if non-nil, zero value otherwise.

### GetCognitiveAgentIdOk

`func (o *CognitiveAgentVariableDto) GetCognitiveAgentIdOk() (*string, bool)`

GetCognitiveAgentIdOk returns a tuple with the CognitiveAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCognitiveAgentId

`func (o *CognitiveAgentVariableDto) SetCognitiveAgentId(v string)`

SetCognitiveAgentId sets CognitiveAgentId field to given value.

### HasCognitiveAgentId

`func (o *CognitiveAgentVariableDto) HasCognitiveAgentId() bool`

HasCognitiveAgentId returns a boolean if a field has been set.

### SetCognitiveAgentIdNil

`func (o *CognitiveAgentVariableDto) SetCognitiveAgentIdNil(b bool)`

 SetCognitiveAgentIdNil sets the value for CognitiveAgentId to be an explicit nil

### UnsetCognitiveAgentId
`func (o *CognitiveAgentVariableDto) UnsetCognitiveAgentId()`

UnsetCognitiveAgentId ensures that no value is present for CognitiveAgentId, not even an explicit nil
### GetKey

`func (o *CognitiveAgentVariableDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CognitiveAgentVariableDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CognitiveAgentVariableDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CognitiveAgentVariableDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CognitiveAgentVariableDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CognitiveAgentVariableDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *CognitiveAgentVariableDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CognitiveAgentVariableDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CognitiveAgentVariableDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CognitiveAgentVariableDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CognitiveAgentVariableDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CognitiveAgentVariableDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTenantId

`func (o *CognitiveAgentVariableDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CognitiveAgentVariableDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CognitiveAgentVariableDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CognitiveAgentVariableDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CognitiveAgentVariableDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CognitiveAgentVariableDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CognitiveAgentVariableDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CognitiveAgentVariableDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CognitiveAgentVariableDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CognitiveAgentVariableDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CognitiveAgentVariableDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CognitiveAgentVariableDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


