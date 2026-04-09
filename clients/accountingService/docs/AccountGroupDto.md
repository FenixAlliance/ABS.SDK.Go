# AccountGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ParentAccountGroupId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAccountGroupDto

`func NewAccountGroupDto() *AccountGroupDto`

NewAccountGroupDto instantiates a new AccountGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupDtoWithDefaults

`func NewAccountGroupDtoWithDefaults() *AccountGroupDto`

NewAccountGroupDtoWithDefaults instantiates a new AccountGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountGroupDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountGroupDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountGroupDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountGroupDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AccountGroupDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AccountGroupDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AccountGroupDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountGroupDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountGroupDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AccountGroupDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AccountGroupDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AccountGroupDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTitle

`func (o *AccountGroupDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AccountGroupDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AccountGroupDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AccountGroupDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AccountGroupDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AccountGroupDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AccountGroupDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountGroupDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountGroupDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountGroupDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AccountGroupDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AccountGroupDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentAccountGroupId

`func (o *AccountGroupDto) GetParentAccountGroupId() string`

GetParentAccountGroupId returns the ParentAccountGroupId field if non-nil, zero value otherwise.

### GetParentAccountGroupIdOk

`func (o *AccountGroupDto) GetParentAccountGroupIdOk() (*string, bool)`

GetParentAccountGroupIdOk returns a tuple with the ParentAccountGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountGroupId

`func (o *AccountGroupDto) SetParentAccountGroupId(v string)`

SetParentAccountGroupId sets ParentAccountGroupId field to given value.

### HasParentAccountGroupId

`func (o *AccountGroupDto) HasParentAccountGroupId() bool`

HasParentAccountGroupId returns a boolean if a field has been set.

### SetParentAccountGroupIdNil

`func (o *AccountGroupDto) SetParentAccountGroupIdNil(b bool)`

 SetParentAccountGroupIdNil sets the value for ParentAccountGroupId to be an explicit nil

### UnsetParentAccountGroupId
`func (o *AccountGroupDto) UnsetParentAccountGroupId()`

UnsetParentAccountGroupId ensures that no value is present for ParentAccountGroupId, not even an explicit nil
### GetTenantId

`func (o *AccountGroupDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AccountGroupDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AccountGroupDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AccountGroupDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AccountGroupDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AccountGroupDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AccountGroupDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AccountGroupDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AccountGroupDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AccountGroupDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AccountGroupDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AccountGroupDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


