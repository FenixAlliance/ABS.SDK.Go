# BillOfLadingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**BillOfLadingNumber** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BillOfLadingType** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**IsNegotiable** | Pointer to **bool** |  | [optional] 
**IsClean** | Pointer to **bool** |  | [optional] 
**NumberOfOriginals** | Pointer to **int32** |  | [optional] 
**FreightPaymentType** | Pointer to **NullableString** |  | [optional] 
**ShippingTerms** | Pointer to **NullableString** |  | [optional] 
**FreightChargesDescription** | Pointer to **NullableString** |  | [optional] 
**DeclaredValueAmount** | Pointer to **float64** |  | [optional] 
**DeclaredValueCurrencyId** | Pointer to **NullableString** |  | [optional] 
**IssuedDate** | Pointer to **NullableTime** |  | [optional] 
**OnBoardDate** | Pointer to **NullableTime** |  | [optional] 
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
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBillOfLadingDto

`func NewBillOfLadingDto() *BillOfLadingDto`

NewBillOfLadingDto instantiates a new BillOfLadingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfLadingDtoWithDefaults

`func NewBillOfLadingDtoWithDefaults() *BillOfLadingDto`

NewBillOfLadingDtoWithDefaults instantiates a new BillOfLadingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillOfLadingDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillOfLadingDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillOfLadingDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillOfLadingDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BillOfLadingDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BillOfLadingDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BillOfLadingDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BillOfLadingDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BillOfLadingDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BillOfLadingDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BillOfLadingDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BillOfLadingDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetBillOfLadingNumber

`func (o *BillOfLadingDto) GetBillOfLadingNumber() string`

GetBillOfLadingNumber returns the BillOfLadingNumber field if non-nil, zero value otherwise.

### GetBillOfLadingNumberOk

`func (o *BillOfLadingDto) GetBillOfLadingNumberOk() (*string, bool)`

GetBillOfLadingNumberOk returns a tuple with the BillOfLadingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingNumber

`func (o *BillOfLadingDto) SetBillOfLadingNumber(v string)`

SetBillOfLadingNumber sets BillOfLadingNumber field to given value.

### HasBillOfLadingNumber

`func (o *BillOfLadingDto) HasBillOfLadingNumber() bool`

HasBillOfLadingNumber returns a boolean if a field has been set.

### SetBillOfLadingNumberNil

`func (o *BillOfLadingDto) SetBillOfLadingNumberNil(b bool)`

 SetBillOfLadingNumberNil sets the value for BillOfLadingNumber to be an explicit nil

### UnsetBillOfLadingNumber
`func (o *BillOfLadingDto) UnsetBillOfLadingNumber()`

UnsetBillOfLadingNumber ensures that no value is present for BillOfLadingNumber, not even an explicit nil
### GetTitle

`func (o *BillOfLadingDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillOfLadingDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillOfLadingDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BillOfLadingDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BillOfLadingDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BillOfLadingDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *BillOfLadingDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillOfLadingDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillOfLadingDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillOfLadingDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BillOfLadingDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BillOfLadingDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBillOfLadingType

`func (o *BillOfLadingDto) GetBillOfLadingType() string`

GetBillOfLadingType returns the BillOfLadingType field if non-nil, zero value otherwise.

### GetBillOfLadingTypeOk

`func (o *BillOfLadingDto) GetBillOfLadingTypeOk() (*string, bool)`

GetBillOfLadingTypeOk returns a tuple with the BillOfLadingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfLadingType

`func (o *BillOfLadingDto) SetBillOfLadingType(v string)`

SetBillOfLadingType sets BillOfLadingType field to given value.

### HasBillOfLadingType

`func (o *BillOfLadingDto) HasBillOfLadingType() bool`

HasBillOfLadingType returns a boolean if a field has been set.

### SetBillOfLadingTypeNil

`func (o *BillOfLadingDto) SetBillOfLadingTypeNil(b bool)`

 SetBillOfLadingTypeNil sets the value for BillOfLadingType to be an explicit nil

### UnsetBillOfLadingType
`func (o *BillOfLadingDto) UnsetBillOfLadingType()`

UnsetBillOfLadingType ensures that no value is present for BillOfLadingType, not even an explicit nil
### GetStatus

`func (o *BillOfLadingDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillOfLadingDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillOfLadingDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillOfLadingDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BillOfLadingDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BillOfLadingDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsNegotiable

`func (o *BillOfLadingDto) GetIsNegotiable() bool`

GetIsNegotiable returns the IsNegotiable field if non-nil, zero value otherwise.

### GetIsNegotiableOk

`func (o *BillOfLadingDto) GetIsNegotiableOk() (*bool, bool)`

GetIsNegotiableOk returns a tuple with the IsNegotiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNegotiable

`func (o *BillOfLadingDto) SetIsNegotiable(v bool)`

SetIsNegotiable sets IsNegotiable field to given value.

### HasIsNegotiable

`func (o *BillOfLadingDto) HasIsNegotiable() bool`

HasIsNegotiable returns a boolean if a field has been set.

### GetIsClean

`func (o *BillOfLadingDto) GetIsClean() bool`

GetIsClean returns the IsClean field if non-nil, zero value otherwise.

### GetIsCleanOk

`func (o *BillOfLadingDto) GetIsCleanOk() (*bool, bool)`

GetIsCleanOk returns a tuple with the IsClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClean

`func (o *BillOfLadingDto) SetIsClean(v bool)`

SetIsClean sets IsClean field to given value.

### HasIsClean

`func (o *BillOfLadingDto) HasIsClean() bool`

HasIsClean returns a boolean if a field has been set.

### GetNumberOfOriginals

`func (o *BillOfLadingDto) GetNumberOfOriginals() int32`

GetNumberOfOriginals returns the NumberOfOriginals field if non-nil, zero value otherwise.

### GetNumberOfOriginalsOk

`func (o *BillOfLadingDto) GetNumberOfOriginalsOk() (*int32, bool)`

GetNumberOfOriginalsOk returns a tuple with the NumberOfOriginals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfOriginals

`func (o *BillOfLadingDto) SetNumberOfOriginals(v int32)`

SetNumberOfOriginals sets NumberOfOriginals field to given value.

### HasNumberOfOriginals

`func (o *BillOfLadingDto) HasNumberOfOriginals() bool`

HasNumberOfOriginals returns a boolean if a field has been set.

### GetFreightPaymentType

`func (o *BillOfLadingDto) GetFreightPaymentType() string`

GetFreightPaymentType returns the FreightPaymentType field if non-nil, zero value otherwise.

### GetFreightPaymentTypeOk

`func (o *BillOfLadingDto) GetFreightPaymentTypeOk() (*string, bool)`

GetFreightPaymentTypeOk returns a tuple with the FreightPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightPaymentType

`func (o *BillOfLadingDto) SetFreightPaymentType(v string)`

SetFreightPaymentType sets FreightPaymentType field to given value.

### HasFreightPaymentType

`func (o *BillOfLadingDto) HasFreightPaymentType() bool`

HasFreightPaymentType returns a boolean if a field has been set.

### SetFreightPaymentTypeNil

`func (o *BillOfLadingDto) SetFreightPaymentTypeNil(b bool)`

 SetFreightPaymentTypeNil sets the value for FreightPaymentType to be an explicit nil

### UnsetFreightPaymentType
`func (o *BillOfLadingDto) UnsetFreightPaymentType()`

UnsetFreightPaymentType ensures that no value is present for FreightPaymentType, not even an explicit nil
### GetShippingTerms

`func (o *BillOfLadingDto) GetShippingTerms() string`

GetShippingTerms returns the ShippingTerms field if non-nil, zero value otherwise.

### GetShippingTermsOk

`func (o *BillOfLadingDto) GetShippingTermsOk() (*string, bool)`

GetShippingTermsOk returns a tuple with the ShippingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingTerms

`func (o *BillOfLadingDto) SetShippingTerms(v string)`

SetShippingTerms sets ShippingTerms field to given value.

### HasShippingTerms

`func (o *BillOfLadingDto) HasShippingTerms() bool`

HasShippingTerms returns a boolean if a field has been set.

### SetShippingTermsNil

`func (o *BillOfLadingDto) SetShippingTermsNil(b bool)`

 SetShippingTermsNil sets the value for ShippingTerms to be an explicit nil

### UnsetShippingTerms
`func (o *BillOfLadingDto) UnsetShippingTerms()`

UnsetShippingTerms ensures that no value is present for ShippingTerms, not even an explicit nil
### GetFreightChargesDescription

`func (o *BillOfLadingDto) GetFreightChargesDescription() string`

GetFreightChargesDescription returns the FreightChargesDescription field if non-nil, zero value otherwise.

### GetFreightChargesDescriptionOk

`func (o *BillOfLadingDto) GetFreightChargesDescriptionOk() (*string, bool)`

GetFreightChargesDescriptionOk returns a tuple with the FreightChargesDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightChargesDescription

`func (o *BillOfLadingDto) SetFreightChargesDescription(v string)`

SetFreightChargesDescription sets FreightChargesDescription field to given value.

### HasFreightChargesDescription

`func (o *BillOfLadingDto) HasFreightChargesDescription() bool`

HasFreightChargesDescription returns a boolean if a field has been set.

### SetFreightChargesDescriptionNil

`func (o *BillOfLadingDto) SetFreightChargesDescriptionNil(b bool)`

 SetFreightChargesDescriptionNil sets the value for FreightChargesDescription to be an explicit nil

### UnsetFreightChargesDescription
`func (o *BillOfLadingDto) UnsetFreightChargesDescription()`

UnsetFreightChargesDescription ensures that no value is present for FreightChargesDescription, not even an explicit nil
### GetDeclaredValueAmount

`func (o *BillOfLadingDto) GetDeclaredValueAmount() float64`

GetDeclaredValueAmount returns the DeclaredValueAmount field if non-nil, zero value otherwise.

### GetDeclaredValueAmountOk

`func (o *BillOfLadingDto) GetDeclaredValueAmountOk() (*float64, bool)`

GetDeclaredValueAmountOk returns a tuple with the DeclaredValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueAmount

`func (o *BillOfLadingDto) SetDeclaredValueAmount(v float64)`

SetDeclaredValueAmount sets DeclaredValueAmount field to given value.

### HasDeclaredValueAmount

`func (o *BillOfLadingDto) HasDeclaredValueAmount() bool`

HasDeclaredValueAmount returns a boolean if a field has been set.

### GetDeclaredValueCurrencyId

`func (o *BillOfLadingDto) GetDeclaredValueCurrencyId() string`

GetDeclaredValueCurrencyId returns the DeclaredValueCurrencyId field if non-nil, zero value otherwise.

### GetDeclaredValueCurrencyIdOk

`func (o *BillOfLadingDto) GetDeclaredValueCurrencyIdOk() (*string, bool)`

GetDeclaredValueCurrencyIdOk returns a tuple with the DeclaredValueCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValueCurrencyId

`func (o *BillOfLadingDto) SetDeclaredValueCurrencyId(v string)`

SetDeclaredValueCurrencyId sets DeclaredValueCurrencyId field to given value.

### HasDeclaredValueCurrencyId

`func (o *BillOfLadingDto) HasDeclaredValueCurrencyId() bool`

HasDeclaredValueCurrencyId returns a boolean if a field has been set.

### SetDeclaredValueCurrencyIdNil

`func (o *BillOfLadingDto) SetDeclaredValueCurrencyIdNil(b bool)`

 SetDeclaredValueCurrencyIdNil sets the value for DeclaredValueCurrencyId to be an explicit nil

### UnsetDeclaredValueCurrencyId
`func (o *BillOfLadingDto) UnsetDeclaredValueCurrencyId()`

UnsetDeclaredValueCurrencyId ensures that no value is present for DeclaredValueCurrencyId, not even an explicit nil
### GetIssuedDate

`func (o *BillOfLadingDto) GetIssuedDate() time.Time`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *BillOfLadingDto) GetIssuedDateOk() (*time.Time, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *BillOfLadingDto) SetIssuedDate(v time.Time)`

SetIssuedDate sets IssuedDate field to given value.

### HasIssuedDate

`func (o *BillOfLadingDto) HasIssuedDate() bool`

HasIssuedDate returns a boolean if a field has been set.

### SetIssuedDateNil

`func (o *BillOfLadingDto) SetIssuedDateNil(b bool)`

 SetIssuedDateNil sets the value for IssuedDate to be an explicit nil

### UnsetIssuedDate
`func (o *BillOfLadingDto) UnsetIssuedDate()`

UnsetIssuedDate ensures that no value is present for IssuedDate, not even an explicit nil
### GetOnBoardDate

`func (o *BillOfLadingDto) GetOnBoardDate() time.Time`

GetOnBoardDate returns the OnBoardDate field if non-nil, zero value otherwise.

### GetOnBoardDateOk

`func (o *BillOfLadingDto) GetOnBoardDateOk() (*time.Time, bool)`

GetOnBoardDateOk returns a tuple with the OnBoardDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBoardDate

`func (o *BillOfLadingDto) SetOnBoardDate(v time.Time)`

SetOnBoardDate sets OnBoardDate field to given value.

### HasOnBoardDate

`func (o *BillOfLadingDto) HasOnBoardDate() bool`

HasOnBoardDate returns a boolean if a field has been set.

### SetOnBoardDateNil

`func (o *BillOfLadingDto) SetOnBoardDateNil(b bool)`

 SetOnBoardDateNil sets the value for OnBoardDate to be an explicit nil

### UnsetOnBoardDate
`func (o *BillOfLadingDto) UnsetOnBoardDate()`

UnsetOnBoardDate ensures that no value is present for OnBoardDate, not even an explicit nil
### GetExpiryDate

`func (o *BillOfLadingDto) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *BillOfLadingDto) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *BillOfLadingDto) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *BillOfLadingDto) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *BillOfLadingDto) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *BillOfLadingDto) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetVesselName

`func (o *BillOfLadingDto) GetVesselName() string`

GetVesselName returns the VesselName field if non-nil, zero value otherwise.

### GetVesselNameOk

`func (o *BillOfLadingDto) GetVesselNameOk() (*string, bool)`

GetVesselNameOk returns a tuple with the VesselName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselName

`func (o *BillOfLadingDto) SetVesselName(v string)`

SetVesselName sets VesselName field to given value.

### HasVesselName

`func (o *BillOfLadingDto) HasVesselName() bool`

HasVesselName returns a boolean if a field has been set.

### SetVesselNameNil

`func (o *BillOfLadingDto) SetVesselNameNil(b bool)`

 SetVesselNameNil sets the value for VesselName to be an explicit nil

### UnsetVesselName
`func (o *BillOfLadingDto) UnsetVesselName()`

UnsetVesselName ensures that no value is present for VesselName, not even an explicit nil
### GetVoyageNumber

`func (o *BillOfLadingDto) GetVoyageNumber() string`

GetVoyageNumber returns the VoyageNumber field if non-nil, zero value otherwise.

### GetVoyageNumberOk

`func (o *BillOfLadingDto) GetVoyageNumberOk() (*string, bool)`

GetVoyageNumberOk returns a tuple with the VoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageNumber

`func (o *BillOfLadingDto) SetVoyageNumber(v string)`

SetVoyageNumber sets VoyageNumber field to given value.

### HasVoyageNumber

`func (o *BillOfLadingDto) HasVoyageNumber() bool`

HasVoyageNumber returns a boolean if a field has been set.

### SetVoyageNumberNil

`func (o *BillOfLadingDto) SetVoyageNumberNil(b bool)`

 SetVoyageNumberNil sets the value for VoyageNumber to be an explicit nil

### UnsetVoyageNumber
`func (o *BillOfLadingDto) UnsetVoyageNumber()`

UnsetVoyageNumber ensures that no value is present for VoyageNumber, not even an explicit nil
### GetShipperContactId

`func (o *BillOfLadingDto) GetShipperContactId() string`

GetShipperContactId returns the ShipperContactId field if non-nil, zero value otherwise.

### GetShipperContactIdOk

`func (o *BillOfLadingDto) GetShipperContactIdOk() (*string, bool)`

GetShipperContactIdOk returns a tuple with the ShipperContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipperContactId

`func (o *BillOfLadingDto) SetShipperContactId(v string)`

SetShipperContactId sets ShipperContactId field to given value.

### HasShipperContactId

`func (o *BillOfLadingDto) HasShipperContactId() bool`

HasShipperContactId returns a boolean if a field has been set.

### SetShipperContactIdNil

`func (o *BillOfLadingDto) SetShipperContactIdNil(b bool)`

 SetShipperContactIdNil sets the value for ShipperContactId to be an explicit nil

### UnsetShipperContactId
`func (o *BillOfLadingDto) UnsetShipperContactId()`

UnsetShipperContactId ensures that no value is present for ShipperContactId, not even an explicit nil
### GetConsigneeContactId

`func (o *BillOfLadingDto) GetConsigneeContactId() string`

GetConsigneeContactId returns the ConsigneeContactId field if non-nil, zero value otherwise.

### GetConsigneeContactIdOk

`func (o *BillOfLadingDto) GetConsigneeContactIdOk() (*string, bool)`

GetConsigneeContactIdOk returns a tuple with the ConsigneeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeContactId

`func (o *BillOfLadingDto) SetConsigneeContactId(v string)`

SetConsigneeContactId sets ConsigneeContactId field to given value.

### HasConsigneeContactId

`func (o *BillOfLadingDto) HasConsigneeContactId() bool`

HasConsigneeContactId returns a boolean if a field has been set.

### SetConsigneeContactIdNil

`func (o *BillOfLadingDto) SetConsigneeContactIdNil(b bool)`

 SetConsigneeContactIdNil sets the value for ConsigneeContactId to be an explicit nil

### UnsetConsigneeContactId
`func (o *BillOfLadingDto) UnsetConsigneeContactId()`

UnsetConsigneeContactId ensures that no value is present for ConsigneeContactId, not even an explicit nil
### GetNotifyPartyContactId

`func (o *BillOfLadingDto) GetNotifyPartyContactId() string`

GetNotifyPartyContactId returns the NotifyPartyContactId field if non-nil, zero value otherwise.

### GetNotifyPartyContactIdOk

`func (o *BillOfLadingDto) GetNotifyPartyContactIdOk() (*string, bool)`

GetNotifyPartyContactIdOk returns a tuple with the NotifyPartyContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyPartyContactId

`func (o *BillOfLadingDto) SetNotifyPartyContactId(v string)`

SetNotifyPartyContactId sets NotifyPartyContactId field to given value.

### HasNotifyPartyContactId

`func (o *BillOfLadingDto) HasNotifyPartyContactId() bool`

HasNotifyPartyContactId returns a boolean if a field has been set.

### SetNotifyPartyContactIdNil

`func (o *BillOfLadingDto) SetNotifyPartyContactIdNil(b bool)`

 SetNotifyPartyContactIdNil sets the value for NotifyPartyContactId to be an explicit nil

### UnsetNotifyPartyContactId
`func (o *BillOfLadingDto) UnsetNotifyPartyContactId()`

UnsetNotifyPartyContactId ensures that no value is present for NotifyPartyContactId, not even an explicit nil
### GetShippingCourierId

`func (o *BillOfLadingDto) GetShippingCourierId() string`

GetShippingCourierId returns the ShippingCourierId field if non-nil, zero value otherwise.

### GetShippingCourierIdOk

`func (o *BillOfLadingDto) GetShippingCourierIdOk() (*string, bool)`

GetShippingCourierIdOk returns a tuple with the ShippingCourierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCourierId

`func (o *BillOfLadingDto) SetShippingCourierId(v string)`

SetShippingCourierId sets ShippingCourierId field to given value.

### HasShippingCourierId

`func (o *BillOfLadingDto) HasShippingCourierId() bool`

HasShippingCourierId returns a boolean if a field has been set.

### SetShippingCourierIdNil

`func (o *BillOfLadingDto) SetShippingCourierIdNil(b bool)`

 SetShippingCourierIdNil sets the value for ShippingCourierId to be an explicit nil

### UnsetShippingCourierId
`func (o *BillOfLadingDto) UnsetShippingCourierId()`

UnsetShippingCourierId ensures that no value is present for ShippingCourierId, not even an explicit nil
### GetPortOfLoadingId

`func (o *BillOfLadingDto) GetPortOfLoadingId() string`

GetPortOfLoadingId returns the PortOfLoadingId field if non-nil, zero value otherwise.

### GetPortOfLoadingIdOk

`func (o *BillOfLadingDto) GetPortOfLoadingIdOk() (*string, bool)`

GetPortOfLoadingIdOk returns a tuple with the PortOfLoadingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoadingId

`func (o *BillOfLadingDto) SetPortOfLoadingId(v string)`

SetPortOfLoadingId sets PortOfLoadingId field to given value.

### HasPortOfLoadingId

`func (o *BillOfLadingDto) HasPortOfLoadingId() bool`

HasPortOfLoadingId returns a boolean if a field has been set.

### SetPortOfLoadingIdNil

`func (o *BillOfLadingDto) SetPortOfLoadingIdNil(b bool)`

 SetPortOfLoadingIdNil sets the value for PortOfLoadingId to be an explicit nil

### UnsetPortOfLoadingId
`func (o *BillOfLadingDto) UnsetPortOfLoadingId()`

UnsetPortOfLoadingId ensures that no value is present for PortOfLoadingId, not even an explicit nil
### GetPortOfDischargeId

`func (o *BillOfLadingDto) GetPortOfDischargeId() string`

GetPortOfDischargeId returns the PortOfDischargeId field if non-nil, zero value otherwise.

### GetPortOfDischargeIdOk

`func (o *BillOfLadingDto) GetPortOfDischargeIdOk() (*string, bool)`

GetPortOfDischargeIdOk returns a tuple with the PortOfDischargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfDischargeId

`func (o *BillOfLadingDto) SetPortOfDischargeId(v string)`

SetPortOfDischargeId sets PortOfDischargeId field to given value.

### HasPortOfDischargeId

`func (o *BillOfLadingDto) HasPortOfDischargeId() bool`

HasPortOfDischargeId returns a boolean if a field has been set.

### SetPortOfDischargeIdNil

`func (o *BillOfLadingDto) SetPortOfDischargeIdNil(b bool)`

 SetPortOfDischargeIdNil sets the value for PortOfDischargeId to be an explicit nil

### UnsetPortOfDischargeId
`func (o *BillOfLadingDto) UnsetPortOfDischargeId()`

UnsetPortOfDischargeId ensures that no value is present for PortOfDischargeId, not even an explicit nil
### GetPlaceOfReceiptId

`func (o *BillOfLadingDto) GetPlaceOfReceiptId() string`

GetPlaceOfReceiptId returns the PlaceOfReceiptId field if non-nil, zero value otherwise.

### GetPlaceOfReceiptIdOk

`func (o *BillOfLadingDto) GetPlaceOfReceiptIdOk() (*string, bool)`

GetPlaceOfReceiptIdOk returns a tuple with the PlaceOfReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfReceiptId

`func (o *BillOfLadingDto) SetPlaceOfReceiptId(v string)`

SetPlaceOfReceiptId sets PlaceOfReceiptId field to given value.

### HasPlaceOfReceiptId

`func (o *BillOfLadingDto) HasPlaceOfReceiptId() bool`

HasPlaceOfReceiptId returns a boolean if a field has been set.

### SetPlaceOfReceiptIdNil

`func (o *BillOfLadingDto) SetPlaceOfReceiptIdNil(b bool)`

 SetPlaceOfReceiptIdNil sets the value for PlaceOfReceiptId to be an explicit nil

### UnsetPlaceOfReceiptId
`func (o *BillOfLadingDto) UnsetPlaceOfReceiptId()`

UnsetPlaceOfReceiptId ensures that no value is present for PlaceOfReceiptId, not even an explicit nil
### GetPlaceOfDeliveryId

`func (o *BillOfLadingDto) GetPlaceOfDeliveryId() string`

GetPlaceOfDeliveryId returns the PlaceOfDeliveryId field if non-nil, zero value otherwise.

### GetPlaceOfDeliveryIdOk

`func (o *BillOfLadingDto) GetPlaceOfDeliveryIdOk() (*string, bool)`

GetPlaceOfDeliveryIdOk returns a tuple with the PlaceOfDeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfDeliveryId

`func (o *BillOfLadingDto) SetPlaceOfDeliveryId(v string)`

SetPlaceOfDeliveryId sets PlaceOfDeliveryId field to given value.

### HasPlaceOfDeliveryId

`func (o *BillOfLadingDto) HasPlaceOfDeliveryId() bool`

HasPlaceOfDeliveryId returns a boolean if a field has been set.

### SetPlaceOfDeliveryIdNil

`func (o *BillOfLadingDto) SetPlaceOfDeliveryIdNil(b bool)`

 SetPlaceOfDeliveryIdNil sets the value for PlaceOfDeliveryId to be an explicit nil

### UnsetPlaceOfDeliveryId
`func (o *BillOfLadingDto) UnsetPlaceOfDeliveryId()`

UnsetPlaceOfDeliveryId ensures that no value is present for PlaceOfDeliveryId, not even an explicit nil
### GetShipmentId

`func (o *BillOfLadingDto) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *BillOfLadingDto) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *BillOfLadingDto) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *BillOfLadingDto) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### SetShipmentIdNil

`func (o *BillOfLadingDto) SetShipmentIdNil(b bool)`

 SetShipmentIdNil sets the value for ShipmentId to be an explicit nil

### UnsetShipmentId
`func (o *BillOfLadingDto) UnsetShipmentId()`

UnsetShipmentId ensures that no value is present for ShipmentId, not even an explicit nil
### GetOrderId

`func (o *BillOfLadingDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BillOfLadingDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BillOfLadingDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *BillOfLadingDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *BillOfLadingDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *BillOfLadingDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetVoyageId

`func (o *BillOfLadingDto) GetVoyageId() string`

GetVoyageId returns the VoyageId field if non-nil, zero value otherwise.

### GetVoyageIdOk

`func (o *BillOfLadingDto) GetVoyageIdOk() (*string, bool)`

GetVoyageIdOk returns a tuple with the VoyageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageId

`func (o *BillOfLadingDto) SetVoyageId(v string)`

SetVoyageId sets VoyageId field to given value.

### HasVoyageId

`func (o *BillOfLadingDto) HasVoyageId() bool`

HasVoyageId returns a boolean if a field has been set.

### SetVoyageIdNil

`func (o *BillOfLadingDto) SetVoyageIdNil(b bool)`

 SetVoyageIdNil sets the value for VoyageId to be an explicit nil

### UnsetVoyageId
`func (o *BillOfLadingDto) UnsetVoyageId()`

UnsetVoyageId ensures that no value is present for VoyageId, not even an explicit nil
### GetMarksAndNumbers

`func (o *BillOfLadingDto) GetMarksAndNumbers() string`

GetMarksAndNumbers returns the MarksAndNumbers field if non-nil, zero value otherwise.

### GetMarksAndNumbersOk

`func (o *BillOfLadingDto) GetMarksAndNumbersOk() (*string, bool)`

GetMarksAndNumbersOk returns a tuple with the MarksAndNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarksAndNumbers

`func (o *BillOfLadingDto) SetMarksAndNumbers(v string)`

SetMarksAndNumbers sets MarksAndNumbers field to given value.

### HasMarksAndNumbers

`func (o *BillOfLadingDto) HasMarksAndNumbers() bool`

HasMarksAndNumbers returns a boolean if a field has been set.

### SetMarksAndNumbersNil

`func (o *BillOfLadingDto) SetMarksAndNumbersNil(b bool)`

 SetMarksAndNumbersNil sets the value for MarksAndNumbers to be an explicit nil

### UnsetMarksAndNumbers
`func (o *BillOfLadingDto) UnsetMarksAndNumbers()`

UnsetMarksAndNumbers ensures that no value is present for MarksAndNumbers, not even an explicit nil
### GetTotalPackages

`func (o *BillOfLadingDto) GetTotalPackages() int32`

GetTotalPackages returns the TotalPackages field if non-nil, zero value otherwise.

### GetTotalPackagesOk

`func (o *BillOfLadingDto) GetTotalPackagesOk() (*int32, bool)`

GetTotalPackagesOk returns a tuple with the TotalPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPackages

`func (o *BillOfLadingDto) SetTotalPackages(v int32)`

SetTotalPackages sets TotalPackages field to given value.

### HasTotalPackages

`func (o *BillOfLadingDto) HasTotalPackages() bool`

HasTotalPackages returns a boolean if a field has been set.

### GetTotalGrossWeightKg

`func (o *BillOfLadingDto) GetTotalGrossWeightKg() float64`

GetTotalGrossWeightKg returns the TotalGrossWeightKg field if non-nil, zero value otherwise.

### GetTotalGrossWeightKgOk

`func (o *BillOfLadingDto) GetTotalGrossWeightKgOk() (*float64, bool)`

GetTotalGrossWeightKgOk returns a tuple with the TotalGrossWeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGrossWeightKg

`func (o *BillOfLadingDto) SetTotalGrossWeightKg(v float64)`

SetTotalGrossWeightKg sets TotalGrossWeightKg field to given value.

### HasTotalGrossWeightKg

`func (o *BillOfLadingDto) HasTotalGrossWeightKg() bool`

HasTotalGrossWeightKg returns a boolean if a field has been set.

### GetTotalVolumeM3

`func (o *BillOfLadingDto) GetTotalVolumeM3() float64`

GetTotalVolumeM3 returns the TotalVolumeM3 field if non-nil, zero value otherwise.

### GetTotalVolumeM3Ok

`func (o *BillOfLadingDto) GetTotalVolumeM3Ok() (*float64, bool)`

GetTotalVolumeM3Ok returns a tuple with the TotalVolumeM3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolumeM3

`func (o *BillOfLadingDto) SetTotalVolumeM3(v float64)`

SetTotalVolumeM3 sets TotalVolumeM3 field to given value.

### HasTotalVolumeM3

`func (o *BillOfLadingDto) HasTotalVolumeM3() bool`

HasTotalVolumeM3 returns a boolean if a field has been set.

### SetTotalVolumeM3Nil

`func (o *BillOfLadingDto) SetTotalVolumeM3Nil(b bool)`

 SetTotalVolumeM3Nil sets the value for TotalVolumeM3 to be an explicit nil

### UnsetTotalVolumeM3
`func (o *BillOfLadingDto) UnsetTotalVolumeM3()`

UnsetTotalVolumeM3 ensures that no value is present for TotalVolumeM3, not even an explicit nil
### GetTenantId

`func (o *BillOfLadingDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BillOfLadingDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BillOfLadingDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BillOfLadingDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BillOfLadingDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BillOfLadingDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *BillOfLadingDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BillOfLadingDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BillOfLadingDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BillOfLadingDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BillOfLadingDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BillOfLadingDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


