# FiscalPeriodUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 
**FiscalYearId** | **string** |  | 

## Methods

### NewFiscalPeriodUpdateDto

`func NewFiscalPeriodUpdateDto(name string, fiscalYearId string, ) *FiscalPeriodUpdateDto`

NewFiscalPeriodUpdateDto instantiates a new FiscalPeriodUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalPeriodUpdateDtoWithDefaults

`func NewFiscalPeriodUpdateDtoWithDefaults() *FiscalPeriodUpdateDto`

NewFiscalPeriodUpdateDtoWithDefaults instantiates a new FiscalPeriodUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FiscalPeriodUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiscalPeriodUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiscalPeriodUpdateDto) SetName(v string)`

SetName sets Name field to given value.


### GetFromDate

`func (o *FiscalPeriodUpdateDto) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *FiscalPeriodUpdateDto) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *FiscalPeriodUpdateDto) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *FiscalPeriodUpdateDto) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *FiscalPeriodUpdateDto) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *FiscalPeriodUpdateDto) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *FiscalPeriodUpdateDto) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *FiscalPeriodUpdateDto) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetFiscalYearId

`func (o *FiscalPeriodUpdateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *FiscalPeriodUpdateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *FiscalPeriodUpdateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


