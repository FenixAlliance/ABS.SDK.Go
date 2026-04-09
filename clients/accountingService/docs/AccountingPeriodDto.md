# AccountingPeriodDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**DateStart** | Pointer to **time.Time** |  | [optional] 
**DateEnd** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccountingPeriodDto

`func NewAccountingPeriodDto() *AccountingPeriodDto`

NewAccountingPeriodDto instantiates a new AccountingPeriodDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingPeriodDtoWithDefaults

`func NewAccountingPeriodDtoWithDefaults() *AccountingPeriodDto`

NewAccountingPeriodDtoWithDefaults instantiates a new AccountingPeriodDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountingPeriodDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingPeriodDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingPeriodDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountingPeriodDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AccountingPeriodDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AccountingPeriodDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AccountingPeriodDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountingPeriodDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountingPeriodDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountingPeriodDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AccountingPeriodDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AccountingPeriodDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *AccountingPeriodDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountingPeriodDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountingPeriodDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountingPeriodDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AccountingPeriodDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AccountingPeriodDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTenantId

`func (o *AccountingPeriodDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccountingPeriodDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccountingPeriodDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AccountingPeriodDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AccountingPeriodDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AccountingPeriodDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AccountingPeriodDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AccountingPeriodDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AccountingPeriodDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AccountingPeriodDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AccountingPeriodDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AccountingPeriodDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetDateStart

`func (o *AccountingPeriodDto) GetDateStart() time.Time`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *AccountingPeriodDto) GetDateStartOk() (*time.Time, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *AccountingPeriodDto) SetDateStart(v time.Time)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *AccountingPeriodDto) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *AccountingPeriodDto) GetDateEnd() time.Time`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *AccountingPeriodDto) GetDateEndOk() (*time.Time, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *AccountingPeriodDto) SetDateEnd(v time.Time)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *AccountingPeriodDto) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


