# WaybillLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ParentDocumentType** | Pointer to **NullableString** |  | [optional] 
**SeawayBillId** | Pointer to **NullableString** |  | [optional] 
**AirwayBillId** | Pointer to **NullableString** |  | [optional] 
**RoadWaybillId** | Pointer to **NullableString** |  | [optional] 
**RailWaybillId** | Pointer to **NullableString** |  | [optional] 
**LineNumber** | Pointer to **int32** |  | [optional] 
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
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWaybillLineDto

`func NewWaybillLineDto() *WaybillLineDto`

NewWaybillLineDto instantiates a new WaybillLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaybillLineDtoWithDefaults

`func NewWaybillLineDtoWithDefaults() *WaybillLineDto`

NewWaybillLineDtoWithDefaults instantiates a new WaybillLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WaybillLineDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WaybillLineDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WaybillLineDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WaybillLineDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WaybillLineDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WaybillLineDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WaybillLineDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WaybillLineDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WaybillLineDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WaybillLineDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *WaybillLineDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *WaybillLineDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetParentDocumentType

`func (o *WaybillLineDto) GetParentDocumentType() string`

GetParentDocumentType returns the ParentDocumentType field if non-nil, zero value otherwise.

### GetParentDocumentTypeOk

`func (o *WaybillLineDto) GetParentDocumentTypeOk() (*string, bool)`

GetParentDocumentTypeOk returns a tuple with the ParentDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDocumentType

`func (o *WaybillLineDto) SetParentDocumentType(v string)`

SetParentDocumentType sets ParentDocumentType field to given value.

### HasParentDocumentType

`func (o *WaybillLineDto) HasParentDocumentType() bool`

HasParentDocumentType returns a boolean if a field has been set.

### SetParentDocumentTypeNil

`func (o *WaybillLineDto) SetParentDocumentTypeNil(b bool)`

 SetParentDocumentTypeNil sets the value for ParentDocumentType to be an explicit nil

### UnsetParentDocumentType
`func (o *WaybillLineDto) UnsetParentDocumentType()`

UnsetParentDocumentType ensures that no value is present for ParentDocumentType, not even an explicit nil
### GetSeawayBillId

`func (o *WaybillLineDto) GetSeawayBillId() string`

GetSeawayBillId returns the SeawayBillId field if non-nil, zero value otherwise.

### GetSeawayBillIdOk

`func (o *WaybillLineDto) GetSeawayBillIdOk() (*string, bool)`

GetSeawayBillIdOk returns a tuple with the SeawayBillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeawayBillId

`func (o *WaybillLineDto) SetSeawayBillId(v string)`

SetSeawayBillId sets SeawayBillId field to given value.

### HasSeawayBillId

`func (o *WaybillLineDto) HasSeawayBillId() bool`

HasSeawayBillId returns a boolean if a field has been set.

### SetSeawayBillIdNil

`func (o *WaybillLineDto) SetSeawayBillIdNil(b bool)`

 SetSeawayBillIdNil sets the value for SeawayBillId to be an explicit nil

### UnsetSeawayBillId
`func (o *WaybillLineDto) UnsetSeawayBillId()`

UnsetSeawayBillId ensures that no value is present for SeawayBillId, not even an explicit nil
### GetAirwayBillId

`func (o *WaybillLineDto) GetAirwayBillId() string`

GetAirwayBillId returns the AirwayBillId field if non-nil, zero value otherwise.

### GetAirwayBillIdOk

`func (o *WaybillLineDto) GetAirwayBillIdOk() (*string, bool)`

GetAirwayBillIdOk returns a tuple with the AirwayBillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwayBillId

`func (o *WaybillLineDto) SetAirwayBillId(v string)`

SetAirwayBillId sets AirwayBillId field to given value.

### HasAirwayBillId

`func (o *WaybillLineDto) HasAirwayBillId() bool`

HasAirwayBillId returns a boolean if a field has been set.

### SetAirwayBillIdNil

`func (o *WaybillLineDto) SetAirwayBillIdNil(b bool)`

 SetAirwayBillIdNil sets the value for AirwayBillId to be an explicit nil

### UnsetAirwayBillId
`func (o *WaybillLineDto) UnsetAirwayBillId()`

UnsetAirwayBillId ensures that no value is present for AirwayBillId, not even an explicit nil
### GetRoadWaybillId

`func (o *WaybillLineDto) GetRoadWaybillId() string`

GetRoadWaybillId returns the RoadWaybillId field if non-nil, zero value otherwise.

### GetRoadWaybillIdOk

`func (o *WaybillLineDto) GetRoadWaybillIdOk() (*string, bool)`

GetRoadWaybillIdOk returns a tuple with the RoadWaybillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadWaybillId

`func (o *WaybillLineDto) SetRoadWaybillId(v string)`

SetRoadWaybillId sets RoadWaybillId field to given value.

### HasRoadWaybillId

`func (o *WaybillLineDto) HasRoadWaybillId() bool`

HasRoadWaybillId returns a boolean if a field has been set.

### SetRoadWaybillIdNil

`func (o *WaybillLineDto) SetRoadWaybillIdNil(b bool)`

 SetRoadWaybillIdNil sets the value for RoadWaybillId to be an explicit nil

### UnsetRoadWaybillId
`func (o *WaybillLineDto) UnsetRoadWaybillId()`

UnsetRoadWaybillId ensures that no value is present for RoadWaybillId, not even an explicit nil
### GetRailWaybillId

`func (o *WaybillLineDto) GetRailWaybillId() string`

GetRailWaybillId returns the RailWaybillId field if non-nil, zero value otherwise.

### GetRailWaybillIdOk

`func (o *WaybillLineDto) GetRailWaybillIdOk() (*string, bool)`

GetRailWaybillIdOk returns a tuple with the RailWaybillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailWaybillId

`func (o *WaybillLineDto) SetRailWaybillId(v string)`

SetRailWaybillId sets RailWaybillId field to given value.

### HasRailWaybillId

`func (o *WaybillLineDto) HasRailWaybillId() bool`

HasRailWaybillId returns a boolean if a field has been set.

### SetRailWaybillIdNil

`func (o *WaybillLineDto) SetRailWaybillIdNil(b bool)`

 SetRailWaybillIdNil sets the value for RailWaybillId to be an explicit nil

### UnsetRailWaybillId
`func (o *WaybillLineDto) UnsetRailWaybillId()`

UnsetRailWaybillId ensures that no value is present for RailWaybillId, not even an explicit nil
### GetLineNumber

`func (o *WaybillLineDto) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *WaybillLineDto) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *WaybillLineDto) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *WaybillLineDto) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetDescription

`func (o *WaybillLineDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WaybillLineDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WaybillLineDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WaybillLineDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WaybillLineDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WaybillLineDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *WaybillLineDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *WaybillLineDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *WaybillLineDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *WaybillLineDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *WaybillLineDto) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *WaybillLineDto) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetGrossWeightKg

`func (o *WaybillLineDto) GetGrossWeightKg() float64`

GetGrossWeightKg returns the GrossWeightKg field if non-nil, zero value otherwise.

### GetGrossWeightKgOk

`func (o *WaybillLineDto) GetGrossWeightKgOk() (*float64, bool)`

GetGrossWeightKgOk returns a tuple with the GrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossWeightKg

`func (o *WaybillLineDto) SetGrossWeightKg(v float64)`

SetGrossWeightKg sets GrossWeightKg field to given value.

### HasGrossWeightKg

`func (o *WaybillLineDto) HasGrossWeightKg() bool`

HasGrossWeightKg returns a boolean if a field has been set.

### SetGrossWeightKgNil

`func (o *WaybillLineDto) SetGrossWeightKgNil(b bool)`

 SetGrossWeightKgNil sets the value for GrossWeightKg to be an explicit nil

### UnsetGrossWeightKg
`func (o *WaybillLineDto) UnsetGrossWeightKg()`

UnsetGrossWeightKg ensures that no value is present for GrossWeightKg, not even an explicit nil
### GetVolumeM3

`func (o *WaybillLineDto) GetVolumeM3() float64`

GetVolumeM3 returns the VolumeM3 field if non-nil, zero value otherwise.

### GetVolumeM3Ok

`func (o *WaybillLineDto) GetVolumeM3Ok() (*float64, bool)`

GetVolumeM3Ok returns a tuple with the VolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeM3

`func (o *WaybillLineDto) SetVolumeM3(v float64)`

SetVolumeM3 sets VolumeM3 field to given value.

### HasVolumeM3

`func (o *WaybillLineDto) HasVolumeM3() bool`

HasVolumeM3 returns a boolean if a field has been set.

### SetVolumeM3Nil

`func (o *WaybillLineDto) SetVolumeM3Nil(b bool)`

 SetVolumeM3Nil sets the value for VolumeM3 to be an explicit nil

### UnsetVolumeM3
`func (o *WaybillLineDto) UnsetVolumeM3()`

UnsetVolumeM3 ensures that no value is present for VolumeM3, not even an explicit nil
### GetPackageType

`func (o *WaybillLineDto) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *WaybillLineDto) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *WaybillLineDto) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *WaybillLineDto) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### SetPackageTypeNil

`func (o *WaybillLineDto) SetPackageTypeNil(b bool)`

 SetPackageTypeNil sets the value for PackageType to be an explicit nil

### UnsetPackageType
`func (o *WaybillLineDto) UnsetPackageType()`

UnsetPackageType ensures that no value is present for PackageType, not even an explicit nil
### GetLengthCm

`func (o *WaybillLineDto) GetLengthCm() float64`

GetLengthCm returns the LengthCm field if non-nil, zero value otherwise.

### GetLengthCmOk

`func (o *WaybillLineDto) GetLengthCmOk() (*float64, bool)`

GetLengthCmOk returns a tuple with the LengthCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthCm

`func (o *WaybillLineDto) SetLengthCm(v float64)`

SetLengthCm sets LengthCm field to given value.

### HasLengthCm

`func (o *WaybillLineDto) HasLengthCm() bool`

HasLengthCm returns a boolean if a field has been set.

### SetLengthCmNil

`func (o *WaybillLineDto) SetLengthCmNil(b bool)`

 SetLengthCmNil sets the value for LengthCm to be an explicit nil

### UnsetLengthCm
`func (o *WaybillLineDto) UnsetLengthCm()`

UnsetLengthCm ensures that no value is present for LengthCm, not even an explicit nil
### GetWidthCm

`func (o *WaybillLineDto) GetWidthCm() float64`

GetWidthCm returns the WidthCm field if non-nil, zero value otherwise.

### GetWidthCmOk

`func (o *WaybillLineDto) GetWidthCmOk() (*float64, bool)`

GetWidthCmOk returns a tuple with the WidthCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthCm

`func (o *WaybillLineDto) SetWidthCm(v float64)`

SetWidthCm sets WidthCm field to given value.

### HasWidthCm

`func (o *WaybillLineDto) HasWidthCm() bool`

HasWidthCm returns a boolean if a field has been set.

### SetWidthCmNil

`func (o *WaybillLineDto) SetWidthCmNil(b bool)`

 SetWidthCmNil sets the value for WidthCm to be an explicit nil

### UnsetWidthCm
`func (o *WaybillLineDto) UnsetWidthCm()`

UnsetWidthCm ensures that no value is present for WidthCm, not even an explicit nil
### GetHeightCm

`func (o *WaybillLineDto) GetHeightCm() float64`

GetHeightCm returns the HeightCm field if non-nil, zero value otherwise.

### GetHeightCmOk

`func (o *WaybillLineDto) GetHeightCmOk() (*float64, bool)`

GetHeightCmOk returns a tuple with the HeightCm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightCm

`func (o *WaybillLineDto) SetHeightCm(v float64)`

SetHeightCm sets HeightCm field to given value.

### HasHeightCm

`func (o *WaybillLineDto) HasHeightCm() bool`

HasHeightCm returns a boolean if a field has been set.

### SetHeightCmNil

`func (o *WaybillLineDto) SetHeightCmNil(b bool)`

 SetHeightCmNil sets the value for HeightCm to be an explicit nil

### UnsetHeightCm
`func (o *WaybillLineDto) UnsetHeightCm()`

UnsetHeightCm ensures that no value is present for HeightCm, not even an explicit nil
### GetHsCode

`func (o *WaybillLineDto) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *WaybillLineDto) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *WaybillLineDto) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *WaybillLineDto) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *WaybillLineDto) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *WaybillLineDto) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil
### GetMarksAndNumbers

`func (o *WaybillLineDto) GetMarksAndNumbers() string`

GetMarksAndNumbers returns the MarksAndNumbers field if non-nil, zero value otherwise.

### GetMarksAndNumbersOk

`func (o *WaybillLineDto) GetMarksAndNumbersOk() (*string, bool)`

GetMarksAndNumbersOk returns a tuple with the MarksAndNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarksAndNumbers

`func (o *WaybillLineDto) SetMarksAndNumbers(v string)`

SetMarksAndNumbers sets MarksAndNumbers field to given value.

### HasMarksAndNumbers

`func (o *WaybillLineDto) HasMarksAndNumbers() bool`

HasMarksAndNumbers returns a boolean if a field has been set.

### SetMarksAndNumbersNil

`func (o *WaybillLineDto) SetMarksAndNumbersNil(b bool)`

 SetMarksAndNumbersNil sets the value for MarksAndNumbers to be an explicit nil

### UnsetMarksAndNumbers
`func (o *WaybillLineDto) UnsetMarksAndNumbers()`

UnsetMarksAndNumbers ensures that no value is present for MarksAndNumbers, not even an explicit nil
### GetDeclaredValue

`func (o *WaybillLineDto) GetDeclaredValue() float64`

GetDeclaredValue returns the DeclaredValue field if non-nil, zero value otherwise.

### GetDeclaredValueOk

`func (o *WaybillLineDto) GetDeclaredValueOk() (*float64, bool)`

GetDeclaredValueOk returns a tuple with the DeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValue

`func (o *WaybillLineDto) SetDeclaredValue(v float64)`

SetDeclaredValue sets DeclaredValue field to given value.

### HasDeclaredValue

`func (o *WaybillLineDto) HasDeclaredValue() bool`

HasDeclaredValue returns a boolean if a field has been set.

### SetDeclaredValueNil

`func (o *WaybillLineDto) SetDeclaredValueNil(b bool)`

 SetDeclaredValueNil sets the value for DeclaredValue to be an explicit nil

### UnsetDeclaredValue
`func (o *WaybillLineDto) UnsetDeclaredValue()`

UnsetDeclaredValue ensures that no value is present for DeclaredValue, not even an explicit nil
### GetDeclaredValueCurrencyId

`func (o *WaybillLineDto) GetDeclaredValueCurrencyId() string`

GetDeclaredValueCurrencyId returns the DeclaredValueCurrencyId field if non-nil, zero value otherwise.

### GetDeclaredValueCurrencyIdOk

`func (o *WaybillLineDto) GetDeclaredValueCurrencyIdOk() (*string, bool)`

GetDeclaredValueCurrencyIdOk returns a tuple with the DeclaredValueCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueCurrencyId

`func (o *WaybillLineDto) SetDeclaredValueCurrencyId(v string)`

SetDeclaredValueCurrencyId sets DeclaredValueCurrencyId field to given value.

### HasDeclaredValueCurrencyId

`func (o *WaybillLineDto) HasDeclaredValueCurrencyId() bool`

HasDeclaredValueCurrencyId returns a boolean if a field has been set.

### SetDeclaredValueCurrencyIdNil

`func (o *WaybillLineDto) SetDeclaredValueCurrencyIdNil(b bool)`

 SetDeclaredValueCurrencyIdNil sets the value for DeclaredValueCurrencyId to be an explicit nil

### UnsetDeclaredValueCurrencyId
`func (o *WaybillLineDto) UnsetDeclaredValueCurrencyId()`

UnsetDeclaredValueCurrencyId ensures that no value is present for DeclaredValueCurrencyId, not even an explicit nil
### GetSealNumber

`func (o *WaybillLineDto) GetSealNumber() string`

GetSealNumber returns the SealNumber field if non-nil, zero value otherwise.

### GetSealNumberOk

`func (o *WaybillLineDto) GetSealNumberOk() (*string, bool)`

GetSealNumberOk returns a tuple with the SealNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealNumber

`func (o *WaybillLineDto) SetSealNumber(v string)`

SetSealNumber sets SealNumber field to given value.

### HasSealNumber

`func (o *WaybillLineDto) HasSealNumber() bool`

HasSealNumber returns a boolean if a field has been set.

### SetSealNumberNil

`func (o *WaybillLineDto) SetSealNumberNil(b bool)`

 SetSealNumberNil sets the value for SealNumber to be an explicit nil

### UnsetSealNumber
`func (o *WaybillLineDto) UnsetSealNumber()`

UnsetSealNumber ensures that no value is present for SealNumber, not even an explicit nil
### GetContainerNumber

`func (o *WaybillLineDto) GetContainerNumber() string`

GetContainerNumber returns the ContainerNumber field if non-nil, zero value otherwise.

### GetContainerNumberOk

`func (o *WaybillLineDto) GetContainerNumberOk() (*string, bool)`

GetContainerNumberOk returns a tuple with the ContainerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNumber

`func (o *WaybillLineDto) SetContainerNumber(v string)`

SetContainerNumber sets ContainerNumber field to given value.

### HasContainerNumber

`func (o *WaybillLineDto) HasContainerNumber() bool`

HasContainerNumber returns a boolean if a field has been set.

### SetContainerNumberNil

`func (o *WaybillLineDto) SetContainerNumberNil(b bool)`

 SetContainerNumberNil sets the value for ContainerNumber to be an explicit nil

### UnsetContainerNumber
`func (o *WaybillLineDto) UnsetContainerNumber()`

UnsetContainerNumber ensures that no value is present for ContainerNumber, not even an explicit nil
### GetChargeableWeightKg

`func (o *WaybillLineDto) GetChargeableWeightKg() float64`

GetChargeableWeightKg returns the ChargeableWeightKg field if non-nil, zero value otherwise.

### GetChargeableWeightKgOk

`func (o *WaybillLineDto) GetChargeableWeightKgOk() (*float64, bool)`

GetChargeableWeightKgOk returns a tuple with the ChargeableWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeableWeightKg

`func (o *WaybillLineDto) SetChargeableWeightKg(v float64)`

SetChargeableWeightKg sets ChargeableWeightKg field to given value.

### HasChargeableWeightKg

`func (o *WaybillLineDto) HasChargeableWeightKg() bool`

HasChargeableWeightKg returns a boolean if a field has been set.

### SetChargeableWeightKgNil

`func (o *WaybillLineDto) SetChargeableWeightKgNil(b bool)`

 SetChargeableWeightKgNil sets the value for ChargeableWeightKg to be an explicit nil

### UnsetChargeableWeightKg
`func (o *WaybillLineDto) UnsetChargeableWeightKg()`

UnsetChargeableWeightKg ensures that no value is present for ChargeableWeightKg, not even an explicit nil
### GetIataRateClass

`func (o *WaybillLineDto) GetIataRateClass() string`

GetIataRateClass returns the IataRateClass field if non-nil, zero value otherwise.

### GetIataRateClassOk

`func (o *WaybillLineDto) GetIataRateClassOk() (*string, bool)`

GetIataRateClassOk returns a tuple with the IataRateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIataRateClass

`func (o *WaybillLineDto) SetIataRateClass(v string)`

SetIataRateClass sets IataRateClass field to given value.

### HasIataRateClass

`func (o *WaybillLineDto) HasIataRateClass() bool`

HasIataRateClass returns a boolean if a field has been set.

### SetIataRateClassNil

`func (o *WaybillLineDto) SetIataRateClassNil(b bool)`

 SetIataRateClassNil sets the value for IataRateClass to be an explicit nil

### UnsetIataRateClass
`func (o *WaybillLineDto) UnsetIataRateClass()`

UnsetIataRateClass ensures that no value is present for IataRateClass, not even an explicit nil
### GetDangerousGoodsClass

`func (o *WaybillLineDto) GetDangerousGoodsClass() string`

GetDangerousGoodsClass returns the DangerousGoodsClass field if non-nil, zero value otherwise.

### GetDangerousGoodsClassOk

`func (o *WaybillLineDto) GetDangerousGoodsClassOk() (*string, bool)`

GetDangerousGoodsClassOk returns a tuple with the DangerousGoodsClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDangerousGoodsClass

`func (o *WaybillLineDto) SetDangerousGoodsClass(v string)`

SetDangerousGoodsClass sets DangerousGoodsClass field to given value.

### HasDangerousGoodsClass

`func (o *WaybillLineDto) HasDangerousGoodsClass() bool`

HasDangerousGoodsClass returns a boolean if a field has been set.

### SetDangerousGoodsClassNil

`func (o *WaybillLineDto) SetDangerousGoodsClassNil(b bool)`

 SetDangerousGoodsClassNil sets the value for DangerousGoodsClass to be an explicit nil

### UnsetDangerousGoodsClass
`func (o *WaybillLineDto) UnsetDangerousGoodsClass()`

UnsetDangerousGoodsClass ensures that no value is present for DangerousGoodsClass, not even an explicit nil
### GetUnHazmatNumber

`func (o *WaybillLineDto) GetUnHazmatNumber() string`

GetUnHazmatNumber returns the UnHazmatNumber field if non-nil, zero value otherwise.

### GetUnHazmatNumberOk

`func (o *WaybillLineDto) GetUnHazmatNumberOk() (*string, bool)`

GetUnHazmatNumberOk returns a tuple with the UnHazmatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnHazmatNumber

`func (o *WaybillLineDto) SetUnHazmatNumber(v string)`

SetUnHazmatNumber sets UnHazmatNumber field to given value.

### HasUnHazmatNumber

`func (o *WaybillLineDto) HasUnHazmatNumber() bool`

HasUnHazmatNumber returns a boolean if a field has been set.

### SetUnHazmatNumberNil

`func (o *WaybillLineDto) SetUnHazmatNumberNil(b bool)`

 SetUnHazmatNumberNil sets the value for UnHazmatNumber to be an explicit nil

### UnsetUnHazmatNumber
`func (o *WaybillLineDto) UnsetUnHazmatNumber()`

UnsetUnHazmatNumber ensures that no value is present for UnHazmatNumber, not even an explicit nil
### GetWagonNumber

`func (o *WaybillLineDto) GetWagonNumber() string`

GetWagonNumber returns the WagonNumber field if non-nil, zero value otherwise.

### GetWagonNumberOk

`func (o *WaybillLineDto) GetWagonNumberOk() (*string, bool)`

GetWagonNumberOk returns a tuple with the WagonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagonNumber

`func (o *WaybillLineDto) SetWagonNumber(v string)`

SetWagonNumber sets WagonNumber field to given value.

### HasWagonNumber

`func (o *WaybillLineDto) HasWagonNumber() bool`

HasWagonNumber returns a boolean if a field has been set.

### SetWagonNumberNil

`func (o *WaybillLineDto) SetWagonNumberNil(b bool)`

 SetWagonNumberNil sets the value for WagonNumber to be an explicit nil

### UnsetWagonNumber
`func (o *WaybillLineDto) UnsetWagonNumber()`

UnsetWagonNumber ensures that no value is present for WagonNumber, not even an explicit nil
### GetTenantId

`func (o *WaybillLineDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WaybillLineDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WaybillLineDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *WaybillLineDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *WaybillLineDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *WaybillLineDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


