# TenantIndustryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ParentBusinessIndustryId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantIndustryDto

`func NewTenantIndustryDto() *TenantIndustryDto`

NewTenantIndustryDto instantiates a new TenantIndustryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantIndustryDtoWithDefaults

`func NewTenantIndustryDtoWithDefaults() *TenantIndustryDto`

NewTenantIndustryDtoWithDefaults instantiates a new TenantIndustryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantIndustryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantIndustryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantIndustryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantIndustryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantIndustryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantIndustryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantIndustryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantIndustryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantIndustryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantIndustryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantIndustryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantIndustryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *TenantIndustryDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantIndustryDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantIndustryDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantIndustryDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantIndustryDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantIndustryDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TenantIndustryDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TenantIndustryDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TenantIndustryDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TenantIndustryDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TenantIndustryDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TenantIndustryDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetName

`func (o *TenantIndustryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantIndustryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantIndustryDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantIndustryDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantIndustryDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantIndustryDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentBusinessIndustryId

`func (o *TenantIndustryDto) GetParentBusinessIndustryId() string`

GetParentBusinessIndustryId returns the ParentBusinessIndustryId field if non-nil, zero value otherwise.

### GetParentBusinessIndustryIdOk

`func (o *TenantIndustryDto) GetParentBusinessIndustryIdOk() (*string, bool)`

GetParentBusinessIndustryIdOk returns a tuple with the ParentBusinessIndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessIndustryId

`func (o *TenantIndustryDto) SetParentBusinessIndustryId(v string)`

SetParentBusinessIndustryId sets ParentBusinessIndustryId field to given value.

### HasParentBusinessIndustryId

`func (o *TenantIndustryDto) HasParentBusinessIndustryId() bool`

HasParentBusinessIndustryId returns a boolean if a field has been set.

### SetParentBusinessIndustryIdNil

`func (o *TenantIndustryDto) SetParentBusinessIndustryIdNil(b bool)`

 SetParentBusinessIndustryIdNil sets the value for ParentBusinessIndustryId to be an explicit nil

### UnsetParentBusinessIndustryId
`func (o *TenantIndustryDto) UnsetParentBusinessIndustryId()`

UnsetParentBusinessIndustryId ensures that no value is present for ParentBusinessIndustryId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


