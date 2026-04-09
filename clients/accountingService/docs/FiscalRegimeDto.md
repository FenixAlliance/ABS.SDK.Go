# FiscalRegimeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FiscalAuthorityId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFiscalRegimeDto

`func NewFiscalRegimeDto() *FiscalRegimeDto`

NewFiscalRegimeDto instantiates a new FiscalRegimeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalRegimeDtoWithDefaults

`func NewFiscalRegimeDtoWithDefaults() *FiscalRegimeDto`

NewFiscalRegimeDtoWithDefaults instantiates a new FiscalRegimeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiscalRegimeDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiscalRegimeDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiscalRegimeDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiscalRegimeDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FiscalRegimeDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FiscalRegimeDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *FiscalRegimeDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FiscalRegimeDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FiscalRegimeDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FiscalRegimeDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *FiscalRegimeDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *FiscalRegimeDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetCode

`func (o *FiscalRegimeDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FiscalRegimeDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FiscalRegimeDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FiscalRegimeDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *FiscalRegimeDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *FiscalRegimeDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *FiscalRegimeDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiscalRegimeDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiscalRegimeDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FiscalRegimeDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FiscalRegimeDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FiscalRegimeDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFiscalAuthorityId

`func (o *FiscalRegimeDto) GetFiscalAuthorityId() string`

GetFiscalAuthorityId returns the FiscalAuthorityId field if non-nil, zero value otherwise.

### GetFiscalAuthorityIdOk

`func (o *FiscalRegimeDto) GetFiscalAuthorityIdOk() (*string, bool)`

GetFiscalAuthorityIdOk returns a tuple with the FiscalAuthorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalAuthorityId

`func (o *FiscalRegimeDto) SetFiscalAuthorityId(v string)`

SetFiscalAuthorityId sets FiscalAuthorityId field to given value.

### HasFiscalAuthorityId

`func (o *FiscalRegimeDto) HasFiscalAuthorityId() bool`

HasFiscalAuthorityId returns a boolean if a field has been set.

### SetFiscalAuthorityIdNil

`func (o *FiscalRegimeDto) SetFiscalAuthorityIdNil(b bool)`

 SetFiscalAuthorityIdNil sets the value for FiscalAuthorityId to be an explicit nil

### UnsetFiscalAuthorityId
`func (o *FiscalRegimeDto) UnsetFiscalAuthorityId()`

UnsetFiscalAuthorityId ensures that no value is present for FiscalAuthorityId, not even an explicit nil
### GetTenantId

`func (o *FiscalRegimeDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FiscalRegimeDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FiscalRegimeDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FiscalRegimeDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *FiscalRegimeDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *FiscalRegimeDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *FiscalRegimeDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *FiscalRegimeDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *FiscalRegimeDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *FiscalRegimeDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *FiscalRegimeDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *FiscalRegimeDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


