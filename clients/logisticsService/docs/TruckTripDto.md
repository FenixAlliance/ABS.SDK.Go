# TruckTripDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TripNumber** | Pointer to **NullableString** |  | [optional] 
**TruckTripStatus** | Pointer to **NullableString** |  | [optional] 
**ContainerNumber** | Pointer to **NullableString** |  | [optional] 
**SealNumber** | Pointer to **NullableString** |  | [optional] 
**DepartureTime** | Pointer to **NullableTime** |  | [optional] 
**ArrivalTime** | Pointer to **NullableTime** |  | [optional] 
**ActualDepartureTime** | Pointer to **NullableTime** |  | [optional] 
**ActualArrivalTime** | Pointer to **NullableTime** |  | [optional] 
**DistanceKm** | Pointer to **NullableFloat64** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**TruckId** | Pointer to **NullableString** |  | [optional] 
**OriginPortId** | Pointer to **NullableString** |  | [optional] 
**OriginLocationId** | Pointer to **NullableString** |  | [optional] 
**DestinationPortId** | Pointer to **NullableString** |  | [optional] 
**DestinationLocationId** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**BillOfLadingId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTruckTripDto

`func NewTruckTripDto() *TruckTripDto`

NewTruckTripDto instantiates a new TruckTripDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruckTripDtoWithDefaults

`func NewTruckTripDtoWithDefaults() *TruckTripDto`

NewTruckTripDtoWithDefaults instantiates a new TruckTripDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TruckTripDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TruckTripDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TruckTripDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TruckTripDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TruckTripDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TruckTripDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TruckTripDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TruckTripDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TruckTripDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TruckTripDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TruckTripDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TruckTripDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTripNumber

`func (o *TruckTripDto) GetTripNumber() string`

GetTripNumber returns the TripNumber field if non-nil, zero value otherwise.

### GetTripNumberOk

`func (o *TruckTripDto) GetTripNumberOk() (*string, bool)`

GetTripNumberOk returns a tuple with the TripNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripNumber

`func (o *TruckTripDto) SetTripNumber(v string)`

SetTripNumber sets TripNumber field to given value.

### HasTripNumber

`func (o *TruckTripDto) HasTripNumber() bool`

HasTripNumber returns a boolean if a field has been set.

### SetTripNumberNil

`func (o *TruckTripDto) SetTripNumberNil(b bool)`

 SetTripNumberNil sets the value for TripNumber to be an explicit nil

### UnsetTripNumber
`func (o *TruckTripDto) UnsetTripNumber()`

UnsetTripNumber ensures that no value is present for TripNumber, not even an explicit nil
### GetTruckTripStatus

`func (o *TruckTripDto) GetTruckTripStatus() string`

GetTruckTripStatus returns the TruckTripStatus field if non-nil, zero value otherwise.

### GetTruckTripStatusOk

`func (o *TruckTripDto) GetTruckTripStatusOk() (*string, bool)`

GetTruckTripStatusOk returns a tuple with the TruckTripStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckTripStatus

`func (o *TruckTripDto) SetTruckTripStatus(v string)`

SetTruckTripStatus sets TruckTripStatus field to given value.

### HasTruckTripStatus

`func (o *TruckTripDto) HasTruckTripStatus() bool`

HasTruckTripStatus returns a boolean if a field has been set.

### SetTruckTripStatusNil

`func (o *TruckTripDto) SetTruckTripStatusNil(b bool)`

 SetTruckTripStatusNil sets the value for TruckTripStatus to be an explicit nil

### UnsetTruckTripStatus
`func (o *TruckTripDto) UnsetTruckTripStatus()`

UnsetTruckTripStatus ensures that no value is present for TruckTripStatus, not even an explicit nil
### GetContainerNumber

`func (o *TruckTripDto) GetContainerNumber() string`

GetContainerNumber returns the ContainerNumber field if non-nil, zero value otherwise.

### GetContainerNumberOk

`func (o *TruckTripDto) GetContainerNumberOk() (*string, bool)`

GetContainerNumberOk returns a tuple with the ContainerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerNumber

`func (o *TruckTripDto) SetContainerNumber(v string)`

SetContainerNumber sets ContainerNumber field to given value.

### HasContainerNumber

`func (o *TruckTripDto) HasContainerNumber() bool`

HasContainerNumber returns a boolean if a field has been set.

### SetContainerNumberNil

`func (o *TruckTripDto) SetContainerNumberNil(b bool)`

 SetContainerNumberNil sets the value for ContainerNumber to be an explicit nil

### UnsetContainerNumber
`func (o *TruckTripDto) UnsetContainerNumber()`

UnsetContainerNumber ensures that no value is present for ContainerNumber, not even an explicit nil
### GetSealNumber

`func (o *TruckTripDto) GetSealNumber() string`

GetSealNumber returns the SealNumber field if non-nil, zero value otherwise.

### GetSealNumberOk

`func (o *TruckTripDto) GetSealNumberOk() (*string, bool)`

GetSealNumberOk returns a tuple with the SealNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealNumber

`func (o *TruckTripDto) SetSealNumber(v string)`

SetSealNumber sets SealNumber field to given value.

### HasSealNumber

`func (o *TruckTripDto) HasSealNumber() bool`

HasSealNumber returns a boolean if a field has been set.

### SetSealNumberNil

`func (o *TruckTripDto) SetSealNumberNil(b bool)`

 SetSealNumberNil sets the value for SealNumber to be an explicit nil

### UnsetSealNumber
`func (o *TruckTripDto) UnsetSealNumber()`

UnsetSealNumber ensures that no value is present for SealNumber, not even an explicit nil
### GetDepartureTime

`func (o *TruckTripDto) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *TruckTripDto) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *TruckTripDto) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.

### HasDepartureTime

`func (o *TruckTripDto) HasDepartureTime() bool`

HasDepartureTime returns a boolean if a field has been set.

### SetDepartureTimeNil

`func (o *TruckTripDto) SetDepartureTimeNil(b bool)`

 SetDepartureTimeNil sets the value for DepartureTime to be an explicit nil

### UnsetDepartureTime
`func (o *TruckTripDto) UnsetDepartureTime()`

UnsetDepartureTime ensures that no value is present for DepartureTime, not even an explicit nil
### GetArrivalTime

`func (o *TruckTripDto) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *TruckTripDto) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *TruckTripDto) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.

### HasArrivalTime

`func (o *TruckTripDto) HasArrivalTime() bool`

HasArrivalTime returns a boolean if a field has been set.

### SetArrivalTimeNil

`func (o *TruckTripDto) SetArrivalTimeNil(b bool)`

 SetArrivalTimeNil sets the value for ArrivalTime to be an explicit nil

### UnsetArrivalTime
`func (o *TruckTripDto) UnsetArrivalTime()`

UnsetArrivalTime ensures that no value is present for ArrivalTime, not even an explicit nil
### GetActualDepartureTime

`func (o *TruckTripDto) GetActualDepartureTime() time.Time`

GetActualDepartureTime returns the ActualDepartureTime field if non-nil, zero value otherwise.

### GetActualDepartureTimeOk

`func (o *TruckTripDto) GetActualDepartureTimeOk() (*time.Time, bool)`

GetActualDepartureTimeOk returns a tuple with the ActualDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDepartureTime

`func (o *TruckTripDto) SetActualDepartureTime(v time.Time)`

SetActualDepartureTime sets ActualDepartureTime field to given value.

### HasActualDepartureTime

`func (o *TruckTripDto) HasActualDepartureTime() bool`

HasActualDepartureTime returns a boolean if a field has been set.

### SetActualDepartureTimeNil

`func (o *TruckTripDto) SetActualDepartureTimeNil(b bool)`

 SetActualDepartureTimeNil sets the value for ActualDepartureTime to be an explicit nil

### UnsetActualDepartureTime
`func (o *TruckTripDto) UnsetActualDepartureTime()`

UnsetActualDepartureTime ensures that no value is present for ActualDepartureTime, not even an explicit nil
### GetActualArrivalTime

`func (o *TruckTripDto) GetActualArrivalTime() time.Time`

GetActualArrivalTime returns the ActualArrivalTime field if non-nil, zero value otherwise.

### GetActualArrivalTimeOk

`func (o *TruckTripDto) GetActualArrivalTimeOk() (*time.Time, bool)`

GetActualArrivalTimeOk returns a tuple with the ActualArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrivalTime

`func (o *TruckTripDto) SetActualArrivalTime(v time.Time)`

SetActualArrivalTime sets ActualArrivalTime field to given value.

### HasActualArrivalTime

`func (o *TruckTripDto) HasActualArrivalTime() bool`

HasActualArrivalTime returns a boolean if a field has been set.

### SetActualArrivalTimeNil

`func (o *TruckTripDto) SetActualArrivalTimeNil(b bool)`

 SetActualArrivalTimeNil sets the value for ActualArrivalTime to be an explicit nil

### UnsetActualArrivalTime
`func (o *TruckTripDto) UnsetActualArrivalTime()`

UnsetActualArrivalTime ensures that no value is present for ActualArrivalTime, not even an explicit nil
### GetDistanceKm

`func (o *TruckTripDto) GetDistanceKm() float64`

GetDistanceKm returns the DistanceKm field if non-nil, zero value otherwise.

### GetDistanceKmOk

`func (o *TruckTripDto) GetDistanceKmOk() (*float64, bool)`

GetDistanceKmOk returns a tuple with the DistanceKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceKm

`func (o *TruckTripDto) SetDistanceKm(v float64)`

SetDistanceKm sets DistanceKm field to given value.

### HasDistanceKm

`func (o *TruckTripDto) HasDistanceKm() bool`

HasDistanceKm returns a boolean if a field has been set.

### SetDistanceKmNil

`func (o *TruckTripDto) SetDistanceKmNil(b bool)`

 SetDistanceKmNil sets the value for DistanceKm to be an explicit nil

### UnsetDistanceKm
`func (o *TruckTripDto) UnsetDistanceKm()`

UnsetDistanceKm ensures that no value is present for DistanceKm, not even an explicit nil
### GetNotes

`func (o *TruckTripDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TruckTripDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TruckTripDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TruckTripDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TruckTripDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TruckTripDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetTruckId

`func (o *TruckTripDto) GetTruckId() string`

GetTruckId returns the TruckId field if non-nil, zero value otherwise.

### GetTruckIdOk

`func (o *TruckTripDto) GetTruckIdOk() (*string, bool)`

GetTruckIdOk returns a tuple with the TruckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckId

`func (o *TruckTripDto) SetTruckId(v string)`

SetTruckId sets TruckId field to given value.

### HasTruckId

`func (o *TruckTripDto) HasTruckId() bool`

HasTruckId returns a boolean if a field has been set.

### SetTruckIdNil

`func (o *TruckTripDto) SetTruckIdNil(b bool)`

 SetTruckIdNil sets the value for TruckId to be an explicit nil

### UnsetTruckId
`func (o *TruckTripDto) UnsetTruckId()`

UnsetTruckId ensures that no value is present for TruckId, not even an explicit nil
### GetOriginPortId

`func (o *TruckTripDto) GetOriginPortId() string`

GetOriginPortId returns the OriginPortId field if non-nil, zero value otherwise.

### GetOriginPortIdOk

`func (o *TruckTripDto) GetOriginPortIdOk() (*string, bool)`

GetOriginPortIdOk returns a tuple with the OriginPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPortId

`func (o *TruckTripDto) SetOriginPortId(v string)`

SetOriginPortId sets OriginPortId field to given value.

### HasOriginPortId

`func (o *TruckTripDto) HasOriginPortId() bool`

HasOriginPortId returns a boolean if a field has been set.

### SetOriginPortIdNil

`func (o *TruckTripDto) SetOriginPortIdNil(b bool)`

 SetOriginPortIdNil sets the value for OriginPortId to be an explicit nil

### UnsetOriginPortId
`func (o *TruckTripDto) UnsetOriginPortId()`

UnsetOriginPortId ensures that no value is present for OriginPortId, not even an explicit nil
### GetOriginLocationId

`func (o *TruckTripDto) GetOriginLocationId() string`

GetOriginLocationId returns the OriginLocationId field if non-nil, zero value otherwise.

### GetOriginLocationIdOk

`func (o *TruckTripDto) GetOriginLocationIdOk() (*string, bool)`

GetOriginLocationIdOk returns a tuple with the OriginLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginLocationId

`func (o *TruckTripDto) SetOriginLocationId(v string)`

SetOriginLocationId sets OriginLocationId field to given value.

### HasOriginLocationId

`func (o *TruckTripDto) HasOriginLocationId() bool`

HasOriginLocationId returns a boolean if a field has been set.

### SetOriginLocationIdNil

`func (o *TruckTripDto) SetOriginLocationIdNil(b bool)`

 SetOriginLocationIdNil sets the value for OriginLocationId to be an explicit nil

### UnsetOriginLocationId
`func (o *TruckTripDto) UnsetOriginLocationId()`

UnsetOriginLocationId ensures that no value is present for OriginLocationId, not even an explicit nil
### GetDestinationPortId

`func (o *TruckTripDto) GetDestinationPortId() string`

GetDestinationPortId returns the DestinationPortId field if non-nil, zero value otherwise.

### GetDestinationPortIdOk

`func (o *TruckTripDto) GetDestinationPortIdOk() (*string, bool)`

GetDestinationPortIdOk returns a tuple with the DestinationPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortId

`func (o *TruckTripDto) SetDestinationPortId(v string)`

SetDestinationPortId sets DestinationPortId field to given value.

### HasDestinationPortId

`func (o *TruckTripDto) HasDestinationPortId() bool`

HasDestinationPortId returns a boolean if a field has been set.

### SetDestinationPortIdNil

`func (o *TruckTripDto) SetDestinationPortIdNil(b bool)`

 SetDestinationPortIdNil sets the value for DestinationPortId to be an explicit nil

### UnsetDestinationPortId
`func (o *TruckTripDto) UnsetDestinationPortId()`

UnsetDestinationPortId ensures that no value is present for DestinationPortId, not even an explicit nil
### GetDestinationLocationId

`func (o *TruckTripDto) GetDestinationLocationId() string`

GetDestinationLocationId returns the DestinationLocationId field if non-nil, zero value otherwise.

### GetDestinationLocationIdOk

`func (o *TruckTripDto) GetDestinationLocationIdOk() (*string, bool)`

GetDestinationLocationIdOk returns a tuple with the DestinationLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLocationId

`func (o *TruckTripDto) SetDestinationLocationId(v string)`

SetDestinationLocationId sets DestinationLocationId field to given value.

### HasDestinationLocationId

`func (o *TruckTripDto) HasDestinationLocationId() bool`

HasDestinationLocationId returns a boolean if a field has been set.

### SetDestinationLocationIdNil

`func (o *TruckTripDto) SetDestinationLocationIdNil(b bool)`

 SetDestinationLocationIdNil sets the value for DestinationLocationId to be an explicit nil

### UnsetDestinationLocationId
`func (o *TruckTripDto) UnsetDestinationLocationId()`

UnsetDestinationLocationId ensures that no value is present for DestinationLocationId, not even an explicit nil
### GetShipmentId

`func (o *TruckTripDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *TruckTripDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *TruckTripDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *TruckTripDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *TruckTripDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *TruckTripDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetBillOfLadingId

`func (o *TruckTripDto) GetBillOfLadingId() string`

GetBillOfLadingId returns the BillOfLadingId field if non-nil, zero value otherwise.

### GetBillOfLadingIdOk

`func (o *TruckTripDto) GetBillOfLadingIdOk() (*string, bool)`

GetBillOfLadingIdOk returns a tuple with the BillOfLadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingId

`func (o *TruckTripDto) SetBillOfLadingId(v string)`

SetBillOfLadingId sets BillOfLadingId field to given value.

### HasBillOfLadingId

`func (o *TruckTripDto) HasBillOfLadingId() bool`

HasBillOfLadingId returns a boolean if a field has been set.

### SetBillOfLadingIdNil

`func (o *TruckTripDto) SetBillOfLadingIdNil(b bool)`

 SetBillOfLadingIdNil sets the value for BillOfLadingId to be an explicit nil

### UnsetBillOfLadingId
`func (o *TruckTripDto) UnsetBillOfLadingId()`

UnsetBillOfLadingId ensures that no value is present for BillOfLadingId, not even an explicit nil
### GetTenantId

`func (o *TruckTripDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TruckTripDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TruckTripDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TruckTripDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TruckTripDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TruckTripDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TruckTripDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TruckTripDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TruckTripDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TruckTripDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TruckTripDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TruckTripDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


