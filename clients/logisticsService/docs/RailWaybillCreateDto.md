# RailWaybillCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewRailWaybillCreateDto

`func NewRailWaybillCreateDto() *RailWaybillCreateDto`

NewRailWaybillCreateDto instantiates a new RailWaybillCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRailWaybillCreateDtoWithDefaults

`func NewRailWaybillCreateDtoWithDefaults() *RailWaybillCreateDto`

NewRailWaybillCreateDtoWithDefaults instantiates a new RailWaybillCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RailWaybillCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RailWaybillCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RailWaybillCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RailWaybillCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *RailWaybillCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RailWaybillCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RailWaybillCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RailWaybillCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *RailWaybillCreateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *RailWaybillCreateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *RailWaybillCreateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *RailWaybillCreateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *RailWaybillCreateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *RailWaybillCreateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetShipperContactId

`func (o *RailWaybillCreateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *RailWaybillCreateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *RailWaybillCreateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *RailWaybillCreateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *RailWaybillCreateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *RailWaybillCreateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *RailWaybillCreateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *RailWaybillCreateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *RailWaybillCreateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *RailWaybillCreateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *RailWaybillCreateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *RailWaybillCreateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetCarrierId

`func (o *RailWaybillCreateDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *RailWaybillCreateDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *RailWaybillCreateDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *RailWaybillCreateDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *RailWaybillCreateDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *RailWaybillCreateDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetRailOperatorName

`func (o *RailWaybillCreateDto) GetRailOperatorName() string`

GetRailOperatorName returns the RailOperatorName field if non-nil, zero value otherwise.

### GetRailOperatorNameOk

`func (o *RailWaybillCreateDto) GetRailOperatorNameOk() (*string, bool)`

GetRailOperatorNameOk returns a tuple with the RailOperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailOperatorName

`func (o *RailWaybillCreateDto) SetRailOperatorName(v string)`

SetRailOperatorName sets RailOperatorName field to given value.

### HasRailOperatorName

`func (o *RailWaybillCreateDto) HasRailOperatorName() bool`

HasRailOperatorName returns a boolean if a field has been set.

### SetRailOperatorNameNil

`func (o *RailWaybillCreateDto) SetRailOperatorNameNil(b bool)`

 SetRailOperatorNameNil sets the value for RailOperatorName to be an explicit nil

### UnsetRailOperatorName
`func (o *RailWaybillCreateDto) UnsetRailOperatorName()`

UnsetRailOperatorName ensures that no value is present for RailOperatorName, not even an explicit nil
### GetStationOfDeparture

`func (o *RailWaybillCreateDto) GetStationOfDeparture() string`

GetStationOfDeparture returns the StationOfDeparture field if non-nil, zero value otherwise.

### GetStationOfDepartureOk

`func (o *RailWaybillCreateDto) GetStationOfDepartureOk() (*string, bool)`

GetStationOfDepartureOk returns a tuple with the StationOfDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDeparture

`func (o *RailWaybillCreateDto) SetStationOfDeparture(v string)`

SetStationOfDeparture sets StationOfDeparture field to given value.

### HasStationOfDeparture

`func (o *RailWaybillCreateDto) HasStationOfDeparture() bool`

HasStationOfDeparture returns a boolean if a field has been set.

### SetStationOfDepartureNil

`func (o *RailWaybillCreateDto) SetStationOfDepartureNil(b bool)`

 SetStationOfDepartureNil sets the value for StationOfDeparture to be an explicit nil

### UnsetStationOfDeparture
`func (o *RailWaybillCreateDto) UnsetStationOfDeparture()`

UnsetStationOfDeparture ensures that no value is present for StationOfDeparture, not even an explicit nil
### GetStationOfDepartureCode

`func (o *RailWaybillCreateDto) GetStationOfDepartureCode() string`

GetStationOfDepartureCode returns the StationOfDepartureCode field if non-nil, zero value otherwise.

### GetStationOfDepartureCodeOk

`func (o *RailWaybillCreateDto) GetStationOfDepartureCodeOk() (*string, bool)`

GetStationOfDepartureCodeOk returns a tuple with the StationOfDepartureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDepartureCode

`func (o *RailWaybillCreateDto) SetStationOfDepartureCode(v string)`

SetStationOfDepartureCode sets StationOfDepartureCode field to given value.

### HasStationOfDepartureCode

`func (o *RailWaybillCreateDto) HasStationOfDepartureCode() bool`

HasStationOfDepartureCode returns a boolean if a field has been set.

### SetStationOfDepartureCodeNil

`func (o *RailWaybillCreateDto) SetStationOfDepartureCodeNil(b bool)`

 SetStationOfDepartureCodeNil sets the value for StationOfDepartureCode to be an explicit nil

### UnsetStationOfDepartureCode
`func (o *RailWaybillCreateDto) UnsetStationOfDepartureCode()`

UnsetStationOfDepartureCode ensures that no value is present for StationOfDepartureCode, not even an explicit nil
### GetStationOfDestination

`func (o *RailWaybillCreateDto) GetStationOfDestination() string`

GetStationOfDestination returns the StationOfDestination field if non-nil, zero value otherwise.

### GetStationOfDestinationOk

`func (o *RailWaybillCreateDto) GetStationOfDestinationOk() (*string, bool)`

GetStationOfDestinationOk returns a tuple with the StationOfDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDestination

`func (o *RailWaybillCreateDto) SetStationOfDestination(v string)`

SetStationOfDestination sets StationOfDestination field to given value.

### HasStationOfDestination

`func (o *RailWaybillCreateDto) HasStationOfDestination() bool`

HasStationOfDestination returns a boolean if a field has been set.

### SetStationOfDestinationNil

`func (o *RailWaybillCreateDto) SetStationOfDestinationNil(b bool)`

 SetStationOfDestinationNil sets the value for StationOfDestination to be an explicit nil

### UnsetStationOfDestination
`func (o *RailWaybillCreateDto) UnsetStationOfDestination()`

UnsetStationOfDestination ensures that no value is present for StationOfDestination, not even an explicit nil
### GetStationOfDestinationCode

`func (o *RailWaybillCreateDto) GetStationOfDestinationCode() string`

GetStationOfDestinationCode returns the StationOfDestinationCode field if non-nil, zero value otherwise.

### GetStationOfDestinationCodeOk

`func (o *RailWaybillCreateDto) GetStationOfDestinationCodeOk() (*string, bool)`

GetStationOfDestinationCodeOk returns a tuple with the StationOfDestinationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationOfDestinationCode

`func (o *RailWaybillCreateDto) SetStationOfDestinationCode(v string)`

SetStationOfDestinationCode sets StationOfDestinationCode field to given value.

### HasStationOfDestinationCode

`func (o *RailWaybillCreateDto) HasStationOfDestinationCode() bool`

HasStationOfDestinationCode returns a boolean if a field has been set.

### SetStationOfDestinationCodeNil

`func (o *RailWaybillCreateDto) SetStationOfDestinationCodeNil(b bool)`

 SetStationOfDestinationCodeNil sets the value for StationOfDestinationCode to be an explicit nil

### UnsetStationOfDestinationCode
`func (o *RailWaybillCreateDto) UnsetStationOfDestinationCode()`

UnsetStationOfDestinationCode ensures that no value is present for StationOfDestinationCode, not even an explicit nil
### GetPrescribedRoute

`func (o *RailWaybillCreateDto) GetPrescribedRoute() string`

GetPrescribedRoute returns the PrescribedRoute field if non-nil, zero value otherwise.

### GetPrescribedRouteOk

`func (o *RailWaybillCreateDto) GetPrescribedRouteOk() (*string, bool)`

GetPrescribedRouteOk returns a tuple with the PrescribedRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrescribedRoute

`func (o *RailWaybillCreateDto) SetPrescribedRoute(v string)`

SetPrescribedRoute sets PrescribedRoute field to given value.

### HasPrescribedRoute

`func (o *RailWaybillCreateDto) HasPrescribedRoute() bool`

HasPrescribedRoute returns a boolean if a field has been set.

### SetPrescribedRouteNil

`func (o *RailWaybillCreateDto) SetPrescribedRouteNil(b bool)`

 SetPrescribedRouteNil sets the value for PrescribedRoute to be an explicit nil

### UnsetPrescribedRoute
`func (o *RailWaybillCreateDto) UnsetPrescribedRoute()`

UnsetPrescribedRoute ensures that no value is present for PrescribedRoute, not even an explicit nil
### GetWagonNumbers

`func (o *RailWaybillCreateDto) GetWagonNumbers() string`

GetWagonNumbers returns the WagonNumbers field if non-nil, zero value otherwise.

### GetWagonNumbersOk

`func (o *RailWaybillCreateDto) GetWagonNumbersOk() (*string, bool)`

GetWagonNumbersOk returns a tuple with the WagonNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWagonNumbers

`func (o *RailWaybillCreateDto) SetWagonNumbers(v string)`

SetWagonNumbers sets WagonNumbers field to given value.

### HasWagonNumbers

`func (o *RailWaybillCreateDto) HasWagonNumbers() bool`

HasWagonNumbers returns a boolean if a field has been set.

### SetWagonNumbersNil

`func (o *RailWaybillCreateDto) SetWagonNumbersNil(b bool)`

 SetWagonNumbersNil sets the value for WagonNumbers to be an explicit nil

### UnsetWagonNumbers
`func (o *RailWaybillCreateDto) UnsetWagonNumbers()`

UnsetWagonNumbers ensures that no value is present for WagonNumbers, not even an explicit nil
### GetDateOfAcceptance

`func (o *RailWaybillCreateDto) GetDateOfAcceptance() time.Time`

GetDateOfAcceptance returns the DateOfAcceptance field if non-nil, zero value otherwise.

### GetDateOfAcceptanceOk

`func (o *RailWaybillCreateDto) GetDateOfAcceptanceOk() (*time.Time, bool)`

GetDateOfAcceptanceOk returns a tuple with the DateOfAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfAcceptance

`func (o *RailWaybillCreateDto) SetDateOfAcceptance(v time.Time)`

SetDateOfAcceptance sets DateOfAcceptance field to given value.

### HasDateOfAcceptance

`func (o *RailWaybillCreateDto) HasDateOfAcceptance() bool`

HasDateOfAcceptance returns a boolean if a field has been set.

### SetDateOfAcceptanceNil

`func (o *RailWaybillCreateDto) SetDateOfAcceptanceNil(b bool)`

 SetDateOfAcceptanceNil sets the value for DateOfAcceptance to be an explicit nil

### UnsetDateOfAcceptance
`func (o *RailWaybillCreateDto) UnsetDateOfAcceptance()`

UnsetDateOfAcceptance ensures that no value is present for DateOfAcceptance, not even an explicit nil
### GetFreightTerms

`func (o *RailWaybillCreateDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *RailWaybillCreateDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *RailWaybillCreateDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *RailWaybillCreateDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *RailWaybillCreateDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *RailWaybillCreateDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *RailWaybillCreateDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *RailWaybillCreateDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *RailWaybillCreateDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *RailWaybillCreateDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *RailWaybillCreateDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *RailWaybillCreateDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *RailWaybillCreateDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *RailWaybillCreateDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *RailWaybillCreateDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *RailWaybillCreateDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *RailWaybillCreateDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *RailWaybillCreateDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalGrossWeightKg

`func (o *RailWaybillCreateDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *RailWaybillCreateDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *RailWaybillCreateDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *RailWaybillCreateDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### SetTotalGrossWeightKgNil

`func (o *RailWaybillCreateDto) SetTotalGrossWeightKgNil(b bool)`

 SetTotalGrossWeightKgNil sets the value for TotalGrossWeightKg to be an explicit nil

### UnsetTotalGrossWeightKg
`func (o *RailWaybillCreateDto) UnsetTotalGrossWeightKg()`

UnsetTotalGrossWeightKg ensures that no value is present for TotalGrossWeightKg, not even an explicit nil
### GetTotalPackages

`func (o *RailWaybillCreateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *RailWaybillCreateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *RailWaybillCreateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *RailWaybillCreateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *RailWaybillCreateDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *RailWaybillCreateDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetTotalVolumeM3

`func (o *RailWaybillCreateDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *RailWaybillCreateDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *RailWaybillCreateDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *RailWaybillCreateDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *RailWaybillCreateDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *RailWaybillCreateDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetCustomsFormalities

`func (o *RailWaybillCreateDto) GetCustomsFormalities() string`

GetCustomsFormalities returns the CustomsFormalities field if non-nil, zero value otherwise.

### GetCustomsFormalitiesOk

`func (o *RailWaybillCreateDto) GetCustomsFormalitiesOk() (*string, bool)`

GetCustomsFormalitiesOk returns a tuple with the CustomsFormalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsFormalities

`func (o *RailWaybillCreateDto) SetCustomsFormalities(v string)`

SetCustomsFormalities sets CustomsFormalities field to given value.

### HasCustomsFormalities

`func (o *RailWaybillCreateDto) HasCustomsFormalities() bool`

HasCustomsFormalities returns a boolean if a field has been set.

### SetCustomsFormalitiesNil

`func (o *RailWaybillCreateDto) SetCustomsFormalitiesNil(b bool)`

 SetCustomsFormalitiesNil sets the value for CustomsFormalities to be an explicit nil

### UnsetCustomsFormalities
`func (o *RailWaybillCreateDto) UnsetCustomsFormalities()`

UnsetCustomsFormalities ensures that no value is present for CustomsFormalities, not even an explicit nil
### GetSpecialInstructions

`func (o *RailWaybillCreateDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *RailWaybillCreateDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *RailWaybillCreateDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *RailWaybillCreateDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *RailWaybillCreateDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *RailWaybillCreateDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *RailWaybillCreateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *RailWaybillCreateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *RailWaybillCreateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *RailWaybillCreateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *RailWaybillCreateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *RailWaybillCreateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *RailWaybillCreateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *RailWaybillCreateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *RailWaybillCreateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *RailWaybillCreateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *RailWaybillCreateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *RailWaybillCreateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


