# ContactEmailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**VerifiedTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Contact** | Pointer to [**ContactDto**](ContactDto.md) |  | [optional] 

## Methods

### NewContactEmailDto

`func NewContactEmailDto() *ContactEmailDto`

NewContactEmailDto instantiates a new ContactEmailDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactEmailDtoWithDefaults

`func NewContactEmailDtoWithDefaults() *ContactEmailDto`

NewContactEmailDtoWithDefaults instantiates a new ContactEmailDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactEmailDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactEmailDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactEmailDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactEmailDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ContactEmailDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ContactEmailDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ContactEmailDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContactEmailDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContactEmailDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContactEmailDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ContactEmailDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ContactEmailDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetContactId

`func (o *ContactEmailDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ContactEmailDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ContactEmailDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ContactEmailDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *ContactEmailDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *ContactEmailDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetTenantId

`func (o *ContactEmailDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ContactEmailDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ContactEmailDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ContactEmailDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ContactEmailDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ContactEmailDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *ContactEmailDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *ContactEmailDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *ContactEmailDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *ContactEmailDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *ContactEmailDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *ContactEmailDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetAddress

`func (o *ContactEmailDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContactEmailDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContactEmailDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContactEmailDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *ContactEmailDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *ContactEmailDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetLabel

`func (o *ContactEmailDto) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ContactEmailDto) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ContactEmailDto) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ContactEmailDto) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ContactEmailDto) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ContactEmailDto) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetIsPrimary

`func (o *ContactEmailDto) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *ContactEmailDto) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *ContactEmailDto) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *ContactEmailDto) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsVerified

`func (o *ContactEmailDto) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *ContactEmailDto) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *ContactEmailDto) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *ContactEmailDto) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetVerifiedTimestamp

`func (o *ContactEmailDto) GetVerifiedTimestamp() time.Time`

GetVerifiedTimestamp returns the VerifiedTimestamp field if non-nil, zero value otherwise.

### GetVerifiedTimestampOk

`func (o *ContactEmailDto) GetVerifiedTimestampOk() (*time.Time, bool)`

GetVerifiedTimestampOk returns a tuple with the VerifiedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedTimestamp

`func (o *ContactEmailDto) SetVerifiedTimestamp(v time.Time)`

SetVerifiedTimestamp sets VerifiedTimestamp field to given value.

### HasVerifiedTimestamp

`func (o *ContactEmailDto) HasVerifiedTimestamp() bool`

HasVerifiedTimestamp returns a boolean if a field has been set.

### SetVerifiedTimestampNil

`func (o *ContactEmailDto) SetVerifiedTimestampNil(b bool)`

 SetVerifiedTimestampNil sets the value for VerifiedTimestamp to be an explicit nil

### UnsetVerifiedTimestamp
`func (o *ContactEmailDto) UnsetVerifiedTimestamp()`

UnsetVerifiedTimestamp ensures that no value is present for VerifiedTimestamp, not even an explicit nil
### GetContact

`func (o *ContactEmailDto) GetContact() ContactDto`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ContactEmailDto) GetContactOk() (*ContactDto, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ContactEmailDto) SetContact(v ContactDto)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ContactEmailDto) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


