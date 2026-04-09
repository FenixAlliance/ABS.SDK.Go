# FiscalPeriodUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFiscalPeriodUpdateDto

`func NewFiscalPeriodUpdateDto() *FiscalPeriodUpdateDto`

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

### HasName

`func (o *FiscalPeriodUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FiscalPeriodUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FiscalPeriodUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### GetTenantId

`func (o *FiscalPeriodUpdateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FiscalPeriodUpdateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FiscalPeriodUpdateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FiscalPeriodUpdateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *FiscalPeriodUpdateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *FiscalPeriodUpdateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *FiscalPeriodUpdateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *FiscalPeriodUpdateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *FiscalPeriodUpdateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *FiscalPeriodUpdateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *FiscalPeriodUpdateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *FiscalPeriodUpdateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
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

### HasFiscalYearId

`func (o *FiscalPeriodUpdateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *FiscalPeriodUpdateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *FiscalPeriodUpdateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


