# TenantTeamDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableString** |  | [optional] 
**OrganizationProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantTeamDto

`func NewTenantTeamDto() *TenantTeamDto`

NewTenantTeamDto instantiates a new TenantTeamDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamDtoWithDefaults

`func NewTenantTeamDtoWithDefaults() *TenantTeamDto`

NewTenantTeamDtoWithDefaults instantiates a new TenantTeamDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantTeamDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantTeamDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantTeamDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantTeamDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantTeamDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *TenantTeamDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantTeamDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantTeamDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantTeamDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantTeamDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantTeamDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *TenantTeamDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *TenantTeamDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *TenantTeamDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *TenantTeamDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *TenantTeamDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *TenantTeamDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetName

`func (o *TenantTeamDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantTeamDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantTeamDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantTeamDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantTeamDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantTeamDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantTeamDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantTeamDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantTeamDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantTeamDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantTeamDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantTeamDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAvatarUrl

`func (o *TenantTeamDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *TenantTeamDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *TenantTeamDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *TenantTeamDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *TenantTeamDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *TenantTeamDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetIsPublic

`func (o *TenantTeamDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *TenantTeamDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *TenantTeamDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *TenantTeamDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *TenantTeamDto) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *TenantTeamDto) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *TenantTeamDto) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *TenantTeamDto) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *TenantTeamDto) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *TenantTeamDto) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetOrganizationProfileId

`func (o *TenantTeamDto) GetOrganizationProfileId() string`

GetOrganizationProfileId returns the OrganizationProfileId field if non-nil, zero value otherwise.

### GetOrganizationProfileIdOk

`func (o *TenantTeamDto) GetOrganizationProfileIdOk() (*string, bool)`

GetOrganizationProfileIdOk returns a tuple with the OrganizationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileId

`func (o *TenantTeamDto) SetOrganizationProfileId(v string)`

SetOrganizationProfileId sets OrganizationProfileId field to given value.

### HasOrganizationProfileId

`func (o *TenantTeamDto) HasOrganizationProfileId() bool`

HasOrganizationProfileId returns a boolean if a field has been set.

### SetOrganizationProfileIdNil

`func (o *TenantTeamDto) SetOrganizationProfileIdNil(b bool)`

 SetOrganizationProfileIdNil sets the value for OrganizationProfileId to be an explicit nil

### UnsetOrganizationProfileId
`func (o *TenantTeamDto) UnsetOrganizationProfileId()`

UnsetOrganizationProfileId ensures that no value is present for OrganizationProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


