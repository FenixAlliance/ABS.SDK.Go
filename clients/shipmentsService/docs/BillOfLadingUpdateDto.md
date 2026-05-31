# BillOfLadingUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**ExpiryDate** | Pointer to **NullableTime** |  | [optional] 
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

### NewBillOfLadingUpdateDto

`func NewBillOfLadingUpdateDto() *BillOfLadingUpdateDto`

NewBillOfLadingUpdateDto instantiates a new BillOfLadingUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfLadingUpdateDtoWithDefaults

`func NewBillOfLadingUpdateDtoWithDefaults() *BillOfLadingUpdateDto`

NewBillOfLadingUpdateDtoWithDefaults instantiates a new BillOfLadingUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillOfLadingNumber

`func (o *BillOfLadingUpdateDto) GetBillOfLadingNumber() string`

GetBillOfLadingNumber returns the BillOfLadingNumber field if non-nil, zero value otherwise.

### GetBillOfLadingNumberOk

`func (o *BillOfLadingUpdateDto) GetBillOfLadingNumberOk() (*string, bool)`

GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingNumber

`func (o *BillOfLadingUpdateDto) SetBillOfLadingNumber(v string)`

SetBillOfLadingNumber sets BillOfLadingNumber field to given value.

### HasBillOfLadingNumber

`func (o *BillOfLadingUpdateDto) HasBillOfLadingNumber() bool`

HasBillOfLadingNumber returns a boolean if a field has been set.

### SetBillOfLadingNumberNil

`func (o *BillOfLadingUpdateDto) SetBillOfLadingNumberNil(b bool)`

 SetBillOfLadingNumberNil sets the value for BillOfLadingNumber to be an explicit nil

### UnsetBillOfLadingNumber
`func (o *BillOfLadingUpdateDto) UnsetBillOfLadingNumber()`

UnsetBillOfLadingNumber ensures that no value is present for BillOfLadingNumber, not even an explicit nil
### GetTitle

`func (o *BillOfLadingUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillOfLadingUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillOfLadingUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BillOfLadingUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BillOfLadingUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BillOfLadingUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *BillOfLadingUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillOfLadingUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillOfLadingUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillOfLadingUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillOfLadingUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillOfLadingUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBillOfLadingType

`func (o *BillOfLadingUpdateDto) GetBillOfLadingType() string`

GetBillOfLadingType returns the BillOfLadingType field if non-nil, zero value otherwise.

### GetBillOfLadingTypeOk

`func (o *BillOfLadingUpdateDto) GetBillOfLadingTypeOk() (*string, bool)`

GetBillOfLadingTypeOk returns a tuple with the BillOfLadingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingType

`func (o *BillOfLadingUpdateDto) SetBillOfLadingType(v string)`

SetBillOfLadingType sets BillOfLadingType field to given value.

### HasBillOfLadingType

`func (o *BillOfLadingUpdateDto) HasBillOfLadingType() bool`

HasBillOfLadingType returns a boolean if a field has been set.

### SetBillOfLadingTypeNil

`func (o *BillOfLadingUpdateDto) SetBillOfLadingTypeNil(b bool)`

 SetBillOfLadingTypeNil sets the value for BillOfLadingType to be an explicit nil

### UnsetBillOfLadingType
`func (o *BillOfLadingUpdateDto) UnsetBillOfLadingType()`

UnsetBillOfLadingType ensures that no value is present for BillOfLadingType, not even an explicit nil
### GetIsNegotiable

`func (o *BillOfLadingUpdateDto) GetIsNegotiable() bool`

GetIsNegotiable returns the IsNegotiable field if non-nil, zero value otherwise.

### GetIsNegotiableOk

`func (o *BillOfLadingUpdateDto) GetIsNegotiableOk() (*bool, bool)`

GetIsNegotiableOk returns a tuple with the IsNegotiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNegotiable

`func (o *BillOfLadingUpdateDto) SetIsNegotiable(v bool)`

SetIsNegotiable sets IsNegotiable field to given value.

### HasIsNegotiable

`func (o *BillOfLadingUpdateDto) HasIsNegotiable() bool`

HasIsNegotiable returns a boolean if a field has been set.

### GetIsClean

`func (o *BillOfLadingUpdateDto) GetIsClean() bool`

GetIsClean returns the IsClean field if non-nil, zero value otherwise.

### GetIsCleanOk

`func (o *BillOfLadingUpdateDto) GetIsCleanOk() (*bool, bool)`

GetIsCleanOk returns a tuple with the IsClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClean

`func (o *BillOfLadingUpdateDto) SetIsClean(v bool)`

SetIsClean sets IsClean field to given value.

### HasIsClean

`func (o *BillOfLadingUpdateDto) HasIsClean() bool`

HasIsClean returns a boolean if a field has been set.

### GetNumberOfOriginals

`func (o *BillOfLadingUpdateDto) GetNumberOfOriginals() int32`

GetNumberOfOriginals returns the NumberOfOriginals field if non-nil, zero value otherwise.

### GetNumberOfOriginalsOk

`func (o *BillOfLadingUpdateDto) GetNumberOfOriginalsOk() (*int32, bool)`

GetNumberOfOriginalsOk returns a tuple with the NumberOfOriginals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOriginals

`func (o *BillOfLadingUpdateDto) SetNumberOfOriginals(v int32)`

SetNumberOfOriginals sets NumberOfOriginals field to given value.

### HasNumberOfOriginals

`func (o *BillOfLadingUpdateDto) HasNumberOfOriginals() bool`

HasNumberOfOriginals returns a boolean if a field has been set.

### GetFreightPaymentType

`func (o *BillOfLadingUpdateDto) GetFreightPaymentType() string`

GetFreightPaymentType returns the FreightPaymentType field if non-nil, zero value otherwise.

### GetFreightPaymentTypeOk

`func (o *BillOfLadingUpdateDto) GetFreightPaymentTypeOk() (*string, bool)`

GetFreightPaymentTypeOk returns a tuple with the FreightPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightPaymentType

`func (o *BillOfLadingUpdateDto) SetFreightPaymentType(v string)`

SetFreightPaymentType sets FreightPaymentType field to given value.

### HasFreightPaymentType

`func (o *BillOfLadingUpdateDto) HasFreightPaymentType() bool`

HasFreightPaymentType returns a boolean if a field has been set.

### SetFreightPaymentTypeNil

`func (o *BillOfLadingUpdateDto) SetFreightPaymentTypeNil(b bool)`

 SetFreightPaymentTypeNil sets the value for FreightPaymentType to be an explicit nil

### UnsetFreightPaymentType
`func (o *BillOfLadingUpdateDto) UnsetFreightPaymentType()`

UnsetFreightPaymentType ensures that no value is present for FreightPaymentType, not even an explicit nil
### GetShippingTerms

`func (o *BillOfLadingUpdateDto) GetShippingTerms() string`

GetShippingTerms returns the ShippingTerms field if non-nil, zero value otherwise.

### GetShippingTermsOk

`func (o *BillOfLadingUpdateDto) GetShippingTermsOk() (*string, bool)`

GetShippingTermsOk returns a tuple with the ShippingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTerms

`func (o *BillOfLadingUpdateDto) SetShippingTerms(v string)`

SetShippingTerms sets ShippingTerms field to given value.

### HasShippingTerms

`func (o *BillOfLadingUpdateDto) HasShippingTerms() bool`

HasShippingTerms returns a boolean if a field has been set.

### SetShippingTermsNil

`func (o *BillOfLadingUpdateDto) SetShippingTermsNil(b bool)`

 SetShippingTermsNil sets the value for ShippingTerms to be an explicit nil

### UnsetShippingTerms
`func (o *BillOfLadingUpdateDto) UnsetShippingTerms()`

UnsetShippingTerms ensures that no value is present for ShippingTerms, not even an explicit nil
### GetFreightChargesDescription

`func (o *BillOfLadingUpdateDto) GetFreightChargesDescription() string`

GetFreightChargesDescription returns the FreightChargesDescription field if non-nil, zero value otherwise.

### GetFreightChargesDescriptionOk

`func (o *BillOfLadingUpdateDto) GetFreightChargesDescriptionOk() (*string, bool)`

GetFreightChargesDescriptionOk returns a tuple with the FreightChargesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightChargesDescription

`func (o *BillOfLadingUpdateDto) SetFreightChargesDescription(v string)`

SetFreightChargesDescription sets FreightChargesDescription field to given value.

### HasFreightChargesDescription

`func (o *BillOfLadingUpdateDto) HasFreightChargesDescription() bool`

HasFreightChargesDescription returns a boolean if a field has been set.

### SetFreightChargesDescriptionNil

`func (o *BillOfLadingUpdateDto) SetFreightChargesDescriptionNil(b bool)`

 SetFreightChargesDescriptionNil sets the value for FreightChargesDescription to be an explicit nil

### UnsetFreightChargesDescription
`func (o *BillOfLadingUpdateDto) UnsetFreightChargesDescription()`

UnsetFreightChargesDescription ensures that no value is present for FreightChargesDescription, not even an explicit nil
### GetDeclaredValueAmount

`func (o *BillOfLadingUpdateDto) GetDeclaredValueAmount() float64`

GetDeclaredValueAmount returns the DeclaredValueAmount field if non-nil, zero value otherwise.

### GetDeclaredValueAmountOk

`func (o *BillOfLadingUpdateDto) GetDeclaredValueAmountOk() (*float64, bool)`

GetDeclaredValueAmountOk returns a tuple with the DeclaredValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueAmount

`func (o *BillOfLadingUpdateDto) SetDeclaredValueAmount(v float64)`

SetDeclaredValueAmount sets DeclaredValueAmount field to given value.

### HasDeclaredValueAmount

`func (o *BillOfLadingUpdateDto) HasDeclaredValueAmount() bool`

HasDeclaredValueAmount returns a boolean if a field has been set.

### GetDeclaredValueCurrencyId

`func (o *BillOfLadingUpdateDto) GetDeclaredValueCurrencyId() string`

GetDeclaredValueCurrencyId returns the DeclaredValueCurrencyId field if non-nil, zero value otherwise.

### GetDeclaredValueCurrencyIdOk

`func (o *BillOfLadingUpdateDto) GetDeclaredValueCurrencyIdOk() (*string, bool)`

GetDeclaredValueCurrencyIdOk returns a tuple with the DeclaredValueCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueCurrencyId

`func (o *BillOfLadingUpdateDto) SetDeclaredValueCurrencyId(v string)`

SetDeclaredValueCurrencyId sets DeclaredValueCurrencyId field to given value.

### HasDeclaredValueCurrencyId

`func (o *BillOfLadingUpdateDto) HasDeclaredValueCurrencyId() bool`

HasDeclaredValueCurrencyId returns a boolean if a field has been set.

### SetDeclaredValueCurrencyIdNil

`func (o *BillOfLadingUpdateDto) SetDeclaredValueCurrencyIdNil(b bool)`

 SetDeclaredValueCurrencyIdNil sets the value for DeclaredValueCurrencyId to be an explicit nil

### UnsetDeclaredValueCurrencyId
`func (o *BillOfLadingUpdateDto) UnsetDeclaredValueCurrencyId()`

UnsetDeclaredValueCurrencyId ensures that no value is present for DeclaredValueCurrencyId, not even an explicit nil
### GetExpiryDate

`func (o *BillOfLadingUpdateDto) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *BillOfLadingUpdateDto) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *BillOfLadingUpdateDto) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *BillOfLadingUpdateDto) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *BillOfLadingUpdateDto) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *BillOfLadingUpdateDto) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetVesselName

`func (o *BillOfLadingUpdateDto) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *BillOfLadingUpdateDto) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *BillOfLadingUpdateDto) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *BillOfLadingUpdateDto) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### SetVesselNameNil

`func (o *BillOfLadingUpdateDto) SetVesselNameNil(b bool)`

 SetVesselNameNil sets the value for VesselName to be an explicit nil

### UnsetVesselName
`func (o *BillOfLadingUpdateDto) UnsetVesselName()`

UnsetVesselName ensures that no value is present for VesselName, not even an explicit nil
### GetVoyageNumber

`func (o *BillOfLadingUpdateDto) GetVoyageNumber() string`

GetVoyageNumber returns the VoyageNumber field if non-nil, zero value otherwise.

### GetVoyageNumberOk

`func (o *BillOfLadingUpdateDto) GetVoyageNumberOk() (*string, bool)`

GetVoyageNumberOk returns a tuple with the VoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNumber

`func (o *BillOfLadingUpdateDto) SetVoyageNumber(v string)`

SetVoyageNumber sets VoyageNumber field to given value.

### HasVoyageNumber

`func (o *BillOfLadingUpdateDto) HasVoyageNumber() bool`

HasVoyageNumber returns a boolean if a field has been set.

### SetVoyageNumberNil

`func (o *BillOfLadingUpdateDto) SetVoyageNumberNil(b bool)`

 SetVoyageNumberNil sets the value for VoyageNumber to be an explicit nil

### UnsetVoyageNumber
`func (o *BillOfLadingUpdateDto) UnsetVoyageNumber()`

UnsetVoyageNumber ensures that no value is present for VoyageNumber, not even an explicit nil
### GetShipperContactId

`func (o *BillOfLadingUpdateDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *BillOfLadingUpdateDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *BillOfLadingUpdateDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *BillOfLadingUpdateDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *BillOfLadingUpdateDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *BillOfLadingUpdateDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *BillOfLadingUpdateDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *BillOfLadingUpdateDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *BillOfLadingUpdateDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *BillOfLadingUpdateDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *BillOfLadingUpdateDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *BillOfLadingUpdateDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetNotifyPartyContactId

`func (o *BillOfLadingUpdateDto) GetNotifyPartyContactId() string`

GetNotifyPartyContactId returns the NotifyPartyContactId field if non-nil, zero value otherwise.

### GetNotifyPartyContactIdOk

`func (o *BillOfLadingUpdateDto) GetNotifyPartyContactIdOk() (*string, bool)`

GetNotifyPartyContactIdOk returns a tuple with the NotifyPartyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPartyContactId

`func (o *BillOfLadingUpdateDto) SetNotifyPartyContactId(v string)`

SetNotifyPartyContactId sets NotifyPartyContactId field to given value.

### HasNotifyPartyContactId

`func (o *BillOfLadingUpdateDto) HasNotifyPartyContactId() bool`

HasNotifyPartyContactId returns a boolean if a field has been set.

### SetNotifyPartyContactIdNil

`func (o *BillOfLadingUpdateDto) SetNotifyPartyContactIdNil(b bool)`

 SetNotifyPartyContactIdNil sets the value for NotifyPartyContactId to be an explicit nil

### UnsetNotifyPartyContactId
`func (o *BillOfLadingUpdateDto) UnsetNotifyPartyContactId()`

UnsetNotifyPartyContactId ensures that no value is present for NotifyPartyContactId, not even an explicit nil
### GetShippingCourierId

`func (o *BillOfLadingUpdateDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *BillOfLadingUpdateDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *BillOfLadingUpdateDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *BillOfLadingUpdateDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *BillOfLadingUpdateDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *BillOfLadingUpdateDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil
### GetPortOfLoadingId

`func (o *BillOfLadingUpdateDto) GetPortOfLoadingId() string`

GetPortOfLoadingId returns the PortOfLoadingId field if non-nil, zero value otherwise.

### GetPortOfLoadingIdOk

`func (o *BillOfLadingUpdateDto) GetPortOfLoadingIdOk() (*string, bool)`

GetPortOfLoadingIdOk returns a tuple with the PortOfLoadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoadingId

`func (o *BillOfLadingUpdateDto) SetPortOfLoadingId(v string)`

SetPortOfLoadingId sets PortOfLoadingId field to given value.

### HasPortOfLoadingId

`func (o *BillOfLadingUpdateDto) HasPortOfLoadingId() bool`

HasPortOfLoadingId returns a boolean if a field has been set.

### SetPortOfLoadingIdNil

`func (o *BillOfLadingUpdateDto) SetPortOfLoadingIdNil(b bool)`

 SetPortOfLoadingIdNil sets the value for PortOfLoadingId to be an explicit nil

### UnsetPortOfLoadingId
`func (o *BillOfLadingUpdateDto) UnsetPortOfLoadingId()`

UnsetPortOfLoadingId ensures that no value is present for PortOfLoadingId, not even an explicit nil
### GetPortOfDischargeId

`func (o *BillOfLadingUpdateDto) GetPortOfDischargeId() string`

GetPortOfDischargeId returns the PortOfDischargeId field if non-nil, zero value otherwise.

### GetPortOfDischargeIdOk

`func (o *BillOfLadingUpdateDto) GetPortOfDischargeIdOk() (*string, bool)`

GetPortOfDischargeIdOk returns a tuple with the PortOfDischargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfDischargeId

`func (o *BillOfLadingUpdateDto) SetPortOfDischargeId(v string)`

SetPortOfDischargeId sets PortOfDischargeId field to given value.

### HasPortOfDischargeId

`func (o *BillOfLadingUpdateDto) HasPortOfDischargeId() bool`

HasPortOfDischargeId returns a boolean if a field has been set.

### SetPortOfDischargeIdNil

`func (o *BillOfLadingUpdateDto) SetPortOfDischargeIdNil(b bool)`

 SetPortOfDischargeIdNil sets the value for PortOfDischargeId to be an explicit nil

### UnsetPortOfDischargeId
`func (o *BillOfLadingUpdateDto) UnsetPortOfDischargeId()`

UnsetPortOfDischargeId ensures that no value is present for PortOfDischargeId, not even an explicit nil
### GetPlaceOfReceiptId

`func (o *BillOfLadingUpdateDto) GetPlaceOfReceiptId() string`

GetPlaceOfReceiptId returns the PlaceOfReceiptId field if non-nil, zero value otherwise.

### GetPlaceOfReceiptIdOk

`func (o *BillOfLadingUpdateDto) GetPlaceOfReceiptIdOk() (*string, bool)`

GetPlaceOfReceiptIdOk returns a tuple with the PlaceOfReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceiptId

`func (o *BillOfLadingUpdateDto) SetPlaceOfReceiptId(v string)`

SetPlaceOfReceiptId sets PlaceOfReceiptId field to given value.

### HasPlaceOfReceiptId

`func (o *BillOfLadingUpdateDto) HasPlaceOfReceiptId() bool`

HasPlaceOfReceiptId returns a boolean if a field has been set.

### SetPlaceOfReceiptIdNil

`func (o *BillOfLadingUpdateDto) SetPlaceOfReceiptIdNil(b bool)`

 SetPlaceOfReceiptIdNil sets the value for PlaceOfReceiptId to be an explicit nil

### UnsetPlaceOfReceiptId
`func (o *BillOfLadingUpdateDto) UnsetPlaceOfReceiptId()`

UnsetPlaceOfReceiptId ensures that no value is present for PlaceOfReceiptId, not even an explicit nil
### GetPlaceOfDeliveryId

`func (o *BillOfLadingUpdateDto) GetPlaceOfDeliveryId() string`

GetPlaceOfDeliveryId returns the PlaceOfDeliveryId field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryIdOk

`func (o *BillOfLadingUpdateDto) GetPlaceOfDeliveryIdOk() (*string, bool)`

GetPlaceOfDeliveryIdOk returns a tuple with the PlaceOfDeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryId

`func (o *BillOfLadingUpdateDto) SetPlaceOfDeliveryId(v string)`

SetPlaceOfDeliveryId sets PlaceOfDeliveryId field to given value.

### HasPlaceOfDeliveryId

`func (o *BillOfLadingUpdateDto) HasPlaceOfDeliveryId() bool`

HasPlaceOfDeliveryId returns a boolean if a field has been set.

### SetPlaceOfDeliveryIdNil

`func (o *BillOfLadingUpdateDto) SetPlaceOfDeliveryIdNil(b bool)`

 SetPlaceOfDeliveryIdNil sets the value for PlaceOfDeliveryId to be an explicit nil

### UnsetPlaceOfDeliveryId
`func (o *BillOfLadingUpdateDto) UnsetPlaceOfDeliveryId()`

UnsetPlaceOfDeliveryId ensures that no value is present for PlaceOfDeliveryId, not even an explicit nil
### GetShipmentId

`func (o *BillOfLadingUpdateDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *BillOfLadingUpdateDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *BillOfLadingUpdateDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *BillOfLadingUpdateDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *BillOfLadingUpdateDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *BillOfLadingUpdateDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetOrderId

`func (o *BillOfLadingUpdateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BillOfLadingUpdateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BillOfLadingUpdateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *BillOfLadingUpdateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *BillOfLadingUpdateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *BillOfLadingUpdateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetVoyageId

`func (o *BillOfLadingUpdateDto) GetVoyageId() string`

GetVoyageId returns the VoyageId field if non-nil, zero value otherwise.

### GetVoyageIdOk

`func (o *BillOfLadingUpdateDto) GetVoyageIdOk() (*string, bool)`

GetVoyageIdOk returns a tuple with the VoyageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageId

`func (o *BillOfLadingUpdateDto) SetVoyageId(v string)`

SetVoyageId sets VoyageId field to given value.

### HasVoyageId

`func (o *BillOfLadingUpdateDto) HasVoyageId() bool`

HasVoyageId returns a boolean if a field has been set.

### SetVoyageIdNil

`func (o *BillOfLadingUpdateDto) SetVoyageIdNil(b bool)`

 SetVoyageIdNil sets the value for VoyageId to be an explicit nil

### UnsetVoyageId
`func (o *BillOfLadingUpdateDto) UnsetVoyageId()`

UnsetVoyageId ensures that no value is present for VoyageId, not even an explicit nil
### GetMarksAndNumbers

`func (o *BillOfLadingUpdateDto) GetMarksAndNumbers() string`

GetMarksAndNumbers returns the MarksAndNumbers field if non-nil, zero value otherwise.

### GetMarksAndNumbersOk

`func (o *BillOfLadingUpdateDto) GetMarksAndNumbersOk() (*string, bool)`

GetMarksAndNumbersOk returns a tuple with the MarksAndNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarksAndNumbers

`func (o *BillOfLadingUpdateDto) SetMarksAndNumbers(v string)`

SetMarksAndNumbers sets MarksAndNumbers field to given value.

### HasMarksAndNumbers

`func (o *BillOfLadingUpdateDto) HasMarksAndNumbers() bool`

HasMarksAndNumbers returns a boolean if a field has been set.

### SetMarksAndNumbersNil

`func (o *BillOfLadingUpdateDto) SetMarksAndNumbersNil(b bool)`

 SetMarksAndNumbersNil sets the value for MarksAndNumbers to be an explicit nil

### UnsetMarksAndNumbers
`func (o *BillOfLadingUpdateDto) UnsetMarksAndNumbers()`

UnsetMarksAndNumbers ensures that no value is present for MarksAndNumbers, not even an explicit nil
### GetTotalPackages

`func (o *BillOfLadingUpdateDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *BillOfLadingUpdateDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *BillOfLadingUpdateDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *BillOfLadingUpdateDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### GetTotalGrossWeightKg

`func (o *BillOfLadingUpdateDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *BillOfLadingUpdateDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *BillOfLadingUpdateDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *BillOfLadingUpdateDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### GetTotalVolumeM3

`func (o *BillOfLadingUpdateDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *BillOfLadingUpdateDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *BillOfLadingUpdateDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *BillOfLadingUpdateDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *BillOfLadingUpdateDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *BillOfLadingUpdateDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


