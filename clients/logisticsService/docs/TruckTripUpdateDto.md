# TruckTripUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TripNumber** | Pointer to **NullableString** |  | [optional] 
**ContainerNumber** | Pointer to **NullableString** |  | [optional] 
**SealNumber** | Pointer to **NullableString** |  | [optional] 
**DepartureTime** | Pointer to **NullableTime** |  | [optional] 
**ArrivalTime** | Pointer to **NullableTime** |  | [optional] 
**DistanceKm** | Pointer to **NullableFloat64** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**OriginPortId** | Pointer to **NullableString** |  | [optional] 
**OriginLocationId** | Pointer to **NullableString** |  | [optional] 
**DestinationPortId** | Pointer to **NullableString** |  | [optional] 
**DestinationLocationId** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**BillOfLadingId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTruckTripUpdateDto

`func NewTruckTripUpdateDto() *TruckTripUpdateDto`

NewTruckTripUpdateDto instantiates a new TruckTripUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruckTripUpdateDtoWithDefaults

`func NewTruckTripUpdateDtoWithDefaults() *TruckTripUpdateDto`

NewTruckTripUpdateDtoWithDefaults instantiates a new TruckTripUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTripNumber

`func (o *TruckTripUpdateDto) GetTripNumber() string`

GetTripNumber returns the TripNumber field if non-nil, zero value otherwise.

### GetTripNumberOk

`func (o *TruckTripUpdateDto) GetTripNumberOk() (*string, bool)`

GetTripNumberOk returns a tuple with the TripNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripNumber

`func (o *TruckTripUpdateDto) SetTripNumber(v string)`

SetTripNumber sets TripNumber field to given value.

### HasTripNumber

`func (o *TruckTripUpdateDto) HasTripNumber() bool`

HasTripNumber returns a boolean if a field has been set.

### SetTripNumberNil

`func (o *TruckTripUpdateDto) SetTripNumberNil(b bool)`

 SetTripNumberNil sets the value for TripNumber to be an explicit nil

### UnsetTripNumber
`func (o *TruckTripUpdateDto) UnsetTripNumber()`

UnsetTripNumber ensures that no value is present for TripNumber, not even an explicit nil
### GetContainerNumber

`func (o *TruckTripUpdateDto) GetContainerNumber() string`

GetContainerNumber returns the ContainerNumber field if non-nil, zero value otherwise.

### GetContainerNumberOk

`func (o *TruckTripUpdateDto) GetContainerNumberOk() (*string, bool)`

GetContainerNumberOk returns a tuple with the ContainerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNumber

`func (o *TruckTripUpdateDto) SetContainerNumber(v string)`

SetContainerNumber sets ContainerNumber field to given value.

### HasContainerNumber

`func (o *TruckTripUpdateDto) HasContainerNumber() bool`

HasContainerNumber returns a boolean if a field has been set.

### SetContainerNumberNil

`func (o *TruckTripUpdateDto) SetContainerNumberNil(b bool)`

 SetContainerNumberNil sets the value for ContainerNumber to be an explicit nil

### UnsetContainerNumber
`func (o *TruckTripUpdateDto) UnsetContainerNumber()`

UnsetContainerNumber ensures that no value is present for ContainerNumber, not even an explicit nil
### GetSealNumber

`func (o *TruckTripUpdateDto) GetSealNumber() string`

GetSealNumber returns the SealNumber field if non-nil, zero value otherwise.

### GetSealNumberOk

`func (o *TruckTripUpdateDto) GetSealNumberOk() (*string, bool)`

GetSealNumberOk returns a tuple with the SealNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealNumber

`func (o *TruckTripUpdateDto) SetSealNumber(v string)`

SetSealNumber sets SealNumber field to given value.

### HasSealNumber

`func (o *TruckTripUpdateDto) HasSealNumber() bool`

HasSealNumber returns a boolean if a field has been set.

### SetSealNumberNil

`func (o *TruckTripUpdateDto) SetSealNumberNil(b bool)`

 SetSealNumberNil sets the value for SealNumber to be an explicit nil

### UnsetSealNumber
`func (o *TruckTripUpdateDto) UnsetSealNumber()`

UnsetSealNumber ensures that no value is present for SealNumber, not even an explicit nil
### GetDepartureTime

`func (o *TruckTripUpdateDto) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *TruckTripUpdateDto) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *TruckTripUpdateDto) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *TruckTripUpdateDto) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### SetDepartureTimeNil

`func (o *TruckTripUpdateDto) SetDepartureTimeNil(b bool)`

 SetDepartureTimeNil sets the value for DepartureTime to be an explicit nil

### UnsetDepartureTime
`func (o *TruckTripUpdateDto) UnsetDepartureTime()`

UnsetDepartureTime ensures that no value is present for DepartureTime, not even an explicit nil
### GetArrivalTime

`func (o *TruckTripUpdateDto) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *TruckTripUpdateDto) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *TruckTripUpdateDto) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *TruckTripUpdateDto) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### SetArrivalTimeNil

`func (o *TruckTripUpdateDto) SetArrivalTimeNil(b bool)`

 SetArrivalTimeNil sets the value for ArrivalTime to be an explicit nil

### UnsetArrivalTime
`func (o *TruckTripUpdateDto) UnsetArrivalTime()`

UnsetArrivalTime ensures that no value is present for ArrivalTime, not even an explicit nil
### GetDistanceKm

`func (o *TruckTripUpdateDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *TruckTripUpdateDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *TruckTripUpdateDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *TruckTripUpdateDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### SetDistanceKmNil

`func (o *TruckTripUpdateDto) SetDistanceKmNil(b bool)`

 SetDistanceKmNil sets the value for DistanceKm to be an explicit nil

### UnsetDistanceKm
`func (o *TruckTripUpdateDto) UnsetDistanceKm()`

UnsetDistanceKm ensures that no value is present for DistanceKm, not even an explicit nil
### GetNotes

`func (o *TruckTripUpdateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TruckTripUpdateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TruckTripUpdateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TruckTripUpdateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TruckTripUpdateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TruckTripUpdateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOriginPortId

`func (o *TruckTripUpdateDto) GetOriginPortId() string`

GetOriginPortId returns the OriginPortId field if non-nil, zero value otherwise.

### GetOriginPortIdOk

`func (o *TruckTripUpdateDto) GetOriginPortIdOk() (*string, bool)`

GetOriginPortIdOk returns a tuple with the OriginPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPortId

`func (o *TruckTripUpdateDto) SetOriginPortId(v string)`

SetOriginPortId sets OriginPortId field to given value.

### HasOriginPortId

`func (o *TruckTripUpdateDto) HasOriginPortId() bool`

HasOriginPortId returns a boolean if a field has been set.

### SetOriginPortIdNil

`func (o *TruckTripUpdateDto) SetOriginPortIdNil(b bool)`

 SetOriginPortIdNil sets the value for OriginPortId to be an explicit nil

### UnsetOriginPortId
`func (o *TruckTripUpdateDto) UnsetOriginPortId()`

UnsetOriginPortId ensures that no value is present for OriginPortId, not even an explicit nil
### GetOriginLocationId

`func (o *TruckTripUpdateDto) GetOriginLocationId() string`

GetOriginLocationId returns the OriginLocationId field if non-nil, zero value otherwise.

### GetOriginLocationIdOk

`func (o *TruckTripUpdateDto) GetOriginLocationIdOk() (*string, bool)`

GetOriginLocationIdOk returns a tuple with the OriginLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginLocationId

`func (o *TruckTripUpdateDto) SetOriginLocationId(v string)`

SetOriginLocationId sets OriginLocationId field to given value.

### HasOriginLocationId

`func (o *TruckTripUpdateDto) HasOriginLocationId() bool`

HasOriginLocationId returns a boolean if a field has been set.

### SetOriginLocationIdNil

`func (o *TruckTripUpdateDto) SetOriginLocationIdNil(b bool)`

 SetOriginLocationIdNil sets the value for OriginLocationId to be an explicit nil

### UnsetOriginLocationId
`func (o *TruckTripUpdateDto) UnsetOriginLocationId()`

UnsetOriginLocationId ensures that no value is present for OriginLocationId, not even an explicit nil
### GetDestinationPortId

`func (o *TruckTripUpdateDto) GetDestinationPortId() string`

GetDestinationPortId returns the DestinationPortId field if non-nil, zero value otherwise.

### GetDestinationPortIdOk

`func (o *TruckTripUpdateDto) GetDestinationPortIdOk() (*string, bool)`

GetDestinationPortIdOk returns a tuple with the DestinationPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortId

`func (o *TruckTripUpdateDto) SetDestinationPortId(v string)`

SetDestinationPortId sets DestinationPortId field to given value.

### HasDestinationPortId

`func (o *TruckTripUpdateDto) HasDestinationPortId() bool`

HasDestinationPortId returns a boolean if a field has been set.

### SetDestinationPortIdNil

`func (o *TruckTripUpdateDto) SetDestinationPortIdNil(b bool)`

 SetDestinationPortIdNil sets the value for DestinationPortId to be an explicit nil

### UnsetDestinationPortId
`func (o *TruckTripUpdateDto) UnsetDestinationPortId()`

UnsetDestinationPortId ensures that no value is present for DestinationPortId, not even an explicit nil
### GetDestinationLocationId

`func (o *TruckTripUpdateDto) GetDestinationLocationId() string`

GetDestinationLocationId returns the DestinationLocationId field if non-nil, zero value otherwise.

### GetDestinationLocationIdOk

`func (o *TruckTripUpdateDto) GetDestinationLocationIdOk() (*string, bool)`

GetDestinationLocationIdOk returns a tuple with the DestinationLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationId

`func (o *TruckTripUpdateDto) SetDestinationLocationId(v string)`

SetDestinationLocationId sets DestinationLocationId field to given value.

### HasDestinationLocationId

`func (o *TruckTripUpdateDto) HasDestinationLocationId() bool`

HasDestinationLocationId returns a boolean if a field has been set.

### SetDestinationLocationIdNil

`func (o *TruckTripUpdateDto) SetDestinationLocationIdNil(b bool)`

 SetDestinationLocationIdNil sets the value for DestinationLocationId to be an explicit nil

### UnsetDestinationLocationId
`func (o *TruckTripUpdateDto) UnsetDestinationLocationId()`

UnsetDestinationLocationId ensures that no value is present for DestinationLocationId, not even an explicit nil
### GetShipmentId

`func (o *TruckTripUpdateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *TruckTripUpdateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *TruckTripUpdateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *TruckTripUpdateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *TruckTripUpdateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *TruckTripUpdateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetBillOfLadingId

`func (o *TruckTripUpdateDto) GetBillOfLadingId() string`

GetBillOfLadingId returns the BillOfLadingId field if non-nil, zero value otherwise.

### GetBillOfLadingIdOk

`func (o *TruckTripUpdateDto) GetBillOfLadingIdOk() (*string, bool)`

GetBillOfLadingIdOk returns a tuple with the BillOfLadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingId

`func (o *TruckTripUpdateDto) SetBillOfLadingId(v string)`

SetBillOfLadingId sets BillOfLadingId field to given value.

### HasBillOfLadingId

`func (o *TruckTripUpdateDto) HasBillOfLadingId() bool`

HasBillOfLadingId returns a boolean if a field has been set.

### SetBillOfLadingIdNil

`func (o *TruckTripUpdateDto) SetBillOfLadingIdNil(b bool)`

 SetBillOfLadingIdNil sets the value for BillOfLadingId to be an explicit nil

### UnsetBillOfLadingId
`func (o *TruckTripUpdateDto) UnsetBillOfLadingId()`

UnsetBillOfLadingId ensures that no value is present for BillOfLadingId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


