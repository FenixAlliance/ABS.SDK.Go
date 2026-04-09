# FiscalYearUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFiscalYearUpdateDto

`func NewFiscalYearUpdateDto() *FiscalYearUpdateDto`

NewFiscalYearUpdateDto instantiates a new FiscalYearUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalYearUpdateDtoWithDefaults

`func NewFiscalYearUpdateDtoWithDefaults() *FiscalYearUpdateDto`

NewFiscalYearUpdateDtoWithDefaults instantiates a new FiscalYearUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FiscalYearUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiscalYearUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiscalYearUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FiscalYearUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FiscalYearUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FiscalYearUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *FiscalYearUpdateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FiscalYearUpdateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FiscalYearUpdateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FiscalYearUpdateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FiscalYearUpdateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FiscalYearUpdateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetClosed

`func (o *FiscalYearUpdateDto) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *FiscalYearUpdateDto) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *FiscalYearUpdateDto) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *FiscalYearUpdateDto) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetEndDate

`func (o *FiscalYearUpdateDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FiscalYearUpdateDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FiscalYearUpdateDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FiscalYearUpdateDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *FiscalYearUpdateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FiscalYearUpdateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FiscalYearUpdateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FiscalYearUpdateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


