# QuoteLineCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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
**TaxCalculationMethod** | Pointer to **string** |  | [optional] 
**CostCalculationMethod** | Pointer to **string** |  | [optional] 
**ForexRatesSnapshot** | Pointer to **NullableString** |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**TotalBaseAmountInUsd** | Pointer to **float64** |  | [optional] 
**TotalProfitInUsd** | Pointer to **float64** |  | [optional] 
**TotalDetailAmountInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseInUsd** | Pointer to **float64** |  | [optional] 
**TotalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalWithheldTaxesInUsd** | Pointer to **float64** |  | [optional] 
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
**CustomGlobalSurchargesAmountCurrencyId** | Pointer to **NullableString** |  | [optional] 
**CustomGlobalDiscountsAmount** | Pointer to **float64** |  | [optional] 
**CustomGlobalDiscountsAmountCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalDetail** | Pointer to **float64** |  | [optional] 
**TotalDetailCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalDiscountsCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalTaxBase** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalSurchargesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalProfit** | Pointer to **float64** |  | [optional] 
**TotalProfitCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalShippingCost** | Pointer to **float64** |  | [optional] 
**TotalShippingCostCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalShippingTax** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalTaxes** | Pointer to **float64** |  | [optional] 
**TotalTaxesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalWithheldTax** | Pointer to **float64** |  | [optional] 
**TotalWithheldTaxCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGlobalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscountsCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGlobalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurchargesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**TotalCurrencyId** | Pointer to **NullableString** |  | [optional] 
**ReturnPolicyId** | Pointer to **NullableString** |  | [optional] 
**RefundPolicyId** | Pointer to **NullableString** |  | [optional] 
**WarrantyPolicyId** | Pointer to **NullableString** |  | [optional] 
**ShipmentPolicyId** | Pointer to **NullableString** |  | [optional] 
**ShippingLocationId** | Pointer to **NullableString** |  | [optional] 
**LocationId** | Pointer to **NullableString** |  | [optional] 
**QuoteItemRecordId** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **NullableString** |  | [optional] 
**ParentBillingItemRecordId** | Pointer to **NullableString** |  | [optional] 
**QuoteId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQuoteLineCreateDto

`func NewQuoteLineCreateDto() *QuoteLineCreateDto`

NewQuoteLineCreateDto instantiates a new QuoteLineCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteLineCreateDtoWithDefaults

`func NewQuoteLineCreateDtoWithDefaults() *QuoteLineCreateDto`

NewQuoteLineCreateDtoWithDefaults instantiates a new QuoteLineCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteLineCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteLineCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteLineCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QuoteLineCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *QuoteLineCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *QuoteLineCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *QuoteLineCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *QuoteLineCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetClosed

`func (o *QuoteLineCreateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *QuoteLineCreateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *QuoteLineCreateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *QuoteLineCreateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetItemId

`func (o *QuoteLineCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuoteLineCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuoteLineCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuoteLineCreateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *QuoteLineCreateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *QuoteLineCreateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemTitle

`func (o *QuoteLineCreateDto) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *QuoteLineCreateDto) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *QuoteLineCreateDto) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *QuoteLineCreateDto) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *QuoteLineCreateDto) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *QuoteLineCreateDto) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetItemShortDescription

`func (o *QuoteLineCreateDto) GetItemShortDescription() string`

GetItemShortDescription returns the ItemShortDescription field if non-nil, zero value otherwise.

### GetItemShortDescriptionOk

`func (o *QuoteLineCreateDto) GetItemShortDescriptionOk() (*string, bool)`

GetItemShortDescriptionOk returns a tuple with the ItemShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemShortDescription

`func (o *QuoteLineCreateDto) SetItemShortDescription(v string)`

SetItemShortDescription sets ItemShortDescription field to given value.

### HasItemShortDescription

`func (o *QuoteLineCreateDto) HasItemShortDescription() bool`

HasItemShortDescription returns a boolean if a field has been set.

### SetItemShortDescriptionNil

`func (o *QuoteLineCreateDto) SetItemShortDescriptionNil(b bool)`

 SetItemShortDescriptionNil sets the value for ItemShortDescription to be an explicit nil

### UnsetItemShortDescription
`func (o *QuoteLineCreateDto) UnsetItemShortDescription()`

UnsetItemShortDescription ensures that no value is present for ItemShortDescription, not even an explicit nil
### GetItemPrimaryImageUrl

`func (o *QuoteLineCreateDto) GetItemPrimaryImageUrl() string`

GetItemPrimaryImageUrl returns the ItemPrimaryImageUrl field if non-nil, zero value otherwise.

### GetItemPrimaryImageUrlOk

`func (o *QuoteLineCreateDto) GetItemPrimaryImageUrlOk() (*string, bool)`

GetItemPrimaryImageUrlOk returns a tuple with the ItemPrimaryImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrimaryImageUrl

`func (o *QuoteLineCreateDto) SetItemPrimaryImageUrl(v string)`

SetItemPrimaryImageUrl sets ItemPrimaryImageUrl field to given value.

### HasItemPrimaryImageUrl

`func (o *QuoteLineCreateDto) HasItemPrimaryImageUrl() bool`

HasItemPrimaryImageUrl returns a boolean if a field has been set.

### SetItemPrimaryImageUrlNil

`func (o *QuoteLineCreateDto) SetItemPrimaryImageUrlNil(b bool)`

 SetItemPrimaryImageUrlNil sets the value for ItemPrimaryImageUrl to be an explicit nil

### UnsetItemPrimaryImageUrl
`func (o *QuoteLineCreateDto) UnsetItemPrimaryImageUrl()`

UnsetItemPrimaryImageUrl ensures that no value is present for ItemPrimaryImageUrl, not even an explicit nil
### GetShippingPolicyId

`func (o *QuoteLineCreateDto) GetShippingPolicyId() string`

GetShippingPolicyId returns the ShippingPolicyId field if non-nil, zero value otherwise.

### GetShippingPolicyIdOk

`func (o *QuoteLineCreateDto) GetShippingPolicyIdOk() (*string, bool)`

GetShippingPolicyIdOk returns a tuple with the ShippingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPolicyId

`func (o *QuoteLineCreateDto) SetShippingPolicyId(v string)`

SetShippingPolicyId sets ShippingPolicyId field to given value.

### HasShippingPolicyId

`func (o *QuoteLineCreateDto) HasShippingPolicyId() bool`

HasShippingPolicyId returns a boolean if a field has been set.

### SetShippingPolicyIdNil

`func (o *QuoteLineCreateDto) SetShippingPolicyIdNil(b bool)`

 SetShippingPolicyIdNil sets the value for ShippingPolicyId to be an explicit nil

### UnsetShippingPolicyId
`func (o *QuoteLineCreateDto) UnsetShippingPolicyId()`

UnsetShippingPolicyId ensures that no value is present for ShippingPolicyId, not even an explicit nil
### GetTenantId

`func (o *QuoteLineCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteLineCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteLineCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *QuoteLineCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *QuoteLineCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *QuoteLineCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *QuoteLineCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *QuoteLineCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *QuoteLineCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *QuoteLineCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *QuoteLineCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *QuoteLineCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCurrencyId

`func (o *QuoteLineCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuoteLineCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuoteLineCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuoteLineCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *QuoteLineCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *QuoteLineCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDescription

`func (o *QuoteLineCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuoteLineCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuoteLineCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuoteLineCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *QuoteLineCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QuoteLineCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *QuoteLineCreateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteLineCreateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteLineCreateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteLineCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetFree

`func (o *QuoteLineCreateDto) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *QuoteLineCreateDto) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *QuoteLineCreateDto) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *QuoteLineCreateDto) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetFreeReason

`func (o *QuoteLineCreateDto) GetFreeReason() string`

GetFreeReason returns the FreeReason field if non-nil, zero value otherwise.

### GetFreeReasonOk

`func (o *QuoteLineCreateDto) GetFreeReasonOk() (*string, bool)`

GetFreeReasonOk returns a tuple with the FreeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeReason

`func (o *QuoteLineCreateDto) SetFreeReason(v string)`

SetFreeReason sets FreeReason field to given value.

### HasFreeReason

`func (o *QuoteLineCreateDto) HasFreeReason() bool`

HasFreeReason returns a boolean if a field has been set.

### SetFreeReasonNil

`func (o *QuoteLineCreateDto) SetFreeReasonNil(b bool)`

 SetFreeReasonNil sets the value for FreeReason to be an explicit nil

### UnsetFreeReason
`func (o *QuoteLineCreateDto) UnsetFreeReason()`

UnsetFreeReason ensures that no value is present for FreeReason, not even an explicit nil
### GetFreeReasonCode

`func (o *QuoteLineCreateDto) GetFreeReasonCode() string`

GetFreeReasonCode returns the FreeReasonCode field if non-nil, zero value otherwise.

### GetFreeReasonCodeOk

`func (o *QuoteLineCreateDto) GetFreeReasonCodeOk() (*string, bool)`

GetFreeReasonCodeOk returns a tuple with the FreeReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeReasonCode

`func (o *QuoteLineCreateDto) SetFreeReasonCode(v string)`

SetFreeReasonCode sets FreeReasonCode field to given value.

### HasFreeReasonCode

`func (o *QuoteLineCreateDto) HasFreeReasonCode() bool`

HasFreeReasonCode returns a boolean if a field has been set.

### SetFreeReasonCodeNil

`func (o *QuoteLineCreateDto) SetFreeReasonCodeNil(b bool)`

 SetFreeReasonCodeNil sets the value for FreeReasonCode to be an explicit nil

### UnsetFreeReasonCode
`func (o *QuoteLineCreateDto) UnsetFreeReasonCode()`

UnsetFreeReasonCode ensures that no value is present for FreeReasonCode, not even an explicit nil
### GetData

`func (o *QuoteLineCreateDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuoteLineCreateDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuoteLineCreateDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *QuoteLineCreateDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *QuoteLineCreateDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *QuoteLineCreateDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *QuoteLineCreateDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *QuoteLineCreateDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *QuoteLineCreateDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *QuoteLineCreateDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *QuoteLineCreateDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *QuoteLineCreateDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *QuoteLineCreateDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *QuoteLineCreateDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *QuoteLineCreateDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *QuoteLineCreateDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *QuoteLineCreateDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *QuoteLineCreateDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *QuoteLineCreateDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *QuoteLineCreateDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *QuoteLineCreateDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *QuoteLineCreateDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *QuoteLineCreateDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *QuoteLineCreateDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *QuoteLineCreateDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *QuoteLineCreateDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *QuoteLineCreateDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *QuoteLineCreateDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *QuoteLineCreateDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *QuoteLineCreateDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *QuoteLineCreateDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *QuoteLineCreateDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *QuoteLineCreateDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *QuoteLineCreateDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *QuoteLineCreateDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *QuoteLineCreateDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *QuoteLineCreateDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *QuoteLineCreateDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *QuoteLineCreateDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *QuoteLineCreateDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *QuoteLineCreateDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *QuoteLineCreateDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *QuoteLineCreateDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *QuoteLineCreateDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *QuoteLineCreateDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *QuoteLineCreateDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *QuoteLineCreateDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *QuoteLineCreateDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *QuoteLineCreateDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *QuoteLineCreateDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *QuoteLineCreateDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *QuoteLineCreateDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *QuoteLineCreateDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *QuoteLineCreateDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *QuoteLineCreateDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *QuoteLineCreateDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *QuoteLineCreateDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *QuoteLineCreateDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *QuoteLineCreateDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *QuoteLineCreateDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *QuoteLineCreateDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *QuoteLineCreateDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *QuoteLineCreateDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *QuoteLineCreateDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *QuoteLineCreateDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *QuoteLineCreateDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *QuoteLineCreateDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *QuoteLineCreateDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *QuoteLineCreateDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *QuoteLineCreateDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *QuoteLineCreateDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *QuoteLineCreateDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *QuoteLineCreateDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *QuoteLineCreateDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *QuoteLineCreateDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *QuoteLineCreateDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *QuoteLineCreateDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *QuoteLineCreateDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *QuoteLineCreateDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *QuoteLineCreateDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *QuoteLineCreateDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *QuoteLineCreateDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *QuoteLineCreateDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *QuoteLineCreateDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *QuoteLineCreateDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *QuoteLineCreateDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *QuoteLineCreateDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *QuoteLineCreateDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *QuoteLineCreateDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *QuoteLineCreateDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *QuoteLineCreateDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *QuoteLineCreateDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *QuoteLineCreateDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *QuoteLineCreateDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *QuoteLineCreateDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *QuoteLineCreateDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *QuoteLineCreateDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *QuoteLineCreateDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *QuoteLineCreateDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *QuoteLineCreateDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *QuoteLineCreateDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *QuoteLineCreateDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *QuoteLineCreateDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *QuoteLineCreateDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *QuoteLineCreateDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *QuoteLineCreateDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *QuoteLineCreateDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *QuoteLineCreateDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *QuoteLineCreateDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *QuoteLineCreateDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *QuoteLineCreateDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *QuoteLineCreateDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *QuoteLineCreateDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *QuoteLineCreateDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *QuoteLineCreateDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *QuoteLineCreateDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *QuoteLineCreateDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *QuoteLineCreateDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *QuoteLineCreateDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *QuoteLineCreateDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetItemPriceId

`func (o *QuoteLineCreateDto) GetItemPriceId() string`

GetItemPriceId returns the ItemPriceId field if non-nil, zero value otherwise.

### GetItemPriceIdOk

`func (o *QuoteLineCreateDto) GetItemPriceIdOk() (*string, bool)`

GetItemPriceIdOk returns a tuple with the ItemPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceId

`func (o *QuoteLineCreateDto) SetItemPriceId(v string)`

SetItemPriceId sets ItemPriceId field to given value.

### HasItemPriceId

`func (o *QuoteLineCreateDto) HasItemPriceId() bool`

HasItemPriceId returns a boolean if a field has been set.

### SetItemPriceIdNil

`func (o *QuoteLineCreateDto) SetItemPriceIdNil(b bool)`

 SetItemPriceIdNil sets the value for ItemPriceId to be an explicit nil

### UnsetItemPriceId
`func (o *QuoteLineCreateDto) UnsetItemPriceId()`

UnsetItemPriceId ensures that no value is present for ItemPriceId, not even an explicit nil
### GetPriceListItemId

`func (o *QuoteLineCreateDto) GetPriceListItemId() string`

GetPriceListItemId returns the PriceListItemId field if non-nil, zero value otherwise.

### GetPriceListItemIdOk

`func (o *QuoteLineCreateDto) GetPriceListItemIdOk() (*string, bool)`

GetPriceListItemIdOk returns a tuple with the PriceListItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListItemId

`func (o *QuoteLineCreateDto) SetPriceListItemId(v string)`

SetPriceListItemId sets PriceListItemId field to given value.

### HasPriceListItemId

`func (o *QuoteLineCreateDto) HasPriceListItemId() bool`

HasPriceListItemId returns a boolean if a field has been set.

### SetPriceListItemIdNil

`func (o *QuoteLineCreateDto) SetPriceListItemIdNil(b bool)`

 SetPriceListItemIdNil sets the value for PriceListItemId to be an explicit nil

### UnsetPriceListItemId
`func (o *QuoteLineCreateDto) UnsetPriceListItemId()`

UnsetPriceListItemId ensures that no value is present for PriceListItemId, not even an explicit nil
### GetUnitId

`func (o *QuoteLineCreateDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *QuoteLineCreateDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *QuoteLineCreateDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *QuoteLineCreateDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *QuoteLineCreateDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *QuoteLineCreateDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *QuoteLineCreateDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *QuoteLineCreateDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *QuoteLineCreateDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *QuoteLineCreateDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *QuoteLineCreateDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *QuoteLineCreateDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetTaxCalculationMethod

`func (o *QuoteLineCreateDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *QuoteLineCreateDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *QuoteLineCreateDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *QuoteLineCreateDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *QuoteLineCreateDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *QuoteLineCreateDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *QuoteLineCreateDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *QuoteLineCreateDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetForexRatesSnapshot

`func (o *QuoteLineCreateDto) GetForexRatesSnapshot() string`

GetForexRatesSnapshot returns the ForexRatesSnapshot field if non-nil, zero value otherwise.

### GetForexRatesSnapshotOk

`func (o *QuoteLineCreateDto) GetForexRatesSnapshotOk() (*string, bool)`

GetForexRatesSnapshotOk returns a tuple with the ForexRatesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRatesSnapshot

`func (o *QuoteLineCreateDto) SetForexRatesSnapshot(v string)`

SetForexRatesSnapshot sets ForexRatesSnapshot field to given value.

### HasForexRatesSnapshot

`func (o *QuoteLineCreateDto) HasForexRatesSnapshot() bool`

HasForexRatesSnapshot returns a boolean if a field has been set.

### SetForexRatesSnapshotNil

`func (o *QuoteLineCreateDto) SetForexRatesSnapshotNil(b bool)`

 SetForexRatesSnapshotNil sets the value for ForexRatesSnapshot to be an explicit nil

### UnsetForexRatesSnapshot
`func (o *QuoteLineCreateDto) UnsetForexRatesSnapshot()`

UnsetForexRatesSnapshot ensures that no value is present for ForexRatesSnapshot, not even an explicit nil
### GetForexRate

`func (o *QuoteLineCreateDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *QuoteLineCreateDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *QuoteLineCreateDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *QuoteLineCreateDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetTotalBaseAmountInUsd

`func (o *QuoteLineCreateDto) GetTotalBaseAmountInUsd() float64`

GetTotalBaseAmountInUsd returns the TotalBaseAmountInUsd field if non-nil, zero value otherwise.

### GetTotalBaseAmountInUsdOk

`func (o *QuoteLineCreateDto) GetTotalBaseAmountInUsdOk() (*float64, bool)`

GetTotalBaseAmountInUsdOk returns a tuple with the TotalBaseAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBaseAmountInUsd

`func (o *QuoteLineCreateDto) SetTotalBaseAmountInUsd(v float64)`

SetTotalBaseAmountInUsd sets TotalBaseAmountInUsd field to given value.

### HasTotalBaseAmountInUsd

`func (o *QuoteLineCreateDto) HasTotalBaseAmountInUsd() bool`

HasTotalBaseAmountInUsd returns a boolean if a field has been set.

### GetTotalProfitInUsd

`func (o *QuoteLineCreateDto) GetTotalProfitInUsd() float64`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *QuoteLineCreateDto) GetTotalProfitInUsdOk() (*float64, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *QuoteLineCreateDto) SetTotalProfitInUsd(v float64)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.

### HasTotalProfitInUsd

`func (o *QuoteLineCreateDto) HasTotalProfitInUsd() bool`

HasTotalProfitInUsd returns a boolean if a field has been set.

### GetTotalDetailAmountInUsd

`func (o *QuoteLineCreateDto) GetTotalDetailAmountInUsd() float64`

GetTotalDetailAmountInUsd returns the TotalDetailAmountInUsd field if non-nil, zero value otherwise.

### GetTotalDetailAmountInUsdOk

`func (o *QuoteLineCreateDto) GetTotalDetailAmountInUsdOk() (*float64, bool)`

GetTotalDetailAmountInUsdOk returns a tuple with the TotalDetailAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailAmountInUsd

`func (o *QuoteLineCreateDto) SetTotalDetailAmountInUsd(v float64)`

SetTotalDetailAmountInUsd sets TotalDetailAmountInUsd field to given value.

### HasTotalDetailAmountInUsd

`func (o *QuoteLineCreateDto) HasTotalDetailAmountInUsd() bool`

HasTotalDetailAmountInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseInUsd

`func (o *QuoteLineCreateDto) GetTotalTaxBaseInUsd() float64`

GetTotalTaxBaseInUsd returns the TotalTaxBaseInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseInUsdOk

`func (o *QuoteLineCreateDto) GetTotalTaxBaseInUsdOk() (*float64, bool)`

GetTotalTaxBaseInUsdOk returns a tuple with the TotalTaxBaseInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseInUsd

`func (o *QuoteLineCreateDto) SetTotalTaxBaseInUsd(v float64)`

SetTotalTaxBaseInUsd sets TotalTaxBaseInUsd field to given value.

### HasTotalTaxBaseInUsd

`func (o *QuoteLineCreateDto) HasTotalTaxBaseInUsd() bool`

HasTotalTaxBaseInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *QuoteLineCreateDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *QuoteLineCreateDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *QuoteLineCreateDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *QuoteLineCreateDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalTaxesInUsd

`func (o *QuoteLineCreateDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *QuoteLineCreateDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *QuoteLineCreateDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *QuoteLineCreateDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalWithheldTaxesInUsd

`func (o *QuoteLineCreateDto) GetTotalWithheldTaxesInUsd() float64`

GetTotalWithheldTaxesInUsd returns the TotalWithheldTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithheldTaxesInUsdOk

`func (o *QuoteLineCreateDto) GetTotalWithheldTaxesInUsdOk() (*float64, bool)`

GetTotalWithheldTaxesInUsdOk returns a tuple with the TotalWithheldTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxesInUsd

`func (o *QuoteLineCreateDto) SetTotalWithheldTaxesInUsd(v float64)`

SetTotalWithheldTaxesInUsd sets TotalWithheldTaxesInUsd field to given value.

### HasTotalWithheldTaxesInUsd

`func (o *QuoteLineCreateDto) HasTotalWithheldTaxesInUsd() bool`

HasTotalWithheldTaxesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *QuoteLineCreateDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *QuoteLineCreateDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *QuoteLineCreateDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *QuoteLineCreateDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxesInUsd

`func (o *QuoteLineCreateDto) GetTotalShippingTaxesInUsd() float64`

GetTotalShippingTaxesInUsd returns the TotalShippingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxesInUsdOk

`func (o *QuoteLineCreateDto) GetTotalShippingTaxesInUsdOk() (*float64, bool)`

GetTotalShippingTaxesInUsdOk returns a tuple with the TotalShippingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxesInUsd

`func (o *QuoteLineCreateDto) SetTotalShippingTaxesInUsd(v float64)`

SetTotalShippingTaxesInUsd sets TotalShippingTaxesInUsd field to given value.

### HasTotalShippingTaxesInUsd

`func (o *QuoteLineCreateDto) HasTotalShippingTaxesInUsd() bool`

HasTotalShippingTaxesInUsd returns a boolean if a field has been set.

### GetTotalWarrantyCostInUsd

`func (o *QuoteLineCreateDto) GetTotalWarrantyCostInUsd() float64`

GetTotalWarrantyCostInUsd returns the TotalWarrantyCostInUsd field if non-nil, zero value otherwise.

### GetTotalWarrantyCostInUsdOk

`func (o *QuoteLineCreateDto) GetTotalWarrantyCostInUsdOk() (*float64, bool)`

GetTotalWarrantyCostInUsdOk returns a tuple with the TotalWarrantyCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWarrantyCostInUsd

`func (o *QuoteLineCreateDto) SetTotalWarrantyCostInUsd(v float64)`

SetTotalWarrantyCostInUsd sets TotalWarrantyCostInUsd field to given value.

### HasTotalWarrantyCostInUsd

`func (o *QuoteLineCreateDto) HasTotalWarrantyCostInUsd() bool`

HasTotalWarrantyCostInUsd returns a boolean if a field has been set.

### GetTotalReturnCostInUsd

`func (o *QuoteLineCreateDto) GetTotalReturnCostInUsd() float64`

GetTotalReturnCostInUsd returns the TotalReturnCostInUsd field if non-nil, zero value otherwise.

### GetTotalReturnCostInUsdOk

`func (o *QuoteLineCreateDto) GetTotalReturnCostInUsdOk() (*float64, bool)`

GetTotalReturnCostInUsdOk returns a tuple with the TotalReturnCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReturnCostInUsd

`func (o *QuoteLineCreateDto) SetTotalReturnCostInUsd(v float64)`

SetTotalReturnCostInUsd sets TotalReturnCostInUsd field to given value.

### HasTotalReturnCostInUsd

`func (o *QuoteLineCreateDto) HasTotalReturnCostInUsd() bool`

HasTotalReturnCostInUsd returns a boolean if a field has been set.

### GetTotalRefundCostInUsd

`func (o *QuoteLineCreateDto) GetTotalRefundCostInUsd() float64`

GetTotalRefundCostInUsd returns the TotalRefundCostInUsd field if non-nil, zero value otherwise.

### GetTotalRefundCostInUsdOk

`func (o *QuoteLineCreateDto) GetTotalRefundCostInUsdOk() (*float64, bool)`

GetTotalRefundCostInUsdOk returns a tuple with the TotalRefundCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRefundCostInUsd

`func (o *QuoteLineCreateDto) SetTotalRefundCostInUsd(v float64)`

SetTotalRefundCostInUsd sets TotalRefundCostInUsd field to given value.

### HasTotalRefundCostInUsd

`func (o *QuoteLineCreateDto) HasTotalRefundCostInUsd() bool`

HasTotalRefundCostInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *QuoteLineCreateDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *QuoteLineCreateDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *QuoteLineCreateDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *QuoteLineCreateDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalAmountInUsd

`func (o *QuoteLineCreateDto) GetTotalAmountInUsd() float64`

GetTotalAmountInUsd returns the TotalAmountInUsd field if non-nil, zero value otherwise.

### GetTotalAmountInUsdOk

`func (o *QuoteLineCreateDto) GetTotalAmountInUsdOk() (*float64, bool)`

GetTotalAmountInUsdOk returns a tuple with the TotalAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInUsd

`func (o *QuoteLineCreateDto) SetTotalAmountInUsd(v float64)`

SetTotalAmountInUsd sets TotalAmountInUsd field to given value.

### HasTotalAmountInUsd

`func (o *QuoteLineCreateDto) HasTotalAmountInUsd() bool`

HasTotalAmountInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsInUsd

`func (o *QuoteLineCreateDto) GetTotalGlobalDiscountsInUsd() float64`

GetTotalGlobalDiscountsInUsd returns the TotalGlobalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsInUsdOk

`func (o *QuoteLineCreateDto) GetTotalGlobalDiscountsInUsdOk() (*float64, bool)`

GetTotalGlobalDiscountsInUsdOk returns a tuple with the TotalGlobalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsInUsd

`func (o *QuoteLineCreateDto) SetTotalGlobalDiscountsInUsd(v float64)`

SetTotalGlobalDiscountsInUsd sets TotalGlobalDiscountsInUsd field to given value.

### HasTotalGlobalDiscountsInUsd

`func (o *QuoteLineCreateDto) HasTotalGlobalDiscountsInUsd() bool`

HasTotalGlobalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesInUsd

`func (o *QuoteLineCreateDto) GetTotalGlobalSurchargesInUsd() float64`

GetTotalGlobalSurchargesInUsd returns the TotalGlobalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesInUsdOk

`func (o *QuoteLineCreateDto) GetTotalGlobalSurchargesInUsdOk() (*float64, bool)`

GetTotalGlobalSurchargesInUsdOk returns a tuple with the TotalGlobalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesInUsd

`func (o *QuoteLineCreateDto) SetTotalGlobalSurchargesInUsd(v float64)`

SetTotalGlobalSurchargesInUsd sets TotalGlobalSurchargesInUsd field to given value.

### HasTotalGlobalSurchargesInUsd

`func (o *QuoteLineCreateDto) HasTotalGlobalSurchargesInUsd() bool`

HasTotalGlobalSurchargesInUsd returns a boolean if a field has been set.

### GetCustomGlobalSurchargesAmount

`func (o *QuoteLineCreateDto) GetCustomGlobalSurchargesAmount() float64`

GetCustomGlobalSurchargesAmount returns the CustomGlobalSurchargesAmount field if non-nil, zero value otherwise.

### GetCustomGlobalSurchargesAmountOk

`func (o *QuoteLineCreateDto) GetCustomGlobalSurchargesAmountOk() (*float64, bool)`

GetCustomGlobalSurchargesAmountOk returns a tuple with the CustomGlobalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalSurchargesAmount

`func (o *QuoteLineCreateDto) SetCustomGlobalSurchargesAmount(v float64)`

SetCustomGlobalSurchargesAmount sets CustomGlobalSurchargesAmount field to given value.

### HasCustomGlobalSurchargesAmount

`func (o *QuoteLineCreateDto) HasCustomGlobalSurchargesAmount() bool`

HasCustomGlobalSurchargesAmount returns a boolean if a field has been set.

### GetCustomGlobalSurchargesAmountCurrencyId

`func (o *QuoteLineCreateDto) GetCustomGlobalSurchargesAmountCurrencyId() string`

GetCustomGlobalSurchargesAmountCurrencyId returns the CustomGlobalSurchargesAmountCurrencyId field if non-nil, zero value otherwise.

### GetCustomGlobalSurchargesAmountCurrencyIdOk

`func (o *QuoteLineCreateDto) GetCustomGlobalSurchargesAmountCurrencyIdOk() (*string, bool)`

GetCustomGlobalSurchargesAmountCurrencyIdOk returns a tuple with the CustomGlobalSurchargesAmountCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalSurchargesAmountCurrencyId

`func (o *QuoteLineCreateDto) SetCustomGlobalSurchargesAmountCurrencyId(v string)`

SetCustomGlobalSurchargesAmountCurrencyId sets CustomGlobalSurchargesAmountCurrencyId field to given value.

### HasCustomGlobalSurchargesAmountCurrencyId

`func (o *QuoteLineCreateDto) HasCustomGlobalSurchargesAmountCurrencyId() bool`

HasCustomGlobalSurchargesAmountCurrencyId returns a boolean if a field has been set.

### SetCustomGlobalSurchargesAmountCurrencyIdNil

`func (o *QuoteLineCreateDto) SetCustomGlobalSurchargesAmountCurrencyIdNil(b bool)`

 SetCustomGlobalSurchargesAmountCurrencyIdNil sets the value for CustomGlobalSurchargesAmountCurrencyId to be an explicit nil

### UnsetCustomGlobalSurchargesAmountCurrencyId
`func (o *QuoteLineCreateDto) UnsetCustomGlobalSurchargesAmountCurrencyId()`

UnsetCustomGlobalSurchargesAmountCurrencyId ensures that no value is present for CustomGlobalSurchargesAmountCurrencyId, not even an explicit nil
### GetCustomGlobalDiscountsAmount

`func (o *QuoteLineCreateDto) GetCustomGlobalDiscountsAmount() float64`

GetCustomGlobalDiscountsAmount returns the CustomGlobalDiscountsAmount field if non-nil, zero value otherwise.

### GetCustomGlobalDiscountsAmountOk

`func (o *QuoteLineCreateDto) GetCustomGlobalDiscountsAmountOk() (*float64, bool)`

GetCustomGlobalDiscountsAmountOk returns a tuple with the CustomGlobalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalDiscountsAmount

`func (o *QuoteLineCreateDto) SetCustomGlobalDiscountsAmount(v float64)`

SetCustomGlobalDiscountsAmount sets CustomGlobalDiscountsAmount field to given value.

### HasCustomGlobalDiscountsAmount

`func (o *QuoteLineCreateDto) HasCustomGlobalDiscountsAmount() bool`

HasCustomGlobalDiscountsAmount returns a boolean if a field has been set.

### GetCustomGlobalDiscountsAmountCurrencyId

`func (o *QuoteLineCreateDto) GetCustomGlobalDiscountsAmountCurrencyId() string`

GetCustomGlobalDiscountsAmountCurrencyId returns the CustomGlobalDiscountsAmountCurrencyId field if non-nil, zero value otherwise.

### GetCustomGlobalDiscountsAmountCurrencyIdOk

`func (o *QuoteLineCreateDto) GetCustomGlobalDiscountsAmountCurrencyIdOk() (*string, bool)`

GetCustomGlobalDiscountsAmountCurrencyIdOk returns a tuple with the CustomGlobalDiscountsAmountCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalDiscountsAmountCurrencyId

`func (o *QuoteLineCreateDto) SetCustomGlobalDiscountsAmountCurrencyId(v string)`

SetCustomGlobalDiscountsAmountCurrencyId sets CustomGlobalDiscountsAmountCurrencyId field to given value.

### HasCustomGlobalDiscountsAmountCurrencyId

`func (o *QuoteLineCreateDto) HasCustomGlobalDiscountsAmountCurrencyId() bool`

HasCustomGlobalDiscountsAmountCurrencyId returns a boolean if a field has been set.

### SetCustomGlobalDiscountsAmountCurrencyIdNil

`func (o *QuoteLineCreateDto) SetCustomGlobalDiscountsAmountCurrencyIdNil(b bool)`

 SetCustomGlobalDiscountsAmountCurrencyIdNil sets the value for CustomGlobalDiscountsAmountCurrencyId to be an explicit nil

### UnsetCustomGlobalDiscountsAmountCurrencyId
`func (o *QuoteLineCreateDto) UnsetCustomGlobalDiscountsAmountCurrencyId()`

UnsetCustomGlobalDiscountsAmountCurrencyId ensures that no value is present for CustomGlobalDiscountsAmountCurrencyId, not even an explicit nil
### GetTotalDetail

`func (o *QuoteLineCreateDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *QuoteLineCreateDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *QuoteLineCreateDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *QuoteLineCreateDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *QuoteLineCreateDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *QuoteLineCreateDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *QuoteLineCreateDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalDiscounts

`func (o *QuoteLineCreateDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *QuoteLineCreateDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *QuoteLineCreateDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *QuoteLineCreateDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *QuoteLineCreateDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *QuoteLineCreateDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *QuoteLineCreateDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalTaxBase

`func (o *QuoteLineCreateDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *QuoteLineCreateDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *QuoteLineCreateDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *QuoteLineCreateDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *QuoteLineCreateDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *QuoteLineCreateDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *QuoteLineCreateDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalSurcharges

`func (o *QuoteLineCreateDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *QuoteLineCreateDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *QuoteLineCreateDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *QuoteLineCreateDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *QuoteLineCreateDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *QuoteLineCreateDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *QuoteLineCreateDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalProfit

`func (o *QuoteLineCreateDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *QuoteLineCreateDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *QuoteLineCreateDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *QuoteLineCreateDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *QuoteLineCreateDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *QuoteLineCreateDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *QuoteLineCreateDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalShippingCost

`func (o *QuoteLineCreateDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *QuoteLineCreateDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *QuoteLineCreateDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *QuoteLineCreateDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *QuoteLineCreateDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *QuoteLineCreateDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *QuoteLineCreateDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingTax

`func (o *QuoteLineCreateDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *QuoteLineCreateDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *QuoteLineCreateDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *QuoteLineCreateDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *QuoteLineCreateDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *QuoteLineCreateDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *QuoteLineCreateDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalTaxes

`func (o *QuoteLineCreateDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *QuoteLineCreateDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *QuoteLineCreateDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *QuoteLineCreateDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *QuoteLineCreateDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *QuoteLineCreateDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *QuoteLineCreateDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalWithheldTax

`func (o *QuoteLineCreateDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *QuoteLineCreateDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *QuoteLineCreateDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *QuoteLineCreateDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *QuoteLineCreateDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *QuoteLineCreateDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *QuoteLineCreateDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalGlobalDiscounts

`func (o *QuoteLineCreateDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *QuoteLineCreateDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *QuoteLineCreateDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *QuoteLineCreateDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *QuoteLineCreateDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *QuoteLineCreateDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *QuoteLineCreateDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotalGlobalSurcharges

`func (o *QuoteLineCreateDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *QuoteLineCreateDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *QuoteLineCreateDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *QuoteLineCreateDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *QuoteLineCreateDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *QuoteLineCreateDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *QuoteLineCreateDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotal

`func (o *QuoteLineCreateDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QuoteLineCreateDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QuoteLineCreateDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QuoteLineCreateDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *QuoteLineCreateDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *QuoteLineCreateDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *QuoteLineCreateDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *QuoteLineCreateDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *QuoteLineCreateDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *QuoteLineCreateDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetReturnPolicyId

`func (o *QuoteLineCreateDto) GetReturnPolicyId() string`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *QuoteLineCreateDto) GetReturnPolicyIdOk() (*string, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *QuoteLineCreateDto) SetReturnPolicyId(v string)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *QuoteLineCreateDto) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### SetReturnPolicyIdNil

`func (o *QuoteLineCreateDto) SetReturnPolicyIdNil(b bool)`

 SetReturnPolicyIdNil sets the value for ReturnPolicyId to be an explicit nil

### UnsetReturnPolicyId
`func (o *QuoteLineCreateDto) UnsetReturnPolicyId()`

UnsetReturnPolicyId ensures that no value is present for ReturnPolicyId, not even an explicit nil
### GetRefundPolicyId

`func (o *QuoteLineCreateDto) GetRefundPolicyId() string`

GetRefundPolicyId returns the RefundPolicyId field if non-nil, zero value otherwise.

### GetRefundPolicyIdOk

`func (o *QuoteLineCreateDto) GetRefundPolicyIdOk() (*string, bool)`

GetRefundPolicyIdOk returns a tuple with the RefundPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPolicyId

`func (o *QuoteLineCreateDto) SetRefundPolicyId(v string)`

SetRefundPolicyId sets RefundPolicyId field to given value.

### HasRefundPolicyId

`func (o *QuoteLineCreateDto) HasRefundPolicyId() bool`

HasRefundPolicyId returns a boolean if a field has been set.

### SetRefundPolicyIdNil

`func (o *QuoteLineCreateDto) SetRefundPolicyIdNil(b bool)`

 SetRefundPolicyIdNil sets the value for RefundPolicyId to be an explicit nil

### UnsetRefundPolicyId
`func (o *QuoteLineCreateDto) UnsetRefundPolicyId()`

UnsetRefundPolicyId ensures that no value is present for RefundPolicyId, not even an explicit nil
### GetWarrantyPolicyId

`func (o *QuoteLineCreateDto) GetWarrantyPolicyId() string`

GetWarrantyPolicyId returns the WarrantyPolicyId field if non-nil, zero value otherwise.

### GetWarrantyPolicyIdOk

`func (o *QuoteLineCreateDto) GetWarrantyPolicyIdOk() (*string, bool)`

GetWarrantyPolicyIdOk returns a tuple with the WarrantyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyPolicyId

`func (o *QuoteLineCreateDto) SetWarrantyPolicyId(v string)`

SetWarrantyPolicyId sets WarrantyPolicyId field to given value.

### HasWarrantyPolicyId

`func (o *QuoteLineCreateDto) HasWarrantyPolicyId() bool`

HasWarrantyPolicyId returns a boolean if a field has been set.

### SetWarrantyPolicyIdNil

`func (o *QuoteLineCreateDto) SetWarrantyPolicyIdNil(b bool)`

 SetWarrantyPolicyIdNil sets the value for WarrantyPolicyId to be an explicit nil

### UnsetWarrantyPolicyId
`func (o *QuoteLineCreateDto) UnsetWarrantyPolicyId()`

UnsetWarrantyPolicyId ensures that no value is present for WarrantyPolicyId, not even an explicit nil
### GetShipmentPolicyId

`func (o *QuoteLineCreateDto) GetShipmentPolicyId() string`

GetShipmentPolicyId returns the ShipmentPolicyId field if non-nil, zero value otherwise.

### GetShipmentPolicyIdOk

`func (o *QuoteLineCreateDto) GetShipmentPolicyIdOk() (*string, bool)`

GetShipmentPolicyIdOk returns a tuple with the ShipmentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentPolicyId

`func (o *QuoteLineCreateDto) SetShipmentPolicyId(v string)`

SetShipmentPolicyId sets ShipmentPolicyId field to given value.

### HasShipmentPolicyId

`func (o *QuoteLineCreateDto) HasShipmentPolicyId() bool`

HasShipmentPolicyId returns a boolean if a field has been set.

### SetShipmentPolicyIdNil

`func (o *QuoteLineCreateDto) SetShipmentPolicyIdNil(b bool)`

 SetShipmentPolicyIdNil sets the value for ShipmentPolicyId to be an explicit nil

### UnsetShipmentPolicyId
`func (o *QuoteLineCreateDto) UnsetShipmentPolicyId()`

UnsetShipmentPolicyId ensures that no value is present for ShipmentPolicyId, not even an explicit nil
### GetShippingLocationId

`func (o *QuoteLineCreateDto) GetShippingLocationId() string`

GetShippingLocationId returns the ShippingLocationId field if non-nil, zero value otherwise.

### GetShippingLocationIdOk

`func (o *QuoteLineCreateDto) GetShippingLocationIdOk() (*string, bool)`

GetShippingLocationIdOk returns a tuple with the ShippingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocationId

`func (o *QuoteLineCreateDto) SetShippingLocationId(v string)`

SetShippingLocationId sets ShippingLocationId field to given value.

### HasShippingLocationId

`func (o *QuoteLineCreateDto) HasShippingLocationId() bool`

HasShippingLocationId returns a boolean if a field has been set.

### SetShippingLocationIdNil

`func (o *QuoteLineCreateDto) SetShippingLocationIdNil(b bool)`

 SetShippingLocationIdNil sets the value for ShippingLocationId to be an explicit nil

### UnsetShippingLocationId
`func (o *QuoteLineCreateDto) UnsetShippingLocationId()`

UnsetShippingLocationId ensures that no value is present for ShippingLocationId, not even an explicit nil
### GetLocationId

`func (o *QuoteLineCreateDto) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *QuoteLineCreateDto) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *QuoteLineCreateDto) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *QuoteLineCreateDto) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *QuoteLineCreateDto) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *QuoteLineCreateDto) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetQuoteItemRecordId

`func (o *QuoteLineCreateDto) GetQuoteItemRecordId() string`

GetQuoteItemRecordId returns the QuoteItemRecordId field if non-nil, zero value otherwise.

### GetQuoteItemRecordIdOk

`func (o *QuoteLineCreateDto) GetQuoteItemRecordIdOk() (*string, bool)`

GetQuoteItemRecordIdOk returns a tuple with the QuoteItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteItemRecordId

`func (o *QuoteLineCreateDto) SetQuoteItemRecordId(v string)`

SetQuoteItemRecordId sets QuoteItemRecordId field to given value.

### HasQuoteItemRecordId

`func (o *QuoteLineCreateDto) HasQuoteItemRecordId() bool`

HasQuoteItemRecordId returns a boolean if a field has been set.

### SetQuoteItemRecordIdNil

`func (o *QuoteLineCreateDto) SetQuoteItemRecordIdNil(b bool)`

 SetQuoteItemRecordIdNil sets the value for QuoteItemRecordId to be an explicit nil

### UnsetQuoteItemRecordId
`func (o *QuoteLineCreateDto) UnsetQuoteItemRecordId()`

UnsetQuoteItemRecordId ensures that no value is present for QuoteItemRecordId, not even an explicit nil
### GetBusinessProfileRecordId

`func (o *QuoteLineCreateDto) GetBusinessProfileRecordId() string`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *QuoteLineCreateDto) GetBusinessProfileRecordIdOk() (*string, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *QuoteLineCreateDto) SetBusinessProfileRecordId(v string)`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *QuoteLineCreateDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### SetBusinessProfileRecordIdNil

`func (o *QuoteLineCreateDto) SetBusinessProfileRecordIdNil(b bool)`

 SetBusinessProfileRecordIdNil sets the value for BusinessProfileRecordId to be an explicit nil

### UnsetBusinessProfileRecordId
`func (o *QuoteLineCreateDto) UnsetBusinessProfileRecordId()`

UnsetBusinessProfileRecordId ensures that no value is present for BusinessProfileRecordId, not even an explicit nil
### GetParentBillingItemRecordId

`func (o *QuoteLineCreateDto) GetParentBillingItemRecordId() string`

GetParentBillingItemRecordId returns the ParentBillingItemRecordId field if non-nil, zero value otherwise.

### GetParentBillingItemRecordIdOk

`func (o *QuoteLineCreateDto) GetParentBillingItemRecordIdOk() (*string, bool)`

GetParentBillingItemRecordIdOk returns a tuple with the ParentBillingItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBillingItemRecordId

`func (o *QuoteLineCreateDto) SetParentBillingItemRecordId(v string)`

SetParentBillingItemRecordId sets ParentBillingItemRecordId field to given value.

### HasParentBillingItemRecordId

`func (o *QuoteLineCreateDto) HasParentBillingItemRecordId() bool`

HasParentBillingItemRecordId returns a boolean if a field has been set.

### SetParentBillingItemRecordIdNil

`func (o *QuoteLineCreateDto) SetParentBillingItemRecordIdNil(b bool)`

 SetParentBillingItemRecordIdNil sets the value for ParentBillingItemRecordId to be an explicit nil

### UnsetParentBillingItemRecordId
`func (o *QuoteLineCreateDto) UnsetParentBillingItemRecordId()`

UnsetParentBillingItemRecordId ensures that no value is present for ParentBillingItemRecordId, not even an explicit nil
### GetQuoteId

`func (o *QuoteLineCreateDto) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteLineCreateDto) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteLineCreateDto) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *QuoteLineCreateDto) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### SetQuoteIdNil

`func (o *QuoteLineCreateDto) SetQuoteIdNil(b bool)`

 SetQuoteIdNil sets the value for QuoteId to be an explicit nil

### UnsetQuoteId
`func (o *QuoteLineCreateDto) UnsetQuoteId()`

UnsetQuoteId ensures that no value is present for QuoteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


