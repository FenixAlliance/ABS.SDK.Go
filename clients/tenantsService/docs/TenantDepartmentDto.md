# TenantDepartmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**OrganizationProfileId** | Pointer to **NullableString** |  | [optional] 
**ParentDepartmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantDepartmentDto

`func NewTenantDepartmentDto() *TenantDepartmentDto`

NewTenantDepartmentDto instantiates a new TenantDepartmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDepartmentDtoWithDefaults

`func NewTenantDepartmentDtoWithDefaults() *TenantDepartmentDto`

NewTenantDepartmentDtoWithDefaults instantiates a new TenantDepartmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantDepartmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantDepartmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantDepartmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantDepartmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantDepartmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantDepartmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantDepartmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantDepartmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantDepartmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantDepartmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantDepartmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantDepartmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *TenantDepartmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantDepartmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantDepartmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantDepartmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantDepartmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantDepartmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TenantDepartmentDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TenantDepartmentDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TenantDepartmentDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TenantDepartmentDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TenantDepartmentDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TenantDepartmentDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetName

`func (o *TenantDepartmentDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantDepartmentDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantDepartmentDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantDepartmentDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantDepartmentDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantDepartmentDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantDepartmentDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantDepartmentDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantDepartmentDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantDepartmentDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantDepartmentDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantDepartmentDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisabled

`func (o *TenantDepartmentDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TenantDepartmentDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TenantDepartmentDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *TenantDepartmentDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetOrganizationProfileId

`func (o *TenantDepartmentDto) GetOrganizationProfileId() string`

GetOrganizationProfileId returns the OrganizationProfileId field if non-nil, zero value otherwise.

### GetOrganizationProfileIdOk

`func (o *TenantDepartmentDto) GetOrganizationProfileIdOk() (*string, bool)`

GetOrganizationProfileIdOk returns a tuple with the OrganizationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileId

`func (o *TenantDepartmentDto) SetOrganizationProfileId(v string)`

SetOrganizationProfileId sets OrganizationProfileId field to given value.

### HasOrganizationProfileId

`func (o *TenantDepartmentDto) HasOrganizationProfileId() bool`

HasOrganizationProfileId returns a boolean if a field has been set.

### SetOrganizationProfileIdNil

`func (o *TenantDepartmentDto) SetOrganizationProfileIdNil(b bool)`

 SetOrganizationProfileIdNil sets the value for OrganizationProfileId to be an explicit nil

### UnsetOrganizationProfileId
`func (o *TenantDepartmentDto) UnsetOrganizationProfileId()`

UnsetOrganizationProfileId ensures that no value is present for OrganizationProfileId, not even an explicit nil
### GetParentDepartmentId

`func (o *TenantDepartmentDto) GetParentDepartmentId() string`

GetParentDepartmentId returns the ParentDepartmentId field if non-nil, zero value otherwise.

### GetParentDepartmentIdOk

`func (o *TenantDepartmentDto) GetParentDepartmentIdOk() (*string, bool)`

GetParentDepartmentIdOk returns a tuple with the ParentDepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDepartmentId

`func (o *TenantDepartmentDto) SetParentDepartmentId(v string)`

SetParentDepartmentId sets ParentDepartmentId field to given value.

### HasParentDepartmentId

`func (o *TenantDepartmentDto) HasParentDepartmentId() bool`

HasParentDepartmentId returns a boolean if a field has been set.

### SetParentDepartmentIdNil

`func (o *TenantDepartmentDto) SetParentDepartmentIdNil(b bool)`

 SetParentDepartmentIdNil sets the value for ParentDepartmentId to be an explicit nil

### UnsetParentDepartmentId
`func (o *TenantDepartmentDto) UnsetParentDepartmentId()`

UnsetParentDepartmentId ensures that no value is present for ParentDepartmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


