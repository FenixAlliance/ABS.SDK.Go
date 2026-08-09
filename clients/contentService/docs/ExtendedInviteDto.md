# ExtendedInviteDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 
**Redeemed** | Pointer to **bool** |  | [optional] 
**RedeemedTimestamp** | Pointer to **NullableTime** |  | [optional] 
**UserEmail** | Pointer to **NullableString** |  | [optional] 
**CreatorEnrollmentId** | Pointer to **NullableString** |  | [optional] 
**RelatedEnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Tenant** | Pointer to [**TenantDto**](TenantDto.md) |  | [optional] 

## Methods

### NewExtendedInviteDto

`func NewExtendedInviteDto() *ExtendedInviteDto`

NewExtendedInviteDto instantiates a new ExtendedInviteDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedInviteDtoWithDefaults

`func NewExtendedInviteDtoWithDefaults() *ExtendedInviteDto`

NewExtendedInviteDtoWithDefaults instantiates a new ExtendedInviteDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtendedInviteDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedInviteDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedInviteDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExtendedInviteDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExtendedInviteDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExtendedInviteDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ExtendedInviteDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedInviteDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedInviteDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtendedInviteDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ExtendedInviteDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ExtendedInviteDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *ExtendedInviteDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExtendedInviteDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExtendedInviteDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ExtendedInviteDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ExtendedInviteDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ExtendedInviteDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRevoked

`func (o *ExtendedInviteDto) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *ExtendedInviteDto) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *ExtendedInviteDto) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *ExtendedInviteDto) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetRedeemed

`func (o *ExtendedInviteDto) GetRedeemed() bool`

GetRedeemed returns the Redeemed field if non-nil, zero value otherwise.

### GetRedeemedOk

`func (o *ExtendedInviteDto) GetRedeemedOk() (*bool, bool)`

GetRedeemedOk returns a tuple with the Redeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemed

`func (o *ExtendedInviteDto) SetRedeemed(v bool)`

SetRedeemed sets Redeemed field to given value.

### HasRedeemed

`func (o *ExtendedInviteDto) HasRedeemed() bool`

HasRedeemed returns a boolean if a field has been set.

### GetRedeemedTimestamp

`func (o *ExtendedInviteDto) GetRedeemedTimestamp() time.Time`

GetRedeemedTimestamp returns the RedeemedTimestamp field if non-nil, zero value otherwise.

### GetRedeemedTimestampOk

`func (o *ExtendedInviteDto) GetRedeemedTimestampOk() (*time.Time, bool)`

GetRedeemedTimestampOk returns a tuple with the RedeemedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedTimestamp

`func (o *ExtendedInviteDto) SetRedeemedTimestamp(v time.Time)`

SetRedeemedTimestamp sets RedeemedTimestamp field to given value.

### HasRedeemedTimestamp

`func (o *ExtendedInviteDto) HasRedeemedTimestamp() bool`

HasRedeemedTimestamp returns a boolean if a field has been set.

### SetRedeemedTimestampNil

`func (o *ExtendedInviteDto) SetRedeemedTimestampNil(b bool)`

 SetRedeemedTimestampNil sets the value for RedeemedTimestamp to be an explicit nil

### UnsetRedeemedTimestamp
`func (o *ExtendedInviteDto) UnsetRedeemedTimestamp()`

UnsetRedeemedTimestamp ensures that no value is present for RedeemedTimestamp, not even an explicit nil
### GetUserEmail

`func (o *ExtendedInviteDto) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ExtendedInviteDto) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ExtendedInviteDto) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *ExtendedInviteDto) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *ExtendedInviteDto) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *ExtendedInviteDto) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetCreatorEnrollmentId

`func (o *ExtendedInviteDto) GetCreatorEnrollmentId() string`

GetCreatorEnrollmentId returns the CreatorEnrollmentId field if non-nil, zero value otherwise.

### GetCreatorEnrollmentIdOk

`func (o *ExtendedInviteDto) GetCreatorEnrollmentIdOk() (*string, bool)`

GetCreatorEnrollmentIdOk returns a tuple with the CreatorEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEnrollmentId

`func (o *ExtendedInviteDto) SetCreatorEnrollmentId(v string)`

SetCreatorEnrollmentId sets CreatorEnrollmentId field to given value.

### HasCreatorEnrollmentId

`func (o *ExtendedInviteDto) HasCreatorEnrollmentId() bool`

HasCreatorEnrollmentId returns a boolean if a field has been set.

### SetCreatorEnrollmentIdNil

`func (o *ExtendedInviteDto) SetCreatorEnrollmentIdNil(b bool)`

 SetCreatorEnrollmentIdNil sets the value for CreatorEnrollmentId to be an explicit nil

### UnsetCreatorEnrollmentId
`func (o *ExtendedInviteDto) UnsetCreatorEnrollmentId()`

UnsetCreatorEnrollmentId ensures that no value is present for CreatorEnrollmentId, not even an explicit nil
### GetRelatedEnrollmentId

`func (o *ExtendedInviteDto) GetRelatedEnrollmentId() string`

GetRelatedEnrollmentId returns the RelatedEnrollmentId field if non-nil, zero value otherwise.

### GetRelatedEnrollmentIdOk

`func (o *ExtendedInviteDto) GetRelatedEnrollmentIdOk() (*string, bool)`

GetRelatedEnrollmentIdOk returns a tuple with the RelatedEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEnrollmentId

`func (o *ExtendedInviteDto) SetRelatedEnrollmentId(v string)`

SetRelatedEnrollmentId sets RelatedEnrollmentId field to given value.

### HasRelatedEnrollmentId

`func (o *ExtendedInviteDto) HasRelatedEnrollmentId() bool`

HasRelatedEnrollmentId returns a boolean if a field has been set.

### SetRelatedEnrollmentIdNil

`func (o *ExtendedInviteDto) SetRelatedEnrollmentIdNil(b bool)`

 SetRelatedEnrollmentIdNil sets the value for RelatedEnrollmentId to be an explicit nil

### UnsetRelatedEnrollmentId
`func (o *ExtendedInviteDto) UnsetRelatedEnrollmentId()`

UnsetRelatedEnrollmentId ensures that no value is present for RelatedEnrollmentId, not even an explicit nil
### GetTenant

`func (o *ExtendedInviteDto) GetTenant() TenantDto`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ExtendedInviteDto) GetTenantOk() (*TenantDto, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ExtendedInviteDto) SetTenant(v TenantDto)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ExtendedInviteDto) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


