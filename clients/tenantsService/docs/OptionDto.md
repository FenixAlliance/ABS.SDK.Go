# OptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**PortalId** | Pointer to **NullableString** |  | [optional] 
**Frozen** | Pointer to **bool** |  | [optional] 
**Autoload** | Pointer to **bool** |  | [optional] 
**Transient** | Pointer to **bool** |  | [optional] 
**Expiration** | Pointer to **int32** |  | [optional] 

## Methods

### NewOptionDto

`func NewOptionDto() *OptionDto`

NewOptionDto instantiates a new OptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionDtoWithDefaults

`func NewOptionDtoWithDefaults() *OptionDto`

NewOptionDtoWithDefaults instantiates a new OptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OptionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OptionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OptionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *OptionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OptionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OptionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OptionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *OptionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *OptionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetKey

`func (o *OptionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OptionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OptionDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OptionDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *OptionDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *OptionDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *OptionDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OptionDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OptionDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OptionDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *OptionDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *OptionDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetPortalId

`func (o *OptionDto) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *OptionDto) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *OptionDto) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *OptionDto) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### SetPortalIdNil

`func (o *OptionDto) SetPortalIdNil(b bool)`

 SetPortalIdNil sets the value for PortalId to be an explicit nil

### UnsetPortalId
`func (o *OptionDto) UnsetPortalId()`

UnsetPortalId ensures that no value is present for PortalId, not even an explicit nil
### GetFrozen

`func (o *OptionDto) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *OptionDto) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *OptionDto) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *OptionDto) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetAutoload

`func (o *OptionDto) GetAutoload() bool`

GetAutoload returns the Autoload field if non-nil, zero value otherwise.

### GetAutoloadOk

`func (o *OptionDto) GetAutoloadOk() (*bool, bool)`

GetAutoloadOk returns a tuple with the Autoload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoload

`func (o *OptionDto) SetAutoload(v bool)`

SetAutoload sets Autoload field to given value.

### HasAutoload

`func (o *OptionDto) HasAutoload() bool`

HasAutoload returns a boolean if a field has been set.

### GetTransient

`func (o *OptionDto) GetTransient() bool`

GetTransient returns the Transient field if non-nil, zero value otherwise.

### GetTransientOk

`func (o *OptionDto) GetTransientOk() (*bool, bool)`

GetTransientOk returns a tuple with the Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransient

`func (o *OptionDto) SetTransient(v bool)`

SetTransient sets Transient field to given value.

### HasTransient

`func (o *OptionDto) HasTransient() bool`

HasTransient returns a boolean if a field has been set.

### GetExpiration

`func (o *OptionDto) GetExpiration() int32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *OptionDto) GetExpirationOk() (*int32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *OptionDto) SetExpiration(v int32)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *OptionDto) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


