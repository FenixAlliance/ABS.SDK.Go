# BankGuaranteeCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
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

### NewBankGuaranteeCreateDto

`func NewBankGuaranteeCreateDto() *BankGuaranteeCreateDto`

NewBankGuaranteeCreateDto instantiates a new BankGuaranteeCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankGuaranteeCreateDtoWithDefaults

`func NewBankGuaranteeCreateDtoWithDefaults() *BankGuaranteeCreateDto`

NewBankGuaranteeCreateDtoWithDefaults instantiates a new BankGuaranteeCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankGuaranteeCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankGuaranteeCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankGuaranteeCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankGuaranteeCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *BankGuaranteeCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BankGuaranteeCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BankGuaranteeCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BankGuaranteeCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMargin

`func (o *BankGuaranteeCreateDto) GetMargin() float64`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *BankGuaranteeCreateDto) GetMarginOk() (*float64, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *BankGuaranteeCreateDto) SetMargin(v float64)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *BankGuaranteeCreateDto) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetCharges

`func (o *BankGuaranteeCreateDto) GetCharges() float64`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *BankGuaranteeCreateDto) GetChargesOk() (*float64, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *BankGuaranteeCreateDto) SetCharges(v float64)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *BankGuaranteeCreateDto) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetBeneficiaryName

`func (o *BankGuaranteeCreateDto) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *BankGuaranteeCreateDto) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *BankGuaranteeCreateDto) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.

### HasBeneficiaryName

`func (o *BankGuaranteeCreateDto) HasBeneficiaryName() bool`

HasBeneficiaryName returns a boolean if a field has been set.

### SetBeneficiaryNameNil

`func (o *BankGuaranteeCreateDto) SetBeneficiaryNameNil(b bool)`

 SetBeneficiaryNameNil sets the value for BeneficiaryName to be an explicit nil

### UnsetBeneficiaryName
`func (o *BankGuaranteeCreateDto) UnsetBeneficiaryName()`

UnsetBeneficiaryName ensures that no value is present for BeneficiaryName, not even an explicit nil
### GetGuaranteeNumber

`func (o *BankGuaranteeCreateDto) GetGuaranteeNumber() string`

GetGuaranteeNumber returns the GuaranteeNumber field if non-nil, zero value otherwise.

### GetGuaranteeNumberOk

`func (o *BankGuaranteeCreateDto) GetGuaranteeNumberOk() (*string, bool)`

GetGuaranteeNumberOk returns a tuple with the GuaranteeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeNumber

`func (o *BankGuaranteeCreateDto) SetGuaranteeNumber(v string)`

SetGuaranteeNumber sets GuaranteeNumber field to given value.

### HasGuaranteeNumber

`func (o *BankGuaranteeCreateDto) HasGuaranteeNumber() bool`

HasGuaranteeNumber returns a boolean if a field has been set.

### SetGuaranteeNumberNil

`func (o *BankGuaranteeCreateDto) SetGuaranteeNumberNil(b bool)`

 SetGuaranteeNumberNil sets the value for GuaranteeNumber to be an explicit nil

### UnsetGuaranteeNumber
`func (o *BankGuaranteeCreateDto) UnsetGuaranteeNumber()`

UnsetGuaranteeNumber ensures that no value is present for GuaranteeNumber, not even an explicit nil
### GetGuaranteeConditions

`func (o *BankGuaranteeCreateDto) GetGuaranteeConditions() string`

GetGuaranteeConditions returns the GuaranteeConditions field if non-nil, zero value otherwise.

### GetGuaranteeConditionsOk

`func (o *BankGuaranteeCreateDto) GetGuaranteeConditionsOk() (*string, bool)`

GetGuaranteeConditionsOk returns a tuple with the GuaranteeConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeConditions

`func (o *BankGuaranteeCreateDto) SetGuaranteeConditions(v string)`

SetGuaranteeConditions sets GuaranteeConditions field to given value.

### HasGuaranteeConditions

`func (o *BankGuaranteeCreateDto) HasGuaranteeConditions() bool`

HasGuaranteeConditions returns a boolean if a field has been set.

### SetGuaranteeConditionsNil

`func (o *BankGuaranteeCreateDto) SetGuaranteeConditionsNil(b bool)`

 SetGuaranteeConditionsNil sets the value for GuaranteeConditions to be an explicit nil

### UnsetGuaranteeConditions
`func (o *BankGuaranteeCreateDto) UnsetGuaranteeConditions()`

UnsetGuaranteeConditions ensures that no value is present for GuaranteeConditions, not even an explicit nil
### GetFixedDepositNumber

`func (o *BankGuaranteeCreateDto) GetFixedDepositNumber() float64`

GetFixedDepositNumber returns the FixedDepositNumber field if non-nil, zero value otherwise.

### GetFixedDepositNumberOk

`func (o *BankGuaranteeCreateDto) GetFixedDepositNumberOk() (*float64, bool)`

GetFixedDepositNumberOk returns a tuple with the FixedDepositNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedDepositNumber

`func (o *BankGuaranteeCreateDto) SetFixedDepositNumber(v float64)`

SetFixedDepositNumber sets FixedDepositNumber field to given value.

### HasFixedDepositNumber

`func (o *BankGuaranteeCreateDto) HasFixedDepositNumber() bool`

HasFixedDepositNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *BankGuaranteeCreateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BankGuaranteeCreateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BankGuaranteeCreateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BankGuaranteeCreateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BankGuaranteeCreateDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BankGuaranteeCreateDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BankGuaranteeCreateDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BankGuaranteeCreateDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetValidityInDays

`func (o *BankGuaranteeCreateDto) GetValidityInDays() int32`

GetValidityInDays returns the ValidityInDays field if non-nil, zero value otherwise.

### GetValidityInDaysOk

`func (o *BankGuaranteeCreateDto) GetValidityInDaysOk() (*int32, bool)`

GetValidityInDaysOk returns a tuple with the ValidityInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityInDays

`func (o *BankGuaranteeCreateDto) SetValidityInDays(v int32)`

SetValidityInDays sets ValidityInDays field to given value.

### HasValidityInDays

`func (o *BankGuaranteeCreateDto) HasValidityInDays() bool`

HasValidityInDays returns a boolean if a field has been set.

### GetTenantId

`func (o *BankGuaranteeCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BankGuaranteeCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BankGuaranteeCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BankGuaranteeCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BankGuaranteeCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BankGuaranteeCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetBankGuaranteeType

`func (o *BankGuaranteeCreateDto) GetBankGuaranteeType() string`

GetBankGuaranteeType returns the BankGuaranteeType field if non-nil, zero value otherwise.

### GetBankGuaranteeTypeOk

`func (o *BankGuaranteeCreateDto) GetBankGuaranteeTypeOk() (*string, bool)`

GetBankGuaranteeTypeOk returns a tuple with the BankGuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankGuaranteeType

`func (o *BankGuaranteeCreateDto) SetBankGuaranteeType(v string)`

SetBankGuaranteeType sets BankGuaranteeType field to given value.

### HasBankGuaranteeType

`func (o *BankGuaranteeCreateDto) HasBankGuaranteeType() bool`

HasBankGuaranteeType returns a boolean if a field has been set.

### GetEnrollmentId

`func (o *BankGuaranteeCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BankGuaranteeCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BankGuaranteeCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BankGuaranteeCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BankGuaranteeCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BankGuaranteeCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetContactId

`func (o *BankGuaranteeCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *BankGuaranteeCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *BankGuaranteeCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *BankGuaranteeCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *BankGuaranteeCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *BankGuaranteeCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetProjectId

`func (o *BankGuaranteeCreateDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BankGuaranteeCreateDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BankGuaranteeCreateDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BankGuaranteeCreateDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BankGuaranteeCreateDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BankGuaranteeCreateDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOrderId

`func (o *BankGuaranteeCreateDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BankGuaranteeCreateDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BankGuaranteeCreateDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *BankGuaranteeCreateDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *BankGuaranteeCreateDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *BankGuaranteeCreateDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetBankProfileId

`func (o *BankGuaranteeCreateDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankGuaranteeCreateDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankGuaranteeCreateDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankGuaranteeCreateDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankGuaranteeCreateDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankGuaranteeCreateDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil
### GetBankAccountId

`func (o *BankGuaranteeCreateDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankGuaranteeCreateDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankGuaranteeCreateDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankGuaranteeCreateDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *BankGuaranteeCreateDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *BankGuaranteeCreateDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetCurrencyId

`func (o *BankGuaranteeCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankGuaranteeCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankGuaranteeCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BankGuaranteeCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *BankGuaranteeCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *BankGuaranteeCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


