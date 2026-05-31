# AppliedItemTaxRecordCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**TaxPolicyId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**TaxAmountInUSD** | Pointer to **float64** |  | [optional] 
**TaxBaseAmountInUSD** | Pointer to **float64** |  | [optional] 
**BillingItemRecordId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppliedItemTaxRecordCreateDto

`func NewAppliedItemTaxRecordCreateDto() *AppliedItemTaxRecordCreateDto`

NewAppliedItemTaxRecordCreateDto instantiates a new AppliedItemTaxRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedItemTaxRecordCreateDtoWithDefaults

`func NewAppliedItemTaxRecordCreateDtoWithDefaults() *AppliedItemTaxRecordCreateDto`

NewAppliedItemTaxRecordCreateDtoWithDefaults instantiates a new AppliedItemTaxRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppliedItemTaxRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppliedItemTaxRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppliedItemTaxRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppliedItemTaxRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AppliedItemTaxRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AppliedItemTaxRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AppliedItemTaxRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AppliedItemTaxRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTaxPolicyId

`func (o *AppliedItemTaxRecordCreateDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *AppliedItemTaxRecordCreateDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *AppliedItemTaxRecordCreateDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *AppliedItemTaxRecordCreateDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *AppliedItemTaxRecordCreateDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *AppliedItemTaxRecordCreateDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetInvoiceId

`func (o *AppliedItemTaxRecordCreateDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *AppliedItemTaxRecordCreateDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *AppliedItemTaxRecordCreateDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *AppliedItemTaxRecordCreateDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *AppliedItemTaxRecordCreateDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *AppliedItemTaxRecordCreateDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetItemId

`func (o *AppliedItemTaxRecordCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AppliedItemTaxRecordCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AppliedItemTaxRecordCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AppliedItemTaxRecordCreateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *AppliedItemTaxRecordCreateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *AppliedItemTaxRecordCreateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetTaxAmountInUSD

`func (o *AppliedItemTaxRecordCreateDto) GetTaxAmountInUSD() float64`

GetTaxAmountInUSD returns the TaxAmountInUSD field if non-nil, zero value otherwise.

### GetTaxAmountInUSDOk

`func (o *AppliedItemTaxRecordCreateDto) GetTaxAmountInUSDOk() (*float64, bool)`

GetTaxAmountInUSDOk returns a tuple with the TaxAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountInUSD

`func (o *AppliedItemTaxRecordCreateDto) SetTaxAmountInUSD(v float64)`

SetTaxAmountInUSD sets TaxAmountInUSD field to given value.

### HasTaxAmountInUSD

`func (o *AppliedItemTaxRecordCreateDto) HasTaxAmountInUSD() bool`

HasTaxAmountInUSD returns a boolean if a field has been set.

### GetTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordCreateDto) GetTaxBaseAmountInUSD() float64`

GetTaxBaseAmountInUSD returns the TaxBaseAmountInUSD field if non-nil, zero value otherwise.

### GetTaxBaseAmountInUSDOk

`func (o *AppliedItemTaxRecordCreateDto) GetTaxBaseAmountInUSDOk() (*float64, bool)`

GetTaxBaseAmountInUSDOk returns a tuple with the TaxBaseAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordCreateDto) SetTaxBaseAmountInUSD(v float64)`

SetTaxBaseAmountInUSD sets TaxBaseAmountInUSD field to given value.

### HasTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordCreateDto) HasTaxBaseAmountInUSD() bool`

HasTaxBaseAmountInUSD returns a boolean if a field has been set.

### GetBillingItemRecordId

`func (o *AppliedItemTaxRecordCreateDto) GetBillingItemRecordId() string`

GetBillingItemRecordId returns the BillingItemRecordId field if non-nil, zero value otherwise.

### GetBillingItemRecordIdOk

`func (o *AppliedItemTaxRecordCreateDto) GetBillingItemRecordIdOk() (*string, bool)`

GetBillingItemRecordIdOk returns a tuple with the BillingItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingItemRecordId

`func (o *AppliedItemTaxRecordCreateDto) SetBillingItemRecordId(v string)`

SetBillingItemRecordId sets BillingItemRecordId field to given value.

### HasBillingItemRecordId

`func (o *AppliedItemTaxRecordCreateDto) HasBillingItemRecordId() bool`

HasBillingItemRecordId returns a boolean if a field has been set.

### SetBillingItemRecordIdNil

`func (o *AppliedItemTaxRecordCreateDto) SetBillingItemRecordIdNil(b bool)`

 SetBillingItemRecordIdNil sets the value for BillingItemRecordId to be an explicit nil

### UnsetBillingItemRecordId
`func (o *AppliedItemTaxRecordCreateDto) UnsetBillingItemRecordId()`

UnsetBillingItemRecordId ensures that no value is present for BillingItemRecordId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


