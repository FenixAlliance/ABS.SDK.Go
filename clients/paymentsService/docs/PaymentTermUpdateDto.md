# PaymentTermUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsTemplate** | Pointer to **NullableBool** |  | [optional] 
**Percentage** | Pointer to **NullableFloat64** |  | [optional] 
**CreditDays** | Pointer to **NullableFloat64** |  | [optional] 
**CreditWeeks** | Pointer to **NullableFloat64** |  | [optional] 
**CreditMonths** | Pointer to **NullableFloat64** |  | [optional] 
**CreditYears** | Pointer to **NullableFloat64** |  | [optional] 
**PaymentModeID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentTermUpdateDto

`func NewPaymentTermUpdateDto() *PaymentTermUpdateDto`

NewPaymentTermUpdateDto instantiates a new PaymentTermUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTermUpdateDtoWithDefaults

`func NewPaymentTermUpdateDtoWithDefaults() *PaymentTermUpdateDto`

NewPaymentTermUpdateDtoWithDefaults instantiates a new PaymentTermUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaymentTermUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentTermUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentTermUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentTermUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PaymentTermUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PaymentTermUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *PaymentTermUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentTermUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentTermUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentTermUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentTermUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentTermUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsTemplate

`func (o *PaymentTermUpdateDto) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *PaymentTermUpdateDto) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *PaymentTermUpdateDto) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.

### HasIsTemplate

`func (o *PaymentTermUpdateDto) HasIsTemplate() bool`

HasIsTemplate returns a boolean if a field has been set.

### SetIsTemplateNil

`func (o *PaymentTermUpdateDto) SetIsTemplateNil(b bool)`

 SetIsTemplateNil sets the value for IsTemplate to be an explicit nil

### UnsetIsTemplate
`func (o *PaymentTermUpdateDto) UnsetIsTemplate()`

UnsetIsTemplate ensures that no value is present for IsTemplate, not even an explicit nil
### GetPercentage

`func (o *PaymentTermUpdateDto) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PaymentTermUpdateDto) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PaymentTermUpdateDto) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *PaymentTermUpdateDto) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *PaymentTermUpdateDto) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *PaymentTermUpdateDto) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetCreditDays

`func (o *PaymentTermUpdateDto) GetCreditDays() float64`

GetCreditDays returns the CreditDays field if non-nil, zero value otherwise.

### GetCreditDaysOk

`func (o *PaymentTermUpdateDto) GetCreditDaysOk() (*float64, bool)`

GetCreditDaysOk returns a tuple with the CreditDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditDays

`func (o *PaymentTermUpdateDto) SetCreditDays(v float64)`

SetCreditDays sets CreditDays field to given value.

### HasCreditDays

`func (o *PaymentTermUpdateDto) HasCreditDays() bool`

HasCreditDays returns a boolean if a field has been set.

### SetCreditDaysNil

`func (o *PaymentTermUpdateDto) SetCreditDaysNil(b bool)`

 SetCreditDaysNil sets the value for CreditDays to be an explicit nil

### UnsetCreditDays
`func (o *PaymentTermUpdateDto) UnsetCreditDays()`

UnsetCreditDays ensures that no value is present for CreditDays, not even an explicit nil
### GetCreditWeeks

`func (o *PaymentTermUpdateDto) GetCreditWeeks() float64`

GetCreditWeeks returns the CreditWeeks field if non-nil, zero value otherwise.

### GetCreditWeeksOk

`func (o *PaymentTermUpdateDto) GetCreditWeeksOk() (*float64, bool)`

GetCreditWeeksOk returns a tuple with the CreditWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditWeeks

`func (o *PaymentTermUpdateDto) SetCreditWeeks(v float64)`

SetCreditWeeks sets CreditWeeks field to given value.

### HasCreditWeeks

`func (o *PaymentTermUpdateDto) HasCreditWeeks() bool`

HasCreditWeeks returns a boolean if a field has been set.

### SetCreditWeeksNil

`func (o *PaymentTermUpdateDto) SetCreditWeeksNil(b bool)`

 SetCreditWeeksNil sets the value for CreditWeeks to be an explicit nil

### UnsetCreditWeeks
`func (o *PaymentTermUpdateDto) UnsetCreditWeeks()`

UnsetCreditWeeks ensures that no value is present for CreditWeeks, not even an explicit nil
### GetCreditMonths

`func (o *PaymentTermUpdateDto) GetCreditMonths() float64`

GetCreditMonths returns the CreditMonths field if non-nil, zero value otherwise.

### GetCreditMonthsOk

`func (o *PaymentTermUpdateDto) GetCreditMonthsOk() (*float64, bool)`

GetCreditMonthsOk returns a tuple with the CreditMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMonths

`func (o *PaymentTermUpdateDto) SetCreditMonths(v float64)`

SetCreditMonths sets CreditMonths field to given value.

### HasCreditMonths

`func (o *PaymentTermUpdateDto) HasCreditMonths() bool`

HasCreditMonths returns a boolean if a field has been set.

### SetCreditMonthsNil

`func (o *PaymentTermUpdateDto) SetCreditMonthsNil(b bool)`

 SetCreditMonthsNil sets the value for CreditMonths to be an explicit nil

### UnsetCreditMonths
`func (o *PaymentTermUpdateDto) UnsetCreditMonths()`

UnsetCreditMonths ensures that no value is present for CreditMonths, not even an explicit nil
### GetCreditYears

`func (o *PaymentTermUpdateDto) GetCreditYears() float64`

GetCreditYears returns the CreditYears field if non-nil, zero value otherwise.

### GetCreditYearsOk

`func (o *PaymentTermUpdateDto) GetCreditYearsOk() (*float64, bool)`

GetCreditYearsOk returns a tuple with the CreditYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditYears

`func (o *PaymentTermUpdateDto) SetCreditYears(v float64)`

SetCreditYears sets CreditYears field to given value.

### HasCreditYears

`func (o *PaymentTermUpdateDto) HasCreditYears() bool`

HasCreditYears returns a boolean if a field has been set.

### SetCreditYearsNil

`func (o *PaymentTermUpdateDto) SetCreditYearsNil(b bool)`

 SetCreditYearsNil sets the value for CreditYears to be an explicit nil

### UnsetCreditYears
`func (o *PaymentTermUpdateDto) UnsetCreditYears()`

UnsetCreditYears ensures that no value is present for CreditYears, not even an explicit nil
### GetPaymentModeID

`func (o *PaymentTermUpdateDto) GetPaymentModeID() string`

GetPaymentModeID returns the PaymentModeID field if non-nil, zero value otherwise.

### GetPaymentModeIDOk

`func (o *PaymentTermUpdateDto) GetPaymentModeIDOk() (*string, bool)`

GetPaymentModeIDOk returns a tuple with the PaymentModeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentModeID

`func (o *PaymentTermUpdateDto) SetPaymentModeID(v string)`

SetPaymentModeID sets PaymentModeID field to given value.

### HasPaymentModeID

`func (o *PaymentTermUpdateDto) HasPaymentModeID() bool`

HasPaymentModeID returns a boolean if a field has been set.

### SetPaymentModeIDNil

`func (o *PaymentTermUpdateDto) SetPaymentModeIDNil(b bool)`

 SetPaymentModeIDNil sets the value for PaymentModeID to be an explicit nil

### UnsetPaymentModeID
`func (o *PaymentTermUpdateDto) UnsetPaymentModeID()`

UnsetPaymentModeID ensures that no value is present for PaymentModeID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


