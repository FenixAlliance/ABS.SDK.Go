# BankGuaranteeUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Margin** | Pointer to **float64** |  | [optional] 
**Charges** | Pointer to **float64** |  | [optional] 
**BeneficiaryName** | Pointer to **NullableString** |  | [optional] 
**GuaranteeNumber** | Pointer to **NullableString** |  | [optional] 
**GuaranteeConditions** | Pointer to **NullableString** |  | [optional] 
**FixedDepositNumber** | Pointer to **float64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**ValidityInDays** | Pointer to **int32** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**BankGuaranteeType** | Pointer to **string** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**BankProfileId** | Pointer to **NullableString** |  | [optional] 
**BankAccountId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBankGuaranteeUpdateDto

`func NewBankGuaranteeUpdateDto() *BankGuaranteeUpdateDto`

NewBankGuaranteeUpdateDto instantiates a new BankGuaranteeUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankGuaranteeUpdateDtoWithDefaults

`func NewBankGuaranteeUpdateDtoWithDefaults() *BankGuaranteeUpdateDto`

NewBankGuaranteeUpdateDtoWithDefaults instantiates a new BankGuaranteeUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMargin

`func (o *BankGuaranteeUpdateDto) GetMargin() float64`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *BankGuaranteeUpdateDto) GetMarginOk() (*float64, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *BankGuaranteeUpdateDto) SetMargin(v float64)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *BankGuaranteeUpdateDto) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetCharges

`func (o *BankGuaranteeUpdateDto) GetCharges() float64`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *BankGuaranteeUpdateDto) GetChargesOk() (*float64, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *BankGuaranteeUpdateDto) SetCharges(v float64)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *BankGuaranteeUpdateDto) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetBeneficiaryName

`func (o *BankGuaranteeUpdateDto) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *BankGuaranteeUpdateDto) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *BankGuaranteeUpdateDto) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.

### HasBeneficiaryName

`func (o *BankGuaranteeUpdateDto) HasBeneficiaryName() bool`

HasBeneficiaryName returns a boolean if a field has been set.

### SetBeneficiaryNameNil

`func (o *BankGuaranteeUpdateDto) SetBeneficiaryNameNil(b bool)`

 SetBeneficiaryNameNil sets the value for BeneficiaryName to be an explicit nil

### UnsetBeneficiaryName
`func (o *BankGuaranteeUpdateDto) UnsetBeneficiaryName()`

UnsetBeneficiaryName ensures that no value is present for BeneficiaryName, not even an explicit nil
### GetGuaranteeNumber

`func (o *BankGuaranteeUpdateDto) GetGuaranteeNumber() string`

GetGuaranteeNumber returns the GuaranteeNumber field if non-nil, zero value otherwise.

### GetGuaranteeNumberOk

`func (o *BankGuaranteeUpdateDto) GetGuaranteeNumberOk() (*string, bool)`

GetGuaranteeNumberOk returns a tuple with the GuaranteeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeNumber

`func (o *BankGuaranteeUpdateDto) SetGuaranteeNumber(v string)`

SetGuaranteeNumber sets GuaranteeNumber field to given value.

### HasGuaranteeNumber

`func (o *BankGuaranteeUpdateDto) HasGuaranteeNumber() bool`

HasGuaranteeNumber returns a boolean if a field has been set.

### SetGuaranteeNumberNil

`func (o *BankGuaranteeUpdateDto) SetGuaranteeNumberNil(b bool)`

 SetGuaranteeNumberNil sets the value for GuaranteeNumber to be an explicit nil

### UnsetGuaranteeNumber
`func (o *BankGuaranteeUpdateDto) UnsetGuaranteeNumber()`

UnsetGuaranteeNumber ensures that no value is present for GuaranteeNumber, not even an explicit nil
### GetGuaranteeConditions

`func (o *BankGuaranteeUpdateDto) GetGuaranteeConditions() string`

GetGuaranteeConditions returns the GuaranteeConditions field if non-nil, zero value otherwise.

### GetGuaranteeConditionsOk

`func (o *BankGuaranteeUpdateDto) GetGuaranteeConditionsOk() (*string, bool)`

GetGuaranteeConditionsOk returns a tuple with the GuaranteeConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeConditions

`func (o *BankGuaranteeUpdateDto) SetGuaranteeConditions(v string)`

SetGuaranteeConditions sets GuaranteeConditions field to given value.

### HasGuaranteeConditions

`func (o *BankGuaranteeUpdateDto) HasGuaranteeConditions() bool`

HasGuaranteeConditions returns a boolean if a field has been set.

### SetGuaranteeConditionsNil

`func (o *BankGuaranteeUpdateDto) SetGuaranteeConditionsNil(b bool)`

 SetGuaranteeConditionsNil sets the value for GuaranteeConditions to be an explicit nil

### UnsetGuaranteeConditions
`func (o *BankGuaranteeUpdateDto) UnsetGuaranteeConditions()`

UnsetGuaranteeConditions ensures that no value is present for GuaranteeConditions, not even an explicit nil
### GetFixedDepositNumber

`func (o *BankGuaranteeUpdateDto) GetFixedDepositNumber() float64`

GetFixedDepositNumber returns the FixedDepositNumber field if non-nil, zero value otherwise.

### GetFixedDepositNumberOk

`func (o *BankGuaranteeUpdateDto) GetFixedDepositNumberOk() (*float64, bool)`

GetFixedDepositNumberOk returns a tuple with the FixedDepositNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedDepositNumber

`func (o *BankGuaranteeUpdateDto) SetFixedDepositNumber(v float64)`

SetFixedDepositNumber sets FixedDepositNumber field to given value.

### HasFixedDepositNumber

`func (o *BankGuaranteeUpdateDto) HasFixedDepositNumber() bool`

HasFixedDepositNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *BankGuaranteeUpdateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BankGuaranteeUpdateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BankGuaranteeUpdateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BankGuaranteeUpdateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BankGuaranteeUpdateDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BankGuaranteeUpdateDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BankGuaranteeUpdateDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BankGuaranteeUpdateDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetValidityInDays

`func (o *BankGuaranteeUpdateDto) GetValidityInDays() int32`

GetValidityInDays returns the ValidityInDays field if non-nil, zero value otherwise.

### GetValidityInDaysOk

`func (o *BankGuaranteeUpdateDto) GetValidityInDaysOk() (*int32, bool)`

GetValidityInDaysOk returns a tuple with the ValidityInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityInDays

`func (o *BankGuaranteeUpdateDto) SetValidityInDays(v int32)`

SetValidityInDays sets ValidityInDays field to given value.

### HasValidityInDays

`func (o *BankGuaranteeUpdateDto) HasValidityInDays() bool`

HasValidityInDays returns a boolean if a field has been set.

### GetTenantId

`func (o *BankGuaranteeUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BankGuaranteeUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BankGuaranteeUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BankGuaranteeUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BankGuaranteeUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BankGuaranteeUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetBankGuaranteeType

`func (o *BankGuaranteeUpdateDto) GetBankGuaranteeType() string`

GetBankGuaranteeType returns the BankGuaranteeType field if non-nil, zero value otherwise.

### GetBankGuaranteeTypeOk

`func (o *BankGuaranteeUpdateDto) GetBankGuaranteeTypeOk() (*string, bool)`

GetBankGuaranteeTypeOk returns a tuple with the BankGuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankGuaranteeType

`func (o *BankGuaranteeUpdateDto) SetBankGuaranteeType(v string)`

SetBankGuaranteeType sets BankGuaranteeType field to given value.

### HasBankGuaranteeType

`func (o *BankGuaranteeUpdateDto) HasBankGuaranteeType() bool`

HasBankGuaranteeType returns a boolean if a field has been set.

### GetEnrollmentId

`func (o *BankGuaranteeUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BankGuaranteeUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BankGuaranteeUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BankGuaranteeUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BankGuaranteeUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BankGuaranteeUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetContactId

`func (o *BankGuaranteeUpdateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *BankGuaranteeUpdateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *BankGuaranteeUpdateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *BankGuaranteeUpdateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *BankGuaranteeUpdateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *BankGuaranteeUpdateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetProjectId

`func (o *BankGuaranteeUpdateDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BankGuaranteeUpdateDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BankGuaranteeUpdateDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BankGuaranteeUpdateDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BankGuaranteeUpdateDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BankGuaranteeUpdateDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOrderId

`func (o *BankGuaranteeUpdateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BankGuaranteeUpdateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BankGuaranteeUpdateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *BankGuaranteeUpdateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *BankGuaranteeUpdateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *BankGuaranteeUpdateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetBankProfileId

`func (o *BankGuaranteeUpdateDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankGuaranteeUpdateDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankGuaranteeUpdateDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankGuaranteeUpdateDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankGuaranteeUpdateDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankGuaranteeUpdateDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil
### GetBankAccountId

`func (o *BankGuaranteeUpdateDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankGuaranteeUpdateDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankGuaranteeUpdateDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankGuaranteeUpdateDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *BankGuaranteeUpdateDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *BankGuaranteeUpdateDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetCurrencyId

`func (o *BankGuaranteeUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankGuaranteeUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankGuaranteeUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BankGuaranteeUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *BankGuaranteeUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *BankGuaranteeUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


