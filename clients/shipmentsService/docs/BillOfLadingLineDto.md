# BillOfLadingLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**PackageType** | Pointer to **NullableString** |  | [optional] 
**GrossWeightKg** | Pointer to **float64** |  | [optional] 
**VolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**MarksAndNumbers** | Pointer to **NullableString** |  | [optional] 
**HsCode** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**BillOfLadingId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBillOfLadingLineDto

`func NewBillOfLadingLineDto() *BillOfLadingLineDto`

NewBillOfLadingLineDto instantiates a new BillOfLadingLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfLadingLineDtoWithDefaults

`func NewBillOfLadingLineDtoWithDefaults() *BillOfLadingLineDto`

NewBillOfLadingLineDtoWithDefaults instantiates a new BillOfLadingLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillOfLadingLineDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillOfLadingLineDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillOfLadingLineDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillOfLadingLineDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BillOfLadingLineDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BillOfLadingLineDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BillOfLadingLineDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BillOfLadingLineDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BillOfLadingLineDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BillOfLadingLineDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BillOfLadingLineDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BillOfLadingLineDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDescription

`func (o *BillOfLadingLineDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillOfLadingLineDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillOfLadingLineDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillOfLadingLineDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillOfLadingLineDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillOfLadingLineDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *BillOfLadingLineDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillOfLadingLineDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillOfLadingLineDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillOfLadingLineDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPackageType

`func (o *BillOfLadingLineDto) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *BillOfLadingLineDto) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *BillOfLadingLineDto) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *BillOfLadingLineDto) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### SetPackageTypeNil

`func (o *BillOfLadingLineDto) SetPackageTypeNil(b bool)`

 SetPackageTypeNil sets the value for PackageType to be an explicit nil

### UnsetPackageType
`func (o *BillOfLadingLineDto) UnsetPackageType()`

UnsetPackageType ensures that no value is present for PackageType, not even an explicit nil
### GetGrossWeightKg

`func (o *BillOfLadingLineDto) GetGrossWeightKg() float64`

GetGrossWeightKg returns the GrossWeightKg field if non-nil, zero value otherwise.

### GetGrossWeightKgOk

`func (o *BillOfLadingLineDto) GetGrossWeightKgOk() (*float64, bool)`

GetGrossWeightKgOk returns a tuple with the GrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossWeightKg

`func (o *BillOfLadingLineDto) SetGrossWeightKg(v float64)`

SetGrossWeightKg sets GrossWeightKg field to given value.

### HasGrossWeightKg

`func (o *BillOfLadingLineDto) HasGrossWeightKg() bool`

HasGrossWeightKg returns a boolean if a field has been set.

### GetVolumeM3

`func (o *BillOfLadingLineDto) GetVolumeM3() float64`

GetVolumeM3 returns the VolumeM3 field if non-nil, zero value otherwise.

### GetVolumeM3Ok

`func (o *BillOfLadingLineDto) GetVolumeM3Ok() (*float64, bool)`

GetVolumeM3Ok returns a tuple with the VolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeM3

`func (o *BillOfLadingLineDto) SetVolumeM3(v float64)`

SetVolumeM3 sets VolumeM3 field to given value.

### HasVolumeM3

`func (o *BillOfLadingLineDto) HasVolumeM3() bool`

HasVolumeM3 returns a boolean if a field has been set.

### SetVolumeM3Nil

`func (o *BillOfLadingLineDto) SetVolumeM3Nil(b bool)`

 SetVolumeM3Nil sets the value for VolumeM3 to be an explicit nil

### UnsetVolumeM3
`func (o *BillOfLadingLineDto) UnsetVolumeM3()`

UnsetVolumeM3 ensures that no value is present for VolumeM3, not even an explicit nil
### GetMarksAndNumbers

`func (o *BillOfLadingLineDto) GetMarksAndNumbers() string`

GetMarksAndNumbers returns the MarksAndNumbers field if non-nil, zero value otherwise.

### GetMarksAndNumbersOk

`func (o *BillOfLadingLineDto) GetMarksAndNumbersOk() (*string, bool)`

GetMarksAndNumbersOk returns a tuple with the MarksAndNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarksAndNumbers

`func (o *BillOfLadingLineDto) SetMarksAndNumbers(v string)`

SetMarksAndNumbers sets MarksAndNumbers field to given value.

### HasMarksAndNumbers

`func (o *BillOfLadingLineDto) HasMarksAndNumbers() bool`

HasMarksAndNumbers returns a boolean if a field has been set.

### SetMarksAndNumbersNil

`func (o *BillOfLadingLineDto) SetMarksAndNumbersNil(b bool)`

 SetMarksAndNumbersNil sets the value for MarksAndNumbers to be an explicit nil

### UnsetMarksAndNumbers
`func (o *BillOfLadingLineDto) UnsetMarksAndNumbers()`

UnsetMarksAndNumbers ensures that no value is present for MarksAndNumbers, not even an explicit nil
### GetHsCode

`func (o *BillOfLadingLineDto) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *BillOfLadingLineDto) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *BillOfLadingLineDto) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *BillOfLadingLineDto) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *BillOfLadingLineDto) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *BillOfLadingLineDto) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil
### GetItemId

`func (o *BillOfLadingLineDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *BillOfLadingLineDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *BillOfLadingLineDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *BillOfLadingLineDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *BillOfLadingLineDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *BillOfLadingLineDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetBillOfLadingId

`func (o *BillOfLadingLineDto) GetBillOfLadingId() string`

GetBillOfLadingId returns the BillOfLadingId field if non-nil, zero value otherwise.

### GetBillOfLadingIdOk

`func (o *BillOfLadingLineDto) GetBillOfLadingIdOk() (*string, bool)`

GetBillOfLadingIdOk returns a tuple with the BillOfLadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingId

`func (o *BillOfLadingLineDto) SetBillOfLadingId(v string)`

SetBillOfLadingId sets BillOfLadingId field to given value.

### HasBillOfLadingId

`func (o *BillOfLadingLineDto) HasBillOfLadingId() bool`

HasBillOfLadingId returns a boolean if a field has been set.

### SetBillOfLadingIdNil

`func (o *BillOfLadingLineDto) SetBillOfLadingIdNil(b bool)`

 SetBillOfLadingIdNil sets the value for BillOfLadingId to be an explicit nil

### UnsetBillOfLadingId
`func (o *BillOfLadingLineDto) UnsetBillOfLadingId()`

UnsetBillOfLadingId ensures that no value is present for BillOfLadingId, not even an explicit nil
### GetTenantId

`func (o *BillOfLadingLineDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillOfLadingLineDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillOfLadingLineDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BillOfLadingLineDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BillOfLadingLineDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BillOfLadingLineDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


