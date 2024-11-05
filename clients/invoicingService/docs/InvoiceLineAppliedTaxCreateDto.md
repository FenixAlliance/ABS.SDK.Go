# InvoiceLineAppliedTaxCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Timestamp** | Pointer to **time.Time** |  | [optional] [readonly] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**TaxPolicyId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceLineAppliedTaxCreateDto

`func NewInvoiceLineAppliedTaxCreateDto() *InvoiceLineAppliedTaxCreateDto`

NewInvoiceLineAppliedTaxCreateDto instantiates a new InvoiceLineAppliedTaxCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineAppliedTaxCreateDtoWithDefaults

`func NewInvoiceLineAppliedTaxCreateDtoWithDefaults() *InvoiceLineAppliedTaxCreateDto`

NewInvoiceLineAppliedTaxCreateDtoWithDefaults instantiates a new InvoiceLineAppliedTaxCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceLineAppliedTaxCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineAppliedTaxCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineAppliedTaxCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceLineAppliedTaxCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *InvoiceLineAppliedTaxCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceLineAppliedTaxCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceLineAppliedTaxCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceLineAppliedTaxCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTenantId

`func (o *InvoiceLineAppliedTaxCreateDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceLineAppliedTaxCreateDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceLineAppliedTaxCreateDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceLineAppliedTaxCreateDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceLineAppliedTaxCreateDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceLineAppliedTaxCreateDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetInvoiceId

`func (o *InvoiceLineAppliedTaxCreateDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceLineAppliedTaxCreateDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceLineAppliedTaxCreateDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceLineAppliedTaxCreateDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *InvoiceLineAppliedTaxCreateDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *InvoiceLineAppliedTaxCreateDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetTaxPolicyId

`func (o *InvoiceLineAppliedTaxCreateDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *InvoiceLineAppliedTaxCreateDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *InvoiceLineAppliedTaxCreateDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *InvoiceLineAppliedTaxCreateDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *InvoiceLineAppliedTaxCreateDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *InvoiceLineAppliedTaxCreateDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceLineAppliedTaxCreateDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceLineAppliedTaxCreateDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceLineAppliedTaxCreateDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceLineAppliedTaxCreateDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceLineAppliedTaxCreateDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceLineAppliedTaxCreateDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


