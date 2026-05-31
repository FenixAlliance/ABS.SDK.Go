# AppliedItemTaxRecordUpdateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxPolicyId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**TaxAmountInUSD** | Pointer to **float64** |  | [optional] 
**TaxBaseAmountInUSD** | Pointer to **float64** |  | [optional] 
**BillingItemRecordId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppliedItemTaxRecordUpdateDto

`func NewAppliedItemTaxRecordUpdateDto() *AppliedItemTaxRecordUpdateDto`

NewAppliedItemTaxRecordUpdateDto instantiates a new AppliedItemTaxRecordUpdateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedItemTaxRecordUpdateDtoWithDefaults

`func NewAppliedItemTaxRecordUpdateDtoWithDefaults() *AppliedItemTaxRecordUpdateDto`

NewAppliedItemTaxRecordUpdateDtoWithDefaults instantiates a new AppliedItemTaxRecordUpdateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxPolicyId

`func (o *AppliedItemTaxRecordUpdateDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *AppliedItemTaxRecordUpdateDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *AppliedItemTaxRecordUpdateDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *AppliedItemTaxRecordUpdateDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *AppliedItemTaxRecordUpdateDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *AppliedItemTaxRecordUpdateDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetInvoiceId

`func (o *AppliedItemTaxRecordUpdateDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *AppliedItemTaxRecordUpdateDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *AppliedItemTaxRecordUpdateDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *AppliedItemTaxRecordUpdateDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *AppliedItemTaxRecordUpdateDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *AppliedItemTaxRecordUpdateDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetItemId

`func (o *AppliedItemTaxRecordUpdateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AppliedItemTaxRecordUpdateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AppliedItemTaxRecordUpdateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AppliedItemTaxRecordUpdateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *AppliedItemTaxRecordUpdateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *AppliedItemTaxRecordUpdateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetTaxAmountInUSD

`func (o *AppliedItemTaxRecordUpdateDto) GetTaxAmountInUSD() float64`

GetTaxAmountInUSD returns the TaxAmountInUSD field if non-nil, zero value otherwise.

### GetTaxAmountInUSDOk

`func (o *AppliedItemTaxRecordUpdateDto) GetTaxAmountInUSDOk() (*float64, bool)`

GetTaxAmountInUSDOk returns a tuple with the TaxAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountInUSD

`func (o *AppliedItemTaxRecordUpdateDto) SetTaxAmountInUSD(v float64)`

SetTaxAmountInUSD sets TaxAmountInUSD field to given value.

### HasTaxAmountInUSD

`func (o *AppliedItemTaxRecordUpdateDto) HasTaxAmountInUSD() bool`

HasTaxAmountInUSD returns a boolean if a field has been set.

### GetTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordUpdateDto) GetTaxBaseAmountInUSD() float64`

GetTaxBaseAmountInUSD returns the TaxBaseAmountInUSD field if non-nil, zero value otherwise.

### GetTaxBaseAmountInUSDOk

`func (o *AppliedItemTaxRecordUpdateDto) GetTaxBaseAmountInUSDOk() (*float64, bool)`

GetTaxBaseAmountInUSDOk returns a tuple with the TaxBaseAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordUpdateDto) SetTaxBaseAmountInUSD(v float64)`

SetTaxBaseAmountInUSD sets TaxBaseAmountInUSD field to given value.

### HasTaxBaseAmountInUSD

`func (o *AppliedItemTaxRecordUpdateDto) HasTaxBaseAmountInUSD() bool`

HasTaxBaseAmountInUSD returns a boolean if a field has been set.

### GetBillingItemRecordId

`func (o *AppliedItemTaxRecordUpdateDto) GetBillingItemRecordId() string`

GetBillingItemRecordId returns the BillingItemRecordId field if non-nil, zero value otherwise.

### GetBillingItemRecordIdOk

`func (o *AppliedItemTaxRecordUpdateDto) GetBillingItemRecordIdOk() (*string, bool)`

GetBillingItemRecordIdOk returns a tuple with the BillingItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingItemRecordId

`func (o *AppliedItemTaxRecordUpdateDto) SetBillingItemRecordId(v string)`

SetBillingItemRecordId sets BillingItemRecordId field to given value.

### HasBillingItemRecordId

`func (o *AppliedItemTaxRecordUpdateDto) HasBillingItemRecordId() bool`

HasBillingItemRecordId returns a boolean if a field has been set.

### SetBillingItemRecordIdNil

`func (o *AppliedItemTaxRecordUpdateDto) SetBillingItemRecordIdNil(b bool)`

 SetBillingItemRecordIdNil sets the value for BillingItemRecordId to be an explicit nil

### UnsetBillingItemRecordId
`func (o *AppliedItemTaxRecordUpdateDto) UnsetBillingItemRecordId()`

UnsetBillingItemRecordId ensures that no value is present for BillingItemRecordId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


