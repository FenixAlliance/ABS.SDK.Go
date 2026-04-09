# FiscalPeriodCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFiscalPeriodCreateDto

`func NewFiscalPeriodCreateDto() *FiscalPeriodCreateDto`

NewFiscalPeriodCreateDto instantiates a new FiscalPeriodCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalPeriodCreateDtoWithDefaults

`func NewFiscalPeriodCreateDtoWithDefaults() *FiscalPeriodCreateDto`

NewFiscalPeriodCreateDtoWithDefaults instantiates a new FiscalPeriodCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiscalPeriodCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiscalPeriodCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiscalPeriodCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiscalPeriodCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *FiscalPeriodCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FiscalPeriodCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FiscalPeriodCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FiscalPeriodCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *FiscalPeriodCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiscalPeriodCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiscalPeriodCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FiscalPeriodCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FiscalPeriodCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FiscalPeriodCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFromDate

`func (o *FiscalPeriodCreateDto) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *FiscalPeriodCreateDto) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *FiscalPeriodCreateDto) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *FiscalPeriodCreateDto) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *FiscalPeriodCreateDto) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *FiscalPeriodCreateDto) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *FiscalPeriodCreateDto) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *FiscalPeriodCreateDto) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetTenantId

`func (o *FiscalPeriodCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FiscalPeriodCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FiscalPeriodCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FiscalPeriodCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *FiscalPeriodCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *FiscalPeriodCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *FiscalPeriodCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *FiscalPeriodCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *FiscalPeriodCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *FiscalPeriodCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *FiscalPeriodCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *FiscalPeriodCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetFiscalYearId

`func (o *FiscalPeriodCreateDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *FiscalPeriodCreateDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *FiscalPeriodCreateDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *FiscalPeriodCreateDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *FiscalPeriodCreateDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *FiscalPeriodCreateDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


