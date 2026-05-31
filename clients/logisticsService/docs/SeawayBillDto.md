# SeawayBillDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**ShipperContactId** | Pointer to **NullableString** |  | [optional] 
**ConsigneeContactId** | Pointer to **NullableString** |  | [optional] 
**NotifyPartyContactId** | Pointer to **NullableString** |  | [optional] 
**CarrierId** | Pointer to **NullableString** |  | [optional] 
**VesselId** | Pointer to **NullableString** |  | [optional] 
**VoyageId** | Pointer to **NullableString** |  | [optional] 
**PortOfLoadingId** | Pointer to **NullableString** |  | [optional] 
**PortOfDischargeId** | Pointer to **NullableString** |  | [optional] 
**PlaceOfReceipt** | Pointer to **NullableString** |  | [optional] 
**PlaceOfDelivery** | Pointer to **NullableString** |  | [optional] 
**DateIssued** | Pointer to **NullableTime** |  | [optional] 
**DateShipped** | Pointer to **NullableTime** |  | [optional] 
**DateDelivered** | Pointer to **NullableTime** |  | [optional] 
**FreightTerms** | Pointer to **NullableString** |  | [optional] 
**FreightAmount** | Pointer to **NullableFloat64** |  | [optional] 
**FreightCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalWeight** | Pointer to **NullableFloat64** |  | [optional] 
**TotalPackages** | Pointer to **NullableInt32** |  | [optional] 
**SpecialInstructions** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Lines** | Pointer to [**[]WaybillLineDto**](WaybillLineDto.md) |  | [optional] 

## Methods

### NewSeawayBillDto

`func NewSeawayBillDto() *SeawayBillDto`

NewSeawayBillDto instantiates a new SeawayBillDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeawayBillDtoWithDefaults

`func NewSeawayBillDtoWithDefaults() *SeawayBillDto`

NewSeawayBillDtoWithDefaults instantiates a new SeawayBillDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SeawayBillDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SeawayBillDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SeawayBillDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SeawayBillDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SeawayBillDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SeawayBillDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *SeawayBillDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SeawayBillDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SeawayBillDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SeawayBillDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *SeawayBillDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *SeawayBillDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDocumentNumber

`func (o *SeawayBillDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *SeawayBillDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *SeawayBillDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *SeawayBillDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *SeawayBillDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *SeawayBillDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetStatus

`func (o *SeawayBillDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SeawayBillDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SeawayBillDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SeawayBillDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SeawayBillDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SeawayBillDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetShipperContactId

`func (o *SeawayBillDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *SeawayBillDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *SeawayBillDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *SeawayBillDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *SeawayBillDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *SeawayBillDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *SeawayBillDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *SeawayBillDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *SeawayBillDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *SeawayBillDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *SeawayBillDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *SeawayBillDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetNotifyPartyContactId

`func (o *SeawayBillDto) GetNotifyPartyContactId() string`

GetNotifyPartyContactId returns the NotifyPartyContactId field if non-nil, zero value otherwise.

### GetNotifyPartyContactIdOk

`func (o *SeawayBillDto) GetNotifyPartyContactIdOk() (*string, bool)`

GetNotifyPartyContactIdOk returns a tuple with the NotifyPartyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPartyContactId

`func (o *SeawayBillDto) SetNotifyPartyContactId(v string)`

SetNotifyPartyContactId sets NotifyPartyContactId field to given value.

### HasNotifyPartyContactId

`func (o *SeawayBillDto) HasNotifyPartyContactId() bool`

HasNotifyPartyContactId returns a boolean if a field has been set.

### SetNotifyPartyContactIdNil

`func (o *SeawayBillDto) SetNotifyPartyContactIdNil(b bool)`

 SetNotifyPartyContactIdNil sets the value for NotifyPartyContactId to be an explicit nil

### UnsetNotifyPartyContactId
`func (o *SeawayBillDto) UnsetNotifyPartyContactId()`

UnsetNotifyPartyContactId ensures that no value is present for NotifyPartyContactId, not even an explicit nil
### GetCarrierId

`func (o *SeawayBillDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *SeawayBillDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *SeawayBillDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *SeawayBillDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *SeawayBillDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *SeawayBillDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetVesselId

`func (o *SeawayBillDto) GetVesselId() string`

GetVesselId returns the VesselId field if non-nil, zero value otherwise.

### GetVesselIdOk

`func (o *SeawayBillDto) GetVesselIdOk() (*string, bool)`

GetVesselIdOk returns a tuple with the VesselId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselId

`func (o *SeawayBillDto) SetVesselId(v string)`

SetVesselId sets VesselId field to given value.

### HasVesselId

`func (o *SeawayBillDto) HasVesselId() bool`

HasVesselId returns a boolean if a field has been set.

### SetVesselIdNil

`func (o *SeawayBillDto) SetVesselIdNil(b bool)`

 SetVesselIdNil sets the value for VesselId to be an explicit nil

### UnsetVesselId
`func (o *SeawayBillDto) UnsetVesselId()`

UnsetVesselId ensures that no value is present for VesselId, not even an explicit nil
### GetVoyageId

`func (o *SeawayBillDto) GetVoyageId() string`

GetVoyageId returns the VoyageId field if non-nil, zero value otherwise.

### GetVoyageIdOk

`func (o *SeawayBillDto) GetVoyageIdOk() (*string, bool)`

GetVoyageIdOk returns a tuple with the VoyageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageId

`func (o *SeawayBillDto) SetVoyageId(v string)`

SetVoyageId sets VoyageId field to given value.

### HasVoyageId

`func (o *SeawayBillDto) HasVoyageId() bool`

HasVoyageId returns a boolean if a field has been set.

### SetVoyageIdNil

`func (o *SeawayBillDto) SetVoyageIdNil(b bool)`

 SetVoyageIdNil sets the value for VoyageId to be an explicit nil

### UnsetVoyageId
`func (o *SeawayBillDto) UnsetVoyageId()`

UnsetVoyageId ensures that no value is present for VoyageId, not even an explicit nil
### GetPortOfLoadingId

`func (o *SeawayBillDto) GetPortOfLoadingId() string`

GetPortOfLoadingId returns the PortOfLoadingId field if non-nil, zero value otherwise.

### GetPortOfLoadingIdOk

`func (o *SeawayBillDto) GetPortOfLoadingIdOk() (*string, bool)`

GetPortOfLoadingIdOk returns a tuple with the PortOfLoadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoadingId

`func (o *SeawayBillDto) SetPortOfLoadingId(v string)`

SetPortOfLoadingId sets PortOfLoadingId field to given value.

### HasPortOfLoadingId

`func (o *SeawayBillDto) HasPortOfLoadingId() bool`

HasPortOfLoadingId returns a boolean if a field has been set.

### SetPortOfLoadingIdNil

`func (o *SeawayBillDto) SetPortOfLoadingIdNil(b bool)`

 SetPortOfLoadingIdNil sets the value for PortOfLoadingId to be an explicit nil

### UnsetPortOfLoadingId
`func (o *SeawayBillDto) UnsetPortOfLoadingId()`

UnsetPortOfLoadingId ensures that no value is present for PortOfLoadingId, not even an explicit nil
### GetPortOfDischargeId

`func (o *SeawayBillDto) GetPortOfDischargeId() string`

GetPortOfDischargeId returns the PortOfDischargeId field if non-nil, zero value otherwise.

### GetPortOfDischargeIdOk

`func (o *SeawayBillDto) GetPortOfDischargeIdOk() (*string, bool)`

GetPortOfDischargeIdOk returns a tuple with the PortOfDischargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfDischargeId

`func (o *SeawayBillDto) SetPortOfDischargeId(v string)`

SetPortOfDischargeId sets PortOfDischargeId field to given value.

### HasPortOfDischargeId

`func (o *SeawayBillDto) HasPortOfDischargeId() bool`

HasPortOfDischargeId returns a boolean if a field has been set.

### SetPortOfDischargeIdNil

`func (o *SeawayBillDto) SetPortOfDischargeIdNil(b bool)`

 SetPortOfDischargeIdNil sets the value for PortOfDischargeId to be an explicit nil

### UnsetPortOfDischargeId
`func (o *SeawayBillDto) UnsetPortOfDischargeId()`

UnsetPortOfDischargeId ensures that no value is present for PortOfDischargeId, not even an explicit nil
### GetPlaceOfReceipt

`func (o *SeawayBillDto) GetPlaceOfReceipt() string`

GetPlaceOfReceipt returns the PlaceOfReceipt field if non-nil, zero value otherwise.

### GetPlaceOfReceiptOk

`func (o *SeawayBillDto) GetPlaceOfReceiptOk() (*string, bool)`

GetPlaceOfReceiptOk returns a tuple with the PlaceOfReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceipt

`func (o *SeawayBillDto) SetPlaceOfReceipt(v string)`

SetPlaceOfReceipt sets PlaceOfReceipt field to given value.

### HasPlaceOfReceipt

`func (o *SeawayBillDto) HasPlaceOfReceipt() bool`

HasPlaceOfReceipt returns a boolean if a field has been set.

### SetPlaceOfReceiptNil

`func (o *SeawayBillDto) SetPlaceOfReceiptNil(b bool)`

 SetPlaceOfReceiptNil sets the value for PlaceOfReceipt to be an explicit nil

### UnsetPlaceOfReceipt
`func (o *SeawayBillDto) UnsetPlaceOfReceipt()`

UnsetPlaceOfReceipt ensures that no value is present for PlaceOfReceipt, not even an explicit nil
### GetPlaceOfDelivery

`func (o *SeawayBillDto) GetPlaceOfDelivery() string`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *SeawayBillDto) GetPlaceOfDeliveryOk() (*string, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *SeawayBillDto) SetPlaceOfDelivery(v string)`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *SeawayBillDto) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### SetPlaceOfDeliveryNil

`func (o *SeawayBillDto) SetPlaceOfDeliveryNil(b bool)`

 SetPlaceOfDeliveryNil sets the value for PlaceOfDelivery to be an explicit nil

### UnsetPlaceOfDelivery
`func (o *SeawayBillDto) UnsetPlaceOfDelivery()`

UnsetPlaceOfDelivery ensures that no value is present for PlaceOfDelivery, not even an explicit nil
### GetDateIssued

`func (o *SeawayBillDto) GetDateIssued() time.Time`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *SeawayBillDto) GetDateIssuedOk() (*time.Time, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *SeawayBillDto) SetDateIssued(v time.Time)`

SetDateIssued sets DateIssued field to given value.

### HasDateIssued

`func (o *SeawayBillDto) HasDateIssued() bool`

HasDateIssued returns a boolean if a field has been set.

### SetDateIssuedNil

`func (o *SeawayBillDto) SetDateIssuedNil(b bool)`

 SetDateIssuedNil sets the value for DateIssued to be an explicit nil

### UnsetDateIssued
`func (o *SeawayBillDto) UnsetDateIssued()`

UnsetDateIssued ensures that no value is present for DateIssued, not even an explicit nil
### GetDateShipped

`func (o *SeawayBillDto) GetDateShipped() time.Time`

GetDateShipped returns the DateShipped field if non-nil, zero value otherwise.

### GetDateShippedOk

`func (o *SeawayBillDto) GetDateShippedOk() (*time.Time, bool)`

GetDateShippedOk returns a tuple with the DateShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateShipped

`func (o *SeawayBillDto) SetDateShipped(v time.Time)`

SetDateShipped sets DateShipped field to given value.

### HasDateShipped

`func (o *SeawayBillDto) HasDateShipped() bool`

HasDateShipped returns a boolean if a field has been set.

### SetDateShippedNil

`func (o *SeawayBillDto) SetDateShippedNil(b bool)`

 SetDateShippedNil sets the value for DateShipped to be an explicit nil

### UnsetDateShipped
`func (o *SeawayBillDto) UnsetDateShipped()`

UnsetDateShipped ensures that no value is present for DateShipped, not even an explicit nil
### GetDateDelivered

`func (o *SeawayBillDto) GetDateDelivered() time.Time`

GetDateDelivered returns the DateDelivered field if non-nil, zero value otherwise.

### GetDateDeliveredOk

`func (o *SeawayBillDto) GetDateDeliveredOk() (*time.Time, bool)`

GetDateDeliveredOk returns a tuple with the DateDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDelivered

`func (o *SeawayBillDto) SetDateDelivered(v time.Time)`

SetDateDelivered sets DateDelivered field to given value.

### HasDateDelivered

`func (o *SeawayBillDto) HasDateDelivered() bool`

HasDateDelivered returns a boolean if a field has been set.

### SetDateDeliveredNil

`func (o *SeawayBillDto) SetDateDeliveredNil(b bool)`

 SetDateDeliveredNil sets the value for DateDelivered to be an explicit nil

### UnsetDateDelivered
`func (o *SeawayBillDto) UnsetDateDelivered()`

UnsetDateDelivered ensures that no value is present for DateDelivered, not even an explicit nil
### GetFreightTerms

`func (o *SeawayBillDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *SeawayBillDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *SeawayBillDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *SeawayBillDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *SeawayBillDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *SeawayBillDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *SeawayBillDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *SeawayBillDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *SeawayBillDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *SeawayBillDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *SeawayBillDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *SeawayBillDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *SeawayBillDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *SeawayBillDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *SeawayBillDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *SeawayBillDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *SeawayBillDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *SeawayBillDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalWeight

`func (o *SeawayBillDto) GetTotalWeight() float64`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SeawayBillDto) GetTotalWeightOk() (*float64, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SeawayBillDto) SetTotalWeight(v float64)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *SeawayBillDto) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### SetTotalWeightNil

`func (o *SeawayBillDto) SetTotalWeightNil(b bool)`

 SetTotalWeightNil sets the value for TotalWeight to be an explicit nil

### UnsetTotalWeight
`func (o *SeawayBillDto) UnsetTotalWeight()`

UnsetTotalWeight ensures that no value is present for TotalWeight, not even an explicit nil
### GetTotalPackages

`func (o *SeawayBillDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *SeawayBillDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *SeawayBillDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *SeawayBillDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *SeawayBillDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *SeawayBillDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetSpecialInstructions

`func (o *SeawayBillDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *SeawayBillDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *SeawayBillDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *SeawayBillDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *SeawayBillDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *SeawayBillDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *SeawayBillDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *SeawayBillDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *SeawayBillDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *SeawayBillDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *SeawayBillDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *SeawayBillDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *SeawayBillDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *SeawayBillDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *SeawayBillDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *SeawayBillDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *SeawayBillDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *SeawayBillDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetTenantId

`func (o *SeawayBillDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SeawayBillDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SeawayBillDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SeawayBillDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *SeawayBillDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *SeawayBillDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *SeawayBillDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *SeawayBillDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *SeawayBillDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *SeawayBillDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *SeawayBillDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *SeawayBillDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetLines

`func (o *SeawayBillDto) GetLines() []WaybillLineDto`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *SeawayBillDto) GetLinesOk() (*[]WaybillLineDto, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *SeawayBillDto) SetLines(v []WaybillLineDto)`

SetLines sets Lines field to given value.

### HasLines

`func (o *SeawayBillDto) HasLines() bool`

HasLines returns a boolean if a field has been set.

### SetLinesNil

`func (o *SeawayBillDto) SetLinesNil(b bool)`

 SetLinesNil sets the value for Lines to be an explicit nil

### UnsetLines
`func (o *SeawayBillDto) UnsetLines()`

UnsetLines ensures that no value is present for Lines, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


