# OrderLineDto

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
**OrderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrderLineDto

`func NewOrderLineDto() *OrderLineDto`

NewOrderLineDto instantiates a new OrderLineDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLineDtoWithDefaults

`func NewOrderLineDtoWithDefaults() *OrderLineDto`

NewOrderLineDtoWithDefaults instantiates a new OrderLineDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderLineDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderLineDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderLineDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderLineDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OrderLineDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OrderLineDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *OrderLineDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrderLineDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrderLineDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OrderLineDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *OrderLineDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *OrderLineDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetClosed

`func (o *OrderLineDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *OrderLineDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *OrderLineDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *OrderLineDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetItemId

`func (o *OrderLineDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *OrderLineDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *OrderLineDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *OrderLineDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *OrderLineDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *OrderLineDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemTitle

`func (o *OrderLineDto) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *OrderLineDto) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *OrderLineDto) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *OrderLineDto) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *OrderLineDto) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *OrderLineDto) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetItemShortDescription

`func (o *OrderLineDto) GetItemShortDescription() string`

GetItemShortDescription returns the ItemShortDescription field if non-nil, zero value otherwise.

### GetItemShortDescriptionOk

`func (o *OrderLineDto) GetItemShortDescriptionOk() (*string, bool)`

GetItemShortDescriptionOk returns a tuple with the ItemShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemShortDescription

`func (o *OrderLineDto) SetItemShortDescription(v string)`

SetItemShortDescription sets ItemShortDescription field to given value.

### HasItemShortDescription

`func (o *OrderLineDto) HasItemShortDescription() bool`

HasItemShortDescription returns a boolean if a field has been set.

### SetItemShortDescriptionNil

`func (o *OrderLineDto) SetItemShortDescriptionNil(b bool)`

 SetItemShortDescriptionNil sets the value for ItemShortDescription to be an explicit nil

### UnsetItemShortDescription
`func (o *OrderLineDto) UnsetItemShortDescription()`

UnsetItemShortDescription ensures that no value is present for ItemShortDescription, not even an explicit nil
### GetItemPrimaryImageUrl

`func (o *OrderLineDto) GetItemPrimaryImageUrl() string`

GetItemPrimaryImageUrl returns the ItemPrimaryImageUrl field if non-nil, zero value otherwise.

### GetItemPrimaryImageUrlOk

`func (o *OrderLineDto) GetItemPrimaryImageUrlOk() (*string, bool)`

GetItemPrimaryImageUrlOk returns a tuple with the ItemPrimaryImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPrimaryImageUrl

`func (o *OrderLineDto) SetItemPrimaryImageUrl(v string)`

SetItemPrimaryImageUrl sets ItemPrimaryImageUrl field to given value.

### HasItemPrimaryImageUrl

`func (o *OrderLineDto) HasItemPrimaryImageUrl() bool`

HasItemPrimaryImageUrl returns a boolean if a field has been set.

### SetItemPrimaryImageUrlNil

`func (o *OrderLineDto) SetItemPrimaryImageUrlNil(b bool)`

 SetItemPrimaryImageUrlNil sets the value for ItemPrimaryImageUrl to be an explicit nil

### UnsetItemPrimaryImageUrl
`func (o *OrderLineDto) UnsetItemPrimaryImageUrl()`

UnsetItemPrimaryImageUrl ensures that no value is present for ItemPrimaryImageUrl, not even an explicit nil
### GetShippingPolicyId

`func (o *OrderLineDto) GetShippingPolicyId() string`

GetShippingPolicyId returns the ShippingPolicyId field if non-nil, zero value otherwise.

### GetShippingPolicyIdOk

`func (o *OrderLineDto) GetShippingPolicyIdOk() (*string, bool)`

GetShippingPolicyIdOk returns a tuple with the ShippingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingPolicyId

`func (o *OrderLineDto) SetShippingPolicyId(v string)`

SetShippingPolicyId sets ShippingPolicyId field to given value.

### HasShippingPolicyId

`func (o *OrderLineDto) HasShippingPolicyId() bool`

HasShippingPolicyId returns a boolean if a field has been set.

### SetShippingPolicyIdNil

`func (o *OrderLineDto) SetShippingPolicyIdNil(b bool)`

 SetShippingPolicyIdNil sets the value for ShippingPolicyId to be an explicit nil

### UnsetShippingPolicyId
`func (o *OrderLineDto) UnsetShippingPolicyId()`

UnsetShippingPolicyId ensures that no value is present for ShippingPolicyId, not even an explicit nil
### GetTenantId

`func (o *OrderLineDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *OrderLineDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *OrderLineDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *OrderLineDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *OrderLineDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *OrderLineDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *OrderLineDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *OrderLineDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *OrderLineDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *OrderLineDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *OrderLineDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *OrderLineDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetCurrencyId

`func (o *OrderLineDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *OrderLineDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *OrderLineDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *OrderLineDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *OrderLineDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *OrderLineDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetDescription

`func (o *OrderLineDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderLineDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderLineDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderLineDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OrderLineDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrderLineDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *OrderLineDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderLineDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderLineDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderLineDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetFree

`func (o *OrderLineDto) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *OrderLineDto) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *OrderLineDto) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *OrderLineDto) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetFreeReason

`func (o *OrderLineDto) GetFreeReason() string`

GetFreeReason returns the FreeReason field if non-nil, zero value otherwise.

### GetFreeReasonOk

`func (o *OrderLineDto) GetFreeReasonOk() (*string, bool)`

GetFreeReasonOk returns a tuple with the FreeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeReason

`func (o *OrderLineDto) SetFreeReason(v string)`

SetFreeReason sets FreeReason field to given value.

### HasFreeReason

`func (o *OrderLineDto) HasFreeReason() bool`

HasFreeReason returns a boolean if a field has been set.

### SetFreeReasonNil

`func (o *OrderLineDto) SetFreeReasonNil(b bool)`

 SetFreeReasonNil sets the value for FreeReason to be an explicit nil

### UnsetFreeReason
`func (o *OrderLineDto) UnsetFreeReason()`

UnsetFreeReason ensures that no value is present for FreeReason, not even an explicit nil
### GetFreeReasonCode

`func (o *OrderLineDto) GetFreeReasonCode() string`

GetFreeReasonCode returns the FreeReasonCode field if non-nil, zero value otherwise.

### GetFreeReasonCodeOk

`func (o *OrderLineDto) GetFreeReasonCodeOk() (*string, bool)`

GetFreeReasonCodeOk returns a tuple with the FreeReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeReasonCode

`func (o *OrderLineDto) SetFreeReasonCode(v string)`

SetFreeReasonCode sets FreeReasonCode field to given value.

### HasFreeReasonCode

`func (o *OrderLineDto) HasFreeReasonCode() bool`

HasFreeReasonCode returns a boolean if a field has been set.

### SetFreeReasonCodeNil

`func (o *OrderLineDto) SetFreeReasonCodeNil(b bool)`

 SetFreeReasonCodeNil sets the value for FreeReasonCode to be an explicit nil

### UnsetFreeReasonCode
`func (o *OrderLineDto) UnsetFreeReasonCode()`

UnsetFreeReasonCode ensures that no value is present for FreeReasonCode, not even an explicit nil
### GetData

`func (o *OrderLineDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderLineDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderLineDto) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *OrderLineDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *OrderLineDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OrderLineDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDataLabel

`func (o *OrderLineDto) GetDataLabel() string`

GetDataLabel returns the DataLabel field if non-nil, zero value otherwise.

### GetDataLabelOk

`func (o *OrderLineDto) GetDataLabelOk() (*string, bool)`

GetDataLabelOk returns a tuple with the DataLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLabel

`func (o *OrderLineDto) SetDataLabel(v string)`

SetDataLabel sets DataLabel field to given value.

### HasDataLabel

`func (o *OrderLineDto) HasDataLabel() bool`

HasDataLabel returns a boolean if a field has been set.

### SetDataLabelNil

`func (o *OrderLineDto) SetDataLabelNil(b bool)`

 SetDataLabelNil sets the value for DataLabel to be an explicit nil

### UnsetDataLabel
`func (o *OrderLineDto) UnsetDataLabel()`

UnsetDataLabel ensures that no value is present for DataLabel, not even an explicit nil
### GetData1

`func (o *OrderLineDto) GetData1() string`

GetData1 returns the Data1 field if non-nil, zero value otherwise.

### GetData1Ok

`func (o *OrderLineDto) GetData1Ok() (*string, bool)`

GetData1Ok returns a tuple with the Data1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1

`func (o *OrderLineDto) SetData1(v string)`

SetData1 sets Data1 field to given value.

### HasData1

`func (o *OrderLineDto) HasData1() bool`

HasData1 returns a boolean if a field has been set.

### SetData1Nil

`func (o *OrderLineDto) SetData1Nil(b bool)`

 SetData1Nil sets the value for Data1 to be an explicit nil

### UnsetData1
`func (o *OrderLineDto) UnsetData1()`

UnsetData1 ensures that no value is present for Data1, not even an explicit nil
### GetData1Label

`func (o *OrderLineDto) GetData1Label() string`

GetData1Label returns the Data1Label field if non-nil, zero value otherwise.

### GetData1LabelOk

`func (o *OrderLineDto) GetData1LabelOk() (*string, bool)`

GetData1LabelOk returns a tuple with the Data1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData1Label

`func (o *OrderLineDto) SetData1Label(v string)`

SetData1Label sets Data1Label field to given value.

### HasData1Label

`func (o *OrderLineDto) HasData1Label() bool`

HasData1Label returns a boolean if a field has been set.

### SetData1LabelNil

`func (o *OrderLineDto) SetData1LabelNil(b bool)`

 SetData1LabelNil sets the value for Data1Label to be an explicit nil

### UnsetData1Label
`func (o *OrderLineDto) UnsetData1Label()`

UnsetData1Label ensures that no value is present for Data1Label, not even an explicit nil
### GetData2

`func (o *OrderLineDto) GetData2() string`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *OrderLineDto) GetData2Ok() (*string, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *OrderLineDto) SetData2(v string)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *OrderLineDto) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### SetData2Nil

`func (o *OrderLineDto) SetData2Nil(b bool)`

 SetData2Nil sets the value for Data2 to be an explicit nil

### UnsetData2
`func (o *OrderLineDto) UnsetData2()`

UnsetData2 ensures that no value is present for Data2, not even an explicit nil
### GetData2Label

`func (o *OrderLineDto) GetData2Label() string`

GetData2Label returns the Data2Label field if non-nil, zero value otherwise.

### GetData2LabelOk

`func (o *OrderLineDto) GetData2LabelOk() (*string, bool)`

GetData2LabelOk returns a tuple with the Data2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2Label

`func (o *OrderLineDto) SetData2Label(v string)`

SetData2Label sets Data2Label field to given value.

### HasData2Label

`func (o *OrderLineDto) HasData2Label() bool`

HasData2Label returns a boolean if a field has been set.

### SetData2LabelNil

`func (o *OrderLineDto) SetData2LabelNil(b bool)`

 SetData2LabelNil sets the value for Data2Label to be an explicit nil

### UnsetData2Label
`func (o *OrderLineDto) UnsetData2Label()`

UnsetData2Label ensures that no value is present for Data2Label, not even an explicit nil
### GetData3

`func (o *OrderLineDto) GetData3() string`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *OrderLineDto) GetData3Ok() (*string, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *OrderLineDto) SetData3(v string)`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *OrderLineDto) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### SetData3Nil

`func (o *OrderLineDto) SetData3Nil(b bool)`

 SetData3Nil sets the value for Data3 to be an explicit nil

### UnsetData3
`func (o *OrderLineDto) UnsetData3()`

UnsetData3 ensures that no value is present for Data3, not even an explicit nil
### GetData3Label

`func (o *OrderLineDto) GetData3Label() string`

GetData3Label returns the Data3Label field if non-nil, zero value otherwise.

### GetData3LabelOk

`func (o *OrderLineDto) GetData3LabelOk() (*string, bool)`

GetData3LabelOk returns a tuple with the Data3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3Label

`func (o *OrderLineDto) SetData3Label(v string)`

SetData3Label sets Data3Label field to given value.

### HasData3Label

`func (o *OrderLineDto) HasData3Label() bool`

HasData3Label returns a boolean if a field has been set.

### SetData3LabelNil

`func (o *OrderLineDto) SetData3LabelNil(b bool)`

 SetData3LabelNil sets the value for Data3Label to be an explicit nil

### UnsetData3Label
`func (o *OrderLineDto) UnsetData3Label()`

UnsetData3Label ensures that no value is present for Data3Label, not even an explicit nil
### GetData4

`func (o *OrderLineDto) GetData4() string`

GetData4 returns the Data4 field if non-nil, zero value otherwise.

### GetData4Ok

`func (o *OrderLineDto) GetData4Ok() (*string, bool)`

GetData4Ok returns a tuple with the Data4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4

`func (o *OrderLineDto) SetData4(v string)`

SetData4 sets Data4 field to given value.

### HasData4

`func (o *OrderLineDto) HasData4() bool`

HasData4 returns a boolean if a field has been set.

### SetData4Nil

`func (o *OrderLineDto) SetData4Nil(b bool)`

 SetData4Nil sets the value for Data4 to be an explicit nil

### UnsetData4
`func (o *OrderLineDto) UnsetData4()`

UnsetData4 ensures that no value is present for Data4, not even an explicit nil
### GetData4Label

`func (o *OrderLineDto) GetData4Label() string`

GetData4Label returns the Data4Label field if non-nil, zero value otherwise.

### GetData4LabelOk

`func (o *OrderLineDto) GetData4LabelOk() (*string, bool)`

GetData4LabelOk returns a tuple with the Data4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData4Label

`func (o *OrderLineDto) SetData4Label(v string)`

SetData4Label sets Data4Label field to given value.

### HasData4Label

`func (o *OrderLineDto) HasData4Label() bool`

HasData4Label returns a boolean if a field has been set.

### SetData4LabelNil

`func (o *OrderLineDto) SetData4LabelNil(b bool)`

 SetData4LabelNil sets the value for Data4Label to be an explicit nil

### UnsetData4Label
`func (o *OrderLineDto) UnsetData4Label()`

UnsetData4Label ensures that no value is present for Data4Label, not even an explicit nil
### GetData5

`func (o *OrderLineDto) GetData5() string`

GetData5 returns the Data5 field if non-nil, zero value otherwise.

### GetData5Ok

`func (o *OrderLineDto) GetData5Ok() (*string, bool)`

GetData5Ok returns a tuple with the Data5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5

`func (o *OrderLineDto) SetData5(v string)`

SetData5 sets Data5 field to given value.

### HasData5

`func (o *OrderLineDto) HasData5() bool`

HasData5 returns a boolean if a field has been set.

### SetData5Nil

`func (o *OrderLineDto) SetData5Nil(b bool)`

 SetData5Nil sets the value for Data5 to be an explicit nil

### UnsetData5
`func (o *OrderLineDto) UnsetData5()`

UnsetData5 ensures that no value is present for Data5, not even an explicit nil
### GetData5Label

`func (o *OrderLineDto) GetData5Label() string`

GetData5Label returns the Data5Label field if non-nil, zero value otherwise.

### GetData5LabelOk

`func (o *OrderLineDto) GetData5LabelOk() (*string, bool)`

GetData5LabelOk returns a tuple with the Data5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData5Label

`func (o *OrderLineDto) SetData5Label(v string)`

SetData5Label sets Data5Label field to given value.

### HasData5Label

`func (o *OrderLineDto) HasData5Label() bool`

HasData5Label returns a boolean if a field has been set.

### SetData5LabelNil

`func (o *OrderLineDto) SetData5LabelNil(b bool)`

 SetData5LabelNil sets the value for Data5Label to be an explicit nil

### UnsetData5Label
`func (o *OrderLineDto) UnsetData5Label()`

UnsetData5Label ensures that no value is present for Data5Label, not even an explicit nil
### GetData6

`func (o *OrderLineDto) GetData6() string`

GetData6 returns the Data6 field if non-nil, zero value otherwise.

### GetData6Ok

`func (o *OrderLineDto) GetData6Ok() (*string, bool)`

GetData6Ok returns a tuple with the Data6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6

`func (o *OrderLineDto) SetData6(v string)`

SetData6 sets Data6 field to given value.

### HasData6

`func (o *OrderLineDto) HasData6() bool`

HasData6 returns a boolean if a field has been set.

### SetData6Nil

`func (o *OrderLineDto) SetData6Nil(b bool)`

 SetData6Nil sets the value for Data6 to be an explicit nil

### UnsetData6
`func (o *OrderLineDto) UnsetData6()`

UnsetData6 ensures that no value is present for Data6, not even an explicit nil
### GetData6Label

`func (o *OrderLineDto) GetData6Label() string`

GetData6Label returns the Data6Label field if non-nil, zero value otherwise.

### GetData6LabelOk

`func (o *OrderLineDto) GetData6LabelOk() (*string, bool)`

GetData6LabelOk returns a tuple with the Data6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData6Label

`func (o *OrderLineDto) SetData6Label(v string)`

SetData6Label sets Data6Label field to given value.

### HasData6Label

`func (o *OrderLineDto) HasData6Label() bool`

HasData6Label returns a boolean if a field has been set.

### SetData6LabelNil

`func (o *OrderLineDto) SetData6LabelNil(b bool)`

 SetData6LabelNil sets the value for Data6Label to be an explicit nil

### UnsetData6Label
`func (o *OrderLineDto) UnsetData6Label()`

UnsetData6Label ensures that no value is present for Data6Label, not even an explicit nil
### GetData7

`func (o *OrderLineDto) GetData7() string`

GetData7 returns the Data7 field if non-nil, zero value otherwise.

### GetData7Ok

`func (o *OrderLineDto) GetData7Ok() (*string, bool)`

GetData7Ok returns a tuple with the Data7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7

`func (o *OrderLineDto) SetData7(v string)`

SetData7 sets Data7 field to given value.

### HasData7

`func (o *OrderLineDto) HasData7() bool`

HasData7 returns a boolean if a field has been set.

### SetData7Nil

`func (o *OrderLineDto) SetData7Nil(b bool)`

 SetData7Nil sets the value for Data7 to be an explicit nil

### UnsetData7
`func (o *OrderLineDto) UnsetData7()`

UnsetData7 ensures that no value is present for Data7, not even an explicit nil
### GetData7Label

`func (o *OrderLineDto) GetData7Label() string`

GetData7Label returns the Data7Label field if non-nil, zero value otherwise.

### GetData7LabelOk

`func (o *OrderLineDto) GetData7LabelOk() (*string, bool)`

GetData7LabelOk returns a tuple with the Data7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData7Label

`func (o *OrderLineDto) SetData7Label(v string)`

SetData7Label sets Data7Label field to given value.

### HasData7Label

`func (o *OrderLineDto) HasData7Label() bool`

HasData7Label returns a boolean if a field has been set.

### SetData7LabelNil

`func (o *OrderLineDto) SetData7LabelNil(b bool)`

 SetData7LabelNil sets the value for Data7Label to be an explicit nil

### UnsetData7Label
`func (o *OrderLineDto) UnsetData7Label()`

UnsetData7Label ensures that no value is present for Data7Label, not even an explicit nil
### GetData8

`func (o *OrderLineDto) GetData8() string`

GetData8 returns the Data8 field if non-nil, zero value otherwise.

### GetData8Ok

`func (o *OrderLineDto) GetData8Ok() (*string, bool)`

GetData8Ok returns a tuple with the Data8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8

`func (o *OrderLineDto) SetData8(v string)`

SetData8 sets Data8 field to given value.

### HasData8

`func (o *OrderLineDto) HasData8() bool`

HasData8 returns a boolean if a field has been set.

### SetData8Nil

`func (o *OrderLineDto) SetData8Nil(b bool)`

 SetData8Nil sets the value for Data8 to be an explicit nil

### UnsetData8
`func (o *OrderLineDto) UnsetData8()`

UnsetData8 ensures that no value is present for Data8, not even an explicit nil
### GetData8Label

`func (o *OrderLineDto) GetData8Label() string`

GetData8Label returns the Data8Label field if non-nil, zero value otherwise.

### GetData8LabelOk

`func (o *OrderLineDto) GetData8LabelOk() (*string, bool)`

GetData8LabelOk returns a tuple with the Data8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData8Label

`func (o *OrderLineDto) SetData8Label(v string)`

SetData8Label sets Data8Label field to given value.

### HasData8Label

`func (o *OrderLineDto) HasData8Label() bool`

HasData8Label returns a boolean if a field has been set.

### SetData8LabelNil

`func (o *OrderLineDto) SetData8LabelNil(b bool)`

 SetData8LabelNil sets the value for Data8Label to be an explicit nil

### UnsetData8Label
`func (o *OrderLineDto) UnsetData8Label()`

UnsetData8Label ensures that no value is present for Data8Label, not even an explicit nil
### GetData9

`func (o *OrderLineDto) GetData9() string`

GetData9 returns the Data9 field if non-nil, zero value otherwise.

### GetData9Ok

`func (o *OrderLineDto) GetData9Ok() (*string, bool)`

GetData9Ok returns a tuple with the Data9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9

`func (o *OrderLineDto) SetData9(v string)`

SetData9 sets Data9 field to given value.

### HasData9

`func (o *OrderLineDto) HasData9() bool`

HasData9 returns a boolean if a field has been set.

### SetData9Nil

`func (o *OrderLineDto) SetData9Nil(b bool)`

 SetData9Nil sets the value for Data9 to be an explicit nil

### UnsetData9
`func (o *OrderLineDto) UnsetData9()`

UnsetData9 ensures that no value is present for Data9, not even an explicit nil
### GetData9Label

`func (o *OrderLineDto) GetData9Label() string`

GetData9Label returns the Data9Label field if non-nil, zero value otherwise.

### GetData9LabelOk

`func (o *OrderLineDto) GetData9LabelOk() (*string, bool)`

GetData9LabelOk returns a tuple with the Data9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData9Label

`func (o *OrderLineDto) SetData9Label(v string)`

SetData9Label sets Data9Label field to given value.

### HasData9Label

`func (o *OrderLineDto) HasData9Label() bool`

HasData9Label returns a boolean if a field has been set.

### SetData9LabelNil

`func (o *OrderLineDto) SetData9LabelNil(b bool)`

 SetData9LabelNil sets the value for Data9Label to be an explicit nil

### UnsetData9Label
`func (o *OrderLineDto) UnsetData9Label()`

UnsetData9Label ensures that no value is present for Data9Label, not even an explicit nil
### GetItemPriceId

`func (o *OrderLineDto) GetItemPriceId() string`

GetItemPriceId returns the ItemPriceId field if non-nil, zero value otherwise.

### GetItemPriceIdOk

`func (o *OrderLineDto) GetItemPriceIdOk() (*string, bool)`

GetItemPriceIdOk returns a tuple with the ItemPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceId

`func (o *OrderLineDto) SetItemPriceId(v string)`

SetItemPriceId sets ItemPriceId field to given value.

### HasItemPriceId

`func (o *OrderLineDto) HasItemPriceId() bool`

HasItemPriceId returns a boolean if a field has been set.

### SetItemPriceIdNil

`func (o *OrderLineDto) SetItemPriceIdNil(b bool)`

 SetItemPriceIdNil sets the value for ItemPriceId to be an explicit nil

### UnsetItemPriceId
`func (o *OrderLineDto) UnsetItemPriceId()`

UnsetItemPriceId ensures that no value is present for ItemPriceId, not even an explicit nil
### GetPriceListItemId

`func (o *OrderLineDto) GetPriceListItemId() string`

GetPriceListItemId returns the PriceListItemId field if non-nil, zero value otherwise.

### GetPriceListItemIdOk

`func (o *OrderLineDto) GetPriceListItemIdOk() (*string, bool)`

GetPriceListItemIdOk returns a tuple with the PriceListItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceListItemId

`func (o *OrderLineDto) SetPriceListItemId(v string)`

SetPriceListItemId sets PriceListItemId field to given value.

### HasPriceListItemId

`func (o *OrderLineDto) HasPriceListItemId() bool`

HasPriceListItemId returns a boolean if a field has been set.

### SetPriceListItemIdNil

`func (o *OrderLineDto) SetPriceListItemIdNil(b bool)`

 SetPriceListItemIdNil sets the value for PriceListItemId to be an explicit nil

### UnsetPriceListItemId
`func (o *OrderLineDto) UnsetPriceListItemId()`

UnsetPriceListItemId ensures that no value is present for PriceListItemId, not even an explicit nil
### GetUnitId

`func (o *OrderLineDto) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *OrderLineDto) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *OrderLineDto) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *OrderLineDto) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *OrderLineDto) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *OrderLineDto) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetUnitGroupId

`func (o *OrderLineDto) GetUnitGroupId() string`

GetUnitGroupId returns the UnitGroupId field if non-nil, zero value otherwise.

### GetUnitGroupIdOk

`func (o *OrderLineDto) GetUnitGroupIdOk() (*string, bool)`

GetUnitGroupIdOk returns a tuple with the UnitGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitGroupId

`func (o *OrderLineDto) SetUnitGroupId(v string)`

SetUnitGroupId sets UnitGroupId field to given value.

### HasUnitGroupId

`func (o *OrderLineDto) HasUnitGroupId() bool`

HasUnitGroupId returns a boolean if a field has been set.

### SetUnitGroupIdNil

`func (o *OrderLineDto) SetUnitGroupIdNil(b bool)`

 SetUnitGroupIdNil sets the value for UnitGroupId to be an explicit nil

### UnsetUnitGroupId
`func (o *OrderLineDto) UnsetUnitGroupId()`

UnsetUnitGroupId ensures that no value is present for UnitGroupId, not even an explicit nil
### GetTaxCalculationMethod

`func (o *OrderLineDto) GetTaxCalculationMethod() string`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *OrderLineDto) GetTaxCalculationMethodOk() (*string, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *OrderLineDto) SetTaxCalculationMethod(v string)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *OrderLineDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *OrderLineDto) GetCostCalculationMethod() string`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *OrderLineDto) GetCostCalculationMethodOk() (*string, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *OrderLineDto) SetCostCalculationMethod(v string)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *OrderLineDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetForexRates

`func (o *OrderLineDto) GetForexRates() ForexRates`

GetForexRates returns the ForexRates field if non-nil, zero value otherwise.

### GetForexRatesOk

`func (o *OrderLineDto) GetForexRatesOk() (*ForexRates, bool)`

GetForexRatesOk returns a tuple with the ForexRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRates

`func (o *OrderLineDto) SetForexRates(v ForexRates)`

SetForexRates sets ForexRates field to given value.

### HasForexRates

`func (o *OrderLineDto) HasForexRates() bool`

HasForexRates returns a boolean if a field has been set.

### GetForexRate

`func (o *OrderLineDto) GetForexRate() float64`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *OrderLineDto) GetForexRateOk() (*float64, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *OrderLineDto) SetForexRate(v float64)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *OrderLineDto) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetTotalDetailInUsd

`func (o *OrderLineDto) GetTotalDetailInUsd() float64`

GetTotalDetailInUsd returns the TotalDetailInUsd field if non-nil, zero value otherwise.

### GetTotalDetailInUsdOk

`func (o *OrderLineDto) GetTotalDetailInUsdOk() (*float64, bool)`

GetTotalDetailInUsdOk returns a tuple with the TotalDetailInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailInUsd

`func (o *OrderLineDto) SetTotalDetailInUsd(v float64)`

SetTotalDetailInUsd sets TotalDetailInUsd field to given value.

### HasTotalDetailInUsd

`func (o *OrderLineDto) HasTotalDetailInUsd() bool`

HasTotalDetailInUsd returns a boolean if a field has been set.

### GetTotalProfitInUsd

`func (o *OrderLineDto) GetTotalProfitInUsd() float64`

GetTotalProfitInUsd returns the TotalProfitInUsd field if non-nil, zero value otherwise.

### GetTotalProfitInUsdOk

`func (o *OrderLineDto) GetTotalProfitInUsdOk() (*float64, bool)`

GetTotalProfitInUsdOk returns a tuple with the TotalProfitInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitInUsd

`func (o *OrderLineDto) SetTotalProfitInUsd(v float64)`

SetTotalProfitInUsd sets TotalProfitInUsd field to given value.

### HasTotalProfitInUsd

`func (o *OrderLineDto) HasTotalProfitInUsd() bool`

HasTotalProfitInUsd returns a boolean if a field has been set.

### GetTotalDiscountsInUsd

`func (o *OrderLineDto) GetTotalDiscountsInUsd() float64`

GetTotalDiscountsInUsd returns the TotalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalDiscountsInUsdOk

`func (o *OrderLineDto) GetTotalDiscountsInUsdOk() (*float64, bool)`

GetTotalDiscountsInUsdOk returns a tuple with the TotalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsInUsd

`func (o *OrderLineDto) SetTotalDiscountsInUsd(v float64)`

SetTotalDiscountsInUsd sets TotalDiscountsInUsd field to given value.

### HasTotalDiscountsInUsd

`func (o *OrderLineDto) HasTotalDiscountsInUsd() bool`

HasTotalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalSurchargesInUsd

`func (o *OrderLineDto) GetTotalSurchargesInUsd() float64`

GetTotalSurchargesInUsd returns the TotalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalSurchargesInUsdOk

`func (o *OrderLineDto) GetTotalSurchargesInUsdOk() (*float64, bool)`

GetTotalSurchargesInUsdOk returns a tuple with the TotalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesInUsd

`func (o *OrderLineDto) SetTotalSurchargesInUsd(v float64)`

SetTotalSurchargesInUsd sets TotalSurchargesInUsd field to given value.

### HasTotalSurchargesInUsd

`func (o *OrderLineDto) HasTotalSurchargesInUsd() bool`

HasTotalSurchargesInUsd returns a boolean if a field has been set.

### GetTotalTaxBaseInUsd

`func (o *OrderLineDto) GetTotalTaxBaseInUsd() float64`

GetTotalTaxBaseInUsd returns the TotalTaxBaseInUsd field if non-nil, zero value otherwise.

### GetTotalTaxBaseInUsdOk

`func (o *OrderLineDto) GetTotalTaxBaseInUsdOk() (*float64, bool)`

GetTotalTaxBaseInUsdOk returns a tuple with the TotalTaxBaseInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseInUsd

`func (o *OrderLineDto) SetTotalTaxBaseInUsd(v float64)`

SetTotalTaxBaseInUsd sets TotalTaxBaseInUsd field to given value.

### HasTotalTaxBaseInUsd

`func (o *OrderLineDto) HasTotalTaxBaseInUsd() bool`

HasTotalTaxBaseInUsd returns a boolean if a field has been set.

### GetTotalTaxesInUsd

`func (o *OrderLineDto) GetTotalTaxesInUsd() float64`

GetTotalTaxesInUsd returns the TotalTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalTaxesInUsdOk

`func (o *OrderLineDto) GetTotalTaxesInUsdOk() (*float64, bool)`

GetTotalTaxesInUsdOk returns a tuple with the TotalTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesInUsd

`func (o *OrderLineDto) SetTotalTaxesInUsd(v float64)`

SetTotalTaxesInUsd sets TotalTaxesInUsd field to given value.

### HasTotalTaxesInUsd

`func (o *OrderLineDto) HasTotalTaxesInUsd() bool`

HasTotalTaxesInUsd returns a boolean if a field has been set.

### GetTotalWithheldTaxesInUsd

`func (o *OrderLineDto) GetTotalWithheldTaxesInUsd() float64`

GetTotalWithheldTaxesInUsd returns the TotalWithheldTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithheldTaxesInUsdOk

`func (o *OrderLineDto) GetTotalWithheldTaxesInUsdOk() (*float64, bool)`

GetTotalWithheldTaxesInUsdOk returns a tuple with the TotalWithheldTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxesInUsd

`func (o *OrderLineDto) SetTotalWithheldTaxesInUsd(v float64)`

SetTotalWithheldTaxesInUsd sets TotalWithheldTaxesInUsd field to given value.

### HasTotalWithheldTaxesInUsd

`func (o *OrderLineDto) HasTotalWithheldTaxesInUsd() bool`

HasTotalWithheldTaxesInUsd returns a boolean if a field has been set.

### GetTotalShippingCostInUsd

`func (o *OrderLineDto) GetTotalShippingCostInUsd() float64`

GetTotalShippingCostInUsd returns the TotalShippingCostInUsd field if non-nil, zero value otherwise.

### GetTotalShippingCostInUsdOk

`func (o *OrderLineDto) GetTotalShippingCostInUsdOk() (*float64, bool)`

GetTotalShippingCostInUsdOk returns a tuple with the TotalShippingCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostInUsd

`func (o *OrderLineDto) SetTotalShippingCostInUsd(v float64)`

SetTotalShippingCostInUsd sets TotalShippingCostInUsd field to given value.

### HasTotalShippingCostInUsd

`func (o *OrderLineDto) HasTotalShippingCostInUsd() bool`

HasTotalShippingCostInUsd returns a boolean if a field has been set.

### GetTotalShippingTaxesInUsd

`func (o *OrderLineDto) GetTotalShippingTaxesInUsd() float64`

GetTotalShippingTaxesInUsd returns the TotalShippingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalShippingTaxesInUsdOk

`func (o *OrderLineDto) GetTotalShippingTaxesInUsdOk() (*float64, bool)`

GetTotalShippingTaxesInUsdOk returns a tuple with the TotalShippingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxesInUsd

`func (o *OrderLineDto) SetTotalShippingTaxesInUsd(v float64)`

SetTotalShippingTaxesInUsd sets TotalShippingTaxesInUsd field to given value.

### HasTotalShippingTaxesInUsd

`func (o *OrderLineDto) HasTotalShippingTaxesInUsd() bool`

HasTotalShippingTaxesInUsd returns a boolean if a field has been set.

### GetTotalWarrantyCostInUsd

`func (o *OrderLineDto) GetTotalWarrantyCostInUsd() float64`

GetTotalWarrantyCostInUsd returns the TotalWarrantyCostInUsd field if non-nil, zero value otherwise.

### GetTotalWarrantyCostInUsdOk

`func (o *OrderLineDto) GetTotalWarrantyCostInUsdOk() (*float64, bool)`

GetTotalWarrantyCostInUsdOk returns a tuple with the TotalWarrantyCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWarrantyCostInUsd

`func (o *OrderLineDto) SetTotalWarrantyCostInUsd(v float64)`

SetTotalWarrantyCostInUsd sets TotalWarrantyCostInUsd field to given value.

### HasTotalWarrantyCostInUsd

`func (o *OrderLineDto) HasTotalWarrantyCostInUsd() bool`

HasTotalWarrantyCostInUsd returns a boolean if a field has been set.

### GetTotalReturnCostInUsd

`func (o *OrderLineDto) GetTotalReturnCostInUsd() float64`

GetTotalReturnCostInUsd returns the TotalReturnCostInUsd field if non-nil, zero value otherwise.

### GetTotalReturnCostInUsdOk

`func (o *OrderLineDto) GetTotalReturnCostInUsdOk() (*float64, bool)`

GetTotalReturnCostInUsdOk returns a tuple with the TotalReturnCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReturnCostInUsd

`func (o *OrderLineDto) SetTotalReturnCostInUsd(v float64)`

SetTotalReturnCostInUsd sets TotalReturnCostInUsd field to given value.

### HasTotalReturnCostInUsd

`func (o *OrderLineDto) HasTotalReturnCostInUsd() bool`

HasTotalReturnCostInUsd returns a boolean if a field has been set.

### GetTotalRefundCostInUsd

`func (o *OrderLineDto) GetTotalRefundCostInUsd() float64`

GetTotalRefundCostInUsd returns the TotalRefundCostInUsd field if non-nil, zero value otherwise.

### GetTotalRefundCostInUsdOk

`func (o *OrderLineDto) GetTotalRefundCostInUsdOk() (*float64, bool)`

GetTotalRefundCostInUsdOk returns a tuple with the TotalRefundCostInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRefundCostInUsd

`func (o *OrderLineDto) SetTotalRefundCostInUsd(v float64)`

SetTotalRefundCostInUsd sets TotalRefundCostInUsd field to given value.

### HasTotalRefundCostInUsd

`func (o *OrderLineDto) HasTotalRefundCostInUsd() bool`

HasTotalRefundCostInUsd returns a boolean if a field has been set.

### GetTotalInUsd

`func (o *OrderLineDto) GetTotalInUsd() float64`

GetTotalInUsd returns the TotalInUsd field if non-nil, zero value otherwise.

### GetTotalInUsdOk

`func (o *OrderLineDto) GetTotalInUsdOk() (*float64, bool)`

GetTotalInUsdOk returns a tuple with the TotalInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInUsd

`func (o *OrderLineDto) SetTotalInUsd(v float64)`

SetTotalInUsd sets TotalInUsd field to given value.

### HasTotalInUsd

`func (o *OrderLineDto) HasTotalInUsd() bool`

HasTotalInUsd returns a boolean if a field has been set.

### GetTotalGlobalDiscountsInUsd

`func (o *OrderLineDto) GetTotalGlobalDiscountsInUsd() float64`

GetTotalGlobalDiscountsInUsd returns the TotalGlobalDiscountsInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsInUsdOk

`func (o *OrderLineDto) GetTotalGlobalDiscountsInUsdOk() (*float64, bool)`

GetTotalGlobalDiscountsInUsdOk returns a tuple with the TotalGlobalDiscountsInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsInUsd

`func (o *OrderLineDto) SetTotalGlobalDiscountsInUsd(v float64)`

SetTotalGlobalDiscountsInUsd sets TotalGlobalDiscountsInUsd field to given value.

### HasTotalGlobalDiscountsInUsd

`func (o *OrderLineDto) HasTotalGlobalDiscountsInUsd() bool`

HasTotalGlobalDiscountsInUsd returns a boolean if a field has been set.

### GetTotalGlobalSurchargesInUsd

`func (o *OrderLineDto) GetTotalGlobalSurchargesInUsd() float64`

GetTotalGlobalSurchargesInUsd returns the TotalGlobalSurchargesInUsd field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesInUsdOk

`func (o *OrderLineDto) GetTotalGlobalSurchargesInUsdOk() (*float64, bool)`

GetTotalGlobalSurchargesInUsdOk returns a tuple with the TotalGlobalSurchargesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesInUsd

`func (o *OrderLineDto) SetTotalGlobalSurchargesInUsd(v float64)`

SetTotalGlobalSurchargesInUsd sets TotalGlobalSurchargesInUsd field to given value.

### HasTotalGlobalSurchargesInUsd

`func (o *OrderLineDto) HasTotalGlobalSurchargesInUsd() bool`

HasTotalGlobalSurchargesInUsd returns a boolean if a field has been set.

### GetCustomGlobalSurchargesAmount

`func (o *OrderLineDto) GetCustomGlobalSurchargesAmount() float64`

GetCustomGlobalSurchargesAmount returns the CustomGlobalSurchargesAmount field if non-nil, zero value otherwise.

### GetCustomGlobalSurchargesAmountOk

`func (o *OrderLineDto) GetCustomGlobalSurchargesAmountOk() (*float64, bool)`

GetCustomGlobalSurchargesAmountOk returns a tuple with the CustomGlobalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalSurchargesAmount

`func (o *OrderLineDto) SetCustomGlobalSurchargesAmount(v float64)`

SetCustomGlobalSurchargesAmount sets CustomGlobalSurchargesAmount field to given value.

### HasCustomGlobalSurchargesAmount

`func (o *OrderLineDto) HasCustomGlobalSurchargesAmount() bool`

HasCustomGlobalSurchargesAmount returns a boolean if a field has been set.

### GetCustomGlobalDiscountsAmount

`func (o *OrderLineDto) GetCustomGlobalDiscountsAmount() float64`

GetCustomGlobalDiscountsAmount returns the CustomGlobalDiscountsAmount field if non-nil, zero value otherwise.

### GetCustomGlobalDiscountsAmountOk

`func (o *OrderLineDto) GetCustomGlobalDiscountsAmountOk() (*float64, bool)`

GetCustomGlobalDiscountsAmountOk returns a tuple with the CustomGlobalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomGlobalDiscountsAmount

`func (o *OrderLineDto) SetCustomGlobalDiscountsAmount(v float64)`

SetCustomGlobalDiscountsAmount sets CustomGlobalDiscountsAmount field to given value.

### HasCustomGlobalDiscountsAmount

`func (o *OrderLineDto) HasCustomGlobalDiscountsAmount() bool`

HasCustomGlobalDiscountsAmount returns a boolean if a field has been set.

### GetReturnPolicyId

`func (o *OrderLineDto) GetReturnPolicyId() string`

GetReturnPolicyId returns the ReturnPolicyId field if non-nil, zero value otherwise.

### GetReturnPolicyIdOk

`func (o *OrderLineDto) GetReturnPolicyIdOk() (*string, bool)`

GetReturnPolicyIdOk returns a tuple with the ReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyId

`func (o *OrderLineDto) SetReturnPolicyId(v string)`

SetReturnPolicyId sets ReturnPolicyId field to given value.

### HasReturnPolicyId

`func (o *OrderLineDto) HasReturnPolicyId() bool`

HasReturnPolicyId returns a boolean if a field has been set.

### SetReturnPolicyIdNil

`func (o *OrderLineDto) SetReturnPolicyIdNil(b bool)`

 SetReturnPolicyIdNil sets the value for ReturnPolicyId to be an explicit nil

### UnsetReturnPolicyId
`func (o *OrderLineDto) UnsetReturnPolicyId()`

UnsetReturnPolicyId ensures that no value is present for ReturnPolicyId, not even an explicit nil
### GetRefundPolicyId

`func (o *OrderLineDto) GetRefundPolicyId() string`

GetRefundPolicyId returns the RefundPolicyId field if non-nil, zero value otherwise.

### GetRefundPolicyIdOk

`func (o *OrderLineDto) GetRefundPolicyIdOk() (*string, bool)`

GetRefundPolicyIdOk returns a tuple with the RefundPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPolicyId

`func (o *OrderLineDto) SetRefundPolicyId(v string)`

SetRefundPolicyId sets RefundPolicyId field to given value.

### HasRefundPolicyId

`func (o *OrderLineDto) HasRefundPolicyId() bool`

HasRefundPolicyId returns a boolean if a field has been set.

### SetRefundPolicyIdNil

`func (o *OrderLineDto) SetRefundPolicyIdNil(b bool)`

 SetRefundPolicyIdNil sets the value for RefundPolicyId to be an explicit nil

### UnsetRefundPolicyId
`func (o *OrderLineDto) UnsetRefundPolicyId()`

UnsetRefundPolicyId ensures that no value is present for RefundPolicyId, not even an explicit nil
### GetWarrantyPolicyId

`func (o *OrderLineDto) GetWarrantyPolicyId() string`

GetWarrantyPolicyId returns the WarrantyPolicyId field if non-nil, zero value otherwise.

### GetWarrantyPolicyIdOk

`func (o *OrderLineDto) GetWarrantyPolicyIdOk() (*string, bool)`

GetWarrantyPolicyIdOk returns a tuple with the WarrantyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyPolicyId

`func (o *OrderLineDto) SetWarrantyPolicyId(v string)`

SetWarrantyPolicyId sets WarrantyPolicyId field to given value.

### HasWarrantyPolicyId

`func (o *OrderLineDto) HasWarrantyPolicyId() bool`

HasWarrantyPolicyId returns a boolean if a field has been set.

### SetWarrantyPolicyIdNil

`func (o *OrderLineDto) SetWarrantyPolicyIdNil(b bool)`

 SetWarrantyPolicyIdNil sets the value for WarrantyPolicyId to be an explicit nil

### UnsetWarrantyPolicyId
`func (o *OrderLineDto) UnsetWarrantyPolicyId()`

UnsetWarrantyPolicyId ensures that no value is present for WarrantyPolicyId, not even an explicit nil
### GetShipmentPolicyId

`func (o *OrderLineDto) GetShipmentPolicyId() string`

GetShipmentPolicyId returns the ShipmentPolicyId field if non-nil, zero value otherwise.

### GetShipmentPolicyIdOk

`func (o *OrderLineDto) GetShipmentPolicyIdOk() (*string, bool)`

GetShipmentPolicyIdOk returns a tuple with the ShipmentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentPolicyId

`func (o *OrderLineDto) SetShipmentPolicyId(v string)`

SetShipmentPolicyId sets ShipmentPolicyId field to given value.

### HasShipmentPolicyId

`func (o *OrderLineDto) HasShipmentPolicyId() bool`

HasShipmentPolicyId returns a boolean if a field has been set.

### SetShipmentPolicyIdNil

`func (o *OrderLineDto) SetShipmentPolicyIdNil(b bool)`

 SetShipmentPolicyIdNil sets the value for ShipmentPolicyId to be an explicit nil

### UnsetShipmentPolicyId
`func (o *OrderLineDto) UnsetShipmentPolicyId()`

UnsetShipmentPolicyId ensures that no value is present for ShipmentPolicyId, not even an explicit nil
### GetShippingLocationId

`func (o *OrderLineDto) GetShippingLocationId() string`

GetShippingLocationId returns the ShippingLocationId field if non-nil, zero value otherwise.

### GetShippingLocationIdOk

`func (o *OrderLineDto) GetShippingLocationIdOk() (*string, bool)`

GetShippingLocationIdOk returns a tuple with the ShippingLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLocationId

`func (o *OrderLineDto) SetShippingLocationId(v string)`

SetShippingLocationId sets ShippingLocationId field to given value.

### HasShippingLocationId

`func (o *OrderLineDto) HasShippingLocationId() bool`

HasShippingLocationId returns a boolean if a field has been set.

### SetShippingLocationIdNil

`func (o *OrderLineDto) SetShippingLocationIdNil(b bool)`

 SetShippingLocationIdNil sets the value for ShippingLocationId to be an explicit nil

### UnsetShippingLocationId
`func (o *OrderLineDto) UnsetShippingLocationId()`

UnsetShippingLocationId ensures that no value is present for ShippingLocationId, not even an explicit nil
### GetLocationId

`func (o *OrderLineDto) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *OrderLineDto) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *OrderLineDto) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *OrderLineDto) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *OrderLineDto) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *OrderLineDto) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetQuoteItemRecordId

`func (o *OrderLineDto) GetQuoteItemRecordId() string`

GetQuoteItemRecordId returns the QuoteItemRecordId field if non-nil, zero value otherwise.

### GetQuoteItemRecordIdOk

`func (o *OrderLineDto) GetQuoteItemRecordIdOk() (*string, bool)`

GetQuoteItemRecordIdOk returns a tuple with the QuoteItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteItemRecordId

`func (o *OrderLineDto) SetQuoteItemRecordId(v string)`

SetQuoteItemRecordId sets QuoteItemRecordId field to given value.

### HasQuoteItemRecordId

`func (o *OrderLineDto) HasQuoteItemRecordId() bool`

HasQuoteItemRecordId returns a boolean if a field has been set.

### SetQuoteItemRecordIdNil

`func (o *OrderLineDto) SetQuoteItemRecordIdNil(b bool)`

 SetQuoteItemRecordIdNil sets the value for QuoteItemRecordId to be an explicit nil

### UnsetQuoteItemRecordId
`func (o *OrderLineDto) UnsetQuoteItemRecordId()`

UnsetQuoteItemRecordId ensures that no value is present for QuoteItemRecordId, not even an explicit nil
### GetBusinessProfileRecordId

`func (o *OrderLineDto) GetBusinessProfileRecordId() string`

GetBusinessProfileRecordId returns the BusinessProfileRecordId field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIdOk

`func (o *OrderLineDto) GetBusinessProfileRecordIdOk() (*string, bool)`

GetBusinessProfileRecordIdOk returns a tuple with the BusinessProfileRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordId

`func (o *OrderLineDto) SetBusinessProfileRecordId(v string)`

SetBusinessProfileRecordId sets BusinessProfileRecordId field to given value.

### HasBusinessProfileRecordId

`func (o *OrderLineDto) HasBusinessProfileRecordId() bool`

HasBusinessProfileRecordId returns a boolean if a field has been set.

### SetBusinessProfileRecordIdNil

`func (o *OrderLineDto) SetBusinessProfileRecordIdNil(b bool)`

 SetBusinessProfileRecordIdNil sets the value for BusinessProfileRecordId to be an explicit nil

### UnsetBusinessProfileRecordId
`func (o *OrderLineDto) UnsetBusinessProfileRecordId()`

UnsetBusinessProfileRecordId ensures that no value is present for BusinessProfileRecordId, not even an explicit nil
### GetParentBillingItemRecordId

`func (o *OrderLineDto) GetParentBillingItemRecordId() string`

GetParentBillingItemRecordId returns the ParentBillingItemRecordId field if non-nil, zero value otherwise.

### GetParentBillingItemRecordIdOk

`func (o *OrderLineDto) GetParentBillingItemRecordIdOk() (*string, bool)`

GetParentBillingItemRecordIdOk returns a tuple with the ParentBillingItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBillingItemRecordId

`func (o *OrderLineDto) SetParentBillingItemRecordId(v string)`

SetParentBillingItemRecordId sets ParentBillingItemRecordId field to given value.

### HasParentBillingItemRecordId

`func (o *OrderLineDto) HasParentBillingItemRecordId() bool`

HasParentBillingItemRecordId returns a boolean if a field has been set.

### SetParentBillingItemRecordIdNil

`func (o *OrderLineDto) SetParentBillingItemRecordIdNil(b bool)`

 SetParentBillingItemRecordIdNil sets the value for ParentBillingItemRecordId to be an explicit nil

### UnsetParentBillingItemRecordId
`func (o *OrderLineDto) UnsetParentBillingItemRecordId()`

UnsetParentBillingItemRecordId ensures that no value is present for ParentBillingItemRecordId, not even an explicit nil
### GetCurrency

`func (o *OrderLineDto) GetCurrency() CurrencyId`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderLineDto) GetCurrencyOk() (*CurrencyId, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderLineDto) SetCurrency(v CurrencyId)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderLineDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTotalDetail

`func (o *OrderLineDto) GetTotalDetail() float64`

GetTotalDetail returns the TotalDetail field if non-nil, zero value otherwise.

### GetTotalDetailOk

`func (o *OrderLineDto) GetTotalDetailOk() (*float64, bool)`

GetTotalDetailOk returns a tuple with the TotalDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetail

`func (o *OrderLineDto) SetTotalDetail(v float64)`

SetTotalDetail sets TotalDetail field to given value.

### HasTotalDetail

`func (o *OrderLineDto) HasTotalDetail() bool`

HasTotalDetail returns a boolean if a field has been set.

### GetTotalDetailCurrencyId

`func (o *OrderLineDto) GetTotalDetailCurrencyId() string`

GetTotalDetailCurrencyId returns the TotalDetailCurrencyId field if non-nil, zero value otherwise.

### GetTotalDetailCurrencyIdOk

`func (o *OrderLineDto) GetTotalDetailCurrencyIdOk() (*string, bool)`

GetTotalDetailCurrencyIdOk returns a tuple with the TotalDetailCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailCurrencyId

`func (o *OrderLineDto) SetTotalDetailCurrencyId(v string)`

SetTotalDetailCurrencyId sets TotalDetailCurrencyId field to given value.

### HasTotalDetailCurrencyId

`func (o *OrderLineDto) HasTotalDetailCurrencyId() bool`

HasTotalDetailCurrencyId returns a boolean if a field has been set.

### SetTotalDetailCurrencyIdNil

`func (o *OrderLineDto) SetTotalDetailCurrencyIdNil(b bool)`

 SetTotalDetailCurrencyIdNil sets the value for TotalDetailCurrencyId to be an explicit nil

### UnsetTotalDetailCurrencyId
`func (o *OrderLineDto) UnsetTotalDetailCurrencyId()`

UnsetTotalDetailCurrencyId ensures that no value is present for TotalDetailCurrencyId, not even an explicit nil
### GetTotalDetailAmount

`func (o *OrderLineDto) GetTotalDetailAmount() Money`

GetTotalDetailAmount returns the TotalDetailAmount field if non-nil, zero value otherwise.

### GetTotalDetailAmountOk

`func (o *OrderLineDto) GetTotalDetailAmountOk() (*Money, bool)`

GetTotalDetailAmountOk returns a tuple with the TotalDetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailAmount

`func (o *OrderLineDto) SetTotalDetailAmount(v Money)`

SetTotalDetailAmount sets TotalDetailAmount field to given value.

### HasTotalDetailAmount

`func (o *OrderLineDto) HasTotalDetailAmount() bool`

HasTotalDetailAmount returns a boolean if a field has been set.

### GetTotalProfit

`func (o *OrderLineDto) GetTotalProfit() float64`

GetTotalProfit returns the TotalProfit field if non-nil, zero value otherwise.

### GetTotalProfitOk

`func (o *OrderLineDto) GetTotalProfitOk() (*float64, bool)`

GetTotalProfitOk returns a tuple with the TotalProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfit

`func (o *OrderLineDto) SetTotalProfit(v float64)`

SetTotalProfit sets TotalProfit field to given value.

### HasTotalProfit

`func (o *OrderLineDto) HasTotalProfit() bool`

HasTotalProfit returns a boolean if a field has been set.

### GetTotalProfitCurrencyId

`func (o *OrderLineDto) GetTotalProfitCurrencyId() string`

GetTotalProfitCurrencyId returns the TotalProfitCurrencyId field if non-nil, zero value otherwise.

### GetTotalProfitCurrencyIdOk

`func (o *OrderLineDto) GetTotalProfitCurrencyIdOk() (*string, bool)`

GetTotalProfitCurrencyIdOk returns a tuple with the TotalProfitCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitCurrencyId

`func (o *OrderLineDto) SetTotalProfitCurrencyId(v string)`

SetTotalProfitCurrencyId sets TotalProfitCurrencyId field to given value.

### HasTotalProfitCurrencyId

`func (o *OrderLineDto) HasTotalProfitCurrencyId() bool`

HasTotalProfitCurrencyId returns a boolean if a field has been set.

### SetTotalProfitCurrencyIdNil

`func (o *OrderLineDto) SetTotalProfitCurrencyIdNil(b bool)`

 SetTotalProfitCurrencyIdNil sets the value for TotalProfitCurrencyId to be an explicit nil

### UnsetTotalProfitCurrencyId
`func (o *OrderLineDto) UnsetTotalProfitCurrencyId()`

UnsetTotalProfitCurrencyId ensures that no value is present for TotalProfitCurrencyId, not even an explicit nil
### GetTotalProfitAmount

`func (o *OrderLineDto) GetTotalProfitAmount() Money`

GetTotalProfitAmount returns the TotalProfitAmount field if non-nil, zero value otherwise.

### GetTotalProfitAmountOk

`func (o *OrderLineDto) GetTotalProfitAmountOk() (*Money, bool)`

GetTotalProfitAmountOk returns a tuple with the TotalProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProfitAmount

`func (o *OrderLineDto) SetTotalProfitAmount(v Money)`

SetTotalProfitAmount sets TotalProfitAmount field to given value.

### HasTotalProfitAmount

`func (o *OrderLineDto) HasTotalProfitAmount() bool`

HasTotalProfitAmount returns a boolean if a field has been set.

### GetTotalDiscounts

`func (o *OrderLineDto) GetTotalDiscounts() float64`

GetTotalDiscounts returns the TotalDiscounts field if non-nil, zero value otherwise.

### GetTotalDiscountsOk

`func (o *OrderLineDto) GetTotalDiscountsOk() (*float64, bool)`

GetTotalDiscountsOk returns a tuple with the TotalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscounts

`func (o *OrderLineDto) SetTotalDiscounts(v float64)`

SetTotalDiscounts sets TotalDiscounts field to given value.

### HasTotalDiscounts

`func (o *OrderLineDto) HasTotalDiscounts() bool`

HasTotalDiscounts returns a boolean if a field has been set.

### GetTotalDiscountsCurrencyId

`func (o *OrderLineDto) GetTotalDiscountsCurrencyId() string`

GetTotalDiscountsCurrencyId returns the TotalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalDiscountsCurrencyIdOk

`func (o *OrderLineDto) GetTotalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalDiscountsCurrencyIdOk returns a tuple with the TotalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsCurrencyId

`func (o *OrderLineDto) SetTotalDiscountsCurrencyId(v string)`

SetTotalDiscountsCurrencyId sets TotalDiscountsCurrencyId field to given value.

### HasTotalDiscountsCurrencyId

`func (o *OrderLineDto) HasTotalDiscountsCurrencyId() bool`

HasTotalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalDiscountsCurrencyIdNil

`func (o *OrderLineDto) SetTotalDiscountsCurrencyIdNil(b bool)`

 SetTotalDiscountsCurrencyIdNil sets the value for TotalDiscountsCurrencyId to be an explicit nil

### UnsetTotalDiscountsCurrencyId
`func (o *OrderLineDto) UnsetTotalDiscountsCurrencyId()`

UnsetTotalDiscountsCurrencyId ensures that no value is present for TotalDiscountsCurrencyId, not even an explicit nil
### GetTotalDiscountsAmount

`func (o *OrderLineDto) GetTotalDiscountsAmount() Money`

GetTotalDiscountsAmount returns the TotalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalDiscountsAmountOk

`func (o *OrderLineDto) GetTotalDiscountsAmountOk() (*Money, bool)`

GetTotalDiscountsAmountOk returns a tuple with the TotalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscountsAmount

`func (o *OrderLineDto) SetTotalDiscountsAmount(v Money)`

SetTotalDiscountsAmount sets TotalDiscountsAmount field to given value.

### HasTotalDiscountsAmount

`func (o *OrderLineDto) HasTotalDiscountsAmount() bool`

HasTotalDiscountsAmount returns a boolean if a field has been set.

### GetTotalSurcharges

`func (o *OrderLineDto) GetTotalSurcharges() float64`

GetTotalSurcharges returns the TotalSurcharges field if non-nil, zero value otherwise.

### GetTotalSurchargesOk

`func (o *OrderLineDto) GetTotalSurchargesOk() (*float64, bool)`

GetTotalSurchargesOk returns a tuple with the TotalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurcharges

`func (o *OrderLineDto) SetTotalSurcharges(v float64)`

SetTotalSurcharges sets TotalSurcharges field to given value.

### HasTotalSurcharges

`func (o *OrderLineDto) HasTotalSurcharges() bool`

HasTotalSurcharges returns a boolean if a field has been set.

### GetTotalSurchargesCurrencyId

`func (o *OrderLineDto) GetTotalSurchargesCurrencyId() string`

GetTotalSurchargesCurrencyId returns the TotalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalSurchargesCurrencyIdOk

`func (o *OrderLineDto) GetTotalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalSurchargesCurrencyIdOk returns a tuple with the TotalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesCurrencyId

`func (o *OrderLineDto) SetTotalSurchargesCurrencyId(v string)`

SetTotalSurchargesCurrencyId sets TotalSurchargesCurrencyId field to given value.

### HasTotalSurchargesCurrencyId

`func (o *OrderLineDto) HasTotalSurchargesCurrencyId() bool`

HasTotalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalSurchargesCurrencyIdNil

`func (o *OrderLineDto) SetTotalSurchargesCurrencyIdNil(b bool)`

 SetTotalSurchargesCurrencyIdNil sets the value for TotalSurchargesCurrencyId to be an explicit nil

### UnsetTotalSurchargesCurrencyId
`func (o *OrderLineDto) UnsetTotalSurchargesCurrencyId()`

UnsetTotalSurchargesCurrencyId ensures that no value is present for TotalSurchargesCurrencyId, not even an explicit nil
### GetTotalSurchargesAmount

`func (o *OrderLineDto) GetTotalSurchargesAmount() Money`

GetTotalSurchargesAmount returns the TotalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalSurchargesAmountOk

`func (o *OrderLineDto) GetTotalSurchargesAmountOk() (*Money, bool)`

GetTotalSurchargesAmountOk returns a tuple with the TotalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSurchargesAmount

`func (o *OrderLineDto) SetTotalSurchargesAmount(v Money)`

SetTotalSurchargesAmount sets TotalSurchargesAmount field to given value.

### HasTotalSurchargesAmount

`func (o *OrderLineDto) HasTotalSurchargesAmount() bool`

HasTotalSurchargesAmount returns a boolean if a field has been set.

### GetTotalTaxBase

`func (o *OrderLineDto) GetTotalTaxBase() float64`

GetTotalTaxBase returns the TotalTaxBase field if non-nil, zero value otherwise.

### GetTotalTaxBaseOk

`func (o *OrderLineDto) GetTotalTaxBaseOk() (*float64, bool)`

GetTotalTaxBaseOk returns a tuple with the TotalTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBase

`func (o *OrderLineDto) SetTotalTaxBase(v float64)`

SetTotalTaxBase sets TotalTaxBase field to given value.

### HasTotalTaxBase

`func (o *OrderLineDto) HasTotalTaxBase() bool`

HasTotalTaxBase returns a boolean if a field has been set.

### GetTotalTaxBaseCurrencyId

`func (o *OrderLineDto) GetTotalTaxBaseCurrencyId() string`

GetTotalTaxBaseCurrencyId returns the TotalTaxBaseCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxBaseCurrencyIdOk

`func (o *OrderLineDto) GetTotalTaxBaseCurrencyIdOk() (*string, bool)`

GetTotalTaxBaseCurrencyIdOk returns a tuple with the TotalTaxBaseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseCurrencyId

`func (o *OrderLineDto) SetTotalTaxBaseCurrencyId(v string)`

SetTotalTaxBaseCurrencyId sets TotalTaxBaseCurrencyId field to given value.

### HasTotalTaxBaseCurrencyId

`func (o *OrderLineDto) HasTotalTaxBaseCurrencyId() bool`

HasTotalTaxBaseCurrencyId returns a boolean if a field has been set.

### SetTotalTaxBaseCurrencyIdNil

`func (o *OrderLineDto) SetTotalTaxBaseCurrencyIdNil(b bool)`

 SetTotalTaxBaseCurrencyIdNil sets the value for TotalTaxBaseCurrencyId to be an explicit nil

### UnsetTotalTaxBaseCurrencyId
`func (o *OrderLineDto) UnsetTotalTaxBaseCurrencyId()`

UnsetTotalTaxBaseCurrencyId ensures that no value is present for TotalTaxBaseCurrencyId, not even an explicit nil
### GetTotalTaxBaseAmount

`func (o *OrderLineDto) GetTotalTaxBaseAmount() Money`

GetTotalTaxBaseAmount returns the TotalTaxBaseAmount field if non-nil, zero value otherwise.

### GetTotalTaxBaseAmountOk

`func (o *OrderLineDto) GetTotalTaxBaseAmountOk() (*Money, bool)`

GetTotalTaxBaseAmountOk returns a tuple with the TotalTaxBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxBaseAmount

`func (o *OrderLineDto) SetTotalTaxBaseAmount(v Money)`

SetTotalTaxBaseAmount sets TotalTaxBaseAmount field to given value.

### HasTotalTaxBaseAmount

`func (o *OrderLineDto) HasTotalTaxBaseAmount() bool`

HasTotalTaxBaseAmount returns a boolean if a field has been set.

### GetTotalTaxes

`func (o *OrderLineDto) GetTotalTaxes() float64`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *OrderLineDto) GetTotalTaxesOk() (*float64, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *OrderLineDto) SetTotalTaxes(v float64)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *OrderLineDto) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetTotalTaxesCurrencyId

`func (o *OrderLineDto) GetTotalTaxesCurrencyId() string`

GetTotalTaxesCurrencyId returns the TotalTaxesCurrencyId field if non-nil, zero value otherwise.

### GetTotalTaxesCurrencyIdOk

`func (o *OrderLineDto) GetTotalTaxesCurrencyIdOk() (*string, bool)`

GetTotalTaxesCurrencyIdOk returns a tuple with the TotalTaxesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesCurrencyId

`func (o *OrderLineDto) SetTotalTaxesCurrencyId(v string)`

SetTotalTaxesCurrencyId sets TotalTaxesCurrencyId field to given value.

### HasTotalTaxesCurrencyId

`func (o *OrderLineDto) HasTotalTaxesCurrencyId() bool`

HasTotalTaxesCurrencyId returns a boolean if a field has been set.

### SetTotalTaxesCurrencyIdNil

`func (o *OrderLineDto) SetTotalTaxesCurrencyIdNil(b bool)`

 SetTotalTaxesCurrencyIdNil sets the value for TotalTaxesCurrencyId to be an explicit nil

### UnsetTotalTaxesCurrencyId
`func (o *OrderLineDto) UnsetTotalTaxesCurrencyId()`

UnsetTotalTaxesCurrencyId ensures that no value is present for TotalTaxesCurrencyId, not even an explicit nil
### GetTotalTaxesAmount

`func (o *OrderLineDto) GetTotalTaxesAmount() Money`

GetTotalTaxesAmount returns the TotalTaxesAmount field if non-nil, zero value otherwise.

### GetTotalTaxesAmountOk

`func (o *OrderLineDto) GetTotalTaxesAmountOk() (*Money, bool)`

GetTotalTaxesAmountOk returns a tuple with the TotalTaxesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxesAmount

`func (o *OrderLineDto) SetTotalTaxesAmount(v Money)`

SetTotalTaxesAmount sets TotalTaxesAmount field to given value.

### HasTotalTaxesAmount

`func (o *OrderLineDto) HasTotalTaxesAmount() bool`

HasTotalTaxesAmount returns a boolean if a field has been set.

### GetTotalShippingCost

`func (o *OrderLineDto) GetTotalShippingCost() float64`

GetTotalShippingCost returns the TotalShippingCost field if non-nil, zero value otherwise.

### GetTotalShippingCostOk

`func (o *OrderLineDto) GetTotalShippingCostOk() (*float64, bool)`

GetTotalShippingCostOk returns a tuple with the TotalShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCost

`func (o *OrderLineDto) SetTotalShippingCost(v float64)`

SetTotalShippingCost sets TotalShippingCost field to given value.

### HasTotalShippingCost

`func (o *OrderLineDto) HasTotalShippingCost() bool`

HasTotalShippingCost returns a boolean if a field has been set.

### GetTotalShippingCostCurrencyId

`func (o *OrderLineDto) GetTotalShippingCostCurrencyId() string`

GetTotalShippingCostCurrencyId returns the TotalShippingCostCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingCostCurrencyIdOk

`func (o *OrderLineDto) GetTotalShippingCostCurrencyIdOk() (*string, bool)`

GetTotalShippingCostCurrencyIdOk returns a tuple with the TotalShippingCostCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostCurrencyId

`func (o *OrderLineDto) SetTotalShippingCostCurrencyId(v string)`

SetTotalShippingCostCurrencyId sets TotalShippingCostCurrencyId field to given value.

### HasTotalShippingCostCurrencyId

`func (o *OrderLineDto) HasTotalShippingCostCurrencyId() bool`

HasTotalShippingCostCurrencyId returns a boolean if a field has been set.

### SetTotalShippingCostCurrencyIdNil

`func (o *OrderLineDto) SetTotalShippingCostCurrencyIdNil(b bool)`

 SetTotalShippingCostCurrencyIdNil sets the value for TotalShippingCostCurrencyId to be an explicit nil

### UnsetTotalShippingCostCurrencyId
`func (o *OrderLineDto) UnsetTotalShippingCostCurrencyId()`

UnsetTotalShippingCostCurrencyId ensures that no value is present for TotalShippingCostCurrencyId, not even an explicit nil
### GetTotalShippingCostAmount

`func (o *OrderLineDto) GetTotalShippingCostAmount() Money`

GetTotalShippingCostAmount returns the TotalShippingCostAmount field if non-nil, zero value otherwise.

### GetTotalShippingCostAmountOk

`func (o *OrderLineDto) GetTotalShippingCostAmountOk() (*Money, bool)`

GetTotalShippingCostAmountOk returns a tuple with the TotalShippingCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingCostAmount

`func (o *OrderLineDto) SetTotalShippingCostAmount(v Money)`

SetTotalShippingCostAmount sets TotalShippingCostAmount field to given value.

### HasTotalShippingCostAmount

`func (o *OrderLineDto) HasTotalShippingCostAmount() bool`

HasTotalShippingCostAmount returns a boolean if a field has been set.

### GetTotalShippingTax

`func (o *OrderLineDto) GetTotalShippingTax() float64`

GetTotalShippingTax returns the TotalShippingTax field if non-nil, zero value otherwise.

### GetTotalShippingTaxOk

`func (o *OrderLineDto) GetTotalShippingTaxOk() (*float64, bool)`

GetTotalShippingTaxOk returns a tuple with the TotalShippingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTax

`func (o *OrderLineDto) SetTotalShippingTax(v float64)`

SetTotalShippingTax sets TotalShippingTax field to given value.

### HasTotalShippingTax

`func (o *OrderLineDto) HasTotalShippingTax() bool`

HasTotalShippingTax returns a boolean if a field has been set.

### GetTotalShippingTaxCurrencyId

`func (o *OrderLineDto) GetTotalShippingTaxCurrencyId() string`

GetTotalShippingTaxCurrencyId returns the TotalShippingTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalShippingTaxCurrencyIdOk

`func (o *OrderLineDto) GetTotalShippingTaxCurrencyIdOk() (*string, bool)`

GetTotalShippingTaxCurrencyIdOk returns a tuple with the TotalShippingTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxCurrencyId

`func (o *OrderLineDto) SetTotalShippingTaxCurrencyId(v string)`

SetTotalShippingTaxCurrencyId sets TotalShippingTaxCurrencyId field to given value.

### HasTotalShippingTaxCurrencyId

`func (o *OrderLineDto) HasTotalShippingTaxCurrencyId() bool`

HasTotalShippingTaxCurrencyId returns a boolean if a field has been set.

### SetTotalShippingTaxCurrencyIdNil

`func (o *OrderLineDto) SetTotalShippingTaxCurrencyIdNil(b bool)`

 SetTotalShippingTaxCurrencyIdNil sets the value for TotalShippingTaxCurrencyId to be an explicit nil

### UnsetTotalShippingTaxCurrencyId
`func (o *OrderLineDto) UnsetTotalShippingTaxCurrencyId()`

UnsetTotalShippingTaxCurrencyId ensures that no value is present for TotalShippingTaxCurrencyId, not even an explicit nil
### GetTotalShippingTaxAmount

`func (o *OrderLineDto) GetTotalShippingTaxAmount() Money`

GetTotalShippingTaxAmount returns the TotalShippingTaxAmount field if non-nil, zero value otherwise.

### GetTotalShippingTaxAmountOk

`func (o *OrderLineDto) GetTotalShippingTaxAmountOk() (*Money, bool)`

GetTotalShippingTaxAmountOk returns a tuple with the TotalShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalShippingTaxAmount

`func (o *OrderLineDto) SetTotalShippingTaxAmount(v Money)`

SetTotalShippingTaxAmount sets TotalShippingTaxAmount field to given value.

### HasTotalShippingTaxAmount

`func (o *OrderLineDto) HasTotalShippingTaxAmount() bool`

HasTotalShippingTaxAmount returns a boolean if a field has been set.

### GetTotalWithheldTax

`func (o *OrderLineDto) GetTotalWithheldTax() float64`

GetTotalWithheldTax returns the TotalWithheldTax field if non-nil, zero value otherwise.

### GetTotalWithheldTaxOk

`func (o *OrderLineDto) GetTotalWithheldTaxOk() (*float64, bool)`

GetTotalWithheldTaxOk returns a tuple with the TotalWithheldTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTax

`func (o *OrderLineDto) SetTotalWithheldTax(v float64)`

SetTotalWithheldTax sets TotalWithheldTax field to given value.

### HasTotalWithheldTax

`func (o *OrderLineDto) HasTotalWithheldTax() bool`

HasTotalWithheldTax returns a boolean if a field has been set.

### GetTotalWithheldTaxCurrencyId

`func (o *OrderLineDto) GetTotalWithheldTaxCurrencyId() string`

GetTotalWithheldTaxCurrencyId returns the TotalWithheldTaxCurrencyId field if non-nil, zero value otherwise.

### GetTotalWithheldTaxCurrencyIdOk

`func (o *OrderLineDto) GetTotalWithheldTaxCurrencyIdOk() (*string, bool)`

GetTotalWithheldTaxCurrencyIdOk returns a tuple with the TotalWithheldTaxCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxCurrencyId

`func (o *OrderLineDto) SetTotalWithheldTaxCurrencyId(v string)`

SetTotalWithheldTaxCurrencyId sets TotalWithheldTaxCurrencyId field to given value.

### HasTotalWithheldTaxCurrencyId

`func (o *OrderLineDto) HasTotalWithheldTaxCurrencyId() bool`

HasTotalWithheldTaxCurrencyId returns a boolean if a field has been set.

### SetTotalWithheldTaxCurrencyIdNil

`func (o *OrderLineDto) SetTotalWithheldTaxCurrencyIdNil(b bool)`

 SetTotalWithheldTaxCurrencyIdNil sets the value for TotalWithheldTaxCurrencyId to be an explicit nil

### UnsetTotalWithheldTaxCurrencyId
`func (o *OrderLineDto) UnsetTotalWithheldTaxCurrencyId()`

UnsetTotalWithheldTaxCurrencyId ensures that no value is present for TotalWithheldTaxCurrencyId, not even an explicit nil
### GetTotalWithheldTaxAmount

`func (o *OrderLineDto) GetTotalWithheldTaxAmount() Money`

GetTotalWithheldTaxAmount returns the TotalWithheldTaxAmount field if non-nil, zero value otherwise.

### GetTotalWithheldTaxAmountOk

`func (o *OrderLineDto) GetTotalWithheldTaxAmountOk() (*Money, bool)`

GetTotalWithheldTaxAmountOk returns a tuple with the TotalWithheldTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithheldTaxAmount

`func (o *OrderLineDto) SetTotalWithheldTaxAmount(v Money)`

SetTotalWithheldTaxAmount sets TotalWithheldTaxAmount field to given value.

### HasTotalWithheldTaxAmount

`func (o *OrderLineDto) HasTotalWithheldTaxAmount() bool`

HasTotalWithheldTaxAmount returns a boolean if a field has been set.

### GetTotalGlobalDiscounts

`func (o *OrderLineDto) GetTotalGlobalDiscounts() float64`

GetTotalGlobalDiscounts returns the TotalGlobalDiscounts field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsOk

`func (o *OrderLineDto) GetTotalGlobalDiscountsOk() (*float64, bool)`

GetTotalGlobalDiscountsOk returns a tuple with the TotalGlobalDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscounts

`func (o *OrderLineDto) SetTotalGlobalDiscounts(v float64)`

SetTotalGlobalDiscounts sets TotalGlobalDiscounts field to given value.

### HasTotalGlobalDiscounts

`func (o *OrderLineDto) HasTotalGlobalDiscounts() bool`

HasTotalGlobalDiscounts returns a boolean if a field has been set.

### GetTotalGlobalDiscountsCurrencyId

`func (o *OrderLineDto) GetTotalGlobalDiscountsCurrencyId() string`

GetTotalGlobalDiscountsCurrencyId returns the TotalGlobalDiscountsCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsCurrencyIdOk

`func (o *OrderLineDto) GetTotalGlobalDiscountsCurrencyIdOk() (*string, bool)`

GetTotalGlobalDiscountsCurrencyIdOk returns a tuple with the TotalGlobalDiscountsCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsCurrencyId

`func (o *OrderLineDto) SetTotalGlobalDiscountsCurrencyId(v string)`

SetTotalGlobalDiscountsCurrencyId sets TotalGlobalDiscountsCurrencyId field to given value.

### HasTotalGlobalDiscountsCurrencyId

`func (o *OrderLineDto) HasTotalGlobalDiscountsCurrencyId() bool`

HasTotalGlobalDiscountsCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalDiscountsCurrencyIdNil

`func (o *OrderLineDto) SetTotalGlobalDiscountsCurrencyIdNil(b bool)`

 SetTotalGlobalDiscountsCurrencyIdNil sets the value for TotalGlobalDiscountsCurrencyId to be an explicit nil

### UnsetTotalGlobalDiscountsCurrencyId
`func (o *OrderLineDto) UnsetTotalGlobalDiscountsCurrencyId()`

UnsetTotalGlobalDiscountsCurrencyId ensures that no value is present for TotalGlobalDiscountsCurrencyId, not even an explicit nil
### GetTotalGlobalDiscountsAmount

`func (o *OrderLineDto) GetTotalGlobalDiscountsAmount() Money`

GetTotalGlobalDiscountsAmount returns the TotalGlobalDiscountsAmount field if non-nil, zero value otherwise.

### GetTotalGlobalDiscountsAmountOk

`func (o *OrderLineDto) GetTotalGlobalDiscountsAmountOk() (*Money, bool)`

GetTotalGlobalDiscountsAmountOk returns a tuple with the TotalGlobalDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalDiscountsAmount

`func (o *OrderLineDto) SetTotalGlobalDiscountsAmount(v Money)`

SetTotalGlobalDiscountsAmount sets TotalGlobalDiscountsAmount field to given value.

### HasTotalGlobalDiscountsAmount

`func (o *OrderLineDto) HasTotalGlobalDiscountsAmount() bool`

HasTotalGlobalDiscountsAmount returns a boolean if a field has been set.

### GetTotalGlobalSurcharges

`func (o *OrderLineDto) GetTotalGlobalSurcharges() float64`

GetTotalGlobalSurcharges returns the TotalGlobalSurcharges field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesOk

`func (o *OrderLineDto) GetTotalGlobalSurchargesOk() (*float64, bool)`

GetTotalGlobalSurchargesOk returns a tuple with the TotalGlobalSurcharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurcharges

`func (o *OrderLineDto) SetTotalGlobalSurcharges(v float64)`

SetTotalGlobalSurcharges sets TotalGlobalSurcharges field to given value.

### HasTotalGlobalSurcharges

`func (o *OrderLineDto) HasTotalGlobalSurcharges() bool`

HasTotalGlobalSurcharges returns a boolean if a field has been set.

### GetTotalGlobalSurchargesCurrencyId

`func (o *OrderLineDto) GetTotalGlobalSurchargesCurrencyId() string`

GetTotalGlobalSurchargesCurrencyId returns the TotalGlobalSurchargesCurrencyId field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesCurrencyIdOk

`func (o *OrderLineDto) GetTotalGlobalSurchargesCurrencyIdOk() (*string, bool)`

GetTotalGlobalSurchargesCurrencyIdOk returns a tuple with the TotalGlobalSurchargesCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesCurrencyId

`func (o *OrderLineDto) SetTotalGlobalSurchargesCurrencyId(v string)`

SetTotalGlobalSurchargesCurrencyId sets TotalGlobalSurchargesCurrencyId field to given value.

### HasTotalGlobalSurchargesCurrencyId

`func (o *OrderLineDto) HasTotalGlobalSurchargesCurrencyId() bool`

HasTotalGlobalSurchargesCurrencyId returns a boolean if a field has been set.

### SetTotalGlobalSurchargesCurrencyIdNil

`func (o *OrderLineDto) SetTotalGlobalSurchargesCurrencyIdNil(b bool)`

 SetTotalGlobalSurchargesCurrencyIdNil sets the value for TotalGlobalSurchargesCurrencyId to be an explicit nil

### UnsetTotalGlobalSurchargesCurrencyId
`func (o *OrderLineDto) UnsetTotalGlobalSurchargesCurrencyId()`

UnsetTotalGlobalSurchargesCurrencyId ensures that no value is present for TotalGlobalSurchargesCurrencyId, not even an explicit nil
### GetTotalGlobalSurchargesAmount

`func (o *OrderLineDto) GetTotalGlobalSurchargesAmount() Money`

GetTotalGlobalSurchargesAmount returns the TotalGlobalSurchargesAmount field if non-nil, zero value otherwise.

### GetTotalGlobalSurchargesAmountOk

`func (o *OrderLineDto) GetTotalGlobalSurchargesAmountOk() (*Money, bool)`

GetTotalGlobalSurchargesAmountOk returns a tuple with the TotalGlobalSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGlobalSurchargesAmount

`func (o *OrderLineDto) SetTotalGlobalSurchargesAmount(v Money)`

SetTotalGlobalSurchargesAmount sets TotalGlobalSurchargesAmount field to given value.

### HasTotalGlobalSurchargesAmount

`func (o *OrderLineDto) HasTotalGlobalSurchargesAmount() bool`

HasTotalGlobalSurchargesAmount returns a boolean if a field has been set.

### GetTotal

`func (o *OrderLineDto) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderLineDto) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderLineDto) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OrderLineDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalCurrencyId

`func (o *OrderLineDto) GetTotalCurrencyId() string`

GetTotalCurrencyId returns the TotalCurrencyId field if non-nil, zero value otherwise.

### GetTotalCurrencyIdOk

`func (o *OrderLineDto) GetTotalCurrencyIdOk() (*string, bool)`

GetTotalCurrencyIdOk returns a tuple with the TotalCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrencyId

`func (o *OrderLineDto) SetTotalCurrencyId(v string)`

SetTotalCurrencyId sets TotalCurrencyId field to given value.

### HasTotalCurrencyId

`func (o *OrderLineDto) HasTotalCurrencyId() bool`

HasTotalCurrencyId returns a boolean if a field has been set.

### SetTotalCurrencyIdNil

`func (o *OrderLineDto) SetTotalCurrencyIdNil(b bool)`

 SetTotalCurrencyIdNil sets the value for TotalCurrencyId to be an explicit nil

### UnsetTotalCurrencyId
`func (o *OrderLineDto) UnsetTotalCurrencyId()`

UnsetTotalCurrencyId ensures that no value is present for TotalCurrencyId, not even an explicit nil
### GetTotalAmount

`func (o *OrderLineDto) GetTotalAmount() Money`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *OrderLineDto) GetTotalAmountOk() (*Money, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *OrderLineDto) SetTotalAmount(v Money)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *OrderLineDto) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderLineDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderLineDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderLineDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderLineDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *OrderLineDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *OrderLineDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


