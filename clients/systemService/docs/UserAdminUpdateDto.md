# UserAdminUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**Handler** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**PublicName** | Pointer to **NullableString** |  | [optional] 
**About** | Pointer to **NullableString** |  | [optional] 
**TwoFactorEnabled** | Pointer to **bool** |  | [optional] 
**LockoutEnabled** | Pointer to **bool** |  | [optional] 
**LockoutEnd** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewUserAdminUpdateDto

`func NewUserAdminUpdateDto() *UserAdminUpdateDto`

NewUserAdminUpdateDto instantiates a new UserAdminUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAdminUpdateDtoWithDefaults

`func NewUserAdminUpdateDtoWithDefaults() *UserAdminUpdateDto`

NewUserAdminUpdateDtoWithDefaults instantiates a new UserAdminUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserAdminUpdateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAdminUpdateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAdminUpdateDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAdminUpdateDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserAdminUpdateDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserAdminUpdateDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUserName

`func (o *UserAdminUpdateDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserAdminUpdateDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserAdminUpdateDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserAdminUpdateDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UserAdminUpdateDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserAdminUpdateDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetHandler

`func (o *UserAdminUpdateDto) GetHandler() string`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *UserAdminUpdateDto) GetHandlerOk() (*string, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *UserAdminUpdateDto) SetHandler(v string)`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *UserAdminUpdateDto) HasHandler() bool`

HasHandler returns a boolean if a field has been set.

### SetHandlerNil

`func (o *UserAdminUpdateDto) SetHandlerNil(b bool)`

 SetHandlerNil sets the value for Handler to be an explicit nil

### UnsetHandler
`func (o *UserAdminUpdateDto) UnsetHandler()`

UnsetHandler ensures that no value is present for Handler, not even an explicit nil
### GetName

`func (o *UserAdminUpdateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAdminUpdateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAdminUpdateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserAdminUpdateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserAdminUpdateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserAdminUpdateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLastName

`func (o *UserAdminUpdateDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserAdminUpdateDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserAdminUpdateDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserAdminUpdateDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserAdminUpdateDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserAdminUpdateDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPublicName

`func (o *UserAdminUpdateDto) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *UserAdminUpdateDto) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *UserAdminUpdateDto) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *UserAdminUpdateDto) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *UserAdminUpdateDto) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *UserAdminUpdateDto) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetAbout

`func (o *UserAdminUpdateDto) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *UserAdminUpdateDto) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *UserAdminUpdateDto) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *UserAdminUpdateDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### SetAboutNil

`func (o *UserAdminUpdateDto) SetAboutNil(b bool)`

 SetAboutNil sets the value for About to be an explicit nil

### UnsetAbout
`func (o *UserAdminUpdateDto) UnsetAbout()`

UnsetAbout ensures that no value is present for About, not even an explicit nil
### GetTwoFactorEnabled

`func (o *UserAdminUpdateDto) GetTwoFactorEnabled() bool`

GetTwoFactorEnabled returns the TwoFactorEnabled field if non-nil, zero value otherwise.

### GetTwoFactorEnabledOk

`func (o *UserAdminUpdateDto) GetTwoFactorEnabledOk() (*bool, bool)`

GetTwoFactorEnabledOk returns a tuple with the TwoFactorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorEnabled

`func (o *UserAdminUpdateDto) SetTwoFactorEnabled(v bool)`

SetTwoFactorEnabled sets TwoFactorEnabled field to given value.

### HasTwoFactorEnabled

`func (o *UserAdminUpdateDto) HasTwoFactorEnabled() bool`

HasTwoFactorEnabled returns a boolean if a field has been set.

### GetLockoutEnabled

`func (o *UserAdminUpdateDto) GetLockoutEnabled() bool`

GetLockoutEnabled returns the LockoutEnabled field if non-nil, zero value otherwise.

### GetLockoutEnabledOk

`func (o *UserAdminUpdateDto) GetLockoutEnabledOk() (*bool, bool)`

GetLockoutEnabledOk returns a tuple with the LockoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnabled

`func (o *UserAdminUpdateDto) SetLockoutEnabled(v bool)`

SetLockoutEnabled sets LockoutEnabled field to given value.

### HasLockoutEnabled

`func (o *UserAdminUpdateDto) HasLockoutEnabled() bool`

HasLockoutEnabled returns a boolean if a field has been set.

### GetLockoutEnd

`func (o *UserAdminUpdateDto) GetLockoutEnd() time.Time`

GetLockoutEnd returns the LockoutEnd field if non-nil, zero value otherwise.

### GetLockoutEndOk

`func (o *UserAdminUpdateDto) GetLockoutEndOk() (*time.Time, bool)`

GetLockoutEndOk returns a tuple with the LockoutEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutEnd

`func (o *UserAdminUpdateDto) SetLockoutEnd(v time.Time)`

SetLockoutEnd sets LockoutEnd field to given value.

### HasLockoutEnd

`func (o *UserAdminUpdateDto) HasLockoutEnd() bool`

HasLockoutEnd returns a boolean if a field has been set.

### SetLockoutEndNil

`func (o *UserAdminUpdateDto) SetLockoutEndNil(b bool)`

 SetLockoutEndNil sets the value for LockoutEnd to be an explicit nil

### UnsetLockoutEnd
`func (o *UserAdminUpdateDto) UnsetLockoutEnd()`

UnsetLockoutEnd ensures that no value is present for LockoutEnd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


