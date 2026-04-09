# CommissionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BaseAmount** | Pointer to **float64** |  | [optional] 
**AddedPercent** | Pointer to **float64** |  | [optional] 
**AddedAmount** | Pointer to **float64** |  | [optional] 
**TaxComission** | Pointer to **float64** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SalaryId** | Pointer to **NullableString** |  | [optional] 
**EmisorWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**ReceiverWalletAccountId** | Pointer to **NullableString** |  | [optional] 
**EmisorContactId** | Pointer to **NullableString** |  | [optional] 
**ReceiverContactId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCommissionCreateDto

`func NewCommissionCreateDto() *CommissionCreateDto`

NewCommissionCreateDto instantiates a new CommissionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionCreateDtoWithDefaults

`func NewCommissionCreateDtoWithDefaults() *CommissionCreateDto`

NewCommissionCreateDtoWithDefaults instantiates a new CommissionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommissionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommissionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommissionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommissionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *CommissionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CommissionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CommissionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CommissionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *CommissionCreateDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CommissionCreateDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CommissionCreateDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CommissionCreateDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CommissionCreateDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CommissionCreateDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CommissionCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommissionCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommissionCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommissionCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommissionCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommissionCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBaseAmount

`func (o *CommissionCreateDto) GetBaseAmount() float64`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *CommissionCreateDto) GetBaseAmountOk() (*float64, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *CommissionCreateDto) SetBaseAmount(v float64)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *CommissionCreateDto) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetAddedPercent

`func (o *CommissionCreateDto) GetAddedPercent() float64`

GetAddedPercent returns the AddedPercent field if non-nil, zero value otherwise.

### GetAddedPercentOk

`func (o *CommissionCreateDto) GetAddedPercentOk() (*float64, bool)`

GetAddedPercentOk returns a tuple with the AddedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedPercent

`func (o *CommissionCreateDto) SetAddedPercent(v float64)`

SetAddedPercent sets AddedPercent field to given value.

### HasAddedPercent

`func (o *CommissionCreateDto) HasAddedPercent() bool`

HasAddedPercent returns a boolean if a field has been set.

### GetAddedAmount

`func (o *CommissionCreateDto) GetAddedAmount() float64`

GetAddedAmount returns the AddedAmount field if non-nil, zero value otherwise.

### GetAddedAmountOk

`func (o *CommissionCreateDto) GetAddedAmountOk() (*float64, bool)`

GetAddedAmountOk returns a tuple with the AddedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAmount

`func (o *CommissionCreateDto) SetAddedAmount(v float64)`

SetAddedAmount sets AddedAmount field to given value.

### HasAddedAmount

`func (o *CommissionCreateDto) HasAddedAmount() bool`

HasAddedAmount returns a boolean if a field has been set.

### GetTaxComission

`func (o *CommissionCreateDto) GetTaxComission() float64`

GetTaxComission returns the TaxComission field if non-nil, zero value otherwise.

### GetTaxComissionOk

`func (o *CommissionCreateDto) GetTaxComissionOk() (*float64, bool)`

GetTaxComissionOk returns a tuple with the TaxComission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxComission

`func (o *CommissionCreateDto) SetTaxComission(v float64)`

SetTaxComission sets TaxComission field to given value.

### HasTaxComission

`func (o *CommissionCreateDto) HasTaxComission() bool`

HasTaxComission returns a boolean if a field has been set.

### GetTenantId

`func (o *CommissionCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CommissionCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CommissionCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CommissionCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CommissionCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CommissionCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CommissionCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CommissionCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CommissionCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CommissionCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CommissionCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CommissionCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSalaryId

`func (o *CommissionCreateDto) GetSalaryId() string`

GetSalaryId returns the SalaryId field if non-nil, zero value otherwise.

### GetSalaryIdOk

`func (o *CommissionCreateDto) GetSalaryIdOk() (*string, bool)`

GetSalaryIdOk returns a tuple with the SalaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryId

`func (o *CommissionCreateDto) SetSalaryId(v string)`

SetSalaryId sets SalaryId field to given value.

### HasSalaryId

`func (o *CommissionCreateDto) HasSalaryId() bool`

HasSalaryId returns a boolean if a field has been set.

### SetSalaryIdNil

`func (o *CommissionCreateDto) SetSalaryIdNil(b bool)`

 SetSalaryIdNil sets the value for SalaryId to be an explicit nil

### UnsetSalaryId
`func (o *CommissionCreateDto) UnsetSalaryId()`

UnsetSalaryId ensures that no value is present for SalaryId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *CommissionCreateDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *CommissionCreateDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *CommissionCreateDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *CommissionCreateDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *CommissionCreateDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *CommissionCreateDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *CommissionCreateDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *CommissionCreateDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *CommissionCreateDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *CommissionCreateDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *CommissionCreateDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *CommissionCreateDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil
### GetEmisorContactId

`func (o *CommissionCreateDto) GetEmisorContactId() string`

GetEmisorContactId returns the EmisorContactId field if non-nil, zero value otherwise.

### GetEmisorContactIdOk

`func (o *CommissionCreateDto) GetEmisorContactIdOk() (*string, bool)`

GetEmisorContactIdOk returns a tuple with the EmisorContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorContactId

`func (o *CommissionCreateDto) SetEmisorContactId(v string)`

SetEmisorContactId sets EmisorContactId field to given value.

### HasEmisorContactId

`func (o *CommissionCreateDto) HasEmisorContactId() bool`

HasEmisorContactId returns a boolean if a field has been set.

### SetEmisorContactIdNil

`func (o *CommissionCreateDto) SetEmisorContactIdNil(b bool)`

 SetEmisorContactIdNil sets the value for EmisorContactId to be an explicit nil

### UnsetEmisorContactId
`func (o *CommissionCreateDto) UnsetEmisorContactId()`

UnsetEmisorContactId ensures that no value is present for EmisorContactId, not even an explicit nil
### GetReceiverContactId

`func (o *CommissionCreateDto) GetReceiverContactId() string`

GetReceiverContactId returns the ReceiverContactId field if non-nil, zero value otherwise.

### GetReceiverContactIdOk

`func (o *CommissionCreateDto) GetReceiverContactIdOk() (*string, bool)`

GetReceiverContactIdOk returns a tuple with the ReceiverContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverContactId

`func (o *CommissionCreateDto) SetReceiverContactId(v string)`

SetReceiverContactId sets ReceiverContactId field to given value.

### HasReceiverContactId

`func (o *CommissionCreateDto) HasReceiverContactId() bool`

HasReceiverContactId returns a boolean if a field has been set.

### SetReceiverContactIdNil

`func (o *CommissionCreateDto) SetReceiverContactIdNil(b bool)`

 SetReceiverContactIdNil sets the value for ReceiverContactId to be an explicit nil

### UnsetReceiverContactId
`func (o *CommissionCreateDto) UnsetReceiverContactId()`

UnsetReceiverContactId ensures that no value is present for ReceiverContactId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


