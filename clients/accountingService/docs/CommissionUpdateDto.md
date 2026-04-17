# CommissionUpdateDto

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

## Methods

### NewCommissionUpdateDto

`func NewCommissionUpdateDto() *CommissionUpdateDto`

NewCommissionUpdateDto instantiates a new CommissionUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionUpdateDtoWithDefaults

`func NewCommissionUpdateDtoWithDefaults() *CommissionUpdateDto`

NewCommissionUpdateDtoWithDefaults instantiates a new CommissionUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CommissionUpdateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CommissionUpdateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CommissionUpdateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CommissionUpdateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CommissionUpdateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CommissionUpdateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CommissionUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommissionUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommissionUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommissionUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommissionUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommissionUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBaseAmount

`func (o *CommissionUpdateDto) GetBaseAmount() float64`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *CommissionUpdateDto) GetBaseAmountOk() (*float64, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *CommissionUpdateDto) SetBaseAmount(v float64)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *CommissionUpdateDto) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetAddedPercent

`func (o *CommissionUpdateDto) GetAddedPercent() float64`

GetAddedPercent returns the AddedPercent field if non-nil, zero value otherwise.

### GetAddedPercentOk

`func (o *CommissionUpdateDto) GetAddedPercentOk() (*float64, bool)`

GetAddedPercentOk returns a tuple with the AddedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedPercent

`func (o *CommissionUpdateDto) SetAddedPercent(v float64)`

SetAddedPercent sets AddedPercent field to given value.

### HasAddedPercent

`func (o *CommissionUpdateDto) HasAddedPercent() bool`

HasAddedPercent returns a boolean if a field has been set.

### GetAddedAmount

`func (o *CommissionUpdateDto) GetAddedAmount() float64`

GetAddedAmount returns the AddedAmount field if non-nil, zero value otherwise.

### GetAddedAmountOk

`func (o *CommissionUpdateDto) GetAddedAmountOk() (*float64, bool)`

GetAddedAmountOk returns a tuple with the AddedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAmount

`func (o *CommissionUpdateDto) SetAddedAmount(v float64)`

SetAddedAmount sets AddedAmount field to given value.

### HasAddedAmount

`func (o *CommissionUpdateDto) HasAddedAmount() bool`

HasAddedAmount returns a boolean if a field has been set.

### GetTaxComission

`func (o *CommissionUpdateDto) GetTaxComission() float64`

GetTaxComission returns the TaxComission field if non-nil, zero value otherwise.

### GetTaxComissionOk

`func (o *CommissionUpdateDto) GetTaxComissionOk() (*float64, bool)`

GetTaxComissionOk returns a tuple with the TaxComission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxComission

`func (o *CommissionUpdateDto) SetTaxComission(v float64)`

SetTaxComission sets TaxComission field to given value.

### HasTaxComission

`func (o *CommissionUpdateDto) HasTaxComission() bool`

HasTaxComission returns a boolean if a field has been set.

### GetSalaryId

`func (o *CommissionUpdateDto) GetSalaryId() string`

GetSalaryId returns the SalaryId field if non-nil, zero value otherwise.

### GetSalaryIdOk

`func (o *CommissionUpdateDto) GetSalaryIdOk() (*string, bool)`

GetSalaryIdOk returns a tuple with the SalaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryId

`func (o *CommissionUpdateDto) SetSalaryId(v string)`

SetSalaryId sets SalaryId field to given value.

### HasSalaryId

`func (o *CommissionUpdateDto) HasSalaryId() bool`

HasSalaryId returns a boolean if a field has been set.

### SetSalaryIdNil

`func (o *CommissionUpdateDto) SetSalaryIdNil(b bool)`

 SetSalaryIdNil sets the value for SalaryId to be an explicit nil

### UnsetSalaryId
`func (o *CommissionUpdateDto) UnsetSalaryId()`

UnsetSalaryId ensures that no value is present for SalaryId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *CommissionUpdateDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *CommissionUpdateDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *CommissionUpdateDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *CommissionUpdateDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *CommissionUpdateDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *CommissionUpdateDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *CommissionUpdateDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *CommissionUpdateDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *CommissionUpdateDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *CommissionUpdateDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *CommissionUpdateDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *CommissionUpdateDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil
### GetEmisorContactId

`func (o *CommissionUpdateDto) GetEmisorContactId() string`

GetEmisorContactId returns the EmisorContactId field if non-nil, zero value otherwise.

### GetEmisorContactIdOk

`func (o *CommissionUpdateDto) GetEmisorContactIdOk() (*string, bool)`

GetEmisorContactIdOk returns a tuple with the EmisorContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorContactId

`func (o *CommissionUpdateDto) SetEmisorContactId(v string)`

SetEmisorContactId sets EmisorContactId field to given value.

### HasEmisorContactId

`func (o *CommissionUpdateDto) HasEmisorContactId() bool`

HasEmisorContactId returns a boolean if a field has been set.

### SetEmisorContactIdNil

`func (o *CommissionUpdateDto) SetEmisorContactIdNil(b bool)`

 SetEmisorContactIdNil sets the value for EmisorContactId to be an explicit nil

### UnsetEmisorContactId
`func (o *CommissionUpdateDto) UnsetEmisorContactId()`

UnsetEmisorContactId ensures that no value is present for EmisorContactId, not even an explicit nil
### GetReceiverContactId

`func (o *CommissionUpdateDto) GetReceiverContactId() string`

GetReceiverContactId returns the ReceiverContactId field if non-nil, zero value otherwise.

### GetReceiverContactIdOk

`func (o *CommissionUpdateDto) GetReceiverContactIdOk() (*string, bool)`

GetReceiverContactIdOk returns a tuple with the ReceiverContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverContactId

`func (o *CommissionUpdateDto) SetReceiverContactId(v string)`

SetReceiverContactId sets ReceiverContactId field to given value.

### HasReceiverContactId

`func (o *CommissionUpdateDto) HasReceiverContactId() bool`

HasReceiverContactId returns a boolean if a field has been set.

### SetReceiverContactIdNil

`func (o *CommissionUpdateDto) SetReceiverContactIdNil(b bool)`

 SetReceiverContactIdNil sets the value for ReceiverContactId to be an explicit nil

### UnsetReceiverContactId
`func (o *CommissionUpdateDto) UnsetReceiverContactId()`

UnsetReceiverContactId ensures that no value is present for ReceiverContactId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


