# TenantTeamDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**BusinessID** | Pointer to **NullableString** |  | [optional] 
**BusinessProfileRecordID** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AvatarURL** | Pointer to **NullableString** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**BusinessUnitID** | Pointer to **NullableString** |  | [optional] 
**OrganizationProfileID** | Pointer to **NullableString** |  | [optional] 

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
### GetBusinessID

`func (o *TenantTeamDto) GetBusinessID() string`

GetBusinessID returns the BusinessID field if non-nil, zero value otherwise.

### GetBusinessIDOk

`func (o *TenantTeamDto) GetBusinessIDOk() (*string, bool)`

GetBusinessIDOk returns a tuple with the BusinessID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessID

`func (o *TenantTeamDto) SetBusinessID(v string)`

SetBusinessID sets BusinessID field to given value.

### HasBusinessID

`func (o *TenantTeamDto) HasBusinessID() bool`

HasBusinessID returns a boolean if a field has been set.

### SetBusinessIDNil

`func (o *TenantTeamDto) SetBusinessIDNil(b bool)`

 SetBusinessIDNil sets the value for BusinessID to be an explicit nil

### UnsetBusinessID
`func (o *TenantTeamDto) UnsetBusinessID()`

UnsetBusinessID ensures that no value is present for BusinessID, not even an explicit nil
### GetBusinessProfileRecordID

`func (o *TenantTeamDto) GetBusinessProfileRecordID() string`

GetBusinessProfileRecordID returns the BusinessProfileRecordID field if non-nil, zero value otherwise.

### GetBusinessProfileRecordIDOk

`func (o *TenantTeamDto) GetBusinessProfileRecordIDOk() (*string, bool)`

GetBusinessProfileRecordIDOk returns a tuple with the BusinessProfileRecordID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProfileRecordID

`func (o *TenantTeamDto) SetBusinessProfileRecordID(v string)`

SetBusinessProfileRecordID sets BusinessProfileRecordID field to given value.

### HasBusinessProfileRecordID

`func (o *TenantTeamDto) HasBusinessProfileRecordID() bool`

HasBusinessProfileRecordID returns a boolean if a field has been set.

### SetBusinessProfileRecordIDNil

`func (o *TenantTeamDto) SetBusinessProfileRecordIDNil(b bool)`

 SetBusinessProfileRecordIDNil sets the value for BusinessProfileRecordID to be an explicit nil

### UnsetBusinessProfileRecordID
`func (o *TenantTeamDto) UnsetBusinessProfileRecordID()`

UnsetBusinessProfileRecordID ensures that no value is present for BusinessProfileRecordID, not even an explicit nil
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
### GetAvatarURL

`func (o *TenantTeamDto) GetAvatarURL() string`

GetAvatarURL returns the AvatarURL field if non-nil, zero value otherwise.

### GetAvatarURLOk

`func (o *TenantTeamDto) GetAvatarURLOk() (*string, bool)`

GetAvatarURLOk returns a tuple with the AvatarURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarURL

`func (o *TenantTeamDto) SetAvatarURL(v string)`

SetAvatarURL sets AvatarURL field to given value.

### HasAvatarURL

`func (o *TenantTeamDto) HasAvatarURL() bool`

HasAvatarURL returns a boolean if a field has been set.

### SetAvatarURLNil

`func (o *TenantTeamDto) SetAvatarURLNil(b bool)`

 SetAvatarURLNil sets the value for AvatarURL to be an explicit nil

### UnsetAvatarURL
`func (o *TenantTeamDto) UnsetAvatarURL()`

UnsetAvatarURL ensures that no value is present for AvatarURL, not even an explicit nil
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

### GetBusinessUnitID

`func (o *TenantTeamDto) GetBusinessUnitID() string`

GetBusinessUnitID returns the BusinessUnitID field if non-nil, zero value otherwise.

### GetBusinessUnitIDOk

`func (o *TenantTeamDto) GetBusinessUnitIDOk() (*string, bool)`

GetBusinessUnitIDOk returns a tuple with the BusinessUnitID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitID

`func (o *TenantTeamDto) SetBusinessUnitID(v string)`

SetBusinessUnitID sets BusinessUnitID field to given value.

### HasBusinessUnitID

`func (o *TenantTeamDto) HasBusinessUnitID() bool`

HasBusinessUnitID returns a boolean if a field has been set.

### SetBusinessUnitIDNil

`func (o *TenantTeamDto) SetBusinessUnitIDNil(b bool)`

 SetBusinessUnitIDNil sets the value for BusinessUnitID to be an explicit nil

### UnsetBusinessUnitID
`func (o *TenantTeamDto) UnsetBusinessUnitID()`

UnsetBusinessUnitID ensures that no value is present for BusinessUnitID, not even an explicit nil
### GetOrganizationProfileID

`func (o *TenantTeamDto) GetOrganizationProfileID() string`

GetOrganizationProfileID returns the OrganizationProfileID field if non-nil, zero value otherwise.

### GetOrganizationProfileIDOk

`func (o *TenantTeamDto) GetOrganizationProfileIDOk() (*string, bool)`

GetOrganizationProfileIDOk returns a tuple with the OrganizationProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationProfileID

`func (o *TenantTeamDto) SetOrganizationProfileID(v string)`

SetOrganizationProfileID sets OrganizationProfileID field to given value.

### HasOrganizationProfileID

`func (o *TenantTeamDto) HasOrganizationProfileID() bool`

HasOrganizationProfileID returns a boolean if a field has been set.

### SetOrganizationProfileIDNil

`func (o *TenantTeamDto) SetOrganizationProfileIDNil(b bool)`

 SetOrganizationProfileIDNil sets the value for OrganizationProfileID to be an explicit nil

### UnsetOrganizationProfileID
`func (o *TenantTeamDto) UnsetOrganizationProfileID()`

UnsetOrganizationProfileID ensures that no value is present for OrganizationProfileID, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


