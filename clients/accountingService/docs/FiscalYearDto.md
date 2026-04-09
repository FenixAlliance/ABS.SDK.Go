# FiscalYearDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**FiscalAuthorityId** | Pointer to **NullableString** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFiscalYearDto

`func NewFiscalYearDto() *FiscalYearDto`

NewFiscalYearDto instantiates a new FiscalYearDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalYearDtoWithDefaults

`func NewFiscalYearDtoWithDefaults() *FiscalYearDto`

NewFiscalYearDtoWithDefaults instantiates a new FiscalYearDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiscalYearDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiscalYearDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiscalYearDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiscalYearDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FiscalYearDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FiscalYearDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *FiscalYearDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FiscalYearDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FiscalYearDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FiscalYearDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *FiscalYearDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *FiscalYearDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetName

`func (o *FiscalYearDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiscalYearDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiscalYearDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FiscalYearDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FiscalYearDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FiscalYearDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *FiscalYearDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FiscalYearDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FiscalYearDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FiscalYearDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FiscalYearDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FiscalYearDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDefault

`func (o *FiscalYearDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FiscalYearDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FiscalYearDto) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FiscalYearDto) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetTenantId

`func (o *FiscalYearDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FiscalYearDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FiscalYearDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FiscalYearDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *FiscalYearDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *FiscalYearDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *FiscalYearDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *FiscalYearDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *FiscalYearDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *FiscalYearDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *FiscalYearDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *FiscalYearDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetFiscalAuthorityId

`func (o *FiscalYearDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *FiscalYearDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *FiscalYearDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *FiscalYearDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *FiscalYearDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *FiscalYearDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetEndDate

`func (o *FiscalYearDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FiscalYearDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FiscalYearDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FiscalYearDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *FiscalYearDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FiscalYearDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FiscalYearDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FiscalYearDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


