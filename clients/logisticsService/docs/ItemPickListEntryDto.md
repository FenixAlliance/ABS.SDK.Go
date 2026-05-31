# ItemPickListEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**ItemId** | Pointer to **NullableString** |  | [optional] 
**WarehouseId** | Pointer to **NullableString** |  | [optional] 
**ItemPickListId** | Pointer to **NullableString** |  | [optional] 
**OrderItemRecordId** | Pointer to **NullableString** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewItemPickListEntryDto

`func NewItemPickListEntryDto() *ItemPickListEntryDto`

NewItemPickListEntryDto instantiates a new ItemPickListEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPickListEntryDtoWithDefaults

`func NewItemPickListEntryDtoWithDefaults() *ItemPickListEntryDto`

NewItemPickListEntryDtoWithDefaults instantiates a new ItemPickListEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemPickListEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemPickListEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemPickListEntryDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemPickListEntryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ItemPickListEntryDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ItemPickListEntryDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTimestamp

`func (o *ItemPickListEntryDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemPickListEntryDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemPickListEntryDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemPickListEntryDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ItemPickListEntryDto) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ItemPickListEntryDto) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetQuantity

`func (o *ItemPickListEntryDto) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ItemPickListEntryDto) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ItemPickListEntryDto) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ItemPickListEntryDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetItemId

`func (o *ItemPickListEntryDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPickListEntryDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPickListEntryDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemPickListEntryDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ItemPickListEntryDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ItemPickListEntryDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetWarehouseId

`func (o *ItemPickListEntryDto) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ItemPickListEntryDto) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ItemPickListEntryDto) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *ItemPickListEntryDto) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### SetWarehouseIdNil

`func (o *ItemPickListEntryDto) SetWarehouseIdNil(b bool)`

 SetWarehouseIdNil sets the value for WarehouseId to be an explicit nil

### UnsetWarehouseId
`func (o *ItemPickListEntryDto) UnsetWarehouseId()`

UnsetWarehouseId ensures that no value is present for WarehouseId, not even an explicit nil
### GetItemPickListId

`func (o *ItemPickListEntryDto) GetItemPickListId() string`

GetItemPickListId returns the ItemPickListId field if non-nil, zero value otherwise.

### GetItemPickListIdOk

`func (o *ItemPickListEntryDto) GetItemPickListIdOk() (*string, bool)`

GetItemPickListIdOk returns a tuple with the ItemPickListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemPickListId

`func (o *ItemPickListEntryDto) SetItemPickListId(v string)`

SetItemPickListId sets ItemPickListId field to given value.

### HasItemPickListId

`func (o *ItemPickListEntryDto) HasItemPickListId() bool`

HasItemPickListId returns a boolean if a field has been set.

### SetItemPickListIdNil

`func (o *ItemPickListEntryDto) SetItemPickListIdNil(b bool)`

 SetItemPickListIdNil sets the value for ItemPickListId to be an explicit nil

### UnsetItemPickListId
`func (o *ItemPickListEntryDto) UnsetItemPickListId()`

UnsetItemPickListId ensures that no value is present for ItemPickListId, not even an explicit nil
### GetOrderItemRecordId

`func (o *ItemPickListEntryDto) GetOrderItemRecordId() string`

GetOrderItemRecordId returns the OrderItemRecordId field if non-nil, zero value otherwise.

### GetOrderItemRecordIdOk

`func (o *ItemPickListEntryDto) GetOrderItemRecordIdOk() (*string, bool)`

GetOrderItemRecordIdOk returns a tuple with the OrderItemRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemRecordId

`func (o *ItemPickListEntryDto) SetOrderItemRecordId(v string)`

SetOrderItemRecordId sets OrderItemRecordId field to given value.

### HasOrderItemRecordId

`func (o *ItemPickListEntryDto) HasOrderItemRecordId() bool`

HasOrderItemRecordId returns a boolean if a field has been set.

### SetOrderItemRecordIdNil

`func (o *ItemPickListEntryDto) SetOrderItemRecordIdNil(b bool)`

 SetOrderItemRecordIdNil sets the value for OrderItemRecordId to be an explicit nil

### UnsetOrderItemRecordId
`func (o *ItemPickListEntryDto) UnsetOrderItemRecordId()`

UnsetOrderItemRecordId ensures that no value is present for OrderItemRecordId, not even an explicit nil
### GetTenantId

`func (o *ItemPickListEntryDto) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ItemPickListEntryDto) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ItemPickListEntryDto) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ItemPickListEntryDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *ItemPickListEntryDto) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *ItemPickListEntryDto) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


