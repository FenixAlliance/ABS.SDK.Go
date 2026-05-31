# WarehouseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Address1** | Pointer to **NullableString** |  | [optional] 
**Address2** | Pointer to **NullableString** |  | [optional] 
**Address3** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**CountryId** | Pointer to **NullableString** |  | [optional] 
**StateId** | Pointer to **NullableString** |  | [optional] 
**CityId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**IsGroup** | Pointer to **bool** |  | [optional] 
**ShipwireWarehouseId** | Pointer to **NullableString** |  | [optional] 
**ParentWarehouseId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWarehouseDto

`func NewWarehouseDto() *WarehouseDto`

NewWarehouseDto instantiates a new WarehouseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseDtoWithDefaults

`func NewWarehouseDtoWithDefaults() *WarehouseDto`

NewWarehouseDtoWithDefaults instantiates a new WarehouseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WarehouseDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WarehouseDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WarehouseDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WarehouseDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WarehouseDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WarehouseDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WarehouseDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WarehouseDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WarehouseDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *WarehouseDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WarehouseDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WarehouseDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WarehouseDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WarehouseDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WarehouseDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAddress1

`func (o *WarehouseDto) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *WarehouseDto) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *WarehouseDto) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *WarehouseDto) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *WarehouseDto) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *WarehouseDto) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *WarehouseDto) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *WarehouseDto) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *WarehouseDto) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *WarehouseDto) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *WarehouseDto) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *WarehouseDto) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetAddress3

`func (o *WarehouseDto) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *WarehouseDto) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *WarehouseDto) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *WarehouseDto) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### SetAddress3Nil

`func (o *WarehouseDto) SetAddress3Nil(b bool)`

 SetAddress3Nil sets the value for Address3 to be an explicit nil

### UnsetAddress3
`func (o *WarehouseDto) UnsetAddress3()`

UnsetAddress3 ensures that no value is present for Address3, not even an explicit nil
### GetPostalCode

`func (o *WarehouseDto) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *WarehouseDto) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *WarehouseDto) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *WarehouseDto) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *WarehouseDto) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *WarehouseDto) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetPhone

`func (o *WarehouseDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *WarehouseDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *WarehouseDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *WarehouseDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *WarehouseDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *WarehouseDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCountryId

`func (o *WarehouseDto) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *WarehouseDto) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *WarehouseDto) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *WarehouseDto) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### SetCountryIdNil

`func (o *WarehouseDto) SetCountryIdNil(b bool)`

 SetCountryIdNil sets the value for CountryId to be an explicit nil

### UnsetCountryId
`func (o *WarehouseDto) UnsetCountryId()`

UnsetCountryId ensures that no value is present for CountryId, not even an explicit nil
### GetStateId

`func (o *WarehouseDto) GetStateId() string`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *WarehouseDto) GetStateIdOk() (*string, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *WarehouseDto) SetStateId(v string)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *WarehouseDto) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### SetStateIdNil

`func (o *WarehouseDto) SetStateIdNil(b bool)`

 SetStateIdNil sets the value for StateId to be an explicit nil

### UnsetStateId
`func (o *WarehouseDto) UnsetStateId()`

UnsetStateId ensures that no value is present for StateId, not even an explicit nil
### GetCityId

`func (o *WarehouseDto) GetCityId() string`

GetCityId returns the CityId field if non-nil, zero value otherwise.

### GetCityIdOk

`func (o *WarehouseDto) GetCityIdOk() (*string, bool)`

GetCityIdOk returns a tuple with the CityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityId

`func (o *WarehouseDto) SetCityId(v string)`

SetCityId sets CityId field to given value.

### HasCityId

`func (o *WarehouseDto) HasCityId() bool`

HasCityId returns a boolean if a field has been set.

### SetCityIdNil

`func (o *WarehouseDto) SetCityIdNil(b bool)`

 SetCityIdNil sets the value for CityId to be an explicit nil

### UnsetCityId
`func (o *WarehouseDto) UnsetCityId()`

UnsetCityId ensures that no value is present for CityId, not even an explicit nil
### GetTenantId

`func (o *WarehouseDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WarehouseDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WarehouseDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *WarehouseDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *WarehouseDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *WarehouseDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetIsGroup

`func (o *WarehouseDto) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *WarehouseDto) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *WarehouseDto) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *WarehouseDto) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.

### GetShipwireWarehouseId

`func (o *WarehouseDto) GetShipwireWarehouseId() string`

GetShipwireWarehouseId returns the ShipwireWarehouseId field if non-nil, zero value otherwise.

### GetShipwireWarehouseIdOk

`func (o *WarehouseDto) GetShipwireWarehouseIdOk() (*string, bool)`

GetShipwireWarehouseIdOk returns a tuple with the ShipwireWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipwireWarehouseId

`func (o *WarehouseDto) SetShipwireWarehouseId(v string)`

SetShipwireWarehouseId sets ShipwireWarehouseId field to given value.

### HasShipwireWarehouseId

`func (o *WarehouseDto) HasShipwireWarehouseId() bool`

HasShipwireWarehouseId returns a boolean if a field has been set.

### SetShipwireWarehouseIdNil

`func (o *WarehouseDto) SetShipwireWarehouseIdNil(b bool)`

 SetShipwireWarehouseIdNil sets the value for ShipwireWarehouseId to be an explicit nil

### UnsetShipwireWarehouseId
`func (o *WarehouseDto) UnsetShipwireWarehouseId()`

UnsetShipwireWarehouseId ensures that no value is present for ShipwireWarehouseId, not even an explicit nil
### GetParentWarehouseId

`func (o *WarehouseDto) GetParentWarehouseId() string`

GetParentWarehouseId returns the ParentWarehouseId field if non-nil, zero value otherwise.

### GetParentWarehouseIdOk

`func (o *WarehouseDto) GetParentWarehouseIdOk() (*string, bool)`

GetParentWarehouseIdOk returns a tuple with the ParentWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWarehouseId

`func (o *WarehouseDto) SetParentWarehouseId(v string)`

SetParentWarehouseId sets ParentWarehouseId field to given value.

### HasParentWarehouseId

`func (o *WarehouseDto) HasParentWarehouseId() bool`

HasParentWarehouseId returns a boolean if a field has been set.

### SetParentWarehouseIdNil

`func (o *WarehouseDto) SetParentWarehouseIdNil(b bool)`

 SetParentWarehouseIdNil sets the value for ParentWarehouseId to be an explicit nil

### UnsetParentWarehouseId
`func (o *WarehouseDto) UnsetParentWarehouseId()`

UnsetParentWarehouseId ensures that no value is present for ParentWarehouseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


