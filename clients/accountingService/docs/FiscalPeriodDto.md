# FiscalPeriodDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FromDate** | Pointer to **time.Time** |  | [optional] 
**ToDate** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**FiscalYearId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFiscalPeriodDto

`func NewFiscalPeriodDto() *FiscalPeriodDto`

NewFiscalPeriodDto instantiates a new FiscalPeriodDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalPeriodDtoWithDefaults

`func NewFiscalPeriodDtoWithDefaults() *FiscalPeriodDto`

NewFiscalPeriodDtoWithDefaults instantiates a new FiscalPeriodDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiscalPeriodDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiscalPeriodDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiscalPeriodDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiscalPeriodDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FiscalPeriodDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FiscalPeriodDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *FiscalPeriodDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FiscalPeriodDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FiscalPeriodDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FiscalPeriodDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *FiscalPeriodDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *FiscalPeriodDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *FiscalPeriodDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiscalPeriodDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiscalPeriodDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FiscalPeriodDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FiscalPeriodDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FiscalPeriodDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFromDate

`func (o *FiscalPeriodDto) GetFromDate() time.Time`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *FiscalPeriodDto) GetFromDateOk() (*time.Time, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *FiscalPeriodDto) SetFromDate(v time.Time)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *FiscalPeriodDto) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetToDate

`func (o *FiscalPeriodDto) GetToDate() time.Time`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *FiscalPeriodDto) GetToDateOk() (*time.Time, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *FiscalPeriodDto) SetToDate(v time.Time)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *FiscalPeriodDto) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### GetTenantId

`func (o *FiscalPeriodDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FiscalPeriodDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FiscalPeriodDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FiscalPeriodDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *FiscalPeriodDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *FiscalPeriodDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *FiscalPeriodDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *FiscalPeriodDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *FiscalPeriodDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *FiscalPeriodDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *FiscalPeriodDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *FiscalPeriodDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetFiscalYearId

`func (o *FiscalPeriodDto) GetFiscalYearId() string`

GetFiscalYearId returns the FiscalYearId field if non-nil, zero value otherwise.

### GetFiscalYearIdOk

`func (o *FiscalPeriodDto) GetFiscalYearIdOk() (*string, bool)`

GetFiscalYearIdOk returns a tuple with the FiscalYearId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearId

`func (o *FiscalPeriodDto) SetFiscalYearId(v string)`

SetFiscalYearId sets FiscalYearId field to given value.

### HasFiscalYearId

`func (o *FiscalPeriodDto) HasFiscalYearId() bool`

HasFiscalYearId returns a boolean if a field has been set.

### SetFiscalYearIdNil

`func (o *FiscalPeriodDto) SetFiscalYearIdNil(b bool)`

 SetFiscalYearIdNil sets the value for FiscalYearId to be an explicit nil

### UnsetFiscalYearId
`func (o *FiscalPeriodDto) UnsetFiscalYearId()`

UnsetFiscalYearId ensures that no value is present for FiscalYearId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


