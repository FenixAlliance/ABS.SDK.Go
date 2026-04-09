# TenantTeamCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**BusinessID** | **string** |  | 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AvatarURL** | Pointer to **NullableString** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**BusinessUnitID** | Pointer to **NullableString** |  | [optional] 
**OrganizationProfileID** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTenantTeamCreateDto

`func NewTenantTeamCreateDto(businessID string, ) *TenantTeamCreateDto`

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

### GetBusinessID

`func (o *TenantTeamCreateDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantTeamCreateDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantTeamCreateDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.


### GetBusinessProfileRecordID

`func (o *TenantTeamCreateDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantTeamCreateDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantTeamCreateDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *TenantTeamCreateDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *TenantTeamCreateDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *TenantTeamCreateDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
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
### GetAvatarURL

`func (o *TenantTeamCreateDto) GetAvatarURL() string`

GetAvatarURL returns the AvatarURL field if non-nil, zero value otherwise.

### GetAvatarURLOk

`func (o *TenantTeamCreateDto) GetAvatarURLOk() (*string, bool)`

GetAvatarURLOk returns a tuple with the AvatarURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarURL

`func (o *TenantTeamCreateDto) SetAvatarURL(v string)`

SetAvatarURL sets AvatarURL field to given value.

### HasAvatarURL

`func (o *TenantTeamCreateDto) HasAvatarURL() bool`

HasAvatarURL returns a boolean if a field has been set.

### SetAvatarURLNil

`func (o *TenantTeamCreateDto) SetAvatarURLNil(b bool)`

 SetAvatarURLNil sets the value for AvatarURL to be an explicit nil

### UnsetAvatarURL
`func (o *TenantTeamCreateDto) UnsetAvatarURL()`

UnsetAvatarURL ensures that no value is present for AvatarURL, not even an explicit nil
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

### GetBusinessUnitID

`func (o *TenantTeamCreateDto) GetBusinessUnitID() string`

GetBusinessUnitID returns the BusinessUnitID field if non-nil, zero value otherwise.

### GetBusinessUnitIDOk

`func (o *TenantTeamCreateDto) GetBusinessUnitIDOk() (*string, bool)`

GetBusinessUnitIDOk returns a tuple with the BusinessUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitID

`func (o *TenantTeamCreateDto) SetBusinessUnitID(v string)`

SetBusinessUnitID sets BusinessUnitID field to given value.

### HasBusinessUnitID

`func (o *TenantTeamCreateDto) HasBusinessUnitID() bool`

HasBusinessUnitID returns a boolean if a field has been set.

### SetBusinessUnitIDNil

`func (o *TenantTeamCreateDto) SetBusinessUnitIDNil(b bool)`

 SetBusinessUnitIDNil sets the value for BusinessUnitID to be an explicit nil

### UnsetBusinessUnitID
`func (o *TenantTeamCreateDto) UnsetBusinessUnitID()`

UnsetBusinessUnitID ensures that no value is present for BusinessUnitID, not even an explicit nil
### GetOrganizationProfileID

`func (o *TenantTeamCreateDto) GetOrganizationProfileID() string`

GetOrganizationProfileID returns the OrganizationProfileID field if non-nil, zero value otherwise.

### GetOrganizationProfileIDOk

`func (o *TenantTeamCreateDto) GetOrganizationProfileIDOk() (*string, bool)`

GetOrganizationProfileIDOk returns a tuple with the OrganizationProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileID

`func (o *TenantTeamCreateDto) SetOrganizationProfileID(v string)`

SetOrganizationProfileID sets OrganizationProfileID field to given value.

### HasOrganizationProfileID

`func (o *TenantTeamCreateDto) HasOrganizationProfileID() bool`

HasOrganizationProfileID returns a boolean if a field has been set.

### SetOrganizationProfileIDNil

`func (o *TenantTeamCreateDto) SetOrganizationProfileIDNil(b bool)`

 SetOrganizationProfileIDNil sets the value for OrganizationProfileID to be an explicit nil

### UnsetOrganizationProfileID
`func (o *TenantTeamCreateDto) UnsetOrganizationProfileID()`

UnsetOrganizationProfileID ensures that no value is present for OrganizationProfileID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


