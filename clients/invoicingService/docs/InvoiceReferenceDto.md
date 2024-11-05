# InvoiceReferenceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ReferralInvoiceId** | Pointer to **NullableString** |  | [optional] 
**ReferencedInvoiceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceReferenceDto

`func NewInvoiceReferenceDto() *InvoiceReferenceDto`

NewInvoiceReferenceDto instantiates a new InvoiceReferenceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReferenceDtoWithDefaults

`func NewInvoiceReferenceDtoWithDefaults() *InvoiceReferenceDto`

NewInvoiceReferenceDtoWithDefaults instantiates a new InvoiceReferenceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceReferenceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceReferenceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceReferenceDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceReferenceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InvoiceReferenceDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvoiceReferenceDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *InvoiceReferenceDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceReferenceDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceReferenceDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceReferenceDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InvoiceReferenceDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InvoiceReferenceDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *InvoiceReferenceDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceReferenceDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceReferenceDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceReferenceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceReferenceDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceReferenceDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceReferenceDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceReferenceDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceReferenceDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceReferenceDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceReferenceDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceReferenceDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetReferralInvoiceId

`func (o *InvoiceReferenceDto) GetReferralInvoiceId() string`

GetReferralInvoiceId returns the ReferralInvoiceId field if non-nil, zero value otherwise.

### GetReferralInvoiceIdOk

`func (o *InvoiceReferenceDto) GetReferralInvoiceIdOk() (*string, bool)`

GetReferralInvoiceIdOk returns a tuple with the ReferralInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralInvoiceId

`func (o *InvoiceReferenceDto) SetReferralInvoiceId(v string)`

SetReferralInvoiceId sets ReferralInvoiceId field to given value.

### HasReferralInvoiceId

`func (o *InvoiceReferenceDto) HasReferralInvoiceId() bool`

HasReferralInvoiceId returns a boolean if a field has been set.

### SetReferralInvoiceIdNil

`func (o *InvoiceReferenceDto) SetReferralInvoiceIdNil(b bool)`

 SetReferralInvoiceIdNil sets the value for ReferralInvoiceId to be an explicit nil

### UnsetReferralInvoiceId
`func (o *InvoiceReferenceDto) UnsetReferralInvoiceId()`

UnsetReferralInvoiceId ensures that no value is present for ReferralInvoiceId, not even an explicit nil
### GetReferencedInvoiceId

`func (o *InvoiceReferenceDto) GetReferencedInvoiceId() string`

GetReferencedInvoiceId returns the ReferencedInvoiceId field if non-nil, zero value otherwise.

### GetReferencedInvoiceIdOk

`func (o *InvoiceReferenceDto) GetReferencedInvoiceIdOk() (*string, bool)`

GetReferencedInvoiceIdOk returns a tuple with the ReferencedInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedInvoiceId

`func (o *InvoiceReferenceDto) SetReferencedInvoiceId(v string)`

SetReferencedInvoiceId sets ReferencedInvoiceId field to given value.

### HasReferencedInvoiceId

`func (o *InvoiceReferenceDto) HasReferencedInvoiceId() bool`

HasReferencedInvoiceId returns a boolean if a field has been set.

### SetReferencedInvoiceIdNil

`func (o *InvoiceReferenceDto) SetReferencedInvoiceIdNil(b bool)`

 SetReferencedInvoiceIdNil sets the value for ReferencedInvoiceId to be an explicit nil

### UnsetReferencedInvoiceId
`func (o *InvoiceReferenceDto) UnsetReferencedInvoiceId()`

UnsetReferencedInvoiceId ensures that no value is present for ReferencedInvoiceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


