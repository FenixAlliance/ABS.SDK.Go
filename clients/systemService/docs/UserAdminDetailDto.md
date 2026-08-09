# UserAdminDetailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]UserOrderSummaryDto**](UserOrderSummaryDto.md) |  | [optional] 
**Logins** | Pointer to [**[]UserExternalLoginDto**](UserExternalLoginDto.md) |  | [optional] 
**Enrollment** | Pointer to [**TenantEnrollmentDto**](TenantEnrollmentDto.md) |  | [optional] 
**GrantedRoles** | Pointer to [**[]SecurityRoleDto**](SecurityRoleDto.md) |  | [optional] 
**GrantedPermissions** | Pointer to [**[]SecurityPermissionDto**](SecurityPermissionDto.md) |  | [optional] 
**RoleCatalog** | Pointer to [**[]SecurityRoleDto**](SecurityRoleDto.md) |  | [optional] 
**PermissionCatalog** | Pointer to [**[]SecurityPermissionDto**](SecurityPermissionDto.md) |  | [optional] 

## Methods

### NewUserAdminDetailDto

`func NewUserAdminDetailDto() *UserAdminDetailDto`

NewUserAdminDetailDto instantiates a new UserAdminDetailDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAdminDetailDtoWithDefaults

`func NewUserAdminDetailDtoWithDefaults() *UserAdminDetailDto`

NewUserAdminDetailDtoWithDefaults instantiates a new UserAdminDetailDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *UserAdminDetailDto) GetOrders() []UserOrderSummaryDto`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *UserAdminDetailDto) GetOrdersOk() (*[]UserOrderSummaryDto, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *UserAdminDetailDto) SetOrders(v []UserOrderSummaryDto)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *UserAdminDetailDto) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### SetOrdersNil

`func (o *UserAdminDetailDto) SetOrdersNil(b bool)`

 SetOrdersNil sets the value for Orders to be an explicit nil

### UnsetOrders
`func (o *UserAdminDetailDto) UnsetOrders()`

UnsetOrders ensures that no value is present for Orders, not even an explicit nil
### GetLogins

`func (o *UserAdminDetailDto) GetLogins() []UserExternalLoginDto`

GetLogins returns the Logins field if non-nil, zero value otherwise.

### GetLoginsOk

`func (o *UserAdminDetailDto) GetLoginsOk() (*[]UserExternalLoginDto, bool)`

GetLoginsOk returns a tuple with the Logins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogins

`func (o *UserAdminDetailDto) SetLogins(v []UserExternalLoginDto)`

SetLogins sets Logins field to given value.

### HasLogins

`func (o *UserAdminDetailDto) HasLogins() bool`

HasLogins returns a boolean if a field has been set.

### SetLoginsNil

`func (o *UserAdminDetailDto) SetLoginsNil(b bool)`

 SetLoginsNil sets the value for Logins to be an explicit nil

### UnsetLogins
`func (o *UserAdminDetailDto) UnsetLogins()`

UnsetLogins ensures that no value is present for Logins, not even an explicit nil
### GetEnrollment

`func (o *UserAdminDetailDto) GetEnrollment() TenantEnrollmentDto`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *UserAdminDetailDto) GetEnrollmentOk() (*TenantEnrollmentDto, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *UserAdminDetailDto) SetEnrollment(v TenantEnrollmentDto)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *UserAdminDetailDto) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.

### GetGrantedRoles

`func (o *UserAdminDetailDto) GetGrantedRoles() []SecurityRoleDto`

GetGrantedRoles returns the GrantedRoles field if non-nil, zero value otherwise.

### GetGrantedRolesOk

`func (o *UserAdminDetailDto) GetGrantedRolesOk() (*[]SecurityRoleDto, bool)`

GetGrantedRolesOk returns a tuple with the GrantedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedRoles

`func (o *UserAdminDetailDto) SetGrantedRoles(v []SecurityRoleDto)`

SetGrantedRoles sets GrantedRoles field to given value.

### HasGrantedRoles

`func (o *UserAdminDetailDto) HasGrantedRoles() bool`

HasGrantedRoles returns a boolean if a field has been set.

### SetGrantedRolesNil

`func (o *UserAdminDetailDto) SetGrantedRolesNil(b bool)`

 SetGrantedRolesNil sets the value for GrantedRoles to be an explicit nil

### UnsetGrantedRoles
`func (o *UserAdminDetailDto) UnsetGrantedRoles()`

UnsetGrantedRoles ensures that no value is present for GrantedRoles, not even an explicit nil
### GetGrantedPermissions

`func (o *UserAdminDetailDto) GetGrantedPermissions() []SecurityPermissionDto`

GetGrantedPermissions returns the GrantedPermissions field if non-nil, zero value otherwise.

### GetGrantedPermissionsOk

`func (o *UserAdminDetailDto) GetGrantedPermissionsOk() (*[]SecurityPermissionDto, bool)`

GetGrantedPermissionsOk returns a tuple with the GrantedPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedPermissions

`func (o *UserAdminDetailDto) SetGrantedPermissions(v []SecurityPermissionDto)`

SetGrantedPermissions sets GrantedPermissions field to given value.

### HasGrantedPermissions

`func (o *UserAdminDetailDto) HasGrantedPermissions() bool`

HasGrantedPermissions returns a boolean if a field has been set.

### SetGrantedPermissionsNil

`func (o *UserAdminDetailDto) SetGrantedPermissionsNil(b bool)`

 SetGrantedPermissionsNil sets the value for GrantedPermissions to be an explicit nil

### UnsetGrantedPermissions
`func (o *UserAdminDetailDto) UnsetGrantedPermissions()`

UnsetGrantedPermissions ensures that no value is present for GrantedPermissions, not even an explicit nil
### GetRoleCatalog

`func (o *UserAdminDetailDto) GetRoleCatalog() []SecurityRoleDto`

GetRoleCatalog returns the RoleCatalog field if non-nil, zero value otherwise.

### GetRoleCatalogOk

`func (o *UserAdminDetailDto) GetRoleCatalogOk() (*[]SecurityRoleDto, bool)`

GetRoleCatalogOk returns a tuple with the RoleCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleCatalog

`func (o *UserAdminDetailDto) SetRoleCatalog(v []SecurityRoleDto)`

SetRoleCatalog sets RoleCatalog field to given value.

### HasRoleCatalog

`func (o *UserAdminDetailDto) HasRoleCatalog() bool`

HasRoleCatalog returns a boolean if a field has been set.

### SetRoleCatalogNil

`func (o *UserAdminDetailDto) SetRoleCatalogNil(b bool)`

 SetRoleCatalogNil sets the value for RoleCatalog to be an explicit nil

### UnsetRoleCatalog
`func (o *UserAdminDetailDto) UnsetRoleCatalog()`

UnsetRoleCatalog ensures that no value is present for RoleCatalog, not even an explicit nil
### GetPermissionCatalog

`func (o *UserAdminDetailDto) GetPermissionCatalog() []SecurityPermissionDto`

GetPermissionCatalog returns the PermissionCatalog field if non-nil, zero value otherwise.

### GetPermissionCatalogOk

`func (o *UserAdminDetailDto) GetPermissionCatalogOk() (*[]SecurityPermissionDto, bool)`

GetPermissionCatalogOk returns a tuple with the PermissionCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionCatalog

`func (o *UserAdminDetailDto) SetPermissionCatalog(v []SecurityPermissionDto)`

SetPermissionCatalog sets PermissionCatalog field to given value.

### HasPermissionCatalog

`func (o *UserAdminDetailDto) HasPermissionCatalog() bool`

HasPermissionCatalog returns a boolean if a field has been set.

### SetPermissionCatalogNil

`func (o *UserAdminDetailDto) SetPermissionCatalogNil(b bool)`

 SetPermissionCatalogNil sets the value for PermissionCatalog to be an explicit nil

### UnsetPermissionCatalog
`func (o *UserAdminDetailDto) UnsetPermissionCatalog()`

UnsetPermissionCatalog ensures that no value is present for PermissionCatalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


