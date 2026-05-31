# WaybillLineCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **NullableInt32** |  | [optional] 
**GrossWeightKg** | Pointer to **NullableFloat64** |  | [optional] 
**VolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**PackageType** | Pointer to **NullableString** |  | [optional] 
**LengthCm** | Pointer to **NullableFloat64** |  | [optional] 
**WidthCm** | Pointer to **NullableFloat64** |  | [optional] 
**HeightCm** | Pointer to **NullableFloat64** |  | [optional] 
**HsCode** | Pointer to **NullableString** |  | [optional] 
**MarksAndNumbers** | Pointer to **NullableString** |  | [optional] 
**DeclaredValue** | Pointer to **NullableFloat64** |  | [optional] 
**DeclaredValueCurrencyId** | Pointer to **NullableString** |  | [optional] 
**SealNumber** | Pointer to **NullableString** |  | [optional] 
**ContainerNumber** | Pointer to **NullableString** |  | [optional] 
**ChargeableWeightKg** | Pointer to **NullableFloat64** |  | [optional] 
**IataRateClass** | Pointer to **NullableString** |  | [optional] 
**DangerousGoodsClass** | Pointer to **NullableString** |  | [optional] 
**UnHazmatNumber** | Pointer to **NullableString** |  | [optional] 
**WagonNumber** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWaybillLineCreateDto

`func NewWaybillLineCreateDto() *WaybillLineCreateDto`

NewWaybillLineCreateDto instantiates a new WaybillLineCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaybillLineCreateDtoWithDefaults

`func NewWaybillLineCreateDtoWithDefaults() *WaybillLineCreateDto`

NewWaybillLineCreateDtoWithDefaults instantiates a new WaybillLineCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WaybillLineCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WaybillLineCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WaybillLineCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WaybillLineCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *WaybillLineCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WaybillLineCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WaybillLineCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WaybillLineCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *WaybillLineCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WaybillLineCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WaybillLineCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WaybillLineCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WaybillLineCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WaybillLineCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *WaybillLineCreateDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *WaybillLineCreateDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *WaybillLineCreateDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *WaybillLineCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *WaybillLineCreateDto) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *WaybillLineCreateDto) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetGrossWeightKg

`func (o *WaybillLineCreateDto) GetGrossWeightKg() float64`

GetGrossWeightKg returns the GrossWeightKg field if non-nil, zero value otherwise.

### GetGrossWeightKgOk

`func (o *WaybillLineCreateDto) GetGrossWeightKgOk() (*float64, bool)`

GetGrossWeightKgOk returns a tuple with the GrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossWeightKg

`func (o *WaybillLineCreateDto) SetGrossWeightKg(v float64)`

SetGrossWeightKg sets GrossWeightKg field to given value.

### HasGrossWeightKg

`func (o *WaybillLineCreateDto) HasGrossWeightKg() bool`

HasGrossWeightKg returns a boolean if a field has been set.

### SetGrossWeightKgNil

`func (o *WaybillLineCreateDto) SetGrossWeightKgNil(b bool)`

 SetGrossWeightKgNil sets the value for GrossWeightKg to be an explicit nil

### UnsetGrossWeightKg
`func (o *WaybillLineCreateDto) UnsetGrossWeightKg()`

UnsetGrossWeightKg ensures that no value is present for GrossWeightKg, not even an explicit nil
### GetVolumeM3

`func (o *WaybillLineCreateDto) GetVolumeM3() float64`

GetVolumeM3 returns the VolumeM3 field if non-nil, zero value otherwise.

### GetVolumeM3Ok

`func (o *WaybillLineCreateDto) GetVolumeM3Ok() (*float64, bool)`

GetVolumeM3Ok returns a tuple with the VolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeM3

`func (o *WaybillLineCreateDto) SetVolumeM3(v float64)`

SetVolumeM3 sets VolumeM3 field to given value.

### HasVolumeM3

`func (o *WaybillLineCreateDto) HasVolumeM3() bool`

HasVolumeM3 returns a boolean if a field has been set.

### SetVolumeM3Nil

`func (o *WaybillLineCreateDto) SetVolumeM3Nil(b bool)`

 SetVolumeM3Nil sets the value for VolumeM3 to be an explicit nil

### UnsetVolumeM3
`func (o *WaybillLineCreateDto) UnsetVolumeM3()`

UnsetVolumeM3 ensures that no value is present for VolumeM3, not even an explicit nil
### GetPackageType

`func (o *WaybillLineCreateDto) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *WaybillLineCreateDto) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *WaybillLineCreateDto) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *WaybillLineCreateDto) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### SetPackageTypeNil

`func (o *WaybillLineCreateDto) SetPackageTypeNil(b bool)`

 SetPackageTypeNil sets the value for PackageType to be an explicit nil

### UnsetPackageType
`func (o *WaybillLineCreateDto) UnsetPackageType()`

UnsetPackageType ensures that no value is present for PackageType, not even an explicit nil
### GetLengthCm

`func (o *WaybillLineCreateDto) GetLengthCm() float64`

GetLengthCm returns the LengthCm field if non-nil, zero value otherwise.

### GetLengthCmOk

`func (o *WaybillLineCreateDto) GetLengthCmOk() (*float64, bool)`

GetLengthCmOk returns a tuple with the LengthCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthCm

`func (o *WaybillLineCreateDto) SetLengthCm(v float64)`

SetLengthCm sets LengthCm field to given value.

### HasLengthCm

`func (o *WaybillLineCreateDto) HasLengthCm() bool`

HasLengthCm returns a boolean if a field has been set.

### SetLengthCmNil

`func (o *WaybillLineCreateDto) SetLengthCmNil(b bool)`

 SetLengthCmNil sets the value for LengthCm to be an explicit nil

### UnsetLengthCm
`func (o *WaybillLineCreateDto) UnsetLengthCm()`

UnsetLengthCm ensures that no value is present for LengthCm, not even an explicit nil
### GetWidthCm

`func (o *WaybillLineCreateDto) GetWidthCm() float64`

GetWidthCm returns the WidthCm field if non-nil, zero value otherwise.

### GetWidthCmOk

`func (o *WaybillLineCreateDto) GetWidthCmOk() (*float64, bool)`

GetWidthCmOk returns a tuple with the WidthCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthCm

`func (o *WaybillLineCreateDto) SetWidthCm(v float64)`

SetWidthCm sets WidthCm field to given value.

### HasWidthCm

`func (o *WaybillLineCreateDto) HasWidthCm() bool`

HasWidthCm returns a boolean if a field has been set.

### SetWidthCmNil

`func (o *WaybillLineCreateDto) SetWidthCmNil(b bool)`

 SetWidthCmNil sets the value for WidthCm to be an explicit nil

### UnsetWidthCm
`func (o *WaybillLineCreateDto) UnsetWidthCm()`

UnsetWidthCm ensures that no value is present for WidthCm, not even an explicit nil
### GetHeightCm

`func (o *WaybillLineCreateDto) GetHeightCm() float64`

GetHeightCm returns the HeightCm field if non-nil, zero value otherwise.

### GetHeightCmOk

`func (o *WaybillLineCreateDto) GetHeightCmOk() (*float64, bool)`

GetHeightCmOk returns a tuple with the HeightCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightCm

`func (o *WaybillLineCreateDto) SetHeightCm(v float64)`

SetHeightCm sets HeightCm field to given value.

### HasHeightCm

`func (o *WaybillLineCreateDto) HasHeightCm() bool`

HasHeightCm returns a boolean if a field has been set.

### SetHeightCmNil

`func (o *WaybillLineCreateDto) SetHeightCmNil(b bool)`

 SetHeightCmNil sets the value for HeightCm to be an explicit nil

### UnsetHeightCm
`func (o *WaybillLineCreateDto) UnsetHeightCm()`

UnsetHeightCm ensures that no value is present for HeightCm, not even an explicit nil
### GetHsCode

`func (o *WaybillLineCreateDto) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *WaybillLineCreateDto) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *WaybillLineCreateDto) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *WaybillLineCreateDto) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *WaybillLineCreateDto) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *WaybillLineCreateDto) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil
### GetMarksAndNumbers

`func (o *WaybillLineCreateDto) GetMarksAndNumbers() string`

GetMarksAndNumbers returns the MarksAndNumbers field if non-nil, zero value otherwise.

### GetMarksAndNumbersOk

`func (o *WaybillLineCreateDto) GetMarksAndNumbersOk() (*string, bool)`

GetMarksAndNumbersOk returns a tuple with the MarksAndNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarksAndNumbers

`func (o *WaybillLineCreateDto) SetMarksAndNumbers(v string)`

SetMarksAndNumbers sets MarksAndNumbers field to given value.

### HasMarksAndNumbers

`func (o *WaybillLineCreateDto) HasMarksAndNumbers() bool`

HasMarksAndNumbers returns a boolean if a field has been set.

### SetMarksAndNumbersNil

`func (o *WaybillLineCreateDto) SetMarksAndNumbersNil(b bool)`

 SetMarksAndNumbersNil sets the value for MarksAndNumbers to be an explicit nil

### UnsetMarksAndNumbers
`func (o *WaybillLineCreateDto) UnsetMarksAndNumbers()`

UnsetMarksAndNumbers ensures that no value is present for MarksAndNumbers, not even an explicit nil
### GetDeclaredValue

`func (o *WaybillLineCreateDto) GetDeclaredValue() float64`

GetDeclaredValue returns the DeclaredValue field if non-nil, zero value otherwise.

### GetDeclaredValueOk

`func (o *WaybillLineCreateDto) GetDeclaredValueOk() (*float64, bool)`

GetDeclaredValueOk returns a tuple with the DeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValue

`func (o *WaybillLineCreateDto) SetDeclaredValue(v float64)`

SetDeclaredValue sets DeclaredValue field to given value.

### HasDeclaredValue

`func (o *WaybillLineCreateDto) HasDeclaredValue() bool`

HasDeclaredValue returns a boolean if a field has been set.

### SetDeclaredValueNil

`func (o *WaybillLineCreateDto) SetDeclaredValueNil(b bool)`

 SetDeclaredValueNil sets the value for DeclaredValue to be an explicit nil

### UnsetDeclaredValue
`func (o *WaybillLineCreateDto) UnsetDeclaredValue()`

UnsetDeclaredValue ensures that no value is present for DeclaredValue, not even an explicit nil
### GetDeclaredValueCurrencyId

`func (o *WaybillLineCreateDto) GetDeclaredValueCurrencyId() string`

GetDeclaredValueCurrencyId returns the DeclaredValueCurrencyId field if non-nil, zero value otherwise.

### GetDeclaredValueCurrencyIdOk

`func (o *WaybillLineCreateDto) GetDeclaredValueCurrencyIdOk() (*string, bool)`

GetDeclaredValueCurrencyIdOk returns a tuple with the DeclaredValueCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueCurrencyId

`func (o *WaybillLineCreateDto) SetDeclaredValueCurrencyId(v string)`

SetDeclaredValueCurrencyId sets DeclaredValueCurrencyId field to given value.

### HasDeclaredValueCurrencyId

`func (o *WaybillLineCreateDto) HasDeclaredValueCurrencyId() bool`

HasDeclaredValueCurrencyId returns a boolean if a field has been set.

### SetDeclaredValueCurrencyIdNil

`func (o *WaybillLineCreateDto) SetDeclaredValueCurrencyIdNil(b bool)`

 SetDeclaredValueCurrencyIdNil sets the value for DeclaredValueCurrencyId to be an explicit nil

### UnsetDeclaredValueCurrencyId
`func (o *WaybillLineCreateDto) UnsetDeclaredValueCurrencyId()`

UnsetDeclaredValueCurrencyId ensures that no value is present for DeclaredValueCurrencyId, not even an explicit nil
### GetSealNumber

`func (o *WaybillLineCreateDto) GetSealNumber() string`

GetSealNumber returns the SealNumber field if non-nil, zero value otherwise.

### GetSealNumberOk

`func (o *WaybillLineCreateDto) GetSealNumberOk() (*string, bool)`

GetSealNumberOk returns a tuple with the SealNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealNumber

`func (o *WaybillLineCreateDto) SetSealNumber(v string)`

SetSealNumber sets SealNumber field to given value.

### HasSealNumber

`func (o *WaybillLineCreateDto) HasSealNumber() bool`

HasSealNumber returns a boolean if a field has been set.

### SetSealNumberNil

`func (o *WaybillLineCreateDto) SetSealNumberNil(b bool)`

 SetSealNumberNil sets the value for SealNumber to be an explicit nil

### UnsetSealNumber
`func (o *WaybillLineCreateDto) UnsetSealNumber()`

UnsetSealNumber ensures that no value is present for SealNumber, not even an explicit nil
### GetContainerNumber

`func (o *WaybillLineCreateDto) GetContainerNumber() string`

GetContainerNumber returns the ContainerNumber field if non-nil, zero value otherwise.

### GetContainerNumberOk

`func (o *WaybillLineCreateDto) GetContainerNumberOk() (*string, bool)`

GetContainerNumberOk returns a tuple with the ContainerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNumber

`func (o *WaybillLineCreateDto) SetContainerNumber(v string)`

SetContainerNumber sets ContainerNumber field to given value.

### HasContainerNumber

`func (o *WaybillLineCreateDto) HasContainerNumber() bool`

HasContainerNumber returns a boolean if a field has been set.

### SetContainerNumberNil

`func (o *WaybillLineCreateDto) SetContainerNumberNil(b bool)`

 SetContainerNumberNil sets the value for ContainerNumber to be an explicit nil

### UnsetContainerNumber
`func (o *WaybillLineCreateDto) UnsetContainerNumber()`

UnsetContainerNumber ensures that no value is present for ContainerNumber, not even an explicit nil
### GetChargeableWeightKg

`func (o *WaybillLineCreateDto) GetChargeableWeightKg() float64`

GetChargeableWeightKg returns the ChargeableWeightKg field if non-nil, zero value otherwise.

### GetChargeableWeightKgOk

`func (o *WaybillLineCreateDto) GetChargeableWeightKgOk() (*float64, bool)`

GetChargeableWeightKgOk returns a tuple with the ChargeableWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeableWeightKg

`func (o *WaybillLineCreateDto) SetChargeableWeightKg(v float64)`

SetChargeableWeightKg sets ChargeableWeightKg field to given value.

### HasChargeableWeightKg

`func (o *WaybillLineCreateDto) HasChargeableWeightKg() bool`

HasChargeableWeightKg returns a boolean if a field has been set.

### SetChargeableWeightKgNil

`func (o *WaybillLineCreateDto) SetChargeableWeightKgNil(b bool)`

 SetChargeableWeightKgNil sets the value for ChargeableWeightKg to be an explicit nil

### UnsetChargeableWeightKg
`func (o *WaybillLineCreateDto) UnsetChargeableWeightKg()`

UnsetChargeableWeightKg ensures that no value is present for ChargeableWeightKg, not even an explicit nil
### GetIataRateClass

`func (o *WaybillLineCreateDto) GetIataRateClass() string`

GetIataRateClass returns the IataRateClass field if non-nil, zero value otherwise.

### GetIataRateClassOk

`func (o *WaybillLineCreateDto) GetIataRateClassOk() (*string, bool)`

GetIataRateClassOk returns a tuple with the IataRateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIataRateClass

`func (o *WaybillLineCreateDto) SetIataRateClass(v string)`

SetIataRateClass sets IataRateClass field to given value.

### HasIataRateClass

`func (o *WaybillLineCreateDto) HasIataRateClass() bool`

HasIataRateClass returns a boolean if a field has been set.

### SetIataRateClassNil

`func (o *WaybillLineCreateDto) SetIataRateClassNil(b bool)`

 SetIataRateClassNil sets the value for IataRateClass to be an explicit nil

### UnsetIataRateClass
`func (o *WaybillLineCreateDto) UnsetIataRateClass()`

UnsetIataRateClass ensures that no value is present for IataRateClass, not even an explicit nil
### GetDangerousGoodsClass

`func (o *WaybillLineCreateDto) GetDangerousGoodsClass() string`

GetDangerousGoodsClass returns the DangerousGoodsClass field if non-nil, zero value otherwise.

### GetDangerousGoodsClassOk

`func (o *WaybillLineCreateDto) GetDangerousGoodsClassOk() (*string, bool)`

GetDangerousGoodsClassOk returns a tuple with the DangerousGoodsClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDangerousGoodsClass

`func (o *WaybillLineCreateDto) SetDangerousGoodsClass(v string)`

SetDangerousGoodsClass sets DangerousGoodsClass field to given value.

### HasDangerousGoodsClass

`func (o *WaybillLineCreateDto) HasDangerousGoodsClass() bool`

HasDangerousGoodsClass returns a boolean if a field has been set.

### SetDangerousGoodsClassNil

`func (o *WaybillLineCreateDto) SetDangerousGoodsClassNil(b bool)`

 SetDangerousGoodsClassNil sets the value for DangerousGoodsClass to be an explicit nil

### UnsetDangerousGoodsClass
`func (o *WaybillLineCreateDto) UnsetDangerousGoodsClass()`

UnsetDangerousGoodsClass ensures that no value is present for DangerousGoodsClass, not even an explicit nil
### GetUnHazmatNumber

`func (o *WaybillLineCreateDto) GetUnHazmatNumber() string`

GetUnHazmatNumber returns the UnHazmatNumber field if non-nil, zero value otherwise.

### GetUnHazmatNumberOk

`func (o *WaybillLineCreateDto) GetUnHazmatNumberOk() (*string, bool)`

GetUnHazmatNumberOk returns a tuple with the UnHazmatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnHazmatNumber

`func (o *WaybillLineCreateDto) SetUnHazmatNumber(v string)`

SetUnHazmatNumber sets UnHazmatNumber field to given value.

### HasUnHazmatNumber

`func (o *WaybillLineCreateDto) HasUnHazmatNumber() bool`

HasUnHazmatNumber returns a boolean if a field has been set.

### SetUnHazmatNumberNil

`func (o *WaybillLineCreateDto) SetUnHazmatNumberNil(b bool)`

 SetUnHazmatNumberNil sets the value for UnHazmatNumber to be an explicit nil

### UnsetUnHazmatNumber
`func (o *WaybillLineCreateDto) UnsetUnHazmatNumber()`

UnsetUnHazmatNumber ensures that no value is present for UnHazmatNumber, not even an explicit nil
### GetWagonNumber

`func (o *WaybillLineCreateDto) GetWagonNumber() string`

GetWagonNumber returns the WagonNumber field if non-nil, zero value otherwise.

### GetWagonNumberOk

`func (o *WaybillLineCreateDto) GetWagonNumberOk() (*string, bool)`

GetWagonNumberOk returns a tuple with the WagonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagonNumber

`func (o *WaybillLineCreateDto) SetWagonNumber(v string)`

SetWagonNumber sets WagonNumber field to given value.

### HasWagonNumber

`func (o *WaybillLineCreateDto) HasWagonNumber() bool`

HasWagonNumber returns a boolean if a field has been set.

### SetWagonNumberNil

`func (o *WaybillLineCreateDto) SetWagonNumberNil(b bool)`

 SetWagonNumberNil sets the value for WagonNumber to be an explicit nil

### UnsetWagonNumber
`func (o *WaybillLineCreateDto) UnsetWagonNumber()`

UnsetWagonNumber ensures that no value is present for WagonNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


