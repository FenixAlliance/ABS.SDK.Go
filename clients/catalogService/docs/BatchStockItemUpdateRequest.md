# BatchStockItemUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemIds** | Pointer to **[]string** |  | [optional] 
**Published** | Pointer to **NullableBool** |  | [optional] 
**Taxable** | Pointer to **NullableBool** |  | [optional] 
**AddTaxPolicyIds** | Pointer to **[]string** |  | [optional] 
**RemoveTaxPolicyIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBatchStockItemUpdateRequest

`func NewBatchStockItemUpdateRequest() *BatchStockItemUpdateRequest`

NewBatchStockItemUpdateRequest instantiates a new BatchStockItemUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchStockItemUpdateRequestWithDefaults

`func NewBatchStockItemUpdateRequestWithDefaults() *BatchStockItemUpdateRequest`

NewBatchStockItemUpdateRequestWithDefaults instantiates a new BatchStockItemUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemIds

`func (o *BatchStockItemUpdateRequest) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *BatchStockItemUpdateRequest) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *BatchStockItemUpdateRequest) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *BatchStockItemUpdateRequest) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### SetItemIdsNil

`func (o *BatchStockItemUpdateRequest) SetItemIdsNil(b bool)`

 SetItemIdsNil sets the value for ItemIds to be an explicit nil

### UnsetItemIds
`func (o *BatchStockItemUpdateRequest) UnsetItemIds()`

UnsetItemIds ensures that no value is present for ItemIds, not even an explicit nil
### GetPublished

`func (o *BatchStockItemUpdateRequest) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *BatchStockItemUpdateRequest) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *BatchStockItemUpdateRequest) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *BatchStockItemUpdateRequest) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *BatchStockItemUpdateRequest) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *BatchStockItemUpdateRequest) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil
### GetTaxable

`func (o *BatchStockItemUpdateRequest) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *BatchStockItemUpdateRequest) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *BatchStockItemUpdateRequest) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *BatchStockItemUpdateRequest) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### SetTaxableNil

`func (o *BatchStockItemUpdateRequest) SetTaxableNil(b bool)`

 SetTaxableNil sets the value for Taxable to be an explicit nil

### UnsetTaxable
`func (o *BatchStockItemUpdateRequest) UnsetTaxable()`

UnsetTaxable ensures that no value is present for Taxable, not even an explicit nil
### GetAddTaxPolicyIds

`func (o *BatchStockItemUpdateRequest) GetAddTaxPolicyIds() []string`

GetAddTaxPolicyIds returns the AddTaxPolicyIds field if non-nil, zero value otherwise.

### GetAddTaxPolicyIdsOk

`func (o *BatchStockItemUpdateRequest) GetAddTaxPolicyIdsOk() (*[]string, bool)`

GetAddTaxPolicyIdsOk returns a tuple with the AddTaxPolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTaxPolicyIds

`func (o *BatchStockItemUpdateRequest) SetAddTaxPolicyIds(v []string)`

SetAddTaxPolicyIds sets AddTaxPolicyIds field to given value.

### HasAddTaxPolicyIds

`func (o *BatchStockItemUpdateRequest) HasAddTaxPolicyIds() bool`

HasAddTaxPolicyIds returns a boolean if a field has been set.

### SetAddTaxPolicyIdsNil

`func (o *BatchStockItemUpdateRequest) SetAddTaxPolicyIdsNil(b bool)`

 SetAddTaxPolicyIdsNil sets the value for AddTaxPolicyIds to be an explicit nil

### UnsetAddTaxPolicyIds
`func (o *BatchStockItemUpdateRequest) UnsetAddTaxPolicyIds()`

UnsetAddTaxPolicyIds ensures that no value is present for AddTaxPolicyIds, not even an explicit nil
### GetRemoveTaxPolicyIds

`func (o *BatchStockItemUpdateRequest) GetRemoveTaxPolicyIds() []string`

GetRemoveTaxPolicyIds returns the RemoveTaxPolicyIds field if non-nil, zero value otherwise.

### GetRemoveTaxPolicyIdsOk

`func (o *BatchStockItemUpdateRequest) GetRemoveTaxPolicyIdsOk() (*[]string, bool)`

GetRemoveTaxPolicyIdsOk returns a tuple with the RemoveTaxPolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTaxPolicyIds

`func (o *BatchStockItemUpdateRequest) SetRemoveTaxPolicyIds(v []string)`

SetRemoveTaxPolicyIds sets RemoveTaxPolicyIds field to given value.

### HasRemoveTaxPolicyIds

`func (o *BatchStockItemUpdateRequest) HasRemoveTaxPolicyIds() bool`

HasRemoveTaxPolicyIds returns a boolean if a field has been set.

### SetRemoveTaxPolicyIdsNil

`func (o *BatchStockItemUpdateRequest) SetRemoveTaxPolicyIdsNil(b bool)`

 SetRemoveTaxPolicyIdsNil sets the value for RemoveTaxPolicyIds to be an explicit nil

### UnsetRemoveTaxPolicyIds
`func (o *BatchStockItemUpdateRequest) UnsetRemoveTaxPolicyIds()`

UnsetRemoveTaxPolicyIds ensures that no value is present for RemoveTaxPolicyIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


