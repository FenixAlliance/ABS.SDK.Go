# ItemRestockEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ItemId** | **string** |  | 
**WarehouseId** | **string** |  | 
**ItemRestockId** | **string** |  | 
**Quantity** | Pointer to **float64** |  | [optional] 
**OrderItemRecordId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemRestockEntryCreateDto

`func NewItemRestockEntryCreateDto(itemId string, warehouseId string, itemRestockId string, ) *ItemRestockEntryCreateDto`

NewItemRestockEntryCreateDto instantiates a new ItemRestockEntryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemRestockEntryCreateDtoWithDefaults

`func NewItemRestockEntryCreateDtoWithDefaults() *ItemRestockEntryCreateDto`

NewItemRestockEntryCreateDtoWithDefaults instantiates a new ItemRestockEntryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemRestockEntryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemRestockEntryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemRestockEntryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemRestockEntryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemRestockEntryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemRestockEntryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemRestockEntryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemRestockEntryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetItemId

`func (o *ItemRestockEntryCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemRestockEntryCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemRestockEntryCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetWarehouseId

`func (o *ItemRestockEntryCreateDto) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ItemRestockEntryCreateDto) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ItemRestockEntryCreateDto) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.


### GetItemRestockId

`func (o *ItemRestockEntryCreateDto) GetItemRestockId() string`

GetItemRestockId returns the ItemRestockId field if non-nil, zero value otherwise.

### GetItemRestockIdOk

`func (o *ItemRestockEntryCreateDto) GetItemRestockIdOk() (*string, bool)`

GetItemRestockIdOk returns a tuple with the ItemRestockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemRestockId

`func (o *ItemRestockEntryCreateDto) SetItemRestockId(v string)`

SetItemRestockId sets ItemRestockId field to given value.


### GetQuantity

`func (o *ItemRestockEntryCreateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ItemRestockEntryCreateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ItemRestockEntryCreateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ItemRestockEntryCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetOrderItemRecordId

`func (o *ItemRestockEntryCreateDto) GetOrderItemRecordId() string`

GetOrderItemRecordId returns the OrderItemRecordId field if non-nil, zero value otherwise.

### GetOrderItemRecordIdOk

`func (o *ItemRestockEntryCreateDto) GetOrderItemRecordIdOk() (*string, bool)`

GetOrderItemRecordIdOk returns a tuple with the OrderItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemRecordId

`func (o *ItemRestockEntryCreateDto) SetOrderItemRecordId(v string)`

SetOrderItemRecordId sets OrderItemRecordId field to given value.

### HasOrderItemRecordId

`func (o *ItemRestockEntryCreateDto) HasOrderItemRecordId() bool`

HasOrderItemRecordId returns a boolean if a field has been set.

### SetOrderItemRecordIdNil

`func (o *ItemRestockEntryCreateDto) SetOrderItemRecordIdNil(b bool)`

 SetOrderItemRecordIdNil sets the value for OrderItemRecordId to be an explicit nil

### UnsetOrderItemRecordId
`func (o *ItemRestockEntryCreateDto) UnsetOrderItemRecordId()`

UnsetOrderItemRecordId ensures that no value is present for OrderItemRecordId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


