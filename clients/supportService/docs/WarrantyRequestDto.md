# WarrantyRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**ApprovedTimestamp** | Pointer to **time.Time** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**SupportEntitlementId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**WarrantyPolicyId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWarrantyRequestDto

`func NewWarrantyRequestDto() *WarrantyRequestDto`

NewWarrantyRequestDto instantiates a new WarrantyRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarrantyRequestDtoWithDefaults

`func NewWarrantyRequestDtoWithDefaults() *WarrantyRequestDto`

NewWarrantyRequestDtoWithDefaults instantiates a new WarrantyRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarrantyRequestDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarrantyRequestDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarrantyRequestDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WarrantyRequestDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WarrantyRequestDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WarrantyRequestDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *WarrantyRequestDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WarrantyRequestDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WarrantyRequestDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WarrantyRequestDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTitle

`func (o *WarrantyRequestDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WarrantyRequestDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WarrantyRequestDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WarrantyRequestDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WarrantyRequestDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WarrantyRequestDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *WarrantyRequestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WarrantyRequestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WarrantyRequestDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WarrantyRequestDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WarrantyRequestDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WarrantyRequestDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetApproved

`func (o *WarrantyRequestDto) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *WarrantyRequestDto) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *WarrantyRequestDto) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *WarrantyRequestDto) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetApprovedTimestamp

`func (o *WarrantyRequestDto) GetApprovedTimestamp() time.Time`

GetApprovedTimestamp returns the ApprovedTimestamp field if non-nil, zero value otherwise.

### GetApprovedTimestampOk

`func (o *WarrantyRequestDto) GetApprovedTimestampOk() (*time.Time, bool)`

GetApprovedTimestampOk returns a tuple with the ApprovedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedTimestamp

`func (o *WarrantyRequestDto) SetApprovedTimestamp(v time.Time)`

SetApprovedTimestamp sets ApprovedTimestamp field to given value.

### HasApprovedTimestamp

`func (o *WarrantyRequestDto) HasApprovedTimestamp() bool`

HasApprovedTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *WarrantyRequestDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WarrantyRequestDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WarrantyRequestDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *WarrantyRequestDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *WarrantyRequestDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *WarrantyRequestDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *WarrantyRequestDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *WarrantyRequestDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *WarrantyRequestDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *WarrantyRequestDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *WarrantyRequestDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *WarrantyRequestDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetSupportEntitlementId

`func (o *WarrantyRequestDto) GetSupportEntitlementId() string`

GetSupportEntitlementId returns the SupportEntitlementId field if non-nil, zero value otherwise.

### GetSupportEntitlementIdOk

`func (o *WarrantyRequestDto) GetSupportEntitlementIdOk() (*string, bool)`

GetSupportEntitlementIdOk returns a tuple with the SupportEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEntitlementId

`func (o *WarrantyRequestDto) SetSupportEntitlementId(v string)`

SetSupportEntitlementId sets SupportEntitlementId field to given value.

### HasSupportEntitlementId

`func (o *WarrantyRequestDto) HasSupportEntitlementId() bool`

HasSupportEntitlementId returns a boolean if a field has been set.

### SetSupportEntitlementIdNil

`func (o *WarrantyRequestDto) SetSupportEntitlementIdNil(b bool)`

 SetSupportEntitlementIdNil sets the value for SupportEntitlementId to be an explicit nil

### UnsetSupportEntitlementId
`func (o *WarrantyRequestDto) UnsetSupportEntitlementId()`

UnsetSupportEntitlementId ensures that no value is present for SupportEntitlementId, not even an explicit nil
### GetContactId

`func (o *WarrantyRequestDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *WarrantyRequestDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *WarrantyRequestDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *WarrantyRequestDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *WarrantyRequestDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *WarrantyRequestDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetUserId

`func (o *WarrantyRequestDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WarrantyRequestDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WarrantyRequestDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WarrantyRequestDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *WarrantyRequestDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *WarrantyRequestDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWarrantyPolicyId

`func (o *WarrantyRequestDto) GetWarrantyPolicyId() string`

GetWarrantyPolicyId returns the WarrantyPolicyId field if non-nil, zero value otherwise.

### GetWarrantyPolicyIdOk

`func (o *WarrantyRequestDto) GetWarrantyPolicyIdOk() (*string, bool)`

GetWarrantyPolicyIdOk returns a tuple with the WarrantyPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyPolicyId

`func (o *WarrantyRequestDto) SetWarrantyPolicyId(v string)`

SetWarrantyPolicyId sets WarrantyPolicyId field to given value.

### HasWarrantyPolicyId

`func (o *WarrantyRequestDto) HasWarrantyPolicyId() bool`

HasWarrantyPolicyId returns a boolean if a field has been set.

### SetWarrantyPolicyIdNil

`func (o *WarrantyRequestDto) SetWarrantyPolicyIdNil(b bool)`

 SetWarrantyPolicyIdNil sets the value for WarrantyPolicyId to be an explicit nil

### UnsetWarrantyPolicyId
`func (o *WarrantyRequestDto) UnsetWarrantyPolicyId()`

UnsetWarrantyPolicyId ensures that no value is present for WarrantyPolicyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


