# ItemPackingSlipEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ItemId** | **string** |  | 
**ItemPackingSlipId** | **string** |  | 
**Quantity** | **string** |  | 

## Methods

### NewItemPackingSlipEntryCreateDto

`func NewItemPackingSlipEntryCreateDto(itemId string, itemPackingSlipId string, quantity string, ) *ItemPackingSlipEntryCreateDto`

NewItemPackingSlipEntryCreateDto instantiates a new ItemPackingSlipEntryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPackingSlipEntryCreateDtoWithDefaults

`func NewItemPackingSlipEntryCreateDtoWithDefaults() *ItemPackingSlipEntryCreateDto`

NewItemPackingSlipEntryCreateDtoWithDefaults instantiates a new ItemPackingSlipEntryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemPackingSlipEntryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemPackingSlipEntryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemPackingSlipEntryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemPackingSlipEntryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemPackingSlipEntryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPackingSlipEntryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPackingSlipEntryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPackingSlipEntryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetItemId

`func (o *ItemPackingSlipEntryCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPackingSlipEntryCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPackingSlipEntryCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetItemPackingSlipId

`func (o *ItemPackingSlipEntryCreateDto) GetItemPackingSlipId() string`

GetItemPackingSlipId returns the ItemPackingSlipId field if non-nil, zero value otherwise.

### GetItemPackingSlipIdOk

`func (o *ItemPackingSlipEntryCreateDto) GetItemPackingSlipIdOk() (*string, bool)`

GetItemPackingSlipIdOk returns a tuple with the ItemPackingSlipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPackingSlipId

`func (o *ItemPackingSlipEntryCreateDto) SetItemPackingSlipId(v string)`

SetItemPackingSlipId sets ItemPackingSlipId field to given value.


### GetQuantity

`func (o *ItemPackingSlipEntryCreateDto) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ItemPackingSlipEntryCreateDto) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ItemPackingSlipEntryCreateDto) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


