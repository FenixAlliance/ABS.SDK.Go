# UnitDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**BaseUnitAmount** | Pointer to **float64** |  | [optional] 
**BaseUnitId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUnitDto

`func NewUnitDto() *UnitDto`

NewUnitDto instantiates a new UnitDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitDtoWithDefaults

`func NewUnitDtoWithDefaults() *UnitDto`

NewUnitDtoWithDefaults instantiates a new UnitDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnitDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnitDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnitDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnitDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UnitDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UnitDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *UnitDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnitDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnitDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UnitDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *UnitDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *UnitDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *UnitDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnitDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnitDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnitDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UnitDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UnitDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUnitGroupId

`func (o *UnitDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *UnitDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *UnitDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *UnitDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *UnitDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *UnitDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetBaseUnitAmount

`func (o *UnitDto) GetBaseUnitAmount() float64`

GetBaseUnitAmount returns the BaseUnitAmount field if non-nil, zero value otherwise.

### GetBaseUnitAmountOk

`func (o *UnitDto) GetBaseUnitAmountOk() (*float64, bool)`

GetBaseUnitAmountOk returns a tuple with the BaseUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnitAmount

`func (o *UnitDto) SetBaseUnitAmount(v float64)`

SetBaseUnitAmount sets BaseUnitAmount field to given value.

### HasBaseUnitAmount

`func (o *UnitDto) HasBaseUnitAmount() bool`

HasBaseUnitAmount returns a boolean if a field has been set.

### GetBaseUnitId

`func (o *UnitDto) GetBaseUnitId() string`

GetBaseUnitId returns the BaseUnitId field if non-nil, zero value otherwise.

### GetBaseUnitIdOk

`func (o *UnitDto) GetBaseUnitIdOk() (*string, bool)`

GetBaseUnitIdOk returns a tuple with the BaseUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnitId

`func (o *UnitDto) SetBaseUnitId(v string)`

SetBaseUnitId sets BaseUnitId field to given value.

### HasBaseUnitId

`func (o *UnitDto) HasBaseUnitId() bool`

HasBaseUnitId returns a boolean if a field has been set.

### SetBaseUnitIdNil

`func (o *UnitDto) SetBaseUnitIdNil(b bool)`

 SetBaseUnitIdNil sets the value for BaseUnitId to be an explicit nil

### UnsetBaseUnitId
`func (o *UnitDto) UnsetBaseUnitId()`

UnsetBaseUnitId ensures that no value is present for BaseUnitId, not even an explicit nil
### GetTenantId

`func (o *UnitDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UnitDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UnitDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UnitDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *UnitDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *UnitDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *UnitDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *UnitDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *UnitDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *UnitDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *UnitDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *UnitDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


