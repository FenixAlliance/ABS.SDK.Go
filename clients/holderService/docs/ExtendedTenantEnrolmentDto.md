# ExtendedTenantEnrolmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**IsRoot** | Pointer to **bool** |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**Tenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 
**User** | Pointer to [**UserDto**](UserDto.md) |  | [optional] 

## Methods

### NewExtendedTenantEnrolmentDto

`func NewExtendedTenantEnrolmentDto() *ExtendedTenantEnrolmentDto`

NewExtendedTenantEnrolmentDto instantiates a new ExtendedTenantEnrolmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedTenantEnrolmentDtoWithDefaults

`func NewExtendedTenantEnrolmentDtoWithDefaults() *ExtendedTenantEnrolmentDto`

NewExtendedTenantEnrolmentDtoWithDefaults instantiates a new ExtendedTenantEnrolmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedTenantEnrolmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedTenantEnrolmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedTenantEnrolmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedTenantEnrolmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedTenantEnrolmentDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedTenantEnrolmentDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ExtendedTenantEnrolmentDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedTenantEnrolmentDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedTenantEnrolmentDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtendedTenantEnrolmentDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ExtendedTenantEnrolmentDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ExtendedTenantEnrolmentDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *ExtendedTenantEnrolmentDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtendedTenantEnrolmentDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtendedTenantEnrolmentDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ExtendedTenantEnrolmentDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ExtendedTenantEnrolmentDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ExtendedTenantEnrolmentDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetUserId

`func (o *ExtendedTenantEnrolmentDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExtendedTenantEnrolmentDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExtendedTenantEnrolmentDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ExtendedTenantEnrolmentDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ExtendedTenantEnrolmentDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ExtendedTenantEnrolmentDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetIsRoot

`func (o *ExtendedTenantEnrolmentDto) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *ExtendedTenantEnrolmentDto) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *ExtendedTenantEnrolmentDto) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *ExtendedTenantEnrolmentDto) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetIsOwner

`func (o *ExtendedTenantEnrolmentDto) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *ExtendedTenantEnrolmentDto) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *ExtendedTenantEnrolmentDto) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *ExtendedTenantEnrolmentDto) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsAdmin

`func (o *ExtendedTenantEnrolmentDto) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *ExtendedTenantEnrolmentDto) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *ExtendedTenantEnrolmentDto) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *ExtendedTenantEnrolmentDto) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ExtendedTenantEnrolmentDto) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ExtendedTenantEnrolmentDto) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ExtendedTenantEnrolmentDto) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ExtendedTenantEnrolmentDto) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetTenant

`func (o *ExtendedTenantEnrolmentDto) GetTenant() TenantDto`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ExtendedTenantEnrolmentDto) GetTenantOk() (*TenantDto, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ExtendedTenantEnrolmentDto) SetTenant(v TenantDto)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ExtendedTenantEnrolmentDto) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUser

`func (o *ExtendedTenantEnrolmentDto) GetUser() UserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ExtendedTenantEnrolmentDto) GetUserOk() (*UserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ExtendedTenantEnrolmentDto) SetUser(v UserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ExtendedTenantEnrolmentDto) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


