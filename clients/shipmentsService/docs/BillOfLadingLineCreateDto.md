# BillOfLadingLineCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**PackageType** | Pointer to **NullableString** |  | [optional] 
**GrossWeightKg** | Pointer to **float64** |  | [optional] 
**VolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**MarksAndNumbers** | Pointer to **NullableString** |  | [optional] 
**HsCode** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBillOfLadingLineCreateDto

`func NewBillOfLadingLineCreateDto() *BillOfLadingLineCreateDto`

NewBillOfLadingLineCreateDto instantiates a new BillOfLadingLineCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfLadingLineCreateDtoWithDefaults

`func NewBillOfLadingLineCreateDtoWithDefaults() *BillOfLadingLineCreateDto`

NewBillOfLadingLineCreateDtoWithDefaults instantiates a new BillOfLadingLineCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillOfLadingLineCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillOfLadingLineCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillOfLadingLineCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillOfLadingLineCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BillOfLadingLineCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BillOfLadingLineCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BillOfLadingLineCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BillOfLadingLineCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *BillOfLadingLineCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillOfLadingLineCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillOfLadingLineCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillOfLadingLineCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillOfLadingLineCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillOfLadingLineCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *BillOfLadingLineCreateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillOfLadingLineCreateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillOfLadingLineCreateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillOfLadingLineCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPackageType

`func (o *BillOfLadingLineCreateDto) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *BillOfLadingLineCreateDto) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *BillOfLadingLineCreateDto) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *BillOfLadingLineCreateDto) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### SetPackageTypeNil

`func (o *BillOfLadingLineCreateDto) SetPackageTypeNil(b bool)`

 SetPackageTypeNil sets the value for PackageType to be an explicit nil

### UnsetPackageType
`func (o *BillOfLadingLineCreateDto) UnsetPackageType()`

UnsetPackageType ensures that no value is present for PackageType, not even an explicit nil
### GetGrossWeightKg

`func (o *BillOfLadingLineCreateDto) GetGrossWeightKg() float64`

GetGrossWeightKg returns the GrossWeightKg field if non-nil, zero value otherwise.

### GetGrossWeightKgOk

`func (o *BillOfLadingLineCreateDto) GetGrossWeightKgOk() (*float64, bool)`

GetGrossWeightKgOk returns a tuple with the GrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossWeightKg

`func (o *BillOfLadingLineCreateDto) SetGrossWeightKg(v float64)`

SetGrossWeightKg sets GrossWeightKg field to given value.

### HasGrossWeightKg

`func (o *BillOfLadingLineCreateDto) HasGrossWeightKg() bool`

HasGrossWeightKg returns a boolean if a field has been set.

### GetVolumeM3

`func (o *BillOfLadingLineCreateDto) GetVolumeM3() float64`

GetVolumeM3 returns the VolumeM3 field if non-nil, zero value otherwise.

### GetVolumeM3Ok

`func (o *BillOfLadingLineCreateDto) GetVolumeM3Ok() (*float64, bool)`

GetVolumeM3Ok returns a tuple with the VolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeM3

`func (o *BillOfLadingLineCreateDto) SetVolumeM3(v float64)`

SetVolumeM3 sets VolumeM3 field to given value.

### HasVolumeM3

`func (o *BillOfLadingLineCreateDto) HasVolumeM3() bool`

HasVolumeM3 returns a boolean if a field has been set.

### SetVolumeM3Nil

`func (o *BillOfLadingLineCreateDto) SetVolumeM3Nil(b bool)`

 SetVolumeM3Nil sets the value for VolumeM3 to be an explicit nil

### UnsetVolumeM3
`func (o *BillOfLadingLineCreateDto) UnsetVolumeM3()`

UnsetVolumeM3 ensures that no value is present for VolumeM3, not even an explicit nil
### GetMarksAndNumbers

`func (o *BillOfLadingLineCreateDto) GetMarksAndNumbers() string`

GetMarksAndNumbers returns the MarksAndNumbers field if non-nil, zero value otherwise.

### GetMarksAndNumbersOk

`func (o *BillOfLadingLineCreateDto) GetMarksAndNumbersOk() (*string, bool)`

GetMarksAndNumbersOk returns a tuple with the MarksAndNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarksAndNumbers

`func (o *BillOfLadingLineCreateDto) SetMarksAndNumbers(v string)`

SetMarksAndNumbers sets MarksAndNumbers field to given value.

### HasMarksAndNumbers

`func (o *BillOfLadingLineCreateDto) HasMarksAndNumbers() bool`

HasMarksAndNumbers returns a boolean if a field has been set.

### SetMarksAndNumbersNil

`func (o *BillOfLadingLineCreateDto) SetMarksAndNumbersNil(b bool)`

 SetMarksAndNumbersNil sets the value for MarksAndNumbers to be an explicit nil

### UnsetMarksAndNumbers
`func (o *BillOfLadingLineCreateDto) UnsetMarksAndNumbers()`

UnsetMarksAndNumbers ensures that no value is present for MarksAndNumbers, not even an explicit nil
### GetHsCode

`func (o *BillOfLadingLineCreateDto) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *BillOfLadingLineCreateDto) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *BillOfLadingLineCreateDto) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *BillOfLadingLineCreateDto) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *BillOfLadingLineCreateDto) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *BillOfLadingLineCreateDto) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil
### GetItemId

`func (o *BillOfLadingLineCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *BillOfLadingLineCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *BillOfLadingLineCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *BillOfLadingLineCreateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *BillOfLadingLineCreateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *BillOfLadingLineCreateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


