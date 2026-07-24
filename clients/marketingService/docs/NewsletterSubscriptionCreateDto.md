# NewsletterSubscriptionCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 
**NewsletterId** | Pointer to **NullableString** |  | [optional] 
**ContactId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNewsletterSubscriptionCreateDto

`func NewNewsletterSubscriptionCreateDto() *NewsletterSubscriptionCreateDto`

NewNewsletterSubscriptionCreateDto instantiates a new NewsletterSubscriptionCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsletterSubscriptionCreateDtoWithDefaults

`func NewNewsletterSubscriptionCreateDtoWithDefaults() *NewsletterSubscriptionCreateDto`

NewNewsletterSubscriptionCreateDtoWithDefaults instantiates a new NewsletterSubscriptionCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewsletterSubscriptionCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewsletterSubscriptionCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewsletterSubscriptionCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NewsletterSubscriptionCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *NewsletterSubscriptionCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *NewsletterSubscriptionCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *NewsletterSubscriptionCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *NewsletterSubscriptionCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEmail

`func (o *NewsletterSubscriptionCreateDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NewsletterSubscriptionCreateDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NewsletterSubscriptionCreateDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NewsletterSubscriptionCreateDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *NewsletterSubscriptionCreateDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *NewsletterSubscriptionCreateDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetVerified

`func (o *NewsletterSubscriptionCreateDto) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *NewsletterSubscriptionCreateDto) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *NewsletterSubscriptionCreateDto) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *NewsletterSubscriptionCreateDto) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetNewsletterId

`func (o *NewsletterSubscriptionCreateDto) GetNewsletterId() string`

GetNewsletterId returns the NewsletterId field if non-nil, zero value otherwise.

### GetNewsletterIdOk

`func (o *NewsletterSubscriptionCreateDto) GetNewsletterIdOk() (*string, bool)`

GetNewsletterIdOk returns a tuple with the NewsletterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsletterId

`func (o *NewsletterSubscriptionCreateDto) SetNewsletterId(v string)`

SetNewsletterId sets NewsletterId field to given value.

### HasNewsletterId

`func (o *NewsletterSubscriptionCreateDto) HasNewsletterId() bool`

HasNewsletterId returns a boolean if a field has been set.

### SetNewsletterIdNil

`func (o *NewsletterSubscriptionCreateDto) SetNewsletterIdNil(b bool)`

 SetNewsletterIdNil sets the value for NewsletterId to be an explicit nil

### UnsetNewsletterId
`func (o *NewsletterSubscriptionCreateDto) UnsetNewsletterId()`

UnsetNewsletterId ensures that no value is present for NewsletterId, not even an explicit nil
### GetContactId

`func (o *NewsletterSubscriptionCreateDto) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *NewsletterSubscriptionCreateDto) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *NewsletterSubscriptionCreateDto) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *NewsletterSubscriptionCreateDto) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### SetContactIdNil

`func (o *NewsletterSubscriptionCreateDto) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *NewsletterSubscriptionCreateDto) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


