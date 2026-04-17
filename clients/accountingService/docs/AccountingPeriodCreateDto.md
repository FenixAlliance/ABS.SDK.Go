# AccountingPeriodCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DateStart** | Pointer to **time.Time** |  | [optional] 
**DateEnd** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccountingPeriodCreateDto

`func NewAccountingPeriodCreateDto() *AccountingPeriodCreateDto`

NewAccountingPeriodCreateDto instantiates a new AccountingPeriodCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingPeriodCreateDtoWithDefaults

`func NewAccountingPeriodCreateDtoWithDefaults() *AccountingPeriodCreateDto`

NewAccountingPeriodCreateDtoWithDefaults instantiates a new AccountingPeriodCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountingPeriodCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingPeriodCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingPeriodCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountingPeriodCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AccountingPeriodCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountingPeriodCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountingPeriodCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountingPeriodCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *AccountingPeriodCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountingPeriodCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountingPeriodCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountingPeriodCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AccountingPeriodCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AccountingPeriodCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDateStart

`func (o *AccountingPeriodCreateDto) GetDateStart() time.Time`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *AccountingPeriodCreateDto) GetDateStartOk() (*time.Time, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *AccountingPeriodCreateDto) SetDateStart(v time.Time)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *AccountingPeriodCreateDto) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *AccountingPeriodCreateDto) GetDateEnd() time.Time`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *AccountingPeriodCreateDto) GetDateEndOk() (*time.Time, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *AccountingPeriodCreateDto) SetDateEnd(v time.Time)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *AccountingPeriodCreateDto) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


