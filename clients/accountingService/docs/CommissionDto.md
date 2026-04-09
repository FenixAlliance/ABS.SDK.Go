# CommissionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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

### NewCommissionDto

`func NewCommissionDto() *CommissionDto`

NewCommissionDto instantiates a new CommissionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionDtoWithDefaults

`func NewCommissionDtoWithDefaults() *CommissionDto`

NewCommissionDtoWithDefaults instantiates a new CommissionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommissionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommissionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommissionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommissionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CommissionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommissionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *CommissionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CommissionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CommissionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CommissionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *CommissionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *CommissionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *CommissionDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CommissionDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CommissionDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CommissionDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CommissionDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CommissionDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CommissionDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommissionDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommissionDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommissionDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CommissionDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CommissionDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBaseAmount

`func (o *CommissionDto) GetBaseAmount() float64`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *CommissionDto) GetBaseAmountOk() (*float64, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *CommissionDto) SetBaseAmount(v float64)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *CommissionDto) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetAddedPercent

`func (o *CommissionDto) GetAddedPercent() float64`

GetAddedPercent returns the AddedPercent field if non-nil, zero value otherwise.

### GetAddedPercentOk

`func (o *CommissionDto) GetAddedPercentOk() (*float64, bool)`

GetAddedPercentOk returns a tuple with the AddedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedPercent

`func (o *CommissionDto) SetAddedPercent(v float64)`

SetAddedPercent sets AddedPercent field to given value.

### HasAddedPercent

`func (o *CommissionDto) HasAddedPercent() bool`

HasAddedPercent returns a boolean if a field has been set.

### GetAddedAmount

`func (o *CommissionDto) GetAddedAmount() float64`

GetAddedAmount returns the AddedAmount field if non-nil, zero value otherwise.

### GetAddedAmountOk

`func (o *CommissionDto) GetAddedAmountOk() (*float64, bool)`

GetAddedAmountOk returns a tuple with the AddedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAmount

`func (o *CommissionDto) SetAddedAmount(v float64)`

SetAddedAmount sets AddedAmount field to given value.

### HasAddedAmount

`func (o *CommissionDto) HasAddedAmount() bool`

HasAddedAmount returns a boolean if a field has been set.

### GetTaxComission

`func (o *CommissionDto) GetTaxComission() float64`

GetTaxComission returns the TaxComission field if non-nil, zero value otherwise.

### GetTaxComissionOk

`func (o *CommissionDto) GetTaxComissionOk() (*float64, bool)`

GetTaxComissionOk returns a tuple with the TaxComission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxComission

`func (o *CommissionDto) SetTaxComission(v float64)`

SetTaxComission sets TaxComission field to given value.

### HasTaxComission

`func (o *CommissionDto) HasTaxComission() bool`

HasTaxComission returns a boolean if a field has been set.

### GetTenantId

`func (o *CommissionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *CommissionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *CommissionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *CommissionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *CommissionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *CommissionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *CommissionDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *CommissionDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *CommissionDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *CommissionDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *CommissionDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *CommissionDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSalaryId

`func (o *CommissionDto) GetSalaryId() string`

GetSalaryId returns the SalaryId field if non-nil, zero value otherwise.

### GetSalaryIdOk

`func (o *CommissionDto) GetSalaryIdOk() (*string, bool)`

GetSalaryIdOk returns a tuple with the SalaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryId

`func (o *CommissionDto) SetSalaryId(v string)`

SetSalaryId sets SalaryId field to given value.

### HasSalaryId

`func (o *CommissionDto) HasSalaryId() bool`

HasSalaryId returns a boolean if a field has been set.

### SetSalaryIdNil

`func (o *CommissionDto) SetSalaryIdNil(b bool)`

 SetSalaryIdNil sets the value for SalaryId to be an explicit nil

### UnsetSalaryId
`func (o *CommissionDto) UnsetSalaryId()`

UnsetSalaryId ensures that no value is present for SalaryId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *CommissionDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *CommissionDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *CommissionDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *CommissionDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *CommissionDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *CommissionDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *CommissionDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *CommissionDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *CommissionDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *CommissionDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *CommissionDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *CommissionDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil
### GetEmisorContactId

`func (o *CommissionDto) GetEmisorContactId() string`

GetEmisorContactId returns the EmisorContactId field if non-nil, zero value otherwise.

### GetEmisorContactIdOk

`func (o *CommissionDto) GetEmisorContactIdOk() (*string, bool)`

GetEmisorContactIdOk returns a tuple with the EmisorContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorContactId

`func (o *CommissionDto) SetEmisorContactId(v string)`

SetEmisorContactId sets EmisorContactId field to given value.

### HasEmisorContactId

`func (o *CommissionDto) HasEmisorContactId() bool`

HasEmisorContactId returns a boolean if a field has been set.

### SetEmisorContactIdNil

`func (o *CommissionDto) SetEmisorContactIdNil(b bool)`

 SetEmisorContactIdNil sets the value for EmisorContactId to be an explicit nil

### UnsetEmisorContactId
`func (o *CommissionDto) UnsetEmisorContactId()`

UnsetEmisorContactId ensures that no value is present for EmisorContactId, not even an explicit nil
### GetReceiverContactId

`func (o *CommissionDto) GetReceiverContactId() string`

GetReceiverContactId returns the ReceiverContactId field if non-nil, zero value otherwise.

### GetReceiverContactIdOk

`func (o *CommissionDto) GetReceiverContactIdOk() (*string, bool)`

GetReceiverContactIdOk returns a tuple with the ReceiverContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverContactId

`func (o *CommissionDto) SetReceiverContactId(v string)`

SetReceiverContactId sets ReceiverContactId field to given value.

### HasReceiverContactId

`func (o *CommissionDto) HasReceiverContactId() bool`

HasReceiverContactId returns a boolean if a field has been set.

### SetReceiverContactIdNil

`func (o *CommissionDto) SetReceiverContactIdNil(b bool)`

 SetReceiverContactIdNil sets the value for ReceiverContactId to be an explicit nil

### UnsetReceiverContactId
`func (o *CommissionDto) UnsetReceiverContactId()`

UnsetReceiverContactId ensures that no value is present for ReceiverContactId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


