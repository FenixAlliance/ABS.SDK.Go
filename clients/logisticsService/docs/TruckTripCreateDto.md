# TruckTripCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewTruckTripCreateDto

`func NewTruckTripCreateDto() *TruckTripCreateDto`

NewTruckTripCreateDto instantiates a new TruckTripCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruckTripCreateDtoWithDefaults

`func NewTruckTripCreateDtoWithDefaults() *TruckTripCreateDto`

NewTruckTripCreateDtoWithDefaults instantiates a new TruckTripCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TruckTripCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TruckTripCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TruckTripCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TruckTripCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TruckTripCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TruckTripCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TruckTripCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TruckTripCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTripNumber

`func (o *TruckTripCreateDto) GetTripNumber() string`

GetTripNumber returns the TripNumber field if non-nil, zero value otherwise.

### GetTripNumberOk

`func (o *TruckTripCreateDto) GetTripNumberOk() (*string, bool)`

GetTripNumberOk returns a tuple with the TripNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripNumber

`func (o *TruckTripCreateDto) SetTripNumber(v string)`

SetTripNumber sets TripNumber field to given value.

### HasTripNumber

`func (o *TruckTripCreateDto) HasTripNumber() bool`

HasTripNumber returns a boolean if a field has been set.

### SetTripNumberNil

`func (o *TruckTripCreateDto) SetTripNumberNil(b bool)`

 SetTripNumberNil sets the value for TripNumber to be an explicit nil

### UnsetTripNumber
`func (o *TruckTripCreateDto) UnsetTripNumber()`

UnsetTripNumber ensures that no value is present for TripNumber, not even an explicit nil
### GetContainerNumber

`func (o *TruckTripCreateDto) GetContainerNumber() string`

GetContainerNumber returns the ContainerNumber field if non-nil, zero value otherwise.

### GetContainerNumberOk

`func (o *TruckTripCreateDto) GetContainerNumberOk() (*string, bool)`

GetContainerNumberOk returns a tuple with the ContainerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNumber

`func (o *TruckTripCreateDto) SetContainerNumber(v string)`

SetContainerNumber sets ContainerNumber field to given value.

### HasContainerNumber

`func (o *TruckTripCreateDto) HasContainerNumber() bool`

HasContainerNumber returns a boolean if a field has been set.

### SetContainerNumberNil

`func (o *TruckTripCreateDto) SetContainerNumberNil(b bool)`

 SetContainerNumberNil sets the value for ContainerNumber to be an explicit nil

### UnsetContainerNumber
`func (o *TruckTripCreateDto) UnsetContainerNumber()`

UnsetContainerNumber ensures that no value is present for ContainerNumber, not even an explicit nil
### GetSealNumber

`func (o *TruckTripCreateDto) GetSealNumber() string`

GetSealNumber returns the SealNumber field if non-nil, zero value otherwise.

### GetSealNumberOk

`func (o *TruckTripCreateDto) GetSealNumberOk() (*string, bool)`

GetSealNumberOk returns a tuple with the SealNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealNumber

`func (o *TruckTripCreateDto) SetSealNumber(v string)`

SetSealNumber sets SealNumber field to given value.

### HasSealNumber

`func (o *TruckTripCreateDto) HasSealNumber() bool`

HasSealNumber returns a boolean if a field has been set.

### SetSealNumberNil

`func (o *TruckTripCreateDto) SetSealNumberNil(b bool)`

 SetSealNumberNil sets the value for SealNumber to be an explicit nil

### UnsetSealNumber
`func (o *TruckTripCreateDto) UnsetSealNumber()`

UnsetSealNumber ensures that no value is present for SealNumber, not even an explicit nil
### GetDepartureTime

`func (o *TruckTripCreateDto) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *TruckTripCreateDto) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *TruckTripCreateDto) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *TruckTripCreateDto) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### SetDepartureTimeNil

`func (o *TruckTripCreateDto) SetDepartureTimeNil(b bool)`

 SetDepartureTimeNil sets the value for DepartureTime to be an explicit nil

### UnsetDepartureTime
`func (o *TruckTripCreateDto) UnsetDepartureTime()`

UnsetDepartureTime ensures that no value is present for DepartureTime, not even an explicit nil
### GetArrivalTime

`func (o *TruckTripCreateDto) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *TruckTripCreateDto) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *TruckTripCreateDto) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *TruckTripCreateDto) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### SetArrivalTimeNil

`func (o *TruckTripCreateDto) SetArrivalTimeNil(b bool)`

 SetArrivalTimeNil sets the value for ArrivalTime to be an explicit nil

### UnsetArrivalTime
`func (o *TruckTripCreateDto) UnsetArrivalTime()`

UnsetArrivalTime ensures that no value is present for ArrivalTime, not even an explicit nil
### GetDistanceKm

`func (o *TruckTripCreateDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *TruckTripCreateDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *TruckTripCreateDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *TruckTripCreateDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### SetDistanceKmNil

`func (o *TruckTripCreateDto) SetDistanceKmNil(b bool)`

 SetDistanceKmNil sets the value for DistanceKm to be an explicit nil

### UnsetDistanceKm
`func (o *TruckTripCreateDto) UnsetDistanceKm()`

UnsetDistanceKm ensures that no value is present for DistanceKm, not even an explicit nil
### GetNotes

`func (o *TruckTripCreateDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TruckTripCreateDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TruckTripCreateDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TruckTripCreateDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TruckTripCreateDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TruckTripCreateDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetOriginPortId

`func (o *TruckTripCreateDto) GetOriginPortId() string`

GetOriginPortId returns the OriginPortId field if non-nil, zero value otherwise.

### GetOriginPortIdOk

`func (o *TruckTripCreateDto) GetOriginPortIdOk() (*string, bool)`

GetOriginPortIdOk returns a tuple with the OriginPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPortId

`func (o *TruckTripCreateDto) SetOriginPortId(v string)`

SetOriginPortId sets OriginPortId field to given value.

### HasOriginPortId

`func (o *TruckTripCreateDto) HasOriginPortId() bool`

HasOriginPortId returns a boolean if a field has been set.

### SetOriginPortIdNil

`func (o *TruckTripCreateDto) SetOriginPortIdNil(b bool)`

 SetOriginPortIdNil sets the value for OriginPortId to be an explicit nil

### UnsetOriginPortId
`func (o *TruckTripCreateDto) UnsetOriginPortId()`

UnsetOriginPortId ensures that no value is present for OriginPortId, not even an explicit nil
### GetOriginLocationId

`func (o *TruckTripCreateDto) GetOriginLocationId() string`

GetOriginLocationId returns the OriginLocationId field if non-nil, zero value otherwise.

### GetOriginLocationIdOk

`func (o *TruckTripCreateDto) GetOriginLocationIdOk() (*string, bool)`

GetOriginLocationIdOk returns a tuple with the OriginLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginLocationId

`func (o *TruckTripCreateDto) SetOriginLocationId(v string)`

SetOriginLocationId sets OriginLocationId field to given value.

### HasOriginLocationId

`func (o *TruckTripCreateDto) HasOriginLocationId() bool`

HasOriginLocationId returns a boolean if a field has been set.

### SetOriginLocationIdNil

`func (o *TruckTripCreateDto) SetOriginLocationIdNil(b bool)`

 SetOriginLocationIdNil sets the value for OriginLocationId to be an explicit nil

### UnsetOriginLocationId
`func (o *TruckTripCreateDto) UnsetOriginLocationId()`

UnsetOriginLocationId ensures that no value is present for OriginLocationId, not even an explicit nil
### GetDestinationPortId

`func (o *TruckTripCreateDto) GetDestinationPortId() string`

GetDestinationPortId returns the DestinationPortId field if non-nil, zero value otherwise.

### GetDestinationPortIdOk

`func (o *TruckTripCreateDto) GetDestinationPortIdOk() (*string, bool)`

GetDestinationPortIdOk returns a tuple with the DestinationPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortId

`func (o *TruckTripCreateDto) SetDestinationPortId(v string)`

SetDestinationPortId sets DestinationPortId field to given value.

### HasDestinationPortId

`func (o *TruckTripCreateDto) HasDestinationPortId() bool`

HasDestinationPortId returns a boolean if a field has been set.

### SetDestinationPortIdNil

`func (o *TruckTripCreateDto) SetDestinationPortIdNil(b bool)`

 SetDestinationPortIdNil sets the value for DestinationPortId to be an explicit nil

### UnsetDestinationPortId
`func (o *TruckTripCreateDto) UnsetDestinationPortId()`

UnsetDestinationPortId ensures that no value is present for DestinationPortId, not even an explicit nil
### GetDestinationLocationId

`func (o *TruckTripCreateDto) GetDestinationLocationId() string`

GetDestinationLocationId returns the DestinationLocationId field if non-nil, zero value otherwise.

### GetDestinationLocationIdOk

`func (o *TruckTripCreateDto) GetDestinationLocationIdOk() (*string, bool)`

GetDestinationLocationIdOk returns a tuple with the DestinationLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationId

`func (o *TruckTripCreateDto) SetDestinationLocationId(v string)`

SetDestinationLocationId sets DestinationLocationId field to given value.

### HasDestinationLocationId

`func (o *TruckTripCreateDto) HasDestinationLocationId() bool`

HasDestinationLocationId returns a boolean if a field has been set.

### SetDestinationLocationIdNil

`func (o *TruckTripCreateDto) SetDestinationLocationIdNil(b bool)`

 SetDestinationLocationIdNil sets the value for DestinationLocationId to be an explicit nil

### UnsetDestinationLocationId
`func (o *TruckTripCreateDto) UnsetDestinationLocationId()`

UnsetDestinationLocationId ensures that no value is present for DestinationLocationId, not even an explicit nil
### GetShipmentId

`func (o *TruckTripCreateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *TruckTripCreateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *TruckTripCreateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *TruckTripCreateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *TruckTripCreateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *TruckTripCreateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetBillOfLadingId

`func (o *TruckTripCreateDto) GetBillOfLadingId() string`

GetBillOfLadingId returns the BillOfLadingId field if non-nil, zero value otherwise.

### GetBillOfLadingIdOk

`func (o *TruckTripCreateDto) GetBillOfLadingIdOk() (*string, bool)`

GetBillOfLadingIdOk returns a tuple with the BillOfLadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingId

`func (o *TruckTripCreateDto) SetBillOfLadingId(v string)`

SetBillOfLadingId sets BillOfLadingId field to given value.

### HasBillOfLadingId

`func (o *TruckTripCreateDto) HasBillOfLadingId() bool`

HasBillOfLadingId returns a boolean if a field has been set.

### SetBillOfLadingIdNil

`func (o *TruckTripCreateDto) SetBillOfLadingIdNil(b bool)`

 SetBillOfLadingIdNil sets the value for BillOfLadingId to be an explicit nil

### UnsetBillOfLadingId
`func (o *TruckTripCreateDto) UnsetBillOfLadingId()`

UnsetBillOfLadingId ensures that no value is present for BillOfLadingId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


