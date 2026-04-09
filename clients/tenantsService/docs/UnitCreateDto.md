# UnitCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**BaseUnitAmount** | Pointer to **float64** |  | [optional] 
**BaseUnitId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUnitCreateDto

`func NewUnitCreateDto(name string, ) *UnitCreateDto`

NewUnitCreateDto instantiates a new UnitCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitCreateDtoWithDefaults

`func NewUnitCreateDtoWithDefaults() *UnitCreateDto`

NewUnitCreateDtoWithDefaults instantiates a new UnitCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnitCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnitCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnitCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UnitCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *UnitCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnitCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnitCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UnitCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *UnitCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnitCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnitCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetBaseUnitAmount

`func (o *UnitCreateDto) GetBaseUnitAmount() float64`

GetBaseUnitAmount returns the BaseUnitAmount field if non-nil, zero value otherwise.

### GetBaseUnitAmountOk

`func (o *UnitCreateDto) GetBaseUnitAmountOk() (*float64, bool)`

GetBaseUnitAmountOk returns a tuple with the BaseUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnitAmount

`func (o *UnitCreateDto) SetBaseUnitAmount(v float64)`

SetBaseUnitAmount sets BaseUnitAmount field to given value.

### HasBaseUnitAmount

`func (o *UnitCreateDto) HasBaseUnitAmount() bool`

HasBaseUnitAmount returns a boolean if a field has been set.

### GetBaseUnitId

`func (o *UnitCreateDto) GetBaseUnitId() string`

GetBaseUnitId returns the BaseUnitId field if non-nil, zero value otherwise.

### GetBaseUnitIdOk

`func (o *UnitCreateDto) GetBaseUnitIdOk() (*string, bool)`

GetBaseUnitIdOk returns a tuple with the BaseUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnitId

`func (o *UnitCreateDto) SetBaseUnitId(v string)`

SetBaseUnitId sets BaseUnitId field to given value.

### HasBaseUnitId

`func (o *UnitCreateDto) HasBaseUnitId() bool`

HasBaseUnitId returns a boolean if a field has been set.

### SetBaseUnitIdNil

`func (o *UnitCreateDto) SetBaseUnitIdNil(b bool)`

 SetBaseUnitIdNil sets the value for BaseUnitId to be an explicit nil

### UnsetBaseUnitId
`func (o *UnitCreateDto) UnsetBaseUnitId()`

UnsetBaseUnitId ensures that no value is present for BaseUnitId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


