# ApplicationPrincipalDetailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**PrincipalKind** | Pointer to **string** |  | [optional] 
**PrincipalStatus** | Pointer to **string** |  | [optional] 
**BusinessApplicationId** | Pointer to **NullableString** |  | [optional] 
**BusinessApplicationName** | Pointer to **NullableString** |  | [optional] 
**BusinessApplicationNamespace** | Pointer to **NullableString** |  | [optional] 
**BusinessApplicationDisabled** | Pointer to **bool** |  | [optional] 
**SystemLocked** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentDisabled** | Pointer to **bool** |  | [optional] 
**GrantedPermissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApplicationPrincipalDetailDto

`func NewApplicationPrincipalDetailDto() *ApplicationPrincipalDetailDto`

NewApplicationPrincipalDetailDto instantiates a new ApplicationPrincipalDetailDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPrincipalDetailDtoWithDefaults

`func NewApplicationPrincipalDetailDtoWithDefaults() *ApplicationPrincipalDetailDto`

NewApplicationPrincipalDetailDtoWithDefaults instantiates a new ApplicationPrincipalDetailDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationPrincipalDetailDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationPrincipalDetailDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationPrincipalDetailDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationPrincipalDetailDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ApplicationPrincipalDetailDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ApplicationPrincipalDetailDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ApplicationPrincipalDetailDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationPrincipalDetailDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationPrincipalDetailDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApplicationPrincipalDetailDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ApplicationPrincipalDetailDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ApplicationPrincipalDetailDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDisplayName

`func (o *ApplicationPrincipalDetailDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApplicationPrincipalDetailDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApplicationPrincipalDetailDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApplicationPrincipalDetailDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ApplicationPrincipalDetailDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ApplicationPrincipalDetailDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPrincipalKind

`func (o *ApplicationPrincipalDetailDto) GetPrincipalKind() string`

GetPrincipalKind returns the PrincipalKind field if non-nil, zero value otherwise.

### GetPrincipalKindOk

`func (o *ApplicationPrincipalDetailDto) GetPrincipalKindOk() (*string, bool)`

GetPrincipalKindOk returns a tuple with the PrincipalKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalKind

`func (o *ApplicationPrincipalDetailDto) SetPrincipalKind(v string)`

SetPrincipalKind sets PrincipalKind field to given value.

### HasPrincipalKind

`func (o *ApplicationPrincipalDetailDto) HasPrincipalKind() bool`

HasPrincipalKind returns a boolean if a field has been set.

### GetPrincipalStatus

`func (o *ApplicationPrincipalDetailDto) GetPrincipalStatus() string`

GetPrincipalStatus returns the PrincipalStatus field if non-nil, zero value otherwise.

### GetPrincipalStatusOk

`func (o *ApplicationPrincipalDetailDto) GetPrincipalStatusOk() (*string, bool)`

GetPrincipalStatusOk returns a tuple with the PrincipalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalStatus

`func (o *ApplicationPrincipalDetailDto) SetPrincipalStatus(v string)`

SetPrincipalStatus sets PrincipalStatus field to given value.

### HasPrincipalStatus

`func (o *ApplicationPrincipalDetailDto) HasPrincipalStatus() bool`

HasPrincipalStatus returns a boolean if a field has been set.

### GetBusinessApplicationId

`func (o *ApplicationPrincipalDetailDto) GetBusinessApplicationId() string`

GetBusinessApplicationId returns the BusinessApplicationId field if non-nil, zero value otherwise.

### GetBusinessApplicationIdOk

`func (o *ApplicationPrincipalDetailDto) GetBusinessApplicationIdOk() (*string, bool)`

GetBusinessApplicationIdOk returns a tuple with the BusinessApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessApplicationId

`func (o *ApplicationPrincipalDetailDto) SetBusinessApplicationId(v string)`

SetBusinessApplicationId sets BusinessApplicationId field to given value.

### HasBusinessApplicationId

`func (o *ApplicationPrincipalDetailDto) HasBusinessApplicationId() bool`

HasBusinessApplicationId returns a boolean if a field has been set.

### SetBusinessApplicationIdNil

`func (o *ApplicationPrincipalDetailDto) SetBusinessApplicationIdNil(b bool)`

 SetBusinessApplicationIdNil sets the value for BusinessApplicationId to be an explicit nil

### UnsetBusinessApplicationId
`func (o *ApplicationPrincipalDetailDto) UnsetBusinessApplicationId()`

UnsetBusinessApplicationId ensures that no value is present for BusinessApplicationId, not even an explicit nil
### GetBusinessApplicationName

`func (o *ApplicationPrincipalDetailDto) GetBusinessApplicationName() string`

GetBusinessApplicationName returns the BusinessApplicationName field if non-nil, zero value otherwise.

### GetBusinessApplicationNameOk

`func (o *ApplicationPrincipalDetailDto) GetBusinessApplicationNameOk() (*string, bool)`

GetBusinessApplicationNameOk returns a tuple with the BusinessApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessApplicationName

`func (o *ApplicationPrincipalDetailDto) SetBusinessApplicationName(v string)`

SetBusinessApplicationName sets BusinessApplicationName field to given value.

### HasBusinessApplicationName

`func (o *ApplicationPrincipalDetailDto) HasBusinessApplicationName() bool`

HasBusinessApplicationName returns a boolean if a field has been set.

### SetBusinessApplicationNameNil

`func (o *ApplicationPrincipalDetailDto) SetBusinessApplicationNameNil(b bool)`

 SetBusinessApplicationNameNil sets the value for BusinessApplicationName to be an explicit nil

### UnsetBusinessApplicationName
`func (o *ApplicationPrincipalDetailDto) UnsetBusinessApplicationName()`

UnsetBusinessApplicationName ensures that no value is present for BusinessApplicationName, not even an explicit nil
### GetBusinessApplicationNamespace

`func (o *ApplicationPrincipalDetailDto) GetBusinessApplicationNamespace() string`

GetBusinessApplicationNamespace returns the BusinessApplicationNamespace field if non-nil, zero value otherwise.

### GetBusinessApplicationNamespaceOk

`func (o *ApplicationPrincipalDetailDto) GetBusinessApplicationNamespaceOk() (*string, bool)`

GetBusinessApplicationNamespaceOk returns a tuple with the BusinessApplicationNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessApplicationNamespace

`func (o *ApplicationPrincipalDetailDto) SetBusinessApplicationNamespace(v string)`

SetBusinessApplicationNamespace sets BusinessApplicationNamespace field to given value.

### HasBusinessApplicationNamespace

`func (o *ApplicationPrincipalDetailDto) HasBusinessApplicationNamespace() bool`

HasBusinessApplicationNamespace returns a boolean if a field has been set.

### SetBusinessApplicationNamespaceNil

`func (o *ApplicationPrincipalDetailDto) SetBusinessApplicationNamespaceNil(b bool)`

 SetBusinessApplicationNamespaceNil sets the value for BusinessApplicationNamespace to be an explicit nil

### UnsetBusinessApplicationNamespace
`func (o *ApplicationPrincipalDetailDto) UnsetBusinessApplicationNamespace()`

UnsetBusinessApplicationNamespace ensures that no value is present for BusinessApplicationNamespace, not even an explicit nil
### GetBusinessApplicationDisabled

`func (o *ApplicationPrincipalDetailDto) GetBusinessApplicationDisabled() bool`

GetBusinessApplicationDisabled returns the BusinessApplicationDisabled field if non-nil, zero value otherwise.

### GetBusinessApplicationDisabledOk

`func (o *ApplicationPrincipalDetailDto) GetBusinessApplicationDisabledOk() (*bool, bool)`

GetBusinessApplicationDisabledOk returns a tuple with the BusinessApplicationDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessApplicationDisabled

`func (o *ApplicationPrincipalDetailDto) SetBusinessApplicationDisabled(v bool)`

SetBusinessApplicationDisabled sets BusinessApplicationDisabled field to given value.

### HasBusinessApplicationDisabled

`func (o *ApplicationPrincipalDetailDto) HasBusinessApplicationDisabled() bool`

HasBusinessApplicationDisabled returns a boolean if a field has been set.

### GetSystemLocked

`func (o *ApplicationPrincipalDetailDto) GetSystemLocked() bool`

GetSystemLocked returns the SystemLocked field if non-nil, zero value otherwise.

### GetSystemLockedOk

`func (o *ApplicationPrincipalDetailDto) GetSystemLockedOk() (*bool, bool)`

GetSystemLockedOk returns a tuple with the SystemLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLocked

`func (o *ApplicationPrincipalDetailDto) SetSystemLocked(v bool)`

SetSystemLocked sets SystemLocked field to given value.

### HasSystemLocked

`func (o *ApplicationPrincipalDetailDto) HasSystemLocked() bool`

HasSystemLocked returns a boolean if a field has been set.

### GetTenantId

`func (o *ApplicationPrincipalDetailDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ApplicationPrincipalDetailDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ApplicationPrincipalDetailDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ApplicationPrincipalDetailDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ApplicationPrincipalDetailDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ApplicationPrincipalDetailDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ApplicationPrincipalDetailDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ApplicationPrincipalDetailDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ApplicationPrincipalDetailDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ApplicationPrincipalDetailDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ApplicationPrincipalDetailDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ApplicationPrincipalDetailDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetEnrollmentDisabled

`func (o *ApplicationPrincipalDetailDto) GetEnrollmentDisabled() bool`

GetEnrollmentDisabled returns the EnrollmentDisabled field if non-nil, zero value otherwise.

### GetEnrollmentDisabledOk

`func (o *ApplicationPrincipalDetailDto) GetEnrollmentDisabledOk() (*bool, bool)`

GetEnrollmentDisabledOk returns a tuple with the EnrollmentDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDisabled

`func (o *ApplicationPrincipalDetailDto) SetEnrollmentDisabled(v bool)`

SetEnrollmentDisabled sets EnrollmentDisabled field to given value.

### HasEnrollmentDisabled

`func (o *ApplicationPrincipalDetailDto) HasEnrollmentDisabled() bool`

HasEnrollmentDisabled returns a boolean if a field has been set.

### GetGrantedPermissions

`func (o *ApplicationPrincipalDetailDto) GetGrantedPermissions() []string`

GetGrantedPermissions returns the GrantedPermissions field if non-nil, zero value otherwise.

### GetGrantedPermissionsOk

`func (o *ApplicationPrincipalDetailDto) GetGrantedPermissionsOk() (*[]string, bool)`

GetGrantedPermissionsOk returns a tuple with the GrantedPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedPermissions

`func (o *ApplicationPrincipalDetailDto) SetGrantedPermissions(v []string)`

SetGrantedPermissions sets GrantedPermissions field to given value.

### HasGrantedPermissions

`func (o *ApplicationPrincipalDetailDto) HasGrantedPermissions() bool`

HasGrantedPermissions returns a boolean if a field has been set.

### SetGrantedPermissionsNil

`func (o *ApplicationPrincipalDetailDto) SetGrantedPermissionsNil(b bool)`

 SetGrantedPermissionsNil sets the value for GrantedPermissions to be an explicit nil

### UnsetGrantedPermissions
`func (o *ApplicationPrincipalDetailDto) UnsetGrantedPermissions()`

UnsetGrantedPermissions ensures that no value is present for GrantedPermissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


