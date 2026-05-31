# RoadWaybillDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**RoadWaybillType** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ShipperContactId** | Pointer to **NullableString** |  | [optional] 
**ConsigneeContactId** | Pointer to **NullableString** |  | [optional] 
**CarrierId** | Pointer to **NullableString** |  | [optional] 
**SuccessiveCarriers** | Pointer to **NullableString** |  | [optional] 
**TruckId** | Pointer to **NullableString** |  | [optional] 
**TruckDriverId** | Pointer to **NullableString** |  | [optional] 
**VehicleRegistration** | Pointer to **NullableString** |  | [optional] 
**TrailerRegistration** | Pointer to **NullableString** |  | [optional] 
**PlaceOfTakingOver** | Pointer to **NullableString** |  | [optional] 
**PlaceOfTakingOverPortId** | Pointer to **NullableString** |  | [optional] 
**PlaceOfDelivery** | Pointer to **NullableString** |  | [optional] 
**PlaceOfDeliveryPortId** | Pointer to **NullableString** |  | [optional] 
**DateOfTakingOver** | Pointer to **NullableTime** |  | [optional] 
**DateOfDelivery** | Pointer to **NullableTime** |  | [optional] 
**FreightTerms** | Pointer to **NullableString** |  | [optional] 
**FreightAmount** | Pointer to **NullableFloat64** |  | [optional] 
**FreightCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGrossWeightKg** | Pointer to **NullableFloat64** |  | [optional] 
**TotalPackages** | Pointer to **NullableInt32** |  | [optional] 
**TotalVolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**AdrDangerousGoods** | Pointer to **bool** |  | [optional] 
**SpecialInstructions** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**SenderSignedDate** | Pointer to **NullableTime** |  | [optional] 
**CarrierSignedDate** | Pointer to **NullableTime** |  | [optional] 
**ConsigneeSignedDate** | Pointer to **NullableTime** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**TruckTripId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Lines** | Pointer to [**[]WaybillLineDto**](WaybillLineDto.md) |  | [optional] 

## Methods

### NewRoadWaybillDto

`func NewRoadWaybillDto() *RoadWaybillDto`

NewRoadWaybillDto instantiates a new RoadWaybillDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoadWaybillDtoWithDefaults

`func NewRoadWaybillDtoWithDefaults() *RoadWaybillDto`

NewRoadWaybillDtoWithDefaults instantiates a new RoadWaybillDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoadWaybillDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoadWaybillDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoadWaybillDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoadWaybillDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RoadWaybillDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RoadWaybillDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *RoadWaybillDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RoadWaybillDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RoadWaybillDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RoadWaybillDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *RoadWaybillDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *RoadWaybillDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDocumentNumber

`func (o *RoadWaybillDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *RoadWaybillDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *RoadWaybillDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *RoadWaybillDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *RoadWaybillDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *RoadWaybillDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetRoadWaybillType

`func (o *RoadWaybillDto) GetRoadWaybillType() string`

GetRoadWaybillType returns the RoadWaybillType field if non-nil, zero value otherwise.

### GetRoadWaybillTypeOk

`func (o *RoadWaybillDto) GetRoadWaybillTypeOk() (*string, bool)`

GetRoadWaybillTypeOk returns a tuple with the RoadWaybillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadWaybillType

`func (o *RoadWaybillDto) SetRoadWaybillType(v string)`

SetRoadWaybillType sets RoadWaybillType field to given value.

### HasRoadWaybillType

`func (o *RoadWaybillDto) HasRoadWaybillType() bool`

HasRoadWaybillType returns a boolean if a field has been set.

### SetRoadWaybillTypeNil

`func (o *RoadWaybillDto) SetRoadWaybillTypeNil(b bool)`

 SetRoadWaybillTypeNil sets the value for RoadWaybillType to be an explicit nil

### UnsetRoadWaybillType
`func (o *RoadWaybillDto) UnsetRoadWaybillType()`

UnsetRoadWaybillType ensures that no value is present for RoadWaybillType, not even an explicit nil
### GetStatus

`func (o *RoadWaybillDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoadWaybillDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoadWaybillDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoadWaybillDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RoadWaybillDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RoadWaybillDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetShipperContactId

`func (o *RoadWaybillDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *RoadWaybillDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *RoadWaybillDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *RoadWaybillDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *RoadWaybillDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *RoadWaybillDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *RoadWaybillDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *RoadWaybillDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *RoadWaybillDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *RoadWaybillDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *RoadWaybillDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *RoadWaybillDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetCarrierId

`func (o *RoadWaybillDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *RoadWaybillDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *RoadWaybillDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *RoadWaybillDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *RoadWaybillDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *RoadWaybillDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetSuccessiveCarriers

`func (o *RoadWaybillDto) GetSuccessiveCarriers() string`

GetSuccessiveCarriers returns the SuccessiveCarriers field if non-nil, zero value otherwise.

### GetSuccessiveCarriersOk

`func (o *RoadWaybillDto) GetSuccessiveCarriersOk() (*string, bool)`

GetSuccessiveCarriersOk returns a tuple with the SuccessiveCarriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessiveCarriers

`func (o *RoadWaybillDto) SetSuccessiveCarriers(v string)`

SetSuccessiveCarriers sets SuccessiveCarriers field to given value.

### HasSuccessiveCarriers

`func (o *RoadWaybillDto) HasSuccessiveCarriers() bool`

HasSuccessiveCarriers returns a boolean if a field has been set.

### SetSuccessiveCarriersNil

`func (o *RoadWaybillDto) SetSuccessiveCarriersNil(b bool)`

 SetSuccessiveCarriersNil sets the value for SuccessiveCarriers to be an explicit nil

### UnsetSuccessiveCarriers
`func (o *RoadWaybillDto) UnsetSuccessiveCarriers()`

UnsetSuccessiveCarriers ensures that no value is present for SuccessiveCarriers, not even an explicit nil
### GetTruckId

`func (o *RoadWaybillDto) GetTruckId() string`

GetTruckId returns the TruckId field if non-nil, zero value otherwise.

### GetTruckIdOk

`func (o *RoadWaybillDto) GetTruckIdOk() (*string, bool)`

GetTruckIdOk returns a tuple with the TruckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckId

`func (o *RoadWaybillDto) SetTruckId(v string)`

SetTruckId sets TruckId field to given value.

### HasTruckId

`func (o *RoadWaybillDto) HasTruckId() bool`

HasTruckId returns a boolean if a field has been set.

### SetTruckIdNil

`func (o *RoadWaybillDto) SetTruckIdNil(b bool)`

 SetTruckIdNil sets the value for TruckId to be an explicit nil

### UnsetTruckId
`func (o *RoadWaybillDto) UnsetTruckId()`

UnsetTruckId ensures that no value is present for TruckId, not even an explicit nil
### GetTruckDriverId

`func (o *RoadWaybillDto) GetTruckDriverId() string`

GetTruckDriverId returns the TruckDriverId field if non-nil, zero value otherwise.

### GetTruckDriverIdOk

`func (o *RoadWaybillDto) GetTruckDriverIdOk() (*string, bool)`

GetTruckDriverIdOk returns a tuple with the TruckDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckDriverId

`func (o *RoadWaybillDto) SetTruckDriverId(v string)`

SetTruckDriverId sets TruckDriverId field to given value.

### HasTruckDriverId

`func (o *RoadWaybillDto) HasTruckDriverId() bool`

HasTruckDriverId returns a boolean if a field has been set.

### SetTruckDriverIdNil

`func (o *RoadWaybillDto) SetTruckDriverIdNil(b bool)`

 SetTruckDriverIdNil sets the value for TruckDriverId to be an explicit nil

### UnsetTruckDriverId
`func (o *RoadWaybillDto) UnsetTruckDriverId()`

UnsetTruckDriverId ensures that no value is present for TruckDriverId, not even an explicit nil
### GetVehicleRegistration

`func (o *RoadWaybillDto) GetVehicleRegistration() string`

GetVehicleRegistration returns the VehicleRegistration field if non-nil, zero value otherwise.

### GetVehicleRegistrationOk

`func (o *RoadWaybillDto) GetVehicleRegistrationOk() (*string, bool)`

GetVehicleRegistrationOk returns a tuple with the VehicleRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleRegistration

`func (o *RoadWaybillDto) SetVehicleRegistration(v string)`

SetVehicleRegistration sets VehicleRegistration field to given value.

### HasVehicleRegistration

`func (o *RoadWaybillDto) HasVehicleRegistration() bool`

HasVehicleRegistration returns a boolean if a field has been set.

### SetVehicleRegistrationNil

`func (o *RoadWaybillDto) SetVehicleRegistrationNil(b bool)`

 SetVehicleRegistrationNil sets the value for VehicleRegistration to be an explicit nil

### UnsetVehicleRegistration
`func (o *RoadWaybillDto) UnsetVehicleRegistration()`

UnsetVehicleRegistration ensures that no value is present for VehicleRegistration, not even an explicit nil
### GetTrailerRegistration

`func (o *RoadWaybillDto) GetTrailerRegistration() string`

GetTrailerRegistration returns the TrailerRegistration field if non-nil, zero value otherwise.

### GetTrailerRegistrationOk

`func (o *RoadWaybillDto) GetTrailerRegistrationOk() (*string, bool)`

GetTrailerRegistrationOk returns a tuple with the TrailerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerRegistration

`func (o *RoadWaybillDto) SetTrailerRegistration(v string)`

SetTrailerRegistration sets TrailerRegistration field to given value.

### HasTrailerRegistration

`func (o *RoadWaybillDto) HasTrailerRegistration() bool`

HasTrailerRegistration returns a boolean if a field has been set.

### SetTrailerRegistrationNil

`func (o *RoadWaybillDto) SetTrailerRegistrationNil(b bool)`

 SetTrailerRegistrationNil sets the value for TrailerRegistration to be an explicit nil

### UnsetTrailerRegistration
`func (o *RoadWaybillDto) UnsetTrailerRegistration()`

UnsetTrailerRegistration ensures that no value is present for TrailerRegistration, not even an explicit nil
### GetPlaceOfTakingOver

`func (o *RoadWaybillDto) GetPlaceOfTakingOver() string`

GetPlaceOfTakingOver returns the PlaceOfTakingOver field if non-nil, zero value otherwise.

### GetPlaceOfTakingOverOk

`func (o *RoadWaybillDto) GetPlaceOfTakingOverOk() (*string, bool)`

GetPlaceOfTakingOverOk returns a tuple with the PlaceOfTakingOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfTakingOver

`func (o *RoadWaybillDto) SetPlaceOfTakingOver(v string)`

SetPlaceOfTakingOver sets PlaceOfTakingOver field to given value.

### HasPlaceOfTakingOver

`func (o *RoadWaybillDto) HasPlaceOfTakingOver() bool`

HasPlaceOfTakingOver returns a boolean if a field has been set.

### SetPlaceOfTakingOverNil

`func (o *RoadWaybillDto) SetPlaceOfTakingOverNil(b bool)`

 SetPlaceOfTakingOverNil sets the value for PlaceOfTakingOver to be an explicit nil

### UnsetPlaceOfTakingOver
`func (o *RoadWaybillDto) UnsetPlaceOfTakingOver()`

UnsetPlaceOfTakingOver ensures that no value is present for PlaceOfTakingOver, not even an explicit nil
### GetPlaceOfTakingOverPortId

`func (o *RoadWaybillDto) GetPlaceOfTakingOverPortId() string`

GetPlaceOfTakingOverPortId returns the PlaceOfTakingOverPortId field if non-nil, zero value otherwise.

### GetPlaceOfTakingOverPortIdOk

`func (o *RoadWaybillDto) GetPlaceOfTakingOverPortIdOk() (*string, bool)`

GetPlaceOfTakingOverPortIdOk returns a tuple with the PlaceOfTakingOverPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfTakingOverPortId

`func (o *RoadWaybillDto) SetPlaceOfTakingOverPortId(v string)`

SetPlaceOfTakingOverPortId sets PlaceOfTakingOverPortId field to given value.

### HasPlaceOfTakingOverPortId

`func (o *RoadWaybillDto) HasPlaceOfTakingOverPortId() bool`

HasPlaceOfTakingOverPortId returns a boolean if a field has been set.

### SetPlaceOfTakingOverPortIdNil

`func (o *RoadWaybillDto) SetPlaceOfTakingOverPortIdNil(b bool)`

 SetPlaceOfTakingOverPortIdNil sets the value for PlaceOfTakingOverPortId to be an explicit nil

### UnsetPlaceOfTakingOverPortId
`func (o *RoadWaybillDto) UnsetPlaceOfTakingOverPortId()`

UnsetPlaceOfTakingOverPortId ensures that no value is present for PlaceOfTakingOverPortId, not even an explicit nil
### GetPlaceOfDelivery

`func (o *RoadWaybillDto) GetPlaceOfDelivery() string`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *RoadWaybillDto) GetPlaceOfDeliveryOk() (*string, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *RoadWaybillDto) SetPlaceOfDelivery(v string)`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *RoadWaybillDto) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### SetPlaceOfDeliveryNil

`func (o *RoadWaybillDto) SetPlaceOfDeliveryNil(b bool)`

 SetPlaceOfDeliveryNil sets the value for PlaceOfDelivery to be an explicit nil

### UnsetPlaceOfDelivery
`func (o *RoadWaybillDto) UnsetPlaceOfDelivery()`

UnsetPlaceOfDelivery ensures that no value is present for PlaceOfDelivery, not even an explicit nil
### GetPlaceOfDeliveryPortId

`func (o *RoadWaybillDto) GetPlaceOfDeliveryPortId() string`

GetPlaceOfDeliveryPortId returns the PlaceOfDeliveryPortId field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryPortIdOk

`func (o *RoadWaybillDto) GetPlaceOfDeliveryPortIdOk() (*string, bool)`

GetPlaceOfDeliveryPortIdOk returns a tuple with the PlaceOfDeliveryPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryPortId

`func (o *RoadWaybillDto) SetPlaceOfDeliveryPortId(v string)`

SetPlaceOfDeliveryPortId sets PlaceOfDeliveryPortId field to given value.

### HasPlaceOfDeliveryPortId

`func (o *RoadWaybillDto) HasPlaceOfDeliveryPortId() bool`

HasPlaceOfDeliveryPortId returns a boolean if a field has been set.

### SetPlaceOfDeliveryPortIdNil

`func (o *RoadWaybillDto) SetPlaceOfDeliveryPortIdNil(b bool)`

 SetPlaceOfDeliveryPortIdNil sets the value for PlaceOfDeliveryPortId to be an explicit nil

### UnsetPlaceOfDeliveryPortId
`func (o *RoadWaybillDto) UnsetPlaceOfDeliveryPortId()`

UnsetPlaceOfDeliveryPortId ensures that no value is present for PlaceOfDeliveryPortId, not even an explicit nil
### GetDateOfTakingOver

`func (o *RoadWaybillDto) GetDateOfTakingOver() time.Time`

GetDateOfTakingOver returns the DateOfTakingOver field if non-nil, zero value otherwise.

### GetDateOfTakingOverOk

`func (o *RoadWaybillDto) GetDateOfTakingOverOk() (*time.Time, bool)`

GetDateOfTakingOverOk returns a tuple with the DateOfTakingOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfTakingOver

`func (o *RoadWaybillDto) SetDateOfTakingOver(v time.Time)`

SetDateOfTakingOver sets DateOfTakingOver field to given value.

### HasDateOfTakingOver

`func (o *RoadWaybillDto) HasDateOfTakingOver() bool`

HasDateOfTakingOver returns a boolean if a field has been set.

### SetDateOfTakingOverNil

`func (o *RoadWaybillDto) SetDateOfTakingOverNil(b bool)`

 SetDateOfTakingOverNil sets the value for DateOfTakingOver to be an explicit nil

### UnsetDateOfTakingOver
`func (o *RoadWaybillDto) UnsetDateOfTakingOver()`

UnsetDateOfTakingOver ensures that no value is present for DateOfTakingOver, not even an explicit nil
### GetDateOfDelivery

`func (o *RoadWaybillDto) GetDateOfDelivery() time.Time`

GetDateOfDelivery returns the DateOfDelivery field if non-nil, zero value otherwise.

### GetDateOfDeliveryOk

`func (o *RoadWaybillDto) GetDateOfDeliveryOk() (*time.Time, bool)`

GetDateOfDeliveryOk returns a tuple with the DateOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfDelivery

`func (o *RoadWaybillDto) SetDateOfDelivery(v time.Time)`

SetDateOfDelivery sets DateOfDelivery field to given value.

### HasDateOfDelivery

`func (o *RoadWaybillDto) HasDateOfDelivery() bool`

HasDateOfDelivery returns a boolean if a field has been set.

### SetDateOfDeliveryNil

`func (o *RoadWaybillDto) SetDateOfDeliveryNil(b bool)`

 SetDateOfDeliveryNil sets the value for DateOfDelivery to be an explicit nil

### UnsetDateOfDelivery
`func (o *RoadWaybillDto) UnsetDateOfDelivery()`

UnsetDateOfDelivery ensures that no value is present for DateOfDelivery, not even an explicit nil
### GetFreightTerms

`func (o *RoadWaybillDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *RoadWaybillDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *RoadWaybillDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *RoadWaybillDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *RoadWaybillDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *RoadWaybillDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *RoadWaybillDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *RoadWaybillDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *RoadWaybillDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *RoadWaybillDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *RoadWaybillDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *RoadWaybillDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *RoadWaybillDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *RoadWaybillDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *RoadWaybillDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *RoadWaybillDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *RoadWaybillDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *RoadWaybillDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalGrossWeightKg

`func (o *RoadWaybillDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *RoadWaybillDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *RoadWaybillDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *RoadWaybillDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### SetTotalGrossWeightKgNil

`func (o *RoadWaybillDto) SetTotalGrossWeightKgNil(b bool)`

 SetTotalGrossWeightKgNil sets the value for TotalGrossWeightKg to be an explicit nil

### UnsetTotalGrossWeightKg
`func (o *RoadWaybillDto) UnsetTotalGrossWeightKg()`

UnsetTotalGrossWeightKg ensures that no value is present for TotalGrossWeightKg, not even an explicit nil
### GetTotalPackages

`func (o *RoadWaybillDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *RoadWaybillDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *RoadWaybillDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *RoadWaybillDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *RoadWaybillDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *RoadWaybillDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetTotalVolumeM3

`func (o *RoadWaybillDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *RoadWaybillDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *RoadWaybillDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *RoadWaybillDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *RoadWaybillDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *RoadWaybillDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetAdrDangerousGoods

`func (o *RoadWaybillDto) GetAdrDangerousGoods() bool`

GetAdrDangerousGoods returns the AdrDangerousGoods field if non-nil, zero value otherwise.

### GetAdrDangerousGoodsOk

`func (o *RoadWaybillDto) GetAdrDangerousGoodsOk() (*bool, bool)`

GetAdrDangerousGoodsOk returns a tuple with the AdrDangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrDangerousGoods

`func (o *RoadWaybillDto) SetAdrDangerousGoods(v bool)`

SetAdrDangerousGoods sets AdrDangerousGoods field to given value.

### HasAdrDangerousGoods

`func (o *RoadWaybillDto) HasAdrDangerousGoods() bool`

HasAdrDangerousGoods returns a boolean if a field has been set.

### GetSpecialInstructions

`func (o *RoadWaybillDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *RoadWaybillDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *RoadWaybillDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *RoadWaybillDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *RoadWaybillDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *RoadWaybillDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *RoadWaybillDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *RoadWaybillDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *RoadWaybillDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *RoadWaybillDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *RoadWaybillDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *RoadWaybillDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetSenderSignedDate

`func (o *RoadWaybillDto) GetSenderSignedDate() time.Time`

GetSenderSignedDate returns the SenderSignedDate field if non-nil, zero value otherwise.

### GetSenderSignedDateOk

`func (o *RoadWaybillDto) GetSenderSignedDateOk() (*time.Time, bool)`

GetSenderSignedDateOk returns a tuple with the SenderSignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSignedDate

`func (o *RoadWaybillDto) SetSenderSignedDate(v time.Time)`

SetSenderSignedDate sets SenderSignedDate field to given value.

### HasSenderSignedDate

`func (o *RoadWaybillDto) HasSenderSignedDate() bool`

HasSenderSignedDate returns a boolean if a field has been set.

### SetSenderSignedDateNil

`func (o *RoadWaybillDto) SetSenderSignedDateNil(b bool)`

 SetSenderSignedDateNil sets the value for SenderSignedDate to be an explicit nil

### UnsetSenderSignedDate
`func (o *RoadWaybillDto) UnsetSenderSignedDate()`

UnsetSenderSignedDate ensures that no value is present for SenderSignedDate, not even an explicit nil
### GetCarrierSignedDate

`func (o *RoadWaybillDto) GetCarrierSignedDate() time.Time`

GetCarrierSignedDate returns the CarrierSignedDate field if non-nil, zero value otherwise.

### GetCarrierSignedDateOk

`func (o *RoadWaybillDto) GetCarrierSignedDateOk() (*time.Time, bool)`

GetCarrierSignedDateOk returns a tuple with the CarrierSignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSignedDate

`func (o *RoadWaybillDto) SetCarrierSignedDate(v time.Time)`

SetCarrierSignedDate sets CarrierSignedDate field to given value.

### HasCarrierSignedDate

`func (o *RoadWaybillDto) HasCarrierSignedDate() bool`

HasCarrierSignedDate returns a boolean if a field has been set.

### SetCarrierSignedDateNil

`func (o *RoadWaybillDto) SetCarrierSignedDateNil(b bool)`

 SetCarrierSignedDateNil sets the value for CarrierSignedDate to be an explicit nil

### UnsetCarrierSignedDate
`func (o *RoadWaybillDto) UnsetCarrierSignedDate()`

UnsetCarrierSignedDate ensures that no value is present for CarrierSignedDate, not even an explicit nil
### GetConsigneeSignedDate

`func (o *RoadWaybillDto) GetConsigneeSignedDate() time.Time`

GetConsigneeSignedDate returns the ConsigneeSignedDate field if non-nil, zero value otherwise.

### GetConsigneeSignedDateOk

`func (o *RoadWaybillDto) GetConsigneeSignedDateOk() (*time.Time, bool)`

GetConsigneeSignedDateOk returns a tuple with the ConsigneeSignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeSignedDate

`func (o *RoadWaybillDto) SetConsigneeSignedDate(v time.Time)`

SetConsigneeSignedDate sets ConsigneeSignedDate field to given value.

### HasConsigneeSignedDate

`func (o *RoadWaybillDto) HasConsigneeSignedDate() bool`

HasConsigneeSignedDate returns a boolean if a field has been set.

### SetConsigneeSignedDateNil

`func (o *RoadWaybillDto) SetConsigneeSignedDateNil(b bool)`

 SetConsigneeSignedDateNil sets the value for ConsigneeSignedDate to be an explicit nil

### UnsetConsigneeSignedDate
`func (o *RoadWaybillDto) UnsetConsigneeSignedDate()`

UnsetConsigneeSignedDate ensures that no value is present for ConsigneeSignedDate, not even an explicit nil
### GetShipmentId

`func (o *RoadWaybillDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *RoadWaybillDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *RoadWaybillDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *RoadWaybillDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *RoadWaybillDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *RoadWaybillDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetTruckTripId

`func (o *RoadWaybillDto) GetTruckTripId() string`

GetTruckTripId returns the TruckTripId field if non-nil, zero value otherwise.

### GetTruckTripIdOk

`func (o *RoadWaybillDto) GetTruckTripIdOk() (*string, bool)`

GetTruckTripIdOk returns a tuple with the TruckTripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckTripId

`func (o *RoadWaybillDto) SetTruckTripId(v string)`

SetTruckTripId sets TruckTripId field to given value.

### HasTruckTripId

`func (o *RoadWaybillDto) HasTruckTripId() bool`

HasTruckTripId returns a boolean if a field has been set.

### SetTruckTripIdNil

`func (o *RoadWaybillDto) SetTruckTripIdNil(b bool)`

 SetTruckTripIdNil sets the value for TruckTripId to be an explicit nil

### UnsetTruckTripId
`func (o *RoadWaybillDto) UnsetTruckTripId()`

UnsetTruckTripId ensures that no value is present for TruckTripId, not even an explicit nil
### GetTenantId

`func (o *RoadWaybillDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RoadWaybillDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RoadWaybillDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *RoadWaybillDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *RoadWaybillDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *RoadWaybillDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *RoadWaybillDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *RoadWaybillDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *RoadWaybillDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *RoadWaybillDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *RoadWaybillDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *RoadWaybillDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetLines

`func (o *RoadWaybillDto) GetLines() []WaybillLineDto`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *RoadWaybillDto) GetLinesOk() (*[]WaybillLineDto, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *RoadWaybillDto) SetLines(v []WaybillLineDto)`

SetLines sets Lines field to given value.

### HasLines

`func (o *RoadWaybillDto) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *RoadWaybillDto) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *RoadWaybillDto) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


