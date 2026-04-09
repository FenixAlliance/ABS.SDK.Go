# OptionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Key** | **string** |  | 
**Value** | **string** |  | 
**PortalId** | Pointer to **NullableString** |  | [optional] 
**Frozen** | Pointer to **bool** |  | [optional] 
**Autoload** | Pointer to **bool** |  | [optional] 
**Transient** | Pointer to **bool** |  | [optional] 
**Expiration** | Pointer to **int32** |  | [optional] 

## Methods

### NewOptionCreateDto

`func NewOptionCreateDto(key string, value string, ) *OptionCreateDto`

NewOptionCreateDto instantiates a new OptionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionCreateDtoWithDefaults

`func NewOptionCreateDtoWithDefaults() *OptionCreateDto`

NewOptionCreateDtoWithDefaults instantiates a new OptionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OptionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *OptionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OptionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OptionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OptionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetKey

`func (o *OptionCreateDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OptionCreateDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OptionCreateDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *OptionCreateDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OptionCreateDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OptionCreateDto) SetValue(v string)`

SetValue sets Value field to given value.


### GetPortalId

`func (o *OptionCreateDto) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *OptionCreateDto) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *OptionCreateDto) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *OptionCreateDto) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### SetPortalIdNil

`func (o *OptionCreateDto) SetPortalIdNil(b bool)`

 SetPortalIdNil sets the value for PortalId to be an explicit nil

### UnsetPortalId
`func (o *OptionCreateDto) UnsetPortalId()`

UnsetPortalId ensures that no value is present for PortalId, not even an explicit nil
### GetFrozen

`func (o *OptionCreateDto) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *OptionCreateDto) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *OptionCreateDto) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *OptionCreateDto) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetAutoload

`func (o *OptionCreateDto) GetAutoload() bool`

GetAutoload returns the Autoload field if non-nil, zero value otherwise.

### GetAutoloadOk

`func (o *OptionCreateDto) GetAutoloadOk() (*bool, bool)`

GetAutoloadOk returns a tuple with the Autoload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoload

`func (o *OptionCreateDto) SetAutoload(v bool)`

SetAutoload sets Autoload field to given value.

### HasAutoload

`func (o *OptionCreateDto) HasAutoload() bool`

HasAutoload returns a boolean if a field has been set.

### GetTransient

`func (o *OptionCreateDto) GetTransient() bool`

GetTransient returns the Transient field if non-nil, zero value otherwise.

### GetTransientOk

`func (o *OptionCreateDto) GetTransientOk() (*bool, bool)`

GetTransientOk returns a tuple with the Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransient

`func (o *OptionCreateDto) SetTransient(v bool)`

SetTransient sets Transient field to given value.

### HasTransient

`func (o *OptionCreateDto) HasTransient() bool`

HasTransient returns a boolean if a field has been set.

### GetExpiration

`func (o *OptionCreateDto) GetExpiration() int32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *OptionCreateDto) GetExpirationOk() (*int32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *OptionCreateDto) SetExpiration(v int32)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *OptionCreateDto) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


