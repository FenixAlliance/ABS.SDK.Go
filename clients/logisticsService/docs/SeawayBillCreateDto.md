# SeawayBillCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**DocumentNumber** | Pointer to **NullableString** |  | [optional] 
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
**FreightTerms** | Pointer to **NullableString** |  | [optional] 
**FreightAmount** | Pointer to **NullableFloat64** |  | [optional] 
**FreightCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalWeight** | Pointer to **NullableFloat64** |  | [optional] 
**TotalPackages** | Pointer to **NullableInt32** |  | [optional] 
**SpecialInstructions** | Pointer to **NullableString** |  | [optional] 
**Remarks** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSeawayBillCreateDto

`func NewSeawayBillCreateDto() *SeawayBillCreateDto`

NewSeawayBillCreateDto instantiates a new SeawayBillCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeawayBillCreateDtoWithDefaults

`func NewSeawayBillCreateDtoWithDefaults() *SeawayBillCreateDto`

NewSeawayBillCreateDtoWithDefaults instantiates a new SeawayBillCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SeawayBillCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SeawayBillCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SeawayBillCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SeawayBillCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *SeawayBillCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SeawayBillCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SeawayBillCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SeawayBillCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDocumentNumber

`func (o *SeawayBillCreateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *SeawayBillCreateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *SeawayBillCreateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *SeawayBillCreateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *SeawayBillCreateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *SeawayBillCreateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetShipperContactId

`func (o *SeawayBillCreateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *SeawayBillCreateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *SeawayBillCreateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *SeawayBillCreateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *SeawayBillCreateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *SeawayBillCreateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *SeawayBillCreateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *SeawayBillCreateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *SeawayBillCreateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *SeawayBillCreateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *SeawayBillCreateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *SeawayBillCreateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetNotifyPartyContactId

`func (o *SeawayBillCreateDto) GetNotifyPartyContactId() string`

GetNotifyPartyContactId returns the NotifyPartyContactId field if non-nil, zero value otherwise.

### GetNotifyPartyContactIdOk

`func (o *SeawayBillCreateDto) GetNotifyPartyContactIdOk() (*string, bool)`

GetNotifyPartyContactIdOk returns a tuple with the NotifyPartyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPartyContactId

`func (o *SeawayBillCreateDto) SetNotifyPartyContactId(v string)`

SetNotifyPartyContactId sets NotifyPartyContactId field to given value.

### HasNotifyPartyContactId

`func (o *SeawayBillCreateDto) HasNotifyPartyContactId() bool`

HasNotifyPartyContactId returns a boolean if a field has been set.

### SetNotifyPartyContactIdNil

`func (o *SeawayBillCreateDto) SetNotifyPartyContactIdNil(b bool)`

 SetNotifyPartyContactIdNil sets the value for NotifyPartyContactId to be an explicit nil

### UnsetNotifyPartyContactId
`func (o *SeawayBillCreateDto) UnsetNotifyPartyContactId()`

UnsetNotifyPartyContactId ensures that no value is present for NotifyPartyContactId, not even an explicit nil
### GetCarrierId

`func (o *SeawayBillCreateDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *SeawayBillCreateDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *SeawayBillCreateDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *SeawayBillCreateDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *SeawayBillCreateDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *SeawayBillCreateDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetVesselId

`func (o *SeawayBillCreateDto) GetVesselId() string`

GetVesselId returns the VesselId field if non-nil, zero value otherwise.

### GetVesselIdOk

`func (o *SeawayBillCreateDto) GetVesselIdOk() (*string, bool)`

GetVesselIdOk returns a tuple with the VesselId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselId

`func (o *SeawayBillCreateDto) SetVesselId(v string)`

SetVesselId sets VesselId field to given value.

### HasVesselId

`func (o *SeawayBillCreateDto) HasVesselId() bool`

HasVesselId returns a boolean if a field has been set.

### SetVesselIdNil

`func (o *SeawayBillCreateDto) SetVesselIdNil(b bool)`

 SetVesselIdNil sets the value for VesselId to be an explicit nil

### UnsetVesselId
`func (o *SeawayBillCreateDto) UnsetVesselId()`

UnsetVesselId ensures that no value is present for VesselId, not even an explicit nil
### GetVoyageId

`func (o *SeawayBillCreateDto) GetVoyageId() string`

GetVoyageId returns the VoyageId field if non-nil, zero value otherwise.

### GetVoyageIdOk

`func (o *SeawayBillCreateDto) GetVoyageIdOk() (*string, bool)`

GetVoyageIdOk returns a tuple with the VoyageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageId

`func (o *SeawayBillCreateDto) SetVoyageId(v string)`

SetVoyageId sets VoyageId field to given value.

### HasVoyageId

`func (o *SeawayBillCreateDto) HasVoyageId() bool`

HasVoyageId returns a boolean if a field has been set.

### SetVoyageIdNil

`func (o *SeawayBillCreateDto) SetVoyageIdNil(b bool)`

 SetVoyageIdNil sets the value for VoyageId to be an explicit nil

### UnsetVoyageId
`func (o *SeawayBillCreateDto) UnsetVoyageId()`

UnsetVoyageId ensures that no value is present for VoyageId, not even an explicit nil
### GetPortOfLoadingId

`func (o *SeawayBillCreateDto) GetPortOfLoadingId() string`

GetPortOfLoadingId returns the PortOfLoadingId field if non-nil, zero value otherwise.

### GetPortOfLoadingIdOk

`func (o *SeawayBillCreateDto) GetPortOfLoadingIdOk() (*string, bool)`

GetPortOfLoadingIdOk returns a tuple with the PortOfLoadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoadingId

`func (o *SeawayBillCreateDto) SetPortOfLoadingId(v string)`

SetPortOfLoadingId sets PortOfLoadingId field to given value.

### HasPortOfLoadingId

`func (o *SeawayBillCreateDto) HasPortOfLoadingId() bool`

HasPortOfLoadingId returns a boolean if a field has been set.

### SetPortOfLoadingIdNil

`func (o *SeawayBillCreateDto) SetPortOfLoadingIdNil(b bool)`

 SetPortOfLoadingIdNil sets the value for PortOfLoadingId to be an explicit nil

### UnsetPortOfLoadingId
`func (o *SeawayBillCreateDto) UnsetPortOfLoadingId()`

UnsetPortOfLoadingId ensures that no value is present for PortOfLoadingId, not even an explicit nil
### GetPortOfDischargeId

`func (o *SeawayBillCreateDto) GetPortOfDischargeId() string`

GetPortOfDischargeId returns the PortOfDischargeId field if non-nil, zero value otherwise.

### GetPortOfDischargeIdOk

`func (o *SeawayBillCreateDto) GetPortOfDischargeIdOk() (*string, bool)`

GetPortOfDischargeIdOk returns a tuple with the PortOfDischargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfDischargeId

`func (o *SeawayBillCreateDto) SetPortOfDischargeId(v string)`

SetPortOfDischargeId sets PortOfDischargeId field to given value.

### HasPortOfDischargeId

`func (o *SeawayBillCreateDto) HasPortOfDischargeId() bool`

HasPortOfDischargeId returns a boolean if a field has been set.

### SetPortOfDischargeIdNil

`func (o *SeawayBillCreateDto) SetPortOfDischargeIdNil(b bool)`

 SetPortOfDischargeIdNil sets the value for PortOfDischargeId to be an explicit nil

### UnsetPortOfDischargeId
`func (o *SeawayBillCreateDto) UnsetPortOfDischargeId()`

UnsetPortOfDischargeId ensures that no value is present for PortOfDischargeId, not even an explicit nil
### GetPlaceOfReceipt

`func (o *SeawayBillCreateDto) GetPlaceOfReceipt() string`

GetPlaceOfReceipt returns the PlaceOfReceipt field if non-nil, zero value otherwise.

### GetPlaceOfReceiptOk

`func (o *SeawayBillCreateDto) GetPlaceOfReceiptOk() (*string, bool)`

GetPlaceOfReceiptOk returns a tuple with the PlaceOfReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceipt

`func (o *SeawayBillCreateDto) SetPlaceOfReceipt(v string)`

SetPlaceOfReceipt sets PlaceOfReceipt field to given value.

### HasPlaceOfReceipt

`func (o *SeawayBillCreateDto) HasPlaceOfReceipt() bool`

HasPlaceOfReceipt returns a boolean if a field has been set.

### SetPlaceOfReceiptNil

`func (o *SeawayBillCreateDto) SetPlaceOfReceiptNil(b bool)`

 SetPlaceOfReceiptNil sets the value for PlaceOfReceipt to be an explicit nil

### UnsetPlaceOfReceipt
`func (o *SeawayBillCreateDto) UnsetPlaceOfReceipt()`

UnsetPlaceOfReceipt ensures that no value is present for PlaceOfReceipt, not even an explicit nil
### GetPlaceOfDelivery

`func (o *SeawayBillCreateDto) GetPlaceOfDelivery() string`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *SeawayBillCreateDto) GetPlaceOfDeliveryOk() (*string, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *SeawayBillCreateDto) SetPlaceOfDelivery(v string)`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *SeawayBillCreateDto) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### SetPlaceOfDeliveryNil

`func (o *SeawayBillCreateDto) SetPlaceOfDeliveryNil(b bool)`

 SetPlaceOfDeliveryNil sets the value for PlaceOfDelivery to be an explicit nil

### UnsetPlaceOfDelivery
`func (o *SeawayBillCreateDto) UnsetPlaceOfDelivery()`

UnsetPlaceOfDelivery ensures that no value is present for PlaceOfDelivery, not even an explicit nil
### GetDateIssued

`func (o *SeawayBillCreateDto) GetDateIssued() time.Time`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *SeawayBillCreateDto) GetDateIssuedOk() (*time.Time, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *SeawayBillCreateDto) SetDateIssued(v time.Time)`

SetDateIssued sets DateIssued field to given value.

### HasDateIssued

`func (o *SeawayBillCreateDto) HasDateIssued() bool`

HasDateIssued returns a boolean if a field has been set.

### SetDateIssuedNil

`func (o *SeawayBillCreateDto) SetDateIssuedNil(b bool)`

 SetDateIssuedNil sets the value for DateIssued to be an explicit nil

### UnsetDateIssued
`func (o *SeawayBillCreateDto) UnsetDateIssued()`

UnsetDateIssued ensures that no value is present for DateIssued, not even an explicit nil
### GetDateShipped

`func (o *SeawayBillCreateDto) GetDateShipped() time.Time`

GetDateShipped returns the DateShipped field if non-nil, zero value otherwise.

### GetDateShippedOk

`func (o *SeawayBillCreateDto) GetDateShippedOk() (*time.Time, bool)`

GetDateShippedOk returns a tuple with the DateShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateShipped

`func (o *SeawayBillCreateDto) SetDateShipped(v time.Time)`

SetDateShipped sets DateShipped field to given value.

### HasDateShipped

`func (o *SeawayBillCreateDto) HasDateShipped() bool`

HasDateShipped returns a boolean if a field has been set.

### SetDateShippedNil

`func (o *SeawayBillCreateDto) SetDateShippedNil(b bool)`

 SetDateShippedNil sets the value for DateShipped to be an explicit nil

### UnsetDateShipped
`func (o *SeawayBillCreateDto) UnsetDateShipped()`

UnsetDateShipped ensures that no value is present for DateShipped, not even an explicit nil
### GetFreightTerms

`func (o *SeawayBillCreateDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *SeawayBillCreateDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *SeawayBillCreateDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *SeawayBillCreateDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *SeawayBillCreateDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *SeawayBillCreateDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *SeawayBillCreateDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *SeawayBillCreateDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *SeawayBillCreateDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *SeawayBillCreateDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *SeawayBillCreateDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *SeawayBillCreateDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *SeawayBillCreateDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *SeawayBillCreateDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *SeawayBillCreateDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *SeawayBillCreateDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *SeawayBillCreateDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *SeawayBillCreateDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalWeight

`func (o *SeawayBillCreateDto) GetTotalWeight() float64`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SeawayBillCreateDto) GetTotalWeightOk() (*float64, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SeawayBillCreateDto) SetTotalWeight(v float64)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *SeawayBillCreateDto) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### SetTotalWeightNil

`func (o *SeawayBillCreateDto) SetTotalWeightNil(b bool)`

 SetTotalWeightNil sets the value for TotalWeight to be an explicit nil

### UnsetTotalWeight
`func (o *SeawayBillCreateDto) UnsetTotalWeight()`

UnsetTotalWeight ensures that no value is present for TotalWeight, not even an explicit nil
### GetTotalPackages

`func (o *SeawayBillCreateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *SeawayBillCreateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *SeawayBillCreateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *SeawayBillCreateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *SeawayBillCreateDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *SeawayBillCreateDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetSpecialInstructions

`func (o *SeawayBillCreateDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *SeawayBillCreateDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *SeawayBillCreateDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *SeawayBillCreateDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *SeawayBillCreateDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *SeawayBillCreateDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *SeawayBillCreateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *SeawayBillCreateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *SeawayBillCreateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *SeawayBillCreateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *SeawayBillCreateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *SeawayBillCreateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *SeawayBillCreateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *SeawayBillCreateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *SeawayBillCreateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *SeawayBillCreateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *SeawayBillCreateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *SeawayBillCreateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


