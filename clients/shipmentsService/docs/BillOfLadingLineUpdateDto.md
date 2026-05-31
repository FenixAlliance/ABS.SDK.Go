# BillOfLadingLineUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**PackageType** | Pointer to **NullableString** |  | [optional] 
**GrossWeightKg** | Pointer to **float64** |  | [optional] 
**VolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**MarksAndNumbers** | Pointer to **NullableString** |  | [optional] 
**HsCode** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBillOfLadingLineUpdateDto

`func NewBillOfLadingLineUpdateDto() *BillOfLadingLineUpdateDto`

NewBillOfLadingLineUpdateDto instantiates a new BillOfLadingLineUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfLadingLineUpdateDtoWithDefaults

`func NewBillOfLadingLineUpdateDtoWithDefaults() *BillOfLadingLineUpdateDto`

NewBillOfLadingLineUpdateDtoWithDefaults instantiates a new BillOfLadingLineUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BillOfLadingLineUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillOfLadingLineUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillOfLadingLineUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillOfLadingLineUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillOfLadingLineUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillOfLadingLineUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *BillOfLadingLineUpdateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillOfLadingLineUpdateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillOfLadingLineUpdateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillOfLadingLineUpdateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPackageType

`func (o *BillOfLadingLineUpdateDto) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *BillOfLadingLineUpdateDto) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *BillOfLadingLineUpdateDto) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *BillOfLadingLineUpdateDto) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### SetPackageTypeNil

`func (o *BillOfLadingLineUpdateDto) SetPackageTypeNil(b bool)`

 SetPackageTypeNil sets the value for PackageType to be an explicit nil

### UnsetPackageType
`func (o *BillOfLadingLineUpdateDto) UnsetPackageType()`

UnsetPackageType ensures that no value is present for PackageType, not even an explicit nil
### GetGrossWeightKg

`func (o *BillOfLadingLineUpdateDto) GetGrossWeightKg() float64`

GetGrossWeightKg returns the GrossWeightKg field if non-nil, zero value otherwise.

### GetGrossWeightKgOk

`func (o *BillOfLadingLineUpdateDto) GetGrossWeightKgOk() (*float64, bool)`

GetGrossWeightKgOk returns a tuple with the GrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossWeightKg

`func (o *BillOfLadingLineUpdateDto) SetGrossWeightKg(v float64)`

SetGrossWeightKg sets GrossWeightKg field to given value.

### HasGrossWeightKg

`func (o *BillOfLadingLineUpdateDto) HasGrossWeightKg() bool`

HasGrossWeightKg returns a boolean if a field has been set.

### GetVolumeM3

`func (o *BillOfLadingLineUpdateDto) GetVolumeM3() float64`

GetVolumeM3 returns the VolumeM3 field if non-nil, zero value otherwise.

### GetVolumeM3Ok

`func (o *BillOfLadingLineUpdateDto) GetVolumeM3Ok() (*float64, bool)`

GetVolumeM3Ok returns a tuple with the VolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeM3

`func (o *BillOfLadingLineUpdateDto) SetVolumeM3(v float64)`

SetVolumeM3 sets VolumeM3 field to given value.

### HasVolumeM3

`func (o *BillOfLadingLineUpdateDto) HasVolumeM3() bool`

HasVolumeM3 returns a boolean if a field has been set.

### SetVolumeM3Nil

`func (o *BillOfLadingLineUpdateDto) SetVolumeM3Nil(b bool)`

 SetVolumeM3Nil sets the value for VolumeM3 to be an explicit nil

### UnsetVolumeM3
`func (o *BillOfLadingLineUpdateDto) UnsetVolumeM3()`

UnsetVolumeM3 ensures that no value is present for VolumeM3, not even an explicit nil
### GetMarksAndNumbers

`func (o *BillOfLadingLineUpdateDto) GetMarksAndNumbers() string`

GetMarksAndNumbers returns the MarksAndNumbers field if non-nil, zero value otherwise.

### GetMarksAndNumbersOk

`func (o *BillOfLadingLineUpdateDto) GetMarksAndNumbersOk() (*string, bool)`

GetMarksAndNumbersOk returns a tuple with the MarksAndNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarksAndNumbers

`func (o *BillOfLadingLineUpdateDto) SetMarksAndNumbers(v string)`

SetMarksAndNumbers sets MarksAndNumbers field to given value.

### HasMarksAndNumbers

`func (o *BillOfLadingLineUpdateDto) HasMarksAndNumbers() bool`

HasMarksAndNumbers returns a boolean if a field has been set.

### SetMarksAndNumbersNil

`func (o *BillOfLadingLineUpdateDto) SetMarksAndNumbersNil(b bool)`

 SetMarksAndNumbersNil sets the value for MarksAndNumbers to be an explicit nil

### UnsetMarksAndNumbers
`func (o *BillOfLadingLineUpdateDto) UnsetMarksAndNumbers()`

UnsetMarksAndNumbers ensures that no value is present for MarksAndNumbers, not even an explicit nil
### GetHsCode

`func (o *BillOfLadingLineUpdateDto) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *BillOfLadingLineUpdateDto) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *BillOfLadingLineUpdateDto) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *BillOfLadingLineUpdateDto) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *BillOfLadingLineUpdateDto) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *BillOfLadingLineUpdateDto) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil
### GetItemId

`func (o *BillOfLadingLineUpdateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *BillOfLadingLineUpdateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *BillOfLadingLineUpdateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *BillOfLadingLineUpdateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *BillOfLadingLineUpdateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *BillOfLadingLineUpdateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


