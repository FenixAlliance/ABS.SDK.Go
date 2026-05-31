# ItemRetainSampleCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**WarehouseId** | **string** |  | 
**ItemId** | **string** |  | 

## Methods

### NewItemRetainSampleCreateDto

`func NewItemRetainSampleCreateDto(warehouseId string, itemId string, ) *ItemRetainSampleCreateDto`

NewItemRetainSampleCreateDto instantiates a new ItemRetainSampleCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemRetainSampleCreateDtoWithDefaults

`func NewItemRetainSampleCreateDtoWithDefaults() *ItemRetainSampleCreateDto`

NewItemRetainSampleCreateDtoWithDefaults instantiates a new ItemRetainSampleCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ItemRetainSampleCreateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemRetainSampleCreateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemRetainSampleCreateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ItemRetainSampleCreateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ItemRetainSampleCreateDto) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ItemRetainSampleCreateDto) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ItemRetainSampleCreateDto) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ItemRetainSampleCreateDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetWarehouseId

`func (o *ItemRetainSampleCreateDto) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *ItemRetainSampleCreateDto) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *ItemRetainSampleCreateDto) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.


### GetItemId

`func (o *ItemRetainSampleCreateDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemRetainSampleCreateDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemRetainSampleCreateDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


