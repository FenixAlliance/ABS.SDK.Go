# ApplicationPrincipalDto

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
**SystemLocked** | Pointer to **bool** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentDisabled** | Pointer to **bool** |  | [optional] 
**GrantedPermissionsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewApplicationPrincipalDto

`func NewApplicationPrincipalDto() *ApplicationPrincipalDto`

NewApplicationPrincipalDto instantiates a new ApplicationPrincipalDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPrincipalDtoWithDefaults

`func NewApplicationPrincipalDtoWithDefaults() *ApplicationPrincipalDto`

NewApplicationPrincipalDtoWithDefaults instantiates a new ApplicationPrincipalDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationPrincipalDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationPrincipalDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationPrincipalDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationPrincipalDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ApplicationPrincipalDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ApplicationPrincipalDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ApplicationPrincipalDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationPrincipalDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationPrincipalDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApplicationPrincipalDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ApplicationPrincipalDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ApplicationPrincipalDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetDisplayName

`func (o *ApplicationPrincipalDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApplicationPrincipalDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApplicationPrincipalDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApplicationPrincipalDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ApplicationPrincipalDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ApplicationPrincipalDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPrincipalKind

`func (o *ApplicationPrincipalDto) GetPrincipalKind() string`

GetPrincipalKind returns the PrincipalKind field if non-nil, zero value otherwise.

### GetPrincipalKindOk

`func (o *ApplicationPrincipalDto) GetPrincipalKindOk() (*string, bool)`

GetPrincipalKindOk returns a tuple with the PrincipalKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalKind

`func (o *ApplicationPrincipalDto) SetPrincipalKind(v string)`

SetPrincipalKind sets PrincipalKind field to given value.

### HasPrincipalKind

`func (o *ApplicationPrincipalDto) HasPrincipalKind() bool`

HasPrincipalKind returns a boolean if a field has been set.

### GetPrincipalStatus

`func (o *ApplicationPrincipalDto) GetPrincipalStatus() string`

GetPrincipalStatus returns the PrincipalStatus field if non-nil, zero value otherwise.

### GetPrincipalStatusOk

`func (o *ApplicationPrincipalDto) GetPrincipalStatusOk() (*string, bool)`

GetPrincipalStatusOk returns a tuple with the PrincipalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalStatus

`func (o *ApplicationPrincipalDto) SetPrincipalStatus(v string)`

SetPrincipalStatus sets PrincipalStatus field to given value.

### HasPrincipalStatus

`func (o *ApplicationPrincipalDto) HasPrincipalStatus() bool`

HasPrincipalStatus returns a boolean if a field has been set.

### GetBusinessApplicationId

`func (o *ApplicationPrincipalDto) GetBusinessApplicationId() string`

GetBusinessApplicationId returns the BusinessApplicationId field if non-nil, zero value otherwise.

### GetBusinessApplicationIdOk

`func (o *ApplicationPrincipalDto) GetBusinessApplicationIdOk() (*string, bool)`

GetBusinessApplicationIdOk returns a tuple with the BusinessApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessApplicationId

`func (o *ApplicationPrincipalDto) SetBusinessApplicationId(v string)`

SetBusinessApplicationId sets BusinessApplicationId field to given value.

### HasBusinessApplicationId

`func (o *ApplicationPrincipalDto) HasBusinessApplicationId() bool`

HasBusinessApplicationId returns a boolean if a field has been set.

### SetBusinessApplicationIdNil

`func (o *ApplicationPrincipalDto) SetBusinessApplicationIdNil(b bool)`

 SetBusinessApplicationIdNil sets the value for BusinessApplicationId to be an explicit nil

### UnsetBusinessApplicationId
`func (o *ApplicationPrincipalDto) UnsetBusinessApplicationId()`

UnsetBusinessApplicationId ensures that no value is present for BusinessApplicationId, not even an explicit nil
### GetBusinessApplicationName

`func (o *ApplicationPrincipalDto) GetBusinessApplicationName() string`

GetBusinessApplicationName returns the BusinessApplicationName field if non-nil, zero value otherwise.

### GetBusinessApplicationNameOk

`func (o *ApplicationPrincipalDto) GetBusinessApplicationNameOk() (*string, bool)`

GetBusinessApplicationNameOk returns a tuple with the BusinessApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessApplicationName

`func (o *ApplicationPrincipalDto) SetBusinessApplicationName(v string)`

SetBusinessApplicationName sets BusinessApplicationName field to given value.

### HasBusinessApplicationName

`func (o *ApplicationPrincipalDto) HasBusinessApplicationName() bool`

HasBusinessApplicationName returns a boolean if a field has been set.

### SetBusinessApplicationNameNil

`func (o *ApplicationPrincipalDto) SetBusinessApplicationNameNil(b bool)`

 SetBusinessApplicationNameNil sets the value for BusinessApplicationName to be an explicit nil

### UnsetBusinessApplicationName
`func (o *ApplicationPrincipalDto) UnsetBusinessApplicationName()`

UnsetBusinessApplicationName ensures that no value is present for BusinessApplicationName, not even an explicit nil
### GetSystemLocked

`func (o *ApplicationPrincipalDto) GetSystemLocked() bool`

GetSystemLocked returns the SystemLocked field if non-nil, zero value otherwise.

### GetSystemLockedOk

`func (o *ApplicationPrincipalDto) GetSystemLockedOk() (*bool, bool)`

GetSystemLockedOk returns a tuple with the SystemLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLocked

`func (o *ApplicationPrincipalDto) SetSystemLocked(v bool)`

SetSystemLocked sets SystemLocked field to given value.

### HasSystemLocked

`func (o *ApplicationPrincipalDto) HasSystemLocked() bool`

HasSystemLocked returns a boolean if a field has been set.

### GetTenantId

`func (o *ApplicationPrincipalDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ApplicationPrincipalDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ApplicationPrincipalDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ApplicationPrincipalDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ApplicationPrincipalDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ApplicationPrincipalDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ApplicationPrincipalDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ApplicationPrincipalDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ApplicationPrincipalDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ApplicationPrincipalDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ApplicationPrincipalDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ApplicationPrincipalDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetEnrollmentDisabled

`func (o *ApplicationPrincipalDto) GetEnrollmentDisabled() bool`

GetEnrollmentDisabled returns the EnrollmentDisabled field if non-nil, zero value otherwise.

### GetEnrollmentDisabledOk

`func (o *ApplicationPrincipalDto) GetEnrollmentDisabledOk() (*bool, bool)`

GetEnrollmentDisabledOk returns a tuple with the EnrollmentDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentDisabled

`func (o *ApplicationPrincipalDto) SetEnrollmentDisabled(v bool)`

SetEnrollmentDisabled sets EnrollmentDisabled field to given value.

### HasEnrollmentDisabled

`func (o *ApplicationPrincipalDto) HasEnrollmentDisabled() bool`

HasEnrollmentDisabled returns a boolean if a field has been set.

### GetGrantedPermissionsCount

`func (o *ApplicationPrincipalDto) GetGrantedPermissionsCount() int32`

GetGrantedPermissionsCount returns the GrantedPermissionsCount field if non-nil, zero value otherwise.

### GetGrantedPermissionsCountOk

`func (o *ApplicationPrincipalDto) GetGrantedPermissionsCountOk() (*int32, bool)`

GetGrantedPermissionsCountOk returns a tuple with the GrantedPermissionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedPermissionsCount

`func (o *ApplicationPrincipalDto) SetGrantedPermissionsCount(v int32)`

SetGrantedPermissionsCount sets GrantedPermissionsCount field to given value.

### HasGrantedPermissionsCount

`func (o *ApplicationPrincipalDto) HasGrantedPermissionsCount() bool`

HasGrantedPermissionsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


