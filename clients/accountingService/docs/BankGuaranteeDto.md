# BankGuaranteeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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

### NewBankGuaranteeDto

`func NewBankGuaranteeDto() *BankGuaranteeDto`

NewBankGuaranteeDto instantiates a new BankGuaranteeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankGuaranteeDtoWithDefaults

`func NewBankGuaranteeDtoWithDefaults() *BankGuaranteeDto`

NewBankGuaranteeDtoWithDefaults instantiates a new BankGuaranteeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankGuaranteeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankGuaranteeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankGuaranteeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankGuaranteeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BankGuaranteeDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BankGuaranteeDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *BankGuaranteeDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BankGuaranteeDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BankGuaranteeDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BankGuaranteeDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *BankGuaranteeDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *BankGuaranteeDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetMargin

`func (o *BankGuaranteeDto) GetMargin() float64`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *BankGuaranteeDto) GetMarginOk() (*float64, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *BankGuaranteeDto) SetMargin(v float64)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *BankGuaranteeDto) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetCharges

`func (o *BankGuaranteeDto) GetCharges() float64`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *BankGuaranteeDto) GetChargesOk() (*float64, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *BankGuaranteeDto) SetCharges(v float64)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *BankGuaranteeDto) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetBeneficiaryName

`func (o *BankGuaranteeDto) GetBeneficiaryName() string`

GetBeneficiaryName returns the BeneficiaryName field if non-nil, zero value otherwise.

### GetBeneficiaryNameOk

`func (o *BankGuaranteeDto) GetBeneficiaryNameOk() (*string, bool)`

GetBeneficiaryNameOk returns a tuple with the BeneficiaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryName

`func (o *BankGuaranteeDto) SetBeneficiaryName(v string)`

SetBeneficiaryName sets BeneficiaryName field to given value.

### HasBeneficiaryName

`func (o *BankGuaranteeDto) HasBeneficiaryName() bool`

HasBeneficiaryName returns a boolean if a field has been set.

### SetBeneficiaryNameNil

`func (o *BankGuaranteeDto) SetBeneficiaryNameNil(b bool)`

 SetBeneficiaryNameNil sets the value for BeneficiaryName to be an explicit nil

### UnsetBeneficiaryName
`func (o *BankGuaranteeDto) UnsetBeneficiaryName()`

UnsetBeneficiaryName ensures that no value is present for BeneficiaryName, not even an explicit nil
### GetGuaranteeNumber

`func (o *BankGuaranteeDto) GetGuaranteeNumber() string`

GetGuaranteeNumber returns the GuaranteeNumber field if non-nil, zero value otherwise.

### GetGuaranteeNumberOk

`func (o *BankGuaranteeDto) GetGuaranteeNumberOk() (*string, bool)`

GetGuaranteeNumberOk returns a tuple with the GuaranteeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeNumber

`func (o *BankGuaranteeDto) SetGuaranteeNumber(v string)`

SetGuaranteeNumber sets GuaranteeNumber field to given value.

### HasGuaranteeNumber

`func (o *BankGuaranteeDto) HasGuaranteeNumber() bool`

HasGuaranteeNumber returns a boolean if a field has been set.

### SetGuaranteeNumberNil

`func (o *BankGuaranteeDto) SetGuaranteeNumberNil(b bool)`

 SetGuaranteeNumberNil sets the value for GuaranteeNumber to be an explicit nil

### UnsetGuaranteeNumber
`func (o *BankGuaranteeDto) UnsetGuaranteeNumber()`

UnsetGuaranteeNumber ensures that no value is present for GuaranteeNumber, not even an explicit nil
### GetGuaranteeConditions

`func (o *BankGuaranteeDto) GetGuaranteeConditions() string`

GetGuaranteeConditions returns the GuaranteeConditions field if non-nil, zero value otherwise.

### GetGuaranteeConditionsOk

`func (o *BankGuaranteeDto) GetGuaranteeConditionsOk() (*string, bool)`

GetGuaranteeConditionsOk returns a tuple with the GuaranteeConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeConditions

`func (o *BankGuaranteeDto) SetGuaranteeConditions(v string)`

SetGuaranteeConditions sets GuaranteeConditions field to given value.

### HasGuaranteeConditions

`func (o *BankGuaranteeDto) HasGuaranteeConditions() bool`

HasGuaranteeConditions returns a boolean if a field has been set.

### SetGuaranteeConditionsNil

`func (o *BankGuaranteeDto) SetGuaranteeConditionsNil(b bool)`

 SetGuaranteeConditionsNil sets the value for GuaranteeConditions to be an explicit nil

### UnsetGuaranteeConditions
`func (o *BankGuaranteeDto) UnsetGuaranteeConditions()`

UnsetGuaranteeConditions ensures that no value is present for GuaranteeConditions, not even an explicit nil
### GetFixedDepositNumber

`func (o *BankGuaranteeDto) GetFixedDepositNumber() float64`

GetFixedDepositNumber returns the FixedDepositNumber field if non-nil, zero value otherwise.

### GetFixedDepositNumberOk

`func (o *BankGuaranteeDto) GetFixedDepositNumberOk() (*float64, bool)`

GetFixedDepositNumberOk returns a tuple with the FixedDepositNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedDepositNumber

`func (o *BankGuaranteeDto) SetFixedDepositNumber(v float64)`

SetFixedDepositNumber sets FixedDepositNumber field to given value.

### HasFixedDepositNumber

`func (o *BankGuaranteeDto) HasFixedDepositNumber() bool`

HasFixedDepositNumber returns a boolean if a field has been set.

### GetStartDate

`func (o *BankGuaranteeDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BankGuaranteeDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BankGuaranteeDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BankGuaranteeDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BankGuaranteeDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BankGuaranteeDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BankGuaranteeDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BankGuaranteeDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetValidityInDays

`func (o *BankGuaranteeDto) GetValidityInDays() int32`

GetValidityInDays returns the ValidityInDays field if non-nil, zero value otherwise.

### GetValidityInDaysOk

`func (o *BankGuaranteeDto) GetValidityInDaysOk() (*int32, bool)`

GetValidityInDaysOk returns a tuple with the ValidityInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityInDays

`func (o *BankGuaranteeDto) SetValidityInDays(v int32)`

SetValidityInDays sets ValidityInDays field to given value.

### HasValidityInDays

`func (o *BankGuaranteeDto) HasValidityInDays() bool`

HasValidityInDays returns a boolean if a field has been set.

### GetTenantId

`func (o *BankGuaranteeDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BankGuaranteeDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BankGuaranteeDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *BankGuaranteeDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *BankGuaranteeDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *BankGuaranteeDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetBankGuaranteeType

`func (o *BankGuaranteeDto) GetBankGuaranteeType() string`

GetBankGuaranteeType returns the BankGuaranteeType field if non-nil, zero value otherwise.

### GetBankGuaranteeTypeOk

`func (o *BankGuaranteeDto) GetBankGuaranteeTypeOk() (*string, bool)`

GetBankGuaranteeTypeOk returns a tuple with the BankGuaranteeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankGuaranteeType

`func (o *BankGuaranteeDto) SetBankGuaranteeType(v string)`

SetBankGuaranteeType sets BankGuaranteeType field to given value.

### HasBankGuaranteeType

`func (o *BankGuaranteeDto) HasBankGuaranteeType() bool`

HasBankGuaranteeType returns a boolean if a field has been set.

### GetEnrollmentId

`func (o *BankGuaranteeDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *BankGuaranteeDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *BankGuaranteeDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *BankGuaranteeDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *BankGuaranteeDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *BankGuaranteeDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetContactId

`func (o *BankGuaranteeDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *BankGuaranteeDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *BankGuaranteeDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *BankGuaranteeDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *BankGuaranteeDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *BankGuaranteeDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetProjectId

`func (o *BankGuaranteeDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BankGuaranteeDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BankGuaranteeDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BankGuaranteeDto) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BankGuaranteeDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BankGuaranteeDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOrderId

`func (o *BankGuaranteeDto) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BankGuaranteeDto) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BankGuaranteeDto) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *BankGuaranteeDto) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *BankGuaranteeDto) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *BankGuaranteeDto) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetBankProfileId

`func (o *BankGuaranteeDto) GetBankProfileId() string`

GetBankProfileId returns the BankProfileId field if non-nil, zero value otherwise.

### GetBankProfileIdOk

`func (o *BankGuaranteeDto) GetBankProfileIdOk() (*string, bool)`

GetBankProfileIdOk returns a tuple with the BankProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankProfileId

`func (o *BankGuaranteeDto) SetBankProfileId(v string)`

SetBankProfileId sets BankProfileId field to given value.

### HasBankProfileId

`func (o *BankGuaranteeDto) HasBankProfileId() bool`

HasBankProfileId returns a boolean if a field has been set.

### SetBankProfileIdNil

`func (o *BankGuaranteeDto) SetBankProfileIdNil(b bool)`

 SetBankProfileIdNil sets the value for BankProfileId to be an explicit nil

### UnsetBankProfileId
`func (o *BankGuaranteeDto) UnsetBankProfileId()`

UnsetBankProfileId ensures that no value is present for BankProfileId, not even an explicit nil
### GetBankAccountId

`func (o *BankGuaranteeDto) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankGuaranteeDto) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankGuaranteeDto) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankGuaranteeDto) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *BankGuaranteeDto) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *BankGuaranteeDto) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetCurrencyId

`func (o *BankGuaranteeDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankGuaranteeDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankGuaranteeDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BankGuaranteeDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *BankGuaranteeDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *BankGuaranteeDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


