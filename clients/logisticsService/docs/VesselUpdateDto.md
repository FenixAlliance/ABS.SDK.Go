# VesselUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ImoNumber** | Pointer to **NullableString** |  | [optional] 
**MmsiNumber** | Pointer to **NullableString** |  | [optional] 
**CallSign** | Pointer to **NullableString** |  | [optional] 
**FlagCountryId** | Pointer to **NullableString** |  | [optional] 
**VesselType** | Pointer to **NullableString** |  | [optional] 
**VesselStatus** | Pointer to **NullableString** |  | [optional] 
**GrossTonnage** | Pointer to **NullableFloat64** |  | [optional] 
**DeadweightTonnage** | Pointer to **NullableFloat64** |  | [optional] 
**TeuCapacity** | Pointer to **NullableInt32** |  | [optional] 
**LengthMeters** | Pointer to **NullableFloat64** |  | [optional] 
**BeamMeters** | Pointer to **NullableFloat64** |  | [optional] 
**DraftMeters** | Pointer to **NullableFloat64** |  | [optional] 
**YearBuilt** | Pointer to **NullableInt32** |  | [optional] 
**ShippingCourierId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVesselUpdateDto

`func NewVesselUpdateDto() *VesselUpdateDto`

NewVesselUpdateDto instantiates a new VesselUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVesselUpdateDtoWithDefaults

`func NewVesselUpdateDtoWithDefaults() *VesselUpdateDto`

NewVesselUpdateDtoWithDefaults instantiates a new VesselUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VesselUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VesselUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VesselUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VesselUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VesselUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VesselUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImoNumber

`func (o *VesselUpdateDto) GetImoNumber() string`

GetImoNumber returns the ImoNumber field if non-nil, zero value otherwise.

### GetImoNumberOk

`func (o *VesselUpdateDto) GetImoNumberOk() (*string, bool)`

GetImoNumberOk returns a tuple with the ImoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImoNumber

`func (o *VesselUpdateDto) SetImoNumber(v string)`

SetImoNumber sets ImoNumber field to given value.

### HasImoNumber

`func (o *VesselUpdateDto) HasImoNumber() bool`

HasImoNumber returns a boolean if a field has been set.

### SetImoNumberNil

`func (o *VesselUpdateDto) SetImoNumberNil(b bool)`

 SetImoNumberNil sets the value for ImoNumber to be an explicit nil

### UnsetImoNumber
`func (o *VesselUpdateDto) UnsetImoNumber()`

UnsetImoNumber ensures that no value is present for ImoNumber, not even an explicit nil
### GetMmsiNumber

`func (o *VesselUpdateDto) GetMmsiNumber() string`

GetMmsiNumber returns the MmsiNumber field if non-nil, zero value otherwise.

### GetMmsiNumberOk

`func (o *VesselUpdateDto) GetMmsiNumberOk() (*string, bool)`

GetMmsiNumberOk returns a tuple with the MmsiNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsiNumber

`func (o *VesselUpdateDto) SetMmsiNumber(v string)`

SetMmsiNumber sets MmsiNumber field to given value.

### HasMmsiNumber

`func (o *VesselUpdateDto) HasMmsiNumber() bool`

HasMmsiNumber returns a boolean if a field has been set.

### SetMmsiNumberNil

`func (o *VesselUpdateDto) SetMmsiNumberNil(b bool)`

 SetMmsiNumberNil sets the value for MmsiNumber to be an explicit nil

### UnsetMmsiNumber
`func (o *VesselUpdateDto) UnsetMmsiNumber()`

UnsetMmsiNumber ensures that no value is present for MmsiNumber, not even an explicit nil
### GetCallSign

`func (o *VesselUpdateDto) GetCallSign() string`

GetCallSign returns the CallSign field if non-nil, zero value otherwise.

### GetCallSignOk

`func (o *VesselUpdateDto) GetCallSignOk() (*string, bool)`

GetCallSignOk returns a tuple with the CallSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSign

`func (o *VesselUpdateDto) SetCallSign(v string)`

SetCallSign sets CallSign field to given value.

### HasCallSign

`func (o *VesselUpdateDto) HasCallSign() bool`

HasCallSign returns a boolean if a field has been set.

### SetCallSignNil

`func (o *VesselUpdateDto) SetCallSignNil(b bool)`

 SetCallSignNil sets the value for CallSign to be an explicit nil

### UnsetCallSign
`func (o *VesselUpdateDto) UnsetCallSign()`

UnsetCallSign ensures that no value is present for CallSign, not even an explicit nil
### GetFlagCountryId

`func (o *VesselUpdateDto) GetFlagCountryId() string`

GetFlagCountryId returns the FlagCountryId field if non-nil, zero value otherwise.

### GetFlagCountryIdOk

`func (o *VesselUpdateDto) GetFlagCountryIdOk() (*string, bool)`

GetFlagCountryIdOk returns a tuple with the FlagCountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCountryId

`func (o *VesselUpdateDto) SetFlagCountryId(v string)`

SetFlagCountryId sets FlagCountryId field to given value.

### HasFlagCountryId

`func (o *VesselUpdateDto) HasFlagCountryId() bool`

HasFlagCountryId returns a boolean if a field has been set.

### SetFlagCountryIdNil

`func (o *VesselUpdateDto) SetFlagCountryIdNil(b bool)`

 SetFlagCountryIdNil sets the value for FlagCountryId to be an explicit nil

### UnsetFlagCountryId
`func (o *VesselUpdateDto) UnsetFlagCountryId()`

UnsetFlagCountryId ensures that no value is present for FlagCountryId, not even an explicit nil
### GetVesselType

`func (o *VesselUpdateDto) GetVesselType() string`

GetVesselType returns the VesselType field if non-nil, zero value otherwise.

### GetVesselTypeOk

`func (o *VesselUpdateDto) GetVesselTypeOk() (*string, bool)`

GetVesselTypeOk returns a tuple with the VesselType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselType

`func (o *VesselUpdateDto) SetVesselType(v string)`

SetVesselType sets VesselType field to given value.

### HasVesselType

`func (o *VesselUpdateDto) HasVesselType() bool`

HasVesselType returns a boolean if a field has been set.

### SetVesselTypeNil

`func (o *VesselUpdateDto) SetVesselTypeNil(b bool)`

 SetVesselTypeNil sets the value for VesselType to be an explicit nil

### UnsetVesselType
`func (o *VesselUpdateDto) UnsetVesselType()`

UnsetVesselType ensures that no value is present for VesselType, not even an explicit nil
### GetVesselStatus

`func (o *VesselUpdateDto) GetVesselStatus() string`

GetVesselStatus returns the VesselStatus field if non-nil, zero value otherwise.

### GetVesselStatusOk

`func (o *VesselUpdateDto) GetVesselStatusOk() (*string, bool)`

GetVesselStatusOk returns a tuple with the VesselStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselStatus

`func (o *VesselUpdateDto) SetVesselStatus(v string)`

SetVesselStatus sets VesselStatus field to given value.

### HasVesselStatus

`func (o *VesselUpdateDto) HasVesselStatus() bool`

HasVesselStatus returns a boolean if a field has been set.

### SetVesselStatusNil

`func (o *VesselUpdateDto) SetVesselStatusNil(b bool)`

 SetVesselStatusNil sets the value for VesselStatus to be an explicit nil

### UnsetVesselStatus
`func (o *VesselUpdateDto) UnsetVesselStatus()`

UnsetVesselStatus ensures that no value is present for VesselStatus, not even an explicit nil
### GetGrossTonnage

`func (o *VesselUpdateDto) GetGrossTonnage() float64`

GetGrossTonnage returns the GrossTonnage field if non-nil, zero value otherwise.

### GetGrossTonnageOk

`func (o *VesselUpdateDto) GetGrossTonnageOk() (*float64, bool)`

GetGrossTonnageOk returns a tuple with the GrossTonnage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossTonnage

`func (o *VesselUpdateDto) SetGrossTonnage(v float64)`

SetGrossTonnage sets GrossTonnage field to given value.

### HasGrossTonnage

`func (o *VesselUpdateDto) HasGrossTonnage() bool`

HasGrossTonnage returns a boolean if a field has been set.

### SetGrossTonnageNil

`func (o *VesselUpdateDto) SetGrossTonnageNil(b bool)`

 SetGrossTonnageNil sets the value for GrossTonnage to be an explicit nil

### UnsetGrossTonnage
`func (o *VesselUpdateDto) UnsetGrossTonnage()`

UnsetGrossTonnage ensures that no value is present for GrossTonnage, not even an explicit nil
### GetDeadweightTonnage

`func (o *VesselUpdateDto) GetDeadweightTonnage() float64`

GetDeadweightTonnage returns the DeadweightTonnage field if non-nil, zero value otherwise.

### GetDeadweightTonnageOk

`func (o *VesselUpdateDto) GetDeadweightTonnageOk() (*float64, bool)`

GetDeadweightTonnageOk returns a tuple with the DeadweightTonnage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadweightTonnage

`func (o *VesselUpdateDto) SetDeadweightTonnage(v float64)`

SetDeadweightTonnage sets DeadweightTonnage field to given value.

### HasDeadweightTonnage

`func (o *VesselUpdateDto) HasDeadweightTonnage() bool`

HasDeadweightTonnage returns a boolean if a field has been set.

### SetDeadweightTonnageNil

`func (o *VesselUpdateDto) SetDeadweightTonnageNil(b bool)`

 SetDeadweightTonnageNil sets the value for DeadweightTonnage to be an explicit nil

### UnsetDeadweightTonnage
`func (o *VesselUpdateDto) UnsetDeadweightTonnage()`

UnsetDeadweightTonnage ensures that no value is present for DeadweightTonnage, not even an explicit nil
### GetTeuCapacity

`func (o *VesselUpdateDto) GetTeuCapacity() int32`

GetTeuCapacity returns the TeuCapacity field if non-nil, zero value otherwise.

### GetTeuCapacityOk

`func (o *VesselUpdateDto) GetTeuCapacityOk() (*int32, bool)`

GetTeuCapacityOk returns a tuple with the TeuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeuCapacity

`func (o *VesselUpdateDto) SetTeuCapacity(v int32)`

SetTeuCapacity sets TeuCapacity field to given value.

### HasTeuCapacity

`func (o *VesselUpdateDto) HasTeuCapacity() bool`

HasTeuCapacity returns a boolean if a field has been set.

### SetTeuCapacityNil

`func (o *VesselUpdateDto) SetTeuCapacityNil(b bool)`

 SetTeuCapacityNil sets the value for TeuCapacity to be an explicit nil

### UnsetTeuCapacity
`func (o *VesselUpdateDto) UnsetTeuCapacity()`

UnsetTeuCapacity ensures that no value is present for TeuCapacity, not even an explicit nil
### GetLengthMeters

`func (o *VesselUpdateDto) GetLengthMeters() float64`

GetLengthMeters returns the LengthMeters field if non-nil, zero value otherwise.

### GetLengthMetersOk

`func (o *VesselUpdateDto) GetLengthMetersOk() (*float64, bool)`

GetLengthMetersOk returns a tuple with the LengthMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthMeters

`func (o *VesselUpdateDto) SetLengthMeters(v float64)`

SetLengthMeters sets LengthMeters field to given value.

### HasLengthMeters

`func (o *VesselUpdateDto) HasLengthMeters() bool`

HasLengthMeters returns a boolean if a field has been set.

### SetLengthMetersNil

`func (o *VesselUpdateDto) SetLengthMetersNil(b bool)`

 SetLengthMetersNil sets the value for LengthMeters to be an explicit nil

### UnsetLengthMeters
`func (o *VesselUpdateDto) UnsetLengthMeters()`

UnsetLengthMeters ensures that no value is present for LengthMeters, not even an explicit nil
### GetBeamMeters

`func (o *VesselUpdateDto) GetBeamMeters() float64`

GetBeamMeters returns the BeamMeters field if non-nil, zero value otherwise.

### GetBeamMetersOk

`func (o *VesselUpdateDto) GetBeamMetersOk() (*float64, bool)`

GetBeamMetersOk returns a tuple with the BeamMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamMeters

`func (o *VesselUpdateDto) SetBeamMeters(v float64)`

SetBeamMeters sets BeamMeters field to given value.

### HasBeamMeters

`func (o *VesselUpdateDto) HasBeamMeters() bool`

HasBeamMeters returns a boolean if a field has been set.

### SetBeamMetersNil

`func (o *VesselUpdateDto) SetBeamMetersNil(b bool)`

 SetBeamMetersNil sets the value for BeamMeters to be an explicit nil

### UnsetBeamMeters
`func (o *VesselUpdateDto) UnsetBeamMeters()`

UnsetBeamMeters ensures that no value is present for BeamMeters, not even an explicit nil
### GetDraftMeters

`func (o *VesselUpdateDto) GetDraftMeters() float64`

GetDraftMeters returns the DraftMeters field if non-nil, zero value otherwise.

### GetDraftMetersOk

`func (o *VesselUpdateDto) GetDraftMetersOk() (*float64, bool)`

GetDraftMetersOk returns a tuple with the DraftMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftMeters

`func (o *VesselUpdateDto) SetDraftMeters(v float64)`

SetDraftMeters sets DraftMeters field to given value.

### HasDraftMeters

`func (o *VesselUpdateDto) HasDraftMeters() bool`

HasDraftMeters returns a boolean if a field has been set.

### SetDraftMetersNil

`func (o *VesselUpdateDto) SetDraftMetersNil(b bool)`

 SetDraftMetersNil sets the value for DraftMeters to be an explicit nil

### UnsetDraftMeters
`func (o *VesselUpdateDto) UnsetDraftMeters()`

UnsetDraftMeters ensures that no value is present for DraftMeters, not even an explicit nil
### GetYearBuilt

`func (o *VesselUpdateDto) GetYearBuilt() int32`

GetYearBuilt returns the YearBuilt field if non-nil, zero value otherwise.

### GetYearBuiltOk

`func (o *VesselUpdateDto) GetYearBuiltOk() (*int32, bool)`

GetYearBuiltOk returns a tuple with the YearBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearBuilt

`func (o *VesselUpdateDto) SetYearBuilt(v int32)`

SetYearBuilt sets YearBuilt field to given value.

### HasYearBuilt

`func (o *VesselUpdateDto) HasYearBuilt() bool`

HasYearBuilt returns a boolean if a field has been set.

### SetYearBuiltNil

`func (o *VesselUpdateDto) SetYearBuiltNil(b bool)`

 SetYearBuiltNil sets the value for YearBuilt to be an explicit nil

### UnsetYearBuilt
`func (o *VesselUpdateDto) UnsetYearBuilt()`

UnsetYearBuilt ensures that no value is present for YearBuilt, not even an explicit nil
### GetShippingCourierId

`func (o *VesselUpdateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *VesselUpdateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *VesselUpdateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *VesselUpdateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *VesselUpdateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *VesselUpdateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


