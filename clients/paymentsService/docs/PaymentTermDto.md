# PaymentTermDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**Percentage** | Pointer to **float64** |  | [optional] 
**CreditDays** | Pointer to **float64** |  | [optional] 
**CreditWeeks** | Pointer to **float64** |  | [optional] 
**CreditMonths** | Pointer to **float64** |  | [optional] 
**CreditYears** | Pointer to **float64** |  | [optional] 
**PaymentModeID** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentTermDto

`func NewPaymentTermDto() *PaymentTermDto`

NewPaymentTermDto instantiates a new PaymentTermDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTermDtoWithDefaults

`func NewPaymentTermDtoWithDefaults() *PaymentTermDto`

NewPaymentTermDtoWithDefaults instantiates a new PaymentTermDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentTermDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentTermDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentTermDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentTermDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PaymentTermDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentTermDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *PaymentTermDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentTermDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentTermDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentTermDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *PaymentTermDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *PaymentTermDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *PaymentTermDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentTermDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentTermDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentTermDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PaymentTermDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PaymentTermDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PaymentTermDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentTermDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentTermDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentTermDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentTermDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentTermDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsTemplate

`func (o *PaymentTermDto) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *PaymentTermDto) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *PaymentTermDto) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *PaymentTermDto) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetPercentage

`func (o *PaymentTermDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PaymentTermDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PaymentTermDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *PaymentTermDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCreditDays

`func (o *PaymentTermDto) GetCreditDays() float64`

GetCreditDays returns the CreditDays field if non-nil, zero value otherwise.

### GetCreditDaysOk

`func (o *PaymentTermDto) GetCreditDaysOk() (*float64, bool)`

GetCreditDaysOk returns a tuple with the CreditDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDays

`func (o *PaymentTermDto) SetCreditDays(v float64)`

SetCreditDays sets CreditDays field to given value.

### HasCreditDays

`func (o *PaymentTermDto) HasCreditDays() bool`

HasCreditDays returns a boolean if a field has been set.

### GetCreditWeeks

`func (o *PaymentTermDto) GetCreditWeeks() float64`

GetCreditWeeks returns the CreditWeeks field if non-nil, zero value otherwise.

### GetCreditWeeksOk

`func (o *PaymentTermDto) GetCreditWeeksOk() (*float64, bool)`

GetCreditWeeksOk returns a tuple with the CreditWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditWeeks

`func (o *PaymentTermDto) SetCreditWeeks(v float64)`

SetCreditWeeks sets CreditWeeks field to given value.

### HasCreditWeeks

`func (o *PaymentTermDto) HasCreditWeeks() bool`

HasCreditWeeks returns a boolean if a field has been set.

### GetCreditMonths

`func (o *PaymentTermDto) GetCreditMonths() float64`

GetCreditMonths returns the CreditMonths field if non-nil, zero value otherwise.

### GetCreditMonthsOk

`func (o *PaymentTermDto) GetCreditMonthsOk() (*float64, bool)`

GetCreditMonthsOk returns a tuple with the CreditMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMonths

`func (o *PaymentTermDto) SetCreditMonths(v float64)`

SetCreditMonths sets CreditMonths field to given value.

### HasCreditMonths

`func (o *PaymentTermDto) HasCreditMonths() bool`

HasCreditMonths returns a boolean if a field has been set.

### GetCreditYears

`func (o *PaymentTermDto) GetCreditYears() float64`

GetCreditYears returns the CreditYears field if non-nil, zero value otherwise.

### GetCreditYearsOk

`func (o *PaymentTermDto) GetCreditYearsOk() (*float64, bool)`

GetCreditYearsOk returns a tuple with the CreditYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditYears

`func (o *PaymentTermDto) SetCreditYears(v float64)`

SetCreditYears sets CreditYears field to given value.

### HasCreditYears

`func (o *PaymentTermDto) HasCreditYears() bool`

HasCreditYears returns a boolean if a field has been set.

### GetPaymentModeID

`func (o *PaymentTermDto) GetPaymentModeID() string`

GetPaymentModeID returns the PaymentModeID field if non-nil, zero value otherwise.

### GetPaymentModeIDOk

`func (o *PaymentTermDto) GetPaymentModeIDOk() (*string, bool)`

GetPaymentModeIDOk returns a tuple with the PaymentModeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentModeID

`func (o *PaymentTermDto) SetPaymentModeID(v string)`

SetPaymentModeID sets PaymentModeID field to given value.

### HasPaymentModeID

`func (o *PaymentTermDto) HasPaymentModeID() bool`

HasPaymentModeID returns a boolean if a field has been set.

### SetPaymentModeIDNil

`func (o *PaymentTermDto) SetPaymentModeIDNil(b bool)`

 SetPaymentModeIDNil sets the value for PaymentModeID to be an explicit nil

### UnsetPaymentModeID
`func (o *PaymentTermDto) UnsetPaymentModeID()`

UnsetPaymentModeID ensures that no value is present for PaymentModeID, not even an explicit nil
### GetTenantId

`func (o *PaymentTermDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PaymentTermDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PaymentTermDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PaymentTermDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *PaymentTermDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *PaymentTermDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *PaymentTermDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *PaymentTermDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *PaymentTermDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *PaymentTermDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *PaymentTermDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *PaymentTermDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


