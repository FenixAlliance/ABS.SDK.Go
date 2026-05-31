# AirwayBillDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**AirwayBillType** | Pointer to **NullableString** |  | [optional] 
**MasterAwbNumber** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
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
**DateDelivered** | Pointer to **NullableTime** |  | [optional] 
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
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Lines** | Pointer to [**[]WaybillLineDto**](WaybillLineDto.md) |  | [optional] 

## Methods

### NewAirwayBillDto

`func NewAirwayBillDto() *AirwayBillDto`

NewAirwayBillDto instantiates a new AirwayBillDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirwayBillDtoWithDefaults

`func NewAirwayBillDtoWithDefaults() *AirwayBillDto`

NewAirwayBillDtoWithDefaults instantiates a new AirwayBillDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AirwayBillDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AirwayBillDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AirwayBillDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AirwayBillDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AirwayBillDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AirwayBillDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AirwayBillDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AirwayBillDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AirwayBillDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AirwayBillDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AirwayBillDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AirwayBillDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDocumentNumber

`func (o *AirwayBillDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *AirwayBillDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *AirwayBillDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *AirwayBillDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *AirwayBillDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *AirwayBillDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetAirwayBillType

`func (o *AirwayBillDto) GetAirwayBillType() string`

GetAirwayBillType returns the AirwayBillType field if non-nil, zero value otherwise.

### GetAirwayBillTypeOk

`func (o *AirwayBillDto) GetAirwayBillTypeOk() (*string, bool)`

GetAirwayBillTypeOk returns a tuple with the AirwayBillType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirwayBillType

`func (o *AirwayBillDto) SetAirwayBillType(v string)`

SetAirwayBillType sets AirwayBillType field to given value.

### HasAirwayBillType

`func (o *AirwayBillDto) HasAirwayBillType() bool`

HasAirwayBillType returns a boolean if a field has been set.

### SetAirwayBillTypeNil

`func (o *AirwayBillDto) SetAirwayBillTypeNil(b bool)`

 SetAirwayBillTypeNil sets the value for AirwayBillType to be an explicit nil

### UnsetAirwayBillType
`func (o *AirwayBillDto) UnsetAirwayBillType()`

UnsetAirwayBillType ensures that no value is present for AirwayBillType, not even an explicit nil
### GetMasterAwbNumber

`func (o *AirwayBillDto) GetMasterAwbNumber() string`

GetMasterAwbNumber returns the MasterAwbNumber field if non-nil, zero value otherwise.

### GetMasterAwbNumberOk

`func (o *AirwayBillDto) GetMasterAwbNumberOk() (*string, bool)`

GetMasterAwbNumberOk returns a tuple with the MasterAwbNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAwbNumber

`func (o *AirwayBillDto) SetMasterAwbNumber(v string)`

SetMasterAwbNumber sets MasterAwbNumber field to given value.

### HasMasterAwbNumber

`func (o *AirwayBillDto) HasMasterAwbNumber() bool`

HasMasterAwbNumber returns a boolean if a field has been set.

### SetMasterAwbNumberNil

`func (o *AirwayBillDto) SetMasterAwbNumberNil(b bool)`

 SetMasterAwbNumberNil sets the value for MasterAwbNumber to be an explicit nil

### UnsetMasterAwbNumber
`func (o *AirwayBillDto) UnsetMasterAwbNumber()`

UnsetMasterAwbNumber ensures that no value is present for MasterAwbNumber, not even an explicit nil
### GetStatus

`func (o *AirwayBillDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AirwayBillDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AirwayBillDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AirwayBillDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AirwayBillDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AirwayBillDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetShipperContactId

`func (o *AirwayBillDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *AirwayBillDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *AirwayBillDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *AirwayBillDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *AirwayBillDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *AirwayBillDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *AirwayBillDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *AirwayBillDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *AirwayBillDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *AirwayBillDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *AirwayBillDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *AirwayBillDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetNotifyPartyContactId

`func (o *AirwayBillDto) GetNotifyPartyContactId() string`

GetNotifyPartyContactId returns the NotifyPartyContactId field if non-nil, zero value otherwise.

### GetNotifyPartyContactIdOk

`func (o *AirwayBillDto) GetNotifyPartyContactIdOk() (*string, bool)`

GetNotifyPartyContactIdOk returns a tuple with the NotifyPartyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPartyContactId

`func (o *AirwayBillDto) SetNotifyPartyContactId(v string)`

SetNotifyPartyContactId sets NotifyPartyContactId field to given value.

### HasNotifyPartyContactId

`func (o *AirwayBillDto) HasNotifyPartyContactId() bool`

HasNotifyPartyContactId returns a boolean if a field has been set.

### SetNotifyPartyContactIdNil

`func (o *AirwayBillDto) SetNotifyPartyContactIdNil(b bool)`

 SetNotifyPartyContactIdNil sets the value for NotifyPartyContactId to be an explicit nil

### UnsetNotifyPartyContactId
`func (o *AirwayBillDto) UnsetNotifyPartyContactId()`

UnsetNotifyPartyContactId ensures that no value is present for NotifyPartyContactId, not even an explicit nil
### GetCarrierId

`func (o *AirwayBillDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *AirwayBillDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *AirwayBillDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *AirwayBillDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *AirwayBillDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *AirwayBillDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetAirlineCode

`func (o *AirwayBillDto) GetAirlineCode() string`

GetAirlineCode returns the AirlineCode field if non-nil, zero value otherwise.

### GetAirlineCodeOk

`func (o *AirwayBillDto) GetAirlineCodeOk() (*string, bool)`

GetAirlineCodeOk returns a tuple with the AirlineCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineCode

`func (o *AirwayBillDto) SetAirlineCode(v string)`

SetAirlineCode sets AirlineCode field to given value.

### HasAirlineCode

`func (o *AirwayBillDto) HasAirlineCode() bool`

HasAirlineCode returns a boolean if a field has been set.

### SetAirlineCodeNil

`func (o *AirwayBillDto) SetAirlineCodeNil(b bool)`

 SetAirlineCodeNil sets the value for AirlineCode to be an explicit nil

### UnsetAirlineCode
`func (o *AirwayBillDto) UnsetAirlineCode()`

UnsetAirlineCode ensures that no value is present for AirlineCode, not even an explicit nil
### GetFlightNumber

`func (o *AirwayBillDto) GetFlightNumber() string`

GetFlightNumber returns the FlightNumber field if non-nil, zero value otherwise.

### GetFlightNumberOk

`func (o *AirwayBillDto) GetFlightNumberOk() (*string, bool)`

GetFlightNumberOk returns a tuple with the FlightNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightNumber

`func (o *AirwayBillDto) SetFlightNumber(v string)`

SetFlightNumber sets FlightNumber field to given value.

### HasFlightNumber

`func (o *AirwayBillDto) HasFlightNumber() bool`

HasFlightNumber returns a boolean if a field has been set.

### SetFlightNumberNil

`func (o *AirwayBillDto) SetFlightNumberNil(b bool)`

 SetFlightNumberNil sets the value for FlightNumber to be an explicit nil

### UnsetFlightNumber
`func (o *AirwayBillDto) UnsetFlightNumber()`

UnsetFlightNumber ensures that no value is present for FlightNumber, not even an explicit nil
### GetAirportOfDepartureCode

`func (o *AirwayBillDto) GetAirportOfDepartureCode() string`

GetAirportOfDepartureCode returns the AirportOfDepartureCode field if non-nil, zero value otherwise.

### GetAirportOfDepartureCodeOk

`func (o *AirwayBillDto) GetAirportOfDepartureCodeOk() (*string, bool)`

GetAirportOfDepartureCodeOk returns a tuple with the AirportOfDepartureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirportOfDepartureCode

`func (o *AirwayBillDto) SetAirportOfDepartureCode(v string)`

SetAirportOfDepartureCode sets AirportOfDepartureCode field to given value.

### HasAirportOfDepartureCode

`func (o *AirwayBillDto) HasAirportOfDepartureCode() bool`

HasAirportOfDepartureCode returns a boolean if a field has been set.

### SetAirportOfDepartureCodeNil

`func (o *AirwayBillDto) SetAirportOfDepartureCodeNil(b bool)`

 SetAirportOfDepartureCodeNil sets the value for AirportOfDepartureCode to be an explicit nil

### UnsetAirportOfDepartureCode
`func (o *AirwayBillDto) UnsetAirportOfDepartureCode()`

UnsetAirportOfDepartureCode ensures that no value is present for AirportOfDepartureCode, not even an explicit nil
### GetAirportOfDestinationCode

`func (o *AirwayBillDto) GetAirportOfDestinationCode() string`

GetAirportOfDestinationCode returns the AirportOfDestinationCode field if non-nil, zero value otherwise.

### GetAirportOfDestinationCodeOk

`func (o *AirwayBillDto) GetAirportOfDestinationCodeOk() (*string, bool)`

GetAirportOfDestinationCodeOk returns a tuple with the AirportOfDestinationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirportOfDestinationCode

`func (o *AirwayBillDto) SetAirportOfDestinationCode(v string)`

SetAirportOfDestinationCode sets AirportOfDestinationCode field to given value.

### HasAirportOfDestinationCode

`func (o *AirwayBillDto) HasAirportOfDestinationCode() bool`

HasAirportOfDestinationCode returns a boolean if a field has been set.

### SetAirportOfDestinationCodeNil

`func (o *AirwayBillDto) SetAirportOfDestinationCodeNil(b bool)`

 SetAirportOfDestinationCodeNil sets the value for AirportOfDestinationCode to be an explicit nil

### UnsetAirportOfDestinationCode
`func (o *AirwayBillDto) UnsetAirportOfDestinationCode()`

UnsetAirportOfDestinationCode ensures that no value is present for AirportOfDestinationCode, not even an explicit nil
### GetDepartureDate

`func (o *AirwayBillDto) GetDepartureDate() time.Time`

GetDepartureDate returns the DepartureDate field if non-nil, zero value otherwise.

### GetDepartureDateOk

`func (o *AirwayBillDto) GetDepartureDateOk() (*time.Time, bool)`

GetDepartureDateOk returns a tuple with the DepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureDate

`func (o *AirwayBillDto) SetDepartureDate(v time.Time)`

SetDepartureDate sets DepartureDate field to given value.

### HasDepartureDate

`func (o *AirwayBillDto) HasDepartureDate() bool`

HasDepartureDate returns a boolean if a field has been set.

### SetDepartureDateNil

`func (o *AirwayBillDto) SetDepartureDateNil(b bool)`

 SetDepartureDateNil sets the value for DepartureDate to be an explicit nil

### UnsetDepartureDate
`func (o *AirwayBillDto) UnsetDepartureDate()`

UnsetDepartureDate ensures that no value is present for DepartureDate, not even an explicit nil
### GetArrivalDate

`func (o *AirwayBillDto) GetArrivalDate() time.Time`

GetArrivalDate returns the ArrivalDate field if non-nil, zero value otherwise.

### GetArrivalDateOk

`func (o *AirwayBillDto) GetArrivalDateOk() (*time.Time, bool)`

GetArrivalDateOk returns a tuple with the ArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalDate

`func (o *AirwayBillDto) SetArrivalDate(v time.Time)`

SetArrivalDate sets ArrivalDate field to given value.

### HasArrivalDate

`func (o *AirwayBillDto) HasArrivalDate() bool`

HasArrivalDate returns a boolean if a field has been set.

### SetArrivalDateNil

`func (o *AirwayBillDto) SetArrivalDateNil(b bool)`

 SetArrivalDateNil sets the value for ArrivalDate to be an explicit nil

### UnsetArrivalDate
`func (o *AirwayBillDto) UnsetArrivalDate()`

UnsetArrivalDate ensures that no value is present for ArrivalDate, not even an explicit nil
### GetDateIssued

`func (o *AirwayBillDto) GetDateIssued() time.Time`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *AirwayBillDto) GetDateIssuedOk() (*time.Time, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *AirwayBillDto) SetDateIssued(v time.Time)`

SetDateIssued sets DateIssued field to given value.

### HasDateIssued

`func (o *AirwayBillDto) HasDateIssued() bool`

HasDateIssued returns a boolean if a field has been set.

### SetDateIssuedNil

`func (o *AirwayBillDto) SetDateIssuedNil(b bool)`

 SetDateIssuedNil sets the value for DateIssued to be an explicit nil

### UnsetDateIssued
`func (o *AirwayBillDto) UnsetDateIssued()`

UnsetDateIssued ensures that no value is present for DateIssued, not even an explicit nil
### GetDateDelivered

`func (o *AirwayBillDto) GetDateDelivered() time.Time`

GetDateDelivered returns the DateDelivered field if non-nil, zero value otherwise.

### GetDateDeliveredOk

`func (o *AirwayBillDto) GetDateDeliveredOk() (*time.Time, bool)`

GetDateDeliveredOk returns a tuple with the DateDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDelivered

`func (o *AirwayBillDto) SetDateDelivered(v time.Time)`

SetDateDelivered sets DateDelivered field to given value.

### HasDateDelivered

`func (o *AirwayBillDto) HasDateDelivered() bool`

HasDateDelivered returns a boolean if a field has been set.

### SetDateDeliveredNil

`func (o *AirwayBillDto) SetDateDeliveredNil(b bool)`

 SetDateDeliveredNil sets the value for DateDelivered to be an explicit nil

### UnsetDateDelivered
`func (o *AirwayBillDto) UnsetDateDelivered()`

UnsetDateDelivered ensures that no value is present for DateDelivered, not even an explicit nil
### GetFreightTerms

`func (o *AirwayBillDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *AirwayBillDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *AirwayBillDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *AirwayBillDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *AirwayBillDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *AirwayBillDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *AirwayBillDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *AirwayBillDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *AirwayBillDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *AirwayBillDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *AirwayBillDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *AirwayBillDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *AirwayBillDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *AirwayBillDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *AirwayBillDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *AirwayBillDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *AirwayBillDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *AirwayBillDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetChargeableWeightKg

`func (o *AirwayBillDto) GetChargeableWeightKg() float64`

GetChargeableWeightKg returns the ChargeableWeightKg field if non-nil, zero value otherwise.

### GetChargeableWeightKgOk

`func (o *AirwayBillDto) GetChargeableWeightKgOk() (*float64, bool)`

GetChargeableWeightKgOk returns a tuple with the ChargeableWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeableWeightKg

`func (o *AirwayBillDto) SetChargeableWeightKg(v float64)`

SetChargeableWeightKg sets ChargeableWeightKg field to given value.

### HasChargeableWeightKg

`func (o *AirwayBillDto) HasChargeableWeightKg() bool`

HasChargeableWeightKg returns a boolean if a field has been set.

### SetChargeableWeightKgNil

`func (o *AirwayBillDto) SetChargeableWeightKgNil(b bool)`

 SetChargeableWeightKgNil sets the value for ChargeableWeightKg to be an explicit nil

### UnsetChargeableWeightKg
`func (o *AirwayBillDto) UnsetChargeableWeightKg()`

UnsetChargeableWeightKg ensures that no value is present for ChargeableWeightKg, not even an explicit nil
### GetTotalGrossWeightKg

`func (o *AirwayBillDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *AirwayBillDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *AirwayBillDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *AirwayBillDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### SetTotalGrossWeightKgNil

`func (o *AirwayBillDto) SetTotalGrossWeightKgNil(b bool)`

 SetTotalGrossWeightKgNil sets the value for TotalGrossWeightKg to be an explicit nil

### UnsetTotalGrossWeightKg
`func (o *AirwayBillDto) UnsetTotalGrossWeightKg()`

UnsetTotalGrossWeightKg ensures that no value is present for TotalGrossWeightKg, not even an explicit nil
### GetTotalPackages

`func (o *AirwayBillDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *AirwayBillDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *AirwayBillDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *AirwayBillDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *AirwayBillDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *AirwayBillDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetTotalVolumeM3

`func (o *AirwayBillDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *AirwayBillDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *AirwayBillDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *AirwayBillDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *AirwayBillDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *AirwayBillDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetDeclaredValueForCarriage

`func (o *AirwayBillDto) GetDeclaredValueForCarriage() float64`

GetDeclaredValueForCarriage returns the DeclaredValueForCarriage field if non-nil, zero value otherwise.

### GetDeclaredValueForCarriageOk

`func (o *AirwayBillDto) GetDeclaredValueForCarriageOk() (*float64, bool)`

GetDeclaredValueForCarriageOk returns a tuple with the DeclaredValueForCarriage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueForCarriage

`func (o *AirwayBillDto) SetDeclaredValueForCarriage(v float64)`

SetDeclaredValueForCarriage sets DeclaredValueForCarriage field to given value.

### HasDeclaredValueForCarriage

`func (o *AirwayBillDto) HasDeclaredValueForCarriage() bool`

HasDeclaredValueForCarriage returns a boolean if a field has been set.

### SetDeclaredValueForCarriageNil

`func (o *AirwayBillDto) SetDeclaredValueForCarriageNil(b bool)`

 SetDeclaredValueForCarriageNil sets the value for DeclaredValueForCarriage to be an explicit nil

### UnsetDeclaredValueForCarriage
`func (o *AirwayBillDto) UnsetDeclaredValueForCarriage()`

UnsetDeclaredValueForCarriage ensures that no value is present for DeclaredValueForCarriage, not even an explicit nil
### GetDeclaredValueForCustoms

`func (o *AirwayBillDto) GetDeclaredValueForCustoms() float64`

GetDeclaredValueForCustoms returns the DeclaredValueForCustoms field if non-nil, zero value otherwise.

### GetDeclaredValueForCustomsOk

`func (o *AirwayBillDto) GetDeclaredValueForCustomsOk() (*float64, bool)`

GetDeclaredValueForCustomsOk returns a tuple with the DeclaredValueForCustoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueForCustoms

`func (o *AirwayBillDto) SetDeclaredValueForCustoms(v float64)`

SetDeclaredValueForCustoms sets DeclaredValueForCustoms field to given value.

### HasDeclaredValueForCustoms

`func (o *AirwayBillDto) HasDeclaredValueForCustoms() bool`

HasDeclaredValueForCustoms returns a boolean if a field has been set.

### SetDeclaredValueForCustomsNil

`func (o *AirwayBillDto) SetDeclaredValueForCustomsNil(b bool)`

 SetDeclaredValueForCustomsNil sets the value for DeclaredValueForCustoms to be an explicit nil

### UnsetDeclaredValueForCustoms
`func (o *AirwayBillDto) UnsetDeclaredValueForCustoms()`

UnsetDeclaredValueForCustoms ensures that no value is present for DeclaredValueForCustoms, not even an explicit nil
### GetInsuranceAmount

`func (o *AirwayBillDto) GetInsuranceAmount() float64`

GetInsuranceAmount returns the InsuranceAmount field if non-nil, zero value otherwise.

### GetInsuranceAmountOk

`func (o *AirwayBillDto) GetInsuranceAmountOk() (*float64, bool)`

GetInsuranceAmountOk returns a tuple with the InsuranceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuranceAmount

`func (o *AirwayBillDto) SetInsuranceAmount(v float64)`

SetInsuranceAmount sets InsuranceAmount field to given value.

### HasInsuranceAmount

`func (o *AirwayBillDto) HasInsuranceAmount() bool`

HasInsuranceAmount returns a boolean if a field has been set.

### SetInsuranceAmountNil

`func (o *AirwayBillDto) SetInsuranceAmountNil(b bool)`

 SetInsuranceAmountNil sets the value for InsuranceAmount to be an explicit nil

### UnsetInsuranceAmount
`func (o *AirwayBillDto) UnsetInsuranceAmount()`

UnsetInsuranceAmount ensures that no value is present for InsuranceAmount, not even an explicit nil
### GetSpecialHandlingCodes

`func (o *AirwayBillDto) GetSpecialHandlingCodes() string`

GetSpecialHandlingCodes returns the SpecialHandlingCodes field if non-nil, zero value otherwise.

### GetSpecialHandlingCodesOk

`func (o *AirwayBillDto) GetSpecialHandlingCodesOk() (*string, bool)`

GetSpecialHandlingCodesOk returns a tuple with the SpecialHandlingCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialHandlingCodes

`func (o *AirwayBillDto) SetSpecialHandlingCodes(v string)`

SetSpecialHandlingCodes sets SpecialHandlingCodes field to given value.

### HasSpecialHandlingCodes

`func (o *AirwayBillDto) HasSpecialHandlingCodes() bool`

HasSpecialHandlingCodes returns a boolean if a field has been set.

### SetSpecialHandlingCodesNil

`func (o *AirwayBillDto) SetSpecialHandlingCodesNil(b bool)`

 SetSpecialHandlingCodesNil sets the value for SpecialHandlingCodes to be an explicit nil

### UnsetSpecialHandlingCodes
`func (o *AirwayBillDto) UnsetSpecialHandlingCodes()`

UnsetSpecialHandlingCodes ensures that no value is present for SpecialHandlingCodes, not even an explicit nil
### GetSpecialInstructions

`func (o *AirwayBillDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *AirwayBillDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *AirwayBillDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *AirwayBillDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *AirwayBillDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *AirwayBillDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *AirwayBillDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *AirwayBillDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *AirwayBillDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *AirwayBillDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *AirwayBillDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *AirwayBillDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *AirwayBillDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *AirwayBillDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *AirwayBillDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *AirwayBillDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *AirwayBillDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *AirwayBillDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetTenantId

`func (o *AirwayBillDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AirwayBillDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AirwayBillDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AirwayBillDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AirwayBillDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AirwayBillDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AirwayBillDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AirwayBillDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AirwayBillDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AirwayBillDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AirwayBillDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AirwayBillDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetLines

`func (o *AirwayBillDto) GetLines() []WaybillLineDto`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *AirwayBillDto) GetLinesOk() (*[]WaybillLineDto, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *AirwayBillDto) SetLines(v []WaybillLineDto)`

SetLines sets Lines field to given value.

### HasLines

`func (o *AirwayBillDto) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *AirwayBillDto) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *AirwayBillDto) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


