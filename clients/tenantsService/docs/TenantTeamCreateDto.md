# TenantTeamCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AvatarUrl** | Pointer to **NullableString** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**BusinessUnitId** | Pointer to **NullableString** |  | [optional] 
**OrganizationProfileId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantTeamCreateDto

`func NewTenantTeamCreateDto() *TenantTeamCreateDto`

NewTenantTeamCreateDto instantiates a new TenantTeamCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantTeamCreateDtoWithDefaults

`func NewTenantTeamCreateDtoWithDefaults() *TenantTeamCreateDto`

NewTenantTeamCreateDtoWithDefaults instantiates a new TenantTeamCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantTeamCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantTeamCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantTeamCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantTeamCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *TenantTeamCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantTeamCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantTeamCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantTeamCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetName

`func (o *TenantTeamCreateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantTeamCreateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantTeamCreateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantTeamCreateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TenantTeamCreateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TenantTeamCreateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *TenantTeamCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TenantTeamCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TenantTeamCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TenantTeamCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TenantTeamCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TenantTeamCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAvatarUrl

`func (o *TenantTeamCreateDto) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *TenantTeamCreateDto) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *TenantTeamCreateDto) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *TenantTeamCreateDto) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *TenantTeamCreateDto) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *TenantTeamCreateDto) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetIsPublic

`func (o *TenantTeamCreateDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *TenantTeamCreateDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *TenantTeamCreateDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *TenantTeamCreateDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *TenantTeamCreateDto) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *TenantTeamCreateDto) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *TenantTeamCreateDto) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *TenantTeamCreateDto) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *TenantTeamCreateDto) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *TenantTeamCreateDto) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetOrganizationProfileId

`func (o *TenantTeamCreateDto) GetOrganizationProfileId() string`

GetOrganizationProfileId returns the OrganizationProfileId field if non-nil, zero value otherwise.

### GetOrganizationProfileIdOk

`func (o *TenantTeamCreateDto) GetOrganizationProfileIdOk() (*string, bool)`

GetOrganizationProfileIdOk returns a tuple with the OrganizationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileId

`func (o *TenantTeamCreateDto) SetOrganizationProfileId(v string)`

SetOrganizationProfileId sets OrganizationProfileId field to given value.

### HasOrganizationProfileId

`func (o *TenantTeamCreateDto) HasOrganizationProfileId() bool`

HasOrganizationProfileId returns a boolean if a field has been set.

### SetOrganizationProfileIdNil

`func (o *TenantTeamCreateDto) SetOrganizationProfileIdNil(b bool)`

 SetOrganizationProfileIdNil sets the value for OrganizationProfileId to be an explicit nil

### UnsetOrganizationProfileId
`func (o *TenantTeamCreateDto) UnsetOrganizationProfileId()`

UnsetOrganizationProfileId ensures that no value is present for OrganizationProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


