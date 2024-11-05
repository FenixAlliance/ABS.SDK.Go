# InvoiceReferenceCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**ReferralInvoiceId** | Pointer to **NullableString** |  | [optional] 
**ReferencedInvoiceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceReferenceCreateDto

`func NewInvoiceReferenceCreateDto() *InvoiceReferenceCreateDto`

NewInvoiceReferenceCreateDto instantiates a new InvoiceReferenceCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReferenceCreateDtoWithDefaults

`func NewInvoiceReferenceCreateDtoWithDefaults() *InvoiceReferenceCreateDto`

NewInvoiceReferenceCreateDtoWithDefaults instantiates a new InvoiceReferenceCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceReferenceCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceReferenceCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceReferenceCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceReferenceCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *InvoiceReferenceCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceReferenceCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceReferenceCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceReferenceCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *InvoiceReferenceCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceReferenceCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceReferenceCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceReferenceCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceReferenceCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceReferenceCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceReferenceCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceReferenceCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceReferenceCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceReferenceCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceReferenceCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceReferenceCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetReferralInvoiceId

`func (o *InvoiceReferenceCreateDto) GetReferralInvoiceId() string`

GetReferralInvoiceId returns the ReferralInvoiceId field if non-nil, zero value otherwise.

### GetReferralInvoiceIdOk

`func (o *InvoiceReferenceCreateDto) GetReferralInvoiceIdOk() (*string, bool)`

GetReferralInvoiceIdOk returns a tuple with the ReferralInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralInvoiceId

`func (o *InvoiceReferenceCreateDto) SetReferralInvoiceId(v string)`

SetReferralInvoiceId sets ReferralInvoiceId field to given value.

### HasReferralInvoiceId

`func (o *InvoiceReferenceCreateDto) HasReferralInvoiceId() bool`

HasReferralInvoiceId returns a boolean if a field has been set.

### SetReferralInvoiceIdNil

`func (o *InvoiceReferenceCreateDto) SetReferralInvoiceIdNil(b bool)`

 SetReferralInvoiceIdNil sets the value for ReferralInvoiceId to be an explicit nil

### UnsetReferralInvoiceId
`func (o *InvoiceReferenceCreateDto) UnsetReferralInvoiceId()`

UnsetReferralInvoiceId ensures that no value is present for ReferralInvoiceId, not even an explicit nil
### GetReferencedInvoiceId

`func (o *InvoiceReferenceCreateDto) GetReferencedInvoiceId() string`

GetReferencedInvoiceId returns the ReferencedInvoiceId field if non-nil, zero value otherwise.

### GetReferencedInvoiceIdOk

`func (o *InvoiceReferenceCreateDto) GetReferencedInvoiceIdOk() (*string, bool)`

GetReferencedInvoiceIdOk returns a tuple with the ReferencedInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedInvoiceId

`func (o *InvoiceReferenceCreateDto) SetReferencedInvoiceId(v string)`

SetReferencedInvoiceId sets ReferencedInvoiceId field to given value.

### HasReferencedInvoiceId

`func (o *InvoiceReferenceCreateDto) HasReferencedInvoiceId() bool`

HasReferencedInvoiceId returns a boolean if a field has been set.

### SetReferencedInvoiceIdNil

`func (o *InvoiceReferenceCreateDto) SetReferencedInvoiceIdNil(b bool)`

 SetReferencedInvoiceIdNil sets the value for ReferencedInvoiceId to be an explicit nil

### UnsetReferencedInvoiceId
`func (o *InvoiceReferenceCreateDto) UnsetReferencedInvoiceId()`

UnsetReferencedInvoiceId ensures that no value is present for ReferencedInvoiceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


