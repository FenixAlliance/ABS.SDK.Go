# OrderLineUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Closed** | Pointer to **bool** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**ItemTitle** | Pointer to **NullableString** |  | [optional] 
**ItemShortDescription** | Pointer to **NullableString** |  | [optional] 
**ItemPrimaryImageUrl** | Pointer to **NullableString** |  | [optional] 
**ShippingPolicyId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**Free** | Pointer to **bool** |  | [optional] 
**FreeReason** | Pointer to **NullableString** |  | [optional] 
**FreeReasonCode** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**DataLabel** | Pointer to **NullableString** |  | [optional] 
**Data1** | Pointer to **NullableString** |  | [optional] 
**Data1Label** | Pointer to **NullableString** |  | [optional] 
**Data2** | Pointer to **NullableString** |  | [optional] 
**Data2Label** | Pointer to **NullableString** |  | [optional] 
**Data3** | Pointer to **NullableString** |  | [optional] 
**Data3Label** | Pointer to **NullableString** |  | [optional] 
**Data4** | Pointer to **NullableString** |  | [optional] 
**Data4Label** | Pointer to **NullableString** |  | [optional] 
**Data5** | Pointer to **NullableString** |  | [optional] 
**Data5Label** | Pointer to **NullableString** |  | [optional] 
**Data6** | Pointer to **NullableString** |  | [optional] 
**Data6Label** | Pointer to **NullableString** |  | [optional] 
**Data7** | Pointer to **NullableString** |  | [optional] 
**Data7Label** | Pointer to **NullableString** |  | [optional] 
**Data8** | Pointer to **NullableString** |  | [optional] 
**Data8Label** | Pointer to **NullableString** |  | [optional] 
**Data9** | Pointer to **NullableString** |  | [optional] 
**Data9Label** | Pointer to **NullableString** |  | [optional] 
**ItemPriceId** | Pointer to **NullableString** |  | [optional] 
**PriceListItemId** | Pointer to **NullableString** |  | [optional] 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**UnitGroupId** | Pointer to **NullableString** |  | [optional] 
**TaxCalculationMethod** | Pointer to **int32** |  | [optional] 
**CostCalculationMethod** | Pointer to **int32** |  | [optional] 
**ForexRatesSnapshot** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**TotalBaseAmountInUsd** | Pointer to **float64** |  | [optional] 
**TotalProfitInUsd** | Pointer to **float64** |  | [optional] 
**TotalDetailAmountInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseInUsd** | Pointer to **float64** |  | [optional] 
**TotalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalWithholdingTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalWarrantyCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalReturnCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalRefundCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**TotalAmountInUsd** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**CustomGlobalSurchargesAmount** | Pointer to **float64** |  | [optional] 
**CustomGlobalDiscountsAmount** | Pointer to **float64** |  | [optional] 
**CustomBaseAmount** | Pointer to **float64** |  | [optional] 
**CustomDetailAmount** | Pointer to **float64** |  | [optional] 
**CustomDiscountsAmount** | Pointer to **float64** |  | [optional] 
**CustomTaxBase** | Pointer to **float64** |  | [optional] 
**CustomSurchargesAmount** | Pointer to **float64** |  | [optional] 
**CustomProfitAmount** | Pointer to **float64** |  | [optional] 
**CustomShippingCostAmount** | Pointer to **float64** |  | [optional] 
**CustomShippingTaxAmount** | Pointer to **float64** |  | [optional] 
**CustomTaxAmount** | Pointer to **float64** |  | [optional] 
**CustomWithholdingTaxAmount** | Pointer to **float64** |  | [optional] 
**CustomTotalAmount** | Pointer to **float64** |  | [optional] 
**ReturnPolicyId** | Pointer to **NullableString** |  | [optional] 
**RefundPolicyId** | Pointer to **NullableString** |  | [optional] 
**WarrantyPolicyId** | Pointer to **NullableString** |  | [optional] 
**ShipmentPolicyId** | Pointer to **NullableString** |  | [optional] 
**ShippingLocationId** | Pointer to **NullableString** |  | [optional] 
**LocationId** | Pointer to **NullableString** |  | [optional] 
**QuoteItemRecordId** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **NullableString** |  | [optional] 
**ParentBillingItemRecordId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrderLineUpdateDto

`func NewOrderLineUpdateDto() *OrderLineUpdateDto`

NewOrderLineUpdateDto instantiates a new OrderLineUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLineUpdateDtoWithDefaults

`func NewOrderLineUpdateDtoWithDefaults() *OrderLineUpdateDto`

NewOrderLineUpdateDtoWithDefaults instantiates a new OrderLineUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosed

`func (o *OrderLineUpdateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *OrderLineUpdateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *OrderLineUpdateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *OrderLineUpdateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetItemId

`func (o *OrderLineUpdateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *OrderLineUpdateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *OrderLineUpdateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *OrderLineUpdateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *OrderLineUpdateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *OrderLineUpdateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemTitle

`func (o *OrderLineUpdateDto) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *OrderLineUpdateDto) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *OrderLineUpdateDto) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *OrderLineUpdateDto) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *OrderLineUpdateDto) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *OrderLineUpdateDto) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetItemShortDescription

`func (o *OrderLineUpdateDto) GetItemShortDescription() string`

GetItemShortDescription returns the ItemShortDescription field if non-nil, zero value otherwise.

### GetItemShortDescriptionOk

`func (o *OrderLineUpdateDto) GetItemShortDescriptionOk() (*string, bool)`

GetItemShortDescriptionOk returns a tuple with the ItemShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemShortDescription

`func (o *OrderLineUpdateDto) SetItemShortDescription(v string)`

SetItemShortDescription sets ItemShortDescription field to given value.

### HasItemShortDescription

`func (o *OrderLineUpdateDto) HasItemShortDescription() bool`

HasItemShortDescription returns a boolean if a field has been set.

### SetItemShortDescriptionNil

`func (o *OrderLineUpdateDto) SetItemShortDescriptionNil(b bool)`

 SetItemShortDescriptionNil sets the value for ItemShortDescription to be an explicit nil

### UnsetItemShortDescription
`func (o *OrderLineUpdateDto) UnsetItemShortDescription()`

UnsetItemShortDescription ensures that no value is present for ItemShortDescription, not even an explicit nil
### GetItemPrimaryImageUrl

`func (o *OrderLineUpdateDto) GetItemPrimaryImageUrl() string`

GetItemPrimaryImageUrl returns the ItemPrimaryImageUrl field if non-nil, zero value otherwise.

### GetItemPrimaryImageUrlOk

`func (o *OrderLineUpdateDto) GetItemPrimaryImageUrlOk() (*string, bool)`

GetItemPrimaryImageUrlOk returns a tuple with the ItemPrimaryImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrimaryImageUrl

`func (o *OrderLineUpdateDto) SetItemPrimaryImageUrl(v string)`

SetItemPrimaryImageUrl sets ItemPrimaryImageUrl field to given value.

### HasItemPrimaryImageUrl

`func (o *OrderLineUpdateDto) HasItemPrimaryImageUrl() bool`

HasItemPrimaryImageUrl returns a boolean if a field has been set.

### SetItemPrimaryImageUrlNil

`func (o *OrderLineUpdateDto) SetItemPrimaryImageUrlNil(b bool)`

 SetItemPrimaryImageUrlNil sets the value for ItemPrimaryImageUrl to be an explicit nil

### UnsetItemPrimaryImageUrl
`func (o *OrderLineUpdateDto) UnsetItemPrimaryImageUrl()`

UnsetItemPrimaryImageUrl ensures that no value is present for ItemPrimaryImageUrl, not even an explicit nil
### GetShippingPolicyId

`func (o *OrderLineUpdateDto) GetShippingPolicyId() string`

GetShippingPolicyId returns the ShippingPolicyId field if non-nil, zero value otherwise.

### GetShippingPolicyIdOk

`func (o *OrderLineUpdateDto) GetShippingPolicyIdOk() (*string, bool)`

GetShippingPolicyIdOk returns a tuple with the ShippingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPolicyId

`func (o *OrderLineUpdateDto) SetShippingPolicyId(v string)`

SetShippingPolicyId sets ShippingPolicyId field to given value.

### HasShippingPolicyId

`func (o *OrderLineUpdateDto) HasShippingPolicyId() bool`

HasShippingPolicyId returns a boolean if a field has been set.

### SetShippingPolicyIdNil

`func (o *OrderLineUpdateDto) SetShippingPolicyIdNil(b bool)`

 SetShippingPolicyIdNil sets the value for ShippingPolicyId to be an explicit nil

### UnsetShippingPolicyId
`func (o *OrderLineUpdateDto) UnsetShippingPolicyId()`

UnsetShippingPolicyId ensures that no value is present for ShippingPolicyId, not even an explicit nil
### GetTenantId

`func (o *OrderLineUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OrderLineUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OrderLineUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *OrderLineUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *OrderLineUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *OrderLineUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *OrderLineUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *OrderLineUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *OrderLineUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *OrderLineUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *OrderLineUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *OrderLineUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCurrencyId

`func (o *OrderLineUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *OrderLineUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *OrderLineUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *OrderLineUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *OrderLineUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *OrderLineUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDescription

`func (o *OrderLineUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderLineUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderLineUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderLineUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OrderLineUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrderLineUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *OrderLineUpdateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderLineUpdateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderLineUpdateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderLineUpdateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetFree

`func (o *OrderLineUpdateDto) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *OrderLineUpdateDto) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *OrderLineUpdateDto) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *OrderLineUpdateDto) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetFreeReason

`func (o *OrderLineUpdateDto) GetFreeReason() string`

GetFreeReason returns the FreeReason field if non-nil, zero value otherwise.

### GetFreeReasonOk

`func (o *OrderLineUpdateDto) GetFreeReasonOk() (*string, bool)`

GetFreeReasonOk returns a tuple with the FreeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeReason

`func (o *OrderLineUpdateDto) SetFreeReason(v string)`

SetFreeReason sets FreeReason field to given value.

### HasFreeReason

`func (o *OrderLineUpdateDto) HasFreeReason() bool`

HasFreeReason returns a boolean if a field has been set.

### SetFreeReasonNil

`func (o *OrderLineUpdateDto) SetFreeReasonNil(b bool)`

 SetFreeReasonNil sets the value for FreeReason to be an explicit nil

### UnsetFreeReason
`func (o *OrderLineUpdateDto) UnsetFreeReason()`

UnsetFreeReason ensures that no value is present for FreeReason, not even an explicit nil
### GetFreeReasonCode

`func (o *OrderLineUpdateDto) GetFreeReasonCode() string`

GetFreeReasonCode returns the FreeReasonCode field if non-nil, zero value otherwise.

### GetFreeReasonCodeOk

`func (o *OrderLineUpdateDto) GetFreeReasonCodeOk() (*string, bool)`

GetFreeReasonCodeOk returns a tuple with the FreeReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeReasonCode

`func (o *OrderLineUpdateDto) SetFreeReasonCode(v string)`

SetFreeReasonCode sets FreeReasonCode field to given value.

### HasFreeReasonCode

`func (o *OrderLineUpdateDto) HasFreeReasonCode() bool`

HasFreeReasonCode returns a boolean if a field has been set.

### SetFreeReasonCodeNil

`func (o *OrderLineUpdateDto) SetFreeReasonCodeNil(b bool)`

 SetFreeReasonCodeNil sets the value for FreeReasonCode to be an explicit nil

### UnsetFreeReasonCode
`func (o *OrderLineUpdateDto) UnsetFreeReasonCode()`

UnsetFreeReasonCode ensures that no value is present for FreeReasonCode, not even an explicit nil
### GetData

`func (o *OrderLineUpdateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderLineUpdateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderLineUpdateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *OrderLineUpdateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *OrderLineUpdateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OrderLineUpdateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *OrderLineUpdateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *OrderLineUpdateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *OrderLineUpdateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *OrderLineUpdateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *OrderLineUpdateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *OrderLineUpdateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *OrderLineUpdateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *OrderLineUpdateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *OrderLineUpdateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *OrderLineUpdateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *OrderLineUpdateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *OrderLineUpdateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *OrderLineUpdateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *OrderLineUpdateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *OrderLineUpdateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *OrderLineUpdateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *OrderLineUpdateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *OrderLineUpdateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *OrderLineUpdateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *OrderLineUpdateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *OrderLineUpdateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *OrderLineUpdateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *OrderLineUpdateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *OrderLineUpdateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *OrderLineUpdateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *OrderLineUpdateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *OrderLineUpdateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *OrderLineUpdateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *OrderLineUpdateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *OrderLineUpdateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *OrderLineUpdateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *OrderLineUpdateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *OrderLineUpdateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *OrderLineUpdateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *OrderLineUpdateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *OrderLineUpdateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *OrderLineUpdateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *OrderLineUpdateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *OrderLineUpdateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *OrderLineUpdateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *OrderLineUpdateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *OrderLineUpdateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *OrderLineUpdateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *OrderLineUpdateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *OrderLineUpdateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *OrderLineUpdateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *OrderLineUpdateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *OrderLineUpdateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *OrderLineUpdateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *OrderLineUpdateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *OrderLineUpdateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *OrderLineUpdateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *OrderLineUpdateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *OrderLineUpdateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *OrderLineUpdateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *OrderLineUpdateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *OrderLineUpdateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *OrderLineUpdateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *OrderLineUpdateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *OrderLineUpdateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *OrderLineUpdateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *OrderLineUpdateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *OrderLineUpdateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *OrderLineUpdateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *OrderLineUpdateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *OrderLineUpdateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *OrderLineUpdateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *OrderLineUpdateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *OrderLineUpdateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *OrderLineUpdateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *OrderLineUpdateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *OrderLineUpdateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *OrderLineUpdateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *OrderLineUpdateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *OrderLineUpdateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *OrderLineUpdateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *OrderLineUpdateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *OrderLineUpdateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *OrderLineUpdateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *OrderLineUpdateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *OrderLineUpdateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *OrderLineUpdateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *OrderLineUpdateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *OrderLineUpdateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *OrderLineUpdateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *OrderLineUpdateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *OrderLineUpdateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *OrderLineUpdateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *OrderLineUpdateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *OrderLineUpdateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *OrderLineUpdateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *OrderLineUpdateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *OrderLineUpdateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *OrderLineUpdateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *OrderLineUpdateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *OrderLineUpdateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *OrderLineUpdateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *OrderLineUpdateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *OrderLineUpdateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *OrderLineUpdateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *OrderLineUpdateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *OrderLineUpdateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *OrderLineUpdateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *OrderLineUpdateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *OrderLineUpdateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *OrderLineUpdateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *OrderLineUpdateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *OrderLineUpdateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *OrderLineUpdateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *OrderLineUpdateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *OrderLineUpdateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *OrderLineUpdateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *OrderLineUpdateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *OrderLineUpdateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetItemPriceId

`func (o *OrderLineUpdateDto) GetItemPriceId() string`

GetItemPriceId returns the ItemPriceId field if non-nil, zero value otherwise.

### GetItemPriceIdOk

`func (o *OrderLineUpdateDto) GetItemPriceIdOk() (*string, bool)`

GetItemPriceIdOk returns a tuple with the ItemPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceId

`func (o *OrderLineUpdateDto) SetItemPriceId(v string)`

SetItemPriceId sets ItemPriceId field to given value.

### HasItemPriceId

`func (o *OrderLineUpdateDto) HasItemPriceId() bool`

HasItemPriceId returns a boolean if a field has been set.

### SetItemPriceIdNil

`func (o *OrderLineUpdateDto) SetItemPriceIdNil(b bool)`

 SetItemPriceIdNil sets the value for ItemPriceId to be an explicit nil

### UnsetItemPriceId
`func (o *OrderLineUpdateDto) UnsetItemPriceId()`

UnsetItemPriceId ensures that no value is present for ItemPriceId, not even an explicit nil
### GetPriceListItemId

`func (o *OrderLineUpdateDto) GetPriceListItemId() string`

GetPriceListItemId returns the PriceListItemId field if non-nil, zero value otherwise.

### GetPriceListItemIdOk

`func (o *OrderLineUpdateDto) GetPriceListItemIdOk() (*string, bool)`

GetPriceListItemIdOk returns a tuple with the PriceListItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListItemId

`func (o *OrderLineUpdateDto) SetPriceListItemId(v string)`

SetPriceListItemId sets PriceListItemId field to given value.

### HasPriceListItemId

`func (o *OrderLineUpdateDto) HasPriceListItemId() bool`

HasPriceListItemId returns a boolean if a field has been set.

### SetPriceListItemIdNil

`func (o *OrderLineUpdateDto) SetPriceListItemIdNil(b bool)`

 SetPriceListItemIdNil sets the value for PriceListItemId to be an explicit nil

### UnsetPriceListItemId
`func (o *OrderLineUpdateDto) UnsetPriceListItemId()`

UnsetPriceListItemId ensures that no value is present for PriceListItemId, not even an explicit nil
### GetUnitId

`func (o *OrderLineUpdateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *OrderLineUpdateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *OrderLineUpdateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *OrderLineUpdateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *OrderLineUpdateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *OrderLineUpdateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *OrderLineUpdateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *OrderLineUpdateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *OrderLineUpdateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *OrderLineUpdateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *OrderLineUpdateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *OrderLineUpdateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetTaxCalculationMethod

`func (o *OrderLineUpdateDto) GetTaxCalculationMethod() int32`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *OrderLineUpdateDto) GetTaxCalculationMethodOk() (*int32, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *OrderLineUpdateDto) SetTaxCalculationMethod(v int32)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *OrderLineUpdateDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *OrderLineUpdateDto) GetCostCalculationMethod() int32`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *OrderLineUpdateDto) GetCostCalculationMethodOk() (*int32, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *OrderLineUpdateDto) SetCostCalculationMethod(v int32)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *OrderLineUpdateDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetForexRatesSnapshot

`func (o *OrderLineUpdateDto) GetForexRatesSnapshot() string`

GetForexRatesSnapshot returns the ForexRatesSnapshot field if non-nil, zero value otherwise.

### GetForexRatesSnapshotOk

`func (o *OrderLineUpdateDto) GetForexRatesSnapshotOk() (*string, bool)`

GetForexRatesSnapshotOk returns a tuple with the ForexRatesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRatesSnapshot

`func (o *OrderLineUpdateDto) SetForexRatesSnapshot(v string)`

SetForexRatesSnapshot sets ForexRatesSnapshot field to given value.

### HasForexRatesSnapshot

`func (o *OrderLineUpdateDto) HasForexRatesSnapshot() bool`

HasForexRatesSnapshot returns a boolean if a field has been set.

### SetForexRatesSnapshotNil

`func (o *OrderLineUpdateDto) SetForexRatesSnapshotNil(b bool)`

 SetForexRatesSnapshotNil sets the value for ForexRatesSnapshot to be an explicit nil

### UnsetForexRatesSnapshot
`func (o *OrderLineUpdateDto) UnsetForexRatesSnapshot()`

UnsetForexRatesSnapshot ensures that no value is present for ForexRatesSnapshot, not even an explicit nil
### GetForexRate

`func (o *OrderLineUpdateDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *OrderLineUpdateDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *OrderLineUpdateDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *OrderLineUpdateDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetTotalBaseAmountInUsd

`func (o *OrderLineUpdateDto) GetTotalBaseAmountInUsd() float64`

GetTotalBaseAmountInUsd returns the TotalBaseAmountInUsd field if non-nil, zero value otherwise.

### GetTotalBaseAmountInUsdOk

`func (o *OrderLineUpdateDto) GetTotalBaseAmountInUsdOk() (*float64, bool)`

GetTotalBaseAmountInUsdOk returns a tuple with the TotalBaseAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBaseAmountInUsd

`func (o *OrderLineUpdateDto) SetTotalBaseAmountInUsd(v float64)`

SetTotalBaseAmountInUsd sets TotalBaseAmountInUsd field to given value.

### HasTotalBaseAmountInUsd

`func (o *OrderLineUpdateDto) HasTotalBaseAmountInUsd() bool`

HasTotalBaseAmountInUsd returns a boolean if a field has been set.

### GetTotalProfitInUsd

`func (o *OrderLineUpdateDto) GetTotalProfitInUsd() float64`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *OrderLineUpdateDto) GetTotalProfitInUsdOk() (*float64, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *OrderLineUpdateDto) SetTotalProfitInUsd(v float64)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.

### HasTotalProfitInUsd

`func (o *OrderLineUpdateDto) HasTotalProfitInUsd() bool`

HasTotalProfitInUsd returns a boolean if a field has been set.

### GetTotalDetailAmountInUsd

`func (o *OrderLineUpdateDto) GetTotalDetailAmountInUsd() float64`

GetTotalDetailAmountInUsd returns the TotalDetailAmountInUsd field if non-nil, zero value otherwise.

### GetTotalDetailAmountInUsdOk

`func (o *OrderLineUpdateDto) GetTotalDetailAmountInUsdOk() (*float64, bool)`

GetTotalDetailAmountInUsdOk returns a tuple with the TotalDetailAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailAmountInUsd

`func (o *OrderLineUpdateDto) SetTotalDetailAmountInUsd(v float64)`

SetTotalDetailAmountInUsd sets TotalDetailAmountInUsd field to given value.

### HasTotalDetailAmountInUsd

`func (o *OrderLineUpdateDto) HasTotalDetailAmountInUsd() bool`

HasTotalDetailAmountInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseInUsd

`func (o *OrderLineUpdateDto) GetTotalTaxBaseInUsd() float64`

GetTotalTaxBaseInUsd returns the TotalTaxBaseInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseInUsdOk

`func (o *OrderLineUpdateDto) GetTotalTaxBaseInUsdOk() (*float64, bool)`

GetTotalTaxBaseInUsdOk returns a tuple with the TotalTaxBaseInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseInUsd

`func (o *OrderLineUpdateDto) SetTotalTaxBaseInUsd(v float64)`

SetTotalTaxBaseInUsd sets TotalTaxBaseInUsd field to given value.

### HasTotalTaxBaseInUsd

`func (o *OrderLineUpdateDto) HasTotalTaxBaseInUsd() bool`

HasTotalTaxBaseInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *OrderLineUpdateDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *OrderLineUpdateDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *OrderLineUpdateDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *OrderLineUpdateDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalTaxesInUsd

`func (o *OrderLineUpdateDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *OrderLineUpdateDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *OrderLineUpdateDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *OrderLineUpdateDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalWithholdingTaxesInUsd

`func (o *OrderLineUpdateDto) GetTotalWithholdingTaxesInUsd() float64`

GetTotalWithholdingTaxesInUsd returns the TotalWithholdingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithholdingTaxesInUsdOk

`func (o *OrderLineUpdateDto) GetTotalWithholdingTaxesInUsdOk() (*float64, bool)`

GetTotalWithholdingTaxesInUsdOk returns a tuple with the TotalWithholdingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithholdingTaxesInUsd

`func (o *OrderLineUpdateDto) SetTotalWithholdingTaxesInUsd(v float64)`

SetTotalWithholdingTaxesInUsd sets TotalWithholdingTaxesInUsd field to given value.

### HasTotalWithholdingTaxesInUsd

`func (o *OrderLineUpdateDto) HasTotalWithholdingTaxesInUsd() bool`

HasTotalWithholdingTaxesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *OrderLineUpdateDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *OrderLineUpdateDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *OrderLineUpdateDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *OrderLineUpdateDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxesInUsd

`func (o *OrderLineUpdateDto) GetTotalShippingTaxesInUsd() float64`

GetTotalShippingTaxesInUsd returns the TotalShippingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxesInUsdOk

`func (o *OrderLineUpdateDto) GetTotalShippingTaxesInUsdOk() (*float64, bool)`

GetTotalShippingTaxesInUsdOk returns a tuple with the TotalShippingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxesInUsd

`func (o *OrderLineUpdateDto) SetTotalShippingTaxesInUsd(v float64)`

SetTotalShippingTaxesInUsd sets TotalShippingTaxesInUsd field to given value.

### HasTotalShippingTaxesInUsd

`func (o *OrderLineUpdateDto) HasTotalShippingTaxesInUsd() bool`

HasTotalShippingTaxesInUsd returns a boolean if a field has been set.

### GetTotalWarrantyCostInUsd

`func (o *OrderLineUpdateDto) GetTotalWarrantyCostInUsd() float64`

GetTotalWarrantyCostInUsd returns the TotalWarrantyCostInUsd field if non-nil, zero value otherwise.

### GetTotalWarrantyCostInUsdOk

`func (o *OrderLineUpdateDto) GetTotalWarrantyCostInUsdOk() (*float64, bool)`

GetTotalWarrantyCostInUsdOk returns a tuple with the TotalWarrantyCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWarrantyCostInUsd

`func (o *OrderLineUpdateDto) SetTotalWarrantyCostInUsd(v float64)`

SetTotalWarrantyCostInUsd sets TotalWarrantyCostInUsd field to given value.

### HasTotalWarrantyCostInUsd

`func (o *OrderLineUpdateDto) HasTotalWarrantyCostInUsd() bool`

HasTotalWarrantyCostInUsd returns a boolean if a field has been set.

### GetTotalReturnCostInUsd

`func (o *OrderLineUpdateDto) GetTotalReturnCostInUsd() float64`

GetTotalReturnCostInUsd returns the TotalReturnCostInUsd field if non-nil, zero value otherwise.

### GetTotalReturnCostInUsdOk

`func (o *OrderLineUpdateDto) GetTotalReturnCostInUsdOk() (*float64, bool)`

GetTotalReturnCostInUsdOk returns a tuple with the TotalReturnCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReturnCostInUsd

`func (o *OrderLineUpdateDto) SetTotalReturnCostInUsd(v float64)`

SetTotalReturnCostInUsd sets TotalReturnCostInUsd field to given value.

### HasTotalReturnCostInUsd

`func (o *OrderLineUpdateDto) HasTotalReturnCostInUsd() bool`

HasTotalReturnCostInUsd returns a boolean if a field has been set.

### GetTotalRefundCostInUsd

`func (o *OrderLineUpdateDto) GetTotalRefundCostInUsd() float64`

GetTotalRefundCostInUsd returns the TotalRefundCostInUsd field if non-nil, zero value otherwise.

### GetTotalRefundCostInUsdOk

`func (o *OrderLineUpdateDto) GetTotalRefundCostInUsdOk() (*float64, bool)`

GetTotalRefundCostInUsdOk returns a tuple with the TotalRefundCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRefundCostInUsd

`func (o *OrderLineUpdateDto) SetTotalRefundCostInUsd(v float64)`

SetTotalRefundCostInUsd sets TotalRefundCostInUsd field to given value.

### HasTotalRefundCostInUsd

`func (o *OrderLineUpdateDto) HasTotalRefundCostInUsd() bool`

HasTotalRefundCostInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *OrderLineUpdateDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *OrderLineUpdateDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *OrderLineUpdateDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *OrderLineUpdateDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalAmountInUsd

`func (o *OrderLineUpdateDto) GetTotalAmountInUsd() float64`

GetTotalAmountInUsd returns the TotalAmountInUsd field if non-nil, zero value otherwise.

### GetTotalAmountInUsdOk

`func (o *OrderLineUpdateDto) GetTotalAmountInUsdOk() (*float64, bool)`

GetTotalAmountInUsdOk returns a tuple with the TotalAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInUsd

`func (o *OrderLineUpdateDto) SetTotalAmountInUsd(v float64)`

SetTotalAmountInUsd sets TotalAmountInUsd field to given value.

### HasTotalAmountInUsd

`func (o *OrderLineUpdateDto) HasTotalAmountInUsd() bool`

HasTotalAmountInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsInUsd

`func (o *OrderLineUpdateDto) GetTotalGlobalDiscountsInUsd() float64`

GetTotalGlobalDiscountsInUsd returns the TotalGlobalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsInUsdOk

`func (o *OrderLineUpdateDto) GetTotalGlobalDiscountsInUsdOk() (*float64, bool)`

GetTotalGlobalDiscountsInUsdOk returns a tuple with the TotalGlobalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsInUsd

`func (o *OrderLineUpdateDto) SetTotalGlobalDiscountsInUsd(v float64)`

SetTotalGlobalDiscountsInUsd sets TotalGlobalDiscountsInUsd field to given value.

### HasTotalGlobalDiscountsInUsd

`func (o *OrderLineUpdateDto) HasTotalGlobalDiscountsInUsd() bool`

HasTotalGlobalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesInUsd

`func (o *OrderLineUpdateDto) GetTotalGlobalSurchargesInUsd() float64`

GetTotalGlobalSurchargesInUsd returns the TotalGlobalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesInUsdOk

`func (o *OrderLineUpdateDto) GetTotalGlobalSurchargesInUsdOk() (*float64, bool)`

GetTotalGlobalSurchargesInUsdOk returns a tuple with the TotalGlobalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesInUsd

`func (o *OrderLineUpdateDto) SetTotalGlobalSurchargesInUsd(v float64)`

SetTotalGlobalSurchargesInUsd sets TotalGlobalSurchargesInUsd field to given value.

### HasTotalGlobalSurchargesInUsd

`func (o *OrderLineUpdateDto) HasTotalGlobalSurchargesInUsd() bool`

HasTotalGlobalSurchargesInUsd returns a boolean if a field has been set.

### GetCustomGlobalSurchargesAmount

`func (o *OrderLineUpdateDto) GetCustomGlobalSurchargesAmount() float64`

GetCustomGlobalSurchargesAmount returns the CustomGlobalSurchargesAmount field if non-nil, zero value otherwise.

### GetCustomGlobalSurchargesAmountOk

`func (o *OrderLineUpdateDto) GetCustomGlobalSurchargesAmountOk() (*float64, bool)`

GetCustomGlobalSurchargesAmountOk returns a tuple with the CustomGlobalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalSurchargesAmount

`func (o *OrderLineUpdateDto) SetCustomGlobalSurchargesAmount(v float64)`

SetCustomGlobalSurchargesAmount sets CustomGlobalSurchargesAmount field to given value.

### HasCustomGlobalSurchargesAmount

`func (o *OrderLineUpdateDto) HasCustomGlobalSurchargesAmount() bool`

HasCustomGlobalSurchargesAmount returns a boolean if a field has been set.

### GetCustomGlobalDiscountsAmount

`func (o *OrderLineUpdateDto) GetCustomGlobalDiscountsAmount() float64`

GetCustomGlobalDiscountsAmount returns the CustomGlobalDiscountsAmount field if non-nil, zero value otherwise.

### GetCustomGlobalDiscountsAmountOk

`func (o *OrderLineUpdateDto) GetCustomGlobalDiscountsAmountOk() (*float64, bool)`

GetCustomGlobalDiscountsAmountOk returns a tuple with the CustomGlobalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalDiscountsAmount

`func (o *OrderLineUpdateDto) SetCustomGlobalDiscountsAmount(v float64)`

SetCustomGlobalDiscountsAmount sets CustomGlobalDiscountsAmount field to given value.

### HasCustomGlobalDiscountsAmount

`func (o *OrderLineUpdateDto) HasCustomGlobalDiscountsAmount() bool`

HasCustomGlobalDiscountsAmount returns a boolean if a field has been set.

### GetCustomBaseAmount

`func (o *OrderLineUpdateDto) GetCustomBaseAmount() float64`

GetCustomBaseAmount returns the CustomBaseAmount field if non-nil, zero value otherwise.

### GetCustomBaseAmountOk

`func (o *OrderLineUpdateDto) GetCustomBaseAmountOk() (*float64, bool)`

GetCustomBaseAmountOk returns a tuple with the CustomBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBaseAmount

`func (o *OrderLineUpdateDto) SetCustomBaseAmount(v float64)`

SetCustomBaseAmount sets CustomBaseAmount field to given value.

### HasCustomBaseAmount

`func (o *OrderLineUpdateDto) HasCustomBaseAmount() bool`

HasCustomBaseAmount returns a boolean if a field has been set.

### GetCustomDetailAmount

`func (o *OrderLineUpdateDto) GetCustomDetailAmount() float64`

GetCustomDetailAmount returns the CustomDetailAmount field if non-nil, zero value otherwise.

### GetCustomDetailAmountOk

`func (o *OrderLineUpdateDto) GetCustomDetailAmountOk() (*float64, bool)`

GetCustomDetailAmountOk returns a tuple with the CustomDetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDetailAmount

`func (o *OrderLineUpdateDto) SetCustomDetailAmount(v float64)`

SetCustomDetailAmount sets CustomDetailAmount field to given value.

### HasCustomDetailAmount

`func (o *OrderLineUpdateDto) HasCustomDetailAmount() bool`

HasCustomDetailAmount returns a boolean if a field has been set.

### GetCustomDiscountsAmount

`func (o *OrderLineUpdateDto) GetCustomDiscountsAmount() float64`

GetCustomDiscountsAmount returns the CustomDiscountsAmount field if non-nil, zero value otherwise.

### GetCustomDiscountsAmountOk

`func (o *OrderLineUpdateDto) GetCustomDiscountsAmountOk() (*float64, bool)`

GetCustomDiscountsAmountOk returns a tuple with the CustomDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDiscountsAmount

`func (o *OrderLineUpdateDto) SetCustomDiscountsAmount(v float64)`

SetCustomDiscountsAmount sets CustomDiscountsAmount field to given value.

### HasCustomDiscountsAmount

`func (o *OrderLineUpdateDto) HasCustomDiscountsAmount() bool`

HasCustomDiscountsAmount returns a boolean if a field has been set.

### GetCustomTaxBase

`func (o *OrderLineUpdateDto) GetCustomTaxBase() float64`

GetCustomTaxBase returns the CustomTaxBase field if non-nil, zero value otherwise.

### GetCustomTaxBaseOk

`func (o *OrderLineUpdateDto) GetCustomTaxBaseOk() (*float64, bool)`

GetCustomTaxBaseOk returns a tuple with the CustomTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaxBase

`func (o *OrderLineUpdateDto) SetCustomTaxBase(v float64)`

SetCustomTaxBase sets CustomTaxBase field to given value.

### HasCustomTaxBase

`func (o *OrderLineUpdateDto) HasCustomTaxBase() bool`

HasCustomTaxBase returns a boolean if a field has been set.

### GetCustomSurchargesAmount

`func (o *OrderLineUpdateDto) GetCustomSurchargesAmount() float64`

GetCustomSurchargesAmount returns the CustomSurchargesAmount field if non-nil, zero value otherwise.

### GetCustomSurchargesAmountOk

`func (o *OrderLineUpdateDto) GetCustomSurchargesAmountOk() (*float64, bool)`

GetCustomSurchargesAmountOk returns a tuple with the CustomSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargesAmount

`func (o *OrderLineUpdateDto) SetCustomSurchargesAmount(v float64)`

SetCustomSurchargesAmount sets CustomSurchargesAmount field to given value.

### HasCustomSurchargesAmount

`func (o *OrderLineUpdateDto) HasCustomSurchargesAmount() bool`

HasCustomSurchargesAmount returns a boolean if a field has been set.

### GetCustomProfitAmount

`func (o *OrderLineUpdateDto) GetCustomProfitAmount() float64`

GetCustomProfitAmount returns the CustomProfitAmount field if non-nil, zero value otherwise.

### GetCustomProfitAmountOk

`func (o *OrderLineUpdateDto) GetCustomProfitAmountOk() (*float64, bool)`

GetCustomProfitAmountOk returns a tuple with the CustomProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProfitAmount

`func (o *OrderLineUpdateDto) SetCustomProfitAmount(v float64)`

SetCustomProfitAmount sets CustomProfitAmount field to given value.

### HasCustomProfitAmount

`func (o *OrderLineUpdateDto) HasCustomProfitAmount() bool`

HasCustomProfitAmount returns a boolean if a field has been set.

### GetCustomShippingCostAmount

`func (o *OrderLineUpdateDto) GetCustomShippingCostAmount() float64`

GetCustomShippingCostAmount returns the CustomShippingCostAmount field if non-nil, zero value otherwise.

### GetCustomShippingCostAmountOk

`func (o *OrderLineUpdateDto) GetCustomShippingCostAmountOk() (*float64, bool)`

GetCustomShippingCostAmountOk returns a tuple with the CustomShippingCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShippingCostAmount

`func (o *OrderLineUpdateDto) SetCustomShippingCostAmount(v float64)`

SetCustomShippingCostAmount sets CustomShippingCostAmount field to given value.

### HasCustomShippingCostAmount

`func (o *OrderLineUpdateDto) HasCustomShippingCostAmount() bool`

HasCustomShippingCostAmount returns a boolean if a field has been set.

### GetCustomShippingTaxAmount

`func (o *OrderLineUpdateDto) GetCustomShippingTaxAmount() float64`

GetCustomShippingTaxAmount returns the CustomShippingTaxAmount field if non-nil, zero value otherwise.

### GetCustomShippingTaxAmountOk

`func (o *OrderLineUpdateDto) GetCustomShippingTaxAmountOk() (*float64, bool)`

GetCustomShippingTaxAmountOk returns a tuple with the CustomShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShippingTaxAmount

`func (o *OrderLineUpdateDto) SetCustomShippingTaxAmount(v float64)`

SetCustomShippingTaxAmount sets CustomShippingTaxAmount field to given value.

### HasCustomShippingTaxAmount

`func (o *OrderLineUpdateDto) HasCustomShippingTaxAmount() bool`

HasCustomShippingTaxAmount returns a boolean if a field has been set.

### GetCustomTaxAmount

`func (o *OrderLineUpdateDto) GetCustomTaxAmount() float64`

GetCustomTaxAmount returns the CustomTaxAmount field if non-nil, zero value otherwise.

### GetCustomTaxAmountOk

`func (o *OrderLineUpdateDto) GetCustomTaxAmountOk() (*float64, bool)`

GetCustomTaxAmountOk returns a tuple with the CustomTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaxAmount

`func (o *OrderLineUpdateDto) SetCustomTaxAmount(v float64)`

SetCustomTaxAmount sets CustomTaxAmount field to given value.

### HasCustomTaxAmount

`func (o *OrderLineUpdateDto) HasCustomTaxAmount() bool`

HasCustomTaxAmount returns a boolean if a field has been set.

### GetCustomWithholdingTaxAmount

`func (o *OrderLineUpdateDto) GetCustomWithholdingTaxAmount() float64`

GetCustomWithholdingTaxAmount returns the CustomWithholdingTaxAmount field if non-nil, zero value otherwise.

### GetCustomWithholdingTaxAmountOk

`func (o *OrderLineUpdateDto) GetCustomWithholdingTaxAmountOk() (*float64, bool)`

GetCustomWithholdingTaxAmountOk returns a tuple with the CustomWithholdingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomWithholdingTaxAmount

`func (o *OrderLineUpdateDto) SetCustomWithholdingTaxAmount(v float64)`

SetCustomWithholdingTaxAmount sets CustomWithholdingTaxAmount field to given value.

### HasCustomWithholdingTaxAmount

`func (o *OrderLineUpdateDto) HasCustomWithholdingTaxAmount() bool`

HasCustomWithholdingTaxAmount returns a boolean if a field has been set.

### GetCustomTotalAmount

`func (o *OrderLineUpdateDto) GetCustomTotalAmount() float64`

GetCustomTotalAmount returns the CustomTotalAmount field if non-nil, zero value otherwise.

### GetCustomTotalAmountOk

`func (o *OrderLineUpdateDto) GetCustomTotalAmountOk() (*float64, bool)`

GetCustomTotalAmountOk returns a tuple with the CustomTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTotalAmount

`func (o *OrderLineUpdateDto) SetCustomTotalAmount(v float64)`

SetCustomTotalAmount sets CustomTotalAmount field to given value.

### HasCustomTotalAmount

`func (o *OrderLineUpdateDto) HasCustomTotalAmount() bool`

HasCustomTotalAmount returns a boolean if a field has been set.

### GetReturnPolicyId

`func (o *OrderLineUpdateDto) GetReturnPolicyId() string`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *OrderLineUpdateDto) GetReturnPolicyIdOk() (*string, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *OrderLineUpdateDto) SetReturnPolicyId(v string)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *OrderLineUpdateDto) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### SetReturnPolicyIdNil

`func (o *OrderLineUpdateDto) SetReturnPolicyIdNil(b bool)`

 SetReturnPolicyIdNil sets the value for ReturnPolicyId to be an explicit nil

### UnsetReturnPolicyId
`func (o *OrderLineUpdateDto) UnsetReturnPolicyId()`

UnsetReturnPolicyId ensures that no value is present for ReturnPolicyId, not even an explicit nil
### GetRefundPolicyId

`func (o *OrderLineUpdateDto) GetRefundPolicyId() string`

GetRefundPolicyId returns the RefundPolicyId field if non-nil, zero value otherwise.

### GetRefundPolicyIdOk

`func (o *OrderLineUpdateDto) GetRefundPolicyIdOk() (*string, bool)`

GetRefundPolicyIdOk returns a tuple with the RefundPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPolicyId

`func (o *OrderLineUpdateDto) SetRefundPolicyId(v string)`

SetRefundPolicyId sets RefundPolicyId field to given value.

### HasRefundPolicyId

`func (o *OrderLineUpdateDto) HasRefundPolicyId() bool`

HasRefundPolicyId returns a boolean if a field has been set.

### SetRefundPolicyIdNil

`func (o *OrderLineUpdateDto) SetRefundPolicyIdNil(b bool)`

 SetRefundPolicyIdNil sets the value for RefundPolicyId to be an explicit nil

### UnsetRefundPolicyId
`func (o *OrderLineUpdateDto) UnsetRefundPolicyId()`

UnsetRefundPolicyId ensures that no value is present for RefundPolicyId, not even an explicit nil
### GetWarrantyPolicyId

`func (o *OrderLineUpdateDto) GetWarrantyPolicyId() string`

GetWarrantyPolicyId returns the WarrantyPolicyId field if non-nil, zero value otherwise.

### GetWarrantyPolicyIdOk

`func (o *OrderLineUpdateDto) GetWarrantyPolicyIdOk() (*string, bool)`

GetWarrantyPolicyIdOk returns a tuple with the WarrantyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyPolicyId

`func (o *OrderLineUpdateDto) SetWarrantyPolicyId(v string)`

SetWarrantyPolicyId sets WarrantyPolicyId field to given value.

### HasWarrantyPolicyId

`func (o *OrderLineUpdateDto) HasWarrantyPolicyId() bool`

HasWarrantyPolicyId returns a boolean if a field has been set.

### SetWarrantyPolicyIdNil

`func (o *OrderLineUpdateDto) SetWarrantyPolicyIdNil(b bool)`

 SetWarrantyPolicyIdNil sets the value for WarrantyPolicyId to be an explicit nil

### UnsetWarrantyPolicyId
`func (o *OrderLineUpdateDto) UnsetWarrantyPolicyId()`

UnsetWarrantyPolicyId ensures that no value is present for WarrantyPolicyId, not even an explicit nil
### GetShipmentPolicyId

`func (o *OrderLineUpdateDto) GetShipmentPolicyId() string`

GetShipmentPolicyId returns the ShipmentPolicyId field if non-nil, zero value otherwise.

### GetShipmentPolicyIdOk

`func (o *OrderLineUpdateDto) GetShipmentPolicyIdOk() (*string, bool)`

GetShipmentPolicyIdOk returns a tuple with the ShipmentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentPolicyId

`func (o *OrderLineUpdateDto) SetShipmentPolicyId(v string)`

SetShipmentPolicyId sets ShipmentPolicyId field to given value.

### HasShipmentPolicyId

`func (o *OrderLineUpdateDto) HasShipmentPolicyId() bool`

HasShipmentPolicyId returns a boolean if a field has been set.

### SetShipmentPolicyIdNil

`func (o *OrderLineUpdateDto) SetShipmentPolicyIdNil(b bool)`

 SetShipmentPolicyIdNil sets the value for ShipmentPolicyId to be an explicit nil

### UnsetShipmentPolicyId
`func (o *OrderLineUpdateDto) UnsetShipmentPolicyId()`

UnsetShipmentPolicyId ensures that no value is present for ShipmentPolicyId, not even an explicit nil
### GetShippingLocationId

`func (o *OrderLineUpdateDto) GetShippingLocationId() string`

GetShippingLocationId returns the ShippingLocationId field if non-nil, zero value otherwise.

### GetShippingLocationIdOk

`func (o *OrderLineUpdateDto) GetShippingLocationIdOk() (*string, bool)`

GetShippingLocationIdOk returns a tuple with the ShippingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocationId

`func (o *OrderLineUpdateDto) SetShippingLocationId(v string)`

SetShippingLocationId sets ShippingLocationId field to given value.

### HasShippingLocationId

`func (o *OrderLineUpdateDto) HasShippingLocationId() bool`

HasShippingLocationId returns a boolean if a field has been set.

### SetShippingLocationIdNil

`func (o *OrderLineUpdateDto) SetShippingLocationIdNil(b bool)`

 SetShippingLocationIdNil sets the value for ShippingLocationId to be an explicit nil

### UnsetShippingLocationId
`func (o *OrderLineUpdateDto) UnsetShippingLocationId()`

UnsetShippingLocationId ensures that no value is present for ShippingLocationId, not even an explicit nil
### GetLocationId

`func (o *OrderLineUpdateDto) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *OrderLineUpdateDto) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *OrderLineUpdateDto) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *OrderLineUpdateDto) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *OrderLineUpdateDto) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *OrderLineUpdateDto) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetQuoteItemRecordId

`func (o *OrderLineUpdateDto) GetQuoteItemRecordId() string`

GetQuoteItemRecordId returns the QuoteItemRecordId field if non-nil, zero value otherwise.

### GetQuoteItemRecordIdOk

`func (o *OrderLineUpdateDto) GetQuoteItemRecordIdOk() (*string, bool)`

GetQuoteItemRecordIdOk returns a tuple with the QuoteItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteItemRecordId

`func (o *OrderLineUpdateDto) SetQuoteItemRecordId(v string)`

SetQuoteItemRecordId sets QuoteItemRecordId field to given value.

### HasQuoteItemRecordId

`func (o *OrderLineUpdateDto) HasQuoteItemRecordId() bool`

HasQuoteItemRecordId returns a boolean if a field has been set.

### SetQuoteItemRecordIdNil

`func (o *OrderLineUpdateDto) SetQuoteItemRecordIdNil(b bool)`

 SetQuoteItemRecordIdNil sets the value for QuoteItemRecordId to be an explicit nil

### UnsetQuoteItemRecordId
`func (o *OrderLineUpdateDto) UnsetQuoteItemRecordId()`

UnsetQuoteItemRecordId ensures that no value is present for QuoteItemRecordId, not even an explicit nil
### GetBusinessProfileRecordId

`func (o *OrderLineUpdateDto) GetBusinessProfileRecordId() string`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *OrderLineUpdateDto) GetBusinessProfileRecordIdOk() (*string, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *OrderLineUpdateDto) SetBusinessProfileRecordId(v string)`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *OrderLineUpdateDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### SetBusinessProfileRecordIdNil

`func (o *OrderLineUpdateDto) SetBusinessProfileRecordIdNil(b bool)`

 SetBusinessProfileRecordIdNil sets the value for BusinessProfileRecordId to be an explicit nil

### UnsetBusinessProfileRecordId
`func (o *OrderLineUpdateDto) UnsetBusinessProfileRecordId()`

UnsetBusinessProfileRecordId ensures that no value is present for BusinessProfileRecordId, not even an explicit nil
### GetParentBillingItemRecordId

`func (o *OrderLineUpdateDto) GetParentBillingItemRecordId() string`

GetParentBillingItemRecordId returns the ParentBillingItemRecordId field if non-nil, zero value otherwise.

### GetParentBillingItemRecordIdOk

`func (o *OrderLineUpdateDto) GetParentBillingItemRecordIdOk() (*string, bool)`

GetParentBillingItemRecordIdOk returns a tuple with the ParentBillingItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBillingItemRecordId

`func (o *OrderLineUpdateDto) SetParentBillingItemRecordId(v string)`

SetParentBillingItemRecordId sets ParentBillingItemRecordId field to given value.

### HasParentBillingItemRecordId

`func (o *OrderLineUpdateDto) HasParentBillingItemRecordId() bool`

HasParentBillingItemRecordId returns a boolean if a field has been set.

### SetParentBillingItemRecordIdNil

`func (o *OrderLineUpdateDto) SetParentBillingItemRecordIdNil(b bool)`

 SetParentBillingItemRecordIdNil sets the value for ParentBillingItemRecordId to be an explicit nil

### UnsetParentBillingItemRecordId
`func (o *OrderLineUpdateDto) UnsetParentBillingItemRecordId()`

UnsetParentBillingItemRecordId ensures that no value is present for ParentBillingItemRecordId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


