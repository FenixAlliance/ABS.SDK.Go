# VesselCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewVesselCreateDto

`func NewVesselCreateDto() *VesselCreateDto`

NewVesselCreateDto instantiates a new VesselCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVesselCreateDtoWithDefaults

`func NewVesselCreateDtoWithDefaults() *VesselCreateDto`

NewVesselCreateDtoWithDefaults instantiates a new VesselCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VesselCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VesselCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VesselCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VesselCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *VesselCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VesselCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VesselCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VesselCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *VesselCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VesselCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VesselCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VesselCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VesselCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VesselCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImoNumber

`func (o *VesselCreateDto) GetImoNumber() string`

GetImoNumber returns the ImoNumber field if non-nil, zero value otherwise.

### GetImoNumberOk

`func (o *VesselCreateDto) GetImoNumberOk() (*string, bool)`

GetImoNumberOk returns a tuple with the ImoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImoNumber

`func (o *VesselCreateDto) SetImoNumber(v string)`

SetImoNumber sets ImoNumber field to given value.

### HasImoNumber

`func (o *VesselCreateDto) HasImoNumber() bool`

HasImoNumber returns a boolean if a field has been set.

### SetImoNumberNil

`func (o *VesselCreateDto) SetImoNumberNil(b bool)`

 SetImoNumberNil sets the value for ImoNumber to be an explicit nil

### UnsetImoNumber
`func (o *VesselCreateDto) UnsetImoNumber()`

UnsetImoNumber ensures that no value is present for ImoNumber, not even an explicit nil
### GetMmsiNumber

`func (o *VesselCreateDto) GetMmsiNumber() string`

GetMmsiNumber returns the MmsiNumber field if non-nil, zero value otherwise.

### GetMmsiNumberOk

`func (o *VesselCreateDto) GetMmsiNumberOk() (*string, bool)`

GetMmsiNumberOk returns a tuple with the MmsiNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsiNumber

`func (o *VesselCreateDto) SetMmsiNumber(v string)`

SetMmsiNumber sets MmsiNumber field to given value.

### HasMmsiNumber

`func (o *VesselCreateDto) HasMmsiNumber() bool`

HasMmsiNumber returns a boolean if a field has been set.

### SetMmsiNumberNil

`func (o *VesselCreateDto) SetMmsiNumberNil(b bool)`

 SetMmsiNumberNil sets the value for MmsiNumber to be an explicit nil

### UnsetMmsiNumber
`func (o *VesselCreateDto) UnsetMmsiNumber()`

UnsetMmsiNumber ensures that no value is present for MmsiNumber, not even an explicit nil
### GetCallSign

`func (o *VesselCreateDto) GetCallSign() string`

GetCallSign returns the CallSign field if non-nil, zero value otherwise.

### GetCallSignOk

`func (o *VesselCreateDto) GetCallSignOk() (*string, bool)`

GetCallSignOk returns a tuple with the CallSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSign

`func (o *VesselCreateDto) SetCallSign(v string)`

SetCallSign sets CallSign field to given value.

### HasCallSign

`func (o *VesselCreateDto) HasCallSign() bool`

HasCallSign returns a boolean if a field has been set.

### SetCallSignNil

`func (o *VesselCreateDto) SetCallSignNil(b bool)`

 SetCallSignNil sets the value for CallSign to be an explicit nil

### UnsetCallSign
`func (o *VesselCreateDto) UnsetCallSign()`

UnsetCallSign ensures that no value is present for CallSign, not even an explicit nil
### GetFlagCountryId

`func (o *VesselCreateDto) GetFlagCountryId() string`

GetFlagCountryId returns the FlagCountryId field if non-nil, zero value otherwise.

### GetFlagCountryIdOk

`func (o *VesselCreateDto) GetFlagCountryIdOk() (*string, bool)`

GetFlagCountryIdOk returns a tuple with the FlagCountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCountryId

`func (o *VesselCreateDto) SetFlagCountryId(v string)`

SetFlagCountryId sets FlagCountryId field to given value.

### HasFlagCountryId

`func (o *VesselCreateDto) HasFlagCountryId() bool`

HasFlagCountryId returns a boolean if a field has been set.

### SetFlagCountryIdNil

`func (o *VesselCreateDto) SetFlagCountryIdNil(b bool)`

 SetFlagCountryIdNil sets the value for FlagCountryId to be an explicit nil

### UnsetFlagCountryId
`func (o *VesselCreateDto) UnsetFlagCountryId()`

UnsetFlagCountryId ensures that no value is present for FlagCountryId, not even an explicit nil
### GetVesselType

`func (o *VesselCreateDto) GetVesselType() string`

GetVesselType returns the VesselType field if non-nil, zero value otherwise.

### GetVesselTypeOk

`func (o *VesselCreateDto) GetVesselTypeOk() (*string, bool)`

GetVesselTypeOk returns a tuple with the VesselType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselType

`func (o *VesselCreateDto) SetVesselType(v string)`

SetVesselType sets VesselType field to given value.

### HasVesselType

`func (o *VesselCreateDto) HasVesselType() bool`

HasVesselType returns a boolean if a field has been set.

### SetVesselTypeNil

`func (o *VesselCreateDto) SetVesselTypeNil(b bool)`

 SetVesselTypeNil sets the value for VesselType to be an explicit nil

### UnsetVesselType
`func (o *VesselCreateDto) UnsetVesselType()`

UnsetVesselType ensures that no value is present for VesselType, not even an explicit nil
### GetVesselStatus

`func (o *VesselCreateDto) GetVesselStatus() string`

GetVesselStatus returns the VesselStatus field if non-nil, zero value otherwise.

### GetVesselStatusOk

`func (o *VesselCreateDto) GetVesselStatusOk() (*string, bool)`

GetVesselStatusOk returns a tuple with the VesselStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselStatus

`func (o *VesselCreateDto) SetVesselStatus(v string)`

SetVesselStatus sets VesselStatus field to given value.

### HasVesselStatus

`func (o *VesselCreateDto) HasVesselStatus() bool`

HasVesselStatus returns a boolean if a field has been set.

### SetVesselStatusNil

`func (o *VesselCreateDto) SetVesselStatusNil(b bool)`

 SetVesselStatusNil sets the value for VesselStatus to be an explicit nil

### UnsetVesselStatus
`func (o *VesselCreateDto) UnsetVesselStatus()`

UnsetVesselStatus ensures that no value is present for VesselStatus, not even an explicit nil
### GetGrossTonnage

`func (o *VesselCreateDto) GetGrossTonnage() float64`

GetGrossTonnage returns the GrossTonnage field if non-nil, zero value otherwise.

### GetGrossTonnageOk

`func (o *VesselCreateDto) GetGrossTonnageOk() (*float64, bool)`

GetGrossTonnageOk returns a tuple with the GrossTonnage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossTonnage

`func (o *VesselCreateDto) SetGrossTonnage(v float64)`

SetGrossTonnage sets GrossTonnage field to given value.

### HasGrossTonnage

`func (o *VesselCreateDto) HasGrossTonnage() bool`

HasGrossTonnage returns a boolean if a field has been set.

### SetGrossTonnageNil

`func (o *VesselCreateDto) SetGrossTonnageNil(b bool)`

 SetGrossTonnageNil sets the value for GrossTonnage to be an explicit nil

### UnsetGrossTonnage
`func (o *VesselCreateDto) UnsetGrossTonnage()`

UnsetGrossTonnage ensures that no value is present for GrossTonnage, not even an explicit nil
### GetDeadweightTonnage

`func (o *VesselCreateDto) GetDeadweightTonnage() float64`

GetDeadweightTonnage returns the DeadweightTonnage field if non-nil, zero value otherwise.

### GetDeadweightTonnageOk

`func (o *VesselCreateDto) GetDeadweightTonnageOk() (*float64, bool)`

GetDeadweightTonnageOk returns a tuple with the DeadweightTonnage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadweightTonnage

`func (o *VesselCreateDto) SetDeadweightTonnage(v float64)`

SetDeadweightTonnage sets DeadweightTonnage field to given value.

### HasDeadweightTonnage

`func (o *VesselCreateDto) HasDeadweightTonnage() bool`

HasDeadweightTonnage returns a boolean if a field has been set.

### SetDeadweightTonnageNil

`func (o *VesselCreateDto) SetDeadweightTonnageNil(b bool)`

 SetDeadweightTonnageNil sets the value for DeadweightTonnage to be an explicit nil

### UnsetDeadweightTonnage
`func (o *VesselCreateDto) UnsetDeadweightTonnage()`

UnsetDeadweightTonnage ensures that no value is present for DeadweightTonnage, not even an explicit nil
### GetTeuCapacity

`func (o *VesselCreateDto) GetTeuCapacity() int32`

GetTeuCapacity returns the TeuCapacity field if non-nil, zero value otherwise.

### GetTeuCapacityOk

`func (o *VesselCreateDto) GetTeuCapacityOk() (*int32, bool)`

GetTeuCapacityOk returns a tuple with the TeuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeuCapacity

`func (o *VesselCreateDto) SetTeuCapacity(v int32)`

SetTeuCapacity sets TeuCapacity field to given value.

### HasTeuCapacity

`func (o *VesselCreateDto) HasTeuCapacity() bool`

HasTeuCapacity returns a boolean if a field has been set.

### SetTeuCapacityNil

`func (o *VesselCreateDto) SetTeuCapacityNil(b bool)`

 SetTeuCapacityNil sets the value for TeuCapacity to be an explicit nil

### UnsetTeuCapacity
`func (o *VesselCreateDto) UnsetTeuCapacity()`

UnsetTeuCapacity ensures that no value is present for TeuCapacity, not even an explicit nil
### GetLengthMeters

`func (o *VesselCreateDto) GetLengthMeters() float64`

GetLengthMeters returns the LengthMeters field if non-nil, zero value otherwise.

### GetLengthMetersOk

`func (o *VesselCreateDto) GetLengthMetersOk() (*float64, bool)`

GetLengthMetersOk returns a tuple with the LengthMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthMeters

`func (o *VesselCreateDto) SetLengthMeters(v float64)`

SetLengthMeters sets LengthMeters field to given value.

### HasLengthMeters

`func (o *VesselCreateDto) HasLengthMeters() bool`

HasLengthMeters returns a boolean if a field has been set.

### SetLengthMetersNil

`func (o *VesselCreateDto) SetLengthMetersNil(b bool)`

 SetLengthMetersNil sets the value for LengthMeters to be an explicit nil

### UnsetLengthMeters
`func (o *VesselCreateDto) UnsetLengthMeters()`

UnsetLengthMeters ensures that no value is present for LengthMeters, not even an explicit nil
### GetBeamMeters

`func (o *VesselCreateDto) GetBeamMeters() float64`

GetBeamMeters returns the BeamMeters field if non-nil, zero value otherwise.

### GetBeamMetersOk

`func (o *VesselCreateDto) GetBeamMetersOk() (*float64, bool)`

GetBeamMetersOk returns a tuple with the BeamMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamMeters

`func (o *VesselCreateDto) SetBeamMeters(v float64)`

SetBeamMeters sets BeamMeters field to given value.

### HasBeamMeters

`func (o *VesselCreateDto) HasBeamMeters() bool`

HasBeamMeters returns a boolean if a field has been set.

### SetBeamMetersNil

`func (o *VesselCreateDto) SetBeamMetersNil(b bool)`

 SetBeamMetersNil sets the value for BeamMeters to be an explicit nil

### UnsetBeamMeters
`func (o *VesselCreateDto) UnsetBeamMeters()`

UnsetBeamMeters ensures that no value is present for BeamMeters, not even an explicit nil
### GetDraftMeters

`func (o *VesselCreateDto) GetDraftMeters() float64`

GetDraftMeters returns the DraftMeters field if non-nil, zero value otherwise.

### GetDraftMetersOk

`func (o *VesselCreateDto) GetDraftMetersOk() (*float64, bool)`

GetDraftMetersOk returns a tuple with the DraftMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftMeters

`func (o *VesselCreateDto) SetDraftMeters(v float64)`

SetDraftMeters sets DraftMeters field to given value.

### HasDraftMeters

`func (o *VesselCreateDto) HasDraftMeters() bool`

HasDraftMeters returns a boolean if a field has been set.

### SetDraftMetersNil

`func (o *VesselCreateDto) SetDraftMetersNil(b bool)`

 SetDraftMetersNil sets the value for DraftMeters to be an explicit nil

### UnsetDraftMeters
`func (o *VesselCreateDto) UnsetDraftMeters()`

UnsetDraftMeters ensures that no value is present for DraftMeters, not even an explicit nil
### GetYearBuilt

`func (o *VesselCreateDto) GetYearBuilt() int32`

GetYearBuilt returns the YearBuilt field if non-nil, zero value otherwise.

### GetYearBuiltOk

`func (o *VesselCreateDto) GetYearBuiltOk() (*int32, bool)`

GetYearBuiltOk returns a tuple with the YearBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearBuilt

`func (o *VesselCreateDto) SetYearBuilt(v int32)`

SetYearBuilt sets YearBuilt field to given value.

### HasYearBuilt

`func (o *VesselCreateDto) HasYearBuilt() bool`

HasYearBuilt returns a boolean if a field has been set.

### SetYearBuiltNil

`func (o *VesselCreateDto) SetYearBuiltNil(b bool)`

 SetYearBuiltNil sets the value for YearBuilt to be an explicit nil

### UnsetYearBuilt
`func (o *VesselCreateDto) UnsetYearBuilt()`

UnsetYearBuilt ensures that no value is present for YearBuilt, not even an explicit nil
### GetShippingCourierId

`func (o *VesselCreateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *VesselCreateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *VesselCreateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *VesselCreateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *VesselCreateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *VesselCreateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


