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

`func (o *OrderLineDto) GetTaxCalculationMethod() int32`

GetTaxCalculationMethod returns the TaxCalculationMethod field if non-nil, zero value otherwise.

### GetTaxCalculationMethodOk

`func (o *OrderLineDto) GetTaxCalculationMethodOk() (*int32, bool)`

GetTaxCalculationMethodOk returns a tuple with the TaxCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculationMethod

`func (o *OrderLineDto) SetTaxCalculationMethod(v int32)`

SetTaxCalculationMethod sets TaxCalculationMethod field to given value.

### HasTaxCalculationMethod

`func (o *OrderLineDto) HasTaxCalculationMethod() bool`

HasTaxCalculationMethod returns a boolean if a field has been set.

### GetCostCalculationMethod

`func (o *OrderLineDto) GetCostCalculationMethod() int32`

GetCostCalculationMethod returns the CostCalculationMethod field if non-nil, zero value otherwise.

### GetCostCalculationMethodOk

`func (o *OrderLineDto) GetCostCalculationMethodOk() (*int32, bool)`

GetCostCalculationMethodOk returns a tuple with the CostCalculationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCalculationMethod

`func (o *OrderLineDto) SetCostCalculationMethod(v int32)`

SetCostCalculationMethod sets CostCalculationMethod field to given value.

### HasCostCalculationMethod

`func (o *OrderLineDto) HasCostCalculationMethod() bool`

HasCostCalculationMethod returns a boolean if a field has been set.

### GetForexRatesSnapshot

`func (o *OrderLineDto) GetForexRatesSnapshot() string`

GetForexRatesSnapshot returns the ForexRatesSnapshot field if non-nil, zero value otherwise.

### GetForexRatesSnapshotOk

`func (o *OrderLineDto) GetForexRatesSnapshotOk() (*string, bool)`

GetForexRatesSnapshotOk returns a tuple with the ForexRatesSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRatesSnapshot

`func (o *OrderLineDto) SetForexRatesSnapshot(v string)`

SetForexRatesSnapshot sets ForexRatesSnapshot field to given value.

### HasForexRatesSnapshot

`func (o *OrderLineDto) HasForexRatesSnapshot() bool`

HasForexRatesSnapshot returns a boolean if a field has been set.

### SetForexRatesSnapshotNil

`func (o *OrderLineDto) SetForexRatesSnapshotNil(b bool)`

 SetForexRatesSnapshotNil sets the value for ForexRatesSnapshot to be an explicit nil

### UnsetForexRatesSnapshot
`func (o *OrderLineDto) UnsetForexRatesSnapshot()`

UnsetForexRatesSnapshot ensures that no value is present for ForexRatesSnapshot, not even an explicit nil
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

### GetTotalBaseAmountInUsd

`func (o *OrderLineDto) GetTotalBaseAmountInUsd() float64`

GetTotalBaseAmountInUsd returns the TotalBaseAmountInUsd field if non-nil, zero value otherwise.

### GetTotalBaseAmountInUsdOk

`func (o *OrderLineDto) GetTotalBaseAmountInUsdOk() (*float64, bool)`

GetTotalBaseAmountInUsdOk returns a tuple with the TotalBaseAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBaseAmountInUsd

`func (o *OrderLineDto) SetTotalBaseAmountInUsd(v float64)`

SetTotalBaseAmountInUsd sets TotalBaseAmountInUsd field to given value.

### HasTotalBaseAmountInUsd

`func (o *OrderLineDto) HasTotalBaseAmountInUsd() bool`

HasTotalBaseAmountInUsd returns a boolean if a field has been set.

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

### GetTotalDetailAmountInUsd

`func (o *OrderLineDto) GetTotalDetailAmountInUsd() float64`

GetTotalDetailAmountInUsd returns the TotalDetailAmountInUsd field if non-nil, zero value otherwise.

### GetTotalDetailAmountInUsdOk

`func (o *OrderLineDto) GetTotalDetailAmountInUsdOk() (*float64, bool)`

GetTotalDetailAmountInUsdOk returns a tuple with the TotalDetailAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDetailAmountInUsd

`func (o *OrderLineDto) SetTotalDetailAmountInUsd(v float64)`

SetTotalDetailAmountInUsd sets TotalDetailAmountInUsd field to given value.

### HasTotalDetailAmountInUsd

`func (o *OrderLineDto) HasTotalDetailAmountInUsd() bool`

HasTotalDetailAmountInUsd returns a boolean if a field has been set.

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

### GetTotalWithholdingTaxesInUsd

`func (o *OrderLineDto) GetTotalWithholdingTaxesInUsd() float64`

GetTotalWithholdingTaxesInUsd returns the TotalWithholdingTaxesInUsd field if non-nil, zero value otherwise.

### GetTotalWithholdingTaxesInUsdOk

`func (o *OrderLineDto) GetTotalWithholdingTaxesInUsdOk() (*float64, bool)`

GetTotalWithholdingTaxesInUsdOk returns a tuple with the TotalWithholdingTaxesInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithholdingTaxesInUsd

`func (o *OrderLineDto) SetTotalWithholdingTaxesInUsd(v float64)`

SetTotalWithholdingTaxesInUsd sets TotalWithholdingTaxesInUsd field to given value.

### HasTotalWithholdingTaxesInUsd

`func (o *OrderLineDto) HasTotalWithholdingTaxesInUsd() bool`

HasTotalWithholdingTaxesInUsd returns a boolean if a field has been set.

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

### GetTotalAmountInUsd

`func (o *OrderLineDto) GetTotalAmountInUsd() float64`

GetTotalAmountInUsd returns the TotalAmountInUsd field if non-nil, zero value otherwise.

### GetTotalAmountInUsdOk

`func (o *OrderLineDto) GetTotalAmountInUsdOk() (*float64, bool)`

GetTotalAmountInUsdOk returns a tuple with the TotalAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInUsd

`func (o *OrderLineDto) SetTotalAmountInUsd(v float64)`

SetTotalAmountInUsd sets TotalAmountInUsd field to given value.

### HasTotalAmountInUsd

`func (o *OrderLineDto) HasTotalAmountInUsd() bool`

HasTotalAmountInUsd returns a boolean if a field has been set.

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

### GetCustomBaseAmount

`func (o *OrderLineDto) GetCustomBaseAmount() float64`

GetCustomBaseAmount returns the CustomBaseAmount field if non-nil, zero value otherwise.

### GetCustomBaseAmountOk

`func (o *OrderLineDto) GetCustomBaseAmountOk() (*float64, bool)`

GetCustomBaseAmountOk returns a tuple with the CustomBaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBaseAmount

`func (o *OrderLineDto) SetCustomBaseAmount(v float64)`

SetCustomBaseAmount sets CustomBaseAmount field to given value.

### HasCustomBaseAmount

`func (o *OrderLineDto) HasCustomBaseAmount() bool`

HasCustomBaseAmount returns a boolean if a field has been set.

### GetCustomDetailAmount

`func (o *OrderLineDto) GetCustomDetailAmount() float64`

GetCustomDetailAmount returns the CustomDetailAmount field if non-nil, zero value otherwise.

### GetCustomDetailAmountOk

`func (o *OrderLineDto) GetCustomDetailAmountOk() (*float64, bool)`

GetCustomDetailAmountOk returns a tuple with the CustomDetailAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDetailAmount

`func (o *OrderLineDto) SetCustomDetailAmount(v float64)`

SetCustomDetailAmount sets CustomDetailAmount field to given value.

### HasCustomDetailAmount

`func (o *OrderLineDto) HasCustomDetailAmount() bool`

HasCustomDetailAmount returns a boolean if a field has been set.

### GetCustomDiscountsAmount

`func (o *OrderLineDto) GetCustomDiscountsAmount() float64`

GetCustomDiscountsAmount returns the CustomDiscountsAmount field if non-nil, zero value otherwise.

### GetCustomDiscountsAmountOk

`func (o *OrderLineDto) GetCustomDiscountsAmountOk() (*float64, bool)`

GetCustomDiscountsAmountOk returns a tuple with the CustomDiscountsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDiscountsAmount

`func (o *OrderLineDto) SetCustomDiscountsAmount(v float64)`

SetCustomDiscountsAmount sets CustomDiscountsAmount field to given value.

### HasCustomDiscountsAmount

`func (o *OrderLineDto) HasCustomDiscountsAmount() bool`

HasCustomDiscountsAmount returns a boolean if a field has been set.

### GetCustomTaxBase

`func (o *OrderLineDto) GetCustomTaxBase() float64`

GetCustomTaxBase returns the CustomTaxBase field if non-nil, zero value otherwise.

### GetCustomTaxBaseOk

`func (o *OrderLineDto) GetCustomTaxBaseOk() (*float64, bool)`

GetCustomTaxBaseOk returns a tuple with the CustomTaxBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaxBase

`func (o *OrderLineDto) SetCustomTaxBase(v float64)`

SetCustomTaxBase sets CustomTaxBase field to given value.

### HasCustomTaxBase

`func (o *OrderLineDto) HasCustomTaxBase() bool`

HasCustomTaxBase returns a boolean if a field has been set.

### GetCustomSurchargesAmount

`func (o *OrderLineDto) GetCustomSurchargesAmount() float64`

GetCustomSurchargesAmount returns the CustomSurchargesAmount field if non-nil, zero value otherwise.

### GetCustomSurchargesAmountOk

`func (o *OrderLineDto) GetCustomSurchargesAmountOk() (*float64, bool)`

GetCustomSurchargesAmountOk returns a tuple with the CustomSurchargesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargesAmount

`func (o *OrderLineDto) SetCustomSurchargesAmount(v float64)`

SetCustomSurchargesAmount sets CustomSurchargesAmount field to given value.

### HasCustomSurchargesAmount

`func (o *OrderLineDto) HasCustomSurchargesAmount() bool`

HasCustomSurchargesAmount returns a boolean if a field has been set.

### GetCustomProfitAmount

`func (o *OrderLineDto) GetCustomProfitAmount() float64`

GetCustomProfitAmount returns the CustomProfitAmount field if non-nil, zero value otherwise.

### GetCustomProfitAmountOk

`func (o *OrderLineDto) GetCustomProfitAmountOk() (*float64, bool)`

GetCustomProfitAmountOk returns a tuple with the CustomProfitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProfitAmount

`func (o *OrderLineDto) SetCustomProfitAmount(v float64)`

SetCustomProfitAmount sets CustomProfitAmount field to given value.

### HasCustomProfitAmount

`func (o *OrderLineDto) HasCustomProfitAmount() bool`

HasCustomProfitAmount returns a boolean if a field has been set.

### GetCustomShippingCostAmount

`func (o *OrderLineDto) GetCustomShippingCostAmount() float64`

GetCustomShippingCostAmount returns the CustomShippingCostAmount field if non-nil, zero value otherwise.

### GetCustomShippingCostAmountOk

`func (o *OrderLineDto) GetCustomShippingCostAmountOk() (*float64, bool)`

GetCustomShippingCostAmountOk returns a tuple with the CustomShippingCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShippingCostAmount

`func (o *OrderLineDto) SetCustomShippingCostAmount(v float64)`

SetCustomShippingCostAmount sets CustomShippingCostAmount field to given value.

### HasCustomShippingCostAmount

`func (o *OrderLineDto) HasCustomShippingCostAmount() bool`

HasCustomShippingCostAmount returns a boolean if a field has been set.

### GetCustomShippingTaxAmount

`func (o *OrderLineDto) GetCustomShippingTaxAmount() float64`

GetCustomShippingTaxAmount returns the CustomShippingTaxAmount field if non-nil, zero value otherwise.

### GetCustomShippingTaxAmountOk

`func (o *OrderLineDto) GetCustomShippingTaxAmountOk() (*float64, bool)`

GetCustomShippingTaxAmountOk returns a tuple with the CustomShippingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomShippingTaxAmount

`func (o *OrderLineDto) SetCustomShippingTaxAmount(v float64)`

SetCustomShippingTaxAmount sets CustomShippingTaxAmount field to given value.

### HasCustomShippingTaxAmount

`func (o *OrderLineDto) HasCustomShippingTaxAmount() bool`

HasCustomShippingTaxAmount returns a boolean if a field has been set.

### GetCustomTaxAmount

`func (o *OrderLineDto) GetCustomTaxAmount() float64`

GetCustomTaxAmount returns the CustomTaxAmount field if non-nil, zero value otherwise.

### GetCustomTaxAmountOk

`func (o *OrderLineDto) GetCustomTaxAmountOk() (*float64, bool)`

GetCustomTaxAmountOk returns a tuple with the CustomTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTaxAmount

`func (o *OrderLineDto) SetCustomTaxAmount(v float64)`

SetCustomTaxAmount sets CustomTaxAmount field to given value.

### HasCustomTaxAmount

`func (o *OrderLineDto) HasCustomTaxAmount() bool`

HasCustomTaxAmount returns a boolean if a field has been set.

### GetCustomWithholdingTaxAmount

`func (o *OrderLineDto) GetCustomWithholdingTaxAmount() float64`

GetCustomWithholdingTaxAmount returns the CustomWithholdingTaxAmount field if non-nil, zero value otherwise.

### GetCustomWithholdingTaxAmountOk

`func (o *OrderLineDto) GetCustomWithholdingTaxAmountOk() (*float64, bool)`

GetCustomWithholdingTaxAmountOk returns a tuple with the CustomWithholdingTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomWithholdingTaxAmount

`func (o *OrderLineDto) SetCustomWithholdingTaxAmount(v float64)`

SetCustomWithholdingTaxAmount sets CustomWithholdingTaxAmount field to given value.

### HasCustomWithholdingTaxAmount

`func (o *OrderLineDto) HasCustomWithholdingTaxAmount() bool`

HasCustomWithholdingTaxAmount returns a boolean if a field has been set.

### GetCustomTotalAmount

`func (o *OrderLineDto) GetCustomTotalAmount() float64`

GetCustomTotalAmount returns the CustomTotalAmount field if non-nil, zero value otherwise.

### GetCustomTotalAmountOk

`func (o *OrderLineDto) GetCustomTotalAmountOk() (*float64, bool)`

GetCustomTotalAmountOk returns a tuple with the CustomTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTotalAmount

`func (o *OrderLineDto) SetCustomTotalAmount(v float64)`

SetCustomTotalAmount sets CustomTotalAmount field to given value.

### HasCustomTotalAmount

`func (o *OrderLineDto) HasCustomTotalAmount() bool`

HasCustomTotalAmount returns a boolean if a field has been set.

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

