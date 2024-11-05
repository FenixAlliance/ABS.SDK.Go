# InvoiceLineAppliedTaxDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentId** | Pointer to **NullableString** |  | [optional] 
**InvoiceLineId** | Pointer to **NullableString** |  | [optional] 
**TaxPolicyId** | Pointer to **NullableString** |  | [optional] 
**ItemPriceId** | Pointer to **NullableString** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**TaxAmountInUsd** | Pointer to **float64** |  | [optional] 
**TaxBaseAmountInUsd** | Pointer to **float64** |  | [optional] 
**TaxPolicyName** | Pointer to **NullableString** |  | [optional] 
**TaxPolicyDescription** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvoiceLineAppliedTaxDto

`func NewInvoiceLineAppliedTaxDto() *InvoiceLineAppliedTaxDto`

NewInvoiceLineAppliedTaxDto instantiates a new InvoiceLineAppliedTaxDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineAppliedTaxDtoWithDefaults

`func NewInvoiceLineAppliedTaxDtoWithDefaults() *InvoiceLineAppliedTaxDto`

NewInvoiceLineAppliedTaxDtoWithDefaults instantiates a new InvoiceLineAppliedTaxDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceLineAppliedTaxDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineAppliedTaxDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineAppliedTaxDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceLineAppliedTaxDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InvoiceLineAppliedTaxDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvoiceLineAppliedTaxDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *InvoiceLineAppliedTaxDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InvoiceLineAppliedTaxDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InvoiceLineAppliedTaxDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InvoiceLineAppliedTaxDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InvoiceLineAppliedTaxDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InvoiceLineAppliedTaxDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTenantId

`func (o *InvoiceLineAppliedTaxDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *InvoiceLineAppliedTaxDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *InvoiceLineAppliedTaxDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *InvoiceLineAppliedTaxDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *InvoiceLineAppliedTaxDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *InvoiceLineAppliedTaxDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
### GetInvoiceId

`func (o *InvoiceLineAppliedTaxDto) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceLineAppliedTaxDto) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceLineAppliedTaxDto) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *InvoiceLineAppliedTaxDto) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *InvoiceLineAppliedTaxDto) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *InvoiceLineAppliedTaxDto) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetEnrollmentId

`func (o *InvoiceLineAppliedTaxDto) GetEnrollmentId() string`

GetEnrollmentId returns the EnrollmentId field if non-nil, zero value otherwise.

### GetEnrollmentIdOk

`func (o *InvoiceLineAppliedTaxDto) GetEnrollmentIdOk() (*string, bool)`

GetEnrollmentIdOk returns a tuple with the EnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentId

`func (o *InvoiceLineAppliedTaxDto) SetEnrollmentId(v string)`

SetEnrollmentId sets EnrollmentId field to given value.

### HasEnrollmentId

`func (o *InvoiceLineAppliedTaxDto) HasEnrollmentId() bool`

HasEnrollmentId returns a boolean if a field has been set.

### SetEnrollmentIdNil

`func (o *InvoiceLineAppliedTaxDto) SetEnrollmentIdNil(b bool)`

 SetEnrollmentIdNil sets the value for EnrollmentId to be an explicit nil

### UnsetEnrollmentId
`func (o *InvoiceLineAppliedTaxDto) UnsetEnrollmentId()`

UnsetEnrollmentId ensures that no value is present for EnrollmentId, not even an explicit nil
### GetInvoiceLineId

`func (o *InvoiceLineAppliedTaxDto) GetInvoiceLineId() string`

GetInvoiceLineId returns the InvoiceLineId field if non-nil, zero value otherwise.

### GetInvoiceLineIdOk

`func (o *InvoiceLineAppliedTaxDto) GetInvoiceLineIdOk() (*string, bool)`

GetInvoiceLineIdOk returns a tuple with the InvoiceLineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLineId

`func (o *InvoiceLineAppliedTaxDto) SetInvoiceLineId(v string)`

SetInvoiceLineId sets InvoiceLineId field to given value.

### HasInvoiceLineId

`func (o *InvoiceLineAppliedTaxDto) HasInvoiceLineId() bool`

HasInvoiceLineId returns a boolean if a field has been set.

### SetInvoiceLineIdNil

`func (o *InvoiceLineAppliedTaxDto) SetInvoiceLineIdNil(b bool)`

 SetInvoiceLineIdNil sets the value for InvoiceLineId to be an explicit nil

### UnsetInvoiceLineId
`func (o *InvoiceLineAppliedTaxDto) UnsetInvoiceLineId()`

UnsetInvoiceLineId ensures that no value is present for InvoiceLineId, not even an explicit nil
### GetTaxPolicyId

`func (o *InvoiceLineAppliedTaxDto) GetTaxPolicyId() string`

GetTaxPolicyId returns the TaxPolicyId field if non-nil, zero value otherwise.

### GetTaxPolicyIdOk

`func (o *InvoiceLineAppliedTaxDto) GetTaxPolicyIdOk() (*string, bool)`

GetTaxPolicyIdOk returns a tuple with the TaxPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyId

`func (o *InvoiceLineAppliedTaxDto) SetTaxPolicyId(v string)`

SetTaxPolicyId sets TaxPolicyId field to given value.

### HasTaxPolicyId

`func (o *InvoiceLineAppliedTaxDto) HasTaxPolicyId() bool`

HasTaxPolicyId returns a boolean if a field has been set.

### SetTaxPolicyIdNil

`func (o *InvoiceLineAppliedTaxDto) SetTaxPolicyIdNil(b bool)`

 SetTaxPolicyIdNil sets the value for TaxPolicyId to be an explicit nil

### UnsetTaxPolicyId
`func (o *InvoiceLineAppliedTaxDto) UnsetTaxPolicyId()`

UnsetTaxPolicyId ensures that no value is present for TaxPolicyId, not even an explicit nil
### GetItemPriceId

`func (o *InvoiceLineAppliedTaxDto) GetItemPriceId() string`

GetItemPriceId returns the ItemPriceId field if non-nil, zero value otherwise.

### GetItemPriceIdOk

`func (o *InvoiceLineAppliedTaxDto) GetItemPriceIdOk() (*string, bool)`

GetItemPriceIdOk returns a tuple with the ItemPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPriceId

`func (o *InvoiceLineAppliedTaxDto) SetItemPriceId(v string)`

SetItemPriceId sets ItemPriceId field to given value.

### HasItemPriceId

`func (o *InvoiceLineAppliedTaxDto) HasItemPriceId() bool`

HasItemPriceId returns a boolean if a field has been set.

### SetItemPriceIdNil

`func (o *InvoiceLineAppliedTaxDto) SetItemPriceIdNil(b bool)`

 SetItemPriceIdNil sets the value for ItemPriceId to be an explicit nil

### UnsetItemPriceId
`func (o *InvoiceLineAppliedTaxDto) UnsetItemPriceId()`

UnsetItemPriceId ensures that no value is present for ItemPriceId, not even an explicit nil
### GetItemId

`func (o *InvoiceLineAppliedTaxDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InvoiceLineAppliedTaxDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InvoiceLineAppliedTaxDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *InvoiceLineAppliedTaxDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *InvoiceLineAppliedTaxDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *InvoiceLineAppliedTaxDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetTaxAmountInUsd

`func (o *InvoiceLineAppliedTaxDto) GetTaxAmountInUsd() float64`

GetTaxAmountInUsd returns the TaxAmountInUsd field if non-nil, zero value otherwise.

### GetTaxAmountInUsdOk

`func (o *InvoiceLineAppliedTaxDto) GetTaxAmountInUsdOk() (*float64, bool)`

GetTaxAmountInUsdOk returns a tuple with the TaxAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountInUsd

`func (o *InvoiceLineAppliedTaxDto) SetTaxAmountInUsd(v float64)`

SetTaxAmountInUsd sets TaxAmountInUsd field to given value.

### HasTaxAmountInUsd

`func (o *InvoiceLineAppliedTaxDto) HasTaxAmountInUsd() bool`

HasTaxAmountInUsd returns a boolean if a field has been set.

### GetTaxBaseAmountInUsd

`func (o *InvoiceLineAppliedTaxDto) GetTaxBaseAmountInUsd() float64`

GetTaxBaseAmountInUsd returns the TaxBaseAmountInUsd field if non-nil, zero value otherwise.

### GetTaxBaseAmountInUsdOk

`func (o *InvoiceLineAppliedTaxDto) GetTaxBaseAmountInUsdOk() (*float64, bool)`

GetTaxBaseAmountInUsdOk returns a tuple with the TaxBaseAmountInUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBaseAmountInUsd

`func (o *InvoiceLineAppliedTaxDto) SetTaxBaseAmountInUsd(v float64)`

SetTaxBaseAmountInUsd sets TaxBaseAmountInUsd field to given value.

### HasTaxBaseAmountInUsd

`func (o *InvoiceLineAppliedTaxDto) HasTaxBaseAmountInUsd() bool`

HasTaxBaseAmountInUsd returns a boolean if a field has been set.

### GetTaxPolicyName

`func (o *InvoiceLineAppliedTaxDto) GetTaxPolicyName() string`

GetTaxPolicyName returns the TaxPolicyName field if non-nil, zero value otherwise.

### GetTaxPolicyNameOk

`func (o *InvoiceLineAppliedTaxDto) GetTaxPolicyNameOk() (*string, bool)`

GetTaxPolicyNameOk returns a tuple with the TaxPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyName

`func (o *InvoiceLineAppliedTaxDto) SetTaxPolicyName(v string)`

SetTaxPolicyName sets TaxPolicyName field to given value.

### HasTaxPolicyName

`func (o *InvoiceLineAppliedTaxDto) HasTaxPolicyName() bool`

HasTaxPolicyName returns a boolean if a field has been set.

### SetTaxPolicyNameNil

`func (o *InvoiceLineAppliedTaxDto) SetTaxPolicyNameNil(b bool)`

 SetTaxPolicyNameNil sets the value for TaxPolicyName to be an explicit nil

### UnsetTaxPolicyName
`func (o *InvoiceLineAppliedTaxDto) UnsetTaxPolicyName()`

UnsetTaxPolicyName ensures that no value is present for TaxPolicyName, not even an explicit nil
### GetTaxPolicyDescription

`func (o *InvoiceLineAppliedTaxDto) GetTaxPolicyDescription() string`

GetTaxPolicyDescription returns the TaxPolicyDescription field if non-nil, zero value otherwise.

### GetTaxPolicyDescriptionOk

`func (o *InvoiceLineAppliedTaxDto) GetTaxPolicyDescriptionOk() (*string, bool)`

GetTaxPolicyDescriptionOk returns a tuple with the TaxPolicyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPolicyDescription

`func (o *InvoiceLineAppliedTaxDto) SetTaxPolicyDescription(v string)`

SetTaxPolicyDescription sets TaxPolicyDescription field to given value.

### HasTaxPolicyDescription

`func (o *InvoiceLineAppliedTaxDto) HasTaxPolicyDescription() bool`

HasTaxPolicyDescription returns a boolean if a field has been set.

### SetTaxPolicyDescriptionNil

`func (o *InvoiceLineAppliedTaxDto) SetTaxPolicyDescriptionNil(b bool)`

 SetTaxPolicyDescriptionNil sets the value for TaxPolicyDescription to be an explicit nil

### UnsetTaxPolicyDescription
`func (o *InvoiceLineAppliedTaxDto) UnsetTaxPolicyDescription()`

UnsetTaxPolicyDescription ensures that no value is present for TaxPolicyDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


