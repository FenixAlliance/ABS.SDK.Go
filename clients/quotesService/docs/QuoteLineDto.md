# QuoteLineDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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
**ForexRates** | Pointer to [**ForexRates**](ForexRates.md) |  | [optional] 
**ForexRate** | Pointer to **float64** |  | [optional] 
**TotalDetailInUsd** | Pointer to **float64** |  | [optional] 
**TotalProfitInUsd** | Pointer to **float64** |  | [optional] 
**TotalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseInUsd** | Pointer to **float64** |  | [optional] 
**TotalTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalWithheldTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxesInUsd** | Pointer to **float64** |  | [optional] 
**TotalWarrantyCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalReturnCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalRefundCostInUsd** | Pointer to **float64** |  | [optional] 
**TotalInUsd** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscountsInUsd** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurchargesInUsd** | Pointer to **float64** |  | [optional] 
**CustomGlobalSurchargesAmount** | Pointer to **float64** |  | [optional] 
**CustomGlobalDiscountsAmount** | Pointer to **float64** |  | [optional] 
**ReturnPolicyId** | Pointer to **NullableString** |  | [optional] 
**RefundPolicyId** | Pointer to **NullableString** |  | [optional] 
**WarrantyPolicyId** | Pointer to **NullableString** |  | [optional] 
**ShipmentPolicyId** | Pointer to **NullableString** |  | [optional] 
**ShippingLocationId** | Pointer to **NullableString** |  | [optional] 
**LocationId** | Pointer to **NullableString** |  | [optional] 
**QuoteItemRecordId** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordId** | Pointer to **NullableString** |  | [optional] 
**ParentBillingItemRecordId** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to [**CurrencyId**](CurrencyId.md) |  | [optional] 
**TotalDetail** | Pointer to **float64** |  | [optional] 
**TotalDetailCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalDetailAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalProfit** | Pointer to **float64** |  | [optional] 
**TotalProfitCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalProfitAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalDiscountsCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalDiscountsAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalSurchargesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalSurchargesAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxBase** | Pointer to **float64** |  | [optional] 
**TotalTaxBaseCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalTaxBaseAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalTaxes** | Pointer to **float64** |  | [optional] 
**TotalTaxesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalTaxesAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalShippingCost** | Pointer to **float64** |  | [optional] 
**TotalShippingCostCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalShippingCostAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalShippingTax** | Pointer to **float64** |  | [optional] 
**TotalShippingTaxCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalShippingTaxAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalWithheldTax** | Pointer to **float64** |  | [optional] 
**TotalWithheldTaxCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalWithheldTaxAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalGlobalDiscounts** | Pointer to **float64** |  | [optional] 
**TotalGlobalDiscountsCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGlobalDiscountsAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**TotalGlobalSurcharges** | Pointer to **float64** |  | [optional] 
**TotalGlobalSurchargesCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalGlobalSurchargesAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**Total** | Pointer to **float64** |  | [optional] 
**TotalCurrencyId** | Pointer to **NullableString** |  | [optional] 
**TotalAmount** | Pointer to [**Money**](Money.md) |  | [optional] 
**QuoteId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQuoteLineDto

`func NewQuoteLineDto() *QuoteLineDto`

NewQuoteLineDto instantiates a new QuoteLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteLineDtoWithDefaults

`func NewQuoteLineDtoWithDefaults() *QuoteLineDto`

NewQuoteLineDtoWithDefaults instantiates a new QuoteLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteLineDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteLineDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteLineDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QuoteLineDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *QuoteLineDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *QuoteLineDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *QuoteLineDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *QuoteLineDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *QuoteLineDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *QuoteLineDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *QuoteLineDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *QuoteLineDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetClosed

`func (o *QuoteLineDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *QuoteLineDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *QuoteLineDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *QuoteLineDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetItemId

`func (o *QuoteLineDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuoteLineDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuoteLineDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuoteLineDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *QuoteLineDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *QuoteLineDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemTitle

`func (o *QuoteLineDto) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *QuoteLineDto) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *QuoteLineDto) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *QuoteLineDto) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *QuoteLineDto) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *QuoteLineDto) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetItemShortDescription

`func (o *QuoteLineDto) GetItemShortDescription() string`

GetItemShortDescription returns the ItemShortDescription field if non-nil, zero value otherwise.

### GetItemShortDescriptionOk

`func (o *QuoteLineDto) GetItemShortDescriptionOk() (*string, bool)`

GetItemShortDescriptionOk returns a tuple with the ItemShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemShortDescription

`func (o *QuoteLineDto) SetItemShortDescription(v string)`

SetItemShortDescription sets ItemShortDescription field to given value.

### HasItemShortDescription

`func (o *QuoteLineDto) HasItemShortDescription() bool`

HasItemShortDescription returns a boolean if a field has been set.

### SetItemShortDescriptionNil

`func (o *QuoteLineDto) SetItemShortDescriptionNil(b bool)`

 SetItemShortDescriptionNil sets the value for ItemShortDescription to be an explicit nil

### UnsetItemShortDescription
`func (o *QuoteLineDto) UnsetItemShortDescription()`

UnsetItemShortDescription ensures that no value is present for ItemShortDescription, not even an explicit nil
### GetItemPrimaryImageUrl

`func (o *QuoteLineDto) GetItemPrimaryImageUrl() string`

GetItemPrimaryImageUrl returns the ItemPrimaryImageUrl field if non-nil, zero value otherwise.

### GetItemPrimaryImageUrlOk

`func (o *QuoteLineDto) GetItemPrimaryImageUrlOk() (*string, bool)`

GetItemPrimaryImageUrlOk returns a tuple with the ItemPrimaryImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrimaryImageUrl

`func (o *QuoteLineDto) SetItemPrimaryImageUrl(v string)`

SetItemPrimaryImageUrl sets ItemPrimaryImageUrl field to given value.

### HasItemPrimaryImageUrl

`func (o *QuoteLineDto) HasItemPrimaryImageUrl() bool`

HasItemPrimaryImageUrl returns a boolean if a field has been set.

### SetItemPrimaryImageUrlNil

`func (o *QuoteLineDto) SetItemPrimaryImageUrlNil(b bool)`

 SetItemPrimaryImageUrlNil sets the value for ItemPrimaryImageUrl to be an explicit nil

### UnsetItemPrimaryImageUrl
`func (o *QuoteLineDto) UnsetItemPrimaryImageUrl()`

UnsetItemPrimaryImageUrl ensures that no value is present for ItemPrimaryImageUrl, not even an explicit nil
### GetShippingPolicyId

`func (o *QuoteLineDto) GetShippingPolicyId() string`

GetShippingPolicyId returns the ShippingPolicyId field if non-nil, zero value otherwise.

### GetShippingPolicyIdOk

`func (o *QuoteLineDto) GetShippingPolicyIdOk() (*string, bool)`

GetShippingPolicyIdOk returns a tuple with the ShippingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPolicyId

`func (o *QuoteLineDto) SetShippingPolicyId(v string)`

SetShippingPolicyId sets ShippingPolicyId field to given value.

### HasShippingPolicyId

`func (o *QuoteLineDto) HasShippingPolicyId() bool`

HasShippingPolicyId returns a boolean if a field has been set.

### SetShippingPolicyIdNil

`func (o *QuoteLineDto) SetShippingPolicyIdNil(b bool)`

 SetShippingPolicyIdNil sets the value for ShippingPolicyId to be an explicit nil

### UnsetShippingPolicyId
`func (o *QuoteLineDto) UnsetShippingPolicyId()`

UnsetShippingPolicyId ensures that no value is present for ShippingPolicyId, not even an explicit nil
### GetTenantId

`func (o *QuoteLineDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteLineDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteLineDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *QuoteLineDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *QuoteLineDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *QuoteLineDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *QuoteLineDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *QuoteLineDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *QuoteLineDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *QuoteLineDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *QuoteLineDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *QuoteLineDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCurrencyId

`func (o *QuoteLineDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *QuoteLineDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *QuoteLineDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *QuoteLineDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *QuoteLineDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *QuoteLineDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDescription

`func (o *QuoteLineDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuoteLineDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuoteLineDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuoteLineDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *QuoteLineDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QuoteLineDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *QuoteLineDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteLineDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteLineDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteLineDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetFree

`func (o *QuoteLineDto) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *QuoteLineDto) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *QuoteLineDto) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *QuoteLineDto) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetFreeReason

`func (o *QuoteLineDto) GetFreeReason() string`

GetFreeReason returns the FreeReason field if non-nil, zero value otherwise.

### GetFreeReasonOk

`func (o *QuoteLineDto) GetFreeReasonOk() (*string, bool)`

GetFreeReasonOk returns a tuple with the FreeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeReason

`func (o *QuoteLineDto) SetFreeReason(v string)`

SetFreeReason sets FreeReason field to given value.

### HasFreeReason

`func (o *QuoteLineDto) HasFreeReason() bool`

HasFreeReason returns a boolean if a field has been set.

### SetFreeReasonNil

`func (o *QuoteLineDto) SetFreeReasonNil(b bool)`

 SetFreeReasonNil sets the value for FreeReason to be an explicit nil

### UnsetFreeReason
`func (o *QuoteLineDto) UnsetFreeReason()`

UnsetFreeReason ensures that no value is present for FreeReason, not even an explicit nil
### GetFreeReasonCode

`func (o *QuoteLineDto) GetFreeReasonCode() string`

GetFreeReasonCode returns the FreeReasonCode field if non-nil, zero value otherwise.

### GetFreeReasonCodeOk

`func (o *QuoteLineDto) GetFreeReasonCodeOk() (*string, bool)`

GetFreeReasonCodeOk returns a tuple with the FreeReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeReasonCode

`func (o *QuoteLineDto) SetFreeReasonCode(v string)`

SetFreeReasonCode sets FreeReasonCode field to given value.

### HasFreeReasonCode

`func (o *QuoteLineDto) HasFreeReasonCode() bool`

HasFreeReasonCode returns a boolean if a field has been set.

### SetFreeReasonCodeNil

`func (o *QuoteLineDto) SetFreeReasonCodeNil(b bool)`

 SetFreeReasonCodeNil sets the value for FreeReasonCode to be an explicit nil

### UnsetFreeReasonCode
`func (o *QuoteLineDto) UnsetFreeReasonCode()`

UnsetFreeReasonCode ensures that no value is present for FreeReasonCode, not even an explicit nil
### GetData

`func (o *QuoteLineDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QuoteLineDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QuoteLineDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *QuoteLineDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *QuoteLineDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *QuoteLineDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *QuoteLineDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *QuoteLineDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *QuoteLineDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *QuoteLineDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *QuoteLineDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *QuoteLineDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *QuoteLineDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *QuoteLineDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *QuoteLineDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *QuoteLineDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *QuoteLineDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *QuoteLineDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *QuoteLineDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *QuoteLineDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *QuoteLineDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *QuoteLineDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *QuoteLineDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *QuoteLineDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *QuoteLineDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *QuoteLineDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *QuoteLineDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *QuoteLineDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *QuoteLineDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *QuoteLineDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *QuoteLineDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *QuoteLineDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *QuoteLineDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *QuoteLineDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *QuoteLineDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *QuoteLineDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *QuoteLineDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *QuoteLineDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *QuoteLineDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *QuoteLineDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *QuoteLineDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *QuoteLineDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *QuoteLineDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *QuoteLineDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *QuoteLineDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *QuoteLineDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *QuoteLineDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *QuoteLineDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *QuoteLineDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *QuoteLineDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *QuoteLineDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *QuoteLineDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *QuoteLineDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *QuoteLineDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *QuoteLineDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *QuoteLineDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *QuoteLineDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *QuoteLineDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *QuoteLineDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *QuoteLineDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *QuoteLineDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *QuoteLineDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *QuoteLineDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *QuoteLineDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *QuoteLineDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *QuoteLineDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *QuoteLineDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *QuoteLineDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *QuoteLineDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *QuoteLineDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *QuoteLineDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *QuoteLineDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *QuoteLineDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *QuoteLineDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *QuoteLineDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *QuoteLineDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *QuoteLineDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *QuoteLineDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *QuoteLineDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *QuoteLineDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *QuoteLineDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *QuoteLineDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *QuoteLineDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *QuoteLineDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *QuoteLineDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *QuoteLineDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *QuoteLineDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *QuoteLineDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *QuoteLineDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *QuoteLineDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *QuoteLineDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *QuoteLineDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *QuoteLineDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *QuoteLineDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *QuoteLineDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *QuoteLineDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *QuoteLineDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *QuoteLineDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *QuoteLineDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *QuoteLineDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *QuoteLineDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *QuoteLineDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *QuoteLineDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *QuoteLineDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *QuoteLineDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *QuoteLineDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *QuoteLineDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *QuoteLineDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *QuoteLineDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *QuoteLineDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *QuoteLineDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *QuoteLineDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *QuoteLineDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *QuoteLineDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *QuoteLineDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *QuoteLineDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *QuoteLineDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *QuoteLineDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *QuoteLineDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *QuoteLineDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetItemPriceId

`func (o *QuoteLineDto) GetItemPriceId() string`

GetItemPriceId returns the ItemPriceId field if non-nil, zero value otherwise.

### GetItemPriceIdOk

`func (o *QuoteLineDto) GetItemPriceIdOk() (*string, bool)`

GetItemPriceIdOk returns a tuple with the ItemPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceId

`func (o *QuoteLineDto) SetItemPriceId(v string)`

SetItemPriceId sets ItemPriceId field to given value.

### HasItemPriceId

`func (o *QuoteLineDto) HasItemPriceId() bool`

HasItemPriceId returns a boolean if a field has been set.

### SetItemPriceIdNil

`func (o *QuoteLineDto) SetItemPriceIdNil(b bool)`

 SetItemPriceIdNil sets the value for ItemPriceId to be an explicit nil

### UnsetItemPriceId
`func (o *QuoteLineDto) UnsetItemPriceId()`

UnsetItemPriceId ensures that no value is present for ItemPriceId, not even an explicit nil
### GetPriceListItemId

`func (o *QuoteLineDto) GetPriceListItemId() string`

GetPriceListItemId returns the PriceListItemId field if non-nil, zero value otherwise.

### GetPriceListItemIdOk

`func (o *QuoteLineDto) GetPriceListItemIdOk() (*string, bool)`

GetPriceListItemIdOk returns a tuple with the PriceListItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListItemId

`func (o *QuoteLineDto) SetPriceListItemId(v string)`

SetPriceListItemId sets PriceListItemId field to given value.

### HasPriceListItemId

`func (o *QuoteLineDto) HasPriceListItemId() bool`

HasPriceListItemId returns a boolean if a field has been set.

### SetPriceListItemIdNil

`func (o *QuoteLineDto) SetPriceListItemIdNil(b bool)`

 SetPriceListItemIdNil sets the value for PriceListItemId to be an explicit nil

### UnsetPriceListItemId
`func (o *QuoteLineDto) UnsetPriceListItemId()`

UnsetPriceListItemId ensures that no value is present for PriceListItemId, not even an explicit nil
### GetUnitId

`func (o *QuoteLineDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *QuoteLineDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *QuoteLineDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *QuoteLineDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *QuoteLineDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *QuoteLineDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *QuoteLineDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *QuoteLineDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *QuoteLineDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *QuoteLineDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *QuoteLineDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *QuoteLineDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetTaxCalculationMethod

`func (o *QuoteLineDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *QuoteLineDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *QuoteLineDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *QuoteLineDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *QuoteLineDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *QuoteLineDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *QuoteLineDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *QuoteLineDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetForexRates

`func (o *QuoteLineDto) GetForexRates() ForexRates`

GetForexRates returns the ForexRates field if non-nil, zero value otherwise.

### GetForexRatesOk

`func (o *QuoteLineDto) GetForexRatesOk() (*ForexRates, bool)`

GetForexRatesOk returns a tuple with the ForexRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRates

`func (o *QuoteLineDto) SetForexRates(v ForexRates)`

SetForexRates sets ForexRates field to given value.

### HasForexRates

`func (o *QuoteLineDto) HasForexRates() bool`

HasForexRates returns a boolean if a field has been set.

### GetForexRate

`func (o *QuoteLineDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *QuoteLineDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *QuoteLineDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *QuoteLineDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetTotalDetailInUsd

`func (o *QuoteLineDto) GetTotalDetailInUsd() float64`

GetTotalDetailInUsd returns the TotalDetailInUsd field if non-nil, zero value otherwise.

### GetTotalDetailInUsdOk

`func (o *QuoteLineDto) GetTotalDetailInUsdOk() (*float64, bool)`

GetTotalDetailInUsdOk returns a tuple with the TotalDetailInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailInUsd

`func (o *QuoteLineDto) SetTotalDetailInUsd(v float64)`

SetTotalDetailInUsd sets TotalDetailInUsd field to given value.

### HasTotalDetailInUsd

`func (o *QuoteLineDto) HasTotalDetailInUsd() bool`

HasTotalDetailInUsd returns a boolean if a field has been set.

### GetTotalProfitInUsd

`func (o *QuoteLineDto) GetTotalProfitInUsd() float64`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *QuoteLineDto) GetTotalProfitInUsdOk() (*float64, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *QuoteLineDto) SetTotalProfitInUsd(v float64)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.

### HasTotalProfitInUsd

`func (o *QuoteLineDto) HasTotalProfitInUsd() bool`

HasTotalProfitInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *QuoteLineDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *QuoteLineDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *QuoteLineDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *QuoteLineDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *QuoteLineDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *QuoteLineDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *QuoteLineDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *QuoteLineDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseInUsd

`func (o *QuoteLineDto) GetTotalTaxBaseInUsd() float64`

GetTotalTaxBaseInUsd returns the TotalTaxBaseInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseInUsdOk

`func (o *QuoteLineDto) GetTotalTaxBaseInUsdOk() (*float64, bool)`

GetTotalTaxBaseInUsdOk returns a tuple with the TotalTaxBaseInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseInUsd

`func (o *QuoteLineDto) SetTotalTaxBaseInUsd(v float64)`

SetTotalTaxBaseInUsd sets TotalTaxBaseInUsd field to given value.

### HasTotalTaxBaseInUsd

`func (o *QuoteLineDto) HasTotalTaxBaseInUsd() bool`

HasTotalTaxBaseInUsd returns a boolean if a field has been set.

### GetTotalTaxesInUsd

`func (o *QuoteLineDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *QuoteLineDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *QuoteLineDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *QuoteLineDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalWithheldTaxesInUsd

`func (o *QuoteLineDto) GetTotalWithheldTaxesInUsd() float64`

GetTotalWithheldTaxesInUsd returns the TotalWithheldTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithheldTaxesInUsdOk

`func (o *QuoteLineDto) GetTotalWithheldTaxesInUsdOk() (*float64, bool)`

GetTotalWithheldTaxesInUsdOk returns a tuple with the TotalWithheldTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxesInUsd

`func (o *QuoteLineDto) SetTotalWithheldTaxesInUsd(v float64)`

SetTotalWithheldTaxesInUsd sets TotalWithheldTaxesInUsd field to given value.

### HasTotalWithheldTaxesInUsd

`func (o *QuoteLineDto) HasTotalWithheldTaxesInUsd() bool`

HasTotalWithheldTaxesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *QuoteLineDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *QuoteLineDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *QuoteLineDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *QuoteLineDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxesInUsd

`func (o *QuoteLineDto) GetTotalShippingTaxesInUsd() float64`

GetTotalShippingTaxesInUsd returns the TotalShippingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxesInUsdOk

`func (o *QuoteLineDto) GetTotalShippingTaxesInUsdOk() (*float64, bool)`

GetTotalShippingTaxesInUsdOk returns a tuple with the TotalShippingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxesInUsd

`func (o *QuoteLineDto) SetTotalShippingTaxesInUsd(v float64)`

SetTotalShippingTaxesInUsd sets TotalShippingTaxesInUsd field to given value.

### HasTotalShippingTaxesInUsd

`func (o *QuoteLineDto) HasTotalShippingTaxesInUsd() bool`

HasTotalShippingTaxesInUsd returns a boolean if a field has been set.

### GetTotalWarrantyCostInUsd

`func (o *QuoteLineDto) GetTotalWarrantyCostInUsd() float64`

GetTotalWarrantyCostInUsd returns the TotalWarrantyCostInUsd field if non-nil, zero value otherwise.

### GetTotalWarrantyCostInUsdOk

`func (o *QuoteLineDto) GetTotalWarrantyCostInUsdOk() (*float64, bool)`

GetTotalWarrantyCostInUsdOk returns a tuple with the TotalWarrantyCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWarrantyCostInUsd

`func (o *QuoteLineDto) SetTotalWarrantyCostInUsd(v float64)`

SetTotalWarrantyCostInUsd sets TotalWarrantyCostInUsd field to given value.

### HasTotalWarrantyCostInUsd

`func (o *QuoteLineDto) HasTotalWarrantyCostInUsd() bool`

HasTotalWarrantyCostInUsd returns a boolean if a field has been set.

### GetTotalReturnCostInUsd

`func (o *QuoteLineDto) GetTotalReturnCostInUsd() float64`

GetTotalReturnCostInUsd returns the TotalReturnCostInUsd field if non-nil, zero value otherwise.

### GetTotalReturnCostInUsdOk

`func (o *QuoteLineDto) GetTotalReturnCostInUsdOk() (*float64, bool)`

GetTotalReturnCostInUsdOk returns a tuple with the TotalReturnCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReturnCostInUsd

`func (o *QuoteLineDto) SetTotalReturnCostInUsd(v float64)`

SetTotalReturnCostInUsd sets TotalReturnCostInUsd field to given value.

### HasTotalReturnCostInUsd

`func (o *QuoteLineDto) HasTotalReturnCostInUsd() bool`

HasTotalReturnCostInUsd returns a boolean if a field has been set.

### GetTotalRefundCostInUsd

`func (o *QuoteLineDto) GetTotalRefundCostInUsd() float64`

GetTotalRefundCostInUsd returns the TotalRefundCostInUsd field if non-nil, zero value otherwise.

### GetTotalRefundCostInUsdOk

`func (o *QuoteLineDto) GetTotalRefundCostInUsdOk() (*float64, bool)`

GetTotalRefundCostInUsdOk returns a tuple with the TotalRefundCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRefundCostInUsd

`func (o *QuoteLineDto) SetTotalRefundCostInUsd(v float64)`

SetTotalRefundCostInUsd sets TotalRefundCostInUsd field to given value.

### HasTotalRefundCostInUsd

`func (o *QuoteLineDto) HasTotalRefundCostInUsd() bool`

HasTotalRefundCostInUsd returns a boolean if a field has been set.

### GetTotalInUsd

`func (o *QuoteLineDto) GetTotalInUsd() float64`

GetTotalInUsd returns the TotalInUsd field if non-nil, zero value otherwise.

### GetTotalInUsdOk

`func (o *QuoteLineDto) GetTotalInUsdOk() (*float64, bool)`

GetTotalInUsdOk returns a tuple with the TotalInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInUsd

`func (o *QuoteLineDto) SetTotalInUsd(v float64)`

SetTotalInUsd sets TotalInUsd field to given value.

### HasTotalInUsd

`func (o *QuoteLineDto) HasTotalInUsd() bool`

HasTotalInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsInUsd

`func (o *QuoteLineDto) GetTotalGlobalDiscountsInUsd() float64`

GetTotalGlobalDiscountsInUsd returns the TotalGlobalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsInUsdOk

`func (o *QuoteLineDto) GetTotalGlobalDiscountsInUsdOk() (*float64, bool)`

GetTotalGlobalDiscountsInUsdOk returns a tuple with the TotalGlobalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsInUsd

`func (o *QuoteLineDto) SetTotalGlobalDiscountsInUsd(v float64)`

SetTotalGlobalDiscountsInUsd sets TotalGlobalDiscountsInUsd field to given value.

### HasTotalGlobalDiscountsInUsd

`func (o *QuoteLineDto) HasTotalGlobalDiscountsInUsd() bool`

HasTotalGlobalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesInUsd

`func (o *QuoteLineDto) GetTotalGlobalSurchargesInUsd() float64`

GetTotalGlobalSurchargesInUsd returns the TotalGlobalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesInUsdOk

`func (o *QuoteLineDto) GetTotalGlobalSurchargesInUsdOk() (*float64, bool)`

GetTotalGlobalSurchargesInUsdOk returns a tuple with the TotalGlobalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesInUsd

`func (o *QuoteLineDto) SetTotalGlobalSurchargesInUsd(v float64)`

SetTotalGlobalSurchargesInUsd sets TotalGlobalSurchargesInUsd field to given value.

### HasTotalGlobalSurchargesInUsd

`func (o *QuoteLineDto) HasTotalGlobalSurchargesInUsd() bool`

HasTotalGlobalSurchargesInUsd returns a boolean if a field has been set.

### GetCustomGlobalSurchargesAmount

`func (o *QuoteLineDto) GetCustomGlobalSurchargesAmount() float64`

GetCustomGlobalSurchargesAmount returns the CustomGlobalSurchargesAmount field if non-nil, zero value otherwise.

### GetCustomGlobalSurchargesAmountOk

`func (o *QuoteLineDto) GetCustomGlobalSurchargesAmountOk() (*float64, bool)`

GetCustomGlobalSurchargesAmountOk returns a tuple with the CustomGlobalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalSurchargesAmount

`func (o *QuoteLineDto) SetCustomGlobalSurchargesAmount(v float64)`

SetCustomGlobalSurchargesAmount sets CustomGlobalSurchargesAmount field to given value.

### HasCustomGlobalSurchargesAmount

`func (o *QuoteLineDto) HasCustomGlobalSurchargesAmount() bool`

HasCustomGlobalSurchargesAmount returns a boolean if a field has been set.

### GetCustomGlobalDiscountsAmount

`func (o *QuoteLineDto) GetCustomGlobalDiscountsAmount() float64`

GetCustomGlobalDiscountsAmount returns the CustomGlobalDiscountsAmount field if non-nil, zero value otherwise.

### GetCustomGlobalDiscountsAmountOk

`func (o *QuoteLineDto) GetCustomGlobalDiscountsAmountOk() (*float64, bool)`

GetCustomGlobalDiscountsAmountOk returns a tuple with the CustomGlobalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalDiscountsAmount

`func (o *QuoteLineDto) SetCustomGlobalDiscountsAmount(v float64)`

SetCustomGlobalDiscountsAmount sets CustomGlobalDiscountsAmount field to given value.

### HasCustomGlobalDiscountsAmount

`func (o *QuoteLineDto) HasCustomGlobalDiscountsAmount() bool`

HasCustomGlobalDiscountsAmount returns a boolean if a field has been set.

### GetReturnPolicyId

`func (o *QuoteLineDto) GetReturnPolicyId() string`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *QuoteLineDto) GetReturnPolicyIdOk() (*string, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *QuoteLineDto) SetReturnPolicyId(v string)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *QuoteLineDto) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### SetReturnPolicyIdNil

`func (o *QuoteLineDto) SetReturnPolicyIdNil(b bool)`

 SetReturnPolicyIdNil sets the value for ReturnPolicyId to be an explicit nil

### UnsetReturnPolicyId
`func (o *QuoteLineDto) UnsetReturnPolicyId()`

UnsetReturnPolicyId ensures that no value is present for ReturnPolicyId, not even an explicit nil
### GetRefundPolicyId

`func (o *QuoteLineDto) GetRefundPolicyId() string`

GetRefundPolicyId returns the RefundPolicyId field if non-nil, zero value otherwise.

### GetRefundPolicyIdOk

`func (o *QuoteLineDto) GetRefundPolicyIdOk() (*string, bool)`

GetRefundPolicyIdOk returns a tuple with the RefundPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPolicyId

`func (o *QuoteLineDto) SetRefundPolicyId(v string)`

SetRefundPolicyId sets RefundPolicyId field to given value.

### HasRefundPolicyId

`func (o *QuoteLineDto) HasRefundPolicyId() bool`

HasRefundPolicyId returns a boolean if a field has been set.

### SetRefundPolicyIdNil

`func (o *QuoteLineDto) SetRefundPolicyIdNil(b bool)`

 SetRefundPolicyIdNil sets the value for RefundPolicyId to be an explicit nil

### UnsetRefundPolicyId
`func (o *QuoteLineDto) UnsetRefundPolicyId()`

UnsetRefundPolicyId ensures that no value is present for RefundPolicyId, not even an explicit nil
### GetWarrantyPolicyId

`func (o *QuoteLineDto) GetWarrantyPolicyId() string`

GetWarrantyPolicyId returns the WarrantyPolicyId field if non-nil, zero value otherwise.

### GetWarrantyPolicyIdOk

`func (o *QuoteLineDto) GetWarrantyPolicyIdOk() (*string, bool)`

GetWarrantyPolicyIdOk returns a tuple with the WarrantyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyPolicyId

`func (o *QuoteLineDto) SetWarrantyPolicyId(v string)`

SetWarrantyPolicyId sets WarrantyPolicyId field to given value.

### HasWarrantyPolicyId

`func (o *QuoteLineDto) HasWarrantyPolicyId() bool`

HasWarrantyPolicyId returns a boolean if a field has been set.

### SetWarrantyPolicyIdNil

`func (o *QuoteLineDto) SetWarrantyPolicyIdNil(b bool)`

 SetWarrantyPolicyIdNil sets the value for WarrantyPolicyId to be an explicit nil

### UnsetWarrantyPolicyId
`func (o *QuoteLineDto) UnsetWarrantyPolicyId()`

UnsetWarrantyPolicyId ensures that no value is present for WarrantyPolicyId, not even an explicit nil
### GetShipmentPolicyId

`func (o *QuoteLineDto) GetShipmentPolicyId() string`

GetShipmentPolicyId returns the ShipmentPolicyId field if non-nil, zero value otherwise.

### GetShipmentPolicyIdOk

`func (o *QuoteLineDto) GetShipmentPolicyIdOk() (*string, bool)`

GetShipmentPolicyIdOk returns a tuple with the ShipmentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentPolicyId

`func (o *QuoteLineDto) SetShipmentPolicyId(v string)`

SetShipmentPolicyId sets ShipmentPolicyId field to given value.

### HasShipmentPolicyId

`func (o *QuoteLineDto) HasShipmentPolicyId() bool`

HasShipmentPolicyId returns a boolean if a field has been set.

### SetShipmentPolicyIdNil

`func (o *QuoteLineDto) SetShipmentPolicyIdNil(b bool)`

 SetShipmentPolicyIdNil sets the value for ShipmentPolicyId to be an explicit nil

### UnsetShipmentPolicyId
`func (o *QuoteLineDto) UnsetShipmentPolicyId()`

UnsetShipmentPolicyId ensures that no value is present for ShipmentPolicyId, not even an explicit nil
### GetShippingLocationId

`func (o *QuoteLineDto) GetShippingLocationId() string`

GetShippingLocationId returns the ShippingLocationId field if non-nil, zero value otherwise.

### GetShippingLocationIdOk

`func (o *QuoteLineDto) GetShippingLocationIdOk() (*string, bool)`

GetShippingLocationIdOk returns a tuple with the ShippingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocationId

`func (o *QuoteLineDto) SetShippingLocationId(v string)`

SetShippingLocationId sets ShippingLocationId field to given value.

### HasShippingLocationId

`func (o *QuoteLineDto) HasShippingLocationId() bool`

HasShippingLocationId returns a boolean if a field has been set.

### SetShippingLocationIdNil

`func (o *QuoteLineDto) SetShippingLocationIdNil(b bool)`

 SetShippingLocationIdNil sets the value for ShippingLocationId to be an explicit nil

### UnsetShippingLocationId
`func (o *QuoteLineDto) UnsetShippingLocationId()`

UnsetShippingLocationId ensures that no value is present for ShippingLocationId, not even an explicit nil
### GetLocationId

`func (o *QuoteLineDto) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *QuoteLineDto) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *QuoteLineDto) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *QuoteLineDto) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *QuoteLineDto) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *QuoteLineDto) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetQuoteItemRecordId

`func (o *QuoteLineDto) GetQuoteItemRecordId() string`

GetQuoteItemRecordId returns the QuoteItemRecordId field if non-nil, zero value otherwise.

### GetQuoteItemRecordIdOk

`func (o *QuoteLineDto) GetQuoteItemRecordIdOk() (*string, bool)`

GetQuoteItemRecordIdOk returns a tuple with the QuoteItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteItemRecordId

`func (o *QuoteLineDto) SetQuoteItemRecordId(v string)`

SetQuoteItemRecordId sets QuoteItemRecordId field to given value.

### HasQuoteItemRecordId

`func (o *QuoteLineDto) HasQuoteItemRecordId() bool`

HasQuoteItemRecordId returns a boolean if a field has been set.

### SetQuoteItemRecordIdNil

`func (o *QuoteLineDto) SetQuoteItemRecordIdNil(b bool)`

 SetQuoteItemRecordIdNil sets the value for QuoteItemRecordId to be an explicit nil

### UnsetQuoteItemRecordId
`func (o *QuoteLineDto) UnsetQuoteItemRecordId()`

UnsetQuoteItemRecordId ensures that no value is present for QuoteItemRecordId, not even an explicit nil
### GetBusinessProfileRecordId

`func (o *QuoteLineDto) GetBusinessProfileRecordId() string`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *QuoteLineDto) GetBusinessProfileRecordIdOk() (*string, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *QuoteLineDto) SetBusinessProfileRecordId(v string)`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *QuoteLineDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### SetBusinessProfileRecordIdNil

`func (o *QuoteLineDto) SetBusinessProfileRecordIdNil(b bool)`

 SetBusinessProfileRecordIdNil sets the value for BusinessProfileRecordId to be an explicit nil

### UnsetBusinessProfileRecordId
`func (o *QuoteLineDto) UnsetBusinessProfileRecordId()`

UnsetBusinessProfileRecordId ensures that no value is present for BusinessProfileRecordId, not even an explicit nil
### GetParentBillingItemRecordId

`func (o *QuoteLineDto) GetParentBillingItemRecordId() string`

GetParentBillingItemRecordId returns the ParentBillingItemRecordId field if non-nil, zero value otherwise.

### GetParentBillingItemRecordIdOk

`func (o *QuoteLineDto) GetParentBillingItemRecordIdOk() (*string, bool)`

GetParentBillingItemRecordIdOk returns a tuple with the ParentBillingItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBillingItemRecordId

`func (o *QuoteLineDto) SetParentBillingItemRecordId(v string)`

SetParentBillingItemRecordId sets ParentBillingItemRecordId field to given value.

### HasParentBillingItemRecordId

`func (o *QuoteLineDto) HasParentBillingItemRecordId() bool`

HasParentBillingItemRecordId returns a boolean if a field has been set.

### SetParentBillingItemRecordIdNil

`func (o *QuoteLineDto) SetParentBillingItemRecordIdNil(b bool)`

 SetParentBillingItemRecordIdNil sets the value for ParentBillingItemRecordId to be an explicit nil

### UnsetParentBillingItemRecordId
`func (o *QuoteLineDto) UnsetParentBillingItemRecordId()`

UnsetParentBillingItemRecordId ensures that no value is present for ParentBillingItemRecordId, not even an explicit nil
### GetCurrency

`func (o *QuoteLineDto) GetCurrency() CurrencyId`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QuoteLineDto) GetCurrencyOk() (*CurrencyId, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QuoteLineDto) SetCurrency(v CurrencyId)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *QuoteLineDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotalDetail

`func (o *QuoteLineDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *QuoteLineDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *QuoteLineDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *QuoteLineDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *QuoteLineDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *QuoteLineDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *QuoteLineDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *QuoteLineDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *QuoteLineDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *QuoteLineDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalDetailAmount

`func (o *QuoteLineDto) GetTotalDetailAmount() Money`

GetTotalDetailAmount returns the TotalDetailAmount field if non-nil, zero value otherwise.

### GetTotalDetailAmountOk

`func (o *QuoteLineDto) GetTotalDetailAmountOk() (*Money, bool)`

GetTotalDetailAmountOk returns a tuple with the TotalDetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailAmount

`func (o *QuoteLineDto) SetTotalDetailAmount(v Money)`

SetTotalDetailAmount sets TotalDetailAmount field to given value.

### HasTotalDetailAmount

`func (o *QuoteLineDto) HasTotalDetailAmount() bool`

HasTotalDetailAmount returns a boolean if a field has been set.

### GetTotalProfit

`func (o *QuoteLineDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *QuoteLineDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *QuoteLineDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *QuoteLineDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *QuoteLineDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *QuoteLineDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *QuoteLineDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *QuoteLineDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *QuoteLineDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *QuoteLineDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalProfitAmount

`func (o *QuoteLineDto) GetTotalProfitAmount() Money`

GetTotalProfitAmount returns the TotalProfitAmount field if non-nil, zero value otherwise.

### GetTotalProfitAmountOk

`func (o *QuoteLineDto) GetTotalProfitAmountOk() (*Money, bool)`

GetTotalProfitAmountOk returns a tuple with the TotalProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitAmount

`func (o *QuoteLineDto) SetTotalProfitAmount(v Money)`

SetTotalProfitAmount sets TotalProfitAmount field to given value.

### HasTotalProfitAmount

`func (o *QuoteLineDto) HasTotalProfitAmount() bool`

HasTotalProfitAmount returns a boolean if a field has been set.

### GetTotalDiscounts

`func (o *QuoteLineDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *QuoteLineDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *QuoteLineDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *QuoteLineDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *QuoteLineDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *QuoteLineDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *QuoteLineDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *QuoteLineDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *QuoteLineDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *QuoteLineDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalDiscountsAmount

`func (o *QuoteLineDto) GetTotalDiscountsAmount() Money`

GetTotalDiscountsAmount returns the TotalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalDiscountsAmountOk

`func (o *QuoteLineDto) GetTotalDiscountsAmountOk() (*Money, bool)`

GetTotalDiscountsAmountOk returns a tuple with the TotalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsAmount

`func (o *QuoteLineDto) SetTotalDiscountsAmount(v Money)`

SetTotalDiscountsAmount sets TotalDiscountsAmount field to given value.

### HasTotalDiscountsAmount

`func (o *QuoteLineDto) HasTotalDiscountsAmount() bool`

HasTotalDiscountsAmount returns a boolean if a field has been set.

### GetTotalSurcharges

`func (o *QuoteLineDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *QuoteLineDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *QuoteLineDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *QuoteLineDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *QuoteLineDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *QuoteLineDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *QuoteLineDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *QuoteLineDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *QuoteLineDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *QuoteLineDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalSurchargesAmount

`func (o *QuoteLineDto) GetTotalSurchargesAmount() Money`

GetTotalSurchargesAmount returns the TotalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalSurchargesAmountOk

`func (o *QuoteLineDto) GetTotalSurchargesAmountOk() (*Money, bool)`

GetTotalSurchargesAmountOk returns a tuple with the TotalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesAmount

`func (o *QuoteLineDto) SetTotalSurchargesAmount(v Money)`

SetTotalSurchargesAmount sets TotalSurchargesAmount field to given value.

### HasTotalSurchargesAmount

`func (o *QuoteLineDto) HasTotalSurchargesAmount() bool`

HasTotalSurchargesAmount returns a boolean if a field has been set.

### GetTotalTaxBase

`func (o *QuoteLineDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *QuoteLineDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *QuoteLineDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *QuoteLineDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *QuoteLineDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *QuoteLineDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *QuoteLineDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *QuoteLineDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *QuoteLineDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *QuoteLineDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalTaxBaseAmount

`func (o *QuoteLineDto) GetTotalTaxBaseAmount() Money`

GetTotalTaxBaseAmount returns the TotalTaxBaseAmount field if non-nil, zero value otherwise.

### GetTotalTaxBaseAmountOk

`func (o *QuoteLineDto) GetTotalTaxBaseAmountOk() (*Money, bool)`

GetTotalTaxBaseAmountOk returns a tuple with the TotalTaxBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseAmount

`func (o *QuoteLineDto) SetTotalTaxBaseAmount(v Money)`

SetTotalTaxBaseAmount sets TotalTaxBaseAmount field to given value.

### HasTotalTaxBaseAmount

`func (o *QuoteLineDto) HasTotalTaxBaseAmount() bool`

HasTotalTaxBaseAmount returns a boolean if a field has been set.

### GetTotalTaxes

`func (o *QuoteLineDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *QuoteLineDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *QuoteLineDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *QuoteLineDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *QuoteLineDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *QuoteLineDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *QuoteLineDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *QuoteLineDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *QuoteLineDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *QuoteLineDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalTaxesAmount

`func (o *QuoteLineDto) GetTotalTaxesAmount() Money`

GetTotalTaxesAmount returns the TotalTaxesAmount field if non-nil, zero value otherwise.

### GetTotalTaxesAmountOk

`func (o *QuoteLineDto) GetTotalTaxesAmountOk() (*Money, bool)`

GetTotalTaxesAmountOk returns a tuple with the TotalTaxesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesAmount

`func (o *QuoteLineDto) SetTotalTaxesAmount(v Money)`

SetTotalTaxesAmount sets TotalTaxesAmount field to given value.

### HasTotalTaxesAmount

`func (o *QuoteLineDto) HasTotalTaxesAmount() bool`

HasTotalTaxesAmount returns a boolean if a field has been set.

### GetTotalShippingCost

`func (o *QuoteLineDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *QuoteLineDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *QuoteLineDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *QuoteLineDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *QuoteLineDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *QuoteLineDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *QuoteLineDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *QuoteLineDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *QuoteLineDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *QuoteLineDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingCostAmount

`func (o *QuoteLineDto) GetTotalShippingCostAmount() Money`

GetTotalShippingCostAmount returns the TotalShippingCostAmount field if non-nil, zero value otherwise.

### GetTotalShippingCostAmountOk

`func (o *QuoteLineDto) GetTotalShippingCostAmountOk() (*Money, bool)`

GetTotalShippingCostAmountOk returns a tuple with the TotalShippingCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostAmount

`func (o *QuoteLineDto) SetTotalShippingCostAmount(v Money)`

SetTotalShippingCostAmount sets TotalShippingCostAmount field to given value.

### HasTotalShippingCostAmount

`func (o *QuoteLineDto) HasTotalShippingCostAmount() bool`

HasTotalShippingCostAmount returns a boolean if a field has been set.

### GetTotalShippingTax

`func (o *QuoteLineDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *QuoteLineDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *QuoteLineDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *QuoteLineDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *QuoteLineDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *QuoteLineDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *QuoteLineDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *QuoteLineDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *QuoteLineDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *QuoteLineDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalShippingTaxAmount

`func (o *QuoteLineDto) GetTotalShippingTaxAmount() Money`

GetTotalShippingTaxAmount returns the TotalShippingTaxAmount field if non-nil, zero value otherwise.

### GetTotalShippingTaxAmountOk

`func (o *QuoteLineDto) GetTotalShippingTaxAmountOk() (*Money, bool)`

GetTotalShippingTaxAmountOk returns a tuple with the TotalShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxAmount

`func (o *QuoteLineDto) SetTotalShippingTaxAmount(v Money)`

SetTotalShippingTaxAmount sets TotalShippingTaxAmount field to given value.

### HasTotalShippingTaxAmount

`func (o *QuoteLineDto) HasTotalShippingTaxAmount() bool`

HasTotalShippingTaxAmount returns a boolean if a field has been set.

### GetTotalWithheldTax

`func (o *QuoteLineDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *QuoteLineDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *QuoteLineDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *QuoteLineDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *QuoteLineDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *QuoteLineDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *QuoteLineDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *QuoteLineDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *QuoteLineDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *QuoteLineDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalWithheldTaxAmount

`func (o *QuoteLineDto) GetTotalWithheldTaxAmount() Money`

GetTotalWithheldTaxAmount returns the TotalWithheldTaxAmount field if non-nil, zero value otherwise.

### GetTotalWithheldTaxAmountOk

`func (o *QuoteLineDto) GetTotalWithheldTaxAmountOk() (*Money, bool)`

GetTotalWithheldTaxAmountOk returns a tuple with the TotalWithheldTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxAmount

`func (o *QuoteLineDto) SetTotalWithheldTaxAmount(v Money)`

SetTotalWithheldTaxAmount sets TotalWithheldTaxAmount field to given value.

### HasTotalWithheldTaxAmount

`func (o *QuoteLineDto) HasTotalWithheldTaxAmount() bool`

HasTotalWithheldTaxAmount returns a boolean if a field has been set.

### GetTotalGlobalDiscounts

`func (o *QuoteLineDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *QuoteLineDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *QuoteLineDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *QuoteLineDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *QuoteLineDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *QuoteLineDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *QuoteLineDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *QuoteLineDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *QuoteLineDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *QuoteLineDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotalGlobalDiscountsAmount

`func (o *QuoteLineDto) GetTotalGlobalDiscountsAmount() Money`

GetTotalGlobalDiscountsAmount returns the TotalGlobalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsAmountOk

`func (o *QuoteLineDto) GetTotalGlobalDiscountsAmountOk() (*Money, bool)`

GetTotalGlobalDiscountsAmountOk returns a tuple with the TotalGlobalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsAmount

`func (o *QuoteLineDto) SetTotalGlobalDiscountsAmount(v Money)`

SetTotalGlobalDiscountsAmount sets TotalGlobalDiscountsAmount field to given value.

### HasTotalGlobalDiscountsAmount

`func (o *QuoteLineDto) HasTotalGlobalDiscountsAmount() bool`

HasTotalGlobalDiscountsAmount returns a boolean if a field has been set.

### GetTotalGlobalSurcharges

`func (o *QuoteLineDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *QuoteLineDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *QuoteLineDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *QuoteLineDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *QuoteLineDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *QuoteLineDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *QuoteLineDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *QuoteLineDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *QuoteLineDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *QuoteLineDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotalGlobalSurchargesAmount

`func (o *QuoteLineDto) GetTotalGlobalSurchargesAmount() Money`

GetTotalGlobalSurchargesAmount returns the TotalGlobalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesAmountOk

`func (o *QuoteLineDto) GetTotalGlobalSurchargesAmountOk() (*Money, bool)`

GetTotalGlobalSurchargesAmountOk returns a tuple with the TotalGlobalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesAmount

`func (o *QuoteLineDto) SetTotalGlobalSurchargesAmount(v Money)`

SetTotalGlobalSurchargesAmount sets TotalGlobalSurchargesAmount field to given value.

### HasTotalGlobalSurchargesAmount

`func (o *QuoteLineDto) HasTotalGlobalSurchargesAmount() bool`

HasTotalGlobalSurchargesAmount returns a boolean if a field has been set.

### GetTotal

`func (o *QuoteLineDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QuoteLineDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QuoteLineDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QuoteLineDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *QuoteLineDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *QuoteLineDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *QuoteLineDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *QuoteLineDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *QuoteLineDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *QuoteLineDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetTotalAmount

`func (o *QuoteLineDto) GetTotalAmount() Money`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QuoteLineDto) GetTotalAmountOk() (*Money, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QuoteLineDto) SetTotalAmount(v Money)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *QuoteLineDto) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetQuoteId

`func (o *QuoteLineDto) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QuoteLineDto) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QuoteLineDto) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *QuoteLineDto) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### SetQuoteIdNil

`func (o *QuoteLineDto) SetQuoteIdNil(b bool)`

 SetQuoteIdNil sets the value for QuoteId to be an explicit nil

### UnsetQuoteId
`func (o *QuoteLineDto) UnsetQuoteId()`

UnsetQuoteId ensures that no value is present for QuoteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


