# RoadWaybillCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**RoadWaybillType** | Pointer to **NullableString** |  | [optional] 
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
**FreightTerms** | Pointer to **NullableString** |  | [optional] 
**FreightAmount** | Pointer to **NullableFloat64** |  | [optional] 
**FreightCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGrossWeightKg** | Pointer to **NullableFloat64** |  | [optional] 
**TotalPackages** | Pointer to **NullableInt32** |  | [optional] 
**TotalVolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**AdrDangerousGoods** | Pointer to **bool** |  | [optional] 
**SpecialInstructions** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**TruckTripId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRoadWaybillCreateDto

`func NewRoadWaybillCreateDto() *RoadWaybillCreateDto`

NewRoadWaybillCreateDto instantiates a new RoadWaybillCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoadWaybillCreateDtoWithDefaults

`func NewRoadWaybillCreateDtoWithDefaults() *RoadWaybillCreateDto`

NewRoadWaybillCreateDtoWithDefaults instantiates a new RoadWaybillCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoadWaybillCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoadWaybillCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoadWaybillCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoadWaybillCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *RoadWaybillCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RoadWaybillCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RoadWaybillCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RoadWaybillCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *RoadWaybillCreateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *RoadWaybillCreateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *RoadWaybillCreateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *RoadWaybillCreateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *RoadWaybillCreateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *RoadWaybillCreateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetRoadWaybillType

`func (o *RoadWaybillCreateDto) GetRoadWaybillType() string`

GetRoadWaybillType returns the RoadWaybillType field if non-nil, zero value otherwise.

### GetRoadWaybillTypeOk

`func (o *RoadWaybillCreateDto) GetRoadWaybillTypeOk() (*string, bool)`

GetRoadWaybillTypeOk returns a tuple with the RoadWaybillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadWaybillType

`func (o *RoadWaybillCreateDto) SetRoadWaybillType(v string)`

SetRoadWaybillType sets RoadWaybillType field to given value.

### HasRoadWaybillType

`func (o *RoadWaybillCreateDto) HasRoadWaybillType() bool`

HasRoadWaybillType returns a boolean if a field has been set.

### SetRoadWaybillTypeNil

`func (o *RoadWaybillCreateDto) SetRoadWaybillTypeNil(b bool)`

 SetRoadWaybillTypeNil sets the value for RoadWaybillType to be an explicit nil

### UnsetRoadWaybillType
`func (o *RoadWaybillCreateDto) UnsetRoadWaybillType()`

UnsetRoadWaybillType ensures that no value is present for RoadWaybillType, not even an explicit nil
### GetShipperContactId

`func (o *RoadWaybillCreateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *RoadWaybillCreateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *RoadWaybillCreateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *RoadWaybillCreateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *RoadWaybillCreateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *RoadWaybillCreateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *RoadWaybillCreateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *RoadWaybillCreateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *RoadWaybillCreateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *RoadWaybillCreateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *RoadWaybillCreateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *RoadWaybillCreateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetCarrierId

`func (o *RoadWaybillCreateDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *RoadWaybillCreateDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *RoadWaybillCreateDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *RoadWaybillCreateDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *RoadWaybillCreateDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *RoadWaybillCreateDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetSuccessiveCarriers

`func (o *RoadWaybillCreateDto) GetSuccessiveCarriers() string`

GetSuccessiveCarriers returns the SuccessiveCarriers field if non-nil, zero value otherwise.

### GetSuccessiveCarriersOk

`func (o *RoadWaybillCreateDto) GetSuccessiveCarriersOk() (*string, bool)`

GetSuccessiveCarriersOk returns a tuple with the SuccessiveCarriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessiveCarriers

`func (o *RoadWaybillCreateDto) SetSuccessiveCarriers(v string)`

SetSuccessiveCarriers sets SuccessiveCarriers field to given value.

### HasSuccessiveCarriers

`func (o *RoadWaybillCreateDto) HasSuccessiveCarriers() bool`

HasSuccessiveCarriers returns a boolean if a field has been set.

### SetSuccessiveCarriersNil

`func (o *RoadWaybillCreateDto) SetSuccessiveCarriersNil(b bool)`

 SetSuccessiveCarriersNil sets the value for SuccessiveCarriers to be an explicit nil

### UnsetSuccessiveCarriers
`func (o *RoadWaybillCreateDto) UnsetSuccessiveCarriers()`

UnsetSuccessiveCarriers ensures that no value is present for SuccessiveCarriers, not even an explicit nil
### GetTruckId

`func (o *RoadWaybillCreateDto) GetTruckId() string`

GetTruckId returns the TruckId field if non-nil, zero value otherwise.

### GetTruckIdOk

`func (o *RoadWaybillCreateDto) GetTruckIdOk() (*string, bool)`

GetTruckIdOk returns a tuple with the TruckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckId

`func (o *RoadWaybillCreateDto) SetTruckId(v string)`

SetTruckId sets TruckId field to given value.

### HasTruckId

`func (o *RoadWaybillCreateDto) HasTruckId() bool`

HasTruckId returns a boolean if a field has been set.

### SetTruckIdNil

`func (o *RoadWaybillCreateDto) SetTruckIdNil(b bool)`

 SetTruckIdNil sets the value for TruckId to be an explicit nil

### UnsetTruckId
`func (o *RoadWaybillCreateDto) UnsetTruckId()`

UnsetTruckId ensures that no value is present for TruckId, not even an explicit nil
### GetTruckDriverId

`func (o *RoadWaybillCreateDto) GetTruckDriverId() string`

GetTruckDriverId returns the TruckDriverId field if non-nil, zero value otherwise.

### GetTruckDriverIdOk

`func (o *RoadWaybillCreateDto) GetTruckDriverIdOk() (*string, bool)`

GetTruckDriverIdOk returns a tuple with the TruckDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckDriverId

`func (o *RoadWaybillCreateDto) SetTruckDriverId(v string)`

SetTruckDriverId sets TruckDriverId field to given value.

### HasTruckDriverId

`func (o *RoadWaybillCreateDto) HasTruckDriverId() bool`

HasTruckDriverId returns a boolean if a field has been set.

### SetTruckDriverIdNil

`func (o *RoadWaybillCreateDto) SetTruckDriverIdNil(b bool)`

 SetTruckDriverIdNil sets the value for TruckDriverId to be an explicit nil

### UnsetTruckDriverId
`func (o *RoadWaybillCreateDto) UnsetTruckDriverId()`

UnsetTruckDriverId ensures that no value is present for TruckDriverId, not even an explicit nil
### GetVehicleRegistration

`func (o *RoadWaybillCreateDto) GetVehicleRegistration() string`

GetVehicleRegistration returns the VehicleRegistration field if non-nil, zero value otherwise.

### GetVehicleRegistrationOk

`func (o *RoadWaybillCreateDto) GetVehicleRegistrationOk() (*string, bool)`

GetVehicleRegistrationOk returns a tuple with the VehicleRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleRegistration

`func (o *RoadWaybillCreateDto) SetVehicleRegistration(v string)`

SetVehicleRegistration sets VehicleRegistration field to given value.

### HasVehicleRegistration

`func (o *RoadWaybillCreateDto) HasVehicleRegistration() bool`

HasVehicleRegistration returns a boolean if a field has been set.

### SetVehicleRegistrationNil

`func (o *RoadWaybillCreateDto) SetVehicleRegistrationNil(b bool)`

 SetVehicleRegistrationNil sets the value for VehicleRegistration to be an explicit nil

### UnsetVehicleRegistration
`func (o *RoadWaybillCreateDto) UnsetVehicleRegistration()`

UnsetVehicleRegistration ensures that no value is present for VehicleRegistration, not even an explicit nil
### GetTrailerRegistration

`func (o *RoadWaybillCreateDto) GetTrailerRegistration() string`

GetTrailerRegistration returns the TrailerRegistration field if non-nil, zero value otherwise.

### GetTrailerRegistrationOk

`func (o *RoadWaybillCreateDto) GetTrailerRegistrationOk() (*string, bool)`

GetTrailerRegistrationOk returns a tuple with the TrailerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerRegistration

`func (o *RoadWaybillCreateDto) SetTrailerRegistration(v string)`

SetTrailerRegistration sets TrailerRegistration field to given value.

### HasTrailerRegistration

`func (o *RoadWaybillCreateDto) HasTrailerRegistration() bool`

HasTrailerRegistration returns a boolean if a field has been set.

### SetTrailerRegistrationNil

`func (o *RoadWaybillCreateDto) SetTrailerRegistrationNil(b bool)`

 SetTrailerRegistrationNil sets the value for TrailerRegistration to be an explicit nil

### UnsetTrailerRegistration
`func (o *RoadWaybillCreateDto) UnsetTrailerRegistration()`

UnsetTrailerRegistration ensures that no value is present for TrailerRegistration, not even an explicit nil
### GetPlaceOfTakingOver

`func (o *RoadWaybillCreateDto) GetPlaceOfTakingOver() string`

GetPlaceOfTakingOver returns the PlaceOfTakingOver field if non-nil, zero value otherwise.

### GetPlaceOfTakingOverOk

`func (o *RoadWaybillCreateDto) GetPlaceOfTakingOverOk() (*string, bool)`

GetPlaceOfTakingOverOk returns a tuple with the PlaceOfTakingOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfTakingOver

`func (o *RoadWaybillCreateDto) SetPlaceOfTakingOver(v string)`

SetPlaceOfTakingOver sets PlaceOfTakingOver field to given value.

### HasPlaceOfTakingOver

`func (o *RoadWaybillCreateDto) HasPlaceOfTakingOver() bool`

HasPlaceOfTakingOver returns a boolean if a field has been set.

### SetPlaceOfTakingOverNil

`func (o *RoadWaybillCreateDto) SetPlaceOfTakingOverNil(b bool)`

 SetPlaceOfTakingOverNil sets the value for PlaceOfTakingOver to be an explicit nil

### UnsetPlaceOfTakingOver
`func (o *RoadWaybillCreateDto) UnsetPlaceOfTakingOver()`

UnsetPlaceOfTakingOver ensures that no value is present for PlaceOfTakingOver, not even an explicit nil
### GetPlaceOfTakingOverPortId

`func (o *RoadWaybillCreateDto) GetPlaceOfTakingOverPortId() string`

GetPlaceOfTakingOverPortId returns the PlaceOfTakingOverPortId field if non-nil, zero value otherwise.

### GetPlaceOfTakingOverPortIdOk

`func (o *RoadWaybillCreateDto) GetPlaceOfTakingOverPortIdOk() (*string, bool)`

GetPlaceOfTakingOverPortIdOk returns a tuple with the PlaceOfTakingOverPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfTakingOverPortId

`func (o *RoadWaybillCreateDto) SetPlaceOfTakingOverPortId(v string)`

SetPlaceOfTakingOverPortId sets PlaceOfTakingOverPortId field to given value.

### HasPlaceOfTakingOverPortId

`func (o *RoadWaybillCreateDto) HasPlaceOfTakingOverPortId() bool`

HasPlaceOfTakingOverPortId returns a boolean if a field has been set.

### SetPlaceOfTakingOverPortIdNil

`func (o *RoadWaybillCreateDto) SetPlaceOfTakingOverPortIdNil(b bool)`

 SetPlaceOfTakingOverPortIdNil sets the value for PlaceOfTakingOverPortId to be an explicit nil

### UnsetPlaceOfTakingOverPortId
`func (o *RoadWaybillCreateDto) UnsetPlaceOfTakingOverPortId()`

UnsetPlaceOfTakingOverPortId ensures that no value is present for PlaceOfTakingOverPortId, not even an explicit nil
### GetPlaceOfDelivery

`func (o *RoadWaybillCreateDto) GetPlaceOfDelivery() string`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *RoadWaybillCreateDto) GetPlaceOfDeliveryOk() (*string, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *RoadWaybillCreateDto) SetPlaceOfDelivery(v string)`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *RoadWaybillCreateDto) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### SetPlaceOfDeliveryNil

`func (o *RoadWaybillCreateDto) SetPlaceOfDeliveryNil(b bool)`

 SetPlaceOfDeliveryNil sets the value for PlaceOfDelivery to be an explicit nil

### UnsetPlaceOfDelivery
`func (o *RoadWaybillCreateDto) UnsetPlaceOfDelivery()`

UnsetPlaceOfDelivery ensures that no value is present for PlaceOfDelivery, not even an explicit nil
### GetPlaceOfDeliveryPortId

`func (o *RoadWaybillCreateDto) GetPlaceOfDeliveryPortId() string`

GetPlaceOfDeliveryPortId returns the PlaceOfDeliveryPortId field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryPortIdOk

`func (o *RoadWaybillCreateDto) GetPlaceOfDeliveryPortIdOk() (*string, bool)`

GetPlaceOfDeliveryPortIdOk returns a tuple with the PlaceOfDeliveryPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryPortId

`func (o *RoadWaybillCreateDto) SetPlaceOfDeliveryPortId(v string)`

SetPlaceOfDeliveryPortId sets PlaceOfDeliveryPortId field to given value.

### HasPlaceOfDeliveryPortId

`func (o *RoadWaybillCreateDto) HasPlaceOfDeliveryPortId() bool`

HasPlaceOfDeliveryPortId returns a boolean if a field has been set.

### SetPlaceOfDeliveryPortIdNil

`func (o *RoadWaybillCreateDto) SetPlaceOfDeliveryPortIdNil(b bool)`

 SetPlaceOfDeliveryPortIdNil sets the value for PlaceOfDeliveryPortId to be an explicit nil

### UnsetPlaceOfDeliveryPortId
`func (o *RoadWaybillCreateDto) UnsetPlaceOfDeliveryPortId()`

UnsetPlaceOfDeliveryPortId ensures that no value is present for PlaceOfDeliveryPortId, not even an explicit nil
### GetDateOfTakingOver

`func (o *RoadWaybillCreateDto) GetDateOfTakingOver() time.Time`

GetDateOfTakingOver returns the DateOfTakingOver field if non-nil, zero value otherwise.

### GetDateOfTakingOverOk

`func (o *RoadWaybillCreateDto) GetDateOfTakingOverOk() (*time.Time, bool)`

GetDateOfTakingOverOk returns a tuple with the DateOfTakingOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfTakingOver

`func (o *RoadWaybillCreateDto) SetDateOfTakingOver(v time.Time)`

SetDateOfTakingOver sets DateOfTakingOver field to given value.

### HasDateOfTakingOver

`func (o *RoadWaybillCreateDto) HasDateOfTakingOver() bool`

HasDateOfTakingOver returns a boolean if a field has been set.

### SetDateOfTakingOverNil

`func (o *RoadWaybillCreateDto) SetDateOfTakingOverNil(b bool)`

 SetDateOfTakingOverNil sets the value for DateOfTakingOver to be an explicit nil

### UnsetDateOfTakingOver
`func (o *RoadWaybillCreateDto) UnsetDateOfTakingOver()`

UnsetDateOfTakingOver ensures that no value is present for DateOfTakingOver, not even an explicit nil
### GetFreightTerms

`func (o *RoadWaybillCreateDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *RoadWaybillCreateDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *RoadWaybillCreateDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *RoadWaybillCreateDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *RoadWaybillCreateDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *RoadWaybillCreateDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *RoadWaybillCreateDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *RoadWaybillCreateDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *RoadWaybillCreateDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *RoadWaybillCreateDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *RoadWaybillCreateDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *RoadWaybillCreateDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *RoadWaybillCreateDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *RoadWaybillCreateDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *RoadWaybillCreateDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *RoadWaybillCreateDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *RoadWaybillCreateDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *RoadWaybillCreateDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalGrossWeightKg

`func (o *RoadWaybillCreateDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *RoadWaybillCreateDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *RoadWaybillCreateDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *RoadWaybillCreateDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### SetTotalGrossWeightKgNil

`func (o *RoadWaybillCreateDto) SetTotalGrossWeightKgNil(b bool)`

 SetTotalGrossWeightKgNil sets the value for TotalGrossWeightKg to be an explicit nil

### UnsetTotalGrossWeightKg
`func (o *RoadWaybillCreateDto) UnsetTotalGrossWeightKg()`

UnsetTotalGrossWeightKg ensures that no value is present for TotalGrossWeightKg, not even an explicit nil
### GetTotalPackages

`func (o *RoadWaybillCreateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *RoadWaybillCreateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *RoadWaybillCreateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *RoadWaybillCreateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *RoadWaybillCreateDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *RoadWaybillCreateDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetTotalVolumeM3

`func (o *RoadWaybillCreateDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *RoadWaybillCreateDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *RoadWaybillCreateDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *RoadWaybillCreateDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *RoadWaybillCreateDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *RoadWaybillCreateDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetAdrDangerousGoods

`func (o *RoadWaybillCreateDto) GetAdrDangerousGoods() bool`

GetAdrDangerousGoods returns the AdrDangerousGoods field if non-nil, zero value otherwise.

### GetAdrDangerousGoodsOk

`func (o *RoadWaybillCreateDto) GetAdrDangerousGoodsOk() (*bool, bool)`

GetAdrDangerousGoodsOk returns a tuple with the AdrDangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrDangerousGoods

`func (o *RoadWaybillCreateDto) SetAdrDangerousGoods(v bool)`

SetAdrDangerousGoods sets AdrDangerousGoods field to given value.

### HasAdrDangerousGoods

`func (o *RoadWaybillCreateDto) HasAdrDangerousGoods() bool`

HasAdrDangerousGoods returns a boolean if a field has been set.

### GetSpecialInstructions

`func (o *RoadWaybillCreateDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *RoadWaybillCreateDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *RoadWaybillCreateDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *RoadWaybillCreateDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *RoadWaybillCreateDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *RoadWaybillCreateDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *RoadWaybillCreateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *RoadWaybillCreateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *RoadWaybillCreateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *RoadWaybillCreateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *RoadWaybillCreateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *RoadWaybillCreateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *RoadWaybillCreateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *RoadWaybillCreateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *RoadWaybillCreateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *RoadWaybillCreateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *RoadWaybillCreateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *RoadWaybillCreateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetTruckTripId

`func (o *RoadWaybillCreateDto) GetTruckTripId() string`

GetTruckTripId returns the TruckTripId field if non-nil, zero value otherwise.

### GetTruckTripIdOk

`func (o *RoadWaybillCreateDto) GetTruckTripIdOk() (*string, bool)`

GetTruckTripIdOk returns a tuple with the TruckTripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckTripId

`func (o *RoadWaybillCreateDto) SetTruckTripId(v string)`

SetTruckTripId sets TruckTripId field to given value.

### HasTruckTripId

`func (o *RoadWaybillCreateDto) HasTruckTripId() bool`

HasTruckTripId returns a boolean if a field has been set.

### SetTruckTripIdNil

`func (o *RoadWaybillCreateDto) SetTruckTripIdNil(b bool)`

 SetTruckTripIdNil sets the value for TruckTripId to be an explicit nil

### UnsetTruckTripId
`func (o *RoadWaybillCreateDto) UnsetTruckTripId()`

UnsetTruckTripId ensures that no value is present for TruckTripId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


