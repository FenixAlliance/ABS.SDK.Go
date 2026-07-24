# NewsletterSubscriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**NewsletterId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewsletterSubscriptionDto

`func NewNewsletterSubscriptionDto() *NewsletterSubscriptionDto`

NewNewsletterSubscriptionDto instantiates a new NewsletterSubscriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsletterSubscriptionDtoWithDefaults

`func NewNewsletterSubscriptionDtoWithDefaults() *NewsletterSubscriptionDto`

NewNewsletterSubscriptionDtoWithDefaults instantiates a new NewsletterSubscriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewsletterSubscriptionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewsletterSubscriptionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewsletterSubscriptionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NewsletterSubscriptionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *NewsletterSubscriptionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NewsletterSubscriptionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *NewsletterSubscriptionDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NewsletterSubscriptionDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NewsletterSubscriptionDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NewsletterSubscriptionDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *NewsletterSubscriptionDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *NewsletterSubscriptionDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetEmail

`func (o *NewsletterSubscriptionDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewsletterSubscriptionDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewsletterSubscriptionDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NewsletterSubscriptionDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *NewsletterSubscriptionDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *NewsletterSubscriptionDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetVerified

`func (o *NewsletterSubscriptionDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *NewsletterSubscriptionDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *NewsletterSubscriptionDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *NewsletterSubscriptionDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetContactId

`func (o *NewsletterSubscriptionDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *NewsletterSubscriptionDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *NewsletterSubscriptionDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *NewsletterSubscriptionDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *NewsletterSubscriptionDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *NewsletterSubscriptionDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetUserId

`func (o *NewsletterSubscriptionDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *NewsletterSubscriptionDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *NewsletterSubscriptionDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *NewsletterSubscriptionDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *NewsletterSubscriptionDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *NewsletterSubscriptionDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetNewsletterId

`func (o *NewsletterSubscriptionDto) GetNewsletterId() string`

GetNewsletterId returns the NewsletterId field if non-nil, zero value otherwise.

### GetNewsletterIdOk

`func (o *NewsletterSubscriptionDto) GetNewsletterIdOk() (*string, bool)`

GetNewsletterIdOk returns a tuple with the NewsletterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsletterId

`func (o *NewsletterSubscriptionDto) SetNewsletterId(v string)`

SetNewsletterId sets NewsletterId field to given value.

### HasNewsletterId

`func (o *NewsletterSubscriptionDto) HasNewsletterId() bool`

HasNewsletterId returns a boolean if a field has been set.

### SetNewsletterIdNil

`func (o *NewsletterSubscriptionDto) SetNewsletterIdNil(b bool)`

 SetNewsletterIdNil sets the value for NewsletterId to be an explicit nil

### UnsetNewsletterId
`func (o *NewsletterSubscriptionDto) UnsetNewsletterId()`

UnsetNewsletterId ensures that no value is present for NewsletterId, not even an explicit nil
### GetTenantId

`func (o *NewsletterSubscriptionDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *NewsletterSubscriptionDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *NewsletterSubscriptionDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *NewsletterSubscriptionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *NewsletterSubscriptionDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *NewsletterSubscriptionDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *NewsletterSubscriptionDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *NewsletterSubscriptionDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *NewsletterSubscriptionDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *NewsletterSubscriptionDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *NewsletterSubscriptionDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *NewsletterSubscriptionDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


