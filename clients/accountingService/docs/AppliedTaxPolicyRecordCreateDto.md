# AppliedTaxPolicyRecordCreateDto

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

## Methods

### NewAppliedTaxPolicyRecordCreateDto

`func NewAppliedTaxPolicyRecordCreateDto() *AppliedTaxPolicyRecordCreateDto`

NewAppliedTaxPolicyRecordCreateDto instantiates a new AppliedTaxPolicyRecordCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedTaxPolicyRecordCreateDtoWithDefaults

`func NewAppliedTaxPolicyRecordCreateDtoWithDefaults() *AppliedTaxPolicyRecordCreateDto`

NewAppliedTaxPolicyRecordCreateDtoWithDefaults instantiates a new AppliedTaxPolicyRecordCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppliedTaxPolicyRecordCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppliedTaxPolicyRecordCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppliedTaxPolicyRecordCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppliedTaxPolicyRecordCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *AppliedTaxPolicyRecordCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AppliedTaxPolicyRecordCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AppliedTaxPolicyRecordCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AppliedTaxPolicyRecordCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTaxPolicyId

`func (o *AppliedTaxPolicyRecordCreateDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *AppliedTaxPolicyRecordCreateDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *AppliedTaxPolicyRecordCreateDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *AppliedTaxPolicyRecordCreateDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *AppliedTaxPolicyRecordCreateDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *AppliedTaxPolicyRecordCreateDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetInvoiceId

`func (o *AppliedTaxPolicyRecordCreateDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *AppliedTaxPolicyRecordCreateDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *AppliedTaxPolicyRecordCreateDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *AppliedTaxPolicyRecordCreateDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *AppliedTaxPolicyRecordCreateDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *AppliedTaxPolicyRecordCreateDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetItemId

`func (o *AppliedTaxPolicyRecordCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AppliedTaxPolicyRecordCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AppliedTaxPolicyRecordCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *AppliedTaxPolicyRecordCreateDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *AppliedTaxPolicyRecordCreateDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *AppliedTaxPolicyRecordCreateDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetTaxAmountInUSD

`func (o *AppliedTaxPolicyRecordCreateDto) GetTaxAmountInUSD() float64`

GetTaxAmountInUSD returns the TaxAmountInUSD field if non-nil, zero value otherwise.

### GetTaxAmountInUSDOk

`func (o *AppliedTaxPolicyRecordCreateDto) GetTaxAmountInUSDOk() (*float64, bool)`

GetTaxAmountInUSDOk returns a tuple with the TaxAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountInUSD

`func (o *AppliedTaxPolicyRecordCreateDto) SetTaxAmountInUSD(v float64)`

SetTaxAmountInUSD sets TaxAmountInUSD field to given value.

### HasTaxAmountInUSD

`func (o *AppliedTaxPolicyRecordCreateDto) HasTaxAmountInUSD() bool`

HasTaxAmountInUSD returns a boolean if a field has been set.

### GetTaxBaseAmountInUSD

`func (o *AppliedTaxPolicyRecordCreateDto) GetTaxBaseAmountInUSD() float64`

GetTaxBaseAmountInUSD returns the TaxBaseAmountInUSD field if non-nil, zero value otherwise.

### GetTaxBaseAmountInUSDOk

`func (o *AppliedTaxPolicyRecordCreateDto) GetTaxBaseAmountInUSDOk() (*float64, bool)`

GetTaxBaseAmountInUSDOk returns a tuple with the TaxBaseAmountInUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBaseAmountInUSD

`func (o *AppliedTaxPolicyRecordCreateDto) SetTaxBaseAmountInUSD(v float64)`

SetTaxBaseAmountInUSD sets TaxBaseAmountInUSD field to given value.

### HasTaxBaseAmountInUSD

`func (o *AppliedTaxPolicyRecordCreateDto) HasTaxBaseAmountInUSD() bool`

HasTaxBaseAmountInUSD returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


