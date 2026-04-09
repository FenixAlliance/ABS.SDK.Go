# PaymentCommissionDto

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
**PaymentID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentCommissionDto

`func NewPaymentCommissionDto() *PaymentCommissionDto`

NewPaymentCommissionDto instantiates a new PaymentCommissionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCommissionDtoWithDefaults

`func NewPaymentCommissionDtoWithDefaults() *PaymentCommissionDto`

NewPaymentCommissionDtoWithDefaults instantiates a new PaymentCommissionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentCommissionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentCommissionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentCommissionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentCommissionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentCommissionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentCommissionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PaymentCommissionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentCommissionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentCommissionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentCommissionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PaymentCommissionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PaymentCommissionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *PaymentCommissionDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentCommissionDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentCommissionDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PaymentCommissionDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PaymentCommissionDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PaymentCommissionDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *PaymentCommissionDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentCommissionDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentCommissionDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentCommissionDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentCommissionDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentCommissionDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBaseAmount

`func (o *PaymentCommissionDto) GetBaseAmount() float64`

GetBaseAmount returns the BaseAmount field if non-nil, zero value otherwise.

### GetBaseAmountOk

`func (o *PaymentCommissionDto) GetBaseAmountOk() (*float64, bool)`

GetBaseAmountOk returns a tuple with the BaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAmount

`func (o *PaymentCommissionDto) SetBaseAmount(v float64)`

SetBaseAmount sets BaseAmount field to given value.

### HasBaseAmount

`func (o *PaymentCommissionDto) HasBaseAmount() bool`

HasBaseAmount returns a boolean if a field has been set.

### GetAddedPercent

`func (o *PaymentCommissionDto) GetAddedPercent() float64`

GetAddedPercent returns the AddedPercent field if non-nil, zero value otherwise.

### GetAddedPercentOk

`func (o *PaymentCommissionDto) GetAddedPercentOk() (*float64, bool)`

GetAddedPercentOk returns a tuple with the AddedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedPercent

`func (o *PaymentCommissionDto) SetAddedPercent(v float64)`

SetAddedPercent sets AddedPercent field to given value.

### HasAddedPercent

`func (o *PaymentCommissionDto) HasAddedPercent() bool`

HasAddedPercent returns a boolean if a field has been set.

### GetAddedAmount

`func (o *PaymentCommissionDto) GetAddedAmount() float64`

GetAddedAmount returns the AddedAmount field if non-nil, zero value otherwise.

### GetAddedAmountOk

`func (o *PaymentCommissionDto) GetAddedAmountOk() (*float64, bool)`

GetAddedAmountOk returns a tuple with the AddedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAmount

`func (o *PaymentCommissionDto) SetAddedAmount(v float64)`

SetAddedAmount sets AddedAmount field to given value.

### HasAddedAmount

`func (o *PaymentCommissionDto) HasAddedAmount() bool`

HasAddedAmount returns a boolean if a field has been set.

### GetTaxComission

`func (o *PaymentCommissionDto) GetTaxComission() float64`

GetTaxComission returns the TaxComission field if non-nil, zero value otherwise.

### GetTaxComissionOk

`func (o *PaymentCommissionDto) GetTaxComissionOk() (*float64, bool)`

GetTaxComissionOk returns a tuple with the TaxComission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxComission

`func (o *PaymentCommissionDto) SetTaxComission(v float64)`

SetTaxComission sets TaxComission field to given value.

### HasTaxComission

`func (o *PaymentCommissionDto) HasTaxComission() bool`

HasTaxComission returns a boolean if a field has been set.

### GetTenantId

`func (o *PaymentCommissionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentCommissionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentCommissionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PaymentCommissionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PaymentCommissionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PaymentCommissionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *PaymentCommissionDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *PaymentCommissionDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *PaymentCommissionDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *PaymentCommissionDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *PaymentCommissionDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *PaymentCommissionDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSalaryId

`func (o *PaymentCommissionDto) GetSalaryId() string`

GetSalaryId returns the SalaryId field if non-nil, zero value otherwise.

### GetSalaryIdOk

`func (o *PaymentCommissionDto) GetSalaryIdOk() (*string, bool)`

GetSalaryIdOk returns a tuple with the SalaryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalaryId

`func (o *PaymentCommissionDto) SetSalaryId(v string)`

SetSalaryId sets SalaryId field to given value.

### HasSalaryId

`func (o *PaymentCommissionDto) HasSalaryId() bool`

HasSalaryId returns a boolean if a field has been set.

### SetSalaryIdNil

`func (o *PaymentCommissionDto) SetSalaryIdNil(b bool)`

 SetSalaryIdNil sets the value for SalaryId to be an explicit nil

### UnsetSalaryId
`func (o *PaymentCommissionDto) UnsetSalaryId()`

UnsetSalaryId ensures that no value is present for SalaryId, not even an explicit nil
### GetEmisorWalletAccountId

`func (o *PaymentCommissionDto) GetEmisorWalletAccountId() string`

GetEmisorWalletAccountId returns the EmisorWalletAccountId field if non-nil, zero value otherwise.

### GetEmisorWalletAccountIdOk

`func (o *PaymentCommissionDto) GetEmisorWalletAccountIdOk() (*string, bool)`

GetEmisorWalletAccountIdOk returns a tuple with the EmisorWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorWalletAccountId

`func (o *PaymentCommissionDto) SetEmisorWalletAccountId(v string)`

SetEmisorWalletAccountId sets EmisorWalletAccountId field to given value.

### HasEmisorWalletAccountId

`func (o *PaymentCommissionDto) HasEmisorWalletAccountId() bool`

HasEmisorWalletAccountId returns a boolean if a field has been set.

### SetEmisorWalletAccountIdNil

`func (o *PaymentCommissionDto) SetEmisorWalletAccountIdNil(b bool)`

 SetEmisorWalletAccountIdNil sets the value for EmisorWalletAccountId to be an explicit nil

### UnsetEmisorWalletAccountId
`func (o *PaymentCommissionDto) UnsetEmisorWalletAccountId()`

UnsetEmisorWalletAccountId ensures that no value is present for EmisorWalletAccountId, not even an explicit nil
### GetReceiverWalletAccountId

`func (o *PaymentCommissionDto) GetReceiverWalletAccountId() string`

GetReceiverWalletAccountId returns the ReceiverWalletAccountId field if non-nil, zero value otherwise.

### GetReceiverWalletAccountIdOk

`func (o *PaymentCommissionDto) GetReceiverWalletAccountIdOk() (*string, bool)`

GetReceiverWalletAccountIdOk returns a tuple with the ReceiverWalletAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverWalletAccountId

`func (o *PaymentCommissionDto) SetReceiverWalletAccountId(v string)`

SetReceiverWalletAccountId sets ReceiverWalletAccountId field to given value.

### HasReceiverWalletAccountId

`func (o *PaymentCommissionDto) HasReceiverWalletAccountId() bool`

HasReceiverWalletAccountId returns a boolean if a field has been set.

### SetReceiverWalletAccountIdNil

`func (o *PaymentCommissionDto) SetReceiverWalletAccountIdNil(b bool)`

 SetReceiverWalletAccountIdNil sets the value for ReceiverWalletAccountId to be an explicit nil

### UnsetReceiverWalletAccountId
`func (o *PaymentCommissionDto) UnsetReceiverWalletAccountId()`

UnsetReceiverWalletAccountId ensures that no value is present for ReceiverWalletAccountId, not even an explicit nil
### GetEmisorContactId

`func (o *PaymentCommissionDto) GetEmisorContactId() string`

GetEmisorContactId returns the EmisorContactId field if non-nil, zero value otherwise.

### GetEmisorContactIdOk

`func (o *PaymentCommissionDto) GetEmisorContactIdOk() (*string, bool)`

GetEmisorContactIdOk returns a tuple with the EmisorContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisorContactId

`func (o *PaymentCommissionDto) SetEmisorContactId(v string)`

SetEmisorContactId sets EmisorContactId field to given value.

### HasEmisorContactId

`func (o *PaymentCommissionDto) HasEmisorContactId() bool`

HasEmisorContactId returns a boolean if a field has been set.

### SetEmisorContactIdNil

`func (o *PaymentCommissionDto) SetEmisorContactIdNil(b bool)`

 SetEmisorContactIdNil sets the value for EmisorContactId to be an explicit nil

### UnsetEmisorContactId
`func (o *PaymentCommissionDto) UnsetEmisorContactId()`

UnsetEmisorContactId ensures that no value is present for EmisorContactId, not even an explicit nil
### GetReceiverContactId

`func (o *PaymentCommissionDto) GetReceiverContactId() string`

GetReceiverContactId returns the ReceiverContactId field if non-nil, zero value otherwise.

### GetReceiverContactIdOk

`func (o *PaymentCommissionDto) GetReceiverContactIdOk() (*string, bool)`

GetReceiverContactIdOk returns a tuple with the ReceiverContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverContactId

`func (o *PaymentCommissionDto) SetReceiverContactId(v string)`

SetReceiverContactId sets ReceiverContactId field to given value.

### HasReceiverContactId

`func (o *PaymentCommissionDto) HasReceiverContactId() bool`

HasReceiverContactId returns a boolean if a field has been set.

### SetReceiverContactIdNil

`func (o *PaymentCommissionDto) SetReceiverContactIdNil(b bool)`

 SetReceiverContactIdNil sets the value for ReceiverContactId to be an explicit nil

### UnsetReceiverContactId
`func (o *PaymentCommissionDto) UnsetReceiverContactId()`

UnsetReceiverContactId ensures that no value is present for ReceiverContactId, not even an explicit nil
### GetPaymentID

`func (o *PaymentCommissionDto) GetPaymentID() string`

GetPaymentID returns the PaymentID field if non-nil, zero value otherwise.

### GetPaymentIDOk

`func (o *PaymentCommissionDto) GetPaymentIDOk() (*string, bool)`

GetPaymentIDOk returns a tuple with the PaymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentID

`func (o *PaymentCommissionDto) SetPaymentID(v string)`

SetPaymentID sets PaymentID field to given value.

### HasPaymentID

`func (o *PaymentCommissionDto) HasPaymentID() bool`

HasPaymentID returns a boolean if a field has been set.

### SetPaymentIDNil

`func (o *PaymentCommissionDto) SetPaymentIDNil(b bool)`

 SetPaymentIDNil sets the value for PaymentID to be an explicit nil

### UnsetPaymentID
`func (o *PaymentCommissionDto) UnsetPaymentID()`

UnsetPaymentID ensures that no value is present for PaymentID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


