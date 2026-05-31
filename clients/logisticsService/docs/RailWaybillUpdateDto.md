# RailWaybillUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
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
**ShipmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRailWaybillUpdateDto

`func NewRailWaybillUpdateDto() *RailWaybillUpdateDto`

NewRailWaybillUpdateDto instantiates a new RailWaybillUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRailWaybillUpdateDtoWithDefaults

`func NewRailWaybillUpdateDtoWithDefaults() *RailWaybillUpdateDto`

NewRailWaybillUpdateDtoWithDefaults instantiates a new RailWaybillUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *RailWaybillUpdateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *RailWaybillUpdateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *RailWaybillUpdateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *RailWaybillUpdateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *RailWaybillUpdateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *RailWaybillUpdateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetShipperContactId

`func (o *RailWaybillUpdateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *RailWaybillUpdateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *RailWaybillUpdateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *RailWaybillUpdateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *RailWaybillUpdateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *RailWaybillUpdateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *RailWaybillUpdateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *RailWaybillUpdateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *RailWaybillUpdateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *RailWaybillUpdateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *RailWaybillUpdateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *RailWaybillUpdateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetCarrierId

`func (o *RailWaybillUpdateDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *RailWaybillUpdateDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *RailWaybillUpdateDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *RailWaybillUpdateDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *RailWaybillUpdateDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *RailWaybillUpdateDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetRailOperatorName

`func (o *RailWaybillUpdateDto) GetRailOperatorName() string`

GetRailOperatorName returns the RailOperatorName field if non-nil, zero value otherwise.

### GetRailOperatorNameOk

`func (o *RailWaybillUpdateDto) GetRailOperatorNameOk() (*string, bool)`

GetRailOperatorNameOk returns a tuple with the RailOperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailOperatorName

`func (o *RailWaybillUpdateDto) SetRailOperatorName(v string)`

SetRailOperatorName sets RailOperatorName field to given value.

### HasRailOperatorName

`func (o *RailWaybillUpdateDto) HasRailOperatorName() bool`

HasRailOperatorName returns a boolean if a field has been set.

### SetRailOperatorNameNil

`func (o *RailWaybillUpdateDto) SetRailOperatorNameNil(b bool)`

 SetRailOperatorNameNil sets the value for RailOperatorName to be an explicit nil

### UnsetRailOperatorName
`func (o *RailWaybillUpdateDto) UnsetRailOperatorName()`

UnsetRailOperatorName ensures that no value is present for RailOperatorName, not even an explicit nil
### GetStationOfDeparture

`func (o *RailWaybillUpdateDto) GetStationOfDeparture() string`

GetStationOfDeparture returns the StationOfDeparture field if non-nil, zero value otherwise.

### GetStationOfDepartureOk

`func (o *RailWaybillUpdateDto) GetStationOfDepartureOk() (*string, bool)`

GetStationOfDepartureOk returns a tuple with the StationOfDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDeparture

`func (o *RailWaybillUpdateDto) SetStationOfDeparture(v string)`

SetStationOfDeparture sets StationOfDeparture field to given value.

### HasStationOfDeparture

`func (o *RailWaybillUpdateDto) HasStationOfDeparture() bool`

HasStationOfDeparture returns a boolean if a field has been set.

### SetStationOfDepartureNil

`func (o *RailWaybillUpdateDto) SetStationOfDepartureNil(b bool)`

 SetStationOfDepartureNil sets the value for StationOfDeparture to be an explicit nil

### UnsetStationOfDeparture
`func (o *RailWaybillUpdateDto) UnsetStationOfDeparture()`

UnsetStationOfDeparture ensures that no value is present for StationOfDeparture, not even an explicit nil
### GetStationOfDepartureCode

`func (o *RailWaybillUpdateDto) GetStationOfDepartureCode() string`

GetStationOfDepartureCode returns the StationOfDepartureCode field if non-nil, zero value otherwise.

### GetStationOfDepartureCodeOk

`func (o *RailWaybillUpdateDto) GetStationOfDepartureCodeOk() (*string, bool)`

GetStationOfDepartureCodeOk returns a tuple with the StationOfDepartureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDepartureCode

`func (o *RailWaybillUpdateDto) SetStationOfDepartureCode(v string)`

SetStationOfDepartureCode sets StationOfDepartureCode field to given value.

### HasStationOfDepartureCode

`func (o *RailWaybillUpdateDto) HasStationOfDepartureCode() bool`

HasStationOfDepartureCode returns a boolean if a field has been set.

### SetStationOfDepartureCodeNil

`func (o *RailWaybillUpdateDto) SetStationOfDepartureCodeNil(b bool)`

 SetStationOfDepartureCodeNil sets the value for StationOfDepartureCode to be an explicit nil

### UnsetStationOfDepartureCode
`func (o *RailWaybillUpdateDto) UnsetStationOfDepartureCode()`

UnsetStationOfDepartureCode ensures that no value is present for StationOfDepartureCode, not even an explicit nil
### GetStationOfDestination

`func (o *RailWaybillUpdateDto) GetStationOfDestination() string`

GetStationOfDestination returns the StationOfDestination field if non-nil, zero value otherwise.

### GetStationOfDestinationOk

`func (o *RailWaybillUpdateDto) GetStationOfDestinationOk() (*string, bool)`

GetStationOfDestinationOk returns a tuple with the StationOfDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDestination

`func (o *RailWaybillUpdateDto) SetStationOfDestination(v string)`

SetStationOfDestination sets StationOfDestination field to given value.

### HasStationOfDestination

`func (o *RailWaybillUpdateDto) HasStationOfDestination() bool`

HasStationOfDestination returns a boolean if a field has been set.

### SetStationOfDestinationNil

`func (o *RailWaybillUpdateDto) SetStationOfDestinationNil(b bool)`

 SetStationOfDestinationNil sets the value for StationOfDestination to be an explicit nil

### UnsetStationOfDestination
`func (o *RailWaybillUpdateDto) UnsetStationOfDestination()`

UnsetStationOfDestination ensures that no value is present for StationOfDestination, not even an explicit nil
### GetStationOfDestinationCode

`func (o *RailWaybillUpdateDto) GetStationOfDestinationCode() string`

GetStationOfDestinationCode returns the StationOfDestinationCode field if non-nil, zero value otherwise.

### GetStationOfDestinationCodeOk

`func (o *RailWaybillUpdateDto) GetStationOfDestinationCodeOk() (*string, bool)`

GetStationOfDestinationCodeOk returns a tuple with the StationOfDestinationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDestinationCode

`func (o *RailWaybillUpdateDto) SetStationOfDestinationCode(v string)`

SetStationOfDestinationCode sets StationOfDestinationCode field to given value.

### HasStationOfDestinationCode

`func (o *RailWaybillUpdateDto) HasStationOfDestinationCode() bool`

HasStationOfDestinationCode returns a boolean if a field has been set.

### SetStationOfDestinationCodeNil

`func (o *RailWaybillUpdateDto) SetStationOfDestinationCodeNil(b bool)`

 SetStationOfDestinationCodeNil sets the value for StationOfDestinationCode to be an explicit nil

### UnsetStationOfDestinationCode
`func (o *RailWaybillUpdateDto) UnsetStationOfDestinationCode()`

UnsetStationOfDestinationCode ensures that no value is present for StationOfDestinationCode, not even an explicit nil
### GetPrescribedRoute

`func (o *RailWaybillUpdateDto) GetPrescribedRoute() string`

GetPrescribedRoute returns the PrescribedRoute field if non-nil, zero value otherwise.

### GetPrescribedRouteOk

`func (o *RailWaybillUpdateDto) GetPrescribedRouteOk() (*string, bool)`

GetPrescribedRouteOk returns a tuple with the PrescribedRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescribedRoute

`func (o *RailWaybillUpdateDto) SetPrescribedRoute(v string)`

SetPrescribedRoute sets PrescribedRoute field to given value.

### HasPrescribedRoute

`func (o *RailWaybillUpdateDto) HasPrescribedRoute() bool`

HasPrescribedRoute returns a boolean if a field has been set.

### SetPrescribedRouteNil

`func (o *RailWaybillUpdateDto) SetPrescribedRouteNil(b bool)`

 SetPrescribedRouteNil sets the value for PrescribedRoute to be an explicit nil

### UnsetPrescribedRoute
`func (o *RailWaybillUpdateDto) UnsetPrescribedRoute()`

UnsetPrescribedRoute ensures that no value is present for PrescribedRoute, not even an explicit nil
### GetWagonNumbers

`func (o *RailWaybillUpdateDto) GetWagonNumbers() string`

GetWagonNumbers returns the WagonNumbers field if non-nil, zero value otherwise.

### GetWagonNumbersOk

`func (o *RailWaybillUpdateDto) GetWagonNumbersOk() (*string, bool)`

GetWagonNumbersOk returns a tuple with the WagonNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagonNumbers

`func (o *RailWaybillUpdateDto) SetWagonNumbers(v string)`

SetWagonNumbers sets WagonNumbers field to given value.

### HasWagonNumbers

`func (o *RailWaybillUpdateDto) HasWagonNumbers() bool`

HasWagonNumbers returns a boolean if a field has been set.

### SetWagonNumbersNil

`func (o *RailWaybillUpdateDto) SetWagonNumbersNil(b bool)`

 SetWagonNumbersNil sets the value for WagonNumbers to be an explicit nil

### UnsetWagonNumbers
`func (o *RailWaybillUpdateDto) UnsetWagonNumbers()`

UnsetWagonNumbers ensures that no value is present for WagonNumbers, not even an explicit nil
### GetDateOfAcceptance

`func (o *RailWaybillUpdateDto) GetDateOfAcceptance() time.Time`

GetDateOfAcceptance returns the DateOfAcceptance field if non-nil, zero value otherwise.

### GetDateOfAcceptanceOk

`func (o *RailWaybillUpdateDto) GetDateOfAcceptanceOk() (*time.Time, bool)`

GetDateOfAcceptanceOk returns a tuple with the DateOfAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfAcceptance

`func (o *RailWaybillUpdateDto) SetDateOfAcceptance(v time.Time)`

SetDateOfAcceptance sets DateOfAcceptance field to given value.

### HasDateOfAcceptance

`func (o *RailWaybillUpdateDto) HasDateOfAcceptance() bool`

HasDateOfAcceptance returns a boolean if a field has been set.

### SetDateOfAcceptanceNil

`func (o *RailWaybillUpdateDto) SetDateOfAcceptanceNil(b bool)`

 SetDateOfAcceptanceNil sets the value for DateOfAcceptance to be an explicit nil

### UnsetDateOfAcceptance
`func (o *RailWaybillUpdateDto) UnsetDateOfAcceptance()`

UnsetDateOfAcceptance ensures that no value is present for DateOfAcceptance, not even an explicit nil
### GetDateOfDelivery

`func (o *RailWaybillUpdateDto) GetDateOfDelivery() time.Time`

GetDateOfDelivery returns the DateOfDelivery field if non-nil, zero value otherwise.

### GetDateOfDeliveryOk

`func (o *RailWaybillUpdateDto) GetDateOfDeliveryOk() (*time.Time, bool)`

GetDateOfDeliveryOk returns a tuple with the DateOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfDelivery

`func (o *RailWaybillUpdateDto) SetDateOfDelivery(v time.Time)`

SetDateOfDelivery sets DateOfDelivery field to given value.

### HasDateOfDelivery

`func (o *RailWaybillUpdateDto) HasDateOfDelivery() bool`

HasDateOfDelivery returns a boolean if a field has been set.

### SetDateOfDeliveryNil

`func (o *RailWaybillUpdateDto) SetDateOfDeliveryNil(b bool)`

 SetDateOfDeliveryNil sets the value for DateOfDelivery to be an explicit nil

### UnsetDateOfDelivery
`func (o *RailWaybillUpdateDto) UnsetDateOfDelivery()`

UnsetDateOfDelivery ensures that no value is present for DateOfDelivery, not even an explicit nil
### GetFreightTerms

`func (o *RailWaybillUpdateDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *RailWaybillUpdateDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *RailWaybillUpdateDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *RailWaybillUpdateDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *RailWaybillUpdateDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *RailWaybillUpdateDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *RailWaybillUpdateDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *RailWaybillUpdateDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *RailWaybillUpdateDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *RailWaybillUpdateDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *RailWaybillUpdateDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *RailWaybillUpdateDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *RailWaybillUpdateDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *RailWaybillUpdateDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *RailWaybillUpdateDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *RailWaybillUpdateDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *RailWaybillUpdateDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *RailWaybillUpdateDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalGrossWeightKg

`func (o *RailWaybillUpdateDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *RailWaybillUpdateDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *RailWaybillUpdateDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *RailWaybillUpdateDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### SetTotalGrossWeightKgNil

`func (o *RailWaybillUpdateDto) SetTotalGrossWeightKgNil(b bool)`

 SetTotalGrossWeightKgNil sets the value for TotalGrossWeightKg to be an explicit nil

### UnsetTotalGrossWeightKg
`func (o *RailWaybillUpdateDto) UnsetTotalGrossWeightKg()`

UnsetTotalGrossWeightKg ensures that no value is present for TotalGrossWeightKg, not even an explicit nil
### GetTotalPackages

`func (o *RailWaybillUpdateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *RailWaybillUpdateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *RailWaybillUpdateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *RailWaybillUpdateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *RailWaybillUpdateDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *RailWaybillUpdateDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetTotalVolumeM3

`func (o *RailWaybillUpdateDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *RailWaybillUpdateDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *RailWaybillUpdateDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *RailWaybillUpdateDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *RailWaybillUpdateDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *RailWaybillUpdateDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetCustomsFormalities

`func (o *RailWaybillUpdateDto) GetCustomsFormalities() string`

GetCustomsFormalities returns the CustomsFormalities field if non-nil, zero value otherwise.

### GetCustomsFormalitiesOk

`func (o *RailWaybillUpdateDto) GetCustomsFormalitiesOk() (*string, bool)`

GetCustomsFormalitiesOk returns a tuple with the CustomsFormalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsFormalities

`func (o *RailWaybillUpdateDto) SetCustomsFormalities(v string)`

SetCustomsFormalities sets CustomsFormalities field to given value.

### HasCustomsFormalities

`func (o *RailWaybillUpdateDto) HasCustomsFormalities() bool`

HasCustomsFormalities returns a boolean if a field has been set.

### SetCustomsFormalitiesNil

`func (o *RailWaybillUpdateDto) SetCustomsFormalitiesNil(b bool)`

 SetCustomsFormalitiesNil sets the value for CustomsFormalities to be an explicit nil

### UnsetCustomsFormalities
`func (o *RailWaybillUpdateDto) UnsetCustomsFormalities()`

UnsetCustomsFormalities ensures that no value is present for CustomsFormalities, not even an explicit nil
### GetSpecialInstructions

`func (o *RailWaybillUpdateDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *RailWaybillUpdateDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *RailWaybillUpdateDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *RailWaybillUpdateDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *RailWaybillUpdateDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *RailWaybillUpdateDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *RailWaybillUpdateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *RailWaybillUpdateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *RailWaybillUpdateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *RailWaybillUpdateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *RailWaybillUpdateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *RailWaybillUpdateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *RailWaybillUpdateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *RailWaybillUpdateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *RailWaybillUpdateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *RailWaybillUpdateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *RailWaybillUpdateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *RailWaybillUpdateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


