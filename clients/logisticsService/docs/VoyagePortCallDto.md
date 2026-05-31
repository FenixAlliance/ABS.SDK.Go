# VoyagePortCallDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**SequenceNumber** | Pointer to **int32** |  | [optional] 
**PortCallStatus** | Pointer to **NullableString** |  | [optional] 
**Eta** | Pointer to **NullableTime** |  | [optional] 
**Ata** | Pointer to **NullableTime** |  | [optional] 
**Etd** | Pointer to **NullableTime** |  | [optional] 
**Atd** | Pointer to **NullableTime** |  | [optional] 
**BerthNumber** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**VoyageId** | Pointer to **NullableString** |  | [optional] 
**PortId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVoyagePortCallDto

`func NewVoyagePortCallDto() *VoyagePortCallDto`

NewVoyagePortCallDto instantiates a new VoyagePortCallDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoyagePortCallDtoWithDefaults

`func NewVoyagePortCallDtoWithDefaults() *VoyagePortCallDto`

NewVoyagePortCallDtoWithDefaults instantiates a new VoyagePortCallDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoyagePortCallDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoyagePortCallDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoyagePortCallDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoyagePortCallDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VoyagePortCallDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VoyagePortCallDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *VoyagePortCallDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VoyagePortCallDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VoyagePortCallDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VoyagePortCallDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *VoyagePortCallDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *VoyagePortCallDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetSequenceNumber

`func (o *VoyagePortCallDto) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *VoyagePortCallDto) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *VoyagePortCallDto) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *VoyagePortCallDto) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetPortCallStatus

`func (o *VoyagePortCallDto) GetPortCallStatus() string`

GetPortCallStatus returns the PortCallStatus field if non-nil, zero value otherwise.

### GetPortCallStatusOk

`func (o *VoyagePortCallDto) GetPortCallStatusOk() (*string, bool)`

GetPortCallStatusOk returns a tuple with the PortCallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCallStatus

`func (o *VoyagePortCallDto) SetPortCallStatus(v string)`

SetPortCallStatus sets PortCallStatus field to given value.

### HasPortCallStatus

`func (o *VoyagePortCallDto) HasPortCallStatus() bool`

HasPortCallStatus returns a boolean if a field has been set.

### SetPortCallStatusNil

`func (o *VoyagePortCallDto) SetPortCallStatusNil(b bool)`

 SetPortCallStatusNil sets the value for PortCallStatus to be an explicit nil

### UnsetPortCallStatus
`func (o *VoyagePortCallDto) UnsetPortCallStatus()`

UnsetPortCallStatus ensures that no value is present for PortCallStatus, not even an explicit nil
### GetEta

`func (o *VoyagePortCallDto) GetEta() time.Time`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *VoyagePortCallDto) GetEtaOk() (*time.Time, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *VoyagePortCallDto) SetEta(v time.Time)`

SetEta sets Eta field to given value.

### HasEta

`func (o *VoyagePortCallDto) HasEta() bool`

HasEta returns a boolean if a field has been set.

### SetEtaNil

`func (o *VoyagePortCallDto) SetEtaNil(b bool)`

 SetEtaNil sets the value for Eta to be an explicit nil

### UnsetEta
`func (o *VoyagePortCallDto) UnsetEta()`

UnsetEta ensures that no value is present for Eta, not even an explicit nil
### GetAta

`func (o *VoyagePortCallDto) GetAta() time.Time`

GetAta returns the Ata field if non-nil, zero value otherwise.

### GetAtaOk

`func (o *VoyagePortCallDto) GetAtaOk() (*time.Time, bool)`

GetAtaOk returns a tuple with the Ata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAta

`func (o *VoyagePortCallDto) SetAta(v time.Time)`

SetAta sets Ata field to given value.

### HasAta

`func (o *VoyagePortCallDto) HasAta() bool`

HasAta returns a boolean if a field has been set.

### SetAtaNil

`func (o *VoyagePortCallDto) SetAtaNil(b bool)`

 SetAtaNil sets the value for Ata to be an explicit nil

### UnsetAta
`func (o *VoyagePortCallDto) UnsetAta()`

UnsetAta ensures that no value is present for Ata, not even an explicit nil
### GetEtd

`func (o *VoyagePortCallDto) GetEtd() time.Time`

GetEtd returns the Etd field if non-nil, zero value otherwise.

### GetEtdOk

`func (o *VoyagePortCallDto) GetEtdOk() (*time.Time, bool)`

GetEtdOk returns a tuple with the Etd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtd

`func (o *VoyagePortCallDto) SetEtd(v time.Time)`

SetEtd sets Etd field to given value.

### HasEtd

`func (o *VoyagePortCallDto) HasEtd() bool`

HasEtd returns a boolean if a field has been set.

### SetEtdNil

`func (o *VoyagePortCallDto) SetEtdNil(b bool)`

 SetEtdNil sets the value for Etd to be an explicit nil

### UnsetEtd
`func (o *VoyagePortCallDto) UnsetEtd()`

UnsetEtd ensures that no value is present for Etd, not even an explicit nil
### GetAtd

`func (o *VoyagePortCallDto) GetAtd() time.Time`

GetAtd returns the Atd field if non-nil, zero value otherwise.

### GetAtdOk

`func (o *VoyagePortCallDto) GetAtdOk() (*time.Time, bool)`

GetAtdOk returns a tuple with the Atd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtd

`func (o *VoyagePortCallDto) SetAtd(v time.Time)`

SetAtd sets Atd field to given value.

### HasAtd

`func (o *VoyagePortCallDto) HasAtd() bool`

HasAtd returns a boolean if a field has been set.

### SetAtdNil

`func (o *VoyagePortCallDto) SetAtdNil(b bool)`

 SetAtdNil sets the value for Atd to be an explicit nil

### UnsetAtd
`func (o *VoyagePortCallDto) UnsetAtd()`

UnsetAtd ensures that no value is present for Atd, not even an explicit nil
### GetBerthNumber

`func (o *VoyagePortCallDto) GetBerthNumber() string`

GetBerthNumber returns the BerthNumber field if non-nil, zero value otherwise.

### GetBerthNumberOk

`func (o *VoyagePortCallDto) GetBerthNumberOk() (*string, bool)`

GetBerthNumberOk returns a tuple with the BerthNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBerthNumber

`func (o *VoyagePortCallDto) SetBerthNumber(v string)`

SetBerthNumber sets BerthNumber field to given value.

### HasBerthNumber

`func (o *VoyagePortCallDto) HasBerthNumber() bool`

HasBerthNumber returns a boolean if a field has been set.

### SetBerthNumberNil

`func (o *VoyagePortCallDto) SetBerthNumberNil(b bool)`

 SetBerthNumberNil sets the value for BerthNumber to be an explicit nil

### UnsetBerthNumber
`func (o *VoyagePortCallDto) UnsetBerthNumber()`

UnsetBerthNumber ensures that no value is present for BerthNumber, not even an explicit nil
### GetRemarks

`func (o *VoyagePortCallDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *VoyagePortCallDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *VoyagePortCallDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *VoyagePortCallDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *VoyagePortCallDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *VoyagePortCallDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetVoyageId

`func (o *VoyagePortCallDto) GetVoyageId() string`

GetVoyageId returns the VoyageId field if non-nil, zero value otherwise.

### GetVoyageIdOk

`func (o *VoyagePortCallDto) GetVoyageIdOk() (*string, bool)`

GetVoyageIdOk returns a tuple with the VoyageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageId

`func (o *VoyagePortCallDto) SetVoyageId(v string)`

SetVoyageId sets VoyageId field to given value.

### HasVoyageId

`func (o *VoyagePortCallDto) HasVoyageId() bool`

HasVoyageId returns a boolean if a field has been set.

### SetVoyageIdNil

`func (o *VoyagePortCallDto) SetVoyageIdNil(b bool)`

 SetVoyageIdNil sets the value for VoyageId to be an explicit nil

### UnsetVoyageId
`func (o *VoyagePortCallDto) UnsetVoyageId()`

UnsetVoyageId ensures that no value is present for VoyageId, not even an explicit nil
### GetPortId

`func (o *VoyagePortCallDto) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *VoyagePortCallDto) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *VoyagePortCallDto) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *VoyagePortCallDto) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### SetPortIdNil

`func (o *VoyagePortCallDto) SetPortIdNil(b bool)`

 SetPortIdNil sets the value for PortId to be an explicit nil

### UnsetPortId
`func (o *VoyagePortCallDto) UnsetPortId()`

UnsetPortId ensures that no value is present for PortId, not even an explicit nil
### GetTenantId

`func (o *VoyagePortCallDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *VoyagePortCallDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *VoyagePortCallDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *VoyagePortCallDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *VoyagePortCallDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *VoyagePortCallDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


