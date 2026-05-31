# BillOfLadingCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BillOfLadingNumber** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BillOfLadingType** | Pointer to **NullableString** |  | [optional] 
**IsNegotiable** | Pointer to **bool** |  | [optional] 
**IsClean** | Pointer to **bool** |  | [optional] 
**NumberOfOriginals** | Pointer to **int32** |  | [optional] 
**FreightPaymentType** | Pointer to **NullableString** |  | [optional] 
**ShippingTerms** | Pointer to **NullableString** |  | [optional] 
**FreightChargesDescription** | Pointer to **NullableString** |  | [optional] 
**DeclaredValueAmount** | Pointer to **float64** |  | [optional] 
**DeclaredValueCurrencyId** | Pointer to **NullableString** |  | [optional] 
**VesselName** | Pointer to **NullableString** |  | [optional] 
**VoyageNumber** | Pointer to **NullableString** |  | [optional] 
**ShipperContactId** | Pointer to **NullableString** |  | [optional] 
**ConsigneeContactId** | Pointer to **NullableString** |  | [optional] 
**NotifyPartyContactId** | Pointer to **NullableString** |  | [optional] 
**ShippingCourierId** | Pointer to **NullableString** |  | [optional] 
**PortOfLoadingId** | Pointer to **NullableString** |  | [optional] 
**PortOfDischargeId** | Pointer to **NullableString** |  | [optional] 
**PlaceOfReceiptId** | Pointer to **NullableString** |  | [optional] 
**PlaceOfDeliveryId** | Pointer to **NullableString** |  | [optional] 
**ShipmentId** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**VoyageId** | Pointer to **NullableString** |  | [optional] 
**MarksAndNumbers** | Pointer to **NullableString** |  | [optional] 
**TotalPackages** | Pointer to **int32** |  | [optional] 
**TotalGrossWeightKg** | Pointer to **float64** |  | [optional] 
**TotalVolumeM3** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewBillOfLadingCreateDto

`func NewBillOfLadingCreateDto() *BillOfLadingCreateDto`

NewBillOfLadingCreateDto instantiates a new BillOfLadingCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfLadingCreateDtoWithDefaults

`func NewBillOfLadingCreateDtoWithDefaults() *BillOfLadingCreateDto`

NewBillOfLadingCreateDtoWithDefaults instantiates a new BillOfLadingCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillOfLadingCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillOfLadingCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillOfLadingCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillOfLadingCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BillOfLadingCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BillOfLadingCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BillOfLadingCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BillOfLadingCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBillOfLadingNumber

`func (o *BillOfLadingCreateDto) GetBillOfLadingNumber() string`

GetBillOfLadingNumber returns the BillOfLadingNumber field if non-nil, zero value otherwise.

### GetBillOfLadingNumberOk

`func (o *BillOfLadingCreateDto) GetBillOfLadingNumberOk() (*string, bool)`

GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingNumber

`func (o *BillOfLadingCreateDto) SetBillOfLadingNumber(v string)`

SetBillOfLadingNumber sets BillOfLadingNumber field to given value.

### HasBillOfLadingNumber

`func (o *BillOfLadingCreateDto) HasBillOfLadingNumber() bool`

HasBillOfLadingNumber returns a boolean if a field has been set.

### SetBillOfLadingNumberNil

`func (o *BillOfLadingCreateDto) SetBillOfLadingNumberNil(b bool)`

 SetBillOfLadingNumberNil sets the value for BillOfLadingNumber to be an explicit nil

### UnsetBillOfLadingNumber
`func (o *BillOfLadingCreateDto) UnsetBillOfLadingNumber()`

UnsetBillOfLadingNumber ensures that no value is present for BillOfLadingNumber, not even an explicit nil
### GetTitle

`func (o *BillOfLadingCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillOfLadingCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillOfLadingCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BillOfLadingCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BillOfLadingCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BillOfLadingCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *BillOfLadingCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillOfLadingCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillOfLadingCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillOfLadingCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillOfLadingCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillOfLadingCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBillOfLadingType

`func (o *BillOfLadingCreateDto) GetBillOfLadingType() string`

GetBillOfLadingType returns the BillOfLadingType field if non-nil, zero value otherwise.

### GetBillOfLadingTypeOk

`func (o *BillOfLadingCreateDto) GetBillOfLadingTypeOk() (*string, bool)`

GetBillOfLadingTypeOk returns a tuple with the BillOfLadingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingType

`func (o *BillOfLadingCreateDto) SetBillOfLadingType(v string)`

SetBillOfLadingType sets BillOfLadingType field to given value.

### HasBillOfLadingType

`func (o *BillOfLadingCreateDto) HasBillOfLadingType() bool`

HasBillOfLadingType returns a boolean if a field has been set.

### SetBillOfLadingTypeNil

`func (o *BillOfLadingCreateDto) SetBillOfLadingTypeNil(b bool)`

 SetBillOfLadingTypeNil sets the value for BillOfLadingType to be an explicit nil

### UnsetBillOfLadingType
`func (o *BillOfLadingCreateDto) UnsetBillOfLadingType()`

UnsetBillOfLadingType ensures that no value is present for BillOfLadingType, not even an explicit nil
### GetIsNegotiable

`func (o *BillOfLadingCreateDto) GetIsNegotiable() bool`

GetIsNegotiable returns the IsNegotiable field if non-nil, zero value otherwise.

### GetIsNegotiableOk

`func (o *BillOfLadingCreateDto) GetIsNegotiableOk() (*bool, bool)`

GetIsNegotiableOk returns a tuple with the IsNegotiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNegotiable

`func (o *BillOfLadingCreateDto) SetIsNegotiable(v bool)`

SetIsNegotiable sets IsNegotiable field to given value.

### HasIsNegotiable

`func (o *BillOfLadingCreateDto) HasIsNegotiable() bool`

HasIsNegotiable returns a boolean if a field has been set.

### GetIsClean

`func (o *BillOfLadingCreateDto) GetIsClean() bool`

GetIsClean returns the IsClean field if non-nil, zero value otherwise.

### GetIsCleanOk

`func (o *BillOfLadingCreateDto) GetIsCleanOk() (*bool, bool)`

GetIsCleanOk returns a tuple with the IsClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClean

`func (o *BillOfLadingCreateDto) SetIsClean(v bool)`

SetIsClean sets IsClean field to given value.

### HasIsClean

`func (o *BillOfLadingCreateDto) HasIsClean() bool`

HasIsClean returns a boolean if a field has been set.

### GetNumberOfOriginals

`func (o *BillOfLadingCreateDto) GetNumberOfOriginals() int32`

GetNumberOfOriginals returns the NumberOfOriginals field if non-nil, zero value otherwise.

### GetNumberOfOriginalsOk

`func (o *BillOfLadingCreateDto) GetNumberOfOriginalsOk() (*int32, bool)`

GetNumberOfOriginalsOk returns a tuple with the NumberOfOriginals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOriginals

`func (o *BillOfLadingCreateDto) SetNumberOfOriginals(v int32)`

SetNumberOfOriginals sets NumberOfOriginals field to given value.

### HasNumberOfOriginals

`func (o *BillOfLadingCreateDto) HasNumberOfOriginals() bool`

HasNumberOfOriginals returns a boolean if a field has been set.

### GetFreightPaymentType

`func (o *BillOfLadingCreateDto) GetFreightPaymentType() string`

GetFreightPaymentType returns the FreightPaymentType field if non-nil, zero value otherwise.

### GetFreightPaymentTypeOk

`func (o *BillOfLadingCreateDto) GetFreightPaymentTypeOk() (*string, bool)`

GetFreightPaymentTypeOk returns a tuple with the FreightPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightPaymentType

`func (o *BillOfLadingCreateDto) SetFreightPaymentType(v string)`

SetFreightPaymentType sets FreightPaymentType field to given value.

### HasFreightPaymentType

`func (o *BillOfLadingCreateDto) HasFreightPaymentType() bool`

HasFreightPaymentType returns a boolean if a field has been set.

### SetFreightPaymentTypeNil

`func (o *BillOfLadingCreateDto) SetFreightPaymentTypeNil(b bool)`

 SetFreightPaymentTypeNil sets the value for FreightPaymentType to be an explicit nil

### UnsetFreightPaymentType
`func (o *BillOfLadingCreateDto) UnsetFreightPaymentType()`

UnsetFreightPaymentType ensures that no value is present for FreightPaymentType, not even an explicit nil
### GetShippingTerms

`func (o *BillOfLadingCreateDto) GetShippingTerms() string`

GetShippingTerms returns the ShippingTerms field if non-nil, zero value otherwise.

### GetShippingTermsOk

`func (o *BillOfLadingCreateDto) GetShippingTermsOk() (*string, bool)`

GetShippingTermsOk returns a tuple with the ShippingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTerms

`func (o *BillOfLadingCreateDto) SetShippingTerms(v string)`

SetShippingTerms sets ShippingTerms field to given value.

### HasShippingTerms

`func (o *BillOfLadingCreateDto) HasShippingTerms() bool`

HasShippingTerms returns a boolean if a field has been set.

### SetShippingTermsNil

`func (o *BillOfLadingCreateDto) SetShippingTermsNil(b bool)`

 SetShippingTermsNil sets the value for ShippingTerms to be an explicit nil

### UnsetShippingTerms
`func (o *BillOfLadingCreateDto) UnsetShippingTerms()`

UnsetShippingTerms ensures that no value is present for ShippingTerms, not even an explicit nil
### GetFreightChargesDescription

`func (o *BillOfLadingCreateDto) GetFreightChargesDescription() string`

GetFreightChargesDescription returns the FreightChargesDescription field if non-nil, zero value otherwise.

### GetFreightChargesDescriptionOk

`func (o *BillOfLadingCreateDto) GetFreightChargesDescriptionOk() (*string, bool)`

GetFreightChargesDescriptionOk returns a tuple with the FreightChargesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightChargesDescription

`func (o *BillOfLadingCreateDto) SetFreightChargesDescription(v string)`

SetFreightChargesDescription sets FreightChargesDescription field to given value.

### HasFreightChargesDescription

`func (o *BillOfLadingCreateDto) HasFreightChargesDescription() bool`

HasFreightChargesDescription returns a boolean if a field has been set.

### SetFreightChargesDescriptionNil

`func (o *BillOfLadingCreateDto) SetFreightChargesDescriptionNil(b bool)`

 SetFreightChargesDescriptionNil sets the value for FreightChargesDescription to be an explicit nil

### UnsetFreightChargesDescription
`func (o *BillOfLadingCreateDto) UnsetFreightChargesDescription()`

UnsetFreightChargesDescription ensures that no value is present for FreightChargesDescription, not even an explicit nil
### GetDeclaredValueAmount

`func (o *BillOfLadingCreateDto) GetDeclaredValueAmount() float64`

GetDeclaredValueAmount returns the DeclaredValueAmount field if non-nil, zero value otherwise.

### GetDeclaredValueAmountOk

`func (o *BillOfLadingCreateDto) GetDeclaredValueAmountOk() (*float64, bool)`

GetDeclaredValueAmountOk returns a tuple with the DeclaredValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueAmount

`func (o *BillOfLadingCreateDto) SetDeclaredValueAmount(v float64)`

SetDeclaredValueAmount sets DeclaredValueAmount field to given value.

### HasDeclaredValueAmount

`func (o *BillOfLadingCreateDto) HasDeclaredValueAmount() bool`

HasDeclaredValueAmount returns a boolean if a field has been set.

### GetDeclaredValueCurrencyId

`func (o *BillOfLadingCreateDto) GetDeclaredValueCurrencyId() string`

GetDeclaredValueCurrencyId returns the DeclaredValueCurrencyId field if non-nil, zero value otherwise.

### GetDeclaredValueCurrencyIdOk

`func (o *BillOfLadingCreateDto) GetDeclaredValueCurrencyIdOk() (*string, bool)`

GetDeclaredValueCurrencyIdOk returns a tuple with the DeclaredValueCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueCurrencyId

`func (o *BillOfLadingCreateDto) SetDeclaredValueCurrencyId(v string)`

SetDeclaredValueCurrencyId sets DeclaredValueCurrencyId field to given value.

### HasDeclaredValueCurrencyId

`func (o *BillOfLadingCreateDto) HasDeclaredValueCurrencyId() bool`

HasDeclaredValueCurrencyId returns a boolean if a field has been set.

### SetDeclaredValueCurrencyIdNil

`func (o *BillOfLadingCreateDto) SetDeclaredValueCurrencyIdNil(b bool)`

 SetDeclaredValueCurrencyIdNil sets the value for DeclaredValueCurrencyId to be an explicit nil

### UnsetDeclaredValueCurrencyId
`func (o *BillOfLadingCreateDto) UnsetDeclaredValueCurrencyId()`

UnsetDeclaredValueCurrencyId ensures that no value is present for DeclaredValueCurrencyId, not even an explicit nil
### GetVesselName

`func (o *BillOfLadingCreateDto) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *BillOfLadingCreateDto) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *BillOfLadingCreateDto) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *BillOfLadingCreateDto) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### SetVesselNameNil

`func (o *BillOfLadingCreateDto) SetVesselNameNil(b bool)`

 SetVesselNameNil sets the value for VesselName to be an explicit nil

### UnsetVesselName
`func (o *BillOfLadingCreateDto) UnsetVesselName()`

UnsetVesselName ensures that no value is present for VesselName, not even an explicit nil
### GetVoyageNumber

`func (o *BillOfLadingCreateDto) GetVoyageNumber() string`

GetVoyageNumber returns the VoyageNumber field if non-nil, zero value otherwise.

### GetVoyageNumberOk

`func (o *BillOfLadingCreateDto) GetVoyageNumberOk() (*string, bool)`

GetVoyageNumberOk returns a tuple with the VoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNumber

`func (o *BillOfLadingCreateDto) SetVoyageNumber(v string)`

SetVoyageNumber sets VoyageNumber field to given value.

### HasVoyageNumber

`func (o *BillOfLadingCreateDto) HasVoyageNumber() bool`

HasVoyageNumber returns a boolean if a field has been set.

### SetVoyageNumberNil

`func (o *BillOfLadingCreateDto) SetVoyageNumberNil(b bool)`

 SetVoyageNumberNil sets the value for VoyageNumber to be an explicit nil

### UnsetVoyageNumber
`func (o *BillOfLadingCreateDto) UnsetVoyageNumber()`

UnsetVoyageNumber ensures that no value is present for VoyageNumber, not even an explicit nil
### GetShipperContactId

`func (o *BillOfLadingCreateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *BillOfLadingCreateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *BillOfLadingCreateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *BillOfLadingCreateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *BillOfLadingCreateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *BillOfLadingCreateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *BillOfLadingCreateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *BillOfLadingCreateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *BillOfLadingCreateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *BillOfLadingCreateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *BillOfLadingCreateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *BillOfLadingCreateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetNotifyPartyContactId

`func (o *BillOfLadingCreateDto) GetNotifyPartyContactId() string`

GetNotifyPartyContactId returns the NotifyPartyContactId field if non-nil, zero value otherwise.

### GetNotifyPartyContactIdOk

`func (o *BillOfLadingCreateDto) GetNotifyPartyContactIdOk() (*string, bool)`

GetNotifyPartyContactIdOk returns a tuple with the NotifyPartyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPartyContactId

`func (o *BillOfLadingCreateDto) SetNotifyPartyContactId(v string)`

SetNotifyPartyContactId sets NotifyPartyContactId field to given value.

### HasNotifyPartyContactId

`func (o *BillOfLadingCreateDto) HasNotifyPartyContactId() bool`

HasNotifyPartyContactId returns a boolean if a field has been set.

### SetNotifyPartyContactIdNil

`func (o *BillOfLadingCreateDto) SetNotifyPartyContactIdNil(b bool)`

 SetNotifyPartyContactIdNil sets the value for NotifyPartyContactId to be an explicit nil

### UnsetNotifyPartyContactId
`func (o *BillOfLadingCreateDto) UnsetNotifyPartyContactId()`

UnsetNotifyPartyContactId ensures that no value is present for NotifyPartyContactId, not even an explicit nil
### GetShippingCourierId

`func (o *BillOfLadingCreateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *BillOfLadingCreateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *BillOfLadingCreateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *BillOfLadingCreateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *BillOfLadingCreateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *BillOfLadingCreateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil
### GetPortOfLoadingId

`func (o *BillOfLadingCreateDto) GetPortOfLoadingId() string`

GetPortOfLoadingId returns the PortOfLoadingId field if non-nil, zero value otherwise.

### GetPortOfLoadingIdOk

`func (o *BillOfLadingCreateDto) GetPortOfLoadingIdOk() (*string, bool)`

GetPortOfLoadingIdOk returns a tuple with the PortOfLoadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoadingId

`func (o *BillOfLadingCreateDto) SetPortOfLoadingId(v string)`

SetPortOfLoadingId sets PortOfLoadingId field to given value.

### HasPortOfLoadingId

`func (o *BillOfLadingCreateDto) HasPortOfLoadingId() bool`

HasPortOfLoadingId returns a boolean if a field has been set.

### SetPortOfLoadingIdNil

`func (o *BillOfLadingCreateDto) SetPortOfLoadingIdNil(b bool)`

 SetPortOfLoadingIdNil sets the value for PortOfLoadingId to be an explicit nil

### UnsetPortOfLoadingId
`func (o *BillOfLadingCreateDto) UnsetPortOfLoadingId()`

UnsetPortOfLoadingId ensures that no value is present for PortOfLoadingId, not even an explicit nil
### GetPortOfDischargeId

`func (o *BillOfLadingCreateDto) GetPortOfDischargeId() string`

GetPortOfDischargeId returns the PortOfDischargeId field if non-nil, zero value otherwise.

### GetPortOfDischargeIdOk

`func (o *BillOfLadingCreateDto) GetPortOfDischargeIdOk() (*string, bool)`

GetPortOfDischargeIdOk returns a tuple with the PortOfDischargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfDischargeId

`func (o *BillOfLadingCreateDto) SetPortOfDischargeId(v string)`

SetPortOfDischargeId sets PortOfDischargeId field to given value.

### HasPortOfDischargeId

`func (o *BillOfLadingCreateDto) HasPortOfDischargeId() bool`

HasPortOfDischargeId returns a boolean if a field has been set.

### SetPortOfDischargeIdNil

`func (o *BillOfLadingCreateDto) SetPortOfDischargeIdNil(b bool)`

 SetPortOfDischargeIdNil sets the value for PortOfDischargeId to be an explicit nil

### UnsetPortOfDischargeId
`func (o *BillOfLadingCreateDto) UnsetPortOfDischargeId()`

UnsetPortOfDischargeId ensures that no value is present for PortOfDischargeId, not even an explicit nil
### GetPlaceOfReceiptId

`func (o *BillOfLadingCreateDto) GetPlaceOfReceiptId() string`

GetPlaceOfReceiptId returns the PlaceOfReceiptId field if non-nil, zero value otherwise.

### GetPlaceOfReceiptIdOk

`func (o *BillOfLadingCreateDto) GetPlaceOfReceiptIdOk() (*string, bool)`

GetPlaceOfReceiptIdOk returns a tuple with the PlaceOfReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceiptId

`func (o *BillOfLadingCreateDto) SetPlaceOfReceiptId(v string)`

SetPlaceOfReceiptId sets PlaceOfReceiptId field to given value.

### HasPlaceOfReceiptId

`func (o *BillOfLadingCreateDto) HasPlaceOfReceiptId() bool`

HasPlaceOfReceiptId returns a boolean if a field has been set.

### SetPlaceOfReceiptIdNil

`func (o *BillOfLadingCreateDto) SetPlaceOfReceiptIdNil(b bool)`

 SetPlaceOfReceiptIdNil sets the value for PlaceOfReceiptId to be an explicit nil

### UnsetPlaceOfReceiptId
`func (o *BillOfLadingCreateDto) UnsetPlaceOfReceiptId()`

UnsetPlaceOfReceiptId ensures that no value is present for PlaceOfReceiptId, not even an explicit nil
### GetPlaceOfDeliveryId

`func (o *BillOfLadingCreateDto) GetPlaceOfDeliveryId() string`

GetPlaceOfDeliveryId returns the PlaceOfDeliveryId field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryIdOk

`func (o *BillOfLadingCreateDto) GetPlaceOfDeliveryIdOk() (*string, bool)`

GetPlaceOfDeliveryIdOk returns a tuple with the PlaceOfDeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryId

`func (o *BillOfLadingCreateDto) SetPlaceOfDeliveryId(v string)`

SetPlaceOfDeliveryId sets PlaceOfDeliveryId field to given value.

### HasPlaceOfDeliveryId

`func (o *BillOfLadingCreateDto) HasPlaceOfDeliveryId() bool`

HasPlaceOfDeliveryId returns a boolean if a field has been set.

### SetPlaceOfDeliveryIdNil

`func (o *BillOfLadingCreateDto) SetPlaceOfDeliveryIdNil(b bool)`

 SetPlaceOfDeliveryIdNil sets the value for PlaceOfDeliveryId to be an explicit nil

### UnsetPlaceOfDeliveryId
`func (o *BillOfLadingCreateDto) UnsetPlaceOfDeliveryId()`

UnsetPlaceOfDeliveryId ensures that no value is present for PlaceOfDeliveryId, not even an explicit nil
### GetShipmentId

`func (o *BillOfLadingCreateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *BillOfLadingCreateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *BillOfLadingCreateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *BillOfLadingCreateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *BillOfLadingCreateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *BillOfLadingCreateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetOrderId

`func (o *BillOfLadingCreateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BillOfLadingCreateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BillOfLadingCreateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *BillOfLadingCreateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *BillOfLadingCreateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *BillOfLadingCreateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetVoyageId

`func (o *BillOfLadingCreateDto) GetVoyageId() string`

GetVoyageId returns the VoyageId field if non-nil, zero value otherwise.

### GetVoyageIdOk

`func (o *BillOfLadingCreateDto) GetVoyageIdOk() (*string, bool)`

GetVoyageIdOk returns a tuple with the VoyageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageId

`func (o *BillOfLadingCreateDto) SetVoyageId(v string)`

SetVoyageId sets VoyageId field to given value.

### HasVoyageId

`func (o *BillOfLadingCreateDto) HasVoyageId() bool`

HasVoyageId returns a boolean if a field has been set.

### SetVoyageIdNil

`func (o *BillOfLadingCreateDto) SetVoyageIdNil(b bool)`

 SetVoyageIdNil sets the value for VoyageId to be an explicit nil

### UnsetVoyageId
`func (o *BillOfLadingCreateDto) UnsetVoyageId()`

UnsetVoyageId ensures that no value is present for VoyageId, not even an explicit nil
### GetMarksAndNumbers

`func (o *BillOfLadingCreateDto) GetMarksAndNumbers() string`

GetMarksAndNumbers returns the MarksAndNumbers field if non-nil, zero value otherwise.

### GetMarksAndNumbersOk

`func (o *BillOfLadingCreateDto) GetMarksAndNumbersOk() (*string, bool)`

GetMarksAndNumbersOk returns a tuple with the MarksAndNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarksAndNumbers

`func (o *BillOfLadingCreateDto) SetMarksAndNumbers(v string)`

SetMarksAndNumbers sets MarksAndNumbers field to given value.

### HasMarksAndNumbers

`func (o *BillOfLadingCreateDto) HasMarksAndNumbers() bool`

HasMarksAndNumbers returns a boolean if a field has been set.

### SetMarksAndNumbersNil

`func (o *BillOfLadingCreateDto) SetMarksAndNumbersNil(b bool)`

 SetMarksAndNumbersNil sets the value for MarksAndNumbers to be an explicit nil

### UnsetMarksAndNumbers
`func (o *BillOfLadingCreateDto) UnsetMarksAndNumbers()`

UnsetMarksAndNumbers ensures that no value is present for MarksAndNumbers, not even an explicit nil
### GetTotalPackages

`func (o *BillOfLadingCreateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *BillOfLadingCreateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *BillOfLadingCreateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *BillOfLadingCreateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### GetTotalGrossWeightKg

`func (o *BillOfLadingCreateDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *BillOfLadingCreateDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *BillOfLadingCreateDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *BillOfLadingCreateDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### GetTotalVolumeM3

`func (o *BillOfLadingCreateDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *BillOfLadingCreateDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *BillOfLadingCreateDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *BillOfLadingCreateDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *BillOfLadingCreateDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *BillOfLadingCreateDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


