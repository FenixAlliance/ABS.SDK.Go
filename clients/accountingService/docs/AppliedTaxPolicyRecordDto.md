# AppliedTaxPolicyRecordDto

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

## Methods

### NewAppliedTaxPolicyRecordDto

`func NewAppliedTaxPolicyRecordDto() *AppliedTaxPolicyRecordDto`

NewAppliedTaxPolicyRecordDto instantiates a new AppliedTaxPolicyRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedTaxPolicyRecordDtoWithDefaults

`func NewAppliedTaxPolicyRecordDtoWithDefaults() *AppliedTaxPolicyRecordDto`

NewAppliedTaxPolicyRecordDtoWithDefaults instantiates a new AppliedTaxPolicyRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppliedTaxPolicyRecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppliedTaxPolicyRecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppliedTaxPolicyRecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppliedTaxPolicyRecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AppliedTaxPolicyRecordDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AppliedTaxPolicyRecordDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *AppliedTaxPolicyRecordDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AppliedTaxPolicyRecordDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AppliedTaxPolicyRecordDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AppliedTaxPolicyRecordDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *AppliedTaxPolicyRecordDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *AppliedTaxPolicyRecordDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *AppliedTaxPolicyRecordDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AppliedTaxPolicyRecordDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AppliedTaxPolicyRecordDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AppliedTaxPolicyRecordDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AppliedTaxPolicyRecordDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AppliedTaxPolicyRecordDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetEnrollmentId

`func (o *AppliedTaxPolicyRecordDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *AppliedTaxPolicyRecordDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *AppliedTaxPolicyRecordDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *AppliedTaxPolicyRecordDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *AppliedTaxPolicyRecordDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *AppliedTaxPolicyRecordDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetTaxPolicyId

`func (o *AppliedTaxPolicyRecordDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *AppliedTaxPolicyRecordDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *AppliedTaxPolicyRecordDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *AppliedTaxPolicyRecordDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *AppliedTaxPolicyRecordDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *AppliedTaxPolicyRecordDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetInvoiceId

`func (o *AppliedTaxPolicyRecordDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *AppliedTaxPolicyRecordDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *AppliedTaxPolicyRecordDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *AppliedTaxPolicyRecordDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *AppliedTaxPolicyRecordDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *AppliedTaxPolicyRecordDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetItemId

`func (o *AppliedTaxPolicyRecordDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AppliedTaxPolicyRecordDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AppliedTaxPolicyRecordDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AppliedTaxPolicyRecordDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *AppliedTaxPolicyRecordDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *AppliedTaxPolicyRecordDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetTaxAmountInUSD

`func (o *AppliedTaxPolicyRecordDto) GetTaxAmountInUSD() float64`

GetTaxAmountInUSD returns the TaxAmountInUSD field if non-nil, zero value otherwise.

### GetTaxAmountInUSDOk

`func (o *AppliedTaxPolicyRecordDto) GetTaxAmountInUSDOk() (*float64, bool)`

GetTaxAmountInUSDOk returns a tuple with the TaxAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountInUSD

`func (o *AppliedTaxPolicyRecordDto) SetTaxAmountInUSD(v float64)`

SetTaxAmountInUSD sets TaxAmountInUSD field to given value.

### HasTaxAmountInUSD

`func (o *AppliedTaxPolicyRecordDto) HasTaxAmountInUSD() bool`

HasTaxAmountInUSD returns a boolean if a field has been set.

### GetTaxBaseAmountInUSD

`func (o *AppliedTaxPolicyRecordDto) GetTaxBaseAmountInUSD() float64`

GetTaxBaseAmountInUSD returns the TaxBaseAmountInUSD field if non-nil, zero value otherwise.

### GetTaxBaseAmountInUSDOk

`func (o *AppliedTaxPolicyRecordDto) GetTaxBaseAmountInUSDOk() (*float64, bool)`

GetTaxBaseAmountInUSDOk returns a tuple with the TaxBaseAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBaseAmountInUSD

`func (o *AppliedTaxPolicyRecordDto) SetTaxBaseAmountInUSD(v float64)`

SetTaxBaseAmountInUSD sets TaxBaseAmountInUSD field to given value.

### HasTaxBaseAmountInUSD

`func (o *AppliedTaxPolicyRecordDto) HasTaxBaseAmountInUSD() bool`

HasTaxBaseAmountInUSD returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


