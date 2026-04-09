# FiscalResponsibilityRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**FiscalResponsibilityId** | Pointer to **NullableString** |  | [optional] 
**BillingProfileId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFiscalResponsibilityRecordCreateDto

`func NewFiscalResponsibilityRecordCreateDto() *FiscalResponsibilityRecordCreateDto`

NewFiscalResponsibilityRecordCreateDto instantiates a new FiscalResponsibilityRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalResponsibilityRecordCreateDtoWithDefaults

`func NewFiscalResponsibilityRecordCreateDtoWithDefaults() *FiscalResponsibilityRecordCreateDto`

NewFiscalResponsibilityRecordCreateDtoWithDefaults instantiates a new FiscalResponsibilityRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiscalResponsibilityRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiscalResponsibilityRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiscalResponsibilityRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiscalResponsibilityRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *FiscalResponsibilityRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FiscalResponsibilityRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FiscalResponsibilityRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FiscalResponsibilityRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetFiscalResponsibilityId

`func (o *FiscalResponsibilityRecordCreateDto) GetFiscalResponsibilityId() string`

GetFiscalResponsibilityId returns the FiscalResponsibilityId field if non-nil, zero value otherwise.

### GetFiscalResponsibilityIdOk

`func (o *FiscalResponsibilityRecordCreateDto) GetFiscalResponsibilityIdOk() (*string, bool)`

GetFiscalResponsibilityIdOk returns a tuple with the FiscalResponsibilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalResponsibilityId

`func (o *FiscalResponsibilityRecordCreateDto) SetFiscalResponsibilityId(v string)`

SetFiscalResponsibilityId sets FiscalResponsibilityId field to given value.

### HasFiscalResponsibilityId

`func (o *FiscalResponsibilityRecordCreateDto) HasFiscalResponsibilityId() bool`

HasFiscalResponsibilityId returns a boolean if a field has been set.

### SetFiscalResponsibilityIdNil

`func (o *FiscalResponsibilityRecordCreateDto) SetFiscalResponsibilityIdNil(b bool)`

 SetFiscalResponsibilityIdNil sets the value for FiscalResponsibilityId to be an explicit nil

### UnsetFiscalResponsibilityId
`func (o *FiscalResponsibilityRecordCreateDto) UnsetFiscalResponsibilityId()`

UnsetFiscalResponsibilityId ensures that no value is present for FiscalResponsibilityId, not even an explicit nil
### GetBillingProfileId

`func (o *FiscalResponsibilityRecordCreateDto) GetBillingProfileId() string`

GetBillingProfileId returns the BillingProfileId field if non-nil, zero value otherwise.

### GetBillingProfileIdOk

`func (o *FiscalResponsibilityRecordCreateDto) GetBillingProfileIdOk() (*string, bool)`

GetBillingProfileIdOk returns a tuple with the BillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProfileId

`func (o *FiscalResponsibilityRecordCreateDto) SetBillingProfileId(v string)`

SetBillingProfileId sets BillingProfileId field to given value.

### HasBillingProfileId

`func (o *FiscalResponsibilityRecordCreateDto) HasBillingProfileId() bool`

HasBillingProfileId returns a boolean if a field has been set.

### SetBillingProfileIdNil

`func (o *FiscalResponsibilityRecordCreateDto) SetBillingProfileIdNil(b bool)`

 SetBillingProfileIdNil sets the value for BillingProfileId to be an explicit nil

### UnsetBillingProfileId
`func (o *FiscalResponsibilityRecordCreateDto) UnsetBillingProfileId()`

UnsetBillingProfileId ensures that no value is present for BillingProfileId, not even an explicit nil
### GetTenantId

`func (o *FiscalResponsibilityRecordCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FiscalResponsibilityRecordCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FiscalResponsibilityRecordCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FiscalResponsibilityRecordCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *FiscalResponsibilityRecordCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *FiscalResponsibilityRecordCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *FiscalResponsibilityRecordCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *FiscalResponsibilityRecordCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *FiscalResponsibilityRecordCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *FiscalResponsibilityRecordCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *FiscalResponsibilityRecordCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *FiscalResponsibilityRecordCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


