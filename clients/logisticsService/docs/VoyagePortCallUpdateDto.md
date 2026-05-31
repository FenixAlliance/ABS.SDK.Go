# VoyagePortCallUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceNumber** | Pointer to **int32** |  | [optional] 
**PortCallStatus** | Pointer to **NullableString** |  | [optional] 
**Eta** | Pointer to **NullableTime** |  | [optional] 
**Ata** | Pointer to **NullableTime** |  | [optional] 
**Etd** | Pointer to **NullableTime** |  | [optional] 
**Atd** | Pointer to **NullableTime** |  | [optional] 
**BerthNumber** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**PortId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVoyagePortCallUpdateDto

`func NewVoyagePortCallUpdateDto() *VoyagePortCallUpdateDto`

NewVoyagePortCallUpdateDto instantiates a new VoyagePortCallUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoyagePortCallUpdateDtoWithDefaults

`func NewVoyagePortCallUpdateDtoWithDefaults() *VoyagePortCallUpdateDto`

NewVoyagePortCallUpdateDtoWithDefaults instantiates a new VoyagePortCallUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceNumber

`func (o *VoyagePortCallUpdateDto) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *VoyagePortCallUpdateDto) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *VoyagePortCallUpdateDto) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *VoyagePortCallUpdateDto) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetPortCallStatus

`func (o *VoyagePortCallUpdateDto) GetPortCallStatus() string`

GetPortCallStatus returns the PortCallStatus field if non-nil, zero value otherwise.

### GetPortCallStatusOk

`func (o *VoyagePortCallUpdateDto) GetPortCallStatusOk() (*string, bool)`

GetPortCallStatusOk returns a tuple with the PortCallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCallStatus

`func (o *VoyagePortCallUpdateDto) SetPortCallStatus(v string)`

SetPortCallStatus sets PortCallStatus field to given value.

### HasPortCallStatus

`func (o *VoyagePortCallUpdateDto) HasPortCallStatus() bool`

HasPortCallStatus returns a boolean if a field has been set.

### SetPortCallStatusNil

`func (o *VoyagePortCallUpdateDto) SetPortCallStatusNil(b bool)`

 SetPortCallStatusNil sets the value for PortCallStatus to be an explicit nil

### UnsetPortCallStatus
`func (o *VoyagePortCallUpdateDto) UnsetPortCallStatus()`

UnsetPortCallStatus ensures that no value is present for PortCallStatus, not even an explicit nil
### GetEta

`func (o *VoyagePortCallUpdateDto) GetEta() time.Time`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *VoyagePortCallUpdateDto) GetEtaOk() (*time.Time, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *VoyagePortCallUpdateDto) SetEta(v time.Time)`

SetEta sets Eta field to given value.

### HasEta

`func (o *VoyagePortCallUpdateDto) HasEta() bool`

HasEta returns a boolean if a field has been set.

### SetEtaNil

`func (o *VoyagePortCallUpdateDto) SetEtaNil(b bool)`

 SetEtaNil sets the value for Eta to be an explicit nil

### UnsetEta
`func (o *VoyagePortCallUpdateDto) UnsetEta()`

UnsetEta ensures that no value is present for Eta, not even an explicit nil
### GetAta

`func (o *VoyagePortCallUpdateDto) GetAta() time.Time`

GetAta returns the Ata field if non-nil, zero value otherwise.

### GetAtaOk

`func (o *VoyagePortCallUpdateDto) GetAtaOk() (*time.Time, bool)`

GetAtaOk returns a tuple with the Ata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAta

`func (o *VoyagePortCallUpdateDto) SetAta(v time.Time)`

SetAta sets Ata field to given value.

### HasAta

`func (o *VoyagePortCallUpdateDto) HasAta() bool`

HasAta returns a boolean if a field has been set.

### SetAtaNil

`func (o *VoyagePortCallUpdateDto) SetAtaNil(b bool)`

 SetAtaNil sets the value for Ata to be an explicit nil

### UnsetAta
`func (o *VoyagePortCallUpdateDto) UnsetAta()`

UnsetAta ensures that no value is present for Ata, not even an explicit nil
### GetEtd

`func (o *VoyagePortCallUpdateDto) GetEtd() time.Time`

GetEtd returns the Etd field if non-nil, zero value otherwise.

### GetEtdOk

`func (o *VoyagePortCallUpdateDto) GetEtdOk() (*time.Time, bool)`

GetEtdOk returns a tuple with the Etd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtd

`func (o *VoyagePortCallUpdateDto) SetEtd(v time.Time)`

SetEtd sets Etd field to given value.

### HasEtd

`func (o *VoyagePortCallUpdateDto) HasEtd() bool`

HasEtd returns a boolean if a field has been set.

### SetEtdNil

`func (o *VoyagePortCallUpdateDto) SetEtdNil(b bool)`

 SetEtdNil sets the value for Etd to be an explicit nil

### UnsetEtd
`func (o *VoyagePortCallUpdateDto) UnsetEtd()`

UnsetEtd ensures that no value is present for Etd, not even an explicit nil
### GetAtd

`func (o *VoyagePortCallUpdateDto) GetAtd() time.Time`

GetAtd returns the Atd field if non-nil, zero value otherwise.

### GetAtdOk

`func (o *VoyagePortCallUpdateDto) GetAtdOk() (*time.Time, bool)`

GetAtdOk returns a tuple with the Atd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtd

`func (o *VoyagePortCallUpdateDto) SetAtd(v time.Time)`

SetAtd sets Atd field to given value.

### HasAtd

`func (o *VoyagePortCallUpdateDto) HasAtd() bool`

HasAtd returns a boolean if a field has been set.

### SetAtdNil

`func (o *VoyagePortCallUpdateDto) SetAtdNil(b bool)`

 SetAtdNil sets the value for Atd to be an explicit nil

### UnsetAtd
`func (o *VoyagePortCallUpdateDto) UnsetAtd()`

UnsetAtd ensures that no value is present for Atd, not even an explicit nil
### GetBerthNumber

`func (o *VoyagePortCallUpdateDto) GetBerthNumber() string`

GetBerthNumber returns the BerthNumber field if non-nil, zero value otherwise.

### GetBerthNumberOk

`func (o *VoyagePortCallUpdateDto) GetBerthNumberOk() (*string, bool)`

GetBerthNumberOk returns a tuple with the BerthNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBerthNumber

`func (o *VoyagePortCallUpdateDto) SetBerthNumber(v string)`

SetBerthNumber sets BerthNumber field to given value.

### HasBerthNumber

`func (o *VoyagePortCallUpdateDto) HasBerthNumber() bool`

HasBerthNumber returns a boolean if a field has been set.

### SetBerthNumberNil

`func (o *VoyagePortCallUpdateDto) SetBerthNumberNil(b bool)`

 SetBerthNumberNil sets the value for BerthNumber to be an explicit nil

### UnsetBerthNumber
`func (o *VoyagePortCallUpdateDto) UnsetBerthNumber()`

UnsetBerthNumber ensures that no value is present for BerthNumber, not even an explicit nil
### GetRemarks

`func (o *VoyagePortCallUpdateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *VoyagePortCallUpdateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *VoyagePortCallUpdateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *VoyagePortCallUpdateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *VoyagePortCallUpdateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *VoyagePortCallUpdateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetPortId

`func (o *VoyagePortCallUpdateDto) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *VoyagePortCallUpdateDto) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *VoyagePortCallUpdateDto) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *VoyagePortCallUpdateDto) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### SetPortIdNil

`func (o *VoyagePortCallUpdateDto) SetPortIdNil(b bool)`

 SetPortIdNil sets the value for PortId to be an explicit nil

### UnsetPortId
`func (o *VoyagePortCallUpdateDto) UnsetPortId()`

UnsetPortId ensures that no value is present for PortId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


