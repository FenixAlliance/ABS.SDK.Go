# VesselDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVesselDto

`func NewVesselDto() *VesselDto`

NewVesselDto instantiates a new VesselDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVesselDtoWithDefaults

`func NewVesselDtoWithDefaults() *VesselDto`

NewVesselDtoWithDefaults instantiates a new VesselDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VesselDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VesselDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VesselDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VesselDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VesselDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VesselDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *VesselDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VesselDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VesselDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VesselDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *VesselDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *VesselDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *VesselDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VesselDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VesselDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VesselDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VesselDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VesselDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImoNumber

`func (o *VesselDto) GetImoNumber() string`

GetImoNumber returns the ImoNumber field if non-nil, zero value otherwise.

### GetImoNumberOk

`func (o *VesselDto) GetImoNumberOk() (*string, bool)`

GetImoNumberOk returns a tuple with the ImoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImoNumber

`func (o *VesselDto) SetImoNumber(v string)`

SetImoNumber sets ImoNumber field to given value.

### HasImoNumber

`func (o *VesselDto) HasImoNumber() bool`

HasImoNumber returns a boolean if a field has been set.

### SetImoNumberNil

`func (o *VesselDto) SetImoNumberNil(b bool)`

 SetImoNumberNil sets the value for ImoNumber to be an explicit nil

### UnsetImoNumber
`func (o *VesselDto) UnsetImoNumber()`

UnsetImoNumber ensures that no value is present for ImoNumber, not even an explicit nil
### GetMmsiNumber

`func (o *VesselDto) GetMmsiNumber() string`

GetMmsiNumber returns the MmsiNumber field if non-nil, zero value otherwise.

### GetMmsiNumberOk

`func (o *VesselDto) GetMmsiNumberOk() (*string, bool)`

GetMmsiNumberOk returns a tuple with the MmsiNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsiNumber

`func (o *VesselDto) SetMmsiNumber(v string)`

SetMmsiNumber sets MmsiNumber field to given value.

### HasMmsiNumber

`func (o *VesselDto) HasMmsiNumber() bool`

HasMmsiNumber returns a boolean if a field has been set.

### SetMmsiNumberNil

`func (o *VesselDto) SetMmsiNumberNil(b bool)`

 SetMmsiNumberNil sets the value for MmsiNumber to be an explicit nil

### UnsetMmsiNumber
`func (o *VesselDto) UnsetMmsiNumber()`

UnsetMmsiNumber ensures that no value is present for MmsiNumber, not even an explicit nil
### GetCallSign

`func (o *VesselDto) GetCallSign() string`

GetCallSign returns the CallSign field if non-nil, zero value otherwise.

### GetCallSignOk

`func (o *VesselDto) GetCallSignOk() (*string, bool)`

GetCallSignOk returns a tuple with the CallSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSign

`func (o *VesselDto) SetCallSign(v string)`

SetCallSign sets CallSign field to given value.

### HasCallSign

`func (o *VesselDto) HasCallSign() bool`

HasCallSign returns a boolean if a field has been set.

### SetCallSignNil

`func (o *VesselDto) SetCallSignNil(b bool)`

 SetCallSignNil sets the value for CallSign to be an explicit nil

### UnsetCallSign
`func (o *VesselDto) UnsetCallSign()`

UnsetCallSign ensures that no value is present for CallSign, not even an explicit nil
### GetFlagCountryId

`func (o *VesselDto) GetFlagCountryId() string`

GetFlagCountryId returns the FlagCountryId field if non-nil, zero value otherwise.

### GetFlagCountryIdOk

`func (o *VesselDto) GetFlagCountryIdOk() (*string, bool)`

GetFlagCountryIdOk returns a tuple with the FlagCountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCountryId

`func (o *VesselDto) SetFlagCountryId(v string)`

SetFlagCountryId sets FlagCountryId field to given value.

### HasFlagCountryId

`func (o *VesselDto) HasFlagCountryId() bool`

HasFlagCountryId returns a boolean if a field has been set.

### SetFlagCountryIdNil

`func (o *VesselDto) SetFlagCountryIdNil(b bool)`

 SetFlagCountryIdNil sets the value for FlagCountryId to be an explicit nil

### UnsetFlagCountryId
`func (o *VesselDto) UnsetFlagCountryId()`

UnsetFlagCountryId ensures that no value is present for FlagCountryId, not even an explicit nil
### GetVesselType

`func (o *VesselDto) GetVesselType() string`

GetVesselType returns the VesselType field if non-nil, zero value otherwise.

### GetVesselTypeOk

`func (o *VesselDto) GetVesselTypeOk() (*string, bool)`

GetVesselTypeOk returns a tuple with the VesselType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselType

`func (o *VesselDto) SetVesselType(v string)`

SetVesselType sets VesselType field to given value.

### HasVesselType

`func (o *VesselDto) HasVesselType() bool`

HasVesselType returns a boolean if a field has been set.

### SetVesselTypeNil

`func (o *VesselDto) SetVesselTypeNil(b bool)`

 SetVesselTypeNil sets the value for VesselType to be an explicit nil

### UnsetVesselType
`func (o *VesselDto) UnsetVesselType()`

UnsetVesselType ensures that no value is present for VesselType, not even an explicit nil
### GetVesselStatus

`func (o *VesselDto) GetVesselStatus() string`

GetVesselStatus returns the VesselStatus field if non-nil, zero value otherwise.

### GetVesselStatusOk

`func (o *VesselDto) GetVesselStatusOk() (*string, bool)`

GetVesselStatusOk returns a tuple with the VesselStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselStatus

`func (o *VesselDto) SetVesselStatus(v string)`

SetVesselStatus sets VesselStatus field to given value.

### HasVesselStatus

`func (o *VesselDto) HasVesselStatus() bool`

HasVesselStatus returns a boolean if a field has been set.

### SetVesselStatusNil

`func (o *VesselDto) SetVesselStatusNil(b bool)`

 SetVesselStatusNil sets the value for VesselStatus to be an explicit nil

### UnsetVesselStatus
`func (o *VesselDto) UnsetVesselStatus()`

UnsetVesselStatus ensures that no value is present for VesselStatus, not even an explicit nil
### GetGrossTonnage

`func (o *VesselDto) GetGrossTonnage() float64`

GetGrossTonnage returns the GrossTonnage field if non-nil, zero value otherwise.

### GetGrossTonnageOk

`func (o *VesselDto) GetGrossTonnageOk() (*float64, bool)`

GetGrossTonnageOk returns a tuple with the GrossTonnage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossTonnage

`func (o *VesselDto) SetGrossTonnage(v float64)`

SetGrossTonnage sets GrossTonnage field to given value.

### HasGrossTonnage

`func (o *VesselDto) HasGrossTonnage() bool`

HasGrossTonnage returns a boolean if a field has been set.

### SetGrossTonnageNil

`func (o *VesselDto) SetGrossTonnageNil(b bool)`

 SetGrossTonnageNil sets the value for GrossTonnage to be an explicit nil

### UnsetGrossTonnage
`func (o *VesselDto) UnsetGrossTonnage()`

UnsetGrossTonnage ensures that no value is present for GrossTonnage, not even an explicit nil
### GetDeadweightTonnage

`func (o *VesselDto) GetDeadweightTonnage() float64`

GetDeadweightTonnage returns the DeadweightTonnage field if non-nil, zero value otherwise.

### GetDeadweightTonnageOk

`func (o *VesselDto) GetDeadweightTonnageOk() (*float64, bool)`

GetDeadweightTonnageOk returns a tuple with the DeadweightTonnage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadweightTonnage

`func (o *VesselDto) SetDeadweightTonnage(v float64)`

SetDeadweightTonnage sets DeadweightTonnage field to given value.

### HasDeadweightTonnage

`func (o *VesselDto) HasDeadweightTonnage() bool`

HasDeadweightTonnage returns a boolean if a field has been set.

### SetDeadweightTonnageNil

`func (o *VesselDto) SetDeadweightTonnageNil(b bool)`

 SetDeadweightTonnageNil sets the value for DeadweightTonnage to be an explicit nil

### UnsetDeadweightTonnage
`func (o *VesselDto) UnsetDeadweightTonnage()`

UnsetDeadweightTonnage ensures that no value is present for DeadweightTonnage, not even an explicit nil
### GetTeuCapacity

`func (o *VesselDto) GetTeuCapacity() int32`

GetTeuCapacity returns the TeuCapacity field if non-nil, zero value otherwise.

### GetTeuCapacityOk

`func (o *VesselDto) GetTeuCapacityOk() (*int32, bool)`

GetTeuCapacityOk returns a tuple with the TeuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeuCapacity

`func (o *VesselDto) SetTeuCapacity(v int32)`

SetTeuCapacity sets TeuCapacity field to given value.

### HasTeuCapacity

`func (o *VesselDto) HasTeuCapacity() bool`

HasTeuCapacity returns a boolean if a field has been set.

### SetTeuCapacityNil

`func (o *VesselDto) SetTeuCapacityNil(b bool)`

 SetTeuCapacityNil sets the value for TeuCapacity to be an explicit nil

### UnsetTeuCapacity
`func (o *VesselDto) UnsetTeuCapacity()`

UnsetTeuCapacity ensures that no value is present for TeuCapacity, not even an explicit nil
### GetLengthMeters

`func (o *VesselDto) GetLengthMeters() float64`

GetLengthMeters returns the LengthMeters field if non-nil, zero value otherwise.

### GetLengthMetersOk

`func (o *VesselDto) GetLengthMetersOk() (*float64, bool)`

GetLengthMetersOk returns a tuple with the LengthMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthMeters

`func (o *VesselDto) SetLengthMeters(v float64)`

SetLengthMeters sets LengthMeters field to given value.

### HasLengthMeters

`func (o *VesselDto) HasLengthMeters() bool`

HasLengthMeters returns a boolean if a field has been set.

### SetLengthMetersNil

`func (o *VesselDto) SetLengthMetersNil(b bool)`

 SetLengthMetersNil sets the value for LengthMeters to be an explicit nil

### UnsetLengthMeters
`func (o *VesselDto) UnsetLengthMeters()`

UnsetLengthMeters ensures that no value is present for LengthMeters, not even an explicit nil
### GetBeamMeters

`func (o *VesselDto) GetBeamMeters() float64`

GetBeamMeters returns the BeamMeters field if non-nil, zero value otherwise.

### GetBeamMetersOk

`func (o *VesselDto) GetBeamMetersOk() (*float64, bool)`

GetBeamMetersOk returns a tuple with the BeamMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeamMeters

`func (o *VesselDto) SetBeamMeters(v float64)`

SetBeamMeters sets BeamMeters field to given value.

### HasBeamMeters

`func (o *VesselDto) HasBeamMeters() bool`

HasBeamMeters returns a boolean if a field has been set.

### SetBeamMetersNil

`func (o *VesselDto) SetBeamMetersNil(b bool)`

 SetBeamMetersNil sets the value for BeamMeters to be an explicit nil

### UnsetBeamMeters
`func (o *VesselDto) UnsetBeamMeters()`

UnsetBeamMeters ensures that no value is present for BeamMeters, not even an explicit nil
### GetDraftMeters

`func (o *VesselDto) GetDraftMeters() float64`

GetDraftMeters returns the DraftMeters field if non-nil, zero value otherwise.

### GetDraftMetersOk

`func (o *VesselDto) GetDraftMetersOk() (*float64, bool)`

GetDraftMetersOk returns a tuple with the DraftMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftMeters

`func (o *VesselDto) SetDraftMeters(v float64)`

SetDraftMeters sets DraftMeters field to given value.

### HasDraftMeters

`func (o *VesselDto) HasDraftMeters() bool`

HasDraftMeters returns a boolean if a field has been set.

### SetDraftMetersNil

`func (o *VesselDto) SetDraftMetersNil(b bool)`

 SetDraftMetersNil sets the value for DraftMeters to be an explicit nil

### UnsetDraftMeters
`func (o *VesselDto) UnsetDraftMeters()`

UnsetDraftMeters ensures that no value is present for DraftMeters, not even an explicit nil
### GetYearBuilt

`func (o *VesselDto) GetYearBuilt() int32`

GetYearBuilt returns the YearBuilt field if non-nil, zero value otherwise.

### GetYearBuiltOk

`func (o *VesselDto) GetYearBuiltOk() (*int32, bool)`

GetYearBuiltOk returns a tuple with the YearBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearBuilt

`func (o *VesselDto) SetYearBuilt(v int32)`

SetYearBuilt sets YearBuilt field to given value.

### HasYearBuilt

`func (o *VesselDto) HasYearBuilt() bool`

HasYearBuilt returns a boolean if a field has been set.

### SetYearBuiltNil

`func (o *VesselDto) SetYearBuiltNil(b bool)`

 SetYearBuiltNil sets the value for YearBuilt to be an explicit nil

### UnsetYearBuilt
`func (o *VesselDto) UnsetYearBuilt()`

UnsetYearBuilt ensures that no value is present for YearBuilt, not even an explicit nil
### GetShippingCourierId

`func (o *VesselDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *VesselDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *VesselDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *VesselDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *VesselDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *VesselDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil
### GetTenantId

`func (o *VesselDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VesselDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VesselDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *VesselDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *VesselDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *VesselDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *VesselDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *VesselDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *VesselDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *VesselDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *VesselDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *VesselDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


