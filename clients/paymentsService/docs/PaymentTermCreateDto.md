# PaymentTermCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsTemplate** | Pointer to **bool** |  | [optional] 
**Percentage** | Pointer to **float64** |  | [optional] 
**CreditDays** | Pointer to **float64** |  | [optional] 
**CreditWeeks** | Pointer to **float64** |  | [optional] 
**CreditMonths** | Pointer to **float64** |  | [optional] 
**CreditYears** | Pointer to **float64** |  | [optional] 
**PaymentModeID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentTermCreateDto

`func NewPaymentTermCreateDto(name string, ) *PaymentTermCreateDto`

NewPaymentTermCreateDto instantiates a new PaymentTermCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTermCreateDtoWithDefaults

`func NewPaymentTermCreateDtoWithDefaults() *PaymentTermCreateDto`

NewPaymentTermCreateDtoWithDefaults instantiates a new PaymentTermCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentTermCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentTermCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentTermCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentTermCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *PaymentTermCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentTermCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentTermCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PaymentTermCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *PaymentTermCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentTermCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentTermCreateDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PaymentTermCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentTermCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentTermCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentTermCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentTermCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentTermCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsTemplate

`func (o *PaymentTermCreateDto) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *PaymentTermCreateDto) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *PaymentTermCreateDto) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *PaymentTermCreateDto) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### GetPercentage

`func (o *PaymentTermCreateDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PaymentTermCreateDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PaymentTermCreateDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *PaymentTermCreateDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetCreditDays

`func (o *PaymentTermCreateDto) GetCreditDays() float64`

GetCreditDays returns the CreditDays field if non-nil, zero value otherwise.

### GetCreditDaysOk

`func (o *PaymentTermCreateDto) GetCreditDaysOk() (*float64, bool)`

GetCreditDaysOk returns a tuple with the CreditDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDays

`func (o *PaymentTermCreateDto) SetCreditDays(v float64)`

SetCreditDays sets CreditDays field to given value.

### HasCreditDays

`func (o *PaymentTermCreateDto) HasCreditDays() bool`

HasCreditDays returns a boolean if a field has been set.

### GetCreditWeeks

`func (o *PaymentTermCreateDto) GetCreditWeeks() float64`

GetCreditWeeks returns the CreditWeeks field if non-nil, zero value otherwise.

### GetCreditWeeksOk

`func (o *PaymentTermCreateDto) GetCreditWeeksOk() (*float64, bool)`

GetCreditWeeksOk returns a tuple with the CreditWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditWeeks

`func (o *PaymentTermCreateDto) SetCreditWeeks(v float64)`

SetCreditWeeks sets CreditWeeks field to given value.

### HasCreditWeeks

`func (o *PaymentTermCreateDto) HasCreditWeeks() bool`

HasCreditWeeks returns a boolean if a field has been set.

### GetCreditMonths

`func (o *PaymentTermCreateDto) GetCreditMonths() float64`

GetCreditMonths returns the CreditMonths field if non-nil, zero value otherwise.

### GetCreditMonthsOk

`func (o *PaymentTermCreateDto) GetCreditMonthsOk() (*float64, bool)`

GetCreditMonthsOk returns a tuple with the CreditMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMonths

`func (o *PaymentTermCreateDto) SetCreditMonths(v float64)`

SetCreditMonths sets CreditMonths field to given value.

### HasCreditMonths

`func (o *PaymentTermCreateDto) HasCreditMonths() bool`

HasCreditMonths returns a boolean if a field has been set.

### GetCreditYears

`func (o *PaymentTermCreateDto) GetCreditYears() float64`

GetCreditYears returns the CreditYears field if non-nil, zero value otherwise.

### GetCreditYearsOk

`func (o *PaymentTermCreateDto) GetCreditYearsOk() (*float64, bool)`

GetCreditYearsOk returns a tuple with the CreditYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditYears

`func (o *PaymentTermCreateDto) SetCreditYears(v float64)`

SetCreditYears sets CreditYears field to given value.

### HasCreditYears

`func (o *PaymentTermCreateDto) HasCreditYears() bool`

HasCreditYears returns a boolean if a field has been set.

### GetPaymentModeID

`func (o *PaymentTermCreateDto) GetPaymentModeID() string`

GetPaymentModeID returns the PaymentModeID field if non-nil, zero value otherwise.

### GetPaymentModeIDOk

`func (o *PaymentTermCreateDto) GetPaymentModeIDOk() (*string, bool)`

GetPaymentModeIDOk returns a tuple with the PaymentModeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentModeID

`func (o *PaymentTermCreateDto) SetPaymentModeID(v string)`

SetPaymentModeID sets PaymentModeID field to given value.

### HasPaymentModeID

`func (o *PaymentTermCreateDto) HasPaymentModeID() bool`

HasPaymentModeID returns a boolean if a field has been set.

### SetPaymentModeIDNil

`func (o *PaymentTermCreateDto) SetPaymentModeIDNil(b bool)`

 SetPaymentModeIDNil sets the value for PaymentModeID to be an explicit nil

### UnsetPaymentModeID
`func (o *PaymentTermCreateDto) UnsetPaymentModeID()`

UnsetPaymentModeID ensures that no value is present for PaymentModeID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


