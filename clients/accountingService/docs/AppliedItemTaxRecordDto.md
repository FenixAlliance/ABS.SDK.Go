# AppliedItemTaxRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**TaxPolicyId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**TaxAmountInUSD** | Pointer to **float64** |  | [optional] 
**TaxBaseAmountInUSD** | Pointer to **float64** |  | [optional] 
**BillingItemRecordId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppliedItemTaxRecordDto

`func NewAppliedItemTaxRecordDto() *AppliedItemTaxRecordDto`

NewAppliedItemTaxRecordDto instantiates a new AppliedItemTaxRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedItemTaxRecordDtoWithDefaults

`func NewAppliedItemTaxRecordDtoWithDefaults() *AppliedItemTaxRecordDto`

NewAppliedItemTaxRecordDtoWithDefaults instantiates a new AppliedItemTaxRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppliedItemTaxRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppliedItemTaxRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppliedItemTaxRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppliedItemTaxRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AppliedItemTaxRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AppliedItemTaxRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AppliedItemTaxRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AppliedItemTaxRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AppliedItemTaxRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AppliedItemTaxRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AppliedItemTaxRecordDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AppliedItemTaxRecordDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *AppliedItemTaxRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AppliedItemTaxRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AppliedItemTaxRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AppliedItemTaxRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AppliedItemTaxRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AppliedItemTaxRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AppliedItemTaxRecordDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AppliedItemTaxRecordDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AppliedItemTaxRecordDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AppliedItemTaxRecordDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AppliedItemTaxRecordDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AppliedItemTaxRecordDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTaxPolicyId

`func (o *AppliedItemTaxRecordDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *AppliedItemTaxRecordDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *AppliedItemTaxRecordDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *AppliedItemTaxRecordDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *AppliedItemTaxRecordDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *AppliedItemTaxRecordDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetInvoiceId

`func (o *AppliedItemTaxRecordDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *AppliedItemTaxRecordDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *AppliedItemTaxRecordDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *AppliedItemTaxRecordDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *AppliedItemTaxRecordDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *AppliedItemTaxRecordDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetItemId

`func (o *AppliedItemTaxRecordDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AppliedItemTaxRecordDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AppliedItemTaxRecordDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AppliedItemTaxRecordDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *AppliedItemTaxRecordDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *AppliedItemTaxRecordDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetTaxAmountInUSD

`func (o *AppliedItemTaxRecordDto) GetTaxAmountInUSD() float64`

GetTaxAmountInUSD returns the TaxAmountInUSD field if non-nil, zero value otherwise.

### GetTaxAmountInUSDOk

`func (o *AppliedItemTaxRecordDto) GetTaxAmountInUSDOk() (*float64, bool)`

GetTaxAmountInUSDOk returns a tuple with the TaxAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountInUSD

`func (o *AppliedItemTaxRecordDto) SetTaxAmountInUSD(v float64)`

SetTaxAmountInUSD sets TaxAmountInUSD field to given value.

### HasTaxAmountInUSD

`func (o *AppliedItemTaxRecordDto) HasTaxAmountInUSD() bool`

HasTaxAmountInUSD returns a boolean if a field has been set.

### GetTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordDto) GetTaxBaseAmountInUSD() float64`

GetTaxBaseAmountInUSD returns the TaxBaseAmountInUSD field if non-nil, zero value otherwise.

### GetTaxBaseAmountInUSDOk

`func (o *AppliedItemTaxRecordDto) GetTaxBaseAmountInUSDOk() (*float64, bool)`

GetTaxBaseAmountInUSDOk returns a tuple with the TaxBaseAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordDto) SetTaxBaseAmountInUSD(v float64)`

SetTaxBaseAmountInUSD sets TaxBaseAmountInUSD field to given value.

### HasTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordDto) HasTaxBaseAmountInUSD() bool`

HasTaxBaseAmountInUSD returns a boolean if a field has been set.

### GetBillingItemRecordId

`func (o *AppliedItemTaxRecordDto) GetBillingItemRecordId() string`

GetBillingItemRecordId returns the BillingItemRecordId field if non-nil, zero value otherwise.

### GetBillingItemRecordIdOk

`func (o *AppliedItemTaxRecordDto) GetBillingItemRecordIdOk() (*string, bool)`

GetBillingItemRecordIdOk returns a tuple with the BillingItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingItemRecordId

`func (o *AppliedItemTaxRecordDto) SetBillingItemRecordId(v string)`

SetBillingItemRecordId sets BillingItemRecordId field to given value.

### HasBillingItemRecordId

`func (o *AppliedItemTaxRecordDto) HasBillingItemRecordId() bool`

HasBillingItemRecordId returns a boolean if a field has been set.

### SetBillingItemRecordIdNil

`func (o *AppliedItemTaxRecordDto) SetBillingItemRecordIdNil(b bool)`

 SetBillingItemRecordIdNil sets the value for BillingItemRecordId to be an explicit nil

### UnsetBillingItemRecordId
`func (o *AppliedItemTaxRecordDto) UnsetBillingItemRecordId()`

UnsetBillingItemRecordId ensures that no value is present for BillingItemRecordId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


