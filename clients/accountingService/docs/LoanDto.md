# LoanDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
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

### NewLoanDto

`func NewLoanDto() *LoanDto`

NewLoanDto instantiates a new LoanDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoanDtoWithDefaults

`func NewLoanDtoWithDefaults() *LoanDto`

NewLoanDtoWithDefaults instantiates a new LoanDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoanDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoanDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoanDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoanDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LoanDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LoanDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *LoanDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LoanDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LoanDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LoanDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *LoanDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *LoanDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetLoanTimestamp

`func (o *LoanDto) GetLoanTimestamp() time.Time`

GetLoanTimestamp returns the LoanTimestamp field if non-nil, zero value otherwise.

### GetLoanTimestampOk

`func (o *LoanDto) GetLoanTimestampOk() (*time.Time, bool)`

GetLoanTimestampOk returns a tuple with the LoanTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTimestamp

`func (o *LoanDto) SetLoanTimestamp(v time.Time)`

SetLoanTimestamp sets LoanTimestamp field to given value.

### HasLoanTimestamp

`func (o *LoanDto) HasLoanTimestamp() bool`

HasLoanTimestamp returns a boolean if a field has been set.

### GetPaymentDeadline

`func (o *LoanDto) GetPaymentDeadline() time.Time`

GetPaymentDeadline returns the PaymentDeadline field if non-nil, zero value otherwise.

### GetPaymentDeadlineOk

`func (o *LoanDto) GetPaymentDeadlineOk() (*time.Time, bool)`

GetPaymentDeadlineOk returns a tuple with the PaymentDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDeadline

`func (o *LoanDto) SetPaymentDeadline(v time.Time)`

SetPaymentDeadline sets PaymentDeadline field to given value.

### HasPaymentDeadline

`func (o *LoanDto) HasPaymentDeadline() bool`

HasPaymentDeadline returns a boolean if a field has been set.

### GetValue

`func (o *LoanDto) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoanDto) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoanDto) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *LoanDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInterestRate

`func (o *LoanDto) GetInterestRate() float64`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *LoanDto) GetInterestRateOk() (*float64, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *LoanDto) SetInterestRate(v float64)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *LoanDto) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetIsCompundInterestRate

`func (o *LoanDto) GetIsCompundInterestRate() bool`

GetIsCompundInterestRate returns the IsCompundInterestRate field if non-nil, zero value otherwise.

### GetIsCompundInterestRateOk

`func (o *LoanDto) GetIsCompundInterestRateOk() (*bool, bool)`

GetIsCompundInterestRateOk returns a tuple with the IsCompundInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompundInterestRate

`func (o *LoanDto) SetIsCompundInterestRate(v bool)`

SetIsCompundInterestRate sets IsCompundInterestRate field to given value.

### HasIsCompundInterestRate

`func (o *LoanDto) HasIsCompundInterestRate() bool`

HasIsCompundInterestRate returns a boolean if a field has been set.

### GetLoanTypeId

`func (o *LoanDto) GetLoanTypeId() string`

GetLoanTypeId returns the LoanTypeId field if non-nil, zero value otherwise.

### GetLoanTypeIdOk

`func (o *LoanDto) GetLoanTypeIdOk() (*string, bool)`

GetLoanTypeIdOk returns a tuple with the LoanTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTypeId

`func (o *LoanDto) SetLoanTypeId(v string)`

SetLoanTypeId sets LoanTypeId field to given value.

### HasLoanTypeId

`func (o *LoanDto) HasLoanTypeId() bool`

HasLoanTypeId returns a boolean if a field has been set.

### SetLoanTypeIdNil

`func (o *LoanDto) SetLoanTypeIdNil(b bool)`

 SetLoanTypeIdNil sets the value for LoanTypeId to be an explicit nil

### UnsetLoanTypeId
`func (o *LoanDto) UnsetLoanTypeId()`

UnsetLoanTypeId ensures that no value is present for LoanTypeId, not even an explicit nil
### GetCurrencyId

`func (o *LoanDto) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *LoanDto) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *LoanDto) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *LoanDto) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### SetCurrencyIdNil

`func (o *LoanDto) SetCurrencyIdNil(b bool)`

 SetCurrencyIdNil sets the value for CurrencyId to be an explicit nil

### UnsetCurrencyId
`func (o *LoanDto) UnsetCurrencyId()`

UnsetCurrencyId ensures that no value is present for CurrencyId, not even an explicit nil
### GetTenantId

`func (o *LoanDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *LoanDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *LoanDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *LoanDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *LoanDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *LoanDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *LoanDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *LoanDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *LoanDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *LoanDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *LoanDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *LoanDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


