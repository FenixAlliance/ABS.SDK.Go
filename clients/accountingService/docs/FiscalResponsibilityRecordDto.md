# FiscalResponsibilityRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**FiscalResponsibilityId** | Pointer to **NullableString** |  | [optional] 
**BillingProfileId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFiscalResponsibilityRecordDto

`func NewFiscalResponsibilityRecordDto() *FiscalResponsibilityRecordDto`

NewFiscalResponsibilityRecordDto instantiates a new FiscalResponsibilityRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalResponsibilityRecordDtoWithDefaults

`func NewFiscalResponsibilityRecordDtoWithDefaults() *FiscalResponsibilityRecordDto`

NewFiscalResponsibilityRecordDtoWithDefaults instantiates a new FiscalResponsibilityRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiscalResponsibilityRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiscalResponsibilityRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiscalResponsibilityRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiscalResponsibilityRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FiscalResponsibilityRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FiscalResponsibilityRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *FiscalResponsibilityRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FiscalResponsibilityRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FiscalResponsibilityRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FiscalResponsibilityRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *FiscalResponsibilityRecordDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *FiscalResponsibilityRecordDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetFiscalResponsibilityId

`func (o *FiscalResponsibilityRecordDto) GetFiscalResponsibilityId() string`

GetFiscalResponsibilityId returns the FiscalResponsibilityId field if non-nil, zero value otherwise.

### GetFiscalResponsibilityIdOk

`func (o *FiscalResponsibilityRecordDto) GetFiscalResponsibilityIdOk() (*string, bool)`

GetFiscalResponsibilityIdOk returns a tuple with the FiscalResponsibilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalResponsibilityId

`func (o *FiscalResponsibilityRecordDto) SetFiscalResponsibilityId(v string)`

SetFiscalResponsibilityId sets FiscalResponsibilityId field to given value.

### HasFiscalResponsibilityId

`func (o *FiscalResponsibilityRecordDto) HasFiscalResponsibilityId() bool`

HasFiscalResponsibilityId returns a boolean if a field has been set.

### SetFiscalResponsibilityIdNil

`func (o *FiscalResponsibilityRecordDto) SetFiscalResponsibilityIdNil(b bool)`

 SetFiscalResponsibilityIdNil sets the value for FiscalResponsibilityId to be an explicit nil

### UnsetFiscalResponsibilityId
`func (o *FiscalResponsibilityRecordDto) UnsetFiscalResponsibilityId()`

UnsetFiscalResponsibilityId ensures that no value is present for FiscalResponsibilityId, not even an explicit nil
### GetBillingProfileId

`func (o *FiscalResponsibilityRecordDto) GetBillingProfileId() string`

GetBillingProfileId returns the BillingProfileId field if non-nil, zero value otherwise.

### GetBillingProfileIdOk

`func (o *FiscalResponsibilityRecordDto) GetBillingProfileIdOk() (*string, bool)`

GetBillingProfileIdOk returns a tuple with the BillingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProfileId

`func (o *FiscalResponsibilityRecordDto) SetBillingProfileId(v string)`

SetBillingProfileId sets BillingProfileId field to given value.

### HasBillingProfileId

`func (o *FiscalResponsibilityRecordDto) HasBillingProfileId() bool`

HasBillingProfileId returns a boolean if a field has been set.

### SetBillingProfileIdNil

`func (o *FiscalResponsibilityRecordDto) SetBillingProfileIdNil(b bool)`

 SetBillingProfileIdNil sets the value for BillingProfileId to be an explicit nil

### UnsetBillingProfileId
`func (o *FiscalResponsibilityRecordDto) UnsetBillingProfileId()`

UnsetBillingProfileId ensures that no value is present for BillingProfileId, not even an explicit nil
### GetTenantId

`func (o *FiscalResponsibilityRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FiscalResponsibilityRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FiscalResponsibilityRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FiscalResponsibilityRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *FiscalResponsibilityRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *FiscalResponsibilityRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *FiscalResponsibilityRecordDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *FiscalResponsibilityRecordDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *FiscalResponsibilityRecordDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *FiscalResponsibilityRecordDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *FiscalResponsibilityRecordDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *FiscalResponsibilityRecordDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


