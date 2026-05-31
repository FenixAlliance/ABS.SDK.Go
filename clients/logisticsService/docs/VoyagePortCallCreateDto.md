# VoyagePortCallCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**SequenceNumber** | Pointer to **int32** |  | [optional] 
**PortCallStatus** | Pointer to **NullableString** |  | [optional] 
**Eta** | Pointer to **NullableTime** |  | [optional] 
**Etd** | Pointer to **NullableTime** |  | [optional] 
**BerthNumber** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**PortId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVoyagePortCallCreateDto

`func NewVoyagePortCallCreateDto() *VoyagePortCallCreateDto`

NewVoyagePortCallCreateDto instantiates a new VoyagePortCallCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoyagePortCallCreateDtoWithDefaults

`func NewVoyagePortCallCreateDtoWithDefaults() *VoyagePortCallCreateDto`

NewVoyagePortCallCreateDtoWithDefaults instantiates a new VoyagePortCallCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoyagePortCallCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoyagePortCallCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoyagePortCallCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoyagePortCallCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *VoyagePortCallCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VoyagePortCallCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VoyagePortCallCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VoyagePortCallCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *VoyagePortCallCreateDto) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *VoyagePortCallCreateDto) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *VoyagePortCallCreateDto) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *VoyagePortCallCreateDto) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetPortCallStatus

`func (o *VoyagePortCallCreateDto) GetPortCallStatus() string`

GetPortCallStatus returns the PortCallStatus field if non-nil, zero value otherwise.

### GetPortCallStatusOk

`func (o *VoyagePortCallCreateDto) GetPortCallStatusOk() (*string, bool)`

GetPortCallStatusOk returns a tuple with the PortCallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortCallStatus

`func (o *VoyagePortCallCreateDto) SetPortCallStatus(v string)`

SetPortCallStatus sets PortCallStatus field to given value.

### HasPortCallStatus

`func (o *VoyagePortCallCreateDto) HasPortCallStatus() bool`

HasPortCallStatus returns a boolean if a field has been set.

### SetPortCallStatusNil

`func (o *VoyagePortCallCreateDto) SetPortCallStatusNil(b bool)`

 SetPortCallStatusNil sets the value for PortCallStatus to be an explicit nil

### UnsetPortCallStatus
`func (o *VoyagePortCallCreateDto) UnsetPortCallStatus()`

UnsetPortCallStatus ensures that no value is present for PortCallStatus, not even an explicit nil
### GetEta

`func (o *VoyagePortCallCreateDto) GetEta() time.Time`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *VoyagePortCallCreateDto) GetEtaOk() (*time.Time, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *VoyagePortCallCreateDto) SetEta(v time.Time)`

SetEta sets Eta field to given value.

### HasEta

`func (o *VoyagePortCallCreateDto) HasEta() bool`

HasEta returns a boolean if a field has been set.

### SetEtaNil

`func (o *VoyagePortCallCreateDto) SetEtaNil(b bool)`

 SetEtaNil sets the value for Eta to be an explicit nil

### UnsetEta
`func (o *VoyagePortCallCreateDto) UnsetEta()`

UnsetEta ensures that no value is present for Eta, not even an explicit nil
### GetEtd

`func (o *VoyagePortCallCreateDto) GetEtd() time.Time`

GetEtd returns the Etd field if non-nil, zero value otherwise.

### GetEtdOk

`func (o *VoyagePortCallCreateDto) GetEtdOk() (*time.Time, bool)`

GetEtdOk returns a tuple with the Etd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtd

`func (o *VoyagePortCallCreateDto) SetEtd(v time.Time)`

SetEtd sets Etd field to given value.

### HasEtd

`func (o *VoyagePortCallCreateDto) HasEtd() bool`

HasEtd returns a boolean if a field has been set.

### SetEtdNil

`func (o *VoyagePortCallCreateDto) SetEtdNil(b bool)`

 SetEtdNil sets the value for Etd to be an explicit nil

### UnsetEtd
`func (o *VoyagePortCallCreateDto) UnsetEtd()`

UnsetEtd ensures that no value is present for Etd, not even an explicit nil
### GetBerthNumber

`func (o *VoyagePortCallCreateDto) GetBerthNumber() string`

GetBerthNumber returns the BerthNumber field if non-nil, zero value otherwise.

### GetBerthNumberOk

`func (o *VoyagePortCallCreateDto) GetBerthNumberOk() (*string, bool)`

GetBerthNumberOk returns a tuple with the BerthNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBerthNumber

`func (o *VoyagePortCallCreateDto) SetBerthNumber(v string)`

SetBerthNumber sets BerthNumber field to given value.

### HasBerthNumber

`func (o *VoyagePortCallCreateDto) HasBerthNumber() bool`

HasBerthNumber returns a boolean if a field has been set.

### SetBerthNumberNil

`func (o *VoyagePortCallCreateDto) SetBerthNumberNil(b bool)`

 SetBerthNumberNil sets the value for BerthNumber to be an explicit nil

### UnsetBerthNumber
`func (o *VoyagePortCallCreateDto) UnsetBerthNumber()`

UnsetBerthNumber ensures that no value is present for BerthNumber, not even an explicit nil
### GetRemarks

`func (o *VoyagePortCallCreateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *VoyagePortCallCreateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *VoyagePortCallCreateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *VoyagePortCallCreateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *VoyagePortCallCreateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *VoyagePortCallCreateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetPortId

`func (o *VoyagePortCallCreateDto) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *VoyagePortCallCreateDto) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *VoyagePortCallCreateDto) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *VoyagePortCallCreateDto) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### SetPortIdNil

`func (o *VoyagePortCallCreateDto) SetPortIdNil(b bool)`

 SetPortIdNil sets the value for PortId to be an explicit nil

### UnsetPortId
`func (o *VoyagePortCallCreateDto) UnsetPortId()`

UnsetPortId ensures that no value is present for PortId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


