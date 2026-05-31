# RailWaybillDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ShipperContactId** | Pointer to **NullableString** |  | [optional] 
**ConsigneeContactId** | Pointer to **NullableString** |  | [optional] 
**CarrierId** | Pointer to **NullableString** |  | [optional] 
**RailOperatorName** | Pointer to **NullableString** |  | [optional] 
**StationOfDeparture** | Pointer to **NullableString** |  | [optional] 
**StationOfDepartureCode** | Pointer to **NullableString** |  | [optional] 
**StationOfDestination** | Pointer to **NullableString** |  | [optional] 
**StationOfDestinationCode** | Pointer to **NullableString** |  | [optional] 
**PrescribedRoute** | Pointer to **NullableString** |  | [optional] 
**WagonNumbers** | Pointer to **NullableString** |  | [optional] 
**DateOfAcceptance** | Pointer to **NullableTime** |  | [optional] 
**DateOfDelivery** | Pointer to **NullableTime** |  | [optional] 
**FreightTerms** | Pointer to **NullableString** |  | [optional] 
**FreightAmount** | Pointer to **NullableFloat64** |  | [optional] 
**FreightCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGrossWeightKg** | Pointer to **NullableFloat64** |  | [optional] 
**TotalPackages** | Pointer to **NullableInt32** |  | [optional] 
**TotalVolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**CustomsFormalities** | Pointer to **NullableString** |  | [optional] 
**SpecialInstructions** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**SenderSignedDate** | Pointer to **NullableTime** |  | [optional] 
**CarrierSignedDate** | Pointer to **NullableTime** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Lines** | Pointer to [**[]WaybillLineDto**](WaybillLineDto.md) |  | [optional] 

## Methods

### NewRailWaybillDto

`func NewRailWaybillDto() *RailWaybillDto`

NewRailWaybillDto instantiates a new RailWaybillDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRailWaybillDtoWithDefaults

`func NewRailWaybillDtoWithDefaults() *RailWaybillDto`

NewRailWaybillDtoWithDefaults instantiates a new RailWaybillDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RailWaybillDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RailWaybillDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RailWaybillDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RailWaybillDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RailWaybillDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RailWaybillDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *RailWaybillDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RailWaybillDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RailWaybillDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RailWaybillDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *RailWaybillDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *RailWaybillDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDocumentNumber

`func (o *RailWaybillDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *RailWaybillDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *RailWaybillDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *RailWaybillDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *RailWaybillDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *RailWaybillDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetStatus

`func (o *RailWaybillDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RailWaybillDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RailWaybillDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RailWaybillDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *RailWaybillDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *RailWaybillDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetShipperContactId

`func (o *RailWaybillDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *RailWaybillDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *RailWaybillDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *RailWaybillDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *RailWaybillDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *RailWaybillDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *RailWaybillDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *RailWaybillDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *RailWaybillDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *RailWaybillDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *RailWaybillDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *RailWaybillDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetCarrierId

`func (o *RailWaybillDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *RailWaybillDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *RailWaybillDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *RailWaybillDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *RailWaybillDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *RailWaybillDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetRailOperatorName

`func (o *RailWaybillDto) GetRailOperatorName() string`

GetRailOperatorName returns the RailOperatorName field if non-nil, zero value otherwise.

### GetRailOperatorNameOk

`func (o *RailWaybillDto) GetRailOperatorNameOk() (*string, bool)`

GetRailOperatorNameOk returns a tuple with the RailOperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailOperatorName

`func (o *RailWaybillDto) SetRailOperatorName(v string)`

SetRailOperatorName sets RailOperatorName field to given value.

### HasRailOperatorName

`func (o *RailWaybillDto) HasRailOperatorName() bool`

HasRailOperatorName returns a boolean if a field has been set.

### SetRailOperatorNameNil

`func (o *RailWaybillDto) SetRailOperatorNameNil(b bool)`

 SetRailOperatorNameNil sets the value for RailOperatorName to be an explicit nil

### UnsetRailOperatorName
`func (o *RailWaybillDto) UnsetRailOperatorName()`

UnsetRailOperatorName ensures that no value is present for RailOperatorName, not even an explicit nil
### GetStationOfDeparture

`func (o *RailWaybillDto) GetStationOfDeparture() string`

GetStationOfDeparture returns the StationOfDeparture field if non-nil, zero value otherwise.

### GetStationOfDepartureOk

`func (o *RailWaybillDto) GetStationOfDepartureOk() (*string, bool)`

GetStationOfDepartureOk returns a tuple with the StationOfDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDeparture

`func (o *RailWaybillDto) SetStationOfDeparture(v string)`

SetStationOfDeparture sets StationOfDeparture field to given value.

### HasStationOfDeparture

`func (o *RailWaybillDto) HasStationOfDeparture() bool`

HasStationOfDeparture returns a boolean if a field has been set.

### SetStationOfDepartureNil

`func (o *RailWaybillDto) SetStationOfDepartureNil(b bool)`

 SetStationOfDepartureNil sets the value for StationOfDeparture to be an explicit nil

### UnsetStationOfDeparture
`func (o *RailWaybillDto) UnsetStationOfDeparture()`

UnsetStationOfDeparture ensures that no value is present for StationOfDeparture, not even an explicit nil
### GetStationOfDepartureCode

`func (o *RailWaybillDto) GetStationOfDepartureCode() string`

GetStationOfDepartureCode returns the StationOfDepartureCode field if non-nil, zero value otherwise.

### GetStationOfDepartureCodeOk

`func (o *RailWaybillDto) GetStationOfDepartureCodeOk() (*string, bool)`

GetStationOfDepartureCodeOk returns a tuple with the StationOfDepartureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDepartureCode

`func (o *RailWaybillDto) SetStationOfDepartureCode(v string)`

SetStationOfDepartureCode sets StationOfDepartureCode field to given value.

### HasStationOfDepartureCode

`func (o *RailWaybillDto) HasStationOfDepartureCode() bool`

HasStationOfDepartureCode returns a boolean if a field has been set.

### SetStationOfDepartureCodeNil

`func (o *RailWaybillDto) SetStationOfDepartureCodeNil(b bool)`

 SetStationOfDepartureCodeNil sets the value for StationOfDepartureCode to be an explicit nil

### UnsetStationOfDepartureCode
`func (o *RailWaybillDto) UnsetStationOfDepartureCode()`

UnsetStationOfDepartureCode ensures that no value is present for StationOfDepartureCode, not even an explicit nil
### GetStationOfDestination

`func (o *RailWaybillDto) GetStationOfDestination() string`

GetStationOfDestination returns the StationOfDestination field if non-nil, zero value otherwise.

### GetStationOfDestinationOk

`func (o *RailWaybillDto) GetStationOfDestinationOk() (*string, bool)`

GetStationOfDestinationOk returns a tuple with the StationOfDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDestination

`func (o *RailWaybillDto) SetStationOfDestination(v string)`

SetStationOfDestination sets StationOfDestination field to given value.

### HasStationOfDestination

`func (o *RailWaybillDto) HasStationOfDestination() bool`

HasStationOfDestination returns a boolean if a field has been set.

### SetStationOfDestinationNil

`func (o *RailWaybillDto) SetStationOfDestinationNil(b bool)`

 SetStationOfDestinationNil sets the value for StationOfDestination to be an explicit nil

### UnsetStationOfDestination
`func (o *RailWaybillDto) UnsetStationOfDestination()`

UnsetStationOfDestination ensures that no value is present for StationOfDestination, not even an explicit nil
### GetStationOfDestinationCode

`func (o *RailWaybillDto) GetStationOfDestinationCode() string`

GetStationOfDestinationCode returns the StationOfDestinationCode field if non-nil, zero value otherwise.

### GetStationOfDestinationCodeOk

`func (o *RailWaybillDto) GetStationOfDestinationCodeOk() (*string, bool)`

GetStationOfDestinationCodeOk returns a tuple with the StationOfDestinationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDestinationCode

`func (o *RailWaybillDto) SetStationOfDestinationCode(v string)`

SetStationOfDestinationCode sets StationOfDestinationCode field to given value.

### HasStationOfDestinationCode

`func (o *RailWaybillDto) HasStationOfDestinationCode() bool`

HasStationOfDestinationCode returns a boolean if a field has been set.

### SetStationOfDestinationCodeNil

`func (o *RailWaybillDto) SetStationOfDestinationCodeNil(b bool)`

 SetStationOfDestinationCodeNil sets the value for StationOfDestinationCode to be an explicit nil

### UnsetStationOfDestinationCode
`func (o *RailWaybillDto) UnsetStationOfDestinationCode()`

UnsetStationOfDestinationCode ensures that no value is present for StationOfDestinationCode, not even an explicit nil
### GetPrescribedRoute

`func (o *RailWaybillDto) GetPrescribedRoute() string`

GetPrescribedRoute returns the PrescribedRoute field if non-nil, zero value otherwise.

### GetPrescribedRouteOk

`func (o *RailWaybillDto) GetPrescribedRouteOk() (*string, bool)`

GetPrescribedRouteOk returns a tuple with the PrescribedRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescribedRoute

`func (o *RailWaybillDto) SetPrescribedRoute(v string)`

SetPrescribedRoute sets PrescribedRoute field to given value.

### HasPrescribedRoute

`func (o *RailWaybillDto) HasPrescribedRoute() bool`

HasPrescribedRoute returns a boolean if a field has been set.

### SetPrescribedRouteNil

`func (o *RailWaybillDto) SetPrescribedRouteNil(b bool)`

 SetPrescribedRouteNil sets the value for PrescribedRoute to be an explicit nil

### UnsetPrescribedRoute
`func (o *RailWaybillDto) UnsetPrescribedRoute()`

UnsetPrescribedRoute ensures that no value is present for PrescribedRoute, not even an explicit nil
### GetWagonNumbers

`func (o *RailWaybillDto) GetWagonNumbers() string`

GetWagonNumbers returns the WagonNumbers field if non-nil, zero value otherwise.

### GetWagonNumbersOk

`func (o *RailWaybillDto) GetWagonNumbersOk() (*string, bool)`

GetWagonNumbersOk returns a tuple with the WagonNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagonNumbers

`func (o *RailWaybillDto) SetWagonNumbers(v string)`

SetWagonNumbers sets WagonNumbers field to given value.

### HasWagonNumbers

`func (o *RailWaybillDto) HasWagonNumbers() bool`

HasWagonNumbers returns a boolean if a field has been set.

### SetWagonNumbersNil

`func (o *RailWaybillDto) SetWagonNumbersNil(b bool)`

 SetWagonNumbersNil sets the value for WagonNumbers to be an explicit nil

### UnsetWagonNumbers
`func (o *RailWaybillDto) UnsetWagonNumbers()`

UnsetWagonNumbers ensures that no value is present for WagonNumbers, not even an explicit nil
### GetDateOfAcceptance

`func (o *RailWaybillDto) GetDateOfAcceptance() time.Time`

GetDateOfAcceptance returns the DateOfAcceptance field if non-nil, zero value otherwise.

### GetDateOfAcceptanceOk

`func (o *RailWaybillDto) GetDateOfAcceptanceOk() (*time.Time, bool)`

GetDateOfAcceptanceOk returns a tuple with the DateOfAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfAcceptance

`func (o *RailWaybillDto) SetDateOfAcceptance(v time.Time)`

SetDateOfAcceptance sets DateOfAcceptance field to given value.

### HasDateOfAcceptance

`func (o *RailWaybillDto) HasDateOfAcceptance() bool`

HasDateOfAcceptance returns a boolean if a field has been set.

### SetDateOfAcceptanceNil

`func (o *RailWaybillDto) SetDateOfAcceptanceNil(b bool)`

 SetDateOfAcceptanceNil sets the value for DateOfAcceptance to be an explicit nil

### UnsetDateOfAcceptance
`func (o *RailWaybillDto) UnsetDateOfAcceptance()`

UnsetDateOfAcceptance ensures that no value is present for DateOfAcceptance, not even an explicit nil
### GetDateOfDelivery

`func (o *RailWaybillDto) GetDateOfDelivery() time.Time`

GetDateOfDelivery returns the DateOfDelivery field if non-nil, zero value otherwise.

### GetDateOfDeliveryOk

`func (o *RailWaybillDto) GetDateOfDeliveryOk() (*time.Time, bool)`

GetDateOfDeliveryOk returns a tuple with the DateOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfDelivery

`func (o *RailWaybillDto) SetDateOfDelivery(v time.Time)`

SetDateOfDelivery sets DateOfDelivery field to given value.

### HasDateOfDelivery

`func (o *RailWaybillDto) HasDateOfDelivery() bool`

HasDateOfDelivery returns a boolean if a field has been set.

### SetDateOfDeliveryNil

`func (o *RailWaybillDto) SetDateOfDeliveryNil(b bool)`

 SetDateOfDeliveryNil sets the value for DateOfDelivery to be an explicit nil

### UnsetDateOfDelivery
`func (o *RailWaybillDto) UnsetDateOfDelivery()`

UnsetDateOfDelivery ensures that no value is present for DateOfDelivery, not even an explicit nil
### GetFreightTerms

`func (o *RailWaybillDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *RailWaybillDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *RailWaybillDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *RailWaybillDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *RailWaybillDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *RailWaybillDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *RailWaybillDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *RailWaybillDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *RailWaybillDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *RailWaybillDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *RailWaybillDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *RailWaybillDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *RailWaybillDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *RailWaybillDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *RailWaybillDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *RailWaybillDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *RailWaybillDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *RailWaybillDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalGrossWeightKg

`func (o *RailWaybillDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *RailWaybillDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *RailWaybillDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *RailWaybillDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### SetTotalGrossWeightKgNil

`func (o *RailWaybillDto) SetTotalGrossWeightKgNil(b bool)`

 SetTotalGrossWeightKgNil sets the value for TotalGrossWeightKg to be an explicit nil

### UnsetTotalGrossWeightKg
`func (o *RailWaybillDto) UnsetTotalGrossWeightKg()`

UnsetTotalGrossWeightKg ensures that no value is present for TotalGrossWeightKg, not even an explicit nil
### GetTotalPackages

`func (o *RailWaybillDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *RailWaybillDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *RailWaybillDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *RailWaybillDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *RailWaybillDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *RailWaybillDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetTotalVolumeM3

`func (o *RailWaybillDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *RailWaybillDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *RailWaybillDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *RailWaybillDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *RailWaybillDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *RailWaybillDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetCustomsFormalities

`func (o *RailWaybillDto) GetCustomsFormalities() string`

GetCustomsFormalities returns the CustomsFormalities field if non-nil, zero value otherwise.

### GetCustomsFormalitiesOk

`func (o *RailWaybillDto) GetCustomsFormalitiesOk() (*string, bool)`

GetCustomsFormalitiesOk returns a tuple with the CustomsFormalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsFormalities

`func (o *RailWaybillDto) SetCustomsFormalities(v string)`

SetCustomsFormalities sets CustomsFormalities field to given value.

### HasCustomsFormalities

`func (o *RailWaybillDto) HasCustomsFormalities() bool`

HasCustomsFormalities returns a boolean if a field has been set.

### SetCustomsFormalitiesNil

`func (o *RailWaybillDto) SetCustomsFormalitiesNil(b bool)`

 SetCustomsFormalitiesNil sets the value for CustomsFormalities to be an explicit nil

### UnsetCustomsFormalities
`func (o *RailWaybillDto) UnsetCustomsFormalities()`

UnsetCustomsFormalities ensures that no value is present for CustomsFormalities, not even an explicit nil
### GetSpecialInstructions

`func (o *RailWaybillDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *RailWaybillDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *RailWaybillDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *RailWaybillDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *RailWaybillDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *RailWaybillDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *RailWaybillDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *RailWaybillDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *RailWaybillDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *RailWaybillDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *RailWaybillDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *RailWaybillDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetSenderSignedDate

`func (o *RailWaybillDto) GetSenderSignedDate() time.Time`

GetSenderSignedDate returns the SenderSignedDate field if non-nil, zero value otherwise.

### GetSenderSignedDateOk

`func (o *RailWaybillDto) GetSenderSignedDateOk() (*time.Time, bool)`

GetSenderSignedDateOk returns a tuple with the SenderSignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderSignedDate

`func (o *RailWaybillDto) SetSenderSignedDate(v time.Time)`

SetSenderSignedDate sets SenderSignedDate field to given value.

### HasSenderSignedDate

`func (o *RailWaybillDto) HasSenderSignedDate() bool`

HasSenderSignedDate returns a boolean if a field has been set.

### SetSenderSignedDateNil

`func (o *RailWaybillDto) SetSenderSignedDateNil(b bool)`

 SetSenderSignedDateNil sets the value for SenderSignedDate to be an explicit nil

### UnsetSenderSignedDate
`func (o *RailWaybillDto) UnsetSenderSignedDate()`

UnsetSenderSignedDate ensures that no value is present for SenderSignedDate, not even an explicit nil
### GetCarrierSignedDate

`func (o *RailWaybillDto) GetCarrierSignedDate() time.Time`

GetCarrierSignedDate returns the CarrierSignedDate field if non-nil, zero value otherwise.

### GetCarrierSignedDateOk

`func (o *RailWaybillDto) GetCarrierSignedDateOk() (*time.Time, bool)`

GetCarrierSignedDateOk returns a tuple with the CarrierSignedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSignedDate

`func (o *RailWaybillDto) SetCarrierSignedDate(v time.Time)`

SetCarrierSignedDate sets CarrierSignedDate field to given value.

### HasCarrierSignedDate

`func (o *RailWaybillDto) HasCarrierSignedDate() bool`

HasCarrierSignedDate returns a boolean if a field has been set.

### SetCarrierSignedDateNil

`func (o *RailWaybillDto) SetCarrierSignedDateNil(b bool)`

 SetCarrierSignedDateNil sets the value for CarrierSignedDate to be an explicit nil

### UnsetCarrierSignedDate
`func (o *RailWaybillDto) UnsetCarrierSignedDate()`

UnsetCarrierSignedDate ensures that no value is present for CarrierSignedDate, not even an explicit nil
### GetShipmentId

`func (o *RailWaybillDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *RailWaybillDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *RailWaybillDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *RailWaybillDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *RailWaybillDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *RailWaybillDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetTenantId

`func (o *RailWaybillDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *RailWaybillDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *RailWaybillDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *RailWaybillDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *RailWaybillDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *RailWaybillDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *RailWaybillDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *RailWaybillDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *RailWaybillDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *RailWaybillDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *RailWaybillDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *RailWaybillDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetLines

`func (o *RailWaybillDto) GetLines() []WaybillLineDto`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *RailWaybillDto) GetLinesOk() (*[]WaybillLineDto, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *RailWaybillDto) SetLines(v []WaybillLineDto)`

SetLines sets Lines field to given value.

### HasLines

`func (o *RailWaybillDto) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *RailWaybillDto) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *RailWaybillDto) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


