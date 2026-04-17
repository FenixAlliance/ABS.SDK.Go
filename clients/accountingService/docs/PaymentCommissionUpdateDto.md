# PaymentCommissionUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BaseAmount** | Pointer to **float64** |  | [optional] 
**AddedPercent** | Pointer to **float64** |  | [optional] 
**AddedAmount** | Pointer to **float64** |  | [optional] 
**TaxComission** | Pointer to **float64** |  | [optional] 
**SalaryId** | Pointer to **NullableString** |  | [optional] 
**EmisorWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**ReceiverWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**EmisorContactId** | Pointer to **NullableString** |  | [optional] 
**ReceiverContactId** | Pointer to **NullableString** |  | [optional] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentCommissionUpdateDto

`func NewPaymentCommissionUpdateDto() *PaymentCommissionUpdateDto`

NewPaymentCommissionUpdateDto instantiates a new PaymentCommissionUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCommissionUpdateDtoWithDefaults

`func NewPaymentCommissionUpdateDtoWithDefaults() *PaymentCommissionUpdateDto`

NewPaymentCommissionUpdateDtoWithDefaults instantiates a new PaymentCommissionUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PaymentCommissionUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentCommissionUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentCommissionUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PaymentCommissionUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PaymentCommissionUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PaymentCommissionUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *PaymentCommissionUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentCommissionUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentCommissionUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentCommissionUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentCommissionUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentCommissionUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBaseAmount

`func (o *PaymentCommissionUpdateDto) GetBaseAmount() float64`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *PaymentCommissionUpdateDto) GetBaseAmountOk() (*float64, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *PaymentCommissionUpdateDto) SetBaseAmount(v float64)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *PaymentCommissionUpdateDto) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetAddedPercent

`func (o *PaymentCommissionUpdateDto) GetAddedPercent() float64`

GetAddedPercent returns the AddedPercent field if non-nil, zero value otherwise.

### GetAddedPercentOk

`func (o *PaymentCommissionUpdateDto) GetAddedPercentOk() (*float64, bool)`

GetAddedPercentOk returns a tuple with the AddedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedPercent

`func (o *PaymentCommissionUpdateDto) SetAddedPercent(v float64)`

SetAddedPercent sets AddedPercent field to given value.

### HasAddedPercent

`func (o *PaymentCommissionUpdateDto) HasAddedPercent() bool`

HasAddedPercent returns a boolean if a field has been set.

### GetAddedAmount

`func (o *PaymentCommissionUpdateDto) GetAddedAmount() float64`

GetAddedAmount returns the AddedAmount field if non-nil, zero value otherwise.

### GetAddedAmountOk

`func (o *PaymentCommissionUpdateDto) GetAddedAmountOk() (*float64, bool)`

GetAddedAmountOk returns a tuple with the AddedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAmount

`func (o *PaymentCommissionUpdateDto) SetAddedAmount(v float64)`

SetAddedAmount sets AddedAmount field to given value.

### HasAddedAmount

`func (o *PaymentCommissionUpdateDto) HasAddedAmount() bool`

HasAddedAmount returns a boolean if a field has been set.

### GetTaxComission

`func (o *PaymentCommissionUpdateDto) GetTaxComission() float64`

GetTaxComission returns the TaxComission field if non-nil, zero value otherwise.

### GetTaxComissionOk

`func (o *PaymentCommissionUpdateDto) GetTaxComissionOk() (*float64, bool)`

GetTaxComissionOk returns a tuple with the TaxComission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxComission

`func (o *PaymentCommissionUpdateDto) SetTaxComission(v float64)`

SetTaxComission sets TaxComission field to given value.

### HasTaxComission

`func (o *PaymentCommissionUpdateDto) HasTaxComission() bool`

HasTaxComission returns a boolean if a field has been set.

### GetSalaryId

`func (o *PaymentCommissionUpdateDto) GetSalaryId() string`

GetSalaryId returns the SalaryId field if non-nil, zero value otherwise.

### GetSalaryIdOk

`func (o *PaymentCommissionUpdateDto) GetSalaryIdOk() (*string, bool)`

GetSalaryIdOk returns a tuple with the SalaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryId

`func (o *PaymentCommissionUpdateDto) SetSalaryId(v string)`

SetSalaryId sets SalaryId field to given value.

### HasSalaryId

`func (o *PaymentCommissionUpdateDto) HasSalaryId() bool`

HasSalaryId returns a boolean if a field has been set.

### SetSalaryIdNil

`func (o *PaymentCommissionUpdateDto) SetSalaryIdNil(b bool)`

 SetSalaryIdNil sets the value for SalaryId to be an explicit nil

### UnsetSalaryId
`func (o *PaymentCommissionUpdateDto) UnsetSalaryId()`

UnsetSalaryId ensures that no value is present for SalaryId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *PaymentCommissionUpdateDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *PaymentCommissionUpdateDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *PaymentCommissionUpdateDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *PaymentCommissionUpdateDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *PaymentCommissionUpdateDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *PaymentCommissionUpdateDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *PaymentCommissionUpdateDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *PaymentCommissionUpdateDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *PaymentCommissionUpdateDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *PaymentCommissionUpdateDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *PaymentCommissionUpdateDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *PaymentCommissionUpdateDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil
### GetEmisorContactId

`func (o *PaymentCommissionUpdateDto) GetEmisorContactId() string`

GetEmisorContactId returns the EmisorContactId field if non-nil, zero value otherwise.

### GetEmisorContactIdOk

`func (o *PaymentCommissionUpdateDto) GetEmisorContactIdOk() (*string, bool)`

GetEmisorContactIdOk returns a tuple with the EmisorContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorContactId

`func (o *PaymentCommissionUpdateDto) SetEmisorContactId(v string)`

SetEmisorContactId sets EmisorContactId field to given value.

### HasEmisorContactId

`func (o *PaymentCommissionUpdateDto) HasEmisorContactId() bool`

HasEmisorContactId returns a boolean if a field has been set.

### SetEmisorContactIdNil

`func (o *PaymentCommissionUpdateDto) SetEmisorContactIdNil(b bool)`

 SetEmisorContactIdNil sets the value for EmisorContactId to be an explicit nil

### UnsetEmisorContactId
`func (o *PaymentCommissionUpdateDto) UnsetEmisorContactId()`

UnsetEmisorContactId ensures that no value is present for EmisorContactId, not even an explicit nil
### GetReceiverContactId

`func (o *PaymentCommissionUpdateDto) GetReceiverContactId() string`

GetReceiverContactId returns the ReceiverContactId field if non-nil, zero value otherwise.

### GetReceiverContactIdOk

`func (o *PaymentCommissionUpdateDto) GetReceiverContactIdOk() (*string, bool)`

GetReceiverContactIdOk returns a tuple with the ReceiverContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverContactId

`func (o *PaymentCommissionUpdateDto) SetReceiverContactId(v string)`

SetReceiverContactId sets ReceiverContactId field to given value.

### HasReceiverContactId

`func (o *PaymentCommissionUpdateDto) HasReceiverContactId() bool`

HasReceiverContactId returns a boolean if a field has been set.

### SetReceiverContactIdNil

`func (o *PaymentCommissionUpdateDto) SetReceiverContactIdNil(b bool)`

 SetReceiverContactIdNil sets the value for ReceiverContactId to be an explicit nil

### UnsetReceiverContactId
`func (o *PaymentCommissionUpdateDto) UnsetReceiverContactId()`

UnsetReceiverContactId ensures that no value is present for ReceiverContactId, not even an explicit nil
### GetPaymentId

`func (o *PaymentCommissionUpdateDto) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentCommissionUpdateDto) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentCommissionUpdateDto) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PaymentCommissionUpdateDto) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *PaymentCommissionUpdateDto) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *PaymentCommissionUpdateDto) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


