# WarehouseCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | **string** |  | 
**Address1** | **string** |  | 
**Address2** | Pointer to **NullableString** |  | [optional] 
**Address3** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**IsGroup** | Pointer to **bool** |  | [optional] 
**ShipwireWarehouseId** | Pointer to **NullableString** |  | [optional] 
**ParentWarehouseId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWarehouseCreateDto

`func NewWarehouseCreateDto(title string, address1 string, ) *WarehouseCreateDto`

NewWarehouseCreateDto instantiates a new WarehouseCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseCreateDtoWithDefaults

`func NewWarehouseCreateDtoWithDefaults() *WarehouseCreateDto`

NewWarehouseCreateDtoWithDefaults instantiates a new WarehouseCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WarehouseCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WarehouseCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WarehouseCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WarehouseCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WarehouseCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *WarehouseCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WarehouseCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WarehouseCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAddress1

`func (o *WarehouseCreateDto) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *WarehouseCreateDto) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *WarehouseCreateDto) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *WarehouseCreateDto) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *WarehouseCreateDto) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *WarehouseCreateDto) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *WarehouseCreateDto) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *WarehouseCreateDto) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *WarehouseCreateDto) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetAddress3

`func (o *WarehouseCreateDto) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *WarehouseCreateDto) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *WarehouseCreateDto) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *WarehouseCreateDto) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### SetAddress3Nil

`func (o *WarehouseCreateDto) SetAddress3Nil(b bool)`

 SetAddress3Nil sets the value for Address3 to be an explicit nil

### UnsetAddress3
`func (o *WarehouseCreateDto) UnsetAddress3()`

UnsetAddress3 ensures that no value is present for Address3, not even an explicit nil
### GetPostalCode

`func (o *WarehouseCreateDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *WarehouseCreateDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *WarehouseCreateDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *WarehouseCreateDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *WarehouseCreateDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *WarehouseCreateDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPhone

`func (o *WarehouseCreateDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *WarehouseCreateDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *WarehouseCreateDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *WarehouseCreateDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *WarehouseCreateDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *WarehouseCreateDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCountryId

`func (o *WarehouseCreateDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *WarehouseCreateDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *WarehouseCreateDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *WarehouseCreateDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *WarehouseCreateDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *WarehouseCreateDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *WarehouseCreateDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *WarehouseCreateDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *WarehouseCreateDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *WarehouseCreateDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *WarehouseCreateDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *WarehouseCreateDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *WarehouseCreateDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *WarehouseCreateDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *WarehouseCreateDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *WarehouseCreateDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *WarehouseCreateDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *WarehouseCreateDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetIsGroup

`func (o *WarehouseCreateDto) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *WarehouseCreateDto) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *WarehouseCreateDto) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *WarehouseCreateDto) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.

### GetShipwireWarehouseId

`func (o *WarehouseCreateDto) GetShipwireWarehouseId() string`

GetShipwireWarehouseId returns the ShipwireWarehouseId field if non-nil, zero value otherwise.

### GetShipwireWarehouseIdOk

`func (o *WarehouseCreateDto) GetShipwireWarehouseIdOk() (*string, bool)`

GetShipwireWarehouseIdOk returns a tuple with the ShipwireWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipwireWarehouseId

`func (o *WarehouseCreateDto) SetShipwireWarehouseId(v string)`

SetShipwireWarehouseId sets ShipwireWarehouseId field to given value.

### HasShipwireWarehouseId

`func (o *WarehouseCreateDto) HasShipwireWarehouseId() bool`

HasShipwireWarehouseId returns a boolean if a field has been set.

### SetShipwireWarehouseIdNil

`func (o *WarehouseCreateDto) SetShipwireWarehouseIdNil(b bool)`

 SetShipwireWarehouseIdNil sets the value for ShipwireWarehouseId to be an explicit nil

### UnsetShipwireWarehouseId
`func (o *WarehouseCreateDto) UnsetShipwireWarehouseId()`

UnsetShipwireWarehouseId ensures that no value is present for ShipwireWarehouseId, not even an explicit nil
### GetParentWarehouseId

`func (o *WarehouseCreateDto) GetParentWarehouseId() string`

GetParentWarehouseId returns the ParentWarehouseId field if non-nil, zero value otherwise.

### GetParentWarehouseIdOk

`func (o *WarehouseCreateDto) GetParentWarehouseIdOk() (*string, bool)`

GetParentWarehouseIdOk returns a tuple with the ParentWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWarehouseId

`func (o *WarehouseCreateDto) SetParentWarehouseId(v string)`

SetParentWarehouseId sets ParentWarehouseId field to given value.

### HasParentWarehouseId

`func (o *WarehouseCreateDto) HasParentWarehouseId() bool`

HasParentWarehouseId returns a boolean if a field has been set.

### SetParentWarehouseIdNil

`func (o *WarehouseCreateDto) SetParentWarehouseIdNil(b bool)`

 SetParentWarehouseIdNil sets the value for ParentWarehouseId to be an explicit nil

### UnsetParentWarehouseId
`func (o *WarehouseCreateDto) UnsetParentWarehouseId()`

UnsetParentWarehouseId ensures that no value is present for ParentWarehouseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


