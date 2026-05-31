# AirwayBillUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**AirwayBillType** | Pointer to **NullableString** |  | [optional] 
**MasterAwbNumber** | Pointer to **NullableString** |  | [optional] 
**ShipperContactId** | Pointer to **NullableString** |  | [optional] 
**ConsigneeContactId** | Pointer to **NullableString** |  | [optional] 
**NotifyPartyContactId** | Pointer to **NullableString** |  | [optional] 
**CarrierId** | Pointer to **NullableString** |  | [optional] 
**AirlineCode** | Pointer to **NullableString** |  | [optional] 
**FlightNumber** | Pointer to **NullableString** |  | [optional] 
**AirportOfDepartureCode** | Pointer to **NullableString** |  | [optional] 
**AirportOfDestinationCode** | Pointer to **NullableString** |  | [optional] 
**DepartureDate** | Pointer to **NullableTime** |  | [optional] 
**ArrivalDate** | Pointer to **NullableTime** |  | [optional] 
**DateIssued** | Pointer to **NullableTime** |  | [optional] 
**FreightTerms** | Pointer to **NullableString** |  | [optional] 
**FreightAmount** | Pointer to **NullableFloat64** |  | [optional] 
**FreightCurrencyId** | Pointer to **NullableString** |  | [optional] 
**ChargeableWeightKg** | Pointer to **NullableFloat64** |  | [optional] 
**TotalGrossWeightKg** | Pointer to **NullableFloat64** |  | [optional] 
**TotalPackages** | Pointer to **NullableInt32** |  | [optional] 
**TotalVolumeM3** | Pointer to **NullableFloat64** |  | [optional] 
**DeclaredValueForCarriage** | Pointer to **NullableFloat64** |  | [optional] 
**DeclaredValueForCustoms** | Pointer to **NullableFloat64** |  | [optional] 
**InsuranceAmount** | Pointer to **NullableFloat64** |  | [optional] 
**SpecialHandlingCodes** | Pointer to **NullableString** |  | [optional] 
**SpecialInstructions** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAirwayBillUpdateDto

`func NewAirwayBillUpdateDto() *AirwayBillUpdateDto`

NewAirwayBillUpdateDto instantiates a new AirwayBillUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirwayBillUpdateDtoWithDefaults

`func NewAirwayBillUpdateDtoWithDefaults() *AirwayBillUpdateDto`

NewAirwayBillUpdateDtoWithDefaults instantiates a new AirwayBillUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *AirwayBillUpdateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *AirwayBillUpdateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *AirwayBillUpdateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *AirwayBillUpdateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *AirwayBillUpdateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *AirwayBillUpdateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetAirwayBillType

`func (o *AirwayBillUpdateDto) GetAirwayBillType() string`

GetAirwayBillType returns the AirwayBillType field if non-nil, zero value otherwise.

### GetAirwayBillTypeOk

`func (o *AirwayBillUpdateDto) GetAirwayBillTypeOk() (*string, bool)`

GetAirwayBillTypeOk returns a tuple with the AirwayBillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwayBillType

`func (o *AirwayBillUpdateDto) SetAirwayBillType(v string)`

SetAirwayBillType sets AirwayBillType field to given value.

### HasAirwayBillType

`func (o *AirwayBillUpdateDto) HasAirwayBillType() bool`

HasAirwayBillType returns a boolean if a field has been set.

### SetAirwayBillTypeNil

`func (o *AirwayBillUpdateDto) SetAirwayBillTypeNil(b bool)`

 SetAirwayBillTypeNil sets the value for AirwayBillType to be an explicit nil

### UnsetAirwayBillType
`func (o *AirwayBillUpdateDto) UnsetAirwayBillType()`

UnsetAirwayBillType ensures that no value is present for AirwayBillType, not even an explicit nil
### GetMasterAwbNumber

`func (o *AirwayBillUpdateDto) GetMasterAwbNumber() string`

GetMasterAwbNumber returns the MasterAwbNumber field if non-nil, zero value otherwise.

### GetMasterAwbNumberOk

`func (o *AirwayBillUpdateDto) GetMasterAwbNumberOk() (*string, bool)`

GetMasterAwbNumberOk returns a tuple with the MasterAwbNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAwbNumber

`func (o *AirwayBillUpdateDto) SetMasterAwbNumber(v string)`

SetMasterAwbNumber sets MasterAwbNumber field to given value.

### HasMasterAwbNumber

`func (o *AirwayBillUpdateDto) HasMasterAwbNumber() bool`

HasMasterAwbNumber returns a boolean if a field has been set.

### SetMasterAwbNumberNil

`func (o *AirwayBillUpdateDto) SetMasterAwbNumberNil(b bool)`

 SetMasterAwbNumberNil sets the value for MasterAwbNumber to be an explicit nil

### UnsetMasterAwbNumber
`func (o *AirwayBillUpdateDto) UnsetMasterAwbNumber()`

UnsetMasterAwbNumber ensures that no value is present for MasterAwbNumber, not even an explicit nil
### GetShipperContactId

`func (o *AirwayBillUpdateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *AirwayBillUpdateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *AirwayBillUpdateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *AirwayBillUpdateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *AirwayBillUpdateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *AirwayBillUpdateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *AirwayBillUpdateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *AirwayBillUpdateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *AirwayBillUpdateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *AirwayBillUpdateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *AirwayBillUpdateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *AirwayBillUpdateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetNotifyPartyContactId

`func (o *AirwayBillUpdateDto) GetNotifyPartyContactId() string`

GetNotifyPartyContactId returns the NotifyPartyContactId field if non-nil, zero value otherwise.

### GetNotifyPartyContactIdOk

`func (o *AirwayBillUpdateDto) GetNotifyPartyContactIdOk() (*string, bool)`

GetNotifyPartyContactIdOk returns a tuple with the NotifyPartyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPartyContactId

`func (o *AirwayBillUpdateDto) SetNotifyPartyContactId(v string)`

SetNotifyPartyContactId sets NotifyPartyContactId field to given value.

### HasNotifyPartyContactId

`func (o *AirwayBillUpdateDto) HasNotifyPartyContactId() bool`

HasNotifyPartyContactId returns a boolean if a field has been set.

### SetNotifyPartyContactIdNil

`func (o *AirwayBillUpdateDto) SetNotifyPartyContactIdNil(b bool)`

 SetNotifyPartyContactIdNil sets the value for NotifyPartyContactId to be an explicit nil

### UnsetNotifyPartyContactId
`func (o *AirwayBillUpdateDto) UnsetNotifyPartyContactId()`

UnsetNotifyPartyContactId ensures that no value is present for NotifyPartyContactId, not even an explicit nil
### GetCarrierId

`func (o *AirwayBillUpdateDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *AirwayBillUpdateDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *AirwayBillUpdateDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *AirwayBillUpdateDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *AirwayBillUpdateDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *AirwayBillUpdateDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetAirlineCode

`func (o *AirwayBillUpdateDto) GetAirlineCode() string`

GetAirlineCode returns the AirlineCode field if non-nil, zero value otherwise.

### GetAirlineCodeOk

`func (o *AirwayBillUpdateDto) GetAirlineCodeOk() (*string, bool)`

GetAirlineCodeOk returns a tuple with the AirlineCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineCode

`func (o *AirwayBillUpdateDto) SetAirlineCode(v string)`

SetAirlineCode sets AirlineCode field to given value.

### HasAirlineCode

`func (o *AirwayBillUpdateDto) HasAirlineCode() bool`

HasAirlineCode returns a boolean if a field has been set.

### SetAirlineCodeNil

`func (o *AirwayBillUpdateDto) SetAirlineCodeNil(b bool)`

 SetAirlineCodeNil sets the value for AirlineCode to be an explicit nil

### UnsetAirlineCode
`func (o *AirwayBillUpdateDto) UnsetAirlineCode()`

UnsetAirlineCode ensures that no value is present for AirlineCode, not even an explicit nil
### GetFlightNumber

`func (o *AirwayBillUpdateDto) GetFlightNumber() string`

GetFlightNumber returns the FlightNumber field if non-nil, zero value otherwise.

### GetFlightNumberOk

`func (o *AirwayBillUpdateDto) GetFlightNumberOk() (*string, bool)`

GetFlightNumberOk returns a tuple with the FlightNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightNumber

`func (o *AirwayBillUpdateDto) SetFlightNumber(v string)`

SetFlightNumber sets FlightNumber field to given value.

### HasFlightNumber

`func (o *AirwayBillUpdateDto) HasFlightNumber() bool`

HasFlightNumber returns a boolean if a field has been set.

### SetFlightNumberNil

`func (o *AirwayBillUpdateDto) SetFlightNumberNil(b bool)`

 SetFlightNumberNil sets the value for FlightNumber to be an explicit nil

### UnsetFlightNumber
`func (o *AirwayBillUpdateDto) UnsetFlightNumber()`

UnsetFlightNumber ensures that no value is present for FlightNumber, not even an explicit nil
### GetAirportOfDepartureCode

`func (o *AirwayBillUpdateDto) GetAirportOfDepartureCode() string`

GetAirportOfDepartureCode returns the AirportOfDepartureCode field if non-nil, zero value otherwise.

### GetAirportOfDepartureCodeOk

`func (o *AirwayBillUpdateDto) GetAirportOfDepartureCodeOk() (*string, bool)`

GetAirportOfDepartureCodeOk returns a tuple with the AirportOfDepartureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirportOfDepartureCode

`func (o *AirwayBillUpdateDto) SetAirportOfDepartureCode(v string)`

SetAirportOfDepartureCode sets AirportOfDepartureCode field to given value.

### HasAirportOfDepartureCode

`func (o *AirwayBillUpdateDto) HasAirportOfDepartureCode() bool`

HasAirportOfDepartureCode returns a boolean if a field has been set.

### SetAirportOfDepartureCodeNil

`func (o *AirwayBillUpdateDto) SetAirportOfDepartureCodeNil(b bool)`

 SetAirportOfDepartureCodeNil sets the value for AirportOfDepartureCode to be an explicit nil

### UnsetAirportOfDepartureCode
`func (o *AirwayBillUpdateDto) UnsetAirportOfDepartureCode()`

UnsetAirportOfDepartureCode ensures that no value is present for AirportOfDepartureCode, not even an explicit nil
### GetAirportOfDestinationCode

`func (o *AirwayBillUpdateDto) GetAirportOfDestinationCode() string`

GetAirportOfDestinationCode returns the AirportOfDestinationCode field if non-nil, zero value otherwise.

### GetAirportOfDestinationCodeOk

`func (o *AirwayBillUpdateDto) GetAirportOfDestinationCodeOk() (*string, bool)`

GetAirportOfDestinationCodeOk returns a tuple with the AirportOfDestinationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirportOfDestinationCode

`func (o *AirwayBillUpdateDto) SetAirportOfDestinationCode(v string)`

SetAirportOfDestinationCode sets AirportOfDestinationCode field to given value.

### HasAirportOfDestinationCode

`func (o *AirwayBillUpdateDto) HasAirportOfDestinationCode() bool`

HasAirportOfDestinationCode returns a boolean if a field has been set.

### SetAirportOfDestinationCodeNil

`func (o *AirwayBillUpdateDto) SetAirportOfDestinationCodeNil(b bool)`

 SetAirportOfDestinationCodeNil sets the value for AirportOfDestinationCode to be an explicit nil

### UnsetAirportOfDestinationCode
`func (o *AirwayBillUpdateDto) UnsetAirportOfDestinationCode()`

UnsetAirportOfDestinationCode ensures that no value is present for AirportOfDestinationCode, not even an explicit nil
### GetDepartureDate

`func (o *AirwayBillUpdateDto) GetDepartureDate() time.Time`

GetDepartureDate returns the DepartureDate field if non-nil, zero value otherwise.

### GetDepartureDateOk

`func (o *AirwayBillUpdateDto) GetDepartureDateOk() (*time.Time, bool)`

GetDepartureDateOk returns a tuple with the DepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureDate

`func (o *AirwayBillUpdateDto) SetDepartureDate(v time.Time)`

SetDepartureDate sets DepartureDate field to given value.

### HasDepartureDate

`func (o *AirwayBillUpdateDto) HasDepartureDate() bool`

HasDepartureDate returns a boolean if a field has been set.

### SetDepartureDateNil

`func (o *AirwayBillUpdateDto) SetDepartureDateNil(b bool)`

 SetDepartureDateNil sets the value for DepartureDate to be an explicit nil

### UnsetDepartureDate
`func (o *AirwayBillUpdateDto) UnsetDepartureDate()`

UnsetDepartureDate ensures that no value is present for DepartureDate, not even an explicit nil
### GetArrivalDate

`func (o *AirwayBillUpdateDto) GetArrivalDate() time.Time`

GetArrivalDate returns the ArrivalDate field if non-nil, zero value otherwise.

### GetArrivalDateOk

`func (o *AirwayBillUpdateDto) GetArrivalDateOk() (*time.Time, bool)`

GetArrivalDateOk returns a tuple with the ArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalDate

`func (o *AirwayBillUpdateDto) SetArrivalDate(v time.Time)`

SetArrivalDate sets ArrivalDate field to given value.

### HasArrivalDate

`func (o *AirwayBillUpdateDto) HasArrivalDate() bool`

HasArrivalDate returns a boolean if a field has been set.

### SetArrivalDateNil

`func (o *AirwayBillUpdateDto) SetArrivalDateNil(b bool)`

 SetArrivalDateNil sets the value for ArrivalDate to be an explicit nil

### UnsetArrivalDate
`func (o *AirwayBillUpdateDto) UnsetArrivalDate()`

UnsetArrivalDate ensures that no value is present for ArrivalDate, not even an explicit nil
### GetDateIssued

`func (o *AirwayBillUpdateDto) GetDateIssued() time.Time`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *AirwayBillUpdateDto) GetDateIssuedOk() (*time.Time, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *AirwayBillUpdateDto) SetDateIssued(v time.Time)`

SetDateIssued sets DateIssued field to given value.

### HasDateIssued

`func (o *AirwayBillUpdateDto) HasDateIssued() bool`

HasDateIssued returns a boolean if a field has been set.

### SetDateIssuedNil

`func (o *AirwayBillUpdateDto) SetDateIssuedNil(b bool)`

 SetDateIssuedNil sets the value for DateIssued to be an explicit nil

### UnsetDateIssued
`func (o *AirwayBillUpdateDto) UnsetDateIssued()`

UnsetDateIssued ensures that no value is present for DateIssued, not even an explicit nil
### GetFreightTerms

`func (o *AirwayBillUpdateDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *AirwayBillUpdateDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *AirwayBillUpdateDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *AirwayBillUpdateDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *AirwayBillUpdateDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *AirwayBillUpdateDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *AirwayBillUpdateDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *AirwayBillUpdateDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *AirwayBillUpdateDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *AirwayBillUpdateDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *AirwayBillUpdateDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *AirwayBillUpdateDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *AirwayBillUpdateDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *AirwayBillUpdateDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *AirwayBillUpdateDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *AirwayBillUpdateDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *AirwayBillUpdateDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *AirwayBillUpdateDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetChargeableWeightKg

`func (o *AirwayBillUpdateDto) GetChargeableWeightKg() float64`

GetChargeableWeightKg returns the ChargeableWeightKg field if non-nil, zero value otherwise.

### GetChargeableWeightKgOk

`func (o *AirwayBillUpdateDto) GetChargeableWeightKgOk() (*float64, bool)`

GetChargeableWeightKgOk returns a tuple with the ChargeableWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeableWeightKg

`func (o *AirwayBillUpdateDto) SetChargeableWeightKg(v float64)`

SetChargeableWeightKg sets ChargeableWeightKg field to given value.

### HasChargeableWeightKg

`func (o *AirwayBillUpdateDto) HasChargeableWeightKg() bool`

HasChargeableWeightKg returns a boolean if a field has been set.

### SetChargeableWeightKgNil

`func (o *AirwayBillUpdateDto) SetChargeableWeightKgNil(b bool)`

 SetChargeableWeightKgNil sets the value for ChargeableWeightKg to be an explicit nil

### UnsetChargeableWeightKg
`func (o *AirwayBillUpdateDto) UnsetChargeableWeightKg()`

UnsetChargeableWeightKg ensures that no value is present for ChargeableWeightKg, not even an explicit nil
### GetTotalGrossWeightKg

`func (o *AirwayBillUpdateDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *AirwayBillUpdateDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *AirwayBillUpdateDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *AirwayBillUpdateDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### SetTotalGrossWeightKgNil

`func (o *AirwayBillUpdateDto) SetTotalGrossWeightKgNil(b bool)`

 SetTotalGrossWeightKgNil sets the value for TotalGrossWeightKg to be an explicit nil

### UnsetTotalGrossWeightKg
`func (o *AirwayBillUpdateDto) UnsetTotalGrossWeightKg()`

UnsetTotalGrossWeightKg ensures that no value is present for TotalGrossWeightKg, not even an explicit nil
### GetTotalPackages

`func (o *AirwayBillUpdateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *AirwayBillUpdateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *AirwayBillUpdateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *AirwayBillUpdateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *AirwayBillUpdateDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *AirwayBillUpdateDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetTotalVolumeM3

`func (o *AirwayBillUpdateDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *AirwayBillUpdateDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *AirwayBillUpdateDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *AirwayBillUpdateDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *AirwayBillUpdateDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *AirwayBillUpdateDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetDeclaredValueForCarriage

`func (o *AirwayBillUpdateDto) GetDeclaredValueForCarriage() float64`

GetDeclaredValueForCarriage returns the DeclaredValueForCarriage field if non-nil, zero value otherwise.

### GetDeclaredValueForCarriageOk

`func (o *AirwayBillUpdateDto) GetDeclaredValueForCarriageOk() (*float64, bool)`

GetDeclaredValueForCarriageOk returns a tuple with the DeclaredValueForCarriage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueForCarriage

`func (o *AirwayBillUpdateDto) SetDeclaredValueForCarriage(v float64)`

SetDeclaredValueForCarriage sets DeclaredValueForCarriage field to given value.

### HasDeclaredValueForCarriage

`func (o *AirwayBillUpdateDto) HasDeclaredValueForCarriage() bool`

HasDeclaredValueForCarriage returns a boolean if a field has been set.

### SetDeclaredValueForCarriageNil

`func (o *AirwayBillUpdateDto) SetDeclaredValueForCarriageNil(b bool)`

 SetDeclaredValueForCarriageNil sets the value for DeclaredValueForCarriage to be an explicit nil

### UnsetDeclaredValueForCarriage
`func (o *AirwayBillUpdateDto) UnsetDeclaredValueForCarriage()`

UnsetDeclaredValueForCarriage ensures that no value is present for DeclaredValueForCarriage, not even an explicit nil
### GetDeclaredValueForCustoms

`func (o *AirwayBillUpdateDto) GetDeclaredValueForCustoms() float64`

GetDeclaredValueForCustoms returns the DeclaredValueForCustoms field if non-nil, zero value otherwise.

### GetDeclaredValueForCustomsOk

`func (o *AirwayBillUpdateDto) GetDeclaredValueForCustomsOk() (*float64, bool)`

GetDeclaredValueForCustomsOk returns a tuple with the DeclaredValueForCustoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueForCustoms

`func (o *AirwayBillUpdateDto) SetDeclaredValueForCustoms(v float64)`

SetDeclaredValueForCustoms sets DeclaredValueForCustoms field to given value.

### HasDeclaredValueForCustoms

`func (o *AirwayBillUpdateDto) HasDeclaredValueForCustoms() bool`

HasDeclaredValueForCustoms returns a boolean if a field has been set.

### SetDeclaredValueForCustomsNil

`func (o *AirwayBillUpdateDto) SetDeclaredValueForCustomsNil(b bool)`

 SetDeclaredValueForCustomsNil sets the value for DeclaredValueForCustoms to be an explicit nil

### UnsetDeclaredValueForCustoms
`func (o *AirwayBillUpdateDto) UnsetDeclaredValueForCustoms()`

UnsetDeclaredValueForCustoms ensures that no value is present for DeclaredValueForCustoms, not even an explicit nil
### GetInsuranceAmount

`func (o *AirwayBillUpdateDto) GetInsuranceAmount() float64`

GetInsuranceAmount returns the InsuranceAmount field if non-nil, zero value otherwise.

### GetInsuranceAmountOk

`func (o *AirwayBillUpdateDto) GetInsuranceAmountOk() (*float64, bool)`

GetInsuranceAmountOk returns a tuple with the InsuranceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceAmount

`func (o *AirwayBillUpdateDto) SetInsuranceAmount(v float64)`

SetInsuranceAmount sets InsuranceAmount field to given value.

### HasInsuranceAmount

`func (o *AirwayBillUpdateDto) HasInsuranceAmount() bool`

HasInsuranceAmount returns a boolean if a field has been set.

### SetInsuranceAmountNil

`func (o *AirwayBillUpdateDto) SetInsuranceAmountNil(b bool)`

 SetInsuranceAmountNil sets the value for InsuranceAmount to be an explicit nil

### UnsetInsuranceAmount
`func (o *AirwayBillUpdateDto) UnsetInsuranceAmount()`

UnsetInsuranceAmount ensures that no value is present for InsuranceAmount, not even an explicit nil
### GetSpecialHandlingCodes

`func (o *AirwayBillUpdateDto) GetSpecialHandlingCodes() string`

GetSpecialHandlingCodes returns the SpecialHandlingCodes field if non-nil, zero value otherwise.

### GetSpecialHandlingCodesOk

`func (o *AirwayBillUpdateDto) GetSpecialHandlingCodesOk() (*string, bool)`

GetSpecialHandlingCodesOk returns a tuple with the SpecialHandlingCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialHandlingCodes

`func (o *AirwayBillUpdateDto) SetSpecialHandlingCodes(v string)`

SetSpecialHandlingCodes sets SpecialHandlingCodes field to given value.

### HasSpecialHandlingCodes

`func (o *AirwayBillUpdateDto) HasSpecialHandlingCodes() bool`

HasSpecialHandlingCodes returns a boolean if a field has been set.

### SetSpecialHandlingCodesNil

`func (o *AirwayBillUpdateDto) SetSpecialHandlingCodesNil(b bool)`

 SetSpecialHandlingCodesNil sets the value for SpecialHandlingCodes to be an explicit nil

### UnsetSpecialHandlingCodes
`func (o *AirwayBillUpdateDto) UnsetSpecialHandlingCodes()`

UnsetSpecialHandlingCodes ensures that no value is present for SpecialHandlingCodes, not even an explicit nil
### GetSpecialInstructions

`func (o *AirwayBillUpdateDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *AirwayBillUpdateDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *AirwayBillUpdateDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *AirwayBillUpdateDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *AirwayBillUpdateDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *AirwayBillUpdateDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *AirwayBillUpdateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *AirwayBillUpdateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *AirwayBillUpdateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *AirwayBillUpdateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *AirwayBillUpdateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *AirwayBillUpdateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *AirwayBillUpdateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *AirwayBillUpdateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *AirwayBillUpdateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *AirwayBillUpdateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *AirwayBillUpdateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *AirwayBillUpdateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


