# LoanUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoanTimestamp** | Pointer to **time.Time** |  | [optional] 
**PaymentDeadline** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to **float64** |  | [optional] 
**InterestRate** | Pointer to **float64** |  | [optional] 
**IsCompundInterestRate** | Pointer to **bool** |  | [optional] 
**LoanTypeId** | Pointer to **NullableString** |  | [optional] 
**CurrencyId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLoanUpdateDto

`func NewLoanUpdateDto() *LoanUpdateDto`

NewLoanUpdateDto instantiates a new LoanUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoanUpdateDtoWithDefaults

`func NewLoanUpdateDtoWithDefaults() *LoanUpdateDto`

NewLoanUpdateDtoWithDefaults instantiates a new LoanUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoanTimestamp

`func (o *LoanUpdateDto) GetLoanTimestamp() time.Time`

GetLoanTimestamp returns the LoanTimestamp field if non-nil, zero value otherwise.

### GetLoanTimestampOk

`func (o *LoanUpdateDto) GetLoanTimestampOk() (*time.Time, bool)`

GetLoanTimestampOk returns a tuple with the LoanTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTimestamp

`func (o *LoanUpdateDto) SetLoanTimestamp(v time.Time)`

SetLoanTimestamp sets LoanTimestamp field to given value.

### HasLoanTimestamp

`func (o *LoanUpdateDto) HasLoanTimestamp() bool`

HasLoanTimestamp returns a boolean if a field has been set.

### GetPaymentDeadline

`func (o *LoanUpdateDto) GetPaymentDeadline() time.Time`

GetPaymentDeadline returns the PaymentDeadline field if non-nil, zero value otherwise.

### GetPaymentDeadlineOk

`func (o *LoanUpdateDto) GetPaymentDeadlineOk() (*time.Time, bool)`

GetPaymentDeadlineOk returns a tuple with the PaymentDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDeadline

`func (o *LoanUpdateDto) SetPaymentDeadline(v time.Time)`

SetPaymentDeadline sets PaymentDeadline field to given value.

### HasPaymentDeadline

`func (o *LoanUpdateDto) HasPaymentDeadline() bool`

HasPaymentDeadline returns a boolean if a field has been set.

### GetValue

`func (o *LoanUpdateDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoanUpdateDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoanUpdateDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *LoanUpdateDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInterestRate

`func (o *LoanUpdateDto) GetInterestRate() float64`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *LoanUpdateDto) GetInterestRateOk() (*float64, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *LoanUpdateDto) SetInterestRate(v float64)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *LoanUpdateDto) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIsCompundInterestRate

`func (o *LoanUpdateDto) GetIsCompundInterestRate() bool`

GetIsCompundInterestRate returns the IsCompundInterestRate field if non-nil, zero value otherwise.

### GetIsCompundInterestRateOk

`func (o *LoanUpdateDto) GetIsCompundInterestRateOk() (*bool, bool)`

GetIsCompundInterestRateOk returns a tuple with the IsCompundInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompundInterestRate

`func (o *LoanUpdateDto) SetIsCompundInterestRate(v bool)`

SetIsCompundInterestRate sets IsCompundInterestRate field to given value.

### HasIsCompundInterestRate

`func (o *LoanUpdateDto) HasIsCompundInterestRate() bool`

HasIsCompundInterestRate returns a boolean if a field has been set.

### GetLoanTypeId

`func (o *LoanUpdateDto) GetLoanTypeId() string`

GetLoanTypeId returns the LoanTypeId field if non-nil, zero value otherwise.

### GetLoanTypeIdOk

`func (o *LoanUpdateDto) GetLoanTypeIdOk() (*string, bool)`

GetLoanTypeIdOk returns a tuple with the LoanTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTypeId

`func (o *LoanUpdateDto) SetLoanTypeId(v string)`

SetLoanTypeId sets LoanTypeId field to given value.

### HasLoanTypeId

`func (o *LoanUpdateDto) HasLoanTypeId() bool`

HasLoanTypeId returns a boolean if a field has been set.

### SetLoanTypeIdNil

`func (o *LoanUpdateDto) SetLoanTypeIdNil(b bool)`

 SetLoanTypeIdNil sets the value for LoanTypeId to be an explicit nil

### UnsetLoanTypeId
`func (o *LoanUpdateDto) UnsetLoanTypeId()`

UnsetLoanTypeId ensures that no value is present for LoanTypeId, not even an explicit nil
### GetCurrencyId

`func (o *LoanUpdateDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *LoanUpdateDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *LoanUpdateDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *LoanUpdateDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *LoanUpdateDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *LoanUpdateDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetEnrollmentId

`func (o *LoanUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LoanUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LoanUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LoanUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LoanUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LoanUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


