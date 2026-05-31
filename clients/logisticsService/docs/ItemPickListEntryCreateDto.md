# ItemPickListEntryCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ItemId** | **string** |  | 
**WarehouseId** | **string** |  | 
**ItemPickListId** | **string** |  | 
**Quantity** | Pointer to **float64** |  | [optional] 
**OrderItemRecordId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemPickListEntryCreateDto

`func NewItemPickListEntryCreateDto(itemId string, warehouseId string, itemPickListId string, ) *ItemPickListEntryCreateDto`

NewItemPickListEntryCreateDto instantiates a new ItemPickListEntryCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPickListEntryCreateDtoWithDefaults

`func NewItemPickListEntryCreateDtoWithDefaults() *ItemPickListEntryCreateDto`

NewItemPickListEntryCreateDtoWithDefaults instantiates a new ItemPickListEntryCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemPickListEntryCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemPickListEntryCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemPickListEntryCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemPickListEntryCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemPickListEntryCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPickListEntryCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPickListEntryCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPickListEntryCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetItemId

`func (o *ItemPickListEntryCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPickListEntryCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPickListEntryCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetWarehouseId

`func (o *ItemPickListEntryCreateDto) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ItemPickListEntryCreateDto) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ItemPickListEntryCreateDto) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.


### GetItemPickListId

`func (o *ItemPickListEntryCreateDto) GetItemPickListId() string`

GetItemPickListId returns the ItemPickListId field if non-nil, zero value otherwise.

### GetItemPickListIdOk

`func (o *ItemPickListEntryCreateDto) GetItemPickListIdOk() (*string, bool)`

GetItemPickListIdOk returns a tuple with the ItemPickListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPickListId

`func (o *ItemPickListEntryCreateDto) SetItemPickListId(v string)`

SetItemPickListId sets ItemPickListId field to given value.


### GetQuantity

`func (o *ItemPickListEntryCreateDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ItemPickListEntryCreateDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ItemPickListEntryCreateDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ItemPickListEntryCreateDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetOrderItemRecordId

`func (o *ItemPickListEntryCreateDto) GetOrderItemRecordId() string`

GetOrderItemRecordId returns the OrderItemRecordId field if non-nil, zero value otherwise.

### GetOrderItemRecordIdOk

`func (o *ItemPickListEntryCreateDto) GetOrderItemRecordIdOk() (*string, bool)`

GetOrderItemRecordIdOk returns a tuple with the OrderItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemRecordId

`func (o *ItemPickListEntryCreateDto) SetOrderItemRecordId(v string)`

SetOrderItemRecordId sets OrderItemRecordId field to given value.

### HasOrderItemRecordId

`func (o *ItemPickListEntryCreateDto) HasOrderItemRecordId() bool`

HasOrderItemRecordId returns a boolean if a field has been set.

### SetOrderItemRecordIdNil

`func (o *ItemPickListEntryCreateDto) SetOrderItemRecordIdNil(b bool)`

 SetOrderItemRecordIdNil sets the value for OrderItemRecordId to be an explicit nil

### UnsetOrderItemRecordId
`func (o *ItemPickListEntryCreateDto) UnsetOrderItemRecordId()`

UnsetOrderItemRecordId ensures that no value is present for OrderItemRecordId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


