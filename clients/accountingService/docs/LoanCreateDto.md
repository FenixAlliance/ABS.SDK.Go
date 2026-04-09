# LoanCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**LoanTimestamp** | Pointer to **time.Time** |  | [optional] 
**PaymentDeadline** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**InterestRate** | Pointer to **float64** |  | [optional] 
**IsCompundInterestRate** | Pointer to **bool** |  | [optional] 
**LoanTypeId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLoanCreateDto

`func NewLoanCreateDto() *LoanCreateDto`

NewLoanCreateDto instantiates a new LoanCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoanCreateDtoWithDefaults

`func NewLoanCreateDtoWithDefaults() *LoanCreateDto`

NewLoanCreateDtoWithDefaults instantiates a new LoanCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoanCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoanCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoanCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoanCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *LoanCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LoanCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LoanCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LoanCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetLoanTimestamp

`func (o *LoanCreateDto) GetLoanTimestamp() time.Time`

GetLoanTimestamp returns the LoanTimestamp field if non-nil, zero value otherwise.

### GetLoanTimestampOk

`func (o *LoanCreateDto) GetLoanTimestampOk() (*time.Time, bool)`

GetLoanTimestampOk returns a tuple with the LoanTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTimestamp

`func (o *LoanCreateDto) SetLoanTimestamp(v time.Time)`

SetLoanTimestamp sets LoanTimestamp field to given value.

### HasLoanTimestamp

`func (o *LoanCreateDto) HasLoanTimestamp() bool`

HasLoanTimestamp returns a boolean if a field has been set.

### GetPaymentDeadline

`func (o *LoanCreateDto) GetPaymentDeadline() time.Time`

GetPaymentDeadline returns the PaymentDeadline field if non-nil, zero value otherwise.

### GetPaymentDeadlineOk

`func (o *LoanCreateDto) GetPaymentDeadlineOk() (*time.Time, bool)`

GetPaymentDeadlineOk returns a tuple with the PaymentDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDeadline

`func (o *LoanCreateDto) SetPaymentDeadline(v time.Time)`

SetPaymentDeadline sets PaymentDeadline field to given value.

### HasPaymentDeadline

`func (o *LoanCreateDto) HasPaymentDeadline() bool`

HasPaymentDeadline returns a boolean if a field has been set.

### GetValue

`func (o *LoanCreateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoanCreateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoanCreateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *LoanCreateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInterestRate

`func (o *LoanCreateDto) GetInterestRate() float64`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *LoanCreateDto) GetInterestRateOk() (*float64, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *LoanCreateDto) SetInterestRate(v float64)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *LoanCreateDto) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIsCompundInterestRate

`func (o *LoanCreateDto) GetIsCompundInterestRate() bool`

GetIsCompundInterestRate returns the IsCompundInterestRate field if non-nil, zero value otherwise.

### GetIsCompundInterestRateOk

`func (o *LoanCreateDto) GetIsCompundInterestRateOk() (*bool, bool)`

GetIsCompundInterestRateOk returns a tuple with the IsCompundInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompundInterestRate

`func (o *LoanCreateDto) SetIsCompundInterestRate(v bool)`

SetIsCompundInterestRate sets IsCompundInterestRate field to given value.

### HasIsCompundInterestRate

`func (o *LoanCreateDto) HasIsCompundInterestRate() bool`

HasIsCompundInterestRate returns a boolean if a field has been set.

### GetLoanTypeId

`func (o *LoanCreateDto) GetLoanTypeId() string`

GetLoanTypeId returns the LoanTypeId field if non-nil, zero value otherwise.

### GetLoanTypeIdOk

`func (o *LoanCreateDto) GetLoanTypeIdOk() (*string, bool)`

GetLoanTypeIdOk returns a tuple with the LoanTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTypeId

`func (o *LoanCreateDto) SetLoanTypeId(v string)`

SetLoanTypeId sets LoanTypeId field to given value.

### HasLoanTypeId

`func (o *LoanCreateDto) HasLoanTypeId() bool`

HasLoanTypeId returns a boolean if a field has been set.

### SetLoanTypeIdNil

`func (o *LoanCreateDto) SetLoanTypeIdNil(b bool)`

 SetLoanTypeIdNil sets the value for LoanTypeId to be an explicit nil

### UnsetLoanTypeId
`func (o *LoanCreateDto) UnsetLoanTypeId()`

UnsetLoanTypeId ensures that no value is present for LoanTypeId, not even an explicit nil
### GetCurrencyId

`func (o *LoanCreateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *LoanCreateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *LoanCreateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *LoanCreateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *LoanCreateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *LoanCreateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *LoanCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LoanCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LoanCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LoanCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LoanCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LoanCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *LoanCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LoanCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LoanCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LoanCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LoanCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LoanCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


