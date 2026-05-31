# RoadWaybillUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**DateOfDelivery** | Pointer to **NullableTime** |  | [optional] 
**FreightTerms** | Pointer to **NullableString** |  | [optional] 
**FreightAmount** | Pointer to **NullableFloat64** |  | [optional] 
**FreightCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGrossWeightKg** | Pointer to **NullableFloat64** |  | [optional] 
**TotalPackages** | Pointer to **NullableInt32** |  | [optional] 
**TotalVolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**AdrDangerousGoods** | Pointer to **NullableBool** |  | [optional] 
**SpecialInstructions** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**TruckTripId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRoadWaybillUpdateDto

`func NewRoadWaybillUpdateDto() *RoadWaybillUpdateDto`

NewRoadWaybillUpdateDto instantiates a new RoadWaybillUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoadWaybillUpdateDtoWithDefaults

`func NewRoadWaybillUpdateDtoWithDefaults() *RoadWaybillUpdateDto`

NewRoadWaybillUpdateDtoWithDefaults instantiates a new RoadWaybillUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *RoadWaybillUpdateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *RoadWaybillUpdateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *RoadWaybillUpdateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *RoadWaybillUpdateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *RoadWaybillUpdateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *RoadWaybillUpdateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetRoadWaybillType

`func (o *RoadWaybillUpdateDto) GetRoadWaybillType() string`

GetRoadWaybillType returns the RoadWaybillType field if non-nil, zero value otherwise.

### GetRoadWaybillTypeOk

`func (o *RoadWaybillUpdateDto) GetRoadWaybillTypeOk() (*string, bool)`

GetRoadWaybillTypeOk returns a tuple with the RoadWaybillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadWaybillType

`func (o *RoadWaybillUpdateDto) SetRoadWaybillType(v string)`

SetRoadWaybillType sets RoadWaybillType field to given value.

### HasRoadWaybillType

`func (o *RoadWaybillUpdateDto) HasRoadWaybillType() bool`

HasRoadWaybillType returns a boolean if a field has been set.

### SetRoadWaybillTypeNil

`func (o *RoadWaybillUpdateDto) SetRoadWaybillTypeNil(b bool)`

 SetRoadWaybillTypeNil sets the value for RoadWaybillType to be an explicit nil

### UnsetRoadWaybillType
`func (o *RoadWaybillUpdateDto) UnsetRoadWaybillType()`

UnsetRoadWaybillType ensures that no value is present for RoadWaybillType, not even an explicit nil
### GetShipperContactId

`func (o *RoadWaybillUpdateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *RoadWaybillUpdateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *RoadWaybillUpdateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *RoadWaybillUpdateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *RoadWaybillUpdateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *RoadWaybillUpdateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *RoadWaybillUpdateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *RoadWaybillUpdateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *RoadWaybillUpdateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *RoadWaybillUpdateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *RoadWaybillUpdateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *RoadWaybillUpdateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetCarrierId

`func (o *RoadWaybillUpdateDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *RoadWaybillUpdateDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *RoadWaybillUpdateDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *RoadWaybillUpdateDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *RoadWaybillUpdateDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *RoadWaybillUpdateDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetSuccessiveCarriers

`func (o *RoadWaybillUpdateDto) GetSuccessiveCarriers() string`

GetSuccessiveCarriers returns the SuccessiveCarriers field if non-nil, zero value otherwise.

### GetSuccessiveCarriersOk

`func (o *RoadWaybillUpdateDto) GetSuccessiveCarriersOk() (*string, bool)`

GetSuccessiveCarriersOk returns a tuple with the SuccessiveCarriers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessiveCarriers

`func (o *RoadWaybillUpdateDto) SetSuccessiveCarriers(v string)`

SetSuccessiveCarriers sets SuccessiveCarriers field to given value.

### HasSuccessiveCarriers

`func (o *RoadWaybillUpdateDto) HasSuccessiveCarriers() bool`

HasSuccessiveCarriers returns a boolean if a field has been set.

### SetSuccessiveCarriersNil

`func (o *RoadWaybillUpdateDto) SetSuccessiveCarriersNil(b bool)`

 SetSuccessiveCarriersNil sets the value for SuccessiveCarriers to be an explicit nil

### UnsetSuccessiveCarriers
`func (o *RoadWaybillUpdateDto) UnsetSuccessiveCarriers()`

UnsetSuccessiveCarriers ensures that no value is present for SuccessiveCarriers, not even an explicit nil
### GetTruckId

`func (o *RoadWaybillUpdateDto) GetTruckId() string`

GetTruckId returns the TruckId field if non-nil, zero value otherwise.

### GetTruckIdOk

`func (o *RoadWaybillUpdateDto) GetTruckIdOk() (*string, bool)`

GetTruckIdOk returns a tuple with the TruckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckId

`func (o *RoadWaybillUpdateDto) SetTruckId(v string)`

SetTruckId sets TruckId field to given value.

### HasTruckId

`func (o *RoadWaybillUpdateDto) HasTruckId() bool`

HasTruckId returns a boolean if a field has been set.

### SetTruckIdNil

`func (o *RoadWaybillUpdateDto) SetTruckIdNil(b bool)`

 SetTruckIdNil sets the value for TruckId to be an explicit nil

### UnsetTruckId
`func (o *RoadWaybillUpdateDto) UnsetTruckId()`

UnsetTruckId ensures that no value is present for TruckId, not even an explicit nil
### GetTruckDriverId

`func (o *RoadWaybillUpdateDto) GetTruckDriverId() string`

GetTruckDriverId returns the TruckDriverId field if non-nil, zero value otherwise.

### GetTruckDriverIdOk

`func (o *RoadWaybillUpdateDto) GetTruckDriverIdOk() (*string, bool)`

GetTruckDriverIdOk returns a tuple with the TruckDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckDriverId

`func (o *RoadWaybillUpdateDto) SetTruckDriverId(v string)`

SetTruckDriverId sets TruckDriverId field to given value.

### HasTruckDriverId

`func (o *RoadWaybillUpdateDto) HasTruckDriverId() bool`

HasTruckDriverId returns a boolean if a field has been set.

### SetTruckDriverIdNil

`func (o *RoadWaybillUpdateDto) SetTruckDriverIdNil(b bool)`

 SetTruckDriverIdNil sets the value for TruckDriverId to be an explicit nil

### UnsetTruckDriverId
`func (o *RoadWaybillUpdateDto) UnsetTruckDriverId()`

UnsetTruckDriverId ensures that no value is present for TruckDriverId, not even an explicit nil
### GetVehicleRegistration

`func (o *RoadWaybillUpdateDto) GetVehicleRegistration() string`

GetVehicleRegistration returns the VehicleRegistration field if non-nil, zero value otherwise.

### GetVehicleRegistrationOk

`func (o *RoadWaybillUpdateDto) GetVehicleRegistrationOk() (*string, bool)`

GetVehicleRegistrationOk returns a tuple with the VehicleRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleRegistration

`func (o *RoadWaybillUpdateDto) SetVehicleRegistration(v string)`

SetVehicleRegistration sets VehicleRegistration field to given value.

### HasVehicleRegistration

`func (o *RoadWaybillUpdateDto) HasVehicleRegistration() bool`

HasVehicleRegistration returns a boolean if a field has been set.

### SetVehicleRegistrationNil

`func (o *RoadWaybillUpdateDto) SetVehicleRegistrationNil(b bool)`

 SetVehicleRegistrationNil sets the value for VehicleRegistration to be an explicit nil

### UnsetVehicleRegistration
`func (o *RoadWaybillUpdateDto) UnsetVehicleRegistration()`

UnsetVehicleRegistration ensures that no value is present for VehicleRegistration, not even an explicit nil
### GetTrailerRegistration

`func (o *RoadWaybillUpdateDto) GetTrailerRegistration() string`

GetTrailerRegistration returns the TrailerRegistration field if non-nil, zero value otherwise.

### GetTrailerRegistrationOk

`func (o *RoadWaybillUpdateDto) GetTrailerRegistrationOk() (*string, bool)`

GetTrailerRegistrationOk returns a tuple with the TrailerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerRegistration

`func (o *RoadWaybillUpdateDto) SetTrailerRegistration(v string)`

SetTrailerRegistration sets TrailerRegistration field to given value.

### HasTrailerRegistration

`func (o *RoadWaybillUpdateDto) HasTrailerRegistration() bool`

HasTrailerRegistration returns a boolean if a field has been set.

### SetTrailerRegistrationNil

`func (o *RoadWaybillUpdateDto) SetTrailerRegistrationNil(b bool)`

 SetTrailerRegistrationNil sets the value for TrailerRegistration to be an explicit nil

### UnsetTrailerRegistration
`func (o *RoadWaybillUpdateDto) UnsetTrailerRegistration()`

UnsetTrailerRegistration ensures that no value is present for TrailerRegistration, not even an explicit nil
### GetPlaceOfTakingOver

`func (o *RoadWaybillUpdateDto) GetPlaceOfTakingOver() string`

GetPlaceOfTakingOver returns the PlaceOfTakingOver field if non-nil, zero value otherwise.

### GetPlaceOfTakingOverOk

`func (o *RoadWaybillUpdateDto) GetPlaceOfTakingOverOk() (*string, bool)`

GetPlaceOfTakingOverOk returns a tuple with the PlaceOfTakingOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfTakingOver

`func (o *RoadWaybillUpdateDto) SetPlaceOfTakingOver(v string)`

SetPlaceOfTakingOver sets PlaceOfTakingOver field to given value.

### HasPlaceOfTakingOver

`func (o *RoadWaybillUpdateDto) HasPlaceOfTakingOver() bool`

HasPlaceOfTakingOver returns a boolean if a field has been set.

### SetPlaceOfTakingOverNil

`func (o *RoadWaybillUpdateDto) SetPlaceOfTakingOverNil(b bool)`

 SetPlaceOfTakingOverNil sets the value for PlaceOfTakingOver to be an explicit nil

### UnsetPlaceOfTakingOver
`func (o *RoadWaybillUpdateDto) UnsetPlaceOfTakingOver()`

UnsetPlaceOfTakingOver ensures that no value is present for PlaceOfTakingOver, not even an explicit nil
### GetPlaceOfTakingOverPortId

`func (o *RoadWaybillUpdateDto) GetPlaceOfTakingOverPortId() string`

GetPlaceOfTakingOverPortId returns the PlaceOfTakingOverPortId field if non-nil, zero value otherwise.

### GetPlaceOfTakingOverPortIdOk

`func (o *RoadWaybillUpdateDto) GetPlaceOfTakingOverPortIdOk() (*string, bool)`

GetPlaceOfTakingOverPortIdOk returns a tuple with the PlaceOfTakingOverPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfTakingOverPortId

`func (o *RoadWaybillUpdateDto) SetPlaceOfTakingOverPortId(v string)`

SetPlaceOfTakingOverPortId sets PlaceOfTakingOverPortId field to given value.

### HasPlaceOfTakingOverPortId

`func (o *RoadWaybillUpdateDto) HasPlaceOfTakingOverPortId() bool`

HasPlaceOfTakingOverPortId returns a boolean if a field has been set.

### SetPlaceOfTakingOverPortIdNil

`func (o *RoadWaybillUpdateDto) SetPlaceOfTakingOverPortIdNil(b bool)`

 SetPlaceOfTakingOverPortIdNil sets the value for PlaceOfTakingOverPortId to be an explicit nil

### UnsetPlaceOfTakingOverPortId
`func (o *RoadWaybillUpdateDto) UnsetPlaceOfTakingOverPortId()`

UnsetPlaceOfTakingOverPortId ensures that no value is present for PlaceOfTakingOverPortId, not even an explicit nil
### GetPlaceOfDelivery

`func (o *RoadWaybillUpdateDto) GetPlaceOfDelivery() string`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *RoadWaybillUpdateDto) GetPlaceOfDeliveryOk() (*string, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *RoadWaybillUpdateDto) SetPlaceOfDelivery(v string)`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *RoadWaybillUpdateDto) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### SetPlaceOfDeliveryNil

`func (o *RoadWaybillUpdateDto) SetPlaceOfDeliveryNil(b bool)`

 SetPlaceOfDeliveryNil sets the value for PlaceOfDelivery to be an explicit nil

### UnsetPlaceOfDelivery
`func (o *RoadWaybillUpdateDto) UnsetPlaceOfDelivery()`

UnsetPlaceOfDelivery ensures that no value is present for PlaceOfDelivery, not even an explicit nil
### GetPlaceOfDeliveryPortId

`func (o *RoadWaybillUpdateDto) GetPlaceOfDeliveryPortId() string`

GetPlaceOfDeliveryPortId returns the PlaceOfDeliveryPortId field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryPortIdOk

`func (o *RoadWaybillUpdateDto) GetPlaceOfDeliveryPortIdOk() (*string, bool)`

GetPlaceOfDeliveryPortIdOk returns a tuple with the PlaceOfDeliveryPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryPortId

`func (o *RoadWaybillUpdateDto) SetPlaceOfDeliveryPortId(v string)`

SetPlaceOfDeliveryPortId sets PlaceOfDeliveryPortId field to given value.

### HasPlaceOfDeliveryPortId

`func (o *RoadWaybillUpdateDto) HasPlaceOfDeliveryPortId() bool`

HasPlaceOfDeliveryPortId returns a boolean if a field has been set.

### SetPlaceOfDeliveryPortIdNil

`func (o *RoadWaybillUpdateDto) SetPlaceOfDeliveryPortIdNil(b bool)`

 SetPlaceOfDeliveryPortIdNil sets the value for PlaceOfDeliveryPortId to be an explicit nil

### UnsetPlaceOfDeliveryPortId
`func (o *RoadWaybillUpdateDto) UnsetPlaceOfDeliveryPortId()`

UnsetPlaceOfDeliveryPortId ensures that no value is present for PlaceOfDeliveryPortId, not even an explicit nil
### GetDateOfTakingOver

`func (o *RoadWaybillUpdateDto) GetDateOfTakingOver() time.Time`

GetDateOfTakingOver returns the DateOfTakingOver field if non-nil, zero value otherwise.

### GetDateOfTakingOverOk

`func (o *RoadWaybillUpdateDto) GetDateOfTakingOverOk() (*time.Time, bool)`

GetDateOfTakingOverOk returns a tuple with the DateOfTakingOver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfTakingOver

`func (o *RoadWaybillUpdateDto) SetDateOfTakingOver(v time.Time)`

SetDateOfTakingOver sets DateOfTakingOver field to given value.

### HasDateOfTakingOver

`func (o *RoadWaybillUpdateDto) HasDateOfTakingOver() bool`

HasDateOfTakingOver returns a boolean if a field has been set.

### SetDateOfTakingOverNil

`func (o *RoadWaybillUpdateDto) SetDateOfTakingOverNil(b bool)`

 SetDateOfTakingOverNil sets the value for DateOfTakingOver to be an explicit nil

### UnsetDateOfTakingOver
`func (o *RoadWaybillUpdateDto) UnsetDateOfTakingOver()`

UnsetDateOfTakingOver ensures that no value is present for DateOfTakingOver, not even an explicit nil
### GetDateOfDelivery

`func (o *RoadWaybillUpdateDto) GetDateOfDelivery() time.Time`

GetDateOfDelivery returns the DateOfDelivery field if non-nil, zero value otherwise.

### GetDateOfDeliveryOk

`func (o *RoadWaybillUpdateDto) GetDateOfDeliveryOk() (*time.Time, bool)`

GetDateOfDeliveryOk returns a tuple with the DateOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfDelivery

`func (o *RoadWaybillUpdateDto) SetDateOfDelivery(v time.Time)`

SetDateOfDelivery sets DateOfDelivery field to given value.

### HasDateOfDelivery

`func (o *RoadWaybillUpdateDto) HasDateOfDelivery() bool`

HasDateOfDelivery returns a boolean if a field has been set.

### SetDateOfDeliveryNil

`func (o *RoadWaybillUpdateDto) SetDateOfDeliveryNil(b bool)`

 SetDateOfDeliveryNil sets the value for DateOfDelivery to be an explicit nil

### UnsetDateOfDelivery
`func (o *RoadWaybillUpdateDto) UnsetDateOfDelivery()`

UnsetDateOfDelivery ensures that no value is present for DateOfDelivery, not even an explicit nil
### GetFreightTerms

`func (o *RoadWaybillUpdateDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *RoadWaybillUpdateDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *RoadWaybillUpdateDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *RoadWaybillUpdateDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *RoadWaybillUpdateDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *RoadWaybillUpdateDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *RoadWaybillUpdateDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *RoadWaybillUpdateDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *RoadWaybillUpdateDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *RoadWaybillUpdateDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *RoadWaybillUpdateDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *RoadWaybillUpdateDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *RoadWaybillUpdateDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *RoadWaybillUpdateDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *RoadWaybillUpdateDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *RoadWaybillUpdateDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *RoadWaybillUpdateDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *RoadWaybillUpdateDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalGrossWeightKg

`func (o *RoadWaybillUpdateDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *RoadWaybillUpdateDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *RoadWaybillUpdateDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *RoadWaybillUpdateDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### SetTotalGrossWeightKgNil

`func (o *RoadWaybillUpdateDto) SetTotalGrossWeightKgNil(b bool)`

 SetTotalGrossWeightKgNil sets the value for TotalGrossWeightKg to be an explicit nil

### UnsetTotalGrossWeightKg
`func (o *RoadWaybillUpdateDto) UnsetTotalGrossWeightKg()`

UnsetTotalGrossWeightKg ensures that no value is present for TotalGrossWeightKg, not even an explicit nil
### GetTotalPackages

`func (o *RoadWaybillUpdateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *RoadWaybillUpdateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *RoadWaybillUpdateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *RoadWaybillUpdateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *RoadWaybillUpdateDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *RoadWaybillUpdateDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetTotalVolumeM3

`func (o *RoadWaybillUpdateDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *RoadWaybillUpdateDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *RoadWaybillUpdateDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *RoadWaybillUpdateDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *RoadWaybillUpdateDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *RoadWaybillUpdateDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetAdrDangerousGoods

`func (o *RoadWaybillUpdateDto) GetAdrDangerousGoods() bool`

GetAdrDangerousGoods returns the AdrDangerousGoods field if non-nil, zero value otherwise.

### GetAdrDangerousGoodsOk

`func (o *RoadWaybillUpdateDto) GetAdrDangerousGoodsOk() (*bool, bool)`

GetAdrDangerousGoodsOk returns a tuple with the AdrDangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrDangerousGoods

`func (o *RoadWaybillUpdateDto) SetAdrDangerousGoods(v bool)`

SetAdrDangerousGoods sets AdrDangerousGoods field to given value.

### HasAdrDangerousGoods

`func (o *RoadWaybillUpdateDto) HasAdrDangerousGoods() bool`

HasAdrDangerousGoods returns a boolean if a field has been set.

### SetAdrDangerousGoodsNil

`func (o *RoadWaybillUpdateDto) SetAdrDangerousGoodsNil(b bool)`

 SetAdrDangerousGoodsNil sets the value for AdrDangerousGoods to be an explicit nil

### UnsetAdrDangerousGoods
`func (o *RoadWaybillUpdateDto) UnsetAdrDangerousGoods()`

UnsetAdrDangerousGoods ensures that no value is present for AdrDangerousGoods, not even an explicit nil
### GetSpecialInstructions

`func (o *RoadWaybillUpdateDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *RoadWaybillUpdateDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *RoadWaybillUpdateDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *RoadWaybillUpdateDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *RoadWaybillUpdateDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *RoadWaybillUpdateDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *RoadWaybillUpdateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *RoadWaybillUpdateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *RoadWaybillUpdateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *RoadWaybillUpdateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *RoadWaybillUpdateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *RoadWaybillUpdateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *RoadWaybillUpdateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *RoadWaybillUpdateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *RoadWaybillUpdateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *RoadWaybillUpdateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *RoadWaybillUpdateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *RoadWaybillUpdateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetTruckTripId

`func (o *RoadWaybillUpdateDto) GetTruckTripId() string`

GetTruckTripId returns the TruckTripId field if non-nil, zero value otherwise.

### GetTruckTripIdOk

`func (o *RoadWaybillUpdateDto) GetTruckTripIdOk() (*string, bool)`

GetTruckTripIdOk returns a tuple with the TruckTripId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruckTripId

`func (o *RoadWaybillUpdateDto) SetTruckTripId(v string)`

SetTruckTripId sets TruckTripId field to given value.

### HasTruckTripId

`func (o *RoadWaybillUpdateDto) HasTruckTripId() bool`

HasTruckTripId returns a boolean if a field has been set.

### SetTruckTripIdNil

`func (o *RoadWaybillUpdateDto) SetTruckTripIdNil(b bool)`

 SetTruckTripIdNil sets the value for TruckTripId to be an explicit nil

### UnsetTruckTripId
`func (o *RoadWaybillUpdateDto) UnsetTruckTripId()`

UnsetTruckTripId ensures that no value is present for TruckTripId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


