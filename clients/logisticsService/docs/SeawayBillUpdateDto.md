# SeawayBillUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSeawayBillUpdateDto

`func NewSeawayBillUpdateDto() *SeawayBillUpdateDto`

NewSeawayBillUpdateDto instantiates a new SeawayBillUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeawayBillUpdateDtoWithDefaults

`func NewSeawayBillUpdateDtoWithDefaults() *SeawayBillUpdateDto`

NewSeawayBillUpdateDtoWithDefaults instantiates a new SeawayBillUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *SeawayBillUpdateDto) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *SeawayBillUpdateDto) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *SeawayBillUpdateDto) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *SeawayBillUpdateDto) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### SetDocumentNumberNil

`func (o *SeawayBillUpdateDto) SetDocumentNumberNil(b bool)`

 SetDocumentNumberNil sets the value for DocumentNumber to be an explicit nil

### UnsetDocumentNumber
`func (o *SeawayBillUpdateDto) UnsetDocumentNumber()`

UnsetDocumentNumber ensures that no value is present for DocumentNumber, not even an explicit nil
### GetShipperContactId

`func (o *SeawayBillUpdateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *SeawayBillUpdateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *SeawayBillUpdateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *SeawayBillUpdateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *SeawayBillUpdateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *SeawayBillUpdateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *SeawayBillUpdateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *SeawayBillUpdateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *SeawayBillUpdateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *SeawayBillUpdateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *SeawayBillUpdateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *SeawayBillUpdateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetNotifyPartyContactId

`func (o *SeawayBillUpdateDto) GetNotifyPartyContactId() string`

GetNotifyPartyContactId returns the NotifyPartyContactId field if non-nil, zero value otherwise.

### GetNotifyPartyContactIdOk

`func (o *SeawayBillUpdateDto) GetNotifyPartyContactIdOk() (*string, bool)`

GetNotifyPartyContactIdOk returns a tuple with the NotifyPartyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPartyContactId

`func (o *SeawayBillUpdateDto) SetNotifyPartyContactId(v string)`

SetNotifyPartyContactId sets NotifyPartyContactId field to given value.

### HasNotifyPartyContactId

`func (o *SeawayBillUpdateDto) HasNotifyPartyContactId() bool`

HasNotifyPartyContactId returns a boolean if a field has been set.

### SetNotifyPartyContactIdNil

`func (o *SeawayBillUpdateDto) SetNotifyPartyContactIdNil(b bool)`

 SetNotifyPartyContactIdNil sets the value for NotifyPartyContactId to be an explicit nil

### UnsetNotifyPartyContactId
`func (o *SeawayBillUpdateDto) UnsetNotifyPartyContactId()`

UnsetNotifyPartyContactId ensures that no value is present for NotifyPartyContactId, not even an explicit nil
### GetCarrierId

`func (o *SeawayBillUpdateDto) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *SeawayBillUpdateDto) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *SeawayBillUpdateDto) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *SeawayBillUpdateDto) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### SetCarrierIdNil

`func (o *SeawayBillUpdateDto) SetCarrierIdNil(b bool)`

 SetCarrierIdNil sets the value for CarrierId to be an explicit nil

### UnsetCarrierId
`func (o *SeawayBillUpdateDto) UnsetCarrierId()`

UnsetCarrierId ensures that no value is present for CarrierId, not even an explicit nil
### GetVesselId

`func (o *SeawayBillUpdateDto) GetVesselId() string`

GetVesselId returns the VesselId field if non-nil, zero value otherwise.

### GetVesselIdOk

`func (o *SeawayBillUpdateDto) GetVesselIdOk() (*string, bool)`

GetVesselIdOk returns a tuple with the VesselId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselId

`func (o *SeawayBillUpdateDto) SetVesselId(v string)`

SetVesselId sets VesselId field to given value.

### HasVesselId

`func (o *SeawayBillUpdateDto) HasVesselId() bool`

HasVesselId returns a boolean if a field has been set.

### SetVesselIdNil

`func (o *SeawayBillUpdateDto) SetVesselIdNil(b bool)`

 SetVesselIdNil sets the value for VesselId to be an explicit nil

### UnsetVesselId
`func (o *SeawayBillUpdateDto) UnsetVesselId()`

UnsetVesselId ensures that no value is present for VesselId, not even an explicit nil
### GetVoyageId

`func (o *SeawayBillUpdateDto) GetVoyageId() string`

GetVoyageId returns the VoyageId field if non-nil, zero value otherwise.

### GetVoyageIdOk

`func (o *SeawayBillUpdateDto) GetVoyageIdOk() (*string, bool)`

GetVoyageIdOk returns a tuple with the VoyageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageId

`func (o *SeawayBillUpdateDto) SetVoyageId(v string)`

SetVoyageId sets VoyageId field to given value.

### HasVoyageId

`func (o *SeawayBillUpdateDto) HasVoyageId() bool`

HasVoyageId returns a boolean if a field has been set.

### SetVoyageIdNil

`func (o *SeawayBillUpdateDto) SetVoyageIdNil(b bool)`

 SetVoyageIdNil sets the value for VoyageId to be an explicit nil

### UnsetVoyageId
`func (o *SeawayBillUpdateDto) UnsetVoyageId()`

UnsetVoyageId ensures that no value is present for VoyageId, not even an explicit nil
### GetPortOfLoadingId

`func (o *SeawayBillUpdateDto) GetPortOfLoadingId() string`

GetPortOfLoadingId returns the PortOfLoadingId field if non-nil, zero value otherwise.

### GetPortOfLoadingIdOk

`func (o *SeawayBillUpdateDto) GetPortOfLoadingIdOk() (*string, bool)`

GetPortOfLoadingIdOk returns a tuple with the PortOfLoadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoadingId

`func (o *SeawayBillUpdateDto) SetPortOfLoadingId(v string)`

SetPortOfLoadingId sets PortOfLoadingId field to given value.

### HasPortOfLoadingId

`func (o *SeawayBillUpdateDto) HasPortOfLoadingId() bool`

HasPortOfLoadingId returns a boolean if a field has been set.

### SetPortOfLoadingIdNil

`func (o *SeawayBillUpdateDto) SetPortOfLoadingIdNil(b bool)`

 SetPortOfLoadingIdNil sets the value for PortOfLoadingId to be an explicit nil

### UnsetPortOfLoadingId
`func (o *SeawayBillUpdateDto) UnsetPortOfLoadingId()`

UnsetPortOfLoadingId ensures that no value is present for PortOfLoadingId, not even an explicit nil
### GetPortOfDischargeId

`func (o *SeawayBillUpdateDto) GetPortOfDischargeId() string`

GetPortOfDischargeId returns the PortOfDischargeId field if non-nil, zero value otherwise.

### GetPortOfDischargeIdOk

`func (o *SeawayBillUpdateDto) GetPortOfDischargeIdOk() (*string, bool)`

GetPortOfDischargeIdOk returns a tuple with the PortOfDischargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfDischargeId

`func (o *SeawayBillUpdateDto) SetPortOfDischargeId(v string)`

SetPortOfDischargeId sets PortOfDischargeId field to given value.

### HasPortOfDischargeId

`func (o *SeawayBillUpdateDto) HasPortOfDischargeId() bool`

HasPortOfDischargeId returns a boolean if a field has been set.

### SetPortOfDischargeIdNil

`func (o *SeawayBillUpdateDto) SetPortOfDischargeIdNil(b bool)`

 SetPortOfDischargeIdNil sets the value for PortOfDischargeId to be an explicit nil

### UnsetPortOfDischargeId
`func (o *SeawayBillUpdateDto) UnsetPortOfDischargeId()`

UnsetPortOfDischargeId ensures that no value is present for PortOfDischargeId, not even an explicit nil
### GetPlaceOfReceipt

`func (o *SeawayBillUpdateDto) GetPlaceOfReceipt() string`

GetPlaceOfReceipt returns the PlaceOfReceipt field if non-nil, zero value otherwise.

### GetPlaceOfReceiptOk

`func (o *SeawayBillUpdateDto) GetPlaceOfReceiptOk() (*string, bool)`

GetPlaceOfReceiptOk returns a tuple with the PlaceOfReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceipt

`func (o *SeawayBillUpdateDto) SetPlaceOfReceipt(v string)`

SetPlaceOfReceipt sets PlaceOfReceipt field to given value.

### HasPlaceOfReceipt

`func (o *SeawayBillUpdateDto) HasPlaceOfReceipt() bool`

HasPlaceOfReceipt returns a boolean if a field has been set.

### SetPlaceOfReceiptNil

`func (o *SeawayBillUpdateDto) SetPlaceOfReceiptNil(b bool)`

 SetPlaceOfReceiptNil sets the value for PlaceOfReceipt to be an explicit nil

### UnsetPlaceOfReceipt
`func (o *SeawayBillUpdateDto) UnsetPlaceOfReceipt()`

UnsetPlaceOfReceipt ensures that no value is present for PlaceOfReceipt, not even an explicit nil
### GetPlaceOfDelivery

`func (o *SeawayBillUpdateDto) GetPlaceOfDelivery() string`

GetPlaceOfDelivery returns the PlaceOfDelivery field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryOk

`func (o *SeawayBillUpdateDto) GetPlaceOfDeliveryOk() (*string, bool)`

GetPlaceOfDeliveryOk returns a tuple with the PlaceOfDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDelivery

`func (o *SeawayBillUpdateDto) SetPlaceOfDelivery(v string)`

SetPlaceOfDelivery sets PlaceOfDelivery field to given value.

### HasPlaceOfDelivery

`func (o *SeawayBillUpdateDto) HasPlaceOfDelivery() bool`

HasPlaceOfDelivery returns a boolean if a field has been set.

### SetPlaceOfDeliveryNil

`func (o *SeawayBillUpdateDto) SetPlaceOfDeliveryNil(b bool)`

 SetPlaceOfDeliveryNil sets the value for PlaceOfDelivery to be an explicit nil

### UnsetPlaceOfDelivery
`func (o *SeawayBillUpdateDto) UnsetPlaceOfDelivery()`

UnsetPlaceOfDelivery ensures that no value is present for PlaceOfDelivery, not even an explicit nil
### GetDateIssued

`func (o *SeawayBillUpdateDto) GetDateIssued() time.Time`

GetDateIssued returns the DateIssued field if non-nil, zero value otherwise.

### GetDateIssuedOk

`func (o *SeawayBillUpdateDto) GetDateIssuedOk() (*time.Time, bool)`

GetDateIssuedOk returns a tuple with the DateIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateIssued

`func (o *SeawayBillUpdateDto) SetDateIssued(v time.Time)`

SetDateIssued sets DateIssued field to given value.

### HasDateIssued

`func (o *SeawayBillUpdateDto) HasDateIssued() bool`

HasDateIssued returns a boolean if a field has been set.

### SetDateIssuedNil

`func (o *SeawayBillUpdateDto) SetDateIssuedNil(b bool)`

 SetDateIssuedNil sets the value for DateIssued to be an explicit nil

### UnsetDateIssued
`func (o *SeawayBillUpdateDto) UnsetDateIssued()`

UnsetDateIssued ensures that no value is present for DateIssued, not even an explicit nil
### GetDateShipped

`func (o *SeawayBillUpdateDto) GetDateShipped() time.Time`

GetDateShipped returns the DateShipped field if non-nil, zero value otherwise.

### GetDateShippedOk

`func (o *SeawayBillUpdateDto) GetDateShippedOk() (*time.Time, bool)`

GetDateShippedOk returns a tuple with the DateShipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateShipped

`func (o *SeawayBillUpdateDto) SetDateShipped(v time.Time)`

SetDateShipped sets DateShipped field to given value.

### HasDateShipped

`func (o *SeawayBillUpdateDto) HasDateShipped() bool`

HasDateShipped returns a boolean if a field has been set.

### SetDateShippedNil

`func (o *SeawayBillUpdateDto) SetDateShippedNil(b bool)`

 SetDateShippedNil sets the value for DateShipped to be an explicit nil

### UnsetDateShipped
`func (o *SeawayBillUpdateDto) UnsetDateShipped()`

UnsetDateShipped ensures that no value is present for DateShipped, not even an explicit nil
### GetFreightTerms

`func (o *SeawayBillUpdateDto) GetFreightTerms() string`

GetFreightTerms returns the FreightTerms field if non-nil, zero value otherwise.

### GetFreightTermsOk

`func (o *SeawayBillUpdateDto) GetFreightTermsOk() (*string, bool)`

GetFreightTermsOk returns a tuple with the FreightTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTerms

`func (o *SeawayBillUpdateDto) SetFreightTerms(v string)`

SetFreightTerms sets FreightTerms field to given value.

### HasFreightTerms

`func (o *SeawayBillUpdateDto) HasFreightTerms() bool`

HasFreightTerms returns a boolean if a field has been set.

### SetFreightTermsNil

`func (o *SeawayBillUpdateDto) SetFreightTermsNil(b bool)`

 SetFreightTermsNil sets the value for FreightTerms to be an explicit nil

### UnsetFreightTerms
`func (o *SeawayBillUpdateDto) UnsetFreightTerms()`

UnsetFreightTerms ensures that no value is present for FreightTerms, not even an explicit nil
### GetFreightAmount

`func (o *SeawayBillUpdateDto) GetFreightAmount() float64`

GetFreightAmount returns the FreightAmount field if non-nil, zero value otherwise.

### GetFreightAmountOk

`func (o *SeawayBillUpdateDto) GetFreightAmountOk() (*float64, bool)`

GetFreightAmountOk returns a tuple with the FreightAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAmount

`func (o *SeawayBillUpdateDto) SetFreightAmount(v float64)`

SetFreightAmount sets FreightAmount field to given value.

### HasFreightAmount

`func (o *SeawayBillUpdateDto) HasFreightAmount() bool`

HasFreightAmount returns a boolean if a field has been set.

### SetFreightAmountNil

`func (o *SeawayBillUpdateDto) SetFreightAmountNil(b bool)`

 SetFreightAmountNil sets the value for FreightAmount to be an explicit nil

### UnsetFreightAmount
`func (o *SeawayBillUpdateDto) UnsetFreightAmount()`

UnsetFreightAmount ensures that no value is present for FreightAmount, not even an explicit nil
### GetFreightCurrencyId

`func (o *SeawayBillUpdateDto) GetFreightCurrencyId() string`

GetFreightCurrencyId returns the FreightCurrencyId field if non-nil, zero value otherwise.

### GetFreightCurrencyIdOk

`func (o *SeawayBillUpdateDto) GetFreightCurrencyIdOk() (*string, bool)`

GetFreightCurrencyIdOk returns a tuple with the FreightCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightCurrencyId

`func (o *SeawayBillUpdateDto) SetFreightCurrencyId(v string)`

SetFreightCurrencyId sets FreightCurrencyId field to given value.

### HasFreightCurrencyId

`func (o *SeawayBillUpdateDto) HasFreightCurrencyId() bool`

HasFreightCurrencyId returns a boolean if a field has been set.

### SetFreightCurrencyIdNil

`func (o *SeawayBillUpdateDto) SetFreightCurrencyIdNil(b bool)`

 SetFreightCurrencyIdNil sets the value for FreightCurrencyId to be an explicit nil

### UnsetFreightCurrencyId
`func (o *SeawayBillUpdateDto) UnsetFreightCurrencyId()`

UnsetFreightCurrencyId ensures that no value is present for FreightCurrencyId, not even an explicit nil
### GetTotalWeight

`func (o *SeawayBillUpdateDto) GetTotalWeight() float64`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *SeawayBillUpdateDto) GetTotalWeightOk() (*float64, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *SeawayBillUpdateDto) SetTotalWeight(v float64)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *SeawayBillUpdateDto) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### SetTotalWeightNil

`func (o *SeawayBillUpdateDto) SetTotalWeightNil(b bool)`

 SetTotalWeightNil sets the value for TotalWeight to be an explicit nil

### UnsetTotalWeight
`func (o *SeawayBillUpdateDto) UnsetTotalWeight()`

UnsetTotalWeight ensures that no value is present for TotalWeight, not even an explicit nil
### GetTotalPackages

`func (o *SeawayBillUpdateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *SeawayBillUpdateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *SeawayBillUpdateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *SeawayBillUpdateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### SetTotalPackagesNil

`func (o *SeawayBillUpdateDto) SetTotalPackagesNil(b bool)`

 SetTotalPackagesNil sets the value for TotalPackages to be an explicit nil

### UnsetTotalPackages
`func (o *SeawayBillUpdateDto) UnsetTotalPackages()`

UnsetTotalPackages ensures that no value is present for TotalPackages, not even an explicit nil
### GetSpecialInstructions

`func (o *SeawayBillUpdateDto) GetSpecialInstructions() string`

GetSpecialInstructions returns the SpecialInstructions field if non-nil, zero value otherwise.

### GetSpecialInstructionsOk

`func (o *SeawayBillUpdateDto) GetSpecialInstructionsOk() (*string, bool)`

GetSpecialInstructionsOk returns a tuple with the SpecialInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstructions

`func (o *SeawayBillUpdateDto) SetSpecialInstructions(v string)`

SetSpecialInstructions sets SpecialInstructions field to given value.

### HasSpecialInstructions

`func (o *SeawayBillUpdateDto) HasSpecialInstructions() bool`

HasSpecialInstructions returns a boolean if a field has been set.

### SetSpecialInstructionsNil

`func (o *SeawayBillUpdateDto) SetSpecialInstructionsNil(b bool)`

 SetSpecialInstructionsNil sets the value for SpecialInstructions to be an explicit nil

### UnsetSpecialInstructions
`func (o *SeawayBillUpdateDto) UnsetSpecialInstructions()`

UnsetSpecialInstructions ensures that no value is present for SpecialInstructions, not even an explicit nil
### GetRemarks

`func (o *SeawayBillUpdateDto) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *SeawayBillUpdateDto) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *SeawayBillUpdateDto) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *SeawayBillUpdateDto) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### SetRemarksNil

`func (o *SeawayBillUpdateDto) SetRemarksNil(b bool)`

 SetRemarksNil sets the value for Remarks to be an explicit nil

### UnsetRemarks
`func (o *SeawayBillUpdateDto) UnsetRemarks()`

UnsetRemarks ensures that no value is present for Remarks, not even an explicit nil
### GetShipmentId

`func (o *SeawayBillUpdateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *SeawayBillUpdateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *SeawayBillUpdateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *SeawayBillUpdateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *SeawayBillUpdateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *SeawayBillUpdateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


