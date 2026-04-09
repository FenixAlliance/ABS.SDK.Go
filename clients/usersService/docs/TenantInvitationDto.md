# TenantInvitationDto

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

## Methods

### NewTenantInvitationDto

`func NewTenantInvitationDto() *TenantInvitationDto`

NewTenantInvitationDto instantiates a new TenantInvitationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantInvitationDtoWithDefaults

`func NewTenantInvitationDtoWithDefaults() *TenantInvitationDto`

NewTenantInvitationDtoWithDefaults instantiates a new TenantInvitationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantInvitationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantInvitationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantInvitationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantInvitationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantInvitationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantInvitationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *TenantInvitationDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TenantInvitationDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TenantInvitationDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TenantInvitationDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TenantInvitationDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TenantInvitationDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *TenantInvitationDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantInvitationDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantInvitationDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantInvitationDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *TenantInvitationDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *TenantInvitationDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetRevoked

`func (o *TenantInvitationDto) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *TenantInvitationDto) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *TenantInvitationDto) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *TenantInvitationDto) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetRedeemed

`func (o *TenantInvitationDto) GetRedeemed() bool`

GetRedeemed returns the Redeemed field if non-nil, zero value otherwise.

### GetRedeemedOk

`func (o *TenantInvitationDto) GetRedeemedOk() (*bool, bool)`

GetRedeemedOk returns a tuple with the Redeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemed

`func (o *TenantInvitationDto) SetRedeemed(v bool)`

SetRedeemed sets Redeemed field to given value.

### HasRedeemed

`func (o *TenantInvitationDto) HasRedeemed() bool`

HasRedeemed returns a boolean if a field has been set.

### GetRedeemedTimestamp

`func (o *TenantInvitationDto) GetRedeemedTimestamp() time.Time`

GetRedeemedTimestamp returns the RedeemedTimestamp field if non-nil, zero value otherwise.

### GetRedeemedTimestampOk

`func (o *TenantInvitationDto) GetRedeemedTimestampOk() (*time.Time, bool)`

GetRedeemedTimestampOk returns a tuple with the RedeemedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedTimestamp

`func (o *TenantInvitationDto) SetRedeemedTimestamp(v time.Time)`

SetRedeemedTimestamp sets RedeemedTimestamp field to given value.

### HasRedeemedTimestamp

`func (o *TenantInvitationDto) HasRedeemedTimestamp() bool`

HasRedeemedTimestamp returns a boolean if a field has been set.

### SetRedeemedTimestampNil

`func (o *TenantInvitationDto) SetRedeemedTimestampNil(b bool)`

 SetRedeemedTimestampNil sets the value for RedeemedTimestamp to be an explicit nil

### UnsetRedeemedTimestamp
`func (o *TenantInvitationDto) UnsetRedeemedTimestamp()`

UnsetRedeemedTimestamp ensures that no value is present for RedeemedTimestamp, not even an explicit nil
### GetUserEmail

`func (o *TenantInvitationDto) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *TenantInvitationDto) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *TenantInvitationDto) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *TenantInvitationDto) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *TenantInvitationDto) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *TenantInvitationDto) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetCreatorEnrollmentId

`func (o *TenantInvitationDto) GetCreatorEnrollmentId() string`

GetCreatorEnrollmentId returns the CreatorEnrollmentId field if non-nil, zero value otherwise.

### GetCreatorEnrollmentIdOk

`func (o *TenantInvitationDto) GetCreatorEnrollmentIdOk() (*string, bool)`

GetCreatorEnrollmentIdOk returns a tuple with the CreatorEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEnrollmentId

`func (o *TenantInvitationDto) SetCreatorEnrollmentId(v string)`

SetCreatorEnrollmentId sets CreatorEnrollmentId field to given value.

### HasCreatorEnrollmentId

`func (o *TenantInvitationDto) HasCreatorEnrollmentId() bool`

HasCreatorEnrollmentId returns a boolean if a field has been set.

### SetCreatorEnrollmentIdNil

`func (o *TenantInvitationDto) SetCreatorEnrollmentIdNil(b bool)`

 SetCreatorEnrollmentIdNil sets the value for CreatorEnrollmentId to be an explicit nil

### UnsetCreatorEnrollmentId
`func (o *TenantInvitationDto) UnsetCreatorEnrollmentId()`

UnsetCreatorEnrollmentId ensures that no value is present for CreatorEnrollmentId, not even an explicit nil
### GetRelatedEnrollmentId

`func (o *TenantInvitationDto) GetRelatedEnrollmentId() string`

GetRelatedEnrollmentId returns the RelatedEnrollmentId field if non-nil, zero value otherwise.

### GetRelatedEnrollmentIdOk

`func (o *TenantInvitationDto) GetRelatedEnrollmentIdOk() (*string, bool)`

GetRelatedEnrollmentIdOk returns a tuple with the RelatedEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEnrollmentId

`func (o *TenantInvitationDto) SetRelatedEnrollmentId(v string)`

SetRelatedEnrollmentId sets RelatedEnrollmentId field to given value.

### HasRelatedEnrollmentId

`func (o *TenantInvitationDto) HasRelatedEnrollmentId() bool`

HasRelatedEnrollmentId returns a boolean if a field has been set.

### SetRelatedEnrollmentIdNil

`func (o *TenantInvitationDto) SetRelatedEnrollmentIdNil(b bool)`

 SetRelatedEnrollmentIdNil sets the value for RelatedEnrollmentId to be an explicit nil

### UnsetRelatedEnrollmentId
`func (o *TenantInvitationDto) UnsetRelatedEnrollmentId()`

UnsetRelatedEnrollmentId ensures that no value is present for RelatedEnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


